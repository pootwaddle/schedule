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
		param := fmt.Sprintf("D:\\ARCHIVE\\scheduleDetail_%s.log", time.Now().Format("20060102"))
		if param == "" {
			logChan <- logMessage{level: "Error", message: "os.Getenv RLOG_LOG_FILE is empty"}
		}
		cmd := exec.Command("CMD", fmt.Sprintf("/C C:\\AUTOJOB\\FORTUN.BAT >>%s", param))
		err := cmd.Run()
		if err != nil {
			logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
		}
	}

	grey := func() {
		funcName := "scheduleGrey"
		logChan <- logMessage{level: "Info", message: funcName}
		param := fmt.Sprintf("D:\\ARCHIVE\\scheduleDetail_%s.log", time.Now().Format("20060102"))
		if param == "" {
			logChan <- logMessage{level: "Error", message: "os.Getenv RLOG_LOG_FILE is empty"}
		}
		cmd := exec.Command("CMD", fmt.Sprintf("/C C:\\AUTOJOB\\GREY.BAT >>%s", param))
		err := cmd.Run()
		if err != nil {
			logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
		}
	}

	/*
		malware1 := func() {
			funcName := "scheduleMalwareBytes"
			logChan <- logMessage{level: "Info", message: funcName}
			param := fmt.Sprintf("D:\\ARCHIVE\\scheduleDetail_%s.log", time.Now().Format("20060102"))
			cmd := exec.Command("CMD", fmt.Sprintf("/C C:\\AUTOJOB\\mw1 >>%s", param))
			err := cmd.Run()
			if err != nil {
				logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
			}
		}
	*/

	reserves := func() {
		funcName := "scheduleReserves"
		logChan <- logMessage{level: "Info", message: funcName}
		param := fmt.Sprintf("D:\\ARCHIVE\\scheduleDetail_%s.log", time.Now().Format("20060102"))
		cmd := exec.Command("CMD", fmt.Sprintf("/C C:\\AUTOJOB\\RESERVES.BAT >>%s", param))
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
	scheduler.Every(3).Minutes().Run(grey) //debug
	scheduler.Every().Day().At("05:05:15").Run(fortune)
	// scheduler.Every().Day().At("06:35:15").Run(malware1)
	scheduler.Every().Monday().At("06:20:15").Run(reserves)
	scheduler.Every().Day().At("00:00:01").Run(rotateLog)

	// Keep the program from exiting.
	wg.Wait()
	//runtime.Goexit()
}
