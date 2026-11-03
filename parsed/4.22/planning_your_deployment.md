---
title: "planning your deployment"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/planning_your_deployment"
source_file: "planning_your_deployment.pdf"
version: "4.22"
chapter: "planning_your_deployment"
---

Red Hat OpenShift Data Foundation
               4.22

    Planning your deployment

    Important considerations when deploying Red Hat OpenShift Data Foundation










    Last Updated: 2026-09-04



Red Hat OpenShift Data Foundation 4.22 Planning your deployment

Important considerations when deploying Red Hat OpenShift Data Foundation

Legal Notice

Copyright © Red Hat.

Except as otherwise noted below, the text of and illustrations in this documentation are licensed by
Red Hat under the Creative Commons Attribution–Share Alike 3.0 Unported license . If you
distribute this document or an adaptation of it, you must provide the URL for the original version.

Red Hat, as the licensor of this document, waives the right to enforce, and agrees not to assert,
Section 4d of CC-BY-SA to the fullest extent permitted by applicable law.

Red Hat, the Red Hat logo, JBoss, Hibernate, and RHCE are trademarks or registered trademarks of
Red Hat, LLC. or its subsidiaries in the United States and other countries.

Linux ® is the registered trademark of Linus Torvalds in the United States and other countries.

XFS is a trademark or registered trademark of Hewlett Packard Enterprise Development LP or its
subsidiaries in the United States and other countries.

The OpenStack ® Word Mark and OpenStack logo are trademarks or registered trademarks of the
Linux Foundation, used under license.

All other trademarks are the property of their respective owners.


    Abstract

    Read this document for important considerations when planning your Red Hat OpenShift Data
    Foundation deployment.

                                                                Table of Contents

      Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
MAKING OPEN SOURCE MORE INCLUSIVE                                               4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION                                     5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. INTRODUCTION TO OPENSHIFT DATA FOUNDATION                            6

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. ARCHITECTURE OF OPENSHIFT DATA FOUNDATION                            7
 2.1. ABOUT OPERATORS                                                           8
 2.2. STORAGE CLUSTER DEPLOYMENT APPROACHES                                     9
  2.2.1. Internal approach                                                      9
  2.2.2. External approach                                                     10
 2.3. NODE TYPES                                                               10

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. INTERNAL STORAGE SERVICES                                           12

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. EXTERNAL STORAGE SERVICES                                           13

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. SECURITY CONSIDERATIONS                                             14
 5.1. FIPS CRYPTOGRAPHY                                                        14
 5.2. PROXY ENVIRONMENT                                                        14
 5.3. DATA ENCRYPTION OPTIONS                                                  14
 5.4. CLUSTER-WIDE ENCRYPTION                                                  15
 5.5. STORAGE CLASS ENCRYPTION                                                 16
 5.5.1. CipherTrust manager                                                    16
 5.6. DATA ENCRYPTION IN-TRANSIT USING RED HAT CEPH STORAGE’S MESSENGER VERSION 2
 PROTOCOL (MSGR2)                                                              17
 5.7. OBJECT STORAGE DATA ENCRYPTION SUPPORTABILITY                            17
 5.8. ENCRYPTION IN TRANSIT                                                    18
 5.9. CENTRALIZING TLS SECURITY PROFILES FOR OPENSHIFT DATA FOUNDATION         18
 5.10. NETWORK ACCESS CONSIDERATIONS FOR MULTICLOUD OBJECT GATEWAY IN CLOUD
 ENVIRONMENTS                                                                  21
 5.11. MULTICLOUD OBJECT GATEWAY PUBLIC ROUTES                                 21

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 6. SUBSCRIPTIONS                                                       22
 6.1. SUBSCRIPTION OFFERINGS                                                   22
 6.2. DISASTER RECOVERY SUBSCRIPTION REQUIREMENT                               22
 6.3. CORES VERSUS VCPUS AND HYPERTHREADING                                    22
 6.4. SPLITTING CORES                                                          22
  6.4.1. Shared Processor Pools for IBM Power                                  23
 6.5. SUBSCRIPTION REQUIREMENTS                                                23

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 7. INFRASTRUCTURE REQUIREMENTS                                         24
 7.1. PLATFORM REQUIREMENTS                                                    24
  7.1.1. Amazon EC2                                                            24
  7.1.2. Bare Metal                                                            24
  7.1.3. VMware vSphere                                                        24
  7.1.4. Microsoft Azure                                                       25
  7.1.5. Google Cloud                                                          25
  7.1.6. Red Hat OpenStack Platform [Technology Preview]                       25
  7.1.7. IBM Power                                                             25
  7.1.8. IBM Z and IBM® LinuxONE                                               26
  7.1.9. ROSA with hosted control planes (HCP)                                 26
  7.1.10. Any platform                                                         26
 7.2. EXTERNAL MODE REQUIREMENT                                                26

                                                                                1

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

     7.2.1. Red Hat Ceph Storage 26
     7.2.2. IBM FlashSystem 26
   7.3. RESOURCE REQUIREMENTS 26
     7.3.1. Resource requirements for IBM Z and IBM LinuxONE infrastructure 28
     7.3.2. Minimum deployment resource requirements 29
     7.3.3. Compact deployment resource requirements 30
     7.3.4. Resource requirements for MCG only deployment 30
     7.3.5. Resource requirements for using Network File System 30
     7.3.6. Resource requirements for performance profiles 31
   7.4. POD PLACEMENT RULES 31
   7.5. STORAGE DEVICE REQUIREMENTS 32
     7.5.1. Dynamic storage devices 32
     7.5.2. Local storage devices 32
     7.5.3. Capacity planning 32
   7.6. ERASURE CODING CONSIDERATIONS 33
     7.6.1. Supported erasure coding profiles 33
     7.6.2. Node requirements 34
     7.6.3. Limitations 34

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 8. NETWORK REQUIREMENTS 35
   8.1. IPV6 SUPPORT 35
   8.2. MULTI NETWORK PLUG-IN (MULTUS) SUPPORT 35

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 9. DISASTER RECOVERY 36

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 10. METRO-DR 37

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 11. REGIONAL-DR 39

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 12. DISASTER RECOVERY WITH STRETCH CLUSTER 41

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 13. DISCONNECTED ENVIRONMENT 42

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 14. PERFORMANCE CONSIDERATIONS AND BENCHMARKING 44

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 15. SUPPORTED AND UNSUPPORTED FEATURES FOR IBM POWER AND IBM Z 45

 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 CHAPTER 16. NEXT STEPS 47










    2

Table of Contents










3

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        MAKING OPEN SOURCE MORE INCLUSIVE

  Red Hat is committed to replacing problematic language in our code, documentation, and web
  properties. We are beginning with these four terms: master, slave, blacklist, and whitelist. Because of the
  enormity of this endeavor, these changes will be implemented gradually over several upcoming releases.
  For more details, see our CTO Chris Wright’s message .










    4

             PROVIDING FEEDBACK ON RED HAT DOCUMENTATION

PROVIDING FEEDBACK ON RED HAT DOCUMENTATION

We appreciate your input on our documentation. Do let us know how we can make it better.

To give feedback, create a Jira ticket:

Procedure

    1.   Log in to the Jira.

    2.   Click Create in the top navigation bar

    3.   Enter a descriptive title in the Summary field.

    4.   Enter your suggestion for improvement in the Description field. Include links to the relevant
         parts of the documentation.

    5.   Select Documentation in the Components field.

    6.   Click Create at the bottom of the dialogue.










         5

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        CHAPTER 1. INTRODUCTION TO OPENSHIFT DATA
        FOUNDATION

  Red Hat OpenShift Data Foundation is a highly integrated collection of cloud storage and data services
  for Red Hat OpenShift Container Platform. It is available as part of the Red Hat OpenShift Container
  Platform Service Catalog, packaged as an operator to facilitate simple deployment and management.

  Red Hat OpenShift Data Foundation services are primarily made available to applications by way of
  storage classes that represent the following components:

        Block storage devices, catering primarily to database workloads. Prime examples include Red
        Hat OpenShift Container Platform logging and monitoring, and PostgreSQL.

        IMPORTANT

        Block storage should be used for any worklaod only when it does not require sharing the
        data across multiple containers.

        Shared and distributed file system, catering primarily to software development, messaging, and
        data aggregation workloads. Examples include Jenkins build sources and artifacts, Wordpress
        uploaded content, Red Hat OpenShift Container Platform registry, and messaging using JBoss
        AMQ.

        Multicloud object storage, featuring a lightweight S3 API endpoint that can abstract the storage
        and retrieval of data from multiple cloud object stores.

        On premises object storage, featuring a robust S3 API endpoint that scales to tens of petabytes
        and billions of objects, primarily targeting data intensive applications. Examples include the
        storage and access of row, columnar, and semi-structured data with applications like Spark,
        Presto, Red Hat AMQ Streams (Kafka), and even machine learning frameworks like TensorFlow
        and Pytorch.

        NOTE

        Running PostgresSQL workload on CephFS persistent volume is not supported and it is
        recommended to use RADOS Block Device (RBD) volume. For more information, see the
        knowledgebase solution on ODF Database Workloads.

  Red Hat OpenShift Data Foundation version 4.x integrates a collection of software projects, including:

        Ceph, providing block storage, a shared and distributed file system, and on-premises object
        storage

        Ceph CSI, to manage provisioning and lifecycle of persistent volumes and claims

        NooBaa, providing a Multicloud Object Gateway

        OpenShift Data Foundation, Rook-Ceph, and NooBaa operators to initialize and manage
        OpenShift Data Foundation services.

  Additional resources

        ODF Database Workloads Must Not Use CephFS PVs/PVCs



    6

        CHAPTER 2. ARCHITECTURE OF OPENSHIFT DATA FOUNDATION

    CHAPTER 2. ARCHITECTURE OF OPENSHIFT DATA
                    FOUNDATION

    Red Hat OpenShift Data Foundation provides services for, and can run internally from the Red Hat
    OpenShift Container Platform.

    Figure 2.1. Red Hat OpenShift Data Foundation architecture










