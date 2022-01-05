package destiny2

import "time"

// ProgressionEntity is an entity in the Destiny.Definitions.DestinyProgressionDefinition contract.
// This represents a measurement of progress through a set of conceptual steps.
type ProgressionEntity struct {
	DisplayProperties DisplayProperties
	// Scope indicates the source of the progression's live data.
	Scope ProgressionScope
	// RepeatLastStep indicates if the last step of a progression can be indefinitely repeated.
	RepeatLastStep bool
	// Source is the localized description of how to earn this progression.
	Source string
	// Steps are the individual steps to complete this progression.
	Steps []ProgressionStep
	// Visible determines if the progression should be shown to users.
	Visible bool
	// FactionHash is the hash of a related FactionEntity.
	FactionHash uint32
	// Color is the RGB-valued related to this progression.
	Color Color
	// RankIcon is the icon used in the Companion app to display a progression's rank value.
	RankIcon string
	// RewardItems are the items rewarded for completing are steps.
	RewardItems []ProgressionRewardItemQuantity
	EntityMetadata
}

// InventoryItemEntity is an entity in the Destiny.Definitions.DestinyInventoryItemDefinition contract.
// This defines items in Destiny which can exist in your inventory.
type InventoryItemEntity struct {
	DisplayProperties DisplayProperties
	// TooltipNotifications are tooltips that come up conditionally for the item.
	TooltipNotifications []ItemTooltipNotification
	// CollectibleHash is the hash of a related CollectibleEntity.
	CollectibleHash uint32
	// IconWatermark is the original release watermark overlay for the icon.
	// If the item has different versions, it can be overriden by the watermark in the quality block or the IconWatermarkShelved field.
	IconWatermark string
	// IconWatermarkShelved is the "shelved" release watermark overlay for the icon.
	// If the item version has a power cap below the season power cap, it should show the "shelved" watermark.
	IconWatermarkShelved string
	// SecondaryIcon is a secondary icon associated with the item; used in places such as emblem nameplates.
	SecondaryIcon string
	// SecondaryOverlay is the background overlay for SecondaryIcon.
	SecondaryOverlay string
	// BackgroundColor is the color used for the icon background of an item.
	BackgroundColor Color
	// Screenshot is the path to an in-game screenshot for this item.
	Screenshot string
	// ItemTypeDisplayName is the localized title/name of the item's type.
	ItemTypeDisplayName string
	// FlavorText is the quote or description of a given item.
	FlavorText string
	// UiItemDisplayStyle is how a UI should render this item in an inventory screen.
	UiItemDisplayStyle string
	// ItemTypeAndTierDisplayName is the localized string that combines the item type and tier description.
	ItemTypeAndTierDisplayName string
	// DisplaySource is the localized string describing how to find this item.
	DisplaySource string
	// TooltipStyle determines what tooltip a UI should show for this item.
	TooltipStyle string
	// Action is an action performed when using the item.
	Action ItemActionBlock
	// Inventory describes this item's relationship with its inventory.
	Inventory ItemInventoryBlock
	// SetData contains data about the steps and mechanics if this item is quest-related.
	SetData ItemSetBlock
	// Stats are the statistics related to equippable items.
	Stats ItemStatBlock
	// EmblemObjectiveHash is the hash of a related ObjectiveEntity.
	EmblemObjectiveHash uint32
	// EquippingBlock describes the conditions under which this item can be equipped.
	EquippingBlock EquippingBlock
	// TranslationBlock contains rendering information for this item.
	TranslationBlock ItemTranslationBlock
	// Preview summarizes items that can be acquired from using this item.
	Preview ItemPreviewBlock
	// Quality describes the item's quality.
	Quality ItemQualityBlock
	// Value describes an item's value; i.e. sale price for a vendor item and dismantle reward.
	Value ItemValueBlock
	// SourceData describes how to acquire this item.
	SourceData ItemSourceBlock
	// Objectives are extra tasks related to this item if it is part of a quest step.
	Objectives ItemObjectiveBlock
	// Metrics are all available metrics for this item.
	Metrics ItemMetricBlock
	// Plug is the information about how to insert a socket-type item.
	Plug ItemPlug
	// GearSet is the set of equipment associated with this item.
	Gearset ItemGearsetBlock
	// Sack describes what items can be obtained by opening this item.
	Sack ItemSackBlock
	// Sockets describes any sockets available on this item.
	Sockets ItemSocketBlock
	// Summary is this item's reward summary.
	Summary ItemSummaryBlock
	// TalentGrid is the information of the talents in a subclass-type item.
	TalentGrid ItemTalentGridBlock
	// InvestmentStats are the raw stats for this item, before any adjustments.
	InvestmentStats []ItemInvestmentStat
	// Perks are intrisic perks of an item.
	Perks []ItemPerkEntry
	// LoreHash is the hash of a related LoreEntity.
	LoreHash uint32
	// SummaryItemHash is the hash of a related InventoryItemEntity.
	SummaryItemHash uint32
	// Animations defines are animations in the game content for this item.
	Animations []AnimationReference
	// AllowActions determines whether or not the API is allowed to take actions which effect this item.
	AllowActions bool
	// Links are any help or informational URLs about this item.
	Links []HyperlinkReference
	// DoesPostmaterPullHaveSideEffects indicates whether something could happen when transferring this item from the postmaster.
	DoesPostmasterPullHaveSideEffects bool
	// NonTransferrable determines whether an item cannot be transferred
	NonTransferrable bool
	// ItemCategoryHashes are hashes of all related ItemCategoryEntity structs.
	ItemCategoryHashes []uint32
	// SpecialItemType is a internal item category from Destiny 1. Use ItemCategoryHashes instead.
	SpecialItemType SpecialItemType
	// ItemType is the base type of this item. Use ItemCategoryHashes instead.
	ItemType ItemType
	// ItemSubType is the sub-type of this item. Use ItemCategoryHashes instead.
	ItemSubType ItemSubType
	// ClassType is the specific class this item is restricted to.
	ClassType Class
	// BreakerType is the type of anti-champion ability granted by using this item.
	BreakerType BreakerTypeEnum
	// BreakerTypeHash is the hash of a related BreakerTypeEntity.
	BreakerTypeHash uint32
	// Equippable indicates whether an item can be equipped.
	Equippable bool
	// DamageTypeHashes are hashes of all related DamageTypeEntity structs.
	DamageTypeHashes []uint32
	// DamageTypes lists all damage types this item can have.
	DamageTypes []DamageType
	// DefaultDamageType is the default damage type of this item.
	DefaultDamageType DamageType
	// DefaultDamageTypeHash is the hash of a related DamageTypeEntity.
	DefaultDamageTypeHash uint32
	// SeasonHash is the hash of a related SeasonEntity.
	SeasonHash uint32
	// IsWrapper determines if this item is a vendor-wrapper that can be refunded before it is unwrapped.
	IsWrapper bool
	// TraidIds are metadata tags applied to this item.
	TraitIds []string
	// TraitHashes are hashes of all related TraitEntity structs.
	TraitHashes []uint32
	EntityMetadata
}

