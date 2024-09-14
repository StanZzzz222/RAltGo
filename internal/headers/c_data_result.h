#ifndef C_DATA_RESULT
#define C_DATA_RESULT

#include <stdint.h>
#include "c_vector3.h"

typedef struct {
    uint8_t tag;
    union {
        uint8_t u8_val;
        uint16_t u16_val;
        uint32_t u32_val;
        uint64_t u64_val;
        int bool_val;
        const char *cstring_val;
        const Vector3 *vector3_val;
    } data;
} CDataResult;

#endif