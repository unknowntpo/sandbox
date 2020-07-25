#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
int main(int argc, char **argv, char **envp)
{
#if 0
    for (int i = 1; argv[i] != NULL; i++) {
        printf("argv[%2d] = %s%s", i, argv[i], (argv[i+1] != NULL) ? " " : "");

        printf("\n");
    }
#endif
#if 1
    int i = 1;
    while (--argc > 0) {
        printf("argv[%2d] = %s%s", i, *++argv, (argc > 1) ? " " : "");
        printf("\n");
        i++;
    }
#endif
    for (int i = 0; *envp; i++, envp++)
         printf("envp[%2d] = %s%s\n", i, *envp, (envp[i+1] != NULL) ? " " : "");
    return 0;
}
