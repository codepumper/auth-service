# Authentication Microservice

## Overview
This microservice is an authentication and authorization solution. It ensures secure user management, token validation, and role-based access control for seamless integration with other microservices, such as the ab API Gateway microservice.

## Features

### User Authentication

User Registration and Login: Enable users to securely register and log in.
Password Security: Implements secure password hashing using industry-standard algorithms.
Session Management: Generates JSON Web Tokens (JWT) for secure session management.

### Token Validation

Validation Services: Verifies JWTs to ensure secure communication between services.
Integration: Supports validation requests from external services, like the API Gateway and Backtesting Engine.

### Role-Based Authorization

Role Assignment: Supports role management (e.g., admin, trader) for users.
Access Control: Enforces granular access restrictions based on user roles.

### APIs
Endpoints:
  * Registration: Allows new users to create accounts.
  * Login: Authenticates users and issues JWTs.
  * Token Validation: Validates existing JWTs for inter-service communication.

Protocol Support: Provides REST and optional gRPC endpoints for seamless service interoperability.

### Possible improvements

* Usage of stronger hashing algo (e.g., bcrypt/Argon2)
* Short-lived tokens, secure storage, RS256/ES256 algorithms, token revocation.
* Optional Permissions and role hierarchy.
* API Security: Rate limiting, input validation.
* Two-factor authentication for enhanced security.
* Blacklist/whitelist for revoked tokens.
* Db Practices: Indexing, schema design, encryption, secure connections.
* Monitoring (Grafana?)
