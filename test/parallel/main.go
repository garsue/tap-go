package main

import (
	"sync"

	"github.com/mndrix/tap-go"
)

func main() {
	t := tap.New()
	t.Header(0)

	var wg sync.WaitGroup
	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c1 := t.Count()
			t.Ok(true, "a test")
			c2 := t.Count()
			t.Ok(c2 > c1, "count up")
			t.Skip(1, "skip")
		}()
	}
	wg.Wait()

	t.AutoPlan()
}
