package destiny2

type ItemTooltipNotification struct {
	DisplayString string
	DisplayStyle  string
}

type ItemActionBlock struct {
	VerbName                string
	VerbDescription         string
	IsPositive              bool
	OverlayScreenName       string
	OverlayIcon             string
	RequiredCooldownSeconds int32
	RequiredItems           []ItemActionRequiredItem
	ProgressionRewards      []ProgressionReward
	ActionTypeLabel         string
	RequiredLocation        string
	RequiredCooldownHash    uint32
	DeleteOnAction          bool
	ConsumeEntireStack      bool
	UseOnAcquire            bool
}

type ItemInventoryBlock struct {
	StackUniqueLabel                         string
	MaxStackSize                             int32
	BucketTypeHash                           uint32
	RecoveryBucketTypeHash                   uint32
	TierTypeHash                             uint32
	IsInstanceItem                           bool
	TierTypeName                             string
	TierType                                 ItemTier
	ExpirationTooltip                        string
	ExpiredInActivityMessage                 string
	ExpiredInOrbitMessage                    string
	SuppressExpirationWhenObjectivesComplete bool
}

type ItemSetBlock struct {
	ItemList                 []ItemSetBlockEntry
	RequireOrderedSetItemAdd bool
	SetIsFeatured            bool
	SetType                  string
	QuestLineName            string
	QuestLineDescription     string
	QuestStepSummary         string
}

type ItemSetBlockEntry struct {
	TrackingValue int32
	ItemHash      uint32
}

type ItemStatBlock struct {
	DisablePrimaryStatDisplay bool
	StatGroupHash             uint32
	Stats                     map[uint32]InventoryItemStat
	HasDisplayableStats       bool
	PrimaryBaseStatHash       uint32
}

type InventoryItemStat struct {
	StatHash       uint32
	Value          int32
	Minimum        int32
	Maximum        int32
	DisplayMaximum int32
}

type EquippingBlock struct {
	GearsetItemHash       uint32
	UniqueLabel           string
	UniqueLabelHash       uint32
	EquipmentSlotTypeHash uint32
	Attributes            EquippingItemBlockAttributes
	AmmoType              AmmunitionType
	DisplayStrings        []string
}

type ItemTranslationBlock struct {
	WeaponPatternIdentifier string
	WeaponPatternHash       uint32
	DefaultDyes             []DyeReference
	LockedDyes              []DyeReference
	CustomDyes              []DyeReference
	Arrangements            []GearArtArrangementReference
	HasGeometry             bool
}

type ItemPreviewBlock struct {
	ScreenStyle           string
	PreviewVendorHash     uint32
	ArtifactHash          uint32
	PreviewActionString   string
	DerivedItemCategories []DerivedItemCategory
}

type DerivedItemCategory struct {
	CategoryDescription string
	Items               []DerivedItem
}

type DerivedItem struct {
	ItemHash        uint32
	ItemName        string
	ItemDescription string
	IconPath        string
	VendorItemIndex int32
}

type ItemQualityBlock struct {
	ItemLevels                      []int32
	QualityLevel                    int32
	InfusionCategoryName            string
	InfusionCategoryHash            uint32
	InfusionCategoryHashes          []uint32
	ProgressionLevelRequirementHash uint32
	CurrentVersion                  uint32
	Versions                        []ItemVersion
	DisplayVersionWatermarkicons    []string
}

type ItemVersion struct {
	PowerCapHash uint32
}

type ItemValueBlock struct {
	ItemValue        []ItemQuantity
	ValueDescription string
}

type ItemQuantity struct {
	ItemHash                 uint32
	ItemInstanceId           int64
	Quantity                 int32
	HasConditionalVisibility bool
}

type ItemSourceBlock struct {
	SourceHashes  []uint32
	Sources       []ItemSource
	Exclusive     BungieMembershipType
	VendorSources []ItemVendorSourceReference
}

