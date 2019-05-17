package main

import (
	"testing"
)

func TestPersonFriendsResolver(t *testing.T) {
	data, err := handle(personFriendsPayload{ID: 2})

	if err != nil {
		t.Errorf("Error must be nil, got: %v", err)

		return
	}

	if len(data) != 2 {
		t.Errorf("Must return 2 friends, got: %d", len(data))
	}
}
