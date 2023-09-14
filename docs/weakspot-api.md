# weakspot API

This document explains the weakspot API, control flow and the contracts behind it. It starts with a high-level overview and then explains every API in detail.

## Overview

Weakspots are implemented with the help of weakspotKit and the weakspot API through the implementation of a weakspot provider. weakspot providers are HTTP servers implementing the weakspot API to describe which weakspots are delivered. The following diagram illustrates who is issuing calls and in what phases.

![UML sequence diagram showing in what order the APIs are called](weakspot-flow.svg)

As can be seen above, the weakspot provider is called by the Steadybit agent in two phases:

- In the weakspot registration phase, Steadybit learns about the supported weakspots. Once this phase is completed, weakspots will be
  scheduled within the agent.
- The weakspot will be called by the agent's scheduler in the execution phase.

The following sections explain the various API endpoints, their responsibilities and structures in more detail.

## Index Response

As the name implies, this is the root of a weakspot provider and returns a list of supported weakspots. Or,
more specifically, HTTP endpoints that the agent should call to learn more about them.

This endpoint needs to be [registered with Steadybit agents](./weakspot-registration.md).

### Example

```json
// Request: GET /

// Response: 200
{
  "weakspots": [
    {
      "path": "/weakspots/com.steadybit.extension_aws.weakspot.aws-single-zone"
    },
    {
      "path": "/weakspots/bad-dogs"
    }
  ]
}
```

### References

- [Go API](https://github.com/steadybit/weakspot-kit/tree/main/go/weakspot_kit_api): `weakspotList`
- [OpenAPI Schema](https://github.com/steadybit/weakspot-kit/tree/main/openapi): `weakspotList`

## Weakspot Description

A weakspot description is required for each weakspot. The HTTP endpoint serving the description is discovered through the index endpoint.

### Example

```json
// Request: GET /weakspots/com.steadybit.extension_aws.weakspot.aws-single-zone

// Response: 200
{
  "id": "com.steadybit.extension_aws.weakspot.aws-single-zone",
  "version": "1.0.0",
  "label": "AWS Single Zone",
  "icon": "data:image/svg+xml;utf8,<svg ...>",
  "assesmentBaseQuery": "target.type = com.steadybit.extension_kubernetes.kubernetes-deployment",
  "assesmentQueryAddon": "aws.zone > 0 AND aws.zone < 2",
  "technology": "AWS",
  "experiments": [
    {
      "name": "{{k8s.cluster-name}}/{{k8s.deployment}} faultless redundancy during single pod failure",
      "hypothesis": "When a single pod fails {{successRate}}% of all requests to {{httpEndpoint}} are successful",
      "environment": "{{environmentName}}",
      "team": "{{teamKey}}",
      "lanes": {
        "steps": [
          {
            "type": "action",
            "ignoreFailure": false,
            "parameters": {
              "duration": "60s",
              "cpuLoad": 100,
              "workers": 0
            },
            "actionType": "com.steadybit.extension_container.stress_cpu",
            "radius": {
              "targetType": "com.steadybit.extension_container.container",
              "predicate": {
                "operator": "AND",
                "predicates": [
                  {
                    "key": "k8s.cluster-name",
                    "operator": "EQUALS",
                    "values": [
                      "{k8s.cluster-name}"
                    ]
                  },
                  {
                    "key": "k8s.namespace",
                    "operator": "EQUALS",
                    "values": [
                      "{k8s.namespace}"
                    ]
                  },
                  {
                    "key": "k8s.deployment",
                    "operator": "EQUALS",
                    "values": [
                      "{k8s.deployment}"
                    ]
                  },
                  {
                    "key": "k8s.container.name",
                    "operator": "EQUALS",
                    "values": [
                      "{k8s.container.name}"
                    ]
                  }
                ]
              },
              "query": null,
              "percentage": 50
            }
          }
        ]
      }
    }
  ],
  "finding": "When availability zone {aws.zones?.[0]} is failing, your service {k8s.pod.name} is not available.",
  "looksGood": "When availability zone {aws.zones?.[0]} is failing, your service {k8s.pod.name} is still available.",
  "guidance": "It is recommended to always split its components into different zones so that in case of a failure of one..",
  "instructions": "It is recommended to always split its components into different zones so that in case of a failure of one..",
  "resources": [
    {
      "name": "AWS Availability Zones",
      "url": "https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html"
    }
  ]
}
```

### References

- [Go API](https://github.com/steadybit/weakspot-kit/tree/main/go/weakspot_kit_api): `weakspotDescription`
- [OpenAPI Schema](https://github.com/steadybit/weakspot-kit/tree/main/openapi): `weakspotDescription`

