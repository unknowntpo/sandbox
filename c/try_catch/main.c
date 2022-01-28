// Ref: https://opensourcedoc.com/applied-c-programming/error-handling/
// Ref: [Exceptions in C with Longjmp and
// Setjmp](http://groups.di.unipi.it/~nids/docs/longjump_try_trow_catch.html)
#include <stdio.h>

#include "try_catch.h"

int main(void)
{
    // gcc -E main.c
    /*
    do {
        jmp_buf ex_buf__;
        switch (setjmp(ex_buf__)) {
        case 0:
            while (1) {
                longjmp(ex_buf__, 2);
                printf("Hello World\n");
                break;
            case 1:
                printf("Something wrong\n");
                break;
            case 2:
                printf("More thing wrong\n");
                break;
            case 3:
                printf("Yet another thing wrong\n");
                break;
            }
        default: {
            printf("Clean resources\n");
            break;
        }
        }
    } while (0);
    */

    // clang-format off
    TRY
        THROW(2);
        printf("Hello World\n");
    CATCH(1)
        printf("Something wrong\n");
    CATCH(2)
        printf("More thing wrong\n");
    CATCH(3)
        printf("Yet another thing wrong\n");
    FINALLY
        printf("Clean resources\n");
    ETRY;

    return 0;
}
