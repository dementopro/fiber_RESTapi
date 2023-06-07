# Tenant Operations Controller

### Build the docker image

```azure
docker build -t <dockerusername>/<dockerimage>:<dockertag> -f Dockerfile .
docker push <dockerusername>/<dockerimage>:<dockertag>
```
### Running the application

```
go run cmd/main.go
```