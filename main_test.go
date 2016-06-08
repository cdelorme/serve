package main

import (
	"os"
	"testing"
)

func init() {
	run = func(_ string) error { return nil }
}

func TestPlacebo(t *testing.T) {
	t.Parallel()
	if !true {
		t.FailNow()
	}
}

func TestGetPortDefault(t *testing.T) {
	os.Unsetenv("PORT")
	os.Args = os.Args[:1]
	if p := getPort(); p != "3000" {
		t.Logf("expected 3000, got %s", p)
		t.FailNow()
	}
}

func TestGetPortEnv(t *testing.T) {
	os.Setenv("PORT", "4000")
	os.Args = os.Args[:1]
	if p := getPort(); p != "4000" {
		t.Logf("expected 4000, got %s", p)
		t.FailNow()
	}
}

func TestGetPortArgs(t *testing.T) {
	os.Unsetenv("PORT")
	os.Args = append(os.Args[:1], "5000")
	if p := getPort(); p != "5000" {
		t.Logf("expected 5000, got %s", p)
		t.FailNow()
	}
}

func TestMain(_ *testing.T) {
	os.Args = append(os.Args[:1], "forced-failure")
	main()
}
