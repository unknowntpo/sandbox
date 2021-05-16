#include <stdio.h>
#include <pthread.h>
//#include <unistd.h>

void *thread(void *);

int main()
{
    pthread_t tid;
    pthread_create(&tid, NULL, thread, NULL);
    pthread_join(tid, NULL);
}

void *thread(void *vargp)
{
    for (int i = 0; i < 10; i++) {
        printf("%d\n", i);
    }

    return NULL;
}



