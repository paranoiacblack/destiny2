package destiny2

import "time"

// InventoryItemEntity is an entity in the Destiny.Definitions.DestinyInventoryItemDefinition contract.
type InventoryItemEntity struct {
	DisplayProperties                 DisplayProperties
	TooltipNotifications              []ItemTooltipNotification
	CollectibleHash                   uint32
	IconWatermark                     string
	IconWatermarkShelved              string
	SecondaryIcon                     string
	SecondaryOverlay                  string
	BackgroundColor                   Color
	Screenshot                        string
	ItemTypeDisplayName               string
	FlavorText                        string
	UiItemDisplayStyle                string
	ItemTypeAndTierDisplayName        string
	DisplaySource                     string
	TooltipStyle                      string
	Action                            ItemActionBlock
	Inventory                         ItemInventoryBlock
	SetData                           ItemSetBlock
	Stats                             ItemStatBlock
	EmblemObjectiveHash               uint32
	EquippingBlock                    EquippingBlock
	TranslationBlock                  ItemTranslationBlock
	Preview                           ItemPreviewBlock
	Quality                           ItemQualityBlock
	Value                             ItemValueBlock
	SourceData                        ItemSourceBlock
	Objectives                        ItemObjectiveBlock
	Metrics                           ItemMetricBlock
	Plug                              ItemPlug
	Gearset                           ItemGearsetBlock
	Sack                              ItemSackBlock
	Sockets                           ItemSocketBlock
	Summary                           ItemSummaryBlock
	TalentGrid                        ItemTalentGridBlock
	InvestmentStats                   []ItemInvestmentStat
	Perks                             []ItemPerkEntry
	LoreHash                          uint32
	SummaryItemHash                   uint32
	Animations                        []AnimationReference
	AllowActions                      bool
	Links                             []HyperlinkReference
	DoesPostmasterPullHaveSideEffects bool
	NonTransferrable                  bool
	ItemCategoryHashes                []uint32
	SpecialItemType                   SpecialItemType
	ItemType                          ItemType
	ItemSubType                       ItemSubType
	BreakerType                       BreakerTypeEnum
	BreakerTypeHash                   uint32
	Equippable                        bool
	DamageTypeHashes                  []uint32
	DamageTypes                       []DamageType
	DefaultDamageType                 DamageType
	DefaultDamageTypeHash             uint32
	SeasonHash                        uint32
	IsWrapper                         bool
	TraitIds                          []string
	TraitHashes                       []uint32
	EntityMetadata
}

// ProgressionEntity is an entity in the Destiny.Definitions.DestinyProgressionDefinition contract.
type ProgressionEntity struct {
	DisplayProperties DisplayProperties
	Scope             ProgressionScope
	RepeatLastStep    bool
	Source            string
	Steps             []ProgressionStep
	Visible           bool
	FactionHash       uint32
	Color             Color
	RankIcon          string
	RewardItems       []ProgressionRewardItemQuantity
	EntityMetadata
}

// InventoryBucketEntity is an entity in the Destiny.Definitions.DestinyInventoryBucketDefinition contract.
type InventoryBucketEntity struct {
	DisplayProperties      DisplayProperties
	Scope                  BucketScope
	Category               BucketCategory
	BucketOrder            int32
	ItemCount              int32
	Location               ItemLocation
	HasTransferDestination bool
	Enabled                bool
	Fifo                   bool
	EntityMetadata
}

// ItemTierTypeEntity is an entity in the Destiny.Definitions.Items.DestinyItemTierTypeDefinition contract.
type ItemTierTypeEntity struct {
	DisplayProperties DisplayProperties
	InfusionProcess   ItemTierTypeInfusionBlock
	EntityMetadata
}

// StatEntity is an entity in the Destiny.Definitions.DestinyStatDefinition contract.
type StatEntity struct {
	DisplayProperties DisplayProperties
	AggregationType   StatAggregationType
	HasComputedBlock  bool
	StatCategory      StatCategory
	EntityMetadata
}

// StatGroupEntity is an entity in the Destiny.Definitions.DestinyStatGroupDefinition contract.
type StatGroupEntity struct {
	MaximumValue int32
	UiPosition   int32
	ScaledStats  []StatDisplay
	Overrides    map[uint32]StatOverride
	EntityMetadata
}

// EquipmentSlotEntity is an entity in the Destiny.Definitions.DestinyEquipmentSlotDefinition contract.
type EquipmentSlotEntity struct {
	EquipmentCategoryHash uint32
	BucketTypeHash        uint32
	ApplyCustomArtDyes    bool
	ArtDyeChannels        []ArtDyeReference
	EntityMetadata
}

