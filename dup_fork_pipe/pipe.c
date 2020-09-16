#include <unistd.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    int fds[2];
    pipe(fds);
    
    pid_t pid = fork();

    if (pid == 0) {
        dup2(fds[0], STDIN_FILENO);
        /* redirection complete, close fds[0], and fds[1] */
        close(fds[0]);
        close(fds[1]);
        char *argv[] = {(char *)"sort", NULL};  // create argument vector
        if(execvp(argv[0], argv) < 0) exit(0);  // // run sort command (exit if something went wrong)
    } else {
        close(fds[0]);
        
        const char *words[] = {"pear", "peach", "apple"};
        size_t numwords = sizeof(words) / sizeof(words[0]);

        /* Write data into fds[1] (which will go to child process's stdin) */
        for(size_t i=0; i < numwords; i++){
            dprintf(fds[1], "%s\n", words[i]);
        }
        /* Send EOF so child can continue
         * (child blocks until all input has been processed) */
        close(fds[1]);
        int status;
        pid_t wpid = waitpid(pid, &status, 0); // wait for child to finish before parent exiting
        printf("status= %d\n", status);
        printf("WIFEXITED(status) = %d\n", WIFEXITED(status));
        printf("WEXITSTATUS(status) = %d\n", WEXITSTATUS(status));
        return wpid = pid && WIFEXITED(status) ? WEXITSTATUS(status) : -1;

    }


}
