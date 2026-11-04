---
title: "deploying openshift data foundation using red hat openshift service on aws with hosted control planes"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/deploying_openshift_data_foundation_using_red_hat_openshift_service_on_aws_with_hosted_control_planes"
source_file: "deploying_openshift_data_foundation_using_red_hat_openshift_service_on_aws_with_hosted_control_planes.pdf"
version: "4.22"
chapter: "deploying_openshift_data_foundation_using_red_hat_openshift_service_on_aws_with_hosted_control_planes"
---

 Red Hat OpenShift Data Foundation
        4.22

Deploying OpenShift Data Foundation using
  Red Hat OpenShift Service on AWS with
        hosted control planes

Instructions for deploying OpenShift Data Foundation using Red Hat OpenShift
     Service on AWS (ROSA) with hosted control planes (HCP)










    Last Updated: 2026-09-10



Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data
Foundation using Red Hat OpenShift Service on AWS with hosted control
planes

Instructions for deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS
(ROSA) with hosted control planes (HCP)

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
Red Hat OpenShift Container Platform on ROSA with hosted control planes (HCP).

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA FOUNDATION 6

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING DYNAMIC STORAGE DEVICES 8
  2.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR 8
  2.2. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE TOKEN AUTHENTICATION METHOD
        10
  2.3. ENABLING CLUSTER-WIDE ENCRYPTION WITH KMS USING THE KUBERNETES AUTHENTICATION
  METHOD 11
  2.4. CREATING OPENSHIFT DATA FOUNDATION CLUSTER 13

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT 19
  3.1. VERIFYING THE STATE OF THE PODS 19
  3.2. VERIFYING THE OPENSHIFT DATA FOUNDATION CLUSTER IS HEALTHY 22
  3.3. VERIFYING THAT THE SPECIFIC STORAGE CLASSES EXIST 22

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. VIEW OPENSHIFT DATA FOUNDATION TOPOLOGY 23

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. UNINSTALLING OPENSHIFT DATA FOUNDATION 24
  5.1. UNINSTALLING OPENSHIFT DATA FOUNDATION IN INTERNAL MODE 24










    1

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w










2

    PREFACE

PREFACE

Deploy OpenShift Data Foundation on Red Hat OpenShift Service on AWS (ROSA) with hosted control
planes using dynamic storage devices to create internal cluster resources.










3

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

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

Red Hat OpenShift Data Foundation supports deployment on Red Hat OpenShift Service on AWS with
Hosted Control Planes (ROSA HCP).

  NOTE

  Only internal OpenShift Data Foundation clusters are supported on AWS. See the
  Planning your deployment guide and the Preparing to deploy OpenShift Data Foundation
  chapter for more information about deployment requirements.

To deploy OpenShift Data Foundation, start with the requirements in the Preparing to deploy OpenShift
Data Foundation chapter and then follow the deployment process in the Deploying using dynamic
storage devices section.

Additional resources

 Planning your deployment

 Preparing to deploy OpenShift Data Foundation










  5

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

  CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA
                       FOUNDATION

 Deploying OpenShift Data Foundation on OpenShift Container Platform using dynamic storage devices
 provides you with the option to create internal cluster resources.

 Before you begin the deployment of Red Hat OpenShift Data Foundation, follow these steps:

  1.   Optional: If you want to enable cluster-wide encryption using the external Key Management
       System (KMS) HashiCorp Vault, follow these steps:

            Ensure that you have a valid Red Hat OpenShift Data Foundation Advanced subscription.
            To know how subscriptions for OpenShift Data Foundation work, see knowledgebase article
            on OpenShift Data Foundation subscriptions.

            When the Token authentication method is selected for encryption then refer to Enabling
            cluster-wide encryption with the Token authentication using KMS.

            When the Kubernetes authentication method is selected for encryption then refer to
            Enabling cluster-wide encryption with the Kubernetes authentication using KMS .

            Ensure that you are using signed certificates on your Vault servers.

  2.   Optional: If you want to enable cluster-wide encryption using the external Key Management
       System (KMS) Thales CipherTrust Manager, you must first enable the Key Management
       Interoperability Protocol (KMIP) and use signed certificates on your server.

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

            i.     Select KMIP Key Management Interoperability Protocol and click Next.

             ii.   Select a free Port.

            iii.   Select Network Interface as all.

             iv.   Select Interface Mode as TLS, verify client cert, user name taken from client cert,
                   auth request is optional.

              v.   (Optional) You can enable hard delete to delete both metadata and material when the
                   key is deleted. It is disabled by default.

             vi.   Select the certificate authority (CA) to be used, and click Save.

6

                      CHAPTER 1. PREPARING TO DEPLOY OPENSHIFT DATA FOUNDATION

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

