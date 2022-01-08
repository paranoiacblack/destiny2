package destiny2

// ProgressionStep defines a single step in a progression.
type ProgressionStep struct {
	// StepName is the localized text describing this step of the progression.
	StepName string
	// DisplayEffectType is a classification of effects to show when this progression goes to the next step.
	DisplayEffectType ProgressionStepDisplayEffect
	// ProgressTotal is the total amount of progression needed to reach this step.
	ProgressTotal int32
	// RewardItems is a list of items rewarded for reaching this step.
	RewardItems []ItemQuantity
	// Icon is a specific icon for this step.
	Icon string
}

// ItemQuantity describes an item stack and its quantity.
type ItemQuantity struct {
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// ItemInstanceID refers to a specific instance of an item.
	ItemInstanceId int64
	// Quantity is the amount of the item.
	Quantity int32
	// HasCondtionalVisibility determines if the item quantity may be condtionally shown or hideen.
	HasConditionalVisibility bool
}

// ItemTooltipNotification is a tooltip that comes up conditionally for an item.
type ItemTooltipNotification struct {
	// DisplayString is a localized tooltip to display.
	DisplayString string
	// DisplayStyle identifies the style used to show this display in a given UI.
	DisplayStyle string
}

// ItemActionBlock defines an action that can be performed on an item.
type ItemActionBlock struct {
	// VerbName is localized text for the verb of this action.
	VerbName string
	// VerbDescription is a localized text describing this action.
	VerbDescription string
	// IsPositive determines if the given item has this action.
	IsPositive bool
	// OverlayScreenName is the name of the overlay screen associated with this action.
	OverlayScreenName string
	// OverlayIcon is the icon of the overlay screen associated with this action.
	OverlayIcon string
	// RequiredCooldownSeconds is the number of seconds to delay before allowing this action to be performed again.
	RequiredCooldownSeconds int32
	// RequiredItems are other items that must exist or be destroyed to perform this action.
	RequiredItems []ItemActionRequiredItem
	// ProgressionRewards are the progressions and rewards earned by performing this action.
	ProgressionRewards []ProgressionReward
	// ActionTypeLabel is the internal identifier for this action.
	ActionTypeLabel string
	// RequiredLocation is a localized string for a hint about the location where this action can be performed.
	RequiredLocation string
	// RequiredCooldownHash is the hash of related Cooldown associated with this action.
	// Cooldown data is not currently available in the Bungie.net API.
	RequiredCooldownHash uint32
	// DeleteOnAction determines the item is deleted when this action is performed.
	DeleteOnAction bool
	// ConsumeEntireStack determines if the entire stack of items is deleted when this action is performed.
	ConsumeEntireStack bool
	// UseOnAcquire determines if this action is performed when the associated item if acquired.
	UseOnAcquire bool
}

// ItemActionRequiredItem defines an item and quantity required in an inventory to perform an action.
type ItemActionRequiredItem struct {
	// Count is the minimum quantity of the item required.
	Count int32
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// DeleteOnAction determines if the quantity of items will be deleted when the action is performed.
	DeleteOnAction bool
}

// ProgressionReward is progression rewarded when an action is performed on an inventory item.
type ProgressionReward struct {
	// ProgressionMappingHash is the hash of a related ProgressionMapping.
	// TODO(paranoiacblack): It is unclear how to obtain a ProgressionMapping through the Bungie.net API.
	ProgressionMappingHash uint32
	// Amount is the amount of experience to give to each of the mapped progressions.
	Amount int32
	// ApplyThrottles determines if the internal mechanisms to throttle progression should be applied.
	ApplyThrottles bool
}

// ItemInventoryBlock defines the basic properties regarding an item's relationship with its inventory.
type ItemInventoryBlock struct {
	// StackUniqueLabel is a unique label for a single stack in a given inventory.
	StackUniqueLabel string
	// MaxStackSize is the maximum quantity of this item that can exist in a stack.
	MaxStackSize int32
	// BucketTypeHash is the hash of a related InventoryBucketEntity.
	BucketTypeHash uint32
	// RecoveryBucketTypeHash is the hash of a related InventoryBucketEntity used for items in the lost loot queue.
	RecoveryBucketTypeHash uint32
	// TierTypeHash is the hash of a related ItemTierTypeEntity.
	TierTypeHash uint32
	// IsInstanceItem determines if this item is instanced.
	IsInstanceItem bool
	// TierTypeName is the localized name of this item's tier type.
	TierTypeName string
	// TierType is the classification of this item's tier type.
	TierType ItemTier
	// ExpirationTooltip is the tooltip to show when this item expires.
	ExpirationTooltip string
	// ExpiredInActivityMessage is the message shown if this item expires while playing in an activity.
	ExpiredInActivityMessage string
	// ExpiredInOrbitMessage is the message shown if this item expires while in orbit.
	ExpiredInOrbitMessage string
	// SuppressExpirationWhenObjectivesComplete determines if expiration messages should be suppressed.
	SuppressExpirationWhenObjectivesComplete bool
}

type ItemTierTypeInfusionBlock struct {
	// BaseQualityTransferRatio is the default portion of quality that will transfer from the infuser to the infusee.
	// (InfuserQuality - InfuseeQuality) * BaseQualityTransferRatio = base quality transferred.
	BaseQualityTransferRatio float32
	// MinimumQualityIncrement is the minimum amount of quality given if InfuserQuality > InfuseeQuality.
	MinimumQualityIncrement int32
}

// ItemSetBlock defines properties related to a quest-item and its quest steps.
type ItemSetBlock struct {
	// ItemList is a collection of hashes of items in this set.
	ItemList []ItemSetBlockEntry
	// RequiredOrderedSetItemAdd determines if items in this set can only be added in increasing order and if adding an item will remove previous items.
	RequireOrderedSetItemAdd bool
	// SetIsFeatured determines if this quest shows up as featured in a given UI.
	SetIsFeatured bool
	// SetType is an identifier of the category of this quest.
	SetType string
	// QuestLineName is the name of the quest line of this quest.
	QuestLineName string
	// QuestLineDescription is the description of the quest line of this quest.
	QuestLineDescription string
	// QuestStepSummary is an additional summary of this step in the quest line.
	QuestStepSummary string
}

// ItemSetBlockEntry is a particular entry in an ItemSet.
type ItemSetBlockEntry struct {
	// TrackingValue tracks which step has been reached in this quest.
	TrackingValue int32
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
}

// ItemStatBlock contains information about this item's calculated stats.
type ItemStatBlock struct {
	// DisablePrimaryStatDisplay determines if the primary stat on this item isn't shown on inspection.
	DisablePrimaryStatDisplay bool
	// StatGroupHash is the hash of a related StatGroupEntity.
	StatGroupHash uint32
	// Stats are precomputed stats on an item by the hash of a the related InventoryItemEntity.
	Stats map[uint32]InventoryItemStat
	// HasDisplayableStats determines if any stat on the item is visible.
	HasDisplayableStats bool
	// PrimaryBaseStatHash is the hash of a related StatEntity.
	PrimaryBaseStatHash uint32
}

