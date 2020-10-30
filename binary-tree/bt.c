/*
 * An simple binary tree implementation in c,
 * followed the tutorial at here:
 * [Stanford - binary
 * tree](http://cslibrary.stanford.edu/110/BinaryTrees.html#s2)
 */
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct node node_t;
struct node {
    int data;
    struct node *left;
    struct node *right;
};

/*
 * Helper function that allocates a new node
 * with the given data and NULL left and right
 * pointers.
 * If failed to allocate new node, NULL pointer will be returned.
 */
node_t *newNode(int data)
{
    node_t *node = malloc(sizeof(node_t));
    if (!node) {
        perror("Malloc failed");
        return NULL;
    }
    node->data = data;
    node->left = NULL;
    node->right = NULL;

    return node;
}
/*
 * Given a binary tree, return true if a node
 * with the target data is found in the tree. Recurs
 * down the tree, chooses the left or right
 * branch by comparing the target to each node.
 */
static bool lookup(node_t *node, int target)
{
    if (!node)
        return false;
    if (target == node->data) {
        return true;
    } else if (target <= node->data) {
        return lookup(node->left, target);
    } else {
        return lookup(node->right, target);
    }
}

/*
 * Give a binary search tree and a number, inserts a new node
 * with the given number in the correct place in the tree.
 * Returns the new root pointer which the caller should
 * then use (the standard trick to avoid using reference
 * parameters).
 */
/* TODO: understand the meaning of return value */
node_t *insert(node_t *node, int data)
{
    if (!node)
        return newNode(data);
    // TODO: Deal with less and equal problem
    if (data < node->data) {
        /* FIXME: What if we simply return this insert()? */
        node->left = insert(node->left, data);
    } else if (data > node->data) {
        node->right = insert(node->right, data);
    }
    /* data == node->data, no need to insert */
    return node;
}

void printTree(node_t *node)
{
    if (!node)
        return;
    printTree(node->left);
    printf("%d ", node->data);
    printTree(node->right);
}

int main()
{
    node_t *root = newNode(5);
    root = insert(root, 3);
    root = insert(root, 6);
    root = insert(root, 4);

    printTree(root);
    puts("");

    int x = 2;
    if (!lookup(root, x)) {
        printf("%d is not exist\n", x);
    }

    return 0;
}
