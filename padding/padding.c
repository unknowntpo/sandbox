
#include <time.h>
int main()
{
    struct Foo {
        char x;
        short y;
        int z;
    };
    struct Foo foo;
    clock_gettime(CLOCK, &start);
    for (unsigned long i = 0; i < RUNS; ++i) {
        foo.z = 1;
        foo.z += 1;
    }
    clock_gettime(CLOCK, &end);
    struct Bar {
        char x;
        short y;
        int z;
    } __attribute__((packed));
    struct Bar bar;
    clock_gettime(CLOCK, &start);
    for (unsigned long i = 0; i < RUNS; ++i) {
        bar.z = 1;
        bar.z += 1;
    }
    clock_gettime(CLOCK, &end);
}
