# ODF Disaster Recovery — Recipe Authoring Guide

## Overview

A Recipe defines how an ACM Discovered Application should be protected and recovered during disaster recovery.

Unlike a generic backup configuration, a Recipe captures the application's recovery requirements by identifying the resources that belong to the application, selecting the PersistentVolumeClaims (PVCs) that require protection, defining any application-specific operations that must be performed before or after backup and restore, and specifying the order in which these operations are executed.

A well-designed Recipe reflects the application's architecture rather than simply listing Kubernetes resources. Applications differ in how they use persistent storage, interact with external systems, initialize components, and recover from failures. These characteristics determine how the Recipe should be designed.

This guide describes a structured approach for authoring Recipes. Instead of beginning with the Recipe schema, it begins by analyzing the application that needs to be protected. The application analysis is then used to determine the Recipe structure, workflows, and supporting components.

The API reference describes the individual fields available in the Recipe specification. This guide focuses on the design decisions involved in creating a complete and maintainable Recipe.

## Recipe Authoring Workflow

Authoring a Recipe is an incremental process. Rather than starting with the Recipe schema, begin by understanding the application that will be protected.

The recommended design process consists of the following stages:

1. Analyze the application architecture.
2. Identify the Kubernetes resources that participate in disaster recovery.
3. Determine whether persistent storage requires protection.
4. Identify application-specific operations that must occur before or after backup or restore.
5. Organize resources into logical groups.
6. Design the backup and restore workflows.
7. Validate the completed Recipe.

The following diagram summarizes the recommended workflow.

```text
Application
    │
    ▼
Analyze Application
    │
    ├── Workloads
    ├── Persistent Storage
    ├── Dependencies
    ├── Consistency Requirements
    └── Recovery Requirements
            │
            ▼
Determine Recipe Components
            │
            ├── Resource Groups
            ├── Volume Selection
            ├── Hooks
            ├── Job Definitions
            └── Workflows
                    │
                    ▼
Validate Recipe
                    │
                    ├── Resource Selection
                    ├── Workflow References
                    ├── Execution Order
                    └── Failure Policies
```

Each stage builds upon the previous one. Decisions made during application analysis determine the structure of the final Recipe.

## Step 1 — Analyze the Application

Before creating a Recipe, understand how the application is deployed and what is required to recover it successfully.

The following questions should be answered before designing the Recipe.

| Question | Design Impact |
|----------|---------------|
| Which Kubernetes workloads make up the application? | Determines the resources that should be protected. |
| Does the application use PersistentVolumeClaims? | Determines whether volume protection is required. |
| Are all resources restored together, or do some components depend on others? | Determines how resource groups and workflows should be organized. |
| Does the application require a consistent state before backup? | Determines whether application-specific hooks are required. |
| Does the application require validation after restore? | Determines whether recovery validation should be performed. |
| Does the application execute custom scripts during backup or restore? | Determines whether Job Hooks are required. |

The answers to these questions determine the overall structure of the Recipe.

## Step 2 — Determine the Recipe Components

After analyzing the application, determine which Recipe components are required to protect and recover it.

Not every application requires every Recipe component. The application architecture and recovery requirements determine which sections should be included in the Recipe.

The following table summarizes the purpose of each Recipe component and when it is typically used.

| Recipe Component | Purpose | Typical Usage |
|------------------|---------|---------------|
| Resource Groups | Identify the Kubernetes resources that belong to the application. | Every Recipe should define one or more resource groups. |
| Volume Selection | Identify the PersistentVolumeClaims (PVCs) that require protection. | Applications that use persistent storage. |
| Hooks | Perform application-specific operations before or after backup and restore. | Databases, distributed applications, applications requiring quiescing or validation. |
| Job Definitions | Define Kubernetes Jobs that are executed by Job Hooks. | Applications requiring standalone scripts or custom recovery tasks. |
| Workflows | Define the order in which resource groups and hooks are executed. | Every Recipe should define backup and restore workflows. |

Rather than beginning with the Recipe schema, first determine which of these components are required by the application. The Recipe can then be constructed by combining only the necessary components.

## Selecting Recipe Components

The following guidelines can be used to determine which Recipe components should be included.

### Resource Groups

