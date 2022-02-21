/*
Copyright © 2022 Chief Of State

*/
package cmd

import (
	"log"
	"os"

	"github.com/chief-of-state/cos-cli/cos"
	cospb "github.com/chief-of-state/cos-cli/gen/chief_of_state/v1"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// offsetCmd represents the offset command
var offsetCmd = &cobra.Command{
	Use:   "offset",
	Short: "offset retrieves the current offset of a readside",
	Long:  `offset can retrieve the readside offset for a given shard or across the whole cluster`,
	Run:   offsetRun,
}

func offsetRun(cmd *cobra.Command, args []string) {
	var offsets []*cospb.ReadSideOffset
	// let us create the cos client
	cosClient, err := cos.NewClient(cmd.Context(), cosHost, cosPort)
	if err != nil {
		panic(err)
	}
	if shardNumber >= 0 {
		resp, err := cosClient.GetLatestOffsetByShard(cmd.Context(), &cospb.GetLatestOffsetByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// handle the response
		offsets = append(offsets, resp.GetOffsets())
	} else {
		resp, err := cosClient.GetLatestOffset(cmd.Context(), &cospb.GetLatestOffsetRequest{
			ReadSideId: readSideID,
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// handle the response
		offsets = append(offsets, resp.GetOffsets()...)
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
