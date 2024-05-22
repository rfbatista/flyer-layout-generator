# flyer-layout-generator

- python -m venv venv
- source ./venv/bin/activate
- ./venv/bin/pip install -U pip setuptools
- ./venv/bin/pip install poetry
- ./venv/bin/poetry requirements
- ./venv/bin/pip install -r requirements.txt

# Installing new packages



# Database migration 

Install atlas

```
curl -sSf https://atlasgo.sh | sh
```

Install sqlc

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

# Log Managment

```
docker plugin install grafana/loki-docker-driver:2.9.1 --alias loki --grant-all-permissions
```

1. sudo mkdir /etc/docker/
2. At location /etc/docker/daemon.json paste the below code:
```
 {
    "log-driver": "loki",
    "log-opts": {
      "loki-url": "http://localhost:3100/loki/api/v1/push",
      "loki-batch-size": "400",
    }
  }
```
3. sudo systemctl restart docker


# Connecting to Lightsail

```
chmod 400 ./ssh-key.pem
make ssh
```

# Code Templates

- nodejs 18.3.0
- npm i hygen -g
