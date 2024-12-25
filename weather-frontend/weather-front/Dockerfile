# Step 1: Build stage
FROM node:18 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock)
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the rest of the project files
COPY . .

# Build the Vite project for production
RUN npm run build

# Step 2: Production stage
FROM nginx:alpine

# Copy the build output from the build stage to the nginx container
COPY --from=build /app/dist /usr/share/nginx/html

# Expose the port nginx is listening on
EXPOSE 80

# Start nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]
