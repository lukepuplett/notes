# Cloud Storage - Concepts

## Key terms

All data belongs in a project.

### Buckets

- Everything must be in a bucket.
- Data access based on buckets.
- There is no limits to the number of buckets you can have in a project or location.
- There are limits to the rate you can create or delete buckets.
- Buckets need a _globally_ unique name and location at creation, which are immutable.

#### Bucket labels

- These are key-value metadata that allow grouping with other GCP resources.
- Max 64 labels per bucket, more limitations: https://cloud.google.com/storage/docs/key-terms#bucket-labels

### Objects

- No limit on objects in a bucket.
- They have object data and object metadata; key-value pairs.

#### Object names

- Name is metadata of less than 1024 bytes of UTF-8, and bucket-unique.
- Storage is flat but tools often break by `/` to present a false heirarchy.

#### Object immutability

- Objects are immutable; i.e. they cannot be appended to (or truncated).
- Wholesale replacement of an object is atomic; the old version is served until write completes.
- A _generation number_ for an object changes with with replacement.

**Note** - Replacement of an object is limited to once-per-second; expect `HTTP 429`.

### Resources

#### Resource names

- Each object has a resource name like this:

    projects/_/buckets/BUCKET_NAME/objects/OBJECT_NAME

- A `#NUMBER` appended to the end of the resource name points to a specific generation.
- `#0` is special and means the most recent.
- Add `#0` when the object name has and ending that can be misinterpreted as a generation number.

### Geo-redundancy

- Geo-redundant data is stored in at least two places with 100 miles between.
- Multi-region and dual-region objects are redundant, regardless of storage class.
- Replication is asynchronous but instantly zone-redundant.
- Some objects can take a while to replicate.
- It seems like failover is automatic with no change to the resource name.

#### Turbo replication

- Premium feature for certain dual-region buckets.
- Predictable _Recovery Point Objective_ or 15 minutes for _any_ new object.

## Request endpoints

- Supports HTTP/1.1, HTTP/2 and HTTP/3 protocols.

### Typical API requests

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

#### Encoding URI path parts

- Percent-encode these characters in the object name or query string of a request URI:

```
  !#$&'()*+,/:;=?@[ ]
```

### Cloud Console endpoints

When using the Cloud Console, use the following URLs:

#### Bucket list for a project

    https://console.cloud.google.com/storage/browser?project=PROJECT_ID

#### Object list for a bucket

    https://console.cloud.google.com/storage/browser/BUCKET_NAME
   
#### Details for an object

    https://console.cloud.google.com/storage/browser/_details/BUCKET_NAME/OBJECT_NAME

### gsutil endpoints

This Python CLI uses the JSON API, but can be reconfigured by setting `prefer_api=xml` in the `.boto` config file.

#### Performance and cost considerations

- The XML API uses the boto framework and so re-reads downloaded files to compute an MD5 if not present.
- For objects that do not include an MD5 in metadata, such as _composite objects_, this double the bandwidth and time.
- If working with composite objects, avoid `prefer_api=xml`.
- The XML API is shit in other ways, too.

### Custom domains

- You can map a _bucket-bound hostname_ with either an `A` or `CNAME`.
- See https://cloud.google.com/storage/docs/request-endpoints#custom-domains.

### Authenticated browser downloads

- Redirects to Google login and then drops a cookie.
- Requires `storage.objects.viewer` permission to download.
- Best use **Storage Object Viewer** role.
- To download an object in your browser, visit:

    https://storage.cloud.google.com/BUCKET_NAME/OBJECT_NAME

- All requests to `storage.cloud.google.com` require authentication, even when `allUsers` have permission.
- Use `storage.googleapis.com` for public anonymous access.
- See Accessing Public Data at: https://cloud.google.com/storage/docs/access-public-data

## Request preconditions

- Only perform the request if the generation or metageneration number meets criteria.
- Read-modify-write semantics to solve lost update problem.
- With a `Match` precondition, the object must have the same generation/metagen number, else you'll get a big fat `HTTP 412`.
- When `Match=0`, the request succeeds only if there are no live objects of that name.
- With a `NotMatch`, you get a `HTTP 304` if the gen/metagen matches (assume this is for a GET).
- Can also use ETags.
- XML API ETags for non-composite objects change only when content changes.
- ETags for composite objects and JSON API resources change whenever the content or metadata changes.
- Supports proper HTTP 1.1 ETags and HTTP conditional headers.

