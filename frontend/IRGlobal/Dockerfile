# Build stage
FROM node:18-alpine AS builder

WORKDIR /app

# Declare build arguments que serão passados de fora
ARG VITE_API_BASE_URL
ARG VITE_API_TIMEOUT # Se você também quiser passar o timeout

# Copy package files
COPY package.json package-lock.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

RUN VITE_API_BASE_URL=$VITE_API_BASE_URL VITE_API_TIMEOUT=$VITE_API_TIMEOUT npm run build

# Production stage
FROM nginx:alpine

# Copy built assets from builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]