Resource Groups identify the Kubernetes resources that participate in disaster recovery.

Every Recipe should define at least one Resource Group. Additional groups should be created when different sets of resources require different backup or restore behaviour.

Examples include:

- Separating infrastructure resources from application resources.
- Separating operator-managed resources from custom resources.
- Restoring dependent components in different stages.

---

### Volume Selection

Define a Volume Selection when the application uses one or more PersistentVolumeClaims (PVCs).

Volume Selection identifies the PVCs that should be protected as part of the application.

Applications that do not use persistent storage typically do not require a Volume Selection.

---

### Hooks

Hooks perform application-specific operations that cannot be achieved by protecting Kubernetes resources alone.

Examples include:

- Preparing an application before backup.
- Validating application health.
- Scaling workloads.
- Running standalone Kubernetes Jobs.

Hooks should only be defined when required by the application.

---

### Job Definitions

Job Definitions are required only when Job Hooks execute Kubernetes Jobs.

Applications that use only Exec, Check, or Scale Hooks do not require Job Definitions.

---

### Workflows

Workflows define the execution order of Resource Groups and Hooks.

Every Recipe should define both a backup workflow and a restore workflow.

The workflow should reflect the application's recovery requirements rather than the order in which the Recipe was written.

## Relating Application Characteristics to Recipe Components

The characteristics of an application determine the structure of the Recipe.

The following table summarizes common application characteristics and the Recipe components that are typically required.

| Application Characteristic | Typical Recipe Components |
|----------------------------|---------------------------|
| Stateless application | Resource Groups, Workflows |
| Application with PVCs | Resource Groups, Volume Selection, Workflows |
| Database | Resource Groups, Volume Selection, Exec Hook, Check Hook, Workflows |
| Application requiring workload quiescing | Scale Hook or Exec Hook |
| Application requiring readiness validation | Check Hook |
| Application requiring standalone scripts | Job Hook, Job Definitions |
| Multi-component application | Multiple Resource Groups, Ordered Workflows |
| Operator-managed application | Multiple Resource Groups, Ordered Workflows, Check Hooks (where appropriate) |

These mappings are guidelines rather than strict rules. The final Recipe should always reflect the recovery requirements of the application being protected.

### Design Checklist

Before continuing, confirm the following:

- The application resources have been identified.
- Persistent storage requirements are understood.
- Required hook types have been identified.
- Resource grouping strategy has been determined.
- Backup and restore workflows have been planned.

## Step 3 — Design Resource Groups

Resource Groups define how application resources are organized during backup and restore.

Rather than thinking of a Resource Group as a collection of Kubernetes resources, think of it as a collection of resources that should be processed together during disaster recovery.

A well-designed Resource Group reflects the logical structure of the application rather than the Kubernetes namespace in which it is deployed.

### When to create multiple Resource Groups

A single Resource Group is sufficient for simple applications whose resources can be backed up and restored together.

Create multiple Resource Groups when different parts of the application require different recovery behaviour.

Common scenarios include:

| Scenario | Recommended Approach |
|----------|----------------------|
| Operator-managed application | Separate operator resources from application resources. |
| Applications with multiple tiers | Create separate groups for each logical tier. |
| Resources restored in different stages | Create one group for each recovery stage. |
| Selective restore requirements | Create separate groups for independently restorable resources. |

### Designing Resource Groups

When defining a Resource Group:

- Include resources that share the same recovery lifecycle.
- Select resources using labels whenever possible.
- Avoid selecting unrelated resources from the same namespace.
- Exclude ephemeral resources that are recreated automatically.
- Keep Resource Groups focused on a single recovery purpose.

The workflow determines when each Resource Group is processed. Resource Groups should therefore be designed before the backup and restore workflows are defined.

### Choosing Resource Selectors

Resource selectors determine which Kubernetes resources participate in disaster recovery.

When possible, use label selectors instead of explicit resource names.

The recommended approach is:

1. Inspect the application's resources.
2. Identify labels common to all application resources.
3. Use the common labels in the Resource Group.
4. Verify that unrelated resources are not selected.

Examples of commonly used labels include:

- app
- app.kubernetes.io/name
- app.kubernetes.io/instance
- app.kubernetes.io/part-of

Avoid creating Resource Groups without a label selector unless the application cannot be identified using labels.

