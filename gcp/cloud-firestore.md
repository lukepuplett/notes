# Cloud Firestore

## Concepts

Notes taken 8h November, 2021.

https://cloud.google.com/firestore/docs/data-model

### Documents and Collections

Collections hold documents and each are created implicitly, no need to pre-create the collections.

Each document is identified by a name you provide or let Firestore make random IDs. They're not JSON but similar and support similar nested arrays and objects (called maps).

**Warning** - Documents are limited to 1MB.

**Warning** - Only the first 1,500 bytes of UTF-8 are considered in queries, maybe len(750).

**Warning** - Arrays cannot nest arrays.

Every document is uniquely identified by its location. Create a reference like this.

    DocumentReference r = db.Collection("users").Document("test");
    // or
    DocumentReference r = db.Document("users/test");

You can also make a reference to a collection, alone, and get access to a different bunch of functionality.

Bizarrely, you can create collections within a document, affording hierarchical organisation up to 100 levels deep.

    var r = db.Collection("top").Document("a").Collection("sub").Document("b");

**Warning** - Deleting a document does not delete its sub-collections!

### Indexes

Firestore has single-field and composite indexes. It uses indexes for all queries. Most indexes can be auto-created as you test your app. I think this happens via error messages from testing during development that contain links through which you can create the missing index.

See https://cloud.google.com/firestore/docs/query-data/indexing?hl=en-AU&_ga=2.186531241.-147496122.1638202233#error-links

New indexes are backfilled using _long-running operations_ and take at least a few minutes even for an empty database. View the operation with:

    gcloud firestore operations list
    gcloud firestore operations describe OP-ID

Single-field indexes are managed by configuring the auto-indexing and exemptions. By default an SFI is made for each field in a document and each sub-field in a document map value. For normal fields, it makes an ascending index and a descending one at collection-scope. It does this also for each map field. Arrays get a collect-scope "array-contains" index.

SFIs with collection-group scope (see below) are not maintained by default.

Use the exemption settings to control indexing; forcing on and turning off it seems. Obviously a map's fields inherit it's parent's setting, but you can define a SFI for a map field.

Google says you can rely on automatic indexing and errors to manage indexes, for most apps, but consider exempting large string fields, high write rates to a collection containing documents with sequential values, large array or map fields.

See https://cloud.google.com/firestore/docs/concepts/index-overview#automatic_indexing

**Note** - if an indexed field's value is sequential as inserted, e.g. a timestamp and IoT data, then max writes is 500/second. Exempt this field if possible.

**Note** - exemptions only impact auto indexes and don't impact manual composite settings.

You have to manually configure composite indexes and they can only contain one array field. Firestore helps you identify candidate composite indexes. You'll get errors for queries that need a CFI and a link to help make it. You can also make indexes in the console or via the Firebase CLI.

Though SFIs are auto or controlled via exemptions and CFIs are controlled explicitly, both need a mode and scope setting.

|Index mode|Description|
|-|-|
|Ascending|`<, <=, ==, =>, >` and `in` clauses and sorting output.|
|Descending|As above.|
|Array-contains|`array-contains` and `array-contains-any` clauses.|

**Collection scope** is the default and return results from a single collection.

**Collection group scope** includes all collections with the same ID via a collection group query. For example, each `city` document has a `landmarks` sub-collection then you can query across all documents in all `landmarks`.

Single-field indexes allow for the following kind of code.

    citiesRef.where('country', 'in', ["USA", "Japan", "China"])

    // Compound equality queries
    citiesRef.where("state", "==", "CO").where("name", "==", "Denver")
    citiesRef.where("country", "==", "USA")
             .where("capital", "==", false)
             .where("state", "==", "CA")
             .where("population", "==", 860000)
    citiesRef.where("regions", "array-contains", "west_coast")
    // array-contains-any and array-contains use the same indexes
    citiesRef.where("regions", "array-contains-any", ["west_coast", "east_coast"])