// InventoryBucketEntity is an entity in the Destiny.Definitions.DestinyInventoryBucketDefinition contract.
// This describes a collection of items in an inventory known as a bucket. For example, "Primary Weapons".
type InventoryBucketEntity struct {
	DisplayProperties DisplayProperties
	// Scope is where this bucket can be found.
	Scope BucketScope
	// Category describes what items can be found in this bucket.
	Category BucketCategory
	// BucketOrder is a recommended ordering for this bucket in a given UI.
	BucketOrder int32
	// ItemCount is the maximum number of item slots in a bucket.
	ItemCount int32
	// Location is the conceptual location of this bucket, such as the vault or postmaster inventory.
	Location ItemLocation
	// HasTransferDestination indicates if items can be transferred to/from this bucket.
	HasTransferDestination bool
	// Enabled determines if this bucket is being used in-game.
	Enabled bool
	// Fifo determines if the bucket will delete the oldest item in a bucket when adding a new one (FIFO)
	// or if this bucket disallows new items until there is space.
	Fifo bool
	EntityMetadata
}

// ItemTierTypeEntity is an entity in the Destiny.Definitions.Items.DestinyItemTierTypeDefinition contract.
// This defines the tier of items such as common, rare, etc, and how they can be infused.
type ItemTierTypeEntity struct {
	DisplayProperties DisplayProperties
	// InfusionProcess defines the process of infusing items of this tier.
	InfusionProcess ItemTierTypeInfusionBlock
	EntityMetadata
}

// StatEntity is an entity in the Destiny.Definitions.DestinyStatDefinition contract.
// This represents a stat that is applied to a character or an item.
type StatEntity struct {
	DisplayProperties DisplayProperties
	// AggregationType is how to aggregate existing stats with this stat.
	AggregationType StatAggregationType
	// HasComputedBlock determines if this stat is computed beforehand or used as a raw value.
	HasComputedBlock bool
	// StatCategory is categorization of this stat.
	StatCategory StatCategory
	EntityMetadata
}

// StatGroupEntity is an entity in the Destiny.Definitions.DestinyStatGroupDefinition contract.
// This defines the grouping a stat belongs to when displaying aggregated stats.
type StatGroupEntity struct {
	// MaximumValue is the highest value any stat in this group can be aggregated to.
	MaximumValue int32
	// UIPosition is information about where to display this stat in a given UI.
	UiPosition int32
	// ScaledStats are all stats in this group that require scaling to be applied before being displayed.
	ScaledStats []StatDisplay
	// Overrides are the localized texts to display for stats in this group.
	// This maps from the hash of a StatEntity to the localized text.
	Overrides map[uint32]StatOverride
	EntityMetadata
}

// EquipmentSlotEntity is an entity in the Destiny.Definitions.DestinyEquipmentSlotDefinition contract.
// This is an indicator that a related bucket can have instanced items equipped in a specific slot on a character.
type EquipmentSlotEntity struct {
	// EquipmentCategoryHash is supposed to be a related EquipmentCategoryEntity which has not been exposed in the Bungie.net API.
	// The specification suggests using this hash as a general way to group equipment by common functionality.
	EquipmentCategoryHash uint32
	// BucketTypeHash is the hash of a related InventoryBucketEntity.
	BucketTypeHash uint32
	// ApplyCustomArtDays determines if equipped items should have custom art dyes applied during rendering.
	ApplyCustomArtDyes bool
	// ArtDyeChannels are all art dyes that apply this equipment slot.
	ArtDyeChannels []ArtDyeReference
	EntityMetadata
}

