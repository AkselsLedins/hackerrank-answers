/*
  Get Nth element from the end in a linked list of integers
  Number of elements in the list will always be greater than N.
  Node is defined as
  struct Node
  {
    int data;
    struct Node *next;
  }
*/
int GetNode(Node *head,int positionFromTail)
{

  int i = 0;
  Node* result = head;

  while (head->next) {
    // we create a gap of positionFromTail between the two pointers
    if (i < positionFromTail) {
    i += 1;
    } else {
    result = result->next;
    }

    head = head->next;
  }

  // when we arrived a the tail of the chained list we ha ve a gap of
  // positionFromTail from the end with the result pointer
  return result->data;
}
