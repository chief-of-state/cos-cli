/*
Copyright © 2022 Chief Of State

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	readSideID  string
	shardNumber int
)

// readsideCmd represents the readside command
var readsideCmd = &cobra.Command{
	Use:   "readside",
	Short: "readside command helps manage the various read-sides",
	Long: `With readside one can:

- List read sides' offsets per shard and across the whole CoS cluster,
- Skip offset per shard and across the whole CoS cluster,
- Pause and Resume read sides per shard and across the whole CoS cluster,
- Restart read sides per shard and across the whole CoS cluster.`,
	Run: readSideRun,
}

// readSideRun helps run the readsideCmd
func readSideRun(cmd *cobra.Command, args []string) {
	for _, x := range args {
		fmt.Println(x)
	}
}

func init() {
	// let us define the flags
	flags := readsideCmd.PersistentFlags()
	flags.StringVar(&readSideID, "id", "", "the read side unique identifier")
	flags.IntVar(&shardNumber, "shard-number", -1, "the CoS cluster shard number")
	_ = readsideCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(readsideCmd)
}
