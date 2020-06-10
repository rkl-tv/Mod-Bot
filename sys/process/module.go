package process

type Module struct {
	path        string
	baseAddress uintptr
}

func NewModule(path string, baseAddress uintptr) *Module {
	return &Module{
		path:        path,
		baseAddress: baseAddress,
	}
}

func (m *Module) GetPath() string {
	return m.path
}

func (m *Module) GetBaseAddress() uintptr {
	return m.baseAddress
}
