package entity

import "C"
import "github.com/StanZzzz222/RAltGo/models"

/*
   Create by zyx
   Date Time: 2024/9/9
   File: entitys.go
*/

type CPlayer struct {
	ID         uint32
	Name       string
	IP         string
	AuthToken  string
	HWIDHash   uint64
	HWIDExHash uint64
	Position   *models.Vector3
	Rotation   *models.Vector3
}

type CVehicle struct {
	ID           uint32
	Model        uint32
	PrimaryColor uint8
	SecondColor  uint8
	Position     *models.Vector3
	Rotation     *models.Vector3
}
