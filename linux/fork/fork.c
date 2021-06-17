/*
 * My practice of using fork(2) system call
 */
#include <errno.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>

pid_t Fork(void); 

int main() {
  pid_t pid_child;
  if ((pid_child = Fork()) == 0) {
    printf("I am child !, my pid is %d, my parent's pid is %d\n", getpid(),
           getppid());
    sleep(30);
  } else {
    printf("I am parent !, my pid is %d, my child's pid is %d\n", getpid(),
           pid_child);
  }
  sleep(30);
}

pid_t Fork(void) {
  pid_t pid;

  if ((pid = fork()) < 0) {
    perror("fork error");
  }

  return pid;
}
