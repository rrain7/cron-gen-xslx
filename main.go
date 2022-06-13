package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/roylee0704/gron"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	log.Println("start running~")
	c := gron.New()
	c.AddFunc(gron.Every(1*time.Hour), func() {
		cmd := exec.Command("./venv/Scripts/python.exe", "main.py")
		out, err := cmd.CombinedOutput()
		log.Println("get exec stdout and stderr")

		if err != nil {
			log.Fatalf("run script error: %s\n%s\n", out, err)
		}

		fmt.Printf("%s\n", string(out))
	})
	
	log.Println("add and start cron job")
	c.Start()
	wg.Wait()
	fmt.Println("stats data to file, please check it and wait for 1 hour ~")
}
