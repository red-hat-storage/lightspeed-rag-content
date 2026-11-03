---
title: "deploying openshift data foundation in external mode"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/deploying_openshift_data_foundation_in_external_mode"
source_file: "deploying_openshift_data_foundation_in_external_mode.pdf"
version: "4.22"
chapter: "deploying_openshift_data_foundation_in_external_mode"
---

                       Red Hat OpenShift Data Foundation
                                      4.22

    Deploying OpenShift Data Foundation in
    external mode

Instructions for deploying OpenShift Data Foundation to use an external Red Hat
                   Ceph Storage cluster and IBM FlashSystem.










    Last Updated: 2026-10-28



Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data
Foundation in external mode

Instructions for deploying OpenShift Data Foundation to use an external Red Hat Ceph Storage
cluster and IBM FlashSystem.

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

Read this document for instructions on installing Red Hat OpenShift Data Foundation to use an
external Red Hat Ceph Storage cluster or IBM FlashSystem.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. OVERVIEW OF DEPLOYING IN EXTERNAL MODE 5
  1.1. DISCONNECTED ENVIRONMENT REQUIREMENTS 5
  1.2. DISASTER RECOVERY REQUIREMENTS 5
  1.3. NETWORK PORTS REQUIRED BETWEEN OPENSHIFT CONTAINER PLATFORM AND CEPH WHEN USING
  EXTERNAL MODE DEPLOYMENT 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE 7
  2.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR 7
  2.2. CREATING AN OPENSHIFT DATA FOUNDATION CLUSTER FOR EXTERNAL CEPH STORAGE SYSTEM 8
  2.3. CONFIGURING TOPOLOGY AWARENESS FOR EXTERNAL MODE 14
  2.4. VERIFYING YOUR OPENSHIFT DATA FOUNDATION INSTALLATION FOR EXTERNAL CEPH STORAGE
  SYSTEM 18
   2.4.1. Verifying the state of the pods 18
   2.4.2. Verifying that the OpenShift Data Foundation cluster is healthy 20
   2.4.3. Verifying that the Multicloud Object Gateway is healthy 21
   2.4.4. Verifying that the storage classes are created and listed 21
   2.4.5. Verifying that Ceph cluster is connected 21
   2.4.6. Verifying that storage cluster is ready 22
   2.4.7. Verifying the the creation of Ceph Object Store CRD 22
   2.4.8. Verifying the state of the pods 22
   2.4.9. Verifying that the OpenShift Data Foundation cluster is healthy 24
   2.4.10. Verifying that the Multicloud Object Gateway is healthy 25
   2.4.11. Verifying that the storage classes are created and listed 25
   2.4.12. Verifying that Ceph cluster is connected 26
   2.4.13. Verifying that storage cluster is ready 26
   2.4.14. Verifying the creation of Ceph Object Store CRD 26

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. DEPLOY OPENSHIFT DATA FOUNDATION USING IBM FLASHSYSTEM 28
  3.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR 28
  3.2. CREATING AN OPENSHIFT DATA FOUNDATION CLUSTER FOR EXTERNAL IBM FLASHSYSTEM
  STORAGE 29

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. UNINSTALLING OPENSHIFT DATA FOUNDATION FROM EXTERNAL STORAGE SYSTEM 36
  4.1. REMOVING MONITORING STACK FROM OPENSHIFT DATA FOUNDATION 39
  4.2. REMOVING OPENSHIFT CONTAINER PLATFORM REGISTRY FROM OPENSHIFT DATA FOUNDATION
        43
  4.3. REMOVING THE CLUSTER LOGGING OPERATOR FROM OPENSHIFT DATA FOUNDATION 44
  4.4. REMOVING EXTERNAL IBM FLASHSYSTEM SECRET 45










    1

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode










2

    PREFACE

PREFACE

This document describes how to deploy Red Hat OpenShift Data Foundation in external mode,
connecting to an existing Red Hat Ceph Storage cluster.










3

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

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

        CHAPTER 1. OVERVIEW OF DEPLOYING IN EXTERNAL MODE

 CHAPTER 1. OVERVIEW OF DEPLOYING IN EXTERNAL MODE

Red Hat OpenShift Data Foundation can make services from an external Red Hat Ceph Storage cluster
or use IBM FlashSystems that are available for consumption through OpenShift Container Platform
clusters running on any platform. See the Planning your deployment guide for more information.

You can configure OpenShift Data Foundation in external mode to use an external IBM Storage Ceph
cluster or IBM FlashSystem cluster using the steps provided in the following sections:

        Deploy OpenShift Data Foundation using Red Hat Ceph Storage.

        Deploy OpenShift Data Foundation using IBM FlashSystem.

For instructions regarding how to install a RHCS cluster, see the installation guide.

1.1. DISCONNECTED ENVIRONMENT REQUIREMENTS

For the list of packages to be mirrored for installing OpenShift Data Foundation in disconnected
environments, see the Disconnected environment section in the Planning your deployment guide.

1.2. DISASTER RECOVERY REQUIREMENTS

Disaster Recovery features supported by Red Hat OpenShift Data Foundation require all of the
following prerequisites to successfully implement a disaster recovery solution:

        A valid Red Hat OpenShift Data Foundation Advanced subscription

        A valid Red Hat Advanced Cluster Management for Kubernetes subscription

For more information, see the knowledgebase article on OpenShift Data Foundation subscriptions.

For detailed disaster recovery solution requirements, see the Configuring OpenShift Data Foundation
Disaster Recovery for OpenShift Workloads guide, and the Requirements and recommendations section
of the Install guide in Red Hat Advanced Cluster Management for Kubernetes documentation.

The following features related to external mode are not supported when using Metro-DR:

        StorageClasses using non-default RADOS namespace

        User created StorageClasses, even when using default RADOS namespace

        Multiple StorageClasses

1.3. NETWORK PORTS REQUIRED BETWEEN OPENSHIFT CONTAINER
PLATFORM AND CEPH WHEN USING EXTERNAL MODE
DEPLOYMENT

For detailed information about the network ports required between OpenShift Container Platform and
Ceph when using external mode deployment, see the Network ports for external mode section in the
OpenShift Data Foundation Architecture guide.

Additional resources

        Planning your deployment

        5

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

Disconnected environment

Deploy OpenShift Data Foundation using Red Hat Ceph Storage

Deploy OpenShift Data Foundation using IBM FlashSystem

Red Hat Ceph Storage installation guide

Knowledgebase article on OpenShift Data Foundation subscriptions

Configuring OpenShift Data Foundation Disaster Recovery for OpenShift Workloads

Red Hat Advanced Cluster Management for Kubernetes - Requirements and recommendations

Network ports for external mode










6

                               CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING
                       RED HAT CEPH STORAGE

Red Hat OpenShift Data Foundation can make services from an external Red Hat Ceph Storage cluster
available for consumption through OpenShift Container Platform clusters. You need to install the
OpenShift Data Foundation operator and then create an OpenShift Data Foundation cluster for the
external Ceph storage system.

2.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.

Prerequisites

     Access to an OpenShift Container Platform cluster using an account with cluster-admin and
     operator installation permissions.

     For additional resource requirements, see the Planning your deployment guide.

             IMPORTANT

                   When you need to override the cluster-wide default node selector for OpenShift
                   Data Foundation, you can use the following command to specify a blank node
                   selector for the openshift-storage namespace (create openshift-storage
                   namespace in this case):

                   $ oc annotate namespace openshift-storage openshift.io/node-selector=

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

      d.  Select Approval Strategy as Automatic or Manual.
          If you select Automatic updates, then the Operator Lifecycle Manager (OLM)
          automatically upgrades the running instance of your Operator without any intervention.

          If you select Manual updates, then the OLM creates an update request. As a cluster
                                                                                                    7

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

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

 2.2. CREATING AN OPENSHIFT DATA FOUNDATION CLUSTER FOR
 EXTERNAL CEPH STORAGE SYSTEM

 Create an OpenShift Data Foundation cluster service in external mode by importing the Ceph cluster
 connection details, which enables OpenShift workloads to consume block, file, and object storage from
 the external Red Hat Ceph Storage cluster.

 You need to create a new OpenShift Data Foundation cluster after you install OpenShift Data
 Foundation operator on OpenShift Container Platform deployed on VMware vSphere or user-
 provisioned bare metal infrastructures.

 Prerequisites

  A valid Red Hat OpenShift Data Foundation Advanced subscription. To know how subscriptions
  for OpenShift Data Foundation work, see knowledgebase article on OpenShift Data Foundation
  subscriptions.

  Ensure the OpenShift Container Platform version is 4.22 or above before deploying OpenShift
  Data Foundation 4.22.

  OpenShift Data Foundation operator must be installed. For more information, see Installing
  OpenShift Data Foundation Operator using the Software Catalog.

  To check the supportability and interoperability of Red Hat Ceph Storage (RHCS) with Red Hat
  OpenShift Data Foundation in external mode, go to the lab Red Hat OpenShift Data Foundation
  Supportability and Interoperability Checker.

        Select Service Type as ODF as Self-Managed Service.

        Select appropriate Version from the drop down.

        On the Versions tab, click the Supported RHCS versions in the External Modetab.

  Red Hat Ceph Storage must have Ceph Dashboard installed and configured. For more
  information, see Ceph Dashboard installation and access .

8

                                   CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

         It is recommended that the external Red Hat Ceph Storage cluster has the PG Autoscaler
         enabled. For more information, see The placement group autoscaler section in the Red Hat
         Ceph Storage documentation.

