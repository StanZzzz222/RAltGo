package common

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/models"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: vector3.go
*/

func NewVector3(x, y, z float32) *models.Vector3 {
	return &models.Vector3{X: x, Y: y, Z: z}
}

func NewVector3ARound(x, y, z float32, rangeVal float32) *models.Vector3 {
	source := rand.NewSource(time.Now().UnixNano())
	source.Seed(time.Now().UnixNano())
	r := rand.New(source)
	position := &models.Vector3{X: x, Y: y, Z: z}
	position.X += r.Float32()*(rangeVal*2) - rangeVal
	position.Y += r.Float32()*(rangeVal*2) - rangeVal
	return position
}

func NewVector3Collection(points [][][]float32) []*models.Vector3 {
	var slice []*models.Vector3
	for _, v := range points {
		for _, point := range v {
			if len(point) >= 3 {
				slice = append(slice, &models.Vector3{
					X: point[0],
					Y: point[1],
					Z: point[2],
				})
			}
		}
	}
	return slice
}

func NewVector3FromStr(position string) (*models.Vector3, error) {
	position = strings.TrimSpace(position)
	arr := strings.Split(position, ",")
	if len(arr) != 3 {
		return nil, fmt.Errorf("invalid position format")
	}
	var coords [3]float32
	for i, str := range arr {
		var err error
		coord, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
		if err != nil {
			return nil, err
		}
		coords[i] = float32(coord)
	}
	return NewVector3(coords[0], coords[1], coords[2]), nil
}

func GetVectoe3Distance(v1, v2 *models.Vector3) float32 {
	return float32(math.Sqrt(float64(
		(v2.X-v1.X)*(v2.X-v1.X) +
			(v2.Y-v1.Y)*(v2.Y-v1.Y) +
			(v2.Z-v1.Z)*(v2.Z-v1.Z))))
}
