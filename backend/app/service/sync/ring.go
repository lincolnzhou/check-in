package sync

import (
	"github.com/labstack/gommon/log"
	"github.com/lincolnzhou/check-in/backend/app/libs/errors"
)

type Ring struct {
	// read
	rn int64
	rp int

	// write
	wn int64
	wp int

	// info
	num  int
	data []Check
}

func NewRing(num int) *Ring {
	r := new(Ring)
	r.data = make([]Check, num)
	r.num = num
	return r
}

func (r *Ring) Get() (check *Check, err error) {
	if r.wn == r.rn {
		return nil, errors.ErrorRingEmpty
	}

	check = &r.data[r.rp]
	return
}

func (r *Ring) GetAdv() {
	if r.rp++; r.rp >= r.num {
		r.rp = 0
	}

	r.rn++
	//if conf.ConfigData.Debug {
	log.Debug("ring rn: %d, rp: %d", r.rn, r.rp)
	//}
}

func (r *Ring) Set() (check *Check, err error) {
	if r.Buffer() >= r.num {
		return nil, errors.ErrorRingFull
	}

	check = &r.data[r.wp]
	return
}

func (r *Ring) SetAdv() {
	if r.wp++; r.wp >= r.num {
		r.wp = 0
	}

	r.wn++
	// if ConfigData.Debug {
	log.Debug("ring wn: %d, wp: %d", r.wn, r.wp)
	//}
}

func (r *Ring) Buffer() int {
	return int(r.wp - r.rp)
}

func (r *Ring) Reset() {
	r.rn = 0
	r.rp = 0
	r.wn = 0
	r.wp = 0
}

func (r *Ring) Debug() {
	log.Debug("ring wn: %d, wp: %d, rn: %d, rp: %d", r.wn, r.wp, r.rn, r.rp)
}
