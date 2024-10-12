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
	"log"
	"os"

	"connectrpc.com/connect"

	"github.com/chief-of-state/cos-cli/cos"
	cospb "github.com/chief-of-state/cos-cli/gen/chief_of_state/v1"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// offsetCmd represents the offset command
var offsetCmd = &cobra.Command{
	Use:   "offset",
	Short: "offset retrieves the current offset of a read-side",
	Long:  `offset can retrieve the read-side offset for a given shard or across the whole cluster`,
	Run:   offsetRun,
}

func offsetRun(cmd *cobra.Command, args []string) {
	var offsets []*cospb.ReadSideOffset
	// let us create the read-side manager
	manager := cos.NewReadSideManager(cmd.Context(), cosHost, cosPort)

	switch {
	case shardNumber >= 0:
		resp, err := manager.GetLatestOffsetByShard(cmd.Context(), connect.NewRequest(&cospb.GetLatestOffsetByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		}))
		if err != nil {
			log.Fatal(err)
		}
		offsets = append(offsets, resp.Msg.GetOffsets())
	default:
		resp, err := manager.GetLatestOffset(cmd.Context(), connect.NewRequest(&cospb.GetLatestOffsetRequest{
			ReadSideId: readSideID,
		}))
		if err != nil {
			log.Fatal(err)
		}
		offsets = append(offsets, resp.Msg.GetOffsets()...)
	}

	// let us display the data
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBright)
	t.AppendHeader(table.Row{"Shard Number", "Offset"})
	rows := make([]table.Row, len(offsets))
	for index, row := range offsets {
		rows[index] = table.Row{
			row.GetClusterShardNumber(),
			row.GetOffset(),
		}
	}
	t.AppendRows(rows)
	t.AppendSeparator()
	t.Render()
}

func init() {
	readsideCmd.AddCommand(offsetCmd)
}
