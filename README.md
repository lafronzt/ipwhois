# IP Whois Go Client

**This is an unofficial client for the [IPWhois.io API](https://ipwhois.io/).**

## Quick Start

Install and Use in Command Line:

```bash
go install github.com/lafronzt/ipwhois/cmd/ipwhois

ipwhois -ip 1.1.1.1
```

## Us as a Go Package

```go
package main

imports (
    "fmt"
    "github.com/lafronzt/ipwhois"
)

var c *ipwhois.Client

func init() {

   c = ipwhois.NewClient()

}

func main() {

    ip := "1.1.1.1"

    whois, err := c.GetIPDetails(IP, nil)
    if err != nil {
      println(err.Error())
      return
    }

    fmt.Printf("%+v\n", whois)
}
```

## To Use with the Pro Version of IPWhois.io use the following

```go
ipwhois.NewClientPro("api-key")
```

## Requirements

- Go >= 1.15
