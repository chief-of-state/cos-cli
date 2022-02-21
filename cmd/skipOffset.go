/*
Copyright © 2022 Chief Of State

*/
package cmd

import (
	"log"

	"github.com/chief-of-state/cos-cli/cos"
	cospb "github.com/chief-of-state/cos-cli/gen/chief_of_state/v1"
	"github.com/spf13/cobra"
)

// skipOffsetCmd represents the skipOffset command
var skipOffsetCmd = &cobra.Command{
	Use:   "skipOffset",
	Short: "skipOffset skips the current offset of a read side and continue with next",
	Long:  `skipOffset will automatically restart the read side. This can be for a given shard or across the whole cluster`,
	Run:   skipRun,
}

func skipRun(cmd *cobra.Command, args []string) {
	skipped := false
	// let us create the cos client
	cosClient, err := cos.NewClient(cmd.Context(), cosHost, cosPort)
	if err != nil {
		panic(err)
	}

	if shardNumber >= 0 {
		resp, err := cosClient.SkipOffsetByShard(cmd.Context(), &cospb.SkipOffsetByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the paused variable
		skipped = resp.GetSuccessful()
	} else {
		resp, err := cosClient.SkipOffset(cmd.Context(), &cospb.SkipOffsetRequest{
			ReadSideId: readSideID,
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the paused variable
		skipped = resp.GetSuccessful()
	}

	if skipped {
		log.Printf("read side=%s offsets successfully skipped\n", readSideID)
		return
	}
	log.Printf("unable to skip read side=%s offsets\n", readSideID)

}

func init() {
	readsideCmd.AddCommand(skipOffsetCmd)
}
