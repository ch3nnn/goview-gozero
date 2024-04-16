FROM nginx:1.25.4

WORKDIR /usr/share/nginx/html/

USER root

COPY ./nginx /etc/nginx/

COPY ./dist  /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
