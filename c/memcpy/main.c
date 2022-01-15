#include <stdio.h>
#include <string.h>

// TODO: Write the test!
void copy(int *dst, int *src, int len)
{
    memcpy(dst, src, len * sizeof(int));
}

int main()
{
    int src[] = {1, 2, 3, 4};
    int len = sizeof(src) / sizeof(int);
    // FIXME: Why C99 compiler does not zero the memory for us ?
    int dst[len];

    // zero the memory.
    memset(dst, 0, len * sizeof(int));

    for (int i = 0; i < len; i++) {
        printf("%d ", dst[i]);
    }

    copy(dst, src, len);
    printf("dst: [ ");
    for (int i = 0; i < len; i++) {
        printf("%d ", dst[i]);
    }
    printf("]\n");
}
