FROM golang:1.20.1-alpine as builder
RUN apk add build-base

WORKDIR /app
ADD .. /app
RUN go build -o /exam_scores

FROM alpine:latest
COPY --from=builder /exam_scores /
EXPOSE 8080
CMD [ "/exam_scores" ]