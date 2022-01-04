package destiny2

// ActivityModeType is the mode of a given activity.
type ActivityModeType int32

const (
	ActivityMode_None  ActivityModeType = iota
	ActivityMode_Story                  = 2
	ActivityMode_Strike
	ActivityMode_Raid
	ActivityMode_AllPvP
	ActivityMode_Patrol
	ActivityMode_AllPvE
	ActivityMode_Control        = 10
	ActivityMode_Clash          = 12
	ActivityMode_CrimsonDoubles = 15
	ActivityMode_Nightfall
	ActivityMode_HeroicNightfall
	ActivityMode_AllStrikes
	ActivityMode_IronBanner
	ActivityMode_AllMayhem = 25
	ActivityMode_Supremacy = 31
	ActivityMode_PrivateMatchesAll
	ActivityMode_Survival = 37
	ActivityMode_Countdown
	ActivityMode_TrialsOfTheNine
	ActivityMode_Social
	ActivityMode_TrialsCountdown
	ActivityMode_TrialsSurvival
	ActivityMode_IronBannerControl
	ActivityMode_IronBannerClash
	ActivityMode_IronBannerSupremacy
	ActivityMode_ScoredNightfall
	ActivityMode_ScoredHeroicNightfall
	ActivityMode_Rumble
	ActivityMode_AllDoubles
	ActivityMode_Doubles
	ActivityMode_PrivateMatchesClash
	ActivityMode_PrivateMatchesControl
	ActivityMode_PrivateMatchesSupremacy
	ActivityMode_PrivateMatchesCountdown
	ActivityMode_PrivateMatchesSurvival
	ActivityMode_PrivateMatchesMayhem
	ActivityMode_PrivateMatchesRumble
	ActivityMode_HeroicAdventure
	ActivityMode_Showdown
	ActivityMode_Lockdown
	ActivityMode_Scorched
	ActivityMode_ScorchedTeam
	ActivityMode_AllPvEEcompetitive
	ActivityMode_Breakthrough
	ActivityMode_BlackArmoryRun
	ActivityMode_Salvage
	ActivityMode_IronBannerSalvage
	ActivityMode_PvPCompetitive
	ActivityMode_PvPQuickplay
	ActivityMode_ClashQuickplay
	ActivityMode_ClashCompetitive
	ActivityMode_ControlQuickplay
	ActivityMode_ControlCompetitive
	ActivityMode_GambitPrime
	ActivityMode_Reckoning
	ActivityMode_Menagerie
	ActivityMode_Elimination
	ActivityMode_Momentum
	ActivityMode_Dungeon
	ActivityMode_Sundial
	ActivityMode_TrialsOfOsiris
	ActivityMode_Dares
)

// ActivityDifficultyTier describes the difficulty tiers for activities.
type ActivityDifficultyTier int32

const (
	ActivityDifficulty_Trivial ActivityDifficultyTier = iota
	AcitivtyDifficulty_Easy
	AcitivtyDifficulty_Normal
	AcitivtyDifficulty_Challenging
	AcitivtyDifficulty_Hard
	AcitivtyDifficulty_Brave
	AcitivtyDifficulty_AlmostImpossible
	AcitivtyDifficulty_Impossible
)

// UnlockValueUIStyle describes the format unlock values should use in the UI.
type UnlockValueUIStyle int32

