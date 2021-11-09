# Mercedes Benz TechGig API

Based on Go > gorila mux

## Requrements 
docker 

## System requrements 
### for best performance 
2 core processor, 512 mb of RAM

### Minimum requrements 
1 core processor

## Steps to run the application

1. sudo docker build -t mbrdi .
2. sudo docker run -p 4000:4000 -it mbrdi

Here you will get a local ip address. Please open that ip address in browser to check if the api is accessable and working successfully (health check route).

Heathcheck api:
```
url: http://localhost:4000/api/healthcheck
method: GET 
```

Use postman to get results
```
url: http://localhost:4000/api/plantrip
method: POST 
content-type: application/json
body: { "vin": <vin>, "source": <source>, "destination": <destination> }
```