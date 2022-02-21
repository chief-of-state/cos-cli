package cos

import (
	"context"

	cospb "github.com/chief-of-state/cos-cli/gen/chief_of_state/v1"
)

// ReadSideManager will be implemented by the CoS client
// to send requests to CoS read side manager service
type ReadSideManager interface {
	// GetLatestOffset retrieves the latest offset across all shards
	GetLatestOffset(ctx context.Context, in *cospb.GetLatestOffsetRequest) (*cospb.GetLatestOffsetResponse, error)
	// GetLatestOffsetByShard retrieves the latest offset given a shard
	GetLatestOffsetByShard(ctx context.Context, in *cospb.GetLatestOffsetByShardRequest) (*cospb.GetLatestOffsetByShardResponse, error)
	// RestartReadSide will clear the read side offset and start it over again
	// from the first offset and this across all shards
	RestartReadSide(ctx context.Context, in *cospb.RestartReadSideRequest) (*cospb.RestartReadSideResponse, error)
	// RestartReadSideByShard will clear the read side offset for the given shard and start it over again from the first offset
	RestartReadSideByShard(ctx context.Context, in *cospb.RestartReadSideByShardRequest) (*cospb.RestartReadSideByShardResponse, error)
	// PauseReadSide pauses a read side. This can be useful when running some data
	// migration and this across all shards
	PauseReadSide(ctx context.Context, in *cospb.PauseReadSideRequest) (*cospb.PauseReadSideResponse, error)
	// PauseReadSideByShard pauses a read side. This can be useful when running some data
	// migration and this for a given shard
	PauseReadSideByShard(ctx context.Context, in *cospb.PauseReadSideByShardRequest) (*cospb.PauseReadSideByShardResponse, error)
	// ResumeReadSide resumes a paused read side and this across all shards
	ResumeReadSide(ctx context.Context, in *cospb.ResumeReadSideRequest) (*cospb.ResumeReadSideResponse, error)
	// ResumeReadSideByShard  resumes a paused read side for a given shard
	ResumeReadSideByShard(ctx context.Context, in *cospb.ResumeReadSideByShardRequest) (*cospb.ResumeReadSideByShardResponse, error)
	// SkipOffset skips the current offset to read across all shards and continue with next. The operation will automatically restart the read side.
	SkipOffset(ctx context.Context, in *cospb.SkipOffsetRequest) (*cospb.SkipOffsetResponse, error)
	// SkipOffsetByShard skips the current offset to read for a given shard and continue with next. The operation will automatically restart the read side.
	SkipOffsetByShard(ctx context.Context, in *cospb.SkipOffsetByShardRequest) (*cospb.SkipOffsetByShardResponse, error)
}
