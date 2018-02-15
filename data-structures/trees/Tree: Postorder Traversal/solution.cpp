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
void postOrder(node *root) {
    if (!root) return;
    postOrder(root->left);
    postOrder(root->right);
    std::cout << root->data << " ";
}