### Cost of preconditions

- You have to pay for them; a GET for the metadata.
- Avoid cost by caching or maintaining state etc.

### Preconditions in the JSON API

- A response containing an object or bucket resource carry `generation` and `metageneration` properties.
- It seems that the JSON API uses bizarre `ifGenerationMatch`, `ifGenerationNotMatch`, `ifMetagenerationMatch` and `ifMetagenerationNotMatch` query parameters for compose, insert or rewrite operations.
- For examples, see https://cloud.google.com/storage/docs/request-preconditions#_JSONAPI

**Note** - Generation and metagen preconditions are not accepted for ACL operations; use the access-control ETag entry resource instead.

**Note** - Preconditions in the XML API are done using custom headers. See https://cloud.google.com/storage/docs/request-preconditions#_XMLAPI

#### HTTP 1.1 ETags

The JSON API supports HTTP 1.1 ETags and the corresponding `If-Match` and `If-None-Match` headers for all resources, buckets, objects, and ACLs. An ETag is returned in the headers and in the content.

### Examples of race conditions and data corruption

#### Simultaneous read-modify-write

- Read-modify-write is well known so I won't repeat it here.

#### Multiple request retries

- Exponential backoff retry is recommended but can cause a problem on simple deletes.
- If network infra fails the delete 1.txt is retried and succeeds, then you create a new 1.txt.
- The infra comes back online and deletes 1.txt. **This is bad**.
- Avoid the above by issuing an `if-generation-match` precondition to ensure the back online request fails.
- Same goes for uploads; request should have `if-generation-match:0` to ensure no file exists (and no overwrites).
- Parallel uploads and large, composed objects are prone to corruption like this if chunk names are no gauranteed unique.

**Note** - `if-generation-match:0` cannot prevent double creation if the object is deleted and there's an upload pending in offline network equipment.

## Retry strategy

Tools, the Cloud Console and most client libraries automatically use retries. If implementing your own retries, consider whether it's safe to retry. See https://cloud.google.com/storage/docs/retry-strategy#build-your-own.

## Request rate and access distribution guidelines

Uses auto-scaling. Here's how to optimise for the way it's designed.

### Auto-scaling

Buckets have an intial IO capacity of (not sure what it means by _initial_):

- 1000 object writes per second, with a 1 write/sec for the same object.
- 5000 object reads per second.

As the request rate grows for a bucket, it's spreads out over >1 servers. This can take minutes so without ramping, can see errors.

#### Object key indexing

Supports consistent object listing (see Consistency, below), allowing data processing workflows to run. An object key index is maintained in lexicographic order, so objects with similar keys increases the chance of contention. The indexing is also auto-scaled and also needs ramping up when objects have similar prefixes.

### Best practices

- Ramp up, doubling every 20 minutes and pausing on errors.
- Exponential backoff for `HTTP 408` and `HTTP 429` and `HTTP 5xx`.
- Name things with wide key distribution.
- Avoid sequential names.
- It's very common for prefixes to be identical, e.g. `/folder1/folder2/1ni23rbo2389bo.txt`
- Random endings with identical prefixes are good but efficient scaling is reduced to it's prefix namespace.
- Randomness after a sequential prefix is not good.
- Randomize the order of items in bulk uploads.

## Sending batch requests

The JSON API supports batching to put several API calls in the same HTTP request which is good for:

- Updating metadata on many objects.
- Deleting many objects (see race conditions above).

Max 100 calls per request and less than 10MB payload.

**Note** - Same syntax as OData batch processing but **semantics differ**.

**Note** - No support for batched uploads or downloads.

More details, here: https://cloud.google.com/storage/docs/batch#details

## Caching

When cached, an object is copied to Google and/or internet CDNs. This poses a problem for stale responses.

### Built-in caching for Cloud Storage

- Cloud Storage behaves like a CDN by default.
- The `Cache-Control` metadata for a _publicly-accessible_ object determines web caching with `max-age` defaulting to 3600s.
- For public objects that are consumed by apps, consider `max-age` of 15-60s.
- Use `Cache-Control: no-store` as needed.

### Cloud Storage with Cloud CDN

- Best performance for public objects.
- To use it, must use external HTTP(S) Load Balancing with buckets as a backend.
- For help setting up, see: https://cloud.google.com/storage/docs/hosting-static-website
- Cloud CDN cache modes allow unified caching config, see: https://cloud.google.com/cdn/docs/caching#cache-modes
- Uses the same `Cache-Control` metadata unless overridden with a cache mode or TTL limit.
- For choosing between the built-in and Cloud CDN, see feature table, here: https://cloud.google.com/storage/docs/caching#with

