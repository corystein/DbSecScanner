//go:generate goversioninfo

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/clagraff/argparse"
	"golang.org/x/sys/windows/svc"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func callback(p *argparse.Parser, ns *argparse.Namespace, leftovers []string, err error) {
	if err != nil {
		switch err.(type) {
		case argparse.ShowHelpErr, argparse.ShowVersionErr:
			// For either ShowHelpErr or ShowVersionErr, the parser has already
			// displayed the necessary text to the user. So we end the program
			// by returning.
			return
		default:
			fmt.Println(err, "\n")
			p.ShowHelp()
		}

		return // Exit program
	}

	name := ns.Get("name").(string)
	upper := ns.Get("upper").(string) == "true"

	if upper == true {
		name = strings.ToUpper(name)
	}

	fmt.Printf("Hello, %s!\n", name)
	if len(leftovers) > 0 {
		fmt.Println("\nUnused args:", leftovers)
	}
}

func main() {
	const svcName = "DbSecScanner"
	const svcDescription = "Database Security Scanner"
	const version = "1.0.0.0"

	p := argparse.NewParser("Output a friendly greeting", callback).Version(version)
	p.AddHelp().AddVersion() // Enable help and version flags

	fmt.Printf("Version: %s\n", version)

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if !isIntSess {
		runService(svcName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		runService(svcName, true)
		return
	case "install":
		err = installService(svcName, svcDescription)
	case "remove":
		err = removeService(svcName)
	case "start":
		err = startService(svcName)
	case "stop":
		err = controlService(svcName, svc.Stop, svc.Stopped)
	case "pause":
		err = controlService(svcName, svc.Pause, svc.Paused)
	case "continue":
		err = controlService(svcName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
	return
}
