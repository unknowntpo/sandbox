#include <stdio.h>
#include <stdlib.h>

#define DEBUG_MODE 1
#define DEBUG(n) printf("fib[%lld]\n", n)

long long do_fib_tail(long long n, long long a, long long b);
long long fib_tail(long long n);
long long fib_iter(long long n);
long long fib_fast(long long n);
void register_fib_method(long long n);
/* Function pointer to selected fib function */
long long (*fp)(long long n);

enum {
    FIB_TAIL,
    FIB_ITER,
    FIB_FAST,
};


long long do_fib_tail(long long n, long long a, long long b)
{
    if (n == 0)
        return a;

    return do_fib_tail(n - 1, b, a + b);
}
/* Calling the actual tail recursion fib funciton */
long long fib_tail(long long n)
{
    return do_fib_tail(n, 0, 1);
}


/* Iterative counting fibonacci numbers */
long long fib_iter(long long n)
{
    long long a = 0;       // f(0)
    long long b = 1;       // f(1)
    long long result = 0;  // result

    if (n == 0)
        return 0;
    if (n == 1)
        return 1;

    for (long long i = 2; i <= n; i++) {
        result = a + b;
        a = b;
        b = result;
    }

    return result;
}
long long fib_fast(long long n)
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
        return fib_fast(k + 1) * fib_fast(k + 1) + fib_fast(k) * fib_fast(k);
    } else {
        long long k = n / 2;
        return fib_fast(k) * (2 * fib_fast(k + 1) - fib_fast(k));
    }
}

int main(int argc, char **argv)
{
    /* TODO: Explicitly report bug if user input wrong method number */
    /* TODO: How to display the Method table automatically? */
    if (argc != 3 || atoi(argv[1]) < 0) {
        printf("<Usage> : ./fastfib [fibnum] [method number], fibnum >= 0\n");
        return 1;
    }
    long long n = atoll(argv[1]);
    int method = atoi(argv[2]); 
    /* Register fib method */
    register_fib_method(method);

    /* Do fib number counting */
    long long result = fp(n);

    printf("%lld\n", result);
    return 0;
}

void register_fib_method(long long n)
{
    /* Register fib method */
    switch (n) {
        case FIB_TAIL:
            fp = fib_tail;
            puts("tail recursion");
            break;
        case FIB_ITER:
            fp = fib_iter;
            puts("iteration");
            break;
        case FIB_FAST:
            fp = fib_fast;
            puts("fast doubling");
            break;
        default:
            fp = fib_fast;
            puts("default method");
    }
}