// VendorEntity is an entity in the Destiny.Definitions.DestinyVendorDefinition contract.
// Vendors range from NPCs that you can buy items from to Kiosks, Collections and the Vault.
type VendorEntity struct {
	DisplayProperties DisplayProperties
	// VendorProgresionType is the type of reward progression this vendor has.
	VendorProgressionType VendorProgressionType
	// BuyString is the localized string describing the buy action for this vendor.
	BuyString string
	// SellString is the localized string describing the sell action for this vendor.
	SellString string
	// DisplayItemHash is the hash for a related InventoryItemEntity.
	DisplayItemHash uint32
	// InhibitBuying determines if a player is not allowed to buy from this vendor.
	InhibitBuying bool
	// InhibitSelling determines if a player is not allowed to sell to this vendor.
	InhibitSelling bool
	// FactionHash is the hash for a related FactionEntity.
	FactionHash uint32
	// ResetIntervalMinutes is the frequency of this vendor's inventory refresh, in minutes.
	ResetIntervalMinutes int32
	// ResetOffsetMinutes is the offest of the frequency of this vendor's inventory refresh, in minutes.
	ResetOffsetMinutes int32
	// FailureStrings are localized strings for a failure to purchase an item from this vendor.
	FailureStrings []string
	// UnlockRanges are the dates with this vendor will be available.
	UnlockRanges []DateRange
	// VendorIdentifier is an internal identifier for this vendor.
	VendorIdentifier string
	// VendorPortrait is the path to the image of this vendor.
	VendorPortrait string
	// VendorBanner is the custom banner image associated with this vendor.
	VendorBanner string
	// Enabled determines if this vendor exists.
	Enabled bool
	// Visible determines if this vendor is visible.
	Visible bool
	// VendorSubcategoryIdentifier is the identifier for this vendor's subcategory.
	VendorSubcategoryIdentifier string
	// ConsolidateCategories determines if similar vendor categories should be consolidated.
	ConsolidateCategories bool
	// Actions describes actions that can be performed at this vendor.
	Actions []VendorAction
	// Categories are the headers of sections of items this vendor is selling.
	Categories []VendorCategoryEntry
	// OriginalCategories are pre-consolidated categories.
	OriginalCategories []VendorCategoryEntry
	// DisplayCategories are hints on how to visually group and display categories in this vendor's UI.
	DisplayCategories []DisplayCategory
	// Interactions are non-selling interactions that can happen with this vendor.
	Interactions []VendorInteraction
	// InventoryFlyouts describes this vendor's UI for showing items from the player's inventory.
	InventoryFlyouts []VendorInventoryFlyout
	// ItemList is the list of items this vendor can sell.
	ItemList []VendorItem
	// Services is a list of flavor text about services that this vendor provides.
	Services []VendorService
	// AcceptedItems is a list of source and destination buckets for transfer, if this vendor is a postmaster or vault.
	AcceptedItems []VendorAcceptedItem
	// ReturnWithVendorRequest determines if information about this vendor is returned from the Bungie.net API.
	ReturnWithVendorRequest bool
	// Locations are the locations this vendor can be depending on the game state.
	Locations []VendorLocation
	// Groups are the collections this vendor are related to by either location or function.
	Groups []VendorGroupReference
	// IgnoreSaleItemHashes are related InventoryItemEntity structs that should be ignored if sold by this vendor.
	IgnoreSaleItemHashes []uint32
	EntityMetadata
}

// SocketTypeEntity is an entity in the Destiny.Definitions.Sockets.DestinySocketTypeDefinition contract.
// All sockets have a type which are a set of common properties that determine when and what plugs can be inserted.
type SocketTypeEntity struct {
	DisplayProperties DisplayProperties
	// InsertAction defines what happens when a plug is inserted into this socket type.
	InsertAction InsertPlugAction
	// PlugWhitelist is a list of plug categories that are allowed to be plugged into this socket type.
	PlugWhitelist []PlugWhitelistEntry
	// SocketCategoryHash is the hash of a related SocketCategoryEntity.
	SocketCategoryHash uint32
	// Visibility is the condition under which this socket type is visible.
	Visibility                 SocketTypeVisibility
	AlwaysRandomizeSockets     bool
	IsPreviewEnabled           bool
	HideDuplicateReusablePlugs bool
	// OverridesUIAppearance determines if this socket type should be overriden by the inserted plug's icon.
	OverridesUiAppearance           bool
	AvoidDuplicatesOnInitialization bool
	CurrencyScalars                 []SocketTypeScalarMaterialRequirementEntry
	EntityMetadata
}

// SocketCategoryEntity is an entity in the Destiny.Definitions.Sockets.DestinySocketCategoryDefinition contract.
// Sockets on an item are organized into categories visually.
type SocketCategoryEntity struct {
	DisplayProperties DisplayProperties
	// UICategoryStyle is a hint about how the sockets in this category should be displayed.
	UiCategoryStyle uint32
	// CategoryStyle is the display style of this category.
	CategoryStyle SocketCategoryStyle
	EntityMetadata
}

// DestinationEntity is an entity in the Destiny.Definitions.DestinyDestinationDefinition contract.
// A destination is a specific area of a larger place.
type DestinationEntity struct {
	DisplayProperties DisplayProperties
	// PlaceHash is the hash of a related PlaceEntity.
	PlaceHash uint32
	// DefaultFreeroamActivityHash is the hash of a related ActivityEntity.
	DefaultFreeroamActivityHash uint32
	// ActivityGraphEntries are the list of maps that are shown in the director for this destination.
	ActivityGraphEntries []ActivityGraphListEntry
	// BubbleSettings are zones within this destination. Deprecated: use Bubbles.
	BubbleSettings []DestinationBubbleSetting
	// Bubbles are unique identifiers for every bubble zone in this destination.
	Bubbles []Bubble
	EntityMetadata
}

