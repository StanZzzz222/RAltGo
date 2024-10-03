#ifndef C_PLAYER_H
#define C_PLAYER_H

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint32_t id;
    const char *name;
    const char *ip;
    const char *auth_token;
    const char *social_name;
    uint64_t social_id;
    uint64_t hwid_hash;
    uint64_t hwid_ex_hash;
    const Vector3 *position;
    const Vector3 *rotation;
} CPlayer;

#endif