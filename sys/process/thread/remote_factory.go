package thread

type RemoteFactory interface {
	New(process interface{}, entryAddress uintptr, argsAddress *uintptr) Remote
}