// InventoryItemStat defines a specific stat value on an item and the range.
type InventoryItemStat struct {
	// StatHash is the hash of a related StatEntity.
	StatHash uint32
	// Value represents that stat value assuming the minimum roll and factoring in mandatory bonuses.
	Value int32
	// Minimum is the minimum possible value for this stat an item can roll.
	Minimum int32
	// Maximum is the maximum possible value for this stat an item can roll.
	Maximum int32
	// DisplayMaximum is the maximum possible value for this stat that is shown in a given UI.
	DisplayMaximum int32
}

// StatDisplay describes the way that an StatEntity is transformed using a related StatGroupEntity.
// This represents the transformation of a stat into a stat displayed in a given UI.
type StatDisplay struct {
	// StatHash is the hash of a related StatEntity.
	StatHash uint32
	// MaximumValue is the upper bound value for displaying a stat.
	MaximumValue int32
	// DisplayAsNumeric determines if this stat should be displayed as a number.
	DisplayAsNumeric bool
	// DisplayInterpolation is an interpolation table represents how the investment stat is transformed into a display stat.
	DisplayInterpolation []InterpolationPoint
}

// StatOverride defines a specific overridden stat.
type StatOverride struct {
	// StatHash is the hash of a related StatEntity.
	StatHash          uint32
	DisplayProperties DisplayProperties
}

// EquippingBlock contains information about how and when an item can be equipped.
type EquippingBlock struct {
	// GearsetItemHash is the hash of a InventoryItemEntity.
	GearsetItemHash uint32
	// UniqueLabel is the label used to check if the item has other items of matching types already equipped.
	UniqueLabel string
	// UniqueLabelHash is the hash of the unique label. This does not point to specific entity.
	UniqueLabelHash uint32
	// EquipmentSlotTypeHash is the hash of a related EquipmentSlotEntity.
	EquipmentSlotTypeHash uint32
	// Attributes are the custom attributes on the equippability of the item.
	Attributes EquippingItemBlockAttributes
	// AmmoType is the classification of ammo a weapon uses.
	AmmoType AmmunitionType
	// DisplayStrings represent the state failure conditions that can occur when trying to equip the item.
	// These match up 1:1 with RequiredUnlockExpressions.
	DisplayStrings []string
}

// ItemTranslationBlock defines the rendering data associated with the item.
type ItemTranslationBlock struct {
	WeaponPatternIdentifier string
	WeaponPatternHash       uint32
	DefaultDyes             []DyeReference
	LockedDyes              []DyeReference
	CustomDyes              []DyeReference
	Arrangements            []GearArtArrangementReference
	HasGeometry             bool
}

// ItemPreviewBlock represents the items that can be obtained when viewing the details of an item sack or box.
type ItemPreviewBlock struct {
	// ScreenStyle is a hint for which detail screen to show for this item preview in a given UI.
	ScreenStyle string
	// PreviewVendorHash is the hash related to a VendorEntity.
	PreviewVendorHash uint32
	// ArtifactHash is the hash related to an ArtifactEntity.
	ArtifactHash uint32
	// PreviewActionString is the localized string for the action associated with this preview.
	PreviewActionString string
	// DerivedItemCategories is the list of items being previewed, categorized in the same way as they presented in the preview UI.
	DerivedItemCategories []DerivedItemCategory
}

// DerivedItemCategory describes the category of items found in a "Preview Vendor".
type DerivedItemCategory struct {
	// CategoryDescription is a localized string for the category title.
	CategoryDescription string
	// Items is the list of all items of this category.
	Items []DerivedItem
}

// DerivedItem is a summary of a specific item that can be obtained by using or acquiring another item.
type DerivedItem struct {
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// ItemName is the name of this derived item.
	ItemName string
	// ItemDetail is an additional detail about this derived item to be added to the description.
	ItemDetail string
	// ItemDescription is a brief description of this item.
	ItemDescription string
	// IconPath is an icon for this item.
	IconPath string
	// VendorItemIndex is the index of this item in the "Preview Vendor" UI.
	VendorItemIndex int32
}

// VendorAction is an action performed by a vendor.
type VendorAction struct {
	Description       string
	ExecuteSeconds    int32
	Icon              string
	Name              string
	Verb              string
	IsPositive        bool
	ActionId          string
	ActionHash        uint32
	AutoPerformAction bool
}

// VendorCategoryEntry defines a single vendor category into which sale items are grouped.
type VendorCategoryEntry struct {
	// CategoryIndex is the index of this category in the original category definitions for the vendor.
	CategoryIndex int32
	// SortValue is used to sort items in this vendor category.
	// Recommended to use VendorCategoryComponent.ItemIndexes instead.
	SortValue int32
	// CategoryHash is the hash for this vendor category.
	CategoryHash uint32
	// QuantityAvailable is the amount of items available when this category is shown.
	QuantityAvailable int32
	// ShowUnavailableItems determines if items not for sale in this category should be shown.
	ShowUnavailableItems bool
	// HideIfNoCurrency determines if items in this category should be shown if there is no currency to buy them.
	HideIfNoCurrency bool
	// HideFromRegularPurchase determines if this category doesn't allow purchases.
	HideFromRegularPurchase bool
	// BuyStringOverride is the localized string for making purchases from this category.
	BuyStringOverride string
	// DisabledDescription is the localized description to show if this category is disabled.
	DisabledDescription string
	// DisplayTitle is the localized title of this category.
	DisplayTitle string
	// Overlay contains details of this category's overlay prompt.
	Overlay VendorCategoryOverlay
	// VendorItemIndexes are the indexes of items sold under this category.
	VendorItemIndexes []int32
	// IsPreview determines if this category is used to preview items rather than sell them.
	IsPreview bool
	// IsDisplayOnly determines if this category only displays items.
	IsDisplayOnly                bool
	ResetIntervalMinutesOverride int32
	ResetOffsetMinutesOverride   int32
}

// VendorCategoryOverlay describes and overlay prompt to show in a given UI.
type VendorCategoryOverlay struct {
	ChoiceDescription string
	Description       string
	Icon              string
	Title             string
	// CurrencyItemHash is the hash for a related InventoryItemEntity.
	CurrencyItemHash uint32
}

// DisplayCategory is a visual grouping and display of categories in a vendor's UI.
type DisplayCategory struct {
	// Index is the position of this display category in the vendor.
	Index int32
	// Identifier is the identifier for this display category.
	Identifier string
	// DisplayCategoryHash is the hash of this display category; meant for grouping.
	DisplayCategoryHash uint32
	DisplayProperties   DisplayProperties
	// DisplayInBanner determines if this display category should be displayed in the banner section of the vendor's UI.
	DisplayInBanner bool
	// ProgressionHash is the hash of a related ProgressionEntity.
	ProgressionHash uint32
	// SortOrder is the sort order of items within this display category.
	SortOrder int32
	// DisplayStyleHash indicates how this display category will be displayed in the UI.
	DisplayStyleHash uint32
	// DisplayStyleIdentifier is the identifier for this display category.
	DisplayStyleIdentifier string
}

