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

package cos

import (
	"context"

	"connectrpc.com/connect"

	cospb "cos-cli/gen/chief_of_state/v1"
)

// ReadSideManager will be implemented by the CoS readSideManager
// to send requests to CoS read side manager service
type ReadSideManager interface {
	// GetLatestOffset retrieves the latest offset across all shards
	GetLatestOffset(ctx context.Context, in *connect.Request[cospb.GetLatestOffsetRequest]) (*connect.Response[cospb.GetLatestOffsetResponse], error)
	// GetLatestOffsetByShard retrieves the latest offset given a shard
	GetLatestOffsetByShard(ctx context.Context, in *connect.Request[cospb.GetLatestOffsetByShardRequest]) (*connect.Response[cospb.GetLatestOffsetByShardResponse], error)
	// RestartReadSide will clear the read side offset and start it over again
	// from the first offset and this across all shards
	RestartReadSide(ctx context.Context, in *connect.Request[cospb.RestartReadSideRequest]) (*connect.Response[cospb.RestartReadSideResponse], error)
	// RestartReadSideByShard will clear the read side offset for the given shard and start it over again from the first offset
	RestartReadSideByShard(ctx context.Context, in *connect.Request[cospb.RestartReadSideByShardRequest]) (*connect.Response[cospb.RestartReadSideByShardResponse], error)
	// PauseReadSide pauses a read side. This can be useful when running some data
	// migration and this across all shards
	PauseReadSide(ctx context.Context, in *connect.Request[cospb.PauseReadSideRequest]) (*connect.Response[cospb.PauseReadSideResponse], error)
	// PauseReadSideByShard pauses a read side. This can be useful when running some data
	// migration and this for a given shard
	PauseReadSideByShard(ctx context.Context, in *connect.Request[cospb.PauseReadSideByShardRequest]) (*connect.Response[cospb.PauseReadSideByShardResponse], error)
	// ResumeReadSide resumes a paused read side and this across all shards
	ResumeReadSide(ctx context.Context, in *connect.Request[cospb.ResumeReadSideRequest]) (*connect.Response[cospb.ResumeReadSideResponse], error)
	// ResumeReadSideByShard  resumes a paused read side for a given shard
	ResumeReadSideByShard(ctx context.Context, in *connect.Request[cospb.ResumeReadSideByShardRequest]) (*connect.Response[cospb.ResumeReadSideByShardResponse], error)
	// SkipOffset skips the current offset to read across all shards and continue with next. The operation will automatically restart the read side.
	SkipOffset(ctx context.Context, in *connect.Request[cospb.SkipOffsetRequest]) (*connect.Response[cospb.SkipOffsetResponse], error)
	// SkipOffsetByShard skips the current offset to read for a given shard and continue with next. The operation will automatically restart the read side.
	SkipOffsetByShard(ctx context.Context, in *connect.Request[cospb.SkipOffsetByShardRequest]) (*connect.Response[cospb.SkipOffsetByShardResponse], error)
}
