# vue
# --- local ---
FROM node:8.15.0-alpine as build-vue
WORKDIR /app
COPY . .
RUN apk update && \
  npm install -g npm && \
  npm install -g yarn && \
  yarn install && \
  yarn run build

# --- production ---
FROM nginx:1.15.12-alpine as production-vue
COPY --from=build-vue /app/dist /usr/share/nginx/html
COPY --from=build-vue /app/nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
