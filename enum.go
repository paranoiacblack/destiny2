package destiny2

// BungieMembershipType describes the types of memberships the accounts system supports.
type BungieMembershipType int32

const (
	Membership_None BungieMembershipType = 0
	TigerXbox                            = 1
	TigerPSN                             = 2
	TigerSteam                           = 3
	TigerBlizzard                        = 4
	TigerStadia                          = 5
	TigerDemon                           = 10
	BungieNext                           = 254
	Membership_All                       = -1 // Membership_All is only valid for searching capabilities.
)

// ProgressionScope is the scope of a progression which determines how it is stored, calculated and used.
type ProgressionScope int32

const (
	Progression_Account               ProgressionScope = 0
	Progression_Character                              = 1
	Progression_Clan                                   = 2
	Progression_Item                                   = 3
	Progression_ImplicitFromEquipment                  = 4
	Progression_Mapped                                 = 5
	Progression_MappedAggregate                        = 6
	Progression_MappedStat                             = 7
	Progression_MappedUnlockValue                      = 8
)

// ProgressionStepDisplayEffect determines whether progression shows visual effects on the character, its item or neither.
type ProgressionStepDisplayEffect int32

const (
	ProgressionStepDisplay_None      ProgressionStepDisplayEffect = 0
	ProgressionStepDisplay_Character                              = 1
	ProgressionStepDisplay_Item                                   = 2
)

// ItemTier is the tier of an item in an inventory.
type ItemTier int32

const (
	ItemTier_Unknown  ItemTier = 0
	ItemTier_Currency          = 1
	ItemTier_Basic             = 2
	ItemTier_Common            = 3
	ItemTier_Rare              = 4
	ItemTier_Superior          = 5
	ItemTier_Exotic            = 6
)

// BucketScope is scope of an inventory bucket.
type BucketScope int32

const (
	BucketScope_Profile   BucketScope = 0
	BucketScope_Character             = 1
)

// BucketCategory describes what categories of items can be found in an inventory bucket.
type BucketCategory int32

const (
	BucketCategory_Invisible  BucketCategory = 0
	BucketCategory_Item                      = 1
	BucketCategory_Currency                  = 2
	BucketCategory_Equippable                = 3
	BucketCategory_Ignored                   = 4
)

// ItemLocation describes where a given item can be found.
type ItemLocation int32

const (
	ItemLocation_Unknown    ItemLocation = 0
	ItemLocation_Inventory               = 1
	ItemLocation_Vault                   = 2
	ItemLocation_Vendor                  = 3
	ItemLocation_Postmaster              = 4
)

// StatAggregationType describes the rule used for determining the level
// and the formula used for aggregation of stats.
type StatAggregationType int32

const (
	// CharacterAverage applies a weighted average using stat grouping and items equipped on a character.
	StatAggregation_CharacterAverage StatAggregationType = 0
	StatAggregation_Character                            = 1 // Use only the character's stat grouping.
	StatAggregation_Item                                 = 2 // Use only the character's equipped items.
)

// StatCategory describes the categories which a given stat affects.
type StatCategory int32

const (
	StatCategory_Gameplay StatCategory = 0
	StatCategory_Weapon                = 1
	StatCategory_Defense               = 2
	StatCategory_Primary               = 3
)

// EquippingItemBlockAttributes are custom attributes on the equippability of the item.
type EquippingItemBlockAttributes int32

const (
	EquippingAttribute_None           EquippingItemBlockAttributes = 0
	EquippingAttribute_EquipOnAcquire                              = 1
)

// AmmunitionType is the ammunition a weapon used.
type AmmunitionType int32

const (
	Ammunition_None    AmmunitionType = 0
	Ammunition_Primary                = 1
	Ammunition_Special                = 2
	Ammunition_Heavy                  = 3
	Ammunition_Unknown                = 4
)

// VendorProgressionType describes the type of progression a vendor has.
type VendorProgressionType int32

