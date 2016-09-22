package throttle_test

import (
	"fmt"
	"time"

	"github.com/boz/go-throttle"
)

const (
	period = 200 * time.Millisecond
)

func ExampleNewThrottle_untrailing() {
	throttle := throttle.NewThrottle(period, false)

	go func() {
		for throttle.Next() {
			fmt.Println("hello not trailing")
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			throttle.Trigger()
			time.Sleep(period / 6)
		}
	}()

	time.Sleep(2 * period)
	throttle.Stop()

	// Output: hello not trailing
}

func ExampleNewThrottle_trailing() {

	throttle := throttle.NewThrottle(period, true)

	go func() {
		for throttle.Next() {
			fmt.Println("hello trailing")
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			throttle.Trigger()
			time.Sleep(period / 4)
		}
	}()

	time.Sleep(2 * period)

	throttle.Stop()

	// Output: hello trailing
	// hello trailing
}

func ExampleThrottleFunc() {
	throttle := throttle.ThrottleFunc(period, false, func() {
		fmt.Println("fun, throttled.")
	})

	go func() {
		for i := 0; i < 5; i++ {
			throttle.Trigger()
			time.Sleep(period / 6)
		}
	}()

	time.Sleep(2 * period)
	throttle.Stop()

	// Output: fun, throttled.
}
