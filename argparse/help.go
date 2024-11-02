package argparse

import (
	"flag"
	"os"
	"fmt"
)

// show the standard help/usage message
func Help() {
	fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
	flag.PrintDefaults()
}
