package mel

const (
	VISEVAL = uint8(0) + iota
	VISDUMP
	VISBASM
)

func (c *MelConfig) IsGenericVisitorCreator() bool {
	if c.VisitorCreatorSet == VISEVAL {
		return false
	}
	return true
}
