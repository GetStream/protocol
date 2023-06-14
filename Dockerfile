FROM --platform=$TARGETPLATFORM swift:latest

ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y --no-install-recommends wget git curl unzip software-properties-common apt-transport-https ca-certificates gnupg lsb-release && rm -rf /var/lib/apt/lists/*

ARG TARGETPLATFORM
ARG TARGETARCH
RUN ARCH=${TARGETPLATFORM#*/} && TARGETARCH=${ARCH}

WORKDIR /home

# Install Go
RUN wget https://golang.org/dl/go1.20.4.linux-$TARGETARCH.tar.gz
RUN tar -C /usr/local -xzf go1.20.4.linux-$TARGETARCH.tar.gz
RUN rm go1.20.4.linux-$TARGETARCH.tar.gz
ENV PATH="$PATH:/usr/local/go/bin"

# Install Dart
#RUN if [ "$TARGETARCH" = "amd64" ]; then \
#      echo "amd64"; \
#      wget https://storage.googleapis.com/dart-archive/channels/be/raw/latest/sdk/dartsdk-linux-x64-release.zip && \
#      unzip dartsdk-linux-x64-release.zip && rm dartsdk-linux-x64-release.zip; \
#    else  \
#      echo "arm64"; \
#      wget https://storage.googleapis.com/dart-archive/channels/be/raw/latest/sdk/dartsdk-linux-arm64-release.zip && \
#      unzip dartsdk-linux-arm64-release.zip && rm dartsdk-linux-arm64-release.zip; \
#    fi
#
#RUN mv dart-sdk /opt/
#ENV PATH="$PATH:/opt/dart-sdk/bin"
#RUN echo 'export PATH="$PATH:/opt/dart-sdk/bin"' >> ~/.bashrc

# Install TypeScript + NodeJS + Yarn
#RUN apt-get update && apt-get install -y --no-install-recommends nodejs && rm -rf /var/lib/apt/lists/*
#RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
#RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
#RUN apt-get update && apt-get install -y --no-install-recommends yarn && rm -rf /var/lib/apt/lists/*

COPY . /home/
RUN ./install.sh
ENTRYPOINT ["/home/docker_entrypoint.sh"]
