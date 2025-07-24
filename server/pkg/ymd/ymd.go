package ymd

import (
	"database/sql/driver"
	"errors"
	"time"
)

const Layout = "2006-01-02"

var (
	ErrInvalidDateFormat = errors.New("invalid date format, expected YYYY-MM-DD")
)

type YMD struct {
	Year  int
	Month time.Month
	Day   int
}

func New(year int, month time.Month, day int) YMD {
	return YMD{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func Parse(s string) (YMD, error) {
	t, err := time.Parse(Layout, s)
	if err != nil {
		return Zero(), ErrInvalidDateFormat
	}
	return FromTime(t), nil
}

func Format(ymd YMD) string {
	return time.Date(ymd.Year, ymd.Month, ymd.Day, 0, 0, 0, 0, time.UTC).Format(Layout)
}

func (ymd YMD) ToTime() time.Time {
	return time.Date(ymd.Year, ymd.Month, ymd.Day, 0, 0, 0, 0, time.UTC)
}

func FromTime(t time.Time) YMD {
	return YMD{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}
}

func (ymd YMD) String() string {
	return time.Date(ymd.Year, ymd.Month, ymd.Day, 0, 0, 0, 0, time.UTC).Format(Layout)
}

var zeroYMD = YMD{}

func Zero() YMD {
	return zeroYMD
}

func (ymd YMD) IsZero() bool {
	return ymd == zeroYMD
}

func (ymd YMD) IsValid() bool {
	if ymd.Year < 1 || ymd.Month < 1 || ymd.Month > 12 || ymd.Day < 1 {
		return false
	}
	daysInMonth := time.Date(ymd.Year, ymd.Month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	return ymd.Day <= daysInMonth
}

func (ymd YMD) AddDays(days int) YMD {
	t := ymd.ToTime().AddDate(0, 0, days)
	return New(t.Year(), t.Month(), t.Day())
}

func (ymd *YMD) Scan(value any) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("ymd: Scan source is not time.Time")
	}
	*ymd = FromTime(t)
	return nil
}

func (ymd YMD) Value() (driver.Value, error) {
	return ymd.ToTime(), nil
}
