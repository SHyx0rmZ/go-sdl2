#include "events.h"
#include <stdio.h>
#include <stdatomic.h>
#include "_cgo_export.h"

void collectEvents(volatile Uint32 *collectEventsQuit) {
  SDL_Event event;
  while(atomic_load(collectEventsQuit) == 0) {

    SDL_PumpEvents();

#if 0
    switch(SDL_PeepEvents(&event, 1, SDL_GETEVENT, SDL_FIRSTEVENT, SDL_LASTEVENT)) {
    case 1:
      if (atomic_load(collectEventsQuit) == 0) {
        propevent(&event);
      }
    }
#else
    if (SDL_WaitEventTimeout(&event, 10) == 1) {
      if (atomic_load(collectEventsQuit) == 0) {
        propevent(&event);
      }
    }
#endif
  }
  atomic_fetch_add(collectEventsQuit, 1);
}
