#include <sys/types.h>
#include <sys/wait.h>

int main()
{
    int status;
    pid_t child_pid = fork();

    if(child_pid == 0) {
        // we are in child
        puts("We are children");
        exit(0);
    } else {
        child_pid = waitpid(child_pid, &status, 0) // unknown 
    }

}