type ItemSource struct {
	Level            int32
	MinQuality       int32
	MaxQuality       int32
	MinLevelRequired int32
	MaxLevelRequired int32
	ComputedStats    map[uint32]InventoryItemStat
	SourceHashes     []uint32
}

type ItemObjectiveBlock struct {
	ObjectiveHashes                []uint32
	DisplayActivityHashes          []uint32
	RequireFullObjectiveCompletion bool
	QuestlineItemHash              uint32
	Narrative                      string
	ObjectiveVerbName              string
	QuestTypeIdentifier            string
	QuestTypeHash                  uint32
	PerObjectiveDisplayProperties  []ObjectiveDisplayProperties
	DisplayAsStatTracker           bool
}

type ObjectiveDisplayProperties struct {
	ActivityHash               uint32
	DisplayOnItemPreviewScreen bool
}

type ObjectiveProgress struct {
	ObjectiveHash   uint32
	DestinationHash uint32
	ActivityHash    uint32
	Progress        int32
	CompletionValue int32
	Complete        bool
	Visible         bool
}

type ItemMetricBlock struct {
	AvailableMetricCategoryNodeHashes []uint32
}

type ItemPlug struct {
	InsertionRules                   []PlugRule
	PlugCategoryIdentifier           string
	PlugCategoryHash                 uint32
	InsertionMaterialRequirementHash uint32
	PreviewItemOverrideHash          uint32
	EnabledMaterialRequirementHash   uint32
	EnabledRules                     []PlugRule
	UiPlugLabels                     string
	PlugStyle                        PlugUIStyles
	PlugAvailable                    PlugAvailabilityMode
	AlternateUiPlugLabel             string
	AlternatePlugStyle               PlugUIStyles
	IsDummyPlug                      bool
	ParentItemOverride               ParentItemOverride
	EnergyCapacity                   EnergyCapacityEntry
	EnergyCost                       EnergyCostEntry
}

type ItemGearsetBlock struct {
	TrackingValueMax int32
	ItemList         []uint32
}

type ItemSackBlock struct {
	DetailAction    string
	OpenAction      string
	SelectItemCount int32
	VendorSackType  string
	OpenOnAcquire   bool
}

type ItemSocketBlock struct {
	Detail           string
	SocketEntries    []ItemSocketEntry
	IntrinsicSockets []ItemIntrinsicSocketEntry
	SocketCategories []ItemSocketCategory
}

type ItemSocketEntry struct {
	SocketTypeHash                        uint32
	SingleInitialItemHash                 uint32
	ReusablePlugItems                     []ItemSocketEntryPlugItem
	PreventInitializationOnVendorPurchase bool
	HidePerksInItemTooltip                bool
	PlugSources                           SocketPlugSources
	ReusablePlugSetHash                   uint32
	RandomizedPlugSetHash                 uint32
	DefaultVisible                        bool
}

type ItemSocketEntryPlugItem struct {
	PlugItemHash uint32
}

type ItemIntrinsicSocketEntry struct {
	PlugItemHash   uint32
	SocketTypeHash uint32
	DefaultVisible bool
}

type ItemSocketCategory struct {
	SocketCategoryHash uint32
	SocketIndexes      []int32
}

type ItemSummaryBlock struct {
	SortPriority int32
}

type ItemTalentGridBlock struct {
	TalentGridHash   uint32
	ItemDetailString string
	BuildName        string
	HudDamageType    DamageType
	HudIcon          string
}

type ItemInvestmentStat struct {
	StatTypeHash          uint32
	Value                 int32
	IsConditionallyActive bool
}

type ItemPerkEntry struct {
	RequirementDisplayString string
	PerkHash                 uint32
	PerkVisibility           ItemPerkVisibility
}

type ItemActionRequiredItem struct {
	Count          int32
	ItemHash       uint32
	DeleteOnAction bool
}

