# Recipe API Specification

**API Version:** `ramendr.openshift.io/v1alpha1`  
**Kind:** `Recipe`  
**Source:** [`api/v1alpha1/recipe_types.go`](https://github.com/RamenDR/recipe/blob/main/api/v1alpha1/recipe_types.go)

---

## Overview

A `Recipe` is a Kubernetes custom resource that defines application-aware backup and restore
workflows for ACM Discovered Applications. It describes how application resources should be
grouped, which PVCs should be protected, which application-specific hooks should be executed,
and in what order operations should run during disaster recovery.

---

## Top-Level Structure

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: <string>
  namespace: <string>
spec:
  appType: <string>
  groups: [...]
  volumes: {...}
  hooks: [...]
  jobs: [...]
  workflows: [...]
```

---

## `metadata`

| Field | Type | Description |
|---|---|---|
| `metadata.name` | `string` | Name of the Recipe resource. Used to reference the Recipe during DR operations. |
| `metadata.namespace` | `string` | Namespace in which the Recipe is created. |

---

## `spec`

Root specification of the Recipe.

| Field | Type | Required | Description |
|---|---|---|---|
| `appType` | `string` | Yes | Type of application the Recipe is designed for. Currently matched against the name of the application CR. |
| `groups` | `[]Group` | No | List of one or more resource groups that define the Kubernetes resources participating in the DR workflow. |
| `volumes` | `Group` | No | Defines which PersistentVolumeClaims (PVCs) to protect. Uses the same `Group` schema. |
| `hooks` | `[]Hook` | No | List of one or more hooks (exec, check, scale, job) to be executed as part of the workflow. |
| `jobs` | `[]map[string]*string` | No | List of Kubernetes Job manifests (as inline YAML strings) referenced by Job Hooks. |
| `workflows` | `[]Workflow` | No | Ordered sequences of groups and hooks that define the backup and restore workflows. |

---

## `spec.groups[]` — Group

Groups define logical collections of Kubernetes resources that participate in the DR workflow.

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | `string` | Yes | Unique name of the group within the Recipe. |
| `parent` | `string` | No | Name of the parent group in the associated Application CR. If unspecified, the implicit default group of the Application CR is used. |
| `backupRef` | `string` | No | Used for groups in restore workflows to refer to another group used in backup workflows. Allows selective restore of resources captured by a specific backup group. |
| `type` | `string` | Yes | Determines the type of group. Allowed values: `volume`, `resource`. |
| `includedResourceTypes` | `[]string` | No | List of Kubernetes resource types to include. If unspecified, all resource types are included. |
| `excludedResourceTypes` | `[]string` | No | List of Kubernetes resource types to exclude. |
| `labelSelector` | `LabelSelector` | No | Selects resources based on their labels. |
| `nameSelector` | `string` | No | If specified, a resource's object name must match this expression. Valid for volume groups only. |
| `selectResource` | `string` | No | Specifies which resource type `labelSelector` and `nameSelector` apply to for selecting PVCs. Valid for volume groups only. Allowed values: `pvc`, `pod`, `deployment`, `statefulset`. Default: `pvc`. |
| `includeClusterResources` | `*bool` | No | Whether to include cluster-scoped resources. If `nil` or `true`, cluster-scoped resources associated with included namespace-scoped resources are included. |
| `includedNamespacesByLabel` | `LabelSelector` | No | Selects namespaces to include by label. |
| `includedNamespaces` | `[]string` | No | Explicit list of namespaces to include. |
| `excludedNamespaces` | `[]string` | No | List of namespaces to exclude. |
| `restoreStatus` | `GroupRestoreStatus` | No | When set, instructs the restore operation to also restore the status of specified resource types. Use `*` to restore status for all resource types. |
| `essential` | `*bool` | No | Defaults to `true`. If set to `false`, a failure in this group is not treated as fatal. |
| `restoreOverwriteResources` | `*bool` | No | Whether to overwrite existing resources during restore. Defaults to `false`. |

### `GroupRestoreStatus`

| Field | Type | Description |
|---|---|---|
| `includedResources` | `[]string` | Resource types whose status should be restored. If unspecified, all types are included. |
| `excludedResources` | `[]string` | Resource types whose status should not be restored. |

---

## `spec.volumes` — Group (Volume Selection)

Uses the same `Group` schema as `spec.groups`. When `spec.volumes.labelSelectors` is specified,
it takes precedence over the `labelSelector` configured in the DRPlacementControl (DRPC).

Typical usage:

```yaml
spec:
  volumes:
    - name: my-app-volumes
      labelSelector:
        matchLabels:
          app: my-app
```

---

## `spec.hooks[]` — Hook

Hooks define application-specific operations executed as part of a workflow.

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | `string` | Yes | Unique name of the hook within the Recipe. |
| `namespace` | `string` | Yes | Namespace in which the hook operates. |
| `type` | `string` | Yes | Hook type. Allowed values: `exec`, `scale`, `check`, `job`. |
| `selectResource` | `string` | No | Kubernetes resource type the hook applies to (e.g., `deployment`, `statefulset`). |
| `labelSelector` | `LabelSelector` | No | If specified, the hook applies to resources matching this label selector. |
| `nameSelector` | `string` | No | If specified, the hook applies to resources whose name matches this expression (supports regex). |
| `singlePodOnly` | `bool` | No | If `true`, executes the command on a single matching pod rather than all matching pods. |
| `onError` | `string` | No | Default behavior when an operation fails. Allowed values: `fail`, `continue`. Default: `fail`. |
| `timeout` | `int` | No | Default timeout in seconds for operations. Default: `30`. |
| `ops` | `[]Operation` | No | Set of exec operations the hook can invoke. Used with `type: exec`. |
| `chks` | `[]Check` | No | Set of check conditions the hook can evaluate. Used with `type: check`. |
| `essential` | `*bool` | No | Defaults to `true`. If set to `false`, a failure is not treated as fatal. |
| `jobs` | `[]JobDetails` | No | List of job references the hook can invoke. Used with `type: job`. |
| `skipHookIfNotPresent` | `bool` | No | If `true`, the hook is skipped when the target resource is not present. Default: `false`. |

> **Note:** `nameSelector` and `labelSelector` use logical **OR** — if both are specified,
> the hook applies to resources matching either selector.

---

## `spec.hooks[].ops[]` — Operation (Exec Hook)

Used with `type: exec`. Defines a command to execute inside a container.

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | `string` | Yes | Unique name of the operation within the hook. |
| `container` | `string` | No | Name of the container in which to execute the command. |
| `command` | `string` | Yes | The command to execute. Must be non-empty. |
| `onError` | `string` | No | How to handle a non-zero exit code. Defaults to `fail`. |
| `timeout` | `int` | No | Maximum time in seconds to wait for the command to complete. |
| `inverseOp` | `string` | No | Name of another operation that reverts the effect of this one (e.g., `quiesce` ↔ `unquiesce`). |

---

## `spec.hooks[].chks[]` — Check (Check Hook)

Used with `type: check`. Defines a condition to validate before the workflow proceeds.

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | `string` | Yes | Unique name of the check within the hook. |
| `condition` | `string` | No | JSONPath-based expression to evaluate against the selected resource (e.g., `{$.spec.replicas} == {$.status.readyReplicas}`). |
| `onError` | `string` | No | How to handle a condition that does not become `true`. Defaults to `fail`. |
| `timeout` | `int` | No | Maximum time in seconds to wait for the condition to be satisfied. |

---

## `spec.hooks[].jobs[]` — JobDetails (Job Hook)

Used with `type: job`. References a Kubernetes Job manifest defined in `spec.jobs`.

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | `string` | Yes | Name of the job manifest defined in `spec.jobs` to execute. |
| `onError` | `string` | No | How to handle a job that returns a non-zero exit code. Defaults to `fail`. |
| `timeout` | `int` | No | Maximum time in seconds to wait for the job to complete. |
| `inverseOp` | `string` | No | Name of another operation that reverts the effect of this job. |
| `forceCreate` | `*bool` | No | If `true`, a new Job is created even if a Job with the same name already exists. Default: `false`. |

---

## `spec.jobs[]`

A list of Kubernetes Job manifests (inline YAML strings) keyed by job name. These manifests are
referenced by `JobDetails.name` inside Job Hooks.

```yaml
spec:
  jobs:
    - my-job-name: |
        apiVersion: batch/v1
        kind: Job
        metadata:
          name: my-job-name
          namespace: my-namespace
        spec:
          template:
            spec:
              restartPolicy: Never
              containers:
                - name: my-container
                  image: registry.k8s.io/busybox:latest
                  command: ["/bin/sh", "-c", "echo hello"]
```

---

## `spec.workflows[]` — Workflow

Workflows define the ordered sequence of groups and hooks executed during backup or restore.

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | `string` | Yes | Name of the workflow. Reserved names: `backup`, `restore`. |
| `sequence` | `[]map[string]string` | Yes | Ordered list of steps. Each step references either a `group` or a `hook`. Format: `group: <group-name>` or `hook: <hook-name>/<op-or-check-name>`. |
| `failOn` | `string` | No | Failure policy for the workflow. Allowed values: `any-error`, `essential-error`, `full-error`. Default: `any-error`. |

### `failOn` Values

| Value | Behaviour |
|---|---|
| `any-error` | Stop the workflow if any step fails (default). |
| `essential-error` | Stop only if an essential step fails; non-essential failures are tolerated. |
| `full-error` | Continue executing all steps regardless of failures; fail only at the end. |

### Workflow Sequence Step Formats

```yaml
sequence:
  - group: <group-name>                          # Process a resource group
  - hook: <hook-name>/<operation-name>           # Exec or Check hook operation
  - hook: <hook-name>/<check-name>               # Check hook condition
```

---

## Reserved Workflow Names

| Name | Purpose |
|---|---|
| `backup` | Executed during capture (backup) operations. |
| `restore` | Executed during recovery (restore) operations. |

---

## Scale Hook — `replicas` Field

Scale Hooks modify the replica count of selected workloads. The desired replica count is
specified directly on the hook via the `replicas` field (not inside `ops`):

```yaml
hooks:
  - name: scale-down
    type: scale
    namespace: my-namespace
    selectResource: deployment
    nameSelector: "my-app.*"
    timeout: 300
    onError: fail
    replicas: 0
```

> **Note:** The `replicas` field is not present in the typed Go struct but is used in YAML
> configuration for Scale Hooks to set the desired replica count.

---

## Full Schema Reference (Go Types)

```
Recipe
  .metadata
    .name          string
    .namespace     string
  .spec            RecipeSpec
    .appType       string
    .groups[]      Group
    .volumes       Group
    .hooks[]       Hook
    .jobs[]        map[string]*string
    .workflows[]   Workflow

Group
  .name                       string
  .parent                     string
  .backupRef                  string
  .type                       string  (volume | resource)
  .includedResourceTypes[]    string
  .excludedResourceTypes[]    string
  .labelSelector              LabelSelector
  .nameSelector               string
  .selectResource             string  (pvc | pod | deployment | statefulset)
  .includeClusterResources    *bool
  .includedNamespacesByLabel  LabelSelector
  .includedNamespaces[]       string
  .excludedNamespaces[]       string
  .restoreStatus              GroupRestoreStatus
    .includedResources[]      string
    .excludedResources[]      string
  .essential                  *bool
  .restoreOverwriteResources  *bool

Hook
  .name                  string
  .namespace             string
  .type                  string  (exec | scale | check | job)
  .selectResource        string
  .labelSelector         LabelSelector
  .nameSelector          string
  .singlePodOnly         bool
  .onError               string  (fail | continue)
  .timeout               int
  .ops[]                 Operation
    .name                string
    .container           string
    .command             string
    .onError             string
    .timeout             int
    .inverseOp           string
  .chks[]                Check
    .name                string
    .condition           string
    .onError             string
    .timeout             int
  .essential             *bool
  .jobs[]                JobDetails
    .name                string
    .onError             string
    .timeout             int
    .inverseOp           string
    .forceCreate         *bool
  .skipHookIfNotPresent  bool

Workflow
  .name       string
  .sequence[] map[string]string
  .failOn     string  (any-error | essential-error | full-error)
```
