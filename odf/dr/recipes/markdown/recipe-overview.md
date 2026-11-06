# ODF Disaster Recovery — Recipes

## Preface

OpenShift Data Foundation (ODF) Disaster Recovery supports protection of two broad categories of
applications:

- **ACM Managed Applications** – Applications that are managed through Advanced Cluster Management
  (ACM). The lifecycle of these applications, including deployment, placement, and recovery, is
  coordinated by ACM.
- **ACM Discovered Applications** – Applications that are deployed independently of ACM and later
  discovered for disaster recovery. Since ACM does not manage their lifecycle, users are responsible
  for ensuring that all application resources are captured and restored correctly during DR
  operations.

Protecting ACM discovered applications can be challenging because many applications require more
than simply backing up Kubernetes resources. Resources often need to be captured in a specific
order, application-specific actions may be required before backup or after restore, and certain
resources may need to be excluded or restored only after dependent components become available.

To address these challenges, ODF Disaster Recovery provides **Recipes**. A Recipe enables users to
define application-specific workflows that describe how an application should be protected and
recovered. Using a Recipe, users can group related resources, execute application-specific hooks,
and control the sequence in which backup and restore operations are performed, making disaster
recovery of ACM discovered applications predictable and reliable.

---

## Introduction

A Recipe is a Kubernetes custom resource that defines the workflow for protecting and recovering an
ACM discovered application. Instead of relying on a generic backup and restore process, a Recipe
allows users to describe the application's specific protection and recovery requirements.

A Recipe provides a structured way to:

- Group related Kubernetes resources that belong to the application.
- Define the order in which resources should be backed up and restored.
- Execute application-specific operations before, during, or after backup and restore.
- Capture application state that cannot be protected using Kubernetes resources alone.
- Ensure dependencies between application components are respected during disaster recovery.

Many enterprise applications require additional preparation before a consistent backup can be taken.
Simply capturing Kubernetes resources may not be sufficient if the application is actively
processing requests, writing data, or waiting for dependent services to become available. Recipes
address these scenarios by supporting **hooks**, which allow custom actions or validations to be
performed as part of the protection and recovery workflow.

Recipes support the following hook types:

- **Exec Hook** – Executes one or more commands inside a container running in a pod. This is useful
  for performing application-specific operations such as flushing in-memory data to disk,
  checkpointing transactions, pausing application writes, or preparing the application for a
  consistent backup. For example, an Exec hook can execute a database flush command immediately
  before the backup is initiated to ensure all pending writes are persisted.
- **Check Hook** – Verifies that a resource or application is in the expected state before the
  workflow proceeds. For example, a Check hook can validate that a database pod is in a healthy
  state, all replicas are ready, or an application has completed initialization before allowing the
  backup to continue.
- **Scale Hook** – Temporarily scales Kubernetes workloads up or down as part of the workflow. This
  can be used to quiesce an application by scaling a deployment to zero replicas before backup and
  restoring it to its original replica count afterward, or to ensure the required number of replicas
  are running during recovery.
- **Job Hook** – Creates and executes a Kubernetes Job to perform tasks that are better suited to an
  isolated execution environment. This is useful for running custom scripts, exporting application
  metadata, preparing external systems, migrating data, or performing other prerequisite operations
  before backup or after recovery.

By combining resource grouping, ordered workflows, and hooks, Recipes enable application-aware
disaster recovery. This ensures that backups are taken only after prerequisite conditions have been
met and that applications are restored in a predictable and consistent manner, reducing the need for
manual intervention during disaster recovery operations.

---

## Recipe Structure

The following diagram illustrates the high-level organization of a Recipe:

```
Recipe
├── Metadata
│
└── Spec
    ├── Groups
    │     └── Defines application resources
    │
    ├── Hooks
    │     ├── Check
    │     ├── Exec
    │     ├── Scale
    │     └── Job
    │
    ├── Volumes
    │     └── Selects PVCs for protection
    │
    └── Workflows
          ├── Backup
          │     ├── References Groups
          │     └── References Hooks
          │
          └── Restore
                ├── References Groups
                └── References Hooks
```

