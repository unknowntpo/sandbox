#include <stdio.h>

#define DEBUG_ON 1
#define MAXSIZE 30

void mergesort(int a[], int i, int j);
void merge(int a[], int i1, int j1, int i2, int j2);


void p_output(int *a, int n)
{
    printf("Output: ");
    printf("[");
    for (int i = 0; i < n; i++)
        printf("%d%s", a[i], (i < n - 1) ? " ": "");
    printf("]");
    printf("\n");
}

int main()
{
    int a[MAXSIZE], n, i;
    printf("Enter no of elements:");
    scanf("%d", &n);
    printf("Enter array elements:");

    for (i = 0; i < n; i++)
        scanf("%d", &a[i]);

    mergesort(a, 0, n - 1);

    p_output(a, n);
    return 0;
}

void mergesort(int a[], int i, int j)
{
    if (i == j)
        return;
    
    int mid = (i + j) / 2;
    
    mergesort(a, i, mid);
    mergesort(a, mid + 1, j);
    merge(a, i, mid, mid + 1, j);
#if DEBUG_ON
    printf("\ti = %d, j = %d\nafter merge", i, j);
    p_output(a, j + 1);
#endif
}
void merge(int a[], int i1, int j1, int i2, int j2)
{
    int k = 0, i = i1, j = i2;
    int temp[MAXSIZE];
#if DEBUG_ON
    printf("\tmerge: i1 = %d, j1 = %d, i2 = %d, j2 = %d\n", i1, j1, i2, j2);
#endif
    if (i1 == i2)
        return;

    while (i <= j1 && j <=j2) {
        if (a[i] < a[j])
            temp[k++] = a[i++];
        else
            temp[k++] = a[j++];

        while (i > j1) {
            while (j <= j2) 
                temp[k++] = a[j++];
            break;
        }
            
        while (j > j2) {
            while (i <= j1) 
                temp[k++] = a[i++];
            break;
        }
    }

    for (i = i1, k = 0; i <= j2; i++, k++)
        a[i] = temp[k];
    return;
}
