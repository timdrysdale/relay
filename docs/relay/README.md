# relay
Secure websocket relay server and clients for sharing video, data, and ssh across firewall boundaries


IN DEVELOPMENT NOW - REPO HERE FOR BACKUP ONLY - DO NOT USE - API IN FLUX


Ignore below here (README.md from crossbar ...)

## Benchmarking

Benchmarking of a mixed-communications relay can be relatively complex, although simple results are enough to give a picture of the performance available. Note that simple calculations on mixed loads, based on the single-producer, single-consumer benchmarks cannot be done accurately because of compiler efficiencies e.g. in duplicating outgoing messages (an order of magnitude less CPU + memory load to duplicate outgoing videos, than it is to receive an incoming video)

### Single-producer, single-consumer benchmarks

On Intel(R) Xeon(R) Silver 4110 CPU @ 2.10GHz

#### Small message
small timeouts (1ms)

3-byte payload: 1000 messages per benchmark (single producer, single consumer)
```
BenchmarkSmallMessage-32    	       8	 134451695 ns/op
```
134.5ms for 1000 messages equates to in excess of 7,000 messages per second

#### 1MByte message

Random data:

Each test is 100 message of 1MB, taking 934ms.
The total data throughput is 8bits/Byte * 100MB = 800Mbits


```
BenchmarkLargeMessage: crossbar_test.go:488: Message size: 1048576 bytes
BenchmarkLargeMessage-32    	       2	 933766152 ns/op
```

for a bandwidth of 800Mbits/0.934sec = 858 Mbps (Mega bits per sec)

This bandwidth of ~850 Mbps is 

  - more than twice my fibre broadband download speed (350 Mbps * 2 = 700 Mbps)
  - close to the maximum theoretical limit of a GiGE network connection (1000 Mbps)

The test used the same data in each packet to avoid dragging the results down with the overhead of generating random data. But we can explore whether the compilier is doing some fancy caching optimisation by running a separate benchmark on the random data generation.
```
BenchmarkLargeMessage
BenchmarkLargeMessage-32                   	       2	 958564581 ns/op
BenchmarkLargeRandomPacketGeneration
BenchmarkLargeRandomPacketGeneration-32    	       2	 618112575 ns/op
BenchmarkLargeRandomMessage
BenchmarkLargeRandomMessage-32             	       1	2421145176 ns/op
```

The adjusted duration for sending and receiving is now 2.42 - 0.62 = 1.8 sec

This gives an adjusted bandwidth, ignoring the data generation time, of 444 Mbps.

This bandwidth of ~450 Mbps is 

  - still in excess of my fibre broadband download speed (350 Mbps)
  - close to the half the maximum theoretical limit of a GiGE network connection (1000 Mbps)

This is still a fairly benign test, in that all the network traffic is internal to my machine.


### Real world test (>=375Mbps)

  - Videos (20) source - Edinburgh, UK
  - Crossbar - London, UK
  - Receivers (>300) - Edinburgh, UK

Running on EC2 c5.(x)large in EU-london I ran a 'real use' test and managted to sustain 20 videos in and >300 videos out, and could not push further because of my home fibre bandwidth limits.
Each video captured live from raspberry pi on my home network (20x)
Individual bandwidth was 1.1Mbps per video, ~30fps 640x480 MPEGTS from ffmpeg for use jsmpeg decoder in browser
Total bandwidth in from video cameras was 22-24Mbps (within the published limit of my home fibre of 30Mbps)
Total bandwith out ~375Mbps (just over the published limit of my home fibre sysm of 350Mbps).

The peak bandwith for the testing was done with 500 video clients, but this hit the rate limit of my fibre to my house, and it is therefore possible that crossbar could do more ....


## TODO

