// Package service implements the Bungie.net Destiny2 API (https://github.com/Bungie-net/api).
package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	applicationspb "github.com/paranoiacblack/destiny2/service/protos/applications"
	userpb "github.com/paranoiacblack/destiny2/service/protos/user"
)

// APIError is the error returned by a client when an API call fails.
type APIError struct {
	ErrorCode          PlatformErrorCode
	ThrottleSeconds    time.Duration
	ErrorStatus        string
	Message            string
	MessageData        map[string]string
	DetailedErrorTrace string
}

func (err APIError) Error() string {
	return fmt.Sprintf(`Bungie Error(%d): "%s: %s"`, err.ErrorCode, err.ErrorStatus, err.Message)
}

// NewBungieContext returns a context which associates an X-API-Key with an API Call.
// Any calls to the Bungie.net API require X-API-Key in the header.
func NewBungieContext(ctx context.Context, APIKey string) context.Context {
	return &bungieContext{ctx, APIKey}
}

type bungieContext struct {
	context.Context
	APIKey string
}

const apiRootPath = "https://www.bungie.net/Platform"

func bungieGet(ctx context.Context, path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", apiRootPath, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return bungieDo(ctx, req)
}

func bungiePost(ctx context.Context, path string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s%s", apiRootPath, path)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	return bungieDo(ctx, req)
}

func bungieDo(ctx context.Context, req *http.Request) ([]byte, error) {
	if bungieContext, ok := ctx.(*bungieContext); ok {
		req.Header.Add("X-API-Key", bungieContext.APIKey)
	}

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var bungieResp struct {
		Response interface{}
		APIError
	}
	if err := json.Unmarshal(data, &bungieResp); err != nil {
		return nil, err
	}

	if bungieResp.ErrorCode != Success {
		return nil, bungieResp.APIError
	}
	return json.Marshal(bungieResp.Response)
}

type proxyFunc func(context.Context, *proxyServer, interface{}) (proto.Message, error)

// proxyFuncs maps each gRPC method name to its implementation via calling the Bungie.net API.
var proxyFuncs = map[string]proxyFunc{
	"/applications.App/GetApplicationAPIUsage": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetApplicationAPIUsage(ctx, in.(*applicationspb.UsageRequest))
	},
	"/applications.App/GetBungieApplications": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetBungieApplications(ctx, in.(*emptypb.Empty))
	},
	"/user.User/GetBungieNetUserByID": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetBungieNetUserByID(ctx, in.(*userpb.MembershipRequest))
	},
	"/user.User/GetCredentialTypesForTargetAccount": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetCredentialTypesForTargetAccount(ctx, in.(*userpb.MembershipRequest))
	},
	"/user.User/GetAvailableThemes": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetAvailableThemes(ctx, in.(*emptypb.Empty))
	},
	"/user.User/GetMembershipDataByID": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetMembershipDataByID(ctx, in.(*userpb.MembershipDataRequest))
	},
	"/user.User/GetMembershipDataForCurrentUser": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetMembershipDataForCurrentUser(ctx, in.(*emptypb.Empty))
	},
	"/user.User/GetMembershipFromHardLinkedCredential": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.GetMembershipFromHardLinkedCredential(ctx, in.(*userpb.HardLinkedCredentialRequest))
	},
	"/user.User/SearchByGlobalName": func(ctx context.Context, proxy *proxyServer, in interface{}) (proto.Message, error) {
		return proxy.SearchByGlobalName(ctx, in.(*userpb.UserSearchRequest))
	},
}

type proxyConn struct {
	srv *proxyServer
}

// NewProxyConn returns a connection to the proxy server.
// The proxy server calls into the Bungie API directly, but accepts
// requests and returns responses as defined protos in destiny2/service/protos.
// Most users of this package will want to use this to connect to the Bungie.net API.
func NewProxyConn() grpc.ClientConnInterface {
	return &proxyConn{srv: new(proxyServer)}
}

func (cc *proxyConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	fn, ok := proxyFuncs[method]
	if !ok {
		return status.Errorf(codes.Unimplemented, "Method %s is not available in the bungie proxy service", method)
	}

	result, err := fn(ctx, cc.srv, args)
	if err != nil {
		return err
	}

	proto.Merge(reply.(proto.Message), result)
	return nil
}

func (cc *proxyConn) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, status.Error(codes.Unimplemented, "Streaming is not supported by the Bungie API proxy")
}

type proxyServer struct {
	userpb.UnimplementedUserServer
}

