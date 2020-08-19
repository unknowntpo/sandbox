#include <stdio.h>
#include <stdlib.h>

#define DEBUG_MODE 1 
#define DEBUG(n) \
    printf("fib[%d]\n", n)

unsigned int fastfib(unsigned n)
{
#if DEBUG_MODE
    DEBUG(n);
#endif
    if (n == 0)
        return 0;
    else if (n <= 2)
        return 1;
    if (n % 2) {
        unsigned int k = (n - 1) / 2;
        return fastfib(k + 1) * fastfib(k + 1) + fastfib(k) * fastfib(k);
    } else {
        unsigned int k = n / 2;
        return fastfib(k) * (2 * fastfib(k + 1) - fastfib(k));
    }
}
 
int main(int argc, char **argv)
{
    if (argc != 2 || atoi(argv[1]) < 0) {
        printf("<Usage> : ./fastfib [fibnum], fibnum >= 0\n");
        return 1;
    }

    unsigned int n = atoi(argv[1]);
    unsigned int result = fastfib(n);

    printf("%d\n", result);
    return 0;
}