const (
	// Automatic means just show the number.
	UnlockValue_Automatic          UnlockValueUIStyle = iota
	UnlockValue_Fraction                              // Show the number as a fractional value; this is relative to the context-specific upper bound.
	UnlockValue_Checkbox                              // Show the number as a checkbox.
	UnlockValue_Percentage                            // Show the number as a percentage; this is relative to the context-specific upper bound.
	UnlockValue_DateTime                              // Show the number as a date and time; the given number number of seconds since the Unix Epoch.
	UnlockValue_FractionFloat                         // Show the number as a floating point value that represents a fraction; this is relative to the context-specific upper bound.
	UnlockValue_Integer                               // Show the number as an integer.
	UnlockValue_TimeDuration                          // Show the number as a time duration in seconds.
	UnlockValue_Hidden                                // Don't show the value.
	UnlockValue_Multiplier                            // Show the value as a multiplier in context, such as "1.5x".
	UnlockValue_GreenPips                             // Show the value as a series of green pips, like wins on a Trials Card.
	UnlockValue_RedPips                               // Show the value as a series of red pips, like losses on a Trials Card.
	UnlockValue_ExplicitPercentage                    // Show the value as a percentage, by appending a %-sign.
	UnlockValue_RawFloat                              // Show the value as a floating-point number with two-decimal precision.
)

// ObjectiveGrantStyle indicates when a given objected-related perk is granted.
type ObjectiveGrantStyle int32

const (
	Objective_WhenIncomplete ObjectiveGrantStyle = iota
	Objective_WhenComplete
	Objective_Always
)

// DamageType describes the general categories of damage in the game.
type DamageType int32

const (
	DamageType_None DamageType = iota
	DamageType_Kinetic
	DamageType_Arc
	DamageType_Thermal
	DamageType_Void
	DamageType_Raid
	DamageType_Stasis
)

// EnergyType represents the socket energy types for Armor 2.0, Ghosts 2.0, and Stasis subclasses.
type EnergyType int32

const (
	EnergyType_Any EnergyType = iota
	EnergyType_Arc
	EnergyType_Thermal
	EnergyType_Void
	EnergyType_Ghost
	EnergyType_Subclass
	EnergyType_Stasis
)

// ActivityNavPointType describes the interaction associated with a navigation point in an activity.
type ActivityNavPointType int32

const (
	ActivityNavPoint_Inactive ActivityNavPointType = iota
	ActivityNavPoint_PrimaryObjective
	ActivityNavPoint_SecondaryObjective
	ActivityNavPoint_TravelObjective
	ActivityNavPoint_PublicEventObjective
	ActivityNavPoint_AmmoCache
	ActivityNavPoint_PointTypeFlag
	ActivityNavPoint_CapturePoint
	ActivityNavPoint_DefensiveEncounter
	ActivityNavPoint_GhostInteraction
	ActivityNavPoint_KillAi
	ActivityNavPoint_QuestItem
	ActivityNavPoint_PatrolMission
	ActivityNavPoint_Incoming
	ActivityNavPoint_ArenaObjective
	ActivityNavPoint_AutomationHint
	ActivityNavPoint_TrackedQuest
)

// GraphNodeState represents the potential state of an activity graph node.
type GraphNodeState int32

const (
	GraphNode_None GraphNodeState = iota
	GraphNode_Visible
	GraphNode_Teaser
	GraphNode_Incomplete
	GraphNode_Completed
)

// ActivityModeCategory describes the categories of a given activity mode.
type ActivityModeCategory int32

const (
	ActivityCategory_None ActivityModeCategory = iota
	ActivityCategory_PvE
	ActivityCategory_PvP
	ActivityCategory_PvECompetitive
)

// BucketCategory describes what categories of items can be found in an inventory bucket.
type BucketCategory int32

const (
	BucketCategory_Invisible BucketCategory = iota
	BucketCategory_Item
	BucketCategory_Currency
	BucketCategory_Equippable
	BucketCategory_Ignored
)

// RewardSourceCategory describes the categories of different reward sources which spawn items.
type RewardSourceCategory int32

const (
	// None indicates that the reward source doesn't fit into the other categories.
	RewardSource_None RewardSourceCategory = iota
	// Rewards gained by playing an activity or set of activities, including quests and other in-game actions.
	RewardSource_Activity
	// Rewards sold by vendors.
	RewardSource_Vendor
	// Aggregate rewards can be gained in multiple ways.
	RewardSource_Aggregate
)

