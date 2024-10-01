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
