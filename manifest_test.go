package destiny2

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUpdate(t *testing.T) {
	manifest, err := NewManifest(nil)
	if err != nil {
		t.Fatal(err)
	}

	updated := false
	onUpdate := func() error {
		updated = true
		return nil
	}
	if err := manifest.Update(onUpdate); err != nil {
		t.Fatal(err)
	}

	if updated {
		t.Error("Update after NewManifest should not be necessary")
	}

	manifest = new(Manifest)
	if err := manifest.Update(onUpdate); err != nil {
		t.Fatal(err)
	}

	if !updated {
		t.Error("Updating empty manifest should be necessary")
	}
}

var testContracts = map[string]Contract{
	"DestinyInventoryItemDefinition": &InventoryItemDefinition{
		2: {ItemTypeDisplayName: "Dummy", Equippable: true},
	},
}

func TestFulfillContract(t *testing.T) {
	reader := &testReader{contracts: testContracts}
	manifest, err := NewManifest(reader)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	var items InventoryItemDefinition
	if err := manifest.FulfillContract(&items); err != nil {
		t.Fatal(err)
	}

	for k, v := range items {
		testEntity := testContracts[items.Name()].Entity(k)
		if diff := cmp.Diff(testEntity, v); diff != "" {
			t.Errorf("TestEntity differs from fulfilled contract at hash %d: %s", k, diff)
		}
	}
}

func TestFulfillContract_Mobile(t *testing.T) {
	bungieReader := new(BungieAPIReader)
	manifest, err := NewManifest(bungieReader)
	if err != nil {
		t.Fatal(err)
	}
	defer bungieReader.Close()

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
	bungieReader := new(BungieAPIReader)
	manifest, err := NewManifest(bungieReader)
	if err != nil {
		t.Fatal(err)
	}
	defer bungieReader.Close()

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
		new(ProgressionDefinition),
		new(InventoryItemDefinition),
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

	bungieReader := new(BungieAPIReader)
	manifest, err := NewManifest(bungieReader)
	if err != nil {
		t.Fatal(err)
	}
	defer bungieReader.Close()

	for _, contract := range allContracts {
		if err := manifest.FulfillContract(contract); err != nil {
			t.Errorf("FulfillContract(%q): %v", contract.Name(), err)
		}
	}
}

// testReader will implement ContractReader to allow table-driven tests.
type testReader struct {
	contracts map[string]Contract
}

func (r *testReader) ReadContract(contract Contract, path string, useMobile bool) ([]byte, error) {
	// Just for rough testing, assume that contracts is filled and return to marshalled version of pre-fulfilled contract.
	return json.Marshal(r.contracts[contract.Name()])
}

func (r *testReader) Close() error { return nil }
