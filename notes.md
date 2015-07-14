
# simplebench notes

## The test server program

It’s a simple HTTP server which listens for a GET request at a certain path, does some trivial work
to generate a JSON response, and returns it.


## Setup

### AWS EC2

We currently run apicard on `c3.large` instances, which have:

* 2 vCPUs
* 7 Elastic Compute Units
* 3.75 GB RAM
* 2 x 16GB SSD

and cost $0.105 per hour / $77 per month

So the server for this test was of the same type.

### Heroku

The closest option they have to EC2 c3.large is Standard 2X dynos, which have:

* 4 CPUs
* 1 GB RAM
* an undefined amount of ephemeral storage of unknown speed

and cost $50 per month each

## The Benchmark

We spun up a `c3.large` EC2 instance to serve as the client.

The command:

    wrk -c <100,200,300> -t 50 -d 10s <url>

was run 5 times for each level of concurrency for each URL.

Notes:

* I used the public IP of the EC2 server
* I did not specify or control for the availability zones used for either EC2 instance
* might want to try out Boom too


## The Results

All results are requests per second from the EC2 client machine to the server under test.

### EC2 c3.large instance ($77/month)

#### 9 July 2015

TO DO: re-run these tests through an ELB

Concurrency | Run 1 | Run 2 | Run 3 | Run 4 | Run 5 | Mean  | Std Dev
----------- | ----- | ----- | ----- | ----- | ----- | ----- | -------
100         | 16765 | 16852 | 16739 | 16874 | 16793 | 16805 | 57
200         | 16968 | 16976 | 17015 | 16964 | 16965 | 16978 | 21
300         | 16574 | 16705 | 16662 | 16696 | 16701 | 16668 | 55


### Heroku Standard 2X dynos ($50/month each)

#### 9 July 2015

Dynos | Concurrency | Run 1 | Run 2 | Run 3 | Run 4 | Run 5 | Mean | Std Dev
----- | ----------- | ----- | ----- | ----- | ----- | ----- | ---- | -------
1     | 100         | 4701  | 4333  | 5351  | 4752  | 3998  | 4627 | 507

#### 13 July 2015

Dynos | Concurrency | Run 1 | Run 2 | Run 3 | Run 4 | Run 5 | Mean | Std Dev
----- | ----------- | ----- | ----- | ----- | ----- | ----- | ---- | -------
2     | 100         | 3023  | 3767  | 6869  | 2789  | 3132  | 3916 | 1690
2     | 200         | 4968  | 5423  | 5140  | 5281  | 5054  | 5173 | 181
2     | 300         | 6967  | 5637  | 8176  | 8434  | 6154  | 7074 | 1223
3     | 300         | 3111  | 3192  | 3609  | 3296  | 3217  | 3285 | 193


### Heroku Performance dynos ($500/month each)

#### 13 July 2015

Dynos | Concurrency | Run 1 | Run 2 | Run 3 | Run 4 | Run 5 | Mean  | Std Dev
----- | ----------- | ----- | ----- | ----- | ----- | ----- | ----- | -------
1     | 100         | 7751  | 9723  | 8656  | 8803  | 7597  | 8506  | 864
1     | 200         | 10979 | 10874 | 10688 | 10758 | 10829 | 10826 | 111
1     | 300         | 11060 | 10599 | 10599 | 10478 | 10665 | 10680 | 223
