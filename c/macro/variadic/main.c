#include <stdio.h>

// Ref: https://gcc.gnu.org/onlinedocs/cpp/Variadic-Macros.html
#define wrong_eprintf(fmt, ...) fprintf(stderr, fmt, __VA_ARGS__)

// TODO: What's the meaning of __VA_OPT__() ?
// TODO: How to detect __VA_OPT__() is supported or not?
// Ref: https://stackoverflow.com/a/48045656
#define opt_eprintf(format, ...) \
    fprintf(stderr, format __VA_OPT__(, ) __VA_ARGS__)

#define correct_eprintf(fmt, ...) fprintf(stderr, fmt, ##__VA_ARGS__)


int main()
{
    // wrong_eprintf("this will failed because of ','");
    correct_eprintf("this will succeed because we use ##, %s\n",
                    "hello there !");
    // opt_eprintf("with __VA_OPT__, %s\n", "Hello");
}
