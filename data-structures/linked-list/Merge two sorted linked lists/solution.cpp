/*
  Merge two sorted lists A and B as one linked list
  Node is defined as
  struct Node
  {
    int data;
    struct Node *next;
  }
*/

/* ************************************************************************** */
/*                        ITERATIVE                                           */
/* ************************************************************************** */

Node* SortingFunctions(Node *a, Node *b) {
  // exceptional cases
  if (!a) return b;
  if (!b) return a;

  return a->data <= b->data ? a : b;
}

void Iterate(Node **nodeWithSmallestValue, Node **a, Node **b) {
  // a node
  if (*nodeWithSmallestValue == *a) {
    *nodeWithSmallestValue = *a;
    *a = (*a)->next;
    return;
  }
  // b node
  *nodeWithSmallestValue = *b;
  *b = (*b)->next;
};

Node* MergeListsIterative(Node *headA, Node *headB)
{
  // exceptional cases
  if (!headA) return headB;
  if (!headB) return headA;

  // keep track of head
  Node *head = SortingFunctions(headA, headB);

  // we will insert there the nodes
  Node *newListIterator = head;

  while (headA && headB) {
    Node* nodeWithSmallestValue;

    // we get the node we have to insert
    nodeWithSmallestValue = SortingFunctions(headA, headB);
    // and we iterate the corresponding list
    Iterate(&nodeWithSmallestValue, &headA, &headB);

    // we skip the first loop
    if (head == nodeWithSmallestValue) continue;

    newListIterator->next = nodeWithSmallestValue;
    newListIterator = nodeWithSmallestValue;
  }

  newListIterator->next = headA == NULL ? headB : headA;

  return head;
}


/* ************************************************************************** */
/*                        Recursive                                           */
/* ************************************************************************** */

Node* MergeLists(Node *headA, Node *headB)
{
  // exceptional cases
  if (!headA) return headB;
  if (!headB) return headA;

  if (headA->data < headB->data) {
    MergeLists(headA->next, headB)
  } else {
    MergeLists(headA, headB->next)
  }

  return head;
}
