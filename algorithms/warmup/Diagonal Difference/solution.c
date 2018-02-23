#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int diagonal2(int n1, int n2, int n, int result) {
    if (n1 > n) return result;
    return result;
}

void absoluteDifferenceOfDiagonals(int n, int **a) {
    int result = 0;

    int diagonal1 = 0;
    int diagonal2 = 0;

    for (int i = 0; i < n; i++) {
        diagonal1 += a[i][i];
        diagonal2 += a[n - i - 1][i];
    }

    printf("%d", abs(diagonal2 - diagonal1));
}

int main(){
    int n;
    scanf("%d",&n);
    int **a;

    a = (int **)malloc(sizeof(int *) * (n));
    for (int i = 0; i < n; i++) {
        a[i] = (int *)malloc(sizeof(int) * (n));
    }

    for(int a_i = 0; a_i < n; a_i++){
       for(int a_j = 0; a_j < n; a_j++){
          scanf("%d",&a[a_i][a_j]);
       }
    }
    absoluteDifferenceOfDiagonals(n, a);

    // free ;(

    return 0;
}
