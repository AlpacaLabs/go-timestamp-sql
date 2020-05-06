package clocksql

import (
	"testing"
	"time"

	clock "github.com/AlpacaLabs/go-timestamp"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_TimestampConversion(t *testing.T) {
	Convey("Given some non-zero time", t, func(c C) {
		now := time.Now()
		pb := clock.TimeToTimestamp(now)
		nt := TimestampToNullTime(pb)
		So(nt.Time, ShouldEqual, now)
		pb2 := TimestampFromNullTime(nt)
		So(pb.Seconds, ShouldEqual, pb2.Seconds)
		So(pb.Nanos, ShouldEqual, pb2.Nanos)
	})
}
