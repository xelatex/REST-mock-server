# REST-mock-server
Mock REST server for E2E test. Run mock server as a seperated server with configuration, mock server responses with mock result to any REST request.

## Build
git clone https://github.com/xelatex/REST-mock-server.git REST-mock-server

cd REST-mock-server

go build

## Get help
./REST-mock-server.git -h

## Run
mkdir config_data

./REST-mock-server.git -c ./config_data

## Usage
Use mock server (default: 0.0.0.0:1080) to get mock result, and use control server (default: 0.0.0.0:1070) to get/add/modify mock server behaviour.

### Use mock server
#### Get mock rules
Issue a REST GET request with header "method:METHOD" to get the current mock rules.

curl -X GET -H "method:GET" -i "http://localhost:1070/abc"

#### Add/modify mock rules
Issue a REST POST request with header "method:METHOD" to designate the type of mock rules, the request body should contains data of the reply.

curl -X POST -H "method:GET" --data '{"Status":200,"Content":"{'name':'john'}","Header":{"A":["a1","a2"],"B":["b"]},"ContentType":"text/json; charset=utf-8"}' -i "http://localhost:1070/abc"

#### Delete mock rules
Issue a REST DELETE request with header "method:METHOD" to delete a mock rule

curl -X DELETE -H "method:GET" -i "http://localhost:1070/abc"


LISCENCE: MIT Lisence
