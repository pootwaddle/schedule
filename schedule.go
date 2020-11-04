package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	//"time"

	"github.com/carlescere/scheduler"
)

func main() {

	fortune := func() {
		fmt.Println("scheduleFortune")
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\FORTUN.BAT")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("whoops!", err)
		}
	}

	malware1 := func() {
		fmt.Println("scheduleMalwareBytes")
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\mw1.bat")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("whoops!", err)
		}
	}

	reserves := func() {
		fmt.Println("scheduleReserves")
		cmd := exec.Command("CMD", "/C C:\\AUTOJOB\\RESERVES.BAT")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("whoops!", err)
		}
	}

	//scheduler.Every(2).Seconds().NotImmediately().Run(job1)
	scheduler.Every().Day().At("05:00:15").Run(fortune)
	scheduler.Every().Day().At("06:35:15").Run(malware1)
	scheduler.Every().Monday().At("06:20:15").Run(reserves)

	// Keep the program from exiting.
	runtime.Goexit()
}
