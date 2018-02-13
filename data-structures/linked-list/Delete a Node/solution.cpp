/*
  Delete Node at a given position in a linked list
  Node is defined as
  struct Node
  {
    int data;
    struct Node *next;
  }
*/
Node* Delete(Node *head, int position)
{
    if (!head) return NULL;

    // specific behaviour for the first position;
    if (position == 0) {
        Node *nextNode = head->next;
        delete(head);
        return nextNode;
    }

    // iterate to the desired position - 1
    Node *ptr = head;
    for (int i = 0; i < position - 1; ++i) ptr = ptr->next;

    // remove the node from the list
    Node *nodeToFree = ptr->next;
    Node *nodeToLink = ptr->next->next;
    ptr->next = nodeToLink;

    // dont forget to free
    delete(nodeToFree);

    return head;
}
