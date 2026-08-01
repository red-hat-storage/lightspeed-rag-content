# ODF Disaster Recovery — Recipe Examples

This document provides step-by-step examples that demonstrate how to create a Recipe for
protecting an ACM discovered application. The examples introduce one concept at a time and
progressively build toward a complete Recipe.

**Assumed application** (deployed in the `recipe-demo` namespace):

- Two Deployments: `busybox-1`, `busybox-2`
- Two PersistentVolumeClaims: `busybox-1-pvc`, `busybox-2-pvc`
- Both PVCs and Deployments are labeled `app: busybox`

---

## Step 1: Defining Resource Groups

**Objective:** Define a resource group that identifies the application resources to be protected.
A resource group selects resources based on namespaces, label selectors, and optionally resource
types.

```yaml
spec:
  appType: busybox

  groups:
    - name: busybox-resources
      backupRef: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox
```

**Explanation:**

- `appType` identifies the application type associated with the Recipe.
- `includedNamespaces` limits resource selection to the `recipe-demo` namespace.
- `labelSelector` selects all resources labeled `app=busybox`.
- `type: resource` indicates that the group contains Kubernetes resources.
- `backupRef` provides a logical identifier for the resources captured from this group. During
  recovery, it can be used to selectively restore only the resources associated with the group
  instead of restoring all captured resources.

If `includedResourceTypes` is not specified, all supported Kubernetes resources matching the
namespace and label selector are included. To restrict the group to specific resource types:

```yaml
includedResourceTypes:
  - deployments
  - services
```

---

## Step 2: Defining Backup and Restore Workflows

**Objective:** Define the capture and recover workflows. Workflows orchestrate the disaster recovery
process by defining the sequence in which resource groups and hooks are executed.

```yaml
spec:
  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - group: busybox-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: busybox-resources
```

**Explanation:**

- `backup` defines the workflow executed during the backup operation.
- `restore` defines the workflow executed during the restore operation.
- `sequence` specifies the ordered list of steps to execute as part of the workflow.
- `group` references the resource group defined in Step 1. All resources associated with the
  referenced group are processed.
- `failOn: any-error` instructs the workflow to stop execution if any step encounters an error.

---

## Step 3: Configuring PVC Selection

**Objective:** Configure the `spec.volumes` section to identify the PersistentVolumeClaims (PVCs)
that should be protected as part of the application.

```yaml
spec:
  volumes:
    - name: busybox-volumes
      labelSelector:
        matchLabels:
          app: busybox
```

**Explanation:**

- `name` uniquely identifies the volume selection within the Recipe.
- `labelSelector` selects all PVCs labeled `app=busybox`.

If `spec.volumes.labelSelector` is specified, it takes precedence over the PVC `labelSelector`
configured in the DRPlacementControl (DRPC). Otherwise, the PVC `labelSelector` in the
DRPlacementControl is used. Using label selectors enables the same Recipe to be reused across
multiple deployments, provided PVCs follow a consistent labeling strategy.

---

## Step 4: Adding a Check Hook

**Objective:** Define a Check Hook that validates the health of application resources before the
workflow proceeds.

```yaml
spec:
  hooks:
    - name: validate-deployments
      type: check
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: "busybox.*"
      timeout: 300
      onError: fail
      chks:
        - name: check-replicas
          condition: '{$.spec.replicas} == {$.status.readyReplicas}'
```

**Explanation:**

- `type: check` defines a Check Hook.
- `selectResource` specifies the Kubernetes resource type on which the check is performed.
- `nameSelector` identifies the resources on which the hook is executed. It accepts a full resource
  name or a regular expression. `busybox.*` matches both `busybox-1` and `busybox-2` Deployments.
- `timeout` specifies the maximum time (in seconds) to wait for the check to succeed.
- `onError: fail` causes the workflow to terminate if the validation fails or times out.
- `chks` contains one or more validation conditions to evaluate.
- `condition` verifies that the number of desired replicas matches the number of ready replicas,
  ensuring that each selected Deployment is healthy.

