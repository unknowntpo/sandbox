#include <stdio.h>
#include <stdlib.h>
int adder(int a, int b)
{
    int sum = a ^ b;
    int carry = a & b;
    if (carry == 0)
        return sum;
    adder(sum, carry << 1);
}

int main(int argc, char **argv)
{
    if (argc != 3) {
        printf("Usage: ./adder a b\n");
        return -1;
    }
    int a=atoi(argv[1]);
    int b=atoi(argv[2]);
    int result = adder(a, b);
    printf("%d + %d = %d", a, b, result);
}
