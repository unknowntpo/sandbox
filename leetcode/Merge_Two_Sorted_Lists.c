
/*
 * Ref: https://leetcode.com/problems/merge-two-sorted-lists/
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
#include <stdio.h>

void output(struct ListNode* l)
{
    /* print out the result */
    printf("Output: ");
    
    for (struct ListNode** iter; (*iter); iter = &(*iter)->next)
        printf("%d%s", (*iter)->val, ((*iter)->next) ? "->" : "");
    
    printf("\n");
}

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

/* Make the list */
void makelist(struct ListNode* l)
{

}

/* Free the list */
void freelist(struct ListNode* l)
{
    
}

/* Merge two list into one */
struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2){

    return l1;
}

int main()
{
    struct ListNode* l1, l2;

    /* Make 2 list */
    makelist(l1);
    makelist(l2);
    
    /* Merge them */
    mergeTwoLists(l1, l2);
    
    /* Print the result */
    output(l1);
    output(l2); 

    /* Free 2 list */
    freelist(l1);
    freelist(l2);
    
    /* Check if lists are freed */
    output(l1);
    output(l2);

    return 0;
}
