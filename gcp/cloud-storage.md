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
