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
Node* Insert(Node *head,int data)
{
  // create new node
  Node *item = new Node();
  item->data = data;
  // link it to the rest of the list
  item->next = head;

  return item;
}
