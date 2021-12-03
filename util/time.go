package util

import "time"

func GetCurrentTime() string {
	return time.Now().Format("YYYY-MM-DD HH:mm:ss")
}
