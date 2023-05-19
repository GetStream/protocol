FROM swift:latest

ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y --no-install-recommends wget git curl unzip software-properties-common apt-transport-https ca-certificates gnupg lsb-release

ARG TARGETPLATFORM
ARG TARGETARCH
RUN ARCH=${TARGETPLATFORM#*/} && TARGETARCH=${ARCH}

# Install Go
RUN wget https://golang.org/dl/go1.20.4.linux-$TARGETARCH.tar.gz
RUN tar -C /usr/local -xzf go1.20.4.linux-$TARGETARCH.tar.gz
RUN rm go1.20.4.linux-$TARGETARCH.tar.gz
ENV PATH="$PATH:/usr/local/go/bin"

# Install Dart
ENV DART_ARCH="arm64"
RUN if [ "$TARGETARCH" = "amd64" ]; then export DART_ARCH="x64"; fi
RUN echo $TARGETARCH $DART_ARCH
RUN wget https://storage.googleapis.com/dart-archive/channels/stable/release/3.0.0/sdk/dartsdk-linux-$DART_ARCH-release.zip
RUN unzip dartsdk-linux-$DART_ARCH-release.zip
RUN rm dartsdk-linux-$DART_ARCH-release.zip
RUN mv dart-sdk /opt/dart-sdk
ENV PATH="$PATH:/opt/dart-sdk/bin"
RUN echo 'export PATH="$PATH:/opt/dart-sdk/bin"' >> ~/.bashrc

# Install TypeScript + NodeJS + Yarn
RUN apt-get update && apt-get install -y --no-install-recommends nodejs
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update && apt-get install -y --no-install-recommends yarn

COPY . /home/
WORKDIR /home
RUN ./install.sh
ENTRYPOINT ["/home/docker_entrypoint.sh"]
