## Redis in Go

### Get Started
1. Setup Redis CLI, which we are going to use to communicate with our redis
```
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

sudo apt-get update
sudo apt-get install redis
```

2. Stop the actual Redis server
```
sudo systemctl stop redis
```

### Tests
1. To make all test file executable 
```
chmod +x test/*.sh
```


### To Know
1. Redis use Redis Serialization Protocol (RESP) to communicate. It can be string, array, numbers or many other way. So, Redis recieves and send messages in RESP format. For example: String "OK" will be "+OK\r\n"

Check out here for [more](https://redis.io/docs/latest/develop/reference/protocol-spec/)

2. goroutine
ou can refer to basic golang documents for this, for example, one item in Google search keyword goroutine is golang concurrency.

3. Redis command are case-insensitive. `ECHO`, `echo` and `EcHo` are all valid commands.

### Possible Queries
1. Why we use `make([]byte, 128)` instead of `[128]byte`?

Ans => Make is used to create slice here. And in Go, arrays have a fixed size, while slices are dynamically-sized and provide more flexibility. What is slice? Slice is a data structure that consists of a pointer to an underlying array, a length, and a capacity. 

Follow up question: What happens if space available for array is no more available but we need to expand?
- A new, larger array is allocated in memory to accommodate the additional elements.
- The contents of the original array are copied over to the new, larger array.
- The slice is then updated to point to the new, larger array.
[Read More Here](https://go101.org/article/container.html#:~:text=The%20length%20and%20capacity%20of,be%20viewed%20as%20dynamic%20arrays.)
