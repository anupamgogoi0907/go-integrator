package core

import "core-esb/pkg/components"

type Flow struct {
	Listener   components.Listener
	Processors []components.Processor
}

type SubFlow struct {
}