Red Hat OpenShift Data Foundation supports deployment into Red Hat OpenShift Container Platform
clusters deployed on installer-provisioned or user-provisioned infrastructure.

For details about these two approaches, see OpenShift Container Platform - Installation process .

To ensure your environment meets the requirements, consult the Red Hat OpenShift Data Foundation
Supportability and Interoperability Checker to verify component compatibility, supported limits,
features, and integration matrices.


    7

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

 For information about the architecture and lifecycle of OpenShift Container Platform, see OpenShift
 Container Platform architecture.

 TIP

 For IBM Power, see Installing on IBM Power .

 2.1. ABOUT OPERATORS

 Red Hat OpenShift Data Foundation comprises of three main operators, which codify administrative
 tasks and custom resources so that you can easily automate the task and resource characteristics.
 Administrators define the desired end state of the cluster, and the OpenShift Data Foundation
 operators ensure the cluster is either in that state, or approaching that state, with minimal administrator
 intervention.

 OpenShift Data Foundation operator

 A meta-operator that draws on other operators in specific tested ways to codify and enforce the
 recommendations and requirements of a supported Red Hat OpenShift Data Foundation deployment.
 The rook-ceph and noobaa operators provide the storage cluster resource that wraps these resources.

 Rook-ceph operator

 This operator automates the packaging, deployment, management, upgrading, and scaling of persistent
 storage and file, block, and object services. It creates block and file storage classes for all environments,
 and creates an object storage class and services Object Bucket Claims (OBCs) made against it in on-
 premises environments.

 Additionally, for internal mode clusters, it provides the ceph cluster resource, which manages the
 deployments and services representing the following:

    Object Storage Daemons (OSDs)

    Monitors (MONs)

    Manager (MGR)

    Metadata servers (MDS)

    RADOS Object Gateways (RGWs) on-premises only

 Multicloud Object Gateway operator

 This operator automates the packaging, deployment, management, upgrading, and scaling of the
 Multicloud Object Gateway (MCG) object service. It creates an object storage class and services the
 OBCs made against it.

 Additionally, it provides the NooBaa cluster resource, which manages the deployments and services for
 NooBaa core, database, and endpoint.

              NOTE

              OpenShift Data Foundation’s default configuration for MCG is optimized for low
              resource consumption and not performance. If you plan to use MCG often, see
              information about increasing resource limits in the knowledebase article Performance
              tuning guide for Multicloud Object Gateway.

8

             CHAPTER 2. ARCHITECTURE OF OPENSHIFT DATA FOUNDATION

2.2. STORAGE CLUSTER DEPLOYMENT APPROACHES

The growing list of operating modalities is an evidence that flexibility is a core tenet of Red Hat
OpenShift Data Foundation. This section provides you with information that will help you to select the
most appropriate approach for your environments.

You can deploy Red Hat OpenShift Data Foundation either entirely within OpenShift Container
Platform (Internal approach) or to make available the services from a cluster running outside of
OpenShift Container Platform (External approach).

2.2.1. Internal approach

Deployment of Red Hat OpenShift Data Foundation entirely within Red Hat OpenShift Container
Platform has all the benefits of operator based deployment and management. You can use the internal-
attached device approach in the graphical user interface (GUI) to deploy Red Hat OpenShift Data
Foundation in internal mode using the local storage operator and local storage devices.

Ease of deployment and management are the highlights of running OpenShift Data Foundation services
internally on OpenShift Container Platform. There are two different deployment modalities available
when Red Hat OpenShift Data Foundation is running entirely within Red Hat OpenShift Container
Platform:

 Simple

 Optimized

Simple deployment

Red Hat OpenShift Data Foundation services run co-resident with applications. The operators in Red
Hat OpenShift Container Platform manages these applications.

A simple deployment is best for situations where,

 Storage requirements are not clear.

 Red Hat OpenShift Data Foundation services runs co-resident with the applications.

 Creating a node instance of a specific size is difficult, for example, on bare metal.

For Red Hat OpenShift Data Foundation to run co-resident with the applications, the nodes must have
local storage devices, or portable storage devices attached to them dynamically, like EBS volumes on
EC2, or vSphere Virtual Volumes on VMware, or SAN volumes.

         NOTE

         PowerVC dynamically provisions the SAN volumes.

Optimized deployment Red Hat OpenShift Data Foundation services run on dedicated infrastructure
nodes. Red Hat OpenShift Container Platform manages these infrastructure nodes.

An optimized approach is best for situations when,

 Storage requirements are clear.

 Red Hat OpenShift Data Foundation services run on dedicated infrastructure nodes.


         9

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

  Creating a node instance of a specific size is easy, for example, on cloud, virtualized
  environment, and so on.

2.2.2. External approach

Red Hat OpenShift Data Foundation exposes the Red Hat Ceph Storage services running outside of the
OpenShift Container Platform cluster as storage classes.

The external approach is best used when,

  Storage requirements are significant (600+ storage devices).

  Multiple OpenShift Container Platform clusters need to consume storage services from a
  common external cluster.

  Another team, Site Reliability Engineering (SRE), storage, and so on, needs to manage the
  external cluster providing storage services. Possibly a pre-existing one.

2.3. NODE TYPES

Nodes run the container runtime, as well as services, to ensure that the containers are running, and
maintain network communication and separation between the pods. In OpenShift Data Foundation,
there are three types of nodes.

Table 2.1. Types of nodes

 Node Type                     Description

 Master                        These nodes run processes that expose the Kubernetes API, watch and
                               schedule newly created pods, maintain node health and quantity, and
                               control interaction with underlying cloud providers.

 Infrastructure (Infra)        Infra nodes run cluster level infrastructure services such as logging,
                               metrics, registry, and routing. These are optional in OpenShift Container
                               Platform clusters. In order to separate OpenShift Data Foundation layer
                               workload from applications, ensure that you use infra nodes for
                               OpenShift Data Foundation in virtualized and cloud environments.

                               To create Infra nodes, you can provision new nodes labeled as infra. For
                               more information, see How to use dedicated worker nodes for Red Hat
                               OpenShift Data Foundation










10

    CHAPTER 2. ARCHITECTURE OF OPENSHIFT DATA FOUNDATION


    Node Type     Description

    Worker        Worker nodes are also known as application nodes since they run
                  applications.

                  When OpenShift Data Foundation is deployed in internal mode, you
                  require a minimal cluster of 3 worker nodes. Make sure that the nodes
                  are spread across 3 different racks, or availability zones, to ensure
                  availability. In order for OpenShift Data Foundation to run on worker
                  nodes, you need to attach the local storage devices, or portable storage
                  devices to the worker nodes dynamically.

                  When OpenShift Data Foundation is deployed in external mode, it runs
                  on multiple nodes. This allows Kubernetes to reschedule on the available
                  nodes in case of a failure.


NOTE

OpenShift Data Foundation requires the same number of subsciptions as OpenShift
Container Platform. However, if OpenShift Data Foundation is running on infra nodes,
OpenShift does not require OpenShift Container Platform subscription for these nodes.
Therefore, the OpenShift Data Foundation control plane does not require additional
OpenShift Container Platform and OpenShift Data Foundation subscriptions. For more
information, see Subscriptions.










                  11

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

 CHAPTER 3. INTERNAL STORAGE SERVICES

Red Hat OpenShift Data Foundation service is available for consumption internally to the Red Hat
OpenShift Container Platform that runs on the following infrastructure:

 Amazon Web Services (AWS)

 Bare metal

 VMware vSphere

 Microsoft Azure

 Google Cloud

 Red Hat OpenStack 13 or higher (installer-provisioned infrastructure) [Technology Preview]

 IBM Power

 IBM Z and IBM® LinuxONE

 ROSA with hosted control planes (HCP)

Creation of an internal cluster resource results in the internal provisioning of the OpenShift Data
Foundation base services, and makes additional storage classes available to the applications.










12

     CHAPTER 4. EXTERNAL STORAGE SERVICES

 CHAPTER 4. EXTERNAL STORAGE SERVICES

This reference describes how to make external Red Hat Ceph Storage services available for
consumption through OpenShift Container Platform. Red Hat OpenShift Data Foundation can make
services from an external Red Hat Ceph Storage cluster available for consumption through OpenShift
Container Platform clusters running on the following platforms:

 VMware vSphere

 Bare metal

 Red Hat OpenStack platform (Technology Preview)

 IBM Power

 IBM Z

