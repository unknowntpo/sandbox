#include <stdio.h>
#include <stdlib.h>

void dec_to_bin(int num);
int bitwise_op(int num);

int bitwise_op(int num)
{
    /* Put the bitwise_op you want to preform here */
    num |= (1 << 0);
    
    return num;
}
void dec_to_bin(int num)
{
    int a[10], i; 
    for (i = 0; num > 0; i++) {
        a[i] = num % 2;
        num = num / 2;
    }
    printf("\nBinary of Given Number is: ");
    for (i = i - 1; i >= 0; i--) {
        printf("%d ", a[i]);
    }
    printf("\n");
}

int main()
{
    int num = 10;
    //char alphabet  = 'a';
    dec_to_bin(num);
    num = bitwise_op(num);
    //alpha = alpha_op(alpha);
    dec_to_bin(num);

    return 0;
}
