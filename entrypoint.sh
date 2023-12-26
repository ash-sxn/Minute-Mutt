#!/bin/bash
# entrypoint.sh

# Update the cron schedule based on the environment variable
echo "$CRON_SCHEDULE cd /app && ./main >> /var/log/cron.log 2>&1" | crontab -

# Start the cron service
service cron start
./main
# Start supervisord to manage the application
/usr/bin/supervisord -c /etc/supervisor/conf.d/supervisord.conf