const (
	// Default is vendor rank progression from token redemption.
	VendorProgression_Default VendorProgressionType = 0
	// Rank progressions from ritual/playlist content such as Crucible, Gambit and Battlegrounds.
	VendorProgression_Ritual = 1
	// Rank progressions that cannot be reset in a given seasons, such as Xur in the Eternity destination.
	VendorProgression_NoSeasonalRefresh = 2
)

// VendorInteractionRewardSelection determines how many rewards are provided on selection.
type VendorInteractionRewardSelection int32

const (
	VendorSelection_None VendorInteractionRewardSelection = 0
	VendorSelection_One                                   = 1
	VendorSelection_All                                   = 2
)

// VendorReplyType determines the type of reply that a vendor will have during an interaction.
type VendorReplyType int32

const (
	VendorReply_Accept   VendorReplyType = 0
	VendorReply_Decline                  = 1
	VendorReply_Complete                 = 2
)

// VendorInteractionType is an enumeration of the known UI interactions for Vendors.
type VendorInteractionType int32

const (
	VendorInteraction_Unknown VendorInteractionType = 0
	// An empty interaction. If this ends up in content, it is probably a game bug.
	VendorInteraction_Undefined = 1
	// An interaction shown when you complete a quest and receive a reward.
	VendorInteraction_QuestComplete = 2
	// An interaction shown when you talk to a Vendor as an intermediary step of a quest.
	VendorInteraction_QuestContinue = 3
	// An interaction shown when you are previewing the vendor's reputation rewards.
	VendorInteraction_ReputationPreview = 4
	// An interaction shown when you rank up with the vendor.
	VendorInteraction_RankUpReward = 5
	// An interaction shown when you have tokens to turn in for the vendor.
	VendorInteraction_TokenTurnin = 6
	// An interaction shown when you're accepting a new quest.
	VendorInteraction_QuestAccept = 7
	VendorInteraction_ProgressTab = 8
	VendorInteraction_End         = 9
	VendorInteraction_Start       = 10
)

// VendorItemRefundPolicy is the action that happens when the user attempts to refund an item.
type VendorItemRefundPolicy int32

const (
	RefundPolicy_NotRefundable  VendorItemRefundPolicy = 0
	RefundPolicy_DeletesItem                           = 1
	RefundPolicy_RevokesLicense                        = 2
)

// GatingScope represents the most restrictive type of gating that is being performed by an entity.
type GatingScope int32

const (
	// No gating on this item.
	Gating_None GatingScope = 0
	// Gating on this item is based on global game state. It will be gated the same for everyone.
	Gating_Global = 1
	// Gating on this item is at the Clan level.
	Gating_Clan = 2
	// Gating includes Profile-specific checks, but not on the Profile's characters.
	Gating_Profile = 3
	// Gating includes Character-specific checks, including character level restrictions.
	Gating_Character = 4
	// Gating includes item-specific checks.
	Gating_Item = 5
	// Gating unlocks and checks are unknown.
	Gating_AssumedWorstCase = 6
)

// SocketTypeActionType indicates the types of actions that can be performed.
type SocketTypeActionType int32

const (
	InsertPlug         SocketTypeActionType = 0
	InfuseItem                              = 1
	ReinitializeSocket                      = 2
)

// SocketVisibility indicates whether a socket is visible.
type SocketVisibility int32

const (
	SocketVisibility_Visible                  SocketVisibility = 0
	SocketVisibility_Hidden                                    = 1
	SocketVisibility_HiddenWhenEmpty                           = 2
	SocketVisibility_HiddenIfNoPlugsAvailable                  = 3
)

// SocketCategoryStyle represents the possible and known UI styles used for rendering socket categories.
type SocketCategoryStyle int32

const (
	SocketCategory_Unknown     SocketCategoryStyle = 0
	SocketCategory_Reusable                        = 1
	SocketCategory_Consumable                      = 2
	SocketCategory_Unlockable                      = 3
	SocketCategory_Intrinsic                       = 4
	SocketCategory_EnergyMeter                     = 5
	SocketCategory_LargePerk                       = 6
	SocketCategory_Abilities                       = 7
	SocketCategory_Supers                          = 8
)