Resources can be selected using either `nameSelector` or `labelSelector`. If both are specified,
the hook applies to resources matching either selector (logical OR).

---

## Step 5: Adding an Exec Hook

**Objective:** Define an Exec Hook that executes one or more commands inside the containers of
selected application resources. Exec Hooks are typically used to prepare an application for backup
by performing application-specific operations.

```yaml
spec:
  hooks:
    - name: flush-database
      type: exec
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: "busybox.*"
      timeout: 300
      onError: fail
      ops:
        - name: flush
          container: busybox
          command: >
            ["/bin/sh", "-c", "sync"]
```

**Explanation:**

- `type: exec` defines an Exec Hook.
- `selectResource` specifies the Kubernetes resource type on which the hook is executed.
- `nameSelector` identifies the target resources. `busybox.*` matches `busybox-1` and `busybox-2`.
- `timeout` specifies the maximum time (in seconds) to wait for the operation to complete.
- `onError: fail` terminates the workflow if the command fails or times out.
- `ops` contains one or more operations to execute.
- `container` specifies the container in which the command is executed.
- `command` defines the command to execute inside the selected container.

Although this example uses a simple `sync` command, Exec Hooks are commonly used to flush pending
database writes, create checkpoints, pause application writes, or execute vendor-provided backup
preparation scripts.

---

## Step 6: Adding a Scale Hook

**Objective:** Define a Scale Hook that temporarily scales Kubernetes workloads as part of a
disaster recovery workflow. Scale Hooks are useful for quiescing an application before backup or
restoring the original replica count after recovery.

```yaml
spec:
  hooks:
    - name: scale-down
      type: scale
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: "busybox.*"
      timeout: 300
      onError: fail
      replicas: 0
```

**Explanation:**

- `type: scale` defines a Scale Hook.
- `selectResource` specifies the Kubernetes resource type to be scaled.
- `nameSelector` identifies the target resources. `busybox.*` matches `busybox-1` and `busybox-2`.
- `replicas` specifies the desired replica count for the selected resources.
- `timeout` specifies the maximum time (in seconds) to wait for the scaling to complete.
- `onError: fail` terminates the workflow if scaling fails or times out.

In this example, both Deployments are scaled to zero before the workflow proceeds. Scale Hooks can
also be used during recovery to restore workloads to their desired replica count after resources
have been recovered.

---

## Step 7: Adding a Job Hook

**Objective:** Define a Job Hook that executes one or more Kubernetes Jobs as part of a disaster
recovery workflow. Job Hooks are useful for performing prerequisite or post-recovery tasks that
require a dedicated execution environment.

```yaml
spec:
  hooks:
    - name: prepare-backup
      type: job
      namespace: recipe-demo
      timeout: 300
      onError: fail
      jobs:
        - name: backup-preparation
          forceCreate: true
          timeout: 300
          onError: fail

  jobs:
    - backup-preparation: |
        apiVersion: batch/v1
        kind: Job
        metadata:
          name: backup-preparation
          namespace: recipe-demo
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
                    - echo "Preparing application for backup"
```

**Explanation:**

- `type: job` defines a Job Hook.
- `jobs` specifies the list of Jobs to execute when the hook is invoked.
- `name` identifies the Job definition to execute (matched against `spec.jobs`).
- `forceCreate: true` ensures a new Job is created even if a Job with the same name already exists.
- `timeout` specifies the maximum time (in seconds) to wait for the Job to complete.
- `onError: fail` terminates the workflow if the Job fails or times out.

The `jobs` section at the spec level contains the Kubernetes Job manifests referenced by Job Hooks.
Each Job is defined once and can be referenced by one or more Job Hooks, promoting reuse across
capture and recover workflows.

Unlike an Exec Hook (which executes commands inside existing application containers), a Job Hook
runs in its own pod. This makes it suitable for custom scripts, exporting metadata, preparing
external systems, or other prerequisite operations that require a dedicated execution environment.

---

## Step 8: Complete Recipe Example

