#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

#define MAX_SIZE 10  // Define the maximum size of the queue

struct stack_map_bucket {
    uint32_t head;
    uint32_t tail;
    uint64_t data[MAX_SIZE];
};

// Function to initialize the queue
void init_queue(struct stack_map_bucket *bucket) {
    bucket->head = 0;
    bucket->tail = 0;
}

// Function to check if the queue is empty
int is_empty(struct stack_map_bucket *bucket) {
    return bucket->head == bucket->tail;
}

// Function to check if the queue is full
int is_full(struct stack_map_bucket *bucket) {
    return ((bucket->tail + 1) % MAX_SIZE) == bucket->head;
}

// Function to enqueue an element into the queue
int enqueue(struct stack_map_bucket *bucket, uint64_t value) {
    if (is_full(bucket)) {
        printf("Queue is full!\n");
        return -1; // Indicate that the queue is full
    }
    bucket->data[bucket->tail] = value;
    bucket->tail = (bucket->tail + 1) % MAX_SIZE;
    return 0; // Indicate success
}

// Function to dequeue an element from the queue
int dequeue(struct stack_map_bucket *bucket, uint64_t *value) {
    if (is_empty(bucket)) {
        printf("Queue is empty!\n");
        return -1; // Indicate that the queue is empty
    }
    *value = bucket->data[bucket->head];
    bucket->head = (bucket->head + 1) % MAX_SIZE;
    return 0; // Indicate success
}

// Function to print the elements of the queue
void print_queue(struct stack_map_bucket *bucket) {
    uint32_t i = bucket->head;
    while (i != bucket->tail) {
        printf("%llu ", bucket->data[i]);
        i = (i + 1) % MAX_SIZE;
    }
    printf("\n");
}

int main() {
    struct stack_map_bucket bucket;
    init_queue(&bucket);

    enqueue(&bucket, 10);
    enqueue(&bucket, 20);
    enqueue(&bucket, 30);
    enqueue(&bucket, 40);

    printf("Queue after enqueuing 10, 20, 30, 40:\n");
    print_queue(&bucket);

    uint64_t value;
    dequeue(&bucket, &value);
    printf("Dequeued value: %llu\n", value);

    printf("Queue after dequeuing an element:\n");
    print_queue(&bucket);

    return 0;
}
