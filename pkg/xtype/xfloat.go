package xtype

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Time be used to MySql timestamp converting.
type XFloat float64

// Scan scan time.
func (jt *XFloat) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case float64:
		*jt = XFloat(src.(float64))
	case float32:
		*jt = XFloat(src.(float32))
	case string:
		var i float64
		i, err = strconv.ParseFloat(sc, 64)
		*jt = XFloat(i)
	case []uint8:
		var i float64
		i, err = strconv.ParseFloat(string(sc), 64)
		*jt = XFloat(i)
	case int64:
		*jt = XFloat(src.(int64))
	case int32:
		*jt = XFloat(src.(int32))
	case int:
		*jt = XFloat(src.(int))
	}
	return
}

// Value get time value.
func (jt XFloat) Value() (driver.Value, error) {
	return float64(jt), nil
}

func (jt XFloat) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.2f", float64(jt))
	b := []byte(s)
	return b, nil
}

func (jt *XFloat) UnMarshalJSON(b []byte) (err error) {
	i, err := strconv.ParseFloat(string(b), 64)
	*jt = XFloat(i)
	return
}

func (jt XFloat) String() string {
	s := fmt.Sprintf("%.2f", float64(jt))
	b := []byte(s)
	return string(b)
}