Whereas composite indexes are needed for queries that mix operators.

    citiesRef.where("country", "==", "USA").orderBy("population", "asc")
    citiesRef.where("country", "==", "USA").where("population", "<", 3800000)
    citiesRef.where("country", "==", "USA").where("population", ">", 690000)
    // in and == clauses use the same index
    citiesRef.where("country", "in", ["USA", "Japan", "China"])
             .where("population", ">", 690000)

And of course, using orderBy() would need a corresponding directional index.

This is what a collection group scope query looks like where landmarks :

    var landmarksGroupRef = db.collectionGroup("landmarks");

    landmarksGroupRef.where("category", "==", "park")
    landmarksGroupRef.where("category", "in", ["park", "museum"])

Firestore can use more than one index in a query. Several good indexes can be combined to cover many queries.

### Index-entry count limit

You can imagine that indexes take up a ton of storage space and cost money.

|Limit|Details|
|-|-|
|Max composite indexes per DB|200|
|Max exemptions per DB|200|
|Max index entries per document|40,000|
|Max size of an entry|7.5KiB|
|Max sum of sizes|8MiB|
|Max size of an index field value|1500 bytes|


## Transactional Data Contention, Serializability, and Isolation

For a transaction to succeed, the docs being read must remain unmodified by operations outside the transaction. If another operation tries to change one of those docs, that op enters a state of contention.

Firestore resolves this by waiting or failing the op and client libraries will auto retry up to a point.

Client-side libs use optimistic concurrency, because locking under unreliable mobile connections is dumb, while server-side libs use pessimistic.

Mobile libs complete only if the docs are not changed during the transaction operation, else it retries a few times. Server libs use locks to get the job done.

Firestore uses serializable isolation by commit time. It assigns each transaction a commit time instant. All changes within a transaction take place in that instant, though actual execution takes a span of time but Firestore guarantees that transactions are committed in order of commit time, and concurrent transactions are isolated with a later commit time.

Queries and reads inside a transaction do not see the results of previous writes inside that transaction. Even if you modify or delete a document within a transaction, all document reads in that transaction return the version of the document at commit time, before the transaction's write operations. Read operations return nothing if the document did not exist then.

**Note** - Reads must come before writes.

To be honest, I don't fully understand this. There are no tangible examples using a query language or code.


## Encryption

All data is encrypted AES 256 on disk and no setup is needed. 


## Datastore mode

You can choose to be fully backward-compatible with the old Datastore product when creating the cloud resource.

This mode removes some previous limitations but then limits some of Firestore's behaviour. Best to look at the docs.


## Best practices

 - Avoid `.` and `..` document IDs and `/` within. Do not use monotonically increasing document IDs, else hotspots. Don’t put stupid characters in field names.
 - Do not use offsets. Instead, use **cursors**. Using an offset only avoids returning the skipped documents, but they're still retrieved internally and you pay for it.
 - Remember to implement retries when not using Google's libs.
 - There's a bunch of hints for Realtime Updates but I'm not planning on using that here. Look up the docs.
 - Enabling a Cloud Function with more than 2,000 Firestore events per second can _temporarily_ increase error rate and _temporarily_ increase latency by several minutes. Ahead of enabling a high-traffic function, contact support to prepare your database for high-traffic functions and avoid the increased latency. I assume this is because of the time it takes to auto scale.
 - Avoid writing to the same document more than once a second, else your application will experience contention, including higher latency, timeouts, and other errors. I looks like this is fine in short bursts.
 - Avoid high read or write rates to lexicographically close documents, or your application will experience contention errors.
 - It seems sequential IDs and failing to ramp up are the main problems for performance.
 - Ramp-up slowly when moving traffic to a new collection. The documented suggestions don't make sense to me. I think a good idea might be to pick a test, like IDs beginning A, then read from the new, then try the old, and move the document to new. Then add more qualifying documents to the test conditions each day. Or add a field to the user entity for migration purposes, with a random future date-time.


## Transactions and batches

You can run a bunch of code within a transaction which either succeeds or fails completely. Reads must all come first. The code within the transaction may run many times on different threads when there is contention, so don't do anything that would be messed up by this.

