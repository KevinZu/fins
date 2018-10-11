package fins

const (
	FINS_ICF = 0  /* Information Control Field				*/
	FINS_RSV = 1  /* Reserved						*/
	FINS_GCT = 2  /* Gateway Counter. Init op 0x07 naar CPU v2.0 of 0x02	*/
	FINS_DNA = 3  /* Destination Network Address (0..127) 0 = lokaal	*/
	FINS_DA1 = 4  /* Destination Node Address (0..254) 0 = lokaal		*/
	FINS_DA2 = 5  /* Destination Unit Address (0..31) 0 = CPU unit	*/
	FINS_SNA = 6  /* Source Network Address (0..127) 0 = lokaal		*/
	FINS_SA1 = 7  /* Source Node Address (0..254) 0 = intern in PLC	*/
	FINS_SA2 = 8  /* Source Unit Address (0..31) 0 = CPU unit		*/
	FINS_SID = 9  /* Service ID, uniek nummer 0..FF per commando		*/
	FINS_MRC = 10 /* Main Request Code					*/
	FINS_SRC = 11 /* Sub Request Code					*/
)

const (
	INVALID_SOCKET = -1
)

const (
	FINS_MRES = 0
	FINS_SRES = 1
)

const (
	/********************************************************/
	/*							*/
	FINS_HEADER_LEN     = 12   /* Length of a FINS header				*/
	FINS_BODY_LEN       = 2000 /* Maximum length of a FINS body			*/
	FINS_MAX_TCP_HEADER = 32   /* Maximum length of a FINS/TCP header			*/
	/*							*/
	/********************************************************/
)

const (
	FINS_PORT_MIN      = 0
	FINS_PORT_RESERVED = 1024
	FINS_PORT_MAX      = 65535

	FINS_TIMEOUT = 60

	FINS_DEFAULT_PORT = 9600 /* Default port for FINS TCP and UDP communications	*/
)

const (
	FINS_MODE_UNKNOWN = 0 /* PLC communication mode unknown			*/
	FINS_MODE_CV      = 1 /* PLC communicates like a CV PLC			*/
	FINS_MODE_CS      = 2 /* PLC communicates like a CS/CJ PLC			*/
)

const (
	FINS_CPU_MODE_PROGRAM = 0 /* The CPU is in program mode				*/
	FINS_CPU_MODE_MONITOR = 2 /* The CPU is in monitor mode				*/
	FINS_CPU_MODE_RUN     = 4 /* The CPU is in run mode				*/
)

const (
	FI_RD   = 0x01
	FI_WR   = 0x02
	FI_FILL = 0x04
	FI_MRD  = 0x08
	FI_TRS  = 0x10
	FI_TRD  = 0x20
	FI_FRC  = 0x40
)

const (
	FINS_MAX_READ_WORDS_SYSWAY      = 269
	FINS_MAX_READ_WORDS_ETHERNET    = 999
	FINS_MAX_READ_WORDS_CLINK       = 999
	FINS_MAX_READ_WORDS_SYSMAC_LINK = 269
	FINS_MAX_READ_WORDS_DEVICENET   = 269

	FINS_MAX_WRITE_WORDS_SYSWAY      = 267
	FINS_MAX_WRITE_WORDS_ETHERNET    = 996
	FINS_MAX_WRITE_WORDS_CLINK       = 996
	FINS_MAX_WRITE_WORDS_SYSMAC_LINK = 267
	FINS_MAX_WRITE_WORDS_DEVICENET   = 267
)

const (
	FINS_COMM_TYPE_UNKNOWN = 0x00 /* No communication protocol has been selected		*/
	FINS_COMM_TYPE_TCP     = 0x01 /* The communication protocol is FINS/TCP		*/
	FINS_COMM_TYPE_UDP     = 0x02 /* The communication protocol is FINS/UDP		*/
)

