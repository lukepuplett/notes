# Firebase Hosting

Notes taken 24th November, 2021.

https://firebase.google.com/docs/hosting

### IMPORTANT

I'm not finished reading the docs, but I don't think Firebase is a good choice for ASP.NET code unless Firebase URL rewrite rules can allow all * URLs to route to a Cloud Run instance.

### Capabilities

- Zero-config SSL.
- Static, dynamic and microservice APIs.
- Files cached on SSDs on CDNs.
- Emulators and sharable preview URLs.
- CLI deployment and simple rollbacks.
- GitHub integration.
- Localized content.
- Multiple sites from same resources.
- Link a project to Cloud Logging.
- URL rewriting.
- Custom headers.
- Integrates with StackBlitz and web IDEs.

### How does it work?

Firebase CLI deploys from local directories to Google's hosting servers. You can host static content, Cloud Functions for Firebase (note this sounds different to Cloud Functions, unless they've rebranded it) or Cloud Run for dynamic contents and APIs. All content is served from the closest edge CDN.

Firebase Local Emulator Suite lets you emulate your app and backend at a local URL, and share previews via a presumably public URL.

### Get started

1. Install the Firebase CLI.
2. Add static assets to a clean project folder, run `firebase init` to connect the folder to your Firebase cloud project.
3. Run `firebase emulators:start` to emulate Hosting and backend resources at a local URL, then share your changes at a preview URL by running `firebase hosting:channel:deploy` and setup GitHub integration to preview iterations.
4. Run `firebase deploy` to upload the latest snapshot.
5. Link your site to a Firebase Web App to use Google Analytics, and use Firebase Performance Monitoring.

### Initialize your project

Run `firebase init hosting` from the root of your local project directory and follow the prompts to select a directory to use as your www root directory, hosting your index.html and 404.html files.

If you select to make a SPA, Firebase automatically adds rewrite configurations. At the end, you'll have a firebase.json file containing configuration, and a .firebaserc which contains your project aliases.

Deploy with `firebase deploy --only hosting` to push just the web resources up, leaving functions and other backend stuff. Your app is viewable from `project_id.web.app` and `project_id.firebaseapp.com`

Google recommend setting up separate projects for testing and developing your app.

Test locally, share changes, then deploy live

### Local

Hosting at dev-time is done by the Firebase Local Emulator Suite. You can interact with your emulated Hosting content and config, plus your emulated functions, databases and rules.

By default, your locally-hosted app interacts with real functions, databases, rules, etc. but you can connect to emulated resources you've configured.

To start testing, run `firebase emulators:start` and stand well back. Visit http://localhost:5000. Add the following to your firebase.json to allow other LAN devices to see.

    "emulators": {
        // ...

        "hosting": {
        "port": 5000
        "host": "0.0.0.0"
        }
    }

There's also `firebase serve` but it's not recommended and the section is collapsed in the docs.

### Preview and share

Obviously, interacting with the preview URL makes calls to your real backend.

Run firebase hosting:channel:deploy channel_id where your channel ID is some string you can make up. Don't pass --only hosting to this command as this only deploys that stuff. The preview URL is something like project_id--channel_id-random_hash.web.app.

Run the command again with the same channel ID to update it with changes. GitHub Actions can publish to a preview channel upon commit.

It doesn't explain which set of backend resources it connects to, i.e. staging. Maybe that's done in my app code that pushed to preview.

### Deploy live

You deploy to your live channel via one of several different options, depending on needs.

Option 1 - Clone a preview channel to live:

	firebase hosting:clone source_site_id:source_channel_id target_site_id:live

For the site IDs, you can use the Firebase project ID and the channel ID is from the preview above, I guess.

Option 2 - From local directory to live:

	firebase deploy --only hosting

Just as we've seen before. Remember this is excluding functions and database rules etc.

### View your changes on your changes on live

The following hostnames will reflect recent changes.

`site_id.web.app` (like `project_id.web.app`) and the same for `.firebaseapp.com`, and any custom domains.

**Note** - it appears that the live site gets its own ID, different from the project ID. Not sure.

### Add a comment for the deploy

	firebase deploy --only hosting -m "Deploying the best new feature ever."