// VendorInteraction is a dialog shown by the vendor other than sale items or transfer screens.
type VendorInteraction struct {
	// InteractionIndex is the position of this interaction in its parent array.
	InteractionIndex int32
	// Replies are potential replies that the player can make in this interaction.
	Replies []VendorInteractionReply
	// VendorCategoryIndex is the category of sale items show along with this vendor interaction.
	VendorCategoryIndex int32
	// QuestlineItemHash is the hash of a related InventoryItemEntity.
	QuestlineItemHash uint32
	// SackInteractionList is the list sacks meant to be shown in this vendor interaction.
	SackInteractionList []VendorInteractionSackEntry
	// UIInteractionType is a hint for the behavior of the interaction screen in a given UI.
	UiInteractionType uint32
	// InteractionType is the classification of UI hints for this vendor interaction.
	InteractionType VendorInteractionType
	// RewardBlockLabel is the localized text header to for displaying rewards in this interaction.
	RewardBlockLabel string
	// RewardVendorCategoryIndex is the category index to show if this interaction displays reward.
	RewardVendorCategoryIndex int32
	// FlavorLineOne is the first part of flavor text in this vendor interaction.
	FlavorLineOne string
	// FlavorLineTwo is the second part of flavor text in this vendor interaction.
	FlavorLineTwo           string
	HeaderDisplayProperties DisplayProperties
	// Instructions is a localized text telling the player what to do when viewing this interaction.
	Instructions string
}

// VendorInteractionReply is a selectable reply in a vendor interaction that obtains a reward.
type VendorInteractionReply struct {
	// ItemRewardsSelection is a classification of the rewards granted when responding to the vendor.
	ItemRewardsSelection VendorInteractionRewardSelection
	// Reply is the localized text for the vendor's reply.
	Reply string
	// ReplyType is a classification of the type of reply being made.
	ReplyType VendorReplyType
}

// VendorInventoryFlyout defines a UI screen that shows a part of a hidden vendor inventory.
type VendorInventoryFlyout struct {
	// LockedDescription is the localized description explaining why this flyout is locked.
	LockedDescription string
	DisplayProperties DisplayProperties
	// Buckets is a list of inventory buckets and other metadata to show on the screen.
	Buckets []VendorInventoryFlyoutBucket
	// FlyoutID is the identifier for the flyout.
	FlyoutId uint32
	// SuppressNewness determines if new item UI elements should be suppressed.
	SuppressNewness bool
	// EquipmentSlotHash is the hash of a related EquipmentSlotEntity.
	EquipmentSlotHash uint32
}

// VendorInventoryFlyoutBucket is information about a single inventory bucket in a vendor flyout UI.
type VendorInventoryFlyoutBucket struct {
	// Collapsible determines if this bucket can be collapsed visually.
	Collapsible bool
	// InventoryBucketHash is the hash of a related InventoryBucketEntity.
	InventoryBucketHash uint32
	// SortItemsBy is how to sort items from the flyout.
	SortItemsBy int32
}

// VendorItem represents an item being sold by the vendor.
type VendorItem struct {
	// VendorItemIndex is the index of this item in VendorEntity.SaleList.
	VendorItemIndex int32
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// Quantity is the amount of the item received when making a purchase.
	Quantity int32
	// FailureIndexes is a list of indexes in VendorEntity.FailureString which indicates possible failure strings for this item.
	FailureIndexes []int32
	// Currencies is a list of different currencies and the quantity needed to purshase this item.
	Currencies []VendorItemQuantity
	// RefundPolicy is a classification of how this item can be refunded and in what time period.
	RefundPolicy VendorItemRefundPolicy
	// RefundTimeLimit is the amount of time before refundability of this item will expire.
	RefundTimeLimit int32
	// CreationLevels are the default levels when this item will spawn in a vendor.
	CreationLevels []ItemCreationEntryLevel
	// DisplayCategoryIndex is an index into the display category for this vendor.
	DisplayCategoryIndex int32
	// CategoryIndex is an index into the category for this vendor.
	CategoryIndex int32
	// OriginalCategory is the original index into the category for this vendor.
	OriginalCategoryIndex int32
	// MinimumLevel is the minimum character level at which this item is available for sale.
	MinimumLevel int32
	// MaximumLevel is the maximum character level at which this item is available for sale.
	MaximumLevel int32
	// Action is the action performed when purchasing this item.
	Action VendorSaleItemActionBlock
	// DisplayCategory is the identifier for the category selling this item.
	DisplayCategory string
	// InventoryBucketHash is the hash of a related InventoryBucketEntity.
	InventoryBucketHash uint32
	// VisibilityScope is the most restrictive scope that determines whether this item is available.
	VisibilityScope GatingScope
	// PurchasableScope is the most restrictive scope that determines whether this item is purchasable.
	PurchasableScope GatingScope
	// Exclusivity is a classification of which platforms this item can be purchased on.
	Exclusivity BungieMembershipType
	// IsOffer determines if this sale can only be performed as the result of an offer check.
	IsOffer bool
	// IsCRM determines if the sale can only performed as the result of receiving a CRM offer.
	IsCrm bool
	// SortValue is the sorting order of this item within its category.
	SortValue int32
	// ExpirationTooltip is the tooltip message that shows when this item expires.
	ExpirationTooltip string
	// RedirectToSaleIndexes are other items that should be purchased when purchasing this item.
	RedirectToSaleIndexes []int32
	SocketOverrides       []VendorItemSocketOverride
	// Unpurchasable determines if this item is a dummy sale item that cannot be purchased.
	Unpurchasable bool
}

// VendorItemQuantity describes item quantity information for vendor prices.
type VendorItemQuantity struct {
	// ItemHash is the hash of a related ItemEntity.
	ItemHash uint32
	// ItemInstanceId is the specific instance ID of this quantity.
	ItemInstanceId int64
	// Quantity is the amount of the item needed/available depending on the context.
	Quantity int32
	// HasConditionalVisibility determines if this item quantity may be conditionally shown or hidden.
	HasConditionalVisibility bool
}

// VendorSaleItemActionBlock is basic cooldown information for a vendor sale.
type VendorSaleItemActionBlock struct {
	// ExecuteSeconds is the amount of time, in seconds, it takes to execute this action.
	ExecuteSeconds float32
	// IsPositive determines if this action is enabled.
	IsPositive bool
}

// VendorItemSocketOverride describes how a vendor purchase should override a given socket with custom plug data.
type VendorItemSocketOverride struct {
	// SingleItemHash is the hash of a related InventoryItemEntity.
	SingleItemHash uint32
	// RandomizedOptionsCount is the number of randomized plugs to set on this socket.
	RandomizedOptionsCount int32
	// SocketTypeHash is the hash of a related SocketTypeEntity.
	SocketTypeHash uint32
}

// InsertPlugAction describes what happens while a plug is being inserted in a given UI.
type InsertPlugAction struct {
	// ActionExecuteSeconds is how much time, in seconds, it takes for this plug to be inserted.
	ActionExecuteSeconds int32
	// ActionType is a classification of the type of action being performed on this socket.
	ActionType SocketTypeActionType
}

// PlugWhitelistEntry defines a plug category that is allowed to be plugged into a socket of this type.
type PlugWhitelistEntry struct {
	// CategoryHash is the hash of this plug's category. Meant to be compared to PlugEntity.PlugCategoryHash.
	CategoryHash uint32
	// CategoryIdentifier is the identifier for this category.
	CategoryIdentifier string
	// ReinitializationPossiblePlugHashes is a list of hashes of InventoryItemEntity structs that describe plug items.
	ReinitializationPossiblePlugHashes []uint32
}

