[![Docker Pulls](https://img.shields.io/docker/pulls/glarfs/default-backend.svg)](https://hub.docker.com/r/glarfs/default-backend/)
[![Go Report Card](https://goreportcard.com/badge/github.com/glarfs/default-backend)](https://goreportcard.com/report/github.com/glarfs/default-backend)
[![license](https://img.shields.io/github/license/glarfs/default-backend.svg)](https://github.com/glarfs/default-backend/blob/master/LICENSE)
# Default backend for nginx-ingress on Kubernetes

This is a custom default backend for [nginx-ingress](https://github.com/kubernetes/ingress-nginx). The only reason for creating this project was to customize the 404 page and provide a better (and branded) user experience.

## Configuration

By default, the backend will listen on port `8080` and serve an `HTML` page with the error file at `/assets/404.html`. Override the default html page by mounting a config map with the html page at a different path.

**Important Note**

The final image is based on `scratch` image, so normal file system operations won't work. See the `Dockerfile` for more details

## Public Builds

Based on https://hub.docker.com/r/nuvo/default-backend

https://hub.docker.com/r/glarfs/default-backend


Build from Source

simple build:
```bash
docker build -t glarfs/default-backend .
```

multiarch build:
```bash
 docker buildx build -t glarfs/default-backend --platform=linux/arm,linux/arm64,linux/amd64 . --push
 ```
