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

| Requirement                                                       | Recommendation      | Comment                                                                                                                                                        |   |   |
|-------------------------------------------------------------------|---------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|
| Accessing public data anonymously                                 | API key             | Only identifies the app.                                                                                                                                       |   |   |
| Accessing private data on behalf of an end user                   | OAuth 2.0 client    | Lets end users authenticate your app with Google and allows your app to access GCP APIs on behalf of the user.                                                 |   |   |
| Accessing private data on behalf of a service account outside GCP | Service account key | Create a service account and download its private key as JSON, then pass it to the client library via the `GOOGLE_APPLICATION_CREDENTIALS` environment variable. |   |   |


## Authenticating as a service account

**Note** - For on-prem or another cloud provider you can use workload identity federation to grant access to external identities without using a service account key. Some client libraries can use Application Default Credentials for external identities.

- For AWS, see https://cloud.google.com/iam/docs/access-resources-aws#generate-automatic.
- For Azure, see https://cloud.google.com/iam/docs/access-resources-azure#generate-automatic.
- For OpenID Connect, see https://cloud.google.com/iam/docs/access-resources-oidc#generate-automatic.

### Finding credentials automatically

If running on GCP and you've attached a service account to that environment (see below), you app can get credentials for the service account automatically. You can attach service accounts to resources for many GCP services.

GCP libraries use ADC to fnd your service account credentials like so:

- First, the `GOOGLE_APPLICATION_CREDENTIALS` environment variable is checked for a service account or configuration file.
- Second, the service account attached to the running resource is used. It may be the default one or one you made.
- Third, give up and crash.

Demo code:

    public object AuthImplicit(string projectId)
    {
        // If you don't specify credentials when constructing the client, the
        // client library will look for credentials in the environment.

        var credential = GoogleCredential.GetApplicationDefault();
        var storage = StorageClient.Create(credential);
        
        // Make an authenticated API request.
        
        var buckets = storage.ListBuckets(projectId);        
        foreach (var bucket in buckets)
        {
            Console.WriteLine(bucket.Name);
        }
        return null;
    } 

### Passing credentials manually

If running in an environment with no service account attached (I'm assuming this means on a local developer device) you must manually create a service account, and one or more service account keys, then pass a key to your app.

#### Creating a service account

    gcloud iam service-accounts create MY_NAME
    
Grant permissions to the account:

    gcloud projects add-iam-policy-binding PROJECT_ID --member="serviceAccount:MY_NAME@PROJECT_ID.iam.gserviceaccount.com" --role="roles/owner"

**Note** - the `--role` flag affects which resources your service account can access in your project and can be changed later. **Don't** grant _Owner_, _Editor_ or _Viewer_ roles when in production, instead grant a predefined role or custom role that meets your needs.

Generate the key file:

    gcloud iam service-accounts keys create FILE_NAME.json --iam-account=MY_NAME@PROJECT_ID.iam.gserviceaccount.com

Note

## Getting started

