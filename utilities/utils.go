package utilities

import (
	"fmt"

	"github.com/fatih/color"
)

//Help for automata cli help menu
func Help() {
	fmt.Println("NAME")
	fmt.Println("\tautomata-micro - Automata micro for Managing the micro service deployments\n ")
	fmt.Println("SYNOPSIS")
	fmt.Println("\tautomata-micro GROUP | COMMAND [--help] [--version,-v]\n ")
	fmt.Println("GLOBAL FLAGS")
	fmt.Println("\t--help\n\t   Display detailed help.\n ")
	fmt.Println("\t--version, -v\n\t   Print version information and exit. This flag is only available at the\n\t   global level.\n ")
	fmt.Println("GROUPS\n\tGROUP is one of the following:\n ")
	fmt.Println("\t deployment\n\t   Manage your deployments.\n ")
	fmt.Println("\t login\n\t   For Logging in the Automata System.\n ")
	fmt.Println("\t logout\n\t  For Logging Out and ending the user session.\n ")
	fmt.Println("\t nginx\n\t   For Ingress Deployments in GKE.\n ")
	fmt.Println("\t signup\n\t  For Signing Up with the Automata System.\n ")
	fmt.Println("\t init \n\t   For initialising Jenkins for Automata.\n ")
	fmt.Println("\t info \n\t   For the Automata System Info.\n ")
	fmt.Println("\t kubeconnect \n\t   For connecting and opening a kubernetes cluster dashboard.\n ")
	fmt.Println("\t ")
}

//Version for automata micro version
func Version() {
	fmt.Println("AWS Automata Micro v1.0.0")
}

//DeploymentHelp for automata micro
func DeploymentHelp() {
	fmt.Println("NAME")
	fmt.Println("\tautomata-aws-micro deployment - to create launchConfig \n ")
	fmt.Println("GROUPS\n\tGROUP is one of the following:\n ")
	fmt.Println("\t create\n\t   to create launchConfig\n ")
	fmt.Println("\t delete\n\t   to delete launchConfig\n ")
	fmt.Println("\t list\n\t   to list the deployment\n ")
}

//DeploymentInfo for deployment info
func DeploymentInfo() {
	//lle := strings.Split(lle_Ip, ":")
	fmt.Println(color.BlueString("AWS Automata Micro LLE Jenkins running at "), color.GreenString("http:10.207.142.2:8080"))
	fmt.Println(color.BlueString("AWS Automata Micro Automata Server running at "), color.GreenString("http:10.207.142.2:8090"))
	fmt.Println(color.BlueString("AWS Automata Micro Automata DB running at "), color.GreenString("10.207.142.2:27017"))
}
