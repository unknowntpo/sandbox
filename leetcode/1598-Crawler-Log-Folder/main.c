// Hello world! Cplayground is an online sandbox that makes it easy to try out
// code.

#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#define CUR_DIR 0
#define PRE_DIR 1
#define NEXT_DIR 2

int parseToken (char *);
void showLog (char **, int);

int minOperations(char **logs, int logsSize){
    int level = 0 ;
    showLog(logs, 5);

    for (int i = 0; i < logsSize; i++) {
        switch (parseToken(logs[i])) {
            case CUR_DIR:
                break;
            case PRE_DIR:
                if (level == 0) break;
                level--;
                break;
            case NEXT_DIR:
                level++;
                break;
        }  
    }
    return level;
}

int parseToken (char *token) {
    if (strcmp(token, "./") == CUR_DIR) return CUR_DIR;
    if (strcmp(token, "../") == PRE_DIR) return PRE_DIR;
    return NEXT_DIR;
}

void showLog(char ** logs, int logSize) {
    for (int i = 0; i < logSize ; i++) {
        printf("logs[%d]: %s\n", i, logs[i]);
    }
}
int main() {
    //logs = ["d1/","d2/","../","d21/","./"]
    char *logs[] = {
        "d1/",
        "d2/",
        "../",
        "d21/",
        "./",
    };
    int r = parseToken(logs[0]);
    printf("parseToken: %d\n", r);

    printf("After init:\n");
    showLog(logs, 5);
    int res = minOperations(logs, 5);
    printf("\nResult: %d\n", res);
}
