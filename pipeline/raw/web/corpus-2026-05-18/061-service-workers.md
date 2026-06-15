---
title: Service Workers
kind: source
created: 2026-05-18
updated: 2026-05-18
sources:
  - raw/web/curated-web-corpus-2026-05-18.md
  - https://www.w3.org/TR/service-workers/
tags: [web-standards, web-corpus]
confidence: medium
---

# Service Workers

## Fetch Metadata

- Corpus item: 61
- Category: Web Standards
- Source URL: https://www.w3.org/TR/service-workers/
- Final URL: https://www.w3.org/TR/service-workers/
- Retrieved: 2026-05-18
- Content-Type: text/html; charset=utf-8
- Fetch status: ok
- Test value: Tests offline and request interception concepts.
- Fetched page title: Service Workers Nightly

## Retrieved Text

W3C

Service Workers Nightly

W3C Candidate Recommendation Draft, 9 April 2026
More details about this document

This version:
https://www.w3.org/TR/2026/CRD-service-workers-20260409/

Latest published version:
https://www.w3.org/TR/service-workers/

Editor's Draft:
https://w3c.github.io/ServiceWorker/

Previous Versions:
https://www.w3.org/TR/2026/CRD-service-workers-20260312/

History:
https://www.w3.org/standards/history/service-workers/

Implementation Report:
https://wpt.fyi/results/service-workers

Test Suite:
https://github.com/web-platform-tests/wpt/tree/master/service-wo
rkers

Feedback:
GitHub
Inline In Spec

Editors:
Monica Chintala (Microsoft)
Yoshisato Yanagisawa (Google)

Former Editors:
Alex Russell (Google)
Jake Archibald (Google)
Jungkee Song (Microsoft‚ represented Samsung until April 2018)
Marijn Kruisselbrink (Google)

Copyright © 2026 World Wide Web Consortium. W3C^® liability, trademark
and permissive document license rules apply.
__________________________________________________________________

Abstract

The core of this specification is a worker that wakes to receive
events. This provides an event destination that can be used when other
destinations would be inappropriate, or no other destination exists.

For example, to allow the developer to decide how a page should be
fetched, an event needs to dispatch potentially before any other
execution contexts exist for that origin. To react to a push message,
or completion of a persistent download, the context that originally
registered interest may no longer exist. In these cases, the service
worker is the ideal event destination.

This specification also provides a fetch event, and a request and
response store similar in design to the HTTP cache, which makes it
easier to build offline-enabled web applications.

Status of this document

This section describes the status of this document at the time of its
publication. A list of current W3C publications and the latest revision
of this technical report can be found in the W3C standards and drafts
index.

This document was published by the Web Applications Working Group as a
Candidate Recommendation Draft using the Recommendation track.
Publication as a Candidate Recommendation does not imply endorsement by
W3C and its Members. A Candidate Recommendation Draft integrates
changes from the previous Candidate Recommendation that the Working
Group intends to include in a subsequent Candidate Recommendation
Snapshot.

This is a draft document and may be updated, replaced or obsoleted by
other documents at any time. It is inappropriate to cite this document
as other than a work in progress.

This is a living document. Readers need to be aware that this
specification may include unimplemented features, and details that may
change. Service Workers 1 is a version that is advancing toward a W3C
Recommendation.

This document was produced by a group operating under the W3C Patent
Policy. W3C maintains a public list of any patent disclosures made in
connection with the deliverables of the group; that page also includes
instructions for disclosing a patent. An individual who has actual
knowledge of a patent that the individual believes contains Essential
Claim(s) must disclose the information in accordance with section 6 of
the W3C Patent Policy.

This document is governed by the 18 August 2025 W3C Process Document.

Table of Contents

1. 1 Motivations
2. 2 Model
1. 2.1 Service Worker
1. 2.1.1 Lifetime
2. 2.1.2 Events
2. 2.2 Service Worker Timing
3. 2.3 Service Worker Registration
1. 2.3.1 Lifetime
4. 2.4 Service Worker Client
5. 2.5 Control and Use
1. 2.5.1 The window client case
2. 2.5.2 The worker client case
6. 2.6 Task Sources
7. 2.7 User Agent Shutdown
3. 3 Client Context
1. 3.1 ServiceWorker
1. 3.1.1 Getting ServiceWorker instances
2. 3.1.2 scriptURL
3. 3.1.3 state
4. 3.1.4 postMessage(message, transfer)
5. 3.1.5 postMessage(message, options)
6. 3.1.6 Event handler
2. 3.2 ServiceWorkerRegistration
1. 3.2.1 Getting ServiceWorkerRegistration instances
2. 3.2.2 installing
3. 3.2.3 waiting
4. 3.2.4 active
5. 3.2.5 navigationPreload
6. 3.2.6 scope
7. 3.2.7 updateViaCache
8. 3.2.8 update()
9. 3.2.9 unregister()
10. 3.2.10 Event handler
3. 3.3 navigator.serviceWorker
4. 3.4 ServiceWorkerContainer
1. 3.4.1 controller
2. 3.4.2 ready
3. 3.4.3 register(scriptURL, options)
4. 3.4.4 getRegistration(clientURL)
5. 3.4.5 getRegistrations()
6. 3.4.6 startMessages()
7. 3.4.7 Event handlers
5. 3.5 Events
6. 3.6 NavigationPreloadManager
1. 3.6.1 enable()
2. 3.6.2 disable()
3. 3.6.3 setHeaderValue(value)
4. 3.6.4 getState()
4. 4 Execution Context
1. 4.1 ServiceWorkerGlobalScope
1. 4.1.1 clients
2. 4.1.2 registration
3. 4.1.3 serviceWorker
4. 4.1.4 skipWaiting()
5. 4.1.5 Event handlers
2. 4.2 Client
1. 4.2.1 url
2. 4.2.2 frameType
3. 4.2.3 id
4. 4.2.4 type
5. 4.2.5 postMessage(message, transfer)
6. 4.2.6 postMessage(message, options)
7. 4.2.7 visibilityState
8. 4.2.8 focused
9. 4.2.9 ancestorOrigins
10. 4.2.10 focus()
11. 4.2.11 navigate(url)
3. 4.3 Clients
1. 4.3.1 get(id)
2. 4.3.2 matchAll(options)
3. 4.3.3 openWindow(url)
4. 4.3.4 claim()
4. 4.4 ExtendableEvent
1. 4.4.1 event.waitUntil(f)
5. 4.5 InstallEvent
1. 4.5.1 event.addRoutes(rules)
6. 4.6 FetchEvent
1. 4.6.1 event.request
2. 4.6.2 event.preloadResponse
3. 4.6.3 event.clientId
4. 4.6.4 event.resultingClientId
5. 4.6.5 event.replacesClientId
6. 4.6.6 event.handled
7. 4.6.7 event.respondWith(r)
7. 4.7 ExtendableMessageEvent
1. 4.7.1 event.data
2. 4.7.2 event.origin
3. 4.7.3 event.lastEventId
4. 4.7.4 event.source
5. 4.7.5 event.ports
8. 4.8 Events
5. 5 Caches
1. 5.1 Constructs
2. 5.2 Understanding Cache Lifetimes
3. 5.3 self.caches
1. 5.3.1 caches
4. 5.4 Cache
1. 5.4.1 match(request, options)
2. 5.4.2 matchAll(request, options)
3. 5.4.3 add(request)
4. 5.4.4 addAll(requests)
5. 5.4.5 put(request, response)
6. 5.4.6 delete(request, options)
7. 5.4.7 keys(request, options)
5. 5.5 CacheStorage
1. 5.5.1 match(request, options)
2. 5.5.2 has(cacheName)
3. 5.5.3 open(cacheName)
4. 5.5.4 delete(cacheName)
5. 5.5.5 keys()
6. 6 Security Considerations
1. 6.1 Secure Context
2. 6.2 Content Security Policy
3. 6.3 Origin Relativity
1. 6.3.1 Origin restriction
2. 6.3.2 importScripts(urls)
4. 6.4 Cross-Origin Resources and CORS
5. 6.5 Path restriction
6. 6.6 Service worker script request
7. 6.7 Implementer Concerns
8. 6.8 Privacy
7. 7 Extensibility
1. 7.1 Define API bound to Service Worker Registration
2. 7.2 Define Functional Event
3. 7.3 Define Event Handler
4. 7.4 Firing Functional Events
8. Appendix A: Algorithms
1. Create Job
2. Schedule Job
3. Run Job
4. Finish Job
5. Resolve Job Promise
6. Reject Job Promise
7. Start Register
8. Register
9. Update
10. Soft Update
11. Install
12. Activate
13. Try Activate
14. Setup ServiceWorkerGlobalScope
15. Run Service Worker
16. All Fetch Listeners Are Empty
17. Terminate Service Worker
18. Handle Fetch
19. Create Fetch Event and Dispatch
20. Parse URL Pattern
21. Verify Router Condition
22. Match Router Condition
23. Check Router Registration Limit
24. Count Router Inner Conditions
25. Get Router Source
26. Should Skip Event
27. Fire Functional Event
28. Handle Service Worker Client Unload
29. Handle User Agent Shutdown
30. Update Service Worker Extended Events Set
31. Unregister
32. Set Registration
33. Clear Registration
34. Try Clear Registration
35. Update Registration State
36. Update Worker State
37. Notify Controller Change
38. Match Service Worker Registration
39. Get Registration
40. Get Newest Worker
41. Service Worker Has No Pending Events
42. Create Client
43. Create Window Client
44. Get Frame Type
45. Resolve Get Client Promise
46. Query Cache
47. Request Matches Cached Item
48. Batch Cache Operations
49. Is Async Module
50. Lookup Race Response
9. Appendix B: Extended HTTP headers
1. Service Worker Script Request
2. Service Worker Script Response
3. Syntax
10. 8 Acknowledgements
11. Conformance
1. Document conventions
2. Conformant Algorithms
12. Index
1. Terms defined by this specification
2. Terms defined by reference
13. References
1. Normative References
2. Informative References
14. IDL Index
15. Issues Index

1. Motivations

This section is non-normative.

Web Applications traditionally assume that the network is reachable.
This assumption pervades the platform. HTML documents are loaded over
HTTP and traditionally fetch all of their sub-resources via subsequent
HTTP requests. This places web content at a disadvantage versus other
technology stacks.

The service worker is designed first to redress this balance by
providing a Web Worker context, which can be started by a runtime when
navigations are about to occur. This event-driven worker is registered
against an origin and a path (or pattern), meaning it can be consulted
when navigations occur to that location. Events that correspond to
network requests are dispatched to the worker and the responses
generated by the worker may override default network stack behavior.
This puts the service worker, conceptually, between the network and a
document renderer, allowing the service worker to provide content for
documents, even while offline.

Web developers familiar with previous attempts to solve the offline
problem have reported a deficit of flexibility in those solutions. As a
result, the service worker is highly procedural, providing a maximum of
flexibility at the price of additional complexity for developers. Part
of this complexity arises from the need to keep service workers
responsive in the face of a single-threaded execution model. As a
result, APIs exposed by service workers are almost entirely
asynchronous, a pattern familiar in other JavaScript contexts but
accentuated here by the need to avoid blocking document and resource
loading.

Developers using the HTML5 Application Cache have also reported that
several attributes of the design contribute to unrecoverable errors. A
key design principle of the service worker is that errors should always
be recoverable. Many details of the update process of service workers
are designed to avoid these hazards.

Service workers are started and kept alive by their relationship to
events, not documents. This design borrows heavily from developer and
vendor experience with shared workers and Chrome Background Pages. A
key lesson from these systems is the necessity to time-limit the
execution of background processing contexts, both to conserve resources
and to ensure that background context loss and restart is top-of-mind
for developers. As a result, service workers bear more than a passing
resemblance to Chrome Event Pages, the successor to Background Pages.
Service workers may be started by user agents without an attached
document and may be killed by the user agent at nearly any time.
Conceptually, service workers can be thought of as Shared Workers that
can start, process events, and die without ever handling messages from
documents. Developers are advised to keep in mind that service workers
may be started and killed many times a second.

Service workers are generic, event-driven, time-limited script contexts
that run at an origin. These properties make them natural endpoints for
a range of runtime services that may outlive the context of a
particular document, e.g. handling push notifications, background data
synchronization, responding to resource requests from other origins, or
receiving centralized updates to expensive-to-calculate data (e.g.,
geolocation or gyroscope).

2. Model

2.1. Service Worker

A service worker is a type of web worker. A service worker executes in
the registering service worker client’s origin.

A service worker has an associated state, which is one of "parsed",
"installing", "installed", "activating", "activated", and "redundant".
It is initially "parsed".

A service worker has an associated script url (a URL).

A service worker has an associated type which is either "classic" or
"module". Unless stated otherwise, it is "classic".

A service worker has an associated containing service worker
registration (a service worker registration), which contains itself.

A service worker has an associated global object (a
ServiceWorkerGlobalScope object or null).

A service worker has an associated script resource (a script), which
represents its own script resource. It is initially set to null.

A script resource has an associated has ever been evaluated flag. It is
initially unset.

A script resource has an associated policy container (a policy
container). It is initially a new policy container.

A service worker has an associated script resource map which is an
ordered map where the keys are URLs and the values are responses.

A service worker has an associated set of used scripts (a set) whose
item is a URL. It is initially a new set.

Note: The set of used scripts is only used to prune unused resources
from a new worker’s map after installation, that were populated based
on the old worker’s map during the update check.

A service worker has an associated skip waiting flag. Unless stated
otherwise it is unset.

A service worker has an associated classic scripts imported flag. It is
initially unset.

A service worker has an associated set of event types to handle (a set)
whose item is an event listener’s event type. It is initially a new
set.

A service worker has an associated set of extended events (a set) whose
item is an ExtendableEvent. It is initially a new set.

A service worker has an associated start status which can be null or a
Completion. It is initially null.

A service worker has an associated all fetch listeners are empty flag.
It is initially unset.

A service worker has an associated list of router rules (a list of
RouterRules). It is initially an empty list.

A service worker is said to be running if its event loop is running.

A service worker has an associated [[service worker queue]] (a parallel
queue).

2.1.1. Lifetime

The lifetime of a service worker is tied to the execution lifetime of
events and not references held by service worker clients to the
ServiceWorker object.

A user agent may terminate service workers at any time it:
* Has no event to handle.
* Detects abnormal operation: such as infinite loops and tasks
exceeding imposed time limits (if any) while handling the events.

2.1.2. Events

The Service Workers specification defines service worker events (each
of which is an event) that include (see the list):
* Lifecycle events: install and activate.
* Functional events: fetch and the events defined by other
specifications that extend the Service Workers specification. (See
the list.)
* message and messageerror.

2.2. Service Worker Timing

Service workers mark certain points in time that are later exposed by
the navigation timing API and resource timing API.

A service worker timing info is a struct. It has the following items:

start time
A DOMHighResTimeStamp, initially 0.

fetch event dispatch time
A DOMHighResTimeStamp, initially 0.

worker router evaluation start
A DOMHighResTimeStamp, initially 0.

worker cache lookup start
A DOMHighResTimeStamp, initially 0.

worker matched router source
A DOMString, initially an empty string.

worker final router source
A DOMString, initially an empty string.

2.3. Service Worker Registration

A service worker registration is a tuple of a scope url, a storage key,
and a set of service workers, an installing worker, a waiting worker,
and an active worker. A user agent may enable many service worker
registrations at a single origin so long as the scope url of the
service worker registration differs. A service worker registration of
an identical scope url when one already exists in the user agent causes
the existing service worker registration to be replaced.

A service worker registration has an associated storage key (a storage
key).

A service worker registration has an associated scope url (a URL).

A service worker registration has an associated installing worker (a
service worker or null) whose state is "installing". It is initially
set to null.

A service worker registration has an associated waiting worker (a
service worker or null) whose state is "installed". It is initially set
to null.

A service worker registration has an associated active worker (a
service worker or null) whose state is either "activating" or
"activated". It is initially set to null.

A service worker registration has an associated last update check time.
It is initially set to null.

A service worker registration is said to be stale if the registration’s
last update check time is non-null and the time difference in seconds
calculated by the current time minus the registration’s last update
check time is greater than 86400.

A service worker registration has an associated update via cache mode,
which is "imports", "all", or "none". It is initially set to "imports".

A service worker registration has one or more task queues that back up
the tasks from its active worker’s event loop’s corresponding task
queues. (The target task sources for this back up operation are the
handle fetch task source and the handle functional event task source.)
The user agent dumps the active worker’s tasks to the service worker
registration’s task queues when the active worker is terminated and
re-queues those tasks to the active worker’s event loop’s corresponding
task queues when the active worker spins off. Unlike the task queues
owned by event loops, the service worker registration’s task queues are
not processed by any event loops in and of itself.

A service worker registration has an associated
NavigationPreloadManager object.

A service worker registration has an associated navigation preload
enabled flag. It is initially unset.

A service worker registration has an associated navigation preload
header value, which is a byte sequence. It is initially set to `true`.

A service worker registration is said to be unregistered if
registration map[this service worker registration’s (storage key,
serialized scope url)] is not this service worker registration.

2.3.1. Lifetime

A user agent must persistently keep a list of registered service worker
registrations unless otherwise they are explicitly unregistered. A user
agent has a registration map that stores the entries of the tuple of
service worker registration’s (storage key, serialized scope url) and
the corresponding service worker registration. The lifetime of service
worker registrations is beyond that of the ServiceWorkerRegistration
objects which represent them within the lifetime of their corresponding
service worker clients.

2.4. Service Worker Client

A service worker client is an environment.

A service worker client has an associated discarded flag. It is
initially unset.

Each service worker client has the following environment discarding
steps:
1. Set client’s discarded flag.

Note: Implementations can discard clients whose discarded flag is set.

A service worker client has an algorithm defined as the origin that
returns the service worker client’s origin if the service worker client
is an environment settings object, and the service worker client’s
creation URL’s origin otherwise.

A window client is a service worker client whose global object is a
Window object.

A dedicated worker client is a service worker client whose global
object is a DedicatedWorkerGlobalScope object.

A shared worker client is a service worker client whose global object
is a SharedWorkerGlobalScope object.

A worker client is either a dedicated worker client or a shared worker
client.

2.5. Control and Use

A service worker client has an active service worker that serves its
own loading and its subresources. When a service worker client has a
non-null active service worker, it is said to be controlled by that
active service worker. When a service worker client is controlled by a
service worker, it is said that the service worker client is using the
service worker’s containing service worker registration. A service
worker client’s active service worker is determined as explained in the
following subsections.

The rest of the section is non-normative.

The behavior in this section is not fully specified yet and will be
specified in HTML Standard. The work is tracked by the issue and the
pull request.

2.5.1. The window client case

A window client is created when a browsing context is created and when
it navigates.

When a window client is created in the process of a browsing context
creation:

If the browsing context’s initial active document’s origin is an opaque
origin, the window client’s active service worker is set to null.
Otherwise, it is set to the creator document’s service worker client’s
active service worker.

When a window client is created in the process of the browsing
context’s navigation:

If the fetch is routed through HTTP fetch, the window client’s active
service worker is set to the result of the service worker registration
matching. Otherwise, if the created document’s origin is an opaque
origin or not the same as its creator document’s origin, the window
client’s active service worker is set to null. Otherwise, it is set to
the creator document’s service worker client’s active service worker.

Note: For an initial replacement navigation, the initial window client
that was created when the browsing context was created is reused, but
the active service worker is determined by the same behavior as above.

Note: Sandboxed iframes without the sandboxing directives,
allow-same-origin and allow-scripts, result in having the active
service worker value of null as their origin is an opaque origin.

2.5.2. The worker client case

A worker client is created when the user agent runs a worker.

When the worker client is created:

When the fetch is routed through HTTP fetch, the worker client’s active
service worker is set to the result of the service worker registration
matching. Otherwise, if the worker client’s origin is an opaque origin,
or the request’s URL is a blob URL and the worker client’s origin is
not the same as the origin of the last item in the worker client’s
global object’s owner set, the worker client’s active service worker is
set to null. Otherwise, it is set to the active service worker of the
environment settings object of the last item in the worker client’s
global object’s owner set.

Note: Window clients and worker clients with a data: URL result in
having the active service worker value of null as their origin is an
opaque origin. Window clients and worker clients with a blob URL can
inherit the active service worker of their creator document or owner,
but if the request’s origin is not the same as the origin of their
creator document or owner, the active service worker is set to null.

2.6. Task Sources

The following additional task sources are used by service workers.

The handle fetch task source
This task source is used for dispatching fetch events to service
workers.

The handle functional event task source
This task source is used for features that dispatch other
functional events, e.g. push events, to service workers.

Note: A user agent may use a separate task source for each
functional event type in order to avoid a head-of-line blocking
phenomenon for certain functional events.

2.7. User Agent Shutdown

A user agent must maintain the state of its stored service worker
registrations across restarts with the following rules:
* An installing worker does not persist but is discarded. If the
installing worker was the only service worker for the service
worker registration, the service worker registration is discarded.
* A waiting worker promotes to an active worker.

To attain this, the user agent must invoke Handle User Agent Shutdown
when it terminates.

3. Client Context

Bootstrapping with a service worker:
// scope defaults to the path the script sits in
// "/" in this example
navigator.serviceWorker.register("/serviceworker.js").then(registration => {
console.log("success!");
if (registration.installing) {
registration.installing.postMessage("Howdy from your installing page.");
}
}, err => {
console.error("Installing the worker failed!", err);
});

3.1. ServiceWorker

[SecureContext, Exposed=(Window,Worker)]
interface ServiceWorker : EventTarget {
readonly attribute USVString scriptURL;
readonly attribute ServiceWorkerState state;
undefined postMessage(any message, sequence<object> transfer);
undefined postMessage(any message, optional StructuredSerializeOptions options
= {});

// event
attribute EventHandler onstatechange;
};
ServiceWorker includes AbstractWorker;

enum ServiceWorkerState {
"parsed",
"installing",
"installed",
"activating",
"activated",
"redundant"
};

A ServiceWorker object represents a service worker. Each ServiceWorker
object is associated with a service worker. Multiple separate objects
implementing the ServiceWorker interface across documents and workers
can all be associated with the same service worker simultaneously.

A ServiceWorker object has an associated ServiceWorkerState object
which is itself associated with service worker’s state.

3.1.1. Getting ServiceWorker instances

An environment settings object has a service worker object map, a map
where the keys are service workers and the values are ServiceWorker
objects.
To get the service worker object representing serviceWorker (a service
worker) in environment (an environment settings object), run these
steps:
1. Let objectMap be environment’s service worker object map.
2. If objectMap[serviceWorker] does not exist, then:
1. Let serviceWorkerObj be a new ServiceWorker in environment’s
Realm, and associate it with serviceWorker.
2. Set serviceWorkerObj’s state to serviceWorker’s state.
3. Set objectMap[serviceWorker] to serviceWorkerObj.
3. Return objectMap[serviceWorker].

3.1.2. scriptURL

The scriptURL getter steps are to return the service worker’s
serialized script url.

For example, consider a document created by a navigation to
https://example.com/app.html which matches via the following
registration call which has been previously executed:
// Script on the page https://example.com/app.html
navigator.serviceWorker.register("/service_worker.js");

The value of navigator.serviceWorker.controller.scriptURL will be
"https://example.com/service_worker.js".

3.1.3. state

The state attribute must return the value (in ServiceWorkerState
enumeration) to which it was last set.

3.1.4. postMessage(message, transfer)

The postMessage(message, transfer) method steps are:
1. Let options be «[ "transfer" → transfer ]».
2. Invoke postMessage(message, options) with message and options as
the arguments.

3.1.5. postMessage(message, options)

The postMessage(message, options) method steps are:
1. Let serviceWorker be the service worker represented by this.
2. Let incumbentSettings be the incumbent settings object.
3. Let incumbentGlobal be incumbentSettings’s global object.
4. Let serializeWithTransferResult be
StructuredSerializeWithTransfer(message, options["transfer"]).
Rethrow any exceptions.
5. If the result of running the Should Skip Event algorithm with
"message" and serviceWorker is true, then return.
6. Run these substeps in parallel:
1. If the result of running the Run Service Worker algorithm with
serviceWorker is failure, then return.
2. Queue a task on the DOM manipulation task source to run the
following steps:
1. Let source be determined by switching on the type of
incumbentGlobal:

ServiceWorkerGlobalScope
The result of getting the service worker
object that represents incumbentGlobal’s
service worker in the relevant settings object
of serviceWorker’s global object.

Window
a new WindowClient object that represents
incumbentGlobal’s relevant settings object.

Otherwise
a new Client object that represents
incumbentGlobal’s associated worker

2. Let origin be incumbentSettings’s origin.
3. Let destination be the ServiceWorkerGlobalScope object
associated with serviceWorker.
4. Let deserializeRecord be
StructuredDeserializeWithTransfer(serializeWithTransferRe
sult, destination’s Realm).
If this throws an exception, let e be the result of
creating an event named messageerror, using
ExtendableMessageEvent, with its origin initialized to
origin and the source attribute initialized to source.
5. Else:
1. Let messageClone be
deserializeRecord.[[Deserialized]].
2. Let newPorts be a new frozen array consisting of all
MessagePort objects in
deserializeRecord.[[TransferredValues]], if any,
maintaining their relative order.
3. Let e be the result of creating an event named
message, using ExtendableMessageEvent, with its
origin initialized to origin, the source attribute
initialized to source, the data attribute
initialized to messageClone, and the ports attribute
initialized to newPorts.
6. Dispatch e at destination.
7. Invoke Update Service Worker Extended Events Set with
serviceWorker and e.

3.1.6. Event handler

The following is the event handler (and its corresponding event handler
event type) that must be supported, as event handler IDL attributes, by
all objects implementing ServiceWorker interface:

event handler event handler event type
onstatechange statechange

3.2. ServiceWorkerRegistration

[SecureContext, Exposed=(Window,Worker)]
interface ServiceWorkerRegistration : EventTarget {
readonly attribute ServiceWorker? installing;
readonly attribute ServiceWorker? waiting;
readonly attribute ServiceWorker? active;
[SameObject] readonly attribute NavigationPreloadManager navigationPreload;

readonly attribute USVString scope;
readonly attribute ServiceWorkerUpdateViaCache updateViaCache;

[NewObject] Promise<ServiceWorkerRegistration> update();
[NewObject] Promise<boolean> unregister();

// event
attribute EventHandler onupdatefound;
};

enum ServiceWorkerUpdateViaCache {
"imports",
"all",
"none"
};

A ServiceWorkerRegistration has a service worker registration (a
service worker registration).

3.2.1. Getting ServiceWorkerRegistration instances

