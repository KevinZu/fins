package fins

import (
	"errors"
	"fmt"
	"github.com/KevinZu/golis"
	//"time"
)

type ProtoCodec struct {
	golis.ProtocalCodec
	//FinsCmdChain *command.CmdChain
}

type FrameCheckStatus uint

const (
	PRE_START FrameCheckStatus = 1
	START                      = 2
)

func getFrame(bs []byte, dataCh chan<- interface{}) int {
	var frame []byte
	var start int
	var end int
	var posAdd = 0
	var status FrameCheckStatus = PRE_START

	for i, v := range bs {
		if v == '@' && status == PRE_START {
			status = START
			start = i
			posAdd = i
		}

		if v == 0x0d && status == START && bs[i-1] == 0x2a {
			status = PRE_START
			end = i
			frame = bs[start:end]

			dataCh <- frame

			posAdd = i + 1
		}
	}

	return posAdd
}

func (p *ProtoCodec) Decode(buffer *golis.Buffer, dataCh chan<- interface{}) error {

	rdPos := buffer.GetReadPos()
	rdLen := buffer.GetWritePos() - buffer.GetReadPos()

	bs, _ := buffer.ReadBytes(rdLen)

	posAdd := getFrame(bs, dataCh) //getFinsFrame(bs, dataCh)

	rdPos += posAdd

	buffer.SetReadPos(rdPos)

	return nil
}

func (*ProtoCodec) Encode(message interface{}) ([]byte, error) {
	// TODO: 封装fins包
	if bs, ok := message.([]byte); ok {
		//dataFrame, err := fins.BuildFinsFrame(bs)
		//return dataFrame, err
		fmt.Println(bs)
	}
	return nil, errors.New("failed")
}
