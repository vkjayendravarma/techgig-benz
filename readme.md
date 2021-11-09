# Mercedes Benz TechGig API

Based on Go > gorila mux

## Requrements 
docker 

## System requrements 
### for optimal performance 
2 core processor, 512 mb of RAM

### Minimum requrements 
1 core processor

## Steps to run the application

1. sudo docker build -t evpitstops .
2. sudo docker run -p 5000:5000 -it evpitstops

Here you will get a local ip address. Please open that ip address in browser to check if the api is accessable and working successfully (health check route).

Use postman to get results
```
url: http://<IP Address>:5000/getrouteplan
method: post 
content-type: application/json
body: { "vin": <vin>, "source": <source>, "destination": <destination> }
```

## Run with nodemon
nodemon --exec go run main.go --signal SIGTERM