// StatCategory describes the categories which a given stat affects.
type StatCategory int32

const (
	StatCategory_Gameplay StatCategory = iota
	StatCategory_Weapon
	StatCategory_Defense
	StatCategory_Primary
)

// PresentationNodeType is the logical group that a given entity belongs to.
type PresentationNodeType int32

const (
	PresentationNode_Default PresentationNodeType = iota
	PresentationNode_Category
	PresentationNode_Collectibles
	PresentationNode_Records
	PresentationNode_Metric
)

// PresentationNodeState is the possible set of states that a presentation node can be in.
// This is meant to be used as a bitmask.
type PresentationNodeState int32

const (
	PresentationState_None PresentationNodeState = iota
	PresentationState_Invisible
	PresentationState_Obscured
)

// PresentationDisplayStyle is a hint for how the presentation node should be displayed when shown in a list.
type PresentationDisplayStyle int32

const (
	PresentationDisplay_Category PresentationDisplayStyle = iota
	PresentationDisplay_Badge
	PresentationDisplay_Medals
	PresentationDisplay_Collectible
	PresentationDisplay_Record
)

// PresentationScreenStyle is a hint for what screen should be shown when this presentation node is clicked into.
type PresentationScreenStyle int32

const (
	PresentationScreen_Default PresentationScreenStyle = iota
	PresentationScreen_CategorySets
	PresentationScreen_Badge
)

// Scope is the generalized scope of a given entity.
type Scope int32

const (
	Scope_Profile Scope = iota
	Scope_Character
)

// BucketScope is scope of an inventory bucket.
type BucketScope int32

const (
	BucketScope_Profile BucketScope = iota
	BucketScope_Character
)

// ProgressionScope is the scope of a progression which determines how it is stored, calculated and used.
type ProgressionScope int32

const (
	Progression_Account ProgressionScope = iota
	Progression_Character
	Progression_Clan
	Progression_Item
	Progression_ImplicitFromEquipment
	Progression_Mapped
	Progression_MappedAggregate
	Progression_MappedStat
	Progression_MappedUnlockValue
)

// GatingScope represents the most restrictive type of gating that is being performed by an entity.
type GatingScope int32

const (
	// No gating on this item.
	Gating_None GatingScope = iota
	// Gating on this item is based on global game state. It will be gated the same for everyone.
	Gating_Global
	// Gating on this item is at the Clan level.
	Gating_Clan
	// Gating includes Profile-specific checks, but not on the Profile's characters.
	Gating_Profile
	// Gating includes Character-specific checks, including character level restrictions.
	Gating_Character
	// Gating includes item-specific checks.
	Gating_Item
	// Gating unlocks and checks are unknown.
	Gating_AssumedWorstCase
)

// RecordValueStyle is how to display a value in a given record.
type RecordValueStyle int32

const (
	RecordValue_Integer RecordValueStyle = iota
	RecordValue_Percentage
	RecordValue_Milliseconds
	RecordValue_Boolean
	RecordValue_Decimal
)

// RecordToastStyle is a hint for how a record completion's toast should be shown in the UI.
type RecordToastStyle int32

const (
	RecordToast_None RecordToastStyle = iota
	RecordToast_Record
	RecordToast_Lore
	RecordToast_Badge
	RecordToast_MetaRecord
	RecordToast_MedalComplete
	RecordToast_SeasonChallengeComplete
	RecordToast_GildedTitleComplete
)

// Gender is the gender of a Destiny 2 character.
type Gender int32

const (
	Gender_Male Gender = iota
	Gender_Female
	Gender_Unknown
)

// GenderName is the gender of a Destiny 2 character as a string.
type GenderName string

const (
	GenderName_Male   GenderName = "Male"
	GenderName_Female GenderName = "Female"
)

// Race is the race of a Destiny 2 character
type Race int32

const (
	Race_Human Race = iota
	Race_Awoken
	Race_Exo
	Race_Unknown
)

