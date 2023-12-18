package common

import (
	"fmt"
	"time"
)

func Exec() func() {
	start := time.Now()
	return func() {
		fmt.Println("Execution time : %v", time.Since(start))
	}
}
