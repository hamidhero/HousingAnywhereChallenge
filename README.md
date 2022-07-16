# HousingAnywhereChallenge

## How to Run with Dockerfile

```bash
docker build --tag dns-location .
```
This above command build an image of the program including running tests beforehand.

and then this command below run the image as a container in a detached mode.

```bash
docker run -d -p 3030:3030 dns-location
```

Now project is up on port 3030 of your machine. You can access its endpoint from postman or running curl below:

```bash
curl --location --request POST '127.0.0.1:3030/api/location' \
--header 'Content-Type: application/json' \
--data-raw '{
    "x": "23.4",
    "y": "13.4",
    "z": "14.1",
    "vel": "19"
}'
```
