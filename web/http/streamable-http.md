# Streamable HTTP (for Model Context Protocol - MCP)

Streamable HTTP, as defined and used by the **Model Context Protocol (MCP)**, is a modern transport mechanism designed for efficient and resilient communication between AI agents and external tools/resources. It builds upon standard HTTP but with key innovations for streaming data.

---

## Core Concept: "Upgrade in Place"

Unlike traditional HTTP where a response is finite, Streamable HTTP allows a standard HTTP connection to **dynamically "upgrade" in place** to a **Server-Sent Events (SSE)** stream. This means:

* A **client (AI agent)** makes a regular HTTP request (often a `POST`).
* The **server (MCP server)**, upon receiving the request, can choose to:
    * Send a finite, one-off HTTP response.
    * **Or, if the interaction requires continuous updates (e.g., streaming LLM output), it sends specific HTTP headers (`Content-Type: text/event-stream`, `Connection: keep-alive`) and then continuously streams data over that *same* HTTP response.** The server code for that request remains active, continuously writing data to the response stream and flushing it to the client.

This approach provides flexibility: simple interactions are handled with regular HTTP, while complex, long-running agent tasks benefit from seamless streaming.

---

## The Underlying TCP Connection & Long-Running Handlers

Regardless of the specific HTTP layer (HTTP/1.1, HTTP/2), the fundamental transport for both Streamable HTTP (and SSE) and WebSockets is a **single, underlying, persistent TCP connection** (Layer 4).

* **TCP Provides the Reliable Pipe:** TCP ensures a reliable, ordered, error-checked, full-duplex byte stream. Once established via a 3-way handshake, this TCP connection acts as the "pipe" for all subsequent application-layer data.
* **The HTTP Layer Defines Behavior:** The HTTP layer (or WebSocket protocol layer) defines *how* data is structured and exchanged *over* that TCP pipe.
* **Long-Running Handlers are Expected:** For streaming responses like SSE, the server-side code handling that specific request is **designed to remain active and execute for an extended period**. It enters a loop, continuously generating and writing data to the response stream, flushing it to the client. This differs from traditional HTTP handlers that complete and exit quickly. This long-running nature is crucial for pushing real-time updates and is well-supported by modern web servers like .NET's Kestrel, as long as `FlushAsync()` is is called and cancellation tokens are respected for graceful disconnections.

---

## Why Streamable HTTP (and SSE) for MCP?

* **Unified Endpoint:** Simplifies architecture by using a single HTTP endpoint for both requests and responses, including streams.
* **Efficiency:** Avoids the overhead of constantly opening/closing TCP connections for every piece of streamed data.
* **Resilience:** Features like `Mcp-Session-Id` headers allow for better session management and recovery from disconnections compared to raw SSE.
* **AI Agent Focus:** Optimized for the primary streaming need in AI agents: the server (model) pushing continuous output to the client.
* **HTTP Compatibility:** Leverages existing HTTP infrastructure, which is generally firewall-friendly and well-understood.

---

## Streamable HTTP vs. WebSockets

| Feature | Streamable HTTP (MCP's Flavor using SSE) | WebSockets |
| :--------------------- | :----------------------------------------------------------------- | :------------------------------------------------- |
| **Communication Flow** | Primarily **unidirectional** (server to client) streaming. Client-to-server data uses separate HTTP requests. | **Bidirectional (full-duplex)** for real-time interaction. |
| **Protocol Basis** | Built **on top of HTTP/1.1 or HTTP/2**, interpreting a long-lived HTTP response. | An **independent protocol** that *upgrades* from an HTTP handshake. |
| **Overhead** | Lower initial handshake overhead than WebSockets. | Higher initial handshake overhead, then lower per-message overhead. |
| **Use Cases** | Live updates, notifications, streaming AI model outputs, scenarios where server pushes dominate. | Chat, gaming, collaborative editing, high-frequency interactive applications. |
| **Connection Limits** | Subject to HTTP/1.1 browser connection limits per domain (typically 6-8) for client-initiated requests. HTTP/2 avoids this by multiplexing. | Single TCP connection can handle multiple logical "channels" (streams). |