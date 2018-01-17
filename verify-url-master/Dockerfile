FROM alpine:3.5
MAINTAINER Kevin Stock <kevin@toolhouse.com>

# SSL CA Root Certs
RUN apk --no-cache add ca-certificates

# Labels: http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="verify-url" \
      org.label-schema.description="A tool for verifying that a URL returns a 200 response" \
      org.label-schema.url="https://github.com/toolhouse/verify-url" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/toolhouse/verify-url" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

WORKDIR /
ADD ./verify-url-linux_amd64 /verify-url
EXPOSE 80

# The environment variables are used to configure the container at runtime:
# ENV URL http://www.example.com/

CMD ["/verify-url"]