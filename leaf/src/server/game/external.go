package game

import (
	"myLeaf/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
//首先，模块会被实例化，这样才能注册到 Leaf 框架中（详见 LeafServer main.go），另外，模块暴露的 ChanRPC 被用于模块间通讯。
//internal.Module 模块中最关键的就是 skeleton（骨架），skeleton 实现了 Module 接口的 Run 方法并提供了：
//- ChanRPC
//- goroutine
//- 定时器