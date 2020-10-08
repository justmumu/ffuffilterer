package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ffuf/ffuf/pkg/output"
	"github.com/olekukonko/tablewriter"
)

func PrintOut(res []output.Result) {
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Host", "Scheme", "Status Code", "Size", "Lines", "Words", "Url"})

	for _, r := range res {
		scheme := strings.ToUpper(strings.Split(r.Url, ":")[0])
		table.Append([]string{green(fmt.Sprintf("[%s]", r.Host)),
			cyan(fmt.Sprintf("[%s]", scheme)),
			magenta(fmt.Sprintf("[%d]", r.StatusCode)),
			fmt.Sprintf("%d", r.ContentLength),
			fmt.Sprintf("%d", r.ContentLines),
			fmt.Sprintf("%d", r.ContentWords),
			r.Url,
		})
	}
	table.Render()
}

func Warning(warning string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	fmt.Printf("%s %s.\n", yellow("[Warning]:"), white(warning))
}

func Err(err error) {
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	fmt.Printf("%s %s.\n", red("[Error]:"), white(err.Error()))
}

func FatalErr(err error) {
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	fmt.Printf("%s %s.\n", red("[Error]:"), white(err.Error()))
	os.Exit(1)
}
