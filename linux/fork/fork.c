#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>
int main()
{
    if (fork() == 0) {
        printf("childish\n");
        execlp("ls", "ls", "-al", NULL);
        printf("after execlp ...\n");  // won't be executed
    } else {
        printf("I'm parental: my pid = %d/n", getpid());
    }
}
