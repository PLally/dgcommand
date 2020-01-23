package snowflake

import (
	"strconv"
	"time"
)

type Snowflake struct {
	ID                uint64
	Increment         uint64
	InternalProcessID uint64
	InternalWorkerID  uint64
	TimestampDiscord  uint64
	TimestampUnix     uint64
	Time              time.Time
}

func NewSnowflake(id string) (Snowflake, error) {
	snowflake, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return Snowflake{}, err
	}

	increment := snowflake & (2 ^ 12 - 1)
	internalProcessID := (snowflake & 0x1F000) >> 12
	internalWorkerID := (snowflake & 0x3E0000) >> 17
	timestampDiscord := snowflake >> 22
	timestampUnix := 1420070400000 + timestampDiscord
	creationTime := time.Unix(int64(timestampUnix/1000), 0)

	return Snowflake{
		snowflake,
		increment,
		internalProcessID,
		internalWorkerID,
		timestampDiscord,
		timestampUnix,
		creationTime,
	}, nil
}
