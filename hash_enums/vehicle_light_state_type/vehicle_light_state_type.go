package vehicle_light_state_type

/*
   Create by zyx
   Date Time: 2024/9/23
   File: vehicle_light_state_type.go
*/

type VehicleLightState uint8

// note1: when using =2 on day it's lowbeam,highbeam, but at night it's lowbeam,lowbeam,highbeam
// note2: when using =0 it's affected by day or night for highbeams don't exist in daytime.
const (
	VehicleLightNormal         VehicleLightState = iota // Vehicle normal lights, off then lowbeams, then highbeams
	VehicleLightOff                              = 1    // Vehicle doesn't have lights, always off
	VehicleLightAlwaysOn                         = 2    // Vehicle has always on lights
	VehicleLightAlwaysOnLarger                   = 3    // or even larger
)
