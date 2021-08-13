// Ref:
// https://blog.gtwang.org/programming/measure-the-execution-time-in-c-language/2/
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>  // clock_gettime 函數所需之標頭檔
double pi(int n)
{
    srand(5);
    int count = 0;
    double x, y;
    for (int i = 0; i < n; ++i) {
        x = (double) rand() / RAND_MAX;
        y = (double) rand() / RAND_MAX;
        if (x * x + y * y <= 1)
            ++count;
    }
    return (double) count / n * 4;
}

struct timespec diff(struct timespec start, struct timespec end)
{
    struct timespec temp;
    if ((end.tv_nsec - start.tv_nsec) < 0) {
        temp.tv_sec = end.tv_sec - start.tv_sec - 1;
        temp.tv_nsec = 1000000000 + end.tv_nsec - start.tv_nsec;
    } else {
        temp.tv_sec = end.tv_sec - start.tv_sec;
        temp.tv_nsec = end.tv_nsec - start.tv_nsec;
    }
    return temp;
}

void count_pi()
{
    printf("pi result:\t   timelapsed (sec)\n");

    for (int i = 1; i < 1e8; i <<= 2) {
        struct timespec start, end;

        clock_gettime(CLOCK_MONOTONIC, &start);
        double result = pi(i);
        clock_gettime(CLOCK_MONOTONIC, &end);


        struct timespec temp = diff(start, end);
        double time_used;

        time_used = temp.tv_sec + (double) temp.tv_nsec / 1000000000.0;

        printf("%f\t%f\n", result, time_used);
    }
}
int main()
{
    count_pi();
}
