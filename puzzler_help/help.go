package puzzler_help

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func Help() {
	var helpArgs string
	if len(flag.Args()) > 0 {
		helpArgs = flag.Args()[0]
	} else {
		helpArgs = ""
	}
	switch helpArgs {

	case "flags":
		openHelptxt("puzzler_help/help_flags.txt")
	default:
		openHelptxt("puzzler_help/help.txt")
	}
}

func openHelptxt(filename string) {
	helpfile, err := os.Open(filename)
	defer helpfile.Close()
	if err != nil {
		panic(err.Error())
	}
	helptxt, err := ioutil.ReadAll(helpfile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(helptxt))
	return
}
