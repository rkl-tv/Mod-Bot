package thread

type MockedRemoteFactory struct {
	NewFunc func(process interface{}, entryAddress uintptr, argsAddress *uintptr) Remote
}

func (f *MockedRemoteFactory) New(process interface{}, entryAddress uintptr, argsAddress *uintptr) Remote {
	if f.NewFunc != nil {
		return f.NewFunc(process, entryAddress, argsAddress)
	}

	return nil
}
