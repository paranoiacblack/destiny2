package destiny2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
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

type fulfillmentTests struct {
	name string
	fn   func(t *testing.T)
}

func TestFulfillContract(t *testing.T) {
	reader := NewTestReader()
	manifest, err := NewManifest(reader)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	tests := []fulfillmentTests{
		{
			name: "Gender DisplayProperties",
			fn: func(t *testing.T) {
				var genders GenderDefinition
				if err := manifest.FulfillContract(&genders); err != nil {
					t.Fatal(err)
				}

				genderProps := map[Gender]DisplayProperties{
					Gender_Male:   DisplayProperties{Name: "Masculine"},
					Gender_Female: DisplayProperties{Name: "Feminine"},
				}

				for _, gender := range genders {
					prop := genderProps[gender.GenderType]
					if diff := cmp.Diff(prop, gender.DisplayProperties); diff != "" {
						t.Errorf("Gender DisplayProperties differ: %s", diff)
					}
				}
			},
		},
		{
			name: "Gjallarhorn Lore",
			fn: func(t *testing.T) {
				const gjallarhornHash = 1363886209

				var items InventoryItemDefinition
				var lore LoreDefinition
				if err := manifest.FulfillContract(&items); err != nil {
					t.Fatal(err)
				}
				if err := manifest.FulfillContract(&lore); err != nil {
					t.Fatal(err)
				}

				horn := items[gjallarhornHash]
				hornLore := lore[horn.LoreHash]
				if hornLore.Subtitle != horn.FlavorText {
					t.Errorf("Gjallarhorn lore is incorrect: want %q, got %q", horn.FlavorText, hornLore.Subtitle)
				}
			},
		},
		{
			name: "Starhorse Location",
			fn: func(t *testing.T) {
				const starhorseHash = 3431983428
				starhorseLocation := DisplayProperties{
					Name:        "Eternity",
					Description: "A strange pocket dimension where multiple realities from across the paraverse converge and overlap.",
				}

				var vendors VendorDefinition
				var destinations DestinationDefinition
				if err := manifest.FulfillContract(&vendors); err != nil {
					t.Fatal(err)
				}
				if err := manifest.FulfillContract(&destinations); err != nil {
					t.Fatal(err)
				}

				starhorse := vendors[starhorseHash]
				destination := destinations[starhorse.Locations[0].DestinationHash]
				if diff := cmp.Diff(destination.DisplayProperties, starhorseLocation); diff != "" {
					t.Errorf("Starhorse's location is incorrect: %s", diff)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}

// Testing that all contracts can be fulfilled without error.
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

	reader := NewTestReader()
	manifest, err := NewManifest(reader)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()

	for _, contract := range allContracts {
		if err := manifest.FulfillContract(contract); err != nil {
			t.Errorf("FulfillContract(%q): %v", contract.Name(), err)
		}
	}
}

// NewTestReader creates a ContractReader for testing.
func NewTestReader() *sqliteReader {
	return &sqliteReader{cachedContracts: map[string][]byte{}}
}

// sqliteReader will implement ContractReader to allow contract fulfillment from a cached database.
type sqliteReader struct {
	cachedContracts map[string][]byte
}

func (r *sqliteReader) ReadContract(contract Contract, path string, useMobile bool) ([]byte, error) {
	name := contract.Name()
	if data, ok := r.cachedContracts[name]; ok {
		return data, nil
	}

	data, err := readContractFromDB("testdata/mobile_manifest_en.sqlite", name)
	if err != nil {
		return nil, err
	}
	r.cachedContracts[name] = data
	return data, nil
}

func (r *sqliteReader) Close() error {
	r.cachedContracts = nil
	return nil
}
