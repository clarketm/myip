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
	-e, --ethernet		# Print (IPv4/IPv6) ethernet IP address.
	-p, --public		# Print (IPv4/IPv6) public IP address.
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
const VERSION = "v1.1.1"

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
			ipv4, ipv6 := getPrivateIP()
			if len(ipv4) > 0 {
				fmt.Printf("%s %v\n", bold("Ethernet (IPv4):"), ipv4)
			}
			if len(ipv6) > 0 {
				fmt.Printf("%s %v\n", bold("Ethernet (IPv6):"), ipv6)
			}
		}
		println()
		if public {
			ipv4, ipv6 := getPublicIP()
			if len(ipv4) > 0 {
				fmt.Printf("%s %v\n", bold("Public (IPv4):"), ipv4)
			}
			if len(ipv6) > 0 {
				fmt.Printf("%s %v\n", bold("Public (IPv6):"), ipv6)
			}
		}
		println()
	}
}

// getPublicIP () (string, string) - get public IP address
func getPublicIP() (string, string) {
	chV4 := make(chan string)
	chV6 := make(chan string)

	makeRequest := func(url string, ch chan string) {
		resp, err := http.Get(url)
		if err != nil {
			close(ch)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			close(ch)
			return
		}
		ch <- string(body)
	}

	go makeRequest("http://v4.ident.me/", chV4)
	go makeRequest("http://v6.ident.me/", chV6)

	ipv4Address := <-chV4
	ipv6Address := <-chV6

	return strings.TrimSpace(ipv4Address), strings.TrimSpace(ipv6Address)
}

// getPrivateIP () (string, string) - get private IP address(es)
func getPrivateIP() (string, string) {
	chEtherV4 := make(chan string, 10)
	chEtherV6 := make(chan string, 10)

	go getInterface("e", chEtherV4, chEtherV6)

	ipv4Addresses := []string{}
	ipv6Addresses := []string{}

	for {
		select {
		case ipV4, ok := <-chEtherV4:
			if !ok {
				chEtherV4 = nil
			} else {
				ipv4Addresses = append(ipv4Addresses, ipV4)
			}
		case ipV6, ok := <-chEtherV6:
			if !ok {
				chEtherV6 = nil
			} else {
				ipv6Addresses = append(ipv6Addresses, ipV6)
			}
		}
		if chEtherV4 == nil && chEtherV6 == nil {
			break
		}
	}

	return strings.TrimSpace(strings.Join(ipv4Addresses, ", ")), strings.TrimSpace(strings.Join(ipv6Addresses, ", "))
}

func getInterface(prefix string, chV4 chan string, chV6 chan string) {
	ifaces, err := net.Interfaces()
	if err != nil {
		close(chV4)
		close(chV6)
		return
	}
	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, prefix) {
			addrs, err := iface.Addrs()
			if err != nil {
				close(chV4)
				close(chV6)
				return
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						chV4 <- ipnet.IP.String()
					} else {
						chV6 <- ipnet.IP.String()
					}
				}
			}

		}
	}
	close(chV4)
	close(chV6)
}
