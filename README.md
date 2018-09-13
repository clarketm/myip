# [myip](https://godoc.org/github.com/clarketm/myip)

Command line utility for displaying public and private IP addresses.

[![release-badge](https://img.shields.io/github/release/clarketm/myip.svg)](https://github.com/clarketm/myip/releases)
[![circleci-badge](https://circleci.com/gh/clarketm/myip.svg?style=shield)](https://circleci.com/gh/clarketm/myip)
[![codacy-badge](https://api.codacy.com/project/badge/Grade/ce4b31a2e23b4959944d44f5add5234b)](https://www.codacy.com/app/clarketm/myip?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=clarketm/myip&amp;utm_campaign=Badge_Grade)

```shell
NAME:
    myip – list IP addresses.

SYNOPSIS:
    myip [ opts... ]

OPTIONS:
    -h, --help          # Show usage.
    -a, --all           # Same as -e|--ethernet, -l|--loopback, -p|--public (default).
    -l, --loopback      # Print (IPv4/IPv6) (l)oopback IP address.
    -e, --ethernet      # Print (IPv4/IPv6) (e)thernet IP address.
    -p, --public        # Print (IPv4/IPv6) public IP address.
    -v, --version       # Show version number.

EXAMPLES:
    myip -a             # list all IP addresses.
```

## Installation

#### Golang
```shell
$ go get -u github.com/clarketm/myip
```

#### Source (Mac/Linux)
```shell
# List of builds: https://github.com/clarketm/myip/releases/

$ BUILD=darwin_amd64.tar.gz     # Mac (64 bit)
# BUILD=linux_amd64.tar.gz      # Linux (64 bit)

$ BIN_DIR=/usr/local/bin        # `bin` install directory
$ mkdir -p $BIN_DIR

$ curl -L https://github.com/clarketm/myip/releases/download/v1.4.1/$BUILD | tar xz -C $BIN_DIR        # install
```

#### Source (Windows)
* https://github.com/clarketm/myip/releases/download/v1.4.1/windows_amd64.zip


## Usage

#### Basic usage 
```shell
$ myip


Ethernet (IPv4): 10.0.0.1
Ethernet (IPv6): fe80::22:a379:7bff:9092, fe80::aede:48ff:fe00:1122

Loopback (IPv4): 127.0.0.1
Loopback (IPv6): ::1, fe80::1

Public (IPv4): 71.217.233.188
```

---

You can see the full reference documentation for the **myip** package at [godoc.org](https://godoc.org/github.com/clarketm/myip), or through go's standard documentation system:
```bash
$ godoc -http=:6060

# Open browser to: "http://localhost:6060/pkg/github.com/clarketm/myip"  to view godoc.
```

## Related
* [public-ip](https://github.com/clarketm/public-ip) – A simple public IP address API

## License
Apache-2.0 &copy; [**Travis Clarke**](https://blog.travismclarke.com/)
