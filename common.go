package destiny2

// DisplayProperties represents common display information for contract entities.
type DisplayProperties struct {
	Description   string
	Name          string
	Icon          string
	IconSequences []IconSequence
	HighResIcon   string
	HasIcon       bool
}

// IconSequence is a sequences of icons represented as frames.
type IconSequence struct {
	Frames []string
}

// Position is 3D position.
type Position struct {
	X, Y, Z int32
}

// EntityMetadata is the meta-properties of a given entity.
type EntityMetadata struct {
	// Hash is the unique identifier for a given entity; guaranteed to be unique within a contract.
	Hash uint32
	// Index is the identifier for a given entity found within the mobile contract tables.
	Index int32
	// Redacted is true if there is an entity, but the Bungie API is not allowed to show it.
	Redacted bool
}
