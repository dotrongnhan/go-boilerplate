package utils

import (
	"database/sql"
	"github.com/spf13/cast"
	"time"
)

const (
	RequestDateTimeLayout  = "2006-01-02T15:04:05-07:00"
	RequestDateTimeLayout2 = "2006-01-02T15:04:05Z"
	RequestDateTimeLayout3 = "2006-01-02T15:04:05+07:00"
	RequestDateTimeLayout4 = "2006-01-02T15:04:05 07:00"
)

var RequestDateTimeLayouts = []string{
	RequestDateTimeLayout,
	RequestDateTimeLayout2,
	RequestDateTimeLayout3,
	RequestDateTimeLayout4,
}

func Mapping(value interface{}) (s string) {
	s = ""

	if value == nil {
		return
	}

	switch value.(type) {
	case time.Time: // Time struct
		return value.(time.Time).Format(RequestDateTimeLayout)

	case *time.Time: // Pointer of Time struct
		return value.(*time.Time).Format(RequestDateTimeLayout)

	case int64, *int64, float64, *float64: // Timestamp
		v := cast.ToInt64(value)
		if v == 0 {
			return
		}

		return time.Unix(v, 0).Format(RequestDateTimeLayout)

	case string: // Time struct after json marshal & unmarshal it came a time string but wrong format
		t, err := time.Parse(RequestDateTimeLayout2, cast.ToString(value))

		if err != nil {
			t, _ = time.Parse(RequestDateTimeLayout3, cast.ToString(value))
		}

		return t.Format(RequestDateTimeLayout)
	}

	return
}

func Reverse(value interface{}) int64 {
	stringDatetime := cast.ToString(value)
	if len(stringDatetime) < 1 {
		return 0
	}

	supportedLayout := []string{
		RequestDateTimeLayout,
		RequestDateTimeLayout2,
		RequestDateTimeLayout3,
		RequestDateTimeLayout4,
	}

	for _, l := range supportedLayout {
		t, err := time.Parse(l, stringDatetime)
		if err == nil {
			return t.Unix()
		}
	}

	return 0
}

func GetTimeFromUnixTimestamp(unix int64) *time.Time {
	if unix == 0 {
		return nil
	}

	datetime := time.Unix(unix, 0)

	return &datetime
}

func GetUnixTimestamp(t *time.Time) int64 {
	if t != nil && !t.IsZero() {
		return t.Unix()
	}

	return 0
}

func GetSqlNullTimeFromUnixTimestamp(unix int64) sql.NullTime {
	if unix == 0 {
		return sql.NullTime{}
	}

	datetime := time.Unix(unix, 0)

	return sql.NullTime{
		Time:  datetime,
		Valid: true,
	}
}

func GetTimeFromUnixPointer(unix *int64) *time.Time {
	if unix == nil {
		return nil
	}

	datetime := time.Unix(*unix, 0)

	return &datetime
}

func GetTimeUTCUnixMicro(unixMicro *int64) *time.Time {
	if unixMicro == nil {
		return nil
	}

	datetime := time.UnixMicro(*unixMicro).UTC()

	return &datetime
}

func GetPointerUnixTimestamp(t *time.Time) *int64 {
	if t != nil && !t.IsZero() {
		unix := t.Unix()
		return &unix
	}

	return nil
}
