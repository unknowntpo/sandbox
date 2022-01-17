#include <stdio.h>
#include "munit.h"

int add(int a, int b)
{
    return a + b;
}

void test_adder()
{
    mu_assert("3+2 are 5", add(3, 2), 5);

    mu_assert("3+2 not equals to 6", add(3, 2), 6);
}

int main()
{
    test_adder();
}
