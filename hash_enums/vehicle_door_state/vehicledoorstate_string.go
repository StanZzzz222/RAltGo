// Code generated by "stringer -type=VehicleDoorState"; DO NOT EDIT.

package vehicle_door_state

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Closed-0]
	_ = x[OpenedLevel1-1]
	_ = x[OpenedLevel2-2]
	_ = x[OpenedLevel3-3]
	_ = x[OpenedLevel4-4]
	_ = x[OpenedLevel5-5]
	_ = x[OpenedLevel6-6]
	_ = x[OpenedLevel7-7]
	_ = x[DoesNotExists-255]
}

const (
	_VehicleDoorState_name_0 = "ClosedOpenedLevel1OpenedLevel2OpenedLevel3OpenedLevel4OpenedLevel5OpenedLevel6OpenedLevel7"
	_VehicleDoorState_name_1 = "DoesNotExists"
)

var (
	_VehicleDoorState_index_0 = [...]uint8{0, 6, 18, 30, 42, 54, 66, 78, 90}
)

func (i VehicleDoorState) String() string {
	switch {
	case i <= 7:
		return _VehicleDoorState_name_0[_VehicleDoorState_index_0[i]:_VehicleDoorState_index_0[i+1]]
	case i == 255:
		return _VehicleDoorState_name_1
	default:
		return "VehicleDoorState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
