package fins

import (
	//	"bytes"
	//	"encoding/binary"
	"errors"
	"fmt"
	//	"net"
	//	"sync"
	//	"sync/atomic"
	"time"
)

////////////////////////////////////////////////////// connect ////////////////////////////////////////////////
func init_system(sys *FinsSysTp, error_max int32) {
	//timeout_val = finslib_monotonic_sec_timer() - 2*FINS_TIMEOUT;
	time_val := time.Now().Unix() - 2*FINS_TIMEOUT
	//if ( finslib_monotonic_sec_timer() > timeout_val ) timeout_val = 0;
	if time.Now().Unix() > time_val {
		time_val = 0
	}

	sys.Address = append(sys.Address, 0)
	//	fmt.Printf("len=%d cap=%d slice=%v\n", len(sys.Address), cap(sys.Address), sys.Address)
	sys.Port = FINS_DEFAULT_PORT
	sys.SocketFd = INVALID_SOCKET
	//sys.Timeout       = timeout_val;
	sys.PlcMode = FINS_MODE_UNKNOWN
	sys.Model = append(sys.Model, 0)
	sys.Version = append(sys.Version, 0)
	sys.Sid = 0
	sys.CommType = FINS_COMM_TYPE_UNKNOWN
	sys.LocalNet = 0
	sys.LocalNode = 0
	sys.LocalUnit = 0
	sys.RemoteNet = 0
	sys.RemoteNode = 0
	sys.RemoteUnit = 0
	sys.ErrorCount = 0
	sys.ErrorMax = error_max
	sys.LastError = FINS_RETVAL_SUCCESS
	sys.ErrorChanged = false
	sys.Timeout = time_val
	sys.CliGroup = GetClientGroup()

} /* init_system */

func (s *FinsSysTp) FinslibTcpConnect(address string, port uint16, local_net uint8, local_node uint8, local_unit uint8, remote_net uint8, remote_node uint8, remote_unit uint8, error_val *int32, error_max int32) error {
	//*error_val = 12
	if time.Now().Unix() < s.Timeout+FINS_TIMEOUT && s.Timeout > 0 {

		if error_val != nil {
			*error_val = FINS_RETVAL_TRY_LATER
		}

		fmt.Println("===== FINS_RETVAL_TRY_LATER! ========")

		return errors.New("FINS_RETVAL_TRY_LATER")
	}

	if port < FINS_PORT_RESERVED || port >= FINS_PORT_MAX {
		port = FINS_DEFAULT_PORT
	}

	addr := []byte(address)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(addr), cap(addr), addr)

	if address == "" || addr[0] == 0 {
		if error_val != nil {
			*error_val = FINS_RETVAL_NO_READ_ADDRESS
		}
		fmt.Println("===== FINS_RETVAL_NO_READ_ADDRESS! ========")
		return errors.New("FINS_RETVAL_NO_READ_ADDRESS")
	}

	init_system(s, error_max)

	s.CommType = FINS_COMM_TYPE_TCP
	s.Port = port
	s.LocalNet = local_net
	s.LocalNode = local_node
	s.LocalUnit = local_unit
	s.RemoteNet = remote_net
	s.RemoteNode = remote_node
	s.RemoteUnit = remote_unit

	s.Address = make([]byte, len(addr))
	copy(s.Address, addr)

	strPort := fmt.Sprintf("%d", port)
	addrInfo := address + ":" + strPort
	fmt.Println("addrinfo: ", addrInfo)
	//go s.Dial("tcp", addrInfo)

	return nil
}