type SocketTypeScalarMaterialRequirementEntry struct {
	CurrencyItemHash uint32
	ScalarValue      int32
}

// VendorAcceptedItem describes an item that a vendor accepts.
// The vault and postmaster are examples of such a vendor.
type VendorAcceptedItem struct {
	// AcceptedInventoryBucketHash is the hash of a related InventoryBucket that acts as the source of this item.
	AcceptedInventoryBucketHash uint32
	// DestinationInventoryBucketHash is the hash of a related InventoryBucket that acts as the destination of this item.
	DestinationInventoryBucketHash uint32
}

// VendorLocation represents a vendor's location and relevant display information.
type VendorLocation struct {
	// DestinationHash is the hash for a related DestinationEntity.
	DestinationHash uint32
	// BackgroundImagePath is the relative path to the background image representing this vendor.
	BackgroundImagePath string
}

// ActivityGraphNode is position and other data related to nodes in the activity graph that launch activities.
type ActivityGraphNode struct {
	// NodeID is a unique identifier for this node within its parent activity graph.
	NodeId uint32
	// OverrideDisplay is a display property that override the active activity's display properties.
	OverrideDisplay DisplayProperties
	// Position is the position on the map for this node.
	Position Position
	// FeaturingStates are a list of possible styles that this node can have.
	FeaturingStates []ActivityGraphNodeFeaturingState
	// Activities are a list possibly active activities for this node.
	Activities []ActivityGraphNodeActivity
	// States are the possible states that this node can be in.
	States []ActivityGraphNodeStateEntry
}

// ActivityGraphNodeActivity is the actual activity to run when this node is clicked on.
type ActivityGraphNodeActivity struct {
	// NodeActivityID is an identifier for this activity within the activity graph.
	NodeActivityId uint32
	// ActivityHash is the hash for a related ActivityEntity.
	ActivityHash uint32
}

// ActivityReward are sets of tooltip-friendly reward data for an activity.
type ActivityReward struct {
	// RewardText is the localized header for this reward set.
	RewardText string
	// RewardItems are the items provided in this reward.
	RewardItems []ItemQuantity
}

// ActivityChallenge represents a reference to a challenge objective.
type ActivityChallenge struct {
	// ObjectiveHash is the hash of a related ObjectiveEntity.
	ObjectiveHash uint32
	// DummyRewards are the rewards for this challenge as shown in the UI.
	DummyRewards []ItemQuantity
}

// ObjectivePerkEntry defines the conditions under which an intrisic perk is applied during an objective.
type ObjectivePerkEntry struct {
	// PerkHash is the hash of a related SandboxPerkEntity.
	PerkHash uint32
	// Style is the classification of when this perk is granted in relation to an objective's progress.
	Style ObjectiveGrantStyle
}

// TalentNodeStepGroups are an attempt to categorize talent node steps by common properties.
type TalentNodeStepGroups struct {
	WeaponPerformance  TalentNodeStepWeaponPerformances
	ImpactEffects      TalentNodeStepImpactEffects
	GuardianAttributes TalentNodeStepGuardianAttributes
	LightAbilities     TalentNodeStepLightAbilities
	DamageTypes        TalentNodeStepDamageTypes
}

// ObjectiveStatEntry defines the conditions under which stat modifications are applied during an objective.
type ObjectiveStatEntry struct {
	// Stat is the stat being modified the value used.
	Stat ItemInvestmentStat
	// Style is the classification of when this stat is granted in relation to an objective's progress.
	Style ObjectiveGrantStyle
}

// ItemInvestmentStat represents a raw investmentment, before calculated stats and stat groups are applied.
type ItemInvestmentStat struct {
	// StatTypeHash is the hash of a related StatEntity.
	StatTypeHash uint32
	// Value is the raw investment value for this stat.
	Value int32
	// IsCondtionallyActive determines if the stat will only be applied on an item in certain game conditions.
	IsConditionallyActive bool
}

// LocationRelease is a specific spot referred to by a location.
type LocationRelease struct {
	DisplayProperties    DisplayProperties
	SmallTransparentIcon string
	MapIcon              string
	LargeTransparentIcon string
	SpawnPoint           uint32
	// DestinationHash is the hash of a related DestinationEntity.
	DestinationHash uint32
	// ActivityHash is the hash of a related ActivityEntity.
	ActivityHash uint32
	// ActivityGraphNodeHash is the activity graph node pointed to by this location.
	ActivityGraphNodeHash uint32
	// ActivityBubbleName is the activity bubble within the Destination.
	ActivityBubbleName      uint32
	ActivityPathBundle      uint32
	ActivityPathDestination uint32
	// NavPointType is a classification of navigation points to this location.
	NavPointType  ActivityNavPointType
	WorldPosition []int32
}

// ActivityPlaylistItem defines a specific entry in an activity playlist.
type ActivityPlaylistItem struct {
	// ActivityHash is the hash of a related ActivityEntity.
	ActivityHash uint32
	// DirectActivityModeHash is the hash of a related ActivityModeEntity.
	DirectActivityModeHash uint32
	// DirectActivityModeType is a classification of the mode related to this activity.
	DirectActivityModeType ActivityModeType
	// ActivityModeHashes are the hash identifiers for all related ActivitymodeEntity structs.
	ActivityModeHashes []uint32
	// ActivityModeTypes are all activity mode's that apply to this activity.
	ActivityModeTypes []ActivityModeType
}

// ActivityMatchmakingBlock is information about matchmaking and party size for an activity.
type ActivityMatchmakingBlock struct {
	// IsMatchmade determines if matchmaking is available for this activity.
	IsMatchmade bool
	// MinParty is the minimum number of players in a fireteam for this activity to launch.
	MinParty int32
	// MaxParty is the maximum number of players in a fireteam for this activity.
	MaxParty int32
	// MaxPlayers is the maximum number of players across all teams in an activity.
	MaxPlayers int32
	// RequiresGuardianOath determines if players are punished for leaving this activity before completion.
	RequiresGuardianOath bool
}

// ActivityGuidedBlock is guided game information for this activity.
type ActivityGuidedBlock struct {
	// GuidedMaxLobbySize is the maximum amount of players allowed in the waiting lobby.
	GuidedMaxLobbySize int32
	// GuidedMinLobbySize is the minimum amount of players allowed in the waiting lobby.
	GuidedMinLobbySize int32
	// GuidedDisbandCount is the total number of votes needed for the guided group to disband.
	GuidedDisbandCount int32
}

type ActivityLoadout struct {
	// EquipmentSlotHash is the hash of a related EquipmentSlotEntity.
	EquipmentSlotHash uint32
	// AllowedEquippedItemHashes are a list of all related InventoryItemEntity structs.
	AllowedEquippedItemHashes []uint32
	// AllowedWeaponSubTypes are a list of all allowed weapon subtypes in activity loadouts.
	AllowedWeaponSubTypes []ItemSubType
}

// ActivityGraphConnection describes how nodes on a graph are visually connected.
type ActivityGraphConnection struct {
	// SourceNodeHash is the hash of the source node.
	SourceNodeHash uint32
	// DestNodeHash is the hash of the destination node.
	DestNodeHash uint32
}

