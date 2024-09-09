package models

/*
   Create by zyx
   Date Time: 2024/9/5
   File: player.go
*/

type IVehicle struct {
	id           uint32
	model        uint32
	primaryColor uint8
	secondColor  uint8
	position     *Vector3
	rotation     *Vector3
}

func (p *IVehicle) NewIVehicle(id, model uint32, primaryColor, secondColor uint8, position, rotation *Vector3) *IVehicle {
	return &IVehicle{
		id:           id,
		model:        model,
		primaryColor: primaryColor,
		secondColor:  secondColor,
		position:     position,
		rotation:     rotation,
	}
}
