
With the following code:

pthread_t incrementingThread;
    pthread_t decrementingThread;

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);  
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);

    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);

    // Get race condition if we create and dont join afterwards, since the threads will run in parallel

However we can (not in a good way split them up)

pthread_t incrementingThread;
    pthread_t decrementingThread;

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL); 
    pthread_join(incrementingThread, NULL);

    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    pthread_join(decrementingThread, NULL);
    
    // But this completely misses the point of concurrency since we are simply processing a single thread at a time.

--------------------------------------Concurrency solution:-----------------------------------

Both a semaphore implementation (maybe using a binary semaphore) as the semaphore as a lock, and mutex implementation will work. Mutexes are simpler and suit our 1 at a time use case very cleanly so that is what i will use in the program.