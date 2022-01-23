FROM --platform=linux/x86_64 openjdk:18-ea-jdk-alpine3.15
ARG JAR_FILE=target/*.jar
COPY ${JAR_FILE} app.jar
COPY src/main/resources/application.properties application.properties

ENTRYPOINT ["java","-jar","/app.jar"]