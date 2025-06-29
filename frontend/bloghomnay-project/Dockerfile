FROM node:20-alpine AS base

WORKDIR /app

COPY package.json ./

RUN corepack enable pnpm
RUN pnpm install

# SAU KHI ĐÃ CÓ .dockerignore, bạn có thể safely COPY . .
COPY . .

RUN pnpm build

FROM nginx:1.24-alpine AS prod

COPY --from=base /app/dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80



CMD ["nginx", "-g", "daemon off;"]