const (
	FINS_DATA_TYPE_INT16       = 1  /* 16 bit signed integer				*/
	FINS_DATA_TYPE_INT32       = 2  /* 32 bit signed integer				*/
	FINS_DATA_TYPE_UINT16      = 3  /* 16 bit unsigned integer				*/
	FINS_DATA_TYPE_UINT32      = 4  /* 32 bit unsigned integer				*/
	FINS_DATA_TYPE_BCD16       = 5  /* Unsigned 16 bit BCD in the range 0..9999		*/
	FINS_DATA_TYPE_BCD32       = 6  /* Unsigned 32 bit BCD in the range 0..99999999		*/
	FINS_DATA_TYPE_SBCD16_0    = 7  /* Signed 16 bit BCD in the range -999..999		*/
	FINS_DATA_TYPE_SBCD16_1    = 8  /* Signed 16 bit BCD in the range -7999..7999		*/
	FINS_DATA_TYPE_SBCD16_2    = 9  /* Signed 16 bit BCD in the range -999..9999		*/
	FINS_DATA_TYPE_SBCD16_3    = 10 /* Signed 16 bit BCD in the range -1999..9999		*/
	FINS_DATA_TYPE_SBCD32_0    = 11 /* Signed 32 bit BCD in the range -9999999..9999999	*/
	FINS_DATA_TYPE_SBCD32_1    = 12 /* Signed 32 bit BCD in the range -79999999..79999999	*/
	FINS_DATA_TYPE_SBCD32_2    = 13 /* Signed 32 bit BCD in the range -9999999..99999999	*/
	FINS_DATA_TYPE_SBCD32_3    = 14 /* Signed 32 bit BCD in the range -19999999..99999999	*/
	FINS_DATA_TYPE_FLOAT       = 15 /* 32 bit floating point value				*/
	FINS_DATA_TYPE_DOUBLE      = 16 /* 64 bit floating point value				*/
	FINS_DATA_TYPE_BIT         = 17 /* Single bit						*/
	FINS_DATA_TYPE_BIT_FORCED  = 18 /* Single bit with forced status			*/
	FINS_DATA_TYPE_WORD_FORCED = 19 /* 16 bit word with for each bit the forced status	*/
)

const (
	FINS_MEMORY_CARD_NONE  = 0
	FINS_MEMORY_CARD_FLASH = 4

	FINS_MSG_0   = 0x01
	FINS_MSG_1   = 0x02
	FINS_MSG_2   = 0x04
	FINS_MSG_3   = 0x08
	FINS_MSG_4   = 0x10
	FINS_MSG_5   = 0x20
	FINS_MSG_6   = 0x40
	FINS_MSG_7   = 0x80
	FINS_MSG_ALL = 0xFF
)

const (
	FINS_FORCE_RESET          = 0x0000 /* Force the bit and reset it				*/
	FINS_FORCE_SET            = 0x0001 /* Force the bit and set it				*/
	FINS_FORCE_RELEASE_TO_OFF = 0x8000 /* Release the force and reset the bit			*/
	FINS_FORCE_RELEASE_TO_ON  = 0x8001 /* Release the force and set the bit			*/
	FINS_FORCE_RELEASE        = 0xFFFF /* Release the force					*/
)

const (
	FINS_WRITE_MODE_NEW_NOT_OVERWRITE = 0x0000
	FINS_WRITE_MODE_NEW_OVERWRITE     = 0x0001
	FINS_WRITE_MODE_ADD_DATA          = 0x0002
	FINS_WRITE_MODE_OVERWRITE         = 0x0003
)

const (
	FINS_PARAM_AREA_ALL                   = 0x8000 /* Pseudo value for all parameter areas			*/
	FINS_PARAM_AREA_PLC_SETUP             = 0x8010 /* Parameter area for the PLC setup			*/
	FINS_PARAM_AREA_IO_TABLE_REGISTRATION = 0x8012 /* Parameter area for the I/O table registration	*/
	FINS_PARAM_AREA_ROUTING_TABLE         = 0x8013 /* Parameter area for the routing table			*/
	FINS_PARAM_AREA_CPU_BUS_UNIT_SETUP    = 0x8002 /* Parameter area for the CPU Bus unit setup		*/
)

const (
	FINS_DISK_MEMORY_CARD    = 0x8000 /* The disk is the flash memory card			*/
	FINS_DISK_EM_FILE_MEMORY = 0x8001 /* The disk is the file system in EM File memory	*/
)