**Objective:** Combine all concepts from the previous steps into a complete Recipe that protects the
`busybox` application.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe
spec:
  appType: busybox

  groups:
    - name: busybox-resources
      backupRef: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      labelSelector:
        matchLabels:
          app: busybox

  hooks:
    - name: validate-deployments
      type: check
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: "busybox.*"
      timeout: 300
      onError: fail
      chks:
        - name: check-replicas
          condition: '{$.spec.replicas} == {$.status.readyReplicas}'

    - name: flush-database
      type: exec
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: "busybox.*"
      timeout: 300
      onError: fail
      ops:
        - name: flush
          container: busybox
          command: >
            ["/bin/sh", "-c", "sync"]

    - name: scale-down
      type: scale
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: "busybox.*"
      timeout: 300
      onError: fail
      replicas: 0

    - name: prepare-backup
      type: job
      namespace: recipe-demo
      timeout: 300
      onError: fail
      jobs:
        - name: backup-preparation
          forceCreate: true
          timeout: 300
          onError: fail

  jobs:
    - backup-preparation: |
        apiVersion: batch/v1
        kind: Job
        metadata:
          name: backup-preparation
          namespace: recipe-demo
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
                    - echo "Preparing application for backup"

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - hook: prepare-backup/backup-preparation
        - hook: flush-database/flush
        - hook: validate-deployments/check-replicas
        - group: busybox-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: busybox-resources
        - hook: validate-deployments/check-replicas
```

**Explanation:**

The `groups` section identifies the Kubernetes resources that make up the application, while the
`volumes` section identifies the PVCs to protect. The `hooks` section defines application-specific
operations, the `jobs` section holds the Kubernetes Job manifests referenced by Job Hooks, and the
`workflows` section orchestrates the execution order during capture and recovery.

**Backup workflow execution order:**

1. Execute the `prepare-backup` Job Hook — perform prerequisite tasks before the backup begins.
2. Execute the `flush-database` Exec Hook — prepare the application for a consistent backup.
3. Execute the `validate-deployments` Check Hook — verify that the selected Deployments are healthy.
4. Capture the application resources defined by the `busybox-resources` group.

**Restore workflow execution order:**

1. Restore the application resources defined by the `busybox-resources` group.
2. Execute the `validate-deployments` Check Hook — verify that restored Deployments have reached
   the expected healthy state.

The execution order of groups and hooks is fully controlled by the workflow sequence. By combining
resource groups, volume selection, hooks, and workflows, Recipes provide a flexible and
application-aware mechanism for protecting and recovering ACM discovered applications.

---

## Additional Example: Selective Resource Types

To protect only specific Kubernetes resource types (e.g., Deployments and ConfigMaps), use
`includedResourceTypes` in the group definition:

```yaml
spec:
  appType: my-app

  groups:
    - name: my-app-resources
      backupRef: my-app-resources
      type: resource
      includedNamespaces:
        - my-app-namespace
      labelSelector:
        matchLabels:
          app: my-app
      includedResourceTypes:
        - deployments
        - configmaps
        - services
        - secrets

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - group: my-app-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: my-app-resources
```

---

## Additional Example: Non-fatal Groups with `essential: false`

To allow the workflow to continue even if a particular group or hook fails, set `essential: false`.
This is useful for optional components that should not block the overall DR operation:

```yaml
spec:
  appType: my-app

  groups:
    - name: core-resources
      backupRef: core-resources
      type: resource
      includedNamespaces:
        - my-app-namespace
      labelSelector:
        matchLabels:
          tier: core

    - name: optional-resources
      backupRef: optional-resources
      type: resource
      essential: false
      includedNamespaces:
        - my-app-namespace
      labelSelector:
        matchLabels:
          tier: optional

  workflows:
    - name: backup
      failOn: essential-error
      sequence:
        - group: core-resources
        - group: optional-resources

    - name: restore
      failOn: essential-error
      sequence:
        - group: core-resources
        - group: optional-resources
