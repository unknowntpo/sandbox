#include "csapp.h"

/* Unix style error */
void unix_error(char *msg)
{
    fprintf(stderr, "%s: %s\n", msg, strerror(errno));
    exit(0);
}

void Stat(const char *filename, struct stat *buf)
{
    if (stat(filename, buf) < 0)
        unix_error("Stat error");
}

void Fstat(int fd, struct stat *buf)
{
    if (fstat(fd, buf) < 0)
        unix_error("Fstat error");
}
