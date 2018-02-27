/* Node is defined as :
typedef struct node
{
    int val;
    struct node* left;
    struct node* right;
    int ht;
} node; */

int height(node *n) {
    if (!n) return -1;
    return n->ht;
}

int getBalanceFactor(node *n) {
    return height(n->left) - height(n->right);
}

node *buildNewNode(int val) {
    node *n = new node();

    n->val   = val;
    n->left   = NULL;
    n->right  = NULL;
    n->ht = 0;

    return(n);
}

int max(int v1, int v2) {
    return v1 > v2 ? v1 : v2;
}

node* rightRotate(node *y) {
    node *x = y->left;
    node *T2 = x->right;

    x->right = y;
    y->left = T2;

    y->ht = max(height(y->left), height(y->right)) + 1;
    x->ht = max(height(x->left), height(x->right)) + 1;

    return x;
}

node *leftRotate(node *x) {
    node *y = x->right;
    node* T2 = y->left;

    y->left = x;
    x->right = T2;

    x->ht = max(height(x->left), height(x->right)) + 1;
    y->ht = max(height(y->left), height(y->right)) + 1;

    return y;
}

node * insert(node * root, int val)
{
    if (!root) return buildNewNode(val);

    if (val < root->val) {
        root->left = insert(root->left, val);
    }
    else if (val > root->val) {
        root->right = insert(root->right, val);
    }
    else {
        return root;
    }

    root->ht = 1 + max(height(root->left), height(root->right));

    int balance = getBalanceFactor(root);

    if (balance > 1 && val < root->left->val) {
        return rightRotate(root);
    }
    if (balance < -1 && val > root->right->val) {
        return leftRotate(root);
    }

    if (balance > 1 && val > root->left->val) {
        root->left = leftRotate(root->left);
        return rightRotate(root);
    }
    if (balance < -1 && val > root->right->val) {
        root->right = rightRotate(root->right);
        return leftRotate(root);
    }

    return root;
}
