#!/bin/bash

sudo su

sed -i 's/us.archive.ubuntu.com/mirrors.us.kernel.org/g' /etc/apt/sources.list

sudo apt-get update
sudo apt-get install -y rabbitmq-server golang

sudo rabbitmq-plugins enable rabbitmq_management
sudo service rabbitmq-server restart
