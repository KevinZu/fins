package fins

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

////////////////////////////////////////////////////// IO //////////////////////////////////////////////
////////////////////////////// io buffer ////////////////////////////

type Buffer struct {
	b    []byte
	rOff int //读取位置
	wOff int //写入位置
}

//创建Buffer
func NewBuffer() *Buffer {
	return new(Buffer)
}

//获取缓存当前容量
func (b *Buffer) Cap() int {
	return cap(b.b)
}

//当前可读长度
func (b *Buffer) ReadLength() int {
	return b.wOff - b.rOff
}

//重置读取位置
func (b *Buffer) ResetRead() {
	b.rOff = 0
}

//获取读取位置
func (b *Buffer) GetReadPos() int {
	return b.rOff
}

//获取写入位置
func (b *Buffer) GetWritePos() int {
	return b.wOff
}

func (b *Buffer) SetReadPos(pos int) error {
	if pos > b.wOff {
		return errors.New("ResetReadAt is out of error")
	}
	b.rOff = pos
	return nil
}

//重置写入位置
func (b *Buffer) ResetWrite() {
	b.wOff = 0
}

//写入bytes数据
func (b *Buffer) PutBytes(buffer []byte) {
	wPos := b.wOff + len(buffer)
	if wPos > len(b.b) {
		b.b = append(b.b, make([]byte, wPos-len(b.b))...)
	}
	copy(b.b[b.wOff:wPos], buffer)
	b.wOff = wPos
}

//指定位置写入,如果指定写入位置超出了wOff位置,则抛出异常
//如果指定位置已经存在数据并写入数据超出wOff位置则覆盖之前数据，wOff变更最新
//如果指定位置已经存在数据并写入数据没有超出wOff位置则覆盖之前数据，wOff不变
func (b *Buffer) PutBytesAt(pos int, buffer []byte) error {
	willPos := pos + len(buffer)
	if pos > b.wOff {
		return errors.New("pos is out of wOff")
	}
	if willPos > b.wOff {
		copy(b.b[b.wOff:], buffer)
		b.wOff = willPos
	} else {
		copy(b.b[b.wOff:], buffer)
	}
	return nil
}

//将int数据存入缓存
func (b *Buffer) PutInt(i int) {
	x := int32(i)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	b.PutBytes(bytesBuffer.Bytes())
}

//将uint32数据放入内存
func (b *Buffer) PutUint32(i uint32) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, i)
	b.PutBytes(bytesBuffer.Bytes())
}

//将字符串存入buffer
func (b *Buffer) PutString(s string) {
	b.PutBytes([]byte(s))
}

//读取指定位置开始，指定长度的bytes数据
//如果读取数据位置超出了写入数据的位置，则返回错误
func (b *Buffer) ReadBytesAt(pos, length int) ([]byte, error) {
	if pos > b.wOff {
		return nil, errors.New("pos is out of wOff")
	}
	buffer := make([]byte, length)
	if pos+length > b.wOff {
		copy(buffer, b.b[pos:b.wOff])
		b.rOff = b.wOff
	} else {
		p := pos + length
		copy(buffer, b.b[pos:p])
		b.rOff = p
	}

	return buffer, nil
}

func (b *Buffer) ReadBytes(length int) ([]byte, error) {
	rpos := b.rOff + length
	if rpos > b.wOff {
		return nil, errors.New("ReadBytes out off wOff")
	}
	buf := make([]byte, length)
	copy(buf, b.b[b.rOff:rpos])
	b.rOff = rpos
	return buf, nil
}

//读取int
func (b *Buffer) ReadInt() (int, error) {
	rpos := b.rOff + 4
	if rpos > b.wOff {
		return 0, errors.New("ReadInt out off wOff")
	}
	bytesBuffer := bytes.NewBuffer(b.b[b.rOff:rpos])
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	b.rOff = rpos
	return int(x), nil
}

//get uint8
func (b *Buffer) ReadUint8() (uint8, error) {
	rpos := b.rOff + 1
	if rpos > b.wOff {
		return 0, errors.New("ReadUint8 out off wOff")
	}
	bytesBuffer := bytes.NewBuffer(b.b[b.rOff:rpos])
	var x uint8
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	b.rOff = rpos
	return x, nil
}

//get uint16
func (b *Buffer) ReadUint16() (uint16, error) {
	rpos := b.rOff + 2
	if rpos > b.wOff {
		return 0, errors.New("ReadUint16 out off wOff")
	}
	bytesBuffer := bytes.NewBuffer(b.b[b.rOff:rpos])
	var x uint16
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	b.rOff = rpos
	return x, nil
}

