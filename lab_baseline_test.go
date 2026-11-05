package main

import "testing"

func TestLabBaselineRedCheck(t *testing.T) {
	t.Fatal("intentional failure: proves a red check does not block the merge")
}
