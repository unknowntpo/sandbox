#include <stdio.h>
#include <stdlib.h>

#define DEBUG_MODE 1 
#define DEBUG(n) \
    printf("fib[%d]\n", n)


/* Demo of Binary Exponentiation with not considering n is odd number
 * Ref: https://www.youtube.com/watch?v=L-Wzglnm4dM
 * at 2:59
 */
long long power_recur(long long a, long long b)
{
    if (b == 0) return 1;
    if (b == 1) return a;

    long long res = power_recur(a, b / 2);  // get a^(b/2)
    res *= res;  // res = (a^(b/2)) ^ 2
    if (b % 2)   // if b is odd
        res = a * res;  // res = a * (a^(b/2)) ^ 2
    return res;
} 
long long power_iter(long long a, long long b)
{
    long long result = 1;

    for(; b != 0; b /= 2) {
        if (b % 2) result *= a;
        a *= a;
    }
    return result;
}
int main(int argc, char **argv)
{
    if (argc != 3 || atoi(argv[1]) <= 0 || atoi(argv[2]) < 0) {
        printf("<Usage> : ./power [base] [power], base > 0, power >= 0\n");
        return 1;
    }

    long long a = atoi(argv[1]);
    long long b = atoi(argv[2]);
#if 0
    long long result = power_recur(a, b);
#endif
    long long result = power_iter(a, b);
    printf("%lld\n", result);
    return 0;
}
