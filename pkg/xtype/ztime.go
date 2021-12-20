package xtype

import (
	"database/sql/driver"
	"strconv"
	xtime "time"
)

const (
	ztimeFormart = "2006-01-02"
)

// ZTime be used to MySql timestamp converting.
type ZTime int64

// Scan scan time.
func (jt *ZTime) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = ZTime(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = ZTime(i)
	case []uint8:
		var i int64
		i, err = strconv.ParseInt(string(sc), 10, 64)
		*jt = ZTime(i)
	case int:
		*jt = ZTime(src.(int))
	case int64:
		*jt = ZTime(src.(int64))
	}
	return
}

// Value get time value.
func (jt ZTime) Value() (driver.Value, error) {
	return int64(jt), nil
}

// ZTime get time.
func (jt ZTime) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

func (jt ZTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(ztimeFormart)+2)
	b = append(b, '"')
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, ztimeFormart)
	b = append(b, '"')
	return b, nil
}

func (jt *ZTime) UnMarshalJSON(b []byte) (err error) {
	i, err := strconv.ParseInt(string(b), 10, 64)
	*jt = ZTime(i)
	return
}

func (jt ZTime) String() string {
	b := make([]byte, 0, len(ztimeFormart)+2)
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, ztimeFormart)
	return string(b)
}