// wrapJSONArray is meant to wrap a JSON array value from the Bungie.net API
// into something that can be marshalled into a protobuf message.
// The issue is that there is no way to write a gRPC service which returns
// a repeated value.
// The standard way to do this is to create a top-level message which contains a single field of repeated values.
// The wrapperName provided is the name of that single field within a top-level message.
// wrapJSONArray adds a container to the JSONArray in data, by wrapping it {'wrapperName': []}.
func wrapJSONArray(data []byte, wrapperName string) []byte {
	wrapper := fmt.Sprintf(`{%q:`, wrapperName)
	wrapped := strings.Join([]string{wrapper, string(data), "}"}, "")
	return []byte(wrapped)
}

func (srv *proxyServer) GetApplicationAPIUsage(ctx context.Context, in *applicationspb.UsageRequest) (*applicationspb.APIUsage, error) {
	v := url.Values{}
	start := in.Start.AsTime()
	end := in.End.AsTime()
	if !start.IsZero() {
		v.Set("start", fmt.Sprintf("%s", in.Start))
	}
	if !end.IsZero() {
		v.Set("end", fmt.Sprintf("%s", in.End))
	}

	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	data, err := bungieGet(ctx, fmt.Sprintf("/App/ApiUsage/%d/"+encoded, in.ApplicationId))
	if err != nil {
		return nil, err
	}

	out := new(applicationspb.APIUsage)
	if err := protojson.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (srv *proxyServer) GetBungieApplications(ctx context.Context, in *emptypb.Empty) (*applicationspb.ApplicationsResponse, error) {
	data, err := bungieGet(ctx, "/App/FirstParty/")
	if err != nil {
		return nil, err
	}
	data = wrapJSONArray(data, "applications")

	out := new(applicationspb.ApplicationsResponse)
	if err := protojson.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (srv *proxyServer) GetBungieNetUserByID(ctx context.Context, in *userpb.MembershipRequest) (*userpb.GeneralUser, error) {
	data, err := bungieGet(ctx, fmt.Sprintf("/User/GetBungieNetUserById/%d/", in.MembershipId))
	if err != nil {
		return nil, err
	}

	out := new(userpb.GeneralUser)
	if err := protojson.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (srv *proxyServer) GetCredentialTypesForTargetAccount(ctx context.Context, in *userpb.MembershipRequest) (*userpb.CredentialTypesForAccountResponse, error) {
	data, err := bungieGet(ctx, fmt.Sprintf("/User/GetCredentialTypesForTargetAccount/%d/", in.MembershipId))
	if err != nil {
		return nil, err
	}
	data = wrapJSONArray(data, "credentials")

	out := new(userpb.CredentialTypesForAccountResponse)
	if err := protojson.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (srv *proxyServer) GetAvailableThemes(ctx context.Context, empty *emptypb.Empty) (*userpb.AvailableThemesResponse, error) {
	data, err := bungieGet(ctx, "/User/GetAvailableThemes/")
	if err != nil {
		return nil, err
	}
	data = wrapJSONArray(data, "themes")

	resp := new(userpb.AvailableThemesResponse)
	if err := protojson.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (srv *proxyServer) GetMembershipsByIxD(ctx context.Context, req *userpb.MembershipDataRequest) (*userpb.UserMembershipData, error) {
	data, err := bungieGet(ctx, fmt.Sprintf("/User/GetMembershipsById/%d/%d/", req.MembershipId, req.MembershipType))
	if err != nil {
		return nil, err
	}

	resp := new(userpb.UserMembershipData)
	if err := protojson.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (srv *proxyServer) GetMembershipDataForCurrentUser(ctx context.Context, empty *emptypb.Empty) (*userpb.UserMembershipData, error) {
	data, err := bungieGet(ctx, "/User/GetMembershipsForCurrentUser/")
	if err != nil {
		return nil, err
	}

	resp := new(userpb.UserMembershipData)
	if err := protojson.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (srv *proxyServer) GetMembershipFromHardLinkedCredential(ctx context.Context, req *userpb.HardLinkedCredentialRequest) (*userpb.HardLinkedUserMembership, error) {
	data, err := bungieGet(ctx, fmt.Sprintf("/User/GetMembershipFromHardLinkedCredential/%s/%d/", req.Credential, req.CrType))
	if err != nil {
		return nil, err
	}

	resp := new(userpb.HardLinkedUserMembership)
	if err := protojson.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (srv *proxyServer) SearchByGlobalName(ctx context.Context, req *userpb.UserSearchRequest) (*userpb.UserSearchResponse, error) {
	body, err := protojson.Marshal(req.Prefix)
	if err != nil {
		return nil, err
	}

	data, err := bungiePost(ctx, fmt.Sprintf("/User/Search/GlobalName/%d/", req.Page), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp := new(userpb.UserSearchResponse)
	if err := protojson.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
