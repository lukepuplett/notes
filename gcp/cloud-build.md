## Overview

- Executed as series of build steps.
- Google's open-source steps.
- Community OS steps.
- Custom ones.
- Each step is run with its container attached to a local Docker network named cloudbuild so steps can talk to each other and share data.
- Can use Docker Hub images.
- Kick off from gcloud or API or built-in triggers, including Cloud Source Repos, GitHub and Bitbucket.
- View builds using gcloud, API or Console.
- Build config in YAML or JSON.
- Stuff published to Artifact Registry.
- You can test run your build locally using cloud-build-local tool.


## Build configuration file schema

- YAML or JSON, e.g. `cloudbuild.yaml`
- No config file is needed to just build a Docker image.
- VSCode Cloud Code extension help.
- If you can containerise it, Cloud Build can run it.
- Defaults to serial execution; use waitFor for concurrent ones.
- Max 100 steps.
- Samples here: https://cloud.google.com/build/docs/build-config-file-schema#json
- `step.name` = specifies cloud builder, which is a container containing common tools; example is gcr.io/cloud-builders/mvn
- `step.args` = passed to the builder, passed to the running tool; if the builder has an ENTRYPOINT then the args are sent to that, else the first arg becomes the entrypoint.
- Max 100 arguments.
- `step.env` = array of environment variables in form `"KEY=VALUE"`
- `step.dir` = changes the working directory to `/workspace/<dir>` unless an absolute path, in which case writes may not be persisted between step executions.
- `step.timeout` = in seconds like `60.01s`, defaults to forever (total build timeout)
- `step.id` = unique for a step, used with waitFor.
- `step.waitFor` = specifies step(s) which must run before; if no value, waits for all.
- `step.entrypoint` = overrides default of the builder.
- `step.secretEnv` = array of variables that must be specified in the build's secrets.
- `step.volumes` = array of Docker container volumes to mount for persisting files between steps (name must be the same), in addition to `/workspace`.
- `root.timeout` = in seconds as before, defaults to 10 minutes.
- `root.queueTtl` = expires build if queued and not run within seconds.
- `root.logsBucket` = override default bucket for logs, e.g. `gs://mybucket`.
- `root.options.env` = global environment variables for all steps.
- `root.options.secretEnv` = as before but global.
- `root.options.volumes` = global and must not conflict with step volumes.
- `root.options.sourceProvenanceHash` = array of algorithm(s) for source provenance, like `"SHA256"`
- `root.options.machineType` = override default single CPU build VM.
- `root.options.diskSizeGb` = max 1000GB, e.g. `"200"` (doc shows string type)
- `root.options.logStreamingOption` = stream logs rather than capture all at end, `"STREAM_ON"`.
- `root.options.logging` = choose to store in Cloud Logging or Cloud Storage, `"GCS_ONLY"`.
- `root.options.dynamic_substitutions` = enable disable bash parameter expansion; when build is run by trigger then this is true; default false.
- `root.options.substitutionOption` = set along with root.substitutions dictionary to specify the behaviour when there is an error in the substitution checks (wtf?!)
- `root.options.pool` = specify the resource name if running on a private pool.
- `root.substitutions` = substitute specific variables when values are not known until build time; errors on missing variable or missing substitution; use `ALLOW_LOOSE` option to skip check, though this is the default when run by trigger.
- `root.tags` = for organization.
- `root.availableSecrets` = see Using secrets: https://cloud.google.com/build/docs/securing-builds/use-secrets
- `root.secrets` = doc says to use availableSecrets.
- `root.serviceAccount` = as it says.
- `root.images` = names of any Docker images to push to the (Artifact) registry.
- `root.artifacts` = complex object specifying "container artifacts" to store in Cloud Storage.

### Using Dockerfiles

- If executing Docker builds using gcloud or build triggers, you must use a Dockerfile. You can provide another build config file.

### Cloud Build network