// ActivityGraphEntity is an entity in the Destiny.Definitions.Director.DestinyActivityGraphDefinition contract.
// This represents a map view in the director.
type ActivityGraphEntity struct {
	// Nodes are the visual nodes on this map's view which correspond to activities that can be clicked on.
	Nodes []ActivityGraphNode
	// ArtElements are special UI elements that appear on this map.
	ArtElements []ActivityGraphArtElement
	// Connections represent connections between Nodes.
	Connections []ActivityGraphConnection
	// DisplayObjectives are objectives that can be displayed on this map.
	DisplayObjectives []ActivityGraphDisplayObjective
	// DisplayProgressions are progressions that can be displayed on this map.
	DisplayProgressions []ActivityGraphDisplayProgression
	// LinkedGraphs links this Activity Graph to others.
	LinkedGraphs []LinkedGraph
	EntityMetadata
}

// ActivityEntity is an entity in the Destiny.Definitions.DestinyActivityDefinition contract.
// This is the static data about a Destiny 2 activity.
type ActivityEntity struct {
	// DisplayProperties are the post-processed display properties for this activity.
	DisplayProperties DisplayProperties
	// OriginalDisplayProperties are the unprocessed display properties for this activity.
	OriginalDisplayProperties DisplayProperties
	// SelectionScreenDisplayProperties are the display properties shown in the selection screen for this activity.
	SelectionScreenDisplayProperties DisplayProperties
	// ReleaseIcon is an icon associated with this activity for a specific content release.
	ReleaseIcon string
	// ReleaseTime is when this activity is visible, in seconds.
	ReleaseTime int32
	// ActivityLightLevel is the recommended light level for this activity.
	ActivityLightLevel int32
	// DestinationHash is the hash for a related DestinationEntity.
	DestinationHash uint32
	// PlaceHash is the hash for a related PlaceEntity.
	PlaceHash uint32
	// ActivityTypeHash is the hash for related ActivityTypeEntity.
	ActivityTypeHash uint32
	// Tier is the difficulty of this activity.
	Tier ActivityDifficultyTier
	// PGCRImage is the image that shows the Post-Game Carnage Report for this activity.
	PgcrImage string
	// Rewards are the possible rewards for this activity.
	Rewards []ActivityReward
	// Modifiers are the modifiers that can be applied to this activity.
	Modifiers []ActivityModifierReference
	// IsPlaylist determines if this activity is in a activity mode's playlist.
	IsPlaylist bool
	// Challenges are the challenges applied to this activity.
	Challenges []ActivityChallenge
	// OptionalUnlockStrings are the statuses related to this activity.
	OptionalUnlockStrings []ActivityUnlockString
	// PlaylistItems represent all the this activity in this playlist activity.
	PlaylistItems []ActivityPlaylistItem
	// ActivityGraphList are the maps to show in the director during this activity.
	ActivityGraphList []ActivityGraphListEntry
	// Matchmaking describes how many people can join an activity and if matchmaking is enabled.
	Matchmaking ActivityMatchmakingBlock
	// GuidedGame provides information about the guided game experience for this activity.
	GuidedGame ActivityGuidedBlock
	// DirectActivityModeHash is the hash of a related ActivityModeEntity.
	DirectActivityModeHash uint32
	// DirectActivityModeType is the mode of this activity.
	DirectActivityModeType ActivityModeType
	// Loadouts are all the possible loadout requirements that could be active for this activity.
	Loadouts []ActivityLoadoutRequirementSet
	// ActivityModeHashes are all ActivityModeEntity structs related to this activity.
	ActivityModeHashes []uint32
	// ActivityModeTypes are all modes related to this activity.
	ActivityModeTypes []ActivityModeType
	// IsPvP determines if this activity is PVP-related.
	IsPvP bool
	// InsertionPoints is a list of phases or points of entry into this activity.
	InsertionPoints []ActivityInsertionPoint
	// ActivityLocationMappings are list of locations affected by this activity.
	ActivityLocationMappings []EnvironmentLocationMapping
	EntityMetadata
}

// ActivityModifierEntity is an entity in the Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition contract.
// Modifiers are just descriptions of the modifiers active on an activity.
type ActivityModifierEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// ObjectiveEntity is an entity in the Destiny.Definitions.DestinyObjectiveDefinition contract.
// An objective is a specific task that can be accomplished in the game.
type ObjectiveEntity struct {
	DisplayProperties DisplayProperties
	// CompletionValue is the value that must be reached in order for this objective be completed.
	CompletionValue int32
	// Scope determines the most restrictive gating this objective is set to use.
	Scope GatingScope
	// LocationHash is the hash of a related LocationEntity.
	LocationHash uint32
	// AllowNegativeValue determines if the CompletionValue is allowed to go negative.
	AllowNegativeValue bool
	// AllowValueChangeWhenCompleted determines whether this objective can be uncompleted by losing progress.
	AllowValueChangeWhenCompleted bool
	// IsCountingDownward determines if completion of this objective means having a value less than or equal to CompletionValue.
	IsCountingDownward bool
	// ValueStyle is the UI style applied to this objective.
	ValueStyle UnlockValueUIStyle
	// ProgressDescription is the description on this objective's progress bar.
	ProgressDescription string
	// Perks are the conditions for enabling this objective's perks.
	Perks ObjectivePerkEntry
	// Stats are the condtions for modifying a player's stats in this objective.
	Stats ObjectiveStatEntry
	// MinimumVisibilityThreshold is the minimum value that this objective's progression should be shown.
	MinimumVisibilityThreshold int32
	// AllowOvercompletion determines if this objective's progress will continue beyond completion requirements.
	AllowOvercompletion bool
	// ShowValueOnComplete determines if the progression value should be shown in the UI after it's complete.
	ShowValueOnComplete bool
	// CompletedValueStyle is the style to use when the objective is completed.
	CompletedValueStyle UnlockValueUIStyle
	// InProgressValueStyle is the style to use when the objective is still in progress.
	InProgressValueStyle UnlockValueUIStyle
	EntityMetadata
}

