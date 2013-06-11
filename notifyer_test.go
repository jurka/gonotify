package gonotify

import "testing"

func TestNotify(t *testing.T) {

	base := New()

	k1, ok := base.Subscribe()
	if !ok {
		t.Error("Subscribe fail for k1")
	}
	k2, ok := base.Subscribe()
	if !ok {
		t.Error("Subscribe fail for k2")
	}
	k3, ok := base.Subscribe()
	if !ok {
		t.Error("Subscribe fail for k3")
	}

	go func() {
		for {
			select {
			case v1 := <-k1:
				if !v1 {
					t.Error("Wrong value on k1 channel")
				}
			case v2 := <-k2:
				if !v2 {
					t.Error("Wrong value on k3 channel")
				}
			case v3 := <-k3:
				if !v3 {
					t.Error("Wrong value on k3 channel")
				}
			}
		}
	}()

	z := base.Notify()
	if !z {
		t.Error("Unable to notify")
	}
	_, ok = base.Subscribe()
	if ok {
		t.Error("Able to subscrite to already run notifyer")
	}

	z2 := base.Notify()
	if z2 {
		t.Error("Notify allowed to wun twice")
	}
}
