package grabber

import "ModBot/sys/process"

type MockedGrabber struct {
	GrabFunc func(processName string) (*process.Process, error)
}

func (g *MockedGrabber) Grab(processName string) (*process.Process, error) {
	if g.GrabFunc != nil {
		return g.GrabFunc(processName)
	}

	return process.NewProcess(1, "dummy", process.ModuleList{}), nil
}
