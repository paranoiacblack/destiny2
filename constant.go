package destiny2

// EnvironmentLocationMapping describes a map location where an item, objective or activity can be used or started.
type EnvironmentLocationMapping struct {
	// LocationHash is the hash of a related LocationEntity.
	LocationHash uint32
	// ActivationSource is a hint a given UI uses to figure out how this location is activated.
	ActivationSource string
	// ItemHash is the hash of a related InventoryItemEntity.
	ItemHash uint32
	// ObjectiveHash is the hash of a related ObjectiveEntity.
	ObjectiveHash uint32
	// ActivityHash is the hash of a related ActivityEntity.
	ActivityHash uint32
}
