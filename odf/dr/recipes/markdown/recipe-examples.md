# ODF Disaster Recovery — Recipe Examples

## Purpose

This document provides canonical examples of ODF Disaster Recovery `Recipe` Custom Resources.

The examples progress from a minimal Recipe to application-specific patterns.

All examples follow these rules:

* `groups` contains Resource Groups with `type: resource`.
* `volumes` contains Volume Groups with `type: volume`.
* Volume Groups are used to select PersistentVolumeClaims for protection.
* Volume Groups are **not referenced in workflow sequences**.
* Workflow sequences reference Resource Groups and Hooks.
* `backup` and `restore` workflows are defined for complete Recipes.
* Hooks are included only when application-specific operations or validation are required.
* Actual application labels should be used when they are known.
* PVC labels must be verified independently rather than assumed to match workload labels.

---

# Example 1 — Minimal Recipe

This example demonstrates the smallest complete Recipe.

It contains:

* one Resource Group
* one backup workflow
* one restore workflow

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource

  workflows:
    - name: backup
      sequence:
        - group: busybox-resources

    - name: restore
      sequence:
        - group: busybox-resources
```

### Explanation

The Resource Group identifies the application resources that participate in disaster recovery.

The backup and restore workflows reference the Resource Group.

No Volume Group is defined because this minimal example does not specify persistent storage.

---

# Example 2 — Selecting Application Resources

This example adds namespace and label-based resource selection.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  workflows:
    - name: backup
      sequence:
        - group: busybox-resources

    - name: restore
      sequence:
        - group: busybox-resources
```

### Explanation

`includedNamespaces` limits resource selection to the application namespace.

`labelSelector` selects resources belonging to the application.

When actual application labels are known, those labels should be used rather than invented labels.

---

# Example 3 — Protecting PersistentVolumeClaims

This example adds a Volume Group for PersistentVolumeClaims.

The application resources and PVCs are selected independently.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      type: volume
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  workflows:
    - name: backup
      sequence:
        - group: busybox-resources

    - name: restore
      sequence:
        - group: busybox-resources
```

### Explanation

The `spec.volumes` section identifies the PersistentVolumeClaims that require protection.

The Volume Group uses the same `Group` schema and therefore specifies:

```yaml
type: volume
```

The PVCs are selected using their own labels.

The workload labels and PVC labels may be different. If PVC labels are known, use those actual labels.

### Important

The Volume Group is **not referenced by either workflow**.

The following is not valid:

```yaml
volumes:
  - name: busybox-pvc
    selector:
      matchLabels:
        app: busybox
```

or:

```yaml
workflows:
  - name: backup
    sequence:
      - group: busybox-resources
      - group: busybox-volumes
```

PVC protection is determined from the Volume Group defined under `spec.volumes`, and label-based volume selection uses `labelSelector`, not `selector`.

---

# Example 3A — Deployment and PVC with the Same Application Label

This example matches a common generation request: one Deployment, one PVC, one namespace, and the same application label on both resources.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: my-app-recipe

spec:
  appType: my-app

  groups:
    - name: app-resources
      type: resource
      includedNamespaces:
        - my-app
      labelSelector:
        matchLabels:
          app: my-app

  volumes:
    - name: app-volumes
      type: volume
      includedNamespaces:
        - my-app
      labelSelector:
        matchLabels:
          app: my-app

  workflows:
    - name: backup
      sequence:
        - group: app-resources

    - name: restore
      sequence:
        - group: app-resources
```

### Important

When the PVC label is known, select the PVC using `labelSelector` under `spec.volumes`.

Do not generate:

```yaml
volumes:
  - name: app-pvc
    persistentVolumeClaimNames:
      - <pvc-name>
```

---

# Example 3B — Name-Based PVC Selection

This example demonstrates name-based PVC selection when the PVC labels are not known.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: payroll-app-recipe

spec:
  appType: payroll-app

  groups:
    - name: payroll-app-resources
      type: resource
      includedNamespaces:
        - payroll
      labelSelector:
        matchLabels:
          app: payroll-app

  volumes:
    - name: payroll-app-volumes
      type: volume
      includedNamespaces:
        - payroll
      nameSelector: payroll-data

  workflows:
    - name: backup
      sequence:
        - group: payroll-app-resources

    - name: restore
      sequence:
        - group: payroll-app-resources
```

### Important

When PVC labels are not known and the PVC name is known, use `nameSelector`.

Do not generate:

```yaml
volumes:
  - name: payroll-app-volumes
    persistentVolumeClaimNames:
      - payroll-data
