package main

import (
	"fmt"
	config "github.com/duchadian/xdg-direct/internal"
	"github.com/ryanuber/go-glob"
	"log"
	"os"
	"os/exec"
)

func usage() string {
	return "./xdg-direct <url>"
}
func main() {
	config.ReadConfig()

	if len(os.Args) == 1 {
		fmt.Println(usage())
		os.Exit(1)
	}

	url := os.Args[1]
	maps := config.UrlMappings()

	if config.DebugMode() {
		log.Printf("Url is %s", url)
		log.Printf("Mapping are %+v\n", config.UrlMappings())
	}

	// find possible match in config
	var command string
	var args []string
	for _, urlMap := range maps {
		if glob.Glob(urlMap.Url, url) {
			if config.DebugMode() {
				log.Printf("Found match for url %s on mapping %s", url, urlMap.Url)
			}
			command = urlMap.Command
			args = urlMap.Args
		}
	}
	args = append(args, url)

	// if no match can be found
	if command == "" {
		command = config.DefaultCommand()
	}

	var cmd *exec.Cmd
	if len(args) > 0 {
		cmd = exec.Command(command, args...)
	} else {
		cmd = exec.Command(command, args...)
	}

	err := cmd.Start()
	if config.DebugMode() {
		if err != nil {
			log.Println(err)
		}
		log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	}
}
