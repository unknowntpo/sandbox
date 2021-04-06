#include <stdlib.h>
#include <stdio.h>

#define ARR_SIZE(arr_ptr) (sizeof(*arr_ptr))

static int cmp(const void *lhs, const void *rhs) {
    if (*(int *) lhs == *(int *) rhs)
        return 0;
    return *(int *) lhs < *(int *) rhs ? -1 : 1;
}

static int *alloc_wrapper(int a, int b, int *returnSize) {
    *returnSize = 2;
    int *res = (int *) malloc(sizeof(int) * 2);
    res[0] = a, res[1] = b;
    return res;
}

int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
    *returnSize = 2;
    int arr[numsSize][2];  /* {value, index} pair */
    for (int i = 0; i < numsSize; ++i) {
        arr[i][0] = nums[i];
        arr[i][1] = i;
    }
    qsort(arr, numsSize, sizeof(arr[0]), cmp);
    for (int i = 0, j = numsSize - 1; i < j; ) {
        if (arr[i][0] + arr[j][0] == target)
            return alloc_wrapper(arr[i][1], arr[j][1], returnSize);
        if (arr[i][0] + arr[j][0] < target)
            ++i;
        else
            --j;
    }
    *returnSize = 0;
    return NULL;
}                  

void show(int *in, int size) {
    printf("[");
    for (int i = 0; i < size; i++)
        printf("%s%d", (i == 0) ? "": " ", *in);
    printf("]\n");
}
int main() {
    int *nums = (int []){2, 7, 11, 15};
    /* What does returnSize do ? */
    int *returnSize;
    int target = 9;
    int *res;

    res = twoSum(nums, ARR_SIZE(nums), target, returnSize);
    /* want [0, 1] */
    show(res, ARR_SIZE(res));
}
