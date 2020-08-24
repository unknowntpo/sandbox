#include <stdio.h>

#define DEBUG_MODE 1
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
    // end condition
    if (i == j)
        return;
    int mid = (i + j) / 2;
    
    mergesort(a, 0, mid);
    mergesort(a, mid + 1, j);
    merge(a, 0, mid, mid + 1, j);
    return;
}

void merge(int a[], int i1, int j1, int i2, int j2)
{
    /* TODO: implement this with leetcode merge sorted array */
    // actually, it's merge sorted array
    int i, j, k;
    int temp[MAXSIZE];
    
    if (i1 == i2)
        return;

    for (k = 0, i = i1, j = i2; (i <= j1) && (j <= j2) ; k++) {
        if (a[i] < a[j]) {
            temp[k] = a[i];
            i++;
        }
        else {
            temp[k] = a[j];
            j++;    
        }
        if (i > j1) {
            // copy rest of ele to temp, range from j to j2 
            for (; j <= j2 ; j++, k++)
                temp[k + 1] = a[j];
            break;
        }
        if (j > j2) {
            // copy rest of ele to temp, range from i to j1
            for (; i <= j1 ; i++, k++)
                temp[k + 1] = a[i];
            break;
        }
    }
    
    /* copy element in temp[] back to a[] */
    for (i = 0; i <= k; i++)
        a[i] = temp[i];
    return;
}
