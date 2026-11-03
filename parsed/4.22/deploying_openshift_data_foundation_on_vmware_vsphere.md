---
title: "deploying openshift data foundation on vmware vsphere"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/deploying_openshift_data_foundation_on_vmware_vsphere"
source_file: "deploying_openshift_data_foundation_on_vmware_vsphere.pdf"
version: "4.22"
chapter: "deploying_openshift_data_foundation_on_vmware_vsphere"
---

                   Red Hat OpenShift Data Foundation
                                  4.22

                 Deploying OpenShift Data Foundation on
                             VMware vSphere

Instructions on deploying OpenShift Data Foundation using VMware vSphere
                             infrastructure










    Last Updated: 2026-10-23



Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data
Foundation on VMware vSphere

Instructions on deploying OpenShift Data Foundation using VMware vSphere infrastructure

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

Read this document for instructions about how to install Red Hat OpenShift Data Foundation using
Red Hat OpenShift Container Platform on VMware vSphere clusters.

       Table of Contents

       Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE                                                                               3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION        4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE                                                                               5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA FOUNDATION                              6
 1.1. RECOMMENDED CONFIGURATION FOR OPTIMIZED PERFORMANCE                             8
 1.1.1. Recommended configuration guidelines                                          8
 1.2. REQUIREMENTS FOR INSTALLING OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES
                                                                                      9

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES        11
 2.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR                          11
 2.2. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE TOKEN AUTHENTICATION METHOD
                                                                                     12
 2.3. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE KUBERNETES AUTHENTICATION
 METHOD                                                                              13
 2.3.1. Enabling and disabling key rotation when using KMS                           16
  2.3.1.1. Enforcing precedence for key rotation                                     16
  2.3.1.2. Enabling key rotation                                                     17
   2.3.1.2.1. Annotating a namespace for key rotation                                17
   2.3.1.2.2. Annotating a storage class for key rotation                            18
   2.3.1.2.3. Annotating a PersistentVolumeClaim for key rotation                    18
  2.3.1.3. Disabling key rotation                                                    19
 2.4. CREATING OPENSHIFT DATA FOUNDATION CLUSTER                                     20

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. DEPLOY USING LOCAL STORAGE DEVICES                                        26
 3.1. INSTALLING LOCAL STORAGE OPERATOR                                              26
 3.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR                          26
 3.3. CREATING OPENSHIFT DATA FOUNDATION CLUSTER ON VMWARE VSPHERE                   28

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT                            35
 4.1. VERIFYING THE STATE OF THE PODS                                                35
 4.2. VERIFYING THE OPENSHIFT DATA FOUNDATION CLUSTER IS HEALTHY                     37
 4.3. VERIFYING THE MULTICLOUD OBJECT GATEWAY IS HEALTHY                             38
 4.4. VERIFYING THAT THE SPECIFIC STORAGE CLASSES EXIST                              38

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY                               40

CHAPTER 6. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING DYNAMIC STORAGE
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
DEVICES                                                                              41
 6.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR                          41
 6.2. CREATING A STANDALONE MULTICLOUD OBJECT GATEWAY                                42

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 7. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING LOCAL STORAGE DEVICES
                                                                                     47
 7.1. INSTALLING LOCAL STORAGE OPERATOR                                              47
 7.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR                          47
 7.3. CREATING A STANDALONE MULTICLOUD OBJECT GATEWAY                                49

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 8. VIEW OPENSHIFT DATA FOUNDATION TOPOLOGY                                   55

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 9. UNINSTALLING OPENSHIFT DATA FOUNDATION                                    56
 9.1. UNINSTALLING OPENSHIFT DATA FOUNDATION IN INTERNAL MODE                        56


   1

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere










2

    PREFACE

    PREFACE

This document describes how to deploy Red Hat OpenShift Data Foundation on VMware vSphere using
dynamic or local storage devices.










3

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

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










4

     PREFACE

     PREFACE

Red Hat OpenShift Data Foundation supports deployment on existing Red Hat OpenShift Container
Platform (RHOCP) vSphere clusters in connected or disconnected environments along with out-of-the-
box support for proxy environments.

To deploy OpenShift Data Foundation, start with the requirements in the Preparing to deploy OpenShift
Data Foundation chapter and then follow the appropriate deployment process:

 Internal mode — Deploy using dynamic storage devices, Deploy using local storage devices, or
 Deploy standalone Multicloud Object Gateway

 External mode — Deploying OpenShift Data Foundation in external mode

Additional resources

 Preparing to deploy OpenShift Data Foundation

 Deploying OpenShift Data Foundation in external mode










 5

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

  CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA
                       FOUNDATION

 Deploying OpenShift Data Foundation on OpenShift Container Platform using dynamic or local storage
 devices provides you with the option to create internal cluster resources. This will result in the internal
 provisioning of the base services, which helps to make additional storage classes available to
 applications.

 Before you begin the deployment of Red Hat OpenShift Data Foundation using dynamic or local
 storage, ensure that your resource requirements are met. See the Resource requirements section in the
 Planning guide.

  1.   Verify the rotational flag on your VMDKs before deploying object storage devices (OSDs) on
       them. For more information, see the knowledgebase article Override device rotational flag in
       ODF environment.

  2.   Optional: If you want to enable cluster-wide encryption using the external Key Management
       System (KMS) HashiCorp Vault, follow these steps:

            Ensure that you have a valid Red Hat OpenShift Data Foundation Advanced subscription.
            To know how subscriptions for OpenShift Data Foundation work, see knowledgebase article
            on OpenShift Data Foundation subscriptions.

            When the Token authentication method is selected for encryption then refer to Enabling
            cluster-wide encryption with KMS using the Token authentication method.

            When the Kubernetes authentication method is selected for encryption then refer to
            Enabling cluster-wide encryption with KMS using the Kubernetes authentication method .

            Ensure that you are using signed certificates on your Vault servers.

  3.   Optional: If you want to enable cluster-wide encryption using the external Key Management
       System (KMS) Thales CipherTrust Manager, you must first enable the Key Management
       Interoperability Protocol (KMIP) and use signed certificates on your server. Follow these steps:

        a.  Create a KMIP client if one does not exist. From the user interface, select KMIP → Client
            Profile → Add Profile.

              i.  Add the CipherTrust username to the Common Name field during profile creation.

        b.  Create a token by navigating to KMIP → Registration Token → New Registration Token.
            Copy the token for the next step.

        c.  To register the client, navigate to KMIP → Registered Clients → Add Client. Specify the
            Name. Paste the Registration Token from the previous step, then click Save.

        d.  Download the Private Key and Client Certificate by clicking Save Private Key and Save
            Certificate respectively.

        e.  To create a new KMIP interface, navigate to Admin Settings → Interfaces → Add
            Interface.

              i.   Select KMIP Key Management Interoperability Protocol and click Next.

             ii.   Select a free Port.

            iii.   Select Network Interface as all.

6

                                              CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA FOUNDATION

           iv.   Select Interface Mode as TLS, verify client cert, user name taken from client cert,
                 auth request is optional.

            v.   (Optional) You can enable hard delete to delete both metadata and material when the
                 key is deleted. It is disabled by default.

           vi.   Select the CA to be used, and click Save.

     f.   To get the server CA certificate, click on the Action menu (⋮) on the right of the newly
          created interface, and click Download Certificate.

     g.   Optional: If StorageClass encryption is to be enabled during deployment, create a key to act
          as the Key Encryption Key (KEK):

            i.   Navigate to Keys → Add Key.

           ii.   Enter Key Name.

          iii.   Set the Algorithm and Size to AES and 256 respectively.

           iv.   Enable Create a key in Pre-Active stateand set the date and time for activation.

            v.   Ensure that Encrypt and Decrypt are enabled under Key Usage.

           vi.   Copy the ID of the newly created Key to be used as the Unique Identifier during
                 deployment.

4.  Minimum starting node requirements
    An OpenShift Data Foundation cluster will be deployed with minimum configuration when the
    standard deployment resource requirement is not met. See Resource requirements section in
    Planning guide.

5.  Disaster recovery requirements
    Disaster Recovery features supported by Red Hat OpenShift Data Foundation require all of the
    following prerequisites to successfully implement a disaster recovery solution:

          A valid Red Hat OpenShift Data Foundation Advanced subscription

          A valid Red Hat Advanced Cluster Management for Kubernetes subscription
          To know how subscriptions for OpenShift Data Foundation work, see knowledgebase article
          on OpenShift Data Foundation subscriptions.

    For detailed requirements, see Configuring OpenShift Data Foundation Disaster Recovery for
    OpenShift Workloads guide, and Requirements and recommendations section of the Install
    guide in Red Hat Advanced Cluster Management for Kubernetes documentation.

6.  For deploying using local storage devices, see requirements for installing OpenShift Data
    Foundation using local storage devices. These are not applicable for deployment using dynamic
    storage devices.

7.  For VMs on VMware, ensure the disk.EnableUUID option is set to TRUE. You need to have
    vCenter account privileges to configure the VMs. For more information, see Required vCenter
    account privileges. To set the disk.EnableUUID option, use the Advanced option of the VM
    Options in the Customize hardware tab. For more information, see Installing on vSphere.
    Optional: If you want to use thick-provisioned storage for flexibility, you must create a storage
    class with zeroedthick or eagerzeroedthick disk format. For information, see VMware vSphere
    object definition.

                                                                                                     7

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

  1.1. RECOMMENDED CONFIGURATION FOR OPTIMIZED
  PERFORMANCE

  Configure VMware vSphere storage settings to optimize OpenShift Data Foundation performance for
  small block I/O workloads such as databases that require low-latency disk flushes.

  When deploying OpenShift Data Foundation to provide storage for workloads, it is important to consider
  applications that write small block size I/Os (4K-16K) and require the data to be flushed to disk before
  acknowledgement back to the client.

  A common example of such a workload is a database workload.

  In these scenarios, it is critical to ensure the storage presented to VMware limits the amount of layers
  between OpenShift Data Foundation and the actual backing physical device providing the persistent
  volume (PV) storage for the OSD.

  Due to the way OpenShift Data Foundation handles write operations, an operation that requires the
  output to be flushed to disk introduces additional latency. When a write is received by the primary OSD,
  it must first be committed to disk. Only then can OpenShift Data Foundation synchronously send two
  additional writes to the peer OSDs, which must also commit to disk before the write is acknowledged
  back to the client.

  Write operations can be further delayed depending on the customer’s infrastructure and platform. For
  instance, if the storage system includes both a virtualization layer and a RAID layer, each further delays
  the completion of the operation. These stacked layers significantly impact performance, especially for
  small block (4K-16K) I/O that require the write to be flushed to disk before being acknowledged.

  Based on the I/O flow just described, deploying OpenShift Data Foundation on VMware vSAN is not
  recommended when the backing storage device being presented into the datastore for vSAN uses
  RAID.

  1.1.1. Recommended configuration guidelines

        Use Passthrough NVMe disks: The optimal setup for OpenShift Data Foundation on VMware is
        to configure it with passthrough NVMe disks. This bypasses the vSAN layer, thereby avoiding
        the additional latency it introduces.

        Leverage LSO for device consumption: Passthrough devices should be consumed using Local
        Storage Operator (LSO).

        Avoid Default vSAN StorageClass for OSDs: In IPI deployments of OpenShift on vSAN, a
        default StorageClass thin-csi is automatically created. Do not use this StorageClass for OSDs,
        as it relies on RAID-1, which adds unnecessary overhead. Instead, use a StorageClass configured
        in thick lazy zeroed format.

        Do not share storage with other workloads: Sharing storage between OpenShift Data
        Foundation and other workloads can significantly degrade performance. A dedicated
        StorageClass should be used exclusively for OpenShift Data Foundation.

        Isolate ODF Storage Network Traffic: To prevent network congestion, it is recommended to
        use an isolated VM network for OpenShift Data Foundation storage traffic.

  Sample impacted workloads:

        MongoDB



    8

                        CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA FOUNDATION

 SQLDB

 Cloud Pak for Data

 Cloud Pak for AI

