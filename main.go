package main

import (
	"flag"
	"os"
	"time"

	"fmt"

	"github.com/automata-aws-cli/deploy"
	"github.com/automata-aws-cli/utilities"
	"github.com/briandowns/spinner"
)

func main() {

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)

	// deployment create
	configfilepath := createCommand.String("config", "", "Path of the "+
		"configuration file for the micro service deployment, use "+
		"--config ./myconfig")
	createEnvName := createCommand.String("env", "", "Provide Environment Name "+
		"usage :"+
		"--env [lle/hle/prod]")
	createSvcName := createCommand.String("svc", "", "Provide Service Name "+
		"usage :"+
		"--svc [mui/msp/browse/loyalty/encrypt]")
	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)
	versionCommand := flag.NewFlagSet("version", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println(" Oops ! Not received any command, list of " +
			"possible " +
			"commands is {deployment,service, configure, init, login, signup, logout}")
		os.Exit(1)
	} else {

		if len(os.Args) < 3 && os.Args[1] == "launchConfig" {
			utilities.DeploymentHelp()
			os.Exit(1)
		}
		if len(os.Args) < 3 && os.Args[1] == "autoScaling" {
			utilities.DeploymentHelp()
			os.Exit(1)
		}

	}
	switch os.Args[1] {
	case "help":
		helpCommand.Parse(os.Args[2:])
		flag.Parse()
	case "version":
		versionCommand.Parse(os.Args[2:])
		flag.Parse()
	case "v":
		versionCommand.Parse(os.Args[2:])
		flag.Parse()
	case "launchConfig":
		switch os.Args[2] {
		case "create":
			createCommand.Parse(os.Args[3:])
		}
	case "autoScaling":
		switch os.Args[2] {
		case "create":
			createCommand.Parse(os.Args[3:])
		}
	default:
		flag.PrintDefaults()
	}

	if createCommand.Parsed() {
		if *createEnvName == "" {
			fmt.Println("Error Occured, Missing Something, please" +
				" verify below , have you given all " +
				"information?")
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		if *configfilepath == "" {
			fmt.Println("Error Occured, Missing Something, please" +
				" verify below , have you given all " +
				"information?")
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		if *createSvcName == "" {
			fmt.Println("Error Occured, Missing Something, please" +
				" verify below , have you given all " +
				"information?")
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		deploy.LaunchConfig(*configfilepath, *createEnvName, *createSvcName)
	}

	if helpCommand.Parsed() {
		utilities.Help()
		os.Exit(0)
	}
	if versionCommand.Parsed() {
		s := spinner.New(spinner.CharSets[35], 350*time.Millisecond)
		s.Color("green")
		s.Start()
		s.Stop()
		utilities.Version()
		os.Exit(0)
	}

}
