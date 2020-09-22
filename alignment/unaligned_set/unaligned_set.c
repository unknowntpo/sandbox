#include <stdio.h>
#include <stdint.h>
#include <stddef.h>
uint8_t unaligned_get8(void *src) {
    uintptr_t csrc = (uintptr_t) src;
    uint32_t v = *(uint32_t *) (csrc & 0xfffffffc);
    v = (v >> (((uint32_t) csrc & 0x3 * 8)) & 0x000000ff);
    return v;
}

int main()
{
    char *s = "hello";
    uint8_t a = unaligned_get8((void *) s);
    printf("%d\n", a);
}
