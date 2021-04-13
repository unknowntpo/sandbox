#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

#define MAXSIZE 10

typedef struct buf {
    char buf[MAXSIZE];
    char *bufptr;
} buf_t;

void rio_readinitb(rio_t *rp, int fd)
{
    rp->rio_fd = fd;

}
int main()
{
    
    int fd = open("foo.txt", O_RDONLY);
    if (fd < 0) {
        perror("opening");
        exit(1);
    }
    int n_read = read(fd, buf_t->buf, 6);
    if (n_read < 0) {
        perror("reading");
        exit(1);
    }
    
    printf("number of byte read: %d\n", n_read);

    for(int i = 0; buf[i] < MAXSIZE; i++){
        printf("buf[%d] = %d\n", i, buf[i]);
    }
    return 0;


}
