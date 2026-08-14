# Recipe API Reference

## Minimal Valid Recipe

The following represents the minimum structure of a Recipe.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: my-recipe

spec:
  appType: my-application

  groups:
    - name: application-resources
      type: resource

  workflows:
    - name: backup
      sequence:
        - group: application-resources

    - name: restore
      sequence:
        - group: application-resources
```

Every generated Recipe should begin with this structure.

Additional sections such as `volumes`, `hooks`, and `jobs` should be added only when required by the application.

## Complete Recipe Structure

A Recipe is composed of the following top-level sections.

| Section | Purpose |
|----------|---------|
| metadata | Identifies the Recipe. |
| appType | Identifies the application type. |
| groups | Defines Resource Groups. |
| volumes | Selects PersistentVolumeClaims. |
| hooks | Defines application operations. |
| jobs | Defines reusable Kubernetes Jobs. |
| workflows | Defines backup and restore execution order. |

The remaining sections describe each component independently.

# Resource Group

## Purpose

A Resource Group defines a logical collection of Kubernetes resources that participate in disaster recovery.

---

## Required Fields

| Field | Description |
|---------|-------------|
| name | Unique group name. |
| type | resource or volume |

---

## Commonly Used Fields

- includedNamespaces
- labelSelector
- includedResourceTypes
- excludedResourceTypes
- backupRef

---

## Optional Fields

- essential
- restoreStatus
- includeClusterResources
- restoreOverwriteResources

---

## Typical Resource Group

```yaml
- name: postgres-resources
  type: resource
  includedNamespaces:
    - postgres
  labelSelector:
    matchLabels:
      app: postgres
```

# Volume Group

## Purpose

A Volume Group identifies PersistentVolumeClaims (PVCs) that are protected as part of the application.

Volume Groups use the same `Group` schema as Resource Groups, with:

```yaml
type: volume
```

A Volume Group is defined under:

```yaml
spec:
  volumes:
```

## Required Fields

| Field  | Description               |
| ------ | ------------------------- |
| `name` | Unique volume group name. |
| `type` | Must be `volume`.         |

## Commonly Used Fields

* `includedNamespaces`
* `labelSelector`
* `nameSelector`
* `includedResourceTypes`
* `excludedResourceTypes`
* `backupRef`

## Label-Based Volume Selection

When PVC labels are known, select the PVCs using `labelSelector`:

```yaml
volumes:
  - name: postgres-volumes
    type: volume
    includedNamespaces:
      - postgres
    labelSelector:
      matchLabels:
        app: postgres
```

The labels used for volume selection must identify the PVCs that belong to the application.

Do not assume that labels on a Deployment or StatefulSet are also present on its PVCs unless that has been verified.

## Name-Based Volume Selection

When name-based selection is required, use `nameSelector`:

```yaml
volumes:
  - name: postgres-volumes
    type: volume
    includedNamespaces:
      - postgres
    nameSelector: postgres-data
```

`nameSelector` is a string.

A complete resource name selects a single resource. A regular expression can select multiple matching resources.

## Generation Notes

* Define Volume Groups only when the application has PersistentVolumeClaims that require protection.
* Use `type: volume`.
* Use `labelSelector` for label-based selection.
* Use `nameSelector` for name-based selection.
* Do not use a separate PVC-specific structure.
* Volume Groups are defined under `spec.volumes`.
* Volume Groups are not workflow steps. The workflow does not reference a Volume Group merely to cause PVC protection.

# Exec Hook

## Purpose

Execute commands inside containers running in application Pods.

---

## Required Fields

| Field |
|--------|
| name |
| type |
| namespace |
| ops |

---

## Commonly Used Fields

- selectResource
- labelSelector
- nameSelector
- timeout
- onError

---

## Operation Structure

```yaml
ops:
  - name: checkpoint
    container: postgres
    command: '["psql","-U","postgres","-c","CHECKPOINT"]'
    timeout: 60
```

# Check Hook

## Purpose

A Check Hook validates that the selected Kubernetes resources have reached the expected state before the workflow proceeds.

Check Hooks are commonly used after restore to verify that workloads are healthy and ready.

---

## Required Fields

| Field       | Description                           |
| ----------- | ------------------------------------- |
| `name`      | Unique Hook name.                     |
| `type`      | Must be `check`.                      |
| `namespace` | Namespace in which the Hook executes. |
| `chks`      | One or more validation conditions.    |

---

## Commonly Used Fields

* `selectResource`
* `labelSelector`
* `nameSelector`
* `timeout`
* `onError`

---

## Check Structure

```yaml
- name: postgres-ready
  type: check
  namespace: postgres
  selectResource: deployment
  nameSelector: postgres
  timeout: 120
  onError: fail
  chks:
    - name: replicasReady
      condition: "{$.spec.replicas} == {$.status.readyReplicas}"
      timeout: 600
      onError: fail