### Add predeploy and postdeploy scripted tasks

Connect shell scripts to the deploy command to, say, notify administrators of deployments or transpile JavaScript. This is done by adding Bash scripts to your firebase.json file.

See https://firebase.google.com/docs/cli#hooks

### Cache deployed content

When you deploy, the CDNs are cleared.


## Deploy to live & preview channels via GitHub pull requests

You can integrate deploys via GitHub Actions to:

- Create a new preview channel for every PR on the repo.
- Add a comment to the PR with the preview URL.
- Optionally, deploy your repo to live when the PR is merged.

**Important** - apps on preview URLs consume real backend resources; remember to use different projects to isolate production from test environments.

https://firebase.google.com/docs/hosting/github-integration


## Share project resources across multiple sites

Setup up to 36 Firebase Hosting sites in a single project. All sites can share resources. Each site has its own hosting configuration, hosts its own collection of content and has one ore more domains.

**Important** - it's recommended to use separate projects for different environments, like staging.

To add a site:

- Use the Hosting page in the web console.
- Or run `firebase hosting:sites:create site_id`
- Or call the Hosting REST API `projects.sites.create`

The site_id ends up as the prefix for web.app and firebaseapp.com domains. Thus, it must qualify as a legal domain name.

You can add custom domains to each subdomain.

It's recommended to setup deploy targets so you can refer to the correct site when issuing commands.

	firebase target:apply hosting target_name resource_identifier
	
Where target_name is a made up nickname for the site and `resource_identifier` is the `site_id` as listed in your Firebase project, e.g.

	firebase target:apply hosting blog myapp-blog

The settings are stored in `.firebaserc`

Then using this target name you can configure `firebase.json` sites using the target key value pair.

See example commands here: https://firebase.google.com/docs/hosting/multisites


## Connect a custom domain

Detailed steps here: https://firebase.google.com/docs/hosting/custom-domain

Free SSL for domains for all sites. Highlight points, here.

- Redirect a.com to www.a.com
- Verify ownership, likely bypassed when Google Domains.
- Supports zero-downtime migration.
- SSL certs are automatically distributed to CDNs.


## Configure hosting behaviour

See docs on globbing here: 

https://firebase.google.com/docs/hosting/full-config#glob_pattern_matching

You can configure:

- Which files to deploy from local directory.
- Custom 404 page.
- Redirects for moved pages.
- Rewrites to show the same content at >1 pages.
- Rewrites to serve a function or access a Cloud Run container from a Hosting URL.
- Rewrites to create a custom domain Dynamic Link, whatever that is.
- Add headers to requests and responses.
- Serve specific content based on language or country (i18n).

Configuration is defined in `firebase.json`. There's a full example file here:

https://firebase.google.com/docs/hosting/full-config#firebase-json_example

Note that `firebase.json` can also contain config for other Firebase services.

### Priority order of Hosting response

Config may overlap. Responses are determined in this order:

- Reserved namespace that begin with a `/_ _/*` path segment.
- Redirects.
- Exact-match static content.
- Rewrites.
- Custom 404.
- Default 404.

Note that the first redirect or rewrite rule with a matching URL pattern is applied; so order your rules carefully. Redirects beat rewrites.

If using "i18n rewrites" (internationalisation?) the exact-match and 404 handling order are expanded in scope to accommodate your i18n content.

### Specify which files to deploy

Use the `public` and `ignore` arrays to define file globbing to define which files in your project directory to deploy.

By default the public key is set to the whole folder you setup during init. The ignore key is optional.

### Configure a custom 404

Just deploy a 404.html file into the root.

### Redirects

Use redirects when you've moved stuff or for shortened URLs.

In the `hosting` object, add a `redirects` array of rules. Each rule has a `source` pattern and a `destination` URL and a `type` of either 301 or 302 for the HTTP code. Order matters.

You can capture segments of the URL using a colon prefix like source /blog/:post* which captures the whole URL segment beginning at post and destination https://blog.myapp.com/:post

For a Regex pattern, use named or unnamed capture groups. See documentation here:

https://firebase.google.com/docs/hosting/full-config#capture-url-segments-for-redirects

### Rewrites

