# advice API

This document explains the Advice API, control flow, and its contracts. It starts with a high-level overview
and then explains every API in detail.

## Overview

Advice is implemented with the help of AdviceKit and the Advice API through the implementation of an advice provider.
Advice providers are HTTP servers that implement the Advice API to describe which advice is delivered. The following
diagram illustrates who is issuing calls and in what phases.

![UML sequence diagram showing in what order the APIs are called](advice-flow.svg)

As can be seen above, the advice provider is called by the Steadybit agent in two phases:

- In the advice **registration phase**, Steadybit learns about the supported advice. Once this phase is completed, calling for advice will be scheduled with the agent.
- The agent's scheduler will call for advice in the **execution phase**.

The following sections explain the various API endpoints, their responsibilities, and structures in more detail.

## Index Response

As the name implies, this is the root of a advice provider and returns a list of supported Advice. Or,
more specifically, HTTP endpoints that the agent should call to learn more about them.

All paths will be resolved relative to the URL used to register the extension at the agent. For example, if `https://extension/some-path` was used to register and this endpoint returns `/advice/bad-dogs,` the agent will make the request to `https://extension/some-path/advice/bad-dogs.` This allows extensions to run behind reverse proxies, rewriting the path. 

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

A advice description is required for each advice. The HTTP endpoint serving the description is discovered through the
index endpoint.

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
        "summary": "When availability zone ${target.attr('aws.zone')} is failing, your service ${target.attr('k8s.pod.name')} is not available.",
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
            "name": "${target.attr('k8s.cluster-name')}/${target.attr('k8s.deployment')} faultless redundancy during single pod failure",
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
                            "${target.attr('k8s.cluster-name')}"
                          ]
                        },
                        {
                          "key": "k8s.namespace",
                          "operator": "EQUALS",
                          "values": [
                            "${target.attr('k8s.namespace')}"
                          ]
                        },
                        {
                          "key": "k8s.deployment",
                          "operator": "EQUALS",
                          "values": [
                            "${target.attr('k8s.deployment')}"
                          ]
                        },
                        {
                          "key": "k8s.container.name",
                          "operator": "EQUALS",
                          "values": [
                            "${target.attr('k8s.container.name')}"
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
        "summary": "When availability zone ${target.attr('aws.zones',0)} is failing, your service ${target.attr('k8s.pod.name')} is still available."
      }
    }
  }
}
```

### Advice Markdown Text and its substitutions

All texts of the advice definition can contain markdown texts. The markdown texts can contain placeholders which will be
substituted with the values of the target. This is done using Apache FreeMarker templates.

The experiment.json in the validation section will also be post-processed by FreeMarker.

The following syntax is supported:
* `${target.attr('key')}` --> will output the value of the given attribute. If no value is found, it will output `<unknown>`. If multiple values are found, it will output `<multiple values found>`
* `${target.attr('key', 'default')}` --> will output the value of the given attribute. If no value is found, it will output the specified default value. If multiple values are found, it will output `<multiple values found>`
* `${target.attr('key', 0)}` --> will output the value of the given attribute at the specified index. If no value is found, it will output `<unknown>`
* `${target.attr('key', 0, 'default')}` --> will output the value of the given attribute at the specified index. If no value is found, it will output the specified default value.

You can also use FreeMarker directives, e.g.:
* `<#list target.attrs('key') as item>${item}<#sep>, </#list>` -> will output all values as comma-separated list
* `<#if target.id.type=='com.steadybit.extension_kubernetes.kubernetes-deployment'>com.steadybit.extension_kubernetes.pod_count_check<#elseif target.id.type=='com.steadybit.extension_kubernetes.kubernetes-statefulset'>com.steadybit.extension_kubernetes.pod_count_check_statefulset<#else>com.steadybit.extension_kubernetes.pod_count_check_daemonset</#if>` -> conditional evaluation

`attr` and `attrs` are convenience methods to access the target attributes and provide proper defaults instead of [throwing exceptions](https://freemarker.apache.org/docs/app_faq.html#faq_picky_about_missing_vars). You can still use the default syntax, but we wouldn't recommend it.
* `${target.attributes['key'][0]}` --> will output the value of the given attribute. If no value is found, the template processing will fail and wil output the plain template text. 

### References

- [FreeMarker](https://freemarker.apache.org/docs/index.html)
- [Go API](https://github.com/steadybit/advice-kit/tree/main/go/advice_kit_api): `adviceDefinition`
- [OpenAPI Schema](https://github.com/steadybit/advice-kit/tree/main/openapi): `weakspotDescription`

