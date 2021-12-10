## Developing your service

Code requirements

- Be contained.
- Listen on PORT environment variable.
- It must not run background work (I bet it can).
- It can have a web server.
	
If you "bring your own binaries" make sure they are configured for Linux ABI x86_64.


## Concepts

### Resource model

- GCP project > Service > Revision > Instance
- Services are zone redundant.
- A project can run many services in different regions.
- Traffic auto routed to latest revision ASAP.
- Auto scaled by default.

### Concurrency

- Default 80 concurrent requests, max 1,000.
- High CPU instances may not be routed to.
- Set to 1 if requests eat big CPU or code is single-threaded.
- Lower threshold to throttle access to resources beyond, e.g. DB.

### Autoscaling

- Default 0.
- Scaling is impacted by rate of incoming requests, and…
- CPU usage maintained under 60%.
- Max concurrency setting.
- Max instances setting.
- Min instances setting.
- When maxed out, requests are queued for up to 60s, else HTTP 429.
- Can breach max limit for short time when big spike comes in (set safety margin).
- Upon deployment, max limit exists twice, so can exceed during phase-over.
- Idle timer can be up to 15 mins.
- Use min-instance setting but incurs continual cost.

### Execution environments

- Gen-1 by default.
- Gen-2 environment is full Linux compatibility, rather than system call emulation.
- Gen-2 is faster and better all round.
- But Gen-2 needs min 512MiB RAM.
- But Gen-2 is slower to start.
- Specify on deploy.


## Connecting to GCP services

When using GCP libraries, you don't need to provide credentials because Cloud Run instances use a default runtime service account with the Project > Editor role, so it's like God. A bit concerning.

You can assign a service account with minimum permissions, e.g. the Firestore User IAM role, see docs.

https://cloud.google.com/run/docs/securing/service-identity#per-service-identity


## Using network file systems

Use network file system to share and persist data between containers on the Gen-2 execution environment.

If not persisting beyond the life of the instance, use the built-in memory file system, but this eats allocated RAM, see doc:

https://cloud.google.com/run/docs/reference/container-contract#filesystem

You can use Filestore or self-managed NFS, NDB, 9P, CIFS/Samba, and Ceph network file systems with Cloud Run.

Cloud Storage FUSE can make blobs look like a file system.

Cloud Run doesn't support NFS locking, when mounting NFS, use the -o nolock option. Design your application to avoid concurrent writes to the same file. You may use a different system, such as Redis or Memorystore to store a mutex.

You'll need to setup a file server on Compute Engine, see docs:

https://cloud.google.com/run/docs/using-network-file-systems
https://cloud.google.com/solutions/filers-on-compute-engine


## Tips

- Configure always-allocated CPU for background activities.
- Background threads are suspended and resumed on the next invocation, so it can cause weird behaviour.
- Disk storage is in-memory, shared with your process, so clean-up immediately.
- Don't crash the process, else the next invocation is a cold start.
- Report error properly, see doc: https://cloud.google.com/run/docs/error-reporting
- Use minimum instances to reduce cold starts.
- Lazy load code.
- Cache objects in statics/memory but don't count on them.
- Configure max concurrency for >1 request handling.
- Remember to scale memory for concurrent requests.
- Use official, secured Docker base images.
- Regularly rebuild to get patched base images.
- Reduce surface area with a super lean image.
- Use Docker USER statement to run not-as-root; some images may already have a user configured.
- Enable container vulnerability scanning if your registry has it.


## Building containers

Cloud Runs accepts any images so long as they listen on HTTP on the injected PORT environment variable.

You can build using Googles tools or Docker. See docs: 

https://cloud.google.com/run/docs/building/containers

Building using Buildpacks

Google provides CNCF-compatible Buildpacks that build code into images designed for Google Cloud container runners. .NET Core is supported. See docs:

https://cloud.google.com/run/docs/building/containers

Testing locally in VSCode

https://cloud.google.com/code/docs/vscode/developing-a-cloud-run-service


## Deploying container images

See docs for permissions to deploy.

- No size limit.
- You can deploy from
    - Google Container Registry (retired product, use Artifact Registry)
    - Artifact Registry (Cloud Build?)
    - Images stored in same project (uh?)
    - Images from other GCP projects.
    - Seems "unsupported" container registries can be used but you have to `docker push` the image to Artifact Registry first.

**Note** - Container Registry is free except for storage bytes and network egress. It's also retired.

