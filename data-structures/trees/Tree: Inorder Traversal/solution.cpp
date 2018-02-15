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
void inOrderRecursive(node *root) {
    if (!root) return;
    inOrderRecursive(root->left);
    std::cout << root->data << " ";
    inOrderRecursive(root->right);
}


/* ************************************************************************** */
/*                        Iterative (Morris traversal)                        */
/* https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion-and-without-stack/ */
/* ************************************************************************** */
void inOrder(node *root) {
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
                current = current->left;
            }
            else {
                pre->right = NULL;
                std::cout << current->data << " ";
                current = current->right;
            }
        }
    }
}
