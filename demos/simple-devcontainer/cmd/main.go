package main

import (
	"flag"
	"fmt"

	"github.com/nych/docker-in-a-nutshell/demos/simple-devcontainer/internal/greeter"
)

var (
	nameFlag string
)

func init() {
	flag.StringVar(&nameFlag, "name", "John Doe", "tells the application who to greet")
	flag.Parse()
}

func main() {
	fmt.Println(greeter.Greetings(nameFlag))
}