// ActivityGraphDisplayObjective defines active objectives in a graph.
type ActivityGraphDisplayObjective struct {
	// ID is the identifier of this objective in the UI.
	Id uint32
	// ObjectiveHash is the hash of a related ObjectiveEntity.
	ObjectiveHash uint32
}

// ActivityGraphDisplayProgression defines active progressions in a graph.
type ActivityGraphDisplayProgression struct {
	// ID is the identifier of this progression in the UI.
	Id uint32
	// ProgressionHash is the hash of a related ProgressionEntity.
	ProgressionHash uint32
}

// LinkedGraph describes links between the current graph and others.
type LinkedGraph struct {
	Description      string
	Name             string
	UnlockExpression UnlockExpression
	LinkedGraphId    uint32
	// LinkedGraphs are a list of all entries in this linked graph.
	LinkedGraphs []LinkedGraphEntry
	Overview     string
}

// Bubble is basic identifying data about a destination bubble.
type Bubble struct {
	// Hash is the identifier for the bubble which is unique within a destination.
	Hash              uint32
	DisplayProperties DisplayProperties
}

// FactionVendor represents faction vendors at different points in the game.
type FactionVendor struct {
	// VendorHash is the hash of a related VendorEntity.
	VendorHash uint32
	// DestinationHash is the hash of a related DestinationEntity.
	DestinationHash uint32
	// BackgroundImagePath is the relative path to the background image representing this vendor.
	BackgroundImagePath string
}

type ArtifactTier struct {
	// TierHash is an unique identifer for this tier within the artifact.
	TierHash uint32
	// DisplayTitle is the localized title of this tier.
	DisplayTitle string
	// ProgressRequirementMessage represents the localized minimum requirement for this tier.
	ProgressRequirementMessage string
	// Items are the items that can earned within this tier.
	Items []ArtifactTierItem
	// MinimumUnlockPointsUsedRequirement is the minimum number of unlock points needed to unlock this artifact tier.
	MinimumUnlockPointsUsedRequirement int32
}

// ItemQualityBlock is used to determine an item's calculated stats based on quality.
type ItemQualityBlock struct {
	// ItemLevels are the base levels of an item as defined by a season.
	ItemLevels []int32
	// QualityLevel is used in combination with the item's level to calculate stats.
	QualityLevel int32
	// InfusionCategoryName is an identifier for this item's infusability. Deprecated.
	InfusionCategoryName string
	// InfusionCategoryHash is the hash identifier for the infusion. Deprecated.
	InfusionCategoryHash uint32
	// InfusionCategoryHashes are a set of infusion category hashes. If any of these hashes match an entry in
	// InventoryItemEntity.InfusionCategoryHashes, the two can infuse with each other.
	InfusionCategoryHashes []uint32
	// ProgressionLevelRequirementHash is the hash of a related ProgressionLevelRequirementEntity.
	ProgressionLevelRequirementHash uint32
	// CurrentVersion is the latest version available for this item.
	CurrentVersion uint32
	// Versions are the list of versions available for this item.
	Versions []ItemVersion
	// DisplayVersionWatermarkIcons are icon overlays that show the item version and power cap.
	DisplayVersionWatermarkicons []string
}

// ItemValueBlock defines an item's value contextually.
// For example, items sold at a vendor use this to determine the sale price of an item.
type ItemValueBlock struct {
	// ItemValue refers to items that make up this item's value and the quantity.
	ItemValue []ItemQuantity
	// ValueDescription is a localized text description of the value.
	ValueDescription string
}

// ItemSourceBlock describes how an item can be obtained.
type ItemSourceBlock struct {
	// SourceHashes are the hashes for all related RewardSourceEntity structs.
	SourceHashes []uint32
	// Sources are a collection of details about ways the item can be spawned.
	Sources []ItemSource
	// Exclusive is a classification of the platforms this item is exclusive to.
	Exclusive BungieMembershipType
	// VendorSources reference vendors that potentially sell this item.
	VendorSources []ItemVendorSourceReference
}

// ItemSource represents all information about how an item spawns and where a player can obtain it.
type ItemSource struct {
	// Level is the level this item spawns at.
	Level int32
	// MinQuality is the minimum quality this item spawns at.
	MinQuality int32
	// MaxQuality is the maximum quality this item spawns at.
	MaxQuality int32
	// MinLevelRequired is the minimum character level required for equipping the item.
	MinLevelRequired int32
	// MaxLevelRequired is the maximum character level required for equipping the item.
	MaxLevelRequired int32
	// ComputedStats are the stats computed for this level/quality range.
	ComputedStats map[uint32]InventoryItemStat
	// SourceHashes are the hashes for all related RewardSourceEntity structs.
	SourceHashes []uint32
}

// ItemObjectiveBlock represents an item that has objectives associated with it.
type ItemObjectiveBlock struct {
	// ObjectiveHashes are the hashes of all related ObjectiveEntity structs.
	ObjectiveHashes []uint32
	// DisplayActivityHashes are the hashes of all related ActivityEntity structs.
	DisplayActivityHashes []uint32
	// RequireFullObjectiveCompletion determines if all objectives must be completed for the step to be completed.
	RequireFullObjectiveCompletion bool
	// QuestlineItemHash is the hash of a related InventoryItemEntity.
	QuestlineItemHash uint32
	// Narrative is the localized string for narrative text related to this quest step.
	Narrative string
	// ObjectiveVerbName is a localized string describing an action to be performed for this objective.
	ObjectiveVerbName string
	// QuestTypeIdentifier identifies the type of quest being performed.
	QuestTypeIdentifier string
	// QuestTypeHash is a unique hash this quest type.
	QuestTypeHash uint32
	// PerObjectiveDisplayProperties are related display information per objective.
	PerObjectiveDisplayProperties []ObjectiveDisplayProperties
	DisplayAsStatTracker          bool
}

// ObjectiveDisplayProperties are display properties for an item with objectives.
type ObjectiveDisplayProperties struct {
	// ActivityHash is the hash for a related ActivityEntity.
	ActivityHash uint32
	// DisplayOnItemPreviewScreen determines whether this objective is shown on item preview screens.
	DisplayOnItemPreviewScreen bool
}

// PresentationNodeChildrenBlock represents each set of properties that could be a children of a presentation node.
type PresentationNodeChildrenBlock struct {
	PresentationNodes []PresentationNodeChildEntry
	Collectibles      []PresentationNodeCollectibleChildEntry
	Records           []PresentationNodeRecordChildEntry
	Metrics           []PresentationNodeMetricChildEntry
}

// MaterialRequirement defines requirements for spending materials to perform an action.
type MaterialRequirement struct {
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// DeleteOnAction determines if the material will be removed from the inventory when the action is performed.
	DeleteOnAction bool
	// Count is the amount of the material required.
	Count int32
	// OmitFromRequirements determines if this material requirement is shown in a given UI.
	OmitFromRequirements bool
}

type CollectibleStateBlock struct {
	// ObscuredOverrideItemHash is the hash of a related InventoryItemEntity.
	ObscuredOverrideItemHash uint32
	// Requirements are to show this collectible's state.
	Requirements PresentationNodeRequirementsBlock
}