The OpenShift Data Foundation operators create and manage services to satisfy Persistent Volume
(PV) and Object Bucket Claims (OBCs) against the external services. External cluster can serve block,
file and object storage classes for applications that run on OpenShift Container Platform. The operators
do not deploy or manage the external clusters.










 13

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        CHAPTER 5. SECURITY CONSIDERATIONS

  5.1. FIPS CRYPTOGRAPHY

  The Federal Information Processing Standard Publication 140-2/140-3 (FIPS 140-2/140-3) is a
  standard that defines a set of security requirements for the use of cryptographic modules. Law
  mandates this standard for the US government agencies and contractors and is also referenced in other
  international and industry specific standards.

  Red Hat OpenShift Data Foundation is designed for FIPS. When running Red Hat Enterprise Linux
  (RHEL) or Red Hat Enterprise Linux CoreOS (RHCOS) booted in FIPS mode, OpenShift Container
  Platform core components, including OpenShift Data Foundation, use the RHEL cryptographic libraries
  that have been submitted to NIST for FIPS 140-2/140-3 Validation on only the x86_64, ppc64le, and
  s390x architectures. For more information, see the Support for FIPS cryptography section in OpenShift
  documentation.

  Currently, the Cryptographic Module Validation Program (CMVP) processes the cryptography modules.
  You can see the state of these modules at Modules in Process List . For more up-to-date information,
  see the Red Hat Knowledgebase solution RHEL core crypto components .

  5.2. PROXY ENVIRONMENT

  A proxy environment is a production environment that denies direct access to the internet and provides
  an available HTTP or HTTPS proxy instead. Red Hat Openshift Container Platform is configured to use a
  proxy by modifying the proxy object for existing clusters or by configuring the proxy settings in the
  install-config.yaml file for new clusters.

  Red Hat supports deployment of OpenShift Data Foundation in proxy environments when OpenShift
  Container Platform has been configured according to configuring the cluster-wide proxy.

  5.3. DATA ENCRYPTION OPTIONS

  Encryption lets you encode your data to make it impossible to read without the required encryption keys.
  This mechanism protects the confidentiality of your data in the event of a physical security breach that
  results in a physical media to escape your custody. The per-PV encryption also provides access
  protection from other namespaces inside the same OpenShift Container Platform cluster. Data is
  encrypted when it is written to the disk, and decrypted when it is read from the disk. Working with
  encrypted data might incur a small penalty to performance.

  Encryption is only supported for new clusters deployed using Red Hat OpenShift Data Foundation 4.6 or
  higher. An existing encrypted cluster that is not using an external Key Management System (KMS)
  cannot be migrated to use an external KMS.

  Previously, HashiCorp Vault was the only supported KMS for Cluster-wide and Persistent Volume
  encryptions. With OpenShift Data Foundation 4.7.0 and 4.7.1, only HashiCorp Vault Key/Value (KV)
  secret engine API, version 1 is supported. Starting with OpenShift Data Foundation 4.7.2, HashiCorp
  Vault KV secret engine API, versions 1 and 2 are supported. As of OpenShift Data Foundation 4.12,
  Thales CipherTrust Manager has been introduced as an additional supported KMS.







    IMPORTANT
    14

                                                                        CHAPTER 5. SECURITY CONSIDERATIONS

     IMPORTANT

                    KMS is required for StorageClass encryption, and is optional for cluster-wide
                    encryption.

                    To start with, Storage class encryption requires a valid Red Hat OpenShift Data
                    Foundation Advanced subscription. For more information, see the
                    knowledgebase article on OpenShift Data Foundation subscriptions.

Red Hat works with the technology partners to provide this documentation as a service to the
customers. However, Red Hat does not provide support for the Hashicorp product. For technical
assistance with this product, contact Hashicorp.

Additional resources

 Knowledgebase article on OpenShift Data Foundation subscriptions

 Hashicorp technical support services and policies

5.4. CLUSTER-WIDE ENCRYPTION

Red Hat OpenShift Data Foundation supports cluster-wide encryption (encryption-at-rest) for all the
disks and Multicloud Object Gateway operations in the storage cluster. OpenShift Data Foundation uses
Linux Unified Key System (LUKS) version 2 based encryption with a key size of 512 bits and the aes-xts-
plain64 cipher where each device has a different encryption key. The keys are stored using a Kubernetes
Secret or an external KMS. Both methods are mutually exclusive and you can not migrate between
methods.

Encryption is disabled by default for block and file storage. You can enable encryption for the cluster at
the time of deployment. The MultiCloud Object Gateway supports encryption by default. See the
deployment guides for more information.

OpenShift Data Foundation supports cluster wide encryption with and without Key Management
System (KMS). Cluster wide encryption with KMS is supported using the following service providers:

 HashiCorp Vault - Tested with Version 1.9.1 and 1.21.1

 Thales Cipher Trust Manager - Tested with Version 2.9.0+7637, crypto version: 1.5.0

Security common practices require periodic encryption key rotation. OpenShift Data Foundation
automatically rotates encryption keys stored in Kubernetes Secret (non-KMS) and Vault on a weekly
basis. However, key rotation for Vault KMS must be enabled after the storage cluster creation and does
not happen by default. For more information refer to the deployment guides.

        NOTE

        Using KMS requires a valid Red Hat OpenShift Data Foundation Advanced subscription.
        To know how subscriptions for OpenShift Data Foundation work, see the knowledgebase
        article on OpenShift Data Foundation subscriptions.

Cluster wide encryption with HashiCorp Vault KMS provides two authentication methods:

 Token: This method allows authentication using vault tokens. A Kubernetes Secret containing
 the vault token is created in the openshift-storage namespace and is used for authentication. If
 this authentication method is selected then the administrator has to provide the vault token

                                                                                                        15

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        that provides access to the backend path in Vault, where the encryption keys are stored.

        Kubernetes: This method allows authentication with vault using serviceaccounts. If this
        authentication method is selected then the administrator has to provide the name of the role
        configured in Vault that provides access to the backend path, where the encryption keys are
        stored. The value of this role is then added to the ocs-kms-connection-details config map.

        NOTE

        OpenShift Data Foundation on IBM Cloud platform supports Hyper Protect Crypto
        Services (HPCS) Key Management Services (KMS) as the encryption solution in addition
        to HashiCorp Vault KMS.

        IMPORTANT

        Red Hat works with the technology partners to provide this documentation as a service to
        the customers. However, Red Hat does not provide support for the Hashicorp product.
        For technical assistance with this product, contact Hashicorp.

  Additional resources

        Knowledgebase article on OpenShift Data Foundation subscriptions

        Hashicorp technical support services and policies

  5.5. STORAGE CLASS ENCRYPTION

  You can encrypt persistent volumes (block only) with storage class encryption using an external Key
  Management System (KMS) to store device encryption keys. Persistent volume encryption is only
  available for RADOS Block Device (RBD) persistent volumes. For information about how to create a
  storage class with persistent volume encryption, see the managing and allocating storage resources
  guide.

  Storage class encryption is supported in OpenShift Data Foundation 4.7 or higher with HashiCorp Vault
  KMS. Storage class encryption is supported in OpenShift Data Foundation 4.12 or higher with both
  HashiCorp Vault KMS and Thales CipherTrust Manager KMS.

        NOTE

        Requires a valid Red Hat OpenShift Data Foundation Advanced subscription. To know
        how subscriptions for OpenShift Data Foundation work, see the knowledgebase article on
        OpenShift Data Foundation subscriptions.

  5.5.1. CipherTrust manager

  Red Hat OpenShift Data Foundation version 4.12 introduced Thales CipherTrust Manager as an
  additional Key Management System (KMS) provider for your deployment. Thales CipherTrust Manager
  provides centralized key lifecycle management. CipherTrust Manager supports Key Management
  Interoperability Protocol (KMIP), which enables communication between key management systems.

  CipherTrust Manager is enabled during deployment.

  Additional resources


    16

          CHAPTER 5. SECURITY CONSIDERATIONS

      Knowledgebase article on OpenShift Data Foundation subscriptions

      Creating a storage class with persistent volume encryption

5.6. DATA ENCRYPTION IN-TRANSIT USING RED HAT CEPH
STORAGE’S MESSENGER VERSION 2 PROTOCOL (MSGR2)

Starting with OpenShift Data Foundation version 4.14, Red Hat Ceph Storage’s messenger version 2
protocol can be used to encrypt data in-transit. This provides an important security requirement for your
infrastructure.

In‑transit encryption can be enabled during deployment while the cluster is being created, and it can also
be enabled on existing clusters after deployment. For deployment‑time configuration, see the relevant
deployment guide for your environment. For post‑deployment configuration, see the managing and
allocating storage resources guide.

The msgr2 protocol supports two connection modes:

crc

      Provides strong initial authentication when a connection is established with cephx.

      Provides a crc32c integrity check to protect against bit flips.

      Does not provide protection against a malicious man-in-the-middle attack.

      Does not prevent an eavesdropper from seeing all post-authentication traffic.

secure

      Provides strong initial authentication when a connection is established with cephx.

      Provides full encryption of all post-authentication traffic.

      Provides a cryptographic integrity check.

The default mode is crc.

Additional resources

      Red Hat OpenShift Data Foundation deployment guides

      Enabling in‑transit encryption after deployment

5.7. OBJECT STORAGE DATA ENCRYPTION SUPPORTABILITY

Red Hat OpenShift Data Foundation supports Server-Side Encryption (SSE) to ensure the
confidentiality of data at rest. The supportability of specific SSE modes (SSE-C, SSE-S3, and SSE-
KMS) varies between the Multicloud Object Gateway and the Rados Gateway.

