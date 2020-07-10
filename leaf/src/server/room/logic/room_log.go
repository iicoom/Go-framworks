package logic

import (
	"fmt"
	"github.com/name5566/leaf/log"
)

func (room *Room) Info(format string, a ...interface{}) {
	log.Release(fmt.Sprintf("<ROOM%v> %v", room.index, format), a...)
}

func (room *Room) Error(format string, a ...interface{}) {
	log.Error(fmt.Sprintf("<ROOM%v> %v", room.index, format), a...)
}

func (room *Room) Fatal(format string, a ...interface{}) {
	log.Fatal(fmt.Sprintf("<ROOM%v> %v", room.index, format), a...)
}

