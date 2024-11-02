package main

import (
	"flag"

	"github.com/tomas-sesztak/go-utils/argparse"
)

func main() {
	flag.Int("testInt", 8, "value of testing int")

	argparse.ParseArgs(false, true)
}
