package taskrunner

const (
	VideoPath = "./videos/"
	ReadyToDispatch = "d"
	ReadtToExecute = "e"
	Close = "c"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error