/* test with pointer to pointer and linked list */

/* TODO: Write the function description */
/* TODO: Functional programming paradiam
 * Ref: Functional Programming 風格的 C 語言實作
 */

/*
 * check if q is null pointer, if q is NULL, return 1, if q is not NULL,
 * return 0.
 */
#include "ll.h"
#define NULL_GUARD(q) \
    do {              \
        if (!q)       \
            return 1; \
    } while (0)

/* err check: if err = 0 means no error, if err == 1 means error occured. */
#define ERR_CHECK(err)               \
    do {                             \
        if (err)                     \
            perror("error occured"); \
    } while (0)

bool err = 0;

bool q_show(queue_t *q)
{
    if (!q) {
        printf("[]\n");
        return 0;
    }
    list_ele_t **p = &q->head;
    printf("[");
    while (*p) {
        printf("%d ", (*p)->val);
        p = &(*p)->next;
    }
    printf("]");

    printf("\n");
    return 0;
}
bool q_insert(queue_t *q, int val)
{
    NULL_GUARD(q);

    /* build new node */
    list_ele_t *newh = malloc(sizeof(list_ele_t));
    NULL_GUARD(newh);

    newh->val = val;

    /* update new head pos */
    newh->next = q->head;
    q->head = newh;

    if (q->size <= 1)
        q->tail = newh;

    q->size++;

    return 0;
}
queue_t *q_new()
{
    queue_t *q = malloc(sizeof(queue_t));
    q->head = NULL;
    q->tail = NULL;
    q->size = 0;
    return q;
}

bool q_free(queue_t *q)
{
    if (!q)
        return 0;
    for (list_ele_t *tmp = q->head; tmp; tmp = q->head) {
        q->head = q->head->next;
        free(tmp);
    }

    q->head = NULL;
    q->tail = NULL;
    q->size = 0;
    free(q);

    return 0;
}

int main()
{
    queue_t *q = q_new();
    NULL_GUARD(q);
    q_insert(q, 1);
    q_insert(q, 2);
    err = q_show(q);
    // if error, set err to 1
    ERR_CHECK(err);

    err = q_free(q);
    ERR_CHECK(err);
    err = q_show(q);
    ERR_CHECK(err);
    /* demo of the usage of report.c */

    return 0;
}
