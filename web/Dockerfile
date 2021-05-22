FROM node:15 AS BUILDER

WORKDIR /app

COPY ./yarn.lock ./
COPY ./package.json ./

RUN yarn install --prod

COPY . .

RUN yarn build

FROM nginx:stable-alpine AS RUNNER

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY --from=BUILDER /app/dist /var/www/html/