#include "csapp.h"

// From learn c the hard way ex19
#define log_err(M, ...)                                                       \
    fprintf(stderr, "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__, \
            clean_errno(), ##__VA_ARGS__)

void Fstat(int fd, struct stat *buf)
{
    if (fstat(fd, buf) < 0)
        log_err("Fstat error");
}
