package logic

import (
	"github.com/name5566/leaf/gate"
	sample "myLeaf/msg/MyGame/Sample"
	"sync"
	"time"
)

type roomRPCHandler func(agent gate.Agent, rpc *sample.RpcToServer)
type roomRawMsgHandler func(agent gate.Agent, msg []byte)
type roomBufferedMsg struct {
	msg        []byte
	receiverID string
}

type Room struct {
	playersMutex sync.RWMutex

	PlayerBuffPool *sync.Pool

	agentsMutex sync.RWMutex
	agents      map[int]gate.Agent
	isReleased  bool

	index        int

	builderPool             *sync.Pool
	bufferPool              *sync.Pool
	rpcHandlers             map[string]roomRPCHandler
	bufferedRpcs            []*roomBufferedMsg
	playerWaypointHandler   roomRawMsgHandler

	timers       map[chan bool]struct{}
	timersMutex  sync.Mutex
	tickers      map[chan bool]struct{}
	tickersMutex sync.Mutex

	lastWriteAlarmTime time.Time
}

func (room *Room) RegisterRPC(name string, rpcHandler roomRPCHandler) {
	if _, yes := room.rpcHandlers[name]; yes {
		room.Error("a rpc function %v already registered", name)
		return
	}
	room.rpcHandlers[name] = rpcHandler
}