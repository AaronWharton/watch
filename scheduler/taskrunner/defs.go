package taskrunner

const (
	ReadyToDispatch = "d"
	ReadtToExecute = "e"
	Close = "c"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error