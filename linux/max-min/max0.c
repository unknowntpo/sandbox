#include <stdio.h>
/* MAX0 */
#define max(x, y) (x) > (y) ? (x) : (y)

/* Demonstrate problem of double evaluation */
void doOneTime()
{
    printf("called doOneTime!\n");
}

int f1()
{
    doOneTime();
    return 0;
}

int f2()
{
    doOneTime();
    return 1;
}

void demo_double_eval()
{
    int result = max(f1(), f2());
}

int main()
{
    demo_double_eval();
}
