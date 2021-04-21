#include <stdio.h>
#define SIZE 2
typedef int array[SIZE][SIZE];

void printarray(array a, int size) {
    for (int i = 0; i < SIZE; i++)
        for (int j = 0; j < SIZE; j++)
            printf("a[%d][%d] = %d\n", i, j, a[i][j]);
}
int main() {
    int a[SIZE][SIZE] = {{0, 1}, {2, 3}};
    printarray(a, SIZE);

}