An environment settings object has a service worker registration object
map, a map where the keys are service worker registrations and the
values are ServiceWorkerRegistration objects.
To get the service worker registration object representing registration
(a service worker registration) in environment (an environment settings
object), run these steps:
1. Let objectMap be environment’s service worker registration object
map.
2. If objectMap[registration] does not exist, then:
1. Let registrationObject be a new ServiceWorkerRegistration in
environment’s Realm.
2. Set registrationObject’s service worker registration to
registration.
3. Set registrationObject’s installing attribute to null.
4. Set registrationObject’s waiting attribute to null.
5. Set registrationObject’s active attribute to null.
6. If registration’s installing worker is not null, then set
registrationObject’s installing attribute to the result of
getting the service worker object that represents
registration’s installing worker in environment.
7. If registration’s waiting worker is not null, then set
registrationObject’s waiting attribute to the result of
getting the service worker object that represents
registration’s waiting worker in environment.
8. If registration’s active worker is not null, then set
registrationObject’s active attribute to the result of getting
the service worker object that represents registration’s
active worker in environment.
9. Set objectMap[registration] to registrationObject.
3. Return objectMap[registration].

3.2.2. installing

installing attribute must return the value to which it was last set.

Note: Within a Realm, there is only one ServiceWorker object per
associated service worker.

3.2.3. waiting

waiting attribute must return the value to which it was last set.

Note: Within a Realm, there is only one ServiceWorker object per
associated service worker.

3.2.4. active

active attribute must return the value to which it was last set.

Note: Within a Realm, there is only one ServiceWorker object per
associated service worker.

3.2.5. navigationPreload

The navigationPreload getter steps are to return the service worker
registration’s NavigationPreloadManager object.

3.2.6. scope

The scope getter steps are to return the service worker registration’s
serialized scope url.

In the example in § 3.1.2 scriptURL, the value of registration.scope,
obtained from navigator.serviceWorker.ready.then(registration =>
console.log(registration.scope)) for example, will be
"https://example.com/".

3.2.7. updateViaCache

The updateViaCache getter steps are to return the service worker
registration’s update via cache mode.

3.2.8. update()

The update() method steps are:
1. Let registration be the service worker registration.
2. Let newestWorker be the result of running Get Newest Worker
algorithm passing registration as its argument.
3. If newestWorker is null, return a promise rejected with an
"InvalidStateError" DOMException and abort these steps.
4. If this’s relevant global object globalObject is a
ServiceWorkerGlobalScope object, and globalObject’s associated
service worker’s state is "installing", return a promise rejected
with an "InvalidStateError" DOMException and abort these steps.
5. Let promise be a promise.
6. Let job be the result of running Create Job with update,
registration’s storage key, registration’s scope url,
newestWorker’s script url, promise, and this’s relevant settings
object.
7. Set job’s worker type to newestWorker’s type.
8. Invoke Schedule Job with job.
9. Return promise.

3.2.9. unregister()

Note: The unregister() method unregisters the service worker
registration. It is important to note that the currently controlled
service worker client’s active service worker’s containing service
worker registration is effective until all the service worker clients
(including itself) using this service worker registration unload. That
is, the unregister() method only affects subsequent navigations.

The unregister() method steps are:
1. Let registration be the service worker registration.
2. Let promise be a new promise.
3. Let job be the result of running Create Job with unregister,
registration’s storage key, registration’s scope url, null,
promise, and this’s relevant settings object.
4. Invoke Schedule Job with job.
5. Return promise.

3.2.10. Event handler

The following is the event handler (and its corresponding event handler
event type) that must be supported, as event handler IDL attributes, by
all objects implementing ServiceWorkerRegistration interface:

event handler event handler event type
onupdatefound updatefound

3.3. navigator.serviceWorker

partial interface Navigator {
[SecureContext, SameObject] readonly attribute ServiceWorkerContainer serviceW
orker;
};

partial interface WorkerNavigator {
[SecureContext, SameObject] readonly attribute ServiceWorkerContainer serviceW
orker;
};

The serviceWorker getter steps are to return the ServiceWorkerContainer
object that is associated with this.

3.4. ServiceWorkerContainer

[SecureContext, Exposed=(Window,Worker)]
interface ServiceWorkerContainer : EventTarget {
readonly attribute ServiceWorker? controller;
readonly attribute Promise<ServiceWorkerRegistration> ready;

[NewObject] Promise<ServiceWorkerRegistration> register((TrustedScriptURL or U
SVString) scriptURL, optional RegistrationOptions options = {});

[NewObject] Promise<(ServiceWorkerRegistration or undefined)> getRegistration(
optional USVString clientURL = "");
[NewObject] Promise<FrozenArray<ServiceWorkerRegistration>> getRegistrations()
;

undefined startMessages();

// events
attribute EventHandler oncontrollerchange;
attribute EventHandler onmessage; // event.source of message events is Service
Worker object
attribute EventHandler onmessageerror;
};

dictionary RegistrationOptions {
USVString scope;
WorkerType type = "classic";
ServiceWorkerUpdateViaCache updateViaCache = "imports";
};

The user agent must create a ServiceWorkerContainer object when a
Navigator object or a WorkerNavigator object is created and associate
it with that object.

A ServiceWorkerContainer provides capabilities to register, unregister,
and update the service worker registrations, and provides access to the
state of the service worker registrations and their associated service
workers.

A ServiceWorkerContainer has an associated service worker client, which
is a service worker client whose global object is associated with the
Navigator object or the WorkerNavigator object that the
ServiceWorkerContainer is retrieved from.

A ServiceWorkerContainer object has an associated ready promise (a
promise or null). It is initially null.

A ServiceWorkerContainer object has a task source called the client
message queue, initially empty. A client message queue can be enabled
or disabled, and is initially disabled. When a ServiceWorkerContainer
object’s client message queue is enabled, the event loop must use it as
one of its task sources. When the ServiceWorkerContainer object’s
relevant global object is a Window object, all tasks queued on its
client message queue must be associated with its relevant settings
object’s associated document.

3.4.1. controller

controller attribute must run these steps:
1. Let client be this’s service worker client.
2. If client’s active service worker is null, then return null.
3. Return the result of getting the service worker object that
represents client’s active service worker in this’s relevant
settings object.

Note: navigator.serviceWorker.controller returns null if the request is
a force refresh (shift+refresh).

3.4.2. ready

ready attribute must run these steps:
1. If this’s ready promise is null, then set this’s ready promise to a
new promise.
2. Let readyPromise be this’s ready promise.
3. If readyPromise is pending, run the following substeps in parallel:
1. Let client by this’s service worker client.
2. Let storage key be the result of running obtain a storage key
given client.
3. Let registration be the result of running Match Service Worker
Registration given storage key and client’s creation URL.
4. If registration is not null, and registration’s active worker
is not null, queue a task on readyPromise’s relevant settings
object’s responsible event loop, using the DOM manipulation
task source, to resolve readyPromise with the result of
getting the service worker registration object that represents
registration in readyPromise’s relevant settings object.
4. Return readyPromise.

Note: The returned ready promise will never reject. If it does not
resolve in this algorithm, it will eventually resolve when a matching
service worker registration is registered and its active worker is set.
(See the relevant Activate algorithm step.)

3.4.3. register(scriptURL, options)

Note: The register(scriptURL, options) method creates or updates a
service worker registration for the given scope url. If successful, a
service worker registration ties the provided scriptURL to a scope url,
which is subsequently used for navigation matching.

The register(scriptURL, options) method steps are:
1. Let p be a promise.
2. Set scriptURL to the result of invoking Get Trusted Type compliant
string with TrustedScriptURL, this’s relevant global object,
scriptURL, "ServiceWorkerContainer register", and "script".
3. Let client be this’s service worker client.
4. Let scriptURL be the result of parsing scriptURL with this’s
relevant settings object’s API base URL.
5. Let scopeURL be null.
6. If options["scope"] exists, set scopeURL to the result of parsing
options["scope"] with this’s relevant settings object’s API base
URL.
7. Invoke Start Register with scopeURL, scriptURL, p, client, client’s
creation URL, options["type"], and options["updateViaCache"].
8. Return p.

3.4.4. getRegistration(clientURL)

getRegistration(clientURL) method steps are:
1. Let client be this’s service worker client.
2. Let storage key be the result of running obtain a storage key given
client.
3. Let clientURL be the result of parsing clientURL with this’s
relevant settings object’s API base URL.
4. If clientURL is failure, return a promise rejected with a
TypeError.
5. Set clientURL’s fragment to null.
6. If the origin of clientURL is not client’s origin, return a promise
rejected with a "SecurityError" DOMException.
7. Let promise be a new promise.
8. Run the following substeps in parallel:
1. Let registration be the result of running Match Service Worker
Registration given storage key and clientURL.
2. If registration is null, resolve promise with undefined and
abort these steps.
3. Resolve promise with the result of getting the service worker
registration object that represents registration in promise’s
relevant settings object.
9. Return promise.

3.4.5. getRegistrations()

getRegistrations() method steps are:
1. Let client be this’s service worker client.
2. Let client storage key be the result of running obtain a storage
key given client.
3. Let promise be a new promise.
4. Run the following steps in parallel:
1. Let registrations be a new list.
2. For each (storage key, scope) → registration of registration
map:
1. If storage key equals client storage key, then append
registration to registrations.
3. Queue a task on promise’s relevant settings object’s
responsible event loop, using the DOM manipulation task
source, to run the following steps:
1. Let registrationObjects be a new list.
2. For each registration of registrations:
1. Let registrationObj be the result of getting the
service worker registration object that represents
registration in promise’s relevant settings object.
2. Append registrationObj to registrationObjects.
3. Resolve promise with a new frozen array of
registrationObjects in promise’s relevant Realm.
5. Return promise.

3.4.6. startMessages()

The startMessages() method steps are to enable this’s client message
queue if it is not enabled.

3.4.7. Event handlers

The following are the event handlers (and their corresponding event
handler event types) that must be supported, as event handler IDL
attributes, by all objects implementing the ServiceWorkerContainer
interface:

event handler event handler event type
oncontrollerchange controllerchange
onmessage message
onmessageerror messageerror

The first time the onmessage setter steps are performed, enable this’s
client message queue.

3.5. Events

The following event is dispatched on ServiceWorker object:

Event name Interface Dispatched when…
statechange Event The state attribute of the ServiceWorker object is
changed.

The following event is dispatched on ServiceWorkerRegistration object:

Event name Interface Dispatched when…
updatefound Event The service worker registration’s installing worker
changes. (See step 8 of the Install algorithm.)

The following events are dispatched on ServiceWorkerContainer object:

Event name Interface Dispatched when…
controllerchange Event The service worker client’s active service
worker changes. (See step 9.2 of the Activate algorithm. The skip
waiting flag of a service worker causes activation of the service
worker registration to occur while service worker clients are using the
service worker registration, navigator.serviceWorker.controller
immediately reflects the active worker as the service worker that
controls the service worker client.)
message Event The service worker client receives a message from a
service worker. See postMessage(message, options).
messageerror Event The service worker client is sent a message that
cannot be deserialized from a service worker. See postMessage(message,
options).

3.6. NavigationPreloadManager

[SecureContext, Exposed=(Window,Worker)]
interface NavigationPreloadManager {
Promise<undefined> enable();
Promise<undefined> disable();
Promise<undefined> setHeaderValue(ByteString value);
Promise<NavigationPreloadState> getState();
};

dictionary NavigationPreloadState {
boolean enabled = false;
ByteString headerValue;
};

3.6.1. enable()

The enable() method steps are:
1. Let promise be a new promise.
2. Run the following steps in parallel:
1. Let registration be this’s associated service worker
registration.
2. If registration’s active worker is null, reject promise with
an "InvalidStateError" DOMException, and abort these steps.
3. Set registration’s navigation preload enabled flag.
4. Resolve promise with undefined.
3. Return promise.

3.6.2. disable()

The disable() method steps are:
1. Let promise be a new promise.
2. Run the following steps in parallel:
1. Let registration be this’s associated service worker
registration.
2. If registration’s active worker is null, reject promise with
an "InvalidStateError" DOMException, and abort these steps.
3. Unset registration’s navigation preload enabled flag.
4. Resolve promise with undefined.
3. Return promise.

3.6.3. setHeaderValue(value)

The setHeaderValue(value) method steps are:
1. Let promise be a new promise.
2. Run the following steps in parallel:
1. Let registration be this’s associated service worker
registration.
2. If registration’s active worker is null, reject promise with
an "InvalidStateError" DOMException, and abort these steps.
3. Set registration’s navigation preload header value to value.
4. Resolve promise with undefined.
3. Return promise.

3.6.4. getState()

The getState() method steps are:
1. Let promise be a new promise.
2. Run the following steps in parallel:
1. Let registration be this’s associated service worker
registration.
2. Let state be a new NavigationPreloadState dictionary.
3. If registration’s navigation preload enabled flag is set, set
state["enabled"] to true.
4. Set state["headerValue"] to registration’s navigation preload
header value.
5. Resolve promise with state.
3. Return promise.

4. Execution Context

Serving Cached Resources:
// caching.js
self.addEventListener("install", event => {
event.waitUntil(
// Open a cache of resources.
caches.open("shell-v1").then(cache => {
// Begins the process of fetching them. Succeeds only once all
// resources have been stored. Even just one failing resource
// causes the entire operation to fail.
return cache.addAll([
"/app.html",
"/assets/v1/base.css",
"/assets/v1/app.js",
"/assets/v1/logo.png",
"/assets/v1/intro_video.webm"
]);
})
);
});

self.addEventListener("fetch", event => {
// No "fetch" events are dispatched to the service worker until it
// successfully installs and activates.

// All operations on caches are async, including matching URLs, so we use
// promises heavily. e.respondWith() even takes promises to enable this:
event.respondWith(
caches.match(e.request).then(response => {
return response || fetch(e.request);
}).catch(() => {
return caches.match("/fallback.html");
})
);
});

4.1. ServiceWorkerGlobalScope

[Global=(Worker,ServiceWorker), Exposed=ServiceWorker, SecureContext]
interface ServiceWorkerGlobalScope : WorkerGlobalScope {
[SameObject] readonly attribute Clients clients;
[SameObject] readonly attribute ServiceWorkerRegistration registration;
[SameObject] readonly attribute ServiceWorker serviceWorker;

[NewObject] Promise<undefined> skipWaiting();

attribute EventHandler oninstall;
attribute EventHandler onactivate;
attribute EventHandler onfetch;

attribute EventHandler onmessage;
attribute EventHandler onmessageerror;
};

A ServiceWorkerGlobalScope object represents the global execution
context of a service worker.

A ServiceWorkerGlobalScope object has an associated service worker (a
service worker).

A ServiceWorkerGlobalScope object has an associated force bypass cache
for import scripts flag. It is initially unset.

A ServiceWorkerGlobalScope object has an associated race response map
which is an ordered map where the keys are requests and the values are
race response.

A race response is a struct used to contain the network response when
"race-network-and-fetch-handler" performs. It has a value, which is a
response, "pending", or null.

Note: ServiceWorkerGlobalScope object provides generic, event-driven,
time-limited script execution contexts that run at an origin. Once
successfully registered, a service worker is started, kept alive and
killed by their relationship to events, not service worker clients. Any
type of synchronous requests must not be initiated inside of a service
worker.

4.1.1. clients

The clients getter steps are to return the Clients object that is
associated with this.

4.1.2. registration

The registration getter steps are to return the result of getting the
service worker registration object representing this’s service worker’s
containing service worker registration in this’s relevant settings
object.

4.1.3. serviceWorker

The serviceWorker getter steps are to return the result of getting the
service worker object that represents this’s service worker in this’s
relevant settings object.

4.1.4. skipWaiting()

Note: The skipWaiting() method allows this service worker to progress
from the registration’s waiting position to active even while service
worker clients are using the registration.

The skipWaiting() method steps are:
1. Let promise be a new promise.
2. Run the following substeps in parallel:
1. Set service worker’s skip waiting flag.
2. Invoke Try Activate with service worker’s containing service
worker registration.
3. Resolve promise with undefined.
3. Return promise.

4.1.5. Event handlers

The following are the event handlers (and their corresponding event
handler event types) that must be supported, as event handler IDL
attributes, by all objects implementing the ServiceWorkerGlobalScope
interface:

event handler event handler event type
oninstall install
onactivate activate
onfetch fetch
onmessage message
onmessageerror messageerror

4.2. Client

[Exposed=ServiceWorker]
interface Client {
readonly attribute USVString url;
readonly attribute FrameType frameType;
readonly attribute DOMString id;
readonly attribute ClientType type;
undefined postMessage(any message, sequence<object> transfer);
undefined postMessage(any message, optional StructuredSerializeOptions options
= {});
};

[Exposed=ServiceWorker]
interface WindowClient : Client {
readonly attribute VisibilityState visibilityState;
readonly attribute boolean focused;
[SameObject] readonly attribute FrozenArray<USVString> ancestorOrigins;
[NewObject] Promise<WindowClient> focus();
[NewObject] Promise<WindowClient?> navigate(USVString url);
};

enum FrameType {
"auxiliary",
"top-level",
"nested",
"none"
};

A Client object has an associated service worker client (a service
worker client).

A Client object has an associated frame type, which is one of
"auxiliary", "top-level", "nested", and "none". Unless stated otherwise
it is "none".

A WindowClient object has an associated browsing context, which is its
service worker client’s global object’s browsing context.

A WindowClient object has an associated visibility state, which is one
of visibilityState attribute value.

A WindowClient object has an associated focus state, which is either
true or false (initially false).

A WindowClient object has an associated ancestor origins array.

4.2.1. url

The url getter steps are to return this’s associated service worker
client’s serialized creation URL.

4.2.2. frameType

The frameType getter steps are to return this’s frame type.

4.2.3. id

The id getter steps are to return this’s associated service worker
client’s id.

4.2.4. type

The type getter steps are:
1. Let client be this’s service worker client.
2. If client is an environment settings object, then:
1. If client is a window client, return "window".
2. Else if client is a dedicated worker client, return "worker".
3. Else if client is a shared worker client, return
"sharedworker".
3. Else:
1. Return "window".

4.2.5. postMessage(message, transfer)

The postMessage(message, transfer) method steps are:
1. Let options be «[ "transfer" → transfer ]».
2. Invoke postMessage(message, options) with message and options as
the arguments.

4.2.6. postMessage(message, options)

The postMessage(message, options) method steps are:
1. Let contextObject be this.
2. Let sourceSettings be the contextObject’s relevant settings object.
3. Let serializeWithTransferResult be
StructuredSerializeWithTransfer(message, options["transfer"]).
Rethrow any exceptions.
4. Run the following steps in parallel:
1. Let targetClient be null.
2. For each service worker client client:
1. If client is the contextObject’s service worker client,
set targetClient to client, and break.
3. If targetClient is null, return.
4. Let destination be the ServiceWorkerContainer object whose
associated service worker client is targetClient.
5. Add a task that runs the following steps to destination’s
client message queue:
1. Let origin be sourceSettings’s origin.
2. Let source be the result of getting the service worker
object that represents contextObject’s relevant global
object’s service worker in targetClient.
3. Let deserializeRecord be
StructuredDeserializeWithTransfer(serializeWithTransferRe
sult, destination’s relevant Realm).
If this throws an exception, catch it, fire an event
named messageerror at destination, using MessageEvent,
with its origin initialized to origin and the source
attribute initialized to source, and then abort these
steps.
4. Let messageClone be deserializeRecord.[[Deserialized]].
5. Let newPorts be a new frozen array consisting of all
MessagePort objects in
deserializeRecord.[[TransferredValues]], if any.
6. Dispatch an event named message at destination, using
MessageEvent, with its origin initialized to origin, the
source attribute initialized to source, the data
attribute initialized to messageClone, and the ports
attribute initialized to newPorts.

4.2.7. visibilityState

The visibilityState getter steps are to return this’s visibility state.

4.2.8. focused

The focused getter steps are to return this’s focus state.

4.2.9. ancestorOrigins

The ancestorOrigins getter steps are to return this’s associated
ancestor origins array.

4.2.10. focus()

The focus() method steps are:
1. If no Window in this origin has transient activation, return a
promise rejected with an "InvalidAccessError" DOMException.
2. Let serviceWorkerEventLoop be the surrounding agent’s event loop.
3. Let promise be a new promise.
4. Queue a task to run the following steps on this’s associated
service worker client’s responsible event loop using the user
interaction task source:
1. Run the focusing steps with this’s browsing context.
2. Let frameType be the result of running Get Frame Type with
this’s browsing context.
3. Let visibilityState be this’s browsing context’s active
document’s visibilityState attribute value.
4. Let focusState be the result of running the has focus steps
with this’s browsing context’s active document.
5. Let ancestorOriginsList be this’s browsing context’s active
document’s ancestor origins list’s associated list.
6. Queue a task to run the following steps on
serviceWorkerEventLoop using the DOM manipulation task source:
1. Let windowClient be the result of running Create Window
Client with this’s associated service worker client,
frameType, visibilityState, focusState, and
ancestorOriginsList.
2. If windowClient’s focus state is true, resolve promise
with windowClient.
3. Else, reject promise with a TypeError.
5. Return promise.

4.2.11. navigate(url)

The navigate(url) method steps are:
1. Let url be the result of parsing url with this’s relevant settings
object’s API base URL.
2. If url is failure, return a promise rejected with a TypeError.
3. If url is about:blank, return a promise rejected with a TypeError.
4. If this’s associated service worker client’s active service worker
is not this’s relevant global object’s service worker, return a
promise rejected with a TypeError.
5. Let serviceWorkerEventLoop be the current global object’s event
loop.
6. Let promise be a new promise.
7. Queue a task to run the following steps on this’s associated
service worker client’s responsible event loop using the user
interaction task source:
1. Let browsingContext be this’s browsing context.
2. If browsingContext’s associated document is not fully active,
queue a task to reject promise with a TypeError, on
serviceWorkerEventLoop using the DOM manipulation task source,
and abort these steps.
3. HandleNavigate: Navigate browsingContext to url, using
browsingContext’s associated document, with exceptionsEnabled
true.
4. If the algorithm steps invoked in the step labeled
HandleNavigate throws an exception, queue a task to reject
promise with the exception, on serviceWorkerEventLoop using
the DOM manipulation task source, and abort these steps.
5. Let frameType be the result of running Get Frame Type with
browsingContext.
6. Let visibilityState be browsingContext’s active document’s
visibilityState attribute value.
7. Let focusState be the result of running the has focus steps
with browsingContext’s active document.
8. Let ancestorOriginsList be browsingContext’s active document’s
ancestor origins list’s associated list.
9. Queue a task to run the following steps on
serviceWorkerEventLoop using the DOM manipulation task source:
1. If browsingContext’s Window object’s environment settings
object’s creation URL’s origin is not the same as the
service worker’s origin, resolve promise with null and
abort these steps.
2. Let windowClient be the result of running Create Window
Client with this’s service worker client, frameType,
visibilityState, focusState, and ancestorOriginsList.
3. Resolve promise with windowClient.
8. Return promise.

4.3. Clients

[Exposed=ServiceWorker]
interface Clients {
// The objects returned will be new instances every time
[NewObject] Promise<(Client or undefined)> get(DOMString id);
[NewObject] Promise<FrozenArray<Client>> matchAll(optional ClientQueryOptions
options = {});
[NewObject] Promise<WindowClient?> openWindow(USVString url);
[NewObject] Promise<undefined> claim();
};

dictionary ClientQueryOptions {
boolean includeUncontrolled = false;
ClientType type = "window";
};

enum ClientType {
"window",
"worker",
"sharedworker",
"all"
};

The user agent must create a Clients object when a
ServiceWorkerGlobalScope object is created and associate it with that
object.

4.3.1. get(id)

The get(id) method steps are:
1. Let promise be a new promise.
2. Run these substeps in parallel:
1. For each service worker client client where the result of
running obtain a storage key given client equals the
associated service worker’s containing service worker
registration’s storage key:
1. If client’s id is not id, continue.
2. Wait for either client’s execution ready flag to be set
or for client’s discarded flag to be set.
3. If client’s execution ready flag is set, then invoke
Resolve Get Client Promise with client and promise, and
abort these steps.
2. Resolve promise with undefined.
3. Return promise.

4.3.2. matchAll(options)

The matchAll(options) method steps are:
1. Let promise be a new promise.
2. Run the following steps in parallel:
1. Let targetClients be a new list.
2. For each service worker client client where the result of
running obtain a storage key given client equals the
associated service worker’s containing service worker
registration’s storage key:
1. If client’s execution ready flag is unset or client’s
discarded flag is set, continue.
2. If client is not a secure context, continue.
3. If options["includeUncontrolled"] is false, and if
client’s active service worker is not the associated
service worker, continue.
4. Add client to targetClients.
3. Let matchedWindowData be a new list.
4. Let matchedClients be a new list.
5. For each service worker client client in targetClients:
1. If options["type"] is "window" or "all", and client is
not an environment settings object or is a window client,
then:
1. Let windowData be «[ "client" → client,
"ancestorOriginsList" → a new list ]».
2. Let browsingContext be null.
3. Let isClientEnumerable be true.
4. If client is an environment settings object, set
browsingContext to client’s global object’s browsing
context.
5. Else, set browsingContext to client’s target
browsing context.
6. Queue a task task to run the following substeps on
browsingContext’s event loop using the user
interaction task source:
1. If browsingContext has been discarded, then set
isClientEnumerable to false and abort these
steps.
2. If client is a window client and client’s
associated document is not browsingContext’s
active document, then set isClientEnumerable to
false and abort these steps.
3. Set windowData["frameType"] to the result of
running Get Frame Type with browsingContext.
4. Set windowData["visibilityState"] to
browsingContext’s active document’s
visibilityState attribute value.
5. Set windowData["focusState"] to the result of
running the has focus steps with
browsingContext’s active document as the
argument.
6. If client is a window client, then set
windowData["ancestorOriginsList"] to
browsingContext’s active document’s ancestor
origins list’s associated list.
7. Wait for task to have executed.
Note: Wait is a blocking wait, but implementers may
run the iterations in parallel as long as the state
is not broken.
8. If isClientEnumerable is true, then:
1. Add windowData to matchedWindowData.
2. Else if options["type"] is "worker" or "all" and client
is a dedicated worker client, or options["type"] is
"sharedworker" or "all" and client is a shared worker
client, then:
1. Add client to matchedClients.
6. Queue a task to run the following steps on promise’s relevant
settings object’s responsible event loop using the DOM
manipulation task source:
1. Let clientObjects be a new list.
2. For each windowData in matchedWindowData:
1. Let windowClient be the result of running Create
Window Client algorithm with windowData["client"],
windowData["frameType"],
windowData["visibilityState"],
windowData["focusState"], and
windowData["ancestorOriginsList"] as the arguments.
2. Append windowClient to clientObjects.
3. For each client in matchedClients:
1. Let clientObject be the result of running Create
Client algorithm with client as the argument.
2. Append clientObject to clientObjects.
4. Sort clientObjects such that:
# WindowClient objects whose browsing context has been
focused are placed first, sorted in the most
recently focused order.
# WindowClient objects whose browsing context has
never been focused are placed next, sorted in their
service worker client’s creation order.
# Client objects whose associated service worker
client is a worker client are placed next, sorted in
their service worker client’s creation order.
Note: Window clients are always placed before worker
clients.
5. Resolve promise with a new frozen array of clientObjects
in promise’s relevant Realm.
3. Return promise.

