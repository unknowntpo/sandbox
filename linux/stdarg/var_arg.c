#include <stdarg.h>
#include <stdio.h>

void foo(char *fmt, ...)
{
    va_list ap;
    int d;
    char c, *s;

    va_start(ap, fmt);
    while (*fmt) {
        switch (*fmt++) {
        case 's': /* string */
            s = va_arg(ap, char *);
            printf("string %s\n", s);
            break;
        case 'd': /* int */
            d = va_arg(ap, int);
            printf("int %d\n", d);
            break;
        case 'c': /* char */
            /* need a cast here since va_arg only
               takes fully promoted types */
            c = (char) va_arg(ap, int);
            printf("char %c\n", c);
            break;
        }
    }
    va_end(ap);
}

int main()
{
    foo("s", "apple");
    foo("d", 1);
    foo("c", "c");
    
}
