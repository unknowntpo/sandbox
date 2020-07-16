#include <limits.h>
#include <stdio.h>


#if LONG_MAX == 2147483647L
#define DETECT(X) \
    (((X) - 0x01010101) & ~(X) & 0x80808080)
#else
#if LONG_MAX == 9223372036854775807L
#define DETECT(X) \
    (((X) - 0x0101010101010101) & ~(X) & 0x8080808080808080)
#else
#error long int is not a 32bit or 64bit type.
#endif
#endif

void show_binary(long int i)
{
    /* TODO: figure out how to print binary */
    if (i == 0)
        return;
    printf("%ld: ", i);
    while (i != 0) {
        printf("%ld ", i % 2);
        i >> 1;
    }
    return;
}


int main()
{
    long int i = 10;
    show_binary(i);
    printf("LONG_MAX = %ld\n", LONG_MAX);
    printf("detect null = %ld\n", DETECT(i));
    return 0;
}
