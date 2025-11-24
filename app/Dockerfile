# vue
# --- local ---
FROM node:20-alpine as build-vue
WORKDIR /app
COPY . .
RUN corepack enable && \
  yarn install && \
  NODE_OPTIONS=--openssl-legacy-provider yarn run build

# --- production ---
FROM nginx:alpine as production-vue
COPY --from=build-vue /app/dist /usr/share/nginx/html
COPY --from=build-vue /app/nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
