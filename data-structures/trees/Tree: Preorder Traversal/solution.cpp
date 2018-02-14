/* you only have to complete the function given below.
Node is defined as

struct node
{
    int data;
    node* left;
    node* right;
};

*/


/* ************************************************************************** */
/*                        Recursive                                           */
/* ************************************************************************** */
void preOrderRecursive(node *root) {
    if (!root) return;
    std::cout << root->data << " ";
    preOrderRecursive(root->left);
    preOrderRecursive(root->right);
}


/* ************************************************************************** */
/*                        Iterative (Morris traversal)                        */
/* https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion-and-without-stack/ */
/* ************************************************************************** */
void preOrder(node *root) {
    if (!root) return;

    node *current = root;
    node *pre = NULL;
    while (current) {
        if (!current->left) {
            std::cout << current->data << " ";
            current = current->right;
        }
        else {
            pre = current->left;
            while (pre->right && pre->right != current) pre = pre->right;

            if (!pre->right) {
                pre->right = current;
                std::cout << current->data << " ";
                current = current->left;
            }
            else {
                pre->right = NULL;
                current = current->right;
            }
        }
    }
}
