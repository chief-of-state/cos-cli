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
	"github.com/spf13/cobra"

	"cos-cli/cos"
	cospb "cos-cli/gen/chief_of_state/v1"
)

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "pause will pause a read-side.",
	Long:  `pause can be run for a given shard or across the whole cluster`,
	Run:   pauseRun,
}

func pauseRun(cmd *cobra.Command, args []string) {
	paused := false
	// let us create the read-side manager
	manager := cos.NewReadSideManager(cmd.Context(), cosHost, cosPort)

	switch {
	case shardNumber >= 0:
		resp, err := manager.PauseReadSideByShard(cmd.Context(), connect.NewRequest(&cospb.PauseReadSideByShardRequest{
			ReadSideId:         readSideID,
			ClusterShardNumber: uint64(shardNumber),
		}))
		if err != nil {
			log.Fatal(err)
		}
		paused = resp.Msg.GetSuccessful()
	default:
		resp, err := manager.PauseReadSide(cmd.Context(), connect.NewRequest(&cospb.PauseReadSideRequest{
			ReadSideId: readSideID,
		}))
		if err != nil {
			log.Fatal(err)
		}
		paused = resp.Msg.GetSuccessful()
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