```

---

## Additional Example: Quiesce and Unquiesce with inverseOp

To quiesce an application before backup and automatically unquiesce it afterward, use the
`inverseOp` field on an Exec Hook operation:

```yaml
spec:
  appType: my-db

  hooks:
    - name: db-quiesce
      type: exec
      namespace: my-db-namespace
      selectResource: deployment
      nameSelector: "my-db"
      timeout: 60
      onError: fail
      ops:
        - name: quiesce
          container: database
          command: '["/bin/sh", "-c", "mysql -e \"FLUSH TABLES WITH READ LOCK;\""]'
          inverseOp: unquiesce

        - name: unquiesce
          container: database
          command: '["/bin/sh", "-c", "mysql -e \"UNLOCK TABLES;\""]'

  groups:
    - name: db-resources
      backupRef: db-resources
      type: resource
      includedNamespaces:
        - my-db-namespace
      labelSelector:
        matchLabels:
          app: my-db

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - hook: db-quiesce/quiesce
        - group: db-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: db-resources
```

---

## Additional Example: Restoring Status with `restoreStatus`

To restore the status subresource of specific resource types during recovery:

```yaml
spec:
  appType: my-operator-app

  groups:
    - name: operator-resources
      backupRef: operator-resources
      type: resource
      includedNamespaces:
        - my-operator-namespace
      labelSelector:
        matchLabels:
          app: my-operator-app
      restoreStatus:
        includedResources:
          - mycustomresources

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - group: operator-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: operator-resources
```

---

## Real-World Example: PostgreSQL (Image-based Deployment)

**Workload description:**
A single PostgreSQL instance deployed as a Kubernetes `Deployment` (not an operator). The database
container runs as `postgres` user and the PVC holds all data. Before backup, a `CHECKPOINT` is
issued inside the primary pod to flush dirty pages to disk, ensuring a crash-consistent snapshot.
After restore, the recipe waits until the deployment reaches its desired replica count before
declaring success.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: postgres-image-based-backup-restore-recipe
  namespace: ramen-system
spec:
  appType: postgres
  groups:
    - name: postgres-volumes
      type: volume
    - name: postgres-resources
      type: resource
      excludedResourceTypes:
        - pods
        - replicasets
  hooks:
  - name: postgres-deployment-check
    type: check
    namespace: ${GROUP.postgres-resources.namespace}
    selectResource: deployment
    nameSelector: postgres-deployment
    timeout: 120
    onError: fail
    chks:
    - name: replicasReady
      timeout: 600
      onError: fail
      condition: "{$.spec.replicas} == {$.status.readyReplicas}"
  - name: postgres-instance-exec
    type: exec
    namespace: ${GROUP.postgres-resources.namespace}
    labelSelector: run=postgres
    singlePodOnly: true
    timeout: 120
    onError: fail
    ops:
    - name: checkpoint
      command: "psql -U postgres -c CHECKPOINT"
      container: postgres
      timeout: 60
  workflows:
  - name: backup
    sequence:
    - group: postgres-resources
    - hook: postgres-instance-exec/checkpoint
    - group: postgres-volumes
  - name: restore
    sequence:
    - group: postgres-volumes
    - group: postgres-resources
    - hook: postgres-deployment-check/replicasReady
```

**Sample prompt:**
> Generate an ODF DR Recipe for a PostgreSQL database deployed as a plain Kubernetes Deployment
> (not an operator) in the `postgres` namespace. Before backup, run `psql -U postgres -c CHECKPOINT`
> in the `postgres` container of the pod labeled `run=postgres` to flush dirty pages. After restore,
> wait for the `postgres-deployment` Deployment to reach its desired replica count before completing.
> Exclude ephemeral resource types (pods, replicasets) from the resource group.

**Validation:**
To validate the generated recipe, copy it into a file (e.g. `recipe.yaml`) and run:
```bash
ramenctl validate recipe recipe.yaml [--verbose]
```

---

## Real-World Example: MongoDB (Standalone — OT MongoDB Operator)

**Workload description:**
A MongoDB standalone instance managed by the OT MongoDB Operator in the `mongodb` namespace. The
operator manages a StatefulSet labeled `app=mongodb-standalone`. Before capturing volumes the
recipe locks the database with `db.fsyncLock()`, takes the snapshot, then calls `db.fsyncUnlock()`.
After restore the recipe waits for both the operator Deployment and the standalone StatefulSet to
reach their desired replica counts.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: mongodb-community-standalone-backup-restore-recipe
  namespace: ramen-system
