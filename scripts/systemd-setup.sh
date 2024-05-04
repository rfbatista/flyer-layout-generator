#!/bin/bash
cp ./server.service /etc/systemd/system/server.service
sudo systemctl enable url-shortener
sudo systemctl start url-shortener
sudo systemctl status url-shortener
sudo systemctl restart url-shortener
sudo systemctl daemon-reload
