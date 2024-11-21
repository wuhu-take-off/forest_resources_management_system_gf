package common

import "time"

const dateStr = "Mon Jan 02 2006 15:04:05 GMT-0700 (MST)"

func DateConversion(date any) any {
	switch date.(type) {
	case string:
		t, _ := time.Parse(date.(string), dateStr)
		return t.Unix()
	case int64:
		// 假设时间戳
		timestamp := int64(1730899200)

		// 将时间戳转换为 time.Time
		t := time.Unix(timestamp, 0)

		// 设置时区为中国标准时间 (CST)
		loc, _ := time.LoadLocation("Asia/Shanghai")
		// 转换为中国标准时间
		t = t.In(loc)

		// 格式化为自定义的日期字符串格式
		formattedDate := t.Format(dateStr)
		return formattedDate
	default:
		return nil
	}
}
