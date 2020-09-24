#include <stdio.h>
#include <time.h>
#include <unistd.h>

int main()
{
    struct timespec start, end;
    /* Store time at beginning into struct start */
    clock_gettime(CLOCK_MONOTONIC, &start);

    /* Do some work */
    for (int i = 0; i < 10000; i++)
        ;
    /* Store time at beginning into struct end */
    clock_gettime(CLOCK_MONOTONIC, &end);

    /* Printout the time spent */
    // why using type long long ?
    long long time_spent = (long long) (end.tv_sec * 1e9 + end.tv_nsec) -
                           (long long) (end.tv_sec * 1e9 - start.tv_nsec);
    printf("time spent: %lld nanoseconds\n", time_spent);
}
