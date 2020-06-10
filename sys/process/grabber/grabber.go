package grabber

import "ModBot/sys/process"

type Grabber interface {
	Grab(processName string) (*process.Process, error)
}
