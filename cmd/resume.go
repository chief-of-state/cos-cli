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

// resumeCmd represents the resume command
var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "resume resumes a paused read side",
	Long:  `resume can be run for a given shard or across the whole cluster`,
	Run:   resumeRun,
}

func resumeRun(cmd *cobra.Command, args []string) {
	resumed := false
	// let us create the cos client
	cosClient, err := cos.NewClient(cmd.Context(), cosHost, cosPort)
	if err != nil {
		panic(err)
	}

	if shardNumber >= 0 {
		resp, err := cosClient.ResumeReadSideByShard(cmd.Context(), &cospb.ResumeReadSideByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the resumed variable
		resumed = resp.GetSuccessful()
	} else {
		resp, err := cosClient.ResumeReadSide(cmd.Context(), &cospb.ResumeReadSideRequest{
			ReadSideId: readSideID,
		})
		// handle the error
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		// set the resumed variable
		resumed = resp.GetSuccessful()
	}

	if resumed {
		log.Printf("read side=%s successfully resumed\n", readSideID)
		return
	}
	log.Printf("unable to resume read side=%s\n", readSideID)

}

func init() {
	readsideCmd.AddCommand(resumeCmd)
}
