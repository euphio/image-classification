FROM golang:1.22

WORKDIR /app

ENV PATH="/go/bin:$PATH"

COPY . .

RUN go mod tidy

EXPOSE 8080
CMD ["go", "run", "main.go"]
