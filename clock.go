package clocksql

import (
	"time"

	clock "github.com/AlpacaLabs/go-timestamp"
	"github.com/guregu/null"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func TimestampToNullTime(in *timestamp.Timestamp) null.Time {
	t := clock.TimestampToTime(in)

	var nt null.Time
	if t.IsZero() {
		nt = null.NewTime(time.Time{}, false)
	} else {
		nt = null.TimeFrom(t)
	}
	return nt
}

func TimestampFromNullTime(in null.Time) *timestamp.Timestamp {
	return clock.TimeToTimestamp(in.ValueOrZero())
}