Procedure

    1.   In the OpenShift Web Console, click Storage → External Systems.

    2.   Click Connect external systems.

    3.   Select Red Hat/IBM Ceph Cluster and click Next.

    4.   In the Connect Red Hat Ceph storage page, provide the necessary information:

         a.  Click on the Download Script link to download the python script for extracting Ceph cluster
             details, or extract it using command line:

              $ oc get cm rook-ceph-external-cluster-script-config -n openshift-storage -o
              jsonpath='{.data.script}' | base64 --decode > ceph-external-cluster-details-exporter.py

         b.  For extracting the Red Hat Ceph Storage (RHCS) cluster details, contact the RHCS
             administrator to run the downloaded python script on a Red Hat Ceph Storage node with
             the admin key.

              i.   Run the following command on the RHCS node to view the list of available arguments:

                   # python3 ceph-external-cluster-details-exporter.py --help

                   You can also run the script from inside a MON container (containerized deployment) or
                   from a MON node (RPM deployment).

                           NOTE

                           Use the yum install cephadm command and then the cephadm
                           command to deploy your RHCS cluster using containers. You must pull
                           the RHCS container images using the cephadm command, rather than
                           using yum for installing the Ceph packages onto nodes. For more
                           information, see RHCS product documentation .

             ii.   To retrieve the external cluster details from the RHCS cluster, choose one of the
                   following two options, either a configuration file or command-line flags.

                   A.  Configuration file
                       Use config-file flag. This stores the parameters used during deployment.

                       In new deployments, you can save the parameters used during deployment in a
                       configuration file. This file can then be used during upgrade to preserve the
                       parameters as well as adding any additional parameters. Use config-file to set the
                       path to the configuration file.

                       Example configuration file saved in /config.ini:

                        [Configurations]
                        format = bash
                        cephfs-filesystem-name = <filesystem-name>

                                                                                                        9

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

     rbd-data-pool-name = <pool_name>
     ...

    Set the path to the config.ini file using config-file:

     # python3 ceph-external-cluster-details-exporter.py --config-file /config.ini

B.  Command-line flags
    Retrieve the external cluster details from the RHCS cluster, and pass the
    parameters for your deployment.

     # python3 ceph-external-cluster-details-exporter.py \
     --rbd-data-pool-name <rbd block pool name> [optional arguments]

    For example:

     # python3 ceph-external-cluster-details-exporter.py --rbd-data-pool-name ceph-
     rbd --monitoring-endpoint xxx.xxx.xxx.xxx --monitoring-endpoint-port xxxx --rgw-
     endpoint xxx.xxx.xxx.xxx:xxxx --run-as-user client.ocs

    RBD parameters

    rbd-data-pool-name
     A mandatory parameter that is used for providing block storage in OpenShift
     Data Foundation.
    rados-namespace
     Divides an RBD data pool into separate logical namespaces, used for creating
     RBD PVC in a radosNamespace. Flags required with rados-namespace are
     restricted-auth-permission and k8s-cluster-name.
    rbd-metadata-ec-pool-name
     (Optional) The name of the erasure coded RBD metadata pool.

    RGW parameters

    rgw-endpoint
     (Optional) This parameter is required only if the object storage is to be
     provisioned through Ceph Rados Gateway for OpenShift Data Foundation.
     Provide the endpoint in the following format: <ip_address>:<port>

                   NOTE

                   A fully-qualified domain name (FQDN) is also supported in the
                   format <FQDN>:<PORT>.

    rgw-pool-prefix
     (Optional) The prefix of the RGW pools. If not specified, the default prefix is
     default.
    rgw-tls-cert-path
     (Optional) The file path of the RADOS Gateway endpoint TLS certificate.

        To provide the TLS certificate and RGW endpoint details to the helper
10

CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

  To provide the TLS certificate and RGW endpoint details to the helper
  script, ceph-external-cluster-details-exporter.py, run the following
  command:

   # python3 ceph-external-clustergw-endpoint r-details-exporter.py --rbd-
   data-pool-name <rbd block pool name> --rgw-endpoint <ip_address>:
   <port> --rgw-tls-cert-path <file path containing cert>

  This creates a resource to create a Ceph Object Store CR such as
  Kubernetes Secret containing the TLS certificate. All the intermediate
  certificates including private keys need to be stored in the certificate file.

rgw-skip-tls
 (Optional) This parameter ignores the TLS certification validation when a self-
 signed certificate is provided (not recommended).

Monitoring parameters

monitoring-endpoint
 (Optional) This parameter accepts comma-separated list of IP addresses of
 active and standby mgrs reachable from the OpenShift Container Platform
 cluster. If not provided, the value is automatically populated.
monitoring-endpoint-port
 (Optional) It is the port associated with the ceph-mgr Prometheus exporter
 specified by --monitoring-endpoint. If not provided, the value is automatically
 populated.

Ceph parameters

ceph-conf
 (Optional) The name of the Ceph configuration file.
run-as-user
 (Optional) This parameter is used for providing name for the Ceph user which is
 created by the script. If this parameter is not specified, a default user name
 client.healthchecker is created. The permissions for the new user is set as:

  caps: [mgr] allow command config

  caps: [mon] allow r, allow command quorum_status, allow command version

  caps: [osd] allow rwx pool=RGW_POOL_PREFIX.rgw.meta, allow r
  pool=.rgw.root, allow rw pool= RGW_POOL_PREFIX.rgw.control, allow rx
  pool=RGW_POOL_PREFIX.rgw.log, allow x
  pool=RGW_POOL_PREFIX.rgw.buckets.index


CephFS parameters

cephfs-metadata-pool-name
(Optional) The name of the CephFS metadata pool.
cephfs-data-pool-name
(Optional) The name of the CephFS data pool.


   11

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

 cephfs-filesystem-name
  (Optional) The name of the CephFS filesystem.

 Output parameters

 dry-run
  (Optional) This parameter helps to print the executed commands without
  running them.
 output
  (Optional) The file where the output is required to be stored.

 Multicluster parameters

 k8s-cluster-name
  (Optional) Kubernetes cluster name.
 cluster-name
  (Optional) The Ceph cluster name.
 restricted-auth-permission
  (Optional) This parameter restricts cephCSIKeyrings auth permissions to
  specific pools and clusters. Mandatory flags that need to be set with this are
  rbd-data-pool-name and cluster-name. You can also pass the cephfs-
  filesystem-name flag if there is CephFS user restriction so that permission is
  restricted to a particular CephFS filesystem.

             NOTE

             This parameter must be applied only for the new deployments. To
             restrict csi-users per pool and per cluster, you need to create new
             csi-users and new secrets for those csi-users.

 Example with restricted auth permission:

  # python3 /etc/ceph/create-external-cluster-resources.py --cephfs-filesystem-
  name myfs --rbd-data-pool-name replicapool --cluster-name rookStorage --
  restricted-auth-permission true

 Example of JSON output generated using the python script:

[{"name": "rook-ceph-mon-endpoints", "kind": "ConfigMap", "data": {"data":
"xxx.xxx.xxx.xxx:xxxx", "maxMonId": "0", "mapping": "{}"}}, {"name": "rook-ceph-mon", "kind":
"Secret", "data": {"admin-secret": "admin-secret", "fsid": "<fs-id>", "mon-secret": "mon-
secret"}}, {"name": "rook-ceph-operator-creds", "kind": "Secret", "data": {"userID": "<user-id>",
"userKey": "<user-key>"}}, {"name": "rook-csi-rbd-node", "kind": "Secret", "data": {"userID":
"csi-rbd-node", "userKey": "<user-key>"}}, {"name": "ceph-rbd", "kind": "StorageClass", "data":
{"pool": "<pool>"}}, {"name": "monitoring-endpoint", "kind": "CephCluster", "data":
{"MonitoringEndpoint": "xxx.xxx.xxx.xxx", "MonitoringPort": "xxxx"}}, {"name": "rook-ceph-
dashboard-link", "kind": "Secret", "data": {"userID": "ceph-dashboard-link", "userKey": "<user-
key>"}}, {"name": "rook-csi-rbd-provisioner", "kind": "Secret", "data": {"userID": "csi-rbd-
provisioner", "userKey": "<user-key>"}}, {"name": "rook-csi-cephfs-provisioner", "kind":
"Secret", "data": {"adminID": "csi-cephfs-provisioner", "adminKey": "<admin-key>"}}, {"name":
"rook-csi-cephfs-node", "kind": "Secret", "data": {"adminID": "csi-cephfs-node", "adminKey": "


12

            CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

      <admin-key>"}}, {"name": "cephfs", "kind": "StorageClass", "data": {"fsName": "cephfs",
      "pool": "cephfs_data"}}, {"name": "ceph-rgw", "kind": "StorageClass", "data": {"endpoint":
      "xxx.xxx.xxx.xxx:xxxx", "poolPrefix": "default"}}, {"name": "rgw-admin-ops-user", "kind":
      "Secret", "data": {"accessKey": "<access-key>", "secretKey": "<secret-key>"}}]

  i.   Save the JSON output to a file with .json extension

             NOTE

             For OpenShift Data Foundation to work seamlessly, ensure that the
             parameters (RGW endpoint, CephFS details, RBD pool, and so on) to be
             uploaded using the JSON file remain unchanged on the RHCS external
             cluster after the storage cluster creation.

 ii.   Run the command when there is a multi-tenant deployment in which the RHCS cluster is
       already connected to OpenShift Data Foundation deployment with a lower version.

           # python3 ceph-external-cluster-details-exporter.py --upgrade

        a.  Click Browse to select and upload the JSON file.
            The content of the JSON file is populated and displayed in the text box.

        b.  (Optional) Select In-transit encryption to enable encryption for entire Ceph traffic
            including data using Ceh msgrv2.

              NOTE

              Verify that the external Red Hat Ceph Storage (RHCS) cluster is
              configured for in‑transit encryption if required. Follow the documentation
              for detailed configuration steps before enabling in‑transit encryption.

        c.  (Optional) Select the advanced configuration options that apply to your environment:

