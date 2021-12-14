## Quickstart

Prerequisites are needed before doing this. See *Configuring secret manager*, below.

### Create and access a secret version

Create a secret via CLI:

	echo -n "my super secret data" | gcloud secrets create my-secret \
        --replication-policy="POLICY" \
        --data-file=-

Where `POLICY` is **automatic** or **user-managed**.

Read a version:

	gcloud secrets versions access 1 --secret="my-secret"

Read latest:

	gcloud secrets versions access latest --secret="my-secret"

**Important** - don't use latest in production; pin the secret version to the deployed code version.

Can also use the Console or code libraries.


## Configuring secret manager

These are prerequisites.

### Enable the Secret Manager API and Cloud SDK

- Enable the Secret Manager API via the Console, best to click Enable the API link in the docs here:

https://cloud.google.com/secret-manager/docs/configuring-secret-manager

- Assign minimum IAM roles to developers, admins and service account.
- The **Secret Manager Secret Accessor** role is good for reading only.

### Accessing the API

- Has REST and gRPC APIs for use in code.
- Obviously, requests need auth; done for you with Cloud SDK and code libraries.
- Use `Authorization` header with `Bearer TOKEN`
- Use `gcloud auth print-access-token` to get one.
- For using with Compute Engine or GKE, instance/node needs **cloud-platform** OAuth scope.
- See doc about OAuth scopes: https://cloud.google.com/secret-manager/docs/accessing-the-api#oauth-scopes

### Creating and accessing secrets

- A secret contains 1+ versions, and metadata like labels and replication info.
- Creating needs **Secret Manager Admin** role.

	    gcloud secrets create secret-id --replication-policy="automatic"

- A version can be enabled, disabled or destroyed.

	    gcloud secrets versions add secret-id --data-file="/path/to/file.txt"
	
	    echo -n "this is my super secret data" | \
            gcloud secrets versions add secret-id --data-file=-
	
	    gcloud secrets create secret-id --data-file="/path/to/file.txt"
	
- When accessing, specify its `version-id` or `"latest"`.
- Requires Secret Manager Secret Accessor.
- Adding a version and immediately accessing is strongly consistent, but other ops are eventual.
- IAM permission changes are eventual.

    	gcloud secrets versions access version-id --secret="secret-id"
	
    	gcloud secrets versions access version-id --secret="secret-id" --format='get(payload.data)' | tr '_-' '/+' | base64 -d
	
- Output in UTF-8 or Base64.

### Managing secrets

- List with `gcloud secrets list`
- Details with `gcloud secrets describe secret-id`
- You can also manage permissions via CLI, see https://cloud.google.com/secret-manager/docs/managing-secrets#managing_access_to_secrets
- Update with `gcloud secrets update secret-id --update-labels=key=value`
- Delete with `gcloud secrets delete secret-id`

### Managing secret versions

- States `enabled`, `disabled` and `destroyed` where the version exists but contents are wiped.
- Details with `gcloud secrets versions describe version-id --secret="secret-id"`
- List versions `gcloud secrets versions list secret-id`
- Disable with `gcloud secrets versions disable version-id --secret="secret-id"`
- Enable `gcloud secrets versions enable version-id --secret="secret-id"`
- Destroys are eventually consistent and permanent; try disabling and monitoring your app before going for it proper. Needs **Secret Admin** role.
- Destroy with `gcloud secret versions destroy version-id --secret="secret-id"`

### Using Secret Manager with other products

- With Cloud Build via environment variables.
- With Cloud Code integrations with some IDEs.
- Functions and Run via environment variables or libraries in code.
- Compute Engine via code libraries.
- And see doc: https://cloud.google.com/secret-manager/docs/using-other-products#google-kubernetes-engine
- And Config Connector, whatever that is.

