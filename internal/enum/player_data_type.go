package enum

/*
   Create by zyx
   Date Time: 2024/9/5
   File: player_data_type.go
*/

type PlayerDataType int32

const (
	Health PlayerDataType = iota
	Model
	Invincible
	Frozen
	Weather
	Collision
	Armour
	DateTime
	Positon
	MaxHealth
	MaxArmour
	Ammo
	MaxAmmo
	CurrentWeapon
	WeaponAmmo
	Dimension
	Rot
	InVehicle
)