type ProgressionReward struct {
	ProgressionMappingHash uint32
	Amount                 int32
	ApplyThrottles         bool
}

type ProgressionStep struct {
	StepName          string
	DisplayEffectType ProgressionStepDisplayEffect
	ProgressTotal     int32
	RewardItems       []ItemQuantity
	Icon              string
}

type ProgressionRewardItemQuantity struct {
	RewardAtProgressionLevel  int32
	AcquisitionBehavior       ProgressionRewardItemAcquisitionBehavior
	UiDisplayStyle            string
	ClaimUnlockDisplayStrings []string
	ItemHash                  uint32
	ItemInstanceId            int64
	Quantity                  int32
	HasConditionalVisibility  bool
}

type ItemTierTypeInfusionBlock struct {
	BaseQualityTransferRadio float32
	MinimumQualityIncrement  int32
}

type StatDisplay struct {
	StatHash             uint32
	MaximumValue         int32
	DisplayAsNumeric     bool
	DisplayInterpolation []InterpolationPoint
}

type StatOverride struct {
	StatHash          uint32
	DisplayProperties DisplayProperties
}

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

type VendorCategoryEntry struct {
	CategoryIndex                int32
	SortValue                    int32
	CategoryHash                 uint32
	QuantityAvailable            int32
	ShowUnavailableItems         bool
	HideIfNoCurrency             bool
	HideFromRegularPurchase      bool
	BuyStringOverride            string
	DisabledDescription          string
	DisplayTitle                 string
	Overlay                      VendorCategoryOverlay
	VendorItemIndexes            []int32
	IsPreview                    bool
	IsDisplayOnly                bool
	ResetIntervalMinutesOverride int32
	ResetOffsetMinutesOverride   int32
}

type VendorCategoryOverlay struct {
	ChoiceDescription string
	Description       string
	Icon              string
	Title             string
	CurrencyItemHash  uint32
}

type DisplayCategory struct {
	Index                  int32
	Identifier             string
	DisplayCategoryHash    uint32
	DisplayProperties      DisplayProperties
	DisplayInBanner        bool
	ProgressionHash        uint32
	SortOrder              int32
	DisplayStyleHash       uint32
	DisplayStyleIdentifier string
}

type VendorInteraction struct {
	InteractionIndex          int32
	Replies                   []VendorInteractionReply
	VendorCategoryIndex       int32
	QuestlineItemHash         uint32
	SackInteractionList       []VendorInteractionSackEntry
	UiInteractionType         uint32
	InteractionType           VendorInteractionType
	RewardBlockLabel          string
	RewardVendorCategoryIndex int32
	FlavorLineOne             string
	FlavorLineTwo             string
	HeaderDisplayProperties   DisplayProperties
	Instructions              string
}

type VendorInteractionReply struct {
	ItemRewardsSelection VendorInteractionRewardSelection
	Reply                string
	ReplyType            VendorReplyType
}

type VendorInteractionSackEntry struct {
	SackType uint32
}

type VendorInventoryFlyout struct {
	LockedDescription string
	DisplayProperties DisplayProperties
	Buckets           []VendorInventoryFlyoutBucket
	Flyoutid          uint32
	SuppressNewness   bool
	EquipmentSlotHash uint32
}

type VendorInventoryFlyoutBucket struct {
	Collapsible         bool
	InventoryBucketHash uint32
	SortItemsBy         int32
}

type VendorItem struct {
	VendorItemIndex       int32
	ItemHash              uint32
	Quantity              int32
	FailureIndexes        []int32
	Currencies            []VendorItemQuantity
	RefundPolicy          VendorItemRefundPolicy
	RefundTimeLimit       int32
	CreationLevels        []ItemCreationEntryLevel
	DisplayCategoryIndex  int32
	CategoryIndex         int32
	OriginalCategoryIndex int32
	MinimumLevel          int32
	MaximumLevel          int32
	Action                VendorSaleItemActionBlock
	DisplayCategory       string
	InventoryBucketHash   uint32
	VisibilityScope       GatingScope
	PurchasableScope      GatingScope
	Exclusivity           BungieMembershipType
	IsOffer               bool
	IsCrm                 bool
	SortValue             int32
	ExpirationTooltip     string
	RedirectToSaleIndexes []int32
	SocketOverrides       []VendorItemSocketOverride
	Unpurchasable         bool
}

