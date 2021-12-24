package destiny2

import "time"

// InventoryItemEntity is an entity in the Destiny.Definitions.DestinyInventoryItemDefinition contract.
type InventoryItemEntity struct {
	DisplayProperties                 DisplayProperties
	TooltipNotifications              []ItemTooltipNotification
	CollectableHash                   uint32
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
	SpecialItemType                   int32
	ItemType                          int32
	ItemSubType                       int32
	BreakerType                       int32
	BreakerTypeHash                   uint32
	Equippable                        bool
	DamageTypeHashes                  []uint32
	DamageTypes                       []int32
	DefaultDamageType                 int32
	DefaultDamageTypeHash             uint32
	SeasonHash                        uint32
	IsWrapper                         bool
	TraitIds                          []string
	TraitHashes                       []uint32
	Hash                              uint32
	Index                             int32
	Redacted                          bool
}

// ProgressionEntity is an entity in the Destiny.Definitions.DestinyProgressionDefinition contract.
type ProgressionEntity struct {
	DisplayProperties DisplayProperties
	Scope             int32
	RepeatLastStep    bool
	Source            string
	Steps             []ProgressionStep
	Visible           bool
	FactionHash       uint32
	Color             Color
	RankIcon          string
	RewardItems       []ProgressionRewardItemQuantity
	Hash              uint32
	Index             int32
	Redacted          bool
}

type InventoryBucketEntity struct {
	DisplayProperties      DisplayProperties
	Scope                  int32
	Category               int32
	BucketOrder            int32
	ItemCount              int32
	Location               int32
	HasTransferDestination bool
	Enabled                bool
	Fifo                   bool
	Hash                   uint32
	Index                  int32
	Redacted               bool
}

type ItemTierTypeEntity struct {
	DisplayProperties DisplayProperties
	InfusionProcess   ItemTierTypeInfusionBlock
	Hash              uint32
	Index             int32
	Redacted          bool
}

type StatEntity struct {
	DisplayProperties DisplayProperties
	AggregationType   int32
	HasComputedBlock  bool
	StatCategory      int32
	Hash              uint32
	Index             int32
	Redacted          bool
}

type StatGroupEntity struct {
	MaximumValue int32
	UiPosition   int32
	ScaledStats  []StatDisplay
	Overrides    map[uint32]StatOverride
	Hash         uint32
	Index        int32
	Redacted     bool
}

type EquipmentSlotEntity struct {
	EquipmentCategoryHash uint32
	BucketTypeHash        uint32
	ApplyCustomArtDyes    bool
	ArtDyeChannels        []ArtDyeReference
	Hash                  uint32
	Index                 int32
	Redacted              bool
}

type VendorEntity struct {
	DisplayProperties           DisplayProperties
	VendorProgressionType       int32
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
	Hash                        uint32
	Index                       int32
	Redacted                    bool
}

type SocketTypeEntity struct {
	DisplayProperties               DisplayProperties
	InsertAction                    InsertPlugAction
	PlugWhitelist                   []PlugWhitelistEntry
	SocketCategoryHash              uint32
	Visibility                      int32
	AlwaysRandomizeSockets          bool
	IsPreviewEnabled                bool
	HideDuplicateReusablePlugs      bool
	OverridesUiAppearance           bool
	AvoidDuplicatesOnInitialization bool
	CurrencyScalars                 []SocketTypeScalarMaterialRequirementEntry
	Hash                            uint32
	Index                           int32
	Redacted                        bool
}

type SocketCategoryEntity struct {
	DisplayProperties DisplayProperties
	UiCategoryStyle   uint32
	CategoryStyle     int32
	Hash              uint32
	Index             int32
	Redacted          bool
}

type DestinationEntity struct {
	DisplayProperties           DisplayProperties
	PlaceHash                   uint32
	DefaultFreeroamActivityHash uint32
	ActivityGraphEntries        []ActivityGraphListEntry
	BubbleSettings              []DestinationBubbleSetting
	Bubbles                     []Bubble
	Hash                        uint32
	Index                       int32
	Redacted                    bool
}

