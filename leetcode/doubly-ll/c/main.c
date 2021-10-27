#include <stdio.h>

typedef struct node {
    int val;
    struct node *prev, *next;
} node_t;

typedef struct List {
    node_t *front, *back;
} list_t;

int main()
{
    printf("hello");
}