The code in the transaction is normal C# within an async anonymous function block. Presumably the SDK begins a transaction, gets a handle, runs the code under that handle and then ends the transaction.

Transactions are designed to fail and automatically retry if the documents involved are changed during the transaction execution.

**Author's Note** - On the surface this is a strange design as my intuition is not to overwrite "their" changes that could influence whether or not to make "our" changes, which could be true, but I guess this design is to support atomic updates such as incrementing a counter; the transaction body would reread the latest changed data and update. To prevent whacking other's changes, I guess the command needs configuring with zero retries somehow.

**Note** - transactions have a max size of 10MiB based on documents and indexes changed.

Incrementing a counter is a good example of something to wrap in a transaction because it involves a read, increment and write.

Batched writes are for changes that don't need reads as part of the change itself. Otherwise they're similar to transactions, but changed to any read documents will not fail the batch.

A batched write can up to 500 operations.

Using **Firestore Security Rules**, mobile and web clients can validate data before it is committed using a `getAfter` rule. This can be used to ensure data across two or more documents is consistent. 


## Offline data

Android and Apple native app client libraries automatically enable offline data. No code changes are needed when reading or writing data.

For Web libraries, offline is disabled by default so to enable it requires calling `enablePersistence`. If you app handles sensitive information then ask the user if they're on a trusted device before enabling the feature.

**Note** - In the sample JavaScript the method called is `enableIndexedDbPersistence`.

The cache sized can be configured during initialization. Firestore caches every document received.

When locally cached documents, collections and queries change, your listeners are notified. When querying a collection offline and no cached documents are available, an empty set is returned, but when fetching a particular document, you'll get an error.

There's much more on this subject here: https://cloud.google.com/firestore/docs/manage-data/enable-offline

## Export and import

You can export data as a form of backup in order to recover from accidental deletions, or to load the data into **Big Query**. You can export and import via the gcloud CLI or the REST or gRPC APIs.

Exports are made to Google Cloud Storage into a bucket. For more information, see https://cloud.google.com/firestore/docs/manage-data/export-import.


## Querying

### Getting data

You can either get the data directly via collections, documents or queries, or setup a listener and get data-change events. I think this is done via a Snapshot type object.

    DocumentReference docRef = db.Collection("cities").Document("BJ");
    DocumentSnapshot snapshot = await docRef.GetSnapshotAsync();
    if (snapshot.Exists)
    {
        City city = snapshot.ConvertTo<City>();
    }

**Note** - The City type is a POCO with its properties decorated with Google's own serializable sort of attribute.

Getting many documents is similar. Here, using Dictionary<string, object>

    Query capitalQuery = db.Collection("cities").WhereEqualTo("Capital", true);
    QuerySnapshot capitalQuerySnapshot = await capitalQuery.GetSnapshotAsync();
    foreach (DocumentSnapshot documentSnapshot in capitalQuerySnapshot.Documents)
    {
        Console.WriteLine("Document data for {0} document:", documentSnapshot.Id);
        Dictionary<string, object> city = documentSnapshot.ToDictionary();
        foreach (KeyValuePair<string, object> pair in city)
        {
            Console.WriteLine("{0}: {1}", pair.Key, pair.Value);
        }
        Console.WriteLine("");
    }

You can see the subcollections for a document using this code which uses an async enumerator.

    DocumentReference cityRef = db.Collection("cities").Document("SF");
    IAsyncEnumerable<CollectionReference> subcollections = cityRef.ListCollectionsAsync();
    IAsyncEnumerator<CollectionReference> subcollectionsEnumerator = subcollections.GetAsyncEnumerator(default);
    while (await subcollectionsEnumerator.MoveNextAsync())
    {
        CollectionReference subcollectionRef = subcollectionsEnumerator.Current;
        Console.WriteLine("Found subcollection with ID: {0}", subcollectionRef.Id);
    }

### Realtime updates

