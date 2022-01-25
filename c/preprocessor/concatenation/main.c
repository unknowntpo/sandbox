#include "stdio.h"

typedef struct command {
    char *name;
    void (*function)(void);
}

// clang-format off
#define COMMAND(NAME) { #NAME, NAME##_command }
// clang-format on


int main()
{
    struct command commands[] = {
        COMMAND(quit),
        COMMAND(help),
    };
}
