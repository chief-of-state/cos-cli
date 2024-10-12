/*
 * MIT License
 *
 * Copyright (c) 2022-2024 chief-of-state
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
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
