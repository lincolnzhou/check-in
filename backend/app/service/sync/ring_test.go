package sync

import "testing"

func TestRing(t *testing.T) {
	r := NewRing(3)
	p0, err := r.Set()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	p0.ID = 1
	r.SetAdv()

	p1, err := r.Set()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	p1.ID = 1
	r.SetAdv()

	r.Debug()
}
