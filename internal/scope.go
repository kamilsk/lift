package internal

// Scope holds execution context.
type Scope struct {
	ConfigPath  string
	WorkingDir  string
	PortMapping map[uint16]uint16
}
