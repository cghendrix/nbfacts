FROM nginx:latest

COPY src/.output/public/ /usr/share/nginx/html/

RUN mv /etc/nginx/nginx.conf /etc/nginx/nginx.conf_orig

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/nbfacts.crt /etc/nginx/nbfacts.crt
COPY nginx/nbfacts.key /etc/nginx/nbfacts.key