# DID - Decentralized Identifiers

https://www.w3.org/TR/did-core/

- An identifier based on PKI with a public-private key pair, not necessarily issued by
  any particular authority.

- Concept of the controller, which is the entity who can make changes to the DID Document, below.

- Concept of authentication which is proving control via PKI.

- URL format: did:[method]:[id_string]

- Method is like a sub-scheme which the public can make up and publish its spec.

- Backed by a DID Document which must be on some form of public ledger (I'm thinking EAS)

```json
{
  "@context": [
    "https://www.w3.org/ns/did/v1",
    "https://w3id.org/security/suites/ed25519-2020/v1"
  ]
  "id": "did:example:123456789abcdefghi",
  "authentication": [{
    
    "id": "did:example:123456789abcdefghi#keys-1",
    "type": "Ed25519VerificationKey2020",
    "controller": "did:example:123456789abcdefghi",
    "publicKeyMultibase": "zH3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
  }]
}
```

- Supports further URL syntax like path, query params.

- Authentication seems limited to `publicKeyJwk` and `publicKeyMultibase` where multibase is not yet a 
  standard.

- Can store the document using different representations so long as they have a known content type and conform to DID RFC, though JSON seems obvious, using `application/did+json`

- See examples, here: https://www.w3.org/TR/did-core/#examples