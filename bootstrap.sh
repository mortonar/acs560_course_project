#!/bin/bash

if [[ $EUID -ne 0 ]]; then
    echo "This script must be run as root"
    exit 1
fi

dnf install -y postgresql{,-contrib,-libs,-server}
/usr/bin/postgresql-setup --initdb
sudo systemctl enable postgresql
sudo systemctl start postgresql

sudo -u postgres psql -c "CREATE ROLE booktracker WITH LOGIN PASSWORD 'booktracker';"
sudo -u postgres psql -c "CREATE DATABASE booktracker;"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE booktracker to booktracker;"
