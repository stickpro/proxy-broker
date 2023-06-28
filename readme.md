## Go lang microservice for updating and initializing ports of microservices

1. Set configs in .env 
2. Set configs in configs/main.yml
3. run command ```Make build```
4. Create systemd service /etc/systemd/system proxybroker.service
```bash
[Unit]
Description=Proxy broker service updater
After=network.target

[Service]
Restart=always
RestartSec=3
WorkingDirectory=/home/server/proxybroker/
ExecStart=/home/server/proxybroker/.bin/app

[Install]
WantedBy=multi-user.target
```
5. ```systemct start proxybroker```