4.3.3. openWindow(url)

The openWindow(url) method steps are:
1. Let url be the result of parsing url with this’s relevant settings
object’s API base URL.
2. If url is failure, return a promise rejected with a TypeError.
3. If url is about:blank, return a promise rejected with a TypeError.
4. If no Window in this origin has transient activation, return a
promise rejected with an "InvalidAccessError" DOMException.
5. Let serviceWorkerEventLoop be the current global object’s event
loop.
6. Let promise be a new promise.
7. Run these substeps in parallel:
1. Let newContext be a new top-level browsing context.
2. Queue a task to run the following steps on newContext’s Window
object’s environment settings object’s responsible event loop
using the user interaction task source:
1. HandleNavigate: Navigate newContext to url with
exceptionsEnabled true, and historyHandling "replace".
2. If the algorithm steps invoked in the step labeled
HandleNavigate throws an exception, queue a task to
reject promise with the exception, on
serviceWorkerEventLoop using the DOM manipulation task
source, and abort these steps.
3. Let frameType be the result of running Get Frame Type
with newContext.
4. Let visibilityState be newContext’s active document’s
visibilityState attribute value.
5. Let focusState be the result of running the has focus
steps with newContext’s active document as the argument.
6. Let ancestorOriginsList be newContext’s active document’s
ancestor origins list’s associated list.
7. Queue a task to run the following steps on
serviceWorkerEventLoop using the DOM manipulation task
source:
1. If the result of running obtain a storage key given
newContext’s Window object’s environment settings
object is not equal to the service worker’s
containing service worker registration’s storage
key, then resolve promise with null and abort these
steps.
2. Let client be the result of running Create Window
Client with newContext’s Window object’s environment
settings object, frameType, visibilityState,
focusState, and ancestorOriginsList as the
arguments.
3. Resolve promise with client.
8. Return promise.

4.3.4. claim()

The claim() method steps are:
1. If the service worker is not an active worker, return a promise
rejected with an "InvalidStateError" DOMException.
2. Let promise be a new promise.
3. Run the following substeps in parallel:
1. For each service worker client client where the result of
running obtain a storage key given client equals the service
worker’s containing service worker registration’s storage key:
1. If client’s execution ready flag is unset or client’s
discarded flag is set, continue.
2. If client is not a secure context, continue.
3. Let storage key be the result of running obtain a storage
key given client.
4. Let registration be the result of running Match Service
Worker Registration given storage key and client’s
creation URL.
5. If registration is not the service worker’s containing
service worker registration, continue.
Note: registration will be null if the service worker’s
containing service worker registration is unregistered.
6. If client’s active service worker is not the service
worker, then:
1. Invoke Handle Service Worker Client Unload with
client as the argument.
2. Set client’s active service worker to service
worker.
3. Invoke Notify Controller Change algorithm with
client as the argument.
2. Resolve promise with undefined.
4. Return promise.

4.4. ExtendableEvent

[Exposed=ServiceWorker]
interface ExtendableEvent : Event {
constructor(DOMString type, optional ExtendableEventInit eventInitDict = {});
undefined waitUntil(Promise<any> f);
};

dictionary ExtendableEventInit : EventInit {
// Defined for the forward compatibility across the derived events
};

An ExtendableEvent object has an associated extend lifetime promises
(an array of promises). It is initially an empty array.

An ExtendableEvent object has an associated pending promises count (the
number of pending promises in the extend lifetime promises). It is
initially set to zero.

An ExtendableEvent object has an associated timed out flag. It is
initially unset, and is set after an optional user agent imposed delay
if the pending promises count is greater than zero.

An ExtendableEvent object is said to be active when its timed out flag
is unset and either its pending promises count is greater than zero or
its dispatch flag is set.

Service workers have two lifecycle events, install and activate.
Service workers use the ExtendableEvent interface for activate event
and install event.

Service worker extensions that define event handlers may also use or
extend the ExtendableEvent interface.

4.4.1. event.waitUntil(f)

Note: waitUntil() method extends the lifetime of the event.

The waitUntil(f) method steps are to add lifetime promise f to this.
To add lifetime promise promise (a promise) to event (an
ExtendableEvent), run these steps:
1. If event’s isTrusted attribute is false, throw an
"InvalidStateError" DOMException.
2. If event is not active, throw an "InvalidStateError" DOMException.
Note: If no lifetime extension promise has been added in the task
that called the event handlers, calling waitUntil() in subsequent
asynchronous tasks will throw.
3. Add promise to event’s extend lifetime promises.
4. Increment event’s pending promises count by one.
Note: The pending promises count is incremented even if the given
promise has already been settled. The corresponding count decrement
is done in the microtask queued by the reaction to the promise.
5. Upon fulfillment or rejection of promise, queue a microtask to run
these substeps:
1. Decrement event’s pending promises count by one.
2. If event’s pending promises count is 0, then:
1. Let registration be the current global object’s
associated service worker’s containing service worker
registration.
2. If registration is unregistered, invoke Try Clear
Registration with registration.
3. If registration is not null, invoke Try Activate with
registration.

The user agent should not terminate a service worker if Service Worker
Has No Pending Events returns false for that service worker.

Service workers and extensions that define event handlers may define
their own behaviors, allowing the extend lifetime promises to suggest
operation length, and the rejected state of any of the promise in
extend lifetime promises to suggest operation failure.

Note: Service workers delay treating the installing worker as
"installed" until all the promises in the install event’s extend
lifetime promises resolve successfully. (See the relevant Install
algorithm step.) If any of the promises rejects, the installation
fails. This is primarily used to ensure that a service worker is not
considered "installed" until all of the core caches it depends on are
populated. Likewise, service workers delay treating the active worker
as "activated" until all the promises in the activate event’s extend
lifetime promises settle. (See the relevant Activate algorithm step.)
This is primarily used to ensure that any functional events are not
dispatched to the service worker until it upgrades database schemas and
deletes the outdated cache entries.

4.5. InstallEvent

[Exposed=ServiceWorker]
interface InstallEvent : ExtendableEvent {
constructor(DOMString type, optional ExtendableEventInit eventInitDict = {});
Promise<undefined> addRoutes((RouterRule or sequence<RouterRule>) rules);
};

dictionary RouterRule {
required RouterCondition condition;
required RouterSource source;
};

dictionary RouterCondition {
URLPatternCompatible urlPattern;
ByteString requestMethod;
RequestMode requestMode;
RequestDestination requestDestination;
RunningStatus runningStatus;

sequence<RouterCondition> _or;
RouterCondition not;
};

typedef (RouterSourceDict or RouterSourceEnum) RouterSource;

dictionary RouterSourceDict {
DOMString cacheName;
};

enum RunningStatus { "running", "not-running" };
enum RouterSourceEnum {
"cache",
"fetch-event",
"network",
"race-network-and-fetch-handler"
};

A count router condition result is a struct that consists of:
* A condition count (a number).
* A quota exceeded (a boolean).

4.5.1. event.addRoutes(rules)

Note: addRoutes(rules) registers rules for this service worker to
offload simple tasks that the fetch event handler ordinarily does.

The addRoutes(rules) method steps are:
1. If rules is a RouterRule dictionary, set rules to « rules ».
2. Let serviceWorker be the current global object’s associated service
worker.
3. For each rule of rules:
1. If running the Verify Router Condition algorithm with
rule["condition"] and serviceWorker returns false, return a
promise rejected with a TypeError.
2. If rule["source"] is either of "fetch-event" or
"race-network-and-fetch-handler", and serviceWorker’s set of
event types to handle does not contain fetch, return a promise
rejected with a TypeError.
4. Let lifetimePromise be a new promise.
5. Add lifetime promise lifetimePromise to this.
Note: event.addRoutes(rules) extends the lifetime of the event by
default as if event.waitUntil(promise) is called.
6. Let promise be a new promise.
7. Upon fulfillment or rejection of promise, resolve lifetimePromise
with undefined.
Note: this step is for making lifetimePromise always fullfilled to
avoid the install event failure.
8. Enqueue the following steps to [[service worker queue]]:
1. Let allRules be a copy of serviceWorker’s list of router
rules.
2. For each rule of rules:
1. Append rule to allRules.
3. If running the Check Router Registration Limit with allRules
returns false, reject promise with a TypeError.
4. Set serviceWorker’s list of router rules to allRules.
5. Let serviceWorkerEventLoop be the current global object’s
event loop.
6. Queue a task to run the following steps on
serviceWorkerEventLoop using the DOM manipulation task source:
1. Resolve promise with undefined.
9. Return promise.

4.6. FetchEvent

[Exposed=ServiceWorker]
interface FetchEvent : ExtendableEvent {
constructor(DOMString type, FetchEventInit eventInitDict);
[SameObject] readonly attribute Request request;
readonly attribute Promise<any> preloadResponse;
readonly attribute DOMString clientId;
readonly attribute DOMString resultingClientId;
readonly attribute DOMString replacesClientId;
readonly attribute Promise<undefined> handled;

undefined respondWith(Promise<Response> r);
};

dictionary FetchEventInit : ExtendableEventInit {
required Request request;
Promise<any> preloadResponse;
DOMString clientId = "";
DOMString resultingClientId = "";
DOMString replacesClientId = "";
Promise<undefined> handled;
};

Service workers have an essential functional event fetch. For fetch
event, service workers use the FetchEvent interface which extends the
ExtendableEvent interface.

Each event using FetchEvent interface has an associated potential
response (a response), initially set to null, and the following
associated flags that are initially unset:
* wait to respond flag
* respond-with entered flag
* respond-with error flag

4.6.1. event.request

request attribute must return the value it was initialized to.

4.6.2. event.preloadResponse

preloadResponse attribute must return the value it was initialized to.
When an event is created the attribute must be initialized to a promise
resolved with undefined.

4.6.3. event.clientId

clientId attribute must return the value it was initialized to. When an
event is created the attribute must be initialized to the empty string.

4.6.4. event.resultingClientId

resultingClientId attribute must return the value it was initialized
to. When an event is created the attribute must be initialized to the
empty string.

4.6.5. event.replacesClientId

replacesClientId attribute must return the value it was initialized to.
When an event is created the attribute must be initialized to the empty
string.

4.6.6. event.handled

handled attribute must return the value it was initialized to. When an
event is created the attribute must be initialized to a pending
promise.

4.6.7. event.respondWith(r)

Note: Developers can set the argument r with either a promise that
resolves with a Response object or a Response object (which is
automatically cast to a promise). Otherwise, a network error is
returned to Fetch. Renderer-side security checks about tainting for
cross-origin content are tied to the types of filtered responses
defined in Fetch.

respondWith(r) method steps are:
1. Let event be this.
2. If event’s dispatch flag is unset, throw an "InvalidStateError"
DOMException.
3. If event’s respond-with entered flag is set, throw an
"InvalidStateError" DOMException.
4. Add lifetime promise r to event.
Note: event.respondWith(r) extends the lifetime of the event by
default as if event.waitUntil(r) is called.
5. Set event’s stop propagation flag and stop immediate propagation
flag.
6. Set event’s respond-with entered flag.
7. Set event’s wait to respond flag.
8. Let targetRealm be event’s relevant Realm.
9. Upon rejection of r:
1. Set event’s respond-with error flag.
2. Unset event’s wait to respond flag.
10. Upon fulfillment of r with response:
1. If response is not a Response object, then set the
respond-with error flag.
Note: If the respond-with error flag is set, a network error
is returned to Fetch through Handle Fetch algorithm. (See the
step 21.1.) Otherwise, the value response is returned to Fetch
through Handle Fetch algorithm. (See the step 22.1.)
2. Else:
1. Let bytes be an empty byte sequence.
2. Let end-of-body be false.
3. Let done be false.
4. Let potentialResponse be a copy of response’s associated
response, except for its body.
5. If response’s body is non-null, run these substeps:
1. Let reader be the result of getting a reader from
response’s body’s stream.
2. Let pullAlgorithm be an action that runs these
steps:
1. Let readRequest be a new read request with the
following items:

chunk steps, given chunk

1. Assert: chunk is a Uint8Array.
2. Append the bytes represented by chunk
to bytes.
3. Increment potentialResponse’s body
info’s encoded size by bytes’s byte
length.
4. Increment potentialResponse’s body
info’s decoded size by bytes’s byte
length.
5. Perform !
DetachArrayBuffer(chunk.[[ViewedArray
Buffer]]).

close steps

1. Set end-of-body to true.

error steps

1. error newStream with a TypeError.

2. Read a chunk from reader given readRequest.
3. Let cancelAlgorithm be an action that cancels
reader.
4. Let highWaterMark be a non-negative, non-NaN number,
chosen by the user agent.
5. Let sizeAlgorithm be an algorithm that accepts a
chunk object and returns a non-negative, non-NaN,
non-infinite number, chosen by the user agent.
6. Let newStream be a new ReadableStream, set up with
the pullAlgorithm pullAlgorithm, the cancelAlgorithm
cancelAlgorithm, the highWaterMark highWaterMark,
and sizeAlgorithm sizeAlgorithm, in targetRealm.
7. Set potentialResponse’s body to a new body whose
stream is newStream.
8. Run these subsubsteps repeatedly in parallel while
done is false:
1. If newStream is errored, then set done to true.
2. Otherwise, if bytes is empty and end-of-body is
true, then close newStream and set done to
true.
3. Otherwise, if bytes is not empty, run these
subsubsubsteps:
1. Let chunk be a subsequence of bytes
starting from the beginning of bytes.
2. Remove chunk from bytes.
3. Let buffer be an ArrayBuffer object
created in targetRealm and containing
chunk.
4. Enqueue a Uint8Array object created in
targetRealm and wrapping buffer to
newStream.
Note: These substeps are meant to produce the observable
equivalent of "piping" response’s body’s stream into
potentialResponse.
Note: The data written by the service worker in chunks
are not guaranteed to read in identical chunks by the
client that receives the data. That is, the client will
read the same data that was written, but it may be
chunked differently by the browser.
6. Set event’s potential response to potentialResponse.
3. Unset event’s wait to respond flag.

4.7. ExtendableMessageEvent

[Exposed=ServiceWorker]
interface ExtendableMessageEvent : ExtendableEvent {
constructor(DOMString type, optional ExtendableMessageEventInit eventInitDict
= {});
readonly attribute any data;
readonly attribute USVString origin;
readonly attribute DOMString lastEventId;
[SameObject] readonly attribute (Client or ServiceWorker or MessagePort)? sour
ce;
readonly attribute FrozenArray<MessagePort> ports;
};

dictionary ExtendableMessageEventInit : ExtendableEventInit {
any data = null;
USVString origin = "";
DOMString lastEventId = "";
(Client or ServiceWorker or MessagePort)? source = null;
sequence<MessagePort> ports = [];
};

Each ExtendableMessageEvent has an origin (an origin, a string, or
null), initially null.

Service workers define the extendable message event to allow extending
the lifetime of the event. For the message event, service workers use
the ExtendableMessageEvent interface which extends the ExtendableEvent
interface.

4.7.1. event.data

The data attribute must return the value it was initialized to. When
the object is created, this attribute must be initialized to null. It
represents the message being sent.

4.7.2. event.origin

The origin attribute represents the origin of the service worker client
that sent the message.

The origin getter steps are:
1. If this’s origin is an origin, then return the serialization of
this’s origin.
2. If this’s origin is null, then return the empty string.
3. Return this’s origin.

When the origin attribute is "initialized" (during
ExtendableMessageEvent’s constructor, for example), the initialization
value is placed into the object’s origin.

4.7.3. event.lastEventId

The lastEventId attribute must return the value it was initialized to.
When the object is created, this attribute must be initialized to the
empty string.

Objects implementing the ExtendableMessageEvent interface’s extract an
origin steps are to return this’s origin if it is an origin; otherwise
null.

4.7.4. event.source

The source attribute must return the value it was initialized to. When
the object is created, this attribute must be initialized to null. It
represents the Client object from which the message is sent.

4.7.5. event.ports

The ports attribute must return the value it was initialized to. When
the object is created, this attribute must be initialized to the empty
array. It represents the MessagePort array being sent.

4.8. Events

The following events, called service worker events, are dispatched on
ServiceWorkerGlobalScope object:

Event name Interface Category Dispatched when…
install InstallEvent Lifecycle The service worker’s containing service
worker registration’s installing worker changes. (See step 11.2 of the
Install algorithm.)
activate ExtendableEvent Lifecycle The service worker’s containing
service worker registration’s active worker changes. (See step 12.2 of
the Activate algorithm.)
fetch FetchEvent Functional The http fetch invokes Handle Fetch with
request. As a result of performing Handle Fetch, the service worker
returns a response to the http fetch. The response, represented by a
Response object, can be retrieved from a Cache object or directly from
network using self.fetch(input, init) method. (A custom Response object
can be another option.)
push PushEvent Functional (See Firing a push event.)
notificationclick NotificationEvent Functional (See Activating a
notification.)
notificationclose NotificationEvent Functional (See Closing a
notification.)
sync SyncEvent Functional (See Firing a sync event.)
canmakepayment CanMakePaymentEvent Functional (See Handling a
CanMakePaymentEvent.)
paymentrequest PaymentRequestEvent Functional (See Handling a
PaymentRequestEvent.)
message ExtendableMessageEvent Legacy When it receives a message.
messageerror ExtendableMessageEvent Legacy When it was sent a message
that cannot be deserialized.

5. Caches

To allow authors to fully manage their content caches for offline use,
the Window and the WorkerGlobalScope provide the asynchronous caching
methods that open and manipulate Cache objects. An origin can have
multiple, named Cache objects, whose contents are entirely under the
control of scripts. Caches are not shared across origins, and they are
completely isolated from the browser’s HTTP cache.

5.1. Constructs

A request response list is a list of tuples consisting of a request (a
request) and a response (a response).

The relevant request response list is the instance that this
represents.

A name to cache map is an ordered map whose entry consists of a key (a
string that represents the name of a request response list) and a value
(a request response list).

The relevant name to cache map for a CacheStorage object is the name to
cache map associated with the result of running obtain a local storage
bottle map with the object’s relevant settings object and "caches".

5.2. Understanding Cache Lifetimes

The Cache instances are not part of the browser’s HTTP cache. The Cache
objects are exactly what authors have to manage themselves. The Cache
objects do not get updated unless authors explicitly request them to
be. The Cache objects do not expire unless authors delete the entries.
The Cache objects do not disappear just because the service worker
script is updated. That is, caches are not updated automatically.
Updates must be manually managed. This implies that authors should
version their caches by name and make sure to use the caches only from
the version of the service worker that can safely operate on.

5.3. self.caches

partial interface mixin WindowOrWorkerGlobalScope {
[SecureContext, SameObject] readonly attribute CacheStorage caches;
};

5.3.1. caches

caches getter steps are to return this’s associated CacheStorage
object.

5.4. Cache

[SecureContext, Exposed=(Window,Worker)]
interface Cache {
[NewObject] Promise<(Response or undefined)> match(RequestInfo request, option
al CacheQueryOptions options = {});
[NewObject] Promise<FrozenArray<Response>> matchAll(optional RequestInfo reque
st, optional CacheQueryOptions options = {});
[NewObject] Promise<undefined> add(RequestInfo request);
[NewObject] Promise<undefined> addAll(sequence<RequestInfo> requests);
[NewObject] Promise<undefined> put(RequestInfo request, Response response);
[NewObject] Promise<boolean> delete(RequestInfo request, optional CacheQueryOp
tions options = {});
[NewObject] Promise<FrozenArray<Request>> keys(optional RequestInfo request, o
ptional CacheQueryOptions options = {});
};

dictionary CacheQueryOptions {
boolean ignoreSearch = false;
boolean ignoreMethod = false;
boolean ignoreVary = false;
};

A Cache object represents a request response list. Multiple separate
objects implementing the Cache interface across documents and workers
can all be associated with the same request response list
simultaneously.

A cache batch operation is a struct that consists of:
* A type ("delete" or "put").
* A request (a request).
* A response (a response).
* An options (a CacheQueryOptions).

5.4.1. match(request, options)

The match(request, options) method steps are:
1. Let promise be a new promise.
2. Run these substeps in parallel:
1. Let p be the result of running the algorithm specified in
matchAll(request, options) method with request and options.
2. Wait until p settles.
3. If p rejects with an exception, then:
1. Reject promise with that exception.
4. Else if p resolves with an array, responses, then:
1. If responses is an empty array, then:
1. Resolve promise with undefined.
2. Else:
1. Resolve promise with the first element of responses.
3. Return promise.

5.4.2. matchAll(request, options)

The matchAll(request, options) method steps are:
1. Let r be null.
2. If the optional argument request is not omitted, then:
1. If request is a Request object, then:
1. Set r to request’s request.
2. If r’s method is not `GET` and options.ignoreMethod is
false, return a promise resolved with an empty array.
2. Else if request is a string, then:
1. Set r to the associated request of the result of invoking
the initial value of Request as constructor with request
as its argument. If this throws an exception, return a
promise rejected with that exception.
3. Let realm be this’s relevant realm.
4. Let promise be a new promise.
5. Run these substeps in parallel:
1. Let responses be an empty list.
2. If the optional argument request is omitted, then:
1. For each requestResponse of the relevant request response
list:
1. Add a copy of requestResponse’s response to
responses.
3. Else:
1. Let requestResponses be the result of running Query Cache
with r and options.
2. For each requestResponse of requestResponses:
1. Add a copy of requestResponse’s response to
responses.
4. For each response of responses:
1. If response’s type is "opaque" and cross-origin resource
policy check with promise’s relevant settings object’s
origin, promise’s relevant settings object, "", and
response’s internal response returns blocked, then reject
promise with a TypeError and abort these steps.
5. Queue a task, on promise’s relevant settings object’s
responsible event loop using the DOM manipulation task source,
to perform the following steps:
1. Let responseList be a list.
2. For each response of responses:
1. Add a new Response object associated with response
and a new Headers object whose guard is "immutable"
to responseList.
3. Resolve promise with a frozen array created from
responseList, in realm.
6. Return promise.

5.4.3. add(request)

The add(request) method steps are:
1. Let requests be an array containing only request.
2. Let responseArrayPromise be the result of running the algorithm
specified in addAll(requests) passing requests as the argument.
3. Return the result of reacting to responseArrayPromise with a
fulfillment handler that returns undefined.

5.4.4. addAll(requests)

The addAll(requests) method steps are:
1. Let responsePromises be an empty list.
2. Let requestList be an empty list.
3. For each request whose type is Request in requests:
1. Let r be request’s request.
2. If r’s url’s scheme is not one of "http" and "https", or r’s
method is not `GET`, return a promise rejected with a
TypeError.
4. Let fetchControllers be a list of fetch controllers.
5. For each request in requests:
1. Let r be the associated request of the result of invoking the
initial value of Request as constructor with request as its
argument. If this throws an exception, return a promise
rejected with that exception.
2. If r’s url’s scheme is not one of "http" and "https", then:
1. For each fetchController of fetchControllers, abort
fetchController.
2. Return a promise rejected with a TypeError.
3. If r’s client’s global object is a ServiceWorkerGlobalScope
object, set request’s service-workers mode to "none".
4. Add r to requestList.
5. Let responsePromise be a new promise.
6. Run the following substeps in parallel:
o Append the result of fetching r.
o To processResponse for response, run these substeps:
1. If response’s type is "error", or response’s status
is not an ok status or is 206, reject
responsePromise with a TypeError.
2. Else if response’s header list contains a header
named `Vary`, then:
1. Let fieldValues be the list containing the
elements corresponding to the field-values of
the Vary header.
2. For each fieldValue of fieldValues:
1. If fieldValue matches "*", then:
1. Reject responsePromise with a
TypeError.
2. For each fetchController of
fetchControllers, abort
fetchController.
3. Abort these steps.
o To processResponseEndOfBody for response, run these
substeps:
1. If response’s aborted flag is set, reject
responsePromise with an "AbortError" DOMException
and abort these steps.
2. Resolve responsePromise with response.
Note: The cache commit is allowed when the response’s
body is fully received.
7. Add responsePromise to responsePromises.
6. Let p be the result of getting a promise to wait for all of
responsePromises.
7. Return the result of reacting to p with a fulfillment handler that,
when called with argument responses, performs the following
substeps:
1. Let operations be an empty list.
2. Let index be zero.
3. For each response in responses:
1. Let operation be a cache batch operation.
2. Set operation’s type to "put".
3. Set operation’s request to requestList[index].
4. Set operation’s response to response.
5. Append operation to operations.
6. Increment index by one.
4. Let realm be this’s relevant realm.
5. Let cacheJobPromise be a new promise.
6. Run the following substeps in parallel:
1. Let errorData be null.
2. Invoke Batch Cache Operations with operations. If this
throws an exception, set errorData to the exception.
3. Queue a task, on cacheJobPromise’s relevant settings
object’s responsible event loop using the DOM
manipulation task source, to perform the following
substeps:
1. If errorData is null, resolve cacheJobPromise with
undefined.
2. Else, reject cacheJobPromise with a new exception
with errorData, in realm.
7. Return cacheJobPromise.

5.4.5. put(request, response)

