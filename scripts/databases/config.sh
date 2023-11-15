#!/bin/bash

# Connect to the MySQL database and run the SQL command
# That's not the best way to do it, but it's a way to give more security to the admin user
mysql -u root -p$MYSQL_ROOT_PASSWORD -D gouserdb -e "

-- Grant all privileges to the new user
GRANT ALL PRIVILEGES ON *.* TO '$MYSQL_USER'@'%' WITH GRANT OPTION;
"