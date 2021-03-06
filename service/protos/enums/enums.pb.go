// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/enums/enums.proto

package enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApplicationScope int32

const (
	ApplicationScope_SCOPES_NONE                       ApplicationScope = 0
	ApplicationScope_READ_BASIC_USER_PROFILE           ApplicationScope = 1
	ApplicationScope_READ_GROUPS                       ApplicationScope = 2
	ApplicationScope_WRITE_GROUPS                      ApplicationScope = 4
	ApplicationScope_ADMIN_GROUPS                      ApplicationScope = 8
	ApplicationScope_BNET_WRITE                        ApplicationScope = 16
	ApplicationScope_MOVE_EQUIP_DESTINY_ITEMS          ApplicationScope = 32
	ApplicationScope_READ_DESTINY_INVENTORY_AND_VAULT  ApplicationScope = 64
	ApplicationScope_READ_USER_DATA                    ApplicationScope = 128
	ApplicationScope_EDIT_USER_DATA                    ApplicationScope = 256
	ApplicationScope_READ_DESTINY_VENDORS_AND_ADVISORS ApplicationScope = 512
	ApplicationScope_READ_AND_APPLY_TOKENS             ApplicationScope = 1024
	ApplicationScope_ADVANCED_WRITE_ACTIONS            ApplicationScope = 2048
	ApplicationScope_PARTNER_OFFER_GRANT               ApplicationScope = 4096
	ApplicationScope_DESTINY_UNLOCK_VALUE_QUERY        ApplicationScope = 8192
	ApplicationScope_USER_PII_READ                     ApplicationScope = 16384
)

// Enum value maps for ApplicationScope.
var (
	ApplicationScope_name = map[int32]string{
		0:     "SCOPES_NONE",
		1:     "READ_BASIC_USER_PROFILE",
		2:     "READ_GROUPS",
		4:     "WRITE_GROUPS",
		8:     "ADMIN_GROUPS",
		16:    "BNET_WRITE",
		32:    "MOVE_EQUIP_DESTINY_ITEMS",
		64:    "READ_DESTINY_INVENTORY_AND_VAULT",
		128:   "READ_USER_DATA",
		256:   "EDIT_USER_DATA",
		512:   "READ_DESTINY_VENDORS_AND_ADVISORS",
		1024:  "READ_AND_APPLY_TOKENS",
		2048:  "ADVANCED_WRITE_ACTIONS",
		4096:  "PARTNER_OFFER_GRANT",
		8192:  "DESTINY_UNLOCK_VALUE_QUERY",
		16384: "USER_PII_READ",
	}
	ApplicationScope_value = map[string]int32{
		"SCOPES_NONE":                       0,
		"READ_BASIC_USER_PROFILE":           1,
		"READ_GROUPS":                       2,
		"WRITE_GROUPS":                      4,
		"ADMIN_GROUPS":                      8,
		"BNET_WRITE":                        16,
		"MOVE_EQUIP_DESTINY_ITEMS":          32,
		"READ_DESTINY_INVENTORY_AND_VAULT":  64,
		"READ_USER_DATA":                    128,
		"EDIT_USER_DATA":                    256,
		"READ_DESTINY_VENDORS_AND_ADVISORS": 512,
		"READ_AND_APPLY_TOKENS":             1024,
		"ADVANCED_WRITE_ACTIONS":            2048,
		"PARTNER_OFFER_GRANT":               4096,
		"DESTINY_UNLOCK_VALUE_QUERY":        8192,
		"USER_PII_READ":                     16384,
	}
)

func (x ApplicationScope) Enum() *ApplicationScope {
	p := new(ApplicationScope)
	*p = x
	return p
}

func (x ApplicationScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplicationScope) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_enums_enums_proto_enumTypes[0].Descriptor()
}

func (ApplicationScope) Type() protoreflect.EnumType {
	return &file_protos_enums_enums_proto_enumTypes[0]
}