```

---

# Example 4 — Backup and Restore Workflows

This example demonstrates workflow sequencing and failure handling.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      type: volume
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

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

### Explanation

Workflow steps execute in sequence.

A workflow may reference:

* Resource Groups
* Hooks

A Volume Group is not a workflow step.

The backup and restore workflows above contain only the Resource Group because no application-specific Hook is required.

---

# Example 5 — Adding an Exec Hook

An Exec Hook can perform an application-specific operation before backup.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      type: volume
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  hooks:
    - name: prepare-backup
      type: exec
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: busybox
      timeout: 120
      onError: fail
      ops:
        - name: flush
          container: busybox
          command: '["/bin/sh", "-c", "sync"]'

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - hook: prepare-backup/flush
        - group: busybox-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: busybox-resources
```

### Explanation

The Exec Hook executes the `sync` command before the Resource Group is processed.

The workflow references the Hook using:

```text
hook: prepare-backup/flush
```

The Volume Group remains outside the workflow.

When an Exec Hook defines named operations under `ops`, reference the specific operation name in the workflow step.

Exec Hooks can be used for operations such as:

* flushing pending writes
* creating application checkpoints
* pausing application writes
* running application-specific preparation commands

---

# Example 6 — Adding a Check Hook

A Check Hook can validate application state after restore.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      type: volume
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  hooks:
    - name: prepare-backup
      type: exec
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: busybox
      timeout: 120
      onError: fail
      ops:
        - name: flush
          container: busybox
          command: '["/bin/sh", "-c", "sync"]'

    - name: validate-application
      type: check
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: busybox
      timeout: 120
      onError: fail
      chks:
        - name: replicasReady
          timeout: 300
          onError: fail
          condition: "{$.spec.replicas} == {$.status.readyReplicas}"

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - hook: prepare-backup/flush
        - group: busybox-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: busybox-resources
        - hook: validate-application/replicasReady
```

### Explanation

The restore workflow restores the application resources and then validates that the Deployment has the expected number of ready replicas.

The Check Hook is referenced using:

```text
hook: validate-application/replicasReady
```

When a Check Hook defines named checks under `chks`, reference the specific check name in the workflow step.

Do not reference only the Hook name when the workflow step must refer to a specific check.

Check Hooks can validate:

* Deployment readiness
* StatefulSet readiness
* operator availability
* application-specific conditions

---

# Example 7 — Quiescing an Application with a Scale Hook

A Scale Hook can temporarily scale a workload down before backup and restore it after recovery.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      type: volume
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  hooks:
    - name: scale-application
      type: scale
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: busybox
      replicas: 0
      timeout: 300
      onError: fail

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - hook: scale-application/down
        - hook: scale-application/sync
        - group: busybox-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: busybox-resources
        - hook: scale-application/up
        - hook: scale-application/sync
```

### Explanation

The backup workflow:

1. Scales the application down.
2. Waits for the scaling operation to complete.
3. Processes the application Resource Group.

The restore workflow:

1. Restores the application Resource Group.
2. Restores the workload to its original replica count.
3. Waits for the scaling operation to complete.

Use the scale operation names shown in the workflow sequence (`down`, `up`, and `sync`).

The Volume Group is not included in either workflow.

---

# Example 8 — Adding a Job Hook

A Job Hook can execute application-specific logic using a Kubernetes Job.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      type: volume
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

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

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - hook: prepare-backup/backup-preparation
        - group: busybox-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: busybox-resources
```

### Explanation

The Job Hook references the Job Definition from `spec.jobs`.

The workflow references the Job Hook using:

```text
hook: prepare-backup/backup-preparation
```

Do not reference only the Hook name when the workflow step must refer to a specific Job.

Job Hooks are useful when application-specific operations require a dedicated execution environment.

---

# Example 8A — Job Hook for Backup Preparation

This example matches a common generation request: run a Kubernetes Job before backup.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: reporting-app-recipe

spec:
  appType: reporting-app

  groups:
    - name: reporting-app-resources
      type: resource
      includedNamespaces:
        - reporting
      labelSelector:
        matchLabels:
          app: reporting-app

  volumes:
    - name: reporting-app-volumes
      type: volume
      includedNamespaces:
        - reporting
      labelSelector:
        matchLabels:
          app: reporting-app

  hooks:
    - name: prepare-backup
      type: job
      namespace: reporting
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
          namespace: reporting
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
      sequence:
        - hook: prepare-backup/backup-preparation
        - group: reporting-app-resources

    - name: restore
      sequence:
        - group: reporting-app-resources
```