Remediation steps

If your deployment currently uses a configuration similar to the one described above, we recommend
reviewing the Knowledge Base article for detailed guidance on migrating from a vSAN-based setup to a
passthrough device configuration.

             IMPORTANT

             If your VMware deployment matches the configuration described in the introduction of
             this section and you choose not to migrate to a setup that minimizes the layers between
             the Software-Defined Storage (SDS) and the physical storage, your environment may
             not meet the I/O performance requirements of certain applications.

Additional resources

 Knowledge Base article: Migrating from vSAN-based setup to passthrough device configuration

1.2. REQUIREMENTS FOR INSTALLING OPENSHIFT DATA
FOUNDATION USING LOCAL STORAGE DEVICES

Before installing OpenShift Data Foundation using local storage devices, your cluster must meet certain
requirements.

Node requirements The cluster must consist of at least three OpenShift Container Platform worker or
infrastructure nodes with locally attached-storage devices on each of them.

 Each of the three selected nodes must have at least one raw block device available. OpenShift
 Data Foundation uses the one or more available raw block devices.

                    NOTE

                    Make sure that the devices have a unique by-id device name for each available
                    raw block device.

 The devices you use must be empty, the disks must not include Physical Volumes (PVs), Volume
 Groups (VGs), or Logical Volumes (LVs) remaining on the disk.

For more information, see the Resource requirements section in the Planning guide.

Disaster recovery requirements

Disaster Recovery features supported by Red Hat OpenShift Data Foundation require all of the
following prerequisites to successfully implement a disaster recovery solution:

 A valid Red Hat OpenShift Data Foundation Advanced subscription.

 A valid Red Hat Advanced Cluster Management (RHACM) for Kubernetes subscription.


To know in detail how subscriptions for OpenShift Data Foundation work, see the knowledgebase article
                        9

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

  To know in detail how subscriptions for OpenShift Data Foundation work, see the knowledgebase article
  on OpenShift Data Foundation subscriptions.

  For detailed disaster recovery solution requirements, see the Configuring OpenShift Data Foundation
  Disaster Recovery for OpenShift Workloads guide, and Requirements and recommendations section of
  the Install guide in Red Hat Advanced Cluster Management for Kubernetes documentation.

  Arbiter stretch cluster requirements

  In this case, a single cluster is stretched across two zones with a third zone as the location for the arbiter.

  To know in detail how subscriptions for OpenShift Data Foundation work, see the knowledgebase article
  on OpenShift Data Foundation subscriptions.

        NOTE

        You cannot enable Flexible scaling and Arbiter both at the same time as they have
        conflicting scaling logic. With Flexible scaling, you can add one node at a time to your
        OpenShift Data Foundation cluster. Whereas, in an Arbiter cluster, you need to add at
        least one node in each of the two data zones.

  Minimum starting node requirements

  An OpenShift Data Foundation cluster is deployed with a minimum configuration when the resource
  requirement for a standard deployment is not met.

  For more information, see the Resource requirements section in the Planning guide.

  Disconnected environment requirements

  For the list of packages to be mirrored for installing OpenShift Data Foundation in disconnected
  environments, see the Disconnected environment section in the Planning your deployment guide.

  Additional resources

        Planning guide - Disconnected environment

        Planning guide - Resource requirements

        Knowledgebase article on OpenShift Data Foundation subscriptions

        Configuring OpenShift Data Foundation Disaster Recovery for OpenShift Workloads

        Red Hat Advanced Cluster Management for Kubernetes - Requirements and recommendations

        Knowledgebase article on OpenShift Data Foundation subscriptions










    10

                                                              CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

Deploying OpenShift Data Foundation on OpenShift Container Platform using dynamic storage devices
provided by VMware vSphere (disk format: thin) provides you with the option to create internal cluster
resources. This will result in the internal provisioning of the base services, which helps to make additional
storage classes available to applications.

             NOTE

             Both internal and external OpenShift Data Foundation clusters are supported on VMware
             vSphere. See Planning your deployment for more information about deployment
             requirements.

Also, ensure that you have addressed the requirements in Preparing to deploy OpenShift Data
Foundation chapter before proceeding with the below steps for deploying using dynamic storage
devices:

1.      Install the Red Hat OpenShift Data Foundation Operator .

2.      Create an OpenShift Data Foundation Cluster .

2.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.

Prerequisites

        Access to an OpenShift Container Platform cluster using an account with cluster-admin and
        operator installation permissions.

        You must have at least three worker or infrastructure nodes in the Red Hat OpenShift
        Container Platform cluster.

        For additional resource requirements, see the Planning your deployment guide.

            IMPORTANT

                 When you need to override the cluster-wide default node selector for OpenShift
                 Data Foundation, you can use the following command to specify a blank node
                 selector for the openshift-storage namespace (create openshift-storage
                 namespace in this case):

                 $ oc annotate namespace openshift-storage openshift.io/node-selector=

                 Taint a node as infra to ensure only Red Hat OpenShift Data Foundation
                 resources are scheduled on that node. This helps you save on subscription costs.
                 For more information, see the How to use dedicated worker nodes for Red Hat
                 OpenShift Data Foundation section in the Managing and Allocating Storage
                 Resources guide.

Procedure

                                                                                                           11

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

1.   Log in to the OpenShift Web Console.

2.   Click Ecosystem → Software Catalog.

3.   Scroll or type OpenShift Data Foundation into the Filter by keyword box to find the
     OpenShift Data Foundation Operator.

4.   Click Install.

5.   Set the following options on the Install Operator page:

      a.   Update Channel as stable-4.22.

      b.   Installation Mode as A specific namespace on the cluster.

      c.   Installed Namespace as Operator recommended namespace openshift-storage. If
           Namespace openshift-storage does not exist, it is created during the operator installation.

      d.   Select Approval Strategy as Automatic or Manual.
           If you select Automatic updates, then the Operator Lifecycle Manager (OLM)
           automatically upgrades the running instance of your Operator without any intervention.

           If you select Manual updates, then the OLM creates an update request. As a cluster
           administrator, you must then manually approve that update request to update the Operator
           to a newer version.

      e.   Ensure that the Enable option is selected for the Console plugin.

      f.   Click Install.

Verification

     After the operator is successfully installed, the web console automatically reloads to apply the
     changes. During this process, a temporary error message might appear on the page and this is
     expected and disappears after the refresh completes.

     In the Web Console:

           Navigate to Installed Operators and verify that the OpenShift Data Foundation Operator
           shows a green tick indicating successful installation.

           Navigate to Storage and verify if the Data Foundation dashboard is available.

2.2. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE
TOKEN AUTHENTICATION METHOD

You can enable the key value backend path and policy in the vault for token authentication.

Prerequisites

     Administrator access to the vault.

     A valid Red Hat OpenShift Data Foundation Advanced subscription. For more information, see
     the knowledgebase article on OpenShift Data Foundation subscriptions .


Carefully, select a unique path name as the backend path that follows the naming convention
12

             CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

       Carefully, select a unique path name as the backend path that follows the naming convention
       since you cannot change it later.

Procedure

  1.   Enable the Key/Value (KV) backend path in the vault.
       For vault KV secret engine API, version 1:

         $ vault secrets enable -path=odf kv

       For vault KV secret engine API, version 2:

         $ vault secrets enable -path=odf kv-v2

  2.   Create a policy to restrict the users to perform a write or delete operation on the secret:

         echo '
         path "odf/*" {
          capabilities = ["create", "read", "update", "delete", "list"]
         }
         path "sys/mounts" {
         capabilities = ["read"]
         }'| vault policy write odf -

  3.   Create a token that matches the above policy:

         $ vault token create -policy=odf -format json

2.3. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE
KUBERNETES AUTHENTICATION METHOD

You can enable the Kubernetes authentication method for cluster-wide encryption using the Key
Management System (KMS).

Prerequisites

       Administrator access to Vault.

       A valid Red Hat OpenShift Data Foundation Advanced subscription. For more information, see
       the knowledgebase article on OpenShift Data Foundation subscriptions .

       The OpenShift Data Foundation operator must be installed from the Software Catalog.

       Select a unique path name as the backend path that follows the naming convention carefully.
       You cannot change this path name later.

Procedure

1.     Create a service account:

         $ oc -n openshift-storage create serviceaccount <serviceaccount_name>

       where, <serviceaccount_name> specifies the name of the service account.

             13

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

    For example:

     $ oc -n openshift-storage create serviceaccount odf-vault-auth

2.  Create clusterrolebindings and clusterroles:

     $ oc -n openshift-storage create clusterrolebinding vault-tokenreview-binding --
     clusterrole=system:auth-delegator --serviceaccount=openshift-
     storage:_<serviceaccount_name>_

    For example:

     $ oc -n openshift-storage create clusterrolebinding vault-tokenreview-binding --
     clusterrole=system:auth-delegator --serviceaccount=openshift-storage:odf-vault-auth

3.  Create a secret for the serviceaccount token and CA certificate.

     $ cat <<EOF | oc create -f -
     apiVersion: v1
     kind: Secret
     metadata:
      name: odf-vault-auth-token
      namespace: openshift-storage
      annotations:
      kubernetes.io/service-account.name: <serviceaccount_name>
     type: kubernetes.io/service-account-token
     data: {}
     EOF

    where, <serviceaccount_name> is the service account created in the earlier step.