The put(request, response) method steps are:
1. Let innerRequest be null.
2. If request is a Request object, then set innerRequest to request’s
request.
3. Else:
1. Let requestObj be the result of invoking Request’s constructor
with request as its argument. If this throws an exception,
return a promise rejected with exception.
2. Set innerRequest to requestObj’s request.
4. If innerRequest’s url’s scheme is not one of "http" and "https", or
innerRequest’s method is not `GET`, return a promise rejected with
a TypeError.
5. Let innerResponse be response’s response.
6. If innerResponse’s status is 206, return a promise rejected with a
TypeError.
7. If innerResponse’s header list contains a header named `Vary`,
then:
1. Let fieldValues be the list containing the items corresponding
to the Vary header’s field-values.
2. For each fieldValue in fieldValues:
1. If fieldValue matches "*", return a promise rejected with
a TypeError.
8. If innerResponse’s body is disturbed or locked, return a promise
rejected with a TypeError.
9. Let clonedResponse be a clone of innerResponse.
10. Let bodyReadPromise be a promise resolved with undefined.
11. If innerResponse’s body is non-null, run these substeps:
1. Let stream be innerResponse’s body’s stream.
2. Let reader be the result of getting a reader for stream.
3. Set bodyReadPromise to the result of reading all bytes from
reader.
Note: This ensures that innerResponse’s body is locked, and we have
a full buffered copy of the body in clonedResponse. An
implementation could optimize by streaming directly to disk rather
than memory.
12. Let operations be an empty list.
13. Let operation be a cache batch operation.
14. Set operation’s type to "put".
15. Set operation’s request to innerRequest.
16. Set operation’s response to clonedResponse.
17. Append operation to operations.
18. Let realm be this’s relevant realm.
19. Return the result of the fulfillment of bodyReadPromise:
1. Let cacheJobPromise be a new promise.
2. Return cacheJobPromise and run these steps in parallel:
1. Let errorData be null.
2. Invoke Batch Cache Operations with operations. If this
throws an exception, set errorData to the exception.
3. Queue a task, on cacheJobPromise’s relevant settings
object’s responsible event loop using the DOM
manipulation task source, to perform the following
substeps:
1. If errorData is null, resolve cacheJobPromise with
undefined.
2. Else, reject cacheJobPromise with a new exception
with errorData, in realm.

5.4.6. delete(request, options)

The delete(request, options) method steps are:
1. Let r be null.
2. If request is a Request object, then:
1. Set r to request’s request.
2. If r’s method is not `GET` and options.ignoreMethod is false,
return a promise resolved with false.
3. Else if request is a string, then:
1. Set r to the associated request of the result of invoking the
initial value of Request as constructor with request as its
argument. If this throws an exception, return a promise
rejected with that exception.
4. Let operations be an empty list.
5. Let operation be a cache batch operation.
6. Set operation’s type to "delete".
7. Set operation’s request to r.
8. Set operation’s options to options.
9. Append operation to operations.
10. Let realm be this’s relevant realm.
11. Let cacheJobPromise be a new promise.
12. Run the following substeps in parallel:
1. Let errorData be null.
2. Let requestResponses be the result of running Batch Cache
Operations with operations. If this throws an exception, set
errorData to the exception.
3. Queue a task, on cacheJobPromise’s relevant settings object’s
responsible event loop using the DOM manipulation task source,
to perform the following substeps:
1. If errorData is null, then:
1. If requestResponses is not empty, resolve
cacheJobPromise with true.
2. Else, resolve cacheJobPromise with false.
2. Else, reject cacheJobPromise with a new exception with
errorData, in realm.
13. Return cacheJobPromise.

5.4.7. keys(request, options)

The keys(request, options) method steps are:
1. Let r be null.
2. If the optional argument request is not omitted, then:
1. If request is a Request object, then:
1. Set r to request’s request.
2. If r’s method is not `GET` and options.ignoreMethod is
false, return a promise resolved with an empty array.
2. Else if request is a string, then:
1. Set r to the associated request of the result of invoking
the initial value of Request as constructor with request
as its argument. If this throws an exception, return a
promise rejected with that exception.
3. Let realm be this’s relevant realm.
4. Let promise be a new promise.
5. Run these substeps in parallel:
1. Let requests be an empty list.
2. If the optional argument request is omitted, then:
1. For each requestResponse of the relevant request response
list:
1. Add requestResponse’s request to requests.
3. Else:
1. Let requestResponses be the result of running Query Cache
with r and options.
2. For each requestResponse of requestResponses:
1. Add requestResponse’s request to requests.
4. Queue a task, on promise’s relevant settings object’s
responsible event loop using the DOM manipulation task source,
to perform the following steps:
1. Let requestList be a list.
2. For each request of requests:
1. Add a new Request object associated with request and
a new associated Headers object whose guard is
"immutable" to requestList.
3. Resolve promise with a frozen array created from
requestList, in realm.
6. Return promise.

5.5. CacheStorage

[SecureContext, Exposed=(Window,Worker)]
interface CacheStorage {
[NewObject] Promise<(Response or undefined)> match(RequestInfo request, option
al MultiCacheQueryOptions options = {});
[NewObject] Promise<boolean> has(DOMString cacheName);
[NewObject] Promise<Cache> open(DOMString cacheName);
[NewObject] Promise<boolean> delete(DOMString cacheName);
[NewObject] Promise<sequence<DOMString>> keys();
};

dictionary MultiCacheQueryOptions : CacheQueryOptions {
DOMString cacheName;
};

Note: CacheStorage interface is designed to largely conform to
ECMAScript 6 Map objects but entirely async, and with additional
convenience methods. The methods, clear, forEach, entries and values,
are intentionally excluded from the scope of the first version
resorting to the ongoing discussion about the async iteration by TC39.

The user agent must create a CacheStorage object when a Window object
or a WorkerGlobalScope object is created.

A CacheStorage object represents the name to cache map associated with
the result of running obtain a local storage bottle map with its
relevant settings object and "caches". Multiple separate objects
implementing the CacheStorage interface across documents and workers
can all be associated with the same name to cache map simultaneously.

5.5.1. match(request, options)

The match(request, options) method steps are:
1. If options["cacheName"] exists, then:
1. Return a new promise promise and run the following substeps in
parallel:
1. For each cacheName → cache of the relevant name to cache
map:
1. If options["cacheName"] matches cacheName, then:
1. Resolve promise with the result of running the
algorithm specified in match(request, options)
method of Cache interface with request and
options (providing cache as thisArgument to the
\[[Call]] internal method of match(request,
options).)
2. Abort these steps.
2. Resolve promise with undefined.
2. Else:
1. Let promise be a promise resolved with undefined.
2. For each cacheName → cache of the relevant name to cache map:
1. Set promise to the result of reacting to itself with a
fulfillment handler that, when called with argument
response, performs the following substeps:
1. If response is not undefined, return response.
2. Return the result of running the algorithm specified
in match(request, options) method of Cache interface
with request and options as the arguments (providing
cache as thisArgument to the \[[Call]] internal
method of match(request, options).)
3. Return promise.

5.5.2. has(cacheName)

The has(cacheName) method steps are:
1. Let promise be a new promise.
2. Run the following substeps in parallel:
1. For each key → value of the relevant name to cache map:
1. If cacheName matches key, resolve promise with true and
abort these steps.
2. Resolve promise with false.
3. Return promise.

5.5.3. open(cacheName)

The open(cacheName) method steps are:
1. Let promise be a new promise.
2. Run the following substeps in parallel:
1. For each key → value of the relevant name to cache map:
1. If cacheName matches key, then:
1. Resolve promise with a new Cache object that
represents value.
2. Abort these steps.
2. Let cache be a new request response list.
3. Set the relevant name to cache map[cacheName] to cache. If
this cache write operation failed due to exceeding the granted
quota limit, reject promise with a QuotaExceededError and
abort these steps.
4. Resolve promise with a new Cache object that represents cache.
3. Return promise.

5.5.4. delete(cacheName)

The delete(cacheName) method steps are:
1. Let promise be the result of running the algorithm specified in
has(cacheName) method with cacheName.
2. Return the result of reacting to promise with a fulfillment handler
that, when called with argument cacheExists, performs the following
substeps:
1. If cacheExists is false, then:
1. Return false.
2. Let cacheJobPromise be a new promise.
3. Run the following substeps in parallel:
1. Remove the relevant name to cache map[cacheName].
2. Resolve cacheJobPromise with true.
Note: After this step, the existing DOM objects (i.e. the
currently referenced Cache, Request, and Response objects)
should remain functional.
4. Return cacheJobPromise.

5.5.5. keys()

The keys() method steps are:
1. Let promise be a new promise.
2. Run the following substeps in parallel:
1. Let cacheKeys be the result of getting the keys of the
relevant name to cache map.
Note: The items in the result ordered set are in the order
that their corresponding entry was added to the name to cache
map.
2. Resolve promise with cacheKeys.
3. Return promise.

6. Security Considerations

6.1. Secure Context

Service workers must execute in secure contexts. Service worker clients
must also be secure contexts to register a service worker registration,
to get access to the service worker registrations and the service
workers, to do messaging with the service workers, and to be
manipulated by the service workers.

Note: This effectively means that service workers and their service
worker clients need to be hosted over HTTPS. A user agent can allow
localhost (see the requirements), 127.0.0.0/8, and ::1/128 for
development purposes. The primary reason for this restriction is to
protect users from the risks associated with insecure contexts.

6.2. Content Security Policy

Whenever a user agent invokes the Run Service Worker algorithm with a
service worker serviceWorker:
* If serviceWorker’s script resource was delivered with a
Content-Security-Policy HTTP header containing the value policy,
the user agent must enforce policy for serviceWorker.
* If serviceWorker’s script resource was delivered with a
Content-Security-Policy-Report-Only HTTP header containing the
value policy, the user agent must monitor policy for serviceWorker.

The primary reason for this restriction is to mitigate a broad class of
content injection vulnerabilities, such as cross-site scripting (XSS).

6.3. Origin Relativity

6.3.1. Origin restriction

This section is non-normative.

A service worker executes in the registering service worker client’s
origin. One of the advanced concerns that major applications would
encounter is whether they can be hosted from a CDN. By definition,
these are servers in other places, often on other origins. Therefore,
service workers cannot be hosted on CDNs. But they can include
resources via importScripts(). The reason for this restriction is that
service workers create the opportunity for a bad actor to turn a bad
day into a bad eternity.

6.3.2. importScripts(urls)

When the importScripts(urls) method is called on a
ServiceWorkerGlobalScope object, the user agent must import scripts
into worker global scope, given this ServiceWorkerGlobalScope object
and urls, and with the following perform the fetch hook steps, given
the request request:
1. Let serviceWorker be request’s client’s global object’s service
worker.
2. Let map be serviceWorker’s script resource map.
3. Let url be request’s url.
4. If serviceWorker’s state is not "parsed" or "installing":
1. Return map[url] if it exists and a network error otherwise.
5. If map[url] exists:
1. Append url to serviceWorker’s set of used scripts.
2. Return map[url].
6. Let registration be serviceWorker’s containing service worker
registration.
7. Set request’s service-workers mode to "none".
8. Set request’s cache mode to "no-cache" if any of the following are
true:
+ registration’s update via cache mode is "none".
+ The current global object’s force bypass cache for import
scripts flag is set.
+ registration is stale.
9. Let response be the result of fetching request.
10. If response’s cache state is not "local", set registration’s last
update check time to the current time.
11. If response’s unsafe response is a bad import script response, then
return a network error.
12. Set map[url] to response.
13. Append url to serviceWorker’s set of used scripts.
14. Set serviceWorker’s classic scripts imported flag.
15. Return response.

6.4. Cross-Origin Resources and CORS

This section is non-normative.

Applications tend to cache items that come from a CDN or other origin.
It is possible to request many of them directly using <script>, <img>,
<video> and <link> elements. It would be hugely limiting if this sort
of runtime collaboration broke when offline. Similarly, it is possible
to fetch many sorts of off-origin resources when appropriate CORS
headers are set. Service workers enable this by allowing Caches to
fetch and cache off-origin items. Some restrictions apply, however.
First, unlike same-origin resources which are managed in the Cache as
Response objects whose corresponding responses are basic filtered
response, the objects stored are Response objects whose corresponding
responses are either CORS filtered responses or opaque filtered
responses. They can be passed to event.respondWith(r) method in the
same manner as the Response objects whose corresponding responses are
basic filtered responses, but cannot be meaningfully created
programmatically. These limitations are necessary to preserve the
security invariants of the platform. Allowing Caches to store them
allows applications to avoid re-architecting in most cases.

6.5. Path restriction

This section is non-normative.

In addition to the origin restriction, service workers are restricted
by the path of the service worker script. For example, a service worker
script at https://www.example.com/~bob/sw.js can be registered for the
scope url https://www.example.com/~bob/ but not for the scope
https://www.example.com/ or https://www.example.com/~alice/. This
provides some protection for sites that host multiple-user content in
separated directories on the same origin. However, the path restriction
is not considered a hard security boundary, as only origins are. Sites
are encouraged to use different origins to securely isolate segments of
the site if appropriate.

Servers can remove the path restriction by setting a
Service-Worker-Allowed header on the service worker script.

6.6. Service worker script request

This section is non-normative.

To further defend against malicious registration of a service worker on
a site, this specification requires that:
* The Service-Worker header is present on service worker script
requests, and
* Service worker scripts are served with a JavaScript MIME type.

6.7. Implementer Concerns

This section is non-normative.

The implementers are encouraged to note:
* Plug-ins should not load via service workers. As plug-ins may get
their security origins from their own urls, the embedding service
worker cannot handle it. For this reason, the Handle Fetch
algorithm makes <embed> and <object> requests immediately fallback
to the network without dispatching fetch event.
* Some of the legacy networking stack code may need to be carefully
audited to understand the ramifications of interactions with
service workers.

6.8. Privacy

Service workers introduce new persistent storage features including
registration map (for service worker registrations and their service
workers), request response list and name to cache map (for caches), and
script resource map (for script resources). In order to protect users
from any potential unsanctioned tracking threat, these persistent
storages should be cleared when users intend to clear them and should
maintain and interoperate with existing user controls e.g. purging all
existing persistent storages.

7. Extensibility

Service Workers specification is extensible from other specifications.

7.1. Define API bound to Service Worker Registration

Specifications may define an API tied to a service worker registration
by using partial interface definition to the ServiceWorkerRegistration
interface where it may define the specification specific attributes and
methods:
partial interface ServiceWorkerRegistration {
// e.g. define an API namespace
readonly attribute APISpaceType APISpace;
// e.g. define a method
Promise<T> methodName(/* list of arguments */);
};

7.2. Define Functional Event

Specifications may define a functional event by extending
ExtendableEvent interface:
// e.g. define FunctionalEvent interface
interface FunctionalEvent : ExtendableEvent {
// add a functional event's own attributes and methods
};

7.3. Define Event Handler

Specifications may define an event handler attribute for the
corresponding functional event using partial interface definition to
the ServiceWorkerGlobalScope interface:
partial interface ServiceWorkerGlobalScope {
attribute EventHandler onfunctionalevent;
};

7.4. Firing Functional Events

To request a functional event dispatch to the active worker of a
service worker registration, specifications should invoke Fire
Functional Event.

Appendix A: Algorithms

The following definitions are the user agent’s internal data structures
used throughout the specification.

A registration map is an ordered map where the keys are (storage key,
serialized scope urls) and the values are service worker registrations.

A job is an abstraction of one of register, update, and unregister
request for a service worker registration.

A job has a job type, which is one of register, update, and unregister.

A job has a storage key (a storage key).

A job has a scope url (a URL).

A job has a script url (a URL).

A job has a worker type ("classic" or "module").

A job has an update via cache mode, which is "imports", "all", or
"none".

A job has a client (a service worker client). It is initially null.

A job has a referrer (a URL or null).

A job has a job promise (a promise). It is initially null.

A job has a containing job queue (a job queue or null). It is initially
null.

A job has a list of equivalent jobs (a list of jobs). It is initially
the empty list.

A job has a force bypass cache flag. It is initially unset.

Two jobs are equivalent when their job type is the same and:
* For register and update jobs, their scope url, script url, worker
type, and update via cache mode are the same.
* For unregister jobs, their scope url is the same.

A job queue is a thread safe queue used to synchronize the set of
concurrent jobs. The job queue contains jobs as its items. A job queue
is initially empty.

A scope to job queue map is an ordered map where the keys are scope
urls, serialized, and the values are job queues.

A bad import script response is a response for which any of the
following conditions are met:
* response’s type is "error"
* response’s status is not an ok status
* The result of extracting a MIME type from response’s header list is
not a JavaScript MIME type
Note: Keep this definition in sync with fetch a classic
worker-imported script.

A race result is a tuple of a routed response and used route.

A race result has an associated routed response (a response).

A race result has an associated used route (a RouterSourceEnum).

Create Job

Input
jobType, a job type

storage key, a storage key

scopeURL, a URL

scriptURL, a URL

promise, a promise

client, a service worker client

Output
job, a job

1. Let job be a new job.
2. Set job’s job type to jobType.
3. Set job’s storage key to storage key.
4. Set job’s scope url to scopeURL.
5. Set job’s script url to scriptURL.
6. Set job’s job promise to promise.
7. Set job’s client to client.
8. If client is not null, set job’s referrer to client’s creation URL.
9. Return job.

Schedule Job

Input
job, a job

Output
none

1. Let jobQueue be null.
2. Let jobScope be job’s scope url, serialized.
3. If scope to job queue map[jobScope] does not exist, set scope to
job queue map[jobScope] to a new job queue.
4. Set jobQueue to scope to job queue map[jobScope].
5. If jobQueue is empty, then:
1. Set job’s containing job queue to jobQueue, and enqueue job to
jobQueue.
2. Invoke Run Job with jobQueue.
6. Else:
1. Let lastJob be the element at the back of jobQueue.
2. If job is equivalent to lastJob and lastJob’s job promise has
not settled, append job to lastJob’s list of equivalent jobs.
3. Else, set job’s containing job queue to jobQueue, and enqueue
job to jobQueue.

Run Job

Input
jobQueue, a job queue

Output
none

1. Assert: jobQueue is not empty.
2. Queue a task to run these steps:
1. Let job be the first item in jobQueue.
2. If job’s job type is register, run Register with job in
parallel.
3. Else if job’s job type is update, run Update with job in
parallel.
Note: For a register job and an update job, the user agent
delays queuing a task for running the job until after a
DOMContentLoaded event has been dispatched to the document
that initiated the job.
4. Else if job’s job type is unregister, run Unregister with job
in parallel.

Finish Job

Input
job, a job

Output
none

1. Let jobQueue be job’s containing job queue.
2. Assert: the first item in jobQueue is job.
3. Dequeue from jobQueue.
4. If jobQueue is not empty, invoke Run Job with jobQueue.

Resolve Job Promise

Input
job, a job

value, any

Output
none

1. If job’s client is not null, queue a task, on job’s client’s
responsible event loop using the DOM manipulation task source, to
run the following substeps:
1. Let convertedValue be null.
2. If job’s job type is either register or update, set
convertedValue to the result of getting the service worker
registration object that represents value in job’s client.
3. Else, set convertedValue to value, in job’s client’s Realm.
4. Resolve job’s job promise with convertedValue.
2. For each equivalentJob in job’s list of equivalent jobs:
1. If equivalentJob’s client is null, continue to the next
iteration of the loop.
2. Queue a task, on equivalentJob’s client’s responsible event
loop using the DOM manipulation task source, to run the
following substeps:
1. Let convertedValue be null.
2. If equivalentJob’s job type is either register or update,
set convertedValue to the result of getting the service
worker registration object that represents value in
equivalentJob’s client.
3. Else, set convertedValue to value, in equivalentJob’s
client’s Realm.
4. Resolve equivalentJob’s job promise with convertedValue.

Reject Job Promise

Input
job, a job

errorData, the information necessary to create an exception

Output
none

1. If job’s client is not null, queue a task, on job’s client’s
responsible event loop using the DOM manipulation task source, to
reject job’s job promise with a new exception with errorData, in
job’s client’s Realm.
2. For each equivalentJob in job’s list of equivalent jobs:
1. If equivalentJob’s client is null, continue.
2. Queue a task, on equivalentJob’s client’s responsible event
loop using the DOM manipulation task source, to reject
equivalentJob’s job promise with a new exception with
errorData, in equivalentJob’s client’s Realm.

Start Register

Input
scopeURL, a URL or failure or null

scriptURL, a URL or failure

promise, a promise

client, a service worker client

referrer, a URL

workerType, a worker type

updateViaCache, an update via cache mode

Output
none

1. If scriptURL is failure, reject promise with a TypeError and abort
these steps.
2. Set scriptURL’s fragment to null.
Note: The user agent does not store the fragment of the script’s
url. This means that the fragment does not have an effect on
identifying service workers.
3. If scriptURL’s scheme is not one of "http" and "https", reject
promise with a TypeError and abort these steps.
4. If any of the strings in scriptURL’s path contains either ASCII
case-insensitive "%2f" or ASCII case-insensitive "%5c", reject
promise with a TypeError and abort these steps.
5. If scopeURL is null, set scopeURL to the result of parsing the
string "./" with scriptURL.
Note: The scope url for the registration is set to the location of
the service worker script by default.
6. If scopeURL is failure, reject promise with a TypeError and abort
these steps.
7. Set scopeURL’s fragment to null.
Note: The user agent does not store the fragment of the scope url.
This means that the fragment does not have an effect on identifying
service worker registrations.
8. If scopeURL’s scheme is not one of "http" and "https", reject
promise with a TypeError and abort these steps.
9. If any of the strings in scopeURL’s path contains either ASCII
case-insensitive "%2f" or ASCII case-insensitive "%5c", reject
promise with a TypeError and abort these steps.
10. Let storage key be the result of running obtain a storage key given
client.
11. Let job be the result of running Create Job with register, storage
key, scopeURL, scriptURL, promise, and client.
12. Set job’s worker type to workerType.
13. Set job’s update via cache mode to updateViaCache.
14. Set job’s referrer to referrer.
15. Invoke Schedule Job with job.

Register

Input
job, a job

Output
none

1. If the result of running potentially trustworthy origin with the
origin of job’s script url as the argument is Not Trusted, then:
1. Invoke Reject Job Promise with job and "SecurityError"
DOMException.
2. Invoke Finish Job with job and abort these steps.
2. If job’s script url’s origin and job’s referrer’s origin are not
same origin, then:
1. Invoke Reject Job Promise with job and "SecurityError"
DOMException.
2. Invoke Finish Job with job and abort these steps.
3. If job’s scope url’s origin and job’s referrer’s origin are not
same origin, then:
1. Invoke Reject Job Promise with job and "SecurityError"
DOMException.
2. Invoke Finish Job with job and abort these steps.
4. Let registration be the result of running Get Registration given
job’s storage key and job’s scope url.
5. If registration is not null, then:
1. Let newestWorker be the result of running the Get Newest
Worker algorithm passing registration as the argument.
2. If newestWorker is not null, job’s script url equals
newestWorker’s script url, job’s worker type equals
newestWorker’s type, and job’s update via cache mode’s value
equals registration’s update via cache mode, then:
1. Invoke Resolve Job Promise with job and registration.
2. Invoke Finish Job with job and abort these steps.
6. Else:
1. Invoke Set Registration algorithm with job’s storage key,
job’s scope url, and job’s update via cache mode.
7. Invoke Update algorithm passing job as the argument.

Update

Input
job, a job

Output
none

1. Let registration be the result of running Get Registration given
job’s storage key and job’s scope url.
2. If registration is null, then:
1. Invoke Reject Job Promise with job and TypeError.
2. Invoke Finish Job with job and abort these steps.
3. Let newestWorker be the result of running Get Newest Worker
algorithm passing registration as the argument.
4. If job’s job type is update, and newestWorker is not null and its
script url does not equal job’s script url, then:
1. Invoke Reject Job Promise with job and TypeError.
2. Invoke Finish Job with job and abort these steps.
5. Let hasUpdatedResources be false.
6. Let updatedResourceMap be an ordered map where the keys are URLs
and the values are responses.
7. Switching on job’s worker type, run these substeps with the
following options:

"classic"
Fetch a classic worker script given job’s serialized
script url, job’s client, "serviceworker", and the
to-be-created environment settings object for this service
worker.

"module"
Fetch a module worker script graph given job’s serialized
script url, job’s client, "serviceworker", "omit", and the
to-be-created environment settings object for this service
worker.

Using the to-be-created environment settings object rather than a
concrete environment settings object. This is used due to the
unique processing model of service workers compared to the
processing model of other web workers. The script fetching
algorithms of HTML standard originally designed for other web
workers require an environment settings object of the execution
environment, but service workers fetch a script separately in the
Update algorithm before the script later runs multiple times
through the Run Service Worker algorithm.
The fetch a classic worker script algorithm and the fetch a module
worker script graph algorithm in HTML take job’s client as an
argument. job’s client is null when passed from the Soft Update
algorithm.
To perform the fetch hook given request, run the following steps:
1. Append `Service-Worker`/`script` to request’s header list.
Note: See the definition of the Service-Worker header in
Appendix B: Extended HTTP headers.
2. Set request’s cache mode to "no-cache" if any of the following
are true:
o registration’s update via cache mode is not "all".
o job’s force bypass cache flag is set.
o newestWorker is not null and registration is stale.
Note: Even if the cache mode is not set to "no-cache", the
user agent obeys Cache-Control header’s max-age value in the
network layer to determine if it should bypass the browser
cache.
3. Set request’s service-workers mode to "none".
4. If the isTopLevel flag is unset, then return the result of
fetching request.
5. Set request’s redirect mode to "error".
6. Fetch request, and asynchronously wait to run the remaining
steps as part of fetch’s processResponse for the response
response.
7. Extract a MIME type from the response’s header list. If this
MIME type (ignoring parameters) is not a JavaScript MIME type,
then:
1. Invoke Reject Job Promise with job and "SecurityError"
DOMException.
2. Asynchronously complete these steps with a network error.
8. Let serviceWorkerAllowed be the result of extracting header
list values given `Service-Worker-Allowed` and response’s
header list.
Note: See the definition of the Service-Worker-Allowed header
in Appendix B: Extended HTTP headers.
9. Set policyContainer to the result of creating a policy
container from a fetch response given response.
10. If serviceWorkerAllowed is failure, then:
1. Asynchronously complete these steps with a network error.
11. Let scopeURL be registration’s scope url.
12. Let maxScopeString be null.
13. If serviceWorkerAllowed is null, then:
1. Let resolvedScope be the result of parsing "./" using
job’s script url as the base URL.
2. Set maxScopeString to "/", followed by the strings in
resolvedScope’s path (including empty strings), separated
from each other by "/".
Note: The final item in resolvedScope’s path will always
be an empty string, so maxScopeString will have a
trailing "/".
14. Else:
1. Let maxScope be the result of parsing
serviceWorkerAllowed using job’s script url as the base
URL.
2. If maxScope’s origin is job’s script url’s origin, then:
1. Set maxScopeString to "/", followed by the strings
in maxScope’s path (including empty strings),
separated from each other by "/".
15. Let scopeString be "/", followed by the strings in scopeURL’s
path (including empty strings), separated from each other by
"/".
16. If maxScopeString is null or scopeString does not start with
maxScopeString, then:
1. Invoke Reject Job Promise with job and "SecurityError"
DOMException.
2. Asynchronously complete these steps with a network error.
17. Let url be request’s url.
18. Set updatedResourceMap[url] to response.
19. If response’s cache state is not "local", set registration’s
last update check time to the current time.
20. Set hasUpdatedResources to true if any of the following are
true:
o newestWorker is null.
o newestWorker’s script url is not url or newestWorker’s
type is not job’s worker type.
o newestWorker’s script resource map[url]'s body is not
byte-for-byte identical with response’s body.
21. If hasUpdatedResources is false and newestWorker’s classic
scripts imported flag is set, then:
Note: The following checks to see if an imported script has
been updated, since the main script has not changed.
1. For each importUrl → storedResponse of newestWorker’s
script resource map:
1. If importUrl is url, then continue.
2. Let importRequest be a new request whose url is
importUrl, client is job’s client, destination is
"script", parser metadata is "not parser-inserted",
and whose use-URL-credentials flag is set.
3. Set importRequest’s cache mode to "no-cache" if any
of the following are true:
@ registration’s update via cache mode is "none".
@ job’s force bypass cache flag is set.
@ registration is stale.
4. Let fetchedResponse be the result of fetching
importRequest.
5. Set updatedResourceMap[importRequest’s url] to
fetchedResponse.
6. Set fetchedResponse to fetchedResponse’s unsafe
response.
7. If fetchedResponse’s cache state is not "local", set
registration’s last update check time to the current
time.
8. If fetchedResponse is a bad import script response,
continue.
Note: Bad responses for importScripts() are ignored
for the purpose of the byte-to-byte check. Only good
responses for the incumbent worker and good
responses for the potential update worker are
considered. See issue #1374 for some rationale.
9. If fetchedResponse’s body is not byte-for-byte
identical with storedResponse’s unsafe response’s
body, set hasUpdatedResources to true.
Note: The control does not break the loop in this
step to continue with all the imported scripts to
populate the cache.
22. Asynchronously complete these steps with response.
When the algorithm asynchronously completes, continue the rest of
these steps, with script being the asynchronous completion value.
8. If script is null or Is Async Module with script’s record, script’s
base URL, and « » is true, then:
1. Invoke Reject Job Promise with job and TypeError.
Note: This will do nothing if Reject Job Promise was
previously invoked with "SecurityError" DOMException.
2. If newestWorker is null, then remove registration
map[(registration’s storage key, serialized scopeURL)].
3. Invoke Finish Job with job and abort these steps.
9. If hasUpdatedResources is false, then:
1. Set registration’s update via cache mode to job’s update via
cache mode.
2. Invoke Resolve Job Promise with job and registration.
3. Invoke Finish Job with job and abort these steps.
10. Let worker be a new service worker.
11. Set worker’s script url to job’s script url, worker’s script
resource to script, worker’s type to job’s worker type, and
worker’s script resource map to updatedResourceMap.
12. Append url to worker’s set of used scripts.
13. Set worker’s script resource’s policy container to policyContainer.
14. Let forceBypassCache be true if job’s force bypass cache flag is
set, and false otherwise.
15. Let runResult be the result of running the Run Service Worker
algorithm with worker and forceBypassCache.
16. If runResult is failure or an abrupt completion, then:
1. Invoke Reject Job Promise with job and TypeError.
2. If newestWorker is null, then remove registration
map[(registration’s storage key, serialized scopeURL)].
3. Invoke Finish Job with job.
17. Else, invoke Install algorithm with job, worker, and registration
as its arguments.

