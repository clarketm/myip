/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

NAME:
	myip – list IP addresses.

SYNOPSIS:
	myip [ opts... ]

OPTIONS:
	-a, --all		# Same as -e, -p.
	-e, --ethernet		# Print ethernet IP address.
	-p, --public		# Print public IP address.

EXAMPLES:
	myip -a			# list all IP addresses (1 per line).

*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var all bool
var ethernet bool
var public bool

func init() {
	const (
		defaultAll      = false
		usageAll        = "Same as --ethernet, --public."
		defaultEthernet = false
		usageEthernet   = "Print ethernet IP address."
		defaultPublic   = false
		usagePublic     = "Print public IP address."
	)
	// -a, --all
	flag.BoolVar(&all, "a", defaultAll, "")
	flag.BoolVar(&all, "all", defaultAll, usageAll)

	// -e, --ethernet
	flag.BoolVar(&ethernet, "e", defaultEthernet, "")
	flag.BoolVar(&ethernet, "ethernet", defaultEthernet, usageEthernet)

	// -p, --public
	flag.BoolVar(&public, "p", defaultPublic, "")
	flag.BoolVar(&public, "public", defaultPublic, usagePublic)

	// Usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	flag.Parse()

	if all {
		ethernet = true
		public = true
	}

	bold := color.New(color.Bold).SprintFunc()
	println()
	if ethernet {
		fmt.Printf("%s %v", bold("Ethernet:"), getPublicIP())
	}
	if public {
		fmt.Printf("%s %v", bold("Public:"), getPrivateIP())
	}
	println()
}

func getPublicIP() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdOut, err = exec.Command("dig", "+short", "@resolver1.opendns.com", "myip.opendns.com").Output()
	if err != nil {
		// fmt.Fprintln(os.Stderr, "There was an error retreiving public IP: ", err)
		// os.Exit(1)
	}
	return string(cmdOut)
}

func getPrivateIP() string {
	var (
		cmdOutSlice []string
		cmdOut      []byte
		err         error
	)

	for i := 0; ; i++ {
		cmdOut, err = exec.Command("ipconfig", "getifaddr", fmt.Sprintf("en%s", strconv.Itoa(i))).Output()
		if err != nil {
			// fmt.Fprintln(os.Stderr, "There was an error retreiving private IP: ", err)
			// os.Exit(1)
			break
		}
		cmdOutSlice = append(cmdOutSlice, string(cmdOut))
	}

	return strings.Join(cmdOutSlice, ",")
}
