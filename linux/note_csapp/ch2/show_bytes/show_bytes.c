#include <stdio.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t len)
{
    for (size_t i = 0; i < len; i++) {
        printf("%.2x ", start[i]);
    }
    printf("\n"); 
}

void show_int(int x)
{
    return show_bytes((byte_pointer)&x, sizeof(int));
}

void show_float(float x)
{
    return show_bytes((byte_pointer)&x, sizeof(float));
}

void show_pointer(void *x)
{
    return show_bytes((byte_pointer)&x, sizeof(void *));
}

void show_endian()
{
    int x = 1;
    byte_pointer byte = (byte_pointer)&x;
    if (*byte != 1)
        puts("big");
    else
        puts("little");
    return;
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

    puts("Check endian...");
    show_endian();
    return 0;
}
