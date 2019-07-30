# --- local ---
# vue
FROM node:8.15.0-alpine as build-vue
WORKDIR /app
COPY . .
RUN apk update && \
  npm install -g npm && \
  npm install -g yarn && \
  yarn install && \
  yarn run build

# --- production ---
# vue
FROM nginx:1.15.12-alpine as production-vue
COPY --from=build-vue /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
