package weapon_hash

/*
   Create by zyx
   Date Time: 2024/9/9
   File: enums.go
*/

// Extend: coreclr-module - PedModel

//go:generate stringer -type=ModelHash
type ModelHash uint32

const (
	AntiqueCavalryDagger            ModelHash = 2460120199
	BaseballBat                     ModelHash = 2508868239
	BrokenBottle                    ModelHash = 4192643659
	Crowbar                         ModelHash = 2227010557
	Fist                            ModelHash = 2725352035
	Flashlight                      ModelHash = 2343591895
	GolfClub                        ModelHash = 1141786504
	Hammer                          ModelHash = 1317494643
	Hatchet                         ModelHash = 4191993645
	BrassKnuckles                   ModelHash = 3638508604
	Knife                           ModelHash = 2578778090
	Machete                         ModelHash = 3713923289
	Switchblade                     ModelHash = 3756226112
	Nightstick                      ModelHash = 1737195953
	PipeWrench                      ModelHash = 419712736
	BattleAxe                       ModelHash = 3441901897
	PoolCue                         ModelHash = 2484171525
	StoneHatchet                    ModelHash = 940833800
	Pistol                          ModelHash = 453432689
	PistolMkII                      ModelHash = 3219281620
	CombatPistol                    ModelHash = 1593441988
	APPistol                        ModelHash = 584646201
	StunGun                         ModelHash = 911657153
	Pistol50                        ModelHash = 2578377531
	SNSPistol                       ModelHash = 3218215474
	SNSPistolMkII                   ModelHash = 2285322324
	HeavyPistol                     ModelHash = 3523564046
	VintagePistol                   ModelHash = 137902532
	FlareGun                        ModelHash = 1198879012
	MarksmanPistol                  ModelHash = 3696079510
	HeavyRevolver                   ModelHash = 3249783761
	HeavyRevolverMkII               ModelHash = 3415619887
	DoubleActionRevolver            ModelHash = 2548703416
	UpnAtomizer                     ModelHash = 2939590305
	MicroSMG                        ModelHash = 324215364
	SMG                             ModelHash = 736523883
	SMGMkII                         ModelHash = 2024373456
	AssaultSMG                      ModelHash = 4024951519
	CombatPDW                       ModelHash = 171789620
	MachinePistol                   ModelHash = 3675956304
	MiniSMG                         ModelHash = 3173288789
	UnholyHellbringer               ModelHash = 1198256469
	PumpShotgun                     ModelHash = 487013001
	PumpShotgunMkII                 ModelHash = 1432025498
	SawedOffShotgun                 ModelHash = 2017895192
	AssaultShotgun                  ModelHash = 3800352039
	BullpupShotgun                  ModelHash = 2640438543
	Musket                          ModelHash = 2828843422
	HeavyShotgun                    ModelHash = 984333226
	DoubleBarrelShotgun             ModelHash = 4019527611
	SweeperShotgun                  ModelHash = 317205821
	AssaultRifle                    ModelHash = 3220176749
	AssaultRifleMkII                ModelHash = 961495388
	CarbineRifle                    ModelHash = 2210333304
	CarbineRifleMkII                ModelHash = 4208062921
	AdvancedRifle                   ModelHash = 2937143193
	SpecialCarbine                  ModelHash = 3231910285
	SpecialCarbineMkII              ModelHash = 2526821735
	BullpupRifle                    ModelHash = 2132975508
	BullpupRifleMkII                ModelHash = 2228681469
	CompactRifle                    ModelHash = 1649403952
	MG                              ModelHash = 2634544996
	CombatMG                        ModelHash = 2144741730
	CombatMGMkII                    ModelHash = 3686625920
	GusenbergSweeper                ModelHash = 1627465347
	SniperRifle                     ModelHash = 100416529
	HeavySniper                     ModelHash = 205991906
	HeavySniperMkII                 ModelHash = 177293209
	MarksmanRifle                   ModelHash = 3342088282
	MarksmanRifleMkII               ModelHash = 1785463520
	RPG                             ModelHash = 2982836145
	GrenadeLauncher                 ModelHash = 2726580491
	GrenadeLauncherSmoke            ModelHash = 1305664598
	Minigun                         ModelHash = 1119849093
	FireworkLauncher                ModelHash = 2138347493
	Railgun                         ModelHash = 1834241177
	HomingLauncher                  ModelHash = 1672152130
	CompactGrenadeLauncherModelHash           = 125959754
	Widowmaker                      ModelHash = 3056410471
	Grenade                         ModelHash = 2481070269
	BZGas                           ModelHash = 2694266206
	MolotovCocktail                 ModelHash = 615608432
	StickyBomb                      ModelHash = 741814745
	ProximityMines                  ModelHash = 2874559379
	Snowballs                       ModelHash = 126349499
	PipeBombs                       ModelHash = 3125143736
	Baseball                        ModelHash = 600439132
	TearGas                         ModelHash = 4256991824
	Flare                           ModelHash = 1233104067
	JerryCan                        ModelHash = 883325847
	Parachute                       ModelHash = 4222310262
	FireExtinguisher                ModelHash = 101631238
	GadgetPistol                    ModelHash = 1470379660
	MilitaryRifle                   ModelHash = 2636060646
	CombatShotgun                   ModelHash = 94989220
	Fertilizercan                   ModelHash = 406929569
	HeavyRifle                      ModelHash = 3347935668
	EMPLauncher                     ModelHash = 3676729658
	CeramicPistol                   ModelHash = 727643628
	NavyRevolver                    ModelHash = 2441047180
	HazardCan                       ModelHash = 3126027122
	TacticalRifle                   ModelHash = 3520460075
	MetalDetector                   ModelHash = 3684886537
	PrecisionRifle                  ModelHash = 1853742572
	CandyCane                       ModelHash = 1703483498
	AcidPackage                     ModelHash = 4159824478
	Wm29Pistol                      ModelHash = 465894841
	RailgunXm3                      ModelHash = 4272043364
	TecPistol                       ModelHash = 350597077  // Tactical SMG
	BattleRifle                     ModelHash = 1924557585 // [0x72B66B11] Battle Rifle
	SnowLauncher                    ModelHash = 62870901   // [0x3BF5575] Snowball Launcher
	HackingDevice                   ModelHash = 485882440  // [0x1CF5FA48] Hacking Device
	Stunrod                         ModelHash = 3670016037 // [0xDAC00025] The Shocker
)
