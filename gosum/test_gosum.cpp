#include <stdio.h>
#include <malloc.h>
#include <string.h>
#include "gosum.h"

// g++ -c test_gosum.cpp -o test_gosum.o -g -std=c++11
// g++ test_gosum.o gosum.dll -o test_gosum.exe -g -std=c++11
int main() {
    printf("Sum(1,2)=%d\n", Sum(1,2));
    return 1;
}
