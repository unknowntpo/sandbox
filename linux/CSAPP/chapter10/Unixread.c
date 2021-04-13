/* p.626
 * read example
 */
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <ctype.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>


int main() {
  char c;

  while (read(STDIN_FILENO, &c, 1) != 0)
    write(STDOUT_FILENO, &c, 1);
  exit(0);
}
