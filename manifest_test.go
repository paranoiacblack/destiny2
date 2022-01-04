package destiny2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUpdate(t *testing.T) {
	manifest, err := NewManifest()
	if err != nil {
		t.Fatal(err)
	}
	defer manifest.Close()

	status, err := manifest.Update()
	if err != nil {
		t.Fatal(err)
	}

	if status != AlreadyUpdated {
		t.Error("Want UpdateStatus AlreadyUpdated got Successful")
	}
}

func TestFulfillContract(t *testing.T) {
	manifest, err := NewManifest()
	if err != nil {
		t.Fatal(err)
	}
	defer manifest.Close()

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
		t.Fatalf("(%T).Entity(%d) returned different entity from direct access: %s", lore, hash, diff)
	}
}

func TestFulfillContract_Mobile(t *testing.T) {
	manifest, err := NewManifest()
	if err != nil {
		t.Fatal(err)
	}
	defer manifest.Close()

	var lore LoreDefinition
	if err := manifest.FulfillContract(&lore, UseMobileManifest(true)); err != nil {
		t.Fatal(err)
	}

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
		t.Fatalf("(%T).Entity(%d) returned different entity from direct access: %s", lore, hash, diff)
	}
}

// For now, just some sanity checking that the mobile manifest matches the component paths.
func TestFulfillContract_Comparison(t *testing.T) {
	manifest, err := NewManifest()
	if err != nil {
		t.Fatal(err)
	}
	defer manifest.Close()

	var lore LoreDefinition
	if err := manifest.FulfillContract(&lore); err != nil {
		t.Fatal(err)
	}

	var mobileLore LoreDefinition
	if err := manifest.FulfillContract(&mobileLore, UseMobileManifest(true)); err != nil {
		t.Fatal(err)
	}

	var hash uint32
	var entity LoreEntity
	for k, v := range lore {
		hash = k
		entity = v
		break
	}

	mobileEntity := mobileLore[hash]
	if diff := cmp.Diff(entity, mobileEntity); diff != "" {
		t.Fatalf("(%T)[%d] returned different entity from mobile manifest: %s", lore, hash, diff)
	}
}

func TestFulfillContract_All(t *testing.T) {
	allContracts := []Contract{
		new(InventoryItemDefinition),
		new(ProgressionDefinition),
		new(InventoryBucketDefinition),
		new(ItemTierTypeDefinition),
		new(StatDefinition),
		new(StatGroupDefinition),
		new(EquipmentSlotDefinition),
		new(SocketTypeDefinition),
		new(SocketCategoryDefinition),
		new(DestinationDefinition),
		new(ActivityGraphDefinition),
		new(ActivityDefinition),
		new(ActivityModifierDefinition),
		new(ObjectiveDefinition),
		new(SandboxPerkDefinition),
		new(LocationDefinition),
		new(ActivityModeDefinition),
		new(PlaceDefinition),
		new(ActivityTypeDefinition),
		new(VendorGroupDefinition),
		new(FactionDefinition),
		new(ArtifactDefinition),
		new(PowerCapDefinition),
		new(ProgressionLevelRequirementDefinition),
		new(RewardSourceDefinition),
		new(TraitDefinition),
		new(TraitCategoryDefinition),
		new(PresentationNodeDefinition),
		new(CollectibleDefinition),
		new(MaterialRequirementSetDefinition),
		new(RecordDefinition),
		new(GenderDefinition),
		new(LoreDefinition),
		new(MetricDefinition),
		new(EnergyTypeDefinition),
		new(PlugSetDefinition),
		new(TalentGridDefinition),
		new(DamageTypeDefinition),
		new(ItemCategoryDefinition),
		new(BreakerTypeDefinition),
		new(SeasonDefinition),
		new(SeasonPassDefinition),
		new(ChecklistDefinition),
		new(RaceDefinition),
		new(ClassDefinition),
		new(MilestoneDefinition),
		new(UnlockDefinition),
		new(ReportReasonCategoryDefinition),
	}

	manifest, err := NewManifest()
	if err != nil {
		t.Fatal(err)
	}
	defer manifest.Close()

	for _, contract := range allContracts {
		if err := manifest.FulfillContract(contract); err != nil {
			t.Errorf("FulfillContract(%q): %v", contract.Name(), err)
		}
	}
}
