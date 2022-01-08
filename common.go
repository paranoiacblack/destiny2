package destiny2

// DisplayProperties represents common display information for entities.
type DisplayProperties struct {
	Description   string
	Name          string
	Icon          string
	IconSequences []IconSequence
	HighResIcon   string
	HasIcon       bool
}

// ProgressionDisplayProperties are common display information for ProgressionEntity structs.
type ProgressionDisplayProperties struct {
	// DisplayUnitsName is a localized string that display how experience is gained.
	DisplayUnitsName string
	DisplayProperties
}

// VendorDisplayProperties are common display information for VendorEntity structs.
type VendorDisplayProperties struct {
	// LargeIcon is an icon that represents the vendor.
	LargeIcon string
	Subtitle  string
	// OriginalIcon is the original icon for the vendor.
	OriginalIcon string
	// RequirementsDisplay are common requirements for interacting with this vendor such as displaying a certain currency.
	RequirementsDisplay []VendorRequirementDisplayEntry
	// SmallTransparentIcon is the icon used for the vendor's waypoint in a given UI.
	SmallTransparentIcon string
	// MapIcon is the icon used in the map overview for the vendor.
	MapIcon string
	// LargeTransparentIcon is the watermark shown for the vendor.
	LargeTransparentIcon string
	DisplayProperties
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
