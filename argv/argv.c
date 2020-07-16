#include <stdio.h>
#if 0
int main(int argc, char **argv)
{
    while (--argc > 0)
        printf("%s%s", *++argv, (argc > 1) ? " " : "");
    printf("\n");
    return 0;
}
#endif

#if 1
int main(int argc, char **argv)
{
    for(int i = 1; i < argc; i++)
        printf("%s%s", argv[i], (i < argc - 1) ? " " : "");
    printf("\n");
    return 0;
}
#endif
