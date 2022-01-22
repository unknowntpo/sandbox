#include <stdio.h>
#include <string.h>
#include <sys/epoll.h>
#include <unistd.h>

int main()
{
    struct epoll_event event;
    int epoll_fd = epoll_create1(0);

    if (epoll_fd == -1) {
        fprintf(stderr, "failed to create epoll fd\n");
        return 1;
    }

    event.events = EPOLLIN;
    event.data.fd = 0;

    if (epoll_ctl(epoll_fd, EPOLL_CTL_ADD, 0, &event)) {
        fprintf(stderr, "failed to add file descripter to epoll\n");
        close(epoll_fd);
        return 1;
    }

    if (close(epoll_fd)) {
        fprintf(stderr, "failed to close epoll fd\n");
        return 1;
    }
    return 0;
}
