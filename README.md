# numbers-service
A REST-ish and gRPC capable service to manipulate numbers.

Things to implement
 - CI/CD
 - Dockerfile
 - gRPC
 - Instrumenting
 - Logging
 - Tracing
 - Circuit Breaking
 - Rate Limiting
 - Authentication
 - Distributed cache (memcached ? redis ?)

## Running in Docker
``` bash
https://github.com/Tiny-Paws/numbers-service.git
cd numbers-service
docker build . -t numbers-service
docker run --rm -it -p 8080:8080 numbers-service
```
Support for environment variables planned

## Testing
```bash
cd test
go test ./numbersservice
```