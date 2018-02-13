/*
  Reverse a linked list and return pointer to the head
  The input list will have at least one element
  Node is defined as
  struct Node
  {
    int data;
    struct Node *next;
  }
*/

Node* ReverseIterative(Node *head)
{
    Node *prev = NULL;
    Node *current = head;
    Node *next;

    while (current) {
        next = current->next;
        current->next = prev;
        prev = current;
        current = next;
    }
    return prev;
}


Node* ReverseRecursive(Node *head)
{
    if (!head || !head->next) return head;

    Node *newHead = ReverseRecursive(head->next);

    head->next->next = head;
    head->next = NULL;
    return newHead;
}


Node* Reverse(Node *head)
{
    return ReverseIterative(head);
}
