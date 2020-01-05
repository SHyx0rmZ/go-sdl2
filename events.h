#ifndef _WITCHES_EVENTS_H
#define _WITCHES_EVENTS_H 1

#include <SDL2/SDL.h>

extern int collectEventsQuit;

void propevent(SDL_Event *);
void collectEvents();

#endif