// PlugUIStyles are specific custom styles for plugs.
type PlugUIStyles int32

const (
	PlugUI_None PlugUIStyles = iota
	PlugUI_Masterwork
)

// PlugAvailabilityMode determines whether the plug is available to be inserted.
type PlugAvailabilityMode int32

const (
	// Normal means that all existing rules for plug insertion apply.
	PlugAvailability_Normal PlugAvailabilityMode = iota
	// UnavailableIfSocketContainsMatchingPlugCategory means that the plug is only
	// available if the socket does NOT match the plug category.
	PlugAvailability_UnavailableIfSocketContainsMatchingPlugCategory
	// AvailableIfSocketContainsMatchingPlugCategory means that the plug is only
	// available if the socket DOES match the plug category.
	PlugAvailability_AvailableIfSocketContainsMatchingPlugCategory
)

// SocketPlugSources are indications of how a socket is populated, and where to look for valid plug data.
type SocketPlugSources int32

const (
	// None means there's no way to detect whether a new plug can be inserted.
	SocketPlug_None SocketPlugSources = iota
	// InventorySourced plugs are found in a player's inventory.
	SocketPlug_InventorySourced
	SocketPlug_ReusablePlugItems
	SocketPlug_ProfilePlugSet
	SocketPlug_CharacterPlugSet
)

// ItemType indicates the generalized category or type of a given item.
type ItemType int32

const (
	Item_None ItemType = iota
	Item_Currency
	Item_Armor
	Item_Weapon
	Item_Message = 7
	Item_Engram
	Item_Consumable
	Item_ExchangeMaterial
	Item_MissionReward
	Item_QuestStep
	Item_QuestStepComplete
	Item_Emblem
	Item_Quest
	Item_Subclass
	Item_ClanBanner
	Item_Aura
	Item_Mod
	Item_Dummy
	Item_Ship
	Item_Vehicle
	Item_Emote
	Item_Ghost
	Item_Package
	Item_Bounty
	Item_Wrapper
	Item_SeasonalArtifact
	Item_Finisher
)

// ItemSubType is a more specific categorization for a given item than ItemType.
type ItemSubType int32

const (
	SubType_None     ItemSubType = iota
	SubType_Crucible             // Deprecated. An item can both be "Crucible" and another subtype.
	SubType_Vanguard             // Deprecated. An item can both be "Vanguard" and another subtype.
	SubType_Exotic   = 5         // Deprecated. An item can both be Exotic and another subtype.
	SubType_AutoRifle
	SubType_Shotgun
	SubType_Machinegun
	SubType_HandCannon
	SubType_RocketLauncher
	SubType_FusionRifle
	SubType_SniperRifle
	SubType_PulseRifle
	SubType_ScoutRifle
	SubType_CRM = 16 // Deprecated. An item can both be CRM and another subtype
	SubType_Sidearm
	SubType_Sword
	SubType_Mask
	SubType_Shader
	SubType_Ornament
	SubType_FusionRifleLine
	SubType_GrenadeLauncher
	SubType_SubmachineGun
	SubType_TraceRifle
	SubType_HelmetArmor
	SubType_GauntletsArmor
	SubType_ChestArmor
	SubType_LegArmor
	SubType_ClassArmor
	SubType_Bow
	SubType_DummyRepeatableBounty
)

// SpecialItemType is an enum retained from Destiny 1 for various internal logic.
// Prefer using ItemCategoryHashes to identify which categories an item belongs to.
type SpecialItemType int32

const (
	SpecialItem_None SpecialItemType = iota
	SpecialItem_SpecialCurrency
	SpecialItem_Armor = 8
	SpecialItem_Weapon
	SpecialItem_Engram = 23
	SpecialItem_Consumable
	SpecialItem_ExchangeMaterial
	SpecialItem_MissionReward
	SpecialItem_Currency = 29
)

// ItemTier is the tier of an item in an inventory.
type ItemTier int32