Can deploy by tag or exact sha256 digest.

Deploy from Console, CLI, Cloud Code (?), Terraform or service.yaml via the CLI. See doc: 

https://cloud.google.com/run/docs/deploying#service

Using this service.yaml

	apiVersion: serving.knative.dev/v1
	kind: Service
	metadata:
	  name: SERVICE
	spec:
	  template:
	    spec:
	      containers:
	      - image: IMAGE

You can deploy with `gcloud run services replace service.yaml`. The service gets a unique immutable URL.

**Note** - In my trials, the `replace` verb required the **Cloud Resource Manager API** to be enabled.

### Deploying a new version

This creates a new service, where none existed. Use `--revision-suffix` to override default rev.

	gcloud run deploy SERVICE --image IMAGE_URL

**Caution** - The following instructions replaces your existing service configuration with the one specified in the YAML file. So if you use YAML to make revision changes, you should avoid also using the Cloud Console or gcloud tool to make configuration changes because those can be overwritten when you use YAML.

This is the same problem as Azure; the Console is not locked when code is used to define services, unlike on Windows using policy files locks the GUI in Windows.

Use this command to save the configuration on an existing service.

	gcloud run services describe SERVICE --format export > service.yaml

Edit the spec.template and deploy.

	gcloud run services replace service.yaml

Deploy from other Google Cloud projects

Possible when you set the right permissions. See doc:

https://cloud.google.com/run/docs/deploying#other-projects


## Continuous deployment from Git using Cloud Build

Use Cloud Build trigger to build and deploy on commit.

Important - Enable the Cloud Build and Cloud Source Repositories APIs.

See doc: https://cloud.google.com/run/docs/continuous-deployment-with-cloud-build

I think you can deploy from source code in the repo by choosing to use Buildpacks.


## Deploying from source code

You can deploy from source, i.e. no container image.

	gcloud run deploy with the --source flag.

It uses Buildpacks behind the scenes. Containers are stored in Artifact Registry, so only works in regions with that.

If no Dockerfile is present, it'll detect the language and make an image.

Supports Go, Node, Python, Java and .NET Core.

Needs these roles:

	- Cloud Build Editor
	- Artifact Registry Admin
	- Storage Admin
	- Cloud Run Admin
	- (Service Account User role) - uh?

Run this from source folder.

	gcloud run deploy SERVICE--source .

Important - Upon completion this revision serves 100% of traffic.


## Rollbacks, gradual rollouts, and traffic migration

You can split traffic by percentage across versions.

For YAML instructions for these, see doc:

https://cloud.google.com/run/docs/rollouts-rollbacks-traffic-migration#yaml_1

Rollback

For CLI use:

	gcloud run services update-traffic SERVICE --to-revisions REVISION=100

Gradual rollout

For CLI use:

	gcloud run deploy --image IMAGE --no-traffic
	
Then:

	gcloud run services update-traffic SERVICE --to-revisions REVISION
	
Splitting

Specify revisions and percentages via comma-delimited list:

	gcloud run services update-traffic SERVICE --to-revisions LIST

e.g. REV1=10,REV2=90

All to latest

Use:

	gcloud run services update-traffic SERVICE --to-latest

Using tags

After deploy, you can create a new revision and assign a tag that lets you access it via a preview URL. You can then migrate onto that tag. This way you can test then make live.

To deploy a new tagged revision:

	gcloud run deploy myservice --image IMAGE_URL --no-traffic --tag TAG_NAME

The preview URL will be:

	https://mytag---myservice-abcdef.a.run.app

To migrate traffic to it:

	gcloud run services update-traffic myservices --to-tags TAG_NAME=PERCENT


## Invoking with an HTTPS request

They're public so long as you allow unauthenticated traffic.

HTTP/2 is downgraded to HTTP/1 unless you override this setting.

Use this to call a private service.

	curl -H "Authorization: Bearer $(gcloud auth print-identity-token)" SERVICE_URL

You use service-to-service authentication for microservice type designs.

And you can use Google's Identity Platform to limit access.


## Using gRPC

- Configure HTTP/2 end-to-end.
- Write and compile a proto file.
- Instantiate a gRPC server listening on PORT

See doc:

https://cloud.google.com/run/docs/triggering/grpc


## Using WebSockets

No additional configuration is required, but they are still subject to the request timeout (default 5 minutes and maximum 60).

