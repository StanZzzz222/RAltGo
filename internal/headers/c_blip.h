#ifndef C_BLIP_H
#define C_BLIP_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    uint32_t blip_type;
    uint32_t sprite_id;
    uint32_t color;
    const char *name;
    const Vector3 *position;
    float rot;
} CBlip;

#endif