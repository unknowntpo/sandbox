#include <stdio.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t len)
{
    size_t i;
    for (i = 0; i < len; i++)
       printf(" %.2x", start[i]);
    printf("\n"); 
}

void show_int(int x)
{
    show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x)
{
    show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void *x)
{
    show_bytes((byte_pointer) &x, sizeof(void *));
}

int main()
{
    int a = 1;
    float b = 1.0;
    int *p = &a;
    printf("show int %d...\n", a);
    show_int(a);
    printf("show float %f...\n", b);
    show_float(b);
    printf("show pointer %p...\n", p);
    show_pointer(p);

    return 0;
}
