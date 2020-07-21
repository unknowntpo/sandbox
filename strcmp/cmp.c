#include <stdio.h>
#include <string.h>

int mystrcmp(char *s, char *t)
{
    while (*s++ == *t++)
        ;
    return (int) (*s-*t);
}


int main()
{
    char *a = "apple";
    char *b = "applepie";
#if 0 // strcmp() in string.h
    int result = strcmp(a, b);
    printf("result = %d\n", result);
#endif
#if 1
    int result = mystrcmp(a, b);
    printf("result = %d\n", result);
#endif
    return 0;
}
