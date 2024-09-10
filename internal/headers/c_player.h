#ifndef C_PLAYER_H
#define C_PLAYER_H

#include <stdint.h>
#include "c_verctor3.h"

typedef struct {
    uint32_t id;
    const char *name;
    const char *ip;
    const char *auth_token;
    uint64_t hwid_hash;
    uint64_t hwid_ex_hash;
    const Vector3 *position;
    const Vector3 *rotation;
} CPlayer;

#endif