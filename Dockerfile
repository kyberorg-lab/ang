FROM node:12.11 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@10.0.5
COPY webapp /webapp
WORKDIR webapp
RUN npm install && ng build --prod

FROM golang:1.14.7-alpine AS GO_BUILD
COPY server /
WORKDIR /

COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy sources from current dir to Container
COPY server /
RUN CGO_ENABLED=0 go build  -a -ldflags '-extldflags "-static"' -o /server main.go

FROM alpine:3.12
COPY --from=ANGULAR_BUILD /webapp/dist/webapp/* ./webapp/dist/webapp/
COPY --from=GO_BUILD /server ./
CMD ./server
EXPOSE 8080
