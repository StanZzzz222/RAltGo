#ifndef C_PED_H
#define C_PED_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    uint32_t model;
    const Vector3 *position;
    const Vector3 *rotation;
} CPed;

#endif