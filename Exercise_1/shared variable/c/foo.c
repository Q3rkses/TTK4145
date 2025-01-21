// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t i_lock;


// Note the return type: void*
void* incrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&i_lock);
        i++;
        pthread_mutex_unlock(&i_lock);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int j = 0; j < 1042069; j++) {
        pthread_mutex_lock(&i_lock);
        i--;
        pthread_mutex_unlock(&i_lock);
    }
    return NULL;
}


int main(){
    
    pthread_t incrementingThread;
    pthread_t decrementingThread;
    pthread_mutex_init(&i_lock, NULL);

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);  
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);

    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);

    pthread_mutex_destroy(&i_lock);
    
    printf("The magic number is: %d\n", i);
    return 0;
}