Soft Update

The user agent may call this as often as it likes to check for updates.

Input
registration, a service worker registration

forceBypassCache, an optional boolean, false by default

Note: Implementers may use forceBypassCache to aid debugging
(e.g. invocations from developer tools), and other
specifications that extend service workers may also use the flag
on their own needs.

Output
None

1. Let newestWorker be the result of running Get Newest Worker
algorithm passing registration as its argument.
2. If newestWorker is null, abort these steps.
3. Let job be the result of running Create Job with update,
registration’s storage key, registration’s scope url,
newestWorker’s script url, null, and null.
4. Set job’s worker type to newestWorker’s type.
5. Set job’s force bypass cache flag if forceBypassCache is true.
6. Invoke Schedule Job with job.

Install

Input
job, a job

worker, a service worker

registration, a service worker registration

Output
none

1. Let installFailed be false.
2. Let newestWorker be the result of running Get Newest Worker
algorithm passing registration as its argument.
3. Set registration’s update via cache mode to job’s update via cache
mode.
4. Run the Update Registration State algorithm passing registration,
"installing" and worker as the arguments.
5. Run the Update Worker State algorithm passing registration’s
installing worker and "installing" as the arguments.
6. Assert: job’s job promise is not null.
7. Invoke Resolve Job Promise with job and registration.
8. Let settingsObjects be all environment settings objects whose
origin is registration’s scope url’s origin.
9. For each settingsObject of settingsObjects, queue a task on
settingsObject’s responsible event loop in the DOM manipulation
task source to run the following steps:
1. Let registrationObjects be every ServiceWorkerRegistration
object in settingsObject’s realm, whose service worker
registration is registration.
2. For each registrationObject of registrationObjects, fire an
event on registrationObject named updatefound.
10. Let installingWorker be registration’s installing worker.
11. If the result of running the Should Skip Event algorithm with
installingWorker and "install" is false, then:
1. Let forceBypassCache be true if job’s force bypass cache flag
is set, and false otherwise.
2. If the result of running the Run Service Worker algorithm with
installingWorker and forceBypassCache is failure, then:
1. Set installFailed to true.
3. Else:
1. Queue a task task on installingWorker’s event loop using
the DOM manipulation task source to run the following
steps:
1. Let e be the result of creating an event with
InstallEvent.
2. Initialize e’s type attribute to install.
3. Dispatch e at installingWorker’s global object.
4. WaitForAsynchronousExtensions: Run the following
substeps in parallel:
1. Wait until e is not active.
2. If e’s timed out flag is set, set installFailed
to true.
3. Let p be the result of getting a promise to
wait for all of e’s extend lifetime promises.
4. Upon rejection of p, set installFailed to true.
If task is discarded, set installFailed to true.
2. Wait for task to have executed or been discarded.
3. Wait for the step labeled WaitForAsynchronousExtensions
to complete.
12. If installFailed is true, then:
1. Run the Update Worker State algorithm passing registration’s
installing worker and "redundant" as the arguments.
2. Run the Update Registration State algorithm passing
registration, "installing" and null as the arguments.
3. If newestWorker is null, then remove registration
map[(registration’s storage key, serialized registration’s
scope url)].
4. Invoke Finish Job with job and abort these steps.
13. Let map be registration’s installing worker’s script resource map.
14. Let usedSet be registration’s installing worker’s set of used
scripts.
15. For each url of map:
1. If usedSet does not contain url, then remove map[url].
16. If registration’s waiting worker is not null, then:
1. Terminate registration’s waiting worker.
2. Run the Update Worker State algorithm passing registration’s
waiting worker and "redundant" as the arguments.
17. Run the Update Registration State algorithm passing registration,
"waiting" and registration’s installing worker as the arguments.
18. Run the Update Registration State algorithm passing registration,
"installing" and null as the arguments.
19. Run the Update Worker State algorithm passing registration’s
waiting worker and "installed" as the arguments.
20. Invoke Finish Job with job.
21. Wait for all the tasks queued by Update Worker State invoked in
this algorithm to have executed.
22. Invoke Try Activate with registration.
Note: If Try Activate does not trigger Activate here, Activate is
tried again when the last client controlled by the existing active
worker is unloaded, skipWaiting() is asynchronously called, or the
extend lifetime promises for the existing active worker settle.

Activate

Input
registration, a service worker registration

Output
None

1. If registration’s waiting worker is null, abort these steps.
2. If registration’s active worker is not null, then:
1. Terminate registration’s active worker.
2. Run the Update Worker State algorithm passing registration’s
active worker and "redundant" as the arguments.
3. Run the Update Registration State algorithm passing registration,
"active" and registration’s waiting worker as the arguments.
4. Run the Update Registration State algorithm passing registration,
"waiting" and null as the arguments.
5. Run the Update Worker State algorithm passing registration’s active
worker and "activating" as the arguments.
Note: Once an active worker is activating, neither a runtime script
error nor a force termination of the active worker prevents the
active worker from getting activated.
Note: Make sure to design activation handlers to do non-essential
work (like cleanup). This is because activation handlers may not
all run to completion, especially in the case of browser
termination during activation. A Service Worker should be designed
to function properly, even if the activation handlers do not all
complete successfully.
6. Let matchedClients be a list of service worker clients whose
creation URL matches registration’s storage key and registration’s
scope url.
7. For each client of matchedClients, queue a task on client’s
responsible event loop, using the DOM manipulation task source, to
run the following substeps:
1. Let readyPromise be client’s global object’s
ServiceWorkerContainer object’s ready promise.
2. If readyPromise is null, then continue.
3. If readyPromise is pending, resolve readyPromise with the the
result of getting the service worker registration object that
represents registration in readyPromise’s relevant settings
object.
8. For each service worker client client who is using registration:
1. Set client’s active worker to registration’s active worker.
2. Invoke Notify Controller Change algorithm with client as the
argument.
9. Let activeWorker be registration’s active worker.
10. If the result of running the Should Skip Event algorithm with
activeWorker and "activate" is false, then:
1. If the result of running the Run Service Worker algorithm with
activeWorker is not failure, then:
1. Queue a task task on activeWorker’s event loop using the
DOM manipulation task source to run the following steps:
1. Let e be the result of creating an event with
ExtendableEvent.
2. Initialize e’s type attribute to activate.
3. Dispatch e at activeWorker’s global object.
4. WaitForAsynchronousExtensions: Wait, in parallel,
until e is not active.
2. Wait for task to have executed or been discarded.
3. Wait for the step labeled WaitForAsynchronousExtensions
to complete.
11. Run the Update Worker State algorithm passing registration’s active
worker and "activated" as the arguments.

Try Activate

Input
registration, a service worker registration

Output
None

1. If registration’s waiting worker is null, return.
2. If registration’s active worker is not null and registration’s
active worker’s state is "activating", return.
Note: If the existing active worker is still in activating state,
the activation of the waiting worker is delayed.
3. Invoke Activate with registration if either of the following is
true:
+ registration’s active worker is null.
+ The result of running Service Worker Has No Pending Events
with registration’s active worker is true, and no service
worker client is using registration or registration’s waiting
worker’s skip waiting flag is set.

Setup ServiceWorkerGlobalScope

Input
serviceWorker, a service worker

Output
a ServiceWorkerGlobalScope object or null

Note: This algorithm returns a ServiceWorkerGlobalScope usable for a
CSP check, or null. If serviceWorker has an active
ServiceWorkerGlobalScope, then it is returned. Otherwise, the object
will be newly created.

This algorithm does the minimal setup for the service worker that is
necessary to create something usable for security checks like CSP and
COEP. This specification ensures that this algorithm is called before
any such checks are performed.

In specifications, such security checks require creating a
ServiceWorkerGlobalScope, a relevant settings object, a realm, and an
agent. In implementations, the amount of work required might be much
less. Therefore, implementations could do less work in their equivalent
of this algorithm, and more work in Run Service Worker, as long as the
results are observably equivalent. (And in particular, as long as all
security checks have the same result.)

1. Let unsafeCreationTime be the unsafe shared current time.
2. If serviceWorker is running, then return serviceWorker’s global
object.
3. If serviceWorker’s state is "redundant", then return null.
4. If serviceWorker’s global object is not null, then return
serviceWorker’s global object.
5. Assert: serviceWorker’s start status is null.
6. Let setupFailed be false.
7. Let globalObject be null.
8. Let agent be the result of obtaining a service worker agent, and
run the following steps in that context:
1. Let realmExecutionContext be the result of creating a new
realm given agent and the following customizations:
o For the global object, create a new
ServiceWorkerGlobalScope object. Let workerGlobalScope be
the created object.
2. Let settingsObject be a new environment settings object whose
algorithms are defined as follows:

The realm execution context
Return realmExecutionContext.

The module map
Return workerGlobalScope’s module map.

The API base URL
Return serviceWorker’s script url.

The origin
Return its registering service worker client’s
origin.

The has cross-site ancestor
Return its registering service worker client’s has
cross-site ancestor.

The policy container
Return workerGlobalScope’s policy container.

The time origin
Return the result of coarsening unsafeCreationTime
given workerGlobalScope’s cross-origin isolated
capability.

3. Set settingsObject’s id to a new unique opaque string,
creation URL to serviceWorker’s script url, top-level creation
URL to null, top-level origin to an implementation-defined
value, target browsing context to null, and active service
worker to null.
4. Set workerGlobalScope’s url to serviceWorker’s script url.
5. Set workerGlobalScope’s policy container to serviceWorker’s
script resource’s policy container.
6. Set workerGlobalScope’s type to serviceWorker’s type.
7. Create a new WorkerLocation object and associate it with
workerGlobalScope.
8. If the run CSP initialization for a global object algorithm
returns "Blocked" when executed upon workerGlobalScope, set
setupFailed to true and abort these steps.
9. Set globalObject to workerGlobalScope.
9. Wait for globalObject is not null, or for setupFailed to be true.
10. If setupFailed is true, then return null.
11. Return globalObject.

Run Service Worker

Input
serviceWorker, a service worker

forceBypassCache, an optional boolean, false by default

Output
a Completion or failure

Note: This algorithm blocks until the service worker is running or
fails to start.
1. If serviceWorker is running, then return serviceWorker’s start
status.
2. If serviceWorker’s state is "redundant", then return failure.
3. Assert: serviceWorker’s start status is null.
4. Let script be serviceWorker’s script resource.
5. Assert: script is not null.
6. Let startFailed be false.
7. Let workerGlobalScope be serviceWorker’s global object.
8. If workerGlobalScope is null:
1. Set workerGlobalScope to be the result of running the Setup
ServiceWorkerGlobalScope algorithm with serviceWorker.
2. If workerGlobalScope is null, then return failure.
3. Set serviceWorker’s global object to workerGlobalScope.
9. Obtain agent for workerGlobalScope’s realm execution context, and
run the following steps in that context:
1. Set workerGlobalScope’s force bypass cache for import scripts
flag if forceBypassCache is true.
2. If serviceWorker is an active worker, and there are any tasks
queued in serviceWorker’s containing service worker
registration’s task queues, queue them to serviceWorker’s
event loop’s task queues in the same order using their
original task sources.
3. Let evaluationStatus be null.
4. If script is a classic script, then:
1. Set evaluationStatus to the result of running the classic
script script.
2. If evaluationStatus.[[Value]] is empty, this means the
script was not evaluated. Set startFailed to true and
abort these steps.
5. Otherwise, if script is a module script, then:
1. Let evaluationPromise be the result of running the module
script script, with report errors set to false.
2. Assert: evaluationPromise.[[PromiseState]] is not
"pending".
3. If evaluationPromise.[[PromiseState]] is "rejected":
1. Set evaluationStatus to
ThrowCompletion(evaluationPromise.[[PromiseResult]])
.
4. Otherwise:
1. Set evaluationStatus to NormalCompletion(undefined).
6. If the script was aborted by the Terminate Service Worker
algorithm, set startFailed to true and abort these steps.
7. Set serviceWorker’s start status to evaluationStatus.
8. If script’s has ever been evaluated flag is unset, then:
1. For each eventType of settingsObject’s global object’s
associated list of event listeners' event types:
1. Append eventType to workerGlobalScope’s associated
service worker’s set of event types to handle.
Note: If the global object’s associated list of event
listeners does not have any event listener added at this
moment, the service worker’s set of event types to handle
remains an empty set.
2. Set script’s has ever been evaluated flag.
3. Unset the serviceWorker’s all fetch listeners are empty
flag.
4. The user agent may, if the All Fetch Listeners Are Empty
algorithm with workerGlobalScope returns true, set
serviceWorker’s all fetch listeners are empty flag.
9. Run the responsible event loop specified by settingsObject
until it is destroyed.
10. Clear workerGlobalScope’s map of active timers.
10. Wait for serviceWorker to be running, or for startFailed to be
true.
11. If startFailed is true, then return failure.
12. Return serviceWorker’s start status.

All Fetch Listeners Are Empty

Input
workerGlobalScope, a global object.

Output
a boolean

1. If workerGlobalScope’s set of event types to handle does not
contain fetch, then return true.
2. Let eventHandler be workerGlobalScope’s event handler
map["onfetch"]'s value.
3. Let eventListenerCallbacks be the result of calling legacy-obtain
service worker fetch event listener callbacks given
workerGlobalScope.
4. For each eventListenerCallback of eventListenerCallbacks:
1. Let callback be null.
2. If eventHandler is not null and eventListenerCallback equals
eventHandler’s listener’s callback, then set callback to the
result of converting to an ECMAScript value eventHandler’s
value.
3. Otherwise, set callback to the result of converting to an
ECMAScript value eventListenerCallback.
4. If IsCallable(callback) is false, then return false.
Note: Callback objects that use handleEvent(event) are assumed
to be non-empty. This avoids calling the handleEvent(event)
getters, which could modify the event listeners during this
check.
5. If callback’s function body is not empty (i.e. either a
statement or declaration exist), then return false.
Note: This detects "fetch" listeners like () => {}. Some sites have
a fetch event listener with empty body to make them recognized by
Chromium as a progressive web application (PWA).
5. Return true.

Note: User agents are encouraged to show a warning indicating that
empty "fetch" listeners are unnecessary, and may have a negative
performance impact.

Terminate Service Worker

Input
serviceWorker, a service worker

Output
None

1. Run the following steps in parallel with serviceWorker’s main loop:
1. Let serviceWorkerGlobalScope be serviceWorker’s global object.
2. Set serviceWorkerGlobalScope’s closing flag to true.
3. Remove all the items from serviceWorker’s set of extended
events.
4. If there are any tasks, whose task source is either the handle
fetch task source or the handle functional event task source,
queued in serviceWorkerGlobalScope’s event loop’s task queues,
queue them to serviceWorker’s containing service worker
registration’s corresponding task queues in the same order
using their original task sources, and discard all the tasks
(including tasks whose task source is neither the handle fetch
task source nor the handle functional event task source) from
serviceWorkerGlobalScope’s event loop’s task queues without
processing them.
Note: This effectively means that the fetch events and the
other functional events such as push events are backed up by
the registration’s task queues while the other tasks including
message events are discarded.
5. Abort the script currently running in serviceWorker.
6. Set serviceWorker’s start status to null.

Handle Fetch

The Handle Fetch algorithm is the entry point for the fetch handling
handed to the service worker context.

Input
request, a request

fetchController, a fetch controller

useHighResPerformanceTimers, a boolean

Output
a response or service worker timing info

1. Let registration be null.
2. Let client be request’s client.
3. Let reservedClient be request’s reserved client.
4. Let preloadResponse be a new promise.
5. Let workerRealm be null.
6. Let timingInfo be a new service worker timing info.
7. Assert: request’s destination is not "serviceworker".
8. If request’s destination is either "embed" or "object", then:
1. Return null.
9. Else if request is a non-subresource request, then:
1. If reservedClient is not null and is an environment settings
object, then:
1. If reservedClient is not a secure context, return null.
2. Else:
1. If request’s url is not a potentially trustworthy URL,
return null.
3. If request is a navigation request and the navigation
triggering it was initiated with a shift+reload or equivalent,
return null.
4. Assert reservedClient is not null.
5. Let storage key be the result of running obtain a storage key
given reservedClient.
6. Set registration to the result of running Match Service Worker
Registration given storage key and request’s url.
7. If registration is null or registration’s active worker is
null, return null.
8. If request’s destination is not "report", set reservedClient’s
active service worker to registration’s active worker.
Note: From this point, the service worker client starts to use its
active service worker’s containing service worker registration.
10. Else if request is a subresource request, then:
1. If client’s active service worker is non-null, set
registration to client’s active service worker’s containing
service worker registration.
2. Else, return null.
11. Let activeWorker be registration’s active worker.
12. Let shouldSoftUpdate be true if any of the following are true, and
false otherwise:
+ request is a non-subresource request.
+ request is a subresource request and registration is stale.
13. If activeWorker’s list of router rules is not empty:
1. Set timingInfo’s worker router evaluation start to the
coarsened shared current time given
useHighResPerformanceTimers.
2. Let source be the result of running the Get Router Source
algorithm with registration’s active worker and request.
3. If source is non-null, then:
1. Set timingInfo’s worker matched router source be set to
source, and worker final router source be set to
"network".
2. If source is "network":
1. If shouldSoftUpdate is true, then in parallel run
the Soft Update algorithm with registration.
2. Return timingInfo.
3. Else if source is "cache", or source["cacheName"] exists,
then:
1. If shouldSoftUpdate is true, then in parallel run
the Soft Update algorithm with registration.
2. Set timingInfo’s worker cache lookup start to the
coarsened shared current time given
useHighResPerformanceTimers.
3. Let environment be null.
4. If request is a non-subresource request, then:
1. Set environment to reservedClient.
5. Else:
1. Set environment to client.
6. Let caches be the result of running obtain a local
storage bottle map with environment and "caches".
7. For each cacheName → cache of caches.
1. If source["cacheName"] exists and
source["cacheName"] is not cacheName, continue.
2. Let requestResponses be the result of running
Query Cache with request, a new
CacheQueryOptions, and cache.
3. If requestResponses is an empty list, return
timingInfo.
4. Else:
1. Let requestResponse be the first element
of requestResponses.
2. Let response be requestResponse’s
response.
3. Let globalObject be activeWorker’s global
object.
4. If globalObject is null:
1. Set globalObject to the result of
running Setup
ServiceWorkerGlobalScope with
activeWorker.
5. If globalObject is null, return
timingInfo.
Note: This only creates a
ServiceWorkerGlobalScope because CORS checks
require that. It is not expected that
implementations will actually create a
ServiceWorkerGlobalScope here.
1. If response’s type is "opaque", and
cross-origin resource policy check with
globalObject’s origin, globalObject, "",
and response’s internal response returns
blocked, then return timingInfo.
2. Set timingInfo’s worker final router
source be set to "cache".
3. Set result’s service worker timing info be
set to timingInfo.
4. Return response.
8. Return timingInfo.
4. Else if source is "race-network-and-fetch-handler", and
request’s method is `GET` then:
1. If shouldSoftUpdate is true, then in parallel run
the Soft Update algorithm with registration.
2. Let queue be an empty queue of race result.
3. Let raceFetchController be null.
4. Let raceResponse be a race response whose value is
"pending".
5. Run the following substeps in parallel:
1. If fetchController’s state is "terminated" or
"aborted", then set raceResponse to a race
response whose value is null, and abort these
steps.
2. Set raceFetchController to the result of
calling fetch given request, with
processResponse set to the following steps
given a response raceNetworkRequestResponse:
1. Run these steps, but abort when
fetchController’s state is "terminated" or
"aborted".
1. Set raceResponse’s value to
raceNetworkRequestResponse.
2. If raceNetworkRequestResponse’s
status is ok status, then:
1. Let raceNetworkResult be a race
result whose routed response is
raceNetworkRequestResponse and used
route is "network".
2. Enqueue raceNetworkResult to queue.
2. If aborted and raceFetchController is not
null, then:
1. Abort raceFetchController.
2. Set raceResponse to a race response
whose value is null.
6. Resolve preloadResponse with undefined.
7. Run the following substeps in parallel:
1. Let fetchHandlerResponse be the result of
Create Fetch Event and Dispatch with request,
registration, useHighResPerformanceTimers,
timingInfo, workerRealm, reservedClient,
preloadResponse, and raceResponse.
2. If fetchHandlerResponse is not null and not a
network error, and raceFetchController is not
null, abort raceFetchController.
3. Let raceFetchHandlerResult be a race result
whose routed response is fetchHandlerResponse
and used route is "fetch-event".
4. Enqueue raceFetchHandlerResult to queue.
8. Wait until queue is not empty.
9. Let result be the result of dequeue queue.
10. Let routedResponse be result’s routed response.
11. If routedResponse is null:
1. Return timingInfo.
12. If result’s used route is "network", then:
1. Set routedResponse’s service worker timing info
be set to timingInfo.
13. Set routedResponse’s service worker timing info’s
worker final router source be set to result’s used
route.
14. Return routedResponse.
5. Assert: source is "fetch-event"
14. Let responseForAutoPreload be null.
15. If request is a navigation request, request’s method is `GET`,
registration’s active worker’s set of event types to handle
contains fetch, and registration’s active worker’s all fetch
listeners are empty flag is not set then:
1. If registration’s navigation preload enabled flag is set then:
Note: If the above is true except registration’s active
worker’s set of event types to handle does not contain fetch,
then the user agent may wish to show a console warning, as the
developer’s intent isn’t clear.
1. Let preloadRequest be the result of cloning the request
request.
2. Let preloadRequestHeaders be preloadRequest’s header
list.
3. Let preloadResponseObject be a new Response object
associated with a new Headers object whose guard is
"immutable".
4. Append to preloadRequestHeaders a new header whose name
is `Service-Worker-Navigation-Preload` and value is
registration’s navigation preload header value.
5. Set preloadRequest’s service-workers mode to "none".
6. Let preloadFetchController be null.
7. Run the following substeps in parallel, but abort when
fetchController’s state is "terminated" or "aborted":
1. Set preloadFetchController to the result of fetching
preloadRequest.
To processResponse for navigationPreloadResponse,
run these substeps:
1. If navigationPreloadResponse’s type is "error",
reject preloadResponse with a TypeError and
terminate these substeps.
2. Associate preloadResponseObject with
navigationPreloadResponse.
3. Resolve preloadResponse with
preloadResponseObject.
8. If aborted, then:
1. Let deserializedError be the result of deserialize a
serialized abort reason given null and workerRealm.
2. Abort preloadFetchController with deserializedError.
2. Else if timingInfo’s worker matched router source is not
"fetch-event", a user agent may run the following substeps:
Note: A user agent may speculatively dispatch a network
request in parallel with creating a fetch event in order to
minimize the bootstrap cost.
1. Assert: timingInfo’s worker matched router source is
null.
2. Let autoPreloadFetchController be null.
3. Set responseForAutoPreload to be a race response whose
value is "pending".
4. Run the following substeps in parallel, but abort when
fetchController’s state is "terminated" or "aborted":
1. Set autoPreloadFetchController to the result of
calling fetch given request, with processResponse
set to the following steps given a response
autoPreloadRequestResponse:
1. Set responseForAutoPreload’s value to
autoPreloadRequestResponse.
5. If aborted and autoPreloadFetchController is not null,
then:
1. Abort autoPreloadFetchController.
2. Set responseForAutoPreload to a race response whose
value is null.
6. Resolve preloadResponse with undefined.
16. Else, resolve preloadResponse with undefined.
17. Let fetchResult be the result of Create Fetch Event and Dispatch
with request, registration, useHighResPerformanceTimers,
timingInfo, workerRealm, reservedClient, preloadResponse, and
responseForAutoPreload.
18. If timingInfo’s worker final router source is not an empty string:
1. Assert timingInfo’s worker final router source is "network".
2. If fetchResult is null, then return timingInfo.
3. Else:
1. Assert fetchResult’s service worker timing info’s worker
final router source is "network".
2. Set fetchResult’s service worker timing info’s worker
final router source to "fetch-event".
19. Return fetchResult.

