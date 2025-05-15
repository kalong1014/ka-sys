# 构建阶段
FROM node:18-alpine as builder

WORKDIR /app

COPY package*.json ./
RUN npm install --production=false

COPY . .
RUN npm run build

# 生产环境
FROM nginx:1.23-alpine

COPY --from=builder /app/dist /var/www/merchant-frontend/dist
COPY nginx/conf.d/merchant.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]