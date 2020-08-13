
/*
Ref: https://leetcode.com/problems/merge-sorted-array/
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/
#include <stdio.h>

#define XOR_SWAP_UNSAFE(a, b)  \
    ((a) ^= (b), (b) ^= (a),   \
     (a) ^= (b))

#define XOR_SWAP(a, b)  \
    (((&a) == &(b)) ? (a)  \
                   : XOR_SWAP_UNSAFE(a, b))

void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n){
    int i, j;
    for (i = 0; i < nums2Size; i++) {
        for (j = 0; j < m; j++) {
            if (nums2[i] < nums1[j])
                XOR_SWAP(nums2[i], nums1[j]);
                // swap(&nums2[i], &nums1[j])
        }
        if (j == m) {
            XOR_SWAP(nums2[i], nums1[j]);
            m++;
        }
    }
}

void output(int* nums1, int nums1Size)
{
    /* print out the result */
    printf("Output: [");
    for (int i = 0; i < nums1Size; i++) {
        printf("%d%s", *nums1, (i < nums1Size - 1) ? "," : "");
        nums1++;
    }
    printf("]");
    printf("\n");
}

int main()
{
    int nums1[] = {1,2,3,0,0,0};
    int nums2[] = {2,5,6};
    int m = 3;
    int n = 3;
    int nums1Size = 6;
    int nums2Size = 3;

    merge(nums1, nums1Size, m, nums2, nums2Size, n);
    output(nums1, nums1Size);

    return 0;
}
