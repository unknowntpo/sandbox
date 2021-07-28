#include <stdio.h>

#define __max(t1, t2, max1, max2, x, y) \
    ({                                  \
        t1 max1 = (x);                  \
        t2 max2 = (y);                  \
        (void) (&max1 == &max2);        \
        max1 > max2 ? max1 : max2;      \
    })
#define max(x, y) \
    __max(typeof(x), typeof(y), __UNIQUE_ID(max1_), __UNIQUE_ID(max2_), x, y)

/* at include/linux/compiler-gcc.h since linux v4.9 */
#define __UNIQUE_ID(prefix) __PASTE(__PASTE(__UNIQUE_ID_, prefix), __COUNTER__)

#define ___PASTE(a, b) a##b
#define __PASTE(a, b) ___PASTE(a, b)

int main()
{
    int a = 3, b = 2;
    printf("max of %d, %d ? %d\n", a, b, max(a, b));
}
