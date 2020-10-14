#include <endian.h>
#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

/* Turn it on to show the verbose output */
#define DEBUG_ON 1

/* Show_byte is the function in CS:APP CH2, showing the binary representation of
 * object */
typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t size)
{
#if __BYTE_ORDER == __LITTLE_ENDIAN
    for (size_t i = size; i > 0; i--)
        printf("%.2x ", start[i - 1]);
#else
    for (size_t i = 0; i < size; i++)
        printf("%.2x ", start[i]);
#endif
    puts("");
}
/* Show the binary representation of long integer */
void show_long(long a)
{
    show_bytes((byte_pointer) &a, sizeof(long));
}

void show_int(int a)
{
    show_bytes((byte_pointer) &a, sizeof(int));
}


/* Cpp ver
 * Ref:
 * https://leetcode.com/problems/divide-two-integers/discuss/142849/C%2B%2BJavaPython-Should-Not-Use-%22long%22-Int
 */
int divide(int A, int B)
{
    if (A == INT_MIN && B == -1)
        return INT_MAX;

    /* Take abs of a and b, dealed with boundary condition */
    int a = (A == INT_MIN) ? (unsigned) A : abs(A);
    int b = (B == INT_MIN) ? (unsigned) B : abs(B);
    unsigned int res = 0;
    for (int x = 31; x >= 0; x--)
        if ((signed) ((unsigned) a >> x) - b >= 0)
            res += 1 << x, a -= b << x;
    return (A > 0) == (B > 0) ? res : -res;
}
#if 0
/* Actual devide function doing the division operation */
int divide(int dividend, int divisor)
{
    /* Deal with special condition */
    if (divisor == 0 || (dividend == INT_MIN && divisor == -1))
        return INT_MAX;
    /* Cast to long integer for computing */
    long dvd = (long) labs((long) dividend);
    long dvs = (long) labs((long) divisor);

    /* Construct for answer */
    int sign = (dividend < 0) ^ (divisor < 0) ? -1 : 1;
    unsigned int ans = 0;
    while (dvd >= dvs) {
        unsigned int temp = dvs;
        unsigned int count = 1;
        while ((temp << 1) <= dvd) {
            temp <<= 1;
            count <<= 1;
        }
        printf("temp = %d\t count = %d\t ans = %d\n", temp, count, ans);
        dvd -= temp;
        ans += count;
    }
    return sign < 0 ? -ans : ans;
}
#endif

int mydivide(int dividend, int divisor)
{
    if (divisor == 0 || (dividend == INT_MIN && divisor == -1))
        return INT_MAX;

    int dvd = abs(dividend);
    int dvs = abs(divisor);
    int count = 0;
    int sign = (dividend > 0) ^ (divisor > 0) ? -1 : 1;
    while (dvd >= dvs) {
        dvd -= dvs;
        count += 1;
    }
    return sign * count;
}

typedef struct {
    int dvd;
    int dvs;
    int want;
} test_t;

int test()
{
    test_t test[] = {{-2147483648, 1, -2147483648},
                     {-2147483648, -1, 2147483647},
                     {10, 3, 3},
                     {7, -3, -2},
                     {0, 1, 0},
                     {1, 1, 1}};

    int t_size = sizeof(test) / sizeof(test[0]);
    int ans;
    for (int i = 0; i < t_size; i++) {
        printf("============================ TEST %d =======================\n",
               i + 1);
        printf("test case: dvd = %15d, dvs = %15d\n", test[i].dvd, test[i].dvs);

        ans = divide(test[i].dvd, test[i].dvs);
        if (ans == test[i].want) {
            printf("\033[1;32m");
            printf("PASS\n");
            printf("\033[0m");

            printf("test case: dvd = %15d, dvs = %15d\n", test[i].dvd,
                   test[i].dvs);
            printf("result: %d\n", ans);
        } else {
            printf("\033[1;31m");
            printf("FAIL\n");
            printf("\033[0m");

            printf("expected :       %15d\n", test[i].want);
            printf("get              %15d\n", ans);
        }
        puts("");
    }
    return 0;
}

int main()
{
    test();
}
