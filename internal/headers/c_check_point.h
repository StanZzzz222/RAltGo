#ifndef C_CHECKPOINT_H
#define C_CHECKPOINT_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    uint8_t check_point_type;
    const Vector3 *position;
} CCheckpoint;

#endif