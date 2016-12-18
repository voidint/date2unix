package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

const (
	defaultLayout = "2006/1/2 15:04:05"
)

func main() {
	app := cli.NewApp()
	app.Name = "date2unix"
	app.Version = fmt.Sprintf("0.0.2\ndate: %s\ncommit: %s", Date, Commit)
	app.Usage = "Convert date string to UNIX timestamp, the number of seconds elapsed since January 1, 1970."
	app.UsageText = "date2unix [global options] [arguments...]"
	app.Authors = []cli.Author{
		cli.Author{Name: "voidnt",
			Email: "voidint@126.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "layout, l",
			Value: defaultLayout,
			Usage: "date pattern",
		},
		cli.BoolFlag{
			Name:  "pretty, p",
			Usage: "pretty output",
		},
		cli.BoolFlag{
			Name:  "utc",
			Usage: "use UTC time zone",
		},
		cli.BoolFlag{
			Name:  "now",
			Usage: "output current time UNIX timestamp",
		},
	}

	app.Action = func(ctx *cli.Context) {
		var val string
		layout := ctx.String("layout")

		if ctx.Bool("now") {
			val = time.Now().Format(layout)
		} else if len(ctx.Args().First()) > 0 {
			val = ctx.Args().First()
		} else {
			bVal, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			val = strings.TrimSpace(string(bVal))
		}

		var loc *time.Location
		if ctx.Bool("utc") {
			loc = time.UTC
		} else {
			loc = time.Local
		}

		t, err := time.ParseInLocation(layout, val, loc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err, val)
			return
		}

		timestamp := t.Unix()

		if ctx.Bool("pretty") {
			fmt.Printf("%s => %d\n", val, timestamp)
		} else {
			fmt.Println(timestamp)
		}

	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