4.  Get the token and the CA certificate from the secret.

     $ SA_JWT_TOKEN=$(oc -n openshift-storage get secret odf-vault-auth-token -o jsonpath="
     {.data['token']}" | base64 --decode; echo)
     $ SA_CA_CRT=$(oc -n openshift-storage get secret odf-vault-auth-token -o jsonpath="
     {.data['ca\.crt']}" | base64 --decode; echo)

5.  Retrieve the OCP cluster endpoint.

     $ OCP_HOST=$(oc config view --minify --flatten -o jsonpath="{.clusters[0].cluster.server}")

6.  Fetch the service account issuer:

     $ oc proxy &
     $ proxy_pid=$!
     $ issuer="$( curl --silent http://127.0.0.1:8001/.well-known/openid-configuration | jq -r
     .issuer)"
     $ AUDIENCE=$(echo '{"apiVersion": "authentication.k8s.io/v1", "kind": "TokenRequest"}' |
     kubectl create f --raw /api/v1/namespaces/default/serviceaccounts/default/token | jq -r
     '.spec.audiences' | jq -r .[0])
     $ kill $proxy_pid



14

                           CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

   Or, for Vault version 1.21.0 and higher:

     $ oc proxy &
     $ proxy_pid=$!
     $ ISSUER=$(echo '{"apiVersion": "authentication.k8s.io/v1", "kind": "TokenRequest"}' \  |
     kubectl create f --raw /api/v1/namespaces/default/serviceaccounts/default/token \ | jq -r
     '.status.token' \ | cut -d . -f2 \ | base64 -d \ | jq '.iss')
     $ AUDIENCE=$(echo '{"apiVersion": "authentication.k8s.io/v1", "kind": "TokenRequest"}' |
     kubectl create f --raw /api/v1/namespaces/default/serviceaccounts/default/token | jq -r
     '.spec.audiences' | jq -r .[0])
     $ kill $proxy_pid

7.  Use the information collected in the previous step to setup the Kubernetes authentication
    method in Vault:

     $ vault auth enable kubernetes

     $ vault write auth/kubernetes/config \
       token_reviewer_jwt="$SA_JWT_TOKEN" \
       kubernetes_host="$OCP_HOST" \
       kubernetes_ca_cert="$SA_CA_CRT" \
       issuer="$issuer"

    Or, for Vault version 1.21.0 and higher:

     $ oc exec -n vault -ti vault-0 – \
      vault write auth/kubernetes/config \
      token_reviewer_jwt="$SA_JWT_TOKEN" \
      kubernetes_host="$OCP_HOST" \
      kubernetes_ca_cert="$SA_CA_CRT" \
      issuer="$ISSUER"\
      token_audience="$AUDIENCE"

        IMPORTANT

        To configure the Kubernetes authentication method in Vault when the issuer is
        empty:

            $ vault write auth/kubernetes/config \
                       token_reviewer_jwt="$SA_JWT_TOKEN" \
                       kubernetes_host="$OCP_HOST" \
                       kubernetes_ca_cert="$SA_CA_CRT"

8.  Enable the Key/Value (KV) backend path in Vault.
    For Vault KV secret engine API, version 1:

     $ vault secrets enable -path=odf kv

    For Vault KV secret engine API, version 2:

     $ vault secrets enable -path=odf kv-v2



                       15

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

 9.  Create a policy to restrict the users to perform a write or delete operation on the secret:

         echo '
         path "odf/*" {
          capabilities = ["create", "read", "update", "delete", "list"]
         }
         path "sys/mounts" {
         capabilities = ["read"]
         }'| vault policy write odf -

10.  Generate the roles:

         $ vault write auth/kubernetes/role/odf-rook-ceph-op \
          bound_service_account_names=rook-ceph-system,rook-ceph-osd,noobaa \
          bound_service_account_namespaces=openshift-storage \
          policies=odf \
          ttl=1440h

     The role odf-rook-ceph-op is later used while you configure the KMS connection details during
     the creation of the storage system.

         $ vault write auth/kubernetes/role/odf-rook-ceph-osd \
          bound_service_account_names=rook-ceph-osd \
          bound_service_account_namespaces=openshift-storage \
          policies=odf \
          ttl=1440h

2.3.1. Enabling and disabling key rotation when using KMS

Security common practices require periodic encryption of key rotation. You can enable or disable key
rotation when using KMS.

2.3.1.1. Enforcing precedence for key rotation

Configure key rotation precedence to control whether the rotation schedule is read from the PVC
annotation or the storage class in OpenShift Data Foundation.

In key rotation, precedence refers to the order in which the system checks for scheduled annotations. In
OpenShift Data Foundation, the default precedence is set to storage class (recommended). This means
the system reads annotations only from the storage class.

However, if you want the system to check the persistent volume claim (PVC) first and then fall back to
the storage class, you can configure this behavior by setting the schedule-precedence to PVC in the
CSI-addons ConfigMap.

Procedure

1.          Define the ConfigMap for csi-addons to set PVC precedence:

         apiVersion: v1
         kind: ConfigMap
         metadata:
          name: csi-addons-config


16

                 CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

         namespace: openshift-storage
         data:
         "schedule-precedence": "pvc"

2.  Restart the csi-addons operator pod for the changes to take effect:

         oc delete po -n openshift-storage -l "app.kubernetes.io/name=csi-addons"

2.3.1.2. Enabling key rotation

Enable key rotation to automatically re-encrypt persistent volume data at a scheduled interval using the
CSI Addons framework in OpenShift Data Foundation.

To enable key rotation, add the annotation keyrotation.csiaddons.openshift.io/schedule: <value> to
a PersistentVolumeClaim, Namespace, or StorageClass (in decreasing order of precedence).

             IMPORTANT

             Key rotation is only supported for RBD backed volumes.

2.3.1.2.1. Annotating a namespace for key rotation

Add the key rotation annotation to a namespace so that all RBD-backed persistent volumes in that
namespace are re-encrypted on the configured schedule.

Prerequisites

     Key rotation is only supported for RBD backed volumes.

Procedure

1.   Verify the namespace exists:

         $ oc get namespace default
         NAME    STATUS AGE
         default Active 5d2h

2.   Annotate the namespace with the desired rotation schedule:

         $ oc annotate namespace default "keyrotation.csiaddons.openshift.io/schedule=@weekly"
         namespace/default annotated

     <value> can be @hourly, @daily, @weekly, @monthly, or @yearly. If <value> is empty, the
     default is @weekly.










             NOTE
                 17

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

              NOTE

                 When any schedule for key rotation is used, including @daily, @weekly,
                 @monthly, and cron expressions such as 0 * * * *, multiple CronJobs are
                 created. To prevent all Jobs from running simultaneously and overloading
                 high‑scale systems, executions are staggered within a 2‑hour window from
                 the scheduled start time (e.g., @weekly would be 00:00, 00:05, 00:10, and
                 so on).

                            If the configured schedule interval is less than 2 hours, the stagger
                            window is reduced to match the interval. For example, a job running
                            every 30 minutes will be staggered within a 00:00–00:30 window rather
                            than 00:00–02:00.

2.3.1.2.2. Annotating a storage class for key rotation

Add the key rotation annotation to a storage class so that all RBD-backed persistent volumes
provisioned from that storage class are re-encrypted on the configured schedule.

Prerequisites

         Key rotation is only supported for RBD backed volumes.

Procedure

    1.   Verify the storage class exists:

          $ oc get storageclass rbd-sc
          NAME   PROVISIONER     RECLAIMPOLICY     VOLUMEBINDINGMODE
          ALLOWVOLUMEEXPANSION     AGE
          rbd-sc rbd.csi.ceph.com     Delete     Immediate     true      5d2h

    2.   Annotate the storage class with the desired rotation schedule:

          $ oc annotate storageclass rbd-sc "keyrotation.csiaddons.openshift.io/schedule=@weekly"
          storageclass.storage.k8s.io/rbd-sc annotated

         <value> can be @hourly, @daily, @weekly, @monthly, or @yearly. If <value> is empty, the
         default is @weekly.

2.3.1.2.3. Annotating a PersistentVolumeClaim for key rotation

Add the key rotation annotation directly to a PersistentVolumeClaim to re-encrypt that volume on the
configured schedule, overriding any namespace or storage class annotation.

Prerequisites

         Key rotation is only supported for RBD backed volumes.

Procedure

  1.     Verify the PersistentVolumeClaim exists:

          $ oc get pvc data-pvc

18

             CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

         NAME    STATUS     VOLUME                            CAPACITY  ACCESS MODES
         STORAGECLASS         AGE
         data-pvc Bound       pvc-f37b8582-4b04-4676-88dd-e1b95c6abf74  1Gi          RWO
         default    20h

  2.  Annotate the PVC with the desired rotation schedule:

         $ oc annotate pvc data-pvc "keyrotation.csiaddons.openshift.io/schedule=@weekly"
         persistentvolumeclaim/data-pvc annotated

      <value> can be @hourly, @daily, @weekly, @monthly, or @yearly. If <value> is empty, the
      default is @weekly.

  3.  Verify the encryption key rotation cron job was created:

         $ oc get encryptionkeyrotationcronjobs.csiaddons.openshift.io
         NAME     SCHEDULE     SUSPEND     ACTIVE     LASTSCHEDULE                   AGE
         data-pvc-1642663516     @weekly        3s

  4.  Optional: Update the schedule by overwriting the annotation:

         $ oc annotate pvc data-pvc "keyrotation.csiaddons.openshift.io/schedule=*/1 * * * *" --
         overwrite=true
         persistentvolumeclaim/data-pvc annotated

  5.  Verify the updated cron job schedule:

         $ oc get encryptionkeyrotationcronjobs.csiaddons.openshift.io
         NAME     SCHEDULE                SUSPEND     ACTIVE  LASTSCHEDULE           AGE
         data-pvc-1642664617 */1 * * * *        3s

2.3.1.3. Disabling key rotation

Disable automatic key rotation for all PVCs in a storage class or for a specific persistent volume claim in
OpenShift Data Foundation.

You can disable key rotation for the following:

      All the persistent volume claims (PVCs) of storage class

      A specific PVC

Disabling key rotation for all PVCs of a storage class

To disable key rotation for all PVCs, update the annotation of the storage class:

 $ oc get storageclass rbd-sc
 NAME          PROVISIONER     RECLAIMPOLICY     VOLUMEBINDINGMODE
 ALLOWVOLUMEEXPANSION     AGE
 rbd-sc   rbd.csi.ceph.com Delete       Immediate     true      5d2h

 $ oc annotate storageclass rbd-sc "keyrotation.csiaddons.openshift.io/enable: false"
 storageclass.storage.k8s.io/rbd-sc annotated


                                                                                     19

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

Disabling key rotation for a specific persistent volume claim

Procedure

    1.   Identify the EncryptionKeyRotationCronJob CR for the PVC you want to disable key rotation
         on:

            $ oc get encryptionkeyrotationcronjob -o jsonpath='{range .items[?
            (@.spec.jobTemplate.spec.target.persistentVolumeClaim=="<PVC_NAME>")]}
            {.metadata.name}{"\n"}{end}'

         Where <PVC_NAME> is the name of the PVC that you want to disable.

    2.   Apply the following to the EncryptionKeyRotationCronJob CR from the previous step to
         disable the key rotation:

         a.  Update the csiaddons.openshift.io/state annotation from managed to unmanaged:

              $ oc annotate encryptionkeyrotationcronjob <encryptionkeyrotationcronjob_name>
              "csiaddons.openshift.io/state=unmanaged" --overwrite=true

             Where <encryptionkeyrotationcronjob_name> is the name of the
             EncryptionKeyRotationCronJob CR.

         b.  Add suspend: true under the spec field:

              $ oc patch encryptionkeyrotationcronjob <encryptionkeyrotationcronjob_name> -p
              '{"spec": {"suspend": true}}' --type=merge.

             The key rotation will be disabled for the PVC.

2.4. CREATING OPENSHIFT DATA FOUNDATION CLUSTER

Create an OpenShift Data Foundation cluster after you install the OpenShift Data Foundation operator.

