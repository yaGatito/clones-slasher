package terminal

import (
	"fmt"
)

type Command struct {
	Paths          []string
	ReportPath     string
	CloneCheck     bool
	NamesakesCheck bool
}

const (
	WatchClones    string = "-c"
	WatchNamesakes string = "-n"
	Paths          string = "-p" // Multiple path support
	Output         string = "-o"
)

// If output path not provided - the first path root folder will be used to save report.
func ParseArgs(args []string) Command {
	fmt.Println(args)

	cmd := Command{
		Paths: make([]string, 0),
	}
	var watchPathArgs bool
	for i, arg := range args {
		switch arg {
		case Paths:
			watchPathArgs = true

		case WatchClones:
			cmd.CloneCheck = true
			watchPathArgs = false

		case WatchNamesakes:
			cmd.NamesakesCheck = true
			watchPathArgs = false

		case Output:
			var outputPath string
			if i+1 <= len(args)-1 {
				outputPath = args[i+1]
			}
			cmd.ReportPath = outputPath
			watchPathArgs = false

		default:
			if watchPathArgs {
				cmd.Paths = append(cmd.Paths, arg)
			} else {
				watchPathArgs = false
				fmt.Println("unparsed data:" + arg)
			}
		}
	}

	return cmd
}
