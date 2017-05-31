package info

import (
	"dna-cluster-cli/common"
	"dna-cluster-cli/net"
	"github.com/urfave/cli"
)

const (
	formatNone = iota
	formatFileID
	formatFilePath
	formatVersion
	formatError
)

func versionAction(c *cli.Context) (err error) {
	resp := net.RequestIPFS("GET", "/version", nil)
	net.FormatIPFSResponse(c, resp)
	return nil
}

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:      "version",
		Usage:     "show ipfs and cluster information",
		UsageText: "retrive the version of ipfs and cluster.",
		ArgsUsage: "[args]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "version, v",
				Usage: "print ipfs and cluster daemon version",
			},
		},
		Action: versionAction,
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			common.PrintError(c, err, "info")
			return cli.NewExitError("", 1)
		},
	}
}