Multicloud Object Gateway (MCG)

The Multicloud Object Gateway provides encryption at rest by default for all backing stores.



SSE-C: Supported. Clients can manage their own encryption keys and provide them with each
          17

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        SSE-C: Supported. Clients can manage their own encryption keys and provide them with each
        read or write request (x-amz-server-side-encryption-customer-key). MCG performs the
        encryption and decryption, but does not store the key.

        SSE-S3: Supported and enabled by default. MCG encrypts data by default using system-
        managed keys, providing behavior consistent with SSE-S3.

        SSE-KMS: Supported. MCG integrates seamlessly with external Key Management Systems
        (KMS) if configured for cluster wide-encryption.

Rados Gateway (RGW)

The Rados Gateway (Ceph Object Storage) supports granular Server-Side Encryption options for S3
clients.

        SSE-C: Supported. Clients can manage their own encryption keys and provide them with each
        read or write request (x-amz-server-side-encryption-customer-key). RGW performs the
        encryption and decryption, but does not store the key.

        SSE-S3: Not supported.

        SSE-KMS: Not supported.

         NOTE

         Enabling encryption with an external KMS requires a valid Red Hat OpenShift Data
         Foundation Advanced subscription. To know how subscriptions for OpenShift Data
         Foundation work, see the knowledgebase article on OpenShift Data Foundation
         subscriptions.

Additional resources

        Knowledgebase article on OpenShift Data Foundation subscriptions

5.8. ENCRYPTION IN TRANSIT

You need to enable IPsec so that all the network traffic between the nodes on the OVN-Kubernetes
Container Network Interface (CNI) cluster network travels through an encrypted tunnel.

By default, IPsec is disabled. You can enable it either during or after installing the cluster. If you need to
enable IPsec after cluster installation, you must first resize your cluster MTU to account for the
overhead of the IPsec ESP IP header.

For more information on how to configure the IPsec encryption, see Configuring IPsec encryption of the
Networking guide in OpenShift Container Platform documentation.

Additional resources

    Configuring IPsec encryption in the OpenShift Container Platform Networking guide

5.9. CENTRALIZING TLS SECURITY PROFILES FOR OPENSHIFT DATA
FOUNDATION

OpenShift Data Foundation version 4.22 introduces a new CRD TLSProfile which allows configuring TLS
servers protocol version, cipher suites, and groups across OpenShift Data Foundation in a centralized

18

                                                           CHAPTER 5. SECURITY CONSIDERATIONS

fashion. This update includes support for Post-Quantum hybrid key exchange mechanisms. The CRD is
installed as a separate package ocs-tls-profiles as OLM dependency in odf-dependencies package.

                  IMPORTANT

                  If the cluster is in FIPS mode, avoid setting a TLSProfile entirely. Forcing a specific profile
                  can conflict with FIPS restrictions.

                  IMPORTANT

                  For disconnected environments, you must include the ocs-tls-profiles package in your
                  mirroring list.

TLSProfile is a namespace-scoped resource:

     spec:
     rules:
     - selectors: [<selector>, ...]
     config:
         version: <TLSv1.2 | TLSv1.3>
         ciphers: [<cipher>, ...]
         groups: [<group>, ...]

A TLSProfile contains an ordered list of rules. Each rule maps one or more selectors to a TLS
configuration (version, ciphers, groups). When a component looks up its configuration, the most specific
matching selector wins.

A selector is a structured string formatted like domain/server or just domain. They are matched in the
following order (from most specific to least specific):

 Form                                                  Matches

 example.io/s3                                         server s3 under example.io

 example.io                                            all servers under example.io

 *.example.io/s3                                       server s3 under any subdomain of example.io

 *.example.io                                          all servers under any subdomain of example.io

 */s3                                                  server s3 under any domain

 *                                                     everything


Version must be exact, not a range. This ensures that what you configure is exactly what runs. API-level
enums and CEL rules automatically reject any incompatible cipher/group combinations for the chosen
version.

TLS constraints enforced at the API level:

     TLS 1.2: only TLS 1.2 ECDHE ciphers and classical groups.

                                                           19

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

 TLS 1.3: only TLS 1.3 ciphers; any group (classical or hybrid post-quantum).

          IMPORTANT

          If your cluster is FIPS-enabled, be aware that the following options are not FIPS 140-2
          approved: * Groups: Hybrid post-quantum groups (X25519MLKEM768,
          SecP256r1MLKEM768, SecP384r1MLKEM1024). * Ciphers: All ChaCha20-Poly1305
          variants. * It is recommended to avoid using this API in FIPS-enabled applications, as FIPS
          already automatically selects vetted cryptography.

Selectors in OpenShift Data Foundation 4.22* Only noobaa.io, rook.io, and * are supported. * \* is a
catch all for all services that support centralized configuration (MCG and RGW). * The CR name must
be ocs-tls-profile and must be in the operator deployed namespace.

Supported TLS configurations

The ciphers are sourced from Mozilla SSL Guidelines under intermediate and modern profiles, with
additional TLS Curves supporting Post Quantum Cryptography (PQC).

 Version   Ciphers                                                Groups

 TLSv1.2   *                                                      * secp256r1 * secp384r1 -
           TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA25                 secp521r1 * X25519
           6 *
           TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA38
           4 *
           TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305
           _SHA256 *
           TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 *
           TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 *
           TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_S
           HA256

 TLSv1.3   * TLS_AES_128_GCM_SHA256 *                             * secp256r1 * secp384r1 *
           TLS_AES_256_GCM_SHA384 *                               secp521r1 * X25519 *
           TLS_CHACHA20_POLY1305_SHA256                           X25519MLKEM768 *
                                                                  SecP256r1MLKEM768 *
                                                                  SecP384r1MLKEM1024


To view the most current list of supported ciphers and groups directly from your cluster, run the
following command:

 $ oc get crd tlsprofiles.ocs.openshift.io -o yaml | grep enum -A 10

Adjust the -A value as needed to display more results.

Example TLSProfile Custom Resource:

 apiVersion: ocs.openshift.io/v1
 kind: TLSProfile
 metadata:
 name: ocs-tls-profile # MUST MATCH THE NAME


20

         CHAPTER 5. SECURITY CONSIDERATIONS

namespace: openshift-storage
spec:
rules:
 # TLS 1.3 with post-quantum
 - selectors:
  - "*"
 config:
  version: TLSv1.3
  ciphers:
     - TLS_AES_128_GCM_SHA256
     - TLS_AES_256_GCM_SHA384
  groups:
     - X25519MLKEM768
     - secp256r1

5.10. NETWORK ACCESS CONSIDERATIONS FOR MULTICLOUD
OBJECT GATEWAY IN CLOUD ENVIRONMENTS

Deployment of Multicloud Object Gateway in cloud environments uses the default network
configuration provided by the cloud platform. These default configurations are typically publicly
accessible but require authentication. Network access to the OpenShift cluster network and applications
can be restricted by applying the access control mechanisms described in the following cloud provider
documentation:

  Azure - Storage network security

  Google - IP filtering overview

  IBM Cloud - Cloud object storage setting a firewall

5.11. MULTICLOUD OBJECT GATEWAY PUBLIC ROUTES

By default, Multicloud Object Gateway (MCG) automatically creates public routes (S3, STS, and
noobaa-mgmt) to enable external access from public networks. This default behavior helps with testing
and initial evaluation of the service by allowing immediate access without additional configuration.

These public routes require authentication but are accessible from outside the OpenShift cluster
network. For production deployments or environments with stricter security requirements, you can
disable these routes after deployment.

For information on how to disable external access routes, see the troubleshooting guide.

Additional resources

  Disabling Multicloud Object Gateway external service after deploying OpenShift Data
  Foundation










     21

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        CHAPTER 6. SUBSCRIPTIONS

  6.1. SUBSCRIPTION OFFERINGS

  Red Hat OpenShift Data Foundation subscription is based on “core-pairs,” similar to Red Hat OpenShift
  Container Platform. The Red Hat OpenShift Data Foundation 2-core subscription is based on the
  number of logical cores on the CPUs in the system where OpenShift Container Platform runs.

  As with OpenShift Container Platform:

        OpenShift Data Foundation subscriptions are stackable to cover larger hosts.

        Cores can be distributed across as many virtual machines (VMs) as needed. For example, ten 2-
        core subscriptions will provide 20 cores and in case of IBM Power a 2-core subscription at SMT
        level of 8 will provide 2 cores or 16 vCPUs that can be used across any number of VMs.

        OpenShift Data Foundation subscriptions are available with Premium or Standard support.

  6.2. DISASTER RECOVERY SUBSCRIPTION REQUIREMENT

  Disaster Recovery features supported by Red Hat OpenShift Data Foundation require all of the
  following prerequisites to successfully implement a disaster recovery solution:

        A valid Red Hat OpenShift Data Foundation Advanced entitlement

        A valid Red Hat Advanced Cluster Management for Kubernetes subscription

  Any Red Hat OpenShift Data Foundation Cluster containing PVs participating in active replication either
  as a source or destination requires OpenShift Data Foundation Advanced entitlement. This subscription
  should be active on both source and destination clusters.

  To know how subscriptions for OpenShift Data Foundation work, see knowledgebase article on
  OpenShift Data Foundation subscriptions.

        IMPORTANT

        OpenShift Data Foundation deployed with Multus networking is not supported for
        Regional Disaster Recovery (Regional-DR) setups.

  6.3. CORES VERSUS VCPUS AND HYPERTHREADING

  Making a determination about whether or not a particular system consumes one or more cores is
  currently dependent on whether or not that system has hyperthreading available. Hyperthreading is only
  a feature of Intel CPUs.

  For more information, see the Self-managed Red Hat OpenShift subscription guide .

  6.4. SPLITTING CORES

  Systems that require an odd number of cores need to consume a full 2-core subscription. For example, a
  system that is calculated to require only 1 core will end up consuming a full 2-core subscription once it is
  registered and subscribed.


    When a single virtual machine (VM) with 2 vCPUs uses hyperthreading resulting in 1 calculated vCPU, a
    22

        CHAPTER 6. SUBSCRIPTIONS

