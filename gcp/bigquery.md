# BigQuery

## What is BigQuery?

- Fully managed data warehouse.
- Does not support foreign keys, but does support joins.
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
- Stream load and batch load via Avro, Parquet, ORC, CSV, JSON Lines, Datastore and Firestore formats.
- BigQuery Data Transfer Service automates ingestion.
- Can use BigQuery federated query engine to query against Cloud Storage, Bigtable, Spanner or even Google Sheets in Google Drive, and "external tables".
- Connect with BI Engine, Data Studio, Looker, Google Sheets, Tableau, Power BI.
- BigQuery ML.
- Jobs let you run actions to load, export, query or copy data.

## Pricing

- First 1TB/month is free! Then it's $5/TB/month.

- Considering the minimum 10MB query size, that would be 100,000 free queries a month.

- Errors are free and so are cached responses.

- Charges are rounded up to the nearest MB, with a minimum 10 MB data processed per table referenced by the query, and with a minimum 10 MB data processed per query.

- Partitioning and clustering can help reduce cost, pruning and reducing scanning.

- Querying external data is charged by bytes read, plus any costs for the host service.

- Flat-rate bulk pricing is done by buying reserved _slots_. This dedicated capacity will be queued if exceeded, so you never get charged extra.

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

- Divide a table into smaller parts based on column values.
- The key can be a time-unit column, an integer or the ingestion time.
- Partitions are divided into buckets; for time-unit or ingestion-time keys, the bucket size can be hourly, daily, monthly or yearly.
- For integers, you specify the range.
- This can speed up queries as they collate matching buckets to load.
- Works best with a few thousand or fewer distinct values of the key.
- Can set partition expiration time and have these buckets auto deleted.

### Clustering

- Colocates similar values in a table based on one or more columns.
- Organized into blocks in BQ storage.
- Inserts cause a re-org.
- Improves query performance when queries have a filter by not reading blocks outside the range.
- Can also improve aggregations based on clustered column values and joins over star schemas.
- Combine with partitioning.

Clustering works best when you have a large number of distinct values in one or more columns and your queries filter or aggregate over a small range of them.

Example shows partitioning by transaction date and clustering on customer ID.

### Ingesting data into BigQuery storage

- Batch load manually or on auto schedule into new or existing table append.
- Stream smaller batches so that data is available in almost real-time.
- Use SQL to insert and append or output results to a table.

Can load from Firestore exports made using [Firestore managed export and import service](https://cloud.google.com/firestore/docs/manage-data/export-import).

### Reading data from BigQuery storage

- `tabledata.list` REST API pages through data but it's single concurrent reader and low throughput.
- Run an [extract job](https://cloud.google.com/bigquery/docs/exporting-data) to save to Cloud Storage in CSV, JSON or Avro form, 1GB per file over multiple files.
- Use the [BigQuery Storage Read API](https://cloud.google.com/bigquery/docs/reference/storage) for parallel reads using its own binary format at $1.1/TB.

### Introduction to tables

- Can define schema when creating table or leave it schemaless and define it at query time, or just before first data is loaded during batch load.

### Table limitations

- Names unique within dataset.
- Cloud console can copy only one table at a time.
- Copy destination dataset must be in same location (EU>US nope).
- When copying >1 source tables to a destination, all must share same schema.
- Export only support Cloud Storage.
- In an API call, enumeration slows towards 50,000 tables in a dataset.
- Cloud console limited to showing 50,000 tables.

## Overview of BigQuery analytics

### Types of analysis

- Ad hoc via SQL queries in the console or 3rd-party tools.
- Geospatial analysis.
- Machine learning via BigQuery ML executed using SQL.
- Business intelligence via BigQuery BI Engine to build interactive dashboards.

### Queries

- Has Standard SQL and legacy SQL.
- Standard SQL:2011 and extensions for geo and ML.

#### Data sources

- Native data is obvious data in BQ storage, loaded or generated by DML statements, or query results to new table.
- External data comes from Cloud Storage, Spanner or Cloud SQL.
- Multi-cloud data comes from AWS or Azure; see BigQuery Omni documentation.
- Public datasets are from the marketplace.

### Types of queries

- Interactive, immediate, via Cloud console, `bq query` CLI, REST API or client libraries.
- Batch, queued and started as soon as idle.

#### Query jobs

- Actions to load, export, query or copy data.
- Cloud console or `bq` CLI creates, schedules and runs jobs.
- Can be done programmatically.
- Poll for status.

#### Saving and sharing queries

- You can save queries and share them, see: https://cloud.google.com/bigquery/docs/saving-sharing-queries

#### Query processing

- When a query runs, an execution tree is made, like a SQL Server query plan.
- Stages contain steps to run in parallel.
- Stages communicate via a fast distributed shuffle tier for storing intermediate data.
- Uses petabit network and RAM where possible.
- **Execution tree** is the plan for parallel stages.
- **Shuffle tier** stores intermediate data.
- **Query plan** is viewable in the console and is made after the tree.
- **Query monitoring and dynamic planning** is a system of workers that monitor and sometimes alter the plan as it runs.
- When the query is complete the results are written to persistent storage and returned to the user so it can serve cached results next time.

#### Query optimization

- View the query plan via console, or `INFORMATION_SCHEMA` view or via Jobs API.
- See more here: https://cloud.google.com/bigquery/query-plan-explanation
- And here: https://cloud.google.com/bigquery/docs/best-practices-performance-overview
- And more about monitoring, here: https://cloud.google.com/bigquery/docs/monitoring
- And audit logs: https://cloud.google.com/bigquery/docs/reference/auditlogs

### Data analytics features

- Supports both _descriptive_ and _predictive_ analytics.

#### Analytics tools integration

- Google Data Studio launched right after a query result in the console will connect and let you explore.
- Connected Sheets and also be launched from the console. This runs queries on request or on schedule and saved in the sheet.
- Looker is a BI platform.


