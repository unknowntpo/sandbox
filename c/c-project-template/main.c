#include <stdio.h>

#include "counter.h"
#include "minunit.h"

#define PASS NULL

int tests_run = 0;

static char *test_counter()
{
    cnt_t cnt = 0;
    add(&cnt);
    int got = cnt;
    int want = 1;

    // TODO: What if we need to provide additional argument ?
    // mu_assert(got == want, "wrong counter value got %d, want %d", got, want);
    mu_assert(got == want, "wrong counter value");
    return PASS;
}

static char *test_add()
{
    mu_assert(3 + 2 == 6, "3 + 2 should be 5");
    return PASS;
}

static char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_counter);
    mu_run_test(test_add);

    return PASS;
}

RUN_TESTS(all_tests);
