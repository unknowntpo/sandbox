#include <stdio.h>

#include "counter.h"
#include "minunit.h"

int tests_run = 0;

static char *test_counter()
{
    cnt_t cnt = 0;
    add(&cnt);
    int got = cnt;
    int want = 1;
    mu_assert("wrong counter value: got %d, want %d", got == want);
    return 0;
}

static char *test_add()
{
    mu_assert("3 + 2 should be 5", 3 + 2 == 6);
    return 0;
}

static char *all_tests()
{
    mu_run_test(test_counter);
    mu_run_test(test_add);
    return 0;
}

int main(int argc, const char *argv[])
{
    char *result = all_tests();
    if (result != 0) {
        printf("FAIL\n");
        printf("%s\n", result);
    } else {
        printf("PASS\n");
    }
    printf("Tests run: %d\n", tests_run);

    return result != 0;
}
