package destiny2

import "time"

type Color struct {
	Red, Green, Blue, Alpha byte
}

type AnimationReference struct {
	AnimName, AnimIdentifier, Path string
}

type HyperlinkReference struct {
	Title, Url string
}

type DyeReference struct {
	ChannelHash, DyeHash uint32
}

type ArtDyeReference struct {
	ArtDyeChannelHash uint32
}

type GearArtArrangementReference struct {
	ClassHash, ArtArrangementHash uint32
}

type ItemVendorSourceReference struct {
	VendorHash        uint32
	VendorItemIndexes []int32
}

type PlugRule struct {
	FailureMessage string
}

type ParentItemOverride struct {
	AdditionalEquipRequirementsDisplayStrings []string
	PipIcon                                   string
}

type EnergyCapacityEntry struct {
	CapacityValue  int32
	EnergyTypeHash uint32
	EnergyType     int32
}

type EnergyCostEntry struct {
	EnergyCost     int32
	EnergyTypeHash uint32
	EnergyType     int32
}

type InterpolationPoint struct {
	Value, Weight int32
}

type InterpolationPointFloat struct {
	Value, Weight float32
}

type DateRange struct {
	Start, End time.Time
}

type VendorService struct {
	Name string
}

type VendorGroupReference struct {
	VendorGroupHash uint32
}

type UnlockExpression struct {
	Scope int32
}

type ActivityModifierReference struct {
	ActivityModifierHash uint32
}

type ActivityUnlockString struct {
	DisplayString string
}

type ActivityGraphListEntry struct {
	ActivityGraphHash uint32
}

type ActivityGraphNodeStateEntry struct {
	State GraphNodeState
}

type ActivityInsertionPoint struct {
	PhaseHash uint32
}

type TalentNodeStep struct {
	WeaponPerformance, ImpactEffects, GuardianAttributes, LightAbilities, DamageTypes int32
}

type ArtifactTierItem struct {
	ItemHash uint32
}

type PresentationNodeChildEntry struct {
	PresentationNodeHash uint32
}

type PresentationNodeCollectibleChildEntry struct {
	CollectibleHash uint32
}

type PresentationNodeRecordChildEntry struct {
	RecordHash uint32
}

type PresentationNodeMetricChildEntry struct {
	MetricHash uint32
}

type PresentationNodeRequirementsBlock struct {
	EntitlementUnavailableMessage string
}

type PresentationNodesComponent struct {
	Nodes []PresentationNodeComponent
}

type NodeSocketReplaceResponse struct {
	SocketTypeHash uint32
	PlugItemHash   uint32
}

type MilestoneQuestRewards struct {
	Items []MilestoneQuestRewardItem
}

type MilestoneActivity struct {
	ConceptualActivityHash uint32
	Variants               map[uint32]MilestoneActivityVariant
}

type MilestoneActivityVariant struct {
	ActivityHash uint32
	Order        int32
}

type MilestoneVendor struct {
	VendorHash uint32
}

type MilestoneValue struct {
	Key               string
	DisplayProperties DisplayProperties
}

type MilestoneChallenge struct {
	ChallengeObjectiveHash uint32
}

type MilestoneChallengeActivityGraphNodeEntry struct {
	ActivityGraphHash     uint32
	ActivityGraphNodeHash uint32
}

type MilestoneChallengeActivityPhase struct {
	PhaseHash uint32
}

type ReportReason struct {
	ReasonHash        uint32
	DisplayProperties DisplayProperties
}
