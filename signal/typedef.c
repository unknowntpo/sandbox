typedef void (*fp)(int);

fp signal(int signum, fp handler);
int main(){};
