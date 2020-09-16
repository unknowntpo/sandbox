#include <stdio.h>

void swap(int *a, int *b)
{
    int temp;
    
    temp = *a;
    *a = *b;
    *b = temp;
}

int main()
{
    int a = 3, b = 6;
    swap(&a, &b);
    return 0;
}
