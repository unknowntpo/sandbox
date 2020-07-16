/* Write your own strlen */

#include <stdio.h>

size_t mystrlen(char *s)
{
    size_t n = 0;
    for(; *s; s++)
        n++;
    return n;
}

int main()
{
    char *s = "hello";
    size_t n = mystrlen(s);
    printf("lens = %ld", n);
}
