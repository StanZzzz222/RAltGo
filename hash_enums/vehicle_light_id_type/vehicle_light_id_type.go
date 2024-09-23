package vehicle_light_id_type

/*
   Create by zyx
   Date Time: 2024/9/23
   File: vehicle_light_id_type.go
*/

type VehicleLightType uint8

const (
	LeftHeadLight       VehicleLightType = iota + 1 // Left headlight
	RightHeadLight                                  // Right headlight (headlight_r)
	LeftTailLight                                   // Left taillight (taillight_l)
	RightTailLight                                  // Right taillight (taillight_r)
	FrontLeftIndicator                              // Front left indicator (indicator_lf)
	FrontRightIndicator                             // Front right indicator (indicator_rf)
	RearLeftIndicator                               // Rear left indicator (indicator_lr)
	RearRightIndicator                              // Rear right indicator (indicator_rr)
	LeftBrakeLight                                  // Left brakelight (brakelight_l)
	RightBrakeLight                                 // Right brakelight (brakelight_r)
	MiddleBrakeLight                                // Middle brakelight (brakelight_m)
	LeftReverseLight                                // Left reverse light (reversinglight_l)
	RightReverseLight                               // Right reverse light (reversinglight_r)
	ExtraLight1                                     // Extra light 1 (extralight_1)
	ExtraLight2                                     // Extra light 2 (extralight_2)
)