Create Fetch Event and Dispatch

Input
request, a request

registration, a service worker registration

useHighResPerformanceTimers, a boolean

timingInfo, a service worker timing info

workerRealm, a relevant realm of the global object

reservedClient, a reserved client

preloadResponse, a promise

raceResponse, a race response or null

Output
a response or null

1. Let response be null.
2. Let eventCanceled be false.
3. Let client be request’s client.
4. Let activeWorker be registration’s active worker.
5. Let eventHandled be null.
6. Let handleFetchFailed be false.
7. Let respondWithEntered be false.
8. Let networkError be a network error.
9. If raceResponse is not null:
1. Set networkError’s service worker timing info to timingInfo.
10. Let shouldSoftUpdate be true if any of the following are true, and
false otherwise:
+ request is a non-subresource request.
+ request is a subresource request and registration is stale.
11. If the result of running the Should Skip Event algorithm with
"fetch" and activeWorker is true, then:
1. If shouldSoftUpdate is true, then in parallel run the Soft
Update algorithm with registration.
2. Return null.
12. If activeWorker’s all fetch listeners are empty flag is set:
1. In parallel:
1. If activeWorker’s state is "activating", then wait for
activeWorker’s state to become "activated".
2. Run the Run Service Worker algorithm with activeWorker.
3. If shouldSoftUpdate is true, then run the Soft Update
algorithm with registration.
2. Return null.
13. If useHighResPerformanceTimers is true, then set
useHighResPerformanceTimers to activeWorker’s global object’s
cross-origin isolated capability.
14. Let timingInfo’s start time be the coarsened shared current time
given useHighResPerformanceTimers.
15. If activeWorker’s state is "activating", wait for activeWorker’s
state to become "activated".
16. If the result of running the Run Service Worker algorithm with
activeWorker is failure, then set handleFetchFailed to true.
17. Else:
1. Set workerRealm to the relevant realm of the activeWorker’s
global object.
2. Set eventHandled to a new promise in workerRealm.
3. If raceResponse is not null, set activeWorker’s global
object’s race response map[request] to raceResponse.
4. Queue a task task to run the following substeps:
1. Let e be the result of creating an event with FetchEvent.
2. Let abortController be a new AbortController object with
workerRealm.
3. Let requestObject be the result of creating a Request
object, given request, a new Headers object’s guard which
is "immutable", abortController’s signal, and
workerRealm.
4. Initialize e’s type attribute to fetch.
5. Initialize e’s cancelable attribute to true.
6. Initialize e’s request attribute to requestObject.
7. Initialize e’s preloadResponse to preloadResponse.
8. Initialize e’s clientId attribute to client’s id.
9. If request is a non-subresource request and request’s
destination is not "report", initialize e’s
resultingClientId attribute to reservedClient’s id, and
to the empty string otherwise.
10. If request is a navigation request, initialize e’s
replacesClientId attribute to request’s replaces client
id, and to the empty string otherwise.
11. Initialize e’s handled to eventHandled.
12. Let timingInfo’s fetch event dispatch time to the
coarsened shared current time given
useHighResPerformanceTimers.
13. Dispatch e at activeWorker’s global object.
14. Invoke Update Service Worker Extended Events Set with
activeWorker and e.
15. If e’s respond-with entered flag is set, set
respondWithEntered to true.
16. If e’s wait to respond flag is set, then:
1. Wait until e’s wait to respond flag is unset.
2. If e’s respond-with error flag is set, set
handleFetchFailed to true.
3. Else, set response to e’s potential response.
17. If response is null, request’s body is not null, and
request’s body’s source is null, then:
1. If request’s body is unusable, set handleFetchFailed
to true.
2. Else, cancel request’s body with undefined.
18. If response is not null, then set response’s service
worker timing info to timingInfo.
19. If e’s canceled flag is set, set eventCanceled to true.
20. If fetchController state is "terminated" or "aborted",
then:
1. Let deserializedError be the result of deserialize a
serialized abort reason given fetchController’s
serialized abort reason and workerRealm.
2. Queue a task to signal abort on abortController with
deserializedError.
If task is discarded, set handleFetchFailed to true.
The task must use activeWorker’s event loop and the handle
fetch task source.
18. Wait for task to have executed or for handleFetchFailed to be true.
19. If shouldSoftUpdate is true, then in parallel run the Soft Update
algorithm with registration.
20. If activeWorker’s global object’s race response map[request]
exists, remove activeWorker’s global object’s race response
map[request].
21. If respondWithEntered is false, then:
1. If eventCanceled is true, then:
1. If eventHandled is not null, then reject eventHandled
with a "NetworkError" DOMException in workerRealm.
2. Return networkError.
2. If eventHandled is not null, then resolve eventHandled.
3. If raceResponse is not null, and raceResponse’s value is not
null, then:
1. Wait until raceResponse’s value is not "pending".
2. If raceResponse’s value is a response, return
raceResponse’s value.
4. Return null.
22. If handleFetchFailed is true, then:
1. If eventHandled is not null, then reject eventHandled with a
"NetworkError" DOMException in workerRealm.
2. Return networkError.
23. If eventHandled is not null, then resolve eventHandled.
24. Return response.

Parse URL Pattern

Input
rawPattern, a URLPatternCompatible

serviceWorker, a service worker

Output
A URL pattern

1. Let baseURL be serviceWorker’s script url.
2. Return the result of building a URL pattern from a Web IDL value
rawPattern given baseURL.

Verify Router Condition

Input
condition, a RouterCondition

serviceWorker, a service worker

Output
a boolean

1. Let hasCondition be false.
2. If condition["urlPattern"] exists, then:
1. Let rawPattern be condition["urlPattern"].
2. Let pattern be the result of running the Parse URL Pattern
algorithm passing rawPattern and serviceWorker. If this throws
an exception, catch it and return false.
3. If pattern has regexp groups, then return false.
Note: Since running a user-defined regular expression has a
security concern, it is prohibited.
4. Set hasCondition to true.
3. If condition["requestMethod"] exists:
1. Let method be condition["requestMethod"].
2. If method is not a method, then return false.
3. If method is a forbidden method, then return false.
4. Set hasCondition to true.
4. If condition["requestMode"] exists, set hasCondition to true.
5. If condition["requestDestination"] exists, set hasCondition to
true.
6. If condition["runningStatus"] exists, set hasCondition to true.
7. If condition["_or"] exists, then:
1. If hasCondition is true, return false.
Note: For ease of understanding the router rule, the "or"
condition is mutually exclusive with other conditions.
2. Let orConditions be condition["_or"].
3. For each orCondition of orConditions:
1. If running the Verify Router Condition algorithm with
orCondition and serviceWorker returns false, return
false.
4. Set hasCondition to true.
8. If condition["not"] exists, then:
1. If hasCondition is true, return false.
Note: For ease of understanding the router rule, the "not"
condition is mutually exclusive with other conditions.
2. If running the Verify Router Condition algorithm with
condition["not"] and serviceWorker returns false, return
false.
3. Set hasCondition to true.
9. Return hasCondition.

Match Router Condition

Input
condition, a RouterCondition

serviceWorker, a service worker

request, a request

Output
a boolean

Note: if there are multiple conditions (e.g. urlPattern, runningStatus,
and requestMethod are set), all conditions need to match for true to be
returned.
1. If condition["or"] exists, then:
1. Let orConditions be condition["or"].
2. For each orCondition of orConditions:
1. If running the Match Router Condition algorithm with
orCondition, serviceWorker and request returns true, then
return true.
3. Return false.
2. If condition["not"] exists, then:
1. If running the Match Router Condition algorithm with
condition["not"], serviceWorker and request returns true, then
return false.
2. Return true.
3. Else:
Note: The Verify Router Condition algorithm guarantees that or,
not, and other conditions are mutually exclusive.
1. If condition["urlPattern"] exists, then:
1. Let rawPattern be condition["urlPattern"].
2. Let pattern be the result of running the Parse URL
Pattern algorithm passing rawPattern and serviceWorker.
3. If running match with pattern and request’s URL returns
null, return false.
2. If condition["requestMethod"] exists, then:
1. Let method be condition["requestMethod"].
2. Normalize method.
3. If request’s method is not method, return false.
3. If condition["requestMode"] exists, then:
1. Let mode be condition["requestMode"].
2. If request’s mode is not mode, return false.
4. If condition["requestDestination"] exists, then:
1. Let destination be condition["requestDestination"].
2. If request’s destination is not destination, return
false.
5. If condition["runningStatus"] exists, then:
1. Let runningStatus be condition["runningStatus"].
2. If runningStatus is "running", and serviceWorker is not
running, return false.
3. If runningStatus is "not-running", and serviceWorker is
running, return false.
6. Return true.

Check Router Registration Limit

Input
routerRules, a list of router rules

Output
a boolean

Note: Router conditions can be complex and nested using _or and not. To
prevent excessive processing, this algorithm introduces two limits.
First, the total number of conditions, counting all nested conditions,
cannot exceed 1024. Second, the nesting depth is limited to 10 levels
to avoid exponential computation.
1. Let result be a count router condition result.
2. Set result’s condition count to 1024.
3. Set result’s quota exceeded to false.
4. For each rule of routerRules:
1. Set result to be the result of running Count Router Inner
Conditions with rule["condition"], result, and 10.
2. If result’s quota exceeded is true, return false.
5. Return true.

Count Router Inner Conditions

Input
condition, a RouterCondition

result, a count router condition result

depth, a number

Output
result, a count router condition result

1. Decrement result’s condition count by one.
2. If result’s condition count is zero, or depth is zero, then:
1. Set result’s quota exceeded to be true.
2. Return result.
3. If condition["_or"] exists, then:
1. Decrement depth by one.
2. For each orCondition of condition["_or"]:
1. Set result to be the result of running Count Router Inner
Conditions with orCondition, result, and depth.
2. If result’s quota exceeded is true, return result.
4. Else if condition["not"] exists, then:
1. Decrement depth by one.
2. Set result to be the result of running Count Router Inner
Conditions with condition["not"], result, and depth.
3. If result’s quota exceeded is true, return result.
5. Return result.

Get Router Source

Input
serviceWorker, a service worker

request, a request

Output
RouterSource or null

1. For each rule of serviceWorker’s list of router rules:
1. If running the Match Router Condition algorithm with
rule["condition"], serviceWorker and request returns true,
then return rule["source"].
2. Return null.

Should Skip Event

Input
eventName, a string

serviceWorker, a service worker

Output
a boolean

Note: To avoid unnecessary delays, this specification permits skipping
event dispatch when no event listeners for the event have been
deterministically added in the service worker’s global during the very
first script execution.
1. If serviceWorker’s set of event types to handle does not contain
eventName, then the user agent may return true.
2. Return false.

Fire Functional Event

Input
eventName, a string

eventConstructor, an event constructor that extends
ExtendableEvent

registration, a service worker registration

initialization, optional property initialization for event,
constructed from eventConstructor

postDispatchSteps, optional steps to run on the active worker’s
event loop, with dispatchedEvent set to the instance of
eventConstructor that was dispatched.

Output
None

1. Assert: registration’s active worker is not null.
2. Let activeWorker be registration’s active worker.
3. If the result of running Should Skip Event with eventName and
activeWorker is true, then:
1. If registration is stale, then in parallel run the Soft Update
algorithm with registration.
2. Return.
4. If activeWorker’s state is "activating", wait for activeWorker’s
state to become "activated".
5. If the result of running the Run Service Worker algorithm with
activeWorker is failure, then:
1. If registration is stale, then in parallel run the Soft Update
algorithm with registration.
2. Return.
6. Queue a task task to run these substeps:
1. Let event be the result of creating an event with
eventConstructor and the relevant realm of activeWorker’s
global object.
2. If initialization is not null, then initialize event with
initialization.
3. Dispatch event on activeWorker’s global object.
4. Invoke Update Service Worker Extended Events Set with
activeWorker and event.
5. If postDispatchSteps is not null, then run postDispatchSteps
passing event as dispatchedEvent.
The task must use activeWorker’s event loop and the handle
functional event task source.
7. Wait for task to have executed or been discarded.
8. If registration is stale, then in parallel run the Soft Update
algorithm with registration.

To fire an "amazingthing" event (which is of type AmazingThingEvent) on
a particular serviceWorkerRegistration, and initialize the event
object’s properties, the prose would be:
1. Fire Functional Event "amazingthing" using AmazingThingEvent on
serviceWorkerRegistration with the following properties:

propertyName
value

anotherPropertyName
anotherValue

Then run the following steps with dispatchedEvent:
1. Do whatever you need to with dispatchedEvent on the service
worker’s event loop.

Note that the initialization steps and post-dispatch steps are
optional. If they aren’t needed, the prose would be:
1. Fire Functional Event "whatever" using ExtendableEvent on
serviceWorkerRegistration.

Handle Service Worker Client Unload

The user agent must run these steps as part of when a service worker
client unloads via unloading document cleanup steps or termination.

Input
client, a service worker client

Output
None

1. Run the following steps atomically.
2. Let registration be the service worker registration used by client.
3. If registration is null, abort these steps.
4. If any other service worker client is using registration, abort
these steps.
5. If registration is unregistered, invoke Try Clear Registration with
registration.
6. Invoke Try Activate with registration.

Handle User Agent Shutdown

Input
None

Output
None

1. For each registration of registration map’s values:
1. If registration’s installing worker is not null, then:
1. If registration’s waiting worker is null and
registration’s active worker is null, invoke Clear
Registration with registration and continue to the next
iteration of the loop.
2. Else, set registration’s installing worker to null.
2. If registration’s waiting worker is not null, then in
parallel:
1. Invoke Activate with registration.

Update Service Worker Extended Events Set

Input
worker, a service worker

event, an ExtendableEvent

Output
None

1. Assert: event’s dispatch flag is unset.
2. For each item of worker’s set of extended events:
1. If item is not active, remove item from worker’s set of
extended events.
3. If event is active, append event to worker’s set of extended
events.

Unregister

Input
job, a job

Output
none

1. Let registration be the result of running Get Registration given
job’s storage key and job’s scope url.
2. If registration is null, then:
1. Invoke Resolve Job Promise with job and false.
2. Invoke Finish Job with job and abort these steps.
3. Remove registration map[(registration’s storage key, job’s scope
url)].
4. Invoke Resolve Job Promise with job and true.
5. Invoke Try Clear Registration with registration.
Note: If Try Clear Registration does not trigger Clear Registration
here, Clear Registration is tried again when the last client using
the registration is unloaded or the extend lifetime promises for
the registration’s service workers settle.
6. Invoke Finish Job with job.

Set Registration

Input
storage key, a storage key

scope, a URL

updateViaCache, an update via cache mode

Output
registration, a service worker registration

1. Run the following steps atomically.
2. Let scopeString be serialized scope with the exclude fragment flag
set.
3. Let registration be a new service worker registration whose storage
key is set to storage key, scope url is set to scope, and update
via cache mode is set to updateViaCache.
4. Set registration map[(storage key, scopeString)] to registration.
5. Return registration.

Clear Registration

Input
registration, a service worker registration

Output
None

1. Run the following steps atomically.
2. If registration’s installing worker is not null, then:
1. Terminate registration’s installing worker.
2. Run the Update Worker State algorithm passing registration’s
installing worker and "redundant" as the arguments.
3. Run the Update Registration State algorithm passing
registration, "installing" and null as the arguments.
3. If registration’s waiting worker is not null, then:
1. Terminate registration’s waiting worker.
2. Run the Update Worker State algorithm passing registration’s
waiting worker and "redundant" as the arguments.
3. Run the Update Registration State algorithm passing
registration, "waiting" and null as the arguments.
4. If registration’s active worker is not null, then:
1. Terminate registration’s active worker.
2. Run the Update Worker State algorithm passing registration’s
active worker and "redundant" as the arguments.
3. Run the Update Registration State algorithm passing
registration, "active" and null as the arguments.

Try Clear Registration

Input
registration, a service worker registration

Output
None

1. Invoke Clear Registration with registration if no service worker
client is using registration and all of the following conditions
are true:
+ registration’s installing worker is null or the result of
running Service Worker Has No Pending Events with
registration’s installing worker is true.
+ registration’s waiting worker is null or the result of running
Service Worker Has No Pending Events with registration’s
waiting worker is true.
+ registration’s active worker is null or the result of running
Service Worker Has No Pending Events with registration’s
active worker is true.

Update Registration State

Input
registration, a service worker registration

target, a string (one of "installing", "waiting", and "active")

source, a service worker or null

Output
None

1. Let registrationObjects be an array containing all the
ServiceWorkerRegistration objects associated with registration.
2. If target is "installing", then:
1. Set registration’s installing worker to source.
2. For each registrationObject in registrationObjects:
1. Queue a task to set the installing attribute of
registrationObject to null if registration’s installing
worker is null, or the result of getting the service
worker object that represents registration’s installing
worker in registrationObject’s relevant settings object.
3. Else if target is "waiting", then:
1. Set registration’s waiting worker to source.
2. For each registrationObject in registrationObjects:
1. Queue a task to set the waiting attribute of
registrationObject to null if registration’s waiting
worker is null, or the result of getting the service
worker object that represents registration’s waiting
worker in registrationObject’s relevant settings object.
4. Else if target is "active", then:
1. Set registration’s active worker to source.
2. For each registrationObject in registrationObjects:
1. Queue a task to set the active attribute of
registrationObject to null if registration’s active
worker is null, or the result of getting the service
worker object that represents registration’s active
worker in registrationObject’s relevant settings object.
The task must use registrationObject’s relevant settings object’s
responsible event loop and the DOM manipulation task source.

Update Worker State

Input
worker, a service worker

state, a service worker state

Output
None

1. Assert: state is not "parsed".
Note: "parsed" is the initial state. A service worker is never
updated to this state.
2. Set worker’s state to state.
3. Let settingsObjects be all environment settings objects whose
origin is worker’s script url’s origin.
4. For each settingsObject of settingsObjects, queue a task on
settingsObject’s responsible event loop in the DOM manipulation
task source to run the following steps:
1. Let objectMap be settingsObject’s service worker object map.
2. If objectMap[worker] does not exist, then abort these steps.
3. Let workerObj be objectMap[worker].
4. Set workerObj’s state to state.
5. Fire an event named statechange at workerObj.

Notify Controller Change

Input
client, a service worker client

Output
None

1. Assert: client is not null.
2. If client is an environment settings object, queue a task to fire
an event named controllerchange at the ServiceWorkerContainer
object that client is associated with.

The task must use client’s responsible event loop and the DOM
manipulation task source.

Match Service Worker Registration

Input
storage key, a storage key

clientURL, a URL

Output
A service worker registration

1. Run the following steps atomically.
2. Let clientURLString be serialized clientURL.
3. Let matchingScopeString be the empty string.
4. Let scopeStringSet be an empty list.
5. For each (entry storage key, entry scope) of registration map’s
keys:
1. If storage key equals entry storage key, then append entry
scope to the end of scopeStringSet.
6. Set matchingScopeString to the longest value in scopeStringSet
which the value of clientURLString starts with, if it exists.
Note: The URL string matching in this step is prefix-based rather
than path-structural. E.g. a client URL string with
"https://example.com/prefix-of/resource.html" will match a
registration for a scope with "https://example.com/prefix". The URL
string comparison is safe for the same-origin security as HTTP(S)
URLs are always serialized with a trailing slash at the end of the
origin part of the URLs.
7. Let matchingScope be null.
8. If matchingScopeString is not the empty string, then:
1. Set matchingScope to the result of parsing
matchingScopeString.
2. Assert: matchingScope’s origin and clientURL’s origin are same
origin.
9. Return the result of running Get Registration given storage key and
matchingScope.

Get Registration

Input
storage key, a storage key

scope, a URL

Output
A service worker registration

1. Run the following steps atomically.
2. Let scopeString be the empty string.
3. If scope is not null, set scopeString to serialized scope with the
exclude fragment flag set.
4. For each (entry storage key, entry scope) → registration of
registration map:
1. If storage key equals entry storage key and scopeString
matches entry scope, then return registration.
5. Return null.

Get Newest Worker

Input
registration, a service worker registration

Output
newestWorker, a service worker

1. Run the following steps atomically.
2. Let newestWorker be null.
3. If registration’s installing worker is not null, set newestWorker
to registration’s installing worker.
4. Else if registration’s waiting worker is not null, set newestWorker
to registration’s waiting worker.
5. Else if registration’s active worker is not null, set newestWorker
to registration’s active worker.
6. Return newestWorker.

Service Worker Has No Pending Events

Input
worker, a service worker

Output
True or false, a boolean

1. For each event of worker’s set of extended events:
1. If event is active, return false.
2. Return true.

Create Client

Input
client, a service worker client

Output
clientObject, a Client object

1. Let clientObject be a new Client object.
2. Set clientObject’s service worker client to client.
3. Return clientObject.

Create Window Client

Input
client, a service worker client

frameType, a string

visibilityState, a string

focusState, a boolean

ancestorOriginsList, a list

Output
windowClient, a WindowClient object

1. Let windowClient be a new WindowClient object.
2. Set windowClient’s service worker client to client.
3. Set windowClient’s frame type to frameType.
4. Set windowClient’s visibility state to visibilityState.
5. Set windowClient’s focus state to focusState.
6. Set windowClient’s ancestor origins array to a frozen array created
from ancestorOriginsList.
7. Return windowClient.

Get Frame Type

Input
navigable, a navigable

Output
frameType, a string

1. If navigable’s parent is not null, then return "nested".
2. If navigable’s active browsing context is an auxiliary browsing
context, then return "auxiliary".
3. Return "top-level".

Resolve Get Client Promise

Input
client, a service worker client

promise, a promise

Output
none

1. If client is an environment settings object, then:
1. If client is not a secure context, queue a task to reject
promise with a "SecurityError" DOMException, on promise’s
relevant settings object’s responsible event loop using the
DOM manipulation task source, and abort these steps.
2. Else:
1. If client’s creation URL is not a potentially trustworthy URL,
queue a task to reject promise with a "SecurityError"
DOMException, on promise’s relevant settings object’s
responsible event loop using the DOM manipulation task source,
and abort these steps.
3. If client is an environment settings object and is not a window
client, then:
1. Let clientObject be the result of running Create Client
algorithm with client as the argument.
2. Queue a task to resolve promise with clientObject, on
promise’s relevant settings object’s responsible event loop
using the DOM manipulation task source, and abort these steps.
4. Else:
1. Let browsingContext be null.
2. If client is an environment settings object, set
browsingContext to client’s global object’s browsing context.
3. Else, set browsingContext to client’s target browsing context.
4. Let navigable be the navigable with a active browsing context
of browsingContext.
5. Queue a task to run the following steps on browsingContext’s
event loop using the user interaction task source:
1. Let frameType be the result of running Get Frame Type
with navigable.
2. Let visibilityState be browsingContext’s active
document’s visibilityState attribute value.
3. Let focusState be the result of running the has focus
steps with browsingContext’s active document as the
argument.
4. Let ancestorOriginsList be the empty list.
5. If client is a window client, set ancestorOriginsList to
browsingContext’s active document’s ancestor origins
list’s associated list.
6. Queue a task to run the following steps on promise’s
relevant settings object’s responsible event loop using
the DOM manipulation task source:
1. If client’s discarded flag is set, resolve promise
with undefined and abort these steps.
2. Let windowClient be the result of running Create
Window Client with client, frameType,
visibilityState, focusState, and
ancestorOriginsList.
3. Resolve promise with windowClient.

Query Cache

Input
requestQuery, a request

options, a CacheQueryOptions object, optional

targetStorage, a request response list, optional

Output
resultList, a request response list

1. Let resultList be an empty list.
2. Let storage be null.
3. If the optional argument targetStorage is omitted, set storage to
the relevant request response list.
4. Else, set storage to targetStorage.
5. For each requestResponse of storage:
1. Let cachedRequest be requestResponse’s request.
2. Let cachedResponse be requestResponse’s response.
3. If Request Matches Cached Item with requestQuery,
cachedRequest, cachedResponse, and options returns true, then:
1. Let requestCopy be a copy of cachedRequest.
2. Let responseCopy be a copy of cachedResponse.
3. Add requestCopy/responseCopy to resultList.
6. Return resultList.

Request Matches Cached Item

Input
requestQuery, a request

request, a request

response, a response or null, optional, defaulting to null

options, a CacheQueryOptions object, optional

Output
a boolean

1. If options["ignoreMethod"] is false and request’s method is not
`GET`, return false.
2. Let queryURL be requestQuery’s url.
3. Let cachedURL be request’s url.
4. If options["ignoreSearch"] is true, then:
1. Set cachedURL’s query to the empty string.
2. Set queryURL’s query to the empty string.
5. If queryURL does not equal cachedURL with the exclude fragment flag
set, then return false.
6. If response is null, options["ignoreVary"] is true, or response’s
header list does not contain `Vary`, then return true.
7. Let fieldValues be the list containing the elements corresponding
to the field-values of the Vary header for the value of the header
with name `Vary`.
8. For each fieldValue in fieldValues:
1. If fieldValue matches "*", or the combined value given
fieldValue and request’s header list does not match the
combined value given fieldValue and requestQuery’s header
list, then return false.
9. Return true.

Batch Cache Operations

Input
operations, a list of cache batch operation objects

Output
resultList, a request response list

1. Let cache be the relevant request response list.
2. Let backupCache be a new request response list that is a copy of
cache.
3. Let addedItems be an empty list.
4. Try running the following substeps atomically:
1. Let resultList be an empty list.
2. For each operation in operations:
1. If operation’s type matches neither "delete" nor "put",
throw a TypeError.
2. If operation’s type matches "delete" and operation’s
response is not null, throw a TypeError.
3. If the result of running Query Cache with operation’s
request, operation’s options, and addedItems is not
empty, throw an "InvalidStateError" DOMException.
4. Let requestResponses be an empty list.
5. If operation’s type matches "delete", then:
1. Set requestResponses to the result of running Query
Cache with operation’s request and operation’s
options.
2. For each requestResponse in requestResponses:
1. Remove the item whose value matches
requestResponse from cache.
6. Else if operation’s type matches "put", then:
1. If operation’s response is null, throw a TypeError.
2. Let r be operation’s request’s associated request.
3. If r’s url’s scheme is not one of "http" and
"https", throw a TypeError.
4. If r’s method is not `GET`, throw a TypeError.
5. If operation’s options is not null, throw a
TypeError.
6. Set requestResponses to the result of running Query
Cache with operation’s request.
7. For each requestResponse of requestResponses:
1. Remove the item whose value matches
requestResponse from cache.
8. Append operation’s request/operation’s response to
cache.
9. If the cache write operation in the previous two
steps failed due to exceeding the granted quota
limit, throw a QuotaExceededError.
10. Append operation’s request/operation’s response to
addedItems.
7. Append operation’s request/operation’s response to
resultList.
3. Return resultList.
5. And then, if an exception was thrown, then:
1. Remove all the items from the relevant request response list.
2. For each requestResponse of backupCache:
1. Append requestResponse to the relevant request response
list.
3. Throw the exception.
Note: When an exception is thrown, the implementation does undo
(roll back) any changes made to the cache storage during the batch
operation job.