type PresentationChildBlock struct {
	// PresentationNodeType is a classification of this presentation node.
	PresentationNodeType PresentationNodeType
	// ParentPresentationNodeHashes are the hashes for all related PresentationNodeEntity structs.
	ParentPresentationNodeHashes []uint32
	// DisplayStyle is a classification of the display style of this presentation node.
	DisplayStyle PresentationDisplayStyle
}

// RecordTitleBlock describes a title for achieving a record.
type RecordTitleBlock struct {
	// HasTitle determines if this record has an associated title.
	HasTitle bool
	// TitlesByGender are localized titles by localized gender name.
	TitlesByGender map[GenderName]string
	// TitlesByGenderHash are localized titles by GenderEntity.
	TitlesByGenderHash map[uint32]string
	// GildingTrackingRecordHash is the hash of a related RecordEntity.
	GildingTrackingRecordHash uint32
}

// RecordCompletionBlock describes a completed record.
type RecordCompletionBlock struct {
	// PartialCompletionObjectiveCountThreshold is the numbers of objectives that must be completed before the record is completed.
	PartialCompletionObjectiveCountThreshold int32
	// ScoreValue is the amount of triumph points awarded for completing this record.
	ScoreValue int32
	// ShouldFireToast determines if completing this record shows a toast-style pop-up in a given UI.
	ShouldFireToast bool
	// ToastStyle is a classification of record-related toast-style pop-ups.
	ToastStyle RecordToastStyle
}

// RecordStateBlock describes the state of a record.
type RecordStateBlock struct {
	FeaturedPriority int32
	// ObscuredString is the localized string to show if this record is in an obscured state.
	ObscuredString string
}

// RecordExpirationBlock describes the expiration time of a record.
type RecordExpirationBlock struct {
	// HasExpiration determines if the record expires.
	HasExpiration bool
	// Description is the localized description of the expiration time of a record.
	Description string
	// Icon is the path to an icon to show when this record expires.
	Icon string
}

// RecordIntervalBlock describes a record with multiple completions in intervals.
type RecordIntervalBlock struct {
	// IntervalObjectives are a list of objectives per interval.
	IntervalObjectives []RecordIntervalObjective
	// IntervalRewards are a list of rewards per interval.
	IntervalRewards                      []RecordIntervalRewards
	OriginalObjectiveArrayInsertionIndex int32
}

// RecordIntervalObjective is an objective that must be completed to fulfill a record interval.
type RecordIntervalObjective struct {
	// IntervalObjectiveHash is the hash of a related ObjectiveEntity.
	IntervalObjectiveHash uint32
	// IntervalScoreValue is the triumph score awarded for completing this interval objective.
	IntervalScoreValue int32
}

// ItemPlug is an item that can be plugged (socketed) into another.
type ItemPlug struct {
	// InsertionRules are the rules around when this plug can be inserted.
	InsertionRules []PlugRule
	// PlugCategoryIdentifier is localized identifier for this plug's category.
	PlugCategoryIdentifier string
	// PlugCategoryHash is a unique hash for this plug's category.
	PlugCategoryHash uint32
	// OnActionRecreateSelf determines if this plug is refunded when socketed.
	OnActionRecreateSelf bool
	// InsertionMaterialRequirementHash is the hash of a related MaterialRequirementSetEntity for inserting this plug.
	InsertionMaterialRequirementHash uint32
	// PreviewItemOverrideHash is the hash of a related InventoryItemEntity.
	PreviewItemOverrideHash uint32
	// EnabledMaterialRequirementHash is the hash of a related MaterialRequirementSetEntity for enabling this plug.
	EnabledMaterialRequirementHash uint32
	// EnabledRules are the rules around this plug is enabled after being socketed.
	EnabledRules []PlugRule
	// UIPlugLabels is a label that describes the style of this plug in a given UI.
	UiPlugLabels string
	// PlugStyle is a classification of the style of this plug in a given UI.
	PlugStyle PlugUIStyles
	// PlugAvailable indicates the rules about when this plug can be used.
	PlugAvailable PlugAvailabilityMode
	// AlternateUIPlugLabel is a label applied to a plug if it meets certain state requirements.
	AlternateUiPlugLabel string
	// AlternatePlugStyle is a classification of the style of this plug in a given UI under certain conditions.
	AlternatePlugStyle PlugUIStyles
	// IsDummyPlug determines if this plug is used for UI display purposes only, with no other effects.
	IsDummyPlug bool
	// ParentItemOverride describes how this plug overrides properties of the item it is inserted in.
	ParentItemOverride ParentItemOverride
	// EnergyCapacity is the energy capacity provided to the item in which this plug is socketed.
	EnergyCapacity EnergyCapacityEntry
	// EnergyCost is the energy cost of this plug.
	EnergyCost EnergyCostEntry
}

// ItemGearsetBlock describes the items in a related gearset.
type ItemGearsetBlock struct {
	// TrackingValueMax is the maximum number of items that can be collected.
	TrackingValueMax int32
	// ItemList is a list of hashes of all related InventoryItemEntity structs.
	ItemList []uint32
}

// ItemSackBlock describes an item sack: an item that can be opened to produce other items.
type ItemSackBlock struct {
	// DetailAction is a localized description of what happens when this sack is opened.
	DetailAction string
	// OpenAction is localized name of the action being performed when this sack is opened.
	OpenAction string
	// SelectItemCount is the amount of the selected item produced by this sack.
	SelectItemCount int32
	// VendorSackType identifies the type of sack being sold by a vendor.
	VendorSackType string
	// OpenOnAcquire determines if this item sack is opened when it is acquired.
	OpenOnAcquire bool
}

// ItemSocketBlock describes the sockets in an item.
type ItemSocketBlock struct {
	// Detail is a localized description about sockets on this item.
	Detail string
	// SocketEntries are all mutable sockets on an item.
	SocketEntries []ItemSocketEntry
	// IntrinsicSockets are the permanent sockets on an item.
	IntrinsicSockets []ItemIntrinsicSocketEntry
	// SocketCategories are the categories that sockets are grouped into.
	SocketCategories []ItemSocketCategory
}

// ItemSocketEntry defines a specific socket on an item.
type ItemSocketEntry struct {
	// SocketTypeHash is the hash of a related SocketTypeEntity.
	SocketTypeHash uint32
	// SingleInitialItemHash is the hash of a related InventoryItemEntity.
	SingleInitialItemHash uint32
	// ReusablePlugItems is a list of pre-determined that can always be plugged into this socket
	// even if the character does not have the plug item in their inventory.
	ReusablePlugItems []ItemSocketEntryPlugItem
	// PreventInitializationOnVendorPurchase determines if a socket will not be intialized with
	// a plug if the item is purchased from a vendor.
	PreventInitializationOnVendorPurchase bool
	// HidePerksInItemTooltip determines is perks provided by this socket should not be shown in the item's tooltip.
	HidePerksInItemTooltip bool
	// PlugSources indicates where plugs for this socket can be obtained.
	PlugSources SocketPlugSources
	// ReusablePlugSetHash is the hash of a related PlugSetEntity that is reusable.
	ReusablePlugSetHash uint32
	// RandomizedPlugSetHash is the hash of a related PlugSetEntity that is randomized.
	RandomizedPlugSetHash uint32
	// DefaultVisible determines if this socket is visible in the item's default state.
	DefaultVisible bool
}

