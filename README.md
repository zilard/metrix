# Table of Contents

- [Prerequisites](#Prerequisites)
- [Install and Deploy](#Install and Deploy)
- [metrix Server Usage](#metrix Server Usage)
- [mx API Client Usage](#mx API Client Usage)



# Prerequisites

Minikube deployed

run the following commands to unset proxies and setup the minikube context

    unset http_proxy
    unset https_proxy
    kubectl config use-context minikube
    eval $(minikube docker-env)




# Install and Deploy

Clone repository

    cd metrix

    make deploy




# metrix Server Usage

Start the server

Listening on default port: 8080

    ./build/metrix/metrix        

OR

Listening on specified port:   

    ./build/metrix/metrix -p <PORT>



You don't need to run the server separately in case you executed "make deploy"
"make deploy" will deploy the server in a kubernetes deployment, and expose it through a NodePort service



# mx API Client Usage

Send metrics using the mx API client:

    ./build/bin/mx send nm -i 192.168.39.210 -o 32202 -n node3 -t 100 -c 67 -m 80

    URL: http://192.168.39.210:32202/v1/metrics/node/node2/

    NODE MEASUREMENT SENT: 

    {
      "timeslice": 360,
      "cpu": 45,
      "mem": 70
    }


Get analytics using the mx API client:


    ./build/bin/mx get na -i 192.168.39.210 -o 32202

    URL: http://192.168.39.210:32202/v1/analytics/nodes/average?timeslice=60.000000

    NODE ANALYTICS: 

    {
      "timeslice": 60,
      "cpu_used": 45,
      "mem_used": 70
    }



If no IP address is specified through "-i" flag and/or no port is specified through "-o" flag
then the mx API Client will use 127.0.0.1 as IP address and default port 8080



