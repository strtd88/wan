package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "wan",
    Short: "wan helps keep track of your local dev stuf and career stuff too",
    Long:  `wan – Your Terminal Sidekick for Developers`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
