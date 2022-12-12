# Backend Service

## To build for prod (full script in deploy.sh)
```
docker buildx build \
--build-arg DB_USER=${DB_USER} \
--build-arg DB_PASS=${DB_PASS} \
--platform linux/amd64 \
--push \
-t gcr.io/nicklebackfacts/nbfactsserver .
```

## To build locally

#### Server
```
docker build --build-arg COMMAND=server --build-arg DB_USER=${DB_USER} --build-arg DB_PASS=${DB_PASS} -t nbfactsserver . && docker run -p 80:80 -it nbfactsserver
```

#### Worker
```
docker build --build-arg COMMAND=worker --build-arg DB_USER=${DB_USER} --build-arg DB_PASS=${DB_PASS} -t nbfactsworker . && docker run -it nbfactsworker
```