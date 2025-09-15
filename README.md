# 🚗 Carcius Rent Car

A modern, full-stack car rental platform built with microservices architecture, featuring a responsive Next.js frontend and Go backend services.

## 🌟 Features

- **User Authentication**: Secure registration and login with JWT
- **Car Listings**: Browse available cars with filters and search
- **Booking System**: Reserve cars for specific date ranges
- **User Profiles**: Manage personal information and view booking history
- **Responsive Design**: Works on desktop, tablet, and mobile devices
- **Microservices**: Scalable architecture with independent services

## 🏗️ Architecture

![Architecture Diagram](https://via.placeholder.com/1200x800/1a202c/ffffff?text=Carcius+Rent+Car+Architecture)

The application follows a microservices architecture with the following components:

- **Frontend**: Modern Next.js application with TypeScript and Tailwind CSS
- **API Gateway**: Single entry point for all frontend requests (Go + Gin)
- **Users Service**: Manages user accounts and authentication (Go + Gin + JWT)
- **Cars Service**: Handles car inventory and details (Go + Gin + GORM)
- **Bookings Service**: Manages reservations and availability (Go + Gin + GORM)
- **PostgreSQL**: Primary database for data persistence
- **Docker**: Containerization for easy deployment and development

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Node.js 18+ (for local frontend development)
- Go 1.21+ (for local backend development)
- Git

### Quick Start with Docker

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/carcius-rent-car.git
   cd carcius-rent-car
   ```

2. Create a `.env` file in the root directory with the following content:
   ```env
   # Database
   DB_HOST=postgres
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=carrental
   
   # JWT
   JWT_SECRET=your_secure_jwt_secret_key_here
   
   # Service URLs
   USERS_SERVICE_URL=http://users-service:8081
   CARS_SERVICE_URL=http://cars-service:8082
   BOOKINGS_SERVICE_URL=http://bookings-service:8083
   
   # Ports
   API_GATEWAY_PORT=8080
   USERS_SERVICE_PORT=8081
   CARS_SERVICE_PORT=8082
   BOOKINGS_SERVICE_PORT=8083
   FRONTEND_PORT=3000
   ```

3. Start all services using Docker Compose:
   ```bash
   docker-compose up --build
   ```

4. Access the application:
   - Frontend: http://localhost:3000
   - API Gateway: http://localhost:8080
   - PostgreSQL: localhost:5432 (user: postgres, password: postgres)

## 🛠️ Development Setup

### Running Services Individually

1. Start the PostgreSQL database:
   ```bash
   docker-compose up -d postgres
   ```

2. Start each service in a separate terminal:
   ```bash
   # API Gateway
   cd services/api-gateway
   go run main.go
   
   # Users Service
   cd services/users-service
   go run main.go
   
   # Cars Service
   cd services/cars-service
   go run main.go
   
   # Bookings Service
   cd services/bookings-service
   go run main.go
   
   # Frontend
   cd frontend
   npm install
   npm run dev
   ```

### API Endpoints

#### Users Service (`/api/users`)
- `POST /register` - Register a new user
- `POST /login` - Authenticate user and get JWT token
- `GET /me` - Get current user profile (requires authentication)
- `PUT /me` - Update user profile (requires authentication)

#### Cars Service (`/api/cars`)
- `GET /` - List all available cars (with filters)
- `GET /:id` - Get car details by ID
- `POST /` - Add a new car (admin only)
- `PUT /:id` - Update car details (admin only)
- `DELETE /:id` - Remove a car (admin only)

#### Bookings Service (`/api/bookings`)
- `GET /` - Get user's bookings (requires authentication)
- `GET /:id` - Get booking details (requires authentication)
- `POST /` - Create a new booking (requires authentication)
- `PUT /:id/cancel` - Cancel a booking (requires authentication)
- `GET /availability` - Check car availability for dates

## 📁 Project Structure

```
carcius-rent-car/
├── frontend/                 # Next.js frontend application
│   ├── public/               # Static files
│   ├── src/                  # Source code
│   │   ├── app/              # Next.js app router pages
│   │   ├── components/       # Reusable UI components
│   │   ├── contexts/         # React contexts
│   │   ├── lib/              # Utility functions and API client
│   │   └── types/            # TypeScript type definitions
│   ├── .env.local            # Frontend environment variables
│   ├── next.config.js        # Next.js configuration
│   └── package.json          # Frontend dependencies
│
├── services/
│   ├── api-gateway/          # API Gateway service
│   ├── users-service/        # Users microservice
│   ├── cars-service/         # Cars microservice
│   └── bookings-service/     # Bookings microservice
│       ├── handlers/         # Request handlers
│       ├── models/           # Database models
│       ├── database/         # Database connection and migrations
│       ├── .env              # Service environment variables
│       └── main.go           # Service entry point
│
├── docker-compose.yml        # Docker Compose configuration
└── README.md                 # This file
```

## 🔧 Configuration

### Environment Variables

#### Common (for all services)
```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=carrental
```

#### API Gateway
```env
PORT=8080
USERS_SERVICE_URL=http://users-service:8081
CARS_SERVICE_URL=http://cars-service:8082
BOOKINGS_SERVICE_URL=http://bookings-service:8083
```

#### Users Service
```env
PORT=8081
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRES_IN=24h
```

#### Frontend
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_GOOGLE_MAPS_API_KEY=your_google_maps_api_key
```

## 🧪 Testing

### Running Tests

1. Navigate to the service directory:
   ```bash
   cd services/users-service  # or any other service
   ```

2. Run the tests:
   ```bash
   go test -v ./...
   ```

### Frontend Tests

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Run the tests:
   ```bash
   npm test
   ```

## 🚀 Deployment

### Production Build

1. Build the frontend for production:
   ```bash
   cd frontend
   npm run build
   ```

2. Start the production stack:
   ```bash
   docker-compose -f docker-compose.prod.yml up --build -d
   ```

### Deployment Options

- **Docker Swarm**: For container orchestration
- **Kubernetes**: For production-grade deployments
- **Cloud Providers**: Deploy to AWS, GCP, or Azure using their container services

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Next.js](https://nextjs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
