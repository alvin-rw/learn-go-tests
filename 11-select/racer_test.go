package racer

import "testing"

func TestRacer(t *testing.T) {
	fastURL := "http://www.facebook.com"
	slowURL := "http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
