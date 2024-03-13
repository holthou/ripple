package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"

	"github.com/rubblelabs/ripple/data"
	"github.com/rubblelabs/ripple/terminal"
	"github.com/rubblelabs/ripple/websockets"
)

const usage = `Usage: lines [ripple address] [options]

Examples:

lines rBxy23n7ZFbUpS699rFVj1V9ZVhAq6EGwC
	Show all trust lines for account rBxy23n7ZFbUpS699rFVj1V9ZVhAq6EGwC

Options:`

var (
	host = flag.String("host", "wss://s1.ripple.com:443", "websockets host")
)

func showUsage() {
	fmt.Println(usage)
	flag.PrintDefaults()
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {

	err := fmt.Errorf("1212qw")
	glog.Errorln("RPC client reconnect failed", "err", err)

	if len(os.Args) == 1 {
		showUsage()
	}
	flag.CommandLine.Parse(os.Args[2:])

	remote, err := websockets.NewRemote(*host)
	checkErr(err)
	account, err := data.NewAccountFromAddress(os.Args[1])
	checkErr(err)
	result, err := remote.AccountLines(*account, "closed", "")
	checkErr(err)
	// fmt.Println(*result.LedgerSequence) //TODO: wait for nikb fix
	fmt.Println(len(result.Lines))
	for _, line := range result.Lines {
		terminal.Println(line, terminal.Default)
	}
}
