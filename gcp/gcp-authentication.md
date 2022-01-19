# Authentication

## Overview

- **Authentication** is who you are, see https://cloud.google.com/iam/docs.
- **Authorization** is what you can do, see https://cloud.google.com/logging/docs/audit.
- **Auditing** is what you did.

This doc is for authentication.

### Principals

An entity, aka an identity, that can be granted access to a resource. There are two types:

- **User accounts** are Google Accounts and represent humans; developers, administrators and anyone else interacting with GCP.
- **Service accounts** are managed by IAM and represent non-human users; applications and automated jobs that need to access resources or perform actions like running App Engine.

More on IAM here, https://cloud.google.com/iam/docs/overview.

### Applications

GCP APIs only accept requests from _registered applications_ which are identifiable via presented application **credentials**.

Valid credential types include:

- API keys
- OAuth 2.0 client credentials
- Service account keys

Service accounts can be used as both an application credential or a principal identity.

Understanding service accounts, here: https://cloud.google.com/iam/docs/understanding-service-accounts

### Authentication strategies

GCP APIs use OAuth 2.0 for user and service accounts. Most GCP APIs also support anonymous access to public data using API keys, but they only identify the application and not the principal, so when using API keys, the principal must be authenticated by other means.

GCP APIs support multiple authentication flows for different runtime environments. GCP SDKs are recommended.

- Use Google's client libraries.
- Determine the correct authentication flow for your app.
- Find or create the app credentials you need.
- Pass the credentials to the client libraries, ideally via Application Default Credentials.

See ADC, here: https://cloud.google.com/docs/authentication/production#automatically

