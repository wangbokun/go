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

func GetSinceMonth() int64{
    var index int
    year := time.Now().Year()
    var months = [...]string{
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    }

    for k,i := range  months{
        if time.Now().Month().String() == i{
            index = k+1
            continue
        }
    }
    return int64((year-1970)*12+index)
}

func GetSinceYear() int64{
   return int64(time.Now().Year() - 1970)
}