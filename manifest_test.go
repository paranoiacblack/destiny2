package destiny2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUpdate(t *testing.T) {
	var manifest Manifest
	if err := manifest.Update(); err != nil {
		t.Fatal(err)
	}
}

func TestFulfillContract(t *testing.T) {
	var manifest Manifest
	if err := manifest.Update(); err != nil {
		t.Fatal(err)
	}

	var lore LoreDefinition
	if err := manifest.FulfillContract(&lore); err != nil {
		t.Fatal(err)
	}

	// Find a random lore entry and compare result to entry found by Contract.Entity.
	// TODO(paranoiacblack): Ideally allow the ability to use an in-memory manifest
	// so these tests can use table-driven tests rather than relying on the manifest
	// entries from when this test is run. Should be easy enough by making Manifest an
	// interface.
	var hash uint32
	var entity LoreEntity
	for k, v := range lore {
		hash = k
		entity = v
		break
	}

	lookup, ok := lore.Entity(hash).(LoreEntity)
	if !ok {
		t.Fatalf("(%T).Entity(%d) returned invalid typed entity", lore, hash)
	}

	if diff := cmp.Diff(entity, lookup); diff != "" {
		t.Fatalf("(%T).Entity(%d) returned different entity from direct access", lore, hash)
	}
}
