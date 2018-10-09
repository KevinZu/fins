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

func (s *FinsSysTp) finslib_tcp_connect(address string, port uint16, local_net uint8, local_node uint8, local_unit uint8, remote_net uint8, remote_node uint8, remote_unit uint8, error_val *int32, error_max int) *FinsSysTp {
	tp := new(FinsSysTp)

	return tp
}