Prerequisites

         The OpenShift Data Foundation operator must be installed from the Software Catalog. For
         more information, see Installing OpenShift Data Foundation Operator .

         For VMs on VMware, if you have not already enabled UUIDs, ensure the disk.EnableUUID
         option is set to TRUE. You need to have vCenter account privileges to configure the VMs. For
         more information, see Required vCenter account privileges . To set the disk.EnableUUID
         option, use the Advanced option of the VM Options in the Customize hardware tab. For more
         information, see Installing on vSphere.

         Optional: If you want to use thick-provisioned storage for flexibility, you must create a storage
         class with zeroedthick or eagerzeroedthick disk format. For information, see VMware vSphere
         object definition.

  Procedure

  1.     In the OpenShift Web Console, click Storage → Data Foundation → Storage Systems →
         Create StorageSystem.

20

                          CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

2.  In the Backing storage page, select the following:

     a.  Select the Use an existing StorageClass option.

     b.  Select the Storage Class.
         By default, it is set to thin-csi. If you have created a storage class with zeroedthick or
         eagerzeroedthick disk format for thick-provisioned storage, then that storage class is listed
         in addition to the default, thin-csi storage class.

     c.  Click Next.

3.  In the Advanced Settings page, you can configure the following optional items:

         Enable Network Files System
         Automatically enable NFS for the cluster

         Use Ceph RBD as the default StorageClass
         Sets Ceph RBD as the default StorageClass, eliminating the need to manually annotate a
         StorageClass.

         Set default StorageClass for virtualization
         Marks the RBD virtualization storage class as the default for KubeVirt VM disks (PVs) after
         installation.

         Use external PostgreSQL [Technology preview]
         Enables the use of an external PostgreSQL instance. This provides a high-availability
         solution for the Multicloud Object Gateway, where the PostgreSQL pod is a single point of
         failure.

                      IMPORTANT

                      OpenShift Data Foundation ships PostgreSQL images maintained by Red
                      Hat, which are used to store metadata for the Multicloud Object Gateway.
                      This PostgreSQL usage is at the application level.

                      As a result, OpenShift Data Foundation does not perform database-level
                      optimizations or in-depth insights.

                      If you have well-maintained and optimized PostgreSQL instance, then it is
                      recommended to use it. OpenShift Data Foundation supports external
                      PostgreSQL instances.

                      Any PostgreSQL-related issues requiring code changes or deep technical
                      analysis may need to be addressed upstream. This could result in longer
                      resolution times.

         a. Provide the following connection details:

                 Username

                 Password

                 Server name and Port

                 Database name



                      21

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

        b.  Select Enable TLS/SSL checkbox to enable encryption for the Postgres server.

       Enable automatic backup
       Allows automatic backup for the Multicloud Object Gateway metadata database.

        a.  Select the backup frequency: Daily, Weekly, or Monthly.

        b.  Specify the number of backups to retain.

                         NOTE

                         This option is disabled if the Use external PostgreSQL option is
                         selected.

a.  Click Next.

    1. In the Capacity and nodes page, provide the necessary information:

   b. Select a value for Requested Capacity from the dropdown list. It is set to 2 TiB by default.

               NOTE

               Once you select the initial storage capacity, cluster expansion is performed only
               using the selected usable capacity (three times of raw storage).

c.  In the Select Nodes section, select at least three available nodes.

d.  In the Configure performance section, select one of the following performance profiles:

       Lean
       Use this in a resource constrained environment with minimum resources that are lower than
       the recommended. This profile minimizes resource consumption by allocating fewer CPUs
       and less memory.

       Balanced (default)
       Use this when recommended resources are available. This profile provides a balance
       between resource consumption and performance for diverse workloads.

       Performance
       Use this in an environment with sufficient resources to get the best performance. This
       profile is tailored for high performance by allocating ample memory and CPUs to ensure
       optimal execution of demanding workloads.

                  NOTE

                  You have the option to configure the performance profile even after the
                  deployment using the Configure performance option from the options
                  menu of the StorageSystems tab.








                  IMPORTANT
22

                            CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

                        IMPORTANT

                        Before selecting a resource profile, make sure to check the current
                        availability of resources within the cluster. Opting for a higher resource
                        profile in a cluster with insufficient resources might lead to installation
                        failures.

           For more information about resource requirements, see Resource requirement for
           performance profiles.

e.   Optional: Select the Taint nodes checkbox to dedicate the selected nodes for OpenShift Data
     Foundation.
     Spread the worker nodes across three different physical nodes, racks, or failure domains for
     high availability.

     Use vCenter anti-affinity to align OpenShift Data Foundation rack labels with physical nodes
     and racks in the data center to avoid scheduling two worker nodes on the same physical chassis.

     If the nodes selected do not match the OpenShift Data Foundation cluster requirement of the
     aggregated 30 CPUs and 72 GiB of RAM, a minimal cluster is deployed. For minimum starting
     node requirements, see the Resource requirements section in the Planning guide.

     Select the Taint nodes checkbox to make selected nodes dedicated for OpenShift Data
     Foundation.

f.   Optional: Select the Enable automatic capacity scaling for your clustercheckbox.
     When automatic capacity scaling is enabled, additional raw capacity equivalent to the
     configured deployment size is automatically added to the cluster when used capacity reaches
     70%. This ensures your deployment scales seamlessly to meet demand.

     This option is disabled in lean profile mode, LSO deployment, and external mode deployment.

                  IMPORTANT

                  This may incur additional costs for the underlying storage.

      i.   Set the cluster expansion limit from the dropdown. This is the maximum the cluster can
           expand in the cloud. Automatic scaling is suspended if this limit is exceeded.

g.   Click Next.

      1.   Optional: In the Security and network page, configure the following based on your
           requirements:

h.   To enable encryption, select Enable data encryption for block and file storage.

      i.   Select either one or both the encryption levels:

            Cluster-wide encryption
            Encrypts the entire cluster (block and file).

            StorageClass encryption
            Creates encrypted persistent volume (block only) using encryption enabled storage
            class.


     ii. Optional: Select the Connect to an external key management servicecheckbox. This is
                            23

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

ii.   Optional: Select the Connect to an external key management servicecheckbox. This is
      optional for cluster-wide encryption.

      A.  From the Key Management Service Providerdrop-down list, select one of the
          following providers and provide the necessary details:

          Vault

          I. Select an Authentication Method.

              Using Token authentication method

               Enter a unique Connection Name, host Address of the Vault server
               ('https://<hostname or ip>'), Port number and Token.

               Expand Advanced Settings to enter additional settings and certificate
               details based on your Vault configuration:

                Enter the Key Value secret path in Backend Path that is dedicated
                and unique to OpenShift Data Foundation.

               Optional: Enter TLS Server Name, Authentication Path, and Vault
               Enterprise Namespace.

               Upload the respective PEM encoded certificate file to provide the CA
               Certificate, Client Certificate and Client Private Key.

               Click Save.

              Using Kubernetes authentication method

               Enter a unique Vault Connection Name, host Address of the Vault
               server ('https://<hostname or ip>'), Port number and Role name.

               Expand Advanced Settings to enter additional settings and certificate
               details based on your Vault configuration:

                Enter the Key Value secret path in Backend Path that is dedicated
                and unique to OpenShift Data Foundation.

                Optional: Enter TLS Server Name, Authentication Path, and
                Vault Enterprise Namespaceif applicable.

                Upload the respective PEM encoded certificate file to provide the
                CA Certificate, Client Certificate and Client Private Key .

               Click Save.

                          NOTE

                          In case you need to enable key rotation for Vault KMS,
                          run the following command in the OpenShift web
                          console after the storage cluster is created:

                                           $ oc patch storagecluster ocs-storagecluster -n
                                           openshift-storage --type=json -p '[{"op": "add",
                                           "path":"/spec/encryption/keyRotation/enable", "value":
                                           true}]'

24

                             CHAPTER 2. DEPLOY USING DYNAMIC STORAGE DEVICES

                     Thales CipherTrust Manager (using KMIP)

                   I.   Enter a unique Connection Name for the Key Management service within the
                        project.

                  II.   In the Address and Port sections, enter the IP of Thales CipherTrust Manager
                        and the port where the KMIP interface is enabled. For example:

                         Address: 123.34.3.2

                         Port: 5696

                 III.   Upload the Client Certificate, CA certificate, and Client Private Key.

                  IV.   If StorageClass encryption is enabled, enter the Unique Identifier to be used for
                        encryption and decryption generated above.

                   V.   The TLS Server field is optional and used when there is no DNS entry for the
                        KMIP endpoint. For example, kmip_all_<port>.ciphertrustmanager.local.

i. To enable in-transit encryption, select In-transit encryption.

 i.   Select a Network.

ii.   Click Next.

            1.   In the Review and create page, review the configuration details.
                 To modify any configuration settings, click Back.

            2.   Click Create StorageSystem.

                    NOTE

                    When your deployment has five or more nodes, racks, or rooms, and when there
                    are five or more number of failure domains present in the deployment, you can
                    configure Ceph monitor counts based on the number of racks or zones. An alert
                    is displayed in the notification panel or Alert Center of the OpenShift Web
                    Console to indicate the option to increase the number of Ceph monitor counts.
                    You can use the Configure option in the alert to configure the Ceph monitor
                    counts. For more information, see Resolving low Ceph monitor count alert .

Verification

To verify the final Status of the installed storage cluster:

 a.   In the OpenShift Web Console, navigate to Storage → Data Foundation → Storage
      System → ocs-storagecluster.

 b.   Verify that Status of StorageCluster is Ready and has a green tick mark next to it.

          1.     To verify that all components for OpenShift Data Foundation are successfully installed,
                 see Verifying your OpenShift Data Foundation deployment .

Additional resources

   Alerts



                         25

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

CHAPTER 3. DEPLOY USING LOCAL STORAGE DEVICES

Deploying OpenShift Data Foundation on OpenShift Container Platform using local storage devices
provides you with the option to create internal cluster resources. This will result in the internal
provisioning of the base services, which helps to make additional storage classes available to
applications.

Use this section to deploy OpenShift Data Foundation on VMware where OpenShift Container Platform
is already installed.

Also, ensure that you have addressed the requirements in Preparing to deploy OpenShift Data
Foundation chapter before proceeding with the next steps.

1.   Installing Local Storage Operator

2.   Install the Red Hat OpenShift Data Foundation Operator .

3.   Create the OpenShift Data Foundation Cluster .

3.1. INSTALLING LOCAL STORAGE OPERATOR

Install the Local Storage Operator from the Software Catalog before creating Red Hat OpenShift Data
Foundation clusters on local storage devices.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Ecosystem → Software Catalog.

3.   Type local storage in the Filter by keyword box to find the Local Storage Operator from the
     list of operators, and click on it.

4.   Set the following options on the Install Operator page:

      a.  Update channel as stable.

      b.  Installation mode as A specific namespace on the cluster.

      c.  Installed Namespace as Operator recommended namespace openshift-local-storage.

      d.  Update approval as Automatic.

5.   Click Install.

Verification

     Verify that the Local Storage Operator shows a green tick indicating successful installation.

3.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.


