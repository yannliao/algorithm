#include <stddef.h>
#include <assert.h>
#include <stdio.h>
#include <stdlib.h>


typedef struct BTree
{
  int data;
  struct BTree *parent;
  struct BTree *lchild;
  struct BTree *rchild;
}BTree, *BTreePtr;


typedef int ElemType;

BTreePtr InsertBST(BTreePtr t, ElemType data)
{

  if (!t) {
    t = (BTreePtr)malloc(sizeof(BTree));
    t->data = data;
    t->lchild = t->rchild = NULL;
    return t;
  } else if (data == t->data) 
    return t;
  if (data < t->data)
    t->lchild = InsertBST(t->lchild, data);
  else
    t->rchild = InsertBST(t->rchild, data);
  return t;
}


int SearchBST(BTreePtr t, BTreePtr *p, ElemType data)
{
  if (!t) {
    return 0;
  } else if (data == t->data) {
    *p = t;
    return 1;
  } else if (data < t->data) {
    return SearchBST(t->lchild, p, data);
  } else {
    return SearchBST(t->rchild, p, data);
  }
}


int BSTPrint(BTreePtr t)
{
  if(!t)
    return 0;
  BSTPrint(t->lchild);
  printf(" %d \n", t->data);
  BSTPrint(t->rchild);
  return 1;
}

int Delete(BTreePtr t) 
{
  BTreePtr q, s;
  if (!t->lchild && !t->rchild)
    free(t);
  else if (!t->lchild) {
    q = t->rchild;
    t->data = t->rchild->data;
    t->lchild = t->rchild->lchild;
    t->rchild = t->rchild->rchild;
    free(q);
  } else if (!t->rchild) {
    q = t->lchild;
    t->data = t->lchild->data;
    t->lchild = t->lchild->lchild;
    t->rchild = t->lchild->rchild;
    free(q);
  } else {
    q = t;
    s = t->lchild;
    while (s->rchild) {
      q = s;
      s = s->rchild;
    }
    t->data = s->data;
    if (q != t)
      q->rchild = s->lchild;
    else 
      q->lchild = s->lchild;
    free(s);
  }
  return 1;  
}

int DeleteBST(BTreePtr t, ElemType data)
{
  if (!t) {
    return 0;
  } else {
    if (data == t->data) 
      return Delete(t);
    else if (data < t->data)
      return DeleteBST(t->lchild, data);
    else 
      return DeleteBST(t->rchild, data);
  }

}
int main (void) 
{
  BTreePtr p;
  BTreePtr t;
  int i,j;
  t = InsertBST(t, 9);
  t = InsertBST(t, 10);
  t = InsertBST(t, 6);
  t = InsertBST(t, 8);
  t = InsertBST(t, 3);
  t = InsertBST(t, 5);
  t = InsertBST(t, 4);
  t = InsertBST(t,7);

  i = BSTPrint(t);
  assert(i == 1);

  printf("\n******************\n");
  DeleteBST(t, 6);
  i = BSTPrint(t); 
}

