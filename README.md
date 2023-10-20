# concurrency-patterns  


Here's an attempt to reproduce the generating patterns described in the AvitoTech YouTube [video](https://www.youtube.com/watch?v=GZSfn-8m-ko&ab_channel=AvitoTech)


![Alt text](src/image.png)  

Generating patterns are patterns that guide the data through one or more execution steps with possible modification of the data  

-- close the channel where we created it. Otherwise, you will need to create additional synchronization logic so that there will not be more than one attempt to close the channel, which may lead to panic

-- reading goroutines should never try to close the channel  

# Generator

The generator pattern is used to create an ordered sequence of values, potentially infinite.  

When you need to read messages and process them in separate goroutines without blocking the reading from the queue (from a message broker or browser WebSocket). The generator will only handle reading from this queue into a buffered channel. Writing won't be blocked as long as there's space in the buffer.

![Alt text](src/generator.png)

# Fan in a.k.a multiplexor  

The Fan In pattern combines multiple inputs into a single output channel, i.e., it multiplexes. The order of output is not guaranteed!

Using the pattern when:

We want to limit concurrently executed requests to an external service, like a read limiter, where only one request is being processed at any given time.

Example: sending an email. Multiple goroutines generate various parts of the email, and one goroutine sends all this data to an external service in a single stream. The order of sending is not guaranteed.

![Alt text](src/fanin.png)  
![Alt text](src/fanin-extended.png)


# Fan out 

A useful pattern for distributing processing among multiple homogeneous goroutines. Not to be confused with publish-subscribe, as a message delivered to one subscriber cannot be delivered to another subscriber.  

Example: distributing processing among different I/O-bound goroutines, such as making network requests or load balancing between channels, as the scheduler will distribute work among different fanOut goroutines.
  
![Alt text](src/fanout.png)