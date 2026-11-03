---
title: "deploying openshift data foundation using ibm power"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/deploying_openshift_data_foundation_using_ibm_power"
source_file: "deploying_openshift_data_foundation_using_ibm_power.pdf"
version: "4.22"
chapter: "deploying_openshift_data_foundation_using_ibm_power"
---

                   Red Hat OpenShift Data Foundation
                                  4.22

               Deploying OpenShift Data Foundation using
                               IBM Power

Instructions on deploying Red Hat OpenShift Data Foundation on IBM Power










    Last Updated: 2026-10-23



Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data
Foundation using IBM Power

Instructions on deploying Red Hat OpenShift Data Foundation on IBM Power

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

    Read this document for instructions about how to install Red Hat OpenShift Data Foundation to use
    local storage on IBM Power.

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
 1.1. REQUIREMENTS FOR INSTALLING OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES
                                                                                      6

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES               8
 2.1. INSTALLING LOCAL STORAGE OPERATOR                                               8
 2.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR                           8
 2.3. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE TOKEN AUTHENTICATION METHOD
                                                                                     10
 2.4. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE KUBERNETES AUTHENTICATION
 METHOD                                                                              11
 2.4.1. Enabling and disabling key rotation when using KMS                           14
  2.4.1.1. Enforcing precedence for key rotation                                     14
  2.4.1.2. Enabling key rotation                                                     14
   2.4.1.2.1. Annotating a namespace for key rotation                                14
   2.4.1.2.2. Annotating a storage class for key rotation                            15
   2.4.1.2.3. Annotating a PersistentVolumeClaim for key rotation                    16
  2.4.1.3. Disabling key rotation                                                    16
 2.5. FINDING AVAILABLE STORAGE DEVICES                                              17
 2.6. CREATING OPENSHIFT DATA FOUNDATION CLUSTER ON IBM POWER                        19

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT FOR INTERNAL MODE          26
 3.1. VERIFYING THE STATE OF THE PODS                                                26
 3.2. VERIFYING THE OPENSHIFT DATA FOUNDATION CLUSTER IS HEALTHY                     29
 3.3. VERIFYING THE MULTICLOUD OBJECT GATEWAY IS HEALTHY                             29
 3.4. VERIFYING THAT THE SPECIFIC STORAGE CLASSES EXIST                              30

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY                               31
 4.1. INSTALLING LOCAL STORAGE OPERATOR                                              31
 4.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR                          32
 4.3. CREATING STANDALONE MULTICLOUD OBJECT GATEWAY ON IBM POWER                     33

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. VIEW OPENSHIFT DATA FOUNDATION TOPOLOGY                                   38

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 6. UNINSTALLING OPENSHIFT DATA FOUNDATION                                    39
 6.1. UNINSTALLING OPENSHIFT DATA FOUNDATION IN INTERNAL MODE                        39










   1

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power










2

PREFACE

PREFACE

This document describes how to deploy Red Hat OpenShift Data Foundation using IBM Power.










3

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

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
Platform (RHOCP) IBM Power clusters in connected or disconnected environments along with out-of-
the-box support for proxy environments.

Both internal and external OpenShift Data Foundation clusters are supported on IBM Power. See the
Planning your deployment guide and the Preparing to deploy OpenShift Data Foundation chapter for
more information about deployment requirements.

To deploy OpenShift Data Foundation, follow the appropriate deployment process based on your
requirement:

 Internal-Attached Devices mode — Deploy using local storage devices, or Deploy standalone
 Multicloud Object Gateway component

 External mode using Red Hat Ceph Storage

Additional resources

 Planning your deployment

 Preparing to deploy OpenShift Data Foundation

 Deploy standalone Multicloud Object Gateway component

 External mode










 5

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

 CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA
           FOUNDATION

When you deploy OpenShift Data Foundation on OpenShift Container Platform using the local storage
devices provided by IBM Power, you can create internal cluster resources. This approach internally
provisions base services so that all the applications can access additional storage classes..

Before you begin the deployment of Red Hat OpenShift Data Foundation using a local storage, ensure
that you meet the resource requirements. See Requirements for installing OpenShift Data Foundation
using local storage devices.

      Optional: If you want to enable cluster-wide encryption using the external Key Management
      System (KMS) follow these steps:

       Ensure that you have a valid Red Hat OpenShift Data Foundation Advanced subscription.
       To know how subscriptions for OpenShift Data Foundation work, see knowledgebase article
       on OpenShift Data Foundation subscriptions.

       When the Token authentication method is selected for encryption then refer to Enabling
       cluster-wide encryption with the Token authentication using KMS.

       When the Kubernetes authentication method is selected for encryption then refer to
       Enabling cluster-wide encryption with the Kubernetes authentication using KMS .

       Ensure that you are using signed certificates on your Vault servers.

After you have addressed the above, perform the following steps:

 1.   Install Local Storage Operator .

 2.   Install the Red Hat OpenShift Data Foundation Operator .

 3.   Find available storage devices .

 4.   Create OpenShift Data Foundation cluster on IBM Power .

1.1. REQUIREMENTS FOR INSTALLING OPENSHIFT DATA
FOUNDATION USING LOCAL STORAGE DEVICES

Before installing OpenShift Data Foundation on IBM Power Systems using local storage devices, your
cluster must meet specific node, disk, and resource requirements.

Node requirements

      The cluster must consist of at least three OpenShift Container Platform worker nodes in the
      cluster with locally attached storage devices on each of them.

       Each of the three selected nodes must have at least one raw block device available to be
       used by OpenShift Data Foundation.

       The devices to be used must be empty, that is, there should be no persistent volumes (PVs),
       volume groups (VGs), or local volumes (LVs) remaining on the disks.

      You must have a minimum of three labeled worker nodes.


      Each node that has local storage devices to be used by OpenShift Data Foundation must
