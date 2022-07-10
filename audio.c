#include <SDL2/SDL.h>
#include "_cgo_export.h"

void callback(void *userdata, Uint8 *stream, int len) {
  audioSpecCallback(userdata, stream, len);
}
