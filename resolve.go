// original: https://jvns.ca/blog/2022/02/01/a-dns-resolver-in-80-lines-of-go/

package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

// 13 DNS root nameservers https://www.iana.org/domains/root/servers

func main() {
	name := os.Args[1]
	if !strings.HasSuffix(name, ".") {
		name = name + "."
	}
	fmt.Println("Result:", resolve(name))
}
