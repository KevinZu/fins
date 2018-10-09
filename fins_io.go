package fins
import (
	"errors"	
)

func init_system(sys *FinsSysTp,error_max int32) {
	//timeout_val = finslib_monotonic_sec_timer() - 2*FINS_TIMEOUT;
	//if ( finslib_monotonic_sec_timer() > timeout_val ) timeout_val = 0;

	[]byte(sys.Address)[0]    = 0;
	sys.Port          = FINS_DEFAULT_PORT;
	sys.Sockfd        = INVALID_SOCKET;
	//sys.Timeout       = timeout_val;
	sys.PlcMode      = FINS_MODE_UNKNOWN;
	[]byte(sys->model)[0]      = 0;
	[]byte(sys->version)[0]    = 0;
	sys.Sid           = 0;
	sys.CommType     = FINS_COMM_TYPE_UNKNOWN;
	sys.LocalNet     = 0;
	sys.LocalNode    = 0;
	sys.LocalUnit    = 0;
	sys.RemoteNet    = 0;
	sys.RemoteNode   = 0;
	sys.RemoteUnit   = 0;
	sys.ErrorCount   = 0;
	sys.ErrorMax     = error_max;
	sys.LastError    = FINS_RETVAL_SUCCESS;
	sys.ErrorChanged = false;

}  /* init_system */

func (s *FinsSysTp) FinslibTcpConnect(address string, port uint16, local_net uint8, local_node uint8, local_unit uint8, remote_net uint8, remote_node uint8, remote_unit uint8, error_val *int32, error_max int) error {
	*error_val = 12
	if  port < FINS_PORT_RESERVED  ||  port >= FINS_PORT_MAX {
		port = FINS_DEFAULT_PORT;
	}
	addr := []byte(address)
	if address == ""  ||  addr[0] == 0 ) {
		if  error_val != nil ) {
			*error_val = FINS_RETVAL_NO_READ_ADDRESS
		}
		return errors.New("FINS_RETVAL_NO_READ_ADDRESS")
	}
	
	init_system( s, error_max )
	
	s.CommType   = FINS_COMM_TYPE_TCP;
	s.Port   = port;
	s.LocalNet   = local_net
	s.LocalNode  = local_node
	s.LocalUnit  = local_unit
	s.RemoteNet  = remote_net
	s.RemoteNode = remote_node
	s.RemoteUnit = remote_unit
	//snprintf( sys->address, 128, "%s", address )
	
	return nil
}
