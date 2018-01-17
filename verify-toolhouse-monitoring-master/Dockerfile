FROM alpine:3.5
MAINTAINER Kevin Stock <kevin@toolhouse.com>

# SSL CA Root Certs
RUN apk --no-cache add ca-certificates

# Labels: http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="verify-toolhouse-monitoring" \
      org.label-schema.description="A tool for verifying the logging endpoints added by Toolhouse.Monitoring" \
      org.label-schema.url="https://github.com/toolhouse/verify-toolhouse-monitoring" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/toolhouse/verify-toolhouse-monitoring" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

WORKDIR /
ADD ./verify-toolhouse-monitoring-linux_amd64 /verify-toolhouse-monitoring
EXPOSE 80

# The environment variables are used to configure the container at runtime:
# ENV URL http://www.example.com/

CMD ["/verify-toolhouse-monitoring"]