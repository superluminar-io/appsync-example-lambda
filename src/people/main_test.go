package main

import (
	"testing"
)

func TestPeopleResolver(t *testing.T) {
	data, err := handle()

	if err != nil {
		t.Errorf("Error must be nil")
	}

	if len(data) != 4 {
		t.Errorf("Must return 4 people, got: %d", len(data))
	}
}
