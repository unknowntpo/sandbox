#include <stdio.h>
#include <stdlib.h>

#define DEBUG_MODE 1
#define DEBUG(n) printf("fib[%lld]\n", n)

long long fib_tail(long long n, long long a, long long b);
long long fib(long long n);
long long fastfib(long long n);
/* Function pointer to selected fib function */
long long (*fp)(long long n);


long long fib_tail(long long n, long long a, long long b)
{
    if (n == 0) return a;
    return fib_tail(n - 1, b , a + b);
}
/* Calling the actual tail recursion fib funciton */
long long fib(long long n)
{
    return fib_tail(n, 0, 1);
}

long long fastfib(long long n)
{
#if DEBUG_MODE
    DEBUG(n);
#endif
    if (n == 0)
        return 0;
    else if (n <= 2)
        return 1;
    if (n % 2) {
        long long k = (n - 1) / 2;
        return fastfib(k + 1) * fastfib(k + 1) + fastfib(k) * fastfib(k);
    } else {
        long long k = n / 2;
        return fastfib(k) * (2 * fastfib(k + 1) - fastfib(k));
    }
}

int main(int argc, char **argv)
{
    if (argc != 2 || atoi(argv[1]) < 0) {
        printf("<Usage> : ./fastfib [fibnum], fibnum >= 0\n");
        return 1;
    }
    /* TODO: Use array of function pointer to switch different fib method */
    //fp = fastfib;
    fp = fib;
    long long n = atoll(argv[1]);
    long long result = fp(n);

    printf("%lld\n", result);
    return 0;
}
