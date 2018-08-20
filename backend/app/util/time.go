package util

import (
	"time"
)

func TimeDayDiff(new, old string) int {
	newTime, _ := time.Parse("2006-01-02 15:04:05", new)
	oldTime, _ := time.Parse("2006-01-02 15:04:05", old)

	hours := newTime.Sub(oldTime).Hours()
	if hours <= 0 {
		return -1
	}

	// sub hours less than 24
	if hours < 24 {
		newy, newm, newd := newTime.Date()
		oldy, oldm, oldd := oldTime.Date()
		isSameDay := (newy == oldy && newm == oldm && newd == oldd)

		if isSameDay {
			return 0
		} else {
			return 1
		}
	} else {
		if (hours/24)-float64(int(hours/24)) == 0 {
			return int(hours / 24)
		} else { // more than 24
			return int(hours/24) + 1
		}
	}
}

func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeNowYMD() string {
	return time.Now().Format("2006-01-02")
}
