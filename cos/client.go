package cos

import (
	"context"
	"fmt"
	"time"

	cospb "github.com/chief-of-state/cos-cli/gen/chief_of_state/v1"
	"google.golang.org/grpc/keepalive"
)

const (
	// KeepAliveTime is the period after which a keepalive ping is sent on the
	// transport
	KeepAliveTime = 1200 * time.Second
)

type client struct {
	remote cospb.ReadSideManagerServiceClient
}

// NewClient creates a new instance of ReadSideManager
func NewClient(ctx context.Context, cosHost string, cosPort int) (ReadSideManager, error) {
	clientBuilder := NewBuilder().
		WithInsecure().
		WithKeepAliveParams(keepalive.ClientParameters{
			Time:                KeepAliveTime,
			PermitWithoutStream: true,
		})

	// get the grpc client connection to CoS
	conn, err := clientBuilder.GetConn(ctx, fmt.Sprintf("%v:%v", cosHost, cosPort))
	if err != nil {
		return nil, err
	}
	return client{
		remote: cospb.NewReadSideManagerServiceClient(conn),
	}, nil
}

// GetLatestOffset retrieves the latest offset across all shards
func (c client) GetLatestOffset(ctx context.Context, in *cospb.GetLatestOffsetRequest) (*cospb.GetLatestOffsetResponse, error) {
	return c.remote.GetLatestOffset(ctx, in)
}

// GetLatestOffsetByShard retrieves the latest offset given a shard
func (c client) GetLatestOffsetByShard(ctx context.Context, in *cospb.GetLatestOffsetByShardRequest) (*cospb.GetLatestOffsetByShardResponse, error) {
	return c.remote.GetLatestOffsetByShard(ctx, in)
}

// RestartReadSide will clear the read side offset and start it over again
// from the first offset and this across all shards
func (c client) RestartReadSide(ctx context.Context, in *cospb.RestartReadSideRequest) (*cospb.RestartReadSideResponse, error) {
	return c.remote.RestartReadSide(ctx, in)
}

// RestartReadSideByShard will clear the read side offset for the given shard and start it over again from the first offset
func (c client) RestartReadSideByShard(ctx context.Context, in *cospb.RestartReadSideByShardRequest) (*cospb.RestartReadSideByShardResponse, error) {
	return c.remote.RestartReadSideByShard(ctx, in)
}

// PauseReadSide pauses a read side. This can be useful when running some data
// migration and this across all shards
func (c client) PauseReadSide(ctx context.Context, in *cospb.PauseReadSideRequest) (*cospb.PauseReadSideResponse, error) {
	return c.remote.PauseReadSide(ctx, in)
}

// PauseReadSideByShard pauses a read side. This can be useful when running some data
// migration and this for a given shard
func (c client) PauseReadSideByShard(ctx context.Context, in *cospb.PauseReadSideByShardRequest) (*cospb.PauseReadSideByShardResponse, error) {
	return c.remote.PauseReadSideByShard(ctx, in)
}

// ResumeReadSide resumes a paused read side and this across all shards
func (c client) ResumeReadSide(ctx context.Context, in *cospb.ResumeReadSideRequest) (*cospb.ResumeReadSideResponse, error) {
	return c.remote.ResumeReadSide(ctx, in)
}

// ResumeReadSideByShard  resumes a paused read side for a given shard
func (c client) ResumeReadSideByShard(ctx context.Context, in *cospb.ResumeReadSideByShardRequest) (*cospb.ResumeReadSideByShardResponse, error) {
	return c.remote.ResumeReadSideByShard(ctx, in)
}

// SkipOffset skips the current offset to read across all shards and continue with next. The operation will automatically restart the read side.
func (c client) SkipOffset(ctx context.Context, in *cospb.SkipOffsetRequest) (*cospb.SkipOffsetResponse, error) {
	return c.remote.SkipOffset(ctx, in)
}

// SkipOffsetByShard skips the current offset to read for a given shard and continue with next. The operation will automatically restart the read side.
func (c client) SkipOffsetByShard(ctx context.Context, in *cospb.SkipOffsetByShardRequest) (*cospb.SkipOffsetByShardResponse, error) {
	return c.remote.SkipOffsetByShard(ctx, in)
}
