package tlds

import "testing"

func TestPopularDomains(t *testing.T) {
	if len(PopularDomains) != 5 {
		t.Error("Unable to find the correct number of known popular domain names")
	}
}

func TestStreamPopularDomains(t *testing.T) {
	var counter = 0
	for _ = range StreamPopularDomains() {
		counter++
	}
	if counter != 5 {
		t.Error("Unable to find the correct number of known popular domain names while streaming")
	}
}

func TestCountryDomains(t *testing.T) {
	if len(CountryDomains) != 252 {
		t.Error("Unable to find the correct number of top country domain names")
	}
}

func TestStreamCountryDomains(t *testing.T) {
	var counter = 0
	for _ = range StreamCountryDomains() {
		counter++
	}
	if counter != 252 {
		t.Error("Unable to find the correct number of top country domain names")
	}
}