You can listen for changes using this code, if you have a long-running server-side application.

    DocumentReference docRef = db.Collection("cities").Document("SF");
    FirestoreChangeListener listener = docRef.Listen(snapshot =>
    {
        Console.WriteLine("Callback received document snapshot.");
        Console.WriteLine("Document exists? {0}", snapshot.Exists);
        if (snapshot.Exists)
        {
            Console.WriteLine("Document data for {0} document:", snapshot.Id);
            Dictionary<string, object> city = snapshot.ToDictionary();
            foreach (KeyValuePair<string, object> pair in city)
            {
                Console.WriteLine("{0}: {1}", pair.Key, pair.Value);
            }
        }
    });

**Note** - when you app process write changes, listeners will be notified before the writes are sent and committed to the Firestore service. Retrieved documents have a `metadata.HasPendingWrites` property but this is not available in the C# library yet.

You can control the granularity of change events. For example, by default you will not get notified of changes to metadata because you'll end up two notifications, one within process locally and another once it's written and the metadata changes. You can opt-in but it's not supported in C#.

**Note** - You pay for all the data you're sent.

Here's how to listen to the results of a query.

    CollectionReference citiesRef = db.Collection("cities");
    Query query = db.Collection("cities").WhereEqualTo("State", "CA");

    FirestoreChangeListener listener = query.Listen(snapshot =>
    {
        Console.WriteLine("Callback received query snapshot.");
        Console.WriteLine("Current cities in California:");
        foreach (DocumentSnapshot documentSnapshot in snapshot.Documents)
        {
            Console.WriteLine(documentSnapshot.Id);
        }
    });

And you stop listening like this.

    await listener.StopAsync();

### Simple queries

Here's a simple query.

    CollectionReference citiesRef = db.Collection("cities");
    Query query = citiesRef.WhereEqualTo("State", "CA");
    QuerySnapshot querySnapshot = await query.GetSnapshotAsync();
    foreach (DocumentSnapshot documentSnapshot in querySnapshot.Documents)
    {
        Console.WriteLine("Document {0} returned by query State=CA", documentSnapshot.Id);
    }

As seen before, called ToDictionary to get the document content.

    Query capitalQuery = db.Collection("cities").WhereEqualTo("Capital", true);
    QuerySnapshot capitalQuerySnapshot = await capitalQuery.GetSnapshotAsync();
    foreach (DocumentSnapshot documentSnapshot in capitalQuerySnapshot.Documents)
    {
        Console.WriteLine("Document data for {0} document:", documentSnapshot.Id);
        Dictionary<string, object> city = documentSnapshot.ToDictionary();
        foreach (KeyValuePair<string, object> pair in city)
        {
            Console.WriteLine("{0}: {1}", pair.Key, pair.Value);
        }
        Console.WriteLine("");
    }

You can use `<. <=, >, >=, !=, array-contains, array-contains-any, in` and `not-in` operators.

**Note** - In C# it looks as if these operators are method names on the collection reference.

**Note** - For `!=` and `not-in` queries, documents with null or missing properties are **not** returned.

**Note** - In a compound query, range `<. <=, >, >=` and `not` equals or `not-in` must all filter on the same field.

**Note** - Limited to one array-contains clause per query and you cannot combine it with array-contains-any.

Use the in operator, WhereIn in C#, to combine up to 10 == clauses on the same field with a logical OR like this.

    Query query = citiesRef.WhereIn("Country", new[] { "USA", "Japan" });

**Note** - The not-in version excludes documents where the field does not exist.

And use array-contains-any to combine up to 10 array-contains clauses on the same field with a logical OR.

    Query query = citiesRef.WhereArrayContainsAny("Regions", new[] { "west_coast", "east_coast" });

Naturally, the documents are de-duped in case of many matches. Also, when the field contains a string, as opposed to an array, there is no match made just on that string.

For `WhereIn`, you can pass in an array but it will look for an exact match.

    Query query = citiesRef.WhereIn("Regions",
        new[] { new[] { "west_coast" }, new[] { "east_coast" } });

**Note** - Limited to one in, not-in or array-contains-any clause per query and you cannot combine these operators in the same query.

**Note** - You can't combine not-in with !=.

**Note** - You can't order by a field included in an equality == or in clause.

### Compound queries