type VendorItemQuantity struct {
	ItemHash                 uint32
	ItemInstanceId           int64
	Quantity                 int32
	HasConditionalVisibility bool
}

type ItemCreationEntryLevel struct {
	Level int32
}

type VendorSaleItemActionBlock struct {
	ExecuteSeconds float32
	IsPositive     bool
}

type VendorItemSocketOverride struct {
	SingleItemHash         uint32
	RandomizedOptionsCount int32
	SocketTypeHash         uint32
}

type VendorAcceptedItem struct {
	AcceptedInventoryBucketHash    uint32
	DestinationInventoryBucketHash uint32
}

type VendorLocation struct {
	DestinationHash     uint32
	BackgroundImagePath string
}

type InsertPlugAction struct {
	ActionExecuteSeconds int32
	ActionType           SocketTypeActionType
}

type PlugWhitelistEntry struct {
	CategoryHash                       uint32
	CategoryIdentifier                 string
	ReinitializationPossiblePlugHashes []uint32
}

type SocketTypeScalarMaterialRequirementEntry struct {
	CurrencyItemHash uint32
	ScalarValue      int32
}

// DestinationBubbleSetting is deprecated. Use Bubble instead.
type DestinationBubbleSetting struct {
	DisplayProperties DisplayProperties
}

type Bubble struct {
	Hash              uint32
	DisplayProperties DisplayProperties
}

type ActivityGraphNode struct {
	NodeId          uint32
	OverrideDisplay DisplayProperties
	Position        Position
	FeaturingStates []ActivityGraphNodeFeaturingState
	Activities      []ActivityGraphNodeActivity
	States          []ActivityGraphNodeStateEntry
}

type ActivityGraphNodeFeaturingState struct {
	HighlightType ActivityGraphNodeHighlightType
}

type ActivityGraphNodeActivity struct {
	NodeActivityId uint32
	ActivityHash   uint32
}

type ActivityGraphArtElement struct {
	Position Position
}

type ActivityGraphConnection struct {
	SourceNodeHash uint32
	DestNodeHash   uint32
}

type ActivityGraphDisplayObjective struct {
	Id            uint32
	ObjectiveHash uint32
}

type ActivityGraphDisplayProgression struct {
	Id              uint32
	ProgressionHash uint32
}

type LinkedGraph struct {
	Description      string
	Name             string
	UnlockExpression UnlockExpression
	LinkedGraphId    uint32
	LinkedGraphs     []LinkedGraphEntry
	Overview         string
}

type LinkedGraphEntry struct {
	ActivityGraphHash uint32
}

type ActivityReward struct {
	RewardText  string
	RewardItems []ItemQuantity
}

type ActivityChallenge struct {
	ObjectiveHash uint32
	DummyRewards  []ItemQuantity
}

type ActivityPlaylistItem struct {
	ActivityHash           uint32
	DirectActivityModeHash uint32
	DirectActivityModeType ActivityModeType
	ActivityModeHashes     []uint32
	ActivityModeTypes      []ActivityModeType
}

type ActivityMatchmakingBlock struct {
	IsMatchmade          bool
	MinParty             int32
	MaxParty             int32
	MaxPlayers           int32
	RequiresGuardianOath bool
}

type ActivityGuidedBlock struct {
	GuidedMaxLobbySize int32
	GuidedMinLobbySize int32
	GuidedDisbandCount int32
}