### Important

A Job Hook uses `type: job` and references Job Definitions from `spec.jobs`.

Do not generate:

```yaml
hooks:
  - name: backup-preparation-hook
    type: exec
    exec:
      jobName: backup-preparation
```

---

# Example 9 — Complete Recipe

This example combines Resource Groups, Volume Selection, Hooks, and Workflows.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: busybox-recipe
  namespace: ramen-system

spec:
  appType: busybox

  groups:
    - name: busybox-resources
      type: resource
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  volumes:
    - name: busybox-volumes
      type: volume
      includedNamespaces:
        - recipe-demo
      labelSelector:
        matchLabels:
          app: busybox

  hooks:
    - name: prepare-backup
      type: exec
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: busybox
      timeout: 120
      onError: fail
      ops:
        - name: flush
          container: busybox
          command: '["/bin/sh", "-c", "sync"]'

    - name: validate-application
      type: check
      namespace: recipe-demo
      selectResource: deployment
      nameSelector: busybox
      timeout: 120
      onError: fail
      chks:
        - name: replicasReady
          timeout: 300
          onError: fail
          condition: "{$.spec.replicas} == {$.status.readyReplicas}"

  workflows:
    - name: backup
      failOn: any-error
      sequence:
        - hook: prepare-backup/flush
        - group: busybox-resources

    - name: restore
      failOn: any-error
      sequence:
        - group: busybox-resources
        - hook: validate-application/replicasReady
```

### Explanation

During backup:

1. The application is prepared using the Exec Hook.
2. The Resource Group is processed.
3. PVCs selected by `spec.volumes` are protected independently.

During restore:

1. The Resource Group is restored.
2. The Check Hook validates application readiness.
3. PVCs selected by `spec.volumes` are restored independently.

The Volume Group is intentionally absent from both workflow sequences.

---

# Recipe Pattern — PostgreSQL Deployment

## Application

This pattern demonstrates a PostgreSQL database deployed as a Kubernetes Deployment with PersistentVolumeClaims.

## Protection requirements

The database:

* stores data on PersistentVolumeClaims
* should execute `CHECKPOINT` before backup
* should be validated after restore

## Recipe

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: postgres-dr-recipe
  namespace: ramen-system

spec:
  appType: postgres

  groups:
    - name: postgres-resources
      type: resource
      includedNamespaces:
        - postgres
      excludedResourceTypes:
        - pods
        - replicasets

  volumes:
    - name: postgres-volumes
      type: volume
      includedNamespaces:
        - postgres
      labelSelector:
        matchLabels:
          app: postgres

  hooks:
    - name: postgres-checkpoint
      type: exec
      namespace: ${GROUP.postgres-resources.namespace}
      selectResource: deployment
      nameSelector: postgres-deployment
      timeout: 120
      onError: fail
      ops:
        - name: checkpoint
          container: postgres
          command: '["psql","-U","postgres","-c","CHECKPOINT"]'
          timeout: 60

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

  workflows:
    - name: backup
      sequence:
        - hook: postgres-checkpoint/checkpoint
        - group: postgres-resources

    - name: restore
      sequence:
        - group: postgres-resources
        - hook: postgres-deployment-check/replicasReady
```

### Key points

* PostgreSQL resources are selected through a Resource Group.
* PostgreSQL PVCs are selected independently through `spec.volumes`.
* The `CHECKPOINT` operation runs before backup.
* The restore workflow validates Deployment readiness.
* The Volume Group is not referenced by either workflow.

---

# Recipe Pattern — MongoDB Operator

## Application

This pattern demonstrates an operator-managed MongoDB application.

Operator-managed applications may require separate Resource Groups so that the operator and the Custom Resources it manages can be restored in the required order.

## Protection requirements

The application:

* has operator resources
* has MongoDB Custom Resources
* has PersistentVolumeClaims
* requires an application-specific database operation before volume protection
* requires validation after restore

## Recipe structure

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: mongodb-dr-recipe
  namespace: ramen-system

