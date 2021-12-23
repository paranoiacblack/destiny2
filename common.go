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
