ARG session_key

FROM golang:1.19

ENV PORT=8085
ENV APP_ENV=local
ENV SESSION_KEY $session_key

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o main cmd/api/main.go

EXPOSE 8085

CMD ["go", "run", "cmd/api/main.go"]