// ActivityGraphNodeHighlightType is the known UI style in which an item can be highlighted.
type ActivityGraphNodeHighlightType int32

const (
	ActivityHighlight_None       ActivityGraphNodeHighlightType = 0
	ActivityHighlight_Normal                                    = 1
	ActivityHighlight_Hyper                                     = 2
	ActivityHighlight_Comet                                     = 3
	ActivityHighlight_RiseOfIron                                = 4
)

// UnlockValueUIStyle describes the format unlock values should use in the UI.
type UnlockValueUIStyle int32

const (
	// Automatic means just show the number.
	UnlockValue_Automatic          UnlockValueUIStyle = 0
	UnlockValue_Fraction                              = 1  // Show the number as a fractional value; this is relative to the context-specific upper bound.
	UnlockValue_Checkbox                              = 2  // Show the number as a checkbox.
	UnlockValue_Percentage                            = 3  // Show the number as a percentage; this is relative to the context-specific upper bound.
	UnlockValue_DateTime                              = 4  // Show the number as a date and time; the given number number of seconds since the Unix Epoch.
	UnlockValue_FractionFloat                         = 5  // Show the number as a floating point value that represents a fraction; this is relative to the context-specific upper bound.
	UnlockValue_Integer                               = 6  // Show the number as an integer.
	UnlockValue_TimeDuration                          = 7  // Show the number as a time duration in seconds.
	UnlockValue_Hidden                                = 8  // Don't show the value.
	UnlockValue_Multiplier                            = 9  // Show the value as a multiplier in context, such as "1.5x".
	UnlockValue_GreenPips                             = 10 // Show the value as a series of green pips, like wins on a Trials Card.
	UnlockValue_RedPips                               = 11 // Show the value as a series of red pips, like losses on a Trials Card.
	UnlockValue_ExplicitPercentage                    = 12 // Show the value as a percentage, by appending a %-sign.
	UnlockValue_RawFloat                              = 13 // Show the value as a floating-point number with two-decimal precision.
)

// ObjectiveGrantStyle indicates when a given objected-related perk is granted.
type ObjectiveGrantStyle int32

const (
	Objective_WhenIncomplete ObjectiveGrantStyle = 0
	Objective_WhenComplete                       = 1
	Objective_Always                             = 2
)

// DamageType describes the general categories of damage in the game.
type DamageType int32

const (
	DamageType_None    DamageType = 0
	DamageType_Kinetic            = 1
	DamageType_Arc                = 2
	DamageType_Thermal            = 3
	DamageType_Void               = 4
	DamageType_Raid               = 5
	DamageType_Stasis             = 6
)

// TalentNodeStepWeaponPerformances are basically the stats on the weapon granted by a weapon talent.
// A given talent can apply multiple performance changes to a weapon so use this as a bitmask.
type TalentNodeStepWeaponPerformances int32

const (
	TalentWeaponPerformance_None                  TalentNodeStepWeaponPerformances = 0
	TalentWeaponPerformance_RateOfFire                                             = 1
	TalentWeaponPerformance_Damage                                                 = 2
	TalentWeaponPerformance_Accuracy                                               = 4
	TalentWeaponPerformance_Range                                                  = 8
	TalentWeaponPerformance_Zoom                                                   = 16
	TalentWeaponPerformance_Recoil                                                 = 32
	TalentWeaponPerformance_Ready                                                  = 64
	TalentWeaponPerformance_Reload                                                 = 128
	TalentWeaponPerformance_HairTrigger                                            = 256
	TalentWeaponPerformance_AmmoAndMagazine                                        = 512
	TalentWeaponPerformance_TrackingAndDetonation                                  = 1024
	TalentWeaponPerformance_ShotgunSpread                                          = 2048
	TalentWeaponPerformance_ChargeTime                                             = 4096
	TalentWeaponPerformance_All                                                    = 8191
)

