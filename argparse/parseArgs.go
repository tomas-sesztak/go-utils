package argparse

import (
	"flag"
)

// Parses args/flags using flag.Parse() and shows services.Help() if requirements are not met
func ParseArgs(needArgs bool, needFlags bool) {
	flag.Parse()
	if needArgs && flag.NArg() == 0 {
		Help()
	}
	if needFlags && flag.NFlag() == 0 {
		Help()
	}
}
