# ODF Disaster Recovery — Recipe Generation Rules

## Purpose

This document defines the rules that must be followed when generating an ODF Disaster Recovery `Recipe` Custom Resource.

This document is a **generation contract**. When a user asks for a Recipe, generate the Recipe according to these rules and the Recipe API.

The Recipe API Reference defines the supported fields and schema. This document defines how those fields must be assembled when generating a Recipe.

---

## 1. Generate a Recipe Custom Resource

When the user asks to generate an ODF Disaster Recovery Recipe, generate the actual `Recipe` Custom Resource.

Do not generate:

* a generic Kubernetes resource
* a Deployment or StatefulSet
* a PersistentVolumeClaim
* a DRPolicy
* a DRPlacementControl
* a VolumeReplicationGroup
* a Velero Backup or Schedule
* instructions for creating a Recipe
* a partially completed template

The generated resource must begin with:

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
```

---

## 2. Use placeholders for unknown values

When generating a Recipe, use values provided by the user or observed from the workload.

If a required application-specific value is unknown, use a clearly identifiable placeholder instead of inventing a value.

This applies to values such as:

- application name
- namespace
- application labels
- PVC labels
- Deployment names
- StatefulSet names
- PVC names
- container names
- commands

Do not infer missing values from the current project, repository, workspace, active file, or environment.

For example, when the application namespace is unknown, use:

```yaml
includedNamespaces:
  - <application-namespace>
```

---

## 3. Canonical Recipe skeleton

Unless the application requires additional components, start from the following valid Recipe structure and fill in application-specific values or placeholders.

```yaml
apiVersion: ramendr.openshift.io/v1alpha1
kind: Recipe
metadata:
  name: <recipe-name>
spec:
  appType: <application-type>

  groups:
    - name: app-resources
      type: resource
      includedNamespaces:
        - <application-namespace>
      labelSelector:
        matchLabels:
          app: <application-label>

  volumes:
    - name: app-volumes
      type: volume
      includedNamespaces:
        - <application-namespace>
      labelSelector:
        matchLabels:
          app: <pvc-label>

  workflows:
    - name: backup
      sequence:
        - group: app-resources

    - name: restore
      sequence:
        - group: app-resources
```

Use this structure as the default generation pattern.

- Resource Groups are defined under `spec.groups` and workflow steps reference them using `group:`.
- Volume selection is defined under `spec.volumes` using `labelSelector` or `nameSelector`.
- Workflow steps reference Resource Groups and Hooks only.
- Volume Groups are never workflow steps.
- Do not use unsupported fields such as `pvcSelector`, `persistentVolumeClaimNames`, `claimName`, or generic `selector`.

---

## 4. Use the Recipe API field names exactly

Use the Recipe API field names exactly as defined by the API.

The main Recipe sections are:

```yaml
spec:
  appType:
  groups:
  volumes:
  hooks:
  jobs:
  workflows:
```

Do not replace an API field with a similarly named field from another Kubernetes resource.

For application resources, use:

```yaml
groups:
```

not `resourceGroups`.

For PVC protection, use:

```yaml
volumes:
```

not a PVC-specific structure.

---

## 5. Resource Groups

Application resources are represented using entries under `spec.groups`.

A Resource Group uses:

```yaml
- name: <resource-group-name>
  type: resource
```

When the application namespace is known, include:

```yaml
includedNamespaces:
  - <application-namespace>
```

When application labels are known, include:

```yaml
labelSelector:
  matchLabels:
    <label-key>: <label-value>
```

Use the actual namespace and labels supplied by the user.

Do not invent application labels or namespaces.

### Resource Group selection

Resource Groups select application resources using the supported Recipe Group selection fields.

For an application consisting of a Deployment, use the Resource Group to select the Deployment. The PersistentVolumeClaim is selected independently through `spec.volumes`.

For example, when the Deployment name is unknown but its application label is known:

```yaml
groups:
  - name: app-resources
    type: resource
    includedNamespaces:
      - <application-namespace>
    labelSelector:
      matchLabels:
        app: <application-label>
```
---

## 6. Volume Groups

PersistentVolumeClaims that require protection are represented under `spec.volumes`.

Each volume entry uses the Recipe Group schema with:

```yaml
volumes:
  - name: <volume-group-name>
    type: volume
    includedNamespaces:
      - <application-namespace>
    labelSelector:
      matchLabels:
        <pvc-label-key>: <pvc-label-value>
```

For label-based PVC selection:

- use `type: volume`
- use `includedNamespaces` when the application namespace is known
- use `labelSelector`, not `selector`
- use the actual PVC labels supplied by the user

For name-based PVC selection, use `nameSelector`.

Do not generate PVC-specific fields such as `selector`, `pvcSelector`, `persistentVolumeClaimNames`, or `claimName`.

Do not generate unsupported structures such as:

```yaml
volumes:
  - name: app-pvc
    selector:
      matchLabels:
        app: my-app
```

or:

```yaml
volumes:
  - name: app-pvc
    persistentVolumeClaimNames:
      - my-app-pvc
```

For example, for label-based selection:

```yaml
volumes:
  - name: app-volumes
    type: volume
    includedNamespaces:
      - my-app
    labelSelector:
      matchLabels:
        app: my-app
