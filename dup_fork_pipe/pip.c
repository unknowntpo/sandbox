#include <unistd.h>

#include <sys/types.h>
#include <sys/wait.h>
#include <stdlib.h>
#include <stdio.h>
int main(int argc, char **argv)
{
    int pipefd[2];
    char *c;
    pid_t cpid;

    /* pipe 2 fd */
    pipe(pipefd);
    // use fork to duplicate process
    // parent: write the argv into pipe
    // child: read the argv from pipe and print them using fprintf

    cpid = fork();
    if (cpid < 0) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (cpid == 0){
        close(pipefd[1]);

        while(--argc > 0)
            dprintf(pipefd[0], "%s", *++argv);
        // in child
        // read the argv from pipe and print them 
    } else {
        // in parent
        // read the argv from stdio
        close(pipefd[0]);

        while(--argc > 0)
            dprintf(pipefd[1], "%s", *++argv);
        close(pipefd[1]);
        int status;
        pid_t wpid = waitpid(cpid, &status, 0);
        if (WIFEXITED(status))
            exit(0);
    }
}
