/*
  Compare two linked lists A and B
  Return 1 if they are identical and 0 if they are not.
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/

#define IDENTICAL 1
#define DIFFERENT 0

int CompareLists(Node *headA, Node* headB)
{
  while (headA && headB) {
    if (headA->data != headB->data) return DIFFERENT;

    headA = headA->next;
    headB = headB->next;
  }

  if (!headA && !headB) return IDENTICAL;

  return DIFFERENT;
}
