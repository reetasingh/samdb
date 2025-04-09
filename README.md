

# SAMDB

## Building Redis from Scratch

### Checking for Port Usage

To check if a server is already using a specific port, run the following command:

```bash
netstat -va | grep -i "LISTEN"
```

### Connecting Redis Client to the Server

To connect the Redis client to the server on port `7380`, execute:

```bash
redis-cli -p 7380
```

---

## Progress Report

### Milestone 1: Working PING Command

Test the connection with the `PING` command:

```bash
reeta@Reetas-MacBook-Pro samdb % redis-cli -p 7380
127.0.0.1:7380> PING
"PONG"
127.0.0.1:7380> PING hi bye
(error) wrong number of arguments to PING cmd
127.0.0.1:7380>
```

---

### Milestone 2: Make It Async to Support Multiple Clients Using Go-Routines

Test the connection with the `PING` command, and observe how multiple clients are handled concurrently:

```bash
client connected: 127.0.0.1:55607
concurrent connections: 1
Received cmd:  COMMAND [DOCS]
Received cmd:  PING []

client connected: 127.0.0.1:55608
concurrent connections: 2
Received cmd:  COMMAND [DOCS]
Received cmd:  PING []
Received cmd:  PING [a]

client connected: 127.0.0.1:55610
concurrent connections: 3
Received cmd:  COMMAND [DOCS]
Received cmd:  PING [dhhd]

client connected: 127.0.0.1:55611
concurrent connections: 4
Received cmd:  COMMAND [DOCS]
```

---

### Benchmark for PING Command

To benchmark the `PING` command:

```bash
reeta@Reetas-MacBook-Pro samdb % redis-benchmark -p 7380 -q -n 100000 PING
WARNING: Could not fetch server CONFIG
PING: 131752.31 requests per second, p50=0.151 msec

reeta@Reetas-MacBook-Pro samdb % redis-benchmark -p 7380 -q -n 100000 PING
WARNING: Could not fetch server CONFIG
PING: 137931.03 requests per second, p50=0.143 msec
```

---

### Milestone 3: Working AsyncIO with Kernel Queue

Benchmark results after implementing async I/O with kernel queue:

```bash
reeta@Reetas-MacBook-Pro samdb % redis-benchmark -p 7380 -q -n 100000 PING
WARNING: Could not fetch server CONFIG
PING: 119047.62 requests per second, p50=0.399 msec
```

---

### Milestone 4: GET, SET, and TTL

Test the `SET`, `GET`, and `TTL` commands:

```bash
reeta@Reetas-MacBook-Pro samdb % redis-cli -p 7380
127.0.0.1:7380> SET db REDIS EX 10
OK
127.0.0.1:7380> GET db
"REDIS"
127.0.0.1:7380> TTL db
(integer) 3
127.0.0.1:7380> TTL db
(nil)
127.0.0.1:7380> GET db
(nil)
127.0.0.1:7380>
```

---

### Milestone 5: TTL, EXPIRE, and DELETE

Test the `TTL`, `EXPIRE`, and `DEL` commands:

```bash
127.0.0.1:7380> SET c d ex 100
OK
127.0.0.1:7380> TTL c
(integer) 98
127.0.0.1:7380> GET c
"d"
127.0.0.1:7380> EXPIRE c 200
(integer) 1
127.0.0.1:7380> TTL c
(integer) 199
127.0.0.1:7380> GET c
"d"
127.0.0.1:7380> DEL c
(integer) 1
127.0.0.1:7380> GET c
(nil)
127.0.0.1:7380> TTL c
(nil)
127.0.0.1:7380>
```

---

### Milestone 6: Adding Unit Tests Coverage

Run tests with coverage:

```bash
Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/d2/6_6kg3355wv25vyyhvpbdjxw0000gn/T/vscode-go7V7JcK/go-code-cover github.com/reetasingh/samdb/pkg/cmd

ok  	github.com/reetasingh/samdb/pkg/cmd	0.246s	coverage: 60.0% of statements
```

---

## Summary of Milestones:

- **Milestone 1**: Implemented basic `PING` command.
- **Milestone 2**: Added support for concurrent connections using Go-routines.
- **Milestone 3**: Optimized with AsyncIO and kernel queue for improved performance.
- **Milestone 4**: Implemented `SET`, `GET`, and `TTL` commands.
- **Milestone 5**: Introduced `TTL`, `EXPIRE`, and `DEL` commands for key expiry management.
- **Milestone 6**: Added unit test coverage to ensure code quality and stability.

---
