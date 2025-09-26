# Zipwire Architecture Overview

## Application Type
Zipwire is a monolithic C# application built using ASP.NET Pages as a traditional multi-page application (MPA). Despite being monolithic, it implements a microservices-within-a-monolith pattern.

## Vertical Organization

The application is organized into **verticals** based on feature areas. Each vertical represents a distinct business domain and operates with clear boundaries, even within the shared monolithic codebase.

### Structure Per Vertical
- **User Interface**: Cross-cutting layer that spans across all verticals (necessary due to monolithic nature)
- **Controllers**: Web page controllers sit within their respective vertical namespaces and C# projects
- **CQRS Pattern**: Each vertical implements Command Query Responsibility Segregation with two primary facades:

#### Commands
- Handle mutating operations
- Contain macro-level functionality
- Bundle large pieces of work that need to be completed together

#### Queries
- Read-only operations
- Implemented as large facade interfaces with corresponding implementation classes
- Handle commonly-performed read operations
- Could benefit from being broken into smaller classes but remain manageable

## Data Storage Architecture

The system uses a dual storage approach:

### Firestore (NoSQL Document Store)
- Google Cloud's automatically scaling database service
- Similar to MongoDB in functionality
- Used for data that requires querying capabilities:
  - Range queries
  - Data that cannot be directly addressed by key
- Stores minimal records to optimize cost and performance
- Example: User account records for querying purposes

### Blob Storage
- Used for richer, non-queryable data
- Stores data in JSON format within blobs
- More cost-effective for large data that doesn't require querying
- Example: Detailed account information that supplements Firestore records

## Repository Pattern

Two types of repositories handle the dual storage system:
- **Blob Repositories**: Handle reading and writing to blob storage
- **Firestore Repositories**: Handle reading and writing to the NoSQL database

## Event-Driven Communication

### In-Memory Event Bus
Since the application runs as a monolith, it uses a custom in-memory event bus rather than implementing a full service bus infrastructure.

### Event Flow
1. **Event Emission**: Verticals (primarily through Commands) emit events when macro-level actions occur
2. **Task Queuing**: Events are converted to C# Tasks (similar to JavaScript Promises) and queued
3. **Event Processing**: Subscriber services observe and process these events asynchronously
4. **Loose Coupling**: Verticals communicate indirectly through events without direct dependencies

### Example: Notifications
The accounts vertical owns email addresses and account data, making it responsible for notifications. Other verticals emit events when actions require notifications, and the accounts vertical processes these events to send emails.

## Data Isolation Principles

### Database Isolation
- All verticals share the same Firestore instance but maintain strict table boundaries
- Verticals are prohibited from accessing other verticals' tables
- This rule is enforced by convention rather than code constraints

### Blob Storage Isolation
- Each vertical has its own blob storage buckets (sometimes multiple buckets per vertical)
- Cross-vertical blob access is prohibited
- Maintains microservices-style isolation within the monolith

## Future Scalability

The architecture is designed for potential future decomposition:

### Splitting Options
- The monolith could be divided into separate services (estimated 8 verticals total)
- Possible configurations: 4-4 split, or extracting individual verticals like accounts
- Would require implementing a real service bus to replace the in-memory event system

### Benefits
- Avoids typical monolithic problems through vertical isolation
- Maintains microservices benefits while keeping operational simplicity
- Enables gradual migration to true microservices architecture when needed
