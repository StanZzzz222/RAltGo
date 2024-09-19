#ifndef C_COLSHAPE_H
#define C_COLSHAPE_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    uint32_t colshape_type;
    const Vector3 *position;
} CColshape;

#endif