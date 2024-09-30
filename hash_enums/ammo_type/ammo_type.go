package ammo_type

/*
   Create by zyx
   Date Time: 2024/9/30
   File: ammo_type.go
*/

// Extend: coreclr-module - AmmoHash

//go:generate stringer -type=AmmoHash
type AmmoHash uint32

const (
	AmmoHashNull                 AmmoHash = 987444055
	AmmoHashShotgun              AmmoHash = 2416459067
	AmmoHashSniper               AmmoHash = 1285032059
	AmmoHashFireExtinguisher     AmmoHash = 1359393852
	AmmoHashGrenadeLauncher      AmmoHash = 1003267566
	AmmoHashSnowball             AmmoHash = 2182627693
	AmmoHashPistol               AmmoHash = 1950175060
	AmmoHashSmg                  AmmoHash = 1820140472
	AmmoHashFertilizerCan        AmmoHash = 1963932634
	AmmoHashBall                 AmmoHash = 4287981158
	AmmoHashMolotov              AmmoHash = 1446246869
	AmmoHashStickyBomb           AmmoHash = 1411692055
	AmmoHashPetrolCan            AmmoHash = 3395492001
	AmmoHashStunGun              AmmoHash = 2955849184
	AmmoHashRifle                AmmoHash = 218444191
	AmmoHashMinigun              AmmoHash = 2680539266
	AmmoHashMG                   AmmoHash = 1788949567
	AmmoHashFlareGun             AmmoHash = 1173416293
	AmmoHashFlare                AmmoHash = 1808594799
	AmmoHashGrenadeLauncherSmoke AmmoHash = 826266432
	AmmoHashHomingLauncher       AmmoHash = 2568293933
	AmmoHashRailgun              AmmoHash = 2034517757
	AmmoHashFirework             AmmoHash = 2938367503
	AmmoHashGrenade              AmmoHash = 1003688881
	AmmoHashBZGas                AmmoHash = 2608103076
	AmmoHashProximityMine        AmmoHash = 2938243239
	AmmoHashRayPistol            AmmoHash = 2768943988
	AmmoHashRpg                  AmmoHash = 1742569970
	AmmoHashPipeBomb             AmmoHash = 357983224
	AmmoHashHazardCan            AmmoHash = 1618528319
	AmmoHashEMPLauncher          AmmoHash = 4057942205
	AmmoHashAcidPackage          AmmoHash = 1003730930
	AmmoHashSmokeGrenade         AmmoHash = 3859679398
	AmmoHashRailgunXm3           AmmoHash = 1322889087
)