When a single virtual machine (VM) with 2 vCPUs uses hyperthreading resulting in 1 calculated vCPU, a
full 2-core subscription is required; a single 2-core subscription may not be split across two VMs with 2
vCPUs using hyperthreading. See section Cores versus vCPUs and hyperthreading for more information.

It is recommended that virtual instances be sized so that they require an even number of cores.

6.4.1. Shared Processor Pools for IBM Power

IBM Power have a notion of shared processor pools. The processors in a shared processor pool can be
shared across the nodes in the cluster. The aggregate compute capacity required for a Red Hat
OpenShift Data Foundation should be a multiple of core-pairs.

6.5. SUBSCRIPTION REQUIREMENTS

Red Hat OpenShift Data Foundation components can run on either OpenShift Container Platform
worker or infrastructure nodes, for which you can use either Red Hat CoreOS (RHCOS) or Red Hat
Enterprise Linux (RHEL) 8.4 as the host operating system. RHEL 7 is now deprecated. OpenShift Data
Foundation subscriptions are required for every OpenShift Container Platform subscribed core with a
ratio of 1:1.

When using infrastructure nodes, the rule to subscribe all OpenShift worker node cores for OpenShift
Data Foundation applies even though they don’t need any OpenShift Container Platform or any
OpenShift Data Foundation subscriptions. You can use labels to state whether a node is a worker or an
infrastructure node.

For more information, see How to use dedicated worker nodes for Red Hat OpenShift Data Foundation
in the Managing and Allocating Storage Resources guide.










    23

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

     CHAPTER 7. INFRASTRUCTURE REQUIREMENTS

7.1. PLATFORM REQUIREMENTS

Red Hat OpenShift Data Foundation 4.22 is supported only on OpenShift Container Platform version
4.22 and its next minor versions.

Bug fixes for previous version of Red Hat OpenShift Data Foundation will be released as bug fix
versions. For more details, see the Red Hat OpenShift Container Platform Life Cycle Policy .

For external cluster subscription requirements, see the Red Hat Knowledgebase article OpenShift Data
Foundation Subscription Guide.

For a complete list of supported platform versions, see the Red Hat OpenShift Data Foundation
Supportability and Interoperability Checker.

7.1.1. Amazon EC2

Supports internal Red Hat OpenShift Data Foundation clusters only.

An Internal cluster must meet both, storage device requirements and have a storage class that provides,
EBS storage via the aws-ebs provisioner.

OpenShift Data Foundation supports gp2-csi and gp3-csi drivers that were introduced by Amazon Web
Services (AWS). These drivers offer better storage expansion capabilities and a reduced monthly price
point (gp3-csi). You can now select the new drivers when selecting your storage class. In case a high
throughput is required, gp3-csi is recommended to be used when deploying OpenShift Data
Foundation.

If you need a high input/output operation per second (IOPS), the recommended EC2 instance types are
D2 or D3.

7.1.2. Bare Metal

Supports internal clusters and consuming external clusters.

An internal cluster must meet both the storage device requirements and have a storage class that
provide local SSD (NVMe/SATA/SAS, SAN) via the Local Storage Operator. For production clusters, a
minimum storage capacity of 0.5 TiB per storage device is required.

7.1.3. VMware vSphere

Supports internal clusters and consuming external clusters.

Recommended versions:

 vSphere 8.0 or later

 VMware Cloud Foundation 9 and VMware vSphere Foundation 9

For more details, see the VMware vSphere infrastructure requirements.




 NOTE
24

      CHAPTER 7. INFRASTRUCTURE REQUIREMENTS

  NOTE

  If VMware ESXi does not recognize its devices as flash, mark them as flash devices.
  Before Red Hat OpenShift Data Foundation deployment, refer to Mark Storage Devices
  as Flash.

Additionally, an Internal cluster must meet both the, storage device requirements and have a storage
class providing either,

 vSAN or VMFS datastore via the vsphere-volume provisioner

 VMDK, RDM, or DirectPath storage devices via the Local Storage Operator.

7.1.4. Microsoft Azure

Supports internal Red Hat OpenShift Data Foundation clusters only.

An internal cluster must meet both, storage device requirements and have a storage class that provides,
an azure disk via the azure-disk provisioner.

Red Hat OpenShift Data Foundation supports StorageClass with performance plus enabled. By enabling
performance plus, the input/output operations per second (IOPS) and throughput limits can be
increased for the following types of disks that are 513 GiB and larger:

 Azure Premium solid-state drives (SSD)

 Standard SSDs

  IMPORTANT

  In order to enable performance plus, you must create a StorageClass with
  enablePerformancePlus=True, and select this StorageClass when deploying OpenShift
  Data Foundation.

7.1.5. Google Cloud

Supports internal Red Hat OpenShift Data Foundation clusters only.

An internal cluster must meet both, storage device requirements and have a storage class that provides,
a GCE Persistent Disk via the gce-pd provisioner.

7.1.6. Red Hat OpenStack Platform [Technology Preview]

Supports internal Red Hat OpenShift Data Foundation clusters and consuming external clusters.

An internal cluster must meet both, storage device requirements and have a storage class that provides
a standard disk via the Cinder provisioner.

7.1.7. IBM Power

Supports internal Red Hat OpenShift Data Foundation clusters and consuming external clusters.

An Internal cluster must meet both, storage device requirements and have a storage class providing
local SSD (NVMe/SATA/SAS, SAN) via the Local Storage Operator.


  25

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

7.1.8. IBM Z and IBM® LinuxONE

Supports internal Red Hat OpenShift Data Foundation clusters. Also, supports external mode where Red
Hat Ceph Storage is running on x86.

An Internal cluster must meet both, storage device requirements and have a storage class providing
local SSD (NVMe/SATA/SAS, SAN) via the Local Storage Operator.

7.1.9. ROSA with hosted control planes (HCP)

Supports internal Red Hat OpenShift Data Foundation clusters only.

An internal cluster must meet both, storage device requirements and have a storage class that provides
AWS EBS volumes via gp3-csi provisioner.

7.1.10. Any platform

Supports internal clusters and consuming external clusters.

An internal cluster must meet both the storage device requirements and have a storage class that
provide local SSD (NVMe/SATA/SAS, SAN) via the Local Storage Operator.

7.2. EXTERNAL MODE REQUIREMENT

7.2.1. Red Hat Ceph Storage

To check the supportability and interoperability of Red Hat Ceph Storage (RHCS) with Red Hat
OpenShift Data Foundation in external mode, go to the lab Red Hat OpenShift Data Foundation
Supportability and Interoperability Checker.

1.   Select Service Type as ODF as Self-Managed Service.

2.   Select appropriate Version from the drop down.

3.   On the Versions tab, click the Supported RHCS Compatibility tab.

For instructions regarding how to install a RHCS cluster, see the installation guide.

7.2.2. IBM FlashSystem

To use IBM FlashSystem as a pluggable external storage on other providers, you need to first deploy it
before you can deploy OpenShift Data Foundation, which would use the IBM FlashSystem storage class
as a backing storage.

For the latest supported FlashSystem storage systems and versions, see IBM ODF FlashSystem driver
documentation.

For instructions on how to deploy OpenShift Data Foundation, see Creating an OpenShift Data
Foundation Cluster for external IBM FlashSystem storage.

7.3. RESOURCE REQUIREMENTS

Red Hat OpenShift Data Foundation services consist of an initial set of base services, and can be
extended with additional device sets. All of these Red Hat OpenShift Data Foundation services pods are

26

        CHAPTER 7. INFRASTRUCTURE REQUIREMENTS

scheduled by kubernetes on OpenShift Container Platform nodes. Expanding the cluster in multiples of
three, one node in each failure domain, is an easy way to satisfy the pod placement rules.

        IMPORTANT

        These requirements relate to OpenShift Data Foundation services only, and not to any
        other services, operators or workloads that are running on these nodes.