type ActivityGraphEntity struct {
	Nodes               []ActivityGraphNode
	ArtElements         []ActivityGraphArtElement
	Connections         []ActivityGraphConnection
	DisplayObjectives   []ActivityGraphDisplayObjective
	DisplayProgressions []ActivityGraphDisplayProgression
	LinkedGraphs        []LinkedGraph
	Hash                uint32
	Index               int32
	Redacted            bool
}

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
	Tier                             int32
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
	Hash                             uint32
	Index                            int32
	Redacted                         bool
}

type ActivityModifierEntity struct {
	DisplayProperties DisplayProperties
	Hash              uint32
	Index             int32
	Redacted          bool
}

type ObjectiveEntity struct {
	DisplayProperties             DisplayProperties
	CompletionValue               int32
	Scope                         int32
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
	CompletedValueStyle           int32
	InProgressValueStyle          int32
	Hash                          uint32
	Index                         int32
	Redacted                      bool
}

type SandboxPerkEntity struct {
	DisplayProperties DisplayProperties
	PerkIdentifier    string
	IsDisplayable     bool
	DamageType        DamageType
	DamageTypeHash    uint32
	PerkGroups        TalentNodeStep
	Hash              uint32
	Index             int32
	Redacted          bool
}

type LocationEntity struct {
	VendorHash       uint32
	LocationReleases []LocationRelease
	Hash             uint32
	Index            int32
	Redacted         bool
}

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
	Hash                 uint32
	Index                int32
	Redacted             bool
}

type PlaceEntity struct {
	DisplayProperties DisplayProperties
	Hash              uint32
	Index             int32
	Redacted          bool
}

type ActivityTypeEntity struct {
	DisplayProperties DisplayProperties
	Hash              uint32
	Index             int32
	Redacted          bool
}

type VendorGroupEntity struct {
	Order        int32
	CategoryName string
	Hash         uint32
	Index        int32
	Redacted     bool
}

type FactionEntity struct {
	DisplayProperties DisplayProperties
	ProgressionHash   uint32
	TokenValues       map[uint32]uint32
	RewardItemHash    uint32
	RewardVendorHash  uint32
	Vendors           []FactionVendor
	Hash              uint32
	Index             int32
	Redacted          bool
}

type ArtifactEntity struct {
	DisplayProperties DisplayProperties
	TranslationBlock  ItemTranslationBlock
	Tiers             []ArtifactTier
	Hash              uint32
	Index             int32
	Redacted          bool
}

type PowerCapEntity struct {
	PowerCap int32
	Hash     uint32
	Index    int32
	Redacted bool
}

type ProgressionLevelRequirementEntity struct {
	RequirementCurve []InterpolationPointFloat
	ProgressionHash  uint32
	Hash             uint32
	Index            int32
	Redacted         bool
}

type RewardSourceEntity struct {
	DisplayProperties DisplayProperties
	Category          RewardSourceCategory
	Hash              uint32
	Index             int32
	Redacted          bool
}

type TraitEntity struct {
	DisplayProperties DisplayProperties
	TraitCategoryId   string
	TraitCategoryHash uint32
	Hash              uint32
	Index             int32
	Redacted          bool
}

type TraitCategoryEntity struct {
	TraitCategoryId string
	TraitHashes     []uint32
	TraitIds        []string
	Hash            uint32
	Index           int32
	Redacted        bool
}

type PresentationNodeEntity struct {
	DisplayProperties               DisplayProperties
	OriginalIcon                    string
	RootViewIcon                    string
	NodeType                        int32
	Scope                           Scope
	ObjectiveHash                   uint32
	CompletionRecordHash            uint32
	Children                        PresentationNodeChildrenBlock
	DisplayStyle                    int32
	ScreenStyle                     int32
	Requirements                    PresentationNodeRequirementsBlock
	DisableChildSubscreenNavigation bool
	MaxCategoryRecordScore          int32
	PresentationNodeType            PresentationNodeType
	TraitIds                        []string
	TraitHashes                     []uint32
	ParentNodeHashes                []uint32
	Hash                            uint32
	Index                           int32
	Redacted                        bool
}

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
	Hash                 uint32
	Index                int32
	Redacted             bool
}