// TalentNodeStepImpactEffects are the effects on bullet impact granted by a talent.
// A given talent can apply multiple effects so use this as a bitmask.
type TalentNodeStepImpactEffects int32

const (
	TalentImpactEffect_None             TalentNodeStepImpactEffects = 0
	TalentImpactEffect_ArmorPiercing                                = 1
	TalentImpactEffect_Ricochet                                     = 2
	TalentImpactEffect_Flinch                                       = 4
	TalentImpactEffect_CollateralDamage                             = 8
	TalentImpactEffect_Disorient                                    = 16
	TalentImpactEffect_HighlightTarget                              = 32
	TalentImpactEffect_All                                          = 63
)

// TalentNodeStepGuardianAttributes are the guardian attributes granted by a talent.
// A given talent can apply multiple guardian attributes so use this as a bitmask.
type TalentNodeStepGuardianAttributes int32

const (
	TalentNodeStepGuardianAttribute_None         TalentNodeStepGuardianAttributes = 0
	TalentNodeStepGuardianAttribute_Stats                                         = 1
	TalentNodeStepGuardianAttribute_Shields                                       = 2
	TalentNodeStepGuardianAttribute_Health                                        = 4
	TalentNodeStepGuardianAttribute_Revive                                        = 8
	TalentNodeStepGuardianAttribute_AimUnderFire                                  = 16
	TalentNodeStepGuardianAttribute_Radar                                         = 32
	TalentNodeStepGuardianAttribute_Invisibility                                  = 64
	TalentNodeStepGuardianAttribute_Reputations                                   = 128
	TalentNodeStepGuardianAttribute_All                                           = 255
)

// TalentNodeStepLightAbilities are the light-related abilities granted by a talent.
// A given talent can apply multiple light abilities so use this as a bitmask.
type TalentNodeStepLightAbilities int32

const (
	TalentLightAbility_None          TalentNodeStepLightAbilities = 0
	TalentLightAbility_Grenades                                   = 1
	TalentLightAbility_Melee                                      = 2
	TalentLightAbility_MovementModes                              = 4
	TalentLightAbility_Orbs                                       = 8
	TalentLightAbility_SuperEnergy                                = 16
	TalentLightAbility_SuperMods                                  = 32
	TalentLightAbility_All                                        = 63
)

// TalentNodeStepDamageTypes are the damage types this talent apply to.
// A given talent can be applied on multiple damage types so use this as a bitmask.
type TalentNodeStepDamageTypes int32

const (
	TalentNodeStepDamageType_None    TalentNodeStepDamageTypes = 0
	TalentNodeStepDamageType_Kinetic                           = 1
	TalentNodeStepDamageType_Arc                               = 2
	TalentNodeStepDamageType_Solar                             = 4
	TalentNodeStepDamageType_Void                              = 8
	TalentNodeStepDamageType_All                               = 15
)

// ActivityModeType is the mode of a given activity.
type ActivityModeType int32