## Step 4 — Design Volume Selection

Applications that store persistent data require volume protection.

Volume Selection identifies the PersistentVolumeClaims (PVCs) that should be included in disaster recovery.

Not every application requires a Volume Selection. Stateless applications that do not own PersistentVolumeClaims typically do not require this section.

### Designing Volume Selection

When selecting PVCs:

- Include only the PVCs that belong to the application.
- Prefer label-based selection over explicit PVC names.
- Ensure that the selected labels are present on every PVC that requires protection.
- Verify that unrelated PVCs are not selected.

Applications managed by Kubernetes Operators sometimes use different labels for application resources and PersistentVolumeClaims.

In these cases, define the Volume Selection independently rather than reusing the labels from the Resource Groups.

### Design Considerations

Consider the following when designing Volume Selection:

- Does every PVC require protection?
- Are temporary or cache volumes excluded?
- Are the PVC labels consistent across deployments?
- Will the Recipe remain reusable across multiple environments?

## Step 5 — Determine Whether Hooks Are Required

Many applications can be protected by backing up Kubernetes resources and PersistentVolumeClaims alone.

Hooks should be introduced only when additional application-specific operations are required to ensure successful backup or recovery.

Typical reasons for using Hooks include:

- Preparing an application before backup.
- Waiting for resources to become ready.
- Temporarily stopping application activity.
- Executing standalone Kubernetes Jobs.
- Performing application-specific validation.

If the application can be safely backed up and restored without these operations, Hooks are usually unnecessary.

### Choosing a Hook Type

| Requirement | Recommended Hook |
|------------|------------------|
| Execute commands inside a running container | Exec Hook |
| Verify application health | Check Hook |
| Scale workloads before or after backup | Scale Hook |
| Execute a standalone Kubernetes Job | Job Hook |

### Hook Design Guidelines

When designing Hooks:

- Keep each Hook focused on a single responsibility.
- Reuse Hooks when the same operation is required by multiple workflows.
- Define appropriate timeout values.
- Choose failure policies that reflect the application's recovery requirements.
- Avoid performing unrelated operations within the same Hook.

## Step 6 — Design the Backup Workflow

The backup workflow defines the order in which Resource Groups and Hooks are executed during application protection.

A well-designed backup workflow prepares the application for backup, captures the required resources, and leaves the application in a consistent operational state.

The workflow should reflect the application's behaviour rather than the order in which the Recipe was authored.

### Designing the Backup Workflow

When designing a backup workflow, consider the following sequence:

1. Prepare the application for backup, if required.
2. Capture application resources.
3. Capture persistent storage.
4. Perform any post-backup operations required to return the application to its normal state.

Not every application requires all of these stages. Stateless applications may only require resource capture, while stateful applications often require additional preparation before persistent storage is protected.

### Common Backup Activities

Depending on the application, the backup workflow may include one or more of the following activities:

- Validate that the application is healthy before backup.
- Pause or quiesce application writes.
- Execute application-specific commands.
- Run prerequisite Kubernetes Jobs.
- Capture Kubernetes resources.
- Capture PersistentVolumeClaims.
- Resume normal application operation.

The workflow should contain only the operations required by the application.

### Workflow Design Guidelines

When designing a backup workflow:

- Execute preparation steps before protecting application data.
- Capture resources in an order that preserves application consistency.
- Avoid unnecessary Hooks.
- Keep the workflow easy to understand and maintain.
- Ensure every referenced Resource Group and Hook has been defined.

### Backup Design Checklist

Before defining the backup workflow, confirm the following:

- Does the application require preparation before backup?
- Are all required Resource Groups available?
- Are PersistentVolumeClaims protected?
- Are prerequisite Hooks executed before data capture?
- Does the workflow return the application to its normal operating state after backup?

## Step 7 — Design the Restore Workflow

The restore workflow defines how an application is reconstructed after a disaster recovery operation.

Unlike the backup workflow, which prepares an application for protection, the restore workflow focuses on rebuilding the application in a usable state.

Applications often require resources to become available in a specific order before dependent components can function correctly.

### Designing the Restore Workflow

When designing a restore workflow, consider the following sequence:

