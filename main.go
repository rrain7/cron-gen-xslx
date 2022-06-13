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

	c := gron.New()
	c.AddFunc(gron.Every(1*time.Hour), func() {
		cmd := exec.Command("./venv/Scripts/python.exe", "main.py")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("运行出错: %s\n%s\n", out, err)
		}
		fmt.Printf("%s\n", string(out))
	})

	c.Start()
	wg.Wait()
}