const (
	ActivityMode_None                    ActivityModeType = 0
	ActivityMode_Story                                    = 2
	ActivityMode_Strike                                   = 3
	ActivityMode_Raid                                     = 4
	ActivityMode_AllPvP                                   = 5
	ActivityMode_Patrol                                   = 6
	ActivityMode_AllPvE                                   = 7
	Reserved9                                             = 9
	ActivityMode_Control                                  = 10
	Reserved11                                            = 11
	ActivityMode_Clash                                    = 12
	Reserved13                                            = 13
	ActivityMode_CrimsonDoubles                           = 15
	ActivityMode_Nightfall                                = 16
	ActivityMode_HeroicNightfall                          = 17
	ActivityMode_AllStrikes                               = 18
	ActivityMode_IronBanner                               = 19
	Reserved20                                            = 20
	Reserved21                                            = 21
	Reserved22                                            = 22
	Reserved24                                            = 24
	ActivityMode_AllMayhem                                = 25
	ActivityMode_Supremacy                                = 31
	Reserved26                                            = 26
	Reserved27                                            = 27
	Reserved28                                            = 28
	Reserved29                                            = 29
	Reserved30                                            = 30
	ActivityMode_PrivateMatchesAll                        = 32
	ActivityMode_Survival                                 = 37
	ActivityMode_Countdown                                = 38
	ActivityMode_TrialsOfTheNine                          = 39
	ActivityMode_Social                                   = 40
	ActivityMode_TrialsCountdown                          = 41
	ActivityMode_TrialsSurvival                           = 42
	ActivityMode_IronBannerControl                        = 43
	ActivityMode_IronBannerClash                          = 44
	ActivityMode_IronBannerSupremacy                      = 45
	ActivityMode_ScoredNightfall                          = 46
	ActivityMode_ScoredHeroicNightfall                    = 47
	ActivityMode_Rumble                                   = 48
	ActivityMode_AllDoubles                               = 49
	ActivityMode_Doubles                                  = 50
	ActivityMode_PrivateMatchesClash                      = 51
	ActivityMode_PrivateMatchesControl                    = 52
	ActivityMode_PrivateMatchesSupremacy                  = 53
	ActivityMode_PrivateMatchesCountdown                  = 54
	ActivityMode_PrivateMatchesSurvival                   = 55
	ActivityMode_PrivateMatchesMayhem                     = 56
	ActivityMode_PrivateMatchesRumble                     = 57
	ActivityMode_HeroicAdventure                          = 58
	ActivityMode_Showdown                                 = 59
	ActivityMode_Lockdown                                 = 60
	ActivityMode_Scorched                                 = 61
	ActivityMode_ScorchedTeam                             = 62
	ActivityMode_Gambit                                   = 63
	ActivityMode_AllPvEEcompetitive                       = 64
	ActivityMode_Breakthrough                             = 65
	ActivityMode_BlackArmoryRun                           = 66
	ActivityMode_Salvage                                  = 67
	ActivityMode_IronBannerSalvage                        = 68
	ActivityMode_PvPCompetitive                           = 69
	ActivityMode_PvPQuickplay                             = 70
	ActivityMode_ClashQuickplay                           = 71
	ActivityMode_ClashCompetitive                         = 72
	ActivityMode_ControlQuickplay                         = 73
	ActivityMode_ControlCompetitive                       = 74
	ActivityMode_GambitPrime                              = 75
	ActivityMode_Reckoning                                = 76
	ActivityMode_Menagerie                                = 77
	ActivityMode_VexOffensive                             = 78
	ActivityMode_NightmareHunt                            = 79
	ActivityMode_Elimination                              = 80
	ActivityMode_Momentum                                 = 81
	ActivityMode_Dungeon                                  = 82
	ActivityMode_Sundial                                  = 83
	ActivityMode_TrialsOfOsiris                           = 84
	ActivityMode_Dares                                    = 85
)

// ActivityNavPointType describes the interaction associated with a navigation point in an activity.
type ActivityNavPointType int32

const (
	ActivityNavPoint_Inactive             ActivityNavPointType = 0
	ActivityNavPoint_PrimaryObjective                          = 1
	ActivityNavPoint_SecondaryObjective                        = 2
	ActivityNavPoint_TravelObjective                           = 3
	ActivityNavPoint_PublicEventObjective                      = 4
	ActivityNavPoint_AmmoCache                                 = 5
	ActivityNavPoint_PointTypeFlag                             = 6
	ActivityNavPoint_CapturePoint                              = 7
	ActivityNavPoint_DefensiveEncounter                        = 8
	ActivityNavPoint_GhostInteraction                          = 9
	ActivityNavPoint_KillAI                                    = 10
	ActivityNavPoint_QuestItem                                 = 11
	ActivityNavPoint_PatrolMission                             = 12
	ActivityNavPoint_Incoming                                  = 13
	ActivityNavPoint_ArenaObjective                            = 14
	ActivityNavPoint_AutomationHint                            = 15
	ActivityNavPoint_TrackedQuest                              = 16
)

// ActivityModeCategory describes the categories of a given activity mode.
type ActivityModeCategory int32

