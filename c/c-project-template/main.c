#include <stdio.h>

#include "counter.h"

int main(int argc, const char *argv[])
{
    cnt_t cnt;
    add(&cnt);
    printf("counter: %d\n", cnt);

    return 0;
}
