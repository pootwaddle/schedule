package main

import (
	"os/exec"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/pootwaddle/logger"
)

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// You can reload logger on log rotation if needed, or just rely on its internal logic.
func rotateLogging(last time.Time) time.Time {
	if !DateEqual(last, time.Now()) {
		logger.ReloadLogger()
		logger.Infof("Log rotated for new day: %s", time.Now().Format("2006-01-02"))
		return time.Now()
	}
	return last
}

func main() {
	defer logger.CloseLogger()
	logger.ReloadLogger() // Ensures env config is loaded at start

	logStartTime := time.Now()

	backup := func() {
		funcName := "Backup 💾"
		logger.Infof("Running %s", funcName)
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\backup_c_drive_to_tech1.bat")
		err := cmd.Run()
		if err != nil {
			logger.Errorf("%s failed: %v", funcName, err)
		}
	}

	deloldlogs := func() {
		funcName := "DelOldLogs 🗑️"
		logger.Infof("Running %s", funcName)
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\DELAGE.BAT")
		err := cmd.Run()
		if err != nil {
			logger.Errorf("%s failed: %v", funcName, err)
		}
	}

	fortune := func() {
		funcName := "Fortune 🥠"
		logger.Infof("Running %s", funcName)
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\FORTUN.BAT")
		err := cmd.Run()
		if err != nil {
			logger.Errorf("%s failed: %v", funcName, err)
		}
	}

	gem := func() {
		funcName := "Gem 💎"
		logger.Infof("Running %s", funcName)
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\gem.bat")
		err := cmd.Run()
		if err != nil {
			logger.Errorf("%s failed: %v", funcName, err)
		}
	}

	grey := func() {
		funcName := "Grey 🦉"
		logger.Infof("Running %s", funcName)
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\greylist.bat")
		err := cmd.Run()
		if err != nil {
			logger.Errorf("%s failed: %v", funcName, err)
		}
	}

	logthings := func() {
		funcName := "LogThings 📋"
		logger.Infof("Running %s", funcName)
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\logthings.bat")
		err := cmd.Run()
		if err != nil {
			logger.Errorf("%s failed: %v", funcName, err)
		}
	}

	reserves := func() {
		funcName := "Reserves 🚒"
		logger.Infof("Running %s", funcName)
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\RESERVES.BAT")
		err := cmd.Run()
		if err != nil {
			logger.Errorf("%s failed: %v", funcName, err)
		}
	}

	// This function will check for new day and rotate logs if needed.
	var lastRotation = logStartTime
	rotateLog := func() {
		funcName := "RotateLog 🔄"
		logger.Infof("Running %s", funcName)

		logger.Info("Checking if log rotation needed...")
		lastRotation = rotateLogging(lastRotation)
	}

	heartbeat := func() {
		logger.Info("Heartbeat - scheduler is alive. 💓")
	}

	scheduler.Every(30).Minutes().Run(heartbeat)
	scheduler.Every(2).Minutes().Run(grey)
	scheduler.Every(31).Minutes().Run(logthings)
	scheduler.Every().Day().At("00:23:23").Run(backup)
	scheduler.Every().Day().At("23:55:55").Run(deloldlogs)
	scheduler.Every().Day().At("03:33:33").Run(fortune)
	scheduler.Every().Monday().At("06:20:15").Run(reserves)
	scheduler.Every().Friday().At("03:33:33").Run(gem)
	scheduler.Every().Day().At("00:00:01").Run(rotateLog)

	// Block forever in idiomatic Go style
	select {}
}
