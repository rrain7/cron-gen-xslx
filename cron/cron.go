package cron

import (
	"log"
	"os/exec"
	"time"

	"github.com/go-co-op/gocron"
)

var task = func() {
	log.Println("start exec python script...")
	cmd := exec.Command("./venv/Scripts/python.exe", "main.py")
	out, err := cmd.CombinedOutput()
	log.Println("get exec stdout and stderr")
	if err != nil {
		log.Fatalf("run script error: %s\n%s\n", out, err)
	}

	log.Println(string(out))
}

func Job() {
	log.Println("Start to run cron job")
	log.Println("The Job will be exec at every hour 1th min")
	log.Println("Please waiting... and Don't close the window")
	s := gocron.NewScheduler(time.UTC)
	s.Cron("1 * * * *").Do(task)

	// s.StartAsync()
	s.StartBlocking()
}
