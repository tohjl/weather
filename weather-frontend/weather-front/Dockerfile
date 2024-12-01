# frontend/Dockerfile
FROM node:18

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the Vue.js app
RUN npm run build

# Install a simple web server to serve the Vue.js build
RUN npm install -g serve

# Expose the port Vue app will run on
EXPOSE 5000

# Command to run the app
CMD ["serve", "-s", "dist", "-l", "5000"]