const (
	FINS_RETVAL_ERRNO_BASE = 0xC000 /* All higher error numbers are errno.h values		*/
	/*							*/
	FINS_RETVAL_SUCCESS            = 0x0000 /* Execution was successful				*/
	FINS_RETVAL_NOT_INITIALIZED    = 0x8001 /* The connection with the PLC was not initialized	*/
	FINS_RETVAL_NOT_CONNECTED      = 0x8002 /* There is no connection with the remote PLC		*/
	FINS_RETVAL_OUT_OF_MEMORY      = 0x8003 /* There was not enough free memory for the action	*/
	FINS_RETVAL_SUCCESS_LAST_DATA  = 0x8004 /* Execution successful and last data delivered		*/
	FINS_RETVAL_INVALID_IP_ADDRESS = 0x8005 /* The IP address passed to inet_pton is invalid	*/
	FINS_RETVAL_MAX_ERROR_COUNT    = 0x8006 /* The connection was closed after reaching max errors	*/
	FINS_RETVAL_SYNC_ERROR         = 0x8007 /* Synchronization error. Some packets probably lost	*/
	/*							*/
	FINS_RETVAL_NO_READ_ADDRESS  = 0x8101 /* No read address in the remote PLC was specified	*/
	FINS_RETVAL_NO_WRITE_ADDRESS = 0x8102 /* No write address in the remote PLC was specified	*/
	FINS_RETVAL_NO_DATA_BLOCK    = 0x8103 /* No local data memory block was provided		*/
	/*							*/
	FINS_RETVAL_INVALID_READ_ADDRESS  = 0x8201 /* An invalid read address string was specified		*/
	FINS_RETVAL_INVALID_WRITE_ADDRESS = 0x8202 /* An invalid write address string was specified	*/
	/*							*/
	FINS_RETVAL_INVALID_READ_AREA  = 0x8301 /* No read area associated with the address		*/
	FINS_RETVAL_INVALID_WRITE_AREA = 0x8302 /* No write area associated with the address		*/
	FINS_RETVAL_INVALID_FILL_AREA  = 0x8303 /* No fill area associated with the address		*/
	/*							*/
	FINS_RETVAL_INVALID_PARAMETER_AREA = 0x8401 /* The parameter area is invalid			*/
	/*							*/
	FINS_RETVAL_INVALID_DATE = 0x8501 /* The provided date is not valid			*/
	/*							*/
	FINS_RETVAL_INVALID_DISK     = 0x8601 /* An invalid disk was specified			*/
	FINS_RETVAL_INVALID_PATH     = 0x8602 /* An invalid path on a disk was specified		*/
	FINS_RETVAL_INVALID_FILENAME = 0x8603 /* An invalid filename was specified			*/
	/*							*/
	FINS_RETVAL_NO_COMMAND                 = 0x8701 /* No command specified when executing a function	*/
	FINS_RETVAL_NO_COMMAND_LENGTH          = 0x8702 /* No command length specified when executing a function*/
	FINS_RETVAL_BODY_TOO_SHORT             = 0x8703 /* Command body length too short			*/
	FINS_RETVAL_BODY_TOO_LONG              = 0x8704 /* The FINS body is longer than allowed			*/
	FINS_RETVAL_HEADER_SEND_ERROR          = 0x8705 /* Error sending complete header			*/
	FINS_RETVAL_COMMAND_SEND_ERROR         = 0x8706 /* Error sending complete command			*/
	FINS_RETVAL_RESPONSE_INCOMPLETE        = 0x8707 /* Response frame is shorter than expected		*/
	FINS_RETVAL_ILLEGAL_FINS_COMMAND       = 0x870B /* Illegal FINS command					*/
	FINS_RETVAL_RESPONSE_HEADER_INCOMPLETE = 0x870C /* The received response header is incomplete		*/
	FINS_RETVAL_INVALID_FORCE_COMMAND      = 0x870D /* An invalid FORCE mode was specified			*/
	/*							*/
	FINS_RETVAL_TRY_LATER = 0x8801 /* Please try again later				*/
	/*							*/
	FINS_RETVAL_CLOSED_BY_REMOTE          = 0x8900 /* TCP connection closed by remote node without error	*/
	FINS_RETVAL_NO_FINS_HEADER            = 0x8901 /* First 4 characters of TCP header are not "FINS"	*/
	FINS_RETVAL_DATA_LENGTH_TOO_LONG      = 0x8902 /* TCP connection data length too long			*/
	FINS_RETVAL_COMMAND_NOT_SUPPORTED     = 0x8903 /* TCP connection command not supported			*/
	FINS_RETVAL_ALL_CONNECTIONS_IN_USE    = 0x8904 /* All TCP connections are in use			*/
	FINS_RETVAL_NODE_ALREADY_CONNECTED    = 0x8905 /* Node is already connected				*/
	FINS_RETVAL_NODE_IP_PROTECTED         = 0x8906 /* IP address of client not in allowed IP adres list	*/
	FINS_RETVAL_CLIENT_NODE_OUT_OF_RANGE  = 0x8907 /* TCP the client node address is out of range		*/
	FINS_RETVAL_SAME_NODE_ADDRESS         = 0x8908 /* TCP client and server have the same node address	*/
	FINS_RETVAL_NO_NODE_ADDRESS_AVAILABLE = 0x8909 /* TCP connection no node address available		*/
	/*							*/
	FINS_RETVAL_WSA_UNRECOGNIZED_ERROR     = 0x8A00 /* Windows WSA returned an unrecognized error code	*/
	FINS_RETVAL_WSA_NOT_INITIALIZED        = 0x8A01 /* Windows WSA was not properly initialized		*/
	FINS_RETVAL_WSA_E_NET_DOWN             = 0x8A02 /* Windows WSA the network subsystem or provided failed	*/
	FINS_RETVAL_WSA_E_AF_NO_SUPPORT        = 0x8A03 /* Windows WSA the address familiy is not supported	*/
	FINS_RETVAL_WSA_E_IN_PROGRESS          = 0x8A04 /* Windows WSA a blocking socket 1.1 call is in progres	*/
	FINS_RETVAL_WSA_E_MFILE                = 0x8A05 /* Windows WSA no more socket descriptors available	*/
	FINS_RETVAL_WSA_E_INVAL                = 0x8A06 /* Windows WSA Invalid argument supplied		*/
	FINS_RETVAL_WSA_E_INVALID_PROVIDER     = 0x8A07 /* Windows WSA Server provider function invalid		*/
	FINS_RETVAL_WSA_E_INVALID_PROCTABLE    = 0x8A08 /* Windows WSA Invalid procedure table			*/
	FINS_RETVAL_WSA_E_NOBUFS               = 0x8A09 /* Windows WSA No buffer space available		*/
	FINS_RETVAL_WSA_E_PROTO_NO_SUPPORT     = 0x8A0A /* Windows WSA the protocol is not supported		*/
	FINS_RETVAL_WSA_E_PROTO_TYPE           = 0x8A0B /* Windows WSA Wrong protocol type for this socket	*/
	FINS_RETVAL_WSA_E_PROVIDER_FAILED_INIT = 0x8A0C /* Windows WSA Provider failed initialization		*/
	FINS_RETVAL_WSA_E_SOCKT_NO_SUPPORT     = 0x8A0D /* Windows WSA The specified socket type not supported	*/
	FINS_RETVAL_WSA_SYS_NOT_READY          = 0x8A0E /* Windows WSA The network subsystem is not ready	*/
	FINS_RETVAL_WSA_VER_NOT_SUPPORTED      = 0x8A0F /* Windows WSA The socket version is not supported	*/
	FINS_RETVAL_WSA_E_PROC_LIM             = 0x8A10 /* Windows WSA Process number limit reached		*/
	FINS_RETVAL_WSA_E_FAULT                = 0x8A11 /* Windows WSA The parameter is not valid		*/
	FINS_RETVAL_WSA_E_NET_RESET            = 0x8A12 /* Windows WSA Connection timeout during Keep Alive	*/
	FINS_RETVAL_WSA_E_NO_PROTO_OPT         = 0x8A13 /* Windows WSA Unsupported option for socket		*/
	FINS_RETVAL_WSA_E_NOT_CONN             = 0x8A14 /* Windows WSA Connection reset during Keep Alive	*/
	FINS_RETVAL_WSA_E_NOT_SOCK             = 0x8A15 /* Windows WSA The descriptor is not a socket		*/
	FINS_RETVAL_WSA_E_ACCES                = 0x8A16 /* Windows WSA Socket access violation			*/
	FINS_RETVAL_WSA_E_ADDR_IN_USE          = 0x8A17 /* Windows WSA The address is already in use		*/
	FINS_RETVAL_WSA_E_ADDR_NOT_AVAIL       = 0x8A18 /* Windows WSA The address is not available		*/
	FINS_RETVAL_WSA_E_INTR                 = 0x8A19 /* Windows WSA The blocking 1.1 call was cancelled	*/
	FINS_RETVAL_WSA_E_ALREADY              = 0x8A1A /* Windows WSA Non blocking call already in progress	*/
	FINS_RETVAL_WSA_E_CONN_REFUSED         = 0x8A1B /* Windows WSA The connection was refused		*/
	FINS_RETVAL_WSA_E_IS_CONN              = 0x8A1C /* Windows WSA Socket is already connected		*/
	FINS_RETVAL_WSA_E_NET_UNREACH          = 0x8A1D /* Windows WSA Network is unreacheable			*/
	FINS_RETVAL_WSA_E_HOST_UNREACH         = 0x8A1E /* Windows WSA Host is unreacheable			*/
	FINS_RETVAL_WSA_E_TIMED_OUT            = 0x8A1F /* Windows WSA The connection timed out			*/
	FINS_RETVAL_WSA_E_WOULD_BLOCK          = 0x8A20 /* Windows WSA Non-blocking connection would block	*/
	/*							*/
	FINS_RETVAL_CANCELED = 0x0001 /* End code 0x0001 The service was canceled		*/
	/*							*/
	FINS_RETVAL_LOCAL_NODE_NOT_IN_NETWORK  = 0x0101 /* End code 0x0101 Local node is not in network		*/
	FINS_RETVAL_LOCAL_TOKEN_TIMEOUT        = 0x0102 /* End code 0x0102 Local node token timeout		*/
	FINS_RETVAL_LOCAL_RETRIES_FAILED       = 0x0103 /* End code 0x0103 Local node retries failed		*/
	FINS_RETVAL_LOCAL_TOO_MANY_SEND_FRAMES = 0x0104 /* End code 0x0104 Local node too many send frames	*/
	FINS_RETVAL_LOCAL_ADDRESS_RANGE_ERROR  = 0x0105 /* End code 0x0105 Local node address range error	*/
	FINS_RETVAL_LOCAL_ADDRESS_DUPLICATION  = 0x0106 /* End code 0x0106 Local node address duplication	*/
	/*							*/
	FINS_RETVAL_DEST_NOT_IN_NETWORK     = 0x0201 /* End code 0x0201 Destination is not in network	*/
	FINS_RETVAL_DEST_UNIT_MISSING       = 0x0202 /* End code 0x0202 Destination unit missing		*/
	FINS_RETVAL_DEST_THIRD_NODE_MISSING = 0x0203 /* End code 0x0203 Destination third node missing	*/
	FINS_RETVAL_DEST_NODE_BUSY          = 0x0204 /* End code 0x0204 Destination node is busy		*/
	FINS_RETVAL_DEST_TIMEOUT            = 0x0205 /* End code 0x0205 Destination response timed out	*/
	/*							*/
	FINS_RETVAL_CONTR_COMM_ERROR        = 0x0301 /* End code 0x0301 Communications controller error	*/
	FINS_RETVAL_CONTR_CPU_UNIT_ERROR    = 0x0302 /* End code 0x0302 CPU Unit error			*/
	FINS_RETVAL_CONTR_BOARD_ERROR       = 0x0303 /* End code 0x0303 Controller board error		*/
	FINS_RETVAL_CONTR_UNIT_NUMBER_ERROR = 0x0304 /* End code 0x0304 Unit number error			*/
	/*							*/
	FINS_RETVAL_UNSUPPORTED_COMMAND = 0x0401 /* End code 0x0401 Undefined command			*/
	FINS_RETVAL_UNSUPPORTED_VERSION = 0x0402 /* End code 0x0402 Not supported by model/version	*/
	/*							*/
	FINS_RETVAL_ROUTING_ADDRESS_ERROR   = 0x0501 /* End code 0x0501 Routing destination address error	*/
	FINS_RETVAL_ROUTING_NO_TABLES       = 0x0502 /* End code 0x0502 No routing tables available		*/
	FINS_RETVAL_ROUTING_TABLE_ERROR     = 0x0503 /* End code 0x0503 Routing table error			*/
	FINS_RETVAL_ROUTING_TOO_MANY_RELAYS = 0x0504 /* End code 0x0504 Too many relays			*/
	/*							*/
	FINS_RETVAL_COMMAND_TOO_LONG         = 0x1001 /* End code 0x1001 Command too long			*/
	FINS_RETVAL_COMMAND_TOO_SHORT        = 0x1002 /* End code 0x1002 Command too short			*/
	FINS_RETVAL_COMMAND_ELEMENT_MISMATCH = 0x1003 /* End code 0x1003 Elements/data don't match		*/
	FINS_RETVAL_COMMAND_FORMAT_ERROR     = 0x1004 /* End code 0x1004 Command format error			*/
	FINS_RETVAL_COMMAND_HEADER_ERROR     = 0x1005 /* End code 0x1005 Command header error			*/
	/*							*/
	FINS_RETVAL_PARAM_AREA_MISSING        = 0x1101 /* End code 0x1101 Parameter area classification missing*/
	FINS_RETVAL_PARAM_ACCESS_SIZE_ERROR   = 0x1102 /* End code 0x1102 Parameter access size wrong		*/
	FINS_RETVAL_PARAM_START_ADDRESS_ERROR = 0x1103 /* End code 0x1103 Start address out of range		*/
	FINS_RETVAL_PARAM_END_ADDRESS_ERROR   = 0x1104 /* End code 0x1104 End address out of range		*/
	FINS_RETVAL_PARAM_PROGRAM_MISSING     = 0x1106 /* End code 0x1106 Program number is missing		*/
	FINS_RETVAL_PARAM_RELATIONAL_ERROR    = 0x1109 /* End code 0x1109 Parameter relational error		*/
	FINS_RETVAL_PARAM_DUPLICATE_ACCESS    = 0x110A /* End code 0x110A Duplicate data access		*/
	FINS_RETVAL_PARAM_RESPONSE_TOO_LONG   = 0x110B /* End code 0x110B Response too long			*/
	FINS_RETVAL_PARAM_PARAMETER_ERROR     = 0x110C /* End code 0x110C Parameter error			*/
	/*							*/
	FINS_RETVAL_RD_ERR_PROTECTED       = 0x2002 /* End code 0x2002 The program area is protected	*/
	FINS_RETVAL_RD_ERR_TABLE_MISSING   = 0x2003 /* End code 0x2003 The table is not existing		*/
	FINS_RETVAL_RD_ERR_DATA_MISSING    = 0x2004 /* End code 0x2004 The search data does not exist	*/
	FINS_RETVAL_RD_ERR_PROGRAM_MISSING = 0x2005 /* End code 0x2005 The program does not exist		*/
	FINS_RETVAL_RD_ERR_FILE_MISSING    = 0x2006 /* End code 0x2006 The file does not exist		*/
	FINS_RETVAL_RD_ERR_DATA_MISMATCH   = 0x2007 /* End code 0x2007 Data comparison failed		*/
	/*							*/
	FINS_RETVAL_WR_ERR_READ_ONLY       = 0x2101 /* End code 0x2101 The specified area is read-only	*/
	FINS_RETVAL_WR_ERR_PROTECTED       = 0x2102 /* End code 0x2102 The program area is protected	*/
	FINS_RETVAL_WR_ERR_CANNOT_REGISTER = 0x2103 /* End code 0x2103 Cannot register file			*/
	FINS_RETVAL_WR_ERR_PROGRAM_MISSING = 0x2105 /* End code 0x2105 Program number is not valid		*/
	FINS_RETVAL_WR_ERR_FILE_MISSING    = 0x2106 /* End code 0x2106 File does not exist			*/
	FINS_RETVAL_WR_ERR_FILE_EXISTS     = 0x2107 /* End code 0x2107 The file already exists		*/
	FINS_RETVAL_WR_ERR_CANNOT_CHANGE   = 0x2108 /* End code 0x2108 Cannot change the data		*/
	/*							*/
	FINS_RETVAL_MODE_NOT_DURING_EXECUTION = 0x2201 /* End code 0x2201 Not possible during execution	*/
	FINS_RETVAL_MODE_NOT_DURING_RUN       = 0x2202 /* End code 0x2202 Not possible while running		*/
	FINS_RETVAL_MODE_IS_PROGRAM           = 0x2203 /* End code 0x2203 Not possible in program mode		*/
	FINS_RETVAL_MODE_IS_DEBUG             = 0x2204 /* End code 0x2204 Not possible in debug mode		*/
	FINS_RETVAL_MODE_IS_MONITOR           = 0x2205 /* End code 0x2205 Not possible in monitor mode		*/
	FINS_RETVAL_MODE_IS_RUN               = 0x2206 /* End code 0x2206 Not possible in run mode		*/
	FINS_RETVAL_MODE_NODE_NOT_POLLING     = 0x2207 /* End code 0x2207 Specified node not in polling mode	*/
	FINS_RETVAL_MODE_NO_STEP              = 0x2208 /* End code 0x2208 Step cannot be executed		*/
	/*							*/
	FINS_RETVAL_DEVICE_FILE_MISSING   = 0x2301 /* End code 0x2301 File device missing			*/
	FINS_RETVAL_DEVICE_MEMORY_MISSING = 0x2302 /* End code 0x2302 There is no file memory		*/
	FINS_RETVAL_DEVICE_CLOCK_MISSING  = 0x2303 /* End code 0x2303 There is no clock			*/
	/*							*/
	FINS_RETVAL_DATALINK_TABLE_MISSING = 0x2401 /* End code 0x2401 Data link table missing or corrupt	*/
	/*							*/
	FINS_RETVAL_UNIT_MEMORY_CONTENT_ERROR   = 0x2502 /* End code 0x2502 Memory content error			*/
	FINS_RETVAL_UNIT_IO_SETTING_ERROR       = 0x2503 /* End code 0x2503 I/O setting error			*/
	FINS_RETVAL_UNIT_TOO_MANY_IO_POINTS     = 0x2504 /* End code 0x2504 Too many I/O points registered	*/
	FINS_RETVAL_UNIT_CPU_BUS_ERROR          = 0x2505 /* End code 0x2505 CPU bus line error			*/
	FINS_RETVAL_UNIT_IO_DUPLICATION         = 0x2506 /* End code 0x2506 Duplicate I/O address		*/
	FINS_RETVAL_UNIT_IO_BUS_ERROR           = 0x2507 /* End code 0x2507 I/O bus error			*/
	FINS_RETVAL_UNIT_SYSMAC_BUS2_ERROR      = 0x2509 /* End code 0x2509 Sysmac bus/2 error			*/
	FINS_RETVAL_UNIT_CPU_BUS_UNIT_ERROR     = 0x250A /* End code 0x250A CPU bus unit error			*/
	FINS_RETVAL_UNIT_SYSMAC_BUS_DUPLICATION = 0x250D /* End code 0x250D Same word is used more than once	*/
	FINS_RETVAL_UNIT_MEMORY_ERROR           = 0x250F /* End code 0x250F Memory error in internal memory	*/
	FINS_RETVAL_UNIT_SYSMAC_BUS_TERMINATOR  = 0x2510 /* End code 0x2510 Sysmac bus terminator missing	*/
	/*							*/
	FINS_RETVAL_COMMAND_NO_PROTECTION       = 0x2601 /* End code 0x2601 The specified area is not protected	*/
	FINS_RETVAL_COMMAND_WRONG_PASSWORD      = 0x2602 /* End code 0x2602 Wrong password specified		*/
	FINS_RETVAL_COMMAND_PROTECTED           = 0x2604 /* End code 0x2604 The specified area is protected	*/
	FINS_RETVAL_COMMAND_SERVICE_EXECUTING   = 0x2605 /* End code 0x2605 Service is already executing		*/
	FINS_RETVAL_COMMAND_SERVICE_STOPPED     = 0x2606 /* End code 0x2606 The service is stopped		*/
	FINS_RETVAL_COMMAND_NO_EXECUTION_RIGHT  = 0x2607 /* End code 0x2607 No execution right			*/
	FINS_RETVAL_COMMAND_SETTINGS_INCOMPLETE = 0x2608 /* End code 0x2608 The settings are not complete	*/
	FINS_RETVAL_COMMAND_ITEMS_NOT_SET       = 0x2609 /* End code 0x2609 Necessary items not set		*/
	FINS_RETVAL_COMMAND_ALREADY_DEFINED     = 0x260A /* End code 0x260A Number already defined		*/
	FINS_RETVAL_COMMAND_ERROR_WONT_CLEAR    = 0x260B /* End code 0x260B Error will not clear			*/
	/*							*/
	FINS_RETVAL_ACCESS_NO_RIGHTS = 0x3001 /* End code 0x3001 No access right			*/
	/*							*/
	FINS_RETVAL_ABORTED = 0x4001 /* End code 0x3001 Service aborted			*/
)

