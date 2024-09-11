package enum

/*
   Create by zyx
   Date Time: 2024/9/5
   File: player_data_type.go
*/

type PlayerDataType int32

const (
	Health PlayerDataType = 0
	Model
	Invincible
	Frozen
	Weather
	Collision
	Armour
	DateTime
	Position
	MaxHealth
	MaxArmour
	Ammo
	MaxAmmo
	CurrentWeapon
	WeaponAmmo
	Dimension
	Visible
	Rot
	InVehicle
)
