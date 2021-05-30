#Stage-1
FROM golang:1.16 as builder
COPY ./app/main.go ./
RUN go build -o /gitops-go-app ./main.go

#Stage-2
FROM gcr.io/distroless/base
EXPOSE 8080
COPY --from=builder /gitops-go-app /.
ENTRYPOINT ["./gitops-go-app"]