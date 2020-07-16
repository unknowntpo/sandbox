#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>

/* SIGINT handler */
void handler(int num)
{
    write(STDOUT_FILENO, "I won't die!\n", 13);
}

int main() {
  while (1) {
    signal(SIGINT, handler);
    printf("Wasting cycles ..., my pid = %d\n", getpid());
    sleep(1);
  }
}
