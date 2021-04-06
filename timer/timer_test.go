package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestTimerLocalSerial(t *testing.T) {
	timer := NewTimer()
	timer.Add(func() {
		time.Sleep(5*time.Second)
		fmt.Println("g1,", time.Now())
	}, 20 * time.Second, WithLocalSerial())

	timer.Add(func() {
		time.Sleep(5*time.Second)
		fmt.Println("g2,", time.Now())
	}, 21 * time.Second, WithLocalSerial())

	timer.Add(func() {
		time.Sleep(5*time.Second)
		fmt.Println("g3,", time.Now())
	}, 22 * time.Second, WithLocalSerial())

	timer.Add(func() {
		time.Sleep(5*time.Second)
		fmt.Println("g4,", time.Now())
	}, 23 * time.Second, WithLocalSerial())
	timer.Start()
	time.Sleep(10 * time.Minute)
}