const (
	ItemTier_Unknown ItemTier = iota
	ItemTier_Currency
	ItemTier_Basic
	ItemTier_Common
	ItemTier_Rare
	ItemTier_Superior
	ItemTier_Exotic
)

// ItemPerkVisibility determines how a perk should be shown in the game UI.
type ItemPerkVisibility int32

const (
	ItemPerk_Visible ItemPerkVisibility = iota
	ItemPerk_Disabled
	ItemPerk_Hidden
)

// ItemLocation describes where a given item can be found.
type ItemLocation int32

const (
	ItemLocation_Unknown ItemLocation = iota
	ItemLocation_Inventory
	ItemLocation_Vault
	ItemLocation_Vendor
	ItemLocation_Postmaster
)

// Class is the Destiny Class of an in-game character.
type Class int32

const (
	Class_Titan Class = iota
	Class_Hunter
	Class_Warlock
	Class_Unknown
)

// BreakerTypeEnum is a special ability granted by activating a certain plug.
// Unfortunately the Bungie.net schema also has a contract named BreakerType
// so we use the awkward suffix Enum for the specific enum.
type BreakerTypeEnum int32

const (
	BreakerType_None BreakerTypeEnum = iota
	BreakerType_ShieldPiercing
	BreakerType_Disruption
	BreakerType_Stagger
)

// ProgressionRewardItemAcquisitionBehavior represents the different kinds of acquisition behavior for progression reward items.
type ProgressionRewardItemAcquisitionBehavior int32

const (
	ProgressionRewardAcquisition_Instant ProgressionRewardItemAcquisitionBehavior = iota
	ProgressionRewardAcquisition_PlayerClaimRequired
)

// MilestoneType describes the basic types of Destiny 2 milestones.
type MilestoneType int32

const (
	Milestone_Unknown  MilestoneType = iota
	Milestone_Tutorial               // One-time milestones that are specifically oriented toward teaching players about new mechanics and gameplay modes.
	Milestone_OneTime                // Milestones that, once completed a single time, can never be repeated.
	Milestone_Weekly                 // Milestones that repeat/reset on a weekly basis.
	Milestone_Daily                  // Milestones that repeat or reset on a daily basis.
	Milestone_Special                // Special indicates that the event is not on a daily/weekly cadence, but does occur more than once.
)

// MilestoneDisplayPreference is a hint for the UI about what display information should be shown for a milestone.
type MilestoneDisplayPreference int32

const (
	// Display properties of MileStoneEntity should be shown.
	MilestoneDisplay_MilestoneDefinition MilestoneDisplayPreference = iota
	// Display properties for currently active quest steps should be shown.
	MilestoneDisplay_CurrentQuestSteps
	// Display properties for activities should be shown.
	MilestoneDisplay_CurrentActivityChallenges
)

// StatAggregationType describes the rule used for determining the level
// and the formula used for aggregation of stats.
type StatAggregationType int32

const (
	// CharacterAverage applies a weighted average using stat grouping and items equipped on a character.
	StatAggregation_CharacterAverage StatAggregationType = iota
	StatAggregation_Character                            // Use only the character's stat grouping.
	StatAggregation_Item                                 // Use only the character's equipped items.

)

// ProgressionStepDisplayEffect determines whether progression shows visual effects on the character, its item or neither.
type ProgressionStepDisplayEffect int32

const (
	ProgressionStepDisplay_None ProgressionStepDisplayEffect = iota
	ProgressionStepDisplay_Character
	ProgressionStepDisplay_Item
)

// EquippingItemBlockAttributes are custom attributes on the equippability of the item.
type EquippingItemBlockAttributes int32

const (
	EquippingAttribute_None EquippingItemBlockAttributes = iota
	EquippingAttribute_EquipOnAcquire
)

// AmmunitionType is the ammunition a weapon used.
type AmmunitionType int32