3.  Minimum starting node requirements: An OpenShift Data Foundation cluster is deployed with
    minimum configuration when the standard deployment resource requirement is not met. See
    Resource requirements in the Planning guide .

           IMPORTANT

           In order perform stop and start node operations, or to create or add a machine pool, it is
           necessary to apply proper labeling. For example:

                  $ rosa edit machinepool --cluster <cluster-name> --labels
                  cluster.ocs.openshift.io/openshift-storage="" <machinepool-name>

           Replace <cluster-name> with the cluster name and <machinepool-name> with the
           machine pool name.

Additional resources

    knowledgebase article on OpenShift Data Foundation subscriptions

    Enabling cluster-wide encryption with the Token authentication using KMS

    Enabling cluster-wide encryption with the Kubernetes authentication using KMS

    Resource requirements










                  7

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

 CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING
                  DYNAMIC STORAGE DEVICES

 You can deploy OpenShift Data Foundation on OpenShift Container Platform using dynamic storage
 devices provided by Amazon Web Services (AWS) EBS (type, gp2-csi or gp3-csi) that provides you with
 the option to create internal cluster resources. This results in the internal provisioning of the base
 services, which helps to make additional storage classes available to applications.

 Although, it is possible to deploy only the Multicloud Object Gateway (MCG) component with
 OpenShift Data Foundation, this deployment method is not supported on ROSA.

              NOTE

              Only internal OpenShift Data Foundation clusters are supported on ROSA. See Planning
              your deployment for more information about deployment requirements.

 Also, ensure that you have addressed the requirements in Preparing to deploy OpenShift Data
 Foundation chapter before proceeding with the below steps for deploying using dynamic storage
 devices:

 1.      Install the Red Hat OpenShift Data Foundation Operator.

 2.      Create the OpenShift Data Foundation Cluster.

 2.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
 OPERATOR

 Install the OpenShift Data Foundation Operator using the Red Hat OpenShift Container Platform
 Software Catalog for ROSA with hosted control planes (HCP) to deploy the storage operator suite.

 Prerequisites

         Access to an OpenShift Container Platform cluster using an account with cluster-admin and
         operator installation permissions.

         You must have at least three worker or infrastructure nodes in the Red Hat OpenShift
         Container Platform cluster.

         For additional resource requirements, see the Planning your deployment guide.

             IMPORTANT

                  When you need to override the cluster-wide default node selector for OpenShift
                  Data Foundation, you can use the following command to specify a blank node
                  selector for the storage namespace:

                  $ oc annotate namespace storage-namespace openshift.io/node-selector=

                  Taint a node as infra to ensure only Red Hat OpenShift Data Foundation
                  resources are scheduled on that node. This helps you save on subscription costs.
                  For more information, see the How to use dedicated worker nodes for Red Hat
                  OpenShift Data Foundation section in the Managing and Allocating Storage
                  Resources guide.

8

            CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING DYNAMIC STORAGE DEVICES

Procedure

  1.   Log in to the OpenShift Web Console.

  2.   Click Ecosystem → Software Catalog.

  3.   Scroll or type OpenShift Data Foundation into the Filter by keyword box to find the
       OpenShift Data Foundation Operator.

  4.   Click Install.

  5.   Set the following options on the Install Operator page:

        a. Fill in role ARN. For instruction to create a Amazon resource name (ARN), see Creating an
        AWS role using a script.

       b.   Update Channel as stable-4.22.

       c.   Installation Mode as A specific namespace on the cluster.

       d.   Installed Namespace as Select a Namespace.

                          NOTE

                          openshift-storage Namespace is not recommended for ROSA deployments.
                          Use a user defined namespace for this deployment. Avoid using “redhat” or
                          “openshift” prefixes in namespaces.

                          IMPORTANT

                          This guide uses <storage-namespace> as an example namespace. Replace
                          <storage-namespace> with your defined namespace in later steps.

       e.   Select Approval Strategy as Automatic or Manual.
            If you select Automatic updates, then the Operator Lifecycle Manager (OLM)
            automatically upgrades the running instance of your Operator without any intervention.

            If you select Manual updates, then the OLM creates an update request. As a cluster
            administrator, you must then manually approve that update request to update the Operator
            to a newer version. Manual updates strategy is recommended for ROSA with hosted control
            planes.

       f.   Ensure that the Enable option is selected for the Console plugin.

       g.   Click Install.










                          IMPORTANT
                              9

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

             IMPORTANT

             Due to a known issue regarding missing toleration on the prometheus-
             operator pod, run the following command after deployment:

                        $ oc patch csv odf-prometheus-operator.v4.22.0-rhodf -n odf-storage --
                        type=json -p='[{"op": "add", "path":
                        "/spec/install/spec/deployments/0/spec/template/spec/tolerations", "value":
                        [

                        {"key": "node.ocs.openshift.io/storage", "operator": "Equal", "value": "true",
                        "effect": "NoSchedule" }
                        ]}]'

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
     since you cannot change it later.