// VendorEntity is an entity in the Destiny.Definitions.DestinyVendorDefinition contract.
type VendorEntity struct {
	DisplayProperties           DisplayProperties
	VendorProgressionType       VendorProgressionType
	BuyString                   string
	SellString                  string
	DisplayItemHash             uint32
	InhibitBuying               bool
	InhibitSelling              bool
	FactionHash                 uint32
	ResetIntervalMinutes        int32
	ResetOffsetMinutes          int32
	FailureStrings              string
	UnlockRanges                []DateRange
	VendorIdentifier            string
	VendorPortrait              string
	VendorBanner                string
	Enabled                     bool
	Visible                     bool
	VendorSubcategoryIdentifier string
	ConsolidateCategories       bool
	Actions                     []VendorAction
	Categories                  []VendorCategoryEntry
	OriginalCategories          []VendorCategoryEntry
	DisplayCategories           []DisplayCategory
	Interactions                []VendorInteraction
	InventoryFlyouts            []VendorInventoryFlyout
	ItemList                    []VendorItem
	Services                    []VendorService
	AcceptedItems               []VendorAcceptedItem
	ReturnWithVendorRequest     bool
	Locations                   []VendorLocation
	Groups                      []VendorGroupReference
	IgnoreSaleItemHashes        []uint32
	EntityMetadata
}

// SocketTypeEntity is an entity in the Destiny.Definitions.Sockets.DestinySocketTypeDefinition contract.
type SocketTypeEntity struct {
	DisplayProperties               DisplayProperties
	InsertAction                    InsertPlugAction
	PlugWhitelist                   []PlugWhitelistEntry
	SocketCategoryHash              uint32
	Visibility                      SocketTypeVisibility
	AlwaysRandomizeSockets          bool
	IsPreviewEnabled                bool
	HideDuplicateReusablePlugs      bool
	OverridesUiAppearance           bool
	AvoidDuplicatesOnInitialization bool
	CurrencyScalars                 []SocketTypeScalarMaterialRequirementEntry
	EntityMetadata
}

// SocketCategoryEntity is an entity in the Destiny.Definitions.Sockets.DestinySocketCategoryDefinition contract.
type SocketCategoryEntity struct {
	DisplayProperties DisplayProperties
	UiCategoryStyle   uint32
	CategoryStyle     SocketCategoryStyle
	EntityMetadata
}

// DestinationEntity is an entity in the Destiny.Definitions.DestinyDestinationDefinition contract.
type DestinationEntity struct {
	DisplayProperties           DisplayProperties
	PlaceHash                   uint32
	DefaultFreeroamActivityHash uint32
	ActivityGraphEntries        []ActivityGraphListEntry
	BubbleSettings              []DestinationBubbleSetting
	Bubbles                     []Bubble
	EntityMetadata
}

// ActivityGraphEntity is an entity in the Destiny.Definitions.Director.DestinyActivityGraphDefinition contract.
type ActivityGraphEntity struct {
	Nodes               []ActivityGraphNode
	ArtElements         []ActivityGraphArtElement
	Connections         []ActivityGraphConnection
	DisplayObjectives   []ActivityGraphDisplayObjective
	DisplayProgressions []ActivityGraphDisplayProgression
	LinkedGraphs        []LinkedGraph
	EntityMetadata
}

// ActivityEntity is an entity in the Destiny.Definitions.DestinyActivityDefinition contract.
type ActivityEntity struct {
	DisplayProperties                DisplayProperties
	OriginalDisplayProperties        DisplayProperties
	SelectionScreenDisplayProperties DisplayProperties
	ReleaseIcon                      string
	ReleaseTime                      int32
	ActivityLightLevel               int32
	DestinationHash                  uint32
	PlaceHash                        uint32
	ActivityTypeHash                 uint32
	Tier                             ActivityDifficultyTier
	PgcrImage                        string
	Rewards                          []ActivityReward
	Modifiers                        []ActivityModifierReference
	IsPlaylist                       bool
	Challenges                       []ActivityChallenge
	OptionalUnlockStrings            []ActivityUnlockString
	PlaylistItems                    []ActivityPlaylistItem
	ActivityGraphList                []ActivityGraphListEntry
	Matchmaking                      ActivityMatchmakingBlock
	GuidedGame                       ActivityGuidedBlock
	DirectActivityModeHash           uint32
	DirectActivityModeType           ActivityModeType
	Loadouts                         []ActivityLoadoutRequirementSet
	ActivityModeHashes               []uint32
	ActivityModeTypes                []ActivityModeType
	IsPvP                            bool
	InsertionPoints                  []ActivityInsertionPoint
	ActivityLocationMappings         []EnvironmentLocationMapping
	EntityMetadata
}

