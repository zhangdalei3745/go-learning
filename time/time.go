package _time

import (
	"fmt"
	"time"
)

func currentTime() string {
	t := time.Now().Unix()
	return fmt.Sprintf(time.Unix(t, 0).String())
}
