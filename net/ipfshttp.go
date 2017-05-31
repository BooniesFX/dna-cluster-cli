package net

import (
	//"DNA/common"
	//"DNA/common/log"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//cluster host
var (
	defaultClusterHost = fmt.Sprintf("127.0.0.1:%d", 9094)
	defaultIPFSHost    = fmt.Sprintf("127.0.0.1:%d", 8080)
	defaultTimeout     = 60
	defaultProtocol    = "http"
)

func out(m string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, m, a...)
}

func checkErr(doing string, err error) {
	if err != nil {
		out("error %s: %s\n", doing, err)
	}

}

func RequestIPFS(method, path string, body io.Reader, args ...string) *http.Response {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(defaultTimeout)*time.Second)
	defer cancel()

	u := defaultProtocol + "://" + defaultIPFSHost + path
	// turn /a/{param0}/{param1} into /a/this/that
	for i, a := range args {
		p := fmt.Sprintf("{param%d}", i)
		u = strings.Replace(u, p, a, 1)
	}
	u = strings.TrimSuffix(u, "/")

	fmt.Printf("%s: %s\n", method, u)

	r, err := http.NewRequest(method, u, body)
	checkErr("creating request", err)
	r.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(r)
	checkErr(fmt.Sprintf("performing request to %s", defaultIPFSHost), err)

	return resp
}

func FormatIPFSResponse(c *cli.Context, r *http.Response) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	checkErr("reading body", err)
	prettyPrint(body)

}

// JSON output is nice and allows users to build on top.
func prettyPrint(buf []byte) {
	var dst bytes.Buffer
	err := json.Indent(&dst, buf, "", "  ")
	checkErr("indenting json", err)
	fmt.Printf("%s", dst.String())
}
