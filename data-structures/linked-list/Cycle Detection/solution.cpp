/*
Detect a cycle in a linked list. Note that the head pointer may be 'NULL' if the list is empty.

A Node is defined as:
    struct Node {
        int data;
        struct Node* next;
    }
*/

// Floyd's tortoise thing
bool has_cycle(Node* head) {
  if (!head || !head->next) return false;

  Node *fastPtr = head;
  Node *slowPtr = head;

  while (fastPtr && slowPtr) {
    slowPtr = slowPtr->next;

    if (!fastPtr->next) return false;
    fastPtr = fastPtr->next->next;

    if (slowPtr == fastPtr) return true;
  }
  return false;
}
