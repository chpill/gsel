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
	app.Usage = "selector infile"

	app.Action = func(c *cli.Context) {
		jsonFile, err := ioutil.ReadFile(c.Args()[1])
		check(err)

		jsonString := string(jsonFile)
		selector := c.Args()[0]

		if c.Bool("debug") {
			jsonselect.EnableLogger()
		}

		parser, _ := jsonselect.CreateParserFromString(jsonString)
		results, _ := parser.GetValues(selector)

		fmt.Print(results)
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "print the very verbose debug logs of json-select",
		},
	}

	app.Run(os.Args)
}