// SandboxPerkEntity is an entity in the Destiny.Definitions.DestinySandboxPerkDefinition contract.
// Perks are modifiers to a character or item that can be applied situationally.
type SandboxPerkEntity struct {
	DisplayProperties DisplayProperties
	// PerkIdentifier is the identifier for this perk.
	PerkIdentifier string
	// IsDisplayable determines if this perk is shown in a given UI.
	IsDisplayable bool
	// DamageType is the damage type granted to a weapon by this perk.
	DamageType DamageType
	// DamageTypeHash is the hash for a related DamageTypeEntity.
	DamageTypeHash uint32
	// PerkGroups are a way to group perks by functionality.
	PerkGroups TalentNodeStep
	EntityMetadata
}

// LocationEntity is an entity in the Destiny.Definitions.DestinyLocationDefinition contract.
// A Location refers to a specific combination of activity, destination, place and bubble.
type LocationEntity struct {
	// VendorHash is the hash for a related VendorEntity.
	VendorHash uint32
	// LocationReleases refer to different spots in the world for this location and which location is valid.
	LocationReleases []LocationRelease
	EntityMetadata
}

// ActivityModeEntity is an entity in the Destiny.Definitions.DestinyActivityModeDefinition contract.
// This represents a collection activities that are played in a certain way.
type ActivityModeEntity struct {
	DisplayProperties DisplayProperties
	// PGCRImage is the Post-Game Carnage Report for this activity mode.
	PgcrImage string
	// ModeType is this activity mode's type.
	ModeType ActivityModeType
	// ActivityModeCategory is the category of this activity mode.
	ActivityModeCategory ActivityModeCategory
	// IsTeamBased determines if this activity mode is team-based, co-op or free-for-all.
	IsTeamBased bool
	// IsAggregateMode determines if this activity mode is an aggregate of other specific modes.
	IsAggregateMode bool
	// ParentHashes are the hashes all ActivityModeEntity structs related this activity mode.
	ParentHashes []uint32
	// FriendlyName is an identifier used for referring to this activity mode.
	FriendlyName string
	// ActivityModeMappings are the ActivityEntity structs and the mode type for this activity mode.
	ActivityModeMappings map[uint32]ActivityModeType
	// Display determines if this activity mode is shown in a given UI.
	Display bool
	// Order is the relative ordering of activity modes.
	Order int32
	EntityMetadata
}

// PlaceEntity is an entity in the Destiny.Definitions.DestinyPlaceDefinition contract.
// Places are synonymous to planets that can be visited via the director.
type PlaceEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// ActivityTypeEntity is an entity in the Destiny.Definitions.DestinyActivityTypeDefinition contract.
// This represents a conceptual categorization of activities.
type ActivityTypeEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// VendorGroupEntity is an entity in the Destiny.Definitions.DestinyVendorGroupDefinition contract.
// This categorizes vendors by similarity.
type VendorGroupEntity struct {
	// Order is the recommended order for rendering the groups in ascending order.
	Order int32
	// CategoryName is the name of this vendor group.
	CategoryName string
	EntityMetadata
}

// FactionEntity is an entity in the Destiny.Definitions.DestinyFactionDefinition contract.
// This represents a given faction in the game.
type FactionEntity struct {
	DisplayProperties DisplayProperties
	// ProgressionHash is the hash of a related ProgressionEntity.
	ProgressionHash uint32
	// TokenValues are the hashes of InventoryItemEntity structs and their progression value for this faction.
	TokenValues map[uint32]uint32
	// RewardItemHash is the hash of a related InventoryItemEntity.
	RewardItemHash uint32
	// RewardVendorHash is the hash of a related VendorEntity.
	RewardVendorHash uint32
	// Vendors are a list of vendors associated with this faction.
	Vendors []FactionVendor
	EntityMetadata
}

// ArtifactEntity is an entity in the Destiny.Definitions.Artifacts.DestinyArtifactDefinition contract.
// Represents known info about a seasonal artifact.
type ArtifactEntity struct {
	DisplayProperties DisplayProperties
	// TranslationBlock is rendering information for this artifact.
	TranslationBlock ItemTranslationBlock
	// Tiers are tier/rank data related to this artifact.
	Tiers []ArtifactTier
	EntityMetadata
}

// PowerCapEntity is an entity in the Destiny.Definitions.PowerCaps.DestinyPowerCapDefinition contract.
// Defines the power cap limit for gear items, based on rarity and season.
type PowerCapEntity struct {
	// PowerCap is the raw value for a power cap.
	PowerCap int32
	EntityMetadata
}

// ProgressionLevelRequirementEntity is an entity in the Destiny.Definitions.Progression.DestinyProgressionLevelRequirementDefinition contract.
// This determines the level requirement for an item given a progression.
type ProgressionLevelRequirementEntity struct {
	// RequirementCurve is the curve of level requirements, weighted by this progression level.
	RequirementCurve []InterpolationPointFloat
	// ProgressionHash is the hash for a related ProgressionEntity.
	ProgressionHash uint32
	EntityMetadata
}

// RewardSourceEntity is an entity in the Destiny.Definitions.DestinyRewardSourceDefinition contract.
// This represents an item source for a given reward.
type RewardSourceEntity struct {
	DisplayProperties DisplayProperties
	// Category is the grouping this reward source.
	Category RewardSourceCategory
	EntityMetadata
}

// TraitEntity is an entity in the Destiny.Definitions.Traits.DestinyTraitDefinition contract.
type TraitEntity struct {
	DisplayProperties DisplayProperties
	TraitCategoryId   string
	// TraitCategoryHash is the hash of a related TraitCategoryEntity.
	TraitCategoryHash uint32
	EntityMetadata
}