Aggregate available resource requirements for Red Hat OpenShift Data Foundation only


    Deployment Mode    Base services        Additional device Set

    Internal           30 CPU (logical)      6 CPU (logical)

                       72 GiB memory         15 GiB memory

                       3 storage devices     3 storage devices


    External           4 CPU (logical)  Not applicable

                       16 GiB memory



    Example: For a 3 node cluster in an internal mode deployment with a single device set, a minimum of 3 x
    10 = 30 units of CPU are required.

    For more information, see Subscriptions and CPU units.

    For additional guidance with designing your Red Hat OpenShift Data Foundation cluster, see the ODF
    Sizing Tool.

    CPU units In this section, 1 CPU Unit maps to the Kubernetes concept of 1 CPU unit.

     1 unit of CPU is equivalent to 1 core for non-hyperthreaded CPUs.

     2 units of CPU are equivalent to 1 core for hyperthreaded CPUs.

     Red Hat OpenShift Data Foundation core-based subscriptions always come in pairs (2 cores).

    Aggregate minimum resource requirements for IBM Power


     Deployment Mode    Base services

    Internal            48 CPU (logical)

                        192 GiB memory

                        3 storage devices, each with additional
                        500GB of disk





                                             27

    Red Hat OpenShift Data Foundation 4.22 Planning your deployment


        Deployment Mode Base services

External 24 CPU (logical)

        48 GiB memory



Example: For a 3 node cluster in an internal-attached devices mode deployment, a minimum of 3 x 16 =
48 units of CPU and 3 x 64 = 192 GB of memory is required.

        NOTE

        Internal OpenShift Data Foundation component resource allocations can be adjusted
        after deployment if requirements change. Adjustments might be needed after scaling
        storage or when the cluster is under sustained high workload conditions. This flexibility
        allows you to adapt resource usage as operational demands evolve. For more
        information, see the knowledgebase article How to Safely Increase Resources on
        OpenShift Data Foundation Storage Nodes.

7.3.1. Resource requirements for IBM Z and IBM LinuxONE infrastructure

Red Hat OpenShift Data Foundation services consist of an initial set of base services, and can be
extended with additional device sets.

All of these Red Hat OpenShift Data Foundation services pods are scheduled by kubernetes on
OpenShift Container Platform nodes . Expanding the cluster in multiples of three, one node in each
failure domain, is an easy way to satisfy the pod placement rules.

Aggregate available resource requirements for Red Hat OpenShift Data Foundation only (IBM Z
and IBM® LinuxONE)


    Deployment Mode  Base services  Additional device Set  IBM Z and IBM®
                                                          LinuxONE minimum
                                                       hardware requirements

    Internal         21 CPU         3 CPU (logical)     1-4 IFL
                     (logical)
                      3 nodes       15 GiB memory
                      with 7        3 storage
                      CPUs          devices
                      (logical)
                      each

                     72 GiB memory

                     3 storage
                     devices










    28

CHAPTER 7. INFRASTRUCTURE REQUIREMENTS


Deployment Mode  Base services   Additional device Set  IBM Z and IBM®
                                                       LinuxONE minimum
                                                    hardware requirements

External         4 CPU (logical) Not applicable        Not applicable

                 16 GiB memory



CPU
 Is the number of virtual cores defined in the hypervisor, IBM Z/VM, Kernel Virtual Machine (KVM), or
 both.
IFL (Integrated Facility for Linux)
 Is the physical core for IBM Z and IBM® LinuxONE.

Minimum system environment

   To operate a minimal cluster, at least one additional IFL is required on top of the entire
   deployment. Depending on the workload, up to 4 IFLs might be necessary.

7.3.2. Minimum deployment resource requirements

An OpenShift Data Foundation cluster will be deployed with minimum configuration when the standard
deployment resource requirement is not met.

Important considerations for minimum deployment

   For production clusters, deploying OpenShift Data Foundation with more than three devices is
   strongly recommended to improve data resiliency.

   DR installation should not use the minimal deployment configuration, as it can result in
   installation failure.

   The minimum deployment requirements relate to OpenShift Data Foundation services only, and
   not to any other services, operators or workloads that are running on these nodes.

Aggregate resource requirements for OpenShift Data Foundation only

       Deployment Mode                             Base services

Internal                                           24 CPU (logical)
                                                   15 CPUs for IBM Z and IBM LinuxONE

                                                   72 GiB memory

                                                   3 storage devices



If you want to add additional device sets, we recommend converting your minimum deployment to
standard deployment.



                                                   29

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

7.3.3. Compact deployment resource requirements

Red Hat OpenShift Data Foundation can be installed on a three-node OpenShift compact bare metal
cluster, where all the workloads run on three strong master nodes. There are no worker or storage
nodes.

            IMPORTANT

            These requirements relate to OpenShift Data Foundation services only, and not to any
            other services, operators or workloads that are running on these nodes.

Table 7.1. Aggregate resource requirements for OpenShift Data Foundation only

    Deployment Mode    Base services        Additional device Set

Internal               24 CPU (logical)                        6 CPU (logical)

                       72 GiB memory                           15 GiB memory

                       3 storage devices                       3 storage devices



To configure OpenShift Container Platform on a compact bare metal cluster, see Configuring a three-
node cluster and Delivering a Three-node Architecture for Edge Deployments .

7.3.4. Resource requirements for MCG only deployment

An OpenShift Data Foundation cluster deployed only with the Multicloud Object Gateway (MCG)
component provides the flexibility in deployment and helps to reduce the resource consumption.

Table 7.2. Aggregate resource requirements for MCG only deployment

Deployment Mode    Core        Database (DB)                       Endpoint

Internal           1 CPU            1 CPU                                   1 CPU

                   4 GiB memory     4 GiB memory                            2 GiB memory

                                                                                 NOTE

                                                                                 The
                                                                                 default
                                                                                 autoscal
                                                                                 e is
                                                                                 betwee
                                                                                 n 1 - 2.



7.3.5. Resource requirements for using Network File System

You can create exports using Network File System (NFS) that can then be accessed externally from the
OpenShift cluster. If you plan to use this feature, the NFS service consumes 6 CPUs and 16Gi of RAM.
NFS is optional and is disabled by default.

30

                                            CHAPTER 7. INFRASTRUCTURE REQUIREMENTS

    The NFS volume can be accessed two ways:

             In-cluster: by an application pod inside of the OpenShift cluster.

             Out of cluster: from outside of the OpenShift cluster.

    For more information about the NFS feature, see Creating exports using NFS

    7.3.6. Resource requirements for performance profiles

    OpenShift Data Foundation provides three performance profiles to enhance the performance of the
    clusters. You can choose one of these profiles based on your available resources and desired
    performance level during deployment or post deployment.

    Table 7.3. Recommended resource requirement for different performance profiles

     Performance profile        CPU         CPU on IBM Z or        Memory
                                                LinuxONE

     Lean                   24              15              72 GiB

     Balanced               30              21              72 GiB

     Performance            45              30              96 GiB


        IMPORTANT

        Make sure to select the profiles based on the available free resources as you might
        already be running other workloads.

7.4. POD PLACEMENT RULES

Kubernetes is responsible for pod placement based on declarative placement rules. The Red Hat
OpenShift Data Foundation base service placement rules for Internal cluster can be summarized as
follows:

        Nodes are labeled with the cluster.ocs.openshift.io/openshift-storage key

        Nodes are sorted into pseudo failure domains if none exist

        Components requiring high availability are spread across failure domains

        A storage device must be accessible in each failure domain

This leads to the requirement that there be at least three nodes, and that nodes be in three distinct rack
or zone failure domains in the case of pre-existing topology labels.

For additional device sets, there must be a storage device, and sufficient resources for the pod
consuming it, in each of the three failure domains. Manual placement rules can be used to override
default placement rules, but generally this approach is only suitable for bare metal deployments.




             NOTE
                                            31

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        NOTE

        Ensure that CPU and RAM resources remain available across all failure domains, even in
        the event of a failure. By default, the allocation is 6 CPUs and 15 GB of RAM. Any changes
        to this default configuration or the addition of backing stores must be considered in
        resource planning.

  7.5. STORAGE DEVICE REQUIREMENTS

  Use this section to understand the different storage capacity requirements that you can consider when
  planning internal mode deployments and upgrades. We generally recommend 12 devices or less per
  node. This recommendation ensures both that nodes stay below cloud provider dynamic storage device
  attachment limits, and to limit the recovery time after node failures with local storage devices.
  Expanding the cluster in multiples of three, one node in each failure domain, is an easy way to satisfy pod
  placement rules.

  Storage nodes should have at least two disks, one for the operating system and the remaining disks for
  OpenShift Data Foundation components.

        NOTE

        You can expand the storage capacity only in the increment of the capacity selected at
        the time of installation.

  7.5.1. Dynamic storage devices

  Red Hat OpenShift Data Foundation permits the selection of either 0.5 TiB, 1 TiB, 2 TiB, 4 TiB, or 8 TiB
  capacities as the request size for dynamic storage device sizes. The number of dynamic storage devices
  that can run per node is a function of the node size, underlying provisioner limits and resource
  requirements.

  7.5.2. Local storage devices

  For local storage deployment, any disk size of 16 TiB or less can be used, and all disks should be of the
  same size and type. The number of local storage devices that can run per node is a function of the node
  size and resource requirements. Expanding the cluster in multiples of three, one node in each failure
  domain, is an easy way to satisfy pod placement rules.

        NOTE

        Disk partitioning is not supported, except for DASD-based deployment on IBM Z.

  7.5.3. Capacity planning

  Always ensure that available storage capacity stays ahead of consumption. Recovery is difficult if
  available storage capacity is completely exhausted, and requires more intervention than simply adding
  capacity or deleting or migrating content.

  Capacity alerts are issued when cluster storage capacity reaches 75% (near-full) and 85% (full) of total
  capacity. Always address capacity warnings promptly, and review your storage regularly to ensure that
  you do not run out of storage space. When you get to 75% (near-full), either free up space or expand
  the cluster. When you get the 85% (full) alert, it indicates that you have run out of storage space
  completely and cannot free up space using standard commands. At this point, contact Red Hat
  Customer Support.

