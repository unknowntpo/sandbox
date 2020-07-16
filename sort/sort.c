/* Create by unknowntpo at 2020.06.12
 * @Description:
 *      sort integer array with different method
 *             - Bubble sort
 *             - Merge sort
 *             - Selection sort
 */
#include "sort.h"


void swap(int *px, int *py)
{
    int tmp;
    tmp = *px;
    *px = *py;
    *py = tmp;
}

int *sort_bubble(int *a, int size)
{
    int i, j;
    for (i = 0; i < size - 1; i++) {
        for (j = 0; j < size - 1; j++) {
            if (a[j] > a[j+1])
                swap(&a[j], &a[j+1]);
        }   
        display(a, size);
    }


    return a;
}

void display(int *p, int size)
{
    /* Display the array */
    for(int i = 0; i < size; i++)
        printf("%d", *p++);    
    printf("\n");
}

int main()
{
    int a[] = {8, 7, 6, 5, 4, 3, 2, 1};
    //int a[] = {3, 1, 6, 5, 4, 2, 8, 7};
    int size = 8;
    int *p = a;
    /* TODO: use callback function to apply more sort method */
    p = sort_bubble(a, size);

    display(p, size);

}
