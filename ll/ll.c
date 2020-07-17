/* test with pointer to pointer and linked list */

/* TODO: deal with error check macro, debug macro? */
/* TODO: Write the function description */
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#define NULL_GUARD(q) \
    do {              \
        if (!q)       \
            return 1; \
    } while (0)
#define ERR_CHECK(err)\
    do { \
        if (err) \
            perror("error occured"); \
    } while (0)

bool err = 0;
typedef struct ELE {
    int val;
    struct ELE *next;
} list_ele_t;
typedef struct {
    list_ele_t *head;
    list_ele_t *tail;
    int size;
} queue_t;


bool q_show(queue_t *q)
{
    list_ele_t **p = &q->head;
    while (*p) {
        printf("%d ", (*p)->val);
        p = &(*p)->next;
    }
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
int main()
{
    queue_t *q = q_new();
    NULL_GUARD(q);
    q_insert(q, 1);
    err = q_show(q);
    // if error, set err to 1
    ERR_CHECK(err);
    return 0;
}