type fins_area_tp struct { /*							*/
	plc_mode  int32  /* CS/CJ or CV mode communication			*/
	name      string /* Text string with the area short code			*/
	bits      int32  /* Number of bits in the data				*/
	length    int32  /* Number of bytes per element				*/
	area      uint8  /* Area code						*/
	low_id    uint32 /* Lowest identificator					*/
	high_id   uint32 /* Highest identificator				*/
	low_addr  uint32 /* Lowest memory address				*/
	high_addr uint32 /* Highest memory address				*/
	access    int32  /* Read or Read/Write access				*/
	force     bool   /* Force status used 					*/
} /*							*/
/********************************************************/
type fins_sys_tp struct {
	address string
	port    uint16
	sockfd  int32
	//time_t		timeout
	error_count   int32
	error_max     int32
	last_error    int32
	error_changed bool
	local_net     uint8
	local_node    uint8
	local_unit    uint8
	remote_net    uint8
	remote_node   uint8
	remote_unit   uint8
	sid           uint8
	comm_type     uint8
	model         string
	version       string
	plc_mode      int32
}

type fins_datetime_tp struct {
	year  int32
	month int32
	day   int32
	hour  int32
	min   int32
	sec   int32
	dow   int32
}

type fins_cycletime_tp struct {
	min uint32
	avg uint32
	max uint32
}