You can chain operators to build logical AND queries, but before combining equality operators with inequality and range operators, you must create a composite index.

    Query chainedQuery = citiesRef
        .WhereEqualTo("State", "CA")
        .WhereEqualTo("Name", "San Francisco");

You can perform range queries only on a single field and you can include max one array clause per compound query.

This is okay.

    Query rangeQuery = citiesRef
        .WhereGreaterThanOrEqualTo("State", "CA")
        .WhereLessThanOrEqualTo("State", "IN");

But this is all kinds of wrong.

    Query invalidRangeQuery = citiesRef
        .WhereGreaterThanOrEqualTo("State", "CA")
        .WhereGreaterThan("Population", 1000000);

### Collection group queries

A collection group consists of collections with the same ID. Normally, queries work on a single collection, but check this out.

This queries within all landmarks subcollections in all documents, no less.

    Query museums = db.CollectionGroup("landmarks").WhereEqualTo("Type", "museum");
    QuerySnapshot querySnapshot = await museums.GetSnapshotAsync();
    foreach (DocumentSnapshot document in querySnapshot.Documents)
    {
        Console.WriteLine($"{document.Reference.Path}: {document.GetValue<string>("Name")}");
    }

Mind you, you'd have to have created the appropriate index first.

Query limitations

 - There's limited support for OR queries. Use the in and array-contains-any operators for logical OR, else make a bunch of queries and merge the results.
 - In a compound query, range operators and not equals, must all filter on the same field.
 - Max one array-contains and it can't be combined with array-contains-any.
 - Max one in, not-in, array-contains-any clause per query and you can't combine in, not-in and array-contains-any in the same query.
 - You can’t order your query by a field included in an equality == or in clause.
 - The sum of filters, orders and parent doc path in a query cannot exceed 100.

### Ordering and limiting

Unless specified, documents are ordered by document ID.

**Note** - The inclusion of an order by clause will omit the document if the field does not exist.

    Query query = citiesRef.OrderBy("Name").Limit(3);

### Ordering limitations

If you include a filter with a range comparison (`<, <=, >, >=`), your first ordering must be on the same field, but you can add secondary ordering.

This code won't work.

    Query query = citiesRef
        .WhereGreaterThan("Population", 2500000)
        .OrderBy("Country");

And you can't order by a field in an equality or in clause.

### Pagination and cursors

You use this to return a subset or paginate, but often you should just use a normal range query. Use `startAt` (inclusive) and `startAfter` (exclusive) and optionally `endAt` and `endAfter`.

    Query query = citiesRef.OrderBy("Population").StartAt(1000000);

You can use a document snapshot to define a query cursor.

    CollectionReference citiesRef = db.Collection("cities");
    DocumentReference docRef = citiesRef.Document("SF");
    DocumentSnapshot snapshot = await docRef.GetSnapshotAsync();
    Query query = citiesRef.OrderBy("Population").StartAt(snapshot);

Then combine this with the limit method to paginate a query.

    CollectionReference citiesRef = db.Collection("cities");
    Query firstQuery = citiesRef.OrderBy("Population").Limit(3);

    QuerySnapshot querySnapshot = await firstQuery.GetSnapshotAsync();
    long lastPopulation = 0;
    foreach (DocumentSnapshot documentSnapshot in querySnapshot.Documents)
    {
        lastPopulation = documentSnapshot.GetValue<long>("Population");
    }

    Query secondQuery = citiesRef.OrderBy("Population").StartAfter(lastPopulation);
    QuerySnapshot secondQuerySnapshot = await secondQuery.GetSnapshotAsync();

When using a cursor based on a field like this, you can make it more precise by adding another field value. Above, if two cities have the same lastPopulation, then the start is ambiguous.

See this strange code.

    Query query1 = db.Collection("cities").OrderBy("Name").OrderBy("State").StartAt("Springfield");
    Query query2 = db.Collection("cities").OrderBy("Name").OrderBy("State").StartAt("Springfield", "Missouri");

**Note** - It's odd because two arguments are sent to the StartAt method, the first is Name and the second is for State but that's not explicit.


