package main

import (
	"os/exec"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/pootwaddle/slogger"
)

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// You can reload logger on log rotation if needed, or just rely on its internal logic.
func rotateLogging(last time.Time) time.Time {
	if !DateEqual(last, time.Now()) {
		slogger.ReloadLogger()
		slogger.With("date", time.Now().Format("2006-01-02")).Info("Log rotated for new day")
		return time.Now()
	}
	return last
}

func main() {
	defer slogger.CloseLogger()
	slogger.ReloadLogger() // Ensures env config is loaded at start

	logStartTime := time.Now()

	backup := func() {
		funcName := "Backup 💾"
		slogger.With("job", funcName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\backup_c_drive_to_tech1.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", funcName, "error", err).Error("Job failed")
		}
	}

	deloldlogs := func() {
		funcName := "DelOldLogs 🗑️"
		slogger.With("job", funcName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\DELAGE.BAT")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", funcName, "error", err).Error("Job failed")
		}
	}

	fortune := func() {
		funcName := "Fortune 🥠"
		slogger.With("job", funcName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\FORTUN.BAT")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", funcName, "error", err).Error("Job failed")
		}
	}

	gem := func() {
		funcName := "Gem 💎"
		slogger.With("job", funcName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\gem.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", funcName, "error", err).Error("Job failed")
		}
	}

	grey := func() {
		funcName := "Grey 🦉"
		slogger.With("job", funcName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\greylist.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", funcName, "error", err).Error("Job failed")
		}
	}

	logthings := func() {
		funcName := "LogThings 📋"
		slogger.With("job", funcName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\logthings.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", funcName, "error", err).Error("Job failed")
		}
	}

	reserves := func() {
		funcName := "Reserves 🚒"
		slogger.With("job", funcName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\RESERVES.BAT")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", funcName, "error", err).Error("Job failed")
		}
	}

	// This function will check for new day and rotate logs if needed.
	var lastRotation = logStartTime
	rotateLog := func() {
		funcName := "RotateLog 🔄"
		slogger.With("job", funcName).Info("Running job")
		lastRotation = rotateLogging(lastRotation)
	}

	heartbeat := func() {
		slogger.With("job", "heartbeat 💓").Info("Running Job")
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
