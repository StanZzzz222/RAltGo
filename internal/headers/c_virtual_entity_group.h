#ifndef C_VIRTUAL_ENTITY_GROUP_H
#define C_VIRTUAL_ENTITY_GROUP_H

#include <stdint.h>

typedef struct {
    uint32_t id;
    uint32_t max_entities_in_stream;
} CVirtualEntityGroup;

#endif