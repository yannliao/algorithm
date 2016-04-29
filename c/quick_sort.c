#include <stdio.h>
#include <stddef.h>

void swap(int *x, int *y)
{
  int temp;
  temp = *x;
  *x = *y;
  *y = temp;
}

size_t partition(int *ary, size_t len, size_t pivot_i)
{
  size_t i = 0;
  size_t store_i = pivot_i;
  int pivot = ary[pivot_i];

  swap(&ary[pivot_i], &ary[pivot_i + (len - 1)]);
  for (; i < len; i++) {
    if (ary[pivot_i + i] < pivot) {
      swap(&ary[pivot_i + i], &ary[store_i]);
      store_i++;
    }
  }
  swap(&ary[pivot_i + (len - 1)], &ary[store_i]);
  return store_i;
}


void quick_sort(int *ary, size_t len)
{
  if(len == 0 || len ==1) return;
  size_t new_i = partition(ary, len, 0);
  quick_sort(ary, new_i);
  quick_sort(&ary[new_i + 1], len - new_i - 1);
}



int main(void)
{
  int i = 0;
  int ary[] = {9,4,7,12,56,7,4,3};
  size_t len = sizeof(ary) / sizeof(int);
  for (; i < 8; i++) {
    printf("%d ", ary[i]);
  }

  quick_sort(ary, len);
  printf("\n"); 
  for (i = 0; i < 8; i++) {
    printf("%d ", ary[i]);
  }
  printf("\n");
  return 0;
}


