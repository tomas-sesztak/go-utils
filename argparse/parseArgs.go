package argparse

import (
	"flag"
	"os"
)

// Parses args/flags using flag.Parse() and shows services.Help() if requirements are not met
func ParseArgs(needArgs bool, needFlags bool) {
	flag.Parse()
	if needArgs && flag.NArg() == 0 {
		Help()
		os.Exit(1)
	}
	if needFlags && flag.NFlag() == 0 {
		Help()
		os.Exit(1)
	}
}
