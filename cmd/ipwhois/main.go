package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/lafronzt/ipwhois"
)

var (
	IP *string = flag.String("ip", "", "IP address to lookup")
	c  *ipwhois.Client
)

func init() {
	flag.Usage = func() {
		println("Usage: ipwhois [options]")
		println("Options:")
		flag.PrintDefaults()
	}

	c = ipwhois.NewClient()

}

func main() {
	flag.Parse()

	if *IP == "" {
		flag.Usage()
		return
	}

	whois, err := c.GetIPDetails(IP, nil)
	if err != nil {
		println(err.Error())
		return
	}

	v := reflect.ValueOf(*whois)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	for i := 0; i < len(values); i++ {
		fmt.Printf("%s: %v\n", v.Type().Field(i).Name, values[i])
	}
}
