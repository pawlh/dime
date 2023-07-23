FROM golang:1.20.6-alpine3.17

WORKDIR /app

#COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN go build -o ./dime ./cmd/dime

RUN apk add --update nodejs yarn
WORKDIR /app/frontend
RUN yarn install && yarn build

WORKDIR /app

CMD [ "./dime" ]
