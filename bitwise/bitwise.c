#include <stdio.h>
#include <stdlib.h>
/* TODO: Refactor with Funtional Programming */

#define XORSWAP_UNSAFE(a, b) \
    ((a) ^= (b), (b) ^= (a), \
     (a) ^= (b)) /* Doesn't work when a and b are the same object - assigns zero \
                    (0) to the object in that case */
#define XORSWAP(a, b) \
         ((&(a) == &(b)) ? (a) \
                        : XORSWAP_UNSAFE(a, b))

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
    for (int i = 0; i < num; i++)
        printf("%d%s", num & (1 << i), (i == (num - 1)) ? "":" ");
    printf("\n");
}
void toggle_cap_small(void)
{
    char x = 'a', y = 'A';
    x ^= ' ';
    y ^= ' ';
    printf("Toggle capital small alphabet...\n");
    printf("%c\n", x);
    printf("%c\n", y);
}

int main()
{
    /* XOR swap demo */
    int a= 1, b = 4;
    
    printf("a = %d, b = %d\n", a, b);
    XORSWAP(a, b);
    printf("a = %d, b = %d\n", a, b);

    /* Bin_Dec transformation */
    int num = 10;
    //char alphabet  = 'a';
    dec_to_bin(num);
    num = bitwise_op(num);
    //alpha = alpha_op(alpha);
    dec_to_bin(num);

    toggle_cap_small();
    
    return 0;
}