32

                                                          CHAPTER 7. INFRASTRUCTURE REQUIREMENTS

    The following tables show example node configurations for Red Hat OpenShift Data Foundation with
    dynamic storage devices.

    Table 7.4. Example initial configurations with 3 nodes

    Storage Device size Storage Devices per               Total capacity  Usable storage
                                node                                         capacity

    0.5 TiB                      1                         1.5 TiB           0.5 TiB

     2 TiB                       1                         6 TiB              2 TiB

     4 TiB                       1                         12 TiB             4 TiB


Table 7.5. Example of expanded configurations with 30 nodes (N)

  Storage Device size (D) Storage Devices per Total capacity (D * M * Usable storage
        node (M) N) capacity (D*M*N/3)

        0.5 TiB 3 45 TiB 15 TiB

        2 TiB 6 360 TiB 120 TiB

        4 TiB 9 1080 TiB 360 TiB


    7.6. ERASURE CODING CONSIDERATIONS

    Erasure coding (EC) is a data protection method that provides cost-effective storage resiliency by
    dividing data into fragments, expanding and encoding them with redundant data pieces, and storing
    them across different locations. This approach offers better storage efficiency compared to traditional
    replication methods, helping to reduce storage costs while maintaining data durability and availability.

    Erasure coding can be configured in the following scenarios:

     Initial deployment: As an advanced option during Red Hat OpenShift Data Foundation cluster
     creation through the UI

     Existing clusters: When creating a new storage class with a new pool through the storage class
     wizard (applicable for block and file storage only)

      NOTE

      Erasure coding for RADOS Gateway (RGW) object storage can only be configured during
      initial deployment and is not supported for existing clusters.

    7.6.1. Supported erasure coding profiles

    The following erasure coding profiles are supported: 2+2, 4+2, 5+2, 6+2, 8+2, 8+3, 8+4, 8+6.

    These profiles are supported for both fresh deployments and post-deployment storage pool creation.

      IMPORTANT
                                                               33

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        IMPORTANT

        For all erasure coding profiles, it is recommended to have at least one additional node
        beyond the minimum required number of chunks. This provides better data protection
        during node maintenance operations, such as node restart, drain, or replacement.

        NOTE

        Only the erasure coding profiles that can be supported by the current number of available
        nodes are displayed to the user.

  7.6.2. Node requirements

  The number of nodes required for each erasure coding profile is calculated as data chunks (k) + parity
  chunks (m). The recommended number of nodes is k + m + 1 to ensure optimal performance and
  availability. While the minimum supported configuration is k + m nodes, deploying with fewer than k + m +
  1 nodes will result in a warning.

  For more information about erasure coding in Ceph, see the Ceph Erasure Code documentation .

  7.6.3. Limitations

  Before implementing erasure coding, consider the following limitations:

        Performance: Erasure coding pools have lower performance compared to replicated pools.

        IPv6 networking: Not supported on clusters using IPv6.

        PV-level encryption: dmcrypt/LUKS encryption is not supported on erasure-coded RBD
        volumes.

        Migration: Migration from replicated pools to erasure-coded pools, and from erasure-coded
        pools to replicated pools, is supported through VolumeSnapshot-based restore.

        Failure domains: Cannot be used when Rack or Zone failure domains are labeled on nodes,
        including stretched clusters and arbiter configurations.










    34

        CHAPTER 8. NETWORK REQUIREMENTS

        CHAPTER 8. NETWORK REQUIREMENTS

OpenShift Data Foundation requires that at least one network interface that is used for the cluster
network to be capable of at least 10 gigabit network speeds. This section further covers different
network considerations for planning deployments.

8.1. IPV6 SUPPORT

Red Hat OpenShift Data Foundation version 4.12 introduced the support of IPv6. IPv6 is supported in
single stack only, and cannot be used simultaneously with IPv4. IPv6 is the default behavior in OpenShift
Data Foundation when IPv6 is turned on in Openshift Container Platform.

Red Hat OpenShift Data Foundation version 4.14 introduces IPv6 auto detection and configuration.
Clusters using IPv6 will automatically be configured accordingly.

OpenShift Container Platform dual stack with Red Hat OpenShift Data Foundation IPv4 is supported
from version 4.13 and later. Dual stack on Red Hat OpenShift Data Foundation IPv6 is not supported.

8.2. MULTI NETWORK PLUG-IN (MULTUS) SUPPORT

OpenShift Data Foundation supports the ability to use multi-network plug-in Multus on bare metal
infrastructures to improve security and performance by isolating the different types of network traffic.
By using Multus, one or more network interfaces on hosts can be reserved for exclusive use of
OpenShift Data Foundation. For more details, see Multus architecture for OpenShift Data Foundation .










    35

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        CHAPTER 9. DISASTER RECOVERY

  Disaster Recovery (DR) helps an organization to recover and resume business critical functions or
  normal operations when there are disruptions or disasters. OpenShift Data Foundation provides High
  Availability (HA) & DR solutions for stateful apps which are broadly categorized into two broad
  categories:

        Metro-DR: Single Region and cross data center protection with no data loss.

        Regional-DR: Cross Region protection with minimal potential data loss.

        Disaster Recovery with stretch cluster: Single OpenShift Data Foundation cluster is stretched
        between two different locations to provide the storage infrastructure with disaster recovery
        capabilities.










    36

                        CHAPTER 10. METRO-DR

                        CHAPTER 10. METRO-DR

Metropolitan disaster recovery (Metro-DR) is composed of Red Hat Advanced Cluster Management for
Kubernetes (RHACM), Red Hat Ceph Storage and OpenShift Data Foundation components to provide
application and data mobility across OpenShift Container Platform clusters.

This release of Metro-DR solution provides volume persistent data and metadata replication across sites
that are geographically dispersed. In the public cloud these would be similar to protecting from an
Availability Zone failure. Metro-DR ensures business continuity during the unavailability of a data center
with no data loss. This solution is entitled with Red Hat Advanced Cluster Management (RHACM) and
OpenShift Data Foundation Advanced SKUs and related bundles.

             IMPORTANT

             You can now easily set up Metropolitan disaster recovery solutions for workloads based
             on OpenShift virtualization technology using OpenShift Data Foundation. For more
             information, see the knowledgebase article on Metropolitan disaster recovery with
             OpenShift Virtualization.

Prerequisites

 Disaster Recovery features supported by Red Hat OpenShift Data Foundation require all of the
 following prerequisites in order to successfully implement a Disaster Recovery solution:

  A valid Red Hat OpenShift Data Foundation Advanced entitlement

  A valid Red Hat Advanced Cluster Management for Kubernetes subscription

 To know how subscriptions for OpenShift Data Foundation work, see the knowledgebase article
 on OpenShift Data Foundation subscriptions.

 Ensure that the primary managed cluster (Site-1) is co-situated with the active RHACM hub
 cluster while the passive hub cluster is situated along with the secondary managed cluster (Site-
 2). Alternatively, the active RHACM hub cluster can be placed in a neutral site (site-3) that is not
 impacted by the failures of either of the primary managed cluster at Site-1 or the secondary
 cluster at Site-2. In this situation, if a passive hub cluster is used it can be placed with the
 secondary cluster at Site-2.

                    NOTE

                    Hub recovery for Metro-DR is a Technology Preview feature and is subject to
                    Technology Preview support limitations.

For detailed solution requirements, see the Metro-DR requirements, the deployment requirements for
Red Hat Ceph Storage stretch cluster with arbiter, and the RHACM requirements documentation.

Additional resources

 Knowledgebase article on Metropolitan disaster recovery with OpenShift Virtualization

 Knowledgebase article on OpenShift Data Foundation subscriptions

 Metro-DR requirements

 Deployment requirements for Red Hat Ceph Storage stretch cluster with arbiter

                        37

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

RHACM requirements










38

                 CHAPTER 11. REGIONAL-DR

                 CHAPTER 11. REGIONAL-DR

Regional disaster recovery (Regional-DR) is composed of Red Hat Advanced Cluster Management for
Kubernetes (RHACM) and OpenShift Data Foundation components to provide application and data
mobility across OpenShift Container Platform clusters. It is built on Asynchronous data replication and
hence could have a potential data loss but provides the protection against a broad set of failures.

Red Hat OpenShift Data Foundation is backed by Ceph as the storage provider, whose lifecycle is
managed by Rook and it’s enhanced with the ability to:

 Enable pools for mirroring.

 Automatically mirror images across RBD pools.

 Provides csi-addons to manage per Persistent Volume Claim mirroring.

This release of Regional-DR supports Multi-Cluster configuration that is deployed across different
regions and data centers. For example, a 2-way replication across two managed clusters located in two
different regions or data centers. This solution is entitled with Red Hat Advanced Cluster Management
(RHACM) and OpenShift Data Foundation Advanced SKUs and related bundles.

             IMPORTANT

             You can now easily set up Regional disaster recovery solutions for workloads based on
             OpenShift virtualization technology using OpenShift Data Foundation. For more
             information, see the knowledgebase article on Regional disaster recovery with OpenShift
             Virtualization.

