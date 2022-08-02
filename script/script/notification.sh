#!/bin/bash
notifications=$(dunstctl history | jq '.["data"][0][] | [.appname.data, .timestamp.data, .body.data, .summary.data] | join("-=-") | @base64')
for row in $(echo $notifications | awk '{for(i=1;i<=20;i++) print $i}')
do
    current_time=$(cat /proc/uptime | head -n1 | awk '{print $1;}')
    current_time=${current_time::-3}
    string=$(echo $row | xargs echo | base64 --decode)
    appname=$(echo $string | awk -F'[-][=][-]' '{print $1}')
    timestamp=$(echo $string | awk -F'[-][=][-]' '{print $2}')
    timestamp=${timestamp::-6}
    message=$(echo $string | awk -F'[-][=][-]' '{print $3}')
    summary=$(echo $string | awk -F'[-][=][-]' '{print $4}')
    echo -en "${appname} - $((current_time - timestamp))s Ago\n${message}\n${summary}\0icon\x1f$(echo $appname | awk '{print tolower($0)}')|"
done
