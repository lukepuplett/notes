# Cloud Storage

## Concepts

### Key terms

All data belongs in a project.

#### Buckets

- Everything must be in a bucket.
- Data access based on buckets.
- There is no limits to the number of buckets you can have in a project or location.
- There are limits to the rate you can create or delete buckets.
- Buckets need a _globally_ unique name and location at creation, which are immutable.

##### Bucket labels

- These are key-value metadata that allow grouping with other GCP resources.
- Max 64 labels per bucket, more limitations: https://cloud.google.com/storage/docs/key-terms#bucket-labels

#### Objects

- No limit on objects in a bucket.
- They have object data and object metadata; key-value pairs.

##### Object names

- Name is metadata of less than 1024 bytes of UTF-8, and bucket-unique.
- Storage is flat but tools often break by `/` to present a false heirarchy.

##### Object immutability

- Objects are immutable; i.e. they cannot be appended to (or truncated).
- Wholesale replacement of an object is atomic; the old version is served until write completes.
- A _generation number_ for an object changes with with replacement.

**Note** - Replacement of an object is limited to once-per-second; expect `HTTP 429`.

#### Resources

##### Resource names

- Each object has a resource name like this:

    projects/_/buckets/BUCKET_NAME/objects/OBJECT_NAME

- A `#NUMBER` appended to the end of the resource name points to a specific generation.
- `#0` is special and means the most recent.
- Add `#0` when the object name has and ending that can be misinterpreted as a generation number.

#### Geo-redundancy

- Geo-redundant data is stored in at least two places with 100 miles between.
- Multi-region and dual-region objects are redundant, regardless of storage class.
- Replication is asynchronous but instantly zone-redundant.
- Some objects can take a while to replicate.
- It seems like failover is automatic with no change to the resource name.

##### Turbo replication

- Premium feature for certain dual-region buckets.
- Predictable _Recovery Point Objective_ or 15 minutes for _any_ new object.

### Request endpoints

- Supports HTTP/1.1, HTTP/2 and HTTP/3 protocols.

#### Typical API requests

- For general JSON API requests, excluding object uploads, use this:

    https://storage.googleapis.com/storage/v1/PATH_TO_RESOURCE

- For uploads:

    https://storage.googleapis.com/upload/storage/v1/b/BUCKET_NAME/o

    (the `o` isn't a typo)

- For batched requests:

    https://storage.googleapis.com/batch/storage/v1/PATH_TO_RESOURCE

- Optionally, for downloads:

    https://storage.googleapis.com/download/storage/v1/b/BUCKET_NAME/o/OBJECT_NAME?alt=media

**Note** - There's also an XML API, see https://cloud.google.com/storage/docs/request-endpoints#typical

##### Encoding URI path parts

- Percent-encode these characters in the object name or query string of a request URI:

```
  !#$&'()*+,/:;=?@[ ]
```

#### Cloud Console endpoints

When using the Cloud Console, use the following URLs:

##### Bucket list for a project

    https://console.cloud.google.com/storage/browser?project=PROJECT_ID

##### Object list for a bucket

    https://console.cloud.google.com/storage/browser/BUCKET_NAME
   
##### Details for an object

    https://console.cloud.google.com/storage/browser/_details/BUCKET_NAME/OBJECT_NAME

#### gsutil endpoints

This Python CLI uses the JSON API, but can be reconfigured by setting `prefer_api=xml` in the `.boto` config file.

##### Performance and cost considerations

- The XML API uses the boto framework and so re-reads downloaded files to compute an MD5 if not present.
- For objects that do not include an MD5 in metadata, such as _composite objects_, this double the bandwidth and time.
- If working with composite objects, avoid `prefer_api=xml`.
- The XML API is shit in other ways, too.

#### Custom domains

- You can map a _bucket-bound hostname_ with either an `A` or `CNAME`.
- See https://cloud.google.com/storage/docs/request-endpoints#custom-domains.

#### Authenticated browser downloads

- Redirects to Google login and then drops a cookie.
- Requires `storage.objects.viewer` permission to download.
- Best use **Storage Object Viewer** role.
- To download an object in your browser, visit:

    https://storage.cloud.google.com/BUCKET_NAME/OBJECT_NAME

- All requests to `storage.cloud.google.com` require authentication, even when `allUsers` have permission.
- Use `storage.googleapis.com` for public anonymous access.
- See Accessing Public Data at: https://cloud.google.com/storage/docs/access-public-data

### Request preconditions

- Only perform the request if the generation or metageneration number meets criteria.
- Read-modify-write semantics to solve lost update problem.
- With a `Match` precondition, the object must have the same generation/metagen number, else you'll get a big fat `HTTP 412`.
- When `Match=0`, the request succeeds only if there are no live objects of that name.
- With a `NotMatch`, you get a `HTTP 304` if the gen/metagen matches (assume this is for a GET).
- Can also use ETags.
- XML API ETags for non-composite objects change only when content changes.
- ETags for composite objects and JSON API resources change whenever the content or metadata changes.
- Supports proper HTTP 1.1 ETags and HTTP conditional headers.

#### Cost of preconditions

- You have to pay for them; a GET for the metadata.
- Avoid cost by caching or maintaining state etc.

#### Preconditions in the JSON API

- A response containing an object or bucket resource carry `generation` and `metageneration` properties.
- It seems that the JSON API uses bizarre `ifGenerationMatch`, `ifGenerationNotMatch`, `ifMetagenerationMatch` and `ifMetagenerationNotMatch` query parameters for compose, insert or rewrite operations.
- For examples, see https://cloud.google.com/storage/docs/request-preconditions#_JSONAPI

**Note** - Generation and metagen preconditions are not accepted for ACL operations; use the access-control ETag entry resource instead.

**Note** - Preconditions in the XML API are done using custom headers. See https://cloud.google.com/storage/docs/request-preconditions#_XMLAPI

##### HTTP 1.1 ETags

The JSON API supports HTTP 1.1 ETags and the corresponding `If-Match` and `If-None-Match` headers for all resources, buckets, objects, and ACLs. An ETag is returned in the headers and in the content.