Prerequisites

 Disaster Recovery features supported by Red Hat OpenShift Data Foundation require all of the
 following prerequisites in order to successfully implement a Disaster Recovery solution:

  A valid Red Hat OpenShift Data Foundation Advanced entitlement

  A valid Red Hat Advanced Cluster Management for Kubernetes subscription

 To know how subscriptions for OpenShift Data Foundation work, see the knowledgebase article
 on OpenShift Data Foundation subscriptions.

 Ensure that the primary managed cluster (Site-1) is co-situated with the active RHACM hub
 cluster while the passive hub cluster is situated along with the secondary managed cluster (Site-
 2). Alternatively, the active RHACM hub cluster can be placed in a neutral site (site-3) that is not
 impacted by the failures of either of the primary managed cluster at Site-1 or the secondary
 cluster at Site-2. In this situation, if a passive hub cluster is used it can be placed with the
 secondary cluster at Site-2.

For detailed solution requirements, see the Regional-DR requirements and RHACM requirements
documentation.

Additional resources

 Knowledgebase article on Regional disaster recovery with OpenShift Virtualization

 Knowledgebase article on OpenShift Data Foundation subscriptions


             39

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

Regional-DR requirements

RHACM requirements










40

          CHAPTER 12. DISASTER RECOVERY WITH STRETCH CLUSTER

CHAPTER 12. DISASTER RECOVERY WITH STRETCH CLUSTER

In this case, a single cluster is stretched across two zones with a third zone as the location for the arbiter.
This feature is currently intended for deployment in the OpenShift Container Platform on-premises and
in the same location. This solution is not recommended for deployments stretching over multiple data
centers. Instead, consider Metro-DR as a first option for no data loss DR solution deployed over multiple
data centers with low latency networks.

      NOTE

      The stretch cluster solution is designed for deployments where latencies do not exceed 10
      ms maximum round-trip time (RTT) between the zones containing data volumes. For
      Arbiter nodes follow the latency requirements specified for etcd. Contact Red Hat
      Customer Support if you are planning to deploy with higher latencies.

To use the stretch cluster,

 You must have a minimum of five nodes across three zones, where:

  Two nodes per zone are used for each data-center zone, and one additional zone with one
  node is used for arbiter zone (the arbiter can be on a master node).

 All the nodes must be manually labeled with the zone labels prior to cluster creation.
 For example, the zones can be labeled as:

  topology.kubernetes.io/zone=arbiter (master or worker node)

  topology.kubernetes.io/zone=datacenter1 (minimum two worker nodes)

  topology.kubernetes.io/zone=datacenter2 (minimum two worker nodes)

For more information on configuring OpenShift Data Foundation for stretch cluster, see the Configuring
OpenShift Data Foundation disaster recovery for OpenShift workloads guide.

      IMPORTANT

      You can now easily set up disaster recovery with stretch cluster for workloads based on
      OpenShift virtualization technology using OpenShift Data Foundation. For more
      information, see the OpenShift Virtualization in OpenShift Container Platform guide.

Additional resources

 Guidance for Red Hat OpenShift Container Platform Clusters - Deployments Spanning Multiple
 Sites

 Red Hat Customer Support

 Knowledgebase article on OpenShift Data Foundation subscriptions

 Configuring OpenShift Data Foundation for stretch cluster

 OpenShift Virtualization in OpenShift Container Platform





      41

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

     CHAPTER 13. DISCONNECTED ENVIRONMENT

Disconnected environment is a network restricted environment where the Operator Lifecycle Manager
(OLM) cannot access the default Software Catalog and image registries, which require internet
connectivity.

Red Hat supports deployment of OpenShift Data Foundation in disconnected environments where you
have installed OpenShift Container Platform in restricted networks.

To install OpenShift Data Foundation in a disconnected environment, see the Using Operator Lifecycle
Manager in disconnected environments topic in OpenShift Container Platform documentation.

             NOTE

             When you install OpenShift Data Foundation in a restricted network environment, apply a
             custom Network Time Protocol (NTP) configuration to the nodes, because by default,
             internet connectivity is assumed in OpenShift Container Platform and chronyd is
             configured to use the *.rhel.pool.ntp.org servers.

             For more information, see the Red Hat Knowledgebase solution A newly deployed OCS 4
             cluster status shows as "Degraded", Why? and Configuring chrony time service of the
             Installing guide in OpenShift Container Platform documentation.

The Agent-based Installer allows you to use a mirror registry for disconnected installations. For more
information, see Preparing to install with Agent-based Installer .

Packages to include for OpenShift Data Foundation

When you prune the redhat-operator index image, include the following operator bundles for the
OpenShift Data Foundation deployment:

 ocs-operator

 odf-operator

 mcg-operator

 odf-csi-addons-operator

 ocs-client-operator

 odf-prometheus-operator

 recipe

 rook-ceph-operator

 cephcsi-operator

 odf-dependencies

 odf-external-snapshotter-operator

 ocs-tls-profiles



42

           CHAPTER 13. DISCONNECTED ENVIRONMENT

Only for local storage deployments:

 local-storage-operator

Only for Regional Disaster Recovery (Regional-DR) configuration or Metro Disaster Recovery
(Metro-DR) configuration:

 odf-multicluster-orchestrator

 odr-cluster-operator

 odr-hub-operator

 odr-volsync-plugin-operator

       IMPORTANT

       Make sure to name the CatalogSource as redhat-operators.

Upgrade requirements

When mirroring OpenShift Data Foundation for y-stream upgrades in a disconnected environment,
ensure the following:

 Include all the OpenShift Data Foundation operator bundles in the mirror configuration file.

 Package both operator versions, the currently installed version and the target upgrade version
 within the same operator catalog index image.
 For example:

  - name: odf-operator
   channels:
   - name: stable-<target-odf-version>
    minVersion : <target-upgrade-version>-rhodf
    maxVersion : <target-upgrade-version>-rhodf
   - name: stable-<current-odf-version>
    minVersion : <current-version>-rhodf
    maxVersion : <current-version>-rhodf
  - name: ocs-operator
   channels:
  [...]










       43

Red Hat OpenShift Data Foundation 4.22 Planning your deployment

        CHAPTER 14. PERFORMANCE CONSIDERATIONS AND
        BENCHMARKING

  To accurately predict performance and identify potential bottlenecks in your OpenShift Data
  Foundation cluster, benchmarking is necessary. This knowledgebase article explains the tools and
  procedures for performing these benchmarks and analyzing the resulting data.










    44

    CHAPTER 15. SUPPORTED AND UNSUPPORTED FEATURES FOR IBM POWER AND IBM Z

                                   CHAPTER 15. SUPPORTED AND UNSUPPORTED FEATURES
                                              FOR IBM POWER AND IBM Z

Table 15.1. List of supported and unsupported features on IBM Power and IBM Z

    Features                           IBM Power        IBM Z

Compact deployment                 Unsupported      Unsupported

Dynamic storage devices            Unsupported      Supported

Federal Information Processing     Unsupported      Unsupported
Standard Publication (FIPS)

Alerts to control overprovision    Supported        Unsupported

Alerts when Ceph Monitor runs      Supported        Unsupported
out of space

Extended OpenShift Data            Unsupported      Unsupported
Foundation control plane, which
allows pluggable external storage
such as IBM FlashSystem

IPv6 support                       Unsupported      Unsupported

Multus                             Unsupported      Supported

Minimum deployment                 Unsupported      Unsupported

Agnostic deployment of             Unsupported      Unsupported
OpenShift Data Foundation on
any Openshift supported platform

Installer-provisioned deployment   Unsupported      Unsupported
of OpenShift Data Foundation
using bare metal infrastructure

OpenShift dual stack with          Unsupported      Unsupported
OpenShift Data Foundation using
IPv4

Ability to disable Multicloud      Unsupported      Unsupported
Object Gateway external service
during deployment

RHACM KubeVirt DR integration      Unsupported      Unsupported
(Virtualization required)




                                                    45

    Red Hat OpenShift Data Foundation 4.22 Planning your deployment


        Features IBM Power IBM Z

Openshift Lightspeed (OLS) via Unsupported Unsupported
BYOK



    Two Nodes Fencing (TNF)   Unsupported     Unsupported

    NVMe-oF Gateway           Unsupported     Unsupported










    46

     CHAPTER 16. NEXT STEPS

     CHAPTER 16. NEXT STEPS

To start deploying your OpenShift Data Foundation, you can use the internal mode within OpenShift
Container Platform or use external mode to make available services from a cluster running outside of
OpenShift Container Platform.

Depending on your requirement, go to the respective deployment guides.

Internal mode * Deploying OpenShift Data Foundation using Amazon web services * Deploying
OpenShift Data Foundation using Bare Metal * Deploying OpenShift Data Foundation using VMWare
vSphere * Deploying OpenShift Data Foundation using Microsoft Azure

 Deploying OpenShift Data Foundation using Google Cloud

 Deploying OpenShift Data Foundation using Red Hat OpenStack Platform [Technology
 Preview]

 Deploying OpenShift Data Foundation on IBM Power

 Deploying OpenShift Data Foundation on IBM Z

 Deploying OpenShift Data Foundation on any platform

External mode * Deploying OpenShift Data Foundation in external mode

Internal or external

For deploying multiple clusters, see Deploying multiple OpenShift Data Foundation clusters .










 47