// TraitCategoryEntity is an entity in the Destiny.Definitions.Traits.DestinyTraitCategoryDefinition contract.
type TraitCategoryEntity struct {
	TraitCategoryId string
	TraitHashes     []uint32
	TraitIds        []string
	EntityMetadata
}

// PresentationNodeEntity is an entity in the Destiny.Definitions.Presentation.DestinyPresentationNodeDefinition contract.
// This represents a logical grouping of other entities visually/organizationally.
type PresentationNodeEntity struct {
	DisplayProperties DisplayProperties
	// OriginalIcon is the original icon for this presentation node.
	OriginalIcon string
	// RootViewIcon is the icon meant to be shown on the entry screen for this presentation node.
	RootViewIcon string
	NodeType     PresentationNodeType
	// Scope indicates the scope of this presentation node's state.
	Scope Scope
	// ObjectiveHash is the hash of a related ObjectiveEntity.
	ObjectiveHash uint32
	// CompletionRecordHash is the hash of a related RecordEntity.
	CompletionRecordHash uint32
	// Children are presentation nodes contained in this presentation node.
	Children PresentationNodeChildrenBlock
	// DisplayStyle is a hint for how to display this presentation node when it's shown in a list.
	DisplayStyle PresentationDisplayStyle
	// ScreenStyle is a hint for how to display this presentation node when it's shown in its own detail screen.
	ScreenStyle PresentationScreenStyle
	// Requirements are the requirements for being able to interact with this presentation node.
	Requirements PresentationNodeRequirementsBlock
	// DisableChildSubscreenNavigation determines if this presentation node's children cannot be inspected.
	DisableChildSubscreenNavigation bool
	MaxCategoryRecordScore          int32
	PresentationNodeType            PresentationNodeType
	TraitIds                        []string
	TraitHashes                     []uint32
	// ParentNodeHashes are the hashes of presentation nodes that this presentation node as a child.
	ParentNodeHashes []uint32
	EntityMetadata
}

// CollectibleEntity is an entity in the Destiny.Definitions.Collectibles.DestinyCollectibleDefinition contract.
type CollectibleEntity struct {
	DisplayProperties DisplayProperties
	// Scope indicates if this collectible is per-character or account-wide.
	Scope Scope
	// SourceString is a hint about how to acquire the item.
	SourceString string
	// SourceHash is supposed to be the hash of a given source. This is meant to be used to group collectibles by source.
	SourceHash uint32
	// ItemHash is the hash for a related InventoryItemEntity.
	ItemHash             uint32
	AcquisitionInfo      CollectibleAcquisitionBlock
	StateInfo            CollectibleStateBlock
	PresentationInfo     PresentationChildBlock
	PresentationNodeType PresentationNodeType
	TraitIds             []string
	TraitHashes          []uint32
	// ParentNodeHashes are the hashes of presentation nodes that this collectible as a child.
	ParentNodeHashes []uint32
	EntityMetadata
}

// MaterialRequirementSetEntity is an entity in the Destiny.Definitions.DestinyMaterialRequirementSetDefinition contract.
// Represents a set of material requirements which need to be owned or need to be consumed to perform an action.
type MaterialRequirementSetEntity struct {
	// Materials are a list of all required materials.
	Materials []MaterialRequirement
	EntityMetadata
}

// RecordEntity is an entity in the Destiny.Definitions.Records.DestinyRecordDefinition contract.
type RecordEntity struct {
	DisplayProperties DisplayProperties
	// Scope indicates if this record's state is per-character or account-wide.
	Scope            Scope
	PresentationInfo PresentationChildBlock
	// LoreHash is the hash of a related LoreEntity.
	LoreHash uint32
	// ObjectiveHashes are the hashes of all related ObjectiveEntity structs.
	ObjectiveHashes  []uint32
	RecordValueStyle RecordValueStyle
	ForTitleGilding  bool
	TitleInfo        RecordTitleBlock
	CompletionInfo   RecordCompletionBlock
	StateInfo        RecordStateBlock
	Requirements     PresentationNodeRequirementsBlock
	ExpirationInfo   RecordExpirationBlock
	// IntervalInfo is the set of interval objectives for this record.
	IntervalInfo RecordIntervalBlock
	// RewardItems are the list of items earned for achieving this record.
	RewardItems          []ItemQuantity
	PresentationNodeType PresentationNodeType
	TraitIds             []string
	TraitHashes          []uint32
	// ParentNodeHashes are the hashes of presentation nodes that this record as a child.
	ParentNodeHashes []uint32
	EntityMetadata
}

// GenderEntity is an entity in the Destiny.Definitions.DestinyGenderDefinition contract.
// This describes the current genders in Destiny 2.
type GenderEntity struct {
	// GenderType is one of the currently defined genders.
	GenderType        Gender
	DisplayProperties DisplayProperties
	EntityMetadata
}

// LoreEntity is an entity in the Destiny.Definitions.Lore.DestinyLoreDefinition contract.
// This represents in-game lore which are meant as narrative enhancements of the game experience.
type LoreEntity struct {
	DisplayProperties DisplayProperties
	Subtitle          string
	EntityMetadata
}

// MetricEntity is an entity in the Destiny.Definitions.Metrics.DestinyMetricDefinition contract.
type MetricEntity struct {
	DisplayProperties DisplayProperties
	// TrackingObjectiveHash is the hash of a related ObjectiveEntity.
	TrackingObjectiveHash uint32
	LowerValueIsBetter    bool
	PresentationNodeType  PresentationNodeType
	TraitIds              []string
	TraitHashes           []uint32
	// ParentNodeHashes are the hashes of presentation nodes that this metric as a child.
	ParentNodeHashes []uint32
	EntityMetadata
}