Is Async Module

Input
record, a Module Record

moduleMap, a module map

base, a URL

seen, a set of URLs

Output
a boolean

1. If record is not a Cyclic Module Record, then:
1. Return false.
2. If record.[[Async]] is true, then:
1. Return true.
3. For each string requested of record.[[RequestedModules]]:
1. Let url be the result of resolving a module specifier given
base and requested.
2. Assert: url is never failure, because resolving a module
specifier must have been previously successful with these same
two arguments.
3. If seen does not contain url, then:
1. Append url to seen.
2. If moduleMap[url] does not have a record, then:
1. Return false.
3. If Is Async Module for moduleMap[url]'s record,
moduleMap, base, and seen is true, then:
1. Return true.
4. Return false.

Lookup Race Response

Input
request, a request

Output
a response or null

1. Let registration be null.
2. If request is a non-subresource request, then:
1. If request’s reserved client is null, return null.
2. Let storage key be the result of running obtain a storage key
given request’s reserved client.
3. Set registration to the result of running Match Service Worker
Registration given storage key and request’s url.
3. Else if request is a subresource request, then:
1. Let client be request’s client.
2. If client’s active service worker is null, return null.
3. Set registration to client’s active service worker’s
containing service worker registration.
4. Otherwise, return null.
5. Let activeWorker be registration’s active worker.
6. Let map be activeWorker’s global object’s race response map.
7. If map[request] exists, then:
1. Let entry be map[request].
2. Wait until entry’s value is not "pending"
3. If entry’s value is a response, return the result of cloning
entry’s value.
8. Return null.

Appendix B: Extended HTTP headers

Service Worker Script Request

An HTTP request to fetch a service worker’s script resource will
include the following header:

`Service-Worker`
Indicates this request is a service worker’s script resource
request.

Note: This header helps administrators log the requests and
detect threats.

Service Worker Script Response

An HTTP response to a service worker’s script resource request can
include the following header:

`Service-Worker-Allowed`
Indicates the user agent will override the path restriction,
which limits the maximum allowed scope url that the script can
control, to the given value.

Note: The value is a URL. If a relative URL is given, it is
parsed against the script’s URL.

Default scope:
// Maximum allowed scope defaults to the path the script sits in
// "/js/" in this example
navigator.serviceWorker.register("/js/sw.js").then(() => {
console.log("Install succeeded with the default scope '/js/'.");
});

Upper path without Service-Worker-Allowed header:
// Set the scope to an upper path of the script location
// Response has no Service-Worker-Allowed header
navigator.serviceWorker.register("/js/sw.js", { scope: "/" }).catch(() => {
console.error("Install failed due to the path restriction violation.");
});

Upper path with Service-Worker-Allowed header:
// Set the scope to an upper path of the script location
// Response included "Service-Worker-Allowed : /"
navigator.serviceWorker.register("/js/sw.js", { scope: "/" }).then(() => {
console.log("Install succeeded as the max allowed scope was overriden to '/'."
);
});