## Object transcoding

- Transcoding is the automatic changing of a file's compression before being served; compressive or decompressive.
- Uses gzip.
- Supports decompressive.
- The stored hash is for the compressed bytes, so decompression invalidates integrity checking.
- Can store compressed to save cost and serve decompressed.
- Must be gzipped when stored and metadata `Content-Encoding: gzip`.
- When this is the case, it's decompressed on serve and `Content-Encoding` and `Content-Length` headers are removed.
- To prevent decompression, include `Accept-Encoding: gzip` or `Cache-Control: no-transform` which forces it, regardless of accepts.
- Useful to reduce egress costs or for validating integrity.

### Content-Type vs. Content-Encoding

- These are object metadata and effect behaviours worth noting.
- `Content-Type` should be included in all uploads.
- There is no check, obviously.
- `Content-Encoding` is optional and can be included for compressed uploads.
- Also no check.
- It's recommended to specify both when the object is gzipped.
- Including `-z/-Z` with gsutil will gzip and upload with `Content-Encoding: gzip` and `Cache-Control: no-transform`.
- Alternatively, `Content-Type: application/gzip` and no `Content-Encoding` set, but hides what's in the content.
- Also becomes ineligible for transcoding.
- Avoid uploading compressed objects and not setting `Content-Encoding`.
- Avoid uploading with just `Content-Encoding`, requests can be rejected.
- Never use `Content-Type: application/gzip` and `Content-Encoding: gzip` as it implies double compression.
- Never set `Content-Encoding: gzip` on non-compressed data.

**Note** - If actually double-compressing, see this: https://cloud.google.com/storage/docs/transcoding#gzip-gzip

### Using the Range header

- When transcoding, the `Range` header is ignored and the whole object is streamed.
- This is because it's impossible, obviously (and presumably Google don't have a streaming decompressor).
- But `Range` requests _do work_ for objects with `Content-Type: application/gzip` and no `Content-Encoding`.
- This is because the object is not transcoded so the request is getting a section of zipped bytes.
- I guess this might be useful for parallel range gets, then assemble and decompress?!

## Hashes and ETags: best practices

- You should validate downloaded bytes using CRC32C or MD5.
- Data can be corrupted by noisy network links, memory errors along the way, bugs, changes to source over a while.
- CRC32C is recommended for integrity since that's supported for composite blobs and XML multipart.
- Returned checksums are for the complete object, not a range.
- ETags are the MD5 when via XML API, and using Google-managed encryption keys, and not composite/not XML multipart uploaded.
- Otherwise, assume nothing about the ETag.
- Same object may have different ETags between JSON and XML.
- Cloud Storage validates when you copy or rewrite within Cloud Storage, or when MD5 or CRC32C is supplied in an upload.
- Can also just get metadata after upload and check.
- Object composition offer no server-side MD5 validation, so DIY.
- The gsutil `cp` and `rsync` operations validate and delete as needed.

For reading XML API hashes, see: https://cloud.google.com/storage/docs/hashes-etags#xml-api

### JSON API

- The `md5Hash` and `crc32c` properties contain base64-encoded hashes.
- Providing either metadata is optional. 
- Providing either as part of a resumable or JSON API multipart upload triggers server-side validation.
- Object will not be created on fail.
- When not provided, no validation happens but Cloud Storage computes them and puts them in metadata.

## Consistency

For cacheable, public objects, the consistency can be controlled. Object data and metadata is strongly consistent for these operations:

- Read-after-write
- Read-after-metadata-update
- Read-after-delete
- Bucket-listing
- Object-listing

Uploads, changes and deletes are immediately available for download from any Google location.

Bucket listing is the same. Create, update and delete to objects and metadata is immediately correct everywhere. Cached-objects may not be, however, depending on configuration.

Bucket configuration changes may take time–Google says 30s–to propagate, like enabling object versioning. Similarly, HMAC key changes can take 3 minutes, so wait before doing anything with it.

**Note** - Access changes are eventually consistent and take about 1 minute or more.

### Cache control and consistency

Obviously, cachable objects will be out in the wild unti their lifetime is up. Use the `Cache-Control` header as appropriate. The default is 60 minutes.

