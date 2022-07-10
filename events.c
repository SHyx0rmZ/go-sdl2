#include "events.h"
#include <stdio.h>
#include "_cgo_export.h"

volatile Uint32 collectEventsQuit;

int catchAppTerminating(void *userdata, SDL_Event *event) {
    if (event->type == SDL_APP_TERMINATING) {
        collectEventsQuit = 1;
    }
    return 1;
}

void collectEvents(volatile Uint32 *collectEventsQuit) {
  SDL_Event event;
  while(*collectEventsQuit == 0) {

    SDL_PumpEvents();

#if 0
    switch(SDL_PeepEvents(&event, 1, SDL_GETEVENT, SDL_FIRSTEVENT, SDL_LASTEVENT)) {
    case 1:
      if (*collectEventsQuit == 0) {
        propevent(&event);
      }
    }
#else
    if (SDL_WaitEventTimeout(&event, 10) == 1) {
      if (*collectEventsQuit == 0) {
        propevent(&event);
      }
    }
#endif
  }
  *collectEventsQuit = 2;
}
