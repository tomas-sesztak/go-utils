package services

import (
	"flag"
	"os"
	"fmt"
)

func Help() {
	fmt.Fprintf(os.Stderr, "Usage of %s", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
