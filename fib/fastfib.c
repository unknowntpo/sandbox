#include <stdio.h>
#include <stdlib.h>

unsigned int fastfib(int n)
{
    if (n == 0)
        return 0;
    else if (n <= 2)
        return 1;
    if (n % 2) {
        unsigned int k = (n - 1) / 2;
        return fastfib(k + 1) * fastfib(k + 1) + fastfib(k) * fastfib(k);
    } else {
        unsigned int k = n / 2;
        return 2 * fastfib(k) * (fastfib(k + 1) - fastfib(k));
    }
}

int main(int argc, char **argv)
{
    /* FIXME: Why fastfib(6) == 0 ? */
    if (argc != 2 || atoi(argv[1]) < 0) {
        printf("<Usage> : ./fastfib [fibnum], fibnum >= 0\n");
        return 1;
    }

    int n = atoi(argv[1]);
    
    unsigned int result = fastfib(n);

    printf("%d\n", result);
    return 0;
}
