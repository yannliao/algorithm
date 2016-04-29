#include <stdio.h>
#include <stddef.h>

void shell_sort(int *ary, size_t len)
{ 
  size_t inc, i, j;
  int temp;
  if (len == 0 || len == 1) return;
  inc = len >> 1;

  while (inc > 0) {
    for (i = inc; i < len; i++) {
      temp = ary[i];
      for (j = i; j >= inc && ary[j - inc] > temp; j -= inc) {
        ary[j] = ary[j - inc];
      }
      ary[j] = temp;
    }
    inc >>= 1;
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

  shell_sort(ary, len);
  printf("\n"); 
  for (i = 0; i < 8; i++) {
    printf("%d ", ary[i]);
  }
  printf("\n");
  return 0;
}


