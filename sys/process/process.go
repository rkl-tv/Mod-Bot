package process

type DWORD uint32
type ModuleList []*Module

type Process struct {
	id      DWORD
	name    string
	modules ModuleList
}

func NewProcess(id DWORD, name string, modules ModuleList) *Process {
	return &Process{
		id:      id,
		name:    name,
		modules: modules,
	}
}

func (p *Process) GetId() DWORD {
	return p.id
}

func (p *Process) GetName() string {
	return p.name
}

func (p *Process) GetModules() ModuleList {
	return p.modules
}
