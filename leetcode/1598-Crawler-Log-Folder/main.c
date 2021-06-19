#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#define CUR_DIR 0
#define PRE_DIR 1
#define NEXT_DIR 2
#define DEBUG 0

int parseToken (char *);
void showLog (char **, int);

int minOperations(char **logs, int logsSize){
    int level = 0 ;
#if DEBUG
    showLog(logs, 5);
#endif
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
#if DEBUG
    printf("inside parseToken: token: %s\n", token);
#endif
    if (strcmp(token, "./") == 0) return CUR_DIR;
    if (strcmp(token, "../") == 0) return PRE_DIR;
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
#if DEBUG
    showLog(logs, 5);
#endif
    int res = minOperations(logs, 5);
    printf("\nResult: %d\n", res);
}
