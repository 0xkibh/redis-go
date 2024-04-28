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