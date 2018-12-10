#!/bin/bash
# --rpm-posttrans
chkconfig --add spinel
touch /var/log/spinel.log
chown nginx /var/log/spinel.log
