/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

NAME:
	myip – list IP addresses.

SYNOPSIS:
	myip [ opts... ]

OPTIONS:
	-h, --help		# Show usage.
	-a, --all		# Same as -e|--ethernet, -l|--loopback, -p|--public (default).
	-l, --loopback		# Print (IPv4/IPv6) (l)oopback IP address.
	-e, --ethernet		# Print (IPv4/IPv6) (e)thernet IP address.
	-p, --public		# Print (IPv4/IPv6) public IP address.
	-v, --version		# Show version number.

EXAMPLES:
	myip -a			# list all IP addresses.

*/

package main

import (
	"bufio"
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
const VERSION = "v1.4.4"

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
	loopback = v
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
var loopback bool
var public bool

// Globals
var statusCode int
var bold = color.New(color.Bold).SprintFunc()

// init () - initialize command-line flags
func init() {
	const (
		usageAll        = "Same as -e|--ethernet, -l|--loopback, -p|--public."
		usageVersion    = "Print version"
		defaultLoopback = false
		usageLoopback   = "Print ethernet IP address."
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

	// -l, --loopback
	flag.BoolVar(&loopback, "l", defaultLoopback, "")
	flag.BoolVar(&loopback, "loopback", defaultLoopback, usageLoopback)

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
	if !ethernet && !loopback && !public {
		statusCode = 0
		flag.Usage()
	} else {
		buffer := bufio.NewWriter(os.Stdout)
		defer buffer.Flush()

		if ethernet {
			buffer.WriteRune('\n')
			ipv4, ipv6 := getEthernetIP()
			if len(ipv4) > 0 {
				fmt.Fprintf(buffer, "%s %v\n", bold("Ethernet (IPv4):"), ipv4)
			}
			if len(ipv6) > 0 {
				fmt.Fprintf(buffer, "%s %v\n", bold("Ethernet (IPv6):"), ipv6)
			}
		}
		if loopback {
			buffer.WriteRune('\n')
			ipv4, ipv6 := getLoopbackIP()
			if len(ipv4) > 0 {
				fmt.Fprintf(buffer, "%s %v\n", bold("Loopback (IPv4):"), ipv4)
			}
			if len(ipv6) > 0 {
				fmt.Fprintf(buffer, "%s %v\n", bold("Loopback (IPv6):"), ipv6)
			}
		}
		if public {
			buffer.WriteRune('\n')
			ipv4, ipv6 := getPublicIP()
			if len(ipv4) > 0 {
				fmt.Fprintf(buffer, "%s %v\n", bold("Public (IPv4):"), ipv4)
			}
			if len(ipv6) > 0 {
				fmt.Fprintf(buffer, "%s %v\n", bold("Public (IPv6):"), ipv6)
			}
		}
		buffer.WriteRune('\n')
	}
}

// getPublicIP () (string, string) - get public IP address
func getPublicIP() (string, string) {
	chV4 := make(chan string, 1)
	chV6 := make(chan string, 1)

	var makeRequest func(urls []string, ch chan string)

	makeRequest = func(urls []string, ch chan string) {
		url, fallbackUrls := urls[0], urls[1:]
		resp, err := http.Get(url)
		if err != nil {
			if len(fallbackUrls) > 0 {
				go makeRequest(fallbackUrls, ch)
				return
			} else {
				close(ch)
				return
			}
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if len(fallbackUrls) > 0 {
				go makeRequest(fallbackUrls, ch)
				return
			} else {
				close(ch)
				return
			}
		}
		ch <- string(body)
	}

	go makeRequest([]string{"http://v4.ident.me/", "https://ipv4.travismclarke.com/"}, chV4)
	go makeRequest([]string{"http://v6.ident.me/", "https://ipv6.travismclarke.com/"}, chV6)

	ipv4Address := <-chV4
	ipv6Address := <-chV6

	return strings.TrimSpace(ipv4Address), strings.TrimSpace(ipv6Address)
}

// getEthernetIP () (string, string) - get private (e)thernet IP address(es)
func getEthernetIP() (string, string) {
	prefixes := map[string]bool{"e": true}
	includeLoopback := false

	return getPrivateIP(prefixes, includeLoopback)
}

// getLoopbackIP () (string, string) - get private (l)oopback IP address(es)
func getLoopbackIP() (string, string) {
	prefixes := map[string]bool{"l": true}
	includeLoopback := true

	return getPrivateIP(prefixes, includeLoopback)
}

// getPrivateIP () (string, string) - get private IP address(es)
func getPrivateIP(prefixes map[string]bool, includeLoopback bool) (string, string) {
	ipv4Addresses := []string{}
	ipv6Addresses := []string{}

	chV4 := make(chan string, 20)
	chV6 := make(chan string, 20)

	go getInterface(prefixes, chV4, chV6, includeLoopback)

	for {
		select {
		case ipv4, ok := <-chV4:
			if !ok {
				chV4 = nil
			} else {
				ipv4Addresses = append(ipv4Addresses, ipv4)
			}
		case ipv6, ok := <-chV6:
			if !ok {
				chV6 = nil
			} else {
				ipv6Addresses = append(ipv6Addresses, ipv6)
			}
		}
		if chV4 == nil && chV6 == nil {
			break
		}
	}

	return joinAddresses(ipv4Addresses), joinAddresses(ipv6Addresses)
}

func getInterface(prefix map[string]bool, chV4 chan string, chV6 chan string, includeLoopback bool) {
	ifaces, err := net.Interfaces()
	if err != nil {
		close(chV4)
		close(chV6)
		return
	}
	for _, iface := range ifaces {
		if _, ok := prefix[string(iface.Name[0])]; ok {
			addrs, err := iface.Addrs()
			if err != nil {
				close(chV4)
				close(chV6)
				return
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok {
					if !includeLoopback && ipnet.IP.IsLoopback() {
						continue
					}
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

func joinAddresses(addresses []string) string {
	return strings.TrimSpace(strings.Join(addresses, ", "))
}
