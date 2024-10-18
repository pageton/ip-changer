package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/pageton/ip-changer/configs"
	"github.com/pageton/ip-changer/internal/scheduler"
	"github.com/pageton/ip-changer/internal/tor"
)

func main() {
	config := configs.LoadConfig("configs/config.yaml")

	torClient := tor.NewClient("127.0.0.1:9051")
	defer torClient.Close()

	interval := config.Scheduler.Interval
	scheduler.ScheduleIPChange(func() {
		torClient.ChangeIP()
		cmd := exec.Command("curl", "--socks5-hostname", "127.0.0.1:9050", "https://api.ipify.org")
		output, err := cmd.Output()
		if err != nil {
			log.Fatalf("Failed to run curl: %v", err)
		}
		fmt.Printf("Current IP: %s\n", output)
	}, interval)
}
