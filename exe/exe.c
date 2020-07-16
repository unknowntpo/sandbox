#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

int main()
{
    if (fork() == 0) {
        printf("childish\n");
        execlp("ls", "ls", "-la", NULL);
        exit(0);
    } else {
        printf("parental\n");
    }

    exit(0);
}
