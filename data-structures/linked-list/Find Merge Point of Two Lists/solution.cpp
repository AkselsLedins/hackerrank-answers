/*
    Find merge point of two linked lists
    Node is defined as
    struct Node
    {
        int data;
        Node* next;
    }
*/

int FindMergeNode(Node *headA, Node *headB) {
  Node *a = headA;
  Node *b = headB;

  // at some point the pointers will collide
  while (a != b) {
    if (!a->next) {
    a = headB;
    } else {
    a = a->next;
    }

    if (!b->next) {
    b = headA;
    } else {
    b = b->next;
    }
  }
  return a->data;
}
