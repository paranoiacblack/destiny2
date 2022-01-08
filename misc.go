package destiny2

import "time"

// Color represents a color with RGBA values represented between 0 and 255.
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

// ItemVendorSourceReference represents that a vendor could sell this item.
type ItemVendorSourceReference struct {
	// VendorHash is the hash of a related VendorEntity.
	VendorHash uint32
	// VendorItemIndexes are the indexes at which this item is sold at this vendor.
	VendorItemIndexes []int32
}

type InterpolationPoint struct {
	Value, Weight int32
}

type InterpolationPointFloat struct {
	Value, Weight float32
}

type VendorRequirementDisplayEntry struct {
	Icon, Name, Source, Type string
}

type DateRange struct {
	Start, End time.Time
}

// VendorInteractionSackEntry is the sack of items to show in a vendor interaction.
type VendorInteractionSackEntry struct {
	// SackType is the sack identifier in InventoryItemEntity.VendorSackType.
	// If these match, this sack is shown in the vendor interaction.
	SackType uint32
}

// ItemCreationEntryLevel is a wrapper for the item level at which an item should spawn.
type ItemCreationEntryLevel struct {
	// Level is the minimum character level the item spawns at.
	Level int32
}

// VendorService is the localized name of a vendor-provided service.
type VendorService struct {
	// Name is the localized name of a service provided.
	Name string
}

// ActivityGraphListEntry is the default map to show for a given activity or destination.
type ActivityGraphListEntry struct {
	// ActivityGraphHash is the hash of a related ActivityGraphEntity.
	ActivityGraphHash uint32
}

// ActivityGraphNodeFeaturingState is a single visual state that a node can be in.
type ActivityGraphNodeFeaturingState struct {
	// HighlightType is a classification of different ways this node feature can be highlighted.
	HighlightType ActivityGraphNodeHighlightType
}

// ActivityModifierReference is a reference to an activity modifier from another entity.
type ActivityModifierReference struct {
	// ActivityModifierHash is the hash of a related ActivityModifierEntity.
	ActivityModifierHash uint32
}

// ActivityUnlockString represents a status string that could be displayed about an activity.
type ActivityUnlockString struct {
	// DisplayString is the localized string to display if certain conditions are met.
	DisplayString string
}

type ActivityLoadoutRequirementSet struct {
	// Requirements are the set of requirements that will be applied on the activity if this set is active.
	Requirements []ActivityLoadout
}

// ActivityInsertionPoint is a point of entry into an activity.
type ActivityInsertionPoint struct {
	// PhaseHash is the unique hash representing the phase.
	PhaseHash uint32
}

// ActivityGraphNodeStateEntry represents a single state that a graph node might end up in.
type ActivityGraphNodeStateEntry struct {
	State GraphNodeState
}

// ActivityGraphArtElement represents one-off visual effects overlayed on the map.
type ActivityGraphArtElement struct {
	// Position is the position on the map of this art element.
	Position Position
}

// UnlockExpression is the foundation of the game's gating mechanics and othe related restrictions.
type UnlockExpression struct {
	Scope GatingScope
}

type LinkedGraphEntry struct {
	// ActivityGraphHash is the hash of a related ActivityGraphEntity.
	ActivityGraphHash uint32
}

// DestinationBubbleSetting is human readable data about a bubble. Deprecated.
type DestinationBubbleSetting struct {
	DisplayProperties DisplayProperties
}

// VendorGroupReference refers to a grouping of vendors.
type VendorGroupReference struct {
	// VendorGroupHash is the hash of a related VendorGroupEntity.
	VendorGroupHash uint32
}

// ArtifactTierItem is a plug item unlocked by activating this item in the artifact.
type ArtifactTierItem struct {
	// ItemHash is the hash for a related InventoryItemEntity.
	ItemHash uint32
}

// ItemVersion refers to the power cap for this item version.
type ItemVersion struct {
	// PowerCapHash is the hash for a related PowerCapEntity.
	PowerCapHash uint32
}

type TalentNodeStep struct {
	WeaponPerformance, ImpactEffects, GuardianAttributes, LightAbilities, DamageTypes int32
}

// ItemMetricBlock represents metrics available for display and selection on an item.
type ItemMetricBlock struct {
	// AvailableMetricCategoryNodeHashes are the hashes of all related PresentationNodeEntity structs.
	AvailableMetricCategoryNodeHashes []uint32
}

type PresentationNodeChildEntry struct {
	// PresentationNodeHash is the hash of a related PresentationNodeEntity.
	PresentationNodeHash uint32
}

type PresentationNodeCollectibleChildEntry struct {
	// CollectibleHash is the hash of a related CollectibleEntity.
	CollectibleHash uint32
}

type CollectibleAcquisitionBlock struct {
	// AcquireMaterialRequirementHash is the hash of a related MaterialRequirementSetEntity.
	AcquireMaterialRequirementHash uint32
	// AcquireTimestampUnlockValueHash is the hash of a related UnlockValueEntity.
	AcquireTimestampUnlockValueHash uint32
}