const (
	Ammunition_None AmmunitionType = iota
	Ammunition_Primary
	Ammunition_Special
	Ammunition_Heavy
	Ammunition_Unknown
)

// BungieMembershipType describes the types of memberships the accounts system supports.
type BungieMembershipType int32

const (
	Membership_None BungieMembershipType = iota
	TigerXbox
	TigerPSN
	TigerSteam
	TigerBlizzard
	TigerStadia
	TigerDemon     = 10
	BungieNext     = 254
	Membership_All = -1 // Membership_All is only valid for searching capabilities.
)

// VendorInteractionType is an enumeration of the known UI interactions for Vendors.
type VendorInteractionType int32

const (
	VendorInteraction_Unknown VendorInteractionType = iota
	// An empty interaction. If this ends up in content, it is probably a game bug.
	VendorInteraction_Undefined
	// An interaction shown when you complete a quest and receive a reward.
	VendorInteraction_QuestComplete
	// An interaction shown when you talk to a Vendor as an intermediary step of a quest.
	VendorInteraction_QuestContinue
	// An interaction shown when you are previewing the vendor's reputation rewards.
	VendorInteraction_ReputationPreview
	// An interaction shown when you rank up with the vendor.
	VendorInteraction_RankUpReward
	// An interaction shown when you have tokens to turn in for the vendor.
	VendorInteraction_TokenTurnin
	// An interaction shown when you're accepting a new quest.
	VendorInteraction_QuestAccept
	VendorInteraction_ProgressTab
	VendorInteraction_End
	VendorInteraction_Start
)

// VendorInteractionRewardSelection determines how many rewards are provided on selection.
type VendorInteractionRewardSelection int32

const (
	VendorSelection_None VendorInteractionRewardSelection = iota
	VendorSelection_One
	VendorSelection_All
)

// VendorItemRefundPolicy is the action that happens when the user attempts to refund an item.
type VendorItemRefundPolicy int32

const (
	RefundPolicy_NotRefundable VendorItemRefundPolicy = iota
	RefundPolicy_DeletesItem
	RefundPolicy_RevokesLicense
)

// VendorProgressionType describes the type of progression a vendor has.
type VendorProgressionType int32

const (
	// Default is vendor rank progression from token redemption.
	VendorProgression_Default VendorProgressionType = iota
	// Rank progressions from ritual/playlist content such as Crucible, Gambit and Battlegrounds.
	VendorProgression_Ritual
	// Rank progressions that cannot be reset in a given seasons, such as Xur in the Eternity destination.
	VendorProgression_NoSeasonalRefresh
)

// VendorReplyType determines the type of reply that a vendor will have during an interaction.
type VendorReplyType int32

const (
	VendorReply_Accept VendorReplyType = iota
	VendorReply_Decline
	VendorReply_Complete
)

// SocketTypeActionType indicates the types of actions that can be performed.
type SocketTypeActionType int32

const (
	InsertPlug SocketTypeActionType = iota
	InfuseItem
	ReinitializeSocket
)

// SocketTypeVisibility indicates whether a socket is visible.
type SocketTypeVisibility int32

const (
	SocketVisibility_Visible SocketTypeVisibility = iota
	SocketVisibility_Hidden
	SocketVisibility_HiddenWhenEmpty
	SocketVisibility_HiddenIfNoPlugsAvailable
)

// SocketCategoryStyle represents the possible and known UI styles used for rendering socket categories.
type SocketCategoryStyle int32

const (
	SocketCategory_Unknown SocketCategoryStyle = iota
	SocketCategory_Reusable
	SocketCategory_Consumable
	SocketCategory_Unlockable
	SocketCategory_Intrinsic
	SocketCategory_EnergyMeter
	SocketCategory_LargePerk
	SocketCategory_Abilities
	SocketCategory_Supers
)

// ActivityGraphNodeHighlightType is the known UI style in which an item can be highlighted.
type ActivityGraphNodeHighlightType int32

