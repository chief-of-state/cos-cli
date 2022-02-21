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

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "pause will pause a readside.",
	Long:  `pause can be run for a given shard or across the whole cluster`,
	Run:   pauseRun,
}

func pauseRun(cmd *cobra.Command, args []string) {
	paused := false
	// let us create the cos client
	cosClient, err := cos.NewClient(cmd.Context(), cosHost, cosPort)
	if err != nil {
		panic(err)
	}

	if shardNumber >= 0 {
		resp, err := cosClient.PauseReadSideByShard(cmd.Context(), &cospb.PauseReadSideByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the paused variable
		paused = resp.GetSuccessful()
	} else {
		resp, err := cosClient.PauseReadSide(cmd.Context(), &cospb.PauseReadSideRequest{
			ReadSideId: readSideID,
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the paused variable
		paused = resp.GetSuccessful()
	}

	if paused {
		log.Printf("read side=%s successfully pause\n", readSideID)
		return
	}
	log.Printf("unable to pause read side=%s\n", readSideID)
}

func init() {
	readsideCmd.AddCommand(pauseCmd)
}
