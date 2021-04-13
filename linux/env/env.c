#include <unistd.h>
#include <stdio.h>
int main(int argc, char **argv, char **envp)
{
    int i = 0;
    while (argv[i]) {
        printf("argv[%2d] = %s\n",i, argv[i]);
        i++;
    }

    return 0;
}
