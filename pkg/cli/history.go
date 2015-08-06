// Copyright (C) 2015 Scaleway. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package cli

import (
	"github.com/scaleway/scaleway-cli/pkg/commands"
	"github.com/scaleway/scaleway-cli/vendor/github.com/Sirupsen/logrus"
)

var cmdHistory = &Command{
	Exec:        runHistory,
	UsageLine:   "history [OPTIONS] IMAGE",
	Description: "Show the history of an image",
	Help:        "Show the history of an image.",
}

func init() {
	cmdHistory.Flag.BoolVar(&historyNoTrunc, []string{"-no-trunc"}, false, "Don't truncate output")
	cmdHistory.Flag.BoolVar(&historyQuiet, []string{"q", "-quiet"}, false, "Only show numeric IDs")
	cmdHistory.Flag.BoolVar(&historyHelp, []string{"h", "-help"}, false, "Print usage")
}

// Flags
var historyNoTrunc bool // --no-trunc flag
var historyQuiet bool   // -q, --quiet flag
var historyHelp bool    // -h, --help flag

func runHistory(cmd *Command, rawArgs []string) {
	if historyHelp {
		cmd.PrintUsage()
	}
	if len(rawArgs) != 1 {
		cmd.PrintShortUsage()
	}

	args := commands.HistoryArgs{
		Quiet:   historyQuiet,
		NoTrunc: historyNoTrunc,
		Image:   rawArgs[0],
	}
	ctx := cmd.GetContext(rawArgs)
	err := commands.RunHistory(ctx, args)
	if err != nil {
		logrus.Fatalf("Cannot execute 'history': %v", err)
	}
}