// EnergyTypeEntity is an entity in the Destiny.Definitions.EnergyTypes.DestinyEnergyTypeDefinition contract.
// Represents type of energy that can be used for costs and payments related to Armor 2.0.
type EnergyTypeEntity struct {
	DisplayProperties DisplayProperties
	// TransparentIconPath is a variant of the icon that is transparent and colorless.
	TransparentIconPath string
	// ShowIcon determines if the energy type's icon should be shown.
	ShowIcon bool
	// EnumValue is the energy type.
	EnumValue EnergyType
	// CapacityStatHash is the hash for a related InvestmentStatEntity for the capacity.
	CapacityStatHash uint32
	// CostStatHash is the hash for a related InvestmentStatEntity for the cost.
	CostStatHash uint32
	EntityMetadata
}

// PlugSetEntity is an entity in the Destiny.Definitions.DestinyPlugSetDefinition contract.
// This describes sets of reusable plugs that can be shared across places they are used.
type PlugSetEntity struct {
	DisplayProperties DisplayProperties
	// ReusuablePlugItems are the list of pre-determined plugs that be inserted in this socket.
	ReusablePlugItems []ItemSocketEntryPlugItemRandomized
	// IsFakePlugSet determines if this plug set if for debugging.
	IsFakePlugSet bool
	EntityMetadata
}

// TalentGridEntity is an entity in the Destiny.Definitions.DestinyTalentGridDefinition contract.
// Talent grids are used to describe all talents/perks granted by subclasses and builds.
type TalentGridEntity struct {
	// MaxGridLevel is the maximum of the talent grid. At the maximum level, any nodes can be activity.
	MaxGridLevel int32
	// GridLevelPerColumn was used to show talent nodes in column by their progression level. Deprecated.
	GridLevelPerColumn int32
	// ProgressionHash is the hash of a related ProgressionEntity.
	ProgressionHash uint32
	// Nodes are a list of talents on the given grid.
	Nodes []TalentNode
	// ExclusiveSets are talent node sets in which only a single node can be activated at a time.
	ExclusiveSets []TalentNodeExclusiveSet
	// IndependentNodeIndexes are the indexes of nodes which are not in exclusive sets.
	IndependentNodeIndexes []int32
	// Groups are the exclusive groups for this talent node.
	Groups map[uint32]TalentExclusiveGroup
	// NodeCategories are talent nodes grouped by similar purpose with localized titles.
	NodeCategories []TalentNodeCategory
	EntityMetadata
}

// DamageTypeEntity is an entity in the Destiny.Definitions.DestinyDamageTypeDefinition contract.
// All damage types that are possible in the game are defined here along with localized info.
type DamageTypeEntity struct {
	DisplayProperties DisplayProperties
	// TransparentIconPath is a variant of the icon that is transparent and colorless
	TransparentIconPath string
	// ShowIcon determines if the UI shows this damage type's icon.
	ShowIcon bool
	// EnumValue is the this damage type's defined value.
	EnumValue DamageType
	EntityMetadata
}

// ItemCategoryEntity is an entity in the Destiny.Definitions.DestinyItemCategoryDefinition contract.
// This is a grouping of items based on properties such as inventory bucket, type name and breaker type.
type ItemCategoryEntity struct {
	DisplayProperties DisplayProperties
	// Visible determines if this item category should be visible in a given UI.
	Visible bool
	// Deprecated determines if this item category is deprecated.
	Deprecated bool
	// ShortTitle is a shorted version of the title of this item category.
	ShortTitle string
	// ItemTypeRegex is a regular expression used against the item type to determine membership in this item category.
	ItemTypeRegex string
	// GrantDestinyBreakerType is the breaker type granted by items in this category.
	GrantDestinyBreakerType BreakerTypeEnum
	// PlugCategoryIdentifier is the identifier of the plug in this item category.
	PlugCategoryIdentifier string
	// ItemTypeRegexNot is a regular expression used against the item type to determine non-membership in this item category.
	ItemTypeRegexNot string
	// OriginBucketIdentifier is the name of the bucket for this item category.
	OriginBucketIdentifier string
	// GrantDestinyItemType is the item type for items in this category.
	GrantDestinyItemType ItemType
	// GrantDestinySubType is the item subtype for items in this category.
	GrantDestinySubType ItemSubType
	// GrantDestinyClass is the class restriction for this item category.
	GrantDestinyClass Class
	// Trait is the identifier of the item trait found in this category.
	TraitId string
	// GroupedCategoryHashes are the hashes of children categories contained within this item category.
	GroupedCategoryHashes []uint32
	// ParentCategoryHashes are the hashes of parent categories which contain this item category.
	ParentCategoryHashes []uint32
	// GroupCategoryOnly determines if this item category is only used for grouping other item categories.
	GroupCategoryOnly bool
	EntityMetadata
}

// BreakerTypeEntity is an entity in the Destiny.Definitions.BreakerTypes.DestinyBreakerTypeDefinition contract.
// This basically describes anti-champion modes of an item or plug.
type BreakerTypeEntity struct {
	DisplayProperties DisplayProperties
	// EnumValue refers to the breaker classification for this breaker type.
	EnumValue BreakerTypeEnum
	EntityMetadata
}

