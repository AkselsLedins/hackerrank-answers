/*
  Insert Node at the end of a linked list
  head pointer input could be NULL as well for empty list
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
Node* InsertNth(Node *head, int data, int position)
{
  // create new node
  Node *item = new Node();
  item->data = data;
  item->next = NULL;

  if (!head || position == 0) {
      item->next = head;
      return item;
  }

  Node *ptr = head;
  for (int currentPosition = 1; ptr->next && currentPosition < position; ++currentPosition) {
      ptr = ptr->next;
  }

  Node *nextNode = ptr->next;

  item->next = nextNode;
  ptr->next = item;

  return head;
}
