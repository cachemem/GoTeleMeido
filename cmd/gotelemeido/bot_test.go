package main

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

type TestRepo struct {
	ownerId int64
}

func (tr *TestRepo) GetOwner() int64 {
	return tr.ownerId
}

func TestHelp(t *testing.T) {
	bot := NewBot(nil, &TestRepo{ownerId: 0})
	want := "* /reverse — reverse whatever text want\n* /8ball — ask a magic 8-ball"
	if got := bot.help(); got != want {
		t.Errorf("Help command is incorrect - got %q, wanted %q", got, want)
	}
}

type GreetingsTest struct {
	realOwnerId int64
	ownerId     int64
	response    string
}

func (gt *GreetingsTest) Generate(rand *rand.Rand, size int) reflect.Value {
	realOwnerId := rand.Int63n(1000)
	ownerId := rand.Int63n(1000)
	var response string
	switch realOwnerId == ownerId {
	case true:
		response = "Welcome back, master!"
	case false:
		response = "Hai~ " + helpText
	}

	return reflect.ValueOf(&GreetingsTest{
		realOwnerId: realOwnerId,
		ownerId:     ownerId,
		response:    response,
	})
}

func TestGreetings(t *testing.T) {
	bot := NewBot(nil, &TestRepo{ownerId: 100500})

	tableTests := []struct {
		input int64
		want  string
	}{
		{123, "Hai~ " + helpText},
		{100500, "Welcome back, master!"},
	}
	for _, test := range tableTests {
		if got := bot.greetings(test.input); got != test.want {
			t.Errorf("Greetings were incorrect: %q, got %q", test.input, got)
		}
	}

	f := func(test *GreetingsTest) bool {
		bot := NewBot(nil, &TestRepo{ownerId: test.realOwnerId})
		return bot.greetings(test.ownerId) == test.response
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 1000000}); err != nil {
		t.Fatal(err)
	}
}

type testR struct{}

func (r testR) Reverse(s string) string {
	return s + "!"
}

func TestReverse(t *testing.T) {
	bot := NewBot(testR{}, &TestRepo{ownerId: 0})
	param := "XyZ"
	want := "XyZ!"
	if got := bot.reverse(&param); got != want {
		t.Errorf("Reverse command is incorrect - got %q, wanted %q", got, want)
	}
}