spec:
  appType: mongodb
  groups:
    - name: mongodb-volumes
      type: volume
      includedNamespaces:
        - mongodb
    - name: mongodb-resources
      type: resource
      includedNamespaces:
        - mongodb
      includeClusterResources: true
      excludedResourceTypes:
        - events
        - mongodbs.opstreelabs.in
    - name: mongodb-instances
      type: resource
      includedNamespaces:
        - mongodb
      includedResourceTypes:
        - mongodbs.opstreelabs.in
  hooks:
  - name: mongodb-operator-check
    type: check
    namespace: mongodb
    selectResource: deployment
    labelSelector: control-plane=mongodb-operator
    timeout: 120
    onError: fail
    chks:
    - name: replicasReady
      timeout: 180
      onError: fail
      condition: "{$.spec.replicas} == {$.status.readyReplicas}"
  - name: mongodb-standalone-check
    type: check
    namespace: mongodb
    selectResource: statefulset
    labelSelector: app=mongodb-standalone,mongodb_setup=standalone,role=standalone
    timeout: 120
    onError: fail
    chks:
    - name: replicasReady
      timeout: 180
      onError: fail
      condition: "{$.spec.replicas} == {$.status.readyReplicas}"
  - name: mongodb-standalone-exec
    labelSelector: app=mongodb-standalone
    timeout: 300
    namespace: mongodb
    onError: fail
    ops:
      - command: >
          ["/bin/bash", "-c", "mongosh -u `printenv MONGO_ROOT_USERNAME` -p `printenv MONGO_ROOT_PASSWORD` --eval \"db.fsyncLock()\""]
        container: mongo
        timeout: 300
        name: fsyncLock
        onError: fail
      - command: >
          ["/bin/bash", "-c", "mongosh -u `printenv MONGO_ROOT_USERNAME` -p `printenv MONGO_ROOT_PASSWORD` --eval \"db.fsyncUnlock()\""]
        container: mongo
        timeout: 300
        name: fsyncUnlock
        onError: fail
    selectResource: pod
    type: exec
  workflows:
  - name: backup
    sequence:
    - group: mongodb-resources
    - group: mongodb-instances
    - hook: mongodb-standalone-exec/fsyncLock
    - group: mongodb-volumes
    - hook: mongodb-standalone-exec/fsyncUnlock
  - name: restore
    sequence:
    - group: mongodb-volumes
    - group: mongodb-resources
    - hook: mongodb-operator-check/replicasReady
    - group: mongodb-instances
    - hook: mongodb-standalone-check/replicasReady
```

**Sample prompt:**
> Generate an ODF DR Recipe for a MongoDB standalone instance managed by the OT MongoDB Operator
> in the `mongodb` namespace. Before capturing volumes, lock the database with `db.fsyncLock()`
> on pods labeled `app=mongodb-standalone`, capture volumes, then immediately call
> `db.fsyncUnlock()`. Credentials are read from the `MONGO_ROOT_USERNAME` and
> `MONGO_ROOT_PASSWORD` environment variables. After restore, confirm both the operator Deployment
> and the standalone StatefulSet reach their desired replica counts. Capture the MongoDB CR
> (`mongodbs.opstreelabs.in`) as a separate group from general resources.

**Validation:**
To validate the generated recipe, copy it into a file (e.g. `recipe.yaml`) and run:
```bash
ramenctl validate recipe recipe.yaml [--verbose]
```

---

## Real-World Example: MinIO (Scale-to-zero Quiesce)

**Workload description:**
A MinIO object-storage instance deployed as a Kubernetes `Deployment`. The recipe uses a Scale Hook
to scale the MinIO Deployment to zero replicas before capturing the PVC snapshot (ensuring no
in-flight writes), then scales it back up after the volume capture completes. After restore, a
Check Hook waits for the `minio` Deployment to reach its desired replica count before declaring
success.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: minio-scale-backup-restore-recipe
  namespace: ramen-system
spec:
  appType: minio
  groups:
    - name: minio-volumes
      type: volume
    - name: minio-resources
      type: resource
  hooks:
  - name: minio-deployment-check
    type: check
    namespace: ${GROUP.minio-resources.namespace}
    selectResource: deployment
    nameSelector: minio
    timeout: 120
    onError: fail
    chks:
    - name: replicasReady
      timeout: 600
      onError: fail
      condition: "{$.spec.replicas} == {$.status.readyReplicas}"
  - name: minio-scale-deployment
    type: scale
    namespace: ${GROUP.minio-resources.namespace}
    selectResource: deployment
    nameSelector: minio
  workflows:
  - name: backup
    sequence:
    - group: minio-resources
    - hook: minio-scale-deployment/down
    - hook: minio-scale-deployment/sync
    - group: minio-volumes
    - hook: minio-scale-deployment/up
    - hook: minio-scale-deployment/sync
  - name: restore
    sequence:
    - group: minio-volumes
    - group: minio-resources
    - hook: minio-deployment-check/replicasReady
```

