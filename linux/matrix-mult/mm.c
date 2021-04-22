#include <inttypes.h>
#include <stdio.h>

#define SIZE 2
typedef int array[SIZE][SIZE];

void printarray(int (*pa)[SIZE])
{
    printf("addr of a: %p\n", (uintptr_t) pa);
    for (int i = 0; i < SIZE; i++)
        for (int j = 0; j < SIZE; j++)
            printf("a[%d][%d] = %d\n", i, j, pa[i][j]);
}
int main()
{
    int a[SIZE][SIZE] = {{0, 1}, {2, 3}};
    int(*pa)[SIZE] = &a;
    printarray(&a);
    printf("addr of a: %p\n", (uintptr_t)(&a));
}
