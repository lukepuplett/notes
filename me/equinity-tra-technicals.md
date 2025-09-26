# System Architecture Overview

**Core Structure**: The original application was re-architected into an event-based microservices system with approximately 18 services running on Microsoft Azure Service Fabric (Microsoft's pre-Kubernetes orchestration platform that originally powered Azure when it was called "Red Dog").

**Technology Stack**: 
- Services written in C# ASP.NET Core
- Angular frontend
- MongoDB for data storage
- RabbitMQ message bus for inter-service communication
- SignalR for real-time websocket connections

## Service Architecture Pattern

**Two-Tier Service Design**: Each workflow step used a "stack" consisting of:
1. **Backend-for-Frontend (BFF) Service** - Flexible layer allowing frontend developers to design their own backends, handling UI-specific calculations and returning view models
2. **Core Market Service** - More stable service containing business logic and handling data storage operations

**Communication Flow**:
- Frontend ↔ BFF: HTTP REST API
- BFF ↔ Core Service: HTTP REST API  
- Inter-service: Event-driven via RabbitMQ message queue

## Event-Driven Architecture

**Event Sourcing**: All events were stored (in RabbitMQ or MongoDB) to enable event replay. When new services were added, the entire event stream could be replayed to bring them up to current state, as if they had always been running.

**Real-time Updates**: SignalR pushed events to the frontend for real-time UI updates, enabling features like disabling buttons when another user was working on the same record.

## Hypermedia API Design

**Self-Describing APIs**: Services implemented true hypermedia REST APIs that returned JSON containing:
- Available actions/links for navigation
- Form definitions with required fields and valid values
- Complete self-documentation without separate API specs

**Custom Tooling**: A command-line hypermedia client called "Hyperdrive" was developed that could:
- Parse hypermedia JSON controls (like a browser parses HTML)
- Drive APIs programmatically using available controls
- Execute BDD (Behavior-Driven Development) tests written in Gherkin syntax

## Organizational Challenge

**Inverted Hierarchy**: The company had an unusual structure where non-technical systems analysts managed technical developers. The hypermedia approach was partly adopted to enable these analysts to interact with APIs using tools like Postman, despite their limited technical capabilities. This organizational dysfunction ultimately proved unsustainable and contributed to the departure from the project.