There's no session affinity on Cloud Run, so sockets can reconnect anywhere.

Also works via Cloud Load Balancing.

The hardest part is synching data between instances, which are stateless and autoscaling. Cloud Run instances do not provision CPU outside of a request so polling or pulling is not much use.

I wonder if it might be best to run all streams to a single instance dedicated to this, but single instances are limited by Google to 1,000 connections.

Google recommends that you use external message queue systems such as Redis Pub/Sub (Memorystore) or Firestore real-time updates that can deliver updates to all instances over connections initiated by the container instance.


## Triggering from Pub/Sub

Pub/Sub will push messages to a Cloud Run endpoint via HTTP.

See docs: https://cloud.google.com/run/docs/triggering/pubsub-push


## Running services on a schedule

Use Cloud Scheduler to trigger. Cloud Scheduler may need a service account created with permission to invoke the service.

See docs: https://cloud.google.com/run/docs/triggering/using-scheduler


## Executing asynchronous tasks

Use Cloud Tasks to enqueue a task to be processed by Cloud Run. Use this for:

- Preserving requests through unexpected indicidents.
- Smoothing traffic spikes by delaying work that is not user-facing.
- Speeding response times (eventual consistency).
- Limiting call rates.

Enable the Cloud Tasks API.

Tasks are pushed via HTTPS.

Looks really useful and simple. See docs:

https://cloud.google.com/run/docs/triggering/using-tasks


## Mapping custom domains

- Firebase hosting in front of Cloud Run.
- Use a load balancer with a serverless network endpoint group(?!)
- Use Cloud Run's own mapping (preview).

A load of regions cannot do this (yet?), including us-west2, 3 and 4. You'd have to use a load balancer.

- Map it, then update DNS.
- A managed certificate is automatically issued and renewed.
- Cannot use own certificate.
- You can map >1 domains.

For Console and CLI method, see doc:

https://cloud.google.com/run/docs/mapping-custom-domains

Note - A domain is verified and owned by a specific GCP user, so other people can't map the domain. See doc above for a workaround.


## Serving static assets with a CDN

- Uses Google External HTTP(S) Load Balancing to route requests.
- Which in turn uses Serverless Network Endpoint Groups (NEGs) which are like load balancer pools and group Google's serverless endpoints where they all share a URL pattern.
- The load balancer and NEG must be in the same region as the serverless app.
- Reserve a static IP address for the load balancer.
- Create the SSL certificate.
- Create the external HTTP(S) load balancer.
- Health checks don’t work with NEGs.
- Add the DNS records for your domain.
- It doesn't say, but it seems like the load balancer is the CDN. Static files are cached at the edges. It also doesn't say how to clear the cache.

See doc:

https://cloud.google.com/cdn/docs/setting-up-cdn-with-serverless


## Serving traffic from multiple regions

Deploy in many regions and route locally using Google's external load balancer product and a secure domain pointing to a global anycast IP address (which does the local routing).

For full setup details, see docs:

https://cloud.google.com/run/docs/multiple-regions

Note - Pub/Sub pushes to same region, which sounds desirable, but can be "fixed", see doc above.


## Managing services

Services are the main resources of Cloud Run with a permanent URL.

- Created by deploying a container for the first time.
- Can copy a service.

See doc:

https://cloud.google.com/run/docs/managing/services


## Managing revisions

- Each deploy makes an immutable revision.
- Configuration changes also make revisions.
- Set traffic %.
- Can tag.
- No need to delete them.
- No billing on old ones.
- 1000 max.
- Override suffix with `--revision-suffix`

Doc:

https://cloud.google.com/run/docs/managing/revisions


## Configure

### Memory

- Instances exceeding memory limits are terminated.
- .exe size, process allocation, files (since are in mem) all count.
- 512MiB default.
- Need 2 vCPU for >4GiB.
- Max 16Gi
- Concurrency requires memory.
- Memory costs money.

### CPU

- Default 1.
- Can set per instance lifetime or per request.
- Configuration change makes new revision.
- Does not state max.
- Preview - always on CPU for background work.
- Unless minimum instances >1 instance is still terminated on 15 mins idle.
- You pay for always on and minimum instance.

### Environment variables

- Injected into container.
- PORT is reserved.
- Can set via YAML or command line. YAML is good because not used by container on dev machine so can store product values without them being accidentally used. Secrets are a problem though; could exclude YAML from repo.
- CLI samples, see https://cloud.google.com/run/docs/configuring/environment-variables
- Else use ENV in Dockerfile but service set ones will override.

