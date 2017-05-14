#include "stdio.h"

// Unions has the memory of the largest data type of its fields
// Union can hold the correct value for only one of its fields at any given point of time
// The latest assigned field value will be consistent and holded in memory for sure. 

// Go has no support for untagged unions, as it would violate its memory safety guarantees.
// https://golang.org/doc/faq#unions

union BigHolder {
    double a;
    float b;
    int c;
} big;

union SmallHolder {
    int a;
    float b;
} small;

int main() {

    if (sizeof(big) != 8 ) {
        printf("Expected 8, got %lu \n", sizeof(big));
        return 1;
    }

    if (sizeof(small) != 4 ) {
        printf("Expected 4, got %lu \n", sizeof(small));
        return 1;
    }

    big.a = 1;
    printf("big.a: %f big.b: %f big.c: %d \n", big.a, big.b, big.c);
    big.b = 2;
    printf("big.a: %f big.b: %f big.c: %d \n", big.a, big.b, big.c);
    big.c = 2147483647;
    printf("big.a: %f big.b: %f big.c: %d \n", big.a, big.b, big.c);

    return 0;
}