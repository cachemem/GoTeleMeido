package main

import (
	"testing"
)

func TestHelp(t *testing.T) {
	bot := NewBot(nil)
	want := "* /reverse — reverse whatever text want\n* /8ball — ask a magic 8-ball"
	if got := bot.help(); got != want {
		t.Errorf("Help command is incorrect - got %q, wanted %q", got, want)
	}
}

type testR struct{}

func (r testR) Reverse(s string) string {
	return s + "!"
}

func TestReverse(t *testing.T) {
	bot := NewBot(testR{})
	param := "XyZ"
	want := "XyZ!"
	if got := bot.reverse(&param); got != want {
		t.Errorf("Reverse command is incorrect - got %q, wanted %q", got, want)
	}
}