type fins_cpustatus_tp struct {
	message_exists                 []bool
	running                        bool
	flash_writing                  bool
	battery_present                bool
	standby                        bool
	fatal_memory_error             bool
	fatal_io_bus_error             bool
	fatal_duplication_error        bool
	fatal_inner_board_error        bool
	fatal_io_point_overflow        bool
	fatal_io_setting_error         bool
	fatal_program_error            bool
	fatal_cycle_time_over          bool
	fatal_fals_error               bool
	fal_error                      bool
	duplex_error                   bool
	interrupt_task_error           bool
	basic_io_unit_error            bool
	plc_setup_error                bool
	io_verification_error          bool
	inner_board_error              bool
	cpu_bus_unit_error             bool
	special_io_unit_error          bool
	sysmac_bus_error               bool
	battery_error                  bool
	cs1_cpu_bus_unit_setting_error bool
	special_io_unit_setting_error  bool
	run_mode                       uint8
	error_code                     uint16
	error_message                  string
}

type fins_cpudata_tp struct {
	model                   string
	version                 string
	system_block            []uint8
	largest_em_bank         int32
	program_area_size       int32
	iom_size                int32
	number_of_dm_words      int32
	timer_counter_size      int32
	em_non_file_memory_size int32
	memory_card_size        int32
	num_sysmac_bus_masters  int32
	num_racks               int32
	bus_unit_id             []uint16
	bus_unit_present        []bool
	dip_switch              []bool
	memory_card_type        uint8
}

