typedef int QElemType
#define OVERFLOW -1

enum Status{OK,ERROR}

typedef struct QNode 
{
  QElemType data;
  struct QNode *next;
}QNode, *QueuePtr;

typedef struct 
{
  QueuePtr front,rear;
}LinkQueue;

void InitQueue(LinkQueue *q)
{
  q->front = q->rear = (QueuePtr)malloc(sizeof(QNode));
  if (!q->front)
    exit(OVERFLOW);
  q->front->next = NULL;
}

void DestroyQueue(LinkQueue *q)
{
  while(q->front) {
    q->rear = q->front->next;
    free(q->front);
    q->front = q->rear;
  }
}

void EnQueue(LinkQueue *q, QElemType e)
{
  QueuePtr p = (QueuePtr)malloc(sizeof(QNode));
  if(!p)
    exit(OVERFLOW);
  p->data = e;
  p->next = NULL;
  q->rear->next = p;
  q->rear = p;
}

Status DeQueue(LinkQueue *q, QElemType *e)
{
  QueuePtr p;
  if(q->front == q->rear)
    return ERROR;
  p = q->front->next;
  *e = p->data;
  q->front->next = p->next;
  if(q->rear == p)
    q->rear = q->front;
  free(p);
  return OK;
}




