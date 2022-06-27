package main

import (
	"fmt"
	"time"

	"github.com/esimov/gogu"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("==================Delay")
	ch := make(chan struct{})
	t := gogu.Delay(time.Millisecond*500, func() {
		fmt.Println("Function executed after 0.5 second.")
		ch <- struct{}{}
	})
	<-ch
	defer t.Stop()

	fmt.Println("==================After")
	sample := []int{1, 2, 3, 4, 5, 6}
	length := len(sample) - 1
	gogu.ForEach[int](sample, func(val int) {
		fmt.Printf("Printing value... %d\n", val)
		gogu.After[string, int](&length, func() {
			time.Sleep(time.Millisecond * 100)
			fmt.Println("save after")
		})
	})

	fmt.Println("==================Before")
	var n int = 3
	gogu.ForEach[int](sample, func(val int) {
		res := gogu.Before[string, int](&n, func() string {
			time.Sleep(time.Millisecond * 100)
			return "memoized function"
		})
		fmt.Printf("Printing value... %d %v\n", val, res)
	})

	fmt.Println("==================Once")
	n = 2
	gogu.ForEach[int](sample, func(val int) {
		res := gogu.Once[string, int](n, func() string {
			time.Sleep(time.Millisecond * 100)
			return "invoked once"
		})
		fmt.Printf("Printing value... %d %v\n", val, res)
	})

	fmt.Println("==================Retry")
	n = 4
	gogu.ForEach[string]([]string{"one", "two", "three"}, func(val string) {
		rt := gogu.RType[string]{Input: val}
		r, e := rt.Retry(n, func(elem string) (err error) {
			if len(elem)%3 != 0 {
				err = fmt.Errorf("retry failed: number of %d attempts exceeded", n)
			}
			return err
		})
		fmt.Println(r, e)
	})

	fmt.Println("==================RetryWithDelay")
	gogu.ForEach[string]([]string{"one", "two", "three"}, func(val string) {
		rt := gogu.RType[string]{Input: val}
		duration, r, e := rt.RetryWithDelay(n, time.Second, func(d time.Duration, elem string) (err error) {
			if len(elem)%3 != 0 {
				err = fmt.Errorf("retry failed: number of %d attempts exceeded", n)
			}
			return err
		})
		fmt.Println(duration.String(), r, e)
	})

	fmt.Println("==================Retry Struct")
	services := []struct {
		service string
		time    int
	}{
		{
			service: "AWS1",
			time:    10,
		}, {
			service: "AWS2",
			time:    20,
		},
	}

	type Service[T ~string, N constraints.Integer] struct {
		Service T
		Time    N
	}

	for _, srv := range services {
		service := Service[string, int]{
			Service: srv.service,
			Time:    srv.time,
		}
		rtyp := gogu.RType[Service[string, int]]{
			Input: service,
		}

		r, e := rtyp.Retry(n, func(srv Service[string, int]) (err error) {
			if srv.Time > 10 {
				err = fmt.Errorf("retry failed: service time exceeded")
			}
			return err
		})
		fmt.Println(r, e)
	}

	fmt.Println("==================Debounce")
	f := func() {
		fmt.Println("DEBOUNCING - might be doing a time consuming operation...")
	}

	debounce, cancel := gogu.NewDebounce(500 * time.Millisecond)
	for i := 0; i < 2; i++ {
		debounce(f)
		time.Sleep(time.Second)
	}
	fmt.Println("FINISHED!")
	cancel()

	fmt.Println("==================Throttle")
	period := time.Second
	throttle, cancel := gogu.NewThrottle(period, true)

	go func() {
		for i := 0; i < 10; i++ {
			throttle(func() {
				fmt.Println("throttled function invoked")
			})
			time.Sleep(period / 2)
		}
	}()
	time.Sleep(2 * period)
	cancel()
}
