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
		fmt.Fprintf(flag.CommandLine.Output(), "  length         Feet <-> Metres\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  temperature    Celsius <-> Fahrenheit <-> Kelvin\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  weight         Pounds <-> Kilograms\n")
		flag.PrintDefaults()
	}

	lengthCommand := flag.NewFlagSet("length", flag.ExitOnError)
	feetFlag := lengthCommand.Float64("feet", 0, "a length value in ft")
	metresFlag := lengthCommand.Float64("metres", 0, "a length value in m")

	temperatureCommand := flag.NewFlagSet("temperature", flag.ExitOnError)
	celsiusFlag := temperatureCommand.Float64("celsius", 0, "a temperature value in °C")
	fahrenheitFlag := temperatureCommand.Float64("fahrenheit", 0, "a temperature value in °F")
	kelvinFlag := temperatureCommand.Float64("kelvin", 0, "a temperature value in K")

	weightCommand := flag.NewFlagSet("weight", flag.ExitOnError)
	kilogramsFlag := weightCommand.Float64("kilograms", 0, "a weight value in kg")
	poundsFlag := weightCommand.Float64("pounds", 0, "a weight value in lb")

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case lengthCommand.Name():
		if errParse := lengthCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing length flags: %v", errParse)
		}
	case temperatureCommand.Name():
		if errParse := temperatureCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing temperature flags: %v", errParse)
		}
	case weightCommand.Name():
		if errParse := weightCommand.Parse(os.Args[2:]); errParse != nil {
			log.Fatalf("error parsing weight flags: %v", errParse)
		}
	default:
		fmt.Printf("'%v' is not a valid command.\n", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}

	if lengthCommand.Parsed() {
		if lengthCommand.NFlag() == 0 {
			fmt.Println("Please specify a length to convert.")
			lengthCommand.PrintDefaults()
			os.Exit(1)
		}

		lengthCommand.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "feet":
				feet := *feetFlag
				fmt.Printf("%vft is %vm.\n", feet, feet*3.281)
			case "metres":
				metres := *metresFlag
				fmt.Printf("%vm is %vft.\n", metres, metres/3.281)
			}
		})
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

	if weightCommand.Parsed() {
		if weightCommand.NFlag() == 0 {
			fmt.Println("Please specify a weight to convert.")
			weightCommand.PrintDefaults()
			os.Exit(1)
		}

		weightCommand.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "kilograms":
				kilograms := *kilogramsFlag
				fmt.Printf("%vkg is %vlb.\n", kilograms, kilograms*2.205)
			case "pounds":
				pounds := *poundsFlag
				fmt.Printf("%vlb is %vkg.\n", pounds, pounds/2.205)
			}
		})
	}
}
