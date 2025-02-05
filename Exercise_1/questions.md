Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> *Parralelism is when multiple threads run simultaneously, however concurrency is when multiple different threads run simultaneously and produce the correct (expected result)*

What is the difference between a *race condition* and a *data race*? 
> *Data race  .*
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> *Scheduler as the name suggests, keeps track of the order of when operations should be executed, and then executres them. When new requests come in it schedules them in the correct place (or i guess one could say at the correct time)* 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *Multiple threads allow us to do more "things" in parallell and if used effectively can make use of more % of the maximum computational power of the CPU. As an example two completely seperate operations that do not interrupt each other and have no dependencies between each other do not need to wait for one another to finish. Threads divide the program up into sub-parts*

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *It depends on the programmer, and the use case.*

What do you think is best - *shared variables* or *message passing*?
> *Both have their time and usefulness. There are times where one cannot choose one, and therefore needs to know how to utilize both.*


