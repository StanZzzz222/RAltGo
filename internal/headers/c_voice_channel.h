#ifndef C_VOICE_CHANNEL_H
#define C_VOICE_CHANNEL_H

#include <stdint.h>

typedef struct {
    uint32_t id;
    uint8_t spatial;
    float max_distance;
} CVoiceChannel;

#endif