type fins_unitdata_tp struct {
	model string
	unit  uint8
}

type fins_msgdata_tp struct {
	text string
	msg  uint8
}

type fins_nodedata_tp struct {
	network uint8
	node    uint8
	unit    uint8
}

type fins_errordata_tp struct {
	error_code []uint16
	year       int32
	month      int32
	day        int32
	hour       int32
	min        int32
	sec        int32
}

type fins_accessdata_tp struct {
	network      uint8
	node         uint8
	unit         uint8
	command_code uint16
	year         int32
	month        int32
	day          int32
	hour         int32
	min          int32
	sec          int32
}

type fins_diskinfo_tp struct {
	volume_label   string
	total_capacity uint32
	free_capacity  uint32
	total_files    uint32
	year           int32
	month          int32
	day            int32
	hour           int32
	min            int32
	sec            int32
}

type fins_fileinfo_tp struct {
	filename     string
	size         uint32
	year         int32
	month        int32
	day          int32
	hour         int32
	min          int32
	sec          int32
	read_only    bool
	hidden       bool
	system       bool
	volume_label bool
	directory    bool
	archive      bool
}

type fins_address_tp struct {
	name         string
	main_address uint32
	sub_address  uint32
}

type fins_forcebit_tp struct {
	address       string
	force_command uint16
}

type fins_bool struct {
	bit     bool
	b_force bool
}

type fins_multidata_tp struct {
	/*
		    char		address[12]
		    int			type
		    union {
			int16_t		int16
			int32_t		int32
			uint16_t	uint16
			uint32_t	uint32
			float		sfloat
			double		dfloat
			finsBool fins_bool
			struct {
			    uint16_t	word
			    uint16_t	w_force
			}
		    }*/
}

type FinsSysTp struct {
	Address    []byte
	Port       uint16
	SocketFd   int32
	LocalNet   uint8
	LocalNode  uint8
	LocalUnit  uint8
	RemoteNet  uint8
	RemoteNode uint8
	RemoteUnit uint8

	Sid      uint8
	CommType uint8
	Model    []byte
	Version  []byte
	PlcMode  int32

	Timeout int64

	CliGroup *ClientGroup
}
