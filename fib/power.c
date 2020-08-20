#include <stdio.h>
#include <stdlib.h>

#define DEBUG_MODE 1 
#define DEBUG(n) \
    printf("fib[%d]\n", n)


/* Demo of Binary Exponentiation with not considering n is odd number
 * Ref: https://www.youtube.com/watch?v=L-Wzglnm4dM
 * at 2:59
 */
long long power(long long a, long long b)
{
    if (b == 0) return 1;
    if (b == 0) return a;

    long long res = power(a, b / 2);
    res *= res;
    if (b % 2)
        res = a * res;
    return res;
} 
int main(int argc, char **argv)
{
    if (argc != 3 || atoi(argv[1]) < 0 || atoi(argv[2]) < 0) {
        printf("<Usage> : ./power [base] [power], base and power are greater than 0\n");
        return 1;
    }

    long long a = atoi(argv[1]);
    long long b = atoi(argv[2]);
    long long result = power(a, b);

    printf("%lld\n", result);
    return 0;
}