const (
	ActivityHighlight_None ActivityGraphNodeHighlightType = iota
	ActivityHighlight_Normal
	ActivityHighlight_Hyper

	// Special highlight states, may no longer be used.
	ActivityHighlight_Comet
	ActivityHighlight_RiseOfIron
)

// TalentNodeStepWeaponPerformances are basically the stats on the weapon granted by a weapon talent.
// A given talent can apply multiple performance changes to a weapon so use this as a bitmask.
type TalentNodeStepWeaponPerformances int32

const (
	TalentWeaponPerformance_RateOfFire TalentNodeStepWeaponPerformances = 1 << iota
	TalentWeaponPerformance_Damage
	TalentWeaponPerformance_Accuracy4
	TalentWeaponPerformance_Range
	TalentWeaponPerformance_Zoom
	TalentWeaponPerformance_Recoil
	TalentWeaponPerformance_Ready
	TalentWeaponPerformance_Reload
	TalentWeaponPerformance_HairTrigger
	TalentWeaponPerformance_AmmoAndMagazine
	TalentWeaponPerformance_TrackingAndDetonation
	TalentWeaponPerformance_ShotgunSpread
	TalentWeaponPerformance_ChargeTime
	TalentWeaponPerformance_None = 0
	TalentWeaponPerformance_All  = 8191
)

// TalentNodeStepImpactEffects are the effects on bullet impact granted by a talent.
// A given talent can apply multiple effects so use this as a bitmask.
type TalentNodeStepImpactEffects int32

const (
	TalentImpactEffect_ArmorPiercing TalentNodeStepImpactEffects = 1 << iota
	TalentImpactEffect_Ricochet
	TalentImpactEffect_Flinch
	TalentImpactEffect_CollateralDamage
	TalentImpactEffect_Disorient
	TalentImpactEffect_HighlightTarget
	TalentImpactEffect_None = 0
	TalentImpactEffect_All  = 63
)

// TalentNodeStepGuardianAttributes are the guardian attributes granted by a talent.
// A given talent can apply multiple guardian attributes so use this as a bitmask.
type TalentNodeStepGuardianAttributes int32

const (
	TalentNodeStepGuardianAttribute_Stats TalentNodeStepGuardianAttributes = 1 << iota
	TalentNodeStepGuardianAttribute_Shields
	TalentNodeStepGuardianAttribute_Health
	TalentNodeStepGuardianAttribute_Revive
	TalentNodeStepGuardianAttribute_AimUnderFire
	TalentNodeStepGuardianAttribute_Radar
	TalentNodeStepGuardianAttribute_Invisibility
	TalentNodeStepGuardianAttribute_Reputations
	TalentNodeStepGuardianAttribute_None = 0
	TalentNodeStepGuardianAttribute_All  = 255
)

// TalentNodeStepLightAbilities are the light-related abilities granted by a talent.
// A given talent can apply multiple light abilities so use this as a bitmask.
type TalentNodeStepLightAbilities int32

const (
	TalentLightAbility_Grenades TalentNodeStepLightAbilities = 1 << iota
	TalentLightAbility_Melee
	TalentLightAbility_MovementModes
	TalentLightAbility_Orbs
	TalentLightAbility_SuperEnergy
	TalentLightAbility_SuperMods
	TalentLightAbility_None = 0
	TalentLightAbility_All  = 63
)

// TalentNodeStepDamageTypes are the damage types this talent apply to.
// A given talent can be applied on multiple damage types so use this as a bitmask.
type TalentNodeStepDamageTypes int32

const (
	TalentNodeStepDamageType_Kinetic TalentNodeStepDamageTypes = 1 << iota
	TalentNodeStepDamageType_Arc
	TalentNodeStepDamageType_Solar
	TalentNodeStepDamageType_Void
	TalentNodeStepDamageType_None = 0
	TalentNodeStepDamageType_All  = 15
)
