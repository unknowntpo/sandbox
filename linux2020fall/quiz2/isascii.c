/*
 * isascii: Check if all the elements at an array is ascii character or not
 * Ref:
 *      2020q3 Homework2 (quiz2) Question 1:
 * https://hackmd.io/@sysprog/2020-quiz2 My note:
 * https://hackmd.io/@unknowntpo/quiz2
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

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <time.h>
bool is_ascii(const char str[], size_t size)
{
    if (size == 0)
        return false;
    int i = 0;
    while ((i + 8) <= size) {
        uint64_t payload;
        memcpy(&payload, str + i, 8);
        if (payload & 0x8080808080808080)  // MMM
            return false;
        i += 8;
    }
    while (i < size) {
        if (str[i] & 0x80)
            return false;
        i++;
    }
    return true;
}

bool is_ascii_one_byte(const char str[], size_t size)
{
    if (size == 0)
        return false;
    int i = 0;
    while (i < size) {
        if (str[i] & 0x80)
            return false;
        i++;
    }
    return true;
}

int main()
{
    char c[] = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of 'de Finibus Bonorum et Malorum' (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, 'Lorem ipsum dolor sit amet..', comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from 'de Finibus Bonorum et Malorum' by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.";
    size_t len = strlen(c);
    bool is_a = true;
    puts("measuring time...");

    for(int i = 0; i < 5000; i++) {
        struct timespec t1, t2, t3, t4;

        clock_gettime(CLOCK_MONOTONIC, &t1);
        is_a = is_ascii(c, len);
        clock_gettime(CLOCK_MONOTONIC, &t2);

        /* (unit: nanosecond) */
        long long time_8_byte =
            (long long) (((long long) t2.tv_sec) * 1e9 + t2.tv_nsec) -
            (long long) (((long long) t1.tv_sec) * 1e9 + t1.tv_nsec);

        clock_gettime(CLOCK_MONOTONIC, &t3);
        is_a = is_ascii_one_byte(c, len);
        clock_gettime(CLOCK_MONOTONIC, &t4);

        /* (unit: nanosecond) */
        long long time_1_byte =
            (long long) (((long long) t4.tv_sec) * 1e9 + t4.tv_nsec) -
            (long long) (((long long) t3.tv_sec) * 1e9 + t3.tv_nsec);

        printf("%i, %lld, %lld\n", i, time_8_byte, time_1_byte);
    }
    return 0;
}
