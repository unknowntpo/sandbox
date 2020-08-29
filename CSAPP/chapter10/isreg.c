#include "csapp.h"
int main(int argc, char **argv)
{
    char *type, *readok;
    struct stat stat;
    
    Stat(argv[1], &stat);
    if (S_ISREG(stat.st_mode))
        type = "regular";
    else if (S_ISDIR(stat.st_mode))
        type = "DIR";
    else if (S_ISSOCK(stat.st_mode))
        type = "SOCKET";
    if ((stat.st_mode & S_IRUSR)) /* Check read access */
        readok = "yes";
    else
        readok = "no";
    printf("%s is %s, read = %s", argv[0], type, readok);
    return 0;
}