iii.   Use Ceph RBD as the default StorageClass
       Sets the RBD StorageClass as the cluster default.

 iv.   Use external PostgreSQL (Technology preview)
       Uses an external PostgreSQL instance instead of the in‑cluster database.

       Additional advanced options may be available in the UI. Select only those that are required
       for the deployment policy and environment.

  v.   Enable automatic backup
       Allows automatic backup for the Multicloud Object Gateway metadata database.

        A.  Select the backup frequency: Daily, Weekly, or Monthly.

        B.  Specify the number of backups to retain.

              NOTE

              This option is disabled if the Use external PostgreSQL option is
              selected.



              13

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

  5.   Review the selections and click Connect.

Verification
To verify the final Status of the installed storage cluster:

  1.   In the OpenShift Web Console, navigate to Storage → Data Foundation → Storage System →
       ocs-external-storagecluster .

  2.   Verify that Status of StorageCluster is Ready and has a green tick.

  3.   To verify that OpenShift Data Foundation, pods and StorageClass are successfully installed, see
       Verifying your external mode OpenShift Data Foundation installation for external Ceph storage
       system.

2.3. CONFIGURING TOPOLOGY AWARENESS FOR EXTERNAL MODE

Topology awareness ensures that data for persistent volumes (PVs) is provisioned from Ceph OSDs
that belong to the same failure domain as the OpenShift node where the application pod runs.

Prerequisites

       You have installed OpenShift Data Foundation in external mode.

       You have access to the external Red Hat Ceph Storage cluster with administrative privileges.

       You know the Ceph failure domain that you want to use for topology-aware provisioning.

       The OpenShift worker nodes that host the application workloads have labels that map to the
       Ceph failure domain values.

Procedure

1.     Identify the Ceph failure domain and the corresponding OpenShift nodes.

       a.  Check the Ceph failure domain:

             $ ceph osd tree

           Example output:

             ID CLASS WEIGHT          TYPE NAME  STATUS REWEIGHT PRI-AFF
             ...
             -9 0.50000     host osd-0
              3 hdd 0.25000     osd.3        up  1.00000 1.00000
             ...

           In this example, the failure domain is host, and the device class is hdd.

       b.  Identify the OpenShift worker nodes and map them to the Ceph failure domains:

             $ oc get nodes

           Example output:



14

            CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

         NAME       STATUS         ROLES      AGE    VERSION
         compute-0  Ready          worker     2d23h  v1.34.1
         compute-1  Ready          worker     2d23h  v1.34.1
         compute-2  Ready          worker     2d23h  v1.34.1

    c.  Verify the node labels that you want to use for topology awareness:

         $ oc get nodes --show-labels

        In this example, the label topology.kubernetes.io/hostname and its values are chosen, for
        example, compute-0, compute-1, and compute-2, when configuring topology-aware
        external cluster resources.

2.  Create separate block pools for each failure domain.
    In this example, a replica-3 pool is created for each failure domain.

       $ ceph osd pool create pool-1-data 32 32 replicated
       $ ceph osd pool set pool-1-data size 3
       $ ceph osd pool create pool-2-data 32 32 replicated
       $ ceph osd pool set pool-2-data size 3
       $ ceph osd pool create pool-3-data 32 32 replicated
       $ ceph osd pool set pool-3-data size 3

    Verify the pools:

       $ ceph osd pool ls detail

3.  Generate the external cluster resources with topology-aware settings.
    Run the following command on the Ceph client:

       $ python3 /etc/rook-external/create-external-cluster-resources.py \
        --rbd-data-pool-name rbd \
        --topology-pools pool-1-data,pool-2-data,pool-3-data \
        --topology-failure-domain-label host \
        --topology-failure-domain-values compute-0,compute-1,compute-2

    where

    --rbd-data-pool-name
       Specifies an existing RBD pool.
    --topology-pools
       Specifies the topology-constrained pools created for each failure domain.
    --topology-failure-domain-label
       Specifies the Ceph failure domain label, for example, host.
    --topology-failure-domain-values
       Specifies the corresponding OpenShift failure domain values, for example, compute-
       0,compute-1,compute-2.

4.  Upload the generated JSON file when creating the external cluster in the OpenShift web
    console.


                                                     15

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

  5. Verify that topology awareness is configured.

       a.  Verify that the operator configuration contains the topology domain label:

           $ oc get operatorconfigs.csi.ceph.io -o yaml | grep topology -A3

           Expected output:

           topology:
            domainLabels:
            - kubernetes.io/hostname

       b.  Verify that the topology-aware storage class contains the topology parameters:

           $ oc get sc ocs-external-storagecluster-ceph-non-resilient-rbd -o yaml

           Example output:

           parameters:
            topologyConstrainedPools: |-
            [
             {
              "poolName": "pool-1-data",
              "domainSegments": [
               {
                "domainLabel": "hostname",
                "value": "compute-0"
               }
              ]
             },
             {
              "poolName": "pool-2-data",
              "domainSegments": [
               {
                "domainLabel": "hostname",
                "value": "compute-1"
               }
              ]
             },
             {
              "poolName": "pool-3-data",
              "domainSegments": [
               {
                "domainLabel": "hostname",
                "value": "compute-2"
               }
              ]
             }
            ]
            topologyFailureDomainLabel: host
            topologyFailureDomainValues: compute-0,compute-1,compute-2
            topologyPools: pool-1-data,pool-2-data,pool-3-data

  6.  Test topology-aware provisioning by creating a persistent volume claim (PVC) and a pod.
      Apply the following YAML:

16

             CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

        apiVersion: v1
        kind: PersistentVolumeClaim
        metadata:
         name: topo-verify-pvc
        spec:
         accessModes:
           - ReadWriteOnce
         resources:
           requests:
            storage: 1Gi
         storageClassName: ocs-external-storagecluster-ceph-non-resilient-rbd
        ---
        apiVersion: v1
        kind: Pod
        metadata:
         name: topo-verify-pod
        spec:
         containers:
         - name: test-container
           image: busybox
           command: [ "sh", "-c", "sleep 3600" ]
           volumeMounts:
           - mountPath: /data
            name: my-volume
         volumes:
         - name: my-volume
           persistentVolumeClaim:
            claimName: topo-verify-pvc

    The WaitForFirstConsumer volume binding mode ensures that the PV is provisioned only after
    the pod is scheduled, so topology constraints can be applied.

7.  Verify data placement in the expected failure domain.

     a.  Get the image name for the provisioned PV:

            $ kubectl get pv $(kubectl get pvc topo-verify-pvc -o jsonpath='{.spec.volumeName}') -o
            jsonpath='{.spec.csi.volumeAttributes.imageName}'

     b.  Check usage from the Ceph side:

            $ rbd du pool-1-data/<image_name>

     c.  Identify the OSDs that store the image data:

            $ ceph osd map pool-1-data rbd_header.<image_name>

         Example output:

            ... -> pg 14.1ed52b0 (14.10) -> up ([4,5,6], p4) acting ([4,5,6], p4)

     d.  Verify that the OSDs belong to the expected failure domains:

            $ ceph osd tree | grep -E "osd.4|osd.5|osd.6"


            17

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

         Example output:

              5  hdd 0.25000           osd.5   up  1.00000 1.00000
              6  hdd 0.25000           osd.6   up  1.00000 1.00000
              4  hdd 0.25000           osd.4   up  1.00000 1.00000

2.4. VERIFYING YOUR OPENSHIFT DATA FOUNDATION
INSTALLATION FOR EXTERNAL CEPH STORAGE SYSTEM

Verify your OpenShift Data Foundation external mode installation to confirm that the storage cluster is
healthy, the expected storage classes are present, and application pods can successfully provision and
use PVCs.

2.4.1. Verifying the state of the pods

 1.   Click Workloads → Podsfrom the left pane of the OpenShift Web Console.

 2.   Select openshift-storage from the Project drop-down list.

                  NOTE

                  If the Show default projects option is disabled, use the toggle button to list all
                  the default projects.

 3.   Verify that the following pods are in running state:
      Pods corresponding to OpenShift Data Foundation components

 Component                                         Corresponding pods

 OpenShift Data Foundation Operator                       ocs-operator-* (1 pod on any worker
                                                          node)

                                                          ocs-metrics-exporter-* (1 pod on any
                                                          worker node)

                                                          odf-operator-controller-manager-* (1
                                                          pod on any worker node)

                                                          odf-console-* (1 pod on any worker node)

                                                          csi-addons-controller-manager-* (1
                                                          pod on any worker node)


Rook-ceph Operator     rook-ceph-operator-*

                       (1 pod on any worker node)









18

CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE


Multicloud Object Gateway     noobaa-operator-* (1 pod on any worker
                              node)

                              noobaa-core-* (1 pod on any worker
                              node)

                              noobaa-db-pg-cluster-1 and noobaa-
                              db-pg-cluster-2 (2 instances of MCG DB
                              pod on any storage node)

                              noobaa-endpoint-* (1 pod on any storage
                              node)

                              cnpg-controller-manager-* (1 pod on
                              any storage node)










                              19

    Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode


    CSI    cephfs
              openshift-
              storage.cephfs.csi.ceph.com-
              ctrlplugin-* (2 pods distributed across
              storage nodes)

              openshift-
              storage.cephfs.csi.ceph.com-
              nodeplugin-* (1 pod on each storage
              node)

               NOTE

               If an MDS is not deployed in the
               external cluster, the openshift-
               storage.cephfs.csi.ceph.com-
               ctrlplugin and openshift-
               storage.cephfs.csi.ceph.com-
               nodeplugin-* pods are not created.

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



