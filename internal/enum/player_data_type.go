package enum

/*
   Create by zyx
   Date Time: 2024/9/5
   File: player_data_type.go
*/

type PlayerDataType int32

const (
	PlayerHealth PlayerDataType = iota
	PlayerModel
	PlayerInvincible
	PlayerFrozen
	PlayerWeather
	PlayerCollision
	PlayerArmour
	PlayerDateTime
	PlayerPosition
	PlayerMaxHealth
	PlayerMaxArmour
	PlayerAmmo
	PlayerMaxAmmo
	PlayerCurrentWeapon
	PlayerWeaponAmmo
	PlayerDimension
	PlayerVisible
	PlayerRotation
	PlayerInVehicle
	PlayerDespawn
	PlayerClearBloodDamage
	PlayerEyeColor
	PlayerHairColor
	PlayerHairHighlightColor
	PlayerSpawn
	PlayerIsEnteringVehicle
	PlayerIsDead
	PlayerIsInVehicle
	PlayerIsAiming
	PlayerIsInCover
	PlayerIsInRagdoll
	PlayerIsShooting
	PlayerIsJumping
	PlayerIsLeavingVehicle
	PlayerIsInMelle
	PlayerVehicle
	PlayerClearProps
	PlayerClearDecorations
	PlayerClearTasks
	PlayerResetHeadBlendData
	PlayerHeadBlendData
	PlayerHeadOverlay
	PlayerHeadOverlayColor
	PlayerHeadBlendPaletteColor
)
