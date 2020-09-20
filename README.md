# ping-monitor

`.env` sample  

```env
LINE_ACCESS_KEY=cavMa*****wkGv532ei******sEzK9dfKFGN****nXF
```

`server_list.json` sample  

```json
[
  {
    "name": "localhost 80 port",
    "host": "localhost",
    "port": 80,
    "network": "tcp",
    "repeat": 3,
    "timeout": 3000,
    "peroid": 180000
  },
  {
    "name": "localhost 443 port",
    "host": "127.0.0.1",
    "port": 443,
    "network": "tcp",
    "repeat": 3,
    "timeout": 1000,
    "peroid": 60000
  }
]
```

timeout and peroid is millisecond
