package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

func main() {
	name := os.Args[1]
	if !strings.HasSuffix(name, ".") {
		name = name + "."
	}
	fmt.Println("Result:", resolve(name))
}
