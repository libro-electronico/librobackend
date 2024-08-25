# Build Stage
FROM maven:3.8.5-openjdk-11 AS builder

# Set working directory
WORKDIR /app

# Copy pom.xml and download dependencies
COPY pom.xml ./
RUN mvn dependency:go-offline

# Copy source code
COPY src ./src

# Package the application (Skip tests for faster builds)
RUN mvn clean package -DskipTests

# Runtime Stage
FROM openjdk:11-jre-slim

# Set working directory
WORKDIR /app

# Copy the jar file from the builder stage
COPY --from=builder /app/target/my-project-1.0-SNAPSHOT.jar /app/my-project.jar

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["java", "-jar", "/app/my-project.jar"]
