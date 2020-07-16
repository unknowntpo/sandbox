#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

pid_t Fork()
{
    pid_t pid;
    if ((pid = fork()) < 0)
        perror("fork failed");
    
    return pid;
}

int main()
{
    pid_t pid = Fork();
    if (pid == 0) {
        printf("childish");
        execlp("ls", "ls", "-la", NULL);
        exit(0);
    }else {
        printf("parental");
        waitpid(pid, NULL, 0);
        exit(0);
    }
}

