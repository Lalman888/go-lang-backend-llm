## Credits
This project is a fork of https://github.com/amitshekhariitbhu/go-backend-clean-architecture

What I (https://github.com/stefanwuthrich) have done is:
- Migrated from Gin to Chi (https://github.com/go-chi/chi). Reason: Chi is more lightweight and based on standard net/http library.
- Updated Go modules to latest versions.

# Go Backend Clean Architecture - Chi/Standard net/http

A Go (Golang) Backend Clean Architecture project with Chi, MongoDB, JWT Authentication Middleware, Test, and Docker.


**You can use this project as a template to build your Backend project in the Go language on top of this project.**


Learn about this project architecture in detail from the blogs mentioned below:

- [Go Backend Clean Architecture](https://outcomeschool.com/blog/go-backend-clean-architecture)
- [Go JWT Authentication Middleware](https://outcomeschool.com/blog/go-jwt-authentication-middleware)
- [Configuration with Viper in Go](https://outcomeschool.com/blog/configuration-with-viper-in-go)
- [Test with Testify and Mockery in Go](https://outcomeschool.com/blog/test-with-testify-and-mockery-in-go)
- [Database Normalization vs Denormalization](https://outcomeschool.com/blog/database-normalization-vs-denormalization)

## Architecture Layers of the project

- Router
- Controller
- Usecase
- Repository
- Domain

![Go Backend Clean Architecture Diagram](https://github.com/altafino/go-backend-clean-architecture-chi/blob/main/assets/go-backend-arch-diagram.png?raw=true)

### About original author


Hi, I am Amit Shekhar, Co-Founder @ [Outcome School](https://outcomeschool.com) • IIT 2010-14 • I have taught and mentored many developers, and their efforts landed them high-paying tech jobs, helped many tech companies in solving their unique problems, and created many open-source libraries being used by top companies. I am passionate about sharing knowledge through open-source, blogs, and videos.

You can connect with me on:

- [Twitter](https://twitter.com/amitiitbhu)
- [YouTube](https://www.youtube.com/@amitshekhar)
- [LinkedIn](https://www.linkedin.com/in/amit-shekhar-iitbhu)
- [GitHub](https://github.com/amitshekhariitbhu)

## System Design Playlist on YouTube

- [What is System Design?](https://www.youtube.com/watch?v=i4YWRY3hsdA)
- [Twitter Timeline Design with Fanout Approach - System Design](https://www.youtube.com/watch?v=_7qHGfwgPz0)
- [HTTP Request vs HTTP Long-Polling vs WebSocket vs Server-Sent Events](https://www.youtube.com/watch?v=8ksWRX4xV-s)
- [YouTube Video Upload Service - System Design](https://www.youtube.com/watch?v=N0vvJTkokZc)
- [What is Consistent Hashing?](https://www.youtube.com/watch?v=dV5cIm9T3ss)
- [Capacity Estimation: Back-of-the-envelope calculation - Twitter](https://www.youtube.com/watch?v=yrbKxzXm6_Q)

## Major Packages used in this project

- **chi**: Chi is an HTTP web framework written in Go (Golang) using standard library net/http. It's lightweight and fast.
- **mongo go driver**: The Official Golang driver for MongoDB.
- **jwt**: JSON Web Tokens are an open, industry-standard RFC 7519 method for representing claims securely between two parties. Used for Access Token and Refresh Token.
- **viper**: For loading configuration from the `.env` file. Go configuration with fangs. Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile, or Java properties formats.
- **bcrypt**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **testify**: A toolkit with common assertions and mocks that plays nicely with the standard library.
- **mockery**: A mock code autogenerator for Golang used in testing.
- Check more packages in `go.mod`.

### Public API Request Flow without JWT Authentication Middleware

![Public API Request Flow](https://github.com/altafino/go-backend-clean-architecture-chi/blob/main/assets/go-arch-public-api-request-flow.png?raw=true)

### Private API Request Flow with JWT Authentication Middleware

> JWT Authentication Middleware for Access Token Validation.

![Private API Request Flow](https://github.com/altafino/go-backend-clean-architecture-chi/blob/main/assets/go-arch-private-api-request-flow.png?raw=true)

### How to run this project?

We can run this Go Backend Clean Architecture project with or without Docker. Here, I am providing both ways to run this project.

- Clone this project

```bash
# Move to your workspace
cd your-workspace

# Clone this project into your workspace
git clone https://github.com/altafino/go-backend-clean-architecture-chi.git

# Move to the project root directory
cd go-backend-clean-architecture
```

#### Run without Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install `go` if not installed on your machine.
- Install `MongoDB` if not installed on your machine.
- Important: Change the `DB_HOST` to `localhost` (`DB_HOST=localhost`) in `.env` configuration file. `DB_HOST=mongodb` is needed only when you run with Docker.
- Run `go run cmd/main.go`.
- Access API using `http://localhost:8080`

#### Run with Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.
- Access API using `http://localhost:8080`

### How to run the test?

```bash
# Run all tests
go test ./...
```

### How to generate the mock code?

In this project, to test, we need to generate mock code for the use-case, repository, and database.

```bash
# Generate mock code for the usecase and repository
mockery --dir=domain --output=domain/mocks --outpkg=mocks --all

# Generate mock code for the database
mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all
```

Whenever you make changes in the interfaces of these use-cases, repositories, or databases, you need to run the corresponding command to regenerate the mock code for testing.

### The Complete Project Folder Structure

```
.
├── Dockerfile
├── api
│   ├── controller
│   │   ├── error.go
│   │   ├── login_controller.go
│   │   ├── profile_controller.go
│   │   ├── profile_controller_test.go
│   │   ├── refresh_token_controller.go
│   │   ├── signup_controller.go
│   │   └── task_controller.go
│   ├── middleware
│   │   └── jwt_auth_middleware.go
│   └── route
│       ├── login_route.go
│       ├── profile_route.go
│       ├── refresh_token_route.go
│       ├── route.go
│       ├── signup_route.go
│       └── task_route.go
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── docker-compose.yaml
├── domain
│   ├── error_response.go
│   ├── jwt_custom.go
│   ├── login.go
│   ├── profile.go
│   ├── refresh_token.go
│   ├── signup.go
│   ├── success_response.go
│   ├── task.go
│   └── user.go
├── go.mod
├── go.sum
├── internal
│   └── tokenutil
│       └── tokenutil.go
├── mongo
│   └── mongo.go
├── repository
│   ├── task_repository.go
│   ├── user_repository.go
│   └── user_repository_test.go
└── usecase
    ├── login_usecase.go
    ├── profile_usecase.go
    ├── refresh_token_usecase.go
    ├── signup_usecase.go
    ├── task_usecase.go
    └── task_usecase_test.go
```

### API Endpoints

- Login
- Profile
- Refresh Token
- Signup
- Task Create


### Example API Request and Response

- signup

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/public/signup' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "email": "test@gmail.com",
    "password": "test",
    "name": "Test Name"
    }'
  ```

  - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- login

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/public/login' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "email": "test@gmail.com",
  "password": "test"
  }'

  ```

  - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- profile

  - Request

  ```
  curl --location --request GET 'http://localhost:8080/protected/profile' \
  --header 'Authorization: Bearer access_token'
  ```

  - Response

  ```json
  {
    "name": "Test Name",
    "email": "test@gmail.com"
  }
  ```

- task create

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/protected/task' \
  --header 'Authorization: Bearer access_token' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "title": "Test Task"
  }'

  ```

  - Response

  ```json
  {
    "message": "Task created successfully"
  }
  ```

- task fetch

  - Request

  ```
  curl --location --request GET 'http://localhost:8080/protected/task' \
  --header 'Authorization: Bearer access_token'
  ```

  - Response

  ```json
  [
    {
      "title": "Test Task"
    },
    {
      "title": "Test Another Task"
    }
  ]
  ```

- refresh token

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/public/refresh' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "refreshToken": "refresh_token"
  }'
  ```

  - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```


### TODO

- Improvement based on feedback.
- Add more test cases.
- Always try to update with the latest version of the packages used.

## If this project helps you in anyway, show your love ❤️ by putting a ⭐ on this project ✌️

### License

```
   Copyright (C) 2024 Amit Shekhar

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

### Contributing to Go Backend Clean Architecture

All pull requests are welcome.
