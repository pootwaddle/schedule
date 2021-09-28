package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/romana/rlog"
)

type logMessage struct {
	level   string
	message string
}

var logChan = make(chan logMessage)

func logWrapper() {
	//initialize logging
	logfileName := fmt.Sprintf("D:\\ARCHIVE\\schedule_%s.log", time.Now().Format("20060102"))
	os.Setenv("RLOG_LOG_FILE", logfileName)
	rlog.UpdateEnv()
	rlog.Info(os.Args[0] + " started")

	for v := range logChan {
		switch v.level {
		case "Error":
			rlog.Error(v.message)
		case "Info":
			rlog.Info(v.message)
		case "Warn":
			rlog.Info(v.message)
		case "Heartbeat":
			rlog.Info(v.message)
		case "Rotate":
			rlog.Info(v.message)
			//initialize logging
			logfileName := fmt.Sprintf("D:\\ARCHIVE\\schedule_%s.log", time.Now().Format("20060102"))
			os.Setenv("RLOG_LOG_FILE", logfileName)
			rlog.UpdateEnv()
			rlog.Info(fmt.Sprintf("new log filename: %s", logfileName))
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1) //we don't ever do WaitGroup.Done, so we will always wait

	go logWrapper()

	fortune := func() {
		funcName := "scheduleFortune"
		logChan <- logMessage{level: "Info", message: funcName}
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\FORTUN.BAT >nul")
		err := cmd.Run()
		if err != nil {
			logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
		}
	}

	malware1 := func() {
		funcName := "scheduleMalwareBytes"
		logChan <- logMessage{level: "Info", message: funcName}
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\mw1.bat")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
		}
	}

	reserves := func() {
		funcName := "scheduleReserves"
		logChan <- logMessage{level: "Info", message: funcName}
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\RESERVES.BAT")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
		}
	}

	rotateLog := func() {
		funcName := "scheduleRotate"
		logChan <- logMessage{level: "Info", message: funcName}
		logChan <- logMessage{level: "Rotate", message: "new day"}
	}

	heartbeat := func() {
		funcName := "Heartbeat"
		logChan <- logMessage{level: "Heartbeat", message: funcName}
	}

	scheduler.Every(30).Minutes().Run(heartbeat)
	// scheduler.Every(5).Minutes().Run(fortune) //debug
	scheduler.Every().Day().At("05:00:15").Run(fortune)
	scheduler.Every().Day().At("06:35:15").Run(malware1)
	scheduler.Every().Monday().At("06:20:15").Run(reserves)
	scheduler.Every().Day().At("00:00:01").Run(rotateLog)

	// Keep the program from exiting.
	wg.Wait()
	//runtime.Goexit()
}
