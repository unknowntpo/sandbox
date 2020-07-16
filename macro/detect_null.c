#include <limits.h>

#define UNALIGNED(X, Y)                                                        \
  (((long)X & (sizeof(long) - 1)) | ((long)Y & (sizeof(long) - 1)))

#if LONG_MAX == 2147483647L
#define DETECTNULL(X) (((X)-0x01010101) & ~(X)&0x80808080)
#else
#if LONG_MAX == 9223372036854775807L
/* Nonzero if X (a long int) contains a NULL byte. */
#define DETECTNULL(X) (((X)-0x0101010101010101) & ~(X)&0x8080808080808080)
#else
#error long int is not a 32bit or 64bit type.
#endif
#endif

int main()
{
    unsigned long int x = 0;
    x = DETECTNULL(x);
    printf("x = %ul", x); 
}
