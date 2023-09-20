a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
    - The "data" struct is our datastructure. It has three fields, seq, ack and message. 

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
    - We use threads. It is not realistic because this implementation is made in a closed circuit, so fx data can't get lost due to channel failure. 

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
    - We would use sequence numbers so the receiver knows which order to put data in. 

d) In case messages can be delayed or lost, how does your implementation handle message loss?
    - We have ack which we validate with if-statements and if it fails we end the handshake. 

e) Why is the 3-way handshake important?
    - TCP is connection oriented so for it to be reliable we check both directions of connection before sending any data. 