Q3:
The thing that happens when you try to run the two threads at once is that the treads continouslty try to interrupt eachother. To fix this i had to make the action inside the for loop atomic e.i. make it so the write sequence does not get interrupted. 

Q4:

I used the pthread mutex. I assume this is the correct one because it worked and it was the only one that i tried. After reading a bit on the web, it seems like the semaphores one makes it possible for the other thread to access the process, whitch is something i dont want to happen.  