package main

import (
	"github.com/BrendanMcCallum/DslReports-CLI"
	"os"
	"fmt"
	"github.com/urfave/cli"
	"strconv"
	"time"
)

// A struct that contains all the runtime information needed to perform and compile results
// of the speedtest
var R *dslr.DslrAppRuntime


// initializes the commandline client, things that are done in initialization:
// display a nice title for the command line utility
// registers all functions defined in the dslr package
func initialize() {
  APPversion := 0.1;
  APPDate := "2016-05-30"
  println("DSLReports.com CLI v" + strconv.FormatFloat(APPversion, 'f', 1, 64) + " - " + APPDate)
	R = &dslr.DslrAppRuntime{}
	R.AppVersion = APPversion;
	dslr.Init(R)
	M := dslr.Dslrmethod{Name: "VerifyClient", Action: dslr.VerifyClient}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "SetupClient", Action: dslr.SetupClient}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "PerformSpeedTest", Action: dslr.PerformSpeedTest}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "StartSpeedTestDownload", Action: dslr.PerformDownloadSpeedTest}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "StartSpeedTestUpload", Action: dslr.PerformUploadSpeedTestAlpha}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "Cleanup", Action: dslr.Cleanup}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "Results", Action: dslr.OutputResults}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "APIauth", Action: dslr.APIauth}
	dslr.Register(R, M)
	M = dslr.Dslrmethod{Name: "APIservers", Action: dslr.APIservers}
	dslr.Register(R, M)
	// Register Alpha features
	M = dslr.Dslrmethod{Name: "PerformUploadSpeedTestAlpha", Action: dslr.PerformUploadSpeedTestAlpha}
	dslr.Register(R, M);
	M = dslr.Dslrmethod{Name: "PushResultstoServer", Action: dslr.PushResultstoServer };
	dslr.Register(R, M);

}


// registers all required flags and their respective handler variables in the runtime
func registerFlags(app *cli.App){
	app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name:  "debug, d",
            Usage: "Enables debug mode",
        },
        cli.StringFlag{
            Name:  "output, o",
            Usage: "Specify type of output . 'json' and 'csv' are currently supported.",
            Value: "default",
            Destination: &R.UserFlags.Output,
        },
        cli.StringFlag{
	      	Name:        "down",
	      	Value:       "12",
	      	Usage:       "Number of streams to use for download",
	      	Destination: &R.UserFlags.DownloadStreams,
	    },
	    cli.StringFlag{
	      	Name:        "up",
	      	Value:       "12",
	      	Usage:       "Number of streams to use for up",
	      	Destination: &R.UserFlags.UploadStreams,
	    },
    }
}


// this function registers all commands with their respective handler functions in the runtime.
// we've restricted the client to only run commands that have been registered in the innitialize phase
// this allows for a modular architecture, if you don't like an implementation of a function simply
// register another one in the initialize phase then call it here using the dslr.Run() call.
func registerCommands(app *cli.App){
	app.Commands = []cli.Command{
		{
			Name:    "setup",
			Aliases: []string{"s"},
			Usage:   "setup the client",
			Action: func(c *cli.Context) {
				dslr.Run(R, "SetupClient")
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run the test",
			SkipFlagParsing: false,
			Flags: []cli.Flag{
	        	cli.BoolFlag{Name: "debug, d"},
	      	},
			Action: func(c *cli.Context) error {
				c.Command.VisibleFlags()
				if c.Bool("debug")  {
					R.DebugEnabled = true;
      				fmt.Println("Debug mode enabled");
      			}      			
				dslr.Run(R, "APIauth")
				dslr.Run(R, "APIservers")
				dslr.Run(R, "StartSpeedTestDownload")
				dslr.Run(R, "StartSpeedTestUpload")
				dslr.Run(R, "PushResultstoServer")
				dslr.Run(R, "Results")
				return nil;
			},
		},
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "run the test",
			Action: func(c *cli.Context) error {
				dslr.Run(R, "APIauth")
				dslr.Run(R, "APIservers")
				dslr.Run(R, "StartSpeedTestUpload")
				return nil;
			},
		},
		{
			Name:    "download",
			Aliases: []string{"u"},
			Usage:   "run the test",
			Action: func(c *cli.Context) error {
				dslr.Run(R, "APIauth")
				dslr.Run(R, "APIservers")
				dslr.Run(R, "StartSpeedTestDownload")
				return nil;
			},
		},
		{
			Name:    "uploaddemo",
			Aliases: []string{"u"},
			Usage:   "run the test",
			Flags: []cli.Flag{
	        	cli.BoolFlag{Name: "debug, d"},
	      	},
			Action: func(c *cli.Context) error {
				if c.Bool("debug")  {
					R.DebugEnabled = true;
      				fmt.Println("Debug mode enabled");
      			}
				dslr.Run(R, "APIauth")
				dslr.Run(R, "APIservers")
				dslr.Run(R, "PerformUploadSpeedTestAlpha")
				return nil;
			},
		},
		{
			Name:    "testpush",
			Aliases: []string{"u"},
			Usage:   "test push results to mothership",
			Flags: []cli.Flag{
	        	cli.BoolFlag{Name: "debug, d"},
	      	},
			Action: func(c *cli.Context) error {
				if c.Bool("debug")  {
					R.DebugEnabled = true;
      				fmt.Println("Debug mode enabled");
      			}
      			dslr.Run(R, "APIauth")
      			dslr.Run(R, "APIservers")
      			dslr.Run(R, "StartSpeedTestDownload")
				dslr.Run(R, "StartSpeedTestUpload")
				dslr.Run(R, "PushResultstoServer")
				dslr.Run(R, "Results")
				return nil;
			},
		},
	}
}


// main
func main() {
	initialize()
	app := cli.NewApp()
	app.Name = "Dslrcli"
	app.Usage = "Test network speed"
	app.Compiled = time.Now()
	app.Version = strconv.FormatFloat(R.AppVersion, 'f', 1, 64)
	app.Authors = []cli.Author{
		cli.Author{
		  Name:  "Abdullah Irfan",
		  Email: "hello@abdullahirfan.com",
		},
	}
	registerFlags(app);
	app.Action = func(c *cli.Context) error {
	    dslr.Run(R, "APIauth")
		dslr.Run(R, "APIservers")
		dslr.Run(R, "StartSpeedTestUpload")
	    return nil;
	}
	registerCommands(app);
	app.Run(os.Args)
}