The `groups`, `hooks`, and `volumes` sections define the resources and actions available to the
Recipe. The `workflows` section references these definitions to determine the order in which
resources are processed and hooks are executed during capture and recovery operations.

---

## Recipe Parameters

This section describes the configurable fields of a Recipe and how each contributes to the
protection and recovery workflow.

### Metadata

The `metadata` section contains the standard Kubernetes object metadata used to identify the
Recipe.

| Parameter | Description |
|---|---|
| `metadata.name` | Specifies the name of the Recipe. This name is used to reference the Recipe during disaster recovery operations. |
| `metadata.namespace` | Specifies the namespace in which the Recipe is created. |

---

### `spec.groups`

Groups define logical collections of Kubernetes resources that participate in the disaster recovery
workflow. Rather than processing every resource individually, related resources can be grouped
together and referenced by the workflow.

Each group can identify resources using criteria such as:

- API group and resource kind
- Label selectors
- Name selectors (where applicable)

Organizing resources into groups makes the Recipe easier to understand, reuse, and maintain while
allowing workflows to process related resources together.

---

### `spec.workflows`

Workflows define the sequence in which resource groups and hooks are executed during protection and
recovery operations.

A Recipe typically contains workflows for capture (backup) and recovery (restore). Each workflow
consists of one or more ordered steps that reference resource groups and, when required, associated
hooks.

Defining workflows allows applications with dependencies between components to be protected and
restored in a controlled and predictable manner.

---

### `spec.hooks`

Hooks enable application-specific operations to be performed as part of a workflow. They allow a
Recipe to prepare an application for backup, validate application state, or perform recovery tasks
after resources have been restored.

The following hook types are supported:

- **Check Hook** – Verifies that a resource or application is in the expected state before the
  workflow proceeds.
- **Exec Hook** – Executes commands inside a container running in a pod to perform
  application-specific operations.
- **Scale Hook** – Scales Kubernetes workloads up or down as part of the workflow.
- **Job Hook** – Creates and executes a Kubernetes Job to perform custom prerequisite or
  post-recovery tasks.

Hooks are associated with workflow steps and are executed at the appropriate stage of the capture or
recovery process.

---

### `spec.volumes`

The `spec.volumes` section identifies the PersistentVolumeClaims (PVCs) that are associated with
the application and should be considered during protection.

PVCs are selected using the label selectors defined in this section. Any PVC whose labels match the
specified selectors is included as part of the application's protection workflow. This enables
Recipes to dynamically identify application volumes without requiring individual PVC names to be
specified.

If one or more `labelSelectors` are specified under `spec.volumes`, they take precedence over the
PVC `labelSelector` configured in the DRPlacementControl (DRPC) during protection. In this case,
ODF Disaster Recovery uses the selectors defined in the Recipe to determine which PVCs are
protected, regardless of the selector configured in the DRPC.

If `spec.volumes.labelSelectors` is not specified, PVC selection falls back to the `labelSelector`
configured in the DRPlacementControl.

Using label selectors makes Recipes reusable across multiple deployments and environments, provided
the PVCs follow a consistent labeling strategy.

---

## Configuring a Recipe

A Recipe is created as a Kubernetes custom resource that defines how an ACM discovered application
should be protected and recovered.

The typical configuration process consists of the following steps:

1. Identify the Kubernetes resources that make up the application.
2. Organize related resources into logical groups.
3. Define the capture workflow.
4. Define the recovery workflow.
5. Configure any required hooks.
6. Identify the application PVCs using `spec.volumes`.
7. Apply the Recipe to the cluster.

Once the Recipe has been created and applied, it can be used to protect ACM discovered applications.
When configuring disaster recovery protection for a discovered application, ensure that the
appropriate Recipe is available in the cluster and selected as part of the protection configuration.
ODF Disaster Recovery uses the selected Recipe to determine the application resources to protect,
execute any configured hooks, identify the PersistentVolumeClaims (PVCs) to be protected, and
orchestrate the backup and restore workflows during disaster recovery operations.
