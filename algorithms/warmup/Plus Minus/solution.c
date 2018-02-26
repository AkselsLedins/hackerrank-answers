#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main(){
    int n;
    scanf("%d",&n);
    int arr[n];

    int positive = 0;
    int negative = 0;
    int zeroes = 0;

    for(int arr_i = 0; arr_i < n; arr_i++){
       scanf("%d",&arr[arr_i]);
       if (arr[arr_i] > 0) { positive += 1; }
       if (arr[arr_i] == 0) { zeroes += 1; }
       if (arr[arr_i] < 0) { negative+= 1; }
    }

    printf("%f\n%f\n%f", (float)positive / n, (float)negative / n, (float)zeroes / n);

    return 0;
}
