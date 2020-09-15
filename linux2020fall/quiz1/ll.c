#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#define NONDIRECT 1
typedef struct __node {
    int value;
    struct __node *next;
} node_t;

void add_entry(node_t **head, int new_value)
{
    node_t **indirect = head;

    node_t *new_node = malloc(sizeof(node_t));
    new_node->value = new_value;
    new_node->next = NULL;

    assert(new_node);

    /* traverse to the end */
    while (*indirect)
        indirect = &(*indirect)->next;
    *indirect = new_node;

    return;
}

node_t *find_entry(node_t *head, int value)
{
    node_t *current = head;
    for (; current && (current->value != value); current = current->next)
        /* Iterate */;
    return current;
}


void remove_entry(node_t **head, node_t *entry)
{
    node_t **indirect = head;

    while ((*indirect) != entry)
        indirect = &(*indirect)->next;

    *indirect = entry->next;
    free(entry);
}

void remove_entry_non_indirect(node_t *head, node_t *entry)
{
    node_t *cur = head;
    node_t *pre = NULL;

    while (cur != entry) {
        pre = cur;
        cur = cur->next;
    }
    /* We found entry here */
    if (!pre)
        head = cur->next;
    pre->next = cur->next;
}

void swap_pair(node_t **indirect_head)
{
    for (node_t **node = indirect_head; *node && (*node)->next;
         node = &(*node)->next->next) {
        node_t *tmp = *node;
        *node = (*node)->next;
        tmp->next = (*node)->next;
        (*node)->next = tmp;
    }
    return;
}

void reverse(node_t **indirect_head)
{
    node_t *cursor = NULL;
    while (*indirect_head) {
        node_t *next = (*indirect_head)->next;
        (*indirect_head)->next = cursor;
        cursor = (*indirect_head);
        (*indirect_head) = next;
    }

    *indirect_head = cursor;

    return;
}

void print_list(node_t *head)
{
    for (node_t *current = head; current; current = current->next)
        printf("%d ", current->value);
    printf("\n");
}

int main(int argc, char const *argv[])
{
    node_t *head = NULL;

    print_list(head);

    puts("adding entry...");
    add_entry(&head, 72);
    add_entry(&head, 101);
    add_entry(&head, 108);
    add_entry(&head, 109);
    add_entry(&head, 110);
    add_entry(&head, 111);

    print_list(head);

    puts("removing entry...");
    node_t *entry = find_entry(head, 101);
#if NONDIRECT
    remove_entry_non_indirect(head, entry);
#else
    remove_entry(&head, entry);
#endif
    print_list(head);

    puts("removing entry...");
    entry = find_entry(head, 111);
#if NONDIRECT
    remove_entry_non_indirect(head, entry);
#else
    remove_entry(&head, entry);
#endif
    print_list(head);
    /* swap pair.
     * See https://leetcode.com/problems/swap-nodes-in-pairs/
     */
    puts("swaping pair...");
    swap_pair(&head);
    print_list(head);
    puts("reversing... ");
    reverse(&head);
    print_list(head);

    return 0;
}