2.4.2. Verifying that the OpenShift Data Foundation cluster is healthy

      1. In the OpenShift Web Console, click Storage → Data Foundation.

     2. In the Status card of the Overview tab, click Storage System and then click the storage
        system link from the pop up that appears.

     3. In the Status card of the Block and File tab, verify that the Storage Cluster has a green tick.

     4. In the Details card, verify that the cluster information is displayed.




    20

           CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

2.4.3. Verifying that the Multicloud Object Gateway is healthy

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and then click the storage
     system link from the pop up that appears.

      a.  In the Status card of the Object tab, verify that both Object Service and Data Resiliency
          have a green tick.

      b.  In the Details card, verify that the Multicloud Object Gateway (MCG) information is
          displayed.

           NOTE

           The RADOS Object Gateway is only listed in case RADOS Object Gateway endpoint
           details are included while deploying OpenShift Data Foundation in external mode.

2.4.4. Verifying that the storage classes are created and listed

1.   Click Storage → Storage Classesfrom the left pane of the OpenShift Web Console.

2.   Verify that the following storage classes are created with the OpenShift Data Foundation
     cluster creation:

          ocs-external-storagecluster-ceph-rbd

          ocs-external-storagecluster-ceph-rgw

          ocs-external-storagecluster-cephfs

          openshift-storage.noobaa.io

           NOTE

                    If an MDS is not deployed in the external cluster, ocs-external-storagecluster-
                    cephfs storage class will not be created.

                    If RGW is not deployed in the external cluster, the ocs-external-storagecluster-
                    ceph-rgw storage class will not be created.

           For more information regarding MDS and RGW, see the Red Hat Ceph Storage
           documentation.

Additional resources

     Monitoring OpenShift Data Foundation

     Red Hat Ceph Storage documentation

2.4.5. Verifying that Ceph cluster is connected

Run the following command to verify if the OpenShift Data Foundation cluster is connected to the
external Red Hat Ceph Storage cluster.



                    21

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

 $ oc get cephcluster -n openshift-storage
 NAME                      DATADIRHOSTPATH     MONCOUNT     AGE   PHASE     MESSAGE
 HEALTH     EXTERNAL
 ocs-external-storagecluster-cephcluster        30m     Connected Cluster connected
 successfully   HEALTH_OK  true

2.4.6. Verifying that storage cluster is ready

Run the following command to verify if the storage cluster is ready and the External option is set to
true.

 $ oc get storagecluster -n openshift-storage
 NAME                AGE   PHASE EXTERNAL     CREATED AT     VERSION
 ocs-external-storagecluster 30m Ready true 2021-11-17T09:09:52Z /4.22/.0

2.4.7. Verifying the the creation of Ceph Object Store CRD

Run the following command to verify if the Ceph Object Store CRD is created in the external Red Hat
Ceph Storage cluster.

 $ oc get cephobjectstore -n openshift-storage"
 NAME           PHASE   ENDPOINT     SECUREENDPOINT     AGE
 object-store1  Ready   <http://IP/FQDN:port>        15m

2.4.8. Verifying the state of the pods

Verify that all expected OpenShift Data Foundation pods are running in the openshift-storage
namespace after deploying in external mode.

Procedure

   1.   Click Workloads → Podsfrom the left pane of the OpenShift Web Console.

   2.   Select openshift-storage from the Project drop-down list.

                     NOTE

                     If the Show default projects option is disabled, use the toggle button to list all
                     the default projects.

   3.   Verify that the following pods are in running state:
        Pods corresponding to OpenShift Data Foundation components

Component                      Corresponding pods










22

CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE


OpenShift Data Foundation Operator     ocs-operator-* (1 pod on any worker
                                       node)

                                       ocs-metrics-exporter-* (1 pod on any
                                       worker node)

                                       odf-operator-controller-manager-* (1
                                       pod on any worker node)

                                       odf-console-* (1 pod on any worker node)

                                       csi-addons-controller-manager-* (1
                                       pod on any worker node)



Rook-ceph Operator           rook-ceph-operator-*

                             (1 pod on any worker node)

Multicloud Object Gateway              noobaa-operator-* (1 pod on any worker
                                       node)

                                       noobaa-core-* (1 pod on any worker
                                       node)

                                       noobaa-db-pg-cluster-1 and noobaa-
                                       db-pg-cluster-2 (2 instances of MCG DB
                                       pod on any storage node)

                                       noobaa-endpoint-* (1 pod on any storage
                                       node)

                                       cnpg-controller-manager-* (1 pod on
                                       any storage node)










                                       23

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode


CSI    cephfs
          openshift-
          storage.cephfs.csi.ceph.com-
          ctrlplugin-* (2 pods distributed across
          storage nodes)

          openshift-
          storage.cephfs.csi.ceph.com-
          nodeplugin-* (1 pod on each storage
          node)

           NOTE

           If an MDS is not deployed in the
           external cluster, the openshift-
           storage.cephfs.csi.ceph.com-
           ctrlplugin and openshift-
           storage.cephfs.csi.ceph.com-
           nodeplugin-* pods are not created.

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



2.4.9. Verifying that the OpenShift Data Foundation cluster is healthy

Verify the health status of the OpenShift Data Foundation storage cluster in the Block and File tab of
the OpenShift web console after deploying in external mode.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and then click the storage
     system link from the pop up that appears.

24

          CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

3.   In the Status card of the Block and File tab, verify that the Storage Cluster has a green tick.

4.   In the Details card, verify that the cluster information is displayed.
     For more information on the health of OpenShift Data Foundation cluster using the Block and
     File dashboard, see the Monitoring OpenShift Data Foundation documentation.

Additional resources

     Monitoring OpenShift Data Foundation

2.4.10. Verifying that the Multicloud Object Gateway is healthy

Verify that the Multicloud Object Gateway is healthy and both Object Service and Data Resiliency show
green status after deploying OpenShift Data Foundation in external mode.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and then click the storage
     system link from the pop up that appears.

      a.  In the Status card of the Object tab, verify that both Object Service and Data Resiliency
          have a green tick.

      b.  In the Details card, verify that the Multicloud Object Gateway (MCG) information is
          displayed.

                      NOTE

                      The RADOS Object Gateway is only listed in case RADOS Object Gateway
                      endpoint details are included while deploying OpenShift Data Foundation in
                      external mode.

          For more information on the health of OpenShift Data Foundation cluster using the object
          dashboard, see the Monitoring OpenShift Data Foundation documentation.

Additional resources

     Monitoring OpenShift Data Foundation

2.4.11. Verifying that the storage classes are created and listed

Verify that the expected OpenShift Data Foundation storage classes are created and listed after
deploying in external mode.

Procedure

1.   Click Storage → Storage Classesfrom the left pane of the OpenShift Web Console.

2.   Verify that the following storage classes are created with the OpenShift Data Foundation
     cluster creation:

         ocs-external-storagecluster-ceph-rbd


                      25

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

          ocs-external-storagecluster-ceph-rgw

          ocs-external-storagecluster-cephfs

          openshift-storage.noobaa.io

                         NOTE

                          If an MDS is not deployed in the external cluster, ocs-external-
                          storagecluster-cephfs storage class will not be created.

                          If RGW is not deployed in the external cluster, the ocs-external-
                          storagecluster-ceph-rgw storage class will not be created.

                         For more information regarding MDS and RGW, see the Red Hat Ceph
                         Storage documentation.

Additional resources

     Red Hat Ceph Storage documentation

2.4.12. Verifying that Ceph cluster is connected

Verify that the OpenShift Data Foundation cluster is connected to the external Red Hat Ceph Storage
cluster after deploying in external mode.

Procedure

1.   Run the following command to verify if the OpenShift Data Foundation cluster is connected to
     the external Red Hat Ceph Storage cluster:

         $ oc get cephcluster -n openshift-storage
         NAME                            DATADIRHOSTPATH    MONCOUNT AGE PHASE
         MESSAGE              HEALTH                EXTERNAL
         ocs-external-storagecluster-cephcluster        30m Connected               Cluster
         connected successfully HEALTH_OK           true

2.4.13. Verifying that storage cluster is ready

Verify that the external OpenShift Data Foundation storage cluster is ready and the External flag is set
to true after deployment.

Procedure

1.   Run the following command to verify if the storage cluster is ready and the External option is
     set to true:

         $ oc get storagecluster -n openshift-storage
         NAME                AGE         PHASE EXTERNAL     CREATED AT              VERSION
    ocs-external-storagecluster 30m Ready true 2021-11-17T09:09:52Z /4.22/.0

2.4.14. Verifying the creation of Ceph Object Store CRD


26

             CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION USING RED HAT CEPH STORAGE

Verify that the Ceph Object Store CRD has been created in the external Red Hat Ceph Storage cluster
after deploying OpenShift Data Foundation in external mode.

Procedure

1.   Run the following command to verify if the Ceph Object Store CRD is created in the external
     Red Hat Ceph Storage cluster:

         $ oc get cephobjectstore -n openshift-storage
         NAME           PHASE   ENDPOINT     SECUREENDPOINT AGE
         object-store1  Ready   <http://IP/FQDN:port>      15m










                                27

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

CHAPTER 3. DEPLOY OPENSHIFT DATA FOUNDATION USING
                       IBM FLASHSYSTEM

