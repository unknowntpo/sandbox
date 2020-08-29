#include <stdio.h>

int x = -1;
unsigned int u = 2147483648; /* 2 to the 31st */

int main()
{
    printf("x = %u = %d\n", x, x);
    printf("u = %u = %d\n", u, u);

    return 0;
}