// ActivityModifierEntity is an entity in the Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition contract.
type ActivityModifierEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// ObjectiveEntity is an entity in the Destiny.Definitions.DestinyObjectiveDefinition contract.
type ObjectiveEntity struct {
	DisplayProperties             DisplayProperties
	CompletionValue               int32
	Scope                         GatingScope
	LocationHash                  uint32
	AllowNegativeValue            bool
	AllowValueChangeWhenCompleted bool
	IsCountingDownward            bool
	ValueStyle                    UnlockValueUIStyle
	ProgressDescription           string
	Perks                         ObjectivePerkEntry
	Stats                         ObjectiveStatEntry
	MinimumVisibilityThreshold    int32
	AllowOvercompletion           bool
	ShowValueOnComplete           bool
	CompletedValueStyle           UnlockValueUIStyle
	InProgressValueStyle          UnlockValueUIStyle
	EntityMetadata
}

// SandboxPerkEntity is an entity in the Destiny.Definitions.DestinySandboxPerkDefinition contract.
type SandboxPerkEntity struct {
	DisplayProperties DisplayProperties
	PerkIdentifier    string
	IsDisplayable     bool
	DamageType        DamageType
	DamageTypeHash    uint32
	PerkGroups        TalentNodeStep
	EntityMetadata
}

// LocationEntity is an entity in the Destiny.Definitions.DestinyLocationDefinition contract.
type LocationEntity struct {
	VendorHash       uint32
	LocationReleases []LocationRelease
	EntityMetadata
}

// ActivityModeEntity is an entity in the Destiny.Definitions.DestinyActivityModeDefinition contract.
type ActivityModeEntity struct {
	DisplayProperties    DisplayProperties
	PgcrImage            string
	ModeType             ActivityModeType
	ActivityModeCategory ActivityModeCategory
	IsTeamBased          bool
	IsAggregateMode      bool
	ParentHashes         []uint32
	FriendlyName         string
	ActivityModeMappings map[uint32]ActivityModeType
	Display              bool
	Order                int32
	EntityMetadata
}

// PlaceEntity is an entity in the Destiny.Definitions.DestinyPlaceDefinition contract.
type PlaceEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// ActivityTypeEntity is an entity in the Destiny.Definitions.DestinyActivityTypeDefinition contract.
type ActivityTypeEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// VendorGroupEntity is an entity in the Destiny.Definitions.DestinyVendorGroupDefinition contract.
type VendorGroupEntity struct {
	Order        int32
	CategoryName string
	EntityMetadata
}

// FactionEntity is an entity in the Destiny.Definitions.DestinyFactionDefinition contract.
type FactionEntity struct {
	DisplayProperties DisplayProperties
	ProgressionHash   uint32
	TokenValues       map[uint32]uint32
	RewardItemHash    uint32
	RewardVendorHash  uint32
	Vendors           []FactionVendor
	EntityMetadata
}

// ArtifactEntity is an entity in the Destiny.Definitions.Artifacts.DestinyArtifactDefinition contract.
type ArtifactEntity struct {
	DisplayProperties DisplayProperties
	TranslationBlock  ItemTranslationBlock
	Tiers             []ArtifactTier
	EntityMetadata
}

// PowerCapEntity is an entity in the Destiny.Definitions.PowerCaps.DestinyPowerCapDefinition contract.
type PowerCapEntity struct {
	PowerCap int32
	EntityMetadata
}

// ProgressionLevelRequirementEntity is an entity in the Destiny.Definitions.Progression.DestinyProgressionLevelRequirementDefinition contract.
type ProgressionLevelRequirementEntity struct {
	RequirementCurve []InterpolationPointFloat
	ProgressionHash  uint32
	EntityMetadata
}

// RewardSourceEntity is an entity in the Destiny.Definitions.DestinyRewardSourceDefinition contract.
type RewardSourceEntity struct {
	DisplayProperties DisplayProperties
	Category          RewardSourceCategory
	EntityMetadata
}

