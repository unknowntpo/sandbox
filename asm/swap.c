#include <stdio.h>

int swap(int *px, int *py)
{
    int temp = *px;
    *px = *py;
    *py = temp;
}
int main()
{
    int x = 3, y = 2;

    swap(&x, &y);
    printf("x = %d, y = %d\n", x, y);
    return 0;
}