spec:
  appType: mongodb

  groups:
    - name: mongodb-operator-resources
      type: resource
      includedNamespaces:
        - mongodb
      labelSelector:
        matchLabels:
          control-plane: mongodb-operator

    - name: mongodb-resources
      type: resource
      includedNamespaces:
        - mongodb
      labelSelector:
        matchLabels:
          app: mongodb

  volumes:
    - name: mongodb-volumes
      type: volume
      includedNamespaces:
        - mongodb
      labelSelector:
        matchLabels:
          app: mongodb

  hooks:
    - name: mongodb-lock
      type: exec
      namespace: mongodb
      selectResource: pod
      labelSelector:
        app: mongodb
      timeout: 300
      onError: fail
      ops:
        - name: fsyncLock
          container: mongo
          command: '["/bin/bash","-c","mongosh --eval \"db.fsyncLock()\""]'
          timeout: 300

        - name: fsyncUnlock
          container: mongo
          command: '["/bin/bash","-c","mongosh --eval \"db.fsyncUnlock()\""]'
          timeout: 300

    - name: mongodb-ready
      type: check
      namespace: mongodb
      selectResource: statefulset
      labelSelector:
        app: mongodb
      timeout: 120
      onError: fail
      chks:
        - name: replicasReady
          timeout: 300
          onError: fail
          condition: "{$.spec.replicas} == {$.status.readyReplicas}"

  workflows:
    - name: backup
      sequence:
        - group: mongodb-operator-resources
        - group: mongodb-resources
        - hook: mongodb-lock/fsyncLock
        - hook: mongodb-lock/fsyncUnlock

    - name: restore
      sequence:
        - group: mongodb-operator-resources
        - group: mongodb-resources
        - hook: mongodb-ready/replicasReady
```

### Key points

* Operator resources and application resources can be separated into different Resource Groups.
* PVC protection is defined independently through `spec.volumes`.
* Application-specific operations are represented by Hooks.
* The workflow references only Resource Groups and Hooks.
* The Volume Group is never added to a workflow sequence.

The exact MongoDB command, selectors, and resource ordering must be adapted to the actual MongoDB deployment.

---

# Recipe Pattern — MinIO with Scale Hook

## Application

This pattern demonstrates an application that can be quiesced by scaling its Deployment down before backup.

## Protection requirements

The application:

* runs as a Deployment
* uses PersistentVolumeClaims
* should be scaled down before backup
* should be restored to its original replica count after recovery
* should be validated after restore

## Recipe

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: minio-dr-recipe
  namespace: ramen-system

spec:
  appType: minio

  groups:
    - name: minio-resources
      type: resource
      includedNamespaces:
        - minio
      labelSelector:
        matchLabels:
          app: minio

  volumes:
    - name: minio-volumes
      type: volume
      includedNamespaces:
        - minio
      labelSelector:
        matchLabels:
          app: minio

  hooks:
    - name: minio-scale
      type: scale
      namespace: ${GROUP.minio-resources.namespace}
      selectResource: deployment
      nameSelector: minio
      replicas: 0
      timeout: 300
      onError: fail

    - name: minio-ready
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

  workflows:
    - name: backup
      sequence:
        - hook: minio-scale/down
        - hook: minio-scale/sync
        - group: minio-resources

    - name: restore
      sequence:
        - group: minio-resources
        - hook: minio-scale/up
        - hook: minio-scale/sync
        - hook: minio-ready/replicasReady
```

### Key points

During backup:

1. Scale the application down.
2. Wait for scaling to complete.
3. Process the application resources.

During restore:

1. Restore the application resources.
2. Restore the workload to its original replica count.
3. Wait for scaling to complete.
4. Validate application readiness.

The Volume Group remains under `spec.volumes` and is not included in either workflow.

---

# Validation

After creating or modifying a Recipe, validate it before using it for disaster recovery.

```bash
ramenctl validate recipe recipe.yaml
```

For detailed validation output:

```bash
ramenctl validate recipe recipe.yaml --verbose
```

Validation should identify schema errors, invalid references, and other Recipe configuration problems before the Recipe is used.

---

# Summary

These examples demonstrate the main Recipe building blocks:

* Resource Groups
* Volume Groups
* Workflows
* Exec Hooks
* Check Hooks
* Scale Hooks
* Job Hooks
* Name-based PVC selection with `nameSelector`

The key structural pattern is:

```yaml
spec:
  groups:
    - name: <resource-group>
      type: resource

  volumes:
    - name: <volume-group>
      type: volume

  workflows:
    - name: backup
      sequence:
        - group: <resource-group>

    - name: restore
      sequence:
        - group: <resource-group>
```

**Resource Groups and Hooks are workflow components. Volume Groups are not.**

Volume Selection under `spec.volumes` independently identifies the PersistentVolumeClaims that ODF Disaster Recovery protects and restores.

Application-specific Hooks should be added only when the application requires preparation, quiescing, validation, or other recovery operations.
