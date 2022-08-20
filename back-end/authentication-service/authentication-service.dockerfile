FROM alpine:latest
RUN mkdir /app
COPY authService /app
WORKDIR /app
CMD ["./authService"]
