Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> *Concurrency is how a computer with one processor will handle tasks that are optimally solved in parralell .e.i. Concurrency makes it seem like it is parallel by interweaving tasks, while parallelism is actually solving tasks simultaniously and separately*

What is the difference between a *race condition* and a *data race*? 
> *Race condition occurs when the program needs tasks to be done in a certain order for it to have a correct outcome. Data race is a spesific type of race condition, it is about when two threads wants to access the same variable atr the same time.* 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> *It makes the program execute tasks in order. It does it by switching context, allocating resources and prioritizing what should come next in the program.* 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *They are there to simulate having separate systems for separate tasks when you have a single processor*

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Fibers are sceduled by user level code, while threads are sceduled by th os kernel. This means that the programmer have "more" control over the fiber scheduling than the threads*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Easier if you do it properly. But if you fuck up scheduling or things like that you can make nesting hell.*

What do you think is best - *shared variables* or *message passing*?
> *I like message passing, depends on system*