// SeasonEntity is an entity in the Destiny.Definitions.Seasons.DestinySeasonDefinition contract.
// A season is a range of a few months where the game highlights certai challenges and events.
type SeasonEntity struct {
	DisplayProperties   DisplayProperties
	BackgroundImagePath string
	SeasonNumber        int32
	StartDate           time.Time
	EndDate             time.Time
	// SeasonPassHash is the hash of a related SeasonPassEntity.
	SeasonPassHash uint32
	// SeasonPassProgressionHash is the hash of a related ProgressionEntity.
	SeasonPassProgressionHash uint32
	// ArtifactItemHash is the hash of a related InventoryItemEntity.
	ArtifactItemHash uint32
	// SealPresentationNodeHash is the hash of a related PresentationNodeEntity for this season's seal.
	SealPresentationNodeHash uint32
	// SeasonalChallengesPresentationNodeHash is the hash of a related PresentationNodeEntity for this season's challenges.
	SeasonalChallengesPresentationNodeHash uint32
	// Preview is an optional preview of promotional text, images and links for this season.
	Preview SeasonPreview
	EntityMetadata
}

// SeasonPassEntity is an entity in the Destiny.Definitions.Seasons.DestinySeasonPassDefinition contract.
type SeasonPassEntity struct {
	DisplayProperties DisplayProperties
	// RewardProgressionHash is the hash for a related ProgressionEntity for this season pass' reward.
	RewardProgressionHash uint32
	// PrestigeProgressionHash is the hash for a related ProgressionEntity for this season pass' rewards after reaching max rank.
	PrestigeProgressionHash uint32
	EntityMetadata
}

// ChecklistEntity is an entity in the Destiny.Definitions.Checklists.DestinyChecklistDefinition contract.
// Checklists are loose sets of things to do/things that have been done that can be tracked.
type ChecklistEntity struct {
	DisplayProperties DisplayProperties
	// ViewActionString is a localized string prompt for viewing this checklist.
	ViewActionString string
	// Scope determines if this checklist if per-character or account-wide.
	Scope Scope
	// Entries are the individual checklist items.
	Entries []ChecklistEntry
	EntityMetadata
}

// RaceEntity is an entity in the Destiny.Definitions.DestinyRaceDefinition contract.
// Races are the various playable species in Destiny 2.
type RaceEntity struct {
	DisplayProperties DisplayProperties
	// RaceType is the classification of this race.
	RaceType Race
	// GenderedRaceNames are the names of this race by gender name.
	GenderedRaceNames map[GenderName]string
	// GenderedRaceNames are the name of this race by the hash of a GenderEntity.
	GenderedRaceNamesByGenderHash map[uint32]string
	EntityMetadata
}

// ClassEntity is an entity in the Destiny.Definitions.DestinyClassDefinition contract.
// This defines a playable character class in Destiny 2.
type ClassEntity struct {
	// ClassType is the classification of this character class.
	ClassType         Class
	DisplayProperties DisplayProperties
	// GenderedClassNames are the names of this class by gender name.
	GenderedClassNames map[GenderName]string
	// GenderedClassNamesByGenderHash are the names of this class by the hash of a GenderEntity.
	GenderedClassNamesByGenderHash map[uint32]string
	// MentorVendorHash is the hash of a related VendorEntity.
	MentorVendorHash uint32
	EntityMetadata
}

// MilestoneEntity is an entity in the Destiny.Definitions.Milestones.DestinyMilestoneDefinition contract.
// Milestones are an in-game concept that hints at what the player can do next.
type MilestoneEntity struct {
	DisplayProperties DisplayProperties
	// DisplayPreference is a hint to a given UI about what DisplayProperties to show when viewing this milestone's live data.
	DisplayPreference MilestoneDisplayPreference
	// Image is a custom image for this milestone.
	Image string
	// MilestoneType is the classification of this milestone.
	MilestoneType MilestoneType
	// Recruitable determines if this milestone is integrated with the API's recruiting feature.
	Recruitable bool
	// FriendlyName is the friendly identifier for this milestone.
	FriendlyName string
	// ShowInMilestones determines if this milestone is shown in the user's personal milestones list.
	ShowInMilestones bool
	// ExplorePrioritizesActivityImage determines if the "Explore Destiny" page prioritizes using the activity image for this milestone.
	ExplorePrioritizesActivityImage bool
	// HashPredictableDates determines if the start and end date of this milestone can be predicted.
	HasPredictableDates bool
	// Quests are the set of all quests related to this milestone.
	Quests map[uint32]MilestoneQuest
	// Rewards are the categories which individual rewards are placed in by the hash of the MilestoneRewordCategory.
	Rewards map[uint32]MilestoneRewardCategory
	// VendorsDisplayTitle is the localized header for vendors related to this milestone.
	VendorsDisplayTitle string
	// Vendors are the vendors which have rewards for this milestone.
	Vendors []MilestoneVendor
	// Values are arbitrary values associated with a given identifier in this milestone.
	Values map[string]MilestoneValue
	// IsInGameMilestone determines if this milestone is an objective that can interacted with in the game.
	IsInGameMilestone bool
	// Activities are the challenge and modifiers related to this milestone.
	Activities   []MilestoneChallengeActivity
	DefaultOrder int32
	EntityMetadata
}

// UnlockEntity is an entity in the Destiny.Definitions.DestinyUnlockDefinition contract.
// Unlock flags are used by the game server for state checks, progress storage, etc.
type UnlockEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// ReportReasonCategoryEntity is an entity in the Destiny.Definitions.Reporting.DestinyReportReasonCategoryDefinition contract.
// When reporting another player for TOS violation a category and reason must be chosen.
type ReportReasonCategoryEntity struct {
	DisplayProperties DisplayProperties
	// Reasons are the specific reasons for the report under this category.
	Reasons map[uint32]ReportReason
	EntityMetadata
}