```

For example, for name-based selection:

```yaml
volumes:
  - name: app-volumes
    type: volume
    includedNamespaces:
      - my-app
    nameSelector: my-app-pvc
```

---

## 7. Volume Selection and Workflows

Volume selection is defined under `spec.volumes`.

A volume group identifies the PersistentVolumeClaims that ODF Disaster Recovery should protect. Volume selection is defined independently from workflow execution.

For label-based PVC selection, use the following structure:

```yaml
volumes:
  - name: <volume-group-name>
    type: volume
    includedNamespaces:
      - <application-namespace>
    labelSelector:
      matchLabels:
        <pvc-label-key>: <pvc-label-value>
```

For an application with one Deployment and one PVC selected by the same application label, use the following structure:

```yaml
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

Workflows operate on Resource Groups and Hooks. A volume group is not a workflow step.

When generating a Recipe, preserve this separation:

- `spec.groups` defines application resources.
- `spec.volumes` defines PVC selection.
- `workflows[].sequence` defines the execution order of Resource Groups and Hooks.

Do not reference a volume group from `workflows[].sequence`.
---

## 8. Application-specific values

Use values provided by the user or observed from the workload.

Use actual:

* application names
* namespaces
* labels
* Deployment names
* StatefulSet names
* PVC names
* container names
* commands

If a required value is unknown, retain the correct Recipe field and use a clearly identifiable placeholder for that value.

For example:

```yaml
includedNamespaces:
  - <application-namespace>

labelSelector:
  matchLabels:
    app: <application-label>
```

When the application namespace is unknown, use `<application-namespace>`. Do not substitute an inferred or environment-specific namespace.

Do not invent a value merely to make the YAML appear complete.

When the user explicitly provides a namespace or label, use that exact value in both `spec.groups` and `spec.volumes` when it applies to both resources and PVCs.

---

## 9. Optional Components

Add `volumes` only when the application has PersistentVolumeClaims that require protection.

Add `hooks` only when application-specific operations or validation are required.

Add `jobs` only when required by a Job Hook.

When generating a Job Hook:

- use `type: job`
- define the Job Hook under `spec.hooks`
- reference one or more jobs using `hooks[].jobs`
- define the corresponding Job manifest under `spec.jobs`
- reference the Job Hook in workflows using `hook: <hook-name>/<job-name>`

Do not model a Job Hook as an Exec Hook that runs a Job.

Do not add optional sections merely to make the Recipe appear more complete.

---

## 10. Workflows

Every complete Recipe must define both a `backup` workflow and a `restore` workflow.

A workflow contains an ordered `sequence` of steps. Each step references a Resource Group or a Hook.

For a basic application without application-specific Hooks, use:

```yaml
workflows:
  - name: backup
    sequence:
      - group: app-resources

  - name: restore
    sequence:
      - group: app-resources
```

For a workflow that requires application-specific operations, Hooks can be added to the sequence:

```yaml
workflows:
  - name: backup
    sequence:
      - hook: prepare-backup/flush
      - group: app-resources

  - name: restore
    sequence:
      - group: app-resources
      - hook: validate-application/ready
```

When a Hook defines named operations, checks, or jobs, the workflow step must reference the specific name, not only the Hook name.

Workflow sequence steps use the following structures:
```yaml
- group: <resource-group-name>
```

or:
```yaml
- hook: <hook-name>/<operation-or-check-name>
```

For Hook workflow steps:

- Exec Hooks use `hook: <hook-name>/<op-name>`
- Check Hooks use `hook: <hook-name>/<check-name>`
- Job Hooks use `hook: <hook-name>/<job-name>`
- Scale Hooks should use the operation names shown in the Recipe examples

The names referenced by workflow steps must correspond to components defined elsewhere in the Recipe.

The spec.volumes section defines PVC selection independently from workflow sequencing. Volume selection is therefore represented under spec.volumes, while workflow sequences reference Resource Groups and Hooks.
---

## 11. Generation Validation

Before returning a generated Recipe, verify all of the following:

1. `apiVersion` is `ramendr.openshift.io/v1alpha1`.
2. `kind` is `Recipe`.
3. `spec.appType` is present.
4. Application resources are defined under `spec.groups`.
5. Resource Groups have `type: resource`.
6. PVC selection, when required, is defined under `spec.volumes`.
7. Every volume entry has `type: volume`.
8. Label-based PVC selection uses `labelSelector`, not `selector` or `pvcSelector`.
9. Name-based PVC selection uses `nameSelector`, not `persistentVolumeClaimNames` or `claimName`.
10. Volume Groups are never referenced by `workflows[].sequence`.
11. Every `group` workflow step references a Resource Group defined under `spec.groups`.
12. Every `hook` workflow step references a valid Hook operation, check, or job.
13. Both `backup` and `restore` workflows are present for a complete Recipe.
14. No unsupported fields such as `resourceGroups`, `claimName`, `pvcSelector`, generic `selector`, or `persistentVolumeClaimNames` are generated.
15. User-provided namespaces and labels are used directly.
16. Unknown values use placeholders rather than invented values.
17. Optional sections are generated only when required.