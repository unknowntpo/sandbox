#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#define AAA 3
#define BBB 6
#define MASK 0x40

uint8_t hexchar2val(uint8_t in)
{
    const uint8_t letter = in & 0x40;
    const uint8_t shift = (letter >> 3) | (letter >> 6);
    return (in + shift) & 0xf;
}

int main()
{
    char a[] = "f";

    for(int i = 0; i < strlen(a);i++) 
        printf("%c %d\n", a[i], hexchar2val(a[i]));
        
    return 0;
}