OpenShift Data Foundation can use IBM FlashSystem storage available for consumption through
OpenShift Container Platform clusters. You need to install the OpenShift Data Foundation operator and
then create an OpenShift Data Foundation cluster for IBM FlashSystem storage.

3.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION
OPERATOR

You can install Red Hat OpenShift Data Foundation Operator using the Red Hat OpenShift Container
Platform Software Catalog.

Prerequisites

     Access to an OpenShift Container Platform cluster using an account with cluster-admin and
     operator installation permissions.

     For additional resource requirements, see the Planning your deployment guide.

             IMPORTANT

                   When you need to override the cluster-wide default node selector for OpenShift
                   Data Foundation, you can use the following command to specify a blank node
                   selector for the openshift-storage namespace (create openshift-storage
                   namespace in this case):

                   $ oc annotate namespace openshift-storage openshift.io/node-selector=

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

      d.  Select Approval Strategy as Automatic or Manual.
          If you select Automatic updates, then the Operator Lifecycle Manager (OLM)
          automatically upgrades the running instance of your Operator without any intervention.

          If you select Manual updates, then the OLM creates an update request. As a cluster
28

                                CHAPTER 3. DEPLOY OPENSHIFT DATA FOUNDATION USING IBM FLASHSYSTEM

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

3.2. CREATING AN OPENSHIFT DATA FOUNDATION CLUSTER FOR
EXTERNAL IBM FLASHSYSTEM STORAGE

You need to create a new OpenShift Data Foundation cluster after you install the OpenShift Data
Foundation operator on the OpenShift Container Platform.

Prerequisites

 A valid Red Hat OpenShift Data Foundation Advanced subscription. For more information, see
 the knowledgebase article on OpenShift Data Foundation subscriptions .

 For Red Hat Enterprise Linux® operating system, ensure that there is iSCSI connectivity and
 then configure Linux multipath devices on the host.

 For Red Hat Enterprise Linux CoreOS or when the packages are already installed, configure
 Linux multipath devices on the host.

 Ensure to configure each worker with storage connectivity according to your storage system
 instructions. For the latest supported FlashSystem storage systems and versions, see IBM ODF
 FlashSystem driver documentation.

 Make sure that that the operator, ibm-storage-odf-operator is installed with the following
 subscription:

      apiVersion: operators.coreos.com/v1alpha1
      kind: Subscription
      metadata:
       name: ibm-storage-odf-operator
       namespace: openshift-storage
      spec:
       channel: stable-v1.8
       installPlanApproval: Automatic
       name: ibm-storage-odf-operator
       source: certified-operators
       sourceNamespace: openshift-marketplace

                                                                                               29

    Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode



