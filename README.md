# BMA: Bare Minimum API

## Prerequisite's
```
Needs docker desktop or minikube along with nginx ingress controller.
* https://docs.docker.com/docker-for-mac/install/
* Install nginx controller : https://docs.nginx.com/nginx-ingress-controller/installation/installation-with-helm/#adding-the-helm-repository
```

## Run it locally on minikube or docker desktop.
```
git clone https://github.com/mohnishanandinc/bare-minimum-api.git && cd bare-minimum-api/charts/bma
helm upgrade -i app . -f values.yaml -n coalationinc --debug
```

## Build the image
```
docker build . -t mohnishanand/bma && docker push mohnishanand/bma:latest && \
date=$(date|tr -d ' '|tr -d :) && echo $date && \
docker tag mohnishanand/bma:latest mohnishanand/bma:$date && \
docker push mohnishanand/bma:$date 
```
## Run it locally without docker-desktop or minikube
```
Make sure you have built the docker image and have it on your local machine.
git clone https://github.com/mohnishanandinc/bare-minimum-api.git && cd bare-minimum-api 
docker build . -t mohnishanand/bma
docker run -p 8080:8080h mohnishanand/bma:latest
``` 
## Access the timestamp API
```
If running on K8s, open the browser and access localhost/timestamp. You can also curl the endpoint.
If running the docker conatiner type : curl localhost:8080/timestamp
```
