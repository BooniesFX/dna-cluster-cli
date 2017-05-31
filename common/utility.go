package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

//PrettyPrint JSON output is nice and allows users to build on top.
func PrettyPrint(buf []byte) {
	var dst bytes.Buffer
	err := json.Indent(&dst, buf, "", "  ")
	checkErr("indenting json", err)
	fmt.Printf("%s", dst.String())
}
func out(m string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, m, a...)
}

func checkErr(doing string, err error) {
	if err != nil {
		out("error %s: %s\n", doing, err)
	}

}
func PrintError(c *cli.Context, err error, cmd string) {
	fmt.Println("Incorrect Usage:", err)
	fmt.Println("")
	cli.ShowCommandHelp(c, cmd)
}
