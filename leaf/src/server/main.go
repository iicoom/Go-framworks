package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	log "github.com/name5566/leaf/log"
	"myLeaf/conf"
	"myLeaf/game"
	"myLeaf/gate"
	"myLeaf/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	//var logger = logrus.New()
	//logger.SetFormatter(&logrus.TextFormatter{
	//	ForceColors: true,
	//	FullTimestamp: true,
	//})
	//logger.SetOutput(colorable.NewColorableStdout())
	//logger.SetReportCaller(true)
	//logger.Level = logrus.TraceLevel
	//
	//logger.Trace("Something very low level.")
	//logger.Debug("Useful debugging information.")
	//logger.Info("Something noteworthy happened!")
	//logger.Warn("You should probably take a look at this.")
	//logger.Error("Something failed but I'm not quitting.")
	//// Calls panic() after logging
	//logger.Panic("I'm bailing.")
	//// Calls os.Exit(1) after logging
	//logger.Fatal("Bye.")
	//
	//logger.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"number": 0,
	//}).Trace("Went to the beach")

	name := "Leaf"

	log.Debug("My name is %v", name)

	leaf.Run(
		game.Module,
		gate.Module,	// 网络请求处理
		login.Module,
	)
	//time.Sleep(5 * time.Second)
}
