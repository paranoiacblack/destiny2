package destiny2

import (
	"encoding/json"
	"strconv"
)

// Contract is a collection of Destiny.Definitions which have their own
// tables in the Manifest Database.
type Contract interface {
	// Name is the name of this contract in the Bungie.Net API.
	Name() string
	// TableName is the name of this contract in the Mobile Manifest Table.
	TableName() string
	// Reference is a link to the Bungie.Net API which describes this contract.
	Reference() string
	// Entity returns a specific entity from this contract with a given hash.
	Entity(hash uint32) interface{}
	// Unmarshal unmarshals the JSON for a contract when calling Manifest.Contract(...).
	// Each contract uses a custom unmarshaller function for convenience.  This is
	// because the JSON data for contracts uses string keys to a given entity.
	// However, many entities and definitions store hash values which reference
	// contracts.  Doing this conversion from string key to hash key on Unmarshal
	// saves the user doing this on every lookup.
	Unmarshal([]byte) error
}

// InventoryItemDefinition is the contract for all Destiny.Definitions.DestinyInventoryItemDefinition entities.
type InventoryItemDefinition map[uint32]InventoryItemEntity

func (InventoryItemDefinition) Name() string {
	return "DestinyInventoryItemDefinition"
}

func (InventoryItemDefinition) TableName() string {
	return "Items"
}

func (InventoryItemDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyInventoryItemDefinition"
}

func (def *InventoryItemDefinition) Unmarshal(data []byte) error {
	m := map[string]InventoryItemEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]InventoryItemEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def InventoryItemDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ProgressionDefinition is the contract for all Destiny.Definitions.DestinyProgressionDefinition entities.
type ProgressionDefinition map[uint32]ProgressionEntity

func (ProgressionDefinition) Name() string {
	return "DestinyProgressionDefinition"
}

func (ProgressionDefinition) TableName() string {
	return "Progressions"
}

func (ProgressionDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyProgressionDefinition"
}