type ActivityLoadoutRequirementSet struct {
	Requirements []ActivityLoadout
}

type ActivityLoadout struct {
	EquipmentSlotHash         uint32
	AllowedEquippedItemHashes []uint32
	AllowedWeaponSubTypes     []int32
}

type ObjectivePerkEntry struct {
	PerkHash uint32
	Style    ObjectiveGrantStyle
}

type ObjectiveStatEntry struct {
	Stat  ItemInvestmentStat
	Style ObjectiveGrantStyle
}

type LocationRelease struct {
	DisplayProperties       DisplayProperties
	SmallTransparentIcon    string
	MapIcon                 string
	LargeTransparentIcon    string
	SpawnPoint              uint32
	DestinationHash         uint32
	ActivityGraphHash       uint32
	ActivityGraphNodeHash   uint32
	ActivityBubbleName      uint32
	ActivityPathBundle      uint32
	ActivityPathDestination uint32
	NavPointType            ActivityNavPointType
	WorldPosition           []int32
}

type FactionVendor struct {
	VendorHash          uint32
	DestinationHash     uint32
	BackgroundImagePath string
}

type ArtifactTier struct {
	TierHash                           uint32
	DisplayTitle                       string
	ProgressRequirementMessage         string
	Items                              []ArtifactTierItem
	MinimumUnlockPointsUsedRequirement int32
}

type PresentationNodeChildrenBlock struct {
	PresentationNodes []PresentationNodeChildEntry
	Collectibles      []PresentationNodeCollectibleChildEntry
	Records           []PresentationNodeRecordChildEntry
	Metrics           []PresentationNodeMetricChildEntry
}

type PresentationChildBlock struct {
	PresentationNodeType          PresentationNodeType
	ParentPresentatiionNodeHashes []uint32
	DisplayStyle                  PresentationDisplayStyle
}

type PresentationNodeComponent struct {
	State               PresentationNodeState
	Objective           ObjectiveProgress
	ProgressValue       int32
	CompletionValue     int32
	RecordCategoryScore int32
}

type CollectibleAcquisitionBlock struct {
	AcquireMaterialRequirementHash  uint32
	AcquireTimestampUnlockValueHash uint32
}

type CollectibleStateBlock struct {
	ObscuredOverrideItemHash uint32
	Requirements             PresentationNodeRequirementsBlock
}

type MaterialRequirement struct {
	ItemHash             uint32
	DeleteOnAction       bool
	Count                int32
	OmitFromRequirements bool
}

type RecordTitleBlock struct {
	HasTitle                  bool
	TitlesByGender            map[string]string
	TitlesByGenderHash        map[uint32]string
	GildingTrackingRecordHash uint32
}

type RecordCompletionBlock struct {
	PartialCompletionObjectiveCountThreshold int32
	ScoreValue                               int32
	ShouldFireToast                          bool
	ToastStyle                               RecordToastStyle
}

type RecordStateBlock struct {
	FeaturedPriority int32
	ObscuredString   string
}

type RecordExpirationBlock struct {
	HasExpiration bool
	Description   string
	Icon          string
}

type RecordIntervalBlock struct {
	IntervalObjectives                   []RecordIntervalObjective
	IntervalRewards                      []RecordIntervalRewards
	OriginalObjectiveArrayInsertionIndex int32
}

type RecordIntervalObjective struct {
	IntervalObjectiveHash uint32
	IntervalScoreValue    int32
}

type RecordIntervalRewards struct {
	IntervalRewardItems []ItemQuantity
}

type ItemSocketEntryPlugItemRandomized struct {
	CurrentlyCanRoll bool
	PlugItemHash     uint32
}

