#include "stdbool.h"
#include "stdio.h"

// clang-format off
#define WARN_IF(EXP) \
    do { \
        if (EXP) \
            fprintf( stderr, "Warning: " #EXP "\n"); \
    } while(0)
// clang-format on


int main()
{
    bool earthquake = true;
    WARN_IF(earthquake);
}
