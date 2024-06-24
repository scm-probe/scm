#include <stdio.h>
#include <unistd.h>
#include <sys/prctl.h>
#include <string.h>

int main() {
    int child_pid;
    char input;

    // char process_name[16] = "test_fork";

    // prctl(PR_SET_NAME, (unsigned long) process_name, 0, 0, 0);

    while (1) {
        printf("Press Enter to fork a new process (or 'q' to quit): ");
        input = getchar();

        if (input == 'q') {
            break;
        }
        while (getchar() != '\n');

        child_pid = fork();

        if (child_pid < 0) {
            perror("Fork failed");
            return 1;
        } else if (child_pid == 0) {
            printf("Child Process ID: %d\n", getpid());
            printf("Parent Process ID from Child: %d\n", getppid());
            return 0;
        } else {
            printf("Parent Process ID: %d\n", getpid());
            printf("Child Process ID from Parent: %d\n", child_pid);
        }
    }

    return 0;
}
