#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

pid_t Fork()
{
    pid_t pid;
    if ((pid = fork()) < 0){
        perror("fork");
        exit(EXIT_FAILURE);
    }
    return pid;
}

int main()
{
    int wstatus;
    pid_t pid;
    if (Fork(pid) == 0) {
        printf("I'm childish\n");
        exit(0);
    } else {
        waitpid(-1, &wstatus, 0);
        
        if (WIFEXITED(wstatus)) {// exit normally
            printf("exit with exit status = %d\n", WIFEXITED(wstatus));  
        } else {
            printf("exit failed with exit status = %d\n", WIFEXITED(wstatus));
            perror("Exit");
        }
        return 0;
    }
    
}
