/* $begin waitpid1 */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>
#include <setjmp.h>
#include <signal.h>
#include <sys/time.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/stat.h>
#include <errno.h>

#define N 2
/* TODO: try with 2 parameter as input of macro */
#if 1
#define PRINT_PID(i, pid) \
    do { \
        printf("pid[%d] = %d\n", i, pid); \
    } while(0)
#endif


pid_t Fork(void)
{
    pid_t pid;

    if ((pid = fork()) < 0)
        perror("Fork error");
    return pid;
}

int main() 
{
    int status, i;
    pid_t pid[N], retpid; // retpid == returnpid
  
    /* Parent creates N children */
    for (i = 0; i < N; i++) {                       //line:ecf:waitpid1:for
	if ((pid[i] = Fork()) == 0) { /* child */     //line:ecf:waitpid1:fork

            pid_t realpid = getpid();
            printf("realpid = %d\n", realpid);
#ifdef PRINT_PID
            PRINT_PID(i, pid[i]);
#endif
    
            exit(100+i);                          //line:ecf:waitpid1:exit
        }

    }
    

    /* Parent reaps N children in particular order */
    i = 0;
    while ((retpid = waitpid(pid[i++], &status, 0)) > 0) { //line:ecf:waitpid1:waitpid
	if (WIFEXITED(status)){                    //line:ecf:waitpid1:wifexited
	    printf("child %d terminated normally with exit status=%d\n",
		   retpid, WEXITSTATUS(status));     //line:ecf:waitpid1:wexitstatus
        }
	else
	    printf("child %d terminated abnormally\n", pid[i]);
    }

    /* The only normal termination is if there are no more children */
    if (errno != ECHILD)                          //line:ecf:waitpid1:errno
	perror("waitpid error");

    exit(0);
}
/* $end waitpid1 */
