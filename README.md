# [myip](https://godoc.org/github.com/clarketm/myip)

Command line utility for displaying public and private IP addresses.

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

You can see the full reference documentation for the **myip** package at [godoc.org](https://godoc.org/github.com/clarketm/myip), or through go's standard documentation system:
```bash
$ godoc -http=:6060

# Open browser to: "http://localhost:6060/pkg/github.com/clarketm/myip"  to view godoc.
```
