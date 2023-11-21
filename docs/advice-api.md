# advice API

This document explains the Advice API, control flow and the contracts behind it. It starts with a high-level overview and then explains every API in detail.

## Overview

Advice are implemented with the help of AdviceKit and the Advice API through the implementation of a advice provider. Advice providers are HTTP servers implementing the Advice API to describe which advice are delivered. The following diagram illustrates who is issuing calls and in what phases.

![UML sequence diagram showing in what order the APIs are called](advice-flow.svg)

As can be seen above, the advice provider is called by the Steadybit agent in two phases:

- In the advice registration phase, Steadybit learns about the supported advice. Once this phase is completed, advice will be
  scheduled within the agent.
- The advice will be called by the agent's scheduler in the execution phase.

The following sections explain the various API endpoints, their responsibilities and structures in more detail.

## Index Response

As the name implies, this is the root of a advice provider and returns a list of supported Advice. Or,
more specifically, HTTP endpoints that the agent should call to learn more about them.

This endpoint needs to be [registered with Steadybit agents](./advice-registration.md).

### Example

```json
// Request: GET /

// Response: 200
{
  "advice": [
    {
      "path": "/advice/com.steadybit.extension_aws.advice.aws-single-zone"
    },
    {
      "path": "/advice/bad-dogs"
    }
  ]
}
```

### References

- [Go API](https://github.com/steadybit/advice-kit/tree/main/go/advice_kit_api): `adviceList`
- [OpenAPI Schema](https://github.com/steadybit/advice-kit/tree/main/openapi): `adviceList`

## Advice Description

A advice description is required for each advice. The HTTP endpoint serving the description is discovered through the index endpoint.

### Example

```json
// Request: GET /advice/com.steadybit.extension_aws.advice.aws-single-zone

// Response: 200
{
  "id": "com.steadybit.extension_aws.advice.aws-single-zone",
  "version": "1.0.0",
  "label": "AWS Single Zone",
  "icon": "data:image/svg+xml;utf8,<svg ...>",
  "assessmentQueryApplicable": "target.type = com.steadybit.extension_kubernetes.kubernetes-deployment",
  "tags": [
    "AWS"
  ],
  "status": {
    "actionNeeded": {
      "assessmentQuery": "aws.zone > 0 AND aws.zone < 2",
      "description": {
        "summary": "When availability zone ${target.aws.zone} is failing, your service ${target.k8s.pod.name} is not available.",
        "motivation": "It is recommended to always split its components into different zones so that in case of a failure of one.",
        "instruction": "Specify the upper limit to be used by defining the   limits   property in your kubernetes manifest: ```...```"
      }
    },
    "validationNeeded": {
      "description": {
        "summary": "You already took action and configured it. Validate your configuration via the experiment."
      },
      "validation": [
        {
          "id": "com.steadybit.extension_aws.advice.aws-single-zone.1",
          "name": "AWS Single Zone Template",
          "type": "experiment",
          "description": "...",
          "experiment": {
            "name": "${target.k8s.cluster-name}/${target.k8s.deployment} faultless redundancy during single pod failure",
            "hypothesis": "When a single pod fails of all requests to an endpoint are successful",
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
                            "${target.k8s.cluster-name}"
                          ]
                        },
                        {
                          "key": "k8s.namespace",
                          "operator": "EQUALS",
                          "values": [
                            "${target.k8s.namespace}"
                          ]
                        },
                        {
                          "key": "k8s.deployment",
                          "operator": "EQUALS",
                          "values": [
                            "${target.k8s.deployment}"
                          ]
                        },
                        {
                          "key": "k8s.container.name",
                          "operator": "EQUALS",
                          "values": [
                            "${target.k8s.container.name}"
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
        },
        {
          "id": "com.steadybit.extension_aws.advice.aws-single-zone.2",
          "name": "Check me",
          "description": "Really check me",
          "type": "text"
        }
      ]
    },
    "implemented": {
      "description": {
        "summary": "When availability zone ${target.aws.zones?.[0]} is failing, your service ${target.k8s.pod.name} is still available."
      }
    }
  }
}
```

### Advice Markdown Text and its substitutions

The `description` field of the advice description is a map of markdown texts. The markdown texts can contain placeholders which will be substituted by the agent with the values of the target. The placeholders are enclosed in `${target.KEY}`. 
The following placeholders are supported:
* ${target.key} --> replaces with the value of the attribute with the given key as `*value*`
* ${target.key:normal} --> replaces with the value of the attribute with the given key as `value`
* ${target.key[0]} --> replaces with the value of the first attribute with the given key as `*value*`
* ${target.key[1]} --> replaces with the value of the second attribute with the given key as `*value2*`
* ${target.key[]} --> replaces with the value of the second attribute with the given key as `*value, value2*`
* ${target.key[1]:normal} --> replaces with the value of the second attribute with the given key as `value2`
* ${target.key:length} --> replaces with the value count of the attribute as e.g. `*42*`
* ${target.key:length:normal} --> replaces with the value count of the attribute as e.g. `42`

### References

- [Go API](https://github.com/steadybit/advice-kit/tree/main/go/advice_kit_api): `adviceDefinition`
- [OpenAPI Schema](https://github.com/steadybit/advice-kit/tree/main/openapi): `weakspotDescription`

