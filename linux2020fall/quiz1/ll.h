#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct __node {
    int value;
    struct __node *next;
} node_t;

void add_entry(node_t **head, int new_value);
node_t *find_entry(node_t *head, int value);
void remove_entry(node_t **head, node_t *entry);
void remove_entry_non_indirect(node_t *head, node_t *entry);
void swap_pair(node_t **in_head);
void reverse(node_t **in_head);
node_t *bubble_sort(node_t *head);
void print_list(node_t *head);
