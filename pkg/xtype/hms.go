package xtype

import (
	"database/sql/driver"
	"strconv"
	xtime "time"
)

const (
	hmsFormart = "15:04"
)

// Hms be used to MySql timestamp converting.
type Hms int64

// Scan scan time.
func (jt *Hms) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = Hms(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = Hms(i)
	case []uint8:
		var i int64
		i, err = strconv.ParseInt(string(sc), 10, 64)
		*jt = Hms(i)
	case int:
		*jt = Hms(src.(int))
	case int64:
		*jt = Hms(src.(int64))
	}
	return
}

// Value get time value.
func (jt Hms) Value() (driver.Value, error) {
	return int64(jt), nil
}

// Hms get time.
func (jt Hms) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

func (jt Hms) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(hmsFormart)+2)
	b = append(b, '"')
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, hmsFormart)
	b = append(b, '"')
	return b, nil
}

func (jt *Hms) UnMarshalJSON(b []byte) (err error) {
	i, err := strconv.ParseInt(string(b), 10, 64)
	*jt = Hms(i)
	return
}