As well as the same content for >1 URLs, rewrites can be used with apps that use HTML5 `pushState` for navigation (JavaScript editing the history stack). Thus, rewrites can be used to accept any URL at the server and then write JavaScript to control what is displayed.

Like redirects, create a `rewrites` array on the `hosting` object with rules that contain either `source` or `regex` and `destination` string values. Order matters.

The browser is sent the content of the destination URL.

Note that rewrite rules are only applied if a file or directory isn't at the requested URL.

### Direct requests to a function

Use rewrites to serve a function.

Important Firebase supports `us-central1` only for these functions.

    "hosting": {
    // Directs all requests from the page `/bigben` to execute the `bigben` function
    "rewrites": [ {
        "source": "/bigben",
        "function": "bigben"
    } ]
    }

Only `GET`, `PUT`, `POST`, `DELETE`, `HEAD`, `PATCH` and `OPTIONS` methods are supported.

### Direct requests to a Cloud Run container

Similarly this works for a container.

    "hosting": {
    "rewrites": [ {
    "source": "/helloworld",
    "run": {
        "serviceId": "helloworld",  // "service name" (from when you deployed the container image)
        "region": "us-central1"  // optional (if omitted, default is us-central1)
    }
    } ]
    }

Important Firebase supports many regions for this.

### Create custom Dynamic Links

Firebase Dynamic Links let users get a different experience depending on their device and can take the user into an app or to an app store.

See docs: https://firebase.google.com/docs/hosting/full-config#rewrite-dynamic-links

### Configure headers

Create a `headers` array in the `hosting` object and add an object with `source` (recommended) or `regex` (less recommended) URL pattern to trigger application of the `headers` array. Each header requires a `key` and `value` string.

Note that pattern matching for headers is comes before rewrite rules are applied.

See also Cache-Control: https://firebase.google.com/docs/hosting/manage-cache

**Important** - Firebase Hosting overwrites any `Strict-Transport-Security` header for default subdomains like *.web.app but any custom domains will serve your configured value.

### Control .html extensions

By adding the `cleanUrls` boolean value to the `hosting` object you can set whether to require the .html extension. Clients are redirected if they add .html.

### Configure internationalization (i18n)

You can serve based on language and country or a combination of those. These depend on the IP address and `Accept-Language` header from the browser.

You'll need to setup a directory for localised content. See this link for more:

https://firebase.google.com/docs/hosting/i18n-rewrites


## Load Firebase SDKs from reserved URLs

URLs beginning `/_ _` are reserved to make it easy to use other Firebase products. These URLs are also available on a local server.

Firebase Hosting is served over HTTP/2 so you can boost performance by loading files from the same origin. For example, the Firebase JavaScript libraries are served from `/_ _/firebase/…`

The docs says something about client-side SDK auto-configuration that I don't understand, here:

https://firebase.google.com/docs/hosting/reserved-urls#sdk_auto-configuration

A bunch of JS SDKs are available at these special URLs.

### Auth helpers

Firebase Authentication uses this URL namespace to serve JS and HTML for Oauth which is more secure (somehow) and you can use your custom domain for the `authDomain` option of `firebase.initializeApp()`.

### Reserved URLs and Service Workers

Something I don't understand about PWAs here: 

https://firebase.google.com/docs/hosting/reserved-urls#reserved_urls_and_service_workers


## Serve dynamic content and host microservices

You can integrate with Cloud Functions for Firebase and Cloud Run by directing HTTPS traffic to your functions and containers.

### Choosing a serverless option

|Consideration|Cloud Functions for Firebase|Cloud Run|
|-|-|-|
|Setup|The CLI bundles tasks into single commands, from initialising to building and deploying.|Containers are more customizable, so setup, build and deploy are more involved.|
|Runtime environment|Require Node.js 10 or 12.|Within container.|
|Language and framework support	JavaScript and TypeScript, and web frameworks.|Within container.|
|Timeout for requests|60 seconds (imposed by Firebase)|60 seconds (imposed by Firebase)|
|Concurrency|1 request per function instance|80 concurrent requests per container instance|
|Billing|Cloud Functions usage.|Cloud Run usage + Container Registry storage.|


