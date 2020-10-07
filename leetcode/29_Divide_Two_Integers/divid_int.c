#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
int divide(int dividend, int divisor)
{
    /* Deal with special condition */
    if (dividend > INT_MAX || dividend < INT_MIN || divisor > INT_MAX ||
        divisor < INT_MIN || divisor == 0)
        return INT_MAX;

    /* Cast to long integer for computing */
    long dvd = (long) abs(dividend);
    long dvs = (long) abs(divisor);

    /* Construct for answer */
    int sign = (dividend < 0) ^ (divisor < 0) ? -1 : 1;
    int ans = 0;
    while (dvd >= dvs) {
        int temp = dvs;
        int count = 1;
        while ((temp << 1) <= dvd) {
            temp <<= 1;
            count <<= 1;
        }
        dvd -= temp;
        ans += count;
    }
    return sign * ans;
}
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
    int a;
    int b;
    int want;
} test_t;


int test()
{
    test_t test[] = {{-2147483648, 1, -2147483648},
                     {10, 3, 3},
                     {7, -3, -2},
                     {0, 1, 0},
                     {1, 1, 1}};

    int t_size = sizeof(test) / sizeof(test[0]);
    int ans;
    for (int i = 0; i < t_size; i++) {
        ans = divide(test[i].a, test[i].b);
        if (ans == test[i].want) {
            puts("======================== PASS ========================");
            printf("test case: a = %15d, b = %15d\n", test[i].a, test[i].b);
            printf("result: %d\n", ans);
        } else {
            puts("======================== FAIL ========================");
            printf("test case: a = %15d, b = %15d\n", test[i].a, test[i].b);
            printf("expected :     %15d\n" , test[i].want);
            printf("get            %15d\n", ans);
        }
        puts("");
    }
    return 0;
}

int main() {
    test();
}