// PresentationNodeRequirementsBlock defines the requirements for showing a presentation node.
type PresentationNodeRequirementsBlock struct {
	// EntitlementUnavailableMessage is the localized string to show if this node is inaccessible due to entitlements.
	EntitlementUnavailableMessage string
}

// PresentationNodeRecordChildEntry is an entry for a record-based presentation node.
type PresentationNodeRecordChildEntry struct {
	// RecordHash is the hash of a related RecordEntity.
	RecordHash uint32
}

// RecordIntervalRewards are items rewarded for completing a record interval.
type RecordIntervalRewards struct {
	// IntervalRewardItems are a list of items and their quantities rewarded for completing this record interval.
	IntervalRewardItems []ItemQuantity
}

// PresentationNodeMetricChildEntry is an entry for a metric-related presentation node.
type PresentationNodeMetricChildEntry struct {
	// MetricHash is the hash of a related MetricEntity.
	MetricHash uint32
}

// PlugRule describes a rule around whether the plug is enabled or insertable.
type PlugRule struct {
	// FailureMessage is the localized string to show if this rule fails.
	FailureMessage string
}

// ParentItemOverride describes how a plug item overrides certain properties of the item it is socketed into.
type ParentItemOverride struct {
	// AdditionalEquipRequirementsDisplayStrings are localized strings
	// that describe additional requirements for equipping the parent item.
	AdditionalEquipRequirementsDisplayStrings []string
	// PIPIcon is the icon to show when viewing an item that has this plug socketed.
	PipIcon string
}

// EnergyCapacityEntry is an entry for a plug item which grants energy capacity to the item is it socketed into.
type EnergyCapacityEntry struct {
	// CapacityValue is the amount of energy capacity this plug provides.
	CapacityValue int32
	// EnergyTypeHash is the hash of a related EnergyTypeEntity.
	EnergyTypeHash uint32
	// EnergyType is a classification of the type of energy capacity granted.
	EnergyType EnergyType
}

// EnergyCostEntry is an entry for a plug item that costs energy to insert into an item.
type EnergyCostEntry struct {
	// EnergyCost is the cost of inserting this plug.
	EnergyCost int32
	// EnergyTypeHash is the hash of a related EnergyTypeEntity.
	EnergyTypeHash uint32
	// EnergyType is a classification of the type of energy that the plug costs.
	EnergyType EnergyType
}

// NodeActivationRequirement describes the requirements for activation a talent node.
type NodeActivationRequirement struct {
	// GridLevel is progression level on the talent grid required to activate this node.
	GridLevel int32
	// MaterialRequirementHashes are the hashes of all related MaterialRequirementSetEntity structs.
	MaterialRequirementHashes []uint32
}

// NodeSocketReplaceResponse describes how a socket on an item is replaced with a new plug on a node step's activation.
type NodeSocketReplaceResponse struct {
	// SocketTypeHash is the hash of a related SocketTypeEntity.
	SocketTypeHash uint32
	// PlugItemHash is the hash of a related InventoryItemEntity.
	PlugItemHash uint32
}

// TalentExclusiveGroup describes a node that exists as part of an exclusive group.
// An exclusive group is a group of nodes that can and cannot be activated together.
type TalentExclusiveGroup struct {
	// GroupHash is the unique identifer for this exclusive group within the talent grid.
	GroupHash uint32
	// LoreHash is the hash of a related LoreEntity.
	LoreHash uint32
	// NodeHashes are the hashes of talent nodes that are part of this group.
	NodeHashes []uint32
	// OpposingGroupHashes are the unique identifiers for all exclusive groups that will be deactivated if any node in this group is activated.
	OpposingGroupHashes []uint32
	// OpposingNodeHashes are the hashs of talent nodes that are deactivated if any node in this group is activated.
	OpposingNodeHashes []uint32
}

// TalentNodeCategory describes a group of talent nodes by functionality.
type TalentNodeCategory struct {
	// Identifier is an identifier for this category.
	Identifier string
	// IsLoreDriven determines if this category has a related LoreEntity.
	IsLoreDriven      bool
	DisplayProperties DisplayProperties
	// NodeHashes are the hashes of all talent nodes in this talent grid that are in this category.
	NodeHashes []uint32
}

// PresentationNodesComponent describes the node components that compose each presentation node.
type PresentationNodesComponent struct {
	// Nodes are presentation components by the hash of a related PresentationNodeEntity.
	Nodes map[uint32]PresentationNodeComponent
}

// PresentationNodeComponent describes the components of a presentation node.
type PresentationNodeComponent struct {
	// State is the state of this component.
	State PresentationNodeState
	// Objective is the progress for an objective in the presentation node.
	Objective ObjectiveProgress
	// ProgressValue is how much of the presentation node is completed so far.
	ProgressValue int32
	// CompletionValue is the value at which the presentation node is considered complete.
	CompletionValue int32
	// RecordCategoryScore is the current score of the record category the presentation node represents.
	RecordCategoryScore int32
}
