#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

long int aVeryBigSum(int n, long int* ar, long int result) {
   if (n - 1 < 0) return result;
   return aVeryBigSum(n - 1, ar, result + ar[n - 1]);
}

int main() {
    int n;
    scanf("%i", &n);
    long int *ar = malloc(sizeof(long int) * n);
    for(int ar_i = 0; ar_i < n; ar_i++){
       scanf("%li",&ar[ar_i]);
    }
    long int result = aVeryBigSum(n, ar, 0);
    printf("%ld\n", result);
    return 0;
}
