package api

import (
	"fmt"
	"net/http"
)

type Fish struct {
	name           string
	desc           string
	healthBenefits string
	taste          string
	quote          string
	popStatus      string
}

func newFish(r *http.Response) *Fish {
	f := Fish{
		name:           r.Header.Get("Species Name"),
		desc:           r.Header.Get("Species Name"),
		healthBenefits: r.Header.Get("Species Name"),
		taste:          r.Header.Get("Species Name"),
		quote:          r.Header.Get("Species Name"),
		popStatus:      r.Header.Get("Species Name"),
	}
	return &f
}

func (f *Fish) getQuote() string {
	return f.quote
}

func (f *Fish) toString() string {
	return fmt.Sprintf("Fish Species: %s\nDescription: %s\nTaste: %s\nHealth Benefits: %s",
		f.name, f.desc, f.taste, f.healthBenefits)
}
