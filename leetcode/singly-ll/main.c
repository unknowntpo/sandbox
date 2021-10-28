#include <stdio.h>
#include <stdlib.h>

typedef struct node node_t;
typedef struct list list_t;

// Forward declaration.
struct node {
    int val;
    struct node *next;
};

struct list {
    node_t *head;
    node_t *tail;
};

list_t *new_list()
{
    list_t *l = malloc(sizeof(list_t));
    if (l == NULL) {
        perror("failed to malloc a new list");
        return NULL;
    }

    node_t *dummy = malloc(sizeof(node_t));
    if (dummy == NULL) {
        perror("failed to malloc a new dummy");
        return NULL;
    }

    dummy->next = NULL;

    l->head = dummy;
    l->tail = dummy;

    return l;
}

int main()
{
    list_t *l = new_list();
    if (l == NULL) {
        perror("failed to create a new list");
        exit(1);
    }
    printf("good!\n");
}
