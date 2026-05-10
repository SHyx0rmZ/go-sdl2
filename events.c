#include "events.h"
#include <stdio.h>
#include <stdatomic.h>
#include "_cgo_export.h"

static const Uint32 COLLECT_EVENTS_SHOULD_RUN = 0;
static const Uint32 COLLECT_EVENTS_SHOULD_QUIT = 1;
static const Uint32 COLLECT_EVENTS_HAS_QUIT = 2;

void collectEvents(volatile Uint32 *collectEventsQuit) {
  SDL_Event event;
  while(atomic_load(collectEventsQuit) != COLLECT_EVENTS_SHOULD_QUIT) {
    SDL_PumpEvents();
#if 0
    switch(SDL_PeepEvents(&event, 1, SDL_GETEVENT, SDL_FIRSTEVENT, SDL_LASTEVENT)) {
    case 1:
      propevent(&event);
    }
#else
    if (SDL_WaitEventTimeout(&event, 10) == 1) {
      propevent(&event);
    }
#endif
  }
  atomic_store(collectEventsQuit, COLLECT_EVENTS_HAS_QUIT);
}