// ItemSocketEntryPlugItem defines a known, reusable plug that can be applied to a socket.
type ItemSocketEntryPlugItem struct {
	// PlugItemHash is the hash of a related InventoryItemEntity.
	PlugItemHash uint32
}

// ItemSocketEntryPlugItemRandomized defines a randomized plug that can be applied to a socket.
type ItemSocketEntryPlugItemRandomized struct {
	// CurrentlyCanRoll determines if this plug can be rolled on the current version of the item.
	CurrentlyCanRoll bool
	// PlugItemHash is the hash of a InventoryItemEntity.
	PlugItemHash uint32
}

// ItemIntrinsicSocketEntry represents a socket that has a plug associated with it intrisically.
type ItemIntrinsicSocketEntry struct {
	// PlugItemHash is the hash of a related InventoryItemEntity.
	PlugItemHash uint32
	// SocketTypeHash is the hash of a related SocketTypeEntity.
	SocketTypeHash uint32
	// DefaultVisible determines if this socket is visible in the item's default state.
	DefaultVisible bool
}

// ItemSocketCategory describes a grouping of item sockets in a given UI.
type ItemSocketCategory struct {
	// SocketCategoryHash is the hash of a related SocketCategoryEntity.
	SocketCategoryHash uint32
	// SocketIndexes are the indexes of sockets in InventoryItemEntity.Sockets.SocketEntries.
	SocketIndexes []int32
}

// ItemSummaryBlock describes information used when rendering rewards.
type ItemSummaryBlock struct {
	// SortPriority is used as the sort priority when rendering an item.
	SortPriority int32
}

// ItemTalentGridBlock defines an item's talent grid.
type ItemTalentGridBlock struct {
	// TalentGridHash is the hash of a related TalentGridEntity.
	TalentGridHash uint32
	// ItemDetailString is the localized subtitle for looking at a talent in the item's grid.
	ItemDetailString string
	// BuildName is a shortcut for the build, if this talent grid has an associated build.
	BuildName string
	// HUDDamageType is the damage type used in this talent grid.
	HudDamageType DamageType
	// HUDIcon is the icon to show for this talent grid in a given UI.
	HudIcon string
}

// TalentNode is a node attached to the talent grid on an item.
type TalentNode struct {
	// NodeIndex is an index into TalentGridEntity.Nodes.
	NodeIndex int32
	// NodeHash is the unique hash for this node.
	NodeHash uint32
	// Row is the visual row where this node should be shown in a given UI.
	Row int32
	// Column is the visual column where this node should be shown in a given UI.
	Column int32
	// PrerequisiteNodeIndexes are indexes into TalentGrid.Entity.Nodes which must be activity before this node can be activated.
	PrerequisiteNodeIndexes []int32
	// BinaryPairNodeIndex is an index into TalentGridEntity.Nodes for a node that deactivates this node if activated.
	BinaryPairNodeIndex int32
	// AutoUnlocks determines if this node will automatically unlock when the talent grid's level reaches a required level.
	AutoUnlocks bool
	// LastStepRepeats determines if this node can be activated multiple times.
	LastStepRepeats bool
	// IsRandom determines if this node's step is randomly chosen.
	IsRandom bool
	// RandomActivationRequirement is the requirement to repurchase and reactivate this node.
	RandomActivationRequirement NodeActivationRequirement
	// IsRandomRepurchasable determines if this node can be re-rolled.
	IsRandomRepurchasable bool
	// Steps are the steps that must be taken to activate this node.
	Steps []NodeStep
	// ExclusiveWithNodeHashes are hashes for nodes in an exclusive set with this node.
	ExclusiveWithNodeHashes []uint32
	// RandomStartProgressionBarAtProgression is the amount of experience in this item's talent grid is the step is randomly selected.
	RandomStartProgressionBarAtProgression int32
	// LayoutIdentifier identifies a custom visual layout to apply to this node.
	LayoutIdentifier string
	// GroupHash is the exclusive group hash that this node belongs to.
	GroupHash uint32
	// LoreHash is the hash of a related LoreEntity.
	LoreHash uint32
	// NodeStyleIdentifier identifies the node style to use in a given UI.
	NodeStyleIdentifier string
	// IgnoreForCompletion determines if this node should be ignored for completion.
	IgnoreForCompletion bool
}

// NodeStep defines the properties of a talent node step.
type NodeStep struct {
	DisplayProperties DisplayProperties
	// StepIndex is the index of this step in the list of steps on TalentNode.
	StepIndex int32
	// NodeStepHash is the hash of this node step, used to uniquely identify this step within a node.
	NodeStepHash uint32
	// InteractionDescription is a localized description of how this step can be interacted with.
	InteractionDescription string
	// DamageType is a classification of the damage type granted by activating this step.
	DamageType DamageType
	// DamageTypeHash is the hash of a related DamageTypeEntity.
	DamageTypeHash uint32
	// NodeActivationRequirement is the requirement for activating this step.
	ActivationRequirement NodeActivationRequirement
	// CanActivateNextStep determines if activating this step also activates the next step.
	CanActivateNextStep bool
	// NextStepIndex is the index of the next step in the list of steps on TalentNode.
	NextStepIndex int32
	// IsNextStepRandom determines if the next chosen step is random.
	IsNextStepRandom bool
	// PerkHashes are the hashes of all related SandboxPerkEntity structs.
	PerkHashes []uint32
	// StartProgressionBarAtProgress is lower bound of the progress bar for this step.
	StartProgressionBarAtProgress int32
	// StatHashes are the hashes of all related StatEntity structs.
	StatHashes []uint32
	// AffectsQuality determines if this node step affects the item's quality.
	AffectsQuality bool
	// StepGroups is the group of talent node steps this step belongs to based on its functionality.
	StepGroups TalentNodeStepGroups
	// AffectsLevel determines if this node step affects the level of the item.
	AffectsLevel bool
	// SocketReplacements is a list of information used to replace socket items with new plugs when this step is activated.
	SocketReplacements []NodeSocketReplaceResponse
}

// TalentNodeExclusiveSet is a list of indexes into TalentGrid.Nodes for nodes in this exclusive set.
type TalentNodeExclusiveSet struct {
	// NodeIndexes are the list of node indexes for this exclusive set.
	NodeIndexes []int32
}

// ItemPerkEntry describes an intrinsic perk on an item and the requirements to active it.
type ItemPerkEntry struct {
	// RequirementDisplayString is the localized string to show if this perk is not active.
	RequirementDisplayString string
	// PerkHash is the hash of a related SandboxPerkDefinition.
	PerkHash uint32
	// PerkVisibility indicates if this perk should be shown.
	PerkVisibility ItemPerkVisibility
}

// SeasonPreview defines the promotional text, images and link to preview a season.
type SeasonPreview struct {
	// Description is a localized description of the season.
	Description string
	// LinkPath is a localized path to learn about the season.
	LinkPath string
	// VideoLink is an optional link to a localized video about this season.
	VideoLink string
	// Images are a list of images to preview the seasonal content.
	Images []SeasonPreviewImage
}

