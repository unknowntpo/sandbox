#include <stdio.h>

int adder(int a, int b)
{
    int sum, carry;

    sum = a ^ b;
    carry = (a & b) << 1;
    if (carry == 0)
        return sum;
    return adder(sum, carry);
}

int main()
{
    int a=3, b=5;
    int result = adder(a, b);
    printf("%d + %d = %d", a, b, result);
}
