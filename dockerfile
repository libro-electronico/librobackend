# Gunakan base image Maven untuk build phase
FROM maven:3.8.5-openjdk-11 AS builder

# Set working directory
WORKDIR /app

# Salin pom.xml dan unduh dependensi
COPY pom.xml ./
RUN mvn dependency:go-offline

# Salin source code
COPY src ./src

# Build aplikasi
RUN mvn clean package -DskipTests

# Gunakan base image yang lebih ringan untuk runtime
FROM openjdk:11-jre-slim

# Salin JAR dari builder
COPY --from=builder /app/target/librobackend-1.0-SNAPSHOT.jar /app/librobackend.jar

# Set command untuk menjalankan aplikasi
CMD ["java", "-jar", "/app/librobackend.jar"]
