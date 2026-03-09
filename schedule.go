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
	slogger.Info("👋 Starting scheduler")

	logStartTime := time.Now()

	backup := func() {
		displayName := "Backup 💾"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\backup_to brooks.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	deloldlogs := func() {
		displayName := "DelOldLogs 🗑️"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\DELAGE.BAT")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	me_archiver := func() {
		displayName := "ME_Archiver 📁"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\me_archiver.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	fortune := func() {
		displayName := "Fortune 🥠"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\FORTUN.BAT")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	birdbuddy := func() {
		displayName := "BirdBuddy 🐦"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\birdbuddy.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	daytmpl := func() {
		displayName := "Daily Templates 📝"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\daytmpl.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	getfit := func() {
		displayName := "GetFit 💪"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\getfit.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	gem := func() {
		displayName := "Gem 💎"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\gem.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	cleangrey := func() {
		displayName := "CleanGrey 🦉"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\grey.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	logsumm := func() {
		displayName := "LogSumm 📋"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\logsumm.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	webby := func() {
		displayName := "WebbyStats 📋"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\webby.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	spamparse := func() {
		displayName := "SpamParse 🧱"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\spamparse.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	logparse := func() {
		displayName := "LogParse 📝"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\logparse.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	delage := func() {
		displayName := "DelAge 🗑️"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\delage.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	reserves := func() {
		displayName := "Reserves 🚒"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\RESERVES.BAT")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	rivercheck := func() {
		displayName := "RiverCheck 🏞️"
		slogger.With("job", displayName).Info("Running job")
		cmd := exec.Command("CMD", "/C", "C:\\AUTOJOB\\river.bat")
		err := cmd.Run()
		if err != nil {
			slogger.With("job", displayName, "error", err).Error("Job failed")
		}
	}

	// This function will check for new day and rotate logs if needed.
	var lastRotation = logStartTime
	rotateLog := func() {
		displayName := "RotateLog 🔄"
		slogger.With("job", displayName).Info("Running job")
		lastRotation = rotateLogging(lastRotation)
	}

	heartbeat := func() {
		slogger.With("job", "Heartbeat 💓").Info("Running Job")
	}

	// daily jobs
	scheduler.Every().Day().At("00:02:01").Run(rotateLog)
	scheduler.Every().Day().At("00:05:05").Run(me_archiver)
	scheduler.Every().Day().At("00:08:08").Run(deloldlogs)
	scheduler.Every().Day().At("00:03:03").Run(delage)
	scheduler.Every().Day().At("00:23:23").Run(backup)
	scheduler.Every().Day().At("03:33:05").Run(birdbuddy)
	scheduler.Every().Day().At("03:33:10").Run(webby)
	scheduler.Every().Day().At("03:33:15").Run(daytmpl)
	scheduler.Every().Day().At("03:33:33").Run(getfit)
	scheduler.Every().Day().At("03:33:35").Run(fortune)

	// email summaries
	scheduler.Every().Day().At("11:59:55").Run(logsumm)
	scheduler.Every().Day().At("23:59:55").Run(logsumm)

	// weekly jobs
	scheduler.Every().Monday().At("06:20:15").Run(reserves)
	scheduler.Every().Friday().At("03:33:38").Run(gem)

	// periodic jobs
	scheduler.Every(2).Minutes().Run(cleangrey)
	scheduler.Every(5).Minutes().Run(logparse)
	scheduler.Every(7).Minutes().Run(spamparse)
	scheduler.Every(9).Minutes().Run(heartbeat)
	scheduler.Every(303).Minutes().Run(rivercheck)

	slogger.Info("✅ All jobs scheduled, waiting for execution...")

	// Block forever in idiomatic Go style
	select {}
}
