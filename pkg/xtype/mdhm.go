package xtype

import (
	"database/sql/driver"
	"strconv"
	xtime "time"
)

const (
	mdhmFormart = "01-02 15:04"
)

// Mdhm be used to MySql timestamp converting.
type Mdhm int64

// Scan scan time.
func (jt *Mdhm) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = Mdhm(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = Mdhm(i)
	case []uint8:
		var i int64
		i, err = strconv.ParseInt(string(sc), 10, 64)
		*jt = Mdhm(i)
	case int:
		*jt = Mdhm(src.(int))
	case int64:
		*jt = Mdhm(src.(int64))
	}
	return
}

// Value get time value.
func (jt Mdhm) Value() (driver.Value, error) {
	return int64(jt), nil
}

// Mdhm get time.
func (jt Mdhm) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

func (jt Mdhm) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(mdhmFormart)+2)
	b = append(b, '"')
	b = xtime.Unix(int64(jt), 0).AppendFormat(b, mdhmFormart)
	b = append(b, '"')
	return b, nil
}

func (jt *Mdhm) UnMarshalJSON(b []byte) (err error) {
	i, err := strconv.ParseInt(string(b), 10, 64)
	*jt = Mdhm(i)
	return
}
