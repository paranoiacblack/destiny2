package service

// ApplicationScope represents a bitmask used to determine which scopes an application has when using the Bungie API.
type ApplicationScope int64

const (
	// Read basic user profile information such as the user's handle, avatar icon, etc.
	ReadBasicUserProfile ApplicationScope = 1
	// Read Group/Clan Forums, Wall, and Members for groups and clans that the user has joined.
	ReadGroups = 2
	// Write Group/Clan Forums, Wall, and Members for groups and clans that the user has joined.
	WriteGroups = 4
	// Administer Group/Clan Forums, Wall, and Members for groups and clans that the user is a founder or an administrator.
	AdminGroups = 8
	// Create new groups, clans, and forum posts, along with other actions that are reserved for Bungie.net elevated scope: not meant to be used by third party applications.
	BnetWrite = 16
	// Move or equip Destiny items
	MoveEquipDestinyItems = 32
	// Read Destiny 1 Inventory and Vault contents. For Destiny 2, this scope is needed to read anything regarded as private. This is the only scope a Destiny 2 app needs for read operations against Destiny 2 data such as inventory, vault, currency, vendors, milestones, progression, etc.
	ReadDestinyInventoryAndVault = 64
	// Read user data such as who they are web notifications, clan/group memberships, recent activity, muted users.
	ReadUserData = 128
	// Edit user data such as preferred language, status, motto, avatar selection and theme.
	EditUserData = 256
	// Access vendor and advisor data specific to a user. OBSOLETE. This scope is only used on the Destiny 1 API.
	ReadDestinyVendorsAndAdvisors = 512
	// Read offer history and claim and apply tokens for the user.
	ReadAndApplyTokens = 1024
	// Can perform actions that will result in a prompt to the user via the Destiny app.
	AdvancedWriteActions = 2048
	// Can user the partner offer api to claim rewards defined for a partner.
	PartnerOfferGrant = 4096
	// Allows an app to query sensitive information like unlock flags and values not available through normal methods.
	DestinyUnlockValueQuery = 8192
	// Allows an app to query sensitive user PII, most notably email information.
	UserPiiRead = 16384
)
