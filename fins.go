package fins

type FinsSysTp struct {
	Address    string
	Port       uint16
	SocketFd   int32
	LocalNet   uint8
	LocalNode  uint8
	LocalUnit  uint8
	RemoteNet  uint8
	RemoteNode uint8
	RemoteUnit uint8
}
