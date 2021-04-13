#if 0
void (*signal(int, void (*fp)(int)))(int);

int main(){};
#endif

typedef void (*fp)(int);

//void (*signal(int, fp))(int);

fp signalhandler(int, fp);
