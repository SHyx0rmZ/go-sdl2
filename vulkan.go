package sdl

// #include <SDL2/SDL_vulkan.h>
import "C"

func VulkanGetDrawableSize(window *Window, w, h *int) {
	var _w, _h C.int
	C.SDL_Vulkan_GetDrawableSize((*C.SDL_Window)(window), &_w, &_h)
	if w != nil {
		*w = int(_w)
	}
	if h != nil {
		*h = int(_h)
	}
}
