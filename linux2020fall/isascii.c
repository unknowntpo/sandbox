/* 
 * isascii: Check if all the elements at an array is ascii character or not
 * Ref:
 *      2020q3 Homework2 (quiz2) Question 1: https://hackmd.io/@sysprog/2020-quiz2
 *      My note: https://hackmd.io/@unknowntpo/quiz2
 */

/* Change to 1 if we want to test non ascii character */
#define TEST_NON_ASCII 1

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
    if (size == 0)
        return false;
    int i = 0;

    /* Check the character 8 bytes (1 word in 64-bit cpu) at a time */
    while ((i + 8) <= size) {
        uint64_t payload;
        memcpy(&payload, str + i, 8);
        if (payload & MMM) // MMM
            return false;
        i += 8;
    }
    /* 
     * Not enough words,
     * so check the character 1 byte at a time 
     */
    while (i < size) {
        if (str[i] & 0x80)
            return false;
        i++;
    }
    return true;
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
