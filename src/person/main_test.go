package main

import (
	"testing"
)

func TestPersonResolver(t *testing.T) {
	data, err := handle(personPayload{ID: 2})

	if err != nil {
		t.Errorf("Error must be nil, got: %v", err)

		return
	}

	if data.ID != 2 {
		t.Errorf("Must return person with id 2, got: %d", data.ID)
	}

	if data.Name != "Paul Gascoigne" {
		t.Errorf("Must return person with name Paul Gascoigne, got: %s", data.Name)
	}
}