### Secrets

- Use GCP's Secret Manager product.
- Mount as a volume so secrets are files.
- Pass via environment variables resolved at instance start.
- See doc: https://cloud.google.com/secret-manager/docs/best-practices
- Must grant Secret Manager Accessor role to Cloud Run service account.
- See doc: https://cloud.google.com/run/docs/configuring/secrets#access-secret
- Can use Console, CLI or YAML.
- Secrets appear to have versions; this might be to link them to container revisions, so rollbacks work (assuming dependencies are also rolled-back).
- Can reference a secret in another project.
- See doc: https://cloud.google.com/run/docs/configuring/secrets#mounting-secrets
	
### Request timeout

- Default 5 minutes.
- Max 60.
- Request terminated and HTTP 504 returned.
- Over 15 minutes and clients should implement retries with all code must resume where left off, despite no session affinity.
- Set via Console, CLI and YAML.
- See doc: https://cloud.google.com/run/docs/configuring/request-timeout#setting

### Concurrency

- Console, CLI, YAML.
- Remember to increase memory.
- See doc: https://cloud.google.com/run/docs/configuring/concurrency

### Max instances

- Limit scaling out.
- Can be exceeded for short period during big spikes.
- Max 1,000 unless personal request made to Google.
- Console, CLI, YAML.
- See doc: https://cloud.google.com/run/docs/configuring/max-instances

### Min instances

- Default 0.
- Limit cold starts.
- Kept running on idle.
- Combine with CPU always allocated.
- But can be restarted at any time by Google.
- Console, CLI, YAML.
- See doc: https://cloud.google.com/run/docs/configuring/min-instances

### Containers

- Can override port, entrypoint command and arguments.
- Use --port to change that.
- Use --command and --args to change those.
- Console, CLI, YAML.
- See docs: https://cloud.google.com/run/docs/configuring/containers

### Execution environment (preview)

- Defaults to "first generation" environment.
- See doc: https://cloud.google.com/run/docs/configuring/execution-environments

### Labels

- Key-value pairs applied to service and revisions.
- Use for cost allocation and billing, teams, staging etc., owners, state, etc.
- Creates a new revision.
- Console, CLI, YAML.
- See doc: https://cloud.google.com/run/docs/configuring/labels

### HTTP/2 requests

- Defaults to HTTP/2 to 1 downgrade.
- Must use HTTP/2 cleartext (h2c) because TLS terminates at Cloud Run.
- Confirm h2c with local test using curl command in doc.
- Console, CLI, YAML.
- See doc: https://cloud.google.com/run/docs/configuring/http2

### Connecting to Cloud SQL

- For SQL Server, must setup private IP within GCP, i.e. VPC.
- Use Cloud SQL Auth proxy locally.
- See SQL Server doc: https://cloud.google.com/sql/docs/sqlserver/connect-run
- See PostgreSQL doc: https://cloud.google.com/sql/docs/postgres/connect-run
- See MySQL doc: https://cloud.google.com/sql/docs/mysql/connect-run

C### onnecting to a VPC network

- Serverless VPC Access connectors are paid expenses.
- Each VPC connector requires its own /28 subnet for its connectors.
- Create the connector using Console, CLI or Terraform.
- Configure your service to use it using Console, CLI, YAML or Terraform.
- Use firewall rules to restrict the connector to the VPC network.
- Can route outbound all requests from Cloud Run service to VPC.
- See doc: https://cloud.google.com/run/docs/configuring/connecting-vpc

### Connecting to a Shared VPC network

- I using a Shared VPC, you can set up Serverless VPC Access connectors in either the service project or the (VPC?) host project.
- Advantages: dedicated bandwidth, charges by connectors are linked to service project enabling easier "chargebacks", principal of least privilege; connectors need access grants to resources in Shared VPC, reduces dependency on host project administrator because teams manage connectors linked to their service project.
- To implement see doc: https://cloud.google.com/run/docs/configuring/connecting-shared-vpc

### Static outbound IP address

- Outbound connections made from dynamic pool.
- Can assign static IP in case of IP whitelisting by external service.
- See doc: https://cloud.google.com/run/docs/configuring/static-outbound-ip

### Service accounts

- Default run as Compute Engine service account.
- Not actually recommended; best use a dedicated least privilege account.
- See doc: https://cloud.google.com/run/docs/securing/service-identity


