# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o app

LABEL maintainer="WeiZhe.Chang <weizhe.chang@gmail.com>" \
  org.label-schema.name="WZ Practice"\
  org.lable-schema.vendor="Wei-Zhe Chang"

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "/bin/sh -c ./app", "-ping" ]
# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app