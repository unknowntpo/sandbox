#include "csapp.h"


int main(int argc, **argv) // why need **argv ptr to ptr?
{
    int n;
    rio_t rio;
    char buf[MAXLINE];

    Rio_readinitb(&rio, STDIN_FILENO); // init the rio buffer
    while((n = Rio_readlineb(&rio, buf,MAXLINE)) != 0) // if not encounter EOF
        Rio_writen(STDOUT_FILENO, buf, n);
}