- Each step's container is attached to a network called cloudbuild which hosts Application Default Credentials to Google Cloud services can automatically find your credentials.
- If you're running nested containers and want to expose ADC to an underlying container, or using gsutil or gcloud in a docker step, use the --network flag in your docker build step, JSON like:

            {
                "steps": [
                    {
                        "name": "gcr.io/cloud-builders/docker",
                        "args": [ "build", "--network=cloudbuild", "." ]
                    }
                ]
            }


## Cloud builders

Publicly available images

Specify the name of any public Docker image URL and use args to specify command(s). This uses the ubuntu image from Docker Hub.

	steps:
	- name: 'ubuntu'
	  args: ['echo', 'hello world']

Other registries need a full URL.

Supported builder images provided by Cloud Build

Pre-built images are in gcr.io/cloud-builders/… e.g. docker, git, gcloud, maven and more at the following URL: The complete list of supported builders for Cloud Build.

### Community-contributed builders

Pre-built images are not available, so you must download them from GitHub and build the image and push it somewhere; probably worth checking Docker Hub for them.

See: https://github.com/GoogleCloudPlatform/cloud-builders-community

Examples include helm, kaniko, docker-compose.

### Writing your own custom builder

You can make your own by building and pushing a container somewhere.


## Cloud Build triggers

- Code repository triggers on push or PR merge to GitHub etc. including GitHub Enterprise on-prem.
- Use the Cloud Build GitHub app to connect and build code in GitHub.
- See: Building repositories from GitHub.
- Can override substitution variables when manually triggered.
- See: Creating manual triggers
- Trigger from Google Pub/Sub via message.
- Webhook trigger via custom URL, including build configuration inline.
- See: Creating webhook triggers
- Time-scheduled via a manual trigger invoked by Cloud Scheduler.
- See: Creating scheduled triggers
- Use Common Expression Language with the variable, build on fields listed in the Build resource to access fields associated with your build event such as your trigger ID, image list or substitution values. Use the filter string to filter build events in your build config file using any field listed in the Build resource.
- See: Use CEL to filter build events
- Some default substitution variables are provided for use with triggers.
- See: Substituting variable values
- Manipulate strings associated with existing variables by using Bash parameter expansions.
- Store part of your trigger's event payload as a substitution variable by using payload bindings, like the author of a PR.
- Can mark a build as pending approval.


## IAM roles and permissions

Too much detail to document here. See doc:

https://cloud.google.com/build/docs/iam-roles-permissions


## Cloud Build service account

- Builds executed under service account.
- [PROJECT_NUMBER]@cloudbuild.gserviceaccount.com
- Has permission to fetch from Cloud Source Repositories and writing to your project's Cloud Storage buckets.
- Can specify your own account to run as.
- Account has Cloud Build Service Account role.
- Default permission here: https://cloud.google.com/build/docs/cloud-build-service-account#default_permissions_of_service_account
- Can have trigger select different account to run as.
- See doc: https://cloud.google.com/build/docs/cloud-build-service-account#build_triggers


## Private pools overview

By default, builds run in a secure hosted environment with internet access. Private pools are dedicated worker VMs that can be heavily customized, including accessing private networks. They are still fully-managed and elastic scale.

Many projects can use a private pool. The service account can be configured to let workers access resource across projects.

You can create a VPC peering connection between your VPC network and the pool's network.

- See doc: https://cloud.google.com/build/docs/private-pools/private-pools-overview
- See also: https://cloud.google.com/build/docs/private-pools/private-pool-config-file-schema
- And: https://cloud.google.com/build/docs/private-pools/use-in-private-network
- And: https://cloud.google.com/build/docs/hybrid/overview


## Quickstart: Build a Docker image

### Create a Docker repository in Artifact Registry

Run this:

	gcloud artifacts repositories create quickstart-docker-repo --repository-format=docker \
	    --location=us-central1 --description="Docker repository"

Check it exists:

	gcloud artifacts repositories list


### Build an image using Dockerfile

You can build an image without needing a Cloud Build config file.

Get your Cloud project ID:

	gcloud config get-value project

Run this from the directory with your Dockerfile:

	gcloud builds submit --tag us-central1-docker.pkg.dev/project-id/quickstart-docker-repo/quickstart-image:tag1

That's it.
