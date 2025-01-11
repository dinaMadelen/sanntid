// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

pthread_mutex_t mutex;

int i = 0;

// Note the return type: void*
void* incrementingThreadFunction(void * arg){
    for (int n = 0; n < 1000000; n++){
        pthread_mutex_lock(&mutex); // Lock the mutex
        i++;                        // Critical section
        pthread_mutex_unlock(&mutex); // Unlock the mutex
    }
    return NULL;
}

void* decrementingThreadFunction(void * arg){
    // TODO: decrement i 1_000_000 times
    for (int n = 0; n < 1000000; n++){
        pthread_mutex_lock(&mutex); // Lock the mutex
        i--;                        // Critical section
        pthread_mutex_unlock(&mutex); // Unlock the mutex
    };

    return NULL;
}


int main(){
    // TODO: 
    //incrementingThreadFunction(NULL);

    pthread_mutex_init(&mutex, NULL);

    pthread_t incrementingThread; 
    pthread_t decrementingThread;

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    printf("The magic number is: %d\n", i);
    pthread_join(decrementingThread, NULL);
    
    
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`    
    
    printf("The magic number is: %d\n", i);

    pthread_mutex_destroy(&mutex);
    return 0;
}
