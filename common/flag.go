package common

import (
	"github.com/urfave/cli"
)

var (
	defaultHost = "127.0.0.1"
)

func IpFlag() cli.Flag {
	return cli.StringFlag{
		Name:        "ip",
		Usage:       "node's ip address",
		Value:       defaultHost,
		Destination: &defaultHost,
	}
}