**Sample prompt:**
> Generate an ODF DR Recipe for a MinIO object-storage instance deployed as a Kubernetes Deployment.
> Before capturing volumes, scale the `minio` Deployment to zero (`scale/down` + `scale/sync`)
> to quiesce all writes. After capturing volumes, scale back up (`scale/up` + `scale/sync`).
> After restore, wait for the `minio` Deployment to reach its desired replica count. Use the
> namespace dynamically from the resource group (`${GROUP.minio-resources.namespace}`).

**Validation:**
To validate the generated recipe, copy it into a file (e.g. `recipe.yaml`) and run:
```bash
ramenctl validate recipe recipe.yaml [--verbose]
```

---

## Guidance: Always Verify Resource Labels Before Writing a Recipe

When writing a Recipe for an application running in a namespace, the `labelSelector` inside each
group determines which resources are captured and restored. If the `labelSelector` is omitted or
too broad, the Recipe may capture unrelated resources that share the same namespace. If it is too
narrow or references labels that do not exist on the actual resources, the group captures nothing
and the DR operation silently protects an empty set.

**Before writing any Recipe, inspect the labels on the application's resources:**

```bash
# List all resources in the namespace with their labels
kubectl get all -n <namespace> --show-labels

# Check labels on PersistentVolumeClaims specifically
kubectl get pvc -n <namespace> --show-labels

# Check labels on a specific resource type
kubectl get deployments -n <namespace> --show-labels
kubectl get statefulsets -n <namespace> --show-labels
```

Choose a label (or label combination) that is present on **all** resources belonging to the
application and absent from resources that should not be protected. Common patterns are:

| Label pattern | Example |
|---|---|
| `app: <name>` | `app: postgres` |
| `app.kubernetes.io/name: <name>` | `app.kubernetes.io/name: mysql` |
| `app.kubernetes.io/part-of: <name>` | `app.kubernetes.io/part-of: my-app` |
| Custom label | `workload: redis-cache` |

**Example — incorrect (no labelSelector, captures the entire namespace):**

```yaml
groups:
  - name: app-resources
    type: resource
    includedNamespaces:
      - my-app-namespace
```

**Example — correct (labelSelector scoped to the application):**

```yaml
groups:
  - name: app-resources
    type: resource
    includedNamespaces:
      - my-app-namespace
    labelSelector:
      matchLabels:
        app: my-app
```

The same rule applies to `spec.volumes`. The `labelSelector` inside a volume group selects
PersistentVolumeClaims. If the PVCs do not carry the expected label, they will not be included in
the snapshot:

```bash
# Verify the labels on each PVC before using them in spec.volumes
kubectl get pvc -n <namespace> --show-labels
```

```yaml
volumes:
  - name: app-volumes
    labelSelector:
      matchLabels:
        app: my-app   # must match a label that actually exists on the PVC objects
```

If the PVCs carry a different label from the rest of the application resources (which is common
with operator-managed databases), use a separate `labelSelector` that matches the PVC label rather
than copying the one used for the resource group.

---