## Connect Cloud Functions to Firebase Hosting

**Important** - At this time you can only use functions in `us-central1`.

Use `firebase init functions` from project root to setup and check for a `(root)/functions` directory.

Open `functions/index.js` and add/replace with this code:

    const functions = require('firebase-functions');

    exports.bigben = functions.https.onRequest((req, res) => {
    const hours = (new Date().getHours() % 12) + 1  // London is UTC + 1hr;
    res.status(200).send(`<!doctype html>
        <head>
        <title>Time</title>
        </head>
        <body>
        ${'BONG '.repeat(hours)}
        </body>
    </html>`);
    });

Test via the local emulator. Run `firebase emulators:start` from project root, then visit this URL.

	http://localhost:5001/PROJECT_ID/us-central1/bigben

See docs about HTTPS functions: https://firebase.google.com/docs/functions/http-events

Open your `firebase.json` and add this:

    "hosting": {
    // ...

    // Add the "rewrites" attribute within "hosting"
    "rewrites": [ {
        "source": "/bigben",
        "function": "bigben"
    } ]
    }

Test again via the emulator by visiting this URL (check port).

	http://localhost:5000/bigben

Run `firebase deploy --only functions,hosting` from project root then visit this URL:

	https://PROJECT_ID.web.app/bigben (or your custom domain)

### Use a web framework

You can use something like Express.js in Cloud Functions to serve dynamic content.

See docs: https://firebase.google.com/docs/hosting/functions#use_a_web_framework


## Serve dynamic content and host microservices with Cloud Run

**Important** - Each instance only handles max 80 concurrent requests, which isn't much for a website. Check instance pricing. Cloud Run auto-scales up and down and you only pay for resources consumed during request handling (check this).

Cloud Billing must be setup first, even for free allowance. Every Firebase project is also a Cloud Run project.

Open the Google APIs console and enable the Cloud Run API, then install and initialize the Cloud SDK, then run `gcloud config list` to check it's setup.

Write an app with a working Dockerfile. Interestingly, Google want you to use their tool to build the container, though you can build it locally.

	gcloud builds submit --tag gcr.io/PROJECT_ID/helloworld

This puts it in the Google Container Registry.

To build with local Docker, try this.

	docker build . --tag IMAGE_URL

Where the image URL is something like:

	docker.pkg.dev/cloudrun/container/hello:latest

Use `gcloud auth configure-docker` to let Docker auth with Google's Container Registry. Then push the container to Google with `docker push IMAGE_URL`

It doesn't mention that you can use your own registry. There are some free ones on the web.

Run this command to deploy your container.

	gcloud run deploy --image gcr.io/PROJECT_ID/helloworld

When prompted, select a region, confirm the service name and hit Y to allow unauthenticated invocations.

**Note** - Remember the region and service ID for use in the firebase.json rewrite rule.

Check the app at https://helloworld-RANDOM_HASH-us-central1.a.run.app - the CLI will spit out the URL once deployed.

Make changes to your firebase.json file.

    "hosting": {
    // ...

    // Add the "rewrites" attribute within "hosting"
    "rewrites": [ {
        "source": "/helloworld",
        "run": {
        "serviceId": "helloworld",  // "service name" (from when you deployed the container image)
        "region": "us-central1"     // optional (if omitted, default is us-central1)
        }
    } ]
    }

Then run `firebase deploy`

See Cloud Run docs about testing locally: https://cloud.google.com/run/docs/testing/local


## Manage cache behaviour

Requested static content is automatically cached on each CDN, and it's cleared on deploy. 

Dynamic content is not cached for obvious reasons, but you can configure caching via the `Cache-Control` header.

### Vary headers

Most of the time you don't need to worry about the Vary header. By adding a header name to the Vary header value, it will be included in (I assume) generating the unique key/hash for the request and thus, if that header changes, it will be cached separately.

**Note** - the headers are set by your code when forming the HTTP response and the Firebase Hosting infrastructure will see it on the way back out.

### Using cookies

When using Firebase Hosting with Cloud Run, **cookies are generally stripped from incoming requests**.

I had to stop after reading that. There's no way I can use Cloud Run via Firebase with no cookies.
