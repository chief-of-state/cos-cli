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
	"time"

	"connectrpc.com/connect"

	cospb "github.com/chief-of-state/cos-cli/gen/chief_of_state/v1"
	cosconnect "github.com/chief-of-state/cos-cli/gen/chief_of_state/v1/chief_of_statev1connect"
)

const (
	// KeepAliveTime is the period after which a keepalive ping is sent on the
	// transport
	KeepAliveTime = 1200 * time.Second
)

type readSideManager struct {
	remote cosconnect.ReadSideManagerServiceClient
}

// NewReadSideManager creates a new instance of ReadSideManager
func NewReadSideManager(ctx context.Context, cosHost string, cosPort int) ReadSideManager {
	return readSideManager{
		remote: cosconnect.NewReadSideManagerServiceClient(
			httpClient(),
			URL(cosHost, cosPort),
			connect.WithGRPC(),
		),
	}
}

// GetLatestOffset retrieves the latest offset across all shards
func (c readSideManager) GetLatestOffset(ctx context.Context, in *connect.Request[cospb.GetLatestOffsetRequest]) (*connect.Response[cospb.GetLatestOffsetResponse], error) {
	return c.remote.GetLatestOffset(ctx, in)
}

// GetLatestOffsetByShard retrieves the latest offset given a shard
func (c readSideManager) GetLatestOffsetByShard(ctx context.Context, in *connect.Request[cospb.GetLatestOffsetByShardRequest]) (*connect.Response[cospb.GetLatestOffsetByShardResponse], error) {
	return c.remote.GetLatestOffsetByShard(ctx, in)
}

// RestartReadSide will clear the read side offset and start it over again
// from the first offset and this across all shards
func (c readSideManager) RestartReadSide(ctx context.Context, in *connect.Request[cospb.RestartReadSideRequest]) (*connect.Response[cospb.RestartReadSideResponse], error) {
	return c.remote.RestartReadSide(ctx, in)
}

// RestartReadSideByShard will clear the read side offset for the given shard and start it over again from the first offset
func (c readSideManager) RestartReadSideByShard(ctx context.Context, in *connect.Request[cospb.RestartReadSideByShardRequest]) (*connect.Response[cospb.RestartReadSideByShardResponse], error) {
	return c.remote.RestartReadSideByShard(ctx, in)
}

// PauseReadSide pauses a read side. This can be useful when running some data
// migration and this across all shards
func (c readSideManager) PauseReadSide(ctx context.Context, in *connect.Request[cospb.PauseReadSideRequest]) (*connect.Response[cospb.PauseReadSideResponse], error) {
	return c.remote.PauseReadSide(ctx, in)
}

// PauseReadSideByShard pauses a read side. This can be useful when running some data
// migration and this for a given shard
func (c readSideManager) PauseReadSideByShard(ctx context.Context, in *connect.Request[cospb.PauseReadSideByShardRequest]) (*connect.Response[cospb.PauseReadSideByShardResponse], error) {
	return c.remote.PauseReadSideByShard(ctx, in)
}

// ResumeReadSide resumes a paused read side and this across all shards
func (c readSideManager) ResumeReadSide(ctx context.Context, in *connect.Request[cospb.ResumeReadSideRequest]) (*connect.Response[cospb.ResumeReadSideResponse], error) {
	return c.remote.ResumeReadSide(ctx, in)
}

// ResumeReadSideByShard  resumes a paused read side for a given shard
func (c readSideManager) ResumeReadSideByShard(ctx context.Context, in *connect.Request[cospb.ResumeReadSideByShardRequest]) (*connect.Response[cospb.ResumeReadSideByShardResponse], error) {
	return c.remote.ResumeReadSideByShard(ctx, in)
}

// SkipOffset skips the current offset to read across all shards and continue with next. The operation will automatically restart the read side.
func (c readSideManager) SkipOffset(ctx context.Context, in *connect.Request[cospb.SkipOffsetRequest]) (*connect.Response[cospb.SkipOffsetResponse], error) {
	return c.remote.SkipOffset(ctx, in)
}

// SkipOffsetByShard skips the current offset to read for a given shard and continue with next. The operation will automatically restart the read side.
func (c readSideManager) SkipOffsetByShard(ctx context.Context, in *connect.Request[cospb.SkipOffsetByShardRequest]) (*connect.Response[cospb.SkipOffsetByShardResponse], error) {
	return c.remote.SkipOffsetByShard(ctx, in)
}
