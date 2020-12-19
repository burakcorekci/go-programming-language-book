# TODO: FIX
FROM adoptopenjdk/openjdk8:alpine

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
        build-essential

RUN wget -q https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz \
    && rm go1.15.6.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

COPY main /app/src/main
COPY gradle /app/src/gradle
COPY build.gradle /app/src/build.gradle
COPY gradlew /app/src/gradlew

CMD ["bash"]