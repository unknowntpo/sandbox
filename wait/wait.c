#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
int main(int argc, char **env)
{
    int status;

    pid_t child_pid = fork();
    
    char *argv[] = {"/bin/ls", "-l", NULL};
    
    if (child_pid == 0) {
        execve("/bin/ls", argv, env);
        printf("child won't comes here\n");
    }
    else {
        waitpid(child_pid, &status, 0);
        exit(0);
    }
} 
