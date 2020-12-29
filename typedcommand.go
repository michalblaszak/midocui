package midocui

type ttypedCommand int

const (
	typedCommandNone    ttypedCommand = iota // Not a command at all (Alt not pressed)
	typedCommandUnknown                      // In a command typing mode (Alt pressed) but command not recognized
	typedCommandResize
	typedCommandMove
	typedCommandAppQuit
	typedCommandAppMenu
)
