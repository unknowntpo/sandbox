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
        printf("%d%s", a[i], (i < n - 1) ? " " : "");
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
}
void merge(int a[], int i1, int j1, int i2, int j2)
{
    if (i1 == i2)
        return;

    int i = i1, j = i2, k = 0;
    int temp[MAXSIZE];

    while (i <= j1 && j <= j2) {
        if (a[i] <= a[j])
            temp[k++] = a[i++];
        else
            temp[k++] = a[j++];
    }

    /* copy rest element */
    while (j <= j2)  // i has exceeded j1
        temp[k++] = a[j++];
    while (i <= j1)  // j has exceeded j2
        temp[k++] = a[i++];

    /* copy element back to a[] */
    for (i = i1, k = 0; i <= j2; i++, k++)
        a[i] = temp[k];

    return;
}
