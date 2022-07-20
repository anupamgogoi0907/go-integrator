package components

type Listener interface {
	Listen()
}

type Processor interface {
	Process()
}
