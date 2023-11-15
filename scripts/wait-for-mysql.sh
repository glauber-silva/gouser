#!/bin/sh
# scripts/wait-for-mysql.sh

set -e

host="$1"
shift
cmd="$@"

until echo 'SELECT 1' | mysql -h"$host" -P"$MYSQL_PORT" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" --silent; do
  echo 'waiting for mysql...'
  sleep 1
done

echo 'mysql is ready, starting user-service...'
$cmd || echo "Failed to execute $cmd"