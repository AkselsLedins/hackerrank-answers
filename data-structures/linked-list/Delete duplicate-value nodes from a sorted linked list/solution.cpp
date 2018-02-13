/*
  Remove all duplicate elements from a sorted linked list
  Node is defined as
  struct Node
  {
    int data;
    struct Node *next;
  }
*/
Node* RemoveDuplicates(Node *head)
{
  if (!head) return NULL;

  Node *savedHead = head;

  while (head) {
    while (head && head->next && head->data == head->next->data) {
    Node *ptrToRemove = head->next;
    // we skip the duplicate node
    head->next = head->next->next;
    // without forgetting to free it
    delete ptrToRemove;
    }

    head = head->next;
  }

  return savedHead;
}
