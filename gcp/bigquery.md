# BigQuery

## What is BigQuery?

- Fully managed data warehouse.
- Does not support foreign keys.
- Uses SQL ANSI 2011 queries.
- Create SQL-like cached views.
- DML, DDL and UDF (user-defined functions).
- Query terabytes in seconds.
- Compute separate from storage, scale independently.
- Has a CLI and web UI.
- Has programming libs.
- Has REST and RPC APIs and even JDBC, ODBC.
- Free usage tier.
- It's a column store called "Capacitor".
- Presents as tables, rows and columns.
- ACID.
- Auto replication for HA.
- Stream load and batch load via Avro, Parquet, ORC, CSV, JSON, Datastore and Firestore formats.
- BigQuery Data Transfer Service automates ingestion.
- Can use BigQuery federated query engine to query against Cloud Storage, Bigtable, Spanner or even Google Sheets in Google Drive, and "external tables".
- Connect with BI Engine, Data Studio, Looker, Google Sheets, Tableau, Power BI.
- BigQuery ML.
- Jobs let you run actions to load, export, query or copy data.

##### More information

[The Definitive Guide](https://www.google.com/books/edition/Google_BigQuery_The_Definitive_Guide/-Jq4DwAAQBAJ)

[Inside Capacitor, BigQuery's next-generation columnar storage format](https://cloud.google.com/blog/products/bigquery/inside-capacitor-bigquerys-next-generation-columnar-storage-format)

## Overview of BigQuery Storage

### Table data

- Uses a petabit network between compute nodes and storage and runs queries in memory.
- 99.999999999% durable/year and encrypted.
- Standard tables contain structured data, with a column datatype schema.
- Tables clones are light writable copies of standards, storing just delta.
- Table snapshots are readonly delta copies, can be used to restore.
- Materialized views are precomputed query results stored in BQ storage.
- External tables have a schema but the data is in e.g. Cloud Storage, they're free.
- Tables etc. are organized into logical containers called _datasets_.

### Metadata

- Free storage of schemas and partitioning information etc. etc.

### Storage layout

- Row-oriented databases are good for looking up one or more rows.
- Columnar databases excel at reading rapidly over select fields, like price.
- Columns also have repeating values and can be compressed, read optimised.
- BigQuery does not support foreign keys, so it's only good for OLAP and DW.
- Supports _nested fields_ (complex object) which have the **STRUCT** data type.
- Supports _repeated fields_ (list) which have the **ARRAY** data type.

Using repeated, nested fields is a great way to denormalise data. E.g. instead of having a table of invoice line items, these can be stored in a BQ table in a repeated, nested field of product, quantity and cost. I'm not sure whether this means in many repeated fields, or in a single, repeated, nested field (combined).

[Working with joins, nexted and repeated data](https://cloud.google.com/blog/topics/developers-practitioners/bigquery-explained-working-joins-nested-repeated-data)

### Partitioning