```

---

## Generation Notes

* Check Hooks are typically used after restore.
* Every Check Hook should contain one or more checks.
* A workflow references a check using:

```text
hook: <hook-name>/<check-name>
```

---

# Scale Hook

## Purpose

A Scale Hook changes the replica count of Kubernetes workloads during backup or restore.

Typical use cases include scaling workloads down before protecting persistent storage and scaling them back up afterward.

---

## Required Fields

| Field       | Description                           |
| ----------- | ------------------------------------- |
| `name`      | Unique Hook name.                     |
| `type`      | Must be `scale`.                      |
| `namespace` | Namespace in which the Hook executes. |
| `replicas`  | Desired replica count.                |

---

## Commonly Used Fields

* `selectResource`
* `labelSelector`
* `nameSelector`
* `timeout`
* `onError`

---

## Typical Scale Hook

```yaml
- name: scale-down
  type: scale
  namespace: my-app
  selectResource: deployment
  nameSelector: my-app
  replicas: 0
  timeout: 300
  onError: fail
```

---

## Generation Notes

* Scale Hooks are commonly used before volume protection.
* Scale Hooks should target workload controllers rather than Pods.
* The desired replica count is specified using the `replicas` field.

---

# Job Hook

## Purpose

A Job Hook executes one or more Kubernetes Jobs as part of a workflow.

Use Job Hooks when application-specific operations require a dedicated Kubernetes Job rather than commands executed inside an existing container.

---

## Required Fields

| Field       | Description                           |
| ----------- | ------------------------------------- |
| `name`      | Unique Hook name.                     |
| `type`      | Must be `job`.                        |
| `namespace` | Namespace in which the Hook executes. |
| `jobs`      | One or more Job references.           |

---

## Commonly Used Fields

* `timeout`
* `onError`

---

## Job Reference Structure

```yaml
jobs:
  - name: backup-preparation
    timeout: 300
    onError: fail
    forceCreate: true
```

---

## Generation Notes

* Job Hooks reference Job Definitions defined in `spec.jobs`.
* Every referenced Job must exist.
* A workflow references a Job Hook using:

```text
hook: <hook-name>/<job-name>
```

---

# Job Definitions

## Purpose

Job Definitions contain reusable Kubernetes Job manifests executed by Job Hooks.

Each Job is defined once and may be referenced by multiple Job Hooks.

---

## Required Fields

| Field        | Description                 |
| ------------ | --------------------------- |
| Job name     | Unique Job identifier.      |
| Job manifest | Inline Kubernetes Job YAML. |

---

## Typical Job Definition

```yaml
jobs:
  - backup-preparation: |
      apiVersion: batch/v1
      kind: Job
      metadata:
        name: backup-preparation
      spec:
        template:
          spec:
            restartPolicy: Never
            containers:
              - name: backup-preparation
                image: registry.k8s.io/busybox:latest
                command:
                  - /bin/sh
                  - -c
                  - echo "Preparing backup"
```

---

## Generation Notes

* Job Definitions are reusable.
* Unreferenced Job Definitions should not be generated.
* Every Job Definition should be referenced by at least one Job Hook.

---

# Workflow

## Purpose

A Workflow defines the order in which Resource Groups and Hooks are executed during backup or restore.

Every Recipe should define both backup and restore workflows.

---

## Required Fields

| Field      | Description                            |
| ---------- | -------------------------------------- |
| `name`     | Workflow name (`backup` or `restore`). |
| `sequence` | Ordered list of workflow steps.        |

---

## Commonly Used Fields

* `failOn`

---

## Backup Workflow

```yaml
- name: backup
  failOn: any-error
  sequence:
    - hook: prepare-backup/backup-preparation
    - hook: postgres-exec/checkpoint
    - group: postgres-resources
```

---

## Restore Workflow

```yaml
- name: restore
  failOn: any-error
  sequence:
    - group: postgres-resources
    - hook: postgres-ready/replicasReady
```

---

## Workflow Step Types

A workflow step may reference either a Resource Group or a Hook.

Resource Group:

```yaml
- group: postgres-resources
```

Hook:

```yaml
- hook: postgres-exec/checkpoint
```

---

## Generation Notes

* Every Recipe should define both backup and restore workflows.
* Workflow steps execute sequentially.
* Every referenced Resource Group and Hook must exist.
* Backup workflows typically prepare the application before protecting resources.
* Restore workflows typically restore resources before performing validation.

---

# Metadata

## Purpose

Metadata identifies the Recipe Custom Resource.

---

## Required Fields

| Field           | Description  |
| --------------- | ------------ |
| `metadata.name` | Recipe name. |

---

## Commonly Used Fields

* `metadata.namespace`

---

## Typical Metadata

```yaml
metadata:
  name: postgres-recipe
  namespace: ramen-system
```

---

## Generation Notes

* Choose a descriptive Recipe name.
* The namespace identifies where the Recipe object is created and is not necessarily the application's namespace.

---

# Complete Recipe Hierarchy

```
Recipe
├── metadata
│   ├── name
│   └── namespace
│
└── spec
    ├── appType
    ├── groups
    ├── volumes
    ├── hooks
    │   ├── exec
    │   ├── check
    │   ├── scale
    │   └── job
    ├── jobs
    └── workflows
        ├── backup
        └── restore
```

---

# API Reference Summary

A complete Recipe is constructed from the following building blocks:

1. Metadata
2. Application type
3. Resource Groups
4. Volume Selection (when required)
5. Hooks (when required)
6. Job Definitions (when required)
7. Backup Workflow
8. Restore Workflow

Begin with the minimal valid Recipe structure and add additional components only when they are required by the application's disaster recovery requirements.

`spec.volumes` defines PVC protection independently of workflow sequencing. Volume Groups are not referenced as workflow steps.