```
./pkg/crossbar/crossbar.go:11:	messagesToDistribute := make(chan message, 10) //TODO make buffer length configurable
./pkg/crossbar/models.go~:144:// TODO - remove unused types below this line (some still in use)
./pkg/crossbar/handleconnections.go:37:// TODO
./pkg/crossbar/handleconnections.go:51:// TODO restrict CheckOrigin
./pkg/crossbar/handleconnections.go:260:				// TODO check what impact, if any, this has on jsmpeg memory requirements
./pkg/crossbar/handleconnections.go:262:				// TODO benchmark effect of loading on message queuing
./pkg/crossbar/handleconnections.go:303:	// TODO check authorisation to get stats
./pkg/crossbar/handleconnections.go:315:	// TODO use gobuffalo/packr to embed this file or wait for go1.16
./pkg/crossbar/handleconnections.go:346:			log.Info("ListenAndServe: ", err) //TODO upgrade to fatal once httptest is supported
./pkg/crossbar/crossbar_oldtest.nogo:323:	// TODO we will want to build servers and clients that are polite about reconnecting
./pkg/crossbar/crossbar_oldtest.nogo:725:	//TODO - add support for httptest https://stackoverflow.com/questions/40786526/resetting-http-handlers-in-golang-for-unit-testing
./pkg/crossbar/models.go:134:// TODO - remove unused types below this line (some still in use)
./pkg/crossbar/servews.go:37:		ct = Unsupported //TODO implement shell!
./pkg/crossbar/handleconnections_test.go:3:// TODO refactor with interface in client struct to allow mocking
./pkg/crossbar/crossbar.go~:14:	messagesToDistribute := make(chan message, 10) //TODO make buffer length configurable
./pkg/access/access.go:68:			// TODO - send correct error code, 401 / 403 rather than 500
./pkg/relay/relay.go:14:	messagesToDistribute := make(chan message, 10) //TODO make buffer length configurable

```





#Crossbar
![alt text][logo]
![alt text][status]
![alt text][coverage]
```
Crossbar relays websocket streams

## Why?

Remote laboratories require a custom video and data communications eco-system in order to support their wider adoption. Key factors include:

+ Remote laboratory participants (whether human or machine) are often located behind institutional firewalls and NAT
+ Most instituational networks support particpants sending and receiving video and data to external relays, but not acting as a server
+ Those data streams are typically embedded in websockets, whereas UDP and some TCP protocols are sometimes explicitly blocked
+ Almost all video and data messaging vendors with relays are focused solely on human-human communications
    + often missing apparently-minor features from the API that are key for remote lab experiments (e.g. being able to change camera programmatically)
	+ often require workarounds for remote lab adminstration tasks which are prevented by privacy features in browsers (e.g. identifying cameras)
	+ typically require x10 more expensive computer for the experiment because of the overhead of running graphical operating system and browser
	+ most vendors can - quite rightly - only guarantee long-term support for users that conform to their core use-case

## Features

+ binary and text websockets
+ multiple, independent streams
    + organised by topics
	+ topics set by routing
+ multiple streaming schemas	
    + bidirectional N:N streaming
    + unidirectional 1:N streaming
+ streaming schemas are 
    + set by routing
	+ individually selectable for each stream 
+ statistics are recorded, and available via
    + websocket API in JSON format
	+ human-readable webpage with 5sec update at ```<host>/stats```
	
## What's with the name?
I once had an old IBM p690 compute cluster whose processor cores had crossbar switches ([Core Interface Units](http://www.netlib.org/utk/papers/advanced-computers/power4.html)) that connect any core with any cache, and [do so more efficiently than a standard compute cluster](http://www.piers.org/piersonline/pdf/Vol5No2Page117to120.pdf). It seemed apt, because this relay is about connecting any experiment with any human, more efficiently than existing systems, (in a holistic sense, including total cost and effort of ownership, administration, maintenance etc).

## Performance

After some profiling-led optimisation, I've run the code for four months continuously (as of March 2020) with multiple streams and re-connections of broadcasters and receivers, including a fortnight where a stream was connecting to a new routing every second (in a sequence of four) using a bash script and [timdrysdale/vw](https://github.com/timdrysdale/vw), with no appreciable leakage of memory.

Further quantitative benchmarking is needed to understand how well crossbar scales to handle large numbers of streams and clients simultaneously, but from the perspective of a light user with around four video streams being in use for live demonstrations at any one time, it has performed well and given confidence. Performance at low N is no guarantee of performance at high N, so please conduct your own testing if adopting at this time (and share with me for inclusion here, via email or PR). 

Meanwhile, here's a text-grab from ```top``` on an Amazon EC2 ```c5.Large``` instance showing the same memory usage now as four months ago under the same load.

```
top - 12:57:15 up 232 days, 22:45, 14 users,  load average: 0.04, 0.04, 0.00
Tasks: 123 total,   1 running,  83 sleeping,   0 stopped,   0 zombie
%Cpu(s):  2.0 us,  0.5 sy,  0.0 ni, 97.5 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :  3794284 total,   486436 free,   691080 used,  2616768 buff/cache
KiB Swap:        0 total,        0 free,        0 used.  2831916 avail Mem 

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND             
24363 root      20   0  252616  52060   4720 S   4.0  1.4   1019:36 iftop               
31302 ubuntu    20   0  930664  26504   8856 S   1.7  0.7   3852:37 crossbar            
 5917 www-data  20   0  145712  11396   7852 S   1.3  0.3 123:49.74 nginx               
