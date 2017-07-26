/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

NAME:
	myip – list IP addresses.

SYNOPSIS:
	myip [ opts... ]

OPTIONS:
	-h, --help		# Show usage.
	-a, --all		# Same as -e, -p (default).
	-e, --ethernet		# Print ethernet IP address.
	-p, --public		# Print public IP address.
	-v, --version		# Show version number.

EXAMPLES:
	myip -a			# list all IP addresses.

*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// VERSION - current version number
const VERSION = "v1.0.0"

// allFlag bool
type allFlag bool

func (a *allFlag) IsBoolFlag() bool {
	return true
}

func (a *allFlag) String() string {
	return "false"
}

func (a *allFlag) Set(value string) error {
	v, _ := strconv.ParseBool(value)
	ethernet = v
	public = v
	return nil
}

// versionFlag bool
type versionFlag bool

func (v *versionFlag) IsBoolFlag() bool {
	return true
}

func (v *versionFlag) String() string {
	return "false"
}

func (v *versionFlag) Set(value string) error {
	println()
	fmt.Printf("%s %v", bold("Version:"), VERSION)
	println()
	os.Exit(0)
	return nil
}

// Flags
var all allFlag
var version versionFlag
var ethernet bool
var public bool

// Globals
var statusCode int
var bold = color.New(color.Bold).SprintFunc()

// init () - initialize command-line flags
func init() {
	const (
		usageAll        = "Same as --ethernet, --public."
		usageVersion    = "Print version"
		defaultEthernet = false
		usageEthernet   = "Print ethernet IP address."
		defaultPublic   = false
		usagePublic     = "Print public IP address."
	)
	// -a, --all
	flag.Var(&all, "a", "")
	flag.Var(&all, "all", usageAll)

	// -e, --ethernet
	flag.BoolVar(&ethernet, "e", defaultEthernet, "")
	flag.BoolVar(&ethernet, "ethernet", defaultEthernet, usageEthernet)

	// -p, --public
	flag.BoolVar(&public, "p", defaultPublic, "")
	flag.BoolVar(&public, "public", defaultPublic, usagePublic)

	// -v, --version
	flag.Var(&version, "v", "")
	flag.Var(&version, "version", usageVersion)

	// Usage
	flag.Usage = func() {
		println()
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		println()
		os.Exit(statusCode)
	}
}

// main ()
func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		all.Set("true") // 1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
	}
	if !ethernet && !public {
		statusCode = 0
		flag.Usage()
	} else {
		println()
		if ethernet {
			fmt.Printf("%s %v\n", bold("Ethernet:"), getPublicIP())
		}
		if public {
			fmt.Printf("%s %v\n", bold("Public:"), getPrivateIP())
		}
		println()
	}
}

// getPublicIP () string - get public IP address
func getPublicIP() string {
	var ipAddress string

	checkError := func(err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, "There was an error retreiving public IP: ", err)
			os.Exit(1)
		}
	}

	resp, err := http.Get("http://diagnostic.opendns.com/myip")
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	ipAddress = string(body)

	return strings.TrimSpace(ipAddress)
}

// getPrivateIP () string - get private IP address(es)
func getPrivateIP() string {
	ipAddresses := []string{}

	checkError := func(err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, "There was an error retreiving private IP: ", err)
			os.Exit(1)
		}
	}

	ifaces, err := net.Interfaces()
	checkError(err)
	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "e") {
			addrs, err := iface.Addrs()
			checkError(err)
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ipAddresses = append(ipAddresses, ipnet.IP.String())
					}
				}
			}
		}
	}
	return strings.TrimSpace(strings.Join(ipAddresses, ","))
}
