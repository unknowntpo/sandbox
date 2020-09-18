/* 
 * Implement binary search tree in c
 * Ref: https://www.programiz.com/dsa/binary-tree
 */

#include <stdio.h>
#include <stdlib.h>

/* Forward declaration technique, ref: https://hackmd.io/@sysprog/c-oop */ 
typedef struct node node_t;
struct node {
    int item;
    node_t *left;
    node_t *right; 
};

/* Inorder traversal */
void inorderTraversal(node_t *root)
{
    if (root == NULL) return;
    inorderTraversal(root->left);
    printf("%d ->", root->item);
    inorderTraversal(root->right);
}

/* Preorder tarversal */
void preorderTraversal(node_t *root)
{
    if (root == NULL) return;
    printf("%d ->", root->item);
    preorderTraversal(root->left);
    preorderTraversal(root->right);
}

/* Postorder traversal */
void postorderTraversal(node_t *root)
{
    if (root == NULL) return;
    preorderTraversal(root->left);
    preorderTraversal(root->right);
    printf("%d ->", root->item);
}

/* Create a new node */
node_t *createNode(value)
{
    node_t *newNode = malloc(sizeof(node_t));
    newNode->item = value;
    newNode->left = NULL;
    newNode->right = NULL;

    return newNode;
}

/* Insert on left of node */
node_t *insertLeft(node_t *root, int value)
{
    root->left = createNode(value);
    return root->left;
}

/* Insert on right of node */
node_t *insertRight(node_t *root, int value)
{
    root->right = createNode(value);
    return root->right;
}

int main()
{
    node_t *root = createNode(1);
    insertLeft(root, 2);
    insertRight(root, 3);
    insertLeft(root->left, 4);


printf("Inorder traversal \n");
  inorderTraversal(root);

  printf("\nPreorder traversal \n");
  preorderTraversal(root);

  printf("\nPostorder traversal \n");
  postorderTraversal(root);
  
  printf("\n");
  return 0;
}
