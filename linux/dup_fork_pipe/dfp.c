#include <unistd.h>
#include <fcntl.h>              /* Obtain O_* constant definitions */
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>


pid_t Fork(void)
{
    pid_t pid;
    if ((pid = fork()) < 0)
        perror("fork");
    return pid;
}

int main()
{
    pid_t pid;
    int status;
    int fds[2];
    pipe(fds);
    
    if ((pid = Fork()) == 0) {
        dup2(fds[0], STDIN_FILENO);
        close(fds[0]);
        close(fds[1]);
        char **argv = {(char *)"sort", NULL};  // create argument vector
        if (execvp(argv[0], argv) < 0) {
            perror("execvp");
            exit(EXIT_FAILURE);
        }
    } else {
        // parental
        close(fds[0]);
        const char *words[] = {"pear", "peach", "apple"};

        /* write inputs to writable fd */
        size_t numwords = sizeof(words) / sizeof (words[0]); // What's the data type of words?
        for (size_t i = 0; i < numwords; i++) {
            dprintf(fds[1], "%s\n", words[i]);
        }

        close(fds[1]);

        pid_t wpid = waitpid(pid, &status, 0); // what if it move to the end?
        // what's the return value of waitpid?    
        return wpid == pid && WIFEXITED(status) ? WEXITSTATUS(status) : -1;   
     }
}
