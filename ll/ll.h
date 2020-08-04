#ifndef LL_H
#define LL_H

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct ELE {
    int val;
    struct ELE *next;
} list_ele_t;
typedef struct {
    list_ele_t *head;
    list_ele_t *tail;
    int size;
} queue_t;

bool q_show(queue_t *q);
bool q_insert(queue_t *q, int val);
queue_t *q_new();
bool q_free(queue_t *q);

#endif /* LL_H */
