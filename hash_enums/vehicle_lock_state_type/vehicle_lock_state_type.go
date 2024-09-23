package vehicle_lock_state_type

/*
   Create by zyx
   Date Time: 2024/9/23
   File: vehicle_light_state_type.go
*/

type VehicleLockState uint8

const (
	VehicleLockNone                  VehicleLockState = iota // No specific lock state, vehicle behaves according to the game's default settings.
	VehicleLockUnlocked                               = 1    // Vehicle is fully unlocked, allowing free entry by players and NPCs.
	VehicleLockLocked                                 = 2    // Vehicle is locked, preventing entry by players and NPCs.
	VehicleLockLockoutPlayerOnly                      = 3    // Vehicle locks out only players, allowing NPCs to enter.
	VehicleLockLockedPlayerInside                     = 4    // Vehicle is locked once a player enters, preventing others from entering.
	VehicleLockLockedInitially                        = 5    // Vehicle starts in a locked state, but may be unlocked through game events.
	VehicleLockForceShutDoors                         = 6    // Forces the vehicle's doors to shut and lock.
	VehicleLockLockedButCanBeDamaged                  = 7    // Vehicle is locked but can still be damaged.
	VehicleLockLockedButBootUnlocked                  = 8    // Vehicle is locked, but its trunk/boot remains unlocked.
	VehicleLockLockedNoPassengers                     = 9    // Vehicle is locked and does not allow passengers, except for the driver.
	VehicleLockCannotEnter                            = 10   // Vehicle is completely locked, preventing entry entirely, even if previously inside.
)
