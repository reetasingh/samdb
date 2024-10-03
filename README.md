# SAMDB
## Building Redis from Scratch

### Checking for Port Usage
To check if a server is already using a specific port, run:

```bash
netstat -va | grep -i "LISTEN"
```

### Connecting Redis Client to the Server

To connect the Redis client to the server on port 7380, execute:

```bash
redis-cli -p 7380
```

# Milestone 1: Working PING Command

Test the connection with the PING command:
```bash
reeta@Reetas-MacBook-Pro samdb % redis-cli -p 7380
127.0.0.1:7380> PING
"PONG"
127.0.0.1:7380> PING hi bye
(error) wrong number of arguments to PING cmd
127.0.0.1:7380>
```

# Milestone 2: Make it async to support multiple clients using go-routine

Test the connection with the PING command:
```bash
client connected: 127.0.0.1:55607
concurrent connections: 1Received cmd:  COMMAND [DOCS]
Received cmd:  PING []
client connected: 127.0.0.1:55608
concurrent connections: 2Received cmd:  COMMAND [DOCS]
Received cmd:  PING []
Received cmd:  PING [a]
client connected: 127.0.0.1:55610
concurrent connections: 3Received cmd:  COMMAND [DOCS]
Received cmd:  PING [dhhd]
client connected: 127.0.0.1:55611
concurrent connections: 4Received cmd:  COMMAND [DOCS]
```

Benchmark for ping command
```bash
reeta@Reetas-MacBook-Pro samdb % redis-benchmark -p 7380 -q -n 100000 PING
WARNING: Could not fetch server CONFIG
PING: 131752.31 requests per second, p50=0.151 msec
reeta@Reetas-MacBook-Pro samdb % redis-benchmark -p 7380 -q -n 100000 PING
WARNING: Could not fetch server CONFIG
PING: 137931.03 requests per second, p50=0.143 msec
reeta@Reetas-MacBook-Pro samdb %
```
