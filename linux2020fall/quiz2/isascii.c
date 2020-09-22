/* 
 * isascii: Check if all the elements at an array is ascii character or not
 * Ref:
 *      2020q3 Homework2 (quiz2) Question 1: https://hackmd.io/@sysprog/2020-quiz2
 *      My note: https://hackmd.io/@unknowntpo/quiz2
 */

/* Change to 1 if we want to test non ascii character */
#define TEST_NON_ASCII 0

/*

(a) 0x0080008000800080
(b) 0x8000800080008000
(c) 0x0808080808080808
(d) 0x8080808080808080  <-- right answer
(e) 0x8888888888888888
(f) 0xFFFFFFFFFFFFFFFF

*/
/* Define MMM */
#define MMM 0x8080808080808080

#include <stdio.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <string.h>
bool is_ascii(const char str[], size_t size)
{
    uint64_t payload;
    size_t i;
    if (size == 0) return false;
    for (i = 0; (i + 8) < size; i = i + 8) {
        memcpy(&payload, str + i, 8);
        if ((payload & 0x8080808080808080) > 0)
            return false;
    }

    while (i < size) {
        if ((str[i] & 0x80) > 0)
            return false;
        i++;
    }

    return true;
    // copy ele in str into payload using memcpy
    // use payload & 0x80808080 to check if there's a non ascii byte in str[]
    // if not enough element to copy to payload
    // check byte by byte
}
int main()
{
    char c[] = "12345678";
    /* Insert non ascii character */
#if TEST_NON_ASCII
    c[1] = 128;
#endif
    printf("is %s ascii ? %s\n", c, is_ascii(c, sizeof(c)) ? "true" : "false");

    return 0;
}
