FROM golang:1.18.2 AS BUILD_IMAGE

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

RUN GOOS=linux go build -a -ldflags '-extldflags "-static"' -o challenge-backend

FROM alpine

ENV UID 1005
ENV GID 1005

RUN addgroup -S back -g $GID && \
    adduser -S back -G back -u $UID && \
    chown -R back:back /home/back
    
USER back

WORKDIR /home/back

COPY --from=BUILD_IMAGE --chown=back:back /app/challenge-backend /home/back/challenge-backend

EXPOSE 5050

ENTRYPOINT [ "/home/back/challenge-backend" ]