23179 root      20   0  447328  42808  23328 S   0.7  1.1 292:00.12 containerd          
  816 jvb       20   0 5770988 135244   8936 S   0.3  3.6   1002:19 java                
23051 root      20   0  471992  76156  40564 S   0.3  2.0 181:01.84 dockerd             
    1 root      20   0  225484   9556   7068 S   0.0  0.3  12:45.70 systemd
```

## Getting started

Binaries will be available in the future, but for now it is straightforward to compile.

It's quick and easy to create a working ```go``` system, [following the advice here](https://golang.org/doc/install).

Get the code, test, and install
```
$ go get github.com/timdrysdale/crossbar
$ cd `go list -f '{{.Dir}}' github.com/timdrysdale/crossbar`
$ go test ./cmd
$ go install
```
To run the relay
```
$export CROSSBAR_LISTEN=127.0.0.1:9999
$ crossbar
```

Navigate to the stats page, e.g. ```http://localhost:9999/stats```

You won't see any connections, yet.

You can connect using any of the other tools in the ```practable``` ecosystem, e.g. [timdrysdale/vw](https://github.com/vw/timdrysdale) or if you already have the useful [websocat](https://github.com/vi/websocat) installed, then

```
websocat --text ws://127.0.0.1:8089/expt - 
```
If you type some messages, you will see the contents of the row change to reflect that you have sent messages.


### Multiple clients are supported

If you connect a second or even third or more times from other terminals, you will see the hub relaying your messages to all other clients.
Try typing in each of the terminals, and see that your message makes it to each of the others.

### Streams are independent

Try setting up a pair of terminals that are using a different topic, and notice that messages do not pass from one topic to another.

e.g. connect from two terminals using
```
websocat --text ws://127.0.0.1:8089/sometopic - 
```

Messages sent in a terminal connected to ```<>/sometopic``` will only go to terminals connected to the same route, and not to any other terminal.

Here's a screenshot (note that ```websocat``` does a local echo so the sender can see their message; the echo is not from ```crossbar```)

![alt text][topics]

## Usage examples

The default is bidirectional messaging, as you have seen in the example above. 

### Unidirectional messaging

If you only want to broadcast messages, such as a video stream, then it is nice to have some certainty that one of your clients won't inadvertently mess up the video for others by transmitting some sort of reply. To take advantage of uni-directional messaging, start the server's route with ```/in/``` and the clients' routes with ```/out/```. Note that the rest of the route has to match.

You can try it out yourself.

```
websocat --text ws://127.0.0.1:8089/in/demo - 
```
and
```
websocat --text ws://127.0.0.1:8089/out/demo - 
```

You can see from the local echo that messages attempted to be sent from the clients connected to ```/out/``` are not sent to any other clients - this is enforced by the hub and does not need any special behaviour from the clients (beyond connecting to the right route). Protecting unauthorised users from connecting to the ```/in/``` route is outside the scope of the ```crossbar``` codebase, in line with conventional practice on separating concerns. 

![alt text][unidirectional]

## Applications

### Relaying video and data

Crossbar has been successfully relaying MPEG video and JSON data (on separate topics) for [penduino-ui](https://github.com/timdrysdale/penduino-ui) experiments using [ffmpeg](https://ffmpeg.org) and [timdrysdale/vw](https://github.com/timdrysdale/vw).  

### Relaying shell access

Shell relay is in a separate project because these have different goals, development schedules, and performance targets, even though some of the underlying code and approach is similar. See [timdrysdale/shellbar]( https://github.com/timdrysdale/shellbar).

## Support / Contributions

Get in touch! My efforts are going into system development at present so support for alpha-stage users is by personal arrangement/consultancy at present. If you wish to contribute, or have development resource available to do so, I'd be more than delighted to discuss with you the possibilities.

## Developer stuff

The swagger2.0 spec for the relay-access API can be used to auto-generate the boilerplate for the server code with the command
```
swagger generate server -t pkg/access -f ./api/openapi-spec/access.yml --exclude-main -A access
```

Some links to articles on swagger and authentication: [jwt auth](https://shashankvivek-7.medium.com/go-swagger-user-authentication-securing-api-using-jwt-part-2-c80fdc1a020a); [context](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39).


[status]: https://img.shields.io/badge/alpha-working-green "Alpha status; working"
[coverage]: https://img.shields.io/badge/coverage-54%25-orange "Test coverage 54%"
[logo]: ./img/logo.png "Crossbar logo - hexagons connected in a network to a letter C"
[topics]: ./img/topics.png "Multiple terminals showing bidirectional message flow"
[unidirectional]: ./img/unidirectional.png "Multiple terminals showing unidirectional message flow"