# Cloud Firestore API Client Library for .NET

It seems that there are two SDKs, the normal Google Cloud client libraries for Firestore, and the Firebase Admin SDKs which bundle the former along with some other Firebase features, though there's more limited language support for the latter.

## Getting started

Google's documentation for the C# library provides a great jumping off point, here:

https://cloud.google.com/dotnet/docs/reference/Google.Cloud.Firestore/latest/index

and here:

https://firebase.google.com/docs/firestore/data-model

### Sample code

Here's some useful sample code:

    FirestoreDb db = FirestoreDb.Create(projectId);

    // Create a document with a random ID in the "users" collection.
    CollectionReference collection = db.Collection("users");
    DocumentReference document = await collection.AddAsync(new { Name = new { First = "Ada", Last = "Lovelace" }, Born = 1815 });

    // A DocumentReference doesn't contain the data - it's just a path.
    // Let's fetch the current document.
    DocumentSnapshot snapshot = await document.GetSnapshotAsync();

    // We can access individual fields by dot-separated path
    Console.WriteLine(snapshot.GetValue<string>("Name.First"));
    Console.WriteLine(snapshot.GetValue<string>("Name.Last"));
    Console.WriteLine(snapshot.GetValue<int>("Born"));

    // Or deserialize the whole document into a dictionary
    Dictionary<string, object> data = snapshot.ToDictionary();
    Dictionary<string, object> name = (Dictionary<string, object>) data["Name"];
    Console.WriteLine(name["First"]);
    Console.WriteLine(name["Last"]);

    // See the "data model" guide for more options for data handling.

    // Query the collection for all documents where doc.Born < 1900.
    Query query = collection.WhereLessThan("Born", 1900);
    QuerySnapshot querySnapshot = await query.GetSnapshotAsync();
    foreach (DocumentSnapshot queryResult in querySnapshot.Documents)
    {
        string firstName = queryResult.GetValue<string>("Name.First");
        string lastName = queryResult.GetValue<string>("Name.Last");
        int born = queryResult.GetValue<int>("Born");
        Console.WriteLine($"{firstName} {lastName}; born {born}");
    }

### Connecting to the emulator

 - Connect to the emulator endpoint.
 - Use `ChannelCredentials.Insecure` as the credentials.
 - Specify `Authorization: Bearer <owner>` header on each request.

ProjectId may be null, in which case the project will be automatically detected if possible.

This code only uses the emulator when the following environment variable is set.

    // variable FIRESTORE_EMULATOR_HOST="localhost:8080"

    FirestoreDb db = new FirestoreDbBuilder
    {
        ProjectId = projectId,
        EmulatorDetection = EmulatorDetection.EmulatorOrProduction
    }.Build();
    // Use db as normal

## Data model

Notes taken from https://cloud.google.com/dotnet/docs/reference/Google.Cloud.Firestore/latest/datamodel

Based around collections and documents, which can contain more collections though the document data itself does not contain subcollections. Documents can contain nested data via maps. I can't remember if you can further nest complex objects in map properties or if they're limited to simple key-value pairs.

The C#.NET types used for the supported Firestore data model are obvious except for some specialised Google ones.

#### Byte arrays

Can use `Google.Cloud.Firestore.Blob`, or `byte[]`, or `Google.Protobuf.ByteString` to store 1,048,487 bytes with first 1,500 indexed.

#### Date-times

Can use `Google.Cloud.Firestore.Timestamp`, or `DateTime`, or `DateTimeOffset`, or `Google.Protobuf.WellKnownTypes.Timestamp`. Precise only to microseconds (rounded down). `DateTime` values must be `Utc` to be converted; For `DateTimeOffset` values,  **UTC-offset information is discarded.**

#### Geographical point

Can use `Google.Cloud.Firestore.GeoPoint` or `Google.Type.LatLng`.

#### Reference

Must use `Google.Cloud.Firestore.DocumentReference`.

### Mapping with attributed classes