## Secure

### Access control with IAM

- Default only project owners can manage and even invoke services.
- Only project owners and Cloud Run Admins can edit IAM policies e.g. make service public.
- In addition to the --allow-unauthenticated you can add allUsers to role roles/run.invoker via the CLI.
- Note that to use --allow-unauthenticated you must have run.services.setIamPolicy permission.
- But when creating the service, saying yes will set the permission and make it public.
- See docs for editing permissions from Console or CLI: https://cloud.google.com/run/docs/securing/managing-access

### Understanding service identity

- Default instances run as Compute Engine service account which has Project Editor role.
- Recommended to create own service account with least privilege for each service.
- See doc to do this: https://cloud.google.com/run/docs/securing/service-identity
- In code Cloud Run libraries get auth tokens automatically.
- In code if not using a library, use Google's "metadata server" to fetch tokens.
- See docs on that: https://cloud.google.com/run/docs/securing/service-identity#fetching_identity_and_access_tokens_using_the_metadata_server
	
### Optimizing service accounts with Recommender

- Tool that nags at you.

### Restricting ingress

- You can stop internet traffic from reaching your URL.
- Settings are, All, Internal and Internal and Cloud Load Balancing.
- When calling a service from inside GCP, use the URL as normal.
- Lots to know about VPC requests, see doc: https://cloud.google.com/run/docs/securing/ingress
- Cloud Scheduler and Cloud Tasks cannot call internal services (on a VPC? I assume okay via load balancer)

### Using VPC Service Controls

- It's a GCP feature that guards against data exfiltration.
- See doc: https://cloud.google.com/run/docs/securing/using-vpc-service-controls

### Using Binary Authorization

- Something to do with Anthos.
- See doc: https://cloud.google.com/binary-authorization/docs/run/enabling-binauthz-cloud-run

### Using customer managed encryption keys (CMEK)

- Protect deployed images with Cloud KMS CMEK.
- See docs: https://cloud.google.com/run/docs/securing/using-cmek

### Protect services with Cloud Armour

- DDoS protection.
- Used with a serverless NEG backend pointing to Cloud Run et al.
- See doc: https://cloud.google.com/armor/docs/integrating-cloud-armor#serverless


## Log and monitor

### Monitoring health and performance

- Google Cloud Monitoring provide performance monitoring, metrics, uptime, alerts on thresholds.
- GCP's operation suite pricing applies, which means fully managed Cloud Run is free.
- Automatically setup, nothing to do.

### Logging and viewing logs

- Sent to Cloud Logging.
- Request logs, done automatically.
- Coded logs, obviously not.
- View in Cloud Run page in Console or Cloud Logging Logs Explorer in Console.
- Can also view in Cloud Code in Visual Studio Code.
- Can view via CLI gcloud logging read …
- See docs: https://cloud.google.com/run/docs/logging#container-logs
- And also via Logging API via library or REST.
- Write to stdout or stderr.
- Or in files in /dev/log
- Or via client library.
- Writing JSON is marshalled into jsonPayload and text to textPayload.
- When writing JSON dictionary, some fields are elevated into the main log entry object, e.g. severity, message and others.
- See docs for others: https://cloud.google.com/run/docs/logging#log-resource
- Logs with same trace are parent-child correlated.
- No auto correlation to request logs unless using Cloud Logging library.
- But you can do this with logging.googleapis.com/trace field set to X-Cloud-Trace-Context value.
- Can use logs exclusion feature from Cloud Logging to reduce request logs.
- See also Log Entry reference: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry

### Audit logging

- GCP services write logs for "Who did what, where and when?"
- Cloud Run keeps data access logs and admin activity logs.
- The LogEntry type contains logName, resource, timeStamp and protoPayload which contains a AuditLog type.
- Optional service-specific data can be found in AuditLog.serviceData or metadata on newer "integrations".
- You cannot disable audit logging.
- View logs via Console, CLI and the Logging API.
- Lots more at: https://cloud.google.com/run/docs/audit-logging#gcloud

### Error reporting

- No setup needed.
- All exceptions that include a stack trace in supported languages, and that are sent to stdout, stderr or to other logs.
- "Memory limit exceeded" and "No instances available" service errors.
- View via Metrics tab of the service detail page, included a top errors table.
- Can view in Error Reporting page once they have been aggregated.


