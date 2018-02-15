/*The tree node has data, left child and right child
class Node {
    int data;
    Node* left;
    Node* right;
};

*/

int height(Node* root) {
    if (!root) return -1;

    const int leftHeight = height(root->left);
    const int rightHeight = height(root->right);

    const int biggerOne = leftHeight > rightHeight ? leftHeight : rightHeight;

    return biggerOne + 1;
}
