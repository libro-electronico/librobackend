FROM maven:3.8.1-jdk-11 AS build
WORKDIR /app
COPY . /app
RUN mvn clean package

FROM openjdk:11-jre-slim
COPY --from=build /app/target/yourapp.jar /app/yourapp.jar
CMD ["java", "-jar", "/app/yourapp.jar"]
