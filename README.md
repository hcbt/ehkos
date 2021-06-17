# Ehkos

## Running

Ehkos requires go 1.15 or later, docker and docker-compose to run.

To run vernemq MQTT broker and ehkos server:

```sh
cd ehkos
docker-compose up
```

To run ehkos client:

```sh
cd ehkos/src
go run main.go run client -f files/screenshot.png -b tcp://localhost:1883 -c test_client -t test
```

Client can be used for any file although sample file is provided in src/files folder. 

## Todo

* Option to define topic for server outside docker file so you wouldn't need to rebuild server container everytime after changing it.
* Codebase improvements.
* Explanation for every commandline arguments flag.
* Sending file data through MQTT payload?