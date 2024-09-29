#ifndef C_VIRTUAL_ENTITY_H
#define C_VIRTUAL_ENTITY_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    uint32_t streaming_distance;
    const Vector3 *position;
} CVirtualEntity;

#endif