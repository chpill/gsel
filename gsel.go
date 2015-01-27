package main

import (
	"fmt"
	"github.com/coddingtonbear/go-jsonselect"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gsel"
	app.Usage = "JSON processor with pseudo-css syntax"

	app.Action = func(c *cli.Context) {
		selector := c.Args()[0]
		var jsonFile []byte
		var err error

		if name := c.String("file"); name != "" {
			jsonFile, err = ioutil.ReadFile(name)

		} else {
			jsonFile, err = ioutil.ReadAll(os.Stdin)
		}
		check(err)

		jsonString := string(jsonFile)

		if c.Bool("debug") {
			jsonselect.EnableLogger()
		}

		parser, _ := jsonselect.CreateParserFromString(jsonString)
		results, _ := parser.GetValues(selector)

		fmt.Print(results)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "",
			Usage: "Input file containing valid json. Gsel " +
				"will ignore standard input if this flag is " +
				"defined",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "print the very verbose debug logs of json-select",
		},
	}

	app.Run(os.Args)
}
