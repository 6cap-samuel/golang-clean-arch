FROM openjdk:19-jdk
ARG JAR_FILE=target/*.jar
COPY ${JAR_FILE} app.jar
COPY src/main/resources/application.properties application.properties

ENTRYPOINT ["java","-jar","/app.jar"]