#include <stdio.h>
#include <stddef.h>



void insert_sort(int *ary, size_t len)
{
  size_t i, j;
  int temp;
  if (len == 0 || len == 1) return;

  for (i = 1; i < len; i++) {
    temp = ary[i];
    for (j = i - 1; j >= 0 && ary[j] > temp; j--) {
      ary[j + 1] = ary[j];  
    }
    ary[j + 1] = temp;
  }
}
   


int main(void)
{
  int i = 0;
  int ary[] = {9,4,7,12,56,7,4,3};
  size_t len = sizeof(ary) / sizeof(int);
  for (; i < 8; i++) {
    printf("%d ", ary[i]);
  }

  insert_sort(ary, len);
  printf("\n"); 
  for (i = 0; i < 8; i++) {
    printf("%d ", ary[i]);
  }
  printf("\n");
  return 0;
}


