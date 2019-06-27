package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tempconv "github.com/go-london-user-group/study-group/workspaces/jlucktay/ch2/ex2_1"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Available commands for %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  length         Feet <-> Meters\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  temperature    Celsius <-> Fahrenheit <-> Kelvin\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  weight         Pounds <-> Kilograms\n")
		flag.PrintDefaults()
	}

	temperatureCommand := flag.NewFlagSet("temperature", flag.ExitOnError)
	celsiusFlag := temperatureCommand.Float64("celsius", 0, "a temperature value in °C")
	fahrenheitFlag := temperatureCommand.Float64("fahrenheit", 0, "a temperature value in °F")
	kelvinFlag := temperatureCommand.Float64("kelvin", 0, "a temperature value in K")

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
		if temperatureCommand.NFlag() == 0 {
			fmt.Println("Please specify a temperature to convert.")
			temperatureCommand.PrintDefaults()
			os.Exit(1)
		}

		temperatureCommand.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "celsius":
				celsius := tempconv.Celsius(*celsiusFlag)
				fmt.Printf("%v is %v and %v.\n", celsius, tempconv.CToF(celsius), tempconv.CToK(celsius))
			case "fahrenheit":
				fahrenheit := tempconv.Fahrenheit(*fahrenheitFlag)
				fmt.Printf("%v is %v and %v.\n", fahrenheit, tempconv.FToC(fahrenheit), tempconv.FToK(fahrenheit))
			case "kelvin":
				kelvin := tempconv.Kelvin(*kelvinFlag)
				fmt.Printf("%v is %v and %v.\n", kelvin, tempconv.KToC(kelvin), tempconv.KToF(kelvin))
			}
		})
	}
}
