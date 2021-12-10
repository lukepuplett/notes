## Docker Intro

Dockerfiles have no file extension. They're instructions for building a container.

I have no idea if these are run in the context of the docker command line process or the sandboxed image being built, or something else.

The first line must be FROM, ARG or a parser directive in a # comment. Each line may produce an intermediate container which may be cached for speedy execution next time. I think these are called layers and they are committed.

The 'context' passed to Docker is a folder tree or can be a Git repo tree, including Git submodules. The Dockerfile is normally at the root of the tree. It's recommended to keep the Dockerfile in a folder along with only the files needed during build.

You can place a .dockerignore file in the root. Use Unix globbing syntax. Use # for comments.

https://docs.docker.com/engine/reference/builder/
https://docs.docker.com/develop/develop-images/multistage-build/

|Line|Explainer|
|-|-|
|`FROM node:12.16.3`|Base our container image on the 'node' image, with '12.16.3' tag (assumes 'latest'). Can have many FROM commands in a Dockerfile to make multi-stage builds.|
|`WORKDIR /code`||
|`ENV PORT 80`|Set environment PORT variable to 80 so it can be referenced in the Dockerfile, and in the final container process.|
|`COPY package.json /code/package.json`|Copy file|
|`RUN npm install`|Execute a command in a new shell. The other syntax looks like the CMD example and does not run in a shell.|
|`COPY . /code`|Copy all files in . to /code folder.|
|`CMD ["node", "src/server.js"]`|Optional command to run for executing containers, here in exec form.|

RUN executes command(s) in a new layer and creates a new image. E.g., it is often used for installing software packages.

CMD sets default command and/or parameters, which can be overwritten from command line when docker container runs.

ENTRYPOINT configures a container that will run as an executable



## Docker `run`

Process isolated by namespaces and OS kernel tricks. The host may be local or remote.

	docker run [OPTIONS] IMAGE[:TAG|@DIGEST] [COMMAND] [ARG…]

Must specify an IMAGE to derive a container from. Image developer can define default for:

- Detached or foreground
- Container ID
- Network
- Resource constraints, e.g. CPU

Override container and Docker's own defaults with OPTIONS. Only the person executing run can set:

- Detached or foreground.
- Identification
- IPC settings
- Network settings
- Restart policies
- Clean up
- Runtime constraints on resources
- Runtime privilege and Linux capabilities

Note - The bizarre contradiction that the docs say the developer can define a bunch of defaults for the container and then says only the person executing run can do this!

### Detached

Use `-d=true` or just `-d` to run detached. Such containers will exit when its root process exits.

Running executables that run start something and then exit will exit the container.

Use network connections or shared volumes to I/O with a detached container.

Use docker attach to reattach.

### Foreground

Docker attaches the console to the container's standard I/O. It can even pretend to be a TTY and pass signals.

Without -a Docker attaches just STDOUT and STDERR.

For interactive processes like a shell you must use -i -t in order to allocate a TTY.

### Container identification

Use --name to assign a name to a running container. Identifiers can be the full UUID or its short UUID (like Git) or its name.

Use --cidfile="" to have Docker write the process ID to the file, which is used for automation.

You can use image:tag to specify the version to run, or image@sha256:91a2… etc.

### PID settings

Use --pid="" to set the process namespace mode. Use "container:<name|id>" to join another container's PID namespace and "host" to use the host's.

PID namespace provides separation of processes, so PIDs can be reused at 1.

### UTS settings

Namespace for setting the hostname and domain that's visible to processes. Defaults to each container having their own, even those with --network="host".

### IPC settings

Use --ipc="mode" where mode is:

""	Daemon's default
"none"	Own IPC namespace with /dev/shm not mounted
"private"	Own IPC namespace.
"shareable"	Own with possibility to share with other containers.
"container:<name_or_id>"	Join another container.
"host"	Use host's.
	
### Network settings

--dns=[]	Set custom DNS servers
--network="bridge"	Use bridge, none, container:<name|id>, host or <network-name>|<network-id>
--network-alias=[]	Network-scoped alias for the container.
--add-hosts=""	Add line to /etc/hosts in form host:IP
--mac-address=""	Container's NIC MAC.
--ip=""	Container's NIC IPv4.
--ip6=""	Container's NIC IPv6.
--link-local-ip=[]	Set 1+ container's NIC's link local IPv4/6 addresses.
	
Default to network enabled and outbound allowed. If disabled then use STDIN and STDOUT.

Publishing ports and linking to other containers only works with the default bridge. Prefer using Docker network drivers over linking. Defaults to host's DNS servers.

Default MAC is generated using IP allocated to the container. Docker does not check if a supplied MAC is unique.

See doc: https://docs.docker.com/engine/reference/run/#network-settings

You can create a network using a Docker network driver or an external network driver plugin, then connect >1 containers to it. Then they can all communicate via IP address.

See doc above.

### Restart policies

This seems useful for process crash restarts. Can have a on-failure[:max-retries] setting.

Exponential delay is used to prevent floods. Max delay is 1 minute. Delay is reset on 10 seconds without crash.

	docker run --restart=on-failure:10 redis

### Exit status

When docker run exits with non-zero, it follows chroot standard.

See doc: https://docs.docker.com/engine/reference/run/#exit-status

### Clean up

By default a container's file system persists after container exit to make debugging easier. These retained file systems can pile-up so add --rm to clear down.

### Security configuration

See doc: https://docs.docker.com/engine/reference/run/#security-configuration

### Specify an init process

Use --init to use an init process as PID 1. Bunch of stuff I don't understand.

See doc: https://docs.docker.com/engine/reference/run/#specify-an-init-process

### Runtime constraints on resources

Far too many options to list.

See doc: https://docs.docker.com/engine/reference/run/#runtime-constraints-on-resources

### Logging drivers

Override default Docker daemon driver with a bunch of options, including splunk and AWS.

See doc: https://docs.docker.com/engine/reference/run/#runtime-constraints-on-resources

### Overriding Dockerfile image defaults

Cannot override FROM, MAINTAINER, RUN and ADD. Everything else is up for grabs at run time.



## Docker `build`

	docker build [OPTIONS] PATH | URL | -

Builds an image from a Dockerfile and a "context". The context is just a file tree rooted at the specified path or URL. Use COPY to reference a file in the context.

The URL can be a Git repo, tarball or plain text file. Git repos are first pulled to a local temp folder. To checkout a branch, tag or reference use a URL fragment. After the tag, you can append a colon : and specify a subfolder to set as the context root #main:sub

### Common CLI options

--file, -f	Path to the Dockerfile
--force-rm	Remove intermediate containers.
--label	Set metadata on the image using KEY=VALUE.
--output	Destination in format: type=local,dest=path)
--platform	Set platform if server is multi-platform capable.
--progress	Default auto. Type of progress, auto, plain, tty; use plain to show container output.
--pull	Always attempt to pull a newer version of the (base?) image.
--rm	Default true. Remove intermediate containers on success.
--tag, -t	Name and optionally tag in name:tag format.
--target	Choose stage to build.
