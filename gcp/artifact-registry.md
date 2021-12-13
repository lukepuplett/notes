## Overview

Expands on and replaces the Google Container Registry product.

- Store packages and Docker containers.
- Store artifacts from Cloud Build.
- Deploy artifacts to Google services.
- IAM rights.
- Scan for vulnerabilities.
- Enforce deployment policies with Binary Authorization.
- Place in VPC perimeter.


## Dependency management

### Approaches to including dependencies

- Install from public services, like npm or NuGet.
- Vendoring; downloading and storing in repo.
- Store in private registry like Artifact Registry.

### Version pinning

- Restricting dependency to version or limited range.
- Helps make builds reproducible.
- But security fixes not auto adopted.
- Auto scan and fix can help.

### Signature and hash verification

- Authenticity of dependency.
	
### Lock files and compiled dependencies

- Lock files specify exactly what version of every dependency, down the graph.
- Include signature or hash verification.
- Consistent builds.

### Mixing private and public dependencies

- More vulnerable to dependency confusion attack; name spoofing.
- Sig or hash verification.
- Use discrete build steps when installing 3rd party dependencies and internals.
- Mirror explicit dependencies into your private repo.
- Use trusted sources for container bases.

### Vulnerability scanning

- Container Analysis on auto or on-demand.


## Service account

Artifact Registry Service Agent is created automatically and is granted a role with the same name.


## Audit logging

- Standard GCP logging, see doc: https://cloud.google.com/artifact-registry/docs/audit-logging


## Supported formats

- Docker Image Manifest V2, schemas 1 and 2.
- Open Container Initiative (OCI) Image format Specifications.
- Artifact Registry implements OCI Specification, an standard API protocol.
- Helm 3 chart packages in OCI format.
- Docker and Helm container images.
- Java, Node.js and Python language packages.
- Debian and RPM OS packages.


## Container concepts

### Registries

Stores and distributes container images and artifacts organized by name within repositories, e.g. Docker Hub and Artifact Registry.

### Repositories

Images and artifacts with the same name but different tags are organized into repositories. If not tag is set (not recommended), then latest is applied and the previous version becomes tagless.

The term repository is not always used consistently; within Artifact Registry it is more useful to use parts of the path to the image to identify the project, region or multi-region, and name of the image along with the tag or manifest digest to identify the correct version.

    docker push us-central1-docker.pkg.dev/PROJECT/quickstart-docker-repo/quickstart-image:tag1

- `us-central1` is the location of the repository
- `docker.pkg.dev` is the hostname for Docker repositories.
- `PROJECT` is the namespace created by your Google Cloud project ID.
- `quickstart-docker-repo` is the namespace under your project where you store images. In Artifact Registry, this part of the path is called the repository.
- `quickstart-image` is the repository for all versions of quickstart-image and is often referred to as the image.
- `tag1` is the tag specifying the version of the image.
	
### Images

An artifact can be anything, but an image is a container. Images are pushed and pulled.

	docker pull us-docker.pkg.dev/google-samples/containers/gke/hello-app:1.0

- `us-docker.pkg.dev` is the registry, including the us multi-region
- `/google-samples/containers/gke/` are the namespace and sub-namespaces. In Artifact Registry google-samples is the Google Cloud project and containers is referred to as the Artifact Registry repository
- `hello-app` is the artifact name
- `:1.0` is the tag specifying the version of artifact to be pulled
	
### Layers

Within a registry, images with common layers share them.

## Tags

- Specify version on push and pull.
- Can have >0 tags.
	
**Important** - On push, when a tag is supplied, latest is not applied (as it may have been before when not supplying a tag) so pulling `latest` may pull some old version that was previously pushed without a tag.

### Manifests

- Uniquely specifies the layers in the image.
- Identified by SHA-256 hash called a digest.
- Multiple versions of same image can be pushed to same tag, so the digest is reliable.
- Image referred to by tag may change, remember this when deploying security scanned safe images.


## Quickstart for docker

### Before you begin

- Select or create a project.
- Enable billing.
- Enable Artifact Registry API.

### Choose a shell

- Cloud Shell is hosted in the Console and has Docker and the gcloud SDK.
- Local; install all the bits.

**Important** - On Windows and Linux, the user running Docker needs to be added to the Docker security group to interact with registries.

### Create a Docker repository

- Create a Docker repo to store the sample for this quickstart.

		gcloud artifacts repositories create quickstart-docker-repo \ 
	  		--repository-format=docker \
	  		--location=us-central1 --description="Docker repository"

### Check it worked.

	gcloud artifacts repositories list

### Configure authentication

Must configure Docker to use gcloud to authenticate requests to GCP.

	gcloud auth configure-docker us-central1-docker.pkg.dev

**Note** - the region name is a DNS name prefix; you can ping it to check it's real.

This updates Docker's config JSON.

### Obtain an image to push

- cd into the directory you want to save the image in.
- Pull the same from Docker.

		docker pull us-docker.pkg.dev/google-samples/containers/gke/hello-app:1.0
	
### Add the image to the repo

- First, tag it with the target repo name. This configures docker push to push to a specific location.

		docker tag us-docker.pkg.dev/google-samples/containers/gke/hello-app:1.0 \
		    us-central1-docker.pkg.dev/PROJECT-ID/quickstart-docker-repo/quickstart-image:tag1

- `PROJECT` is your Google Cloud project ID. If your project ID contains a colon (:), see Domain-scoped projects.
- `us-central1` is the repository location.
- `docker.pkg.dev` is the hostname for Docker repositories.
- `quickstart-image` is the image name you want to use in the repository. The image name can be different than the local image name.
- `tag1` is a tag you're adding to the Docker image. If you didn't specify a tag, Docker will apply the default tag latest.
	
- In Docker, an image name is made up slash-separated components, optionally prefixed by a standard DNS registry hostname (:port is allowed). In the absence of a registry hostname, registry-1.docker.io is used by Docker.
- Tags may contain lower and uppercase ACSII letters, digits, underscores, periods and dashes but not start with a period or dash or be longer than 128 chars.

- Then push it with:

		docker push us-central1-docker.pkg.dev/PROJECT-ID/quickstart-docker-repo/quickstart-image:tag1
	
- You can pull it with:

		docker pull us-central1-docker.pkg.dev/PROJECT-ID/quickstart-docker-repo/quickstart-image:tag1
	
- You can delete any test stuff by deleting the rep:

		gcloud artifacts repositories delete quickstart-docker-repo --location=us-central1
	
