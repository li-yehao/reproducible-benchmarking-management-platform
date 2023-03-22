#!/bin/bash

while getopts ":c::l:" opt
do
  case $opt in
    c)
    cmd_line="$OPTARG"
    ;;
    i)
    id="$OPTARG"
    ;;
    ?)
    exit 1
    ;;
  esac
done

# execute job


# grep env
cpu=$(echo '{"cpu": {'$(lscpu | awk -F ': +' '{if(NR>1) printf ",\n"; printf "\"%s\": \"%s\"", $1, $2}' | sed 's/\" *\"/: null/g')'}}' | jq .)
mem=$(awk 'BEGIN{print "{"} NR>0{if(NR>1)printf ",\n"; printf "\"%s\": %d", $1, $2} END{print "\n}"}' /proc/meminfo | jq '{ "mem": . }')
cilium=$(cilium version | jq -Rn 'reduce inputs as $line ({}; . + ($line | split(": ") | {(.[0]): .[1]}))' | jq '{ "cilium": . }')
calico=$(calicoctl version | jq -Rn 'reduce inputs as $line ({}; . + ($line | split(": ") | {(.[0]): .[1]}))' | jq '{ "calico": . }')
pmu=$(jq -n --arg pmu "$(cat /sys/devices/cpu/caps/pmu_name)" '{"pmu":$pmu}')
k8s=$(kubectl version -o json)
node=$(jq '{node: .}' <<< $(kubectl get node -o json))
status=$(jq -s '.' /tmp/perfkitbenchmarker/runs/hotjag/completion_statuses.json | jq '{ "status": . }')
if [ "$status" == "SUCCEEDED" ]; then
  result=$(jq -s '.' /tmp/perfkitbenchmarker/runs/hotjag/perfkitbenchmarker_results.json | jq '{ "result": . }')
else
  result=$(jq -s '.' /tmp/perfkitbenchmarker/runs/hotjag/completion_statuses.json | jq '{ "status_detail": .[0].status_detail }')
fi

json_array=("$cpu" "$mem" "$cilium" "$calico" "$pmu" "$k8s" "$result" "$node" "$status")
merged_json=$(jq -s 'reduce .[] as $item ({}; . * $item)' <<< "${json_array[@]}")


echo "$merged_json" > results.json

summary=$(cat results.json | jq '{ pmu: .pmu, workload: .result[1].test, throughput: .result[1].value,
host: .node.items[0].metadata.name, status: .status[0].status,
ns: .result[1].metadata.namespace_count, os: .result[1].metadata.client_os_info, 
kernel: .result[1].metadata.client_kernel_release, duration: .result[1].metadata.duration, 
request: .result[1].metadata.request_count, rate: .result[1].metadata.rate, 
cri: .node.items[1].status.nodeInfo.containerRuntimeVersion,
threads: .result[1].metadata.threads, connections: .result[1].metadata.connections }')

echo "$summary" > summary.json

echo "$summary" | jq

# workload=$(cat summary.json | jq .summary.workload)
# host=$(cat summary.json | jq .summary.host)
# timestamp=$(cat results.json | jq .result[2].timestamp)
folder="$id"
mkdir -p "$folder"
mv summary.json results.json "$folder"