type TalentNode struct {
	NodeIndex                              int32
	NodeHash                               uint32
	Row                                    int32
	Column                                 int32
	PrerequisiteNodeIndexes                []int32
	BinaryPairNodeIndex                    int32
	AutoUnlocks                            bool
	LastStepRepeats                        bool
	IsRandom                               bool
	RandomActivationRequirement            NodeActivationRequirement
	IsRandomRepurchasable                  bool
	Steps                                  []NodeStep
	ExclusiveWithNodeHashes                []uint32
	RandomStartProgressionBarAtProgression int32
	LayoutIdentifier                       string
	GroupHash                              uint32
	LoreHash                               uint32
	NodeStyleIdentifier                    string
	IgnoreForCompletion                    bool
}

type NodeActivationRequirement struct {
	GridLevel                 int32
	MaterialRequirementHashes []uint32
}

type NodeStep struct {
	DisplayProperties             DisplayProperties
	StepIndex                     int32
	NodeStepHash                  uint32
	InteractionDescription        string
	DamageType                    DamageType
	DamageTypeHash                uint32
	ActivationRequirement         NodeActivationRequirement
	CanActivateNextStep           bool
	NextStepIndex                 int32
	IsNextStepRandom              bool
	PerkHashes                    []uint32
	StartProgressionBarAtProgress int32
	StatHashes                    []uint32
	AffectsQuality                bool
	StepGroups                    TalentNodeStepGroups
	AffectsLevel                  bool
	SocketReplacements            []NodeSocketReplaceResponse
}

type TalentNodeStepGroups struct {
	WeaponPerformance  TalentNodeStepWeaponPerformances
	ImpactEffects      TalentNodeStepImpactEffects
	GuardianAttributes TalentNodeStepGuardianAttributes
	LightAbilities     TalentNodeStepLightAbilities
	DamageTypes        TalentNodeStepDamageTypes
}

type TalentNodeExclusiveSet struct {
	NodeIndexes []int32
}

type TalentExclusiveGroup struct {
	GroupHash           uint32
	LoreHash            uint32
	NodeHashes          []uint32
	OpposingGroupHashes []uint32
	OpposingNodeHashes  []uint32
}

type TalentNodeCategory struct {
	Identifier        string
	IsLoreDriven      bool
	DisplayProperties DisplayProperties
	NodeHashes        []uint32
}

type SeasonPreview struct {
	Description string
	LinkPath    string
	VideoLink   string
	Images      []SeasonPreviewImage
}

type SeasonPreviewImage struct {
	ThumbnailImage string
	HighResImage   string
}

type ChecklistEntry struct {
	Hash                   uint32
	DisplayProperties      DisplayProperties
	DestinationHash        uint32
	LocationHash           uint32
	BubbleHash             uint32
	ActivityHash           uint32
	ItemHash               uint32
	VendorHash             uint32
	VendorInteractionIndex int32
	Scope                  Scope
}

type MilestoneQuest struct {
	QuestItemHash     uint32
	DisplayProperties DisplayProperties
	OverrideImage     string
	QuestRewards      MilestoneQuestRewards
	Activities        map[uint32]MilestoneActivity
	DestinationHash   uint32
}

type MilestoneQuestRewardItem struct {
	VendorHash               uint32
	VendorItemIndex          int32
	ItemHash                 uint32
	ItemInstanceId           int64
	Quantity                 int32
	HasConditionalVisibility bool
}

type MilestoneRewardCategory struct {
	CategoryHash       uint32
	CategoryIdentifier string
	DisplayProperties  DisplayProperties
	RewardEntries      map[uint32]MilestoneRewardEntry
	Order              int32
}

type MilestoneRewardEntry struct {
	RewardEntryHash       uint32
	RewardEntryIdentifier string
	Items                 []ItemQuantity
	VendorHash            uint32
	DisplayProperties     DisplayProperties
	Order                 int32
}

type MilestoneChallengeActivity struct {
	ActivityHash       uint32
	Challenges         []MilestoneChallenge
	ActivityGraphNodes []MilestoneChallengeActivityGraphNodeEntry
	Phases             []MilestoneChallengeActivityPhase
}
