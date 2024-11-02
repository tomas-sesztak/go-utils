package main

import (
	"flag"

	"github.com/tomas-sesztak/go-utils/argparse/services"
)

func main() {
	flag.Int("testInt", 8, "value of testing int")

	services.ParseArgs(false, true)
}
