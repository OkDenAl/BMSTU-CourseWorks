#/bin/bash

# add an auth key for the replica set
echo "12345678" > "/tmp/key.file"
chmod 600 /tmp/key.file
chown 999:999 /tmp/key.file
