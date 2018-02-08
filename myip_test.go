/*

Copyright 2018 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

*/

package main

import (
	"fmt"
	"regexp"
	"testing"
)

var ipRegex, _ = regexp.Compile(`(?:\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})|`)

func TestGetEthernetIP(t *testing.T) {
	ipv4, ipv6 := getEthernetIP()
	fmt.Println(ipRegex.MatchString(ipv4))
	fmt.Println(ipRegex.MatchString(ipv6))

	// Output:
	// true
	// true
}

func TestGetLoopbackIP(t *testing.T) {
	ipv4, ipv6 := getLoopbackIP()
	fmt.Println(ipRegex.MatchString(ipv4))
	fmt.Println(ipRegex.MatchString(ipv6))

	// Output:
	// true
	// true
}

func TestGetPublicI(t *testing.T) {
	ipv4, ipv6 := getPublicIP()
	fmt.Println(ipRegex.MatchString(ipv4))
	fmt.Println(ipRegex.MatchString(ipv6))

	// Output:
	// true
	// true
}
