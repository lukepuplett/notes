# Cloud Scheduler

## Overview

It's cron jobs in the cloud. Each cron job is sent to a _target_ according to a schedule, where the task work is executed.

Targets are:

- Public HTTP/S resources.
- Pub/Sub topics.
- App Engine HTTP/S apps (presumably it means, private URLs).

You create them using the Console or `gcloud` CLI.

Only a single instance of a job should be run at any time. Cloud Schedule is for repeating jobs; for one-offs, consider using **Cloud Tasks**.

**Important** - Expect the same job to be run concurrently; your handler should be idempotent. Code can inspect headers on the request to determine which cron job is running. See https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations.jobs#httptarget

## How-to

### Creating and configuring cron jobs

#### CLI

    gcloud scheduler jobs create http JOB --location-LOCATION --schedule=SCHEDULE --uri=URI [optional flags]

Where `JOB` is its project-unique name. `LOCATION` defaults to the location of the current project's App Engine app, it exists.

The interval can be **unix-cron** format or "every 3 hours". The default HTTP method is **POST**. Specify optional HTTP body content sent as bytes for POST or PUT requests. Default retry values are usually fine.

See CLI reference for timezone, description and other options, including authentication.

https://cloud.google.com/sdk/gcloud/reference/scheduler/jobs/create/http

     gcloud scheduler jobs create http my-http-job --schedule "0 1 * * 0" --uri "http://myproject/my-url.com" --http-method GET

**Note** - If scheduler reports a failure even though it returned okay, adjust the `attempt_deadline` field.

Delete with:

    gcloud scheduler jobs delete [my-job]

#### Console

Visit the Cloud Scheduler console page and click **Create a job** and follow the UI.