Can serialize to and from a POCO with Google's attributes.

    [FirestoreData]
    public class City
    {
        [FirestoreProperty]
        public string Name { get; set; }

        [FirestoreProperty]
        public string State { get; set; }

        [FirestoreProperty]
        public string Country { get; set; }

        [FirestoreProperty("Capital")]
        public bool IsCapital { get; set;  }

        [FirestoreProperty]
        public long Population { get; set; }
    }

Use it like this:

    FirestoreDb db = FirestoreDb.Create(projectId);
    
    // Create a document with a random ID in the "cities" collection.
    CollectionReference collection = db.Collection("cities");
    City city = new City
    {
        Name = "Los Angeles",
        Country = "USA",
        State = "CA",
        IsCapital = false,
        Population = 3900000L
    };
    DocumentReference document = await collection.AddAsync(city);

    // Fetch the data back from the server and deserialize it.
    DocumentSnapshot snapshot = await document.GetSnapshotAsync();
    City citySnapshot = snapshot.ConvertTo<City>();
    Console.WriteLine(citySnapshot.Name); // Los Angeles

There are also several special attributes which appear to be used to get at some document metadata, e.g.:

    [FirestoreData]
    public class ChatRoom
    {
        [FirestoreDocumentId]
        public DocumentReference Reference { get; set; }

        [FirestoreDocumentCreateTimestamp]
        public Timestamp CreateTime { get; set; }

        [FirestoreDocumentUpdateTimestamp]
        public Timestamp UpdateTime { get; set; }

        [FirestoreDocumentReadTimestamp]
        public Timestamp ReadTime { get; set; }

        [FirestoreProperty]
        public string Name { get; set; }

        [FirestoreProperty]
        public bool Public { get; set; }
    }

Where `FirestoreDocumentId` must be a `string` or `DocumentReference` for a complete and useable reference object. The timestamps must be a `DateTime`, `DateTimeOffset`, `Google.Cloud.Firestore.Timestamp` or `Google.Protobuf.WellKnownTypes.Timestamp` type and can be nullable.

### Mapping with dictionaries

Use `Dictionary<string, object>` or replace `object` with a type if the document has uniform values. Use `DocumentSnapshot.ToDictionary` to deserialize, e.g.:

    FirestoreDb db = FirestoreDb.Create(projectId);

    // Create a document with a random ID in the "cities" collection.
    CollectionReference collection = db.Collection("cities");
    Dictionary<string, object> city = new Dictionary<string, object>
    {
        { "Name", "Los Angeles" },
        { "Country", "USA" },
        { "State", "CA" },
        { "Capital", false },
        { "Population", 3900000L }
    };
    DocumentReference document = await collection.AddAsync(city);

    // Fetch the data back from the server and deserialize it.
    DocumentSnapshot snapshot = await document.GetSnapshotAsync();
    Dictionary<string, object> cityData = snapshot.ToDictionary();
    Console.WriteLine(cityData["Name"]); // Los Angeles

### Mapping with anonymous types

These can be used for serialization but not deserialization, for obvious reason. They're great for _partial_ updates.

    FirestoreDb db = FirestoreDb.Create(projectId);

    // Create a document with a random ID in the "cities" collection.
    CollectionReference collection = db.Collection("cities");
    var city = new
    {
        Name = "Los Angeles",
        Country = "USA",
        State = "CA",
        Capital = false,
        Population = 3900000L
    };
    DocumentReference document = await collection.AddAsync(city);

    // Update just the population using another anonymous type
    await document.SetAsync(new { Population = 3900005L }, SetOptions.MergeAll);

    // Fetch the latest document and print the population
    DocumentSnapshot snapshot = await document.GetSnapshotAsync();
    Console.WriteLine(snapshot.GetValue<long>("Population")); // 3900005

### Custom converters

Create a custom converter by implementing the following interface. This might be useful for dealing with `DateTimeOffset` and lost UTC information, but care is needed to ensure it is still correctly indexable, sortable and queryable when in it's in its storage presentation. Perhaps an alternative is to store `DateTimeOffset` in it's ISO format but it would be useless in queries, so additionally store a `DateTime` version in UTC.

    public interface IFirestoreConverter<T>
    {
        object ToFirestore(T value);    // convert to storage version
        T FromFirestore(object value);  // convert to application code version
    }

