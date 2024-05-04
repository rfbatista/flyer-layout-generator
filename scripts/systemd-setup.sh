#!/bin/bash
cp ./server.service /etc/systemd/system/server.service
sudo systemctl enable server
sudo systemctl start server
sudo systemctl status server
sudo systemctl restart server
sudo systemctl daemon-reload
