#include "csapp.h"

int main(int argc, char **argv) // why need **argv ptr to ptr?
{
    int n = strlen("hello") + 1;
    rio_t *rp = malloc(sizeof(rio_t));
    rio_readinitb(rp, STDIN_FILENO);
    char usrbuf[n];

    while(Rio_readlineb(rp, usrbuf, n) != 0) {
        rio_writen(STDOUT_FILENO, usrbuf, n);
    }
    return 0;
}


