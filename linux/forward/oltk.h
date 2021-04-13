#include <stdio.h>

struct oltk_rectangle {
	int x;
	int y;
	int width;
	int height;
};

typedef struct oltk;
struct oltk_button;

typedef oltk_button_cb_click(struct oltk_button *button, void *data);
typedef oltk_button_cb_draw(struct oltk_button *button,
                            struct oltk_rectangle *rect, void *data);

struct oltk_button *oltk_button_add(struct oltk *oltk,
                        int x, int y,
                        int width, int height);