const (
	ActivityCategory_None           ActivityModeCategory = 0
	ActivityCategory_PvE                                 = 1
	ActivityCategory_PvP                                 = 2
	ActivityCategory_PvECompetitive                      = 3
)

// ItemSubType is a more specific categorization for a given item than ItemType.
type ItemSubType int32

const (
	SubType_None                  ItemSubType = 0
	SubType_Crucible                          = 1 // Deprecated. An item can both be "Crucible" and another subtype.
	SubType_Vanguard                          = 2 // Deprecated. An item can both be "Vanguard" and another subtype.
	SubType_Exotic                            = 5 // Deprecated. An item can both be Exotic and another subtype.
	SubType_AutoRifle                         = 6
	SubType_Shotgun                           = 7
	SubType_Machinegun                        = 8
	SubType_HandCannon                        = 9
	SubType_RocketLauncher                    = 10
	SubType_FusionRifle                       = 11
	SubType_SniperRifle                       = 12
	SubType_PulseRifle                        = 13
	SubType_ScoutRifle                        = 14
	SubType_CRM                               = 16 // Deprecated. An item can both be CRM and another subtype
	SubType_Sidearm                           = 17
	SubType_Sword                             = 18
	SubType_Mask                              = 19
	SubType_Shader                            = 20
	SubType_Ornament                          = 21
	SubType_FusionRifleLine                   = 22
	SubType_GrenadeLauncher                   = 23
	SubType_SubmachineGun                     = 24
	SubType_TraceRifle                        = 25
	SubType_HelmetArmor                       = 26
	SubType_GauntletsArmor                    = 27
	SubType_ChestArmor                        = 28
	SubType_LegArmor                          = 29
	SubType_ClassArmor                        = 30
	SubType_Bow                               = 31
	SubType_DummyRepeatableBounty             = 32
)

// GraphNodeState represents the potential state of an activity graph node.
type GraphNodeState int32

const (
	GraphNode_None       GraphNodeState = 0
	GraphNode_Visible                   = 1
	GraphNode_Teaser                    = 2
	GraphNode_Incomplete                = 3
	GraphNode_Completed                 = 4
)

// RewardSourceCategory describes the categories of different reward sources which spawn items.
type RewardSourceCategory int32

const (
	// None indicates that the reward source doesn't fit into the other categories.
	RewardSource_None RewardSourceCategory = 0
	// Rewards gained by playing an activity or set of activities, including quests and other in-game actions.
	RewardSource_Activity = 1
	// Rewards sold by vendors.
	RewardSource_Vendor = 2
	// Aggregate rewards can be gained in multiple ways.
	RewardSource_Aggregate = 3
)

// PresentationNodeType is the logical group that a given entity belongs to.
type PresentationNodeType int32

const (
	PresentationNode_Default      PresentationNodeType = 0
	PresentationNode_Category                          = 1
	PresentationNode_Collectibles                      = 2
	PresentationNode_Records                           = 3
	PresentationNode_Metric                            = 4
)

// Scope is the generalized scope of a given entity.
type Scope int32

const (
	Scope_Profile   Scope = 0
	Scope_Character       = 1
)

// PresentationDisplayStyle is a hint for how the presentation node should be displayed when shown in a list.
type PresentationDisplayStyle int32

const (
	PresentationDisplay_Category    PresentationDisplayStyle = 0
	PresentationDisplay_Badge                                = 1
	PresentationDisplay_Medals                               = 2
	PresentationDisplay_Collectible                          = 3
	PresentationDisplay_Record                               = 4
)

// RecordValueStyle is how to display a value in a given record.
type RecordValueStyle int32

const (
	RecordValue_Integer      RecordValueStyle = 0
	RecordValue_Percentage                    = 1
	RecordValue_Milliseconds                  = 2
	RecordValue_Boolean                       = 3
	RecordValue_Decimal                       = 4
)

// Gender is the gender of a Destiny 2 character.
type Gender int32

const (
	Gender_Male    Gender = 0
	Gender_Female         = 1
	Gender_Unknown        = 2
)

