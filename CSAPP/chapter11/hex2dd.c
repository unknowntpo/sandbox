#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>


/* TODO: ref: inet_aton source code :
 * https://android.googlesource.com/platform/bionic/+/froyo/libc/inet/inet_aton.c
 */
int main(int argc, char **argv)
{
    if (argc < 2) {
        fprintf(stderr, "<Usage>: ./hex2dd [hexidecimal IP]\n");
        exit(EXIT_FAILURE);
    }

    char *ptr = argv[1];
    int c;
    for(ptr = &ptr[2]; *ptr; ptr++) {
        //c = (int)(*ptr) - 48;
        c = (int)(*ptr);
        printf("%d", c);
        /* TODO: convert c to */
    }
    puts("");
}
