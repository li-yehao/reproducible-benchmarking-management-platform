#!/bin/bash

id=$1
cluster=$2
cmd_line=$3

if [ -z $id ]; then
  echo "id is null"
  exit 1
fi

if [ -z $cluster ]; then
  echo "cluster is null"
  exit 1
fi

if [ -z $cmd_line ]; then
  echo "cmd_line is null"
  exit 1
fi

# execute job and grep env at remote host
ssh "$cluster" "bash execute_job.sh -c $cmd_line -i $id"

# transfer 2 fileserver
path=~/rbmp/"$id"
fileserver=gateway:/var/www/html/nfsdata/rbmp/
# scp -3r "$cluster":"$path" "$fileserver"
echo "$cluster":"$path" "$fileserver"