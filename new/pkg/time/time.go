package time_plus

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Local")
}

// ToProcessTime 按照产品需求格式化时间
func ToProcessTime(timeObj time.Time) string {
	timeNowObj := time.Now()
	now := timeNowObj.Unix()
	timestamp := timeObj.Unix()

	var per, formatedStr string
	var diff int64
	if now >= timestamp {
		per = "前"
		diff = now - timestamp
	} else {
		per = "后"
		diff = timestamp - now
	}

	year := timeObj.Year()
	yearNow := timeNowObj.Year()
	today := Today()
	tomorrow := today + 86400
	if diff < 120 {
		if now >= timestamp {
			formatedStr = "刚刚"
		} else {
			formatedStr = "即将"
		}
	} else if diff < 3600 {
		mins := int(math.Floor(float64(diff / 60)))
		formatedStr = strconv.Itoa(mins) + "分钟" + per
	} else if diff < 18000 {
		hours := int(math.Floor(float64(diff / 3600)))
		formatedStr = strconv.Itoa(hours) + "小时" + per
	} else if timestamp >= today && timestamp < tomorrow {
		formatedStr = timeObj.Format("15:04")
		if timestamp >= tomorrow {
			formatedStr = "明日 " + formatedStr
		}
	} else if timestamp < today && timestamp >= today-86400 {
		tmp := timeObj.Format("15:04")
		if tmp == "00:00" {
			tmp = ""
		}
		formatedStr = "昨天" + tmp
	} else if timestamp >= tomorrow && timestamp < tomorrow+86400 {
		formatedStr = "明天" + timeObj.Format("15:04")
	} else if math.Abs(float64(yearNow-year)) < 1 {
		formatedStr = timeObj.Format("01-02")
	} else {
		formatedStr = timeObj.Format("2006-01-02")
	}

	return formatedStr
}

// Today 获取今天零点时间戳
func Today() int64 {
	timeStr := time.Now().Format("2006-01-02")
	ti, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	timeNumber := ti.Unix()

	return timeNumber
}

func GetMonthTime(timeObj time.Time) string {
	timeNowObj := time.Now()

	if timeObj.Year() == timeNowObj.Year() {
		return timeObj.Format("01月")
	} else {
		return timeObj.Format("2006年01月")
	}
}

//GetDurationToClock 获取t到指定整点需要经过的时间
func GetDurationToClock(t time.Time, clock int) time.Duration {
	y, m, d := t.Date()
	clockTime := time.Date(y, m, d, clock, 0, 0, 0, t.Location())

	if t.Before(clockTime) {
		return clockTime.Sub(t)
	} else {
		return clockTime.Add(24 * time.Hour).Sub(t)
	}
}

//GetNextClockTime 获取时间t的下一个clock整点
func GetNextClockTime(t time.Time, clock int) time.Time {
	y, m, d := t.Date()
	clockTime := time.Date(y, m, d, clock, 0, 0, 0, t.Location())

	if t.Before(clockTime) {
		return clockTime
	} else {
		return clockTime.Add(24 * time.Hour)
	}
}

//GetPrevClockTime 获取时间t的上一个clock整点
func GetPrevClockTime(t time.Time, clock int) time.Time {
	y, m, d := t.Date()
	clockTime := time.Date(y, m, d, clock, 0, 0, 0, t.Location())

	if t.Before(clockTime) {
		return clockTime.Add(-24 * time.Hour)
	} else {
		return clockTime
	}
}

// 解析武汉给的生日格式 20111121
func ParseDouyuBirthday(birthday string) (time.Time, error) {
	return time.ParseInLocation("20060102", birthday, loc)
}

func GetAge(birthTime time.Time) int {
	birthYear := birthTime.Year()
	birthMonth := birthTime.Month()
	birthDay := birthTime.Day()
	now := time.Now()
	nowYear := now.Year()
	nowMonth := now.Month()
	nowDay := now.Day()
	age := nowYear - birthYear
	if nowMonth < birthMonth {
		return age - 1
	}
	if nowMonth == birthMonth {
		if nowDay < birthDay {
			return age - 1
		}
	}
	return age
}

func GetDefaultAge() int {
	defaultBirth, _ := time.ParseInLocation("20060102", "19900101", loc)
	return GetAge(defaultBirth)
}

