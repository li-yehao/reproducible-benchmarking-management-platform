#!/bin/bash

while getopts ":c::i::r::l::f:" opt
do
  case $opt in
    c)
    cmd_line="$OPTARG"
    ;;
    l)
    cluster="$OPTARG"
    ;;
    i)
    id="$OPTARG"
    ;;
    r)
    reason="$OPTARG"
    ;;
    f)
    config="$OPTARG"
    ;;
    ?)
    exit 1
    ;;
  esac
done

if [ -z "$id" ]; then
  echo "id is null"
  exit 1
fi

if [ -z "$cluster" ]; then
  echo "cluster is null"
  exit 1
fi

if [ -z "$cmd_line" ]; then
  echo "cmd_line is null"
  exit 1
fi

if [ -z "$reason" ]; then
  echo "reason is null"
  exit 1
fi

# execute job and grep env at remote host, todo: move to ~
echo "$config" > pkbConfig.yaml
scp pkbConfig.yaml "ansible@$cluster":~
mv pkbConfig.yaml
scp ~/lyh/reproducible-benchmarking-management-platform/server/service/Job/execute_job.sh "ansible@$cluster":~
ssh "ansible@$cluster" "bash execute_job.sh -c '$cmd_line' -i $id -r '$reason'"
echo "execute job finish!"

# transfer to fileserver
path=/home/ansible/rbmp/"$id"_"$reason"
fileserver=gateway:/var/www/html/nfsdata/rbmp/
scp -3r "$cluster":"$path" "$fileserver"
echo "transfer files finish!"