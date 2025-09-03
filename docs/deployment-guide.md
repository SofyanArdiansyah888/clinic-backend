# Clinic Backend - Deployment Guide

## Overview
This guide provides comprehensive instructions for deploying the Clinic Backend application to various environments. The application is built with Go Fiber and uses PostgreSQL as the database.

## Prerequisites

### Production Requirements
- **Server**: Linux-based server (Ubuntu 20.04+ recommended)
- **Go**: Version 1.21 or higher
- **PostgreSQL**: Version 13 or higher
- **Nginx**: For reverse proxy and SSL termination
- **SSL Certificate**: For HTTPS (Let's Encrypt recommended)
- **Firewall**: Configured to allow necessary ports

### Optional Components
- **Redis**: For caching and session storage
- **Docker**: For containerized deployment
- **PM2/Systemd**: For process management
- **Monitoring**: Prometheus, Grafana, or similar

## Environment Configuration

### Environment Variables
Create environment-specific configuration files:

#### Production Environment (.env.production)
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=clinic_user
DB_PASSWORD=secure_password_here
DB_NAME=clinic_production
DB_SSL_MODE=require

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
ENV=production

# JWT Configuration
JWT_SECRET=very_secure_jwt_secret_key_here
JWT_EXPIRY=24h

# Security
CORS_ORIGIN=https://yourdomain.com
RATE_LIMIT=100

# Logging
LOG_LEVEL=info
LOG_FILE=/var/log/clinic-backend/app.log

# Redis (if using)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

#### Staging Environment (.env.staging)
```env
# Similar to production but with staging-specific values
DB_NAME=clinic_staging
ENV=staging
LOG_LEVEL=debug
```

## Deployment Methods

### Method 1: Direct Deployment

#### 1. Server Setup
```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install required packages
sudo apt install -y postgresql postgresql-contrib nginx git curl

# Install Go
wget https://golang.org/dl/go1.21.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### 2. Database Setup
```bash
# Create database user
sudo -u postgres psql
CREATE USER clinic_user WITH PASSWORD 'secure_password_here';
CREATE DATABASE clinic_production OWNER clinic_user;
GRANT ALL PRIVILEGES ON DATABASE clinic_production TO clinic_user;
\q

# Configure PostgreSQL for production
sudo nano /etc/postgresql/13/main/postgresql.conf
# Set appropriate values for:
# - max_connections
# - shared_buffers
# - effective_cache_size
# - maintenance_work_mem
# - checkpoint_completion_target
# - wal_buffers
# - default_statistics_target
# - random_page_cost
# - effective_io_concurrency
# - work_mem
# - min_wal_size
# - max_wal_size

# Restart PostgreSQL
sudo systemctl restart postgresql
```

#### 3. Application Deployment
```bash
# Create application directory
sudo mkdir -p /opt/clinic-backend
sudo chown $USER:$USER /opt/clinic-backend

# Clone repository
cd /opt/clinic-backend
git clone <repository-url> .

# Build application
go mod download
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o clinic-backend .

# Create necessary directories
sudo mkdir -p /var/log/clinic-backend
sudo chown $USER:$USER /var/log/clinic-backend

# Copy environment file
cp .env.production .env
```

#### 4. Systemd Service Configuration
Create systemd service file:

```bash
sudo nano /etc/systemd/system/clinic-backend.service
```

```ini
[Unit]
Description=Clinic Backend Service
After=network.target postgresql.service
Wants=postgresql.service

[Service]
Type=simple
User=clinic
Group=clinic
WorkingDirectory=/opt/clinic-backend
ExecStart=/opt/clinic-backend/clinic-backend
Restart=always
RestartSec=5
Environment=ENV=production
StandardOutput=journal
StandardError=journal
SyslogIdentifier=clinic-backend

# Security settings
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/var/log/clinic-backend

[Install]
WantedBy=multi-user.target
```

Enable and start the service:
```bash
sudo systemctl daemon-reload
sudo systemctl enable clinic-backend
sudo systemctl start clinic-backend
sudo systemctl status clinic-backend
```

#### 5. Nginx Configuration
Create Nginx configuration:

```bash
sudo nano /etc/nginx/sites-available/clinic-backend
```

```nginx
server {
    listen 80;
    server_name yourdomain.com www.yourdomain.com;
    
    # Redirect HTTP to HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name yourdomain.com www.yourdomain.com;
    
    # SSL Configuration
    ssl_certificate /etc/letsencrypt/live/yourdomain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/yourdomain.com/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES256-GCM-SHA512:DHE-RSA-AES256-GCM-SHA512:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-GCM-SHA384;
    ssl_prefer_server_ciphers off;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;
    
    # Security headers
    add_header X-Frame-Options DENY;
    add_header X-Content-Type-Options nosniff;
    add_header X-XSS-Protection "1; mode=block";
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
    
    # Rate limiting
    limit_req_zone $binary_remote_addr zone=api:10m rate=10r/s;
    limit_req zone=api burst=20 nodelay;
    
    # Proxy configuration
    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
        proxy_read_timeout 300s;
        proxy_connect_timeout 75s;
    }
    
    # Health check endpoint
    location /health {
        proxy_pass http://localhost:8080/health;
        access_log off;
    }
    
    # Static files (if any)
    location /static/ {
        alias /opt/clinic-backend/static/;
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
}
```

Enable the site:
```bash
sudo ln -s /etc/nginx/sites-available/clinic-backend /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

#### 6. SSL Certificate (Let's Encrypt)
```bash
# Install Certbot
sudo apt install certbot python3-certbot-nginx

# Obtain SSL certificate
sudo certbot --nginx -d yourdomain.com -d www.yourdomain.com

# Set up auto-renewal
sudo crontab -e
# Add: 0 12 * * * /usr/bin/certbot renew --quiet
```

### Method 2: Docker Deployment

#### 1. Dockerfile
Create a Dockerfile in the project root:

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o clinic-backend .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/clinic-backend .

# Create non-root user
RUN addgroup -g 1001 -S clinic && \
    adduser -u 1001 -S clinic -G clinic

# Change ownership
RUN chown clinic:clinic /root/clinic-backend

USER clinic

EXPOSE 8080

CMD ["./clinic-backend"]
```

#### 2. Docker Compose
Create docker-compose.yml:

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=clinic_user
      - DB_PASSWORD=secure_password_here
      - DB_NAME=clinic_production
      - ENV=production
    depends_on:
      - postgres
      - redis
    restart: unless-stopped
    networks:
      - clinic-network

  postgres:
    image: postgres:13-alpine
    environment:
      - POSTGRES_USER=clinic_user
      - POSTGRES_PASSWORD=secure_password_here
      - POSTGRES_DB=clinic_production
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    restart: unless-stopped
    networks:
      - clinic-network

  redis:
    image: redis:7-alpine
    command: redis-server --appendonly yes
    volumes:
      - redis_data:/data
    restart: unless-stopped
    networks:
      - clinic-network

  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - app
    restart: unless-stopped
    networks:
      - clinic-network

volumes:
  postgres_data:
  redis_data:

networks:
  clinic-network:
    driver: bridge
```

#### 3. Deploy with Docker
```bash
# Build and start services
docker-compose up -d

# Check logs
docker-compose logs -f app

# Scale if needed
docker-compose up -d --scale app=3
```

## Monitoring and Logging

### Application Monitoring
```bash
# Check application status
sudo systemctl status clinic-backend

# View logs
sudo journalctl -u clinic-backend -f

# Monitor resource usage
htop
iotop
```

### Database Monitoring
```bash
# Check PostgreSQL status
sudo systemctl status postgresql

# Monitor database connections
sudo -u postgres psql -c "SELECT count(*) FROM pg_stat_activity;"

# Check database size
sudo -u postgres psql -c "SELECT pg_size_pretty(pg_database_size('clinic_production'));"
```

### Nginx Monitoring
```bash
# Check Nginx status
sudo systemctl status nginx

# View access logs
sudo tail -f /var/log/nginx/access.log

# View error logs
sudo tail -f /var/log/nginx/error.log
```

## Backup and Recovery

### Database Backup
```bash
# Create backup script
sudo nano /opt/backup-db.sh
```

```bash
#!/bin/bash
BACKUP_DIR="/opt/backups"
DATE=$(date +%Y%m%d_%H%M%S)
DB_NAME="clinic_production"

# Create backup directory
mkdir -p $BACKUP_DIR

# Create backup
sudo -u postgres pg_dump $DB_NAME > $BACKUP_DIR/backup_$DATE.sql

# Compress backup
gzip $BACKUP_DIR/backup_$DATE.sql

# Keep only last 7 days of backups
find $BACKUP_DIR -name "backup_*.sql.gz" -mtime +7 -delete

echo "Backup completed: backup_$DATE.sql.gz"
```

Make executable and schedule:
```bash
chmod +x /opt/backup-db.sh
sudo crontab -e
# Add: 0 2 * * * /opt/backup-db.sh
```

### Application Backup
```bash
# Backup application files
sudo tar -czf /opt/backups/app_$(date +%Y%m%d_%H%M%S).tar.gz /opt/clinic-backend

# Backup configuration files
sudo tar -czf /opt/backups/config_$(date +%Y%m%d_%H%M%S).tar.gz /etc/nginx/sites-available/clinic-backend /etc/systemd/system/clinic-backend.service
```

## Security Hardening

### Firewall Configuration
```bash
# Install UFW
sudo apt install ufw

# Configure firewall
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow ssh
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable
```

### Application Security
```bash
# Create dedicated user
sudo adduser clinic
sudo usermod -aG clinic www-data

# Set proper permissions
sudo chown -R clinic:clinic /opt/clinic-backend
sudo chmod -R 755 /opt/clinic-backend
```

### Database Security
```bash
# Configure PostgreSQL for security
sudo nano /etc/postgresql/13/main/pg_hba.conf
# Restrict connections to localhost only
# local   all             all                                     peer
# host    all             all             127.0.0.1/32            md5
# host    all             all             ::1/128                 md5
```

## Performance Optimization

### Application Optimization
```bash
# Enable GOMAXPROCS
export GOMAXPROCS=$(nproc)

# Optimize build flags
go build -ldflags="-s -w -X main.Version=1.0.0" -o clinic-backend .
```

### Database Optimization
```bash
# Analyze and vacuum database
sudo -u postgres psql clinic_production -c "VACUUM ANALYZE;"

# Create indexes for frequently queried columns
sudo -u postgres psql clinic_production -c "CREATE INDEX CONCURRENTLY idx_patients_name ON patients(name);"
```

### Nginx Optimization
```nginx
# Add to nginx.conf
worker_processes auto;
worker_connections 1024;
keepalive_timeout 65;
gzip on;
gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;
```

## Troubleshooting

### Common Issues

#### Application Won't Start
```bash
# Check logs
sudo journalctl -u clinic-backend -n 50

# Check permissions
ls -la /opt/clinic-backend/

# Check environment file
cat /opt/clinic-backend/.env
```

#### Database Connection Issues
```bash
# Test database connection
psql -h localhost -U clinic_user -d clinic_production

# Check PostgreSQL logs
sudo tail -f /var/log/postgresql/postgresql-13-main.log
```

#### Nginx Issues
```bash
# Test configuration
sudo nginx -t

# Check error logs
sudo tail -f /var/log/nginx/error.log

# Check access logs
sudo tail -f /var/log/nginx/access.log
```

### Performance Issues
```bash
# Check system resources
top
free -h
df -h

# Check application performance
curl -w "@curl-format.txt" -o /dev/null -s "http://localhost:8080/health"
```

## Rollback Procedures

### Application Rollback
```bash
# Stop application
sudo systemctl stop clinic-backend

# Restore previous version
cd /opt/clinic-backend
git checkout <previous-commit-hash>
go build -o clinic-backend .

# Start application
sudo systemctl start clinic-backend
```

### Database Rollback
```bash
# Restore from backup
sudo -u postgres psql clinic_production < /opt/backups/backup_20240101_120000.sql
```

## Maintenance

### Regular Maintenance Tasks
```bash
# Update system packages
sudo apt update && sudo apt upgrade -y

# Update application
cd /opt/clinic-backend
git pull origin main
go build -o clinic-backend .
sudo systemctl restart clinic-backend

# Clean up old logs
sudo journalctl --vacuum-time=7d

# Clean up old backups
find /opt/backups -name "*.tar.gz" -mtime +30 -delete
```

This deployment guide should be updated as the application evolves and new deployment requirements emerge.
