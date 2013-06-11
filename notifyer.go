package gonotify

import (
	"sync"
)

type Notifyer struct {
	subscribers         []chan bool
	subscriptionAllowed bool
	lock                sync.Mutex
}

func New() *Notifyer {
	return &Notifyer{subscribers: make([]chan bool, 0), subscriptionAllowed: true}
}

func (this *Notifyer) Subscribe() (<-chan bool, bool) {
	z := make(chan bool, 1)
	this.lock.Lock()
	as := this.subscriptionAllowed
	this.lock.Unlock()
	if !as {
		z <- false
		return z, false
	}
	this.subscribers = append(this.subscribers, z)
	return z, true
}

func (this *Notifyer) Notify() bool {
	this.lock.Lock()
	val := this.subscriptionAllowed
	this.subscriptionAllowed = false
	this.lock.Unlock()
	if val {
		for _, z := range this.subscribers {
			z <- true
		}
		return true
	}
	return false
}