26

                                                          CHAPTER 3. DEPLOY USING LOCAL STORAGE DEVICES

    Prerequisites

       Access to an OpenShift Container Platform cluster using an account with cluster-admin and
       operator installation permissions.

       You must have at least three worker or infrastructure nodes in the Red Hat OpenShift
       Container Platform cluster.

       For additional resource requirements, see the Planning your deployment guide.

                IMPORTANT

                     When you need to override the cluster-wide default node selector for OpenShift
                     Data Foundation, you can use the following command to specify a blank node
                     selector for the openshift-storage namespace (create openshift-storage
                     namespace in this case):

                         $ oc annotate namespace openshift-storage openshift.io/node-selector=

                     Taint a node as infra to ensure only Red Hat OpenShift Data Foundation
                     resources are scheduled on that node. This helps you save on subscription costs.
                     For more information, see the How to use dedicated worker nodes for Red Hat
                     OpenShift Data Foundation section in the Managing and Allocating Storage
                     Resources guide.

Procedure

  1.   Log in to the OpenShift Web Console.

  2.   Click Ecosystem → Software Catalog.

  3.   Scroll or type OpenShift Data Foundation into the Filter by keyword box to find the
       OpenShift Data Foundation Operator.

  4.   Click Install.

  5.   Set the following options on the Install Operator page:

       a.   Update Channel as stable-4.22.

       b.   Installation Mode as A specific namespace on the cluster.

       c.   Installed Namespace as Operator recommended namespace openshift-storage. If
            Namespace openshift-storage does not exist, it is created during the operator installation.

       d.   Select Approval Strategy as Automatic or Manual.
            If you select Automatic updates, then the Operator Lifecycle Manager (OLM)
            automatically upgrades the running instance of your Operator without any intervention.

            If you select Manual updates, then the OLM creates an update request. As a cluster
            administrator, you must then manually approve that update request to update the Operator
            to a newer version.

       e.   Ensure that the Enable option is selected for the Console plugin.

       f.   Click Install.

                                                                                                     27

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

Verification

     After the operator is successfully installed, the web console automatically reloads to apply the
     changes. During this process, a temporary error message might appear on the page and this is
     expected and disappears after the refresh completes.

     In the Web Console:

         Navigate to Installed Operators and verify that the OpenShift Data Foundation Operator
         shows a green tick indicating successful installation.

         Navigate to Storage and verify if the Data Foundation dashboard is available.

3.3. CREATING OPENSHIFT DATA FOUNDATION CLUSTER ON
VMWARE VSPHERE

Create an OpenShift Data Foundation storage cluster on VMware vSphere using the OpenShift Web
Console.

VMware vSphere supports the following three types of local storage:

     Virtual machine disk (VMDK)

     Raw device mapping (RDM)

     VMDirectPath I/O

             NOTE

             When using thin-provisioned VMDKs with the Local Storage Operator (LSO) on VMware
             vSphere, all the disk space allocated to those VMDKs appears as fully consumed in the
             VMware vSphere interface as soon as the disks are claimed by LSO to create OSD
             volumes. This is expected behavior. The thin-provisioning concept at the VMware layer
             does not apply once OpenShift Data Foundation takes ownership of the disks.

Prerequisites

     Ensure that all the requirements in the Requirements for installing OpenShift Data Foundation
     using local storage devices section are met.

     You must have a minimum of three worker nodes with the same storage type and size attached
     to each node to use local storage devices on VMware.

     Ensure that the disk type is SSD, which is the only supported disk type.

     For VMs on VMware vSphere, ensure the disk.EnableUUID option is set to TRUE. You need to
     have vCenter account privileges to configure the VMs. For more information, see Required
     vCenter account privileges. To set the disk.EnableUUID option, use the Advanced option of
     the VM Options in the Customize hardware tab. For more information, see Installing on
     vSphere.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation → Storage Systems →
     Create StorageSystem.


28

                                                   CHAPTER 3. DEPLOY USING LOCAL STORAGE DEVICES

2.  In the Backing storage page, select the Create a new StorageClass using the local storage
    devices option and click Next.

3.  In the Optional settings page, you can configure the following optional items:

     Enable Network Files System
     Automatically enable NFS for the cluster

     Use Ceph RBD as the default StorageClass
     Sets Ceph RBD as the default StorageClass, eliminating the need to manually annotate a
     StorageClass.

     Set default StorageClass for virtualization
     Marks the RBD virtualization storage class as the default for KubeVirt VM disks (PVs) after
     installation.

     Use external PostgreSQL [Technology preview]
     Enables the use of an external PostgreSQL instance. This provides a high-availability
     solution for the Multicloud Object Gateway, where the PostgreSQL pod is a single point of
     failure.

                  IMPORTANT

                  OpenShift Data Foundation ships PostgreSQL images maintained by Red
                  Hat, which are used to store metadata for the Multicloud Object Gateway.
                  This PostgreSQL usage is at the application level.

                  As a result, OpenShift Data Foundation does not perform database-level
                  optimizations or in-depth insights.

                  If you have well-maintained and optimized PostgreSQL instance, then it is
                  recommended to use it. OpenShift Data Foundation supports external
                  PostgreSQL instances.

                  Any PostgreSQL-related issues requiring code changes or deep technical
                  analysis may need to be addressed upstream. This could result in longer
                  resolution times.

      a.  Provide the following connection details:

             Username

             Password

             Server name and Port

             Database name

      b.  Select Enable TLS/SSL checkbox to enable encryption for the Postgres server.

     Enable automatic backup
     Allows automatic backup for the Multicloud Object Gateway metadata database.

      a.  Select the backup frequency: Daily, Weekly, or Monthly.

      b.  Specify the number of backups to retain.

                     NOTE
                                                                                              29

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

                       NOTE

                       This option is disabled if the Use external PostgreSQL option is
                       selected.

       c.  Click Next.

4. In the Create local volume set page, provide the following information:

   a.  Enter a name for the LocalVolumeSet and the StorageClass.
       By default, the local volume set name appears for the storage class name. You can change
       the name.

   b.  Select one of the following:

           Disks on all nodes to use the available disks that match the selected filters on all nodes.

           Disks on selected nodesto use the available disks that match the selected filters only
           on selected nodes.

                       IMPORTANT

                             The flexible scaling feature is enabled only when the storage cluster
                             that you created with 3 or more nodes is spread across fewer than
                             the minimum requirement of 3 availability zones.
                             For information about flexible scaling, see knowledgebase article on
                             Scaling OpenShift Data Foundation cluster using YAML when flexible
                             scaling is enabled.

                             Flexible scaling features get enabled at the time of deployment and
                             can not be enabled or disabled later on.

                             If the nodes selected do not match the OpenShift Data Foundation
                             cluster requirement of an aggregated 30 CPUs and 72 GiB of RAM, a
                             minimal cluster is deployed.
                             For minimum starting node requirements, see the Resource
                             requirements section in the Planning guide.

   c.  From the available list of Disk Type, select SSD/NVMe.

   d.  Expand the Advanced section and set the following options:

        Volume        Block is selected by default.
        Mode

        Device        Select one or more device types from the dropdown list.
        Type

        Disk Size     Set a minimum size of 100GB for the device and maximum available size of the
                      device that needs to be included.

        Maximum       This indicates the maximum number of PVs that can be created on a node. If this
        Disks Limit   field is left empty, then PVs are created for all the available disks on the matching
                      nodes.


30

                            CHAPTER 3. DEPLOY USING LOCAL STORAGE DEVICES

   e.   Click Next.
        A pop-up to confirm the creation of LocalVolumeSet is displayed.

   f.   Click Yes to continue.

5. In the Capacity and nodes page, configure the following:

   a.   Available raw capacity is populated with the capacity value based on all the attached disks
        associated with the storage class. This takes some time to show up. The Selected nodes
        list shows the nodes based on the storage class.

   b.   In the Configure performance section, select one of the following performance profiles:

         Lean
         Use this in a resource constrained environment with minimum resources that are lower
         than the recommended. This profile minimizes resource consumption by allocating
         fewer CPUs and less memory.

         Balanced (default)
         Use this when recommended resources are available. This profile provides a balance
         between resource consumption and performance for diverse workloads.

         Performance
         Use this in an environment with sufficient resources to get the best performance. This
         profile is tailored for high performance by allocating ample memory and CPUs to
         ensure optimal execution of demanding workloads.

                        NOTE

                        You have the option to configure the performance profile even after the
                        deployment using the Configure performance option from the options
                        menu of the StorageSystems tab.

                        IMPORTANT

                        Before selecting a resource profile, make sure to check the current
                        availability of resources within the cluster. Opting for a higher resource
                        profile in a cluster with insufficient resources might lead to installation
                        failures.

         For more information about resource requirements, see Resource requirement for
         performance profiles.

   c.   Optional: Select the Taint nodes checkbox to dedicate the selected nodes for OpenShift
        Data Foundation.

   d.   Click Next.

6. Optional: In the Advanced settings page, configure the following based on your requirement:

   a.   Select Enable forceful deployment to override and completely wipe any data on disks used
        to create this storage cluster and type CONFIRM in the text box.






                        31

    Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere


 WARNING
        Any data on the disks used to create this storage cluster will be deleted.
        This cannot be undone. Use this option only when you are certain that
        the disks do not contain any data you need to preserve, such as when
        reusing disks from a previous OpenShift Data Foundation installation or
        when disks contain data from other Ceph clusters.




          b.   Select Use erasure coding to enable erasure coding instead of replication to protect the
               cluster data.
               Erasure coding is supported for Block (RBD), File (CephFS), and RADOS Gateway (RGW)
               object storage.

               For more information, see Erasure coding considerations.

      a.  Click Next.

          1.   Optional: In the Security and network page, configure the following based on your
               requirement:

      b.  To enable encryption, select Enable data encryption for block and file storage.

      c.  Select one of the following Encryption level:

               Cluster-wide encryption to encrypt the entire cluster (block and file).

               StorageClass encryption to create encrypted persistent volume (block only) using
               encryption enabled storage class.

      d.  Optional: Select the Connect to an external key management servicecheckbox. This is
          optional for cluster-wide encryption.

          i.   From the Key Management Service Providerdrop-down list, select one of the following
               providers and provide the necessary details:

               Vault

               A.    Select an Authentication Method.

                   Using Token authentication method

                              Enter a unique Connection Name, host Address of the Vault server
                              ('https://<hostname or ip>'), Port number and Token.

                              Expand Advanced Settings to enter additional settings and certificate
                              details based on your Vault configuration:

                               Enter the Key Value secret path in Backend Path that is dedicated and
                               unique to OpenShift Data Foundation.

                              Optional: Enter TLS Server Name and Vault Enterprise Namespace.

                              Upload the respective PEM encoded certificate file to provide the CA
                              Certificate, Client Certificate and Client Private Key .

    32

                                                    CHAPTER 3. DEPLOY USING LOCAL STORAGE DEVICES

                        Click Save.

                 Using Kubernetes authentication method

                        Enter a unique Vault Connection Name, host Address of the Vault server
                        ('https://<hostname or ip>'), Port number and Role name.

                        Expand Advanced Settings to enter additional settings and certificate
                        details based on your Vault configuration:

                         Enter the Key Value secret path in Backend Path that is dedicated and
                         unique to OpenShift Data Foundation.

                         Optional: Enter TLS Server Name and Authentication Path if
                         applicable.

                         Upload the respective PEM encoded certificate file to provide the CA
                         Certificate, Client Certificate and Client Private Key.

                        Click Save.

                                   NOTE

                                   In case you need to enable key rotation for Vault KMS, run
                                   the following command in the OpenShift web console after
                                   the storage cluster is created:

                                    oc patch storagecluster ocs-storagecluster -n openshift-
                                    storage --type=json -p '[{"op": "add",
                                    "path":"/spec/encryption/keyRotation/enable", "value":
                                    true}]'

