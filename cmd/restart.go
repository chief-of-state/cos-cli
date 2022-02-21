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

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart basically restart a given read side",
	Long: `restart will clear the read side offset and start it over again from the first offset.

restart be run for a single shard or across the whole cluster.
`,
	Run: restartRun,
}

func restartRun(cmd *cobra.Command, args []string) {
	restarted := false
	// let us create the cos client
	cosClient, err := cos.NewClient(cmd.Context(), cosHost, cosPort)
	if err != nil {
		panic(err)
	}

	if shardNumber >= 0 {
		resp, err := cosClient.RestartReadSideByShard(cmd.Context(), &cospb.RestartReadSideByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the restarted variable
		restarted = resp.GetSuccessful()
	} else {
		resp, err := cosClient.RestartReadSide(cmd.Context(), &cospb.RestartReadSideRequest{
			ReadSideId: readSideID,
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the restarted variable
		restarted = resp.GetSuccessful()
	}

	if restarted {
		log.Printf("read side=%s successfully restarted\n", readSideID)
		return
	}
	log.Printf("unable to restart read side=%s\n", readSideID)

}

func init() {
	readsideCmd.AddCommand(restartCmd)
}
