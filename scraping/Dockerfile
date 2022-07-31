FROM golang:1.17

WORKDIR /app
COPY . .
RUN go build -v -o ./bin/app
CMD [ "./bin/app" ]