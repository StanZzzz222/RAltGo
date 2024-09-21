package models

import (
	"encoding/json"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/logger"
	"os"
	"reflect"
)

/*
   Create by zyx
   Date Time: 2024/9/17
   File: mvalue.go
*/

type MValues []any

func NewMValues(args ...any) *MValues {
	mvalues := MValues(args)
	return &mvalues
}

func (mv *MValues) Dump() string {
	var obj []any
	for _, arg := range *mv {
		t := reflect.TypeOf(arg)
		switch t.Kind() {
		case reflect.Ptr:
			switch t.Elem() {
			case reflect.TypeOf((*IPlayer)(nil)).Elem():
				param, ok := arg.(*IPlayer)
				if !ok {
					logger.LogErrorf("Invalid type for Player: %v", param.GetId())
					continue
				}
				obj = append(obj, map[string]any{
					"value": param.GetId(),
					"type":  enum.Player,
				})
				continue
			case reflect.TypeOf((*IVehicle)(nil)).Elem():
				param, ok := arg.(*IVehicle)
				if !ok {
					logger.LogErrorf("Invalid type for Vehicle: %v", param.GetId())
					continue
				}
				obj = append(obj, map[string]any{
					"value": param.GetId(),
					"type":  enum.Vehicle,
				})
				continue
			case reflect.TypeOf((*IBlip)(nil)).Elem():
				param, ok := arg.(*IBlip)
				if !ok {
					logger.LogErrorf("Invalid type for Blip: %v", param.GetId())
					continue
				}
				obj = append(obj, map[string]any{
					"value": param.GetId(),
					"type":  enum.Vehicle,
				})
				continue
			case reflect.TypeOf((*IPed)(nil)).Elem():
				param, ok := arg.(*IPed)
				if !ok {
					logger.LogErrorf("Invalid type for Ped: %v", param.GetId())
					continue
				}
				obj = append(obj, map[string]any{
					"value": param.GetId(),
					"type":  enum.Vehicle,
				})
				continue
			case reflect.TypeOf((*IColshape)(nil)).Elem():
				param, ok := arg.(*IColshape)
				if !ok {
					logger.LogErrorf("Invalid type for Colshape: %v", param.GetId())
					continue
				}
				obj = append(obj, map[string]any{
					"value": param.GetId(),
					"type":  enum.Colshape,
				})
				continue
			}
		case reflect.Invalid, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			obj = append(obj, map[string]any{
				"value": arg,
				"type":  t.Kind().String(),
			})
			continue
		case reflect.Bool, reflect.Float32, reflect.Float64:
			obj = append(obj, map[string]any{
				"value": arg,
				"type":  t.Kind().String(),
			})
			continue
		case reflect.String, reflect.Map, reflect.Struct, reflect.Array, reflect.Slice:
			obj = append(obj, map[string]any{
				"value": arg,
				"type":  t.Kind().String(),
			})
			continue
		default:
			logger.LogErrorf("Unknow MValue type: %v", t.Kind().String())
			os.Exit(1)
		}
	}
	dumpBytes, err := json.Marshal(&obj)
	if err != nil {
		logger.LogErrorf("Dump mvalues falied, %v", err.Error())
		return ""
	}
	return string(dumpBytes)
}
