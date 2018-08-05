#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int main(){
    int flag;
    char buffer[64];
    flag = 0;
    printf("Input Sometext here:\n");
    gets(buffer);
    if(flag != 0) {
        printf("Flag{_________________}\n"); //Flag store here in server
    } else {
      printf("Aww Can't Change?\n");
  }
}