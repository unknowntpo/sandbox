#ifndef TWOSUM_UTIL_H
#define TWOSUM_UTIL_H

/* ARR_SIZE count the size of array */
#define ARR_SIZE(arr) (sizeof(arr))

/* show displays the array of integer of the format: [1 2 3],
 * @in: pointer to the first element of the integer array.
 * @size: size of array.
 */
void show(int *in, int size)
{
    printf("[");
    for (int i = 0; i < size; i++)
        printf("%s%d", (i == 0) ? "" : " ", *(in + i));
    printf("]\n");
}

#endif /* TWOSUM_UTIL_H */