// TraitEntity is an entity in the Destiny.Definitions.Traits.DestinyTraitDefinition contract.
type TraitEntity struct {
	DisplayProperties DisplayProperties
	TraitCategoryId   string
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
type PresentationNodeEntity struct {
	DisplayProperties               DisplayProperties
	OriginalIcon                    string
	RootViewIcon                    string
	NodeType                        PresentationNodeType
	Scope                           Scope
	ObjectiveHash                   uint32
	CompletionRecordHash            uint32
	Children                        PresentationNodeChildrenBlock
	DisplayStyle                    PresentationDisplayStyle
	ScreenStyle                     PresentationScreenStyle
	Requirements                    PresentationNodeRequirementsBlock
	DisableChildSubscreenNavigation bool
	MaxCategoryRecordScore          int32
	PresentationNodeType            PresentationNodeType
	TraitIds                        []string
	TraitHashes                     []uint32
	ParentNodeHashes                []uint32
	EntityMetadata
}

// CollectibleEntity is an entity in the Destiny.Definitions.Collectibles.DestinyCollectibleDefinition contract.
type CollectibleEntity struct {
	DisplayProperties    DisplayProperties
	Scope                Scope
	SourceString         string
	SourceHash           uint32
	ItemHash             uint32
	AcquisitionInfo      CollectibleAcquisitionBlock
	StateInfo            CollectibleStateBlock
	PresentationInfo     PresentationChildBlock
	PresentationNodeType PresentationNodeType
	TraitIds             []string
	TraitHashes          []uint32
	EntityMetadata
}

// MaterialRequirementSetEntity is an entity in the Destiny.Definitions.DestinyMaterialRequirementSetDefinition contract.
type MaterialRequirementSetEntity struct {
	Materials []MaterialRequirement
	EntityMetadata
}

// RecordEntity is an entity in the Destiny.Definitions.Records.DestinyRecordDefinition contract.
type RecordEntity struct {
	DisplayProperties    DisplayProperties
	Scope                Scope
	PresentationInfo     PresentationChildBlock
	LoreHash             uint32
	ObjectiveHashes      []uint32
	RecordValueStyle     RecordValueStyle
	ForTitleGilding      bool
	TitleInfo            RecordTitleBlock
	CompletionInfo       RecordCompletionBlock
	StateInfo            RecordStateBlock
	Requirements         PresentationNodeRequirementsBlock
	ExpirationInfo       RecordExpirationBlock
	IntervalInfo         RecordIntervalBlock
	RewardItems          []ItemQuantity
	PresentationNodeType PresentationNodeType
	TraitIds             []string
	TraitHashes          []uint32
	ParentNodeHashes     []uint32
	EntityMetadata
}

// GenderEntity is an entity in the Destiny.Definitions.DestinyGenderDefinition contract.
type GenderEntity struct {
	GenderType        Gender
	DisplayProperties DisplayProperties
	EntityMetadata
}

// LoreEntity is an entity in the Destiny.Definitions.Lore.DestinyLoreDefinition contract.
type LoreEntity struct {
	DisplayProperties DisplayProperties
	Subtitle          string
	EntityMetadata
}

// MetricEntity is an entity in the Destiny.Definitions.Metrics.DestinyMetricDefinition contract.
type MetricEntity struct {
	DisplayProperties     DisplayProperties
	TrackingObjectiveHash uint32
	LowerValueIsBetter    bool
	PresentationNodeType  PresentationNodeType
	TraitIds              []string
	TraitHashes           []uint32
	ParentNodeHashes      []uint32
	EntityMetadata
}

// EnergyTypeEntity is an entity in the Destiny.Definitions.EnergyTypes.DestinyEnergyTypeDefinition contract.
type EnergyTypeEntity struct {
	DisplayProperties   DisplayProperties
	TransparentIconPath string
	ShowIcon            bool
	EnumValue           EnergyType
	CapacityStatHash    uint32
	CostStatHash        uint32
	EntityMetadata
}

// PlugSetEntity is an entity in the Destiny.Definitions.DestinyPlugSetDefinition contract.
type PlugSetEntity struct {
	DisplayProperties DisplayProperties
	ReusablePlugItems []ItemSocketEntryPlugItemRandomized
	IsFakePlugSet     bool
	EntityMetadata
}

// TalentGridEntity is an entity in the Destiny.Definitions.DestinyTalentGridDefinition contract.
type TalentGridEntity struct {
	MaxGridLevel           int32
	GridLevelPerColumn     int32
	ProgressionHash        uint32
	Nodes                  []TalentNode
	ExclusiveSets          []TalentNodeExclusiveSet
	IndependentNodeIndexes []int32
	Groups                 TalentExclusiveGroup
	NodeCategories         []TalentNodeCategory
	EntityMetadata
}

// DamageTypeEntity is an entity in the Destiny.Definitions.DestinyDamageTypeDefinition contract.
type DamageTypeEntity struct {
	DisplayProperties   DisplayProperties
	TransparentIconPath string
	ShowIcon            bool
	EnumValue           DamageType
	EntityMetadata
}

// ItemCategoryEntity is an entity in the Destiny.Definitions.DestinyItemCategoryDefinition contract.
type ItemCategoryEntity struct {
	DisplayProperties       DisplayProperties
	Visible                 bool
	Deprecated              bool
	ShortTitle              string
	ItemTypeRegex           string
	GrantDestinyBreakerType BreakerTypeEnum
	PlugCategoryIdentifier  string
	ItemTypeRegexNot        string
	OriginBucketIdentifier  string
	GrantDestinyItemType    ItemType
	GrantDestinySubType     ItemSubType
	GrantDestinyClass       Class
	TraitId                 string
	GroupedCategoryHashes   []uint32
	ParentCategoryHashes    []uint32
	GroupCategoryOnly       bool
	EntityMetadata
}

// BreakerTypeEntity is an entity in the Destiny.Definitions.BreakerTypes.DestinyBreakerTypeDefinition contract.
type BreakerTypeEntity struct {
	DisplayProperties DisplayProperties
	EnumValue         BreakerTypeEnum
	EntityMetadata
}

// SeasonEntity is an entity in the Destiny.Definitions.Seasons.DestinySeasonDefinition contract.
type SeasonEntity struct {
	DisplayProperties                      DisplayProperties
	BackgroundImagePath                    string
	SeasonNumber                           int32
	StartDate                              time.Time
	EndDate                                time.Time
	SeasonPassHash                         uint32
	SeasonPassProgressionHash              uint32
	ArtifactItemHash                       uint32
	SealPresentationNodeHash               uint32
	SeasonalChallengesPresentationNodeHash uint32
	Preview                                SeasonPreview
	EntityMetadata
}

// SeasonPassEntity is an entity in the Destiny.Definitions.Seasons.DestinySeasonPassDefinition contract.
type SeasonPassEntity struct {
	DisplayProperties       DisplayProperties
	RewardProgressionHash   uint32
	PrestigeProgressionHash uint32
	EntityMetadata
}

// ChecklistEntity is an entity in the Destiny.Definitions.Checklists.DestinyChecklistDefinition contract.
type ChecklistEntity struct {
	DisplayProperties DisplayProperties
	ViewActionString  string
	Scope             Scope
	Entries           []ChecklistEntry
	EntityMetadata
}

// RaceEntity is an entity in the Destiny.Definitions.DestinyRaceDefinition contract.
type RaceEntity struct {
	DisplayProperties             DisplayProperties
	RaceType                      Race
	GenderedRaceNames             map[GenderName]string
	GenderedRaceNamesByGenderHash map[uint32]string
	EntityMetadata
}

// ClassEntity is an entity in the Destiny.Definitions.DestinyClassDefinition contract.
type ClassEntity struct {
	ClassType                      Class
	DisplayProperties              DisplayProperties
	GenderedClassNames             map[GenderName]string
	GenderedClassNamesByGenderHash map[uint32]string
	MentorVendorHash               uint32
	EntityMetadata
}

// MilestoneEntity is an entity in the Destiny.Definitions.Milestones.DestinyMilestoneDefinition contract.
type MilestoneEntity struct {
	DisplayProperties               DisplayProperties
	DisplayPreference               MilestoneDisplayPreference
	Imagine                         string
	MilestoneType                   MilestoneType
	Recruitable                     bool
	FriendlyName                    string
	ShowInMilestones                bool
	ExplorePrioritizesActivityImage bool
	HasPredictableDates             bool
	Quests                          map[uint32]MilestoneQuest
	Rewards                         map[uint32]MilestoneRewardCategory
	VendorsDisplayTitle             string
	Vendors                         []MilestoneVendor
	Values                          map[string]MilestoneValue
	IsInGameMilestone               bool
	Activities                      []MilestoneChallengeActivity
	DefaultOrder                    int32
	EntityMetadata
}

// UnlockEntity is an entity in the Destiny.Definitions.DestinyUnlockDefinition contract.
type UnlockEntity struct {
	DisplayProperties DisplayProperties
	EntityMetadata
}

// ReportReasonCategoryEntity is an entity in the Destiny.Definitions.Reporting.DestinyReportReasonCategoryDefinition contract.
type ReportReasonCategoryEntity struct {
	DisplayProperties DisplayProperties
	Reasons           map[uint32]ReportReason
	EntityMetadata
}
