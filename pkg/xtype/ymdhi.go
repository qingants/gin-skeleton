package xtype

import (
	"database/sql/driver"
	"strconv"
	xtime "time"
)

const (
	ymdhiFormart = "2006-01-02 15:04"
)

// Ymdhi be used to MySql timestamp converting.
type Ymdhi int64

// Scan scan time.
func (jt *Ymdhi) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = Ymdhi(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = Ymdhi(i)
	case []uint8:
		var i int64
		i, err = strconv.ParseInt(string(sc), 10, 64)
		*jt = Ymdhi(i)
	case int:
		*jt = Ymdhi(src.(int))
	case int64:
		*jt = Ymdhi(src.(int64))
	}
	return
}

// Value get time value.
func (jt Ymdhi) Value() (driver.Value, error) {
	return int64(jt), nil
}

// Ymdhi get time.
func (jt Ymdhi) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

func (jt Ymdhi) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(ymdhiFormart)+2)
	b = append(b, '"')
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, ymdhiFormart)
	b = append(b, '"')
	return b, nil
}

func (jt *Ymdhi) UnMarshalJSON(b []byte) (err error) {
	i, err := strconv.ParseInt(string(b), 10, 64)
	*jt = Ymdhi(i)
	return
}

func (jt Ymdhi) String() string {
	b := make([]byte, 0, len(ymdhiFormart)+2)
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, ymdhiFormart)
	return string(b)
}
