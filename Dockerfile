FROM golang:1.20.5-alpine3.18

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o ./dime ./cmd/dime

RUN apk add --update nodejs yarn
WORKDIR /app/frontend
RUN yarn install && yarn build

WORKDIR /app

CMD [ "./dime" ]
