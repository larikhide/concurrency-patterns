# concurrency-patterns  


Here's an attempt to reproduce the generating patterns described in the AvitoTech YouTube [video](https://www.youtube.com/watch?v=GZSfn-8m-ko&ab_channel=AvitoTech)


![Alt text](src/image.png)  

Generating patterns are patterns that guide the data through one or more execution steps with possible modification of the data  

-- close the channel where we created it. Otherwise, you will need to create additional synchronization logic so that there will not be more than one attempt to close the channel, which may lead to panic

-- reading goroutines should never try to close the channel  

#Generator

![Alt text](src/generator.png)
