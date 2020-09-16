#include <unistd.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
int main()
{
    
    char *argv[] = {(char *)"sort", NULL};
    dprintf(0, "%s ", "banana");
    dprintf(0, "%s ", "apple");
    dprintf(0, "%s ", "cherry");
    if (execvp(argv[0], argv) < 0) exit(1);
    
    //printf("errno = %d", errno);
    return 0;
}
