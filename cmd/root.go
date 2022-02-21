/*
Copyright © 2022 Chief Of State

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cosHost string
	cosPort int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cos-cli",
	Short: "Chief Of State Command Line Tool",
	Long:  `cos-cli is command line tool that helps send commands to a running CoS to manage the various read sides.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// print the error to the stand err
		fmt.Println(err)
		// exit the program
		os.Exit(1)
	}
}

func init() {
	// Let us define the persistent flags cosHost and cosPort that will be used for all sub commands
	pflags := rootCmd.PersistentFlags()
	pflags.StringVar(&cosHost, "cosHost", "", "CoS service host address")
	pflags.IntVar(&cosPort, "cosPort", 9000, "CoS service port")
	_ = rootCmd.MarkFlagRequired("cosHost")
	_ = rootCmd.MarkFlagRequired("costPort")
}
