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

	"connectrpc.com/connect"

	"cos-cli/cos"
	cospb "cos-cli/gen/chief_of_state/v1"

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
	// let us create the read-side manager
	manager := cos.NewReadSideManager(cmd.Context(), cosHost, cosPort)

	switch {
	case shardNumber >= 0:
		resp, err := manager.ResumeReadSideByShard(cmd.Context(), connect.NewRequest(&cospb.ResumeReadSideByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		}))
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		resumed = resp.Msg.GetSuccessful()
	default:
		resp, err := manager.ResumeReadSide(cmd.Context(), connect.NewRequest(&cospb.ResumeReadSideRequest{
			ReadSideId: readSideID,
		}))
		if err != nil {
			// TODO it is good to panic or not
			log.Panic(err)
		}
		resumed = resp.Msg.GetSuccessful()
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
