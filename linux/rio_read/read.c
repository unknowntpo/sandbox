#include "csapp.h"

int main(int argc, char **argv)
{
    int n;
    rio_t rio;
    char buf[MAXLINE];

    int fd1;
    fd1 = Open(argv[1], O_RDWR, 0);

    Rio_readinitb(&rio, fd1);
    while((n = Rio_readlineb(&rio, buf, MAXLINE)) != 0)
        Rio_writen(STDOUT_FILENO, buf, n);

    Close(fd1);
}
