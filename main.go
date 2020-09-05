package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"ping-monitor/monitor"
	"time"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	serverListFile := flag.String("server_list_file", pwd+"/server_list.json", "default: "+pwd+"/server_list.json")
	envFile := flag.String("env_file", pwd+"/.env", "default: "+pwd+"/.env")
	flag.Parse()

	if err := godotenv.Load(*envFile); err != nil {
		log.Println(err)
	}

	monitor.Monitoring(*serverListFile)

	for {
		log.Println("monitoring...")
		time.Sleep(time.Hour)
	}
}
