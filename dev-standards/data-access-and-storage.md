# Data access and storage

From `Tz/src/.cursor/rules/data-access.mdc`, `Tz/src/AGENTS.md`, `tests/README.md` (in-memory patterns).

---

## Storage

- **Primary:** GCP — Firestore (documents), Google Cloud Storage (blobs).
- **Legacy:** Azure Storage (phasing out).
- **Tests:** in-memory repos when configured (see DI toggle in test/app config).

---

## Repository pattern

- Base types: `FirestoreApiRepository<TAppDocPoco, TFirestoreDocPoco>`, `GoogleIndexedObjectStore<TObjectDto>`.
- Factories: `GoogleStorageClientFactory`, `FirestoreClientFactory`.
- One repository focus per entity; register repos at DI **seam**.
- async/await for all I/O; strong typing; storage model separate from domain model.

---

## Document models

- Firestore: `[FirestoreData]`, `[FirestoreProperty]`.
- Implement `IDocInfo`, `IConnectionDoc` where applicable.
- Support **data shredding** for privacy.
- Validators: `*Validator` inheriting `DocValidator<T>`, override `HasValidProperties` (required fields, UTC dates, business rules, relationships); pass validator into repository constructor.

---

## OpContext enforcement

- **Restricted `op`:** caller may only access own account key (`OpContext.CanActOn`). Cross-account access throws `UnauthorizedDataAccessException` (in-memory and Firestore repos).
- **Derestrict** only at controllers/event handlers **after** auth — not in Commands/Queries.
- Workplace admin: restricted validate → resolve workplace → `dOp` for scoped work.
- Audit sensitive database operations.

**Debugging empty/wrong query results:**

1. Is data actually persisted? (`FindByKeyAsync`)
2. Is `op` still restricted? (try `op.Derestricted()` at the right layer only)
3. Account/doc key format correct? (e.g. `accn` prefix, two-part keys — see [testing.md](./testing.md))

Full guide: `Tz/docs/architecture/OP_CONTEXT_DATA_PROTECTION.md`.

---

## In-memory tests (Firestore POCOs)

**Always create `InMemoryDocRef` before document POCOs:**

```csharp
var docKey = "00_accn_test123";
var path = $"accounts/{docKey}";
var docRef = new InMemoryDocRef(docKey, path);
var account = new AccountDoc(docRef) { AccountKey = docKey, /* ... */ };
await repository.InsertAsync(op, account);
```

Without `DocRef`, `GetDocKey()` / inserts fail (“Firestore document reference is null”). See `SeedDataLoader.cs`, `tests/README.md`.

Seed JSON uses explicit `*Key` fields (e.g. `activityKey: "00_jnla_001"`), not legacy Azure partitionKey/rowKey.

---

## General data-access practices

- Efficient indexing and LINQ; compiled queries for hot paths.
- Concurrency, retries for transient failures.
- Parameterized queries where SQL exists.
- Encrypt sensitive fields; access control at repository/op layer.