// GetConstellation 通过用户生日获取星座
func GetConstellation(t time.Time) (star string) {
	month := t.Month()
	day := t.Day()
	switch {
	case month <= 0, month >= 13, day <= 0, day >= 32:
		star = "摩羯座"
	case (month == 1 && day >= 20), (month == 2 && day <= 18):
		star = "水瓶座"
	case (month == 2 && day >= 19), (month == 3 && day <= 20):
		star = "双鱼座"
	case (month == 3 && day >= 21), (month == 4 && day <= 19):
		star = "白羊座"
	case (month == 4 && day >= 20), (month == 5 && day <= 20):
		star = "金牛座"
	case (month == 5 && day >= 21), (month == 6 && day <= 21):
		star = "双子座"
	case (month == 6 && day >= 22), (month == 7 && day <= 22):
		star = "巨蟹座"
	case (month == 7 && day >= 23), (month == 8 && day <= 22):
		star = "狮子座"
	case (month == 8 && day >= 23), (month == 9 && day <= 22):
		star = "处女座"
	case (month == 9 && day >= 23), (month == 10 && day <= 22):
		star = "天秤座"
	case (month == 10 && day >= 23), (month == 11 && day <= 21):
		star = "天蝎座"
	case (month == 11 && day >= 22), (month == 12 && day <= 21):
		star = "射手座"
	case (month == 12 && day >= 22), (month == 1 && day <= 19):
		star = "摩羯座"
	}
	return
}

//GetZodiac 生肖
func GetZodiac(t time.Time) (zodiac string) {
	year := t.Year()
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

//一周剩下的秒数
func TheRestOfWeek() time.Duration {
	t := time.Now()
	day := 7 - int(t.Weekday()) //一周剩下的天数,包括今天
	t2, _ := time.ParseInLocation("2006-01-02", t.Format("2006-01-02"), time.Local)
	lastDayUnix := t2.AddDate(0, 0, day).Unix()
	res := time.Duration(lastDayUnix - t.Unix() - 1)

	return res
}

//一天剩下的秒数
func TheRestOfDay() time.Duration {
	timeStr := time.Now().Format("2006-01-02")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	res := time.Duration(t2.Unix() - time.Now().Unix())

	return res
}

// 获取n天前的起始时间
func BeginOfPrevNDay(n int) time.Time {
	now := time.Now()
	h := time.Duration(-now.Hour()) * time.Hour
	beginDay := now.Truncate(time.Hour).Add(h)
	return beginDay.Add(time.Duration(0-n) * 24 * time.Hour)
}

// 获取某天末尾时间
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// TheRestOfMonth 获取本月余下的秒数
func TheRestOfMonth() time.Duration {
	now := time.Now()
	lastDayOfM := LastDayOfMonth().Format("2006-01-02")
	lastDayOfM1, _ := time.ParseInLocation("2006-01-02 15:04:05", lastDayOfM+" 23:59:59", time.Local)

	return time.Duration(lastDayOfM1.Unix() - now.Unix())
}

// LastDayOfMonth 获取本月最后一天的时间对象
func LastDayOfMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.Local)
}

// NextYearOfToday 获取明年的今天
func NextYearOfToday() time.Time {
	now := time.Now()
	return time.Date(now.Year()+1, now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local)
}

// parseFloat64ToStringTime 将float64 的时间转换为分秒
func ParseFloat64ToStringTime(duration float64) string {
	minute := math.Floor(duration / 60)
	second := math.Ceil(duration - minute*60)

	var strTime string
	if minute < 10 {
		strTime = "0" + fmt.Sprint(minute) + ":"
	} else {
		strTime = fmt.Sprint(minute) + ":"
	}

	if second < 10 {
		strTime += "0" + fmt.Sprint(second)
	} else {
		strTime += fmt.Sprint(second)
	}

	return strTime
}

// WeekDayToHan returns the chinese name of the day
func WeekDayToHan(wd int) string {
	m := [7]string{0: "日", 1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六"}
	if wd >= int(time.Sunday) && wd <= int(time.Saturday) {
		return m[wd]
	}

	return m[0]
}

// FirstDayOfMonth 本月第一天
func FirstDayOfMonth() time.Time {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()

	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Local)
}

// CheckTime 判断某个时间是否是6个月前
func CheckTime(scoreTime time.Time) bool {
	var b bool
	if (time.Now().Unix() - scoreTime.Unix()) > 15552000 { // 6*30*86400
		b = true
	}
	return b
}
