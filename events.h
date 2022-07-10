#ifndef _WITCHES_EVENTS_H
#define _WITCHES_EVENTS_H 1

#include <SDL2/SDL.h>

extern volatile Uint32 collectEventsQuit;

void collectEvents(volatile Uint32 *);

#endif