Thales CipherTrust Manager (using KMIP)

            A.  Enter a unique Connection Name for the Key Management service within the
                project.

            B.  In the Address and Port sections, enter the IP of Thales CipherTrust Manager and
                the port where the KMIP interface is enabled. For example:

                 Address: 123.34.3.2

                 Port: 5696

            C.  Upload the Client Certificate, CA certificate, and Client Private Key.

            D.  If StorageClass encryption is enabled, enter the Unique Identifier to be used for
                encryption and decryption generated above.

            E.  The TLS Server field is optional and used when there is no DNS entry for the KMIP
                endpoint. For example, kmip_all_<port>.ciphertrustmanager.local.

ii. Select a Network.

e. Click Next.

   1. In the Review and create page, review the configuration details.

                                                                                               33

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere
1.   In the Review and create page, review the configuration details.
             To modify any configuration settings, click Back to go back to the previous
             configuration page.

2.   Click Create StorageSystem.

                    NOTE

                    When your deployment has five or more nodes, racks, or rooms, and when there
                    are five or more number of failure domains present in the deployment, you can
                    configure Ceph monitor counts based on the number of racks or zones. An alert
                    is displayed in the notification panel or Alert Center of the OpenShift Web
                    Console to indicate the option to increase the number of Ceph monitor counts.
                    You can use the Configure option in the alert to configure the Ceph monitor
                    counts. For more information, see Resolving low Ceph monitor count alert .

Verification

To verify the final Status of the installed storage cluster:

a.   In the OpenShift Web Console, navigate to Storage → Data Foundation → Storage
     System → ocs-storagecluster.

b.   Verify that the Status of StorageCluster is Ready and has a green tick mark next to it.

To verify if flexible scaling is enabled on your storage cluster, perform the following steps (for
arbiter mode, flexible scaling is disabled):

1.   In the OpenShift Web Console, navigate to Storage → Data Foundation → Storage
     System.

2.   From the Action menu, select Edit Storage System.

3.   In the YAML tab, search for the keys flexibleScaling in spec section and failureDomain in
     status section. If flexible scaling is true and failureDomain is set to host, the flexible
     scaling feature is enabled.

            spec:
            flexibleScaling: true
            […]
            status:
            failureDomain: host

To verify that all components for OpenShift Data Foundation are successfully installed, see
Verifying your OpenShift Data Foundation deployment .

Additional resources

Scaling Storage guide










34

                  CHAPTER 4. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT

              CHAPTER 4. VERIFYING OPENSHIFT DATA FOUNDATION
                                DEPLOYMENT

Use this section to verify that OpenShift Data Foundation is deployed correctly.

4.1. VERIFYING THE STATE OF THE PODS

Verify that all required OpenShift Data Foundation pods are in a Running state after deployment.

Procedure

   1.   Click Workloads → Podsfrom the OpenShift Web Console.

   2.   Select openshift-storage from the Project drop-down list.

              NOTE

              If the Show default projects option is disabled, use the toggle button to list all
              the default projects.

        For more information on the expected number of pods for each component and how it varies
        depending on the number of nodes, see the following table:

 1.     Set filter for Running and Completed pods to verify that the following pods are in Running and
        Completed state:

              NOTE

              The available pods depend on the cluster configuration. When the cluster is
              deployed as a standalone Multicloud Object Gateway, the rook-ceph-operator-*
              pods are not available. Similarly, when the cluster is deployed without the
              Multicloud Object Gateway, noobaa-* pods are not available.

        Pods corresponding to OpenShift Data Foundation cluster

 Component        Corresponding pods

 OpenShift Data Foundation Operator     ocs-operator-* (1 pod on any storage
                                        node)

                                        ocs-metrics-exporter-* (1 pod on any
                                        storage node)

                                        odf-operator-controller-manager-* (1
                                        pod on any storage node)

                                        odf-console-* (1 pod on any storage
                                        node)

                                        csi-addons-controller-manager-* (1
                                        pod on any storage node)



                                        35

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere


Rook-ceph Operator     rook-ceph-operator-*

                       (1 pod on any storage node)


Multicloud Object Gateway     noobaa-operator-* (1 pod on any storage
                              node)

                              noobaa-core-* (1 pod on any storage
                              node)

                              noobaa-db-pg-cluster-1 and noobaa-
                              db-pg-cluster-2 (2 instances of MCG DB
                              pod on any storage node)

                              noobaa-endpoint-* (1 pod on any storage
                              node)

                              cnpg-controller-manager-* (1 pod on
                              any storage node)


MON                    rook-ceph-mon-*

                       (3 pods distributed across storage nodes)

MGR                    rook-ceph-mgr-*

                       (2 pods on different storage nodes, one active, one
                       standby)

MDS                    rook-ceph-mds-ocs-storagecluster-
                       cephfilesystem-*

                       (2 pods distributed across storage nodes)

RGW                    rook-ceph-rgw-ocs-storagecluster-
                       cephobjectstore-* (1 pod on any storage node)










36

CHAPTER 4. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT


CSI     cephfs
           openshift-
           storage.cephfs.csi.ceph.com-
           ctrlplugin-* (2 pods distributed across
           storage nodes)

           openshift-
           storage.cephfs.csi.ceph.com-
           nodeplugin-* (1 pod on each storage
           node)

        nfs
           openshift-
           storage.nfs.csi.ceph.com-
           ctrlplugin-* (2 pods distributed across
           storage nodes)

           openshift-
           storage.nfs.csi.ceph.com-
           nodeplugin-* (1 pod on each storage
           node)

        rbd
           openshift-
           storage.rbd.csi.ceph.com-
           ctrlplugin-* (2 pods distributed across
           storage nodes)

           openshift-
           storage.rbd.csi.ceph.com-
           nodeplugin-* (1 pod on each storage
           node)


rook-ceph-crashcollector     rook-ceph-crashcollector-*

                             (1 pod on each storage node)

OSD                           rook-ceph-osd-* (1 pod for each device)

                              rook-ceph-osd-prepare-ocs-* (1 pod for
                              each device)


odf-blackbox-exporter        odf-blackbox-exporter-* (1 pod on any storage
                             node) NOTE: This does not apply to clusters without
                             storage, for example, Multicloud Object
                             Gateway‑only deployments or external mode
                             deployments.


4.2. VERIFYING THE OPENSHIFT DATA FOUNDATION CLUSTER IS
HEALTHY


                              37

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

Verify that the OpenShift Data Foundation cluster is healthy after installation.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and then click the storage
     system link from the pop up that appears.

3.   In the Status card of the Block and File tab, verify that the Storage Cluster has a green tick.

4.   In the Details card, verify that the cluster information is displayed.
     For more information on the health of the OpenShift Data Foundation cluster using the Block
     and File dashboard, see Monitoring OpenShift Data Foundation .

4.3. VERIFYING THE MULTICLOUD OBJECT GATEWAY IS HEALTHY

Verify that the Multicloud Object Gateway is healthy after deployment to confirm that the object
service and data resiliency are operational and that the NooBaa database has a recent backup.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and then click the storage
     system link from the pop up that appears.

      a.  In the Status card of the Object tab, verify that both Object Service and Data Resiliency
          have a green tick.

      b.  In the Details card, verify that the MCG information is displayed.
          For more information on the health of the OpenShift Data Foundation cluster using the
          object service dashboard, see Monitoring OpenShift Data Foundation .

                      IMPORTANT

                      To avoid data loss, it is recommended to take a backup of NooBaa DB PVC
                      regularly. If NooBaa DB fails and cannot be recovered, then you can revert to
                      the latest backed-up version. For instructions on backing up your NooBaa
                      DB, follow the steps in the knowledgebase article, Perform a One-Time
                      Backup of the Database for the Multicloud Object Gateway.

4.4. VERIFYING THAT THE SPECIFIC STORAGE CLASSES EXIST

Verify that the expected OpenShift Data Foundation storage classes are present after deployment.

Procedure

1.   Click Storage → Storage Classesfrom the left pane of the OpenShift Web Console.

2.   Verify that the following storage classes are created with the OpenShift Data Foundation
     cluster creation:

         ocs-storagecluster-ceph-rbd


38

CHAPTER 4. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT

ocs-storagecluster-cephfs

openshift-storage.noobaa.io

ocs-storagecluster-ceph-rgw










39

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

    CHAPTER 5. DEPLOY STANDALONE MULTICLOUD OBJECT
        GATEWAY

  Deploying only the Multicloud Object Gateway component with the OpenShift Data Foundation
  provides the flexibility in deployment and helps to reduce the resource consumption. You can deploy
  the Multicloud Object Gateway component either using dynamic storage devices or using the local
  storage devices.










    40

CHAPTER 6. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING DYNAMIC STORAGE DEVICES

                    CHAPTER 6. DEPLOY STANDALONE MULTICLOUD OBJECT
                        GATEWAY USING DYNAMIC STORAGE DEVICES

Use this section to deploy only the standalone Multicloud Object Gateway component, which involves
the following steps:

     Installing Red Hat OpenShift Data Foundation Operator

     Creating standalone Multicloud Object Gateway

After deploying the component, you can create and manage buckets using MCG object browser. For
more information, see Creating and managing buckets using MCG object browser .

6.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.

Prerequisites

     Access to an OpenShift Container Platform cluster using an account with cluster-admin and
     operator installation permissions.

     You must have at least three worker or infrastructure nodes in the Red Hat OpenShift
     Container Platform cluster.

     For additional resource requirements, see the Planning your deployment guide.

         IMPORTANT

                    When you need to override the cluster-wide default node selector for OpenShift
                    Data Foundation, you can use the following command to specify a blank node
                    selector for the openshift-storage namespace (create openshift-storage
                    namespace in this case):

                    $ oc annotate namespace openshift-storage openshift.io/node-selector=

                    Taint a node as infra to ensure only Red Hat OpenShift Data Foundation
                    resources are scheduled on that node. This helps you save on subscription costs.
                    For more information, see the How to use dedicated worker nodes for Red Hat
                    OpenShift Data Foundation section in the Managing and Allocating Storage
                    Resources guide.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Ecosystem → Software Catalog.

3.   Scroll or type OpenShift Data Foundation into the Filter by keyword box to find the
     OpenShift Data Foundation Operator.



                    41

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

     4.  Click Install.

     5.  Set the following options on the Install Operator page:

          a.   Update Channel as stable-4.22.

          b.   Installation Mode as A specific namespace on the cluster.

          c.   Installed Namespace as Operator recommended namespace openshift-storage. If
               Namespace openshift-storage does not exist, it is created during the operator installation.

          d.   Select Approval Strategy as Automatic or Manual.
               If you select Automatic updates, then the Operator Lifecycle Manager (OLM)
               automatically upgrades the running instance of your Operator without any intervention.

               If you select Manual updates, then the OLM creates an update request. As a cluster
               administrator, you must then manually approve that update request to update the Operator
               to a newer version.

          e.   Ensure that the Enable option is selected for the Console plugin.

          f.   Click Install.

