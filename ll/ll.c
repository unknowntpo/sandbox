/*
 * This is the basic operation following the tutorial at 
 * [Linked List Bubble Sort](http://faculty.salina.k-state.edu/tim/CMST302/study_guide/topic7/bubble.html)
 */

#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>

typedef struct list {
    int data;
    struct list *next;
} LIST;

LIST *append(LIST *, int);
LIST *sort(LIST *);
LIST *list_switch(LIST *, LIST *);
void print_list(LIST *);

int main()
{
    LIST *try;

    /* Test insertion, sort, and display result */
    try = NULL;
    try = append(try, 5);
    try = append(try, 2);
    try = append(try, 9);
    try = append(try, 8);
    try = append(try, 1);
    try = append(try, 7);
    
    printf("Original list:\n");
    print_list( try );
    try = sort( try );
    printf("Sorted list:\n");
    print_list( try );
    
    return 0;
}

/* Print lists with pointer to pointer */
void print_list(LIST *t)
{
    LIST **indir_head;
    for (indir_head = &t; *indir_head != NULL; indir_head = &(*indir_head)->next)
        printf("%d\n", (*indir_head)->data);
}

LIST *append(LIST *start, int newdata)
{
    LIST *new, *end, *ret;

    if ((new = malloc(sizeof(LIST))) == NULL) {
        fprintf(stderr, "Error: Memory allocation\n");
        exit(1);
    }

    /* Do append */
    if (start == NULL)
        ret = new;
    else {
        ret = start;
        end = start;
        while(end->next != NULL) end = end->next;
        end->next = new;
    }
    
    new->data = newdata;
    new->next = NULL;

    return ret;
}

LIST *sort( LIST *start )
{
    LIST *p, *q, *top;
    int changed = 1;

    /*
    * We need an extra item at the top of the list just to help
    * with assigning switched data to the 'next' of a previous item.
    * It (top) gets deleted after the data is sorted.
    */

    if( (top = (LIST *)malloc(sizeof(LIST))) == NULL) {
        fprintf( stderr, "Memory Allocation error.\n" );
        // In Windows, replace following with a return statement.
        exit(1);
    }

    top->next = start;
    if( start != NULL && start->next != NULL ) {
        /*
        * This is a survival technique with the variable changed.
        *
        * Variable q is always one item behind p. We need q, so
        * that we can make the assignment q->next = list_switch( ... ).
        */

        while( changed ) {
            changed = 0;
            q = top;
            p = top->next;
            while( p->next != NULL ) {
                /* push bigger items down */
                if( p->data > p->next->data ) {
                    q->next = list_switch( p, p->next );
                    changed = 1;
                }
                q = p;
                if( p->next != NULL )
                    p = p->next;
            }
        }
    }
    p = top->next;
    free( top );
    return p;
}

LIST *list_switch( LIST *l1, LIST *l2 )
{
    l1->next = l2->next;
    l2->next = l1;
    return l2;
}
