package fins

import (
	"container/list"
	"fmt"
	"github.com/KevinZu/golis"
	"sync"
	"time"
)

var lock sync.Mutex
var clientGroupOnce sync.Once

type ClientAddr struct {
	Ip   string
	Port uint
}

type ClientGroup struct {
	Dch chan bool
	//	ClientAddrs map[string]*ClientAddr
	Clients map[*ClientInfo]*golis.Iosession
	Timer   *time.Ticker
}

var clientGroupInstance *ClientGroup

func GetClientGroup() *ClientGroup {
	clientGroupOnce.Do(func() {
		clientGroupInstance = &ClientGroup{}
		clientGroupInstance.Clients = make(map[*ClientInfo]*golis.Iosession)
		//	clientGroupInstance.ClientAddrs = make(map[string]*ClientAddr)
		clientGroupInstance.Dch = make(chan bool)
		clientGroupInstance.Timer = time.NewTicker(10 * time.Second)
		go clientGroupInstance.TimerHandle(clientGroupInstance.Timer)
	})
	return clientGroupInstance
}

type ConnectStatus int

var (
	CONNECT    ConnectStatus = 0
	DISCONNECT ConnectStatus = 1
)

type ClientInfo struct {
	Cli         *golis.Client
	ConnStatus  ConnectStatus
	conAddr     string
	Session     *golis.Iosession
	LastTryTime int64

	ErrorCount   int32
	ErrorMax     int32
	LastError    int32
	ErrorChanged bool
	//	SyncCmd     map[byte]string // key: SID  val: FinsCmd
	//CommandList *list.List
}

func (cli *ClientInfo) Init(error_max int32) {
	cli.ErrorMax = error_max
	cli.LastError = FINS_RETVAL_SUCCESS
	cli.ErrorCount = 0
	cli.ErrorChanged = false
}

func GetCurTimeMs() int64 {
	return time.Now().UnixNano() / 1e6
}

func PushReq(l *list.List, v interface{}) {
	defer lock.Unlock()
	lock.Lock()
	l.PushFront(v)
}

func popReq(l *list.List) interface{} {
	defer lock.Unlock()
	lock.Lock()
	iter := l.Back()
	v := iter.Value
	l.Remove(iter)
	return v
}

func (cg *ClientGroup) AddNewClient(cli *ClientAddr,error_max int32) (error, *ClientInfo) {
	filter := &MsgFilter{}
	c := golis.NewClient()
	c.FilterChain().AddLast("clientFilter", filter)
	codec := &ProtoCodec{}

	c.SetCodecer(codec)
	connectAddr := fmt.Sprintf("%s:%d", cli.Ip, cli.Port)
	fmt.Println("== addr: ", connectAddr)
	//cg.ClientAddrs[connectAddr] = cli

	cliInfo := ClientInfo{}
	cliInfo.Init(error_max)
	cliInfo.Cli = c
	cliInfo.conAddr = connectAddr
	//	cliInfo.CommandList = list.New()

	err := c.Dial("tcp", connectAddr)
	if err != nil {
		//c.Session.Close()
		cliInfo.ConnStatus = DISCONNECT
		cg.Clients[&cliInfo] = nil
		//cliInfo.LastTryTime = GetCurTimeMs()
		return err, nil

	} else {
		cliInfo.ConnStatus = CONNECT
		fmt.Printf("*** %p\n", c.Session) //
		c.Session.SetExtraData("connectAddr", connectAddr)
		cg.Clients[&cliInfo] = c.Session
	}

	fmt.Printf("+++    session: %p\n", c.Session)
	fmt.Printf("+++    cliInfo: %p\n", &cliInfo)

	return nil, &cliInfo
}

func (cg *ClientGroup) DelClient(cli *ClientInfo) error {
	for info, session := range cg.Clients {
		if info == cli {
			if info.ConnStatus == CONNECT {
				session.Close()
			}

			delete(cg.Clients, info)
			break
		}
	}
	return nil
}

func (cg *ClientGroup) BuildClients(clients []*ClientAddr, error_max int32) error {

	//cg.ClientAddrs = clients
	//list.New()
	filter := &MsgFilter{}

	for _, v := range clients {

		c := golis.NewClient()
		c.FilterChain().AddLast("clientFilter", filter)
		codec := &ProtoCodec{}

		c.SetCodecer(codec)
		connectAddr := fmt.Sprintf("%s:%d", v.Ip, v.Port)
		fmt.Println("== addr: ", connectAddr)
		//cg.ClientAddrs[connectAddr] = v

		cliInfo := ClientInfo{}
		cliInfo.Init(error_max)
		cliInfo.Cli = c
		cliInfo.conAddr = connectAddr
		//	cliInfo.CommandList = list.New()

		err := c.Dial("tcp", connectAddr)
		if err != nil {
			//c.Session.Close()
			cliInfo.ConnStatus = DISCONNECT
			cg.Clients[&cliInfo] = nil
			//cliInfo.LastTryTime = GetCurTimeMs()

		} else {
			cliInfo.ConnStatus = CONNECT
			fmt.Printf("*** %p\n", c.Session) //
			c.Session.SetExtraData("connectAddr", connectAddr)
			cg.Clients[&cliInfo] = c.Session
		}

		fmt.Printf("+++    session: %p\n", c.Session)
		fmt.Printf("+++    cliInfo: %p\n", &cliInfo)
	}

	select {
	case <-cg.Dch:
		fmt.Println("关闭连接")
	}

	return nil
}

func (cg *ClientGroup) TimerHandle(ticker *time.Ticker) {
	//curTime := GetCurTimeMs()
	//fmt.Println("=== cur: ", curTime)

	for {
		<-ticker.C
		//fmt.Println("&&&        map len: ", len(cg.Clients))
		for info, session := range cg.Clients {
			//if curTime-info.LastTryTime > 1000 {
			//fmt.Println(info, " : ", info.ConnStatus)
			if info.ConnStatus == DISCONNECT {
				fmt.Println("====   try connect: ", info.conAddr)
				//go func() {
				err := info.Cli.ReDial("tcp", info.conAddr)
				if err != nil {
					//info.ConnStatus = DISCONNECT
					info.LastTryTime = GetCurTimeMs()
				} else {
					//fmt.Printf("==+=== session:%p    client:%p\n", info.Cli.Session, info.Cli)
					info.ConnStatus = CONNECT
					info.Cli.Session.SetExtraData("connectAddr", info.conAddr)
					cg.Clients[info] = info.Cli.Session
				}
				//}()

			} else {
				heartBeat := fmt.Sprintf("heart beat!\n")
				if session != nil {
					err := session.Write([]byte(heartBeat))
					if err != nil {
						//info.ConnStatus = DISCONNECT
						session.Close()
						cg.Clients[info] = nil
					}
				}

			}
		}
	}

}
