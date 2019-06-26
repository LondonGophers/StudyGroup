package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tempconv "github.com/go-london-user-group/study-group/workspaces/jlucktay/ch2/ex2_1"
)

func main() {
	temperatureCommand := flag.NewFlagSet("temperature", flag.ExitOnError)
	celsiusFlag := temperatureCommand.Float64("celsius", 0, "a temperature value in °C")

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case temperatureCommand.Name():
		if errParse := temperatureCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing temperature flags: %v", errParse)
		}
	default:
		log.Fatalf("'%v' is not a valid command.\n", os.Args[1])
	}

	if temperatureCommand.Parsed() {
		if temperatureCommand.NFlag() > 0 {
			if celsiusFlag != nil {
				celsius := tempconv.Celsius(*celsiusFlag)
				fmt.Printf("%v is %v.\n", celsius, tempconv.CToF(celsius))
			}
		} else {
			fmt.Println("Please specify a temperature to convert.")
		}
	}
}
