package monitor

import (
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"
)

type exitHook struct {
	Name  string
	Order int
	Call  func()
}
type exitHooks []exitHook

func (h exitHooks) Len() int           { return len(h) }
func (h exitHooks) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h exitHooks) Less(i, j int) bool { return h[i].Order < h[j].Order }

var registExitHooks exitHooks

func (hooks exitHooks) exec() {
	log.Println("start secure exit...")
	bt := time.Now()
	for _, hook := range registExitHooks {
		hbt := time.Now()
		hook.Call()
		log.Printf("safeexit hook (%s) exec cost: %v\n", hook.Name, time.Now().Sub(hbt))
	}
	log.Printf("\033[043;1m[SECURE EXIT]\033[0m cost: %v\n", time.Now().Sub(bt))
	os.Exit(0)
}

func monitor() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	for {
		msg := <-ch
		switch msg {
		case syscall.SIGHUP:
			log.Print("\033[043;1m[SIGHUP]\033[0m")
		case syscall.SIGTERM:
			registExitHooks.exec()
		case syscall.SIGINT:
			registExitHooks.exec()
		}
	}
}

func HookOnExit(name string, f func(), order ...int) {
	h := exitHook{Name: name, Order: 100, Call: f}
	if len(order) > 0 {
		h.Order = order[0]
	}
	registExitHooks = append(registExitHooks, h)
	sort.Sort(registExitHooks)
}

func Startup() {
	go monitor()
}