Procedure

  1.   In the OpenShift Web Console, click Storage → External Systems.

  2.   Click Connect external systems.

  3.   Select IBM FlashSystem and click Next.

  4.   In the Create storage class page, provide the following information:

       a.  Enter the following details of IBM FlashSystem connection:

                 IP address

                 User name

                 Password

                 Pool name

       b.  Select thick or thin for the Volume mode.

       c.  Click Next.

  5.   In the Capacity and nodes page, provide the necessary details:

       a.  Select a value for Requested capacity.
           The available options are 0.5 TiB, 2 TiB, and 4 TiB. The requested capacity is dynamically
           allocated on the infrastructure storage class.

       b.  Select at least three nodes in three different zones.
           It is recommended to start with at least 14 CPUs and 34 GiB of RAM per node. If the nodes
           selected do not match the OpenShift Data Foundation cluster requirement of an
           aggregated 30 CPUs and 72 GiB of RAM, a minimal cluster will be deployed. For minimum
           starting node requirements, see the Resource requirements section in the Planning guide.

       c.  Click Next.

    6. Optional: In the Security and network page, provide the necessary details:

       a.  To enable encryption, select Enable data encryption for block and file storage.

            i.   Choose any one or both Encryption level:

                      Cluster-wide encryption to encrypt the entire cluster (block and file).

                      StorageClass encryption to create encrypted persistent volume (block only) using
                      encryption enabled storage class.

           ii.   Select the Connect to an external key management servicecheckbox. This is optional
                 for cluster-wide encryption.

                  A.  Key Management Service Provider is set to Vault by default.

                  B.  Enter Vault Service Name, host Address of Vault server ('https://<hostname or
                      ip>'), Port number, and Token.


    30

                        CHAPTER 3. DEPLOY OPENSHIFT DATA FOUNDATION USING IBM FLASHSYSTEM

            iii.   Expand Advanced Settings to enter additional settings and certificate details based on
                   your Vault configuration:

                    A.  Enter the Key Value secret path in the Backend Path that is dedicated and unique
                        to OpenShift Data Foundation.

                    B.  Optional: Enter TLS Server Name and Vault Enterprise Namespace.

                    C.  Provide CA Certificate, Client Certificate, and Client Private Key by uploading the
                        respective PEM encoded certificate file.

             iv.   Click Save.

    v. Select Default (SDN) if you are using a single network or Custom (Multus) if you are
    using multiple network interfaces.

                    A.  Select a Public Network Interface from the dropdown.

                    B.  Select a Cluster Network Interface from the dropdown. NOTE: If you are using only
                        one additional network interface, select the single NetworkAttachementDefinition,
                        that is, ocs-public-cluster for the Public Network Interface, and leave the Cluster
                        Network Interface blank.

             vi.   Click Next.

      b.    To enable in-transit encryption, select In-transit encryption.

              i.   Select a Network.

             ii.   Click Next.

 7.  In the Review and create page, review if all the details are correct:

            To modify any configuration settings, click Back to go back to the previous configuration
            page.

 8.  Click Create StorageSystem.

Verification

Verifying the state of the pods

      1.   Click Workloads → Pods from the left pane of the OpenShift Web Console.

     2.    Select openshift-storage from the Project drop-down list.

                         NOTE

                         If the Show default projects option is disabled, use the toggle button to list
                         all the default projects.

           Pods corresponding to OpenShift Data Foundation components







                         31

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode


Component        Corresponding pods

OpenShift Data Foundation   ocs-operator-* (1 pod on any worker node)
Operator
                            ocs-metrics-exporter-* (1 pod on any worker node)

                            odf-operator-controller-manager-* (1 pod on any worker node)

                            odf-console-* (1 pod on any worker node)

                            csi-addons-controller-manager-* (1 pod on any worker node)



ibm-storage-odf-            ibm-storage-odf-operator-* (2 pods on any worker nodes)
operator
                            ibm-odf-console-*


ibm-flashsystem-           ibm-flashsystem-storage-* (1 pod on any worker node)
storage

rook-ceph Operator         rook-ceph-operator-* (1 pod on any worker node)

Multicloud Object Gateway   noobaa-operator-* (1 pod on any worker node)

                            noobaa-core-* (1 pod on any worker node)

                            noobaa-db-pg-cluster-1 and noobaa-db-pg-cluster-2 (2
                            instances of MCG DB pod on any storage node)

                            noobaa-endpoint-* (1 pod on any storage node)

                            cnpg-controller-manager-* (1 pod on any storage node)


CSI                        ibm-block-csi-* (1 pod on any worker node)



Verifying that the OpenShift Data Foundation cluster is healthy

  1.   In the Web Console, click Storage → Data Foundation.

2.     In the Status card of the Overview tab, verify that the Storage System has a green tick mark.

3.     In the Details card, verify that the cluster information is displayed.

For more information on the health of the OpenShift Data Foundation cluster using the Block and File
dashboard, see Monitoring OpenShift Data Foundation .

Verifying that the Multicloud Object Gateway is healthy

1.     In the Web Console, click Storage → Data Foundation.

32

                                         CHAPTER 3. DEPLOY OPENSHIFT DATA FOUNDATION USING IBM FLASHSYSTEM

  2.  In the Status card of the Overview tab, click Storage System and then click the storage
      system link from the pop up.

  3.  In the Status card of the Object tab, verify that both Object Service and Data Resiliency
      have a green tick.

  4.  In the Details card, verify that the MCG information is displayed.

For more information on the health of the OpenShift Data Foundation cluster using the object
dashboard, see Monitoring OpenShift Data Foundation .

Verifying that IBM FlashSystem is connected and the storage cluster is ready

      Run the following command to verify if the OpenShift Data Foundation cluster is connected
      to the external IBM FlashSystem.

 $ oc get flashsystemclusters.odf.ibm.com
 NAME                       AGE   PHASE  CREATED AT
  ibm-flashsystem-storage-fab3p-60 46h Ready 2025-06-30T12:17:06Z

Verifying the subscription of the IBM operator

      Run the following command to verify the subscription:

 $ oc get subscriptions.operators.coreos.com
 NAME                                        PACKAGE                                SOURCE
 CHANNEL
 cephcsi-operator-stable-4.22-redhat-operators-openshift-marketplace                cephcsi-operator
 redhat-operators       stable-4.22
 ibm-block-csi-operator-stable-certified-operators-openshift-marketplace           ibm-block-csi-operator
 certified-operators    stable
 ibm-storage-odf-operator                    ibm-storage-odf-operator                        ibm-storage-
 odf-catalog     stable-v1.7
 mcg-operator-stable-4.22-redhat-operators-openshift-marketplace                   mcg-operator
 redhat-operators       stable-4.22
 ocs-client-operator-stable-4.22-redhat-operators-openshift-marketplace             ocs-client-operator
 redhat-operators       stable-4.22
 ocs-operator-stable-4.22-redhat-operators-openshift-marketplace      ocs-operator
 redhat-operators       stable-4.22
 odf-csi-addons-operator-stable-4.22-redhat-operators-openshift-marketplace         odf-csi-addons-
 operator  redhat-operators     stable-4.22
 odf-dependencies                            odf-dependencies                            redhat-operators
 stable-4.22
 odf-operator                                                       odf-operator   redhat-operators
 stable-4.22
 odf-prometheus-operator-stable-4.22-redhat-operators-openshift-marketplace             odf-prometheus-
 operator  redhat-operators     stable-4.22
 recipe-stable-4.22-redhat-operators-openshift-marketplace      recipe                            redhat-
 operators     stable-4.22
 rook-ceph-operator-stable-4.22-redhat-operators-openshift-marketplace               rook-ceph-operator
 redhat-operators       stable-4.22          odf-operator                                          odf-
 catalogsource     alpha

Verifying the CSVs


                                                                                    33

    Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

       Run the following command to verify that the CSVs are in the succeeded state.

    $ oc get csv
    NAME                          DISPLAY                         VERSION   REPLACES
    PHASE
    cephcsi-operator.v4.22-rhodf           CephCSI operator                4.22-rhodf  cephcsi-
    operator.v4.22-rhodf       Succeeded
    ibm-block-csi-operator.v1.12.4      IBM block storage CSI driver operator   1.12.4
    Succeeded
    ibm-storage-odf-operator.v1.7.1        IBM Storage ODF operator          1.7.1          ibm-storage-
    odf-operator.v1.7.0       Succeeded
    mcg-operator.v4.22-rhodf         NooBaa Operator                     4.22-rhodf   mcg-
    operator.v4.22-rhodf          Succeeded
    ocs-client-operator.v4.22-rhodf      OpenShift Data Foundation Client  4.22-rhodf        ocs-client-
    operator.v4.22-rhodf      Succeeded
    ocs-operator.v4.22-rhodf         OpenShift Container Storage            4.22-rhodf  ocs-
    operator.v4.22-rhodf          Succeeded
    odf-csi-addons-operator.v4.22-rhodf     CSI Addons                     4.22-rhodf    odf-csi-addons-
    operator.v4.22-rhodf  Succeeded
    odf-dependencies.v4.22-rhodf        Data Foundation Dependencies                 4.22-rhodf  odf-
    dependencies.v4.22-rhodf         Succeeded
    odf-operator.v4.22-rhodf         OpenShift Data Foundation             4.22-rhodf      odf-
    operator.v4.22-rhodf          Succeeded
    odf-prometheus-operator.v4.22-rhodf     Prometheus Operator            4.22-rhodf       odf-
    prometheus-operator.v4.22-rhodf  Succeeded
    recipe.v4.22-rhodf            Recipe        4.22-rhodf                    recipe.v4.22-rhodf
    Succeeded
    rook-ceph-operator.v4.22-rhodf      Rook-Ceph

Verifying the IBM operator and CSI pods

       Run the following command to verify the IBM operator and CSI pods:

    $ oc get pods
    NAME                                              READY  STATUS          RESTARTS            AGE
    ceph-csi-controller-manager-65f4bc4cd5-9qbp4        2/2     Running                     0        46h
    csi-addons-controller-manager-757d9c6748-94nrt      2/2              Running            0        46h
    openshift-storage.cephfs.csi.ceph.com-ctrlplugin-b642k                              3/3      Running
    0     46h
    openshift-storage.cephfs.csi.ceph.com-ctrlplugin-l75pt                             3/3    Running
    0     46h
    openshift-storage.cephfs.csi.ceph.com-nodeplugin-847f7c54d8-52nz9                   7/7      Running
    0     46h
    openshift-storage.cephfs.csi.ceph.com-nodeplugin-847f7c54d8-xtp5j                  7/7    Running
    0     46h
    openshift-storage.cephfs.csi.ceph.com-ctrlplugin-vh485                              3/3      Running
    0     46h
    openshift-storage.rbd.csi.ceph.com-ctrlplugin-9hss7                               4/4     Running
    0     46h
    openshift-storage.rbd.csi.ceph.com-ctrlplugin-9q9qd                               4/4     Running
    0     46h
    openshift-storage.rbd.csi.ceph.com-ctrlplugin-9vkqh                               4/4     Running
    0     46h
    openshift-storage.rbd.csi.ceph.com-nodeplugin-7f4b8d946f-ngz65                    7/7     Running
    0     46h


    34

                                                              CHAPTER 3. DEPLOY OPENSHIFT DATA FOUNDATION USING IBM FLASHSYSTEM

    openshift-storage.rbd.csi.ceph.com-nodeplugin-7f4b8d946f-x8l8t                                                   7/7          Running
    0  46h
    ibm-block-csi-controller-0        8/8                   Running                                             0       41h
    ibm-block-csi-node-6vmpq                          3/3     Running                                           0            46h
    ibm-block-csi-node-fw7xc                          3/3    Running                                              0        41h
    ibm-block-csi-node-k57sv                          3/3    Running                                              0        46h
    ibm-block-csi-operator-867c54df5d-2t968                  1/1                                   Running                  0     46h
    ibm-flashsystem-storage-deployment-5958ffd56b-tj6zw                                            1/1   Running                 0
    46h
    ibm-odf-console-6d64dcc46c-bkknx                        1/1                                  Running               0         41h
    ibm-storage-odf-operator-787c67c7bf-nzp9w                  2/2                                 Running                   0    41h
    noobaa-core-0        2/2                          Running                                               0           46h
    noobaa-db-pg-0        1/1                              Running                                            0         46h
    noobaa-endpoint-5d94d888fc-bj28x                        1/1                                  Running               0         46h
    noobaa-operator-67b569fc76-rk99d                        1/1                                  Running               0         46h
    ocs-metrics-exporter-6f7f79fcf9-xm6pq                   2/3                                  Running               0         46h
    ocs-operator-676c964c58-nsb7g                          1/1    Running                                          0           46h
    odf-console-75dc785696-n97sk                          1/1                                      Running         0                46h
    odf-operator-controller-manager-785ddcd999-tf245              2/2                                  Running                 0     46h
    rook-ceph-crashcollector-worker-node1-b95b5cd54-7qbhb                                          1/1         Running            0
    46h
    rook-ceph-crashcollector-worker-node2-757d746d5f-jsfrx                                         1/1   Running                 0
    46h
    rook-ceph-crashcollector-worker-node3-6d4f8f7848-xvvw2                                         1/1      Running               0
    46h
    rook-ceph-exporter-worker-node1-5dcf74b647-t5rg7                1/1                                Running                 0
    46h
    rook-ceph-exporter-worker-node2-66c694f5b6-n96g5                                               1/1   Running                 0
    46h
    rook-ceph-exporter-worker-node3-58fd7ffcfc-vpzfr             1/1                                     Running             0    46h
    rook-ceph-mds-ocs-storagecluster-cephfilesystem-a-557b4cc9gsx65                                      2/2     Running            0
    46h
    rook-ceph-mds-ocs-storagecluster-cephfilesystem-b-5d8748dfmrhfn                                        2/2   Running           0
    46h
    rook-ceph-mgr-a-f56c6866b-tzl2s                       3/3                                      Running         0                46h
    rook-ceph-mgr-b-796d5f69bf-f54qr                       3/3    Running                                            0         46h
    rook-ceph-mon-a-6b5dc9dd6c-f4dvn                                                         2/2   Running                  0     46h
    rook-ceph-mon-b-8d458fbdf-wjgzm                         2/2                                  Running               0         46h
    rook-ceph-mon-c-5b78749887-mcpjl                                                         2/2   Running                  0     46h
    rook-ceph-operator-578dfdd8dd-mn9vn                       1/1                                  Running                 0      46h
    rook-ceph-osd-0-754499c4bd-2nrs2                        2/2                                  Running               0         46h
    rook-ceph-osd-1-5566ff6d99-dc6l4                       2/2    Running                                            0         46h
    rook-ceph-osd-2-8567c9f9f7-wr66b                        2/2                                  Running               0         46h
rook-ceph-osd-prepare-667012a10f7f0800994523ce6731b5d1-n7fgb 0/1 Completed
    0  46h
rook-ceph-osd-prepare-ab6be2d13c76aad7123b793996562bf0-wpqsv 0/1 Completed
    0  46h
    rook-ceph-osd-prepare-baf6cafc480b4e046c337c49601633b8-zllnx                                            0/1         Completed    0
    46h
    rook-ceph-rgw-ocs-storagecluster-cephobjectstore-a-5f4dfd9wlbvn                                  2/2         Running            0
    46h
    ux-backend-server-db4c5b7bd-hn9jx                                                        2/2   Running              0         46h







                                                                                                                                  35

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

CHAPTER 4. UNINSTALLING OPENSHIFT DATA FOUNDATION
  FROM EXTERNAL STORAGE SYSTEM

Use the steps in this section to uninstall OpenShift Data Foundation. Uninstalling OpenShift Data
Foundation does not remove the RBD pool from the external cluster, or uninstall the external Red Hat
Ceph Storage cluster.

Uninstall Annotations

Annotations on the Storage Cluster are used to change the behavior of the uninstall process. To define
the uninstall behavior, the following two annotations have been introduced in the storage cluster:

 uninstall.ocs.openshift.io/cleanup-policy: delete

 uninstall.ocs.openshift.io/mode: graceful

  NOTE

  The uninstall.ocs.openshift.io/cleanup-policy is not applicable for external mode.

The below table provides information on the different values that can used with these annotations:

uninstall.ocs.openshift.io uninstall annotations descriptions


Annotation       Value        Default     Behavior

cleanup-policy   delete       Yes         Rook cleans up the
                                          physical drives and the
                                          DataDirHostPath

cleanup-policy   retain       No          Rook does not clean up
                                          the physical drives and
                                          the DataDirHostPath

mode             graceful     Yes         Rook and NooBaa
                                          pauses the uninstall
                                          process until the PVCs
                                          and the OBCs are
                                          removed by the
                                          administrator/user

mode             forced       No          Rook and NooBaa
                                          proceeds with uninstall
                                          even if PVCs/OBCs
                                          provisioned using Rook
                                          and NooBaa exist
                                          respectively


You can change the uninstall mode by editing the value of the annotation by using the following
commands:



36

       CHAPTER 4. UNINSTALLING OPENSHIFT DATA FOUNDATION FROM EXTERNAL STORAGE SYSTEM

  $ oc annotate storagecluster ocs-external-storagecluster -n openshift-storage
  uninstall.ocs.openshift.io/mode="forced" --overwrite
  storagecluster.ocs.openshift.io/ocs-external-storagecluster annotated

  Prerequisites

       Ensure that the OpenShift Data Foundation cluster is in a healthy state. The uninstall process
       can fail when some of the pods are not terminated successfully due to insufficient resources or
       nodes. In case the cluster is in an unhealthy state, contact Red Hat Customer Support before
       uninstalling OpenShift Data Foundation.

       Ensure that applications are not consuming persistent volume claims (PVCs) or object bucket
       claims (OBCs) using the storage classes provided by OpenShift Data Foundation.

Procedure

  1.   Delete the volume snapshots that are using OpenShift Data Foundation.

       a.  List the volume snapshots from all the namespaces

               $ oc get volumesnapshot --all-namespaces

       b.  From the output of the previous command, identify and delete the volume snapshots that
           are using OpenShift Data Foundation.

               $ oc delete volumesnapshot <VOLUME-SNAPSHOT-NAME> -n <NAMESPACE>

  2.   Delete PVCs and OBCs that are using OpenShift Data Foundation.
       In the default uninstall mode (graceful), the uninstaller waits till all the PVCs and OBCs that use
       OpenShift Data Foundation are deleted.

       If you wish to delete the Storage Cluster without deleting the PVCs beforehand, you may set
       the uninstall mode annotation to "forced" and skip this step. Doing so will result in orphan PVCs
       and OBCs in the system.

       a.  Delete OpenShift Container Platform monitoring stack PVCs using OpenShift Data
           Foundation.
           See Removing monitoring stack from OpenShift Data Foundation

       b.  Delete OpenShift Container Platform Registry PVCs using OpenShift Data Foundation.
           Removing OpenShift Container Platform registry from OpenShift Data Foundation

       c.  Delete OpenShift Container Platform logging PVCs using OpenShift Data Foundation.
           Removing the cluster logging operator from OpenShift Data Foundation

       d.  Delete other PVCs and OBCs provisioned using OpenShift Data Foundation.

               Given below is a sample script to identify the PVCs and OBCs provisioned using
               OpenShift Data Foundation. The script ignores the PVCs and OBCs that are used
               internally by OpenShift Data Foundation.

                #!/bin/bash

                RBD_PROVISIONER="openshift-storage.rbd.csi.ceph.com"
                CEPHFS_PROVISIONER="openshift-storage.cephfs.csi.ceph.com"

                    37

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

       NOOBAA_PROVISIONER="openshift-storage.noobaa.io/obc"
       RGW_PROVISIONER="openshift-storage.ceph.rook.io/bucket"

       NOOBAA_DB_PG_1_PVC="noobaa-db-pg-cluster-1"
       NOOBAA_DB_PG_2_PVC="noobaa-db-pg-cluster-2"
       NOOBAA_BACKINGSTORE_PVC="noobaa-default-backing-store-noobaa-pvc"

       # Find all the OCS StorageClasses
       OCS_STORAGECLASSES=$(oc get storageclasses | grep -e
       "$RBD_PROVISIONER" -e "$CEPHFS_PROVISIONER" -e
       "$NOOBAA_PROVISIONER" -e "$RGW_PROVISIONER" | awk '{print $1}')

       # List PVCs in each of the StorageClasses
       for SC in $OCS_STORAGECLASSES
       do
                echo
       "====================================================================
       =="
                echo "$SC StorageClass PVCs and OBCs"
                echo
       "====================================================================
       =="
                oc get pvc --all-namespaces --no-headers 2>/dev/null | grep $SC | grep -v -e
       "$NOOBAA_DB_PVC" -e "$NOOBAA_BACKINGSTORE_PVC"
                oc get obc --all-namespaces --no-headers 2>/dev/null | grep $SC
                echo
       done

      Delete the OBCs.

       $ oc delete obc <obc name> -n <project name>

      Delete the PVCs.

       $ oc delete pvc <pvc name> -n <project-name>

      Ensure that you have removed any custom backing stores, bucket classes, and so on
      that are created in the cluster.

3.  Delete the Storage Cluster object and wait for the removal of the associated resources.

     $ oc delete -n openshift-storage storagecluster --all --wait=true

4.  Delete the namespace and wait until the deletion is complete. You will need to switch to another
    project if openshift-storage is the active project.
    For example:

     $ oc project default
     $ oc delete project openshift-storage --wait=true --timeout=5m

    The project is deleted if the following command returns a NotFound error.

     $ oc get project openshift-storage



38

      CHAPTER 4. UNINSTALLING OPENSHIFT DATA FOUNDATION FROM EXTERNAL STORAGE SYSTEM

                     NOTE

                     While uninstalling OpenShift Data Foundation, if the namespace is not deleted
                     completely and remains in Terminating state, perform the steps in
                     Troubleshooting and deleting remaining resources during Uninstall to identify
                     objects that are blocking the namespace from being terminated.

 5.   Confirm all PVs provisioned using OpenShift Data Foundation are deleted. If there is any PV left
      in the Released state, delete it.

          $ oc get pv
          $ oc delete pv <pv name>

 6.   Remove CustomResourceDefinitions.

          $ oc delete crd backingstores.noobaa.io bucketclasses.noobaa.io
          cephblockpools.ceph.rook.io cephclusters.ceph.rook.io cephfilesystems.ceph.rook.io
          cephnfses.ceph.rook.io cephobjectstores.ceph.rook.io cephobjectstoreusers.ceph.rook.io
          noobaas.noobaa.io ocsinitializations.ocs.openshift.io storageclusters.ocs.openshift.io
          cephclients.ceph.rook.io cephobjectrealms.ceph.rook.io cephobjectzonegroups.ceph.rook.io
          cephobjectzones.ceph.rook.io cephrbdmirrors.ceph.rook.io storagesystems.odf.openshift.io --
          wait=true --timeout=5m

 7.   To ensure that OpenShift Data Foundation is uninstalled completely:

       a.  In the OpenShift Container Platform Web Console, click Storage.

       b.  Verify that OpenShift Data Foundation no longer appears under Storage.

4.1. REMOVING MONITORING STACK FROM OPENSHIFT DATA
FOUNDATION

Remove the OpenShift monitoring stack from OpenShift Data Foundation external mode storage
before uninstalling the deployment to reclaim the PVCs used by Prometheus and Alertmanager.

The PVCs that are created as a part of configuring the monitoring stack are in the openshift-
monitoring namespace.

Prerequisites

      PVCs are configured to use the OpenShift Container Platform monitoring stack.
      For information, see configuring monitoring stack.

Procedure

 1.   List the pods and PVCs that are currently running in the openshift-monitoring namespace.

          $ oc get pod,pvc -n openshift-monitoring
          NAME           READY         STATUS     RESTARTS    AGE
          pod/alertmanager-main-0      3/3   Running    0   8d
          pod/alertmanager-main-1      3/3   Running    0   8d
          pod/alertmanager-main-2      3/3   Running    0   8d
          pod/cluster-monitoring-


                                                            39

                                        Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode

     operator-84457656d-pkrxm         1/1                                                  Running    0   8d
     pod/grafana-79ccf6689f-2ll28     2/2                                                Running     0    8d
     pod/kube-state-metrics-
     7d86fb966-rvd9w             3/3     Running   0                                                    8d
     pod/node-exporter-25894      2/2      Running                                                 0       8d
     pod/node-exporter-4dsd7      2/2      Running                                                 0       8d
     pod/node-exporter-6p4zc      2/2      Running                                                 0       8d
     pod/node-exporter-jbjvg      2/2     Running                                                0         8d
     pod/node-exporter-jj4t5       2/2    Running                                              0         6d18h
     pod/node-exporter-k856s      2/2      Running                                                 0      6d18h
     pod/node-exporter-rf8gn      2/2      Running                                               0         8d
     pod/node-exporter-rmb5m                                                        2/2  Running     0    6d18h
     pod/node-exporter-zj7kx      2/2      Running                                               0         8d
     pod/openshift-state-metrics-
     59dbd4f654-4clng            3/3     Running   0                                                    8d
     pod/prometheus-adapter-
     5df5865596-k8dzn              1/1    Running                                              0         7d23h
     pod/prometheus-adapter-
     5df5865596-n2gj9            1/1     Running   0                                                     7d23h
     pod/prometheus-k8s-0         6/6      Running                                               1         8d
     pod/prometheus-k8s-1         6/6      Running                                               1         8d
     pod/prometheus-operator-
     55cfb858c9-c4zd9            1/1     Running   0                                                     6d21h
     pod/telemeter-client-
     78fc8fc97d-2rgfp      3/3          Running  0                                                   8d

     NAME                                           STATUS                                                VOLUME
     CAPACITY     ACCESS MODES          STORAGECLASS                                                         AGE
    persistentvolumeclaim/my-alertmanager-claim-alertmanager-main-0 Bound pvc-0d519c4f-
     15a5-11ea-baa0-026d231574aa         40Gi                                                    RWO     ocs-external-storagecluster-ceph-rbd
    8d
    persistentvolumeclaim/my-alertmanager-claim-alertmanager-main-1 Bound pvc-
     0d5a9825-15a5-11ea-baa0-026d231574aa          40Gi                                                  RWO    ocs-external-
    storagecluster-ceph-rbd 8d
    persistentvolumeclaim/my-alertmanager-claim-alertmanager-main-2 Bound pvc-
     0d6413dc-15a5-11ea-baa0-026d231574aa          40Gi                                                  RWO    ocs-external-
     storagecluster-ceph-rbd   8d
     persistentvolumeclaim/my-prometheus-claim-prometheus-k8s-0                                                 Bound pvc-0b7c19b0-
     15a5-11ea-baa0-026d231574aa         40Gi                                                    RWO     ocs-external-storagecluster-ceph-rbd
     8d
     persistentvolumeclaim/my-prometheus-claim-prometheus-k8s-1                                                 Bound pvc-0b8aed3f-
     15a5-11ea-baa0-026d231574aa         40Gi                                                    RWO     ocs-external-storagecluster-ceph-rbd
     8d

2.  Edit the monitoring configmap.

     $ oc -n openshift-monitoring edit configmap cluster-monitoring-config

    Remove any config sections that reference the OpenShift Data Foundation storage classes as
    shown in the following example and save it.

    Before editing





40

CHAPTER 4. UNINSTALLING OPENSHIFT DATA FOUNDATION FROM EXTERNAL STORAGE SYSTEM



.
.
.
apiVersion: v1
data:
 config.yaml: |
  alertmanagerMain:
  volumeClaimTemplate:
     metadata:
      name: my-alertmanager-claim
     spec:
      resources:
      requests:
         storage: 40Gi
      storageClassName: ocs-external-storagecluster-ceph-rbd
  prometheusK8s:
      volumeClaimTemplate:
     metadata:
      name: my-prometheus-claim
     spec:
      resources:
      requests:
      storage: 40Gi
      storageClassName: ocs-external-storagecluster-ceph-rbd
kind: ConfigMap
metadata:
 creationTimestamp: "2019-12-02T07:47:29Z"
 name: cluster-monitoring-config
 namespace: openshift-monitoring
 resourceVersion: "22110"
 selfLink: /api/v1/namespaces/openshift-monitoring/configmaps/cluster-monitoring-config
 uid: fd6d988b-14d7-11ea-84ff-066035b9efa8


.
.
.



After editing










      41

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode



.
.
.
apiVersion: v1
data:
 config.yaml: |
kind: ConfigMap
metadata:
 creationTimestamp: "2019-11-21T13:07:05Z"
 name: cluster-monitoring-config
 namespace: openshift-monitoring
 resourceVersion: "404352"
 selfLink: /api/v1/namespaces/openshift-monitoring/configmaps/cluster-monitoring-config
 uid: d12c796a-0c5f-11ea-9832-063cd735b81c
.
.
.


    In this example, alertmanagerMain and prometheusK8s monitoring components are using the
    OpenShift Data Foundation PVCs.

3.  List the pods consuming the PVC.
    In this example, the alertmanagerMain and prometheusK8s pods that were consuming the
    PVCs are in the Terminating state. You can delete the PVCs once these pods are no longer
    using OpenShift Data Foundation PVC.

     $ oc get pod,pvc -n openshift-monitoring
     NAME        READY                            STATUS     RESTARTS AGE
     pod/alertmanager-main-0                  3/3   Terminating    0   10h
     pod/alertmanager-main-1                  3/3   Terminating    0   10h
     pod/alertmanager-main-2                  3/3   Terminating    0   10h
    pod/cluster-monitoring-operator-84cd9df668-zhjfn 1/1 Running 0 18h
     pod/grafana-5db6fd97f8-pmtbf             2/2    Running       0    10h
     pod/kube-state-metrics-895899678-z2r9q         3/3      Running    0   10h
     pod/node-exporter-4njxv                 2/2   Running      0    18h
     pod/node-exporter-b8ckz                  2/2   Running     0     11h
     pod/node-exporter-c2vp5                  2/2   Running     0     18h
     pod/node-exporter-cq65n                  2/2   Running     0     18h
     pod/node-exporter-f5sm7                  2/2   Running     0     11h
     pod/node-exporter-f852c                 2/2   Running      0    18h
     pod/node-exporter-l9zn7                 2/2   Running      0    11h
     pod/node-exporter-ngbs8                  2/2   Running     0     18h
     pod/node-exporter-rv4v9                 2/2   Running      0    18h
    pod/openshift-state-metrics-77d5f699d8-69q5x 3/3 Running 0 10h
     pod/prometheus-adapter-765465b56-4tbxx          1/1     Running     0   10h
     pod/prometheus-adapter-765465b56-s2qg2            1/1   Running     0   10h
     pod/prometheus-k8s-0                    6/6   Terminating  1     9m47s
     pod/prometheus-k8s-1                    6/6   Terminating  1     9m47s
     pod/prometheus-operator-cbfd89f9-ldnwc         1/1     Running    0    43m
     pod/telemeter-client-7b5ddb4489-2xfpz         3/3    Running     0     10h

     NAME        STATUS VOLUME
     CAPACITY ACCESS MODES     STORAGECLASS                        AGE

42

         CHAPTER 4. UNINSTALLING OPENSHIFT DATA FOUNDATION FROM EXTERNAL STORAGE SYSTEM

    persistentvolumeclaim/ocs-alertmanager-claim-alertmanager-main-0 Bound pvc-
         2eb79797-1fed-11ea-93e1-0a88476a6a64 40Gi     RWO ocs-external-
    storagecluster-ceph-rbd 19h
    persistentvolumeclaim/ocs-alertmanager-claim-alertmanager-main-1 Bound pvc-
         2ebeee54-1fed-11ea-93e1-0a88476a6a64 40Gi     RWO ocs-external-
    storagecluster-ceph-rbd 19h
    persistentvolumeclaim/ocs-alertmanager-claim-alertmanager-main-2 Bound pvc-2ec6a9cf-
         1fed-11ea-93e1-0a88476a6a64 40Gi     RWO     ocs-external-storagecluster-ceph-rbd
         19h
         persistentvolumeclaim/ocs-prometheus-claim-prometheus-k8s-0     Bound        pvc-3162a80c-
         1fed-11ea-93e1-0a88476a6a64 40Gi     RWO     ocs-external-storagecluster-ceph-rbd
         19h
         persistentvolumeclaim/ocs-prometheus-claim-prometheus-k8s-1     Bound     pvc-
         316e99e2-1fed-11ea-93e1-0a88476a6a64 40Gi     RWO     ocs-external-
         storagecluster-ceph-rbd 19h

4.   Delete relevant PVCs. Make sure you delete all the PVCs that are consuming the storage
     classes.

         $ oc delete -n openshift-monitoring pvc <pvc-name> --wait=true --timeout=5m

4.2. REMOVING OPENSHIFT CONTAINER PLATFORM REGISTRY
FROM OPENSHIFT DATA FOUNDATION

Use this section to clean up the OpenShift Container Platform registry from OpenShift Data
Foundation. If you want to configure an alternative storage, see image registry

The PVCs that are created as a part of configuring OpenShift Container Platform registry are in the
openshift-image-registry namespace.

Prerequisites

     The image registry should have been configured to use an OpenShift Data Foundation PVC.

Procedure

1.   Edit the configs.imageregistry.operator.openshift.io object and remove the content in the
     storage section.

         $ oc edit configs.imageregistry.operator.openshift.io


     Before editing










         43

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation in external mode



.
.
.
storage:
     pvc:
 claim: registry-cephfs-rwx-pvc
.
.
.

After editing


.
.
.
storage:
 emptyDir: {}
.
.
.


      In this example, the PVC is called registry-cephfs-rwx-pvc, which is now safe to delete.

  2.  Delete the PVC.

         $ oc delete pvc <pvc-name> -n openshift-image-registry --wait=true --timeout=5m

4.3. REMOVING THE CLUSTER LOGGING OPERATOR FROM
OPENSHIFT DATA FOUNDATION

Use this section to clean up the cluster logging operator from OpenShift Data Foundation.

The Persistent Volume Claims (PVCs) that are created as a part of configuring the cluster logging
operator are in the openshift-logging namespace.

Prerequisites

      The cluster logging instance should have been configured to use the OpenShift Data
      Foundation PVCs.

Procedure

1.    Remove the ClusterLogging instance in the namespace.

         $ oc delete clusterlogging instance -n openshift-logging --wait=true --timeout=5m

      The PVCs in the openshift-logging namespace are now safe to delete.

2.    Delete the PVCs.

44

        CHAPTER 4. UNINSTALLING OPENSHIFT DATA FOUNDATION FROM EXTERNAL STORAGE SYSTEM

        $ oc delete pvc <pvc-name> -n openshift-logging --wait=true --timeout=5m

        <pvc-name>
        Is the name of the PVC

4.4. REMOVING EXTERNAL IBM FLASHSYSTEM SECRET

You need to clean up the FlashSystem secret from OpenShift Data Foundation while uninstalling. This
secret is created when you configure the external IBM FlashSystem Storage. For more information, see
Creating an OpenShift Data Foundation Cluster for external IBM FlashSystem storage .

Procedure

        Remove the IBM FlashSystem secret by using the following command:

        $ oc delete secret -n openshift-storage ibm-flashsystem-storage










    45
