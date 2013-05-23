#include <stdio.h>
#include <pthread.h>
#define FALSE 0
#define TRUE 1

int busy = FALSE;
int waiting = 0;
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t OKtoread = PTHREAD_COND_INITIALIZER;
pthread_cond_t OKtowrite = PTHREAD_COND_INITIALIZER;
int readercount = 0;

void readerStart(){	
	pthread_mutex_lock(&mutex);
	if (busy == TRUE){
		pthread_cond_wait(&OKtoread,&mutex);
	}
	readercount = readercount + 1;
	pthread_cond_signal(&OKtoread);
	pthread_mutex_unlock(&mutex);
}

void readerEnd(){
	pthread_mutex_lock(&mutex);
	readercount = readercount - 1;
	if (busy == TRUE){
		pthread_cond_signal(&OKtowrite);
	}
	pthread_mutex_unlock(&mutex);
}

void writerStart(){
	pthread_mutex_lock(&mutex);
	waiting = waiting + 1;
	if (busy == TRUE || readercount != 0){
		pthread_cond_wait(&OKtowrite,&mutex);
	}
	busy = TRUE;
	waiting = waiting - 1;
	pthread_mutex_unlock(&mutex);
}

void writerEnd(){
	pthread_mutex_lock(&mutex);
	busy = FALSE;
	if (waiting > 0){
		pthread_cond_signal(&OKtoread);
	}else{
		pthread_cond_signal(&OKtowrite);
	}
	pthread_mutex_unlock(&mutex);
}

int main()
{
	printf("ende");
	return 0;
}