// GenderName is the gender of a Destiny 2 character as a string.
type GenderName string

const (
	GenderName_Male   GenderName = "Male"
	GenderName_Female GenderName = "Female"
)

// RecordToastStyle is a hint for how a record completion's toast should be shown in the UI.
type RecordToastStyle int32

const (
	RecordToast_None                    RecordToastStyle = 0
	RecordToast_Record                                   = 1
	RecordToast_Lore                                     = 2
	RecordToast_Badge                                    = 3
	RecordToast_MetaRecord                               = 4
	RecordToast_MedalComplete                            = 5
	RecordToast_SeasonChallengeComplete                  = 6
	RecordToast_GildedTitleComplete                      = 7
)

// PresentationScreenStyle is a hint for what screen should be shown when this presentation node is clicked into.
type PresentationScreenStyle int32

const (
	PresentationScreen_Default      PresentationScreenStyle = 0
	PresentationScreen_CategorySets                         = 1
	PresentationScreen_Badge                                = 2
)

// PlugUIStyles are specific custom styles for plugs.
type PlugUIStyles int32

const (
	PlugUI_None       PlugUIStyles = 0
	PlugUI_Masterwork              = 1
)

// PlugAvailabilityMode determines whether the plug is available to be inserted.
type PlugAvailabilityMode int32

const (
	// Normal means that all existing rules for plug insertion apply.
	PlugAvailability_Normal PlugAvailabilityMode = 0
	// UnavailableIfSocketContainsMatchingPlugCategory means that the plug is only
	// available if the socket does NOT match the plug category.
	PlugAvailability_UnavailableIfSocketContainsMatchingPlugCategory = 1
	// AvailableIfSocketContainsMatchingPlugCategory means that the plug is only
	// available if the socket DOES match the plug category.
	PlugAvailability_AvailableIfSocketContainsMatchingPlugCategory = 2
)

// EnergyType represents the socket energy types for Armor 2.0, Ghosts 2.0, and Stasis subclasses.
type EnergyType int32

const (
	EnergyType_Any      EnergyType = 0
	EnergyType_Arc                 = 1
	EnergyType_Thermal             = 2
	EnergyType_Void                = 3
	EnergyType_Ghost               = 4
	EnergyType_Subclass            = 5
	EnergyType_Stasis              = 6
)

// SocketPlugSources are indications of how a socket is populated, and where to look for valid plug data.
type SocketPlugSources int32

const (
	// None means there's no way to detect whether a new plug can be inserted.
	SocketPlug_None SocketPlugSources = 0
	// InventorySourced plugs are found in a player's inventory.
	SocketPlug_InventorySourced  = 1
	SocketPlug_ReusablePlugItems = 2
	SocketPlug_ProfilePlugSet    = 3
	SocketPlug_CharacterPlugSet  = 4
)

// ItemPerkVisibility determines how a perk should be shown in the game UI.
type ItemPerkVisibility int32

const (
	ItemPerk_Visible  ItemPerkVisibility = 0
	ItemPerk_Disabled                    = 1
	ItemPerk_Hidden                      = 2
)

// SpecialItemType is an enum retained from Destiny 1 for various internal logic.
// Prefer using ItemCategoryHashes to identify which categories an item belongs to.
type SpecialItemType int32

const (
	SpecialItem_None             SpecialItemType = 0
	SpecialItem_SpecialCurrency                  = 1
	SpecialItem_Armor                            = 8
	SpecialItem_Weapon                           = 9
	SpecialItem_Engram                           = 23
	SpecialItem_Consumable                       = 24
	SpecialItem_ExchangeMaterial                 = 25
	SpecialItem_MissionReward                    = 27
	SpecialItem_Currency                         = 29
)

// ItemType indicates the generalized category or type of a given item.
type ItemType int32

