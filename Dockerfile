FROM scratch

EXPOSE 8080

COPY app /app

ENTRYPOINT ["/app"]