6

        CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA FOUNDATION

        Each node that has local storage devices to be used by OpenShift Data Foundation must
        have a specific label to deploy OpenShift Data Foundation pods. To label the nodes, use the
        following command:

        $ oc label nodes <NodeNames> cluster.ocs.openshift.io/openshift-storage=''

For more information, see the Resource requirements section in the Planning guide.

Disconnected environment requirements

For the list of packages to be mirrored for installing OpenShift Data Foundation in disconnected
environments, see the Disconnected environment section in the Planning your deployment guide.

Disaster recovery requirements

Disaster Recovery features supported by Red Hat OpenShift Data Foundation require all of the
following prerequisites to successfully implement a disaster recovery solution:

        A valid Red Hat OpenShift Data Foundation Advanced subscription.

        A valid Red Hat Advanced Cluster Management (RHACM) for Kubernetes subscription.

To know in detail how subscriptions for OpenShift Data Foundation work, see the knowledgebase article
on OpenShift Data Foundation subscriptions.

For detailed disaster recovery solution requirements, see the Configuring OpenShift Data Foundation
Disaster Recovery for OpenShift Workloads guide, and the Requirements and recommendations section
of the Install guide in Red Hat Advanced Cluster Management for Kubernetes documentation.

Additional resources

        Planning guide - Disconnected environment

        Planning guide - Resource requirements

        Knowledgebase article on OpenShift Data Foundation subscriptions

        Configuring OpenShift Data Foundation Disaster Recovery for OpenShift Workloads

        Red Hat Advanced Cluster Management for Kubernetes - Requirements and recommendations










    7

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING
              LOCAL STORAGE DEVICES

Use this section to deploy OpenShift Data Foundation on IBM Power infrastructure where OpenShift
Container Platform is already installed.

Also, it is possible to deploy only the Multicloud Object Gateway (MCG) component with OpenShift
Data Foundation. For more information, see Deploy standalone Multicloud Object Gateway .

Perform the following steps to deploy OpenShift Data Foundation:

1.   Install the Local Storage Operator .

2.   Install the Red Hat OpenShift Data Foundation Operator .

3.   Find available storage devices .

4.   Create an OpenShift Data Foundation cluster on IBM Power .

2.1. INSTALLING LOCAL STORAGE OPERATOR

Use this procedure to install the Local Storage Operator from the Software Catalog before creating
OpenShift Data Foundation clusters on local storage devices.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Ecosystem → Software Catalog.

3.   Type local storage in the Filter by keyword… box to find the Local Storage Operator from
     the list of operators and click on it.

4.   Set the following options on the Install Operator page:

      a.  Update channel as stable.

      b.  Installation Mode as A specific namespace on the cluster.

      c.  Installed Namespace as Operator recommended namespace openshift-local-storage.

      d.  Approval Strategy as Automatic.

5.   Click Install.

Verification

     Verify that the Local Storage Operator shows a green tick indicating successful installation.

2.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.


8

             CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

For information about the hardware and software requirements, see Planning your deployment.

Prerequisites

     Access to an OpenShift Container Platform cluster using an account with cluster-admin and
     Operator installation permissions.

     You must have at least three worker nodes in the Red Hat OpenShift Container Platform
     cluster.

             IMPORTANT

                   When you need to override the cluster-wide default node selector for OpenShift
                   Data Foundation, you can use the following command in the command line
                   interface to specify a blank node selector for the openshift-storage namespace
                   (create openshift-storage namespace in this case):

             $ oc annotate namespace openshift-storage openshift.io/node-selector=

                   Taint a node as infra to ensure only Red Hat OpenShift Data Foundation
                   resources are scheduled on that node. This helps you save on subscription costs.
                   For more information, see How to use dedicated worker nodes for Red Hat
                   OpenShift Data Foundation chapter in the Managing and Allocating Storage
                   Resources guide.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Ecosystem → Software Catalog.

3.   Scroll or type OpenShift Data Foundation into the Filter by keyword box to find the
     OpenShift Data Foundation Operator.

4.   Click Install.

5.   Set the following options on the Install Operator page:

       a.  Update Channel as stable-4.22.

       b.  Installation Mode as A specific namespace on the cluster.

       c.  Installed Namespace as Operator recommended namespace openshift-storage. If
           Namespace openshift-storage does not exist, it is created during the operator installation.

6.   Select Approval Strategy as Automatic or Manual.
     If you select Automatic updates, then the Operator Lifecycle Manager (OLM) automatically
     upgrades the running instance of your Operator without any intervention.

     If you select Manual updates, then the OLM creates an update request. As a cluster
     administrator, you must then manually approve that update request to update the Operator to a
     newer version.

7.   Ensure that the Enable option is selected for the Console plugin.



                   9

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

8.   Click Install.

Verification

     Verify that the OpenShift Data Foundation Operator shows a green tick indicating successful
     installation.

     After the operator is successfully installed, the web console automatically reloads to apply the
     changes. During this process, a temporary error message might appear on the page and this is
     expected and disappears after the refresh completes.

            In the Web Console, navigate to Storage and verify if Data Foundation is available.

2.3. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE
TOKEN AUTHENTICATION METHOD

You can enable the key value backend path and policy in the vault for token authentication.

Prerequisites

     Administrator access to the vault.

     A valid Red Hat OpenShift Data Foundation Advanced subscription. For more information, see
     the knowledgebase article on OpenShift Data Foundation subscriptions .

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



10

             CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