A path restriction voliation even with Service-Worker-Allowed header:
// Set the scope to an upper path of the script location
// Response included "Service-Worker-Allowed : /foo"
navigator.serviceWorker.register("/foo/bar/sw.js", { scope: "/" }).catch(() => {
console.error("Install failed as the scope is still out of the overriden maxim
um allowed scope.");
});

Syntax

ABNF for the values of the headers used by the service worker’s script
resource requests and responses:
Service-Worker = %x73.63.72.69.70.74 ; "script", case-sensitive

Note: The validation of the Service-Worker-Allowed header’s values is
done by URL parsing algorithm (in Update algorithm) instead of using
ABNF.

8. Acknowledgements

Deep thanks go to Andrew Betts for organizing and hosting a small
workshop of like-minded individuals including: Jake Archibald, Jackson
Gabbard, Tobie Langel, Robin Berjon, Patrick Lauke, Christian Heilmann.
From the clarity of the day’s discussions and the use-cases outlined
there, much has become possible. Further thanks to Andrew for raising
consciousness about the offline problem. His organization of EdgeConf
and inclusion of Offline as a persistent topic there has created many
opportunities and connections that have enabled this work to progress.

Anne van Kesteren has generously lent his encyclopedic knowledge of Web
Platform arcana and standards development experience throughout the
development of the service worker. This specification would be
incomplete without his previous work in describing the real-world
behavior of URLs, HTTP Fetch, Promises, and DOM. Similarly, this
specification would not be possible without Ian Hickson’s rigorous Web
Worker spec. Much thanks to him.

In no particular order, deep gratitude for design guidance and
discussion goes to: Jungkee Song, Alec Flett, David Barrett-Kahn, Aaron
Boodman, Michael Nordman, Tom Ashworth, Kinuko Yasuda, Darin Fisher,
Jonas Sicking, Jesús Leganés Combarro, Mark Christian, Dave Hermann,
Yehuda Katz, François Remy, Ilya Grigorik, Will Chan, Domenic Denicola,
Nikhil Marathe, Yves Lafon, Adam Barth, Greg Simon, Devdatta Akhawe,
Dominic Cooney, Jeffrey Yasskin, Joshua Bell, Boris Zbarsky, Matt
Falkenhagen, Tobie Langel, Gavin Peters, Ben Kelly, Hiroki Nakagawa,
Jake Archibald, Josh Soref, Jinho Bang, Yutaka Hirano, Michael(tm)
Smith, isonmad, Ali Alabbas, Philip Jägenstedt, Mike Pennisi, and Eric
Willigers.

Jason Weber, Chris Wilson, Paul Kinlan, Ehsan Akhgari, and Daniel
Austin have provided valuable, well-timed feedback on requirements and
the standardization process.

The authors would also like to thank Dimitri Glazkov for his scripts
and formatting tools which have been essential in the production of
this specification. The authors are also grateful for his considerable
guidance.

Thanks also to Vivian Cromwell, Greg Simon, Alex Komoroske, Wonsuk Lee,
and Seojin Kim for their considerable professional support.

Conformance

Document conventions

Conformance requirements are expressed with a combination of
descriptive assertions and RFC 2119 terminology. The key words “MUST”,
“MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”,
“RECOMMENDED”, “MAY”, and “OPTIONAL” in the normative parts of this
document are to be interpreted as described in RFC 2119. However, for
readability, these words do not appear in all uppercase letters in this
specification.

All of the text of this specification is normative except sections
explicitly marked as non-normative, examples, and notes. [RFC2119]

Examples in this specification are introduced with the words “for
example” or are set apart from the normative text with class="example",
like this:

This is an example of an informative example.

Informative notes begin with the word “Note” and are set apart from the
normative text with class="note", like this:

Note, this is an informative note.

Conformant Algorithms

Requirements phrased in the imperative as part of algorithms (such as
"strip any leading space characters" or "return false and abort these
steps") are to be interpreted with the meaning of the key word ("must",
"should", "may", etc) used in introducing the algorithm.

Conformance requirements phrased as algorithms or specific steps can be
implemented in any manner, so long as the end result is equivalent. In
particular, the algorithms defined in this specification are intended
to be easy to understand and are not intended to be performant.
Implementers are encouraged to optimize.

Index

Terms defined by this specification

* Activate, in § Unnumbered section
* activate, in § 4.8
* "activated", in § 3.1
* "activating", in § 3.1
* active
+ attribute for ServiceWorkerRegistration, in § 3.2.4
+ dfn for ExtendableEvent, in § 4.4
* active worker, in § 2.3
* addAll(requests), in § 5.4.4
* add lifetime promise, in § 4.4
* add(request), in § 5.4.3
* addRoutes(rules), in § 4.5.1
* "all"
+ enum-value for ClientType, in § 4.3
+ enum-value for ServiceWorkerUpdateViaCache, in § 3.2
* All Fetch Listeners Are Empty, in § Unnumbered section
* all fetch listeners are empty flag, in § 2.1
* ancestorOrigins, in § 4.2.9
* ancestor origins array, in § 4.2
* "auxiliary", in § 4.2
* bad import script response, in § Unnumbered section
* Batch Cache Operations, in § Unnumbered section
* browsing context, in § 4.2
* "cache", in § 4.5
* Cache, in § 5.4
* cache batch operation, in § 5.4
* cacheName
+ dict-member for MultiCacheQueryOptions, in § 5.5
+ dict-member for RouterSourceDict, in § 4.5
* CacheQueryOptions, in § 5.4
* caches, in § 5.3.1
* CacheStorage, in § 5.5
* Check Router Registration Limit, in § Unnumbered section
* claim(), in § 4.3.4
* classic scripts imported flag, in § 2.1
* Clear Registration, in § Unnumbered section
* Client, in § 4.2
* client, in § Unnumbered section
* clientId
+ attribute for FetchEvent, in § 4.6.3
+ dict-member for FetchEventInit, in § 4.6
* client message queue, in § 3.4
* ClientQueryOptions, in § 4.3
* Clients, in § 4.3
* clients, in § 4.1.1
* ClientType, in § 4.3
* condition, in § 4.5
* condition count, in § 4.5
* constructor(type)
+ constructor for ExtendableEvent, in § 4.4
+ constructor for ExtendableMessageEvent, in § 4.7
+ constructor for InstallEvent, in § 4.5
* constructor(type, eventInitDict)
+ constructor for ExtendableEvent, in § 4.4
+ constructor for ExtendableMessageEvent, in § 4.7
+ constructor for FetchEvent, in § 4.6
+ constructor for InstallEvent, in § 4.5
* containing job queue, in § Unnumbered section
* containing service worker registration, in § 2.1
* controlled, in § 2.5
* controller, in § 3.4.1
* controllerchange, in § 3.5
* controlling, in § 2.5
* controls, in § 2.5
* count router condition result, in § 4.5
* Count Router Inner Conditions, in § Unnumbered section
* Create Client, in § Unnumbered section
* Create Fetch Event and Dispatch, in § Unnumbered section
* Create Job, in § Unnumbered section
* Create Window Client, in § Unnumbered section
* data
+ attribute for ExtendableMessageEvent, in § 4.7.1
+ dict-member for ExtendableMessageEventInit, in § 4.7
* dedicated worker client, in § 2.4
* delete(cacheName), in § 5.5.4
* delete(request), in § 5.4.6
* delete(request, options), in § 5.4.6
* disable(), in § 3.6.2
* discarded flag, in § 2.4
* enable(), in § 3.6.1
* enabled, in § 3.6
* equivalent, in § Unnumbered section
* ExtendableEvent, in § 4.4
* ExtendableEventInit, in § 4.4
* ExtendableEvent(type), in § 4.4
* ExtendableEvent(type, eventInitDict), in § 4.4
* ExtendableMessageEvent, in § 4.7
* ExtendableMessageEventInit, in § 4.7
* ExtendableMessageEvent(type), in § 4.7
* ExtendableMessageEvent(type, eventInitDict), in § 4.7
* extend lifetime promises, in § 4.4
* fetch, in § 4.8
* "fetch-event", in § 4.5
* FetchEvent, in § 4.6
* fetch event dispatch time, in § 2.2
* FetchEventInit, in § 4.6
* FetchEvent(type, eventInitDict), in § 4.6
* Finish Job, in § Unnumbered section
* Fire a functional event, in § Unnumbered section
* Fire Functional Event, in § Unnumbered section
* focus(), in § 4.2.10
* focused, in § 4.2.8
* focus state, in § 4.2
* force bypass cache flag, in § Unnumbered section
* force bypass cache for import scripts flag, in § 4.1
* frame type, in § 4.2
* FrameType, in § 4.2
* frameType, in § 4.2.2
* Functional events, in § 2.1.2
* Get Frame Type, in § Unnumbered section
* get(id), in § 4.3.1
* Get Newest Worker, in § Unnumbered section
* Get Registration, in § Unnumbered section
* getRegistration(), in § 3.4.4
* getRegistration(clientURL), in § 3.4.4
* getRegistrations(), in § 3.4.5
* Get Router Source, in § Unnumbered section
* getState(), in § 3.6.4
* get the service worker object, in § 3.1.1
* get the service worker registration object, in § 3.2.1
* getting the service worker object, in § 3.1.1
* getting the service worker registration object, in § 3.2.1
* global object, in § 2.1
* handled
+ attribute for FetchEvent, in § 4.6.6
+ dict-member for FetchEventInit, in § 4.6
* Handle Fetch, in § Unnumbered section
* handle fetch task source, in § 2.6
* handle functional event task source, in § 2.6
* Handle Service Worker Client Unload, in § Unnumbered section
* Handle User Agent Shutdown, in § Unnumbered section
* has(cacheName), in § 5.5.2
* has ever been evaluated flag, in § 2.1
* headerValue, in § 3.6
* id, in § 4.2.3
* ignoreMethod, in § 5.4
* ignoreSearch, in § 5.4
* ignoreVary, in § 5.4
* "imports", in § 3.2
* importScripts(urls), in § 6.3.2
* includeUncontrolled, in § 4.3
* Install, in § Unnumbered section
* install, in § 4.8
* "installed", in § 3.1
* InstallEvent, in § 4.5
* InstallEvent(type), in § 4.5
* InstallEvent(type, eventInitDict), in § 4.5
* "installing", in § 3.1
* installing, in § 3.2.2
* installing worker, in § 2.3
* Is Async Module, in § Unnumbered section
* job, in § Unnumbered section
* job promise, in § Unnumbered section
* job queue, in § Unnumbered section
* job type, in § Unnumbered section
* keys()
+ method for Cache, in § 5.4.7
+ method for CacheStorage, in § 5.5.5
* keys(request), in § 5.4.7
* keys(request, options), in § 5.4.7
* lastEventId
+ attribute for ExtendableMessageEvent, in § 4.7.3
+ dict-member for ExtendableMessageEventInit, in § 4.7
* last update check time, in § 2.3
* Lifecycle events, in § 2.1.2
* list of equivalent jobs, in § Unnumbered section
* list of router rules, in § 2.1
* Lookup Race Response, in § Unnumbered section
* matchAll()
+ method for Cache, in § 5.4.2
+ method for Clients, in § 4.3.2
* matchAll(options), in § 4.3.2
* matchAll(request), in § 5.4.2
* matchAll(request, options), in § 5.4.2
* match(request)
+ method for Cache, in § 5.4.1
+ method for CacheStorage, in § 5.5.1
* match(request, options)
+ method for Cache, in § 5.4.1
+ method for CacheStorage, in § 5.5.1
* Match Router Condition, in § Unnumbered section
* Match Service Worker Registration, in § Unnumbered section
* message
+ event for ServiceWorkerContainer, in § 3.5
+ event for ServiceWorkerGlobalScope, in § 4.8
* messageerror
+ event for ServiceWorkerContainer, in § 3.5
+ event for ServiceWorkerGlobalScope, in § 4.8
* MultiCacheQueryOptions, in § 5.5
* name to cache map, in § 5.1
* navigate(url), in § 4.2.11
* navigationPreload, in § 3.2.5
* navigation preload enabled flag, in § 2.3
* navigation preload header value, in § 2.3
* NavigationPreloadManager
+ (interface), in § 3.6
+ dfn for service worker registration, in § 2.3
* NavigationPreloadState, in § 3.6
* "nested", in § 4.2
* "network", in § 4.5
* "none"
+ enum-value for FrameType, in § 4.2
+ enum-value for ServiceWorkerUpdateViaCache, in § 3.2
* not, in § 4.5
* Notify Controller Change, in § Unnumbered section
* "not-running", in § 4.5
* onactivate, in § 4.1.5
* oncontrollerchange, in § 3.4.7
* onfetch, in § 4.1.5
* oninstall, in § 4.1.5
* onmessage
+ attribute for ServiceWorkerContainer, in § 3.4.7
+ attribute for ServiceWorkerGlobalScope, in § 4.1.5
* onmessageerror
+ attribute for ServiceWorkerContainer, in § 3.4.7
+ attribute for ServiceWorkerGlobalScope, in § 4.1.5
* onstatechange, in § 3.1.6
* onupdatefound, in § 3.2.10
* open(cacheName), in § 5.5.3
* openWindow(url), in § 4.3.3
* options, in § 5.4
* or, in § 4.5
* origin
+ attribute for ExtendableMessageEvent, in § 4.7.2
+ dfn for ExtendableMessageEvent, in § 4.7
+ dfn for service worker client, in § 2.4
+ dict-member for ExtendableMessageEventInit, in § 4.7
* "parsed", in § 3.1
* Parse URL Pattern, in § Unnumbered section
* pending promises count, in § 4.4
* policy container, in § 2.1
* ports
+ attribute for ExtendableMessageEvent, in § 4.7.5
+ dict-member for ExtendableMessageEventInit, in § 4.7
* postMessage(message)
+ method for Client, in § 4.2.6
+ method for ServiceWorker, in § 3.1.5
* postMessage(message, options)
+ method for Client, in § 4.2.6
+ method for ServiceWorker, in § 3.1.5
* postMessage(message, transfer)
+ method for Client, in § 4.2.5
+ method for ServiceWorker, in § 3.1.4
* potential response, in § 4.6
* preloadResponse
+ attribute for FetchEvent, in § 4.6.2
+ dict-member for FetchEventInit, in § 4.6
* put(request, response), in § 5.4.5
* Query Cache, in § Unnumbered section
* quota exceeded, in § 4.5
* "race-network-and-fetch-handler", in § 4.5
* race response, in § 4.1
* race response map, in § 4.1
* race result, in § Unnumbered section
* ready, in § 3.4.2
* ready promise, in § 3.4
* "redundant", in § 3.1
* referrer, in § Unnumbered section
* Register, in § Unnumbered section
* register(scriptURL), in § 3.4.3
* register(scriptURL, options), in § 3.4.3
* registration
+ attribute for ServiceWorkerGlobalScope, in § 4.1.2
+ dfn for service worker, in § 2.1
* registration map, in § Unnumbered section
* RegistrationOptions, in § 3.4
* Reject Job Promise, in § Unnumbered section
* relevant name to cache map, in § 5.1
* relevant request response list, in § 5.1
* replacesClientId
+ attribute for FetchEvent, in § 4.6.5
+ dict-member for FetchEventInit, in § 4.6
* request
+ attribute for FetchEvent, in § 4.6.1
+ dfn for cache batch operation, in § 5.4
+ dict-member for FetchEventInit, in § 4.6
* requestDestination, in § 4.5
* Request Matches Cached Item, in § Unnumbered section
* requestMethod, in § 4.5
* requestMode, in § 4.5
* request response list, in § 5.1
* Resolve Get Client Promise, in § Unnumbered section
* Resolve Job Promise, in § Unnumbered section
* respond-with entered flag, in § 4.6
* respond-with error flag, in § 4.6
* respondWith(r), in § 4.6.7
* response, in § 5.4
* resultingClientId
+ attribute for FetchEvent, in § 4.6.4
+ dict-member for FetchEventInit, in § 4.6
* routed response, in § Unnumbered section
* RouterCondition, in § 4.5
* RouterRule, in § 4.5
* RouterSource, in § 4.5
* RouterSourceDict, in § 4.5
* RouterSourceEnum, in § 4.5
* Run Job, in § Unnumbered section
* "running", in § 4.5
* running, in § 2.1
* RunningStatus, in § 4.5
* runningStatus, in § 4.5
* Run Service Worker, in § Unnumbered section
* Schedule Job, in § Unnumbered section
* scope
+ attribute for ServiceWorkerRegistration, in § 3.2.6
+ dict-member for RegistrationOptions, in § 3.4
* scope to job queue map, in § Unnumbered section
* scope url
+ dfn for job, in § Unnumbered section
+ dfn for service worker registration, in § 2.3
* script resource, in § 2.1
* script resource map, in § 2.1
* script url
+ dfn for job, in § Unnumbered section
+ dfn for service worker, in § 2.1
* scriptURL, in § 3.1.2
* service worker
+ dfn for , in § 2.1
+ dfn for ServiceWorkerGlobalScope, in § 4.1
* Service-Worker, in § Unnumbered section
* ServiceWorker, in § 3.1
* serviceWorker
+ attribute for Navigator, WorkerNavigator, in § 3.3
+ attribute for ServiceWorkerGlobalScope, in § 4.1.3
* Service-Worker-Allowed, in § Unnumbered section
* service worker client
+ dfn for , in § 2.4
+ dfn for Client, in § 4.2
+ dfn for ServiceWorkerContainer, in § 3.4
* ServiceWorkerContainer, in § 3.4
* service worker events, in § 2.1.2
* ServiceWorkerGlobalScope, in § 4.1
* Service Worker Has No Pending Events, in § Unnumbered section
* service worker object map, in § 3.1.1
* [[service worker queue]], in § 2.1
* service worker registration
+ dfn for , in § 2.3
+ dfn for ServiceWorkerRegistration, in § 3.2
* ServiceWorkerRegistration, in § 3.2
* service worker registration object map, in § 3.2.1
* ServiceWorkerState, in § 3.1
* service worker timing info, in § 2.2
* ServiceWorkerUpdateViaCache, in § 3.2
* setHeaderValue(value), in § 3.6.3
* set of event types to handle, in § 2.1
* set of extended events, in § 2.1
* set of used scripts, in § 2.1
* Set Registration, in § Unnumbered section
* Setup ServiceWorkerGlobalScope, in § Unnumbered section
* "sharedworker", in § 4.3
* shared worker client, in § 2.4
* Should Skip Event, in § Unnumbered section
* skipWaiting(), in § 4.1.4
* skip waiting flag, in § 2.1
* Soft Update, in § Unnumbered section
* source
+ attribute for ExtendableMessageEvent, in § 4.7.4
+ dict-member for ExtendableMessageEventInit, in § 4.7
+ dict-member for RouterRule, in § 4.5
* stale, in § 2.3
* startMessages(), in § 3.4.6
* Start Register, in § Unnumbered section
* start status, in § 2.1
* start time, in § 2.2
* state
+ attribute for ServiceWorker, in § 3.1.3
+ dfn for service worker, in § 2.1
* statechange, in § 3.5
* storage key
+ dfn for job, in § Unnumbered section
+ dfn for service worker registration, in § 2.3
* task queues, in § 2.3
* Terminate Service Worker, in § Unnumbered section
* timed out flag, in § 4.4
* "top-level", in § 4.2
* Try Activate, in § Unnumbered section
* Try Clear Registration, in § Unnumbered section
* type
+ attribute for Client, in § 4.2.4
+ dfn for cache batch operation, in § 5.4
+ dfn for service worker, in § 2.1
+ dict-member for ClientQueryOptions, in § 4.3
+ dict-member for RegistrationOptions, in § 3.4
* Unregister, in § Unnumbered section
* unregister(), in § 3.2.9
* unregistered, in § 2.3
* Update, in § Unnumbered section
* update(), in § 3.2.8
* updatefound, in § 3.5
* Update Registration State, in § Unnumbered section
* Update Service Worker Extended Events Set, in § Unnumbered section
* updateViaCache
+ attribute for ServiceWorkerRegistration, in § 3.2.7
+ dict-member for RegistrationOptions, in § 3.4
* update via cache mode
+ dfn for job, in § Unnumbered section
+ dfn for service worker registration, in § 2.3
* Update Worker State, in § Unnumbered section
* url, in § 4.2.1
* urlPattern, in § 4.5
* used, in § 2.5
* used route, in § Unnumbered section
* uses, in § 2.5
* using, in § 2.5
* value, in § 4.1
* Verify Router Condition, in § Unnumbered section
* visibility state, in § 4.2
* visibilityState, in § 4.2.7
* waiting, in § 3.2.3
* waiting worker, in § 2.3
* wait to respond flag, in § 4.6
* waitUntil(f), in § 4.4.1
* "window", in § 4.3
* window client, in § 2.4
* WindowClient, in § 4.2
* "worker", in § 4.3
* worker cache lookup start, in § 2.2
* worker client, in § 2.4
* worker final router source, in § 2.2
* worker matched router source, in § 2.2
* worker router evaluation start, in § 2.2
* worker type, in § Unnumbered section

Terms defined by reference

* [CSP-NEXT] defines the following terms:
+ enforced
* [CSP3] defines the following terms:
+ Run CSP initialization for a global object
* [DOM] defines the following terms:
+ AbortController
+ Event
+ EventInit
+ EventTarget
+ callback
+ cancelable
+ canceled flag
+ creating an event
+ dispatch
+ dispatch flag
+ document
+ event
+ event listener
+ fire an event
+ handleEvent(event)
+ isTrusted
+ legacy-obtain service worker fetch event listener callbacks
+ signal
+ signal abort
+ stop immediate propagation flag
+ stop propagation flag
+ type
* [ECMASCRIPT] defines the following terms:
+ IsCallable
+ NormalCompletion
+ ThrowCompletion
+ abrupt completion
+ agent
+ Assert
+ Completion
+ Cyclic Module Record
+ declaration
+ DetachArrayBuffer
+ function body
+ map objects
+ Module Record
+ promise
+ statement
+ surrounding agent
* [FETCH] defines the following terms:
+ "report"
+ Headers
+ Request
+ RequestDestination
+ RequestInfo
+ RequestMode
+ Response
+ abort
+ aborted flag
+ append
+ basic filtered response
+ body
+ body (for request)
+ body (for response)
+ body info
+ cache mode
+ cache state
+ client
+ clone (for request)
+ clone (for response)
+ combine
+ contains
+ CORS filtered response
+ creating
+ cross-origin resource policy check
+ decoded size
+ deserialize a serialized abort reason
+ destination
+ encoded size
+ extract a MIME type
+ extracting header list values
+ fetch
+ fetch controller
+ fetch(input, init)
+ filtered response
+ forbidden method
+ guard
+ header
+ header list (for request)
+ header list (for response)
+ HTTP fetch
+ internal response
+ method
+ method (for request)
+ mode
+ name
+ navigation request
+ network error
+ non-subresource request
+ normalize
+ ok status
+ opaque filtered response
+ origin
+ parser metadata
+ processResponse
+ processResponseEndOfBody
+ redirect mode
+ replaces client id
+ request
+ request (for Request)
+ reserved client
+ response
+ response (for Response)
+ serialized abort reason
+ service worker timing info
+ service-workers mode
+ source
+ state
+ status
+ stream
+ subresource request
+ type
+ unusable
+ URL
+ use-URL-credentials flag
+ value
* [FileAPI] defines the following terms:
+ blob URL
* [HR-TIME-3] defines the following terms:
+ DOMHighResTimeStamp
+ coarsen time
+ coarsened shared current time
+ unsafe shared current time
* [HTML] defines the following terms:
+ AbstractWorker
+ DOMContentLoaded
+ DedicatedWorkerGlobalScope
+ EventHandler
+ MessageEvent
+ MessagePort
+ Navigator
+ SharedWorkerGlobalScope
+ StructuredDeserializeWithTransfer
+ StructuredSerializeOptions
+ StructuredSerializeWithTransfer
+ VisibilityState
+ Window
+ WindowOrWorkerGlobalScope
+ Worker
+ WorkerGlobalScope
+ WorkerLocation
+ WorkerNavigator
+ WorkerType
+ abort a running script
+ active browsing context
+ active document
+ active service worker
+ ancestor origins list
+ API base URL
+ associated Document
+ auxiliary browsing context
+ base URL
+ browsing context
+ classic script
+ closing
+ create a new realm
+ creating a policy container from a fetch response
+ creation URL
+ cross-origin isolated capability
+ current global object
+ data
+ DOM manipulation task source
+ environment
+ environment discarding steps
+ environment settings object
+ event handler
+ event handler event type
+ event handler IDL attribute
+ event handler map
+ event loop
+ event loop (for agent)
+ exceptionsEnabled
+ execution ready flag
+ extract an origin
+ fetch a classic worker script
+ fetch a classic worker-imported script
+ fetch a module worker script graph
+ focusing steps
+ fully active
+ global object
+ global object (for environment settings object)
+ has cross-site ancestor
+ has focus steps
+ historyHandling
+ id
+ iframe
+ import scripts into worker global scope
+ in parallel
+ incumbent settings object
+ isTopLevel
+ listener
+ map of active timers
+ message
+ module map
+ module map (for WorkerGlobalScope)
+ module map (for environment settings object)
+ module script
+ navigable
+ navigate
+ obtain a service worker agent
+ opaque origin
+ origin
+ origin (for MessageEvent)
+ origin (for environment settings object)
+ owner set
+ parallel queue
+ parent
+ perform the fetch hook
+ policy container
+ policy container (for WorkerGlobalScope)
+ policy container (for environment settings object)
+ ports
+ queue a microtask
+ queue a task
+ realm (for environment settings object)
+ realm (for global object)
+ realm execution context
+ record
+ relevant global object
+ relevant realm
+ relevant settings object
+ resolve a module specifier
+ responsible event loop
+ run a classic script
+ run a module script
+ run a worker
+ same origin
+ sandbox
+ script
+ secure context
+ serialization of an origin
+ source
+ target browsing context
+ task
+ task queues
+ task source
+ time origin
+ top-level browsing context
+ top-level creation URL
+ top-level origin
+ transfer
+ transient activation
+ type
+ unloading document cleanup steps
+ unsafe response
+ url
+ user interaction task source
+ value
+ visibilityState
+ web worker
* [INFRA] defines the following terms:
+ abort when
+ append (for list)
+ append (for set)
+ ASCII case-insensitive
+ break
+ byte sequence
+ clear
+ contain
+ continue
+ dequeue
+ enqueue
+ entry
+ exist
+ for each (for list)
+ for each (for map)
+ get the keys
+ if aborted
+ implementation-defined
+ is
+ is not empty
+ item (for list)
+ item (for struct)
+ key
+ keys
+ list
+ map
+ ordered map
+ ordered set
+ queue
+ remove (for list)
+ remove (for map)
+ set
+ set (for map)
+ struct
+ tuple
+ value
+ values
* [MIMESNIFF] defines the following terms:
+ JavaScript MIME type
* [NAVIGATION-TIMING-2] defines the following terms:
+ PerformanceNavigationTiming
* [PAGE-LIFECYCLE] defines the following terms:
+ discard
* [PUSH] defines the following terms:
+ push
* [RESOURCE-TIMING] defines the following terms:
+ PerformanceResourceTiming
* [RFC7230] defines the following terms:
+ field-value
* [RFC7231] defines the following terms:
+ Vary
* [SCREEN-CAPTURE] defines the following terms:
+ monitor
* [SECURE-CONTEXTS] defines the following terms:
+ potentially trustworthy origin
+ potentially trustworthy URL
* [STORAGE] defines the following terms:
+ equal
+ obtain a local storage bottle map
+ obtain a storage key
+ storage key
* [STREAMS] defines the following terms:
+ ReadableStream
+ cancel (for ReadableStream)
+ cancel (for ReadableStreamDefaultReader)
+ cancelAlgorithm
+ chunk
+ chunk steps
+ close
+ close steps
+ disturbed
+ enqueue
+ error
+ error steps
+ errored
+ getting a reader
+ highWaterMark
+ locked
+ pullAlgorithm
+ read a chunk
+ read request
+ reading all bytes
+ set up
+ sizeAlgorithm
* [TRUSTED-TYPES] defines the following terms:
+ TrustedScriptURL
+ get trusted type compliant string
* [URL] defines the following terms:
+ equal
+ fragment
+ origin
+ path
+ query
+ scheme
+ URL
+ URL parser
+ URL serializer
* [URLPATTERN] defines the following terms:
+ URLPatternCompatible
+ build a URL pattern from a Web IDL value
+ has regexp groups
+ match
+ URL pattern
* [WEBIDL] defines the following terms:
+ AbortError
+ ArrayBuffer
+ ByteString
+ DOMException
+ DOMString
+ Exposed
+ FrozenArray
+ Global
+ InvalidAccessError
+ InvalidStateError
+ NetworkError
+ NewObject
+ Promise
+ QuotaExceededError
+ SameObject
+ SecureContext
+ SecurityError
+ TypeError
+ USVString
+ Uint8Array
+ a new promise
+ a promise rejected with
+ a promise resolved with
+ any
+ boolean
+ byte length
+ converted to an ECMAScript value
+ create
+ create a frozen array
+ exception
+ frozen array type
+ getting a promise to wait for all
+ new
+ object
+ partial interface
+ reacting
+ reject
+ resolve
+ sequence
+ this
+ throw
+ undefined
+ upon fulfillment
+ upon rejection

References

Normative References

[CSP-NEXT]
Scripting Policy. Editor's Draft. URL:
https://wicg.github.io/csp-next/scripting-policy.html

[CSP3]
Mike West; Antonio Sartori. Content Security Policy Level 3. 1
April 2026. WD. URL: https://www.w3.org/TR/CSP3/

[DOM]
Anne van Kesteren. DOM Standard. Living Standard. URL:
https://dom.spec.whatwg.org/

[ECMASCRIPT]
ECMAScript Language Specification. URL:
https://tc39.es/ecma262/multipage/

[FETCH]
Anne van Kesteren. Fetch Standard. Living Standard. URL:
https://fetch.spec.whatwg.org/

[FileAPI]
Marijn Kruisselbrink. File API. 3 December 2025. WD. URL:
https://www.w3.org/TR/FileAPI/

[HR-TIME-3]
Yoav Weiss. High Resolution Time. 24 March 2026. WD. URL:
https://www.w3.org/TR/hr-time-3/

[HTML]
Anne van Kesteren; et al. HTML Standard. Living Standard. URL:
https://html.spec.whatwg.org/multipage/

[INFRA]
Anne van Kesteren; Domenic Denicola. Infra Standard. Living
Standard. URL: https://infra.spec.whatwg.org/

[MIMESNIFF]
Gordon P. Hemsley. MIME Sniffing Standard. Living Standard. URL:
https://mimesniff.spec.whatwg.org/

[NAVIGATION-TIMING-2]
Yoav Weiss; Noam Rosenthal. Navigation Timing Level 2. 25
February 2026. WD. URL:
https://www.w3.org/TR/navigation-timing-2/

[PAGE-LIFECYCLE]
Page Lifecycle. Draft Community Group Report. URL:
https://wicg.github.io/page-lifecycle/

[RESOURCE-TIMING]
Yoav Weiss; Noam Rosenthal. Resource Timing. 23 March 2026. CRD.
URL: https://www.w3.org/TR/resource-timing/

[RFC2119]
S. Bradner. Key words for use in RFCs to Indicate Requirement
Levels. March 1997. Best Current Practice. URL:
https://datatracker.ietf.org/doc/html/rfc2119

[RFC5234]
D. Crocker, Ed.; P. Overell. Augmented BNF for Syntax
Specifications: ABNF. January 2008. Internet Standard. URL:
https://www.rfc-editor.org/rfc/rfc5234

[RFC7230]
R. Fielding, Ed.; J. Reschke, Ed.. Hypertext Transfer Protocol
(HTTP/1.1): Message Syntax and Routing. June 2014. Proposed
Standard. URL: https://httpwg.org/specs/rfc7230.html

[RFC7231]
R. Fielding, Ed.; J. Reschke, Ed.. Hypertext Transfer Protocol
(HTTP/1.1): Semantics and Content. June 2014. Proposed Standard.
URL: https://httpwg.org/specs/rfc7231.html

[SCREEN-CAPTURE]
Jan-Ivar Bruaroey; Elad Alon. Screen Capture. 17 July 2025. WD.
URL: https://www.w3.org/TR/screen-capture/

[SECURE-CONTEXTS]
Mike West. Secure Contexts. 10 November 2023. CRD. URL:
https://www.w3.org/TR/secure-contexts/

[STORAGE]
Anne van Kesteren. Storage Standard. Living Standard. URL:
https://storage.spec.whatwg.org/

[STREAMS]
Adam Rice; et al. Streams Standard. Living Standard. URL:
https://streams.spec.whatwg.org/

[TRUSTED-TYPES]
Krzysztof Kotowicz. Trusted Types. 24 February 2026. WD. URL:
https://www.w3.org/TR/trusted-types/

[URL]
Anne van Kesteren. URL Standard. Living Standard. URL:
https://url.spec.whatwg.org/

[URLPATTERN]
Ben Kelly; Jeremy Roman; 宍戸俊哉 (Shunya Shishido). URL Pattern
Standard. Living Standard. URL:
https://urlpattern.spec.whatwg.org/

[WEBIDL]
Edgar Chen; Timothy Gu. Web IDL Standard. Living Standard. URL:
https://webidl.spec.whatwg.org/

Informative References

[UNSANCTIONED-TRACKING]
Unsanctioned Web Tracking. 17 July 2015. Finding of the W3C TAG.
URL: https://www.w3.org/2001/tag/doc/unsanctioned-tracking/

IDL Index

[SecureContext, Exposed=(Window,Worker)]
interface ServiceWorker : EventTarget {
readonly attribute USVString scriptURL;
readonly attribute ServiceWorkerState state;
undefined postMessage(any message, sequence<object> transfer);
undefined postMessage(any message, optional StructuredSerializeOptions options
= {});

// event
attribute EventHandler onstatechange;
};
ServiceWorker includes AbstractWorker;

enum ServiceWorkerState {
"parsed",
"installing",
"installed",
"activating",
"activated",
"redundant"
};

[SecureContext, Exposed=(Window,Worker)]
interface ServiceWorkerRegistration : EventTarget {
readonly attribute ServiceWorker? installing;
readonly attribute ServiceWorker? waiting;
readonly attribute ServiceWorker? active;
[SameObject] readonly attribute NavigationPreloadManager navigationPreload;

readonly attribute USVString scope;
readonly attribute ServiceWorkerUpdateViaCache updateViaCache;

[NewObject] Promise<ServiceWorkerRegistration> update();
[NewObject] Promise<boolean> unregister();

// event
attribute EventHandler onupdatefound;
};

enum ServiceWorkerUpdateViaCache {
"imports",
"all",
"none"
};

partial interface Navigator {
[SecureContext, SameObject] readonly attribute ServiceWorkerContainer serviceW
orker;
};

partial interface WorkerNavigator {
[SecureContext, SameObject] readonly attribute ServiceWorkerContainer serviceW
orker;
};

[SecureContext, Exposed=(Window,Worker)]
interface ServiceWorkerContainer : EventTarget {
readonly attribute ServiceWorker? controller;
readonly attribute Promise<ServiceWorkerRegistration> ready;

[NewObject] Promise<ServiceWorkerRegistration> register((TrustedScriptURL or U
SVString) scriptURL, optional RegistrationOptions options = {});

[NewObject] Promise<(ServiceWorkerRegistration or undefined)> getRegistration(
optional USVString clientURL = "");
[NewObject] Promise<FrozenArray<ServiceWorkerRegistration>> getRegistrations()
;

undefined startMessages();

// events
attribute EventHandler oncontrollerchange;
attribute EventHandler onmessage; // event.source of message events is Service
Worker object
attribute EventHandler onmessageerror;
};

dictionary RegistrationOptions {
USVString scope;
WorkerType type = "classic";
ServiceWorkerUpdateViaCache updateViaCache = "imports";
};

[SecureContext, Exposed=(Window,Worker)]
interface NavigationPreloadManager {
Promise<undefined> enable();
Promise<undefined> disable();
Promise<undefined> setHeaderValue(ByteString value);
Promise<NavigationPreloadState> getState();
};

dictionary NavigationPreloadState {
boolean enabled = false;
ByteString headerValue;
};

[Global=(Worker,ServiceWorker), Exposed=ServiceWorker, SecureContext]
interface ServiceWorkerGlobalScope : WorkerGlobalScope {
[SameObject] readonly attribute Clients clients;
[SameObject] readonly attribute ServiceWorkerRegistration registration;
[SameObject] readonly attribute ServiceWorker serviceWorker;

[NewObject] Promise<undefined> skipWaiting();

attribute EventHandler oninstall;
attribute EventHandler onactivate;
attribute EventHandler onfetch;

attribute EventHandler onmessage;
attribute EventHandler onmessageerror;
};

[Exposed=ServiceWorker]
interface Client {
readonly attribute USVString url;
readonly attribute FrameType frameType;
readonly attribute DOMString id;
readonly attribute ClientType type;
undefined postMessage(any message, sequence<object> transfer);
undefined postMessage(any message, optional StructuredSerializeOptions options
= {});
};

[Exposed=ServiceWorker]
interface WindowClient : Client {
readonly attribute VisibilityState visibilityState;
readonly attribute boolean focused;
[SameObject] readonly attribute FrozenArray<USVString> ancestorOrigins;
[NewObject] Promise<WindowClient> focus();
[NewObject] Promise<WindowClient?> navigate(USVString url);
};

enum FrameType {
"auxiliary",
"top-level",
"nested",
"none"
};

[Exposed=ServiceWorker]
interface Clients {
// The objects returned will be new instances every time
[NewObject] Promise<(Client or undefined)> get(DOMString id);
[NewObject] Promise<FrozenArray<Client>> matchAll(optional ClientQueryOptions
options = {});
[NewObject] Promise<WindowClient?> openWindow(USVString url);
[NewObject] Promise<undefined> claim();
};

dictionary ClientQueryOptions {
boolean includeUncontrolled = false;
ClientType type = "window";
};

enum ClientType {
"window",
"worker",
"sharedworker",
"all"
};

[Exposed=ServiceWorker]
interface ExtendableEvent : Event {
constructor(DOMString type, optional ExtendableEventInit eventInitDict = {});
undefined waitUntil(Promise<any> f);
};

dictionary ExtendableEventInit : EventInit {
// Defined for the forward compatibility across the derived events
};

[Exposed=ServiceWorker]
interface InstallEvent : ExtendableEvent {
constructor(DOMString type, optional ExtendableEventInit eventInitDict = {});
Promise<undefined> addRoutes((RouterRule or sequence<RouterRule>) rules);
};

dictionary RouterRule {
required RouterCondition condition;
required RouterSource source;
};

dictionary RouterCondition {
URLPatternCompatible urlPattern;
ByteString requestMethod;
RequestMode requestMode;
RequestDestination requestDestination;
RunningStatus runningStatus;

sequence<RouterCondition> _or;
RouterCondition not;
};

typedef (RouterSourceDict or RouterSourceEnum) RouterSource;

dictionary RouterSourceDict {
DOMString cacheName;
};

enum RunningStatus { "running", "not-running" };
enum RouterSourceEnum {
"cache",
"fetch-event",
"network",
"race-network-and-fetch-handler"
};

[Exposed=ServiceWorker]
interface FetchEvent : ExtendableEvent {
constructor(DOMString type, FetchEventInit eventInitDict);
[SameObject] readonly attribute Request request;
readonly attribute Promise<any> preloadResponse;
readonly attribute DOMString clientId;
readonly attribute DOMString resultingClientId;
readonly attribute DOMString replacesClientId;
readonly attribute Promise<undefined> handled;

undefined respondWith(Promise<Response> r);
};

dictionary FetchEventInit : ExtendableEventInit {
required Request request;
Promise<any> preloadResponse;
DOMString clientId = "";
DOMString resultingClientId = "";
DOMString replacesClientId = "";
Promise<undefined> handled;
};

[Exposed=ServiceWorker]
interface ExtendableMessageEvent : ExtendableEvent {
constructor(DOMString type, optional ExtendableMessageEventInit eventInitDict
= {});
readonly attribute any data;
readonly attribute USVString origin;
readonly attribute DOMString lastEventId;
[SameObject] readonly attribute (Client or ServiceWorker or MessagePort)? sour
ce;
readonly attribute FrozenArray<MessagePort> ports;
};

dictionary ExtendableMessageEventInit : ExtendableEventInit {
any data = null;
USVString origin = "";
DOMString lastEventId = "";
(Client or ServiceWorker or MessagePort)? source = null;
sequence<MessagePort> ports = [];
};

partial interface mixin WindowOrWorkerGlobalScope {
[SecureContext, SameObject] readonly attribute CacheStorage caches;
};

[SecureContext, Exposed=(Window,Worker)]
interface Cache {
[NewObject] Promise<(Response or undefined)> match(RequestInfo request, option
al CacheQueryOptions options = {});
[NewObject] Promise<FrozenArray<Response>> matchAll(optional RequestInfo reque
st, optional CacheQueryOptions options = {});
[NewObject] Promise<undefined> add(RequestInfo request);
[NewObject] Promise<undefined> addAll(sequence<RequestInfo> requests);
[NewObject] Promise<undefined> put(RequestInfo request, Response response);
[NewObject] Promise<boolean> delete(RequestInfo request, optional CacheQueryOp
tions options = {});
[NewObject] Promise<FrozenArray<Request>> keys(optional RequestInfo request, o
ptional CacheQueryOptions options = {});
};

dictionary CacheQueryOptions {
boolean ignoreSearch = false;
boolean ignoreMethod = false;
boolean ignoreVary = false;
};

[SecureContext, Exposed=(Window,Worker)]
interface CacheStorage {
[NewObject] Promise<(Response or undefined)> match(RequestInfo request, option
al MultiCacheQueryOptions options = {});
[NewObject] Promise<boolean> has(DOMString cacheName);
[NewObject] Promise<Cache> open(DOMString cacheName);
[NewObject] Promise<boolean> delete(DOMString cacheName);
[NewObject] Promise<sequence<DOMString>> keys();
};

dictionary MultiCacheQueryOptions : CacheQueryOptions {
DOMString cacheName;
};

Issues Index

The behavior in this section is not fully specified yet and will be
specified in HTML Standard. The work is tracked by the issue and the
pull request. ↵
Using the to-be-created environment settings object rather than a
concrete environment settings object. This is used due to the unique
processing model of service workers compared to the processing model of
other web workers. The script fetching algorithms of HTML standard
originally designed for other web workers require an environment
settings object of the execution environment, but service workers fetch
a script separately in the Update algorithm before the script later
runs multiple times through the Run Service Worker algorithm. ↵
The fetch a classic worker script algorithm and the fetch a module
worker script graph algorithm in HTML take job’s client as an argument.
job’s client is null when passed from the Soft Update algorithm. ↵

✔MDN

Cache/add

In all current engines.
Firefox41+Safari11.1+Chrome44+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet4.0+Opera Mobile?

✔MDN

Cache/addAll

In all current engines.
Firefox41+Safari11.1+Chrome46+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Cache/delete

In all current engines.
Firefox41+Safari11.1+Chrome43+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Cache/keys

In all current engines.
Firefox41+Safari11.1+Chrome43+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Cache/match

In all current engines.
Firefox41+Safari11.1+Chrome43+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Cache/matchAll

In all current engines.
Firefox41+Safari11.1+Chrome47+
__________________________________________________________________

Opera34+Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Cache/put

In all current engines.
Firefox41+Safari11.1+Chrome43+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet4.0+Opera Mobile?

✔MDN

Cache

In all current engines.
Firefox41+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet4.0+Opera Mobile?

✔MDN

CacheStorage/delete

In all current engines.
Firefox41+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

CacheStorage/has

In all current engines.
Firefox41+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

CacheStorage/keys

In all current engines.
Firefox41+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

CacheStorage/match

In all current engines.
Firefox41+Safari11.1+Chrome54+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

CacheStorage/open

In all current engines.
Firefox41+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

CacheStorage

In all current engines.
Firefox41+Safari11.1+Chrome43+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Client/frameType

In all current engines.
Firefox44+Safari11.1+Chrome43+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Client/id

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Client/postMessage

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Client/type

In all current engines.
Firefox54+Safari11.1+Chrome60+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Client/url

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Client

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Clients/claim

In all current engines.
Firefox44+Safari11.1+Chrome42+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Clients/get

In all current engines.
Firefox45+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Clients/matchAll

In all current engines.
Firefox54+Safari11.1+Chrome42+
__________________________________________________________________

Opera29+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile29+

✔MDN

Clients/openWindow

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera38+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile41+

✔MDN

Clients

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ExtendableEvent/ExtendableEvent

In all current engines.
Firefox44+Safari11.1+Chrome41+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ExtendableEvent/waitUntil

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ExtendableEvent

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ExtendableMessageEvent/ExtendableMessageEvent

In all current engines.
Firefox44+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ExtendableMessageEvent/data

In all current engines.
Firefox44+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ExtendableMessageEvent/lastEventId

In all current engines.
Firefox44+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ExtendableMessageEvent/origin

In all current engines.
Firefox44+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ExtendableMessageEvent/ports

In all current engines.
Firefox44+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ExtendableMessageEvent/source

In all current engines.
Firefox44+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ExtendableMessageEvent

In all current engines.
Firefox44+Safari11.1+Chrome51+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

FetchEvent/FetchEvent

In all current engines.
Firefox44+Safari11.1+Chrome44+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

FetchEvent/clientId

In all current engines.
Firefox45+Safari11.1+Chrome49+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

FetchEvent/handled

In all current engines.
Firefox84+Safari16+Chrome86+
__________________________________________________________________

Opera?Edge86+
__________________________________________________________________

Edge (Legacy)?IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

FetchEvent/preloadResponse

In all current engines.
Firefox99+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

⚠MDN

FetchEvent/replacesClientId

In no current engines.
FirefoxNoneSafariNoneChromeNone
__________________________________________________________________

Opera?EdgeNone
__________________________________________________________________

Edge (Legacy)?IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

FetchEvent/request

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

FetchEvent/respondWith

In all current engines.
Firefox44+Safari11.1+Chrome42+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

FetchEvent/resultingClientId

In all current engines.
Firefox65+Safari16+Chrome72+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)?IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile50+

✔MDN

FetchEvent

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

NavigationPreloadManager/disable

In all current engines.
Firefox99+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

NavigationPreloadManager/enable

In all current engines.
Firefox99+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

NavigationPreloadManager/getState

In all current engines.
Firefox99+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

NavigationPreloadManager/setHeaderValue

In all current engines.
Firefox99+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

NavigationPreloadManager

In all current engines.
Firefox99+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Navigator/serviceWorker

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorker/postMessage

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorker/postMessage

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorker/scriptURL

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorker/state

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorker/statechange_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorker

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerContainer/controller

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerContainer/controllerchange_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerContainer/getRegistration

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerContainer/getRegistrations

In all current engines.
Firefox44+Safari11.1+Chrome45+
__________________________________________________________________

Opera27+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView40+Samsung Internet4.0+Opera Mobile27+

✔MDN

ServiceWorkerContainer/message_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerContainer/ready

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerContainer/register

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerContainer/startMessages

In all current engines.
Firefox64+Safari11.1+Chrome74+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)?IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile50+

✔MDN

ServiceWorkerContainer

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerGlobalScope/activate_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerGlobalScope/activate_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerGlobalScope/clients

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerGlobalScope/fetch_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerGlobalScope/install_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerGlobalScope/message_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerGlobalScope/message_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerGlobalScope/registration

In all current engines.
Firefox44+Safari11.1+Chrome42+
__________________________________________________________________

Opera26+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile26+

✔MDN

ServiceWorkerGlobalScope/skipWaiting

In all current engines.
Firefox44+Safari11.1+Chrome41+
__________________________________________________________________

Opera25+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile25+

✔MDN

ServiceWorkerGlobalScope

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera24+Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile24+

✔MDN

ServiceWorkerRegistration/active

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerRegistration/installing

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerRegistration/navigationPreload

In all current engines.
Firefox99+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet4.0+Opera Mobile?

✔MDN

ServiceWorkerRegistration/scope

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerRegistration/unregister

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerRegistration/update

In all current engines.
Firefox44+Safari11.1+Chrome45+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet4.0+Opera Mobile?

✔MDN

ServiceWorkerRegistration/updatefound_event

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerRegistration/updateViaCache

In all current engines.
Firefox57+Safari11.1+Chrome68+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerRegistration/waiting

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

ServiceWorkerRegistration

In all current engines.
Firefox44+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

WindowClient/focus

In all current engines.
Firefox44+Safari11.1+Chrome42+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

WindowClient/focused

In all current engines.
Firefox44+Safari11.1+Chrome42+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

WindowClient/navigate

In all current engines.
Firefox50+Safari16+Chrome49+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView42+Samsung Internet4.0+Opera Mobile?

✔MDN

WindowClient/visibilityState

In all current engines.
Firefox44+Safari11.1+Chrome42+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

WindowClient

In all current engines.
Firefox44+Safari11.1+Chrome42+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)17+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

caches

In all current engines.
Firefox41+Safari11.1+Chrome40+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)16+IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?

✔MDN

Headers/Service-Worker-Navigation-Preload

In all current engines.
Firefoxpreview+Safari15.4+Chrome59+
__________________________________________________________________

Opera?Edge79+
__________________________________________________________________

Edge (Legacy)18IENone
__________________________________________________________________

Firefox for Android?iOS Safari?Chrome for Android?Android
WebView?Samsung Internet?Opera Mobile?
