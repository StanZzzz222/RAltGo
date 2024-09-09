#ifndef VEHICLE_H
#define VEHICLE_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    uint32_t model;
    uint8_t primary_color;
    uint8_t second_color;
    const Vector3 *position;
    const Vector3 *rotation;
} CVehicle;

#endif