Verification

         After the operator is successfully installed, the web console automatically reloads to apply the
         changes. During this process, a temporary error message might appear on the page and this is
         expected and disappears after the refresh completes.

         In the Web Console:

               Navigate to Installed Operators and verify that the OpenShift Data Foundation Operator
               shows a green tick indicating successful installation.

               Navigate to Storage and verify if the Data Foundation dashboard is available.

6.2. CREATING A STANDALONE MULTICLOUD OBJECT GATEWAY

You can create only the standalone Multicloud Object Gateway (MCG) component while deploying
OpenShift Data Foundation. After you create the MCG component, you can create and manage buckets
using the MCG object browser. For more information, see Creating and managing buckets using MCG
object browser.

Prerequisites

         Ensure that the OpenShift Data Foundation Operator is installed.

         The Multicloud Object Gateway database uses ReadWriteOncePod (RWOP) access mode to
         safeguard against node failures. If the underlying CSI driver does not support RWOP, the access
         mode must be manually changed to ReadWriteOnce (RWO) before deployment. For more
         information, see the knowledgebase article How to change NooBaa DB access mode from
         RWOP to RWO in standalone MCG deployments?.

Procedure

  1.     In the OpenShift Web Console, navigate to Storage → StorageCluster.

42

CHAPTER 6. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING DYNAMIC STORAGE DEVICES

2.  On the Welcome page, click Configure Data Foundation.

3.  Select Setup Multicloud Object Gateway.

4.  In the Backing storage page, select the following:

     a.  Select the Use an existing StorageClass option.

     b.  Click Next.

5.  In the Advanced Settings page, you can configure the following optional items:

         Use external PostgreSQL [Technology preview]
         Enables the use of an external PostgreSQL instance. This provides a high-availability
         solution for the Multicloud Object Gateway, where the PostgreSQL pod is a single point of
         failure.

                    IMPORTANT

                    OpenShift Data Foundation ships PostgreSQL images maintained by Red
                    Hat, which are used to store metadata for the Multicloud Object Gateway.
                    This PostgreSQL usage is at the application level.

                    As a result, OpenShift Data Foundation does not perform database-level
                    optimizations or in-depth insights.

                    If you have well-maintained and optimized PostgreSQL instance, then it is
                    recommended to use it. OpenShift Data Foundation supports external
                    PostgreSQL instances.

                    Any PostgreSQL-related issues requiring code changes or deep technical
                    analysis may need to be addressed upstream. This could result in longer
                    resolution times.

          a.  Provide the following connection details:

                 Username

                 Password

                 Server name and Port

                 Database name

          b.  Select Enable TLS/SSL checkbox to enable encryption for the Postgres server.

         Enable automatic backup
         Allows automatic backup for the NooBaa database.

          a.  Select the backup frequency: Daily, Weekly, or Monthly.

          b.  Specify the number of backups to retain.

          c.  Select a VolumeSnapshotClass to add a volume snapshot.



                    NOTE
                        43

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

                     NOTE

                     This option is disabled if the Use external PostgreSQL option is
                     selected.

       d. Click Next.

6. In the Security page, select any of the following optional items:

   a.  Select the Connect to an external key management servicecheckbox. This is optional for
       cluster-wide encryption.

   b.  From the Key Management Service Providerdrop-down list, either select Vault or Thales
       CipherTrust Manager (using KMIP). If you selected Vault, go to the next step. If you
       selected Thales CipherTrust Manager (using KMIP), go to step iii.

   c.  Select an Authentication Method.

       Using Token authentication method

          Enter a unique Connection Name, host Address of the Vault server
          ('https://<hostname or ip>'), Port number and Token.

          Expand Advanced Settings to enter additional settings and certificate details
          based on your Vault configuration:

           Enter the Key Value secret path in Backend Path that is dedicated and unique
           to OpenShift Data Foundation.

           Optional: Enter TLS Server Name, Authentication Path, and Vault Enterprise
           Namespace.

           Upload the respective PEM encoded certificate file to provide the CA
           Certificate, Client Certificate and Client Private Key .

           Click Save and skip to step iv.

      Using Kubernetes authentication method

          Enter a unique Vault Connection Name, host Address of the Vault server
          ('https://<hostname or ip>'), Port number and Role name.

          Expand Advanced Settings to enter additional settings and certificate details
          based on your Vault configuration:

           Enter the Key Value secret path in Backend Path that is dedicated and unique
           to OpenShift Data Foundation.

           Optional: Enter TLS Server Name, Authentication Path, and Vault Enterprise
           Namespace if applicable.

           Upload the respective PEM encoded certificate file to provide the CA
           Certificate, Client Certificate and Client Private Key .

           Click Save and skip to step iv.



   d. To use Thales CipherTrust Manager (using KMIP) as the KMS provider, follow the steps
44

CHAPTER 6. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING DYNAMIC STORAGE DEVICES

      d.     To use Thales CipherTrust Manager (using KMIP) as the KMS provider, follow the steps
             below:

               i.   Enter a unique Connection Name for the Key Management service within the project.

              ii.   In the Address and Port sections, enter the IP of Thales CipherTrust Manager and the
                    port where the KMIP interface is enabled. For example:

                     Address: 123.34.3.2

                     Port: 5696

             iii.   Upload the Client Certificate, CA certificate, and Client Private Key.

              iv.   If StorageClass encryption is enabled, enter the Unique Identifier to be used for
                    encryption and decryption generated above.

               v.   The TLS Server field is optional and used when there is no DNS entry for the KMIP
                    endpoint. For example, kmip_all_<port>.ciphertrustmanager.local.

        e.   Select a Network.

        f.   Click Next.

 7.  In the Review and create page, review the configuration details:
     To modify any configuration settings, click Back.

 8.  Click Create StorageSystem.

Verification

Verifying that the OpenShift Data Foundation cluster is healthy

      1.    In the OpenShift Web Console, click Storage → Data Foundation.

     2.     In the Status card of the Overview tab, click Storage System and then click the storage
            system link from the pop up that appears.

             a.    In the Status card of the Object tab, verify that both Object Service and Data Resiliency
                   have a green tick.

             b.    In the Details card, verify that the MCG information is displayed.

Verifying the state of the pods

      1.    Click Workloads → Pods from the OpenShift Web Console.

     2.     Select openshift-storage from the Project drop-down list and verify that the following
            pods are in Running state.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to list
                        all the default projects.




                        45

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere


Component        Corresponding pods

OpenShift Data     ocs-operator-* (1 pod on any storage node)
Foundation Operator
                   ocs-metrics-exporter-* (1 pod on any storage node)

                   odf-operator-controller-manager-* (1 pod on any storage
                   node)

                   odf-console-* (1 pod on any storage node)

                   csi-addons-controller-manager-* (1 pod on any storage
                   node)


Rook-ceph Operator   rook-ceph-operator-*

                     (1 pod on any storage node)

Multicloud Object     noobaa-operator-* (1 pod on any storage node)
Gateway
                      noobaa-core-* (1 pod on any storage node)

                      noobaa-db-pg-cluster-1 and noobaa-db-pg-cluster-2 (2
                      instances of MCG DB pod on any storage node)

                      noobaa-endpoint-* (1 pod on any storage node)

                      cnpg-controller-manager-* (1 pod on any storage node)










46

 CHAPTER 7. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING LOCAL STORAGE DEVICES

 CHAPTER 7. DEPLOY STANDALONE MULTICLOUD OBJECT
            GATEWAY USING LOCAL STORAGE DEVICES

Use this section to deploy only the standalone Multicloud Object Gateway component, which involves
the following steps:

      Installing the Local Storage Operator

      Installing Red Hat OpenShift Data Foundation Operator

      Creating standalone Multicloud Object Gateway

After deploying the MCG component, you can create and manage buckets using MCG object browser.
For more information, see Creating and managing buckets using MCG object browser .

7.1. INSTALLING LOCAL STORAGE OPERATOR

Install the Local Storage Operator from the Software Catalog before creating Red Hat OpenShift Data
Foundation clusters on local storage devices.

Procedure

 1.   Log in to the OpenShift Web Console.

 2.   Click Ecosystem → Software Catalog.

 3.   Type local storage in the Filter by keyword box to find the Local Storage Operator from the
      list of operators, and click on it.

 4.   Set the following options on the Install Operator page:

       a.  Update channel as stable.

       b.  Installation mode as A specific namespace on the cluster.

       c.  Installed Namespace as Operator recommended namespace openshift-local-storage.

       d.  Update approval as Automatic.

 5.   Click Install.

Verification

      Verify that the Local Storage Operator shows a green tick indicating successful installation.

7.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.

Prerequisites


 Access to an OpenShift Container Platform cluster using an account with cluster-admin and
               47

  Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

       Access to an OpenShift Container Platform cluster using an account with cluster-admin and
       operator installation permissions.

       You must have at least three worker or infrastructure nodes in the Red Hat OpenShift
       Container Platform cluster.

       For additional resource requirements, see the Planning your deployment guide.

              IMPORTANT

                     When you need to override the cluster-wide default node selector for OpenShift
                     Data Foundation, you can use the following command to specify a blank node
                     selector for the openshift-storage namespace (create openshift-storage
                     namespace in this case):

                         $ oc annotate namespace openshift-storage openshift.io/node-selector=

                     Taint a node as infra to ensure only Red Hat OpenShift Data Foundation
                     resources are scheduled on that node. This helps you save on subscription costs.
                     For more information, see the How to use dedicated worker nodes for Red Hat
                     OpenShift Data Foundation section in the Managing and Allocating Storage
                     Resources guide.

Procedure

  1.   Log in to the OpenShift Web Console.

  2.   Click Ecosystem → Software Catalog.

  3.   Scroll or type OpenShift Data Foundation into the Filter by keyword box to find the
       OpenShift Data Foundation Operator.

  4.   Click Install.

  5.   Set the following options on the Install Operator page:

       a.   Update Channel as stable-4.22.

       b.   Installation Mode as A specific namespace on the cluster.

       c.   Installed Namespace as Operator recommended namespace openshift-storage. If
            Namespace openshift-storage does not exist, it is created during the operator installation.

       d.   Select Approval Strategy as Automatic or Manual.
            If you select Automatic updates, then the Operator Lifecycle Manager (OLM)
            automatically upgrades the running instance of your Operator without any intervention.

            If you select Manual updates, then the OLM creates an update request. As a cluster
            administrator, you must then manually approve that update request to update the Operator
            to a newer version.

       e.   Ensure that the Enable option is selected for the Console plugin.

       f.   Click Install.

  Verification

  48

                   CHAPTER 7. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING LOCAL STORAGE DEVICES
