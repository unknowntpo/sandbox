#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>

#define AAA 3
#define BBB 6
#define MASK 0x40

uint8_t hexchar2val(uint8_t in)
{
    uint8_t letter = in & 0x40;
    uint8_t offset = (letter >> 3) | (letter >> 6);
    return (in + offset) & 0xf;
}

int main()
{
    char a[] = "123456789AaBbCcDdEeFf";

    for (int i = 0; i < strlen(a); i++)
        printf("%c %d\n", a[i], hexchar2val(a[i]));

    return 0;
}