Can apply `[FirestoreDeserializationConfigurationAttribute]` to `FromFirestore` for added functionality, see documentation below:

https://cloud.google.com/dotnet/docs/reference/Google.Cloud.Firestore/latest/datamodel#custom-converters

To apply the converter see sample below:

    [FirestoreData(ConverterType = typeof(PlayerIdConverter))]
    public class PlayerId
    {
        public string Id { get; }

        public PlayerId(string id)
        {
            Id = id;
        }
    }

    public class PlayerIdConverter : IFirestoreConverter<PlayerId>
    {
        // A PlayerId should be represented in storage as just the string value.
        public object ToFirestore(PlayerId value) => value.Id;

        public PlayerId FromFirestore(object value)
        {
            switch (value)
            {
                // This is the expected path: 
                case string id:
                    return new PlayerId(id);
                // The converter will never be called with a null value from Google.Cloud.Firestore,
                // but throw an appropriate exception anyway.
                case null:
                    throw new ArgumentNullException(nameof(value));
                // The converter may be called with unexpected data if (say) there's a document stored
                // with a field of the wrong type.
                default:
                    throw new ArgumentException($"Invalid type to convert to PlayerId: {value.GetType()}");
            }
        }
    }

Then I guess the `PlayerId` class is used as a property type in the POCO. You can also selectively apply it, like so:

    [FirestoreData]
    public class PlayerWithGuidId
    {
        [FirestoreProperty(ConverterType = typeof(GuidConverter))]
        public Guid Id { get; set; }

        [FirestoreProperty]
        public DateTime LastPlayed { get; set; }

        [FirestoreProperty]
        public int HighScore { get; set; }
    }

    public class GuidConverter : IFirestoreConverter<Guid>
    {
        public object ToFirestore(Guid value) => value.ToString("N");

        public Guid FromFirestore(object value)
        {
            switch (value)
            {
                case string guid: return Guid.ParseExact(guid, "N");
                case null: throw new ArgumentNullException(nameof(value));
                default: throw new ArgumentException($"Unexpected data: {value.GetType()}");
            }
        }
    }

Or you can register the converter with the `FirestoreDb` instance via the `FirestoreDbBuilder.ConverterRegistry` property.

    FirestoreDb db = new FirestoreDbBuilder
    {
        ProjectId = projectId,
        ConverterRegistry = new ConverterRegistry
        {
            new GuidConverter()
        }
    }.Build();
    // Documents serialized and deserialized via db will now use the
    // custom Guid converter.

With a **document converter** you can even convert between a custom type and a Firestore map value. This requires that the `ToFirestore` method return a dictionary or anonymous object, and the `FromFirestore` method should expect an `IDictionary<string, object>` and marshall the values back into the custom type.

    [FirestoreData(ConverterType = typeof(CustomCityConverter))]
    public class CustomCity
    {
        // Note: no [FirestoreProperty] attributes. The converter
        // is doing all the work.
        public string Name { get; }
        public string Country { get; }
        public long Population { get; }

        public CustomCity(string name, string country, long population)
        {
            Name = name;
            Country = country;
            Population = population;
        }
    }

    public class CustomCityConverter : IFirestoreConverter<CustomCity>
    {
        // An anonymous type is a convenient way of serializing a map value.
        // You could create a Dictionary<string, object> instead; they're stored the
        // same way.
        public object ToFirestore(CustomCity value) =>
            new { value.Name, value.Country, value.Population };

        public CustomCity FromFirestore(object value)
        {
            if (value is IDictionary<string, object> map)
            {
                // Any exceptions thrown here will be propagated directly. You may wish to be more
                // careful, if you need to control the exact exception thrown used when the
                // data is incomplete or has the wrong type.
                return new CustomCity((string) map["Name"], (string) map["Country"], (long) map["Population"]);
            }
            throw new ArgumentException($"Unexpected data: {value.GetType()}");
        }
    }

