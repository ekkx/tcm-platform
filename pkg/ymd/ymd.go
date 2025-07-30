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

var jst = mustLoadJST()

func mustLoadJST() *time.Location {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load Asia/Tokyo time zone")
	}
	return loc
}

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
	return time.Date(ymd.Year, ymd.Month, ymd.Day, 0, 0, 0, 0, jst).Format(Layout)
}

func (ymd YMD) ToTime() time.Time {
	return time.Date(ymd.Year, ymd.Month, ymd.Day, 0, 0, 0, 0, jst)
}

func FromTime(t time.Time) YMD {
	t = t.In(jst)
	return YMD{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}
}

func (ymd YMD) String() string {
	return time.Date(ymd.Year, ymd.Month, ymd.Day, 0, 0, 0, 0, jst).Format(Layout)
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
	daysInMonth := time.Date(ymd.Year, ymd.Month+1, 0, 0, 0, 0, 0, jst).Day()
	return ymd.Day <= daysInMonth
}

func (ymd YMD) Equal(other YMD) bool {
	return ymd.Year == other.Year && ymd.Month == other.Month && ymd.Day == other.Day
}

func (ymd YMD) Before(other YMD) bool {
	if ymd.Year < other.Year {
		return true
	}
	if ymd.Year > other.Year {
		return false
	}
	if ymd.Month < other.Month {
		return true
	}
	if ymd.Month > other.Month {
		return false
	}
	return ymd.Day < other.Day
}

func (ymd YMD) After(other YMD) bool {
	if ymd.Year > other.Year {
		return true
	}
	if ymd.Year < other.Year {
		return false
	}
	if ymd.Month > other.Month {
		return true
	}
	if ymd.Month < other.Month {
		return false
	}
	return ymd.Day > other.Day
}

func Today() YMD {
	return FromTime(time.Now().In(jst))
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
