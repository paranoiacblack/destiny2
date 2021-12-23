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

type ActivityModeCategory int32

const (
	ModeCategory_None ActivityModeCategory = iota
	ModeCategory_PvE
	ModeCategory_PvP
	ModeCategory_PvECompetitive
)

type UnlockValueUIStyle int32

const (
	UnlockValue_Automatic UnlockValueUIStyle = iota
	UnlockValue_Fraction
	UnlockValue_Checkbox
	UnlockValue_Percentage
	UnlockValue_DateTime
	UnlockValue_FractionFloat
	UnlockValue_Integer
	UnlockValue_TimeDuration
	UnlockValue_Hidden
	UnlockValue_Multiplier
	UnlockValue_GreenPips
	UnlockValue_RedPips
	UnlockValue_ExplicitPercentage
	UnlockValue_RawFloat
)

type ObjectiveGrantStyle int32

const (
	Objective_WhenIncomplete ObjectiveGrantStyle = iota
	Objective_WhenComplete
	Objective_Always
)

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

type GraphNodeState int32

const (
	GraphNode_None GraphNodeState = iota
	GraphNode_Visible
	GraphNode_Teaser
	GraphNode_Incomplete
	GraphNode_Completed
)

type RewardSourceCategory int32

const (
	RewardSource_None RewardSourceCategory = iota
	RewardSource_Activity
	RewardSource_Vendor
	RewardSource_Aggregate
)

type PresentationNodeType int32

const (
	PresentationNode_Default PresentationNodeType = iota
	PresentationNode_Category
	PresentationNode_Collectibles
	PresentationNode_Records
	PresentationNode_Metric
)

type PresentationNodeState int32

const (
	PresentationNode_None PresentationNodeState = iota
	PresentationNode_Invisible
	PresentationNode_Obscured
)

type Scope int32

const (
	Scope_Profile Scope = iota
	Scope_Character
)

type PresentationDisplayStyle int32

const (
	PresentationDisplay_Category PresentationDisplayStyle = iota
	PresentationDisplay_Badge
	PresentationDisplay_Medals
	PresentationDisplay_Collectible
	PresentationDisplay_Record
)

type RecordValueStyle int32

const (
	RecordValue_Integer RecordValueStyle = iota
	RecordValue_Percentage
	RecordValue_Milliseconds
	RecordValue_Boolean
	RecordValue_Decimal
)

type Gender int32

const (
	Gender_Male Gender = iota
	Gender_Female
	Gender_Unknown
)

type Race int32

const (
	Race_Human Race = iota
	Race_Awoken
	Race_Exo
	Race_Unknown
)

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

type PresentationScreenStyle int32

const (
	PresentationScreen_Default PresentationScreenStyle = iota
	PresentationScreen_CategorySets
	PresentatoinScreen_Badge
)

type PlugUIStyles int32

const (
	PlugUI_None PlugUIStyles = iota
	PlugUI_Masterwork
)

type PlugAvailabilityMode int32

const (
	PlugAvailability_Normal PlugAvailabilityMode = iota
	PlugAvailability_UnavailableIfSocketContainsMatchingPlugCategory
	PlugAvailability_AvailableIfSocketContainsMatchingPlugCategory
)

type SocketPlugSources int32

const (
	SocketPlug_None SocketPlugSources = iota
	SocketPlug_InventorySourced
	SocketPlug_ReusablePlugItems
	SocketPlug_ProfilePlugSet
	SocketPlug_CharacterPlugSet
)

type ItemPerkVisibility int32

const (
	ItemPerk_Visible ItemPerkVisibility = iota
	ItemPerk_Disabled
	ItemPerk_Hidden
)

type SpecialItemType int32

const (
	SpecialItem_None SpecialItemType = iota
	SpecialItem_SpecialCurrency
	SpecialItem_Armor            = 8
	SpecialItem_Weapon           = 9
	SpecialItem_Engram           = 23
	SpecialItem_Consumable       = 24
	SpecialItem_ExchangeMaterial = 25
	SpecialItem_MissionReward    = 27
	SpecialItem_Currency         = 29
)

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

type SubType int32

const (
	SubType_None SubType = iota
	// Deprecated
	SubType_Crucible
	// Deprecated
	SubType_Vanguard
	// Deprecated
	SubType_Exotic = 5
	SubType_AutoRifle
	SubType_Shotgun
	SubType_Machinegun
	SubType_HandCannon
	SubType_RocketLauncher
	SubType_FusionRifle
	SubType_SniperRifle
	SubType_PulseRifle
	SubType_ScoutRifle
	// Deprecated
	SubType_Crm = 16
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

type Class int32

const (
	Class_Titan Class = iota
	Class_Hunter
	Class_Warlock
	Class_Unknown
)

// BreakerTypeEnum ...
// Unfortunately the Bungie.net schema also has a contract named BreakerType
// so we use the awkward suffix Enum for the specific enum.
type BreakerTypeEnum int32

const (
	BreakerType_None BreakerTypeEnum = iota
	BreakerType_ShieldPiercing
	BreakerType_Disruption
	BreakerType_Stagger
)

type ProgressionRewardItemAcquisitionBehavior int32

const (
	ProgressionRewardAcquisition_Instant ProgressionRewardItemAcquisitionBehavior = iota
	ProgressionRewardAcquisition_PlayerClaimRequired
)

type MilestoneDisplayPreference int32

const (
	MilestoneDisplay_MilestoneDefinition MilestoneDisplayPreference = iota
	MilestoneDisplay_CurrentQuestSteps
	MilestoneDisplay_CurrentActivityChallenges
)

type MilestoneType int32

const (
	Milestone_Unknown MilestoneType = iota
	Milestone_Tutorial
	Milestone_OneTime
	Milestone_Weekly
	Milestone_Daily
	Milestone_Special
)

type StatsMergeMethod int32

const (
	MergeMethod_Add StatsMergeMethod = iota
	MergeMethod_Min
	MergeMethod_Max
)

type StatsGroupType int32

const (
	StatsGroup_None StatsGroupType = iota
	StatsGroup_General
	StatsGroup_Weapons
	StatsGroup_Medals
	StatsGroup_ReservedGroups = 100
	StatsGroup_Leaderboard
	StatsGroup_Activity
	StatsGroup_UniqueWeapon
	StatsGroup_Internal
)

type StatsCategoryType int32

const (
	StatsCategory_None StatsCategoryType = iota
	StatsCategory_Kills
	StatsCategory_Assists
	StatsCategory_Deaths
	StatsCategory_Criticals
	StatsCategory_KDA
	StatsCategory_KD
	StatsCategory_Score
	StatsCategory_Entered
	StatsCategory_TimePlayed
	StatsCategory_MedalWins
	StatsCategory_MedalGame
	StatsCategory_MedalSpecialKills
	StatsCategory_MedalSprees
	StatsCategory_MedalMultiKills
	StatsCategory_MedalAbilities
)

type UnitType int32

const (
	Unit_None UnitType = iota
	Unit_Count
	Unit_PerGame
	Unit_Seconds
	Unit_Points
	Unit_Team
	Unit_Distance
	Unit_Percent
	Unit_Ratio
	Unit_Boolean
	Unit_WeaponType
	Unit_Standing
	Unit_Milliseconds
	Unit_CompletionReason
)
