package main

import (
	"DNA/common/log"
	"DNA/crypto"
	"dna-cluster-cli/cmd/info"
	"dna-cluster-cli/common"
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"sort"
	"time"
)

//Version of app,see more info in version.go
var Version string

func main() {
	var path = "./Log/"
	log.CreatePrintLog(path)
	//region crypto&rand is needed in DNA,could be commented
	crypto.SetAlg("P256R1")
	//seed transaction nonce
	rand.Seed(time.Now().UnixNano())
	//endregion crypto&rand
	app := cli.NewApp()
	app.Name = "dna-cluster-cli"
	app.Version = Version
	app.HelpName = "dna-cluster-cli"
	app.Usage = "command line tool for ipfs cluster network"
	app.UsageText = "dna-cluster-cli [global options] command [command options] [args]"
	app.HideHelp = false
	app.HideVersion = false
	//global options
	app.Flags = []cli.Flag{
		common.IpFlag(),
	}
	//commands
	app.Commands = []cli.Command{
		*info.NewCommand(),
		//*swarm.NewCommand(),
		//*repo.NewCommand(),
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
