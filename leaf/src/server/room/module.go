package room

import (
	"github.com/name5566/leaf/module"
	"myLeaf/base"
	"myLeaf/room/logic"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.initChanRPC()
}

func (m *Module) initChanRPC() {
	m.RegisterChanRPC("NewBattle", createNewBattle)
	m.RegisterChanRPC(msg.Ping, handlePing)
}