1. Restore persistent storage.
2. Restore Kubernetes resources.
3. Restore dependent application components.
4. Verify that the application has recovered successfully.

The exact sequence depends on the application's architecture and recovery requirements.

### Common Restore Activities

A restore workflow may include activities such as:

- Restoring PersistentVolumeClaims.
- Restoring Kubernetes resources.
- Waiting for workloads to become ready.
- Executing validation checks.
- Running post-recovery Jobs.

Only include activities that are required by the application.

### Workflow Design Guidelines

When designing a restore workflow:

- Restore dependencies before dependent workloads.
- Validate that workloads become healthy before considering the recovery complete.
- Avoid unnecessary recovery operations.
- Ensure workflow ordering reflects application dependencies.

### Restore Design Checklist

Before finalizing the restore workflow, confirm the following:

- Persistent storage is restored before workloads that depend on it.
- Application dependencies are restored in the required order.
- Recovery validation has been defined where appropriate.
- All referenced Resource Groups and Hooks exist.
- The workflow restores the application to an operational state.

## Step 8 — Validate the Recipe

Before using a Recipe to protect an application, verify that it accurately represents the application's disaster recovery requirements.

Validation should confirm both the correctness of the Recipe structure and the suitability of the design decisions.

A Recipe is complete only when every referenced resource, workflow, and application-specific operation has been defined and the recovery process reflects the application's behaviour.

### Validate Resource Selection

Verify that:

- All application resources that require protection are included.
- Resource selectors identify only the intended resources.
- Unrelated resources are excluded.
- PersistentVolumeClaims that require protection are included.
- Ephemeral resources are excluded unless explicitly required.

### Validate Resource Organization

Verify that:

- Resource Groups represent logical recovery units.
- Resources that should be recovered together belong to the same Resource Group.
- Resources requiring different recovery behaviour are separated into different Resource Groups.

### Validate Hooks

Verify that:

- Hooks are defined only when required.
- Each Hook performs a single responsibility.
- Hook execution order reflects the application lifecycle.
- Timeout and failure policies are appropriate for the application.

### Validate Workflows

Verify that:

- Both backup and restore workflows have been defined.
- Every workflow step references an existing Resource Group or Hook.
- Backup operations occur before resource protection where required.
- Restore operations respect application dependencies.
- Recovery validation occurs after application resources have been restored.

### Validate Recipe Completeness

Before considering the Recipe complete, verify that:

- Every referenced Resource Group exists.
- Every referenced Hook exists.
- Every referenced Job Definition exists.
- The workflow can execute without undefined references.
- The Recipe reflects the application's recovery requirements.

# Recipe Authoring Checklist

Use the following checklist before finalizing a Recipe.

## Application Analysis

- Application architecture has been reviewed.
- Kubernetes workloads have been identified.
- Persistent storage requirements are understood.
- Recovery dependencies have been identified.

## Recipe Structure

- Resource Groups have been defined.
- Volume Selection has been configured where required.
- Required Hooks have been identified.
- Job Definitions have been added where required.

## Workflow Design

- Backup workflow has been defined.
- Restore workflow has been defined.
- Workflow ordering reflects application dependencies.

## Validation

- Resource selectors have been verified.
- All workflow references exist.
- Failure policies have been reviewed.
- Recovery validation has been included where appropriate.

## Final Review

- The Recipe contains only components required by the application.
- The Recipe is reusable across deployments where possible.
- The Recipe accurately represents the application's recovery requirements.

## Recipe Design Principles

Keep the following principles in mind when designing Recipes.

### Design for the application

A Recipe should describe how the application behaves during backup and recovery rather than mirroring the Kubernetes objects that exist in the cluster.

### Keep the Recipe simple

Only include the Resource Groups, Hooks, and workflows required by the application.

Unnecessary components make Recipes more difficult to understand and maintain.

### Prefer reusable selectors

Use label-based selection whenever practical.

Avoid depending on resource names that may vary between deployments.

### Separate independent recovery stages

Use multiple Resource Groups when different components require different recovery behaviour.

### Design workflows around recovery

The workflow should describe the order in which the application should be protected and recovered rather than the order in which the Recipe was authored.

### Validate before use

Always verify that the Recipe selects the intended resources and accurately represents the application's recovery requirements before using it in production.