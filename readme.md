● The .csv file could be very big (billions of entries) - how would your application perform?
I think it depends on the server. Because all records are stored in the memory of the application and time complexity of Get operation is O(1) right now.
But it's still bad practice to handle a large data stream using only the memory.
Looks like Redis could perform it much better.

● Every new file is immutable, that is, you should erase and write the whole storage;
If it's possible to use Redis here we can just update it with new records.
If it's not I think we need to merge two files into a new one using the background task and when it's done just swap the main file with the new one.

● How would your application perform in peak periods (millions of requests per minute)?
Depends on the server again, but there definitely must be load-balancing infrastructure if we are talking about high-load applications.

● How would you operate this app in production (e.g. deployment, scaling, monitoring)?
Deployment - GitHub actions or Gitlab CI
Scaling - k8s (Amazon EKS)
Monitoring - ELK Stack for logs and Prometheus with Grafana for metrics. 
