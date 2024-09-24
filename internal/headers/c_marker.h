#ifndef C_MARKER_H
#define C_MARKER_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    uint8_t marker_type;
    const Vector3 *position;
} CMarker;

#endif