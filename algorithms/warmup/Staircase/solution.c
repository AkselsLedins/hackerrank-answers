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

    for (int i = 0; i < n; i++) {
        for (int x = i + 1; x < n; x++) {
            printf(" ");
        }
        for (int y = i; y >= 0; y--) {
            printf("#");
        }
        printf("\n");
    }

    return 0;
}
