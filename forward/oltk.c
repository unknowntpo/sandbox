#include "oltk.h"
struct oltk {
	struct gr *gr;
	struct oltk_button **zbuttons;
	int n_buttons;
	int zsize;

	int abort;
	struct oltk_button *depressed;

	int fg[N_OLTK_BUTTON_STATES];
	int bg[N_OLTK_BUTTON_STATES];
	int bd[N_OLTK_BUTTON_STATES];
	int color_index;

	struct oltk_rectangle invalid_rect;
};

struct oltk_button {
	struct oltk_rectangle rect;
	int show;
	int sensitive;
	const char *name;

	int state;

	void *callbacks[N_OLTK_BUTTON_CBS];
	void *data[N_OLTK_BUTTON_CBS];

	int fg[N_OLTK_BUTTON_STATES];
	int bg[N_OLTK_BUTTON_STATES];
	int bd[N_OLTK_BUTTON_STATES];

	struct oltk *oltk;
	int update;
};

