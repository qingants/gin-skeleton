package xtype

import (
	"database/sql/driver"
	"strconv"
	xtime "time"
)

// YTime be used to MySql timestamp converting.
type YTime int64

// Scan scan time.
func (jt *YTime) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = YTime(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = YTime(i)
	case []uint8:
		var i int64
		i, err = strconv.ParseInt(string(sc), 10, 64)
		*jt = YTime(i)
	case int:
		*jt = YTime(src.(int))
	case int64:
		*jt = YTime(src.(int64))
	}
	return
}

// Value get time value.
func (jt YTime) Value() (driver.Value, error) {
	return int64(jt), nil
}

// YTime get time.
func (jt YTime) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

func (jt YTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (jt *YTime) UnMarshalJSON(b []byte) (err error) {
	i, err := strconv.ParseInt(string(b), 10, 64)
	*jt = YTime(i)
	return
}

func (jt YTime) String() string {
	b := make([]byte, 0, len(timeFormart)+2)
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, timeFormart)
	return string(b)
}
