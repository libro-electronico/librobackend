# Use Maven image to build the project
FROM maven:3.8.5-openjdk-11 AS builder

# Set working directory
WORKDIR /app

# Copy pom.xml and download dependencies
COPY pom.xml ./
RUN mvn dependency:go-offline

# Copy source code
COPY src ./src

# Build the application
RUN mvn clean package -DskipTests

# Use a lightweight OpenJDK image to run the application
FROM openjdk:11-jre-slim

# Copy the JAR file from the builder stage
COPY --from=builder /app/target/librobackend-1.0-SNAPSHOT.jar /app/librobackend.jar

# Set the command to run the application
CMD ["java", "-jar", "/app/librobackend.jar"]
