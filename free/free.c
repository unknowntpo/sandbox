#if 0
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main()
{
    size_t len = strlen("hello");
    char *p = malloc(len + 1);
    if (!p) {
        perror("Fail to malloc");
        exit(EXIT_FAILURE);
    }
    strncpy(p, "hello", len);
    p[len] = '\0';
    free(p);
    return 0;
}
#endif

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    int *p = malloc(100);

    printf("空間位址：%p\n", p);
    printf("儲存的值：%d\n", *p);

    *p = 200;

    printf("空間位址：%p\n", p);
    printf("儲存的值：%d\n", *p);

    free(p);

    printf("空間位址：%p\n", p);
    return 0;
}
