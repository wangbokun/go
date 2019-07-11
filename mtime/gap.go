package mtime

import(
	"time"
	"fmt"
)

// func Gap(start_time, end_time string) string {
	
// 	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
// 	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
// 	if err == nil && t1.Before(t2) {
// 		diff := t2.Unix() - t1.Unix()

// 		if diff/60 < 60 {
// 			return fmt.Sprintf("%d分钟", diff/60)
// 		}

// 		if diff/60 > 60 {
// 			return fmt.Sprintf("%d小时", diff/60/60)
// 		}

// 		if diff/60 > 1440 {
// 			return fmt.Sprintf("%d天", diff/60/60/24)
// 		}

// 	} else {
// 		return "Null"
// 	}
// 	return "Null"
// }


func Gap(start_time, end_time string) string {
	// var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() 

		if diff/60 < 60 {
			return fmt.Sprintf("%d分钟", diff/60)
		}

		if diff/60 > 60 {
			return fmt.Sprintf("%d小时", diff/60/60)
		}

		if diff/60 > 1440 {
			return fmt.Sprintf("%d天", diff/60/60/24)
		}

	} else {
		return "0秒"
	}
	return "0秒"
}