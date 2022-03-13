package service

import (
	"context"
	"testing"

	applicationspb "github.com/paranoiacblack/destiny2/service/protos/applications"
	userpb "github.com/paranoiacblack/destiny2/service/protos/user"
)

const key = "a243b9c909314c73ac35e641ef1059ca"

// For now, just test methods which don't require OAuth2.
func TestApp(t *testing.T) {
	ctx := NewBungieContext(context.Background(), key)
	client := applicationspb.NewAppClient(NewProxyConn())

	apps, err := client.GetBungieApplications(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	appNames := map[int32]string{
		8750:  "www.bungie.net",
		8751:  "Destiny 2 Companion (Android)",
		8752:  "Destiny 2 Companion (iOS)",
		29065: "iOS test version (debug and release)",
		29066: "iOS App Store version (distribution)",
		29067: "Android test version (debug)",
		29068: "Android Google Play (release)",
	}

	for _, app := range apps.GetApplications() {
		name, ok := appNames[app.GetApplicationId()]
		if !ok {
			t.Errorf("GetApplications(%d) has no mapped name", app.GetApplicationId())
		}

		if app.GetName() != name {
			t.Errorf("GetApplications(%d): got %q, want %q", app.GetApplicationId(), app.GetName(), name)
		}
	}
}

func TestUser(t *testing.T) {
	ctx := NewBungieContext(context.Background(), key)
	client := userpb.NewUserClient(NewProxyConn())

	themes, err := client.GetAvailableThemes(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	// There are currently 41 available themes, as of this commit.
	const totalThemes = 41
	if gotThemes := len(themes.GetThemes()); gotThemes != totalThemes {
		t.Errorf("GetAvailableThemes: got %d total themes, want %d total themes", gotThemes, totalThemes)
	}

	const (
		devMembershipID = 29013018
		wantName        = "CatDog Let's Player#6540"
	)
	membershipReq := &userpb.MembershipRequest{MembershipId: devMembershipID}
	user, err := client.GetBungieNetUserByID(ctx, membershipReq)
	if err != nil {
		t.Fatal(err)
	}

	if gotName := user.GetUniqueName(); gotName != wantName {
		t.Errorf("GetBungieNetUserByID(%d): got %s, want %s", devMembershipID, gotName, wantName)
	}

	// In this case, search for a user name prefix that appears in most names.
	// Check that hasMore is true.
	// TODO(paranoiacblack): Use a test server and table-driven test for given queries.
	searchReq := &userpb.UserSearchRequest{
		Prefix: &userpb.UserSearchPrefixRequest{DisplayNamePrefix: "c"},
		Page:   0,
	}
	searchResp, err := client.SearchByGlobalName(ctx, searchReq)
	if err != nil {
		t.Fatal(err)
	}
	if !searchResp.GetHasMore() {
		t.Errorf("SearchByGlobalName('c'): expected multiple pages, got all results on page 0")
	}
}
