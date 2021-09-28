package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/romana/rlog"
)

func main() {

	//initialize logging
	logfileName := fmt.Sprintf("schedule_%s", time.Now().Format("20060102"))
	// Example of redirecting log output to a new file at runtime
	newLogFile, err := os.OpenFile(filepath.Join("D:\\ARCHIVE\\"+logfileName+".log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		rlog.SetOutput(newLogFile)
		rlog.Info(os.Args[0] + " started")
	}

	fortune := func() {
		rlog.Info("scheduleFortune")
		cmd := exec.Command("CMD", "/C D:\\ARCHIVE\\FORTUN.BAT")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			rlog.Error("scheduleFortune failed: ", err)
		}
	}

	malware1 := func() {
		rlog.Info("scheduleMalwareBytes")
		cmd := exec.Command("CMD", "/C D:\\ARCHIVE\\mw1.bat")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			rlog.Error("scheduleMalwareBytes failed:", err)
		}
	}

	reserves := func() {
		rlog.Info("scheduleReserves")
		cmd := exec.Command("CMD", "/C D:\\ARCHIVE\\RESERVES.BAT")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			rlog.Error("scheduleReserves failed: ", err)
		}
	}

	rotateLog := func() {
		rlog.Info("rotateLog")
		//initialize new log file
		logfileName := fmt.Sprintf("cert-check_%s", time.Now().Format("20060102"))
		// Example of redirecting log output to a new file at runtime
		newLogFile, err := os.OpenFile(filepath.Join("D:\\ARCHIVE\\"+logfileName+".log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err == nil {
			rlog.SetOutput(newLogFile)
			rlog.Info(newLogFile)
		}
	}

	//scheduler.Every(2).Seconds().NotImmediately().Run(job1)
	scheduler.Every().Day().At("05:00:15").Run(fortune)
	scheduler.Every().Day().At("06:35:15").Run(malware1)
	scheduler.Every().Monday().At("06:20:15").Run(reserves)
	scheduler.Every().Day().At("00:00:01").Run(rotateLog)

	// Keep the program from exiting.
	runtime.Goexit()
}
