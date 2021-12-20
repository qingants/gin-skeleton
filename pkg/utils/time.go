package utils

import "time"

func Unix2Time(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}


// 获取当天 零点的unix 时间
func UnixZero()int64{
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.Unix()
}