Verification

     After the operator is successfully installed, the web console automatically reloads to apply the
     changes. During this process, a temporary error message might appear on the page and this is
     expected and disappears after the refresh completes.

     In the Web Console:

          Navigate to Installed Operators and verify that the OpenShift Data Foundation Operator
          shows a green tick indicating successful installation.

          Navigate to Storage and verify if the Data Foundation dashboard is available.

7.3. CREATING A STANDALONE MULTICLOUD OBJECT GATEWAY

You can create only the standalone Multicloud Object Gateway (MCG) component while deploying
OpenShift Data Foundation. After you create the MCG component, you can create and manage buckets
using the MCG object browser. For more information, see Creating and managing buckets using MCG
object browser.

Prerequisites

     Ensure that the OpenShift Data Foundation Operator is installed.

     The Multicloud Object Gateway database uses ReadWriteOncePod (RWOP) access mode to
     safeguard against node failures. If the underlying CSI driver does not support RWOP, the access
     mode must be manually changed to ReadWriteOnce (RWO) before deployment. For more
     information, see the knowledgebase article How to change NooBaa DB access mode from
     RWOP to RWO in standalone MCG deployments?.

Procedure

1.   In the OpenShift Web Console, navigate to Storage → StorageCluster.

2.   On the Welcome page, click Configure Data Foundation.

3.   Select Setup Multicloud Object Gateway.

4.   In the Backing storage page, select the following:

      a.  Select the Create a new StorageClass using the local storage devicesoption.

      b.  Click Next.

                     NOTE

                     You are prompted to install the Local Storage Operator if it is not already
                     installed. Click Install, and follow the procedure as described in Installing
                     Local Storage Operator.

5.   In the Advanced Settings page, you can configure the following optional items:

          Use external PostgreSQL [Technology preview]
          Enables the use of an external PostgreSQL instance. This provides a high-availability
          solution for the Multicloud Object Gateway, where the PostgreSQL pod is a single point of
          failure.

                     IMPORTANT
                                                                                                   49

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

                IMPORTANT

                OpenShift Data Foundation ships PostgreSQL images maintained by Red
                Hat, which are used to store metadata for the Multicloud Object Gateway.
                This PostgreSQL usage is at the application level.

                As a result, OpenShift Data Foundation does not perform database-level
                optimizations or in-depth insights.

                If you have well-maintained and optimized PostgreSQL instance, then it is
                recommended to use it. OpenShift Data Foundation supports external
                PostgreSQL instances.

                Any PostgreSQL-related issues requiring code changes or deep technical
                analysis may need to be addressed upstream. This could result in longer
                resolution times.

        a.  Provide the following connection details:

             Username

             Password

             Server name and Port

             Database name

        b.  Select Enable TLS/SSL checkbox to enable encryption for the Postgres server.

       Enable automatic backup
       Allows automatic backup for the NooBaa database.

        a.  Select the backup frequency: Daily, Weekly, or Monthly.

        b.  Specify the number of backups to retain.

        c.  Select a VolumeSnapshotClass to add a volume snapshot.

                       NOTE

                       This option is disabled if the Use external PostgreSQL option is
                       selected.

        d.  Click Next.

6. In the Create local volume set page, provide the following information:

   a.  Enter a name for the LocalVolumeSet and the StorageClass.
       By default, the local volume set name appears for the storage class name. You can change
       the name.

   b.  Choose one of the following:

            Disks on all nodes
            Uses the available disks that match the selected filters on all the nodes.



50

CHAPTER 7. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING LOCAL STORAGE DEVICES

      Disks on selected nodes
      Uses the available disks that match the selected filters only on the selected nodes.

c.  From the available list of Disk Type, select SSD/NVMe.

d.  Expand the Advanced section and set the following options:

     Volume        Filesystem is selected by default. Always ensure that the Filesystem is selected
     Mode          for Volume Mode.

     Device        Select one or more device types from the dropdown list.
     Type

     Disk Size     Set a minimum size of 100GB for the device and maximum available size of the
                   device that needs to be included.

     Maximum       This indicates the maximum number of PVs that can be created on a node. If this
     Disks Limit   field is left empty, then PVs are created for all the available disks on the matching
                   nodes.


     e.   Click Next.
          A pop-up to confirm the creation of LocalVolumeSet is displayed.

     f.   Click Yes to continue.

7.  In the Capacity and nodes page, configure the following:

     a.   Available raw capacity is populated with the capacity value based on all the attached disks
          associated with the storage class. This takes some time to show up. The Selected nodes
          list shows the nodes based on the storage class.

     b.   Click Next.

8.  In the Security page, select any of the following optional items:

     a.   Select the Connect to an external key management servicecheckbox. This is optional for
          cluster-wide encryption.

     b.   From the Key Management Service Providerdrop-down list, either select Vault or Thales
          CipherTrust Manager (using KMIP). If you selected Vault, go to the next step. If you
          selected Thales CipherTrust Manager (using KMIP), go to step iii.

     c.   Select an Authentication Method.

          Using Token authentication method

           Enter a unique Connection Name, host Address of the Vault server
           ('https://<hostname or ip>'), Port number and Token.

           Expand Advanced Settings to enter additional settings and certificate details
           based on your Vault configuration:

                     Enter the Key Value secret path in Backend Path that is dedicated and unique
                     to OpenShift Data Foundation.


           Optional: Enter TLS Server Name, Authentication Path, and Vault Enterprise
                         51

 Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

                      Optional: Enter TLS Server Name, Authentication Path, and Vault Enterprise
                      Namespace.

                      Upload the respective PEM encoded certificate file to provide the CA
                      Certificate, Client Certificate and Client Private Key .

                      Click Save and skip to step iv.

          Using Kubernetes authentication method

                   Enter a unique Vault Connection Name, host Address of the Vault server
                   ('https://<hostname or ip>'), Port number and Role name.

                   Expand Advanced Settings to enter additional settings and certificate details
                   based on your Vault configuration:

                      Enter the Key Value secret path in Backend Path that is dedicated and unique
                      to OpenShift Data Foundation.

                      Optional: Enter TLS Server Name, Authentication Path, and Vault Enterprise
                      Namespace if applicable.

                      Upload the respective PEM encoded certificate file to provide the CA
                      Certificate, Client Certificate and Client Private Key .

                      Click Save and skip to step iv.

      d.   To use Thales CipherTrust Manager (using KMIP) as the KMS provider, follow the steps
           below:

             i.   Enter a unique Connection Name for the Key Management service within the project.

            ii.   In the Address and Port sections, enter the IP of Thales CipherTrust Manager and the
                  port where the KMIP interface is enabled. For example:

                    Address: 123.34.3.2

                    Port: 5696

           iii.   Upload the Client Certificate, CA certificate, and Client Private Key.

            iv.   If StorageClass encryption is enabled, enter the Unique Identifier to be used for
                  encryption and decryption generated above.

             v.   The TLS Server field is optional and used when there is no DNS entry for the KMIP
                  endpoint. For example, kmip_all_<port>.ciphertrustmanager.local.

      e.   Select a Network.

      f.   Click Next.

 9.  In the Review and create page, review the configuration details:
     To modify any configuration settings, click Back.

10.  Click Create StorageSystem.

 Verification

 52

CHAPTER 7. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY USING LOCAL STORAGE DEVICES
Verification

Verifying that the OpenShift Data Foundation cluster is healthy

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and then click the storage
     system link from the pop up that appears.

            a.  In the Status card of the Object tab, verify that both Object Service and Data Resiliency
                have a green tick.

            b.  In the Details card, verify that the MCG information is displayed.

Verifying the state of the pods

1.   Click Workloads → Pods from the OpenShift Web Console.

2.   Select openshift-storage from the Project drop-down list and verify that the following
     pods are in Running state.

                 NOTE

                 If the Show default projects option is disabled, use the toggle button to list
                 all the default projects.


Component          Corresponding pods

OpenShift Data     ocs-operator-* (1 pod on any storage node)
Foundation Operator
                   ocs-metrics-exporter-* (1 pod on any storage node)

                   odf-operator-controller-manager-* (1 pod on any storage
                   node)

                   odf-console-* (1 pod on any storage node)

                   csi-addons-controller-manager-* (1 pod on any storage
                   node)


            Rook-ceph Operator   rook-ceph-operator-*

                                 (1 pod on any storage node)










                                 53

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere


Component        Corresponding pods

Multicloud Object     noobaa-operator-* (1 pod on any storage node)
Gateway
                      noobaa-core-* (1 pod on any storage node)

                      noobaa-db-pg-cluster-1 and noobaa-db-pg-cluster-2 (2
                      instances of MCG DB pod on any storage node)

                      noobaa-endpoint-* (1 pod on any storage node)

                      cnpg-controller-manager-* (1 pod on any storage node)










54

              CHAPTER 8. VIEW OPENSHIFT DATA FOUNDATION TOPOLOGY

      CHAPTER 8. VIEW OPENSHIFT DATA FOUNDATION
              TOPOLOGY

The topology provides a visual representation of the OpenShift Data Foundation storage cluster at
multiple levels of abstraction and enables interaction with each layer. It illustrates how the cluster
components are organized and related, while offering visibility into Object Storage Devices (OSDs),
including their count, type, and status.

Procedure

1.   On the OpenShift Web Console, navigate to Storage → Data Foundation → Topology.
     The view shows the storage cluster and the topology groupings inside it. You can see the nodes
     depicted by circular entities within the topology groupings, which are indicated by dotted lines.
     The topology groupings are based on the cluster’s topology configuration (for example, racks in
     bare metal or vSphere clusters, or zones in cloud-based clusters). The label of each item or
     resource contains basic information such as status and health or indication for alerts.

2.   Choose a node to view node details on the right-hand panel. You can also access resources or
     deployments within a node by clicking on the search/preview decorator icon.
     The node details panel displays information about the OSDs running on that node, including:

          Number of OSDs on the node

          OSD disk type (for example, SSD, HDD)

          OSD status (for example, Running, Error, CrashLoopBackOff)

3.   To view deployment details

      a.  Click the preview decorator on a node. A modal window appears above the node that
          displays all of the deployments associated with that node along with their statuses.

      b.  Click the Back to main viewbutton in the model’s upper left corner to close and return to
          the previous view.

      c.  Select a specific deployment to see more information about it. All relevant data is shown in
          the side panel.

4.   Click the Resources tab to view the pods information. This tab provides a deeper
     understanding of the problems and offers granularity that aids in better troubleshooting.

5.   Click the OSD information tab to view detailed information about the OSDs. This tab provides
     more comprehensive OSD-specific details, including:

          OSD pod name and status

          Disk device associated with each OSD

          Device class (for example, SSD, HDD)

          Resource utilization metrics

6.   Click the pod links to view the pod information page on OpenShift Container Platform. The link
     opens in a new window.




          55

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on VMware vSphere

CHAPTER 9. UNINSTALLING OPENSHIFT DATA FOUNDATION

9.1. UNINSTALLING OPENSHIFT DATA FOUNDATION IN INTERNAL
MODE

Uninstalling OpenShift Data Foundation in internal mode permanently removes the storage cluster, all
associated PVCs, and the operator.

Procedure

1.   To uninstall OpenShift Data Foundation in internal mode, refer to the knowledgebase article on
     Uninstalling OpenShift Data Foundation.










56
