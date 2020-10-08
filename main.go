package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	File          string   `short:"f" long:"file" description:"File path should be filtered" required:"true"`
	Host          []string `long:"host" description:"Include host filter. (Can be use multiple times)"`
	NotLengths    []int64  `long:"ws" description:"Exclude size option. (Can be use multiple times)"`
	NotWords      []int64  `long:"ww" description:"Exclude words option. (Can be use multiple times)"`
	NotLines      []int64  `long:"wl" description:"Exclude lines option. (Can be use multiple times)"`
	NotCode       []int64  `long:"wc" description:"Exclude status code option. (Can be use multiple times)"`
	ShouldLengths []int64  `long:"fs" description:"Include size option. (Can be use multiple times)"`
	ShouldWords   []int64  `long:"fw" description:"Include words option. (Can be use multiple times)"`
	ShouldLines   []int64  `long:"fl" description:"Include lines option. (Can be use multiple times)"`
	ShouldCode    []int64  `long:"fc" description:"Include status code option. (Can be use multiple times)"`
}

var opts Options
var parser = flags.NewParser(&opts, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case *flags.Error:
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	record, err := LoadResults(opts.File)
	if err != nil {
		FatalErr(err)
	}

	filters := GetFilters(opts)

	results := ApplyFilters(record.Results, filters...)

	PrintOut(results)

}
