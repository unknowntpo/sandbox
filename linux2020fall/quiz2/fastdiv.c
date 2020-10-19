/*
 * This program is a verification of RinHizakura's test program of NCKU
 * linux2020 fall quiz2
 * - [RinHizakura
 * NCKU-System-Software-Quiz2](https://github.com/RinHizakura/NCKU-System-Software-Quiz2/blob/master/fastdiv.c)
 * - [Quiz2](https://hackmd.io/@sysprog/2020-quiz2#%E6%B8%AC%E9%A9%97-3)
 * - [My Note of
 * Quiz2](https://hackmd.io/@unknowntpo/quiz2#%E6%B8%AC%E9%A9%97-3)
 * - [jemalloc div.c](https://github.com/jemalloc/jemalloc/blob/dev/src/div.c)
 */


#include <assert.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

/* Divisor */
const uint32_t D = 3;
#define M ((uint64_t)(UINT64_C(0xFFFFFFFFFFFFFFFF) / (D) + 1))

typedef struct div_info_s div_info_t;
struct div_info_s {
    uint32_t magic;
};

void div_init(div_info_t *div_info, size_t d)
{
    /* What does the special case of d ? */
    assert(d != 0);
    assert(d != 1);

    /* What does this line do? */
    uint64_t two_to_k = ((uint64_t) 1 << 32);
    /* What does magic represent? */
    uint32_t magic = (uint32_t)(two_to_k / d);

    /*  What does this line do? */
    if (two_to_k % d != 0) {
        magic++;
    }
    div_info->magic = magic;
}
/* Why use static inline?  What happeded when not adding storage-class
 * specifiers? */
static inline size_t div_compute(div_info_t *div_info, size_t n)
{
    assert(n <= (uint32_t) -1);
    size_t i = ((uint64_t) n * (uint64_t) div_info->magic) >> 32;
    return i;
}

uint32_t quickdiv(uint32_t n)
{
    uint64_t quotient = ((__uint128_t) M * n) >> 64;
    return quotient;
}

uint32_t quickmod(uint32_t n)
{
    uint64_t quotient = ((__uint128_t) M * n) >> 64;
    return n - quotient * D;
}

bool divisible(uint32_t n)
{
    printf("%lx %lx\n", (n * M), (M - 1));
    return n * M <= M - 1;
}

void test()
{
    struct timespec tt1, tt2;
    clock_gettime(CLOCK_MONOTONIC, &tt1);
    for (uint64_t i = 2; i < 5000; i++) {
        div_info_t DIV;
        div_init(&DIV, i);

        uint64_t num = rand() % 10000;
        num *= i;

        clock_gettime(CLOCK_MONOTONIC, &tt1);
        size_t ans1 = div_compute(&DIV, num);
        clock_gettime(CLOCK_MONOTONIC, &tt2);

        long long time1 = (long long) (tt2.tv_sec * 1e9 + tt2.tv_nsec) -
                          (long long) (tt1.tv_sec * 1e9 + tt1.tv_nsec);


        clock_gettime(CLOCK_MONOTONIC, &tt1);
        size_t ans2 = num / i;
        clock_gettime(CLOCK_MONOTONIC, &tt2);

        long long time2 = (long long) (tt2.tv_sec * 1e9 + tt2.tv_nsec) -
                          (long long) (tt1.tv_sec * 1e9 + tt1.tv_nsec);

        printf("%ld, %lld, %lld, %ld, %ld\n", i, time1, time2, ans1, ans2);
    }

}
int main()
{
    printf("quickmod(5) = %d\n", quickmod(5));
    //test();
}
