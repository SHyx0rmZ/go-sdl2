#include "events.h"

int collectEventsQuit;

void collectEvents() {
  SDL_Event event;
  while(!collectEventsQuit) {
    if (SDL_WaitEvent(&event) == 1) {
      propevent(&event);
    }
  }
}
