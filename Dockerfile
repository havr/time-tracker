FROM golang:1.13 AS backend
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN make build

FROM node:13 AS frontend
WORKDIR /build
COPY ui .
RUN yarn install
RUN make build

FROM alpine:3.10
WORKDIR /app
COPY --from=backend /build/tracker tracker
COPY --from=backend /build/migrations migrations
COPY --from=frontend /build/build ui
ENV TT_STATIC_DIR /app/ui/
ENV TT_MIGRATE_FROM /app/migrations
EXPOSE 8080
CMD ["/app/tracker"]
