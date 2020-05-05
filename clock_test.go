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
		out := TimestampToNullTime(pb)
		So(out.Time, ShouldEqual, now)
	})
}