type MaterialRequirementSetEntity struct {
	Materials []MaterialRequirement
	Hash      uint32
	Index     int32
	Redacted  bool
}

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
	Hash                 uint32
	Index                int32
	Redacted             bool
}

type GenderEntity struct {
	GenderType        Gender
	DisplayProperties DisplayProperties
	Hash              uint32
	Index             int32
	Redacted          bool
}

type LoreEntity struct {
	DisplayProperties DisplayProperties
	Subtitle          string
	Hash              uint32
	Index             int32
	Redacted          bool
}

type MetricEntity struct {
	DisplayProperties     DisplayProperties
	TrackingObjectiveHash uint32
	LowerValueIsBetter    bool
	PresentationNodeType  PresentationNodeType
	TraitIds              []string
	TraitHashes           []uint32
	ParentNodeHashes      []uint32
	Hash                  uint32
	Index                 int32
	Redacted              bool
}

type EnergyTypeEntity struct {
	DisplayProperties   DisplayProperties
	TransparentIconPath string
	ShowIcon            bool
	EnumValue           EnergyType
	CapacityStatHash    uint32
	CostStatHash        uint32
	Hash                uint32
	Index               int32
	Redacted            bool
}

type PlugSetEntity struct {
	DisplayProperties DisplayProperties
	ReusablePlugItems []ItemSocketEntryPlugItemRandomized
	IsFakePlugSet     bool
	Hash              uint32
	Index             int32
	Redacted          bool
}

type TalentGridEntity struct {
	MaxGridLevel           int32
	GridLevelPerColumn     int32
	ProgressionHash        uint32
	Nodes                  []TalentNode
	ExclusiveSets          []TalentNodeExclusiveSet
	IndependentNodeIndexes []int32
	Groups                 TalentExclusiveGroup
	NodeCategories         []TalentNodeCategory
	Hash                   uint32
	Index                  int32
	Redacted               bool
}

type DamageTypeEntity struct {
	DisplayProperties   DisplayProperties
	TransparentIconPath string
	ShowIcon            bool
	EnumValue           DamageType
	Hash                uint32
	Index               int32
	Redacted            bool
}

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
	GrantDestinySubType     SubType
	GrantDestinyClass       Class
	TraitId                 string
	GroupedCategoryHashes   []uint32
	ParentCategoryHashes    uint32
	GroupCategoryOnly       bool
	Hash                    uint32
	Index                   int32
	Redacted                bool
}

type BreakerTypeEntity struct {
	DisplayProperties DisplayProperties
	EnumValue         BreakerTypeEnum
	Hash              uint32
	Index             int32
	Redacted          bool
}

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
	Hash                                   uint32
	Index                                  int32
	Redacted                               bool
}

type SeasonPassEntity struct {
	DisplayProperties       DisplayProperties
	RewardProgressionHash   uint32
	PrestigeProgressionHash uint32
	Hash                    uint32
	Index                   int32
	Redacted                bool
}

type ChecklistEntity struct {
	DisplayProperties DisplayProperties
	ViewActionString  string
	Scope             Scope
	Entries           []ChecklistEntry
	Hash              uint32
	Index             int32
	Redacted          bool
}

type RaceEntity struct {
	DisplayProperties             DisplayProperties
	RaceType                      Race
	GenderedRaceNames             map[Gender]string
	GenderedRaceNamesByGenderHash map[uint32]string
	Hash                          uint32
	Index                         int32
	Redacted                      bool
}

type ClassEntity struct {
	ClassType                      Class
	DisplayProperties              DisplayProperties
	GenderedClassNames             map[Gender]string
	GenderedClassNamesByGenderHash map[uint32]string
	MentorVendorHash               uint32
	Hash                           uint32
	Index                          int32
	Redacted                       bool
}

type MilestoneEntity struct {
	DisplayProperties               DisplayProperties
	DisplayPreference               int32
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
	Hash                            uint32
	Index                           int32
	Redacted                        bool
}

type UnlockEntity struct {
	DisplayProperties DisplayProperties
	Hash              uint32
	Index             int32
	Redacted          bool
}

type ReportReasonCategoryEntity struct {
	DisplayProperties DisplayProperties
	Reasons           map[uint32]ReportReason
	Hash              uint32
	Index             int32
	Redacted          bool
}
