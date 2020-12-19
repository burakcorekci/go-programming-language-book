FROM adoptopenjdk/openjdk8:alpine AS build
RUN mkdir /project
COPY gradlew /project/gradlew
RUN chmod +x /project/gradlew
COPY gradle /project/gradle
COPY settings.gradle /project/settings.gradle
COPY build.gradle /project/build.gradle
RUN cd project && ./gradlew
ADD main /project
RUN cd project && ./gradlew goBuild

FROM alpine:3.12 AS bin
COPY --from=build /project/.gogradle/app /opt/app
CMD /opt/app