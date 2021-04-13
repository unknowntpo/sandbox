#include "ll.h"
#define DEBUG_ON 1
#define NONDIRECT 1
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

void swap_pair(node_t **in_head)  // in_head: pointer to head of the list
{
    for (node_t **node = in_head; *node && (*node)->next;
         node = &(*node)->next->next) {
        node_t *tmp = *node;
        *node = (*node)->next;
        tmp->next = (*node)->next;
        (*node)->next = tmp;
    }
    return;
}

void reverse(node_t **in_head)
{
    node_t *cursor = NULL;
    while (*in_head) {
        node_t *next = (*in_head)->next;
        (*in_head)->next = cursor;
        cursor = (*in_head);
        (*in_head) = next;
    }

    *in_head = cursor;

    return;
}

/*
 * Sort list in ascending order
 * Return if no need to sort
 */
void bubble_sort(node_t **in_head)
{
    // get size
    int size = 0;
    for (node_t *h = *in_head; h; h = h->next)
        size++;

    if (size <= 1)
        return;

    node_t **in_h = in_head;
    for (int i = 0; i < size; i++) {
        in_h = in_head;
        for (int j = 0; j < size - 1 - i; j++) {
            if ((*in_h)->value > (*in_h)->next->value) {
                node_t *tmp = (*in_h)->next;
                (*in_h)->next = tmp->next;
                tmp->next = (*in_h);
                (*in_h) = tmp;  // set head to tmp
            }
            in_h = &(*in_h)->next;
        }
#if DEBUG_ON
        print_list(*in_head);
#endif
    }
    // if no need to sort
    // return
    // set sort variable
    // sort n times, each time has n - 1 comparison

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

    puts("bubble sorting... ");
    bubble_sort(&head);
    print_list(head);
    return 0;
}
