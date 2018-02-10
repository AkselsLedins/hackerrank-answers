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
  item->next = NULL;

  // keep track of head
  Node *savedHead = head;

  // null linked list
  if (!head) return item;

  // iterate to the end of the list
  while (head->next) head = head->next;

  // insert
  head->next = item;
  
  return savedHead;
}
