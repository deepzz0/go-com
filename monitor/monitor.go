package monitor

import (
	"code.ccplaying.com/foundation/log"
	"os"
	"os/signal"
	"sort"
	"syscall"
)

type hook struct {
	Name  string
	Order int
	Exec  func()
}
type hooks []hook

func (h hooks) Len() int           { return len(h) }
func (h hooks) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h hooks) Less(i, j int) bool { return h[i].Order < h[j].Order }

var exitHooks hooks

func init() {
	go schedule()
}

func callHooksAndExit() {
	for _, hook := range exitHooks {
		hook.Exec()
		log.Printf("%s hook exec!\n", hook.Name)
	}
	log.Println("\033[043;1m[SECURE EXIT]\033[0m")
	os.Exit(0)
}

func schedule() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)

	for {
		msg := <-ch
		switch msg {
		case syscall.SIGHUP:
			log.Println("\033[043;1m[SIGHUP]\033[0m")

		case syscall.SIGTERM:
			log.Println("\033[043;1m[SIGTERM]\033[0m")
			callHooksAndExit()
		case syscall.SIGINT:
			log.Println("\033[043;1m[SIGINT]\033[0m")
			callHooksAndExit()
		}
	}
}

func RegistExitFunc(name string, f func(), order ...int) {
	h := hook{Name: name, Order: 100, Exec: f}
	if len(order) > 0 {
		h.Order = order[0]
	}
	exitHooks = append(exitHooks, h)
	sort.Sort(exitHooks)
}
