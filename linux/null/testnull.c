#include <stdio.h>


int main()
{
    char *c = '\0';
#if 1
    printf("%c\n", *c);
#endif
#if 0
    char *null_ptr = 0;
    printf("%c\n", *null_ptr);
#endif
    return 0;
}
