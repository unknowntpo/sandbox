/* Write your own strlen */

#include <stdio.h>

size_t mystrlen(char *s)
{
    size_t len = 0;
    while (*s++)
        len++;
    return len;
}

int main()
{
    char *s = "hello";
    size_t n = mystrlen(s);
    printf("lens = %ld", n);
}
