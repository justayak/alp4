/*
A Multithreaded "Readers/Writers" Problem

This programming assignment is intended to demonstrate the
readers/writers problem. In addition, you will
become familiar with multithreaded programming using POSIX Threads,
concurrency control issues, and
critical section problems.

You are to implement the first readers/writers problem as described in
page 173 of the textbook. The shared
data element (SDE) is a timer (integer for our purposes). A writer
thread has exclusive access to SDE while
reader threads can access it concurrently. Only the last waiting writer
is allowed to update SDE, while the
others return without updating it. Your program should perform the
followings: 

    1.Create 4 writer threads (w1, w2, w3, w4). 
    2.Create 4 reader threads (r1, r2, r3, r4). 
    3.Read the initial value and the 4 writer values from the
command-line. 

      For example:

      $ a.out 0 5 8 2 4

      which means initialize SDE to 0, and use 5, 8, 2, and 4 as writer
values for w1, w2, w3, w4 threads,
      respectively.

    4.Use sleep(3) to simulate reading or writing process. 

In addition, you are required to develop a synchronization mechanism to
implement the following timing
diagram (Figure 1).
*/
/* Anan Tongprasith */
/********************/
#include<stdio.h>
#include<stdlib.h>
#include<pthread.h>

void invoke_write(void *mysde);
void invoke_read(void *num);

/* global variable */
pthread_mutex_t wrt,wrtcount,reader;
int sde,readcount,writecount;

/* Main function */
int main(int argc,char *argv[])
{	char *one="1",*two="2",*three="3",*four="4"; 	/* reader names */
	char *w1="1 ",*w2="2 ",*w3="3 ",*w4="4 ";	/* writer names */
	pthread_t writer1,writer2,writer3,writer4; /* 4 writer threads */
	pthread_t reader1,reader2,reader3,reader4; /* 4 reader threads */

	sde=atoi(argv[1]);			/* initial sde */
	strcat(w1,argv[2]);strcat(w2,argv[3]);  /* attach writer names */
	strcat(w3,argv[4]);strcat(w4,argv[5]);  /* with sde values */
/* Initialize mutex */
	pthread_mutex_init(&wrt,NULL);
	pthread_mutex_init(&wrtcount,NULL);
	pthread_mutex_init(&reader,NULL);
/* Start */
	pthread_create(&reader1,NULL,(void *) &invoke_read,(void *)one);
	sleep(1);	/* reader2 starts after reader1 */
	pthread_create(&reader2,NULL,(void *) &invoke_read,(void *)two);
/* wait until both readers finish then start writer1 */
	pthread_mutex_lock(&wrt); pthread_mutex_unlock(&wrt);
	pthread_create(&writer1,NULL,(void *) &invoke_write,(void *)w1);
	sleep(1);	/* reader3 starts after writer1 */
	pthread_create(&reader3,NULL,(void *) &invoke_read,(void *)three);
	sleep(3);	/* wait until reader3 finishes then start reader4 */
	pthread_mutex_lock(&wrt);pthread_mutex_unlock(&wrt);
	pthread_create(&reader4,NULL,(void *) &invoke_read,(void *)four);
	sleep(1);	/* writer2 starts after reader4 */
	pthread_create(&writer2,NULL,(void *) &invoke_write,(void *)w2);
	sleep(3);
	pthread_create(&writer3,NULL,(void *) &invoke_write,(void *)w3);
	sleep(1);
	pthread_create(&writer4,NULL,(void *) &invoke_write,(void *)w4);
/* Finish */
sleep(5);	/* wait until every thread exit */
	pthread_mutex_destroy(&wrt);pthread_mutex_destroy(&wrtcount);
	pthread_mutex_destroy(&reader);
}

void invoke_write(void *mysde)
{	int temp;char *myname="1",*tempsde="1";

/* use writecount to check who is the last waiting writer */
	pthread_mutex_lock(&wrtcount);
	  writecount=writecount+1;
	  temp=writecount;
	pthread_mutex_unlock(&wrtcount);
/* waiting for critical section */
	pthread_mutex_lock(&wrt);
/* critical section */
	sscanf((char *)mysde,"%c %s",myname,tempsde);
	  printf("writer%s -> entering\n",myname);
	  if(temp==writecount)
	  {
		sde=atoi(tempsde);
		sleep(3);
		printf("writer%s -> writing; SDE=%i\n",myname,sde);
	  }
	printf("writer%s -> exiting\n",myname);
/* exiting the critical section */
	pthread_mutex_unlock(&wrt);
}
void invoke_read(void *num)
{	int readernum;

	readernum=atoi((char *)num);
/* waiting for critical section */
	pthread_mutex_lock(&reader);
		readcount=readcount+1;
		if(readcount==1)  pthread_mutex_lock(&wrt);
	pthread_mutex_unlock(&reader);
/* critical section */
	printf("reader%i -> entering\n",readernum);
	sleep(3);
	printf("reader%i -> reading; SDE=%i\n",readernum,sde);
/* exiting critical section */
	pthread_mutex_lock(&reader);
		readcount=readcount-1;
		printf("reader%i -> exiting\n",readernum);
		if(readcount==0) pthread_mutex_unlock(&wrt);
	pthread_mutex_unlock(&reader);
}