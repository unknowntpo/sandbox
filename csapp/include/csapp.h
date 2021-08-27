#ifndef __CSAPP_H__
#define __CSAPP_H__

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/errno.h>
#include <sys/stat.h>

/* Our own error-handling functions */
void unix_error(char *msg);

/* Unix I/O wrappers */
void Stat(const char *filename, struct stat *buf);
void Fstat(int fd, struct stat *buf);

#endif /* __CSAPP_H__ */