// SeaosnPreviewImage defines the thumbnail icon and high-resolution image for a season preview.
type SeasonPreviewImage struct {
	// ThumbnailImage is the icon path to preview seasonal content.
	ThumbnailImage string
	// HighResImage is a path to a high-resolution image of the seasonal content.
	HighResImage string
}

type ProgressionRewardItemQuantity struct {
	RewardAtProgressionLevel  int32
	AcquisitionBehavior       ProgressionRewardItemAcquisitionBehavior
	UiDisplayStyle            string
	ClaimUnlockDisplayStrings []string
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash                 uint32
	ItemInstanceId           int64
	Quantity                 int32
	HasConditionalVisibility bool
}

// ObjectiveProgress describes a character's status on a given objective.
type ObjectiveProgress struct {
	// ObjectiveHash is the hash of a related ObjectiveEntity.
	ObjectiveHash uint32
	// DestinationHash is the hash of a related DestinationEntity.
	DestinationHash uint32
	// ActivityHash is the hash of a related ActivityEntity.
	ActivityHash uint32
	// Progress is the numeric value of progress on this objective.
	Progress int32
	// CompletionValue is the numeric value of progress needed to complete this objective.
	CompletionValue int32
	// Complete determines if this objective is completed.
	Complete bool
	// Visible determines if this objective is visible in-game.
	Visible bool
}

// ChecklistEntry defines the properties of an individual checklist item.
type ChecklistEntry struct {
	// Hash is the unique identifier of this entry within the checklist.
	Hash              uint32
	DisplayProperties DisplayProperties
	// DestinationHash is the hash of a related DestinationEntity.
	DestinationHash uint32
	// LocationHash is the hash of a related LocationEntity.
	LocationHash uint32
	// BubbleHash is the hash of an activity bubble in LocationEntity.LocationReleases.
	BubbleHash uint32
	// ActivityHash is the hash of a related ActivityEntity.
	ActivityHash uint32
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// VendorHash is the hash of a related VendorEntity.
	VendorHash uint32
	// VendorInteractionIndex is this entry's index in VendorEntity.Interactions.
	VendorInteractionIndex int32
	// Scope is the scope at which this specific entry can be completed.
	Scope Scope
}

// MilestoneQuest describes a quest item that is currently active for a milestone.
type MilestoneQuest struct {
	// QuestItemHash is the hash of a related InventoryItemEntity.
	QuestItemHash     uint32
	DisplayProperties DisplayProperties
	// OverrideImage is the image to show instead of the milestone's image.
	OverrideImage string
	// QuestRewards are rewards for completing this quest.
	QuestRewards MilestoneQuestRewards
	// Activities are milestone-related activities by the hash of a related ActivityEntity.
	Activities map[uint32]MilestoneActivity
	// DestinationHash is the hash of a related DestinationEntity.
	DestinationHash uint32
}

// MilestoneQuestRewards describe rewards given for completing a milestone-related quest.
type MilestoneQuestRewards struct {
	// Items are the items rewarded for completing the quest.
	Items []MilestoneQuestRewardItem
}

// MilestoneQuestRewardItem describes an milestone-related quest reward and the vendor that provides it.
type MilestoneQuestRewardItem struct {
	// VendorHash is the hash of a related VendorEntity.
	VendorHash uint32
	// VendorItemIndex is the index of item being sold in the vendor.
	VendorItemIndex int32
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// ItemInstanceID refers to a specific instance of an item.
	ItemInstanceId int64
	// Quantity is the amount of the item.
	Quantity int32
	// HasCondtionalVisibility determines if the item quantity may be condtionally shown or hideen.
	HasConditionalVisibility bool
}

// MilestoneActivity is a milestone-related activity.
type MilestoneActivity struct {
	// ConceptualActivityHash is the hash of a related ActivityEntity.
	ConceptualActivityHash uint32
	// Variants are the variants of this milestone-related activity by the hash of a related ActivityEntity.
	Variants map[uint32]MilestoneActivityVariant
}

// MilestoneActivityVariant represents a variant on a milestone-related activity.
type MilestoneActivityVariant struct {
	// ActivityHash is the hash of a related ActivityEntity.
	ActivityHash uint32
	// Order is the order to arrange this variant by when displaying a milestone-related activity.
	Order int32
}

// MilestoneRewardCategory defines a category of milestone-related rewards.
type MilestoneRewardCategory struct {
	// CategoryHash is the unique hash of this reward category.
	CategoryHash uint32
	// CategoryIdentifier is unique identifier of this category.
	CategoryIdentifier string
	DisplayProperties  DisplayProperties
	// RewardEntries are milestone-related rewards by the reward category hash.
	RewardEntries map[uint32]MilestoneRewardEntry
	// Order is the recommended order for rendering this category.
	Order int32
}

// MilestoneRewardEntry defines a specific reward in a category of milestone-related rewards.
type MilestoneRewardEntry struct {
	// RewardEntryHash is the unique hash of this reward entry within the milestone.
	RewardEntryHash uint32
	// RewardEntryIdentifier is the unique identifier of this reward entry within the milestone.
	RewardEntryIdentifier string
	// Items are the items and the quantity rewarded.
	Items []ItemQuantity
	// VendorHash is the hash of a related VendorEntity.
	VendorHash        uint32
	DisplayProperties DisplayProperties
	// Order is the ordering of this reward in a given UI.
	Order int32
}

// MilestoneVendor describes a milestone-related vendor.
type MilestoneVendor struct {
	// VendorHash is the hash of a related VendorEntity.
	VendorHash uint32
}

// MilestoneValue contains information related to the key/value pair for a specific milestone or milestone step.
type MilestoneValue struct {
	// Key is the for the specific milestone or milestone step.
	Key               string
	DisplayProperties DisplayProperties
}

// MilestoneChallengeActivity describes a milestone-related challenge.
type MilestoneChallengeActivity struct {
	// ActivityHash is the hash of a related ActivityEntity.
	ActivityHash uint32
	// Challenges are the challenges related to this activity.
	Challenges []MilestoneChallenge
	// ActivityGraphNodes are the nodes this challenge is visible on.
	ActivityGraphNodes []MilestoneChallengeActivityGraphNodeEntry
	// Phases are the phases related to this activity.
	Phases []MilestoneChallengeActivityPhase
}

// MilestoneChallenge describes the objective of a milestone-related challenge.
type MilestoneChallenge struct {
	// ChallengeObjectiveHash is the hash of a related ObjectiveEntity.
	ChallengeObjectiveHash uint32
}

// MilestoneChallengeActivityGraphNodeEntry describes the map of a milestone-related challenge.
type MilestoneChallengeActivityGraphNodeEntry struct {
	// ActivityGraphHash is the hash of a related ActivityGraphEntity.
	ActivityGraphHash uint32
	// ActivityGraphNodeHash is hash of this milestone-related challenge within a map.
	ActivityGraphNodeHash uint32
}

// MilestoneChallengeActivityPhase describes a phase in a milestone-related challenge.
type MilestoneChallengeActivityPhase struct {
	// PhaseHash is the unique hash representing the phase.
	PhaseHash uint32
}

// ReportReason describes a specific reason for being banned.
type ReportReason struct {
	// ReasonHash is the unique hash of this reason in the report category.
	ReasonHash        uint32
	DisplayProperties DisplayProperties
}
