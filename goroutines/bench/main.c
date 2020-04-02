#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define MAX_T 100000
#define DEF_T 1000
#define LOOPS 1000

typedef struct _foo
{
    pthread_mutex_t mutex;
    int counter;
} foo;

typedef struct _args
{
    foo *counter;
} arguments;

pthread_t ntid[MAX_T];

void *example(void *arg)
{
    arguments *data = arg;

    for (int i = 0; i < LOOPS; i++) continue;

    pthread_mutex_lock(&data->counter->mutex);
    data->counter->counter++;
    pthread_mutex_unlock(&data->counter->mutex);

    return NULL;
}

int main(int argc, char *argv[])
{
    int n = (argc > 1) ? atoi(argv[1]) : DEF_T;

    foo *counter = malloc(sizeof(*counter));
    arguments *args = malloc(sizeof(*args));

    counter->mutex = (pthread_mutex_t)PTHREAD_MUTEX_INITIALIZER;
    counter->counter = 0;
    args->counter = counter;

    for (int i = 0; i < n; i++)
    {
        if (pthread_create(&ntid[i], NULL, example, args) != 0)
        {
            free(counter);
            free(args);
            exit(1);
        }
    }

    for (int i = 0; i < n; i++)
    {
        if (pthread_join(ntid[i], NULL) != 0)
        {
            free(counter);
            free(args);
            exit(1);
        }
    }

    printf("> the counter is %d\n", counter->counter);

    free(counter);
    free(args);

    return 0;
}