# Security Gateway System - Technical Overview

## System Architecture

**Company Context**: Retirement Advantage (formerly MGM)  
**Component**: Security Gateway - A custom authentication and authorization system

## Core Design Principles

The Gateway system was built around a **nested group membership model** with the following key components:

### User Management Structure
- **User Accounts**: Central user registry
- **Groups**: Collections of permissions bundled together
- **Group Membership**: Users assigned to multiple groups for hierarchical access control

### Permission System Design

**Permission Objects** contained:
- **Offset**: Index position of a specific bit in a byte array
- **Friendly Name**: Human-readable description of what the permission grants
- **Bit Mapping**: Each permission corresponds to a single bit flag

**Permission Bundling Process**:
1. Permissions grouped together into logical collections
2. Groups assigned to users based on role requirements
3. User login triggers group membership evaluation
4. System performs distinct membership resolution
5. Generates list of applicable bit indexes
6. Creates byte array with relevant permission bits set

## Token Generation & Security

### Custom JWT Implementation
- **Accidental Innovation**: Independently developed a JWT-like system before widespread JWT adoption
- **Structure**: JSON-based token containing:
  - Plain text user attributes (name, profile data)
  - Permissions byte array
  - Expiration date

### Security Layers
1. **HMAC Digest**: Created using special secret key for tamper detection
2. **Encryption**: Entire token encrypted with separate key
3. **Dual-Key System**: 
   - One key for HMAC signing
   - Another key for encryption

## Runtime Operation

### Token Flow
1. Encrypted token sent to frontend after authentication
2. Frontend includes token in subsequent requests
3. Backend middleware decrypts token automatically
4. HMAC verification ensures integrity
5. Friendly security context object populated for application use

### Developer Experience
**Custom Library Features**:
- Simplified permission checking API
- Request context integration
- Boolean permission evaluation (`true`/`false` responses)
- Abstracted complexity from business service developers

### Performance Optimization
- **Rapid Permission Evaluation**: Implemented fast short-circuiting for permission denial
- **Caching Strategy**: Reduced database calls by embedding common user profile data in tokens
- **Efficient Bit Operations**: Byte array approach enabled fast permission lookups

## Intellectual Property
- **Patent Status**: The rapid permission evaluation method was patented
- **Jurisdiction**: UK patent filing
- **Tax Benefits**: Patent provided tax advantages for the company

## Technical Innovation Notes
- Independent development of JWT-like architecture before mainstream adoption
- Bit-level permission system for memory and performance efficiency
- Dual-encryption approach for enhanced security
- Developer-friendly abstraction layer