2.4. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE
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

 1.   Create a service account:

         $ oc -n openshift-storage create serviceaccount <serviceaccount_name>

      where, <serviceaccount_name> specifies the name of the service account.

      For example:

         $ oc -n openshift-storage create serviceaccount odf-vault-auth

 2.   Create clusterrolebindings and clusterroles:

         $ oc -n openshift-storage create clusterrolebinding vault-tokenreview-binding --
         clusterrole=system:auth-delegator --serviceaccount=openshift-
         storage:_<serviceaccount_name>_

      For example:

         $ oc -n openshift-storage create clusterrolebinding vault-tokenreview-binding --
         clusterrole=system:auth-delegator --serviceaccount=openshift-storage:odf-vault-auth

 3.   Create a secret for the serviceaccount token and CA certificate.

         $ cat <<EOF | oc create -f -
         apiVersion: v1
         kind: Secret
         metadata:
          name: odf-vault-auth-token
          namespace: openshift-storage
          annotations:
          kubernetes.io/service-account.name: <serviceaccount_name>




          11

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

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

    Or, for Vault version 1.21.0 and higher:

     $ oc proxy &
     $ proxy_pid=$!
     $ ISSUER=$(echo '{"apiVersion": "authentication.k8s.io/v1", "kind": "TokenRequest"}' \   |
     kubectl create f --raw /api/v1/namespaces/default/serviceaccounts/default/token \    | jq -r
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



12

            CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

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


                         13

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

              policies=odf \
              ttl=1440h

2.4.1. Enabling and disabling key rotation when using KMS

Security common practices require periodic encryption of key rotation. You can enable or disable key
rotation when using KMS.

2.4.1.1. Enforcing precedence for key rotation

Configure key rotation precedence to control whether the rotation schedule is read from the PVC
annotation or the storage class in OpenShift Data Foundation.

In key rotation, precedence refers to the order in which the system checks for scheduled annotations. In
OpenShift Data Foundation, the default precedence is set to storage class (recommended). This means
the system reads annotations only from the storage class.

However, if you want the system to check the persistent volume claim (PVC) first and then fall back to
the storage class, you can configure this behavior by setting the schedule-precedence to PVC in the
CSI-addons ConfigMap.

Procedure

1.   Define the ConfigMap for csi-addons to set PVC precedence:

         apiVersion: v1
         kind: ConfigMap
         metadata:
          name: csi-addons-config
          namespace: openshift-storage
         data:
          "schedule-precedence": "pvc"

2.   Restart the csi-addons operator pod for the changes to take effect:

         oc delete po -n openshift-storage -l "app.kubernetes.io/name=csi-addons"

2.4.1.2. Enabling key rotation

Enable key rotation to automatically re-encrypt persistent volume data at a scheduled interval using the
CSI Addons framework in OpenShift Data Foundation.

To enable key rotation, add the annotation keyrotation.csiaddons.openshift.io/schedule: <value> to
a PersistentVolumeClaim, Namespace, or StorageClass (in decreasing order of precedence).

           IMPORTANT

           Key rotation is only supported for RBD backed volumes.

2.4.1.2.1. Annotating a namespace for key rotation

Add the key rotation annotation to a namespace so that all RBD-backed persistent volumes in that
namespace are re-encrypted on the configured schedule.


14

             CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

Prerequisites

     Key rotation is only supported for RBD backed volumes.

Procedure

1.   Verify the namespace exists:

         $ oc get namespace default
         NAME    STATUS  AGE
         default Active 5d2h

2.   Annotate the namespace with the desired rotation schedule:

         $ oc annotate namespace default "keyrotation.csiaddons.openshift.io/schedule=@weekly"
         namespace/default annotated

     <value> can be @hourly, @daily, @weekly, @monthly, or @yearly. If <value> is empty, the
     default is @weekly.

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

2.4.1.2.2. Annotating a storage class for key rotation

Add the key rotation annotation to a storage class so that all RBD-backed persistent volumes
provisioned from that storage class are re-encrypted on the configured schedule.

Prerequisites

     Key rotation is only supported for RBD backed volumes.

Procedure

1.   Verify the storage class exists:

         $ oc get storageclass rbd-sc
         NAME    PROVISIONER     RECLAIMPOLICY     VOLUMEBINDINGMODE
         ALLOWVOLUMEEXPANSION     AGE
         rbd-sc  rbd.csi.ceph.com    Delete Immediate     true      5d2h

2.   Annotate the storage class with the desired rotation schedule:


                         15

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

         $ oc annotate storageclass rbd-sc "keyrotation.csiaddons.openshift.io/schedule=@weekly"
         storageclass.storage.k8s.io/rbd-sc annotated

     <value> can be @hourly, @daily, @weekly, @monthly, or @yearly. If <value> is empty, the
     default is @weekly.

2.4.1.2.3. Annotating a PersistentVolumeClaim for key rotation

Add the key rotation annotation directly to a PersistentVolumeClaim to re-encrypt that volume on the
configured schedule, overriding any namespace or storage class annotation.

Prerequisites

     Key rotation is only supported for RBD backed volumes.

Procedure

1.   Verify the PersistentVolumeClaim exists:

         $ oc get pvc data-pvc
         NAME    STATUS  VOLUME                                   CAPACITY  ACCESS MODES
         STORAGECLASS     AGE
         data-pvc Bound           pvc-f37b8582-4b04-4676-88dd-e1b95c6abf74  1Gi   RWO
         default 20h

2.   Annotate the PVC with the desired rotation schedule:

         $ oc annotate pvc data-pvc "keyrotation.csiaddons.openshift.io/schedule=@weekly"
         persistentvolumeclaim/data-pvc annotated

     <value> can be @hourly, @daily, @weekly, @monthly, or @yearly. If <value> is empty, the
     default is @weekly.

3.   Verify the encryption key rotation cron job was created:

         $ oc get encryptionkeyrotationcronjobs.csiaddons.openshift.io
         NAME     SCHEDULE     SUSPEND     ACTIVE     LASTSCHEDULE                AGE
         data-pvc-1642663516     @weekly        3s

4.   Optional: Update the schedule by overwriting the annotation:

         $ oc annotate pvc data-pvc "keyrotation.csiaddons.openshift.io/schedule=*/1 * * * *" --
         overwrite=true
         persistentvolumeclaim/data-pvc annotated

5.   Verify the updated cron job schedule:

         $ oc get encryptionkeyrotationcronjobs.csiaddons.openshift.io
         NAME     SCHEDULE     SUSPEND     ACTIVE     LASTSCHEDULE                AGE
         data-pvc-1642664617 */1 * * * *        3s

2.4.1.3. Disabling key rotation

16

            CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

Disable automatic key rotation for all PVCs in a storage class or for a specific persistent volume claim in
OpenShift Data Foundation.

You can disable key rotation for the following:

      All the persistent volume claims (PVCs) of storage class

      A specific PVC

Disabling key rotation for all PVCs of a storage class

To disable key rotation for all PVCs, update the annotation of the storage class:

 $ oc get storageclass rbd-sc
 NAME      PROVISIONER     RECLAIMPOLICY     VOLUMEBINDINGMODE
 ALLOWVOLUMEEXPANSION     AGE
 rbd-sc    rbd.csi.ceph.com  Delete   Immediate     true      5d2h

 $ oc annotate storageclass rbd-sc "keyrotation.csiaddons.openshift.io/enable: false"
 storageclass.storage.k8s.io/rbd-sc annotated

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

2.5. FINDING AVAILABLE STORAGE DEVICES


            17

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

Use this procedure to identify the device names for each of the three or more worker nodes that you
have labeled with the OpenShift Data Foundation label cluster.ocs.openshift.io/openshift-storage=''
before creating PVs for IBM Power.

Procedure

  1.   List and verify the name of the worker nodes with the OpenShift Data Foundation label.

         $ oc get nodes -l cluster.ocs.openshift.io/openshift-storage=

       Example output:

         NAME         STATUS ROLES       AGE       VERSION
         worker-0         Ready  worker  2d11h v1.23.3+e419edf
         worker-1         Ready  worker  2d11h v1.23.3+e419edf
         worker-2         Ready  worker  2d11h v1.23.3+e419edf

  2.   Log in to each worker node that is used for OpenShift Data Foundation resources and find the
       name of the additional disk that you have attached while deploying Openshift Container
       Platform.

         $ oc debug node/<node name>

       Example output:

         $ oc debug node/worker-0
         Starting pod/worker-0-debug ...
         To use host binaries, run `chroot /host`
         Pod IP: 192.168.0.63
         If you don't see a command prompt, try pressing enter.
         sh-4.4#
         sh-4.4# chroot /host
         sh-4.4# lsblk
         NAME      MAJ:MIN RM SIZE RO TYPE MOUNTPOINT
         loop1     7:1      0 500G 0 loop
         sda      8:0   0    500G 0 disk
         sdb      8:16      0 120G 0 disk
         |-sdb1    8:17   0  4M 0 part
         |-sdb3    8:19   0 384M 0 part
         `-sdb4    8:20   0 119.6G 0 part
         sdc      8:32      0 500G 0 disk
         sdd      8:48  0     120G 0 disk
         |-sdd1    8:49   0  4M 0 part
         |-sdd3    8:51   0 384M 0 part
         `-sdd4    8:52   0 119.6G 0 part
         sde      8:64      0 500G 0 disk
         sdf     8:80       0 120G 0 disk
         |-sdf1   8:81      0  4M 0 part
         |-sdf3   8:83      0 384M 0 part
         `-sdf4   8:84  0 119.6G 0 part
         sdg      8:96      0 500G 0 disk
         sdh      8:112 0    120G 0 disk
         |-sdh1    8:113 0      4M 0 part
         |-sdh3    8:115 0   384M 0 part
         `-sdh4    8:116 0 119.6G 0 part

18

                 CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

         sdi     8:128 0  500G 0 disk
         sdj     8:144 0  120G 0 disk
         |-sdj1   8:145 0  4M 0 part
         |-sdj3   8:147 0   384M 0 part
         `-sdj4   8:148 0 119.6G 0 part
         sdk      8:160 0 500G 0 disk
         sdl     8:176 0  120G 0 disk
         |-sdl1   8:177 0  4M 0 part
         |-sdl3   8:179 0 384M 0 part
         `-sdl4   8:180 0 119.6G 0 part /sysroot
         sdm       8:192 0 500G 0 disk
         sdn      8:208 0 120G 0 disk
         |-sdn1    8:209 0  4M 0 part
         |-sdn3    8:211 0  384M 0 part /boot
         `-sdn4    8:212 0 119.6G 0 part
         sdo      8:224 0 500G 0 disk
         sdp      8:240 0 120G 0 disk
         |-sdp1    8:241 0  4M 0 part
         |-sdp3    8:243 0 384M 0 part
         `-sdp4    8:244 0 119.6G 0 part

     In this example, for worker-0, the available local devices of 500G are sda, sdc, sde,sdg, sdi,
     sdk, sdm,sdo.

3.   Repeat the above step for all the other worker nodes that have the storage devices to be used
     by OpenShift Data Foundation. See this Knowledge Base article for more details.

2.6. CREATING OPENSHIFT DATA FOUNDATION CLUSTER ON IBM
POWER

Use this procedure to create an OpenShift Data Foundation cluster after you install the OpenShift Data
Foundation operator.

Prerequisites

     Ensure that all the requirements in the Requirements for installing OpenShift Data Foundation
     using local storage devices section are met.

     You must have a minimum of three worker nodes with the same storage type and size attached
     to each node (for example, 200 GB SSD) to use local storage devices on IBM Power.

     Verify your OpenShift Container Platform worker nodes are labeled for OpenShift Data
     Foundation:

         oc get nodes -l cluster.ocs.openshift.io/openshift-storage -o jsonpath='{range .items[*]}
         {.metadata.name}{"\n"}'

To identify storage devices on each node, refer to Finding available storage devices .

Procedure

1.   Log into the OpenShift Web Console.



2. In openshift-local-storage namespace Click Ecosystem → Installed Operators to view the
                       19

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

2.  In openshift-local-storage namespace Click Ecosystem → Installed Operators to view the
    installed operators.

3.  Click the Local Storage installed operator.

4.  On the Operator Details page, click the Local Volume link.

5.  Click Create Local Volume.

6.  Click on YAML view for configuring Local Volume.

7.  Define a LocalVolume custom resource for block PVs using the following YAML.

       apiVersion: local.storage.openshift.io/v1
       kind: LocalVolume
       metadata:
        name: localblock
        namespace: openshift-local-storage
       spec:
        logLevel: Normal
        managementState: Managed
        nodeSelector:
         nodeSelectorTerms:
                 - matchExpressions:
               - key: kubernetes.io/hostname
                     operator: In
               values:
                      - worker-0
                      - worker-1
                      - worker-2
        storageClassDevices:
         - devicePaths:
               - /dev/sda
             storageClassName: localblock
                  volumeMode: Block

    The above definition selects sda local device from the worker-0, worker-1 and worker-2
    nodes. The localblock storage class is created and persistent volumes are provisioned from
    sda.

                 IMPORTANT

                 Specify appropriate values of nodeSelector as per your environment. The device
                 name should be same on all the worker nodes. You can also specify more than
                 one devicePaths.

8.  Click Create.

9.  Confirm whether diskmaker-manager pods and Persistent Volumes are created.

    a. For Pods

         i.    Click Workloads → Podsfrom the left pane of the OpenShift Web Console.

         ii.   Select openshift-local-storage from the Project drop-down list.


20

                  CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

           iii.   Check if there are diskmaker-manager pods for each of the worker node that you used
                  while creating LocalVolume CR.

       b. For Persistent Volumes

             i.   Click Storage → PersistentVolumes from the left pane of the OpenShift Web Console.

            ii.   Check the Persistent Volumes with the name local-pv-*. Number of Persistent Volumes
                  will be equivalent to the product of number of worker nodes and number of storage
                  devices provisioned while creating localVolume CR.

                  IMPORTANT

                           The flexible scaling feature is enabled only when the storage cluster
                           that you created with three or more nodes are spread across fewer
                           than the minimum requirement of three availability zones.
                           For information about flexible scaling, see knowledgebase article on
                           Scaling OpenShift Data Foundation cluster using YAML when flexible
                           scaling is enabled.

                           Flexible scaling features get enabled at the time of deployment and
                           can not be enabled or disabled later on.

10.   In the OpenShift Web Console, click Storage → Data Foundation → Storage Systems →
      Create StorageSystem.

11.   In the Backing storage page, perform the following:

       a.  Select the Use an existing StorageClass option.

       b.  Select the required Storage Class that you used while installing LocalVolume.
           By default, it is set to none.

       c.  Optional: Select Use Ceph RBD as the default StorageClass. This avoids having to
           manually annotate a StorageClass.

       d.  Optional: Select Use external PostgreSQL checkbox to use an external PostgreSQL
           [Technology preview].
           This provides high availability solution for Multicloud Object Gateway where the
           PostgreSQL pod is a single point of failure.










                  IMPORTANT
                           21

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

                   IMPORTANT

                   OpenShift Data Foundation ships PostgreSQL images maintained by Red
                   Hat, which are used to store metadata for the Multicloud Object Gateway.
                   This PostgreSQL usage is at the application level.

                   As a result, OpenShift Data Foundation does not perform database-level
                   optimizations or in-depth insights.

                   If customers have their own PostgreSQL that is well-maintained and
                   optimized, we recommend using it. OpenShift Data Foundation supports
                   external PostgreSQL instances.

                   Any PostgreSQL-related issues requiring code changes or deep technical
                   analysis may need to be addressed upstream. This could result in longer
                   resolution times.

         i.   Provide the following connection details:

               Username

               Password

               Server name and Port

               Database name

        ii.   Select Enable TLS/SSL checkbox to enable encryption for the Postgres server.

    e.  Click Next.

12. In the Capacity and nodes page, configure the following:

    a.  Available raw capacity is populated with the capacity value based on all the attached disks
        associated with the storage class. This takes some time to show up. The Selected nodes
        list shows the nodes based on the storage class.

    b.  In the Configure performance section, select one of the following performance profiles:

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
22

                 CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

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

      c.  Optional: Select the Taint nodes checkbox to dedicate the selected nodes for OpenShift
          Data Foundation.

      d.  Click Next.

13.  Optional: In the Security and network page, configure the following based on your
     requirements:

      a.  To enable encryption, select Enable data encryption for block and file storage.

           i.   Select either one or both the encryption levels:

                     Cluster-wide encryption
                     Encrypts the entire cluster (block and file).

                     StorageClass encryption
                     Creates encrypted persistent volume (block only) using encryption enabled storage
                     class.

          ii.   Optional: Select the Connect to an external key management servicecheckbox. This
                is optional for cluster-wide encryption.

                 A.  From the Key Management Service Providerdrop-down list, either select Vault
                     or Thales CipherTrust Manager (using KMIP). If you selected Vault, go to the next
                     step. If you selected Thales CipherTrust Manager (using KMIP), go to step iii.

                 B.  Select an Authentication Method.

                     Using Token authentication method

                           Enter a unique Connection Name, host Address of the Vault server
                           ('https://<hostname or ip>'), Port number and Token.

                           Expand Advanced Settings to enter additional settings and certificate
                           details based on your Vault configuration:

                             Enter the Key Value secret path in Backend Path that is dedicated and
                             unique to OpenShift Data Foundation.

                             Optional: Enter TLS Server Name and Vault Enterprise Namespace.



                             23

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

                      Upload the respective PEM encoded certificate file to provide the CA
                      Certificate, Client Certificate and Client Private Key .

                      Click Save and skip to step iv.

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
                      Certificate, Client Certificate and Client Private Key .

                      Click Save and skip to step iv.

       C.  To use Thales CipherTrust Manager (using KMIP) as the KMS provider, follow the
           steps below:

             I.   Enter a unique Connection Name for the Key Management service within the
                  project.

            II.   In the Address and Port sections, enter the IP of Thales CipherTrust Manager
                  and the port where the KMIP interface is enabled. For example:

                    Address: 123.34.3.2

                    Port: 5696

           III.   Upload the Client Certificate, CA certificate, and Client Private Key.

           IV.    If StorageClass encryption is enabled, enter the Unique Identifier to be used for
                  encryption and decryption generated above.

             V.   The TLS Server field is optional and used when there is no DNS entry for the
                  KMIP endpoint. For example, kmip_all_<port>.ciphertrustmanager.local.

       D.  Select a Network.

           I.     Select Default (OVN) network as Multus is not yet supported on OpenShift
                  Data Foundation on IBM Power.

       E.  Click Next.

b. To enable in-transit encryption, select In-transit encryption.

 i.   Select a Network.

ii.   Click Next.


24

            CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING LOCAL STORAGE DEVICES

14. In the Review and create page::

    a.   Review the configurations details. To modify any configuration settings, click Back to go
         back to the previous configuration page.

    b.   Click Create StorageSystem.

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

    To verify if flexible scaling is enabled on your storage cluster, perform the following steps:

    1.   In the OpenShift Web Console, navigate to Storage → Data Foundation → Storage
         System → ocs-storagecluster.

    2.   From the Action menu, select Edit Storage System.

    3.   In the YAML tab, search for the keys flexibleScaling in spec section and failureDomain in
         status section. If flexible scaling is true and failureDomain is set to host, flexible scaling
         feature is enabled.

            spec:
            flexibleScaling: true
            […]
            status:
            failureDomain: host

    To verify that all the components for OpenShift Data Foundation are successfully installed, see
    Verifying your OpenShift Data Foundation deployment .

Additional resources

    Scaling Storage guide









                    25

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION
         DEPLOYMENT FOR INTERNAL MODE

Use this section to verify that OpenShift Data Foundation is deployed correctly.

1.   Verify the state of the pods.

2.   Verify that the OpenShift Data Foundation cluster is healthy .

3.   Verify that the Multicloud Object Gateway is healthy .

4.   Verify that the OpenShift Data Foundation specific storage classes exist .

3.1. VERIFYING THE STATE OF THE PODS

To determine if OpenShift Data Foundation is deployed successfully, you can verify that the pods are in
Running state.

Procedure

1.   Click Workloads → Podsfrom the left pane of the OpenShift Web Console.

2.   Select openshift-storage from the Project drop-down list.

                  NOTE

                  If the Show default projects option is disabled, use the toggle button to list all
                  the default projects.

     For more information on the expected number of pods for each component and how it varies
     depending on the number of nodes, see ???.

3.   Verify that the following pods are in running and completed state by clicking the Running and
     the Completed tabs:

                  NOTE

                  The available pods depend on the cluster configuration. When the cluster is
                  deployed as a standalone Multicloud Object Gateway, the rook-ceph-operator-*
                  pods are not available. Similarly, when the cluster is deployed without the
                  Multicloud Object Gateway, noobaa-* pods are not available.

     Pods corresponding to OpenShift Data Foundation cluster

     Component        Corresponding pods










26

CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT FOR INTERNAL MODE


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




Rook-ceph Operator                         rook-ceph-operator-* (1 pod on any storage
                                           node)

OpenShift Data Foundation Client Operator   ocs-client-operator-console-* (1 pod on
                                            any storage node)

                                            ocs-client-operator-controller-
                                            manager-* (1 pod on any storage node)


UX Backend                                 ux-backend-server-* (1 pod on any storage
                                           node)

Multicloud Object Gateway                   noobaa-operator-* (1 pod on any storage
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


MON                                        rook-ceph-mon-* (3 pods on each storage
                                           node)

MGR                                        rook-ceph-mgr-* (2 pods distributed across
                                           storage nodes)


                                            27

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power


Component     Corresponding pods

MDS           rook-ceph-mds-ocs-storagecluster-
              cephfilesystem-* (2 pods distributed across
              storage node)

RGW           rook-ceph-rgw-ocs-storagecluster-
              cephobjectstore-* (1 pod on any storage
              node)

CSI            cephfs
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



rook-ceph-crashcollector     rook-ceph-crashcollector-* (1 pod on each
                             storage node)

rook-ceph-exporter           rook-ceph-exporter-worker-* (1 pod on each
                             storage node)






28

CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT FOR INTERNAL MODE


Component    Corresponding pods

OSD          rook-ceph-osd-* (1 pod for each device)

             rook-ceph-osd-prepare-* (1 pod for each
             device)



3.2. VERIFYING THE OPENSHIFT DATA FOUNDATION CLUSTER IS
HEALTHY

Verify that the OpenShift Data Foundation cluster on IBM Power Systems is healthy after deployment
by checking the storage system status and data resiliency indicators in the OpenShift Web Console.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   Click the Storage Systems tab and then click on ocs-storagecluster.

 3. In the Status card of Block and File dashboard under Overview tab, verify that both Storage
    Cluster and Data Resiliency has a green tick mark.

4.   In the Details card, verify that the cluster information is displayed.
     For more information on the health of the OpenShift Data Foundation cluster using the Block
     and File dashboard, see Monitoring OpenShift Data Foundation .

3.3. VERIFYING THE MULTICLOUD OBJECT GATEWAY IS HEALTHY

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
                 29

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

                      IMPORTANT

                      To avoid data loss, it is recommended to take a backup of NooBaa DB PVC
                      regularly. If NooBaa DB fails and cannot be recovered, then you can revert to
                      the latest backed-up version. For instructions on backing up your NooBaa
                      DB, follow the steps in the knowledgebase article, Perform a One-Time
                      Backup of the Database for the Multicloud Object Gateway.

3.4. VERIFYING THAT THE SPECIFIC STORAGE CLASSES EXIST

Verify that the expected OpenShift Data Foundation storage classes are present after deployment.

Procedure

1.   Click Storage → Storage Classesfrom the left pane of the OpenShift Web Console.

2.   Verify that the following storage classes are created with the OpenShift Data Foundation
     cluster creation:

         ocs-storagecluster-ceph-rbd

         ocs-storagecluster-cephfs

         openshift-storage.noobaa.io

         ocs-storagecluster-ceph-rgw










30

                CHAPTER 4. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY

            CHAPTER 4. DEPLOY STANDALONE MULTICLOUD OBJECT
                               GATEWAY

Deploying only the Multicloud Object Gateway component with the OpenShift Data Foundation
provides the flexibility in deployment and helps to reduce the resource consumption. After deploying
the MCG component, you can create and manage buckets using MCG object browser. For more
information, see Creating and managing buckets using MCG object browser .

Use this section to deploy only the standalone Multicloud Object Gateway component, which involves
the following steps:

     Installing the Local Storage Operator.

     Installing Red Hat OpenShift Data Foundation Operator

     Creating standalone Multicloud Object Gateway

            IMPORTANT

            To avoid data loss, it is recommended to take a backup of NooBaa DB PVC regularly. If
            NooBaa DB fails and cannot be recovered, then you can revert to the latest backed-up
            version. For instructions on backing up NooBaa DB, follow the steps in the
            knowledgebase article, Backup and Restore for Multicloud Object Gateway database
            (NooBaa DB).

4.1. INSTALLING LOCAL STORAGE OPERATOR

Use this procedure to install the Local Storage Operator from the Software Catalog before creating
OpenShift Data Foundation clusters on local storage devices.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Ecosystem → Software Catalog.

3.   Type local storage in the Filter by keyword… box to find the Local Storage Operator from
     the list of operators and click on it.

4.   Set the following options on the Install Operator page:

      a.  Update channel as stable.

      b.  Installation Mode as A specific namespace on the cluster.

      c.  Installed Namespace as Operator recommended namespace openshift-local-storage.

      d.  Approval Strategy as Automatic.

5.   Click Install.

Verification

     Verify that the Local Storage Operator shows a green tick indicating successful installation.


            31

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

4.2. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.

For information about the hardware and software requirements, see Planning your deployment.

Prerequisites

         Access to an OpenShift Container Platform cluster using an account with cluster-admin and
         Operator installation permissions.

         You must have at least three worker nodes in the Red Hat OpenShift Container Platform
         cluster.

             IMPORTANT

                       When you need to override the cluster-wide default node selector for OpenShift
                       Data Foundation, you can use the following command in the command line
                       interface to specify a blank node selector for the openshift-storage namespace
                       (create openshift-storage namespace in this case):

                 $ oc annotate namespace openshift-storage openshift.io/node-selector=

                       Taint a node as infra to ensure only Red Hat OpenShift Data Foundation
                       resources are scheduled on that node. This helps you save on subscription costs.
                       For more information, see How to use dedicated worker nodes for Red Hat
                       OpenShift Data Foundation chapter in the Managing and Allocating Storage
                       Resources guide.

Procedure

    1.   Log in to the OpenShift Web Console.

    2.   Click Ecosystem → Software Catalog.

    3.   Scroll or type OpenShift Data Foundation into the Filter by keyword box to find the
         OpenShift Data Foundation Operator.

    4.   Click Install.

    5.   Set the following options on the Install Operator page:

         a.  Update Channel as stable-4.22.

         b.  Installation Mode as A specific namespace on the cluster.

         c.  Installed Namespace as Operator recommended namespace openshift-storage. If
             Namespace openshift-storage does not exist, it is created during the operator installation.

    6.   Select Approval Strategy as Automatic or Manual.
         If you select Automatic updates, then the Operator Lifecycle Manager (OLM) automatically
         upgrades the running instance of your Operator without any intervention.


32

             CHAPTER 4. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY

         If you select Manual updates, then the OLM creates an update request. As a cluster
         administrator, you must then manually approve that update request to update the Operator to a
         newer version.

     7.  Ensure that the Enable option is selected for the Console plugin.

     8.  Click Install.

Verification

         Verify that the OpenShift Data Foundation Operator shows a green tick indicating successful
         installation.

         After the operator is successfully installed, the web console automatically reloads to apply the
         changes. During this process, a temporary error message might appear on the page and this is
         expected and disappears after the refresh completes.

         In the Web Console, navigate to Storage and verify if Data Foundation is available.

4.3. CREATING STANDALONE MULTICLOUD OBJECT GATEWAY ON
IBM POWER

You can create only the standalone Multicloud Object Gateway component while deploying OpenShift
Data Foundation.

Prerequisites

         Ensure that the OpenShift Data Foundation Operator is installed.

         The Multicloud Object Gateway database uses ReadWriteOncePod (RWOP) access mode to
         safeguard against node failures. If the underlying CSI driver does not support RWOP, the access
         mode must be manually changed to ReadWriteOnce (RWO) before deployment. For more
         information, see the knowledgebase article How to change NooBaa DB access mode from
         RWOP to RWO in standalone MCG deployments?.

         (For deploying using local storage devices only) Ensure that Local Storage Operator is installed.

To identify storage devices on each node, refer to Finding available storage devices .

 Procedure

1.       Log into the OpenShift Web Console.

     2.  In openshift-local-storage namespace, click Ecosystem → Installed Operators to view the
         installed operators.

     3.  Click the Local Storage installed operator.

     4.  On the Operator Details page, click the Local Volume link.

     5.  Click Create Local Volume.

     6.  Click on YAML view for configuring Local Volume.

     7.  Define a LocalVolume custom resource for filesystem PVs using the following YAML.


         33

   Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

          apiVersion: local.storage.openshift.io/v1
          kind: LocalVolume
          metadata:
           name: localblock
           namespace: openshift-local-storage
          spec:
           logLevel: Normal
           managementState: Managed
           nodeSelector:
           nodeSelectorTerms:
                   - matchExpressions:
               - key: kubernetes.io/hostname
                       operator: In
               values:
                        - worker-0
                        - worker-1
                        - worker-2
           storageClassDevices:
           - devicePaths:
            - /dev/sda
            storageClassName: localblock
            volumeMode: Filesystem

      The above definition selects sda local device from the worker-0, worker-1 and worker-2
      nodes. The localblock storage class is created and persistent volumes are provisioned from
      sda.

                   IMPORTANT

                   Specify appropriate values of nodeSelector as per your environment. The device
                   name should be same on all the worker nodes. You can also specify more than
                   one devicePaths.

 8.   Click Create.

 9.   In the OpenShift Web Console, navigate to Storage → StorageCluster.

10.   On the Welcome page, click Configure Data Foundation.

11.   Select Setup Multicloud Object Gateway.

12.   Select the Use an existing StorageClass option for Backing storage type .

       a.  Select the Storage Class that you used while installing LocalVolume.

13.   Click Next.

14.   Optional: In the Security page, select the Connect to an external key management service
      checkbox. This is optional for cluster-wide encryption.

       a.  From the Key Management Service Providerdrop down list, either select Vault or Thales
           CipherTrust Manager (using KMIP). If you selected Vault, go to the next step. If you
           selected Thales CipherTrust Manager (using KMIP), go to step iii.

       b.  Select an Authentication Method.


   34

                                         CHAPTER 4. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY

Using Token authentication method

            Enter a unique Connection Name, host Address of the Vault server
            ('https://<hostname or ip>'), Port number and Token.

            Expand Advanced Settings to enter additional settings and certificate details
            based on your Vault configuration:

              Enter the Key Value secret path in Backend Path that is dedicated and unique
              to OpenShift Data Foundation.

              Optional: Enter TLS Server Name and Vault Enterprise Namespace.

              Upload the respective PEM encoded certificate file to provide the CA
              Certificate, Client Certificate and Client Private Key .

              Click Save and skip to step iv.

Using Kubernetes authentication method

            Enter a unique Vault Connection Name, host Address of the Vault server
            ('https://<hostname or ip>'), Port number and Role name.

            Expand Advanced Settings to enter additional settings and certificate details
            based on your Vault configuration:

              Enter the Key Value secret path in the Backend Path that is dedicated and
              unique to OpenShift Data Foundation.

              Optional: Enter TLS Server Name and Authentication Path if applicable.

              Upload the respective PEM encoded certificate file to provide the CA
              Certificate, Client Certificate, and Client Private Key.

              Click Save and skip to step iv.

c.  To use Thales CipherTrust Manager (using KMIP) as the KMS provider, follow the steps
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

d. Select a Network.

                                                                                             35

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

       e.    Click Next.

 15.  In the Review and create page, review the configuration details:
      To modify any configuration settings, click Back.

 16.  Click Create StorageSystem.

Verification

Verifying that the OpenShift Data Foundation cluster is healthy

       1.   In the OpenShift Web Console, click Storage → Data Foundation.

      2.    Click the Storage Systems tab and then click on ocs-storagecluster.

             a.  In the Status card of the Object tab, verify that both Object Service and Data Resiliency
                 have a green tick.

             b.  In the Details card, verify that the MCG information is displayed.

Verifying the state of the pods

       1.   Click Workloads → Pods from the OpenShift Web Console.

      2.    Select openshift-storage from the Project drop-down list and verify that the following
            pods are in Running state.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to list
                        all the default projects.


Component               Corresponding pods

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







36

CHAPTER 4. DEPLOY STANDALONE MULTICLOUD OBJECT GATEWAY


Component        Corresponding pods

Multicloud Object     noobaa-operator-* (1 pod on any storage node)
Gateway
                      noobaa-core-* (1 pod on any storage node)

                      noobaa-db-pg-cluster-1 and noobaa-db-pg-cluster-2 (2
                      instances of MCG DB pod on any storage node)

                      noobaa-endpoint-* (1 pod on any storage node)

                      cnpg-controller-manager-* (1 pod on any storage node)

                      noobaa-default-backing-store-noobaa-pod-* (1 pod on
                      any storage node)










                      37

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Power

      CHAPTER 5. VIEW OPENSHIFT DATA FOUNDATION
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




38

         CHAPTER 6. UNINSTALLING OPENSHIFT DATA FOUNDATION

CHAPTER 6. UNINSTALLING OPENSHIFT DATA FOUNDATION

6.1. UNINSTALLING OPENSHIFT DATA FOUNDATION IN INTERNAL
MODE

Uninstalling OpenShift Data Foundation in internal mode permanently removes the storage cluster, all
associated PVCs, and the operator.

Procedure

1.   To uninstall OpenShift Data Foundation in internal mode, refer to the knowledgebase article on
     Uninstalling OpenShift Data Foundation.










     39
