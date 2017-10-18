package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/whytheplatypus/lights/commands"
)

var (
	// BuildTime is supplied by the compiler as the time at which the binary was built
	BuildTime string
	// GitCommit is supplied by the compiler as the most recent commit the binary was built from
	GitCommit string
	// Version is supplied by the compiler as the most recent git tag the binary was built from
	// defaults to 0.0.1
	Version string
	// VersionDescription is a modifier to Version that describes the binary build
	VersionDescription = "dev"

	Subcmds = CmdRegistry{}
)

func init() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Enable for verbose logging")

	flag.Parse()

	if verbose {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

type Runnable interface {
	Run(args []string) int
}

type CmdRegistry map[string]Runnable

func (c CmdRegistry) Register(name string, cmd Runnable) {
	if _, ok := c[name]; ok {
		panic(fmt.Errorf("subcommand %s already registered", name))
	}
	c[name] = cmd
}

func main() {

	Subcmds.Register("render", &commands.Renderer{})
	Subcmds.Register("clear", &commands.Clear{})

	if !flag.Parsed() {
		flag.Parse()
	}
	args := flag.Args()
	if len(args) > 0 {
		c, args := args[0], args[1:]
		cmd, ok := Subcmds[c]
		if !ok {
			fmt.Println("command not found")
			return
		}
		exitStatus := cmd.Run(args)
		os.Exit(exitStatus)
	} else {
		fmt.Print("Commands:")
		for key, _ := range Subcmds {
			fmt.Print(" ", key)
		}
		fmt.Println()
	}
}

// Description returns a string describing the binary build
// bartender <version>(-<VersionDescription>) ( :: commit - <GitCommit> [ :: built @ <BuildTime> ] )
func Description() string {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "bartender %s", Version)
	if VersionDescription != "" {
		fmt.Fprintf(&versionString, "-%s", VersionDescription)
	}

	if GitCommit != "" {
		fmt.Fprintf(&versionString, " :: commit - %s", GitCommit)
	}

	if BuildTime != "" {
		fmt.Fprintf(&versionString, " :: built @ %s", BuildTime)
	}

	return versionString.String()
}

func Short() string {
	if Version != "" {
		return strings.TrimPrefix(Version, "v")
	}

	if GitCommit != "" {
		return GitCommit
	}

	return ""
}
