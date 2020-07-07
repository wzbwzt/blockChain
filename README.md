## 区块链的简单代码实现，基于pow的共识算法
## LevelDB 简介
## DEMO（基于Web服务器的Pow实例）

- 基于gorilla/mux包实现的http请求的demo（格式话返回所有的区块链信息）；

- GET请求实现查看所有的区块链信息；

- POST请求增加新节点；

服务端

```bash
D:\www\Blockchain_go\pow>go run main.go
2020/07/05 22:41:53 REST service on 127.0.0.1:8001
2020/07/05 22:42:20 Add new node success
2020/07/05 22:43:00 Add new node success
```

客户端

```bash
D:\www\Blockchain_go\pow>curl -X GET http://localhost:8001/blockChain/list
[
        {
                "Index": 0,
                "PreHash": "0",
                "HashCode": "38fc3af258669791206922284800c53a0eab92b0d602253e447a71b802f5c77f",
                "TimeStamp": 1593960113,
                "Diff": 3,
                "Data": "创世区块",
                "Nonce": 123
        },
        {
                "Index": 1,
                "PreHash": "38fc3af258669791206922284800c53a0eab92b0d602253e447a71b802f5c77f",
                "HashCode": "0001bed84d1c2c3a2913f067865824effe66d20152eff783af17b532b2fe577d",
                "TimeStamp": 1593960140,
                "Diff": 3,
                "Data": "second node",
                "Nonce": 132996
        },
        {
                "Index": 2,
                "PreHash": "0001bed84d1c2c3a2913f067865824effe66d20152eff783af17b532b2fe577d",
                "HashCode": "0001f89dc9b82c524c0e3e0c36fd56749c97ddc7fb2b1c2d3bb5baeca1142698",
                "TimeStamp": 1593960180,
                "Diff": 3,
                "Data": "third node",
                "Nonce": 105283
        }
]
```

```bash
D:\www\Blockchain_go\pow>curl -X POST -d  "{\"data\":\"third node\"}"   http://localhost:8001/blockChain/add
{
        "result": [
                {
                        "Index": 0,
                        "PreHash": "0",
                        "HashCode": "38fc3af258669791206922284800c53a0eab92b0d602253e447a71b802f5c77f",
                        "TimeStamp": 1593960113,
                        "Diff": 3,
                        "Data": "创世区块",
                        "Nonce": 123
                },
                {
                        "Index": 1,
                        "PreHash": "38fc3af258669791206922284800c53a0eab92b0d602253e447a71b802f5c77f",
                        "HashCode": "0001bed84d1c2c3a2913f067865824effe66d20152eff783af17b532b2fe577d",
                        "TimeStamp": 1593960140,
                        "Diff": 3,
                        "Data": "second node",
                        "Nonce": 132996
                },
                {
                        "Index": 2,
                        "PreHash": "0001bed84d1c2c3a2913f067865824effe66d20152eff783af17b532b2fe577d",
                        "HashCode": "0001f89dc9b82c524c0e3e0c36fd56749c97ddc7fb2b1c2d3bb5baeca1142698",
                        "TimeStamp": 1593960180,
                        "Diff": 3,
                        "Data": "third node",
                        "Nonce": 105283
                }
        ],
        "status": "add success!!!"
}
```



