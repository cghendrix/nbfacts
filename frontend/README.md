# Frontend App + Nginx

## To build for prod (full script in deploy.sh)
```
(cd src; npm i; npm run generate;)

 docker buildx build \
--platform linux/amd64 \
--push \
-t gcr.io/nicklebackfacts/nbfactsnginx .
```

## To build / run Nginx locally

```
docker build -t nbfactsnginx  . && docker run -it nbfactsnginx
```

## To build / run Nuxt app locally (will run on http://localhost:3000/)

```
(cd src; npm i; npm run dev;)
```