func (def *ProgressionDefinition) Unmarshal(data []byte) error {
	m := map[string]ProgressionEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ProgressionEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ProgressionDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// InventoryBucketDefinition is the contract for all Destiny.Definitions.DestinyInventoryBucketDefinition entities.
type InventoryBucketDefinition map[uint32]InventoryBucketEntity

func (InventoryBucketDefinition) Name() string {
	return "DestinyInventoryBucketDefinition"
}

func (InventoryBucketDefinition) TableName() string {
	return "InventoryBuckets"
}

func (InventoryBucketDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyInventoryBucketDefinition"
}

func (def *InventoryBucketDefinition) Unmarshal(data []byte) error {
	m := map[string]InventoryBucketEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]InventoryBucketEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def InventoryBucketDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ItemTierTypeDefinition is the contract for all Destiny.Definitions.DestinyItemTierTypeDefinition entities.
type ItemTierTypeDefinition map[uint32]ItemTierTypeEntity

func (ItemTierTypeDefinition) Name() string {
	return "DestinyItemTierTypeDefinition"
}

func (ItemTierTypeDefinition) TableName() string {
	return "ItemTierTypes"
}

func (ItemTierTypeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyItemTierTypeDefinition"
}

func (def *ItemTierTypeDefinition) Unmarshal(data []byte) error {
	m := map[string]ItemTierTypeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ItemTierTypeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ItemTierTypeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// StatDefinition is the contract for all Destiny.Definitions.DestinyStatDefinition entities.
type StatDefinition map[uint32]StatEntity

func (StatDefinition) Name() string {
	return "DestinyStatDefinition"
}

func (StatDefinition) TableName() string {
	return "Stats"
}

func (StatDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyStatDefinition"
}

func (def *StatDefinition) Unmarshal(data []byte) error {
	m := map[string]StatEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]StatEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def StatDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// StatGroupDefinition is the contract for all Destiny.Definitions.DestinyStatGroupDefinition entities.
type StatGroupDefinition map[uint32]StatGroupEntity

func (StatGroupDefinition) Name() string {
	return "DestinyStatGroupDefinition"
}

func (StatGroupDefinition) TableName() string {
	return "StatGroups"
}

func (StatGroupDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyStatGroupDefinition"
}

func (def *StatGroupDefinition) Unmarshal(data []byte) error {
	m := map[string]StatGroupEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]StatGroupEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def StatGroupDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// EquipmentSlotDefinition is the contract for all Destiny.Definitions.DestinyEquipmentSlotDefinition entities.
type EquipmentSlotDefinition map[uint32]EquipmentSlotEntity

func (EquipmentSlotDefinition) Name() string {
	return "DestinyEquipmentSlotDefinition"
}

func (EquipmentSlotDefinition) TableName() string {
	return "EquipmentSlots"
}

func (EquipmentSlotDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyEquipmentSlotDefinition"
}

func (def *EquipmentSlotDefinition) Unmarshal(data []byte) error {
	m := map[string]EquipmentSlotEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]EquipmentSlotEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def EquipmentSlotDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// SocketTypeDefinition is the contract for all Destiny.Definitions.DestinySocketTypeDefinition entities.
type SocketTypeDefinition map[uint32]SocketTypeEntity

func (SocketTypeDefinition) Name() string {
	return "DestinySocketTypeDefinition"
}

func (SocketTypeDefinition) TableName() string {
	return "SocketTypes"
}

func (SocketTypeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinySocketTypeDefinition"
}

func (def *SocketTypeDefinition) Unmarshal(data []byte) error {
	m := map[string]SocketTypeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]SocketTypeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def SocketTypeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// SocketCategoryDefinition is the contract for all Destiny.Definitions.DestinySocketCategoryDefinition entities.
type SocketCategoryDefinition map[uint32]SocketCategoryEntity

func (SocketCategoryDefinition) Name() string {
	return "DestinySocketCategoryDefinition"
}

func (SocketCategoryDefinition) TableName() string {
	return "SocketCategories"
}

func (SocketCategoryDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinySocketCategoryDefinition"
}

func (def *SocketCategoryDefinition) Unmarshal(data []byte) error {
	m := map[string]SocketCategoryEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]SocketCategoryEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def SocketCategoryDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// DestinationDefinition is the contract for all Destiny.Definitions.Common.DestinyDestinationDefinition entities.
type DestinationDefinition map[uint32]DestinationEntity

func (DestinationDefinition) Name() string {
	return "DestinyDestinationDefinition"
}

func (DestinationDefinition) TableName() string {
	return "Destinations"
}

func (DestinationDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Common.DestinyDestinationDefinition"
}

func (def *DestinationDefinition) Unmarshal(data []byte) error {
	m := map[string]DestinationEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]DestinationEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def DestinationDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ActivityGraphDefinition is the contract for all Destiny.Definitions.DestinyActivityGraphDefinition entities.
type ActivityGraphDefinition map[uint32]ActivityGraphEntity

func (ActivityGraphDefinition) Name() string {
	return "DestinyActivityGraphDefinition"
}

func (ActivityGraphDefinition) TableName() string {
	return "ActivityGraphs"
}

func (ActivityGraphDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Director.Definitions.DestinyActivityGraphDefinition"
}

func (def *ActivityGraphDefinition) Unmarshal(data []byte) error {
	m := map[string]ActivityGraphEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ActivityGraphEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ActivityGraphDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ActivityDefinition is the contract for all Destiny.Definitions.DestinyActivityDefinition entities.
type ActivityDefinition map[uint32]ActivityEntity

func (ActivityDefinition) Name() string {
	return "DestinyActivityDefinition"
}

func (ActivityDefinition) TableName() string {
	return "Activities"
}

func (ActivityDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyActivityDefinition"
}

func (def *ActivityDefinition) Unmarshal(data []byte) error {
	m := map[string]ActivityEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ActivityEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ActivityDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ActivityModifierDefinition is the contract for all Destiny.Definitions.ActiveModifier.DestinyActivityModifierDefinition entities.
type ActivityModifierDefinition map[uint32]ActivityModifierEntity

func (ActivityModifierDefinition) Name() string {
	return "DestinyActivityModifierDefinition"
}

func (ActivityModifierDefinition) TableName() string {
	return "ActivityModifiers"
}

func (ActivityModifierDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.ActiveModifier.DestinyActivityModifierDefinition"
}

func (def *ActivityModifierDefinition) Unmarshal(data []byte) error {
	m := map[string]ActivityModifierEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ActivityModifierEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ActivityModifierDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ObjectiveDefinition is the contract for all Destiny.Definitions.DestinyObjectiveDefinition entities.
type ObjectiveDefinition map[uint32]ObjectiveEntity

func (ObjectiveDefinition) Name() string {
	return "DestinyObjectiveDefinition"
}

func (ObjectiveDefinition) TableName() string {
	return "Objectives"
}

func (ObjectiveDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyObjectiveDefinition"
}

func (def *ObjectiveDefinition) Unmarshal(data []byte) error {
	m := map[string]ObjectiveEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ObjectiveEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ObjectiveDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// SandboxPerkDefinition is the contract for all Destiny.Definitions.DestinySandboxPerkDefinition entities.
type SandboxPerkDefinition map[uint32]SandboxPerkEntity

func (SandboxPerkDefinition) Name() string {
	return "DestinySandboxPerkDefinition"
}

func (SandboxPerkDefinition) TableName() string {
	return "SandboxPerks"
}

func (SandboxPerkDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinySandboxPerkDefinition"
}

func (def *SandboxPerkDefinition) Unmarshal(data []byte) error {
	m := map[string]SandboxPerkEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]SandboxPerkEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def SandboxPerkDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// LocationDefinition is the contract for all Destiny.Definitions.DestinyLocationDefinition entities.
type LocationDefinition map[uint32]LocationEntity

func (LocationDefinition) Name() string {
	return "DestinyLocationDefinition"
}

func (LocationDefinition) TableName() string {
	return "Locations"
}

func (LocationDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyLocationDefinition"
}

func (def *LocationDefinition) Unmarshal(data []byte) error {
	m := map[string]LocationEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]LocationEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def LocationDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ActivityModeDefinition is the contract for all Destiny.Definitions.DestinyActivityModeDefinition entities.
type ActivityModeDefinition map[uint32]ActivityModeEntity

func (ActivityModeDefinition) Name() string {
	return "DestinyActivityModeDefinition"
}

func (ActivityModeDefinition) TableName() string {
	return "ActivityModes"
}

func (ActivityModeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyActivityModeDefinition"
}

func (def *ActivityModeDefinition) Unmarshal(data []byte) error {
	m := map[string]ActivityModeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ActivityModeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ActivityModeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// PlaceDefinition is the contract for all Destiny.Definitions.DestinyPlaceDefinition entities.
type PlaceDefinition map[uint32]PlaceEntity

func (PlaceDefinition) Name() string {
	return "DestinyPlaceDefinition"
}

func (PlaceDefinition) TableName() string {
	return "Places"
}

func (PlaceDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyPlaceDefinition"
}

func (def *PlaceDefinition) Unmarshal(data []byte) error {
	m := map[string]PlaceEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]PlaceEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def PlaceDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ActivityTypeDefinition is the contract for all Destiny.Definitions.DestinyActivityTypeDefinition entities.
type ActivityTypeDefinition map[uint32]ActivityTypeEntity

func (ActivityTypeDefinition) Name() string {
	return "DestinyActivityTypeDefinition"
}

func (ActivityTypeDefinition) TableName() string {
	return "ActivityTypes"
}

func (ActivityTypeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyActivityTypeDefinition"
}

func (def *ActivityTypeDefinition) Unmarshal(data []byte) error {
	m := map[string]ActivityTypeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ActivityTypeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ActivityTypeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// VendorGroupDefinition is the contract for all Destiny.Definitions.DestinyVendorGroupDefinition entities.
type VendorGroupDefinition map[uint32]VendorGroupEntity

func (VendorGroupDefinition) Name() string {
	return "DestinyVendorGroupDefinition"
}

func (VendorGroupDefinition) TableName() string {
	return "VendorGroups"
}

func (VendorGroupDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyVendorGroupDefinition"
}

func (def *VendorGroupDefinition) Unmarshal(data []byte) error {
	m := map[string]VendorGroupEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]VendorGroupEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def VendorGroupDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// FactionDefinition is the contract for all Destiny.Definitions.DestinyFactionDefinition entities.
type FactionDefinition map[uint32]FactionEntity

func (FactionDefinition) Name() string {
	return "DestinyFactionDefinition"
}

func (FactionDefinition) TableName() string {
	return "Factions"
}

func (FactionDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyFactionDefinition"
}

func (def *FactionDefinition) Unmarshal(data []byte) error {
	m := map[string]FactionEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]FactionEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def FactionDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ArtifactDefinition is the contract for all Destiny.Definitions.Artifacts.DestinyArtifactDefinition entities.
type ArtifactDefinition map[uint32]ArtifactEntity

func (ArtifactDefinition) Name() string {
	return "DestinyArtifactDefinition"
}

func (ArtifactDefinition) TableName() string {
	return "Artifacts"
}

func (ArtifactDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Artifacts.DestinyArtifactDefinition"
}

func (def *ArtifactDefinition) Unmarshal(data []byte) error {
	m := map[string]ArtifactEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ArtifactEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ArtifactDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// PowerCapDefinition is the contract for all Destiny.Definitions.PowerCaps.DestinyPowerCapDefinition entities.
type PowerCapDefinition map[uint32]PowerCapEntity

func (PowerCapDefinition) Name() string {
	return "DestinyPowerCapDefinition"
}

func (PowerCapDefinition) TableName() string {
	return "PowerCaps"
}

func (PowerCapDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.PowerCaps.DestinyPowerCapDefinition"
}

func (def *PowerCapDefinition) Unmarshal(data []byte) error {
	m := map[string]PowerCapEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]PowerCapEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def PowerCapDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ProgressionLevelRequirementDefinition is the contract for all Destiny.Definitions.Progression.DestinyProgressionLevelRequirementDefinition entities.
type ProgressionLevelRequirementDefinition map[uint32]ProgressionLevelRequirementEntity

func (ProgressionLevelRequirementDefinition) Name() string {
	return "DestinyProgressionLevelRequirementDefinition"
}

func (ProgressionLevelRequirementDefinition) TableName() string {
	return "ProgressionLevelRequirements"
}

func (ProgressionLevelRequirementDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Progression.DestinyProgressionLevelRequirementDefinition"
}

func (def *ProgressionLevelRequirementDefinition) Unmarshal(data []byte) error {
	m := map[string]ProgressionLevelRequirementEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ProgressionLevelRequirementEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ProgressionLevelRequirementDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// RewardSourceDefinition is the contract for all Destiny.Definitions.DestinyRewardSourceDefinition entities.
type RewardSourceDefinition map[uint32]RewardSourceEntity

func (RewardSourceDefinition) Name() string {
	return "DestinyRewardSourceDefinition"
}

func (RewardSourceDefinition) TableName() string {
	return "RewardSources"
}

func (RewardSourceDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyRewardSourceDefinition"
}

func (def *RewardSourceDefinition) Unmarshal(data []byte) error {
	m := map[string]RewardSourceEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]RewardSourceEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def RewardSourceDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// TraitDefinition is the contract for all Destiny.Definitions.DestinyTraitDefinition entities.
type TraitDefinition map[uint32]TraitEntity

func (TraitDefinition) Name() string {
	return "DestinyTraitDefinition"
}

func (TraitDefinition) TableName() string {
	return "Traits"
}

func (TraitDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyTraitDefinition"
}

func (def *TraitDefinition) Unmarshal(data []byte) error {
	m := map[string]TraitEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]TraitEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def TraitDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// TraitCategoryDefinition is the contract for all Destiny.Definitions.DestinyTraitCategoryDefinition entities.
type TraitCategoryDefinition map[uint32]TraitCategoryEntity

func (TraitCategoryDefinition) Name() string {
	return "DestinyTraitCategoryDefinition"
}

func (TraitCategoryDefinition) TableName() string {
	return "TraitCategories"
}

func (TraitCategoryDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyTraitCategoryDefinition"
}

func (def *TraitCategoryDefinition) Unmarshal(data []byte) error {
	m := map[string]TraitCategoryEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]TraitCategoryEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def TraitCategoryDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// PresentationNodeDefinition is the contract for all Destiny.Definitions.DestinyPresentationNodeDefinition entities.
type PresentationNodeDefinition map[uint32]PresentationNodeEntity

func (PresentationNodeDefinition) Name() string {
	return "DestinyPresentationNodeDefinition"
}

func (PresentationNodeDefinition) TableName() string {
	return "PresentationNodes"
}

func (PresentationNodeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyPresentationNodeDefinition"
}

func (def *PresentationNodeDefinition) Unmarshal(data []byte) error {
	m := map[string]PresentationNodeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]PresentationNodeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def PresentationNodeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// CollectibleDefinition is the contract for all Destiny.Definitions.Collectibles.DestinyCollectibleDefinition entities.
type CollectibleDefinition map[uint32]CollectibleEntity

func (CollectibleDefinition) Name() string {
	return "DestinyCollectibleDefinition"
}

func (CollectibleDefinition) TableName() string {
	return "Collectibles"
}

func (CollectibleDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Collectibles.DestinyCollectibleDefinition"
}

func (def *CollectibleDefinition) Unmarshal(data []byte) error {
	m := map[string]CollectibleEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]CollectibleEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def CollectibleDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// MaterialRequirementSetDefinition is the contract for all Destiny.Definitions.DestinyMaterialRequirementSetDefinition entities.
type MaterialRequirementSetDefinition map[uint32]MaterialRequirementSetEntity

func (MaterialRequirementSetDefinition) Name() string {
	return "DestinyMaterialRequirementSetDefinition"
}

func (MaterialRequirementSetDefinition) TableName() string {
	return "MaterialRequirementSets"
}

func (MaterialRequirementSetDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyMaterialRequirementSetDefinition"
}

func (def *MaterialRequirementSetDefinition) Unmarshal(data []byte) error {
	m := map[string]MaterialRequirementSetEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]MaterialRequirementSetEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def MaterialRequirementSetDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// RecordDefinition is the contract for all Destiny.Definitions.Records.DestinyRecordDefinition entities.
type RecordDefinition map[uint32]RecordEntity

func (RecordDefinition) Name() string {
	return "DestinyRecordDefinition"
}

func (RecordDefinition) TableName() string {
	return "Records"
}

func (RecordDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyRecordDefinition"
}

func (def *RecordDefinition) Unmarshal(data []byte) error {
	m := map[string]RecordEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]RecordEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def RecordDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// GenderDefinition is the contract for all Destiny.Definitions.DestinyGenderDefinition entities.
type GenderDefinition map[uint32]GenderEntity

func (GenderDefinition) Name() string {
	return "DestinyGenderDefinition"
}

func (GenderDefinition) TableName() string {
	return "Genders"
}

func (GenderDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyGenderDefinition"
}

func (def *GenderDefinition) Unmarshal(data []byte) error {
	m := map[string]GenderEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]GenderEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def GenderDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// LoreDefinition is the contract for all Destiny.Definitions.DestinyLoreDefinition entities.
type LoreDefinition map[uint32]LoreEntity

func (LoreDefinition) Name() string {
	return "DestinyLoreDefinition"
}

func (LoreDefinition) TableName() string {
	return "Lore"
}

func (LoreDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyLoreDefinition"
}

func (def *LoreDefinition) Unmarshal(data []byte) error {
	m := map[string]LoreEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]LoreEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def LoreDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// MetricDefinition is the contract for all Destiny.Definitions.DestinyMetricDefinition entities.
type MetricDefinition map[uint32]MetricEntity

func (MetricDefinition) Name() string {
	return "DestinyMetricDefinition"
}

func (MetricDefinition) TablkeName() string {
	return "Metrics"
}

func (MetricDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyMetricDefinition"
}

func (def *MetricDefinition) Unmarshal(data []byte) error {
	m := map[string]MetricEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]MetricEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def MetricDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// EnergyTypeDefinition is the contract for all Destiny.Definitions.EnergyTypes.DestinyEnergyTypeDefinition entities.
type EnergyTypeDefinition map[uint32]EnergyTypeEntity

func (EnergyTypeDefinition) Name() string {
	return "DestinyEnergyTypeDefinition"
}

func (EnergyTypeDefinition) TableName() string {
	return "EnergyTypes"
}

func (EnergyTypeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.EnergyTypes.DestinyEnergyTypeDefinition"
}

func (def *EnergyTypeDefinition) Unmarshal(data []byte) error {
	m := map[string]EnergyTypeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]EnergyTypeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def EnergyTypeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// PlugSetDefinition is the contract for all Destiny.Definitions.DestinyPlugSetDefinition entities.
type PlugSetDefinition map[uint32]PlugSetEntity

func (PlugSetDefinition) Name() string {
	return "DestinyPlugSetDefinition"
}

func (PlugSetDefinition) TableName() string {
	return "PlugSets"
}

func (PlugSetDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyPlugSetDefinition"
}

func (def *PlugSetDefinition) Unmarshal(data []byte) error {
	m := map[string]PlugSetEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]PlugSetEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def PlugSetDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// TalentGridDefinition is the contract for all Destiny.Definitions.DestinyTalentGridDefinition entities.
type TalentGridDefinition map[uint32]TalentGridEntity

func (TalentGridDefinition) Name() string {
	return "DestinyTalentGridDefinition"
}

func (TalentGridDefinition) TableName() string {
	return "Talents"
}

func (TalentGridDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyTalentGridDefinition"
}

func (def *TalentGridDefinition) Unmarshal(data []byte) error {
	m := map[string]TalentGridEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]TalentGridEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def TalentGridDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// DamageTypeDefinition is the contract for all Destiny.Definitions.DestinyDamageTypeDefinition entities.
type DamageTypeDefinition map[uint32]DamageTypeEntity

func (DamageTypeDefinition) Name() string {
	return "DestinyDamageTypeDefinition"
}

func (DamageTypeDefinition) TableName() string {
	return "DamageTypes"
}

func (DamageTypeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyDamageTypeDefinition"
}

func (def *DamageTypeDefinition) Unmarshal(data []byte) error {
	m := map[string]DamageTypeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]DamageTypeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def DamageTypeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ItemCategoryDefinition is the contract for all Destiny.Definitions.DestinyItemCategoryDefinition entities.
type ItemCategoryDefinition map[uint32]ItemCategoryEntity

func (ItemCategoryDefinition) Name() string {
	return "DestinyItemCategoryDefinition"
}

func (ItemCategoryDefinition) TableName() string {
	return "ItemCategories"
}

func (ItemCategoryDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyItemCategoryDefinition"
}

func (def *ItemCategoryDefinition) Unmarshal(data []byte) error {
	m := map[string]ItemCategoryEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ItemCategoryEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ItemCategoryDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// BreakerTypeDefinition is the contract for all Destiny.Definitions.BreakerTypes.DestinyBreakerTypeDefinition entities.
type BreakerTypeDefinition map[uint32]BreakerTypeEntity

func (BreakerTypeDefinition) Name() string {
	return "DestinyBreakerTypeDefinition"
}

func (BreakerTypeDefinition) TableName() string {
	return "BreakerTypes"
}

func (BreakerTypeDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.BreakerTypes.DestinyBreakerTypeDefinition"
}

func (def *BreakerTypeDefinition) Unmarshal(data []byte) error {
	m := map[string]BreakerTypeEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]BreakerTypeEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def BreakerTypeDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// SeasonDefinition is the contract for all Destiny.Definitions.Seasons.DestinySeasonDefinition entities.
type SeasonDefinition map[uint32]SeasonEntity

func (SeasonDefinition) Name() string {
	return "DestinySeasonDefinition"
}

func (SeasonDefinition) TableName() string {
	return "Seasons"
}

func (SeasonDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Seasons.DestinySeasonDefinition"
}

func (def *SeasonDefinition) Unmarshal(data []byte) error {
	m := map[string]SeasonEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]SeasonEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def SeasonDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// SeasonPassDefinition is the contract for all Destiny.Definitions.DestinySeasonPassDefinition entities.
type SeasonPassDefinition map[uint32]SeasonPassEntity

func (SeasonPassDefinition) Name() string {
	return "DestinySeasonPassDefinition"
}

func (SeasonPassDefinition) TableName() string {
	return "SeasonPasses"
}

func (SeasonPassDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinySeasonPassDefinition"
}

func (def *SeasonPassDefinition) Unmarshal(data []byte) error {
	m := map[string]SeasonPassEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]SeasonPassEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def SeasonPassDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ChecklistDefinition is the contract for all Destiny.Definitions.Checklists.DestinyChecklistDefinition entities.
type ChecklistDefinition map[uint32]ChecklistEntity

func (ChecklistDefinition) Name() string {
	return "DestinyChecklistDefinition"
}

func (ChecklistDefinition) TableName() string {
	return "Checklists"
}

func (ChecklistDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Checklists.DestinyChecklistDefinition"
}

func (def *ChecklistDefinition) Unmarshal(data []byte) error {
	m := map[string]ChecklistEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ChecklistEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ChecklistDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// RaceDefinition is the contract for all Destiny.Definitions.DestinyRaceDefinition entities.
type RaceDefinition map[uint32]RaceEntity

func (RaceDefinition) Name() string {
	return "DestinyRaceDefinition"
}

func (RaceDefinition) TableName() string {
	return "Races"
}

func (RaceDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyRaceDefinition"
}

func (def *RaceDefinition) Unmarshal(data []byte) error {
	m := map[string]RaceEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]RaceEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def RaceDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ClassDefinition is the contract for all Destiny.Definitions.DestinyClassDefinition entities.
type ClassDefinition map[uint32]ClassEntity

func (ClassDefinition) Name() string {
	return "DestinyClassDefinition"
}

func (ClassDefinition) TableName() string {
	return "Classes"
}

func (ClassDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyClassDefinition"
}

func (def *ClassDefinition) Unmarshal(data []byte) error {
	m := map[string]ClassEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ClassEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ClassDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// MilestoneDefinition is the contract for all Destiny.Definitions.Milestones.DestinyMilestoneDefinition entities.
type MilestoneDefinition map[uint32]MilestoneEntity

func (MilestoneDefinition) Name() string {
	return "DestinyMilestoneDefinition"
}

func (MilestoneDefinition) TableName() string {
	return "Milestones"
}

func (MilestoneDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Milestones.DestinyMilestoneDefinition"
}

func (def *MilestoneDefinition) Unmarshal(data []byte) error {
	m := map[string]MilestoneEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]MilestoneEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def MilestoneDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// UnlockDefinition is the contract for all Destiny.Definitions.DestinyUnlockDefinition entities.
type UnlockDefinition map[uint32]UnlockEntity

func (UnlockDefinition) Name() string {
	return "DestinyUnlockDefinition"
}

func (UnlockDefinition) TableName() string {
	return "Unlocks"
}

func (UnlockDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyUnlockDefinition"
}

func (def *UnlockDefinition) Unmarshal(data []byte) error {
	m := map[string]UnlockEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]UnlockEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def UnlockDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// ReportReasonCategoryDefinition is the contract for all Destiny.Definitions.Reporting.DestinyReportReasonCategoryDefinition entities.
type ReportReasonCategoryDefinition map[uint32]ReportReasonCategoryEntity

func (ReportReasonCategoryDefinition) Name() string {
	return "DestinyReportReasonCategoryDefinition"
}

func (ReportReasonCategoryDefinition) TableName() string {
	return "ReportReasonCategories"
}

func (ReportReasonCategoryDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.Reporting.DestinyReportReasonCategoryDefinition"
}

func (def *ReportReasonCategoryDefinition) Unmarshal(data []byte) error {
	m := map[string]ReportReasonCategoryEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]ReportReasonCategoryEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def ReportReasonCategoryDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}

// HistoricalStatsDefinition is the contract for all Destiny.Definitions.DestinyHistoricalStatsDefinition entities.
type HistoricalStatsDefinition map[uint32]HistoricalStatsEntity

func (HistoricalStatsDefinition) Name() string {
	return "DestinyHistoricalStatsDefinition"
}

func (HistoricalStatsDefinition) TableName() string {
	return "HistoricalStats"
}

func (HistoricalStatsDefinition) Reference() string {
	return "https://bungie-net.github.io/#/components/schemas/Destiny.Definitions.DestinyHistoricalStatsDefinition"
}

func (def *HistoricalStatsDefinition) Unmarshal(data []byte) error {
	m := map[string]HistoricalStatsEntity{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	hashMap := map[uint32]HistoricalStatsEntity{}
	for key, entity := range m {
		hash, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return err
		}
		hashMap[uint32(hash)] = entity
	}
	*def = hashMap
	return nil
}

func (def HistoricalStatsDefinition) Entity(entityHash uint32) interface{} {
	return def[entityHash]
}