Procedure

1.   Enable the Key/Value (KV) backend path in the vault.
     For vault KV secret engine API, version 1:

         $ vault secrets enable -path=odf kv

     For vault KV secret engine API, version 2:

         $ vault secrets enable -path=odf kv-v2



10

                        CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING DYNAMIC STORAGE DEVICES

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

Enable the Kubernetes authentication method for cluster-wide encryption using the Key Management
System (KMS).

Prerequisites

      Administrator access to Vault.

      A valid Red Hat OpenShift Data Foundation Advanced subscription. For more information, see
      the knowledgebase article on OpenShift Data Foundation subscriptions .

      The OpenShift Data Foundation operator must be installed from the Software Catalog.

      Select a unique path name as the backend path that follows the naming convention carefully.
      You cannot change this path name later.

Procedure

 1.   Create a service account:

         $ oc -n <storage-namespace> create serviceaccount <serviceaccount_name>

      where, <serviceaccount_name> specifies the name of the service account.

      For example:

         $ oc -n <storage-namespace> create serviceaccount odf-vault-auth

 2.   Create clusterrolebindings and clusterroles:

         $ oc -n <storage-namespace> create clusterrolebinding vault-tokenreview-binding --
         clusterrole=system:auth-delegator --serviceaccount=<storage-
         namespace>:_<serviceaccount_name>_

      For example:

         $ oc -n <storage-namespace> create clusterrolebinding vault-tokenreview-binding --

                                                                                               11

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

     clusterrole=system:auth-delegator --serviceaccount=<storage-namespace>:odf-vault-auth

3. Create a secret for the serviceaccount token and CA certificate.

     $ cat <<EOF | oc create -f -
     apiVersion: v1
     kind: Secret
     metadata:
      name: odf-vault-auth-token
      namespace: <storage-namespace>
      annotations:
      kubernetes.io/service-account.name: <serviceaccount_name>
     type: kubernetes.io/service-account-token
     data: {}
     EOF

    where, <serviceaccount_name> is the service account created in the earlier step, and
    <storage-namespace> is the namespace where ODF operator and StorageSystem were
    created.

4.  Get the token and the CA certificate from the secret.

     $ SA_JWT_TOKEN=$(oc -n <storage-namespace> get secret odf-vault-auth-token -o
     jsonpath="{.data['token']}" | base64 --decode; echo)
     $ SA_CA_CRT=$(oc -n <storage-namespace> get secret odf-vault-auth-token -o jsonpath="
     {.data['ca\.crt']}" | base64 --decode; echo)

5.  Retrieve the OCP cluster endpoint.

     $ OCP_HOST=$(oc config view --minify --flatten -o jsonpath="{.clusters[0].cluster.server}")

6.  Fetch the service account issuer:

     $ oc proxy &
     $ proxy_pid=$!
     $ issuer="$( curl --silent http://127.0.0.1:8001/.well-known/openid-configuration | jq -r
     .issuer)"
     $ kill $proxy_pid

7.  Use the information collected in the previous step to setup the Kubernetes authentication
    method in Vault:

     $ vault auth enable kubernetes

     $ vault write auth/kubernetes/config \
            token_reviewer_jwt="$SA_JWT_TOKEN" \
            kubernetes_host="$OCP_HOST" \
            kubernetes_ca_cert="$SA_CA_CRT" \
            issuer="$issuer"






            IMPORTANT
12

                            CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING DYNAMIC STORAGE DEVICES

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
          bound_service_account_namespaces=<storage-namespace> \
          policies=odf \
          ttl=1440h

     The role odf-rook-ceph-op is later used while you configure the KMS connection details during
     the creation of the storage system.

         $ vault write auth/kubernetes/role/odf-rook-ceph-osd \
          bound_service_account_names=rook-ceph-osd \
          bound_service_account_namespaces=<storage-namespace> \
          policies=odf \
          ttl=1440h

2.4. CREATING OPENSHIFT DATA FOUNDATION CLUSTER

Create an OpenShift Data Foundation cluster after you install the OpenShift Data Foundation operator.

Procedure

1.  In the OpenShift Web Console, click Storage → Data Foundation → Storage Systems →
                                                                                                   13

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

1.   In the OpenShift Web Console, click Storage → Data Foundation → Storage Systems →
     Create StorageSystem.

2.   In the Backing storage page, select the following:

      a.  Select the Use an existing StorageClass option.

      b.  Click Next.

3.   In the Advanced Settings page, you can configure the following optional items:

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

