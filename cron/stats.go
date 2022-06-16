package cron

import (
	"log"
	"os/exec"
)

func Stats() {
	cmd := exec.Command("python.exe", "do_one.py")
	out, err := cmd.CombinedOutput()
	log.Println("get exec stdout and stderr")
	if err != nil {
		log.Fatalf("run script error: %s\n%s\n", out, err)
	}

	log.Println(string(out))
}