const (
	Item_None              ItemType = 0
	Item_Currency                   = 1
	Item_Armor                      = 2
	Item_Weapon                     = 3
	Item_Message                    = 7
	Item_Engram                     = 8
	Item_Consumable                 = 9
	Item_ExchangeMaterial           = 10
	Item_MissionReward              = 11
	Item_QuestStep                  = 12
	Item_QuestStepComplete          = 13
	Item_Emblem                     = 14
	Item_Quest                      = 15
	Item_Subclass                   = 16
	Item_ClanBanner                 = 17
	Item_Aura                       = 18
	Item_Mod                        = 19
	Item_Dummy                      = 20
	Item_Ship                       = 21
	Item_Vehicle                    = 22
	Item_Emote                      = 23
	Item_Ghost                      = 24
	Item_Package                    = 25
	Item_Bounty                     = 26
	Item_Wrapper                    = 27
	Item_SeasonalArtifact           = 28
	Item_Finisher                   = 29
)

// Class is the Destiny Class of an in-game character.
type Class int32

const (
	Class_Titan   Class = 0
	Class_Hunter        = 1
	Class_Warlock       = 2
	Class_Unknown       = 3
)

// BreakerTypeEnum is a special ability granted by activating a certain plug.
// Unfortunately the Bungie.net schema also has a contract named BreakerType
// so we use the awkward suffix Enum for the specific enum.
type BreakerTypeEnum int32

const (
	BreakerType_None           BreakerTypeEnum = 0
	BreakerType_ShieldPiercing                 = 1
	BreakerType_Disruption                     = 2
	BreakerType_Stagger                        = 3
)

// ProgressionRewardItemAcquisitionBehavior represents the different kinds of acquisition behavior for progression reward items.
type ProgressionRewardItemAcquisitionBehavior int32

const (
	ProgressionRewardAcquisition_Instant             ProgressionRewardItemAcquisitionBehavior = 0
	ProgressionRewardAcquisition_PlayerClaimRequired                                          = 1
)

// PresentationNodeState is the possible set of states that a presentation node can be in.
// This is meant to be used as a bitmask.
type PresentationNodeState int32

const (
	PresentationState_None      PresentationNodeState = 0
	PresentationState_Invisible                       = 1
	PresentationState_Obscured                        = 2
)

// Race is the race of a Destiny 2 character
type Race int32

const (
	Race_Human   Race = 0
	Race_Awoken       = 1
	Race_Exo          = 2
	Race_Unknown      = 3
)

// MilestoneType describes the basic types of Destiny 2 milestones.
type MilestoneType int32

const (
	Milestone_Unknown  MilestoneType = 0
	Milestone_Tutorial               = 1 // One-time milestones that are specifically oriented toward teaching players about new mechanics and gameplay modes.
	Milestone_OneTime                = 2 // Milestones that, once completed a single time, can never be repeated.
	Milestone_Weekly                 = 3 // Milestones that repeat/reset on a weekly basis.
	Milestone_Daily                  = 4 // Milestones that repeat or reset on a daily basis.
	Milestone_Special                = 5 // Special indicates that the event is not on a daily/weekly cadence, but does occur more than once.
)

// MilestoneDisplayPreference is a hint for the UI about what display information should be shown for a milestone.
type MilestoneDisplayPreference int32

const (
	// Display properties of MileStoneEntity should be shown.
	MilestoneDisplay_MilestoneDefinition MilestoneDisplayPreference = 0
	// Display properties for currently active quest steps should be shown.
	MilestoneDisplay_CurrentQuestSteps = 1
	// Display properties for activities should be shown.
	MilestoneDisplay_CurrentActivityChallenges = 2
)

// ActivityDifficultyTier describes the difficulty tiers for activities.
type ActivityDifficultyTier int32

const (
	ActivityDifficulty_Trivial          ActivityDifficultyTier = 0
	ActivityDifficulty_Easy                                    = 1
	ActivityDifficulty_Normal                                  = 2
	ActivityDifficulty_Challenging                             = 3
	ActivityDifficulty_Hard                                    = 4
	ActivityDifficulty_Brave                                   = 5
	ActivityDifficulty_AlmostImpossible                        = 6
	ActivityDifficulty_Impossible                              = 7
)