//读取uint32数据
func (b *Buffer) ReadUint32() (uint32, error) {
	rpos := b.rOff + 4
	if rpos > b.wOff {
		return 0, errors.New("ReadUint32 out off wOff")
	}
	bytesBuffer := bytes.NewBuffer(b.b[b.rOff:rpos])
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	b.rOff = rpos
	return x, nil
}

//读取字符串
func (b *Buffer) ReadString(length int) (string, error) {
	rpos := b.rOff + length
	if rpos > b.wOff {
		return "", errors.New("ReadString out of wOff")
	}
	s := string(b.b[b.rOff:rpos])
	b.rOff = rpos
	return s, nil
}

//判断是否包含指定[]byte，如果存在则返回位置，如果不存在返回-1
func (b *Buffer) Index(sep []byte) int {
	return bytes.Index(b.b[b.rOff:], sep)
}

////////////////////////////// io session ///////////////////////////

//Iosession
type Iosession struct {
	id        uint64
	serv      *ioserv
	conn      net.Conn
	closed    bool
	dataCh    chan interface{}
	userId    interface{}
	extraData map[string]interface{}
}

func (s *Iosession) Id() uint64 {
	return s.id
}

func (session *Iosession) SetUserId(id interface{}) {
	session.userId = id
}

func (session *Iosession) GetUserId() interface{} {
	return session.userId
}

func (session *Iosession) ExtraData(key string) (value interface{}, ok bool) {
	value, ok = session.extraData[key]
	return
}

func (session *Iosession) SetExtraData(key string, value interface{}) {
	session.extraData[key] = value
}

func (s *Iosession) Conn() net.Conn {
	return s.conn
}

func (s *Iosession) dealDataCh() {
	s.serv.wg.Add(1)
	defer s.serv.wg.Done()
	var msg interface{}
	for s.serv.runnable && !s.closed {
		select {
		case msg = <-s.dataCh:
			fmt.Println("收到消息: ", msg)
			// TODO: process message
			//s.serv.filterChain.msgReceived(s, msg)
		}
	}
}

func (session *Iosession) readData() {
	session.serv.wg.Add(1)
	ioBuffer := NewBuffer()
	buffer := make([]byte, 512)
	defer func() {
		if !session.closed {
			session.Close()
		}
		session.serv.wg.Done()
	}()
	var n int
	var err error
	for session.serv.runnable && !session.closed {
		n, err = session.conn.Read(buffer)
		ioBuffer.PutBytes(buffer[:n])
		if err != nil {
			//session.serv.filterChain.errorCaught(session, err)
			session.Close()
			return
		}
		fmt.Println("读取数据")
		// TODO: read data
		//err = session.serv.codecer.Decode(ioBuffer, session.dataCh)
		//if err != nil {
		//	session.serv.filterChain.errorCaught(session, err)
		//}
	}
}

func (session *Iosession) Write(message interface{}) error {
	if session.serv.runnable && !session.closed {
		//		if msg, ok := session.serv.filterChain.msgSend(session, message); ok {
		//			bs, err := session.serv.codecer.Encode(msg)
		//			if err != nil {
		//				return err
		//			}
		_, err := session.conn.Write(message.([]byte))
		if err != nil {
			//session.serv.filterChain.errorCaught(session, err)
			fmt.Println(err)
		}
		//}
		return nil
	} else {
		err := errors.New("Iosession is closed")
		//session.serv.filterChain.errorCaught(session, err)
		return err
	}
}

//close iosession
func (this *Iosession) Close() {
	if !this.closed {
		//this.serv.filterChain.sessionClosed(this)
		this.closed = true
	}
	this.conn.Close()
}

/////////////////////////////// io service //////////////////////////
type ioserv struct {
	generator_id uint64
	wg           sync.WaitGroup
	runnable     bool
}

//create session
func (serv *ioserv) newIoSession(conn net.Conn) *Iosession {
	session := &Iosession{}
	session.conn = conn
	session.serv = serv
	session.closed = false
	session.extraData = make(map[string]interface{})
	session.dataCh = make(chan interface{}, 16)
	session.id = atomic.AddUint64(&serv.generator_id, 1)
	go session.dealDataCh()
	go session.readData()
	//	go session.serv.filterChain.sessionOpened(session)
	return session
}

//stop serv
func (serv *ioserv) Stop() {
	serv.runnable = false
}

type client struct {
	ioserv
}

func NewClient() *client {
	c := &client{}

	return c
}

// dial to server
func (c *client) Dial(netPro, laddr string) {
	c.runnable = true
	conn, err := net.Dial(netPro, laddr)
	if err != nil {
		fmt.Println(err)
	}
	c.newIoSession(conn)
	time.Sleep(20 * time.Millisecond)
	c.wg.Wait()
}

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

	return nil
}
