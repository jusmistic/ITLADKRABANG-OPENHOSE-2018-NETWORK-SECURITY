#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>
int key();
int main(){
    int match = 0 ;
    
    //Random times
    time_t t;
    srand((unsigned) time(&t));
    int tour = rand();
    //handle is time  = 0
    if (tour <= 0 ) tour = rand();
    tour = tour %1000;
    printf("Tournament %d Team.\n", tour);
    printf("How many matches?\nAns: ");
    scanf("%d", &match);
    if (match == tour-1){
        key();
    } else{
        printf("Incorrect! Try again.");
    }
}

int key(){
    // char text[128] = "Flag{___W3lc0me___T0__Tou2n@m3nt!__}";
    char key[128] = "asdfghjklwertyuiosdfghjzxcvbsdfghjkwertyu";
    int flag[50] = {39, 31, 5, 1, 28, 55, 53, 52, 59, 68, 9, 17, 68, 20, 16, 54, 48, 44, 48, 86, 56, 55, 62, 21, 13, 81, 24, 34, 30, 87, 8, 19, 73, 53, 52, 10, 101};
    for(int i = 0; i<= 36; i++){
        printf("%c", key[i] ^ flag[i]);
    }
}