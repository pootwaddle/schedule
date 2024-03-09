package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
var LogStartTime time.Time

func logWrapper() {
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
			LogStartTime = initLogging(LogStartTime)
		}
	}
}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func initLogging(t1 time.Time) time.Time {
	if !DateEqual(t1, time.Now()) {
		cfgBase := os.Getenv("RLOG_CONF") // C:\AUTOJOB\RLOG or ""
		cfgF := filepath.Join(cfgBase, "schedule.conf")
		logBase := os.Getenv("SCHEDULE_LOG_FILE") // E:\WEBSVC or ""
		fn := filepath.Join(logBase, fmt.Sprintf("schedule_%s.log", time.Now().Format("20060102")))
		err := os.WriteFile(cfgF, []byte("RLOG_LOG_FILE = "+fn), 0755)
		if err != nil {
			rlog.Error(err)
		}
		rlog.SetConfFile(cfgF)
		rlog.Infof("Config from %s", cfgF)
		rlog.Infof("Logging to %s", fn)
		return time.Now()
	}
	return t1
}

func main() {
	LogStartTime = initLogging(LogStartTime)

	var wg sync.WaitGroup
	wg.Add(1) //we don't ever do WaitGroup.Done, so we will always wait
	defer wg.Done()

	go logWrapper()

	fortune := func() {
		funcName := "scheduleFortune"
		logChan <- logMessage{level: "Info", message: funcName}
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\FORTUN.BAT")
		err := cmd.Run()
		if err != nil {
			logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
		}
	}

	grey := func() {
		funcName := "scheduleGrey"
		logChan <- logMessage{level: "Info", message: funcName}
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\grey.bat")
		err := cmd.Run()
		if err != nil {
			logChan <- logMessage{level: "Error", message: fmt.Sprintf("%s failed: %s", funcName, err)}
		}
	}

	reserves := func() {
		funcName := "scheduleReserves"
		logChan <- logMessage{level: "Info", message: funcName}
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\RESERVES.BAT")
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
	scheduler.Every(3).Minutes().Run(grey)
	scheduler.Every().Day().At("05:05:15").Run(fortune)
	scheduler.Every().Monday().At("06:20:15").Run(reserves)
	scheduler.Every().Day().At("00:00:01").Run(rotateLog)

	// Keep the program from exiting.
	wg.Wait()
	//runtime.Goexit()
}
