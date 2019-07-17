package mtime

import (
	"time"
)

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
	secondsPerDay    = 24 * secondsPerHour
	secondsPerWeek   = 7 * secondsPerDay
)


func GetSinceDay() int64{
	
     year := time.Now().Year()
     month := time.Now().Month()
     day := time.Now().Day()
     hour := time.Now().Hour()
     minute := time.Now().Minute()
     second := time.Now().Second()
     nanosecond := time.Now().Nanosecond()

    loc, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Date(year,month,day,hour,minute,second,nanosecond,loc)

	return t.Unix()/int64(secondsPerDay)
}

func GetSinceWeek() int64{
	
    year := time.Now().Year()
    month := time.Now().Month()
    day := time.Now().Day()
    hour := time.Now().Hour()
    minute := time.Now().Minute()
    second := time.Now().Second()
    nanosecond := time.Now().Nanosecond()

   loc, _ := time.LoadLocation("Asia/Shanghai")
   t := time.Date(year,month,day,hour,minute,second,nanosecond,loc)

   return t.Unix()/int64(secondsPerWeek)
}