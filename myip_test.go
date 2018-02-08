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

var ipv4Regex, _ = regexp.Compile(`(?:(?:[0-9]{1,3}\.){1,3}[0-9]{1,3},? ?)+|`)
var ipv6Regex, _ = regexp.Compile(`(?:(?:[A-Fa-f0-9]{0,4}:){1,3}[A-Fa-f0-9]{1,4},? ?)+|`)

func TestGetEthernetIP(t *testing.T) {
	ipv4, ipv6 := getEthernetIP()
	fmt.Println(ipv4Regex.MatchString(ipv4))
	fmt.Println(ipv6Regex.MatchString(ipv6))

	// Output:
	// true
	// true
}

func TestGetLoopbackIP(t *testing.T) {
	ipv4, ipv6 := getLoopbackIP()
	fmt.Println(ipv4Regex.MatchString(ipv4))
	fmt.Println(ipv6Regex.MatchString(ipv6))

	// Output:
	// true
	// true
}

func TestGetPublicI(t *testing.T) {
	ipv4, ipv6 := getPublicIP()
	fmt.Println(ipv4Regex.MatchString(ipv4))
	fmt.Println(ipv6Regex.MatchString(ipv6))

	// Output:
	// true
	// true
}
