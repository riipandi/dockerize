# -----------------------------------------------------------------------------
# Builder
# -----------------------------------------------------------------------------
FROM alpine:latest AS builder
ARG THTTPD_VERSION=2.29
RUN apk add gcc musl-dev make
RUN wget http://www.acme.com/software/thttpd/thttpd-${THTTPD_VERSION}.tar.gz \
  && tar xzf thttpd-${THTTPD_VERSION}.tar.gz \
  && mv /thttpd-${THTTPD_VERSION} /thttpd

# Compile thttpd to a static binary which we can copy around
RUN cd /thttpd \
  && ./configure \
  && make CCOPT='-O2 -s -static' thttpd

# Create a non-root user to own the files and run our server
RUN adduser -D static

# -----------------------------------------------------------------------------
# Runner from scratch
# -----------------------------------------------------------------------------
FROM scratch

# Copy required binaries and files to the scratch image
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /thttpd/thttpd /
COPY ./dist .

USER static
WORKDIR /home/static

EXPOSE 3001

CMD ["/thttpd", "-D", "-h", "0.0.0.0", "-p", "3001", "-d", "/home/static", "-u", "static", "-l", "-", "-M", "60"]