14

               CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING DYNAMIC STORAGE DEVICES

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

                      Before selecting a resource profile, make sure to check the current
                      availability of resources within the cluster. Opting for a higher resource
                      profile in a cluster with insufficient resources might lead to installation
                      failures.



           For more information about resource requirements, see Resource requirement for
                                 15

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

             For more information about resource requirements, see Resource requirement for
             performance profiles.

  e.   Optional: Select the Taint nodes checkbox to dedicate the selected nodes for OpenShift Data
       Foundation.

  f.   Click Next.

        1.   Optional: In the Security and network page, configure the following based on your
             requirements:

  g.   To enable encryption, select Enable data encryption for block and file storage.

        i.   Select either one or both the encryption levels:

                 Cluster-wide encryption
                 Encrypts the entire cluster (block and file).

                 StorageClass encryption
                 Creates encrypted persistent volume (block only) using encryption enabled storage
                 class.

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
16

          CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING DYNAMIC STORAGE DEVICES

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
                                        <storage-namespace> --type=json -p '[{"op": "add",
                                        "path":"/spec/encryption/keyRotation/enable", "value":
                                        true}]'

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

h.                          To enable in-transit encryption, select In-transit encryption.

 i.   Select a Network.

ii.   Click Next.

      1.   In the Review and create page, review the configuration details.
           To modify any configuration settings, click Back.

      2.   Click Create StorageSystem.





           NOTE
                                        17

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

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

 a.  In the OpenShift Web Console, navigate to Storage → Data Foundation → Storage
     System → ocs-storagecluster.

 b.  Verify that Status of StorageCluster is Ready and has a green tick mark next to it.

Additional resources

Alerts










18

                        CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT

                    CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION
                                      DEPLOYMENT

3.1. VERIFYING THE STATE OF THE PODS

Verify that all required OpenShift Data Foundation pods are in a Running state after deployment.

Procedure

  1.   Click Workloads → Podsfrom the OpenShift Web Console.

  2.   Select the name of your defined namespace from the Project drop-down list.

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










                    19

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w


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

                                       ux-backend-server-* (1 pod on any
                                       storage node)


Rook-ceph Operator           rook-ceph-operator-*

                             (1 pod on any storage node)

Multicloud Object Gateway              noobaa-operator-* (1 pod on any storage
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


MON                          rook-ceph-mon-*

                             (3 pods distributed across storage nodes)

MGR                          rook-ceph-mgr-*

                             (2 pods on any storage node)

MDS                          rook-ceph-mds-ocs-storagecluster-
                             cephfilesystem-*

                             (2 pods distributed across storage nodes)





20

CHAPTER 3. VERIFYING OPENSHIFT DATA FOUNDATION DEPLOYMENT


Component    Corresponding pods

CSI          cephfs
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










                              21

    Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w


    Component     Corresponding pods

    Monitoring     prometheus-odf-prometheus-* (1 pod
                   on any storage node)

                   prometheus-operator-* (1 pod on any
                   storage node)

                   alertmanager-odf-alertmanager-* (1
                   pod on any storage node)


    odf-blackbox-exporter     odf-blackbox-exporter-* (1 pod on any
                              storage node) NOTE: This does not apply to
                              clusters without storage, for example, Multicloud
                              Object Gateway‑only deployments or external
                              mode deployments.


3.2. VERIFYING THE OPENSHIFT DATA FOUNDATION CLUSTER IS
HEALTHY

Verify that the OpenShift Data Foundation cluster is healthy after installation.

Procedure

      1. In the OpenShift Web Console, click Storage → Data Foundation.

     2. In the Status card of the Overview tab, click Storage System and then click the storage
        system link from the pop up that appears.

     3. In the Status card of the Block and File tab, verify that the Storage Cluster has a green tick.

     4. In the Details card, verify that the cluster information is displayed.
        For more information on the health of the OpenShift Data Foundation cluster using the Block
        and File dashboard, see Monitoring OpenShift Data Foundation .

3.3. VERIFYING THAT THE SPECIFIC STORAGE CLASSES EXIST

Verify that the expected OpenShift Data Foundation storage classes are present after deployment.

Procedure

      1. Click Storage → Storage Classesfrom the left pane of the OpenShift Web Console.

     2. Verify that the following storage classes are created with the OpenShift Data Foundation
        cluster creation:










    22

              CHAPTER 4. VIEW OPENSHIFT DATA FOUNDATION TOPOLOGY

      CHAPTER 4. VIEW OPENSHIFT DATA FOUNDATION
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




          23

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using Red Hat OpenShift Service on AWS w

CHAPTER 5. UNINSTALLING OPENSHIFT DATA FOUNDATION

5.1. UNINSTALLING OPENSHIFT DATA FOUNDATION IN INTERNAL
MODE

Uninstalling OpenShift Data Foundation in internal mode permanently removes the storage cluster, all
associated PVCs, and the operator.

Procedure

1.   To uninstall OpenShift Data Foundation in internal mode, refer to the knowledgebase article on
     Uninstalling OpenShift Data Foundation.










24