func (x ApplicationScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplicationScope.Descriptor instead.
func (ApplicationScope) EnumDescriptor() ([]byte, []int) {
	return file_protos_enums_enums_proto_rawDescGZIP(), []int{0}
}

type ApplicationStatus int32

const (
	ApplicationStatus_APP_STATUS_NONE ApplicationStatus = 0
	ApplicationStatus_PRIVATE         ApplicationStatus = 1
	ApplicationStatus_PUBLIC          ApplicationStatus = 2
	ApplicationStatus_DISABLED        ApplicationStatus = 3
	ApplicationStatus_BLOCKED         ApplicationStatus = 4
)

// Enum value maps for ApplicationStatus.
var (
	ApplicationStatus_name = map[int32]string{
		0: "APP_STATUS_NONE",
		1: "PRIVATE",
		2: "PUBLIC",
		3: "DISABLED",
		4: "BLOCKED",
	}
	ApplicationStatus_value = map[string]int32{
		"APP_STATUS_NONE": 0,
		"PRIVATE":         1,
		"PUBLIC":          2,
		"DISABLED":        3,
		"BLOCKED":         4,
	}
)

func (x ApplicationStatus) Enum() *ApplicationStatus {
	p := new(ApplicationStatus)
	*p = x
	return p
}

func (x ApplicationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplicationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_enums_enums_proto_enumTypes[1].Descriptor()
}

func (ApplicationStatus) Type() protoreflect.EnumType {
	return &file_protos_enums_enums_proto_enumTypes[1]
}

func (x ApplicationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplicationStatus.Descriptor instead.
func (ApplicationStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_enums_enums_proto_rawDescGZIP(), []int{1}
}

type DeveloperRole int32

const (
	DeveloperRole_ROLE_NONE   DeveloperRole = 0
	DeveloperRole_OWNER       DeveloperRole = 1
	DeveloperRole_TEAM_MEMBER DeveloperRole = 2
)

// Enum value maps for DeveloperRole.
var (
	DeveloperRole_name = map[int32]string{
		0: "ROLE_NONE",
		1: "OWNER",
		2: "TEAM_MEMBER",
	}
	DeveloperRole_value = map[string]int32{
		"ROLE_NONE":   0,
		"OWNER":       1,
		"TEAM_MEMBER": 2,
	}
)

func (x DeveloperRole) Enum() *DeveloperRole {
	p := new(DeveloperRole)
	*p = x
	return p
}

func (x DeveloperRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeveloperRole) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_enums_enums_proto_enumTypes[2].Descriptor()
}

func (DeveloperRole) Type() protoreflect.EnumType {
	return &file_protos_enums_enums_proto_enumTypes[2]
}

func (x DeveloperRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeveloperRole.Descriptor instead.
func (DeveloperRole) EnumDescriptor() ([]byte, []int) {
	return file_protos_enums_enums_proto_rawDescGZIP(), []int{2}
}

type BungieMembershipType int32

const (
	BungieMembershipType_MEMBERSHIP_NONE BungieMembershipType = 0
	BungieMembershipType_TIGER_XBOX      BungieMembershipType = 1
	BungieMembershipType_TIGER_PSN       BungieMembershipType = 2
	BungieMembershipType_TIGER_STEAM     BungieMembershipType = 3
	BungieMembershipType_TIGER_BLIZZARD  BungieMembershipType = 4
	BungieMembershipType_TIGER_STADIA    BungieMembershipType = 5
	BungieMembershipType_TIGER_DEMON     BungieMembershipType = 10
	BungieMembershipType_BUNGIE_NEXT     BungieMembershipType = 254
	BungieMembershipType_ALL             BungieMembershipType = -1
)

// Enum value maps for BungieMembershipType.
var (
	BungieMembershipType_name = map[int32]string{
		0:   "MEMBERSHIP_NONE",
		1:   "TIGER_XBOX",
		2:   "TIGER_PSN",
		3:   "TIGER_STEAM",
		4:   "TIGER_BLIZZARD",
		5:   "TIGER_STADIA",
		10:  "TIGER_DEMON",
		254: "BUNGIE_NEXT",
		-1:  "ALL",
	}
	BungieMembershipType_value = map[string]int32{
		"MEMBERSHIP_NONE": 0,
		"TIGER_XBOX":      1,
		"TIGER_PSN":       2,
		"TIGER_STEAM":     3,
		"TIGER_BLIZZARD":  4,
		"TIGER_STADIA":    5,
		"TIGER_DEMON":     10,
		"BUNGIE_NEXT":     254,
		"ALL":             -1,
	}
)

func (x BungieMembershipType) Enum() *BungieMembershipType {
	p := new(BungieMembershipType)
	*p = x
	return p
}

func (x BungieMembershipType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BungieMembershipType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_enums_enums_proto_enumTypes[3].Descriptor()
}

func (BungieMembershipType) Type() protoreflect.EnumType {
	return &file_protos_enums_enums_proto_enumTypes[3]
}

func (x BungieMembershipType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BungieMembershipType.Descriptor instead.
func (BungieMembershipType) EnumDescriptor() ([]byte, []int) {
	return file_protos_enums_enums_proto_rawDescGZIP(), []int{3}
}

type BungieCredentialType int32

const (
	BungieCredentialType_CREDENTIALS_NONE BungieCredentialType = 0
	BungieCredentialType_XU_ID            BungieCredentialType = 1
	BungieCredentialType_PSN_ID           BungieCredentialType = 2
	BungieCredentialType_WILD             BungieCredentialType = 3
	BungieCredentialType_FAKE             BungieCredentialType = 4
	BungieCredentialType_FACEBOOK         BungieCredentialType = 5
	BungieCredentialType_GOOGLE           BungieCredentialType = 8
	BungieCredentialType_WINDOWS          BungieCredentialType = 9
	BungieCredentialType_DEMON_ID         BungieCredentialType = 10
	BungieCredentialType_STEAM_ID         BungieCredentialType = 12
	BungieCredentialType_BATTLENET_ID     BungieCredentialType = 14
	BungieCredentialType_STADIA_ID        BungieCredentialType = 16
	BungieCredentialType_TWITCH_ID        BungieCredentialType = 18
)

// Enum value maps for BungieCredentialType.
var (
	BungieCredentialType_name = map[int32]string{
		0:  "CREDENTIALS_NONE",
		1:  "XU_ID",
		2:  "PSN_ID",
		3:  "WILD",
		4:  "FAKE",
		5:  "FACEBOOK",
		8:  "GOOGLE",
		9:  "WINDOWS",
		10: "DEMON_ID",
		12: "STEAM_ID",
		14: "BATTLENET_ID",
		16: "STADIA_ID",
		18: "TWITCH_ID",
	}
	BungieCredentialType_value = map[string]int32{
		"CREDENTIALS_NONE": 0,
		"XU_ID":            1,
		"PSN_ID":           2,
		"WILD":             3,
		"FAKE":             4,
		"FACEBOOK":         5,
		"GOOGLE":           8,
		"WINDOWS":          9,
		"DEMON_ID":         10,
		"STEAM_ID":         12,
		"BATTLENET_ID":     14,
		"STADIA_ID":        16,
		"TWITCH_ID":        18,
	}
)

func (x BungieCredentialType) Enum() *BungieCredentialType {
	p := new(BungieCredentialType)
	*p = x
	return p
}

func (x BungieCredentialType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BungieCredentialType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_enums_enums_proto_enumTypes[4].Descriptor()
}

func (BungieCredentialType) Type() protoreflect.EnumType {
	return &file_protos_enums_enums_proto_enumTypes[4]
}

func (x BungieCredentialType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BungieCredentialType.Descriptor instead.
func (BungieCredentialType) EnumDescriptor() ([]byte, []int) {
	return file_protos_enums_enums_proto_rawDescGZIP(), []int{4}
}

var File_protos_enums_enums_proto protoreflect.FileDescriptor

var file_protos_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2a, 0xa4, 0x03, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x53,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x41, 0x44, 0x5f,
	0x42, 0x41, 0x53, 0x49, 0x43, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x53, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4e, 0x45,
	0x54, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x10, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x4f, 0x56,
	0x45, 0x5f, 0x45, 0x51, 0x55, 0x49, 0x50, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x59, 0x5f,
	0x49, 0x54, 0x45, 0x4d, 0x53, 0x10, 0x20, 0x12, 0x24, 0x0a, 0x20, 0x52, 0x45, 0x41, 0x44, 0x5f,
	0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x45, 0x4e, 0x54, 0x4f, 0x52,
	0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x56, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x40, 0x12, 0x13, 0x0a,
	0x0e, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10,
	0x80, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x10, 0x80, 0x02, 0x12, 0x26, 0x0a, 0x21, 0x52, 0x45, 0x41, 0x44, 0x5f,
	0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x59, 0x5f, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x53, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x41, 0x44, 0x56, 0x49, 0x53, 0x4f, 0x52, 0x53, 0x10, 0x80, 0x04, 0x12,
	0x1a, 0x0a, 0x15, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x4c,
	0x59, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x53, 0x10, 0x80, 0x08, 0x12, 0x1b, 0x0a, 0x16, 0x41,
	0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x80, 0x10, 0x12, 0x18, 0x0a, 0x13, 0x50, 0x41, 0x52, 0x54,
	0x4e, 0x45, 0x52, 0x5f, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x5f, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x10,
	0x80, 0x20, 0x12, 0x1f, 0x0a, 0x1a, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x59, 0x5f, 0x55, 0x4e,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59,
	0x10, 0x80, 0x40, 0x12, 0x13, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x49, 0x49, 0x5f,
	0x52, 0x45, 0x41, 0x44, 0x10, 0x80, 0x80, 0x01, 0x2a, 0x5c, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x0a,
	0x0f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x3a, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x02, 0x2a, 0xb6, 0x01, 0x0a, 0x14, 0x42, 0x75, 0x6e, 0x67, 0x69, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x49, 0x47, 0x45, 0x52, 0x5f, 0x58, 0x42, 0x4f, 0x58, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x47, 0x45, 0x52, 0x5f, 0x50, 0x53, 0x4e, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x49, 0x47, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x54, 0x49, 0x47, 0x45, 0x52, 0x5f, 0x42, 0x4c, 0x49, 0x5a, 0x5a, 0x41,
	0x52, 0x44, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x47, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x44, 0x49, 0x41, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x49, 0x47, 0x45, 0x52, 0x5f,
	0x44, 0x45, 0x4d, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0b, 0x42, 0x55, 0x4e, 0x47, 0x49,
	0x45, 0x5f, 0x4e, 0x45, 0x58, 0x54, 0x10, 0xfe, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x4c, 0x4c,
	0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x2a, 0xca, 0x01, 0x0a, 0x14,
	0x42, 0x75, 0x6e, 0x67, 0x69, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49,
	0x41, 0x4c, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x58, 0x55,
	0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x53, 0x4e, 0x5f, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x49, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x41, 0x4b, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41, 0x43, 0x45, 0x42, 0x4f, 0x4f,
	0x4b, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x08, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x4d, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54,
	0x45, 0x41, 0x4d, 0x5f, 0x49, 0x44, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x41, 0x54, 0x54,
	0x4c, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54,
	0x41, 0x44, 0x49, 0x41, 0x5f, 0x49, 0x44, 0x10, 0x10, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x57, 0x49,
	0x54, 0x43, 0x48, 0x5f, 0x49, 0x44, 0x10, 0x12, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6e, 0x6f, 0x69, 0x61, 0x63,
	0x62, 0x6c, 0x61, 0x63, 0x6b, 0x2f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x79, 0x32, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_enums_enums_proto_rawDescOnce sync.Once
	file_protos_enums_enums_proto_rawDescData = file_protos_enums_enums_proto_rawDesc
)

func file_protos_enums_enums_proto_rawDescGZIP() []byte {
	file_protos_enums_enums_proto_rawDescOnce.Do(func() {
		file_protos_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_enums_enums_proto_rawDescData)
	})
	return file_protos_enums_enums_proto_rawDescData
}

var file_protos_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_protos_enums_enums_proto_goTypes = []interface{}{
	(ApplicationScope)(0),     // 0: enums.ApplicationScope
	(ApplicationStatus)(0),    // 1: enums.ApplicationStatus
	(DeveloperRole)(0),        // 2: enums.DeveloperRole
	(BungieMembershipType)(0), // 3: enums.BungieMembershipType
	(BungieCredentialType)(0), // 4: enums.BungieCredentialType
}
var file_protos_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_enums_enums_proto_init() }
func file_protos_enums_enums_proto_init() {
	if File_protos_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_enums_enums_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_enums_enums_proto_goTypes,
		DependencyIndexes: file_protos_enums_enums_proto_depIdxs,
		EnumInfos:         file_protos_enums_enums_proto_enumTypes,
	}.Build()
	File_protos_enums_enums_proto = out.File
	file_protos_enums_enums_proto_rawDesc = nil
	file_protos_enums_enums_proto_goTypes = nil
	file_protos_enums_enums_proto_depIdxs = nil
}
