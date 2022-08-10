package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	done := make(chan struct{})
	for _, ch := range channels {
		go func(ch <-chan interface{}, done chan struct{}) {
			for {
				_, ok := <-ch
				if !ok {
					done <- struct{}{}
					return
				}
			}
		}(ch, done)
	}
	go func() {
		for {
			<-done
			wg.Done()
			return
		}
	}()
	wg.Wait()
	fmt.Println("merge")
	return merge(channels)
}

func merge(chs []<-chan interface{}) <-chan interface{} {
	merged := make(chan interface{})
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))
		for _, ch := range chs {
			go func(ch <-chan interface{}) {
				defer wg.Done()
				for data := range ch {
					merged <- data
				}
			}(ch)
		}
		wg.Wait()
		close(merged)
	}()
	return merged
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(4*time.Second),
		sig(3*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
