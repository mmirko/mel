package outsource

type OutSource interface {
	SendGenotype() ([]byte, error)
	ReceiveGenotype([]byte) error
}

type outSourceServer struct {
	waitList chan []byte
}
