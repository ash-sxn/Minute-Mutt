#!/bin/bash
# entrypoint.sh

# Update the resolv.conf file to use the DNS server of the host machine
echo "nameserver 8.8.8.8" > /etc/resolv.conf

# Update the PATH environment variable & # Update the cron schedule based on the environment variable
echo "$CRON_SCHEDULE export PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin; export OUTPUT_DIR="$OUTPUT_DIR"; export MAX_RESOLUTION="$MAX_RESOLUTION"; cd /app && ./main >> /var/log/cron.log 2>&1" | crontab -


# Start the cron service
service cron start
./main

sleep infinity
