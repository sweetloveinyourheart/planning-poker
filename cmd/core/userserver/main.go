package main

import (
	"time"

	"github.com/spf13/cobra"

	kittens_userserver "github.com/sweetloveinyourheart/exploding-kittens/cmd/exploding-kittens/services/userserver"
	kittens_utils "github.com/sweetloveinyourheart/exploding-kittens/cmd/exploding-kittens/utils"
	"github.com/sweetloveinyourheart/exploding-kittens/pkg/cmdutil"
)

const defaultShortDescription = "Exploding Kittens User Server"

func init() {
	// Always use UTC for generated timestamps
	time.Local = time.UTC
}

//go:generate go run github.com/sweetloveinyourheart/exploding-kittens/cmd/core/userserver generate

func main() {
	cmdutil.ServiceRootCmd.Short = defaultShortDescription
	commands := make([]*cobra.Command, 0)
	commands = append(commands, kittens_userserver.Command(cmdutil.ServiceRootCmd))

	commands = append(commands, kittens_utils.CheckCommand())

	cmdutil.InitializeService(commands...)
}
