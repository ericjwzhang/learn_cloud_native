 #! /bin/bash
 docker ps -a | grep -vE "grep|CONTAINER" |awk '{print $1}'|xargs docker rm --force
