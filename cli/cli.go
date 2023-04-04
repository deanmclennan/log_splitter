package cli

import (
	"fmt"
	"os"

	"dmclennan.com/splitter/splitter"
	// urfave

	"github.com/urfave/cli"
)

// Green color for terminal
const green = "\033[32m"

// Reset color for terminal
const reset = "\033[0m"

func Cli() {
	app := cli.NewApp()
	app.Name = "Log File Splitter"
	app.Usage = "Sorts logs by processes and splits them into separate files."
	app.Version = "1.0"

	app.Authors = []cli.Author{
		{
			Name: "Dean McLennan",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log, l",
			Value: "logs",
			Usage: "Log file to split",
		},

		cli.StringFlag{
			Name:  "root, r",
			Value: "logOutput",
			Usage: "Root folder to store logs",
		},
	}

	app.Action = func(c *cli.Context) error {

		fmt.Println(green + "--------------------------------------------")
		fmt.Println("Log File Splitter" + " v1.0")
		fmt.Println("--------------------------------------------")

		fmt.Println(" ")
		fmt.Println(" ")

		fmt.Println("Splitting logs...")

		// split logs
		splitter.Split_logs(c.String("log"), c.String("root"))

		// if process was successful
		fmt.Println("Logs split successfully!")

		// Logs location message
		fmt.Println("Logs can be found in: " + c.String("root"))

		fmt.Println("--------------------------------------------" + reset)
		fmt.Println(reset)

		return nil
	}

	app.Run(os.Args)

}
