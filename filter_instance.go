package fins

import (
	"fmt"
	"github.com/KevinZu/golis"
)

type MsgFilter struct {
	golis.IoFilterAdapter
	//FinsCmdChain *command.CmdChain
}

func (*MsgFilter) SessionClosed(session *golis.Iosession) bool {
	fmt.Println("session closed")
	clientGroup := GetClientGroup()

	fmt.Printf("---    session: %p\n", session)

	for clientInfo, s := range clientGroup.Clients {
		fmt.Printf("-+-    session: %p    s: %p\n", session, s)
		if s == session {
			fmt.Println("----- addr: ", clientInfo.conAddr)
			clientInfo.ConnStatus = DISCONNECT
			fmt.Println(clientInfo, " : ", clientInfo.ConnStatus)
			//session.Close()
			clientInfo.LastTryTime = GetCurTimeMs()
		}
	}

	// clientInfo := clientGroup.Clients[session]
	// if clientInfo != nil {
	// 	fmt.Println("--- addr: ", clientInfo.conAddr)
	// 	clientInfo.ConnStatus = DISCONNECT
	// 	clientInfo.LastTryTime = GetCurTimeMs()
	// }

	return true
}

func (*MsgFilter) SessionOpened(session *golis.Iosession) bool {

	fmt.Println("====== session opened")

	return true
}

func (m *MsgFilter) MsgReceived(session *golis.Iosession, message interface{}) bool {
	// if bs, ok := message.([]byte); ok {
	// 	//c := make(chan bool)
	// 	fmt.Println("received msg :", string(bs))
	// 	replayMsg := fmt.Sprintf(" -- msg : %v -- ", string(bs))
	// 	//<-c
	// 	session.Write([]byte(replayMsg))
	// }
	//cmdInfo := message.(command.CmdInfo)
	//fmt.Println("cmd name:   ", cmdInfo.CmdName)
	//fmt.Println("cmd:   ", cmdInfo.Cmd)
	//fmt.Printf("cmd len:%d", len(cmdInfo.Cmd))
	connectAddr, ok := session.ExtraData("connectAddr")
	if ok {
		fmt.Printf("cmd connect addr:%s", connectAddr.(string))
		//cmdInfo.ConnectAddr = connectAddr.(string)
	}
	//m.FinsCmdChain.InvokeCmd(cmdInfo)
	//rep := cmdInfo.Cmd[0:]
	replayMsg := fmt.Sprintf(" -- msg : %v -- ", message.(string))
	//<-c
	session.Write([]byte(replayMsg))
	return true
}

func (*MsgFilter) MsgSent(session *golis.Iosession, message interface{}) bool {
	fmt.Println("client msg sent")
	return true
}
