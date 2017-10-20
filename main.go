package main

import (
	"bytes"
	"errors"
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

	// Subcmds are the possible excutable sub-commands for this program
	Subcmds = CmdRegistry{}
)

func main() {
	Subcmds.Register("render", &commands.Renderer{})
	Subcmds.Register("clear", &commands.Clear{})
	Subcmds.Register("version", RunFunc(func(args []string) error {
		fmt.Println(Description())
		return nil
	}))

	var verbose bool
	cmdFlag := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cmdFlag.BoolVar(&verbose, "v", false, "Enable for verbose logging")
	if err := cmdFlag.Parse(os.Args[1:]); err != nil {
		Subcmds.Usage()
		os.Exit(1)
	}

	if verbose {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	args := flag.Args()

	if err := Subcmds.Run(args); err != nil {
		Subcmds.Usage()
		os.Exit(1)
	}

	os.Exit(0)
}

type Runnable interface {
	Run(args []string) error
}

type RunFunc func(args []string) error

func (r RunFunc) Run(args []string) error {
	return r(args)
}

type CmdRegistry map[string]Runnable

func (c CmdRegistry) Register(name string, cmd Runnable) {
	if _, ok := c[name]; ok {
		panic(fmt.Errorf("subcommand %s already registered", name))
	}
	c[name] = cmd
}

var ErrCommandNotFound = errors.New("no subcommand registered with that name")
var ErrNoSubcommandSupplied = errors.New("no subcommand supplied")

func (c CmdRegistry) Run(args []string) error {
	if len(args) < 1 {
		return ErrNoSubcommandSupplied
	}
	cn, args := args[0], args[1:]
	cmd, ok := Subcmds[cn]
	if !ok {
		return ErrCommandNotFound
	}
	return cmd.Run(args)
}

func (c CmdRegistry) Usage() {
	fmt.Println("Subcommands: ")
	for key, _ := range c {
		fmt.Printf("%s\n", key)
	}
}

// Description returns a string describing the binary build
// bartender <version>(-<VersionDescription>) ( :: commit - <GitCommit> [ :: built @ <BuildTime> ] )
func Description() string {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "lights %s", Version)
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
