---
title: "deploying multiple openshift data foundation storage clusters"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/deploying_multiple_openshift_data_foundation_storage_clusters"
source_file: "deploying_multiple_openshift_data_foundation_storage_clusters.pdf"
version: "4.22"
chapter: "deploying_multiple_openshift_data_foundation_storage_clusters"
---

                     Red Hat OpenShift Data Foundation
                                    4.22

                Deploying multiple OpenShift Data Foundation
                              storage clusters

    Instructions on how to deploy an external OpenShift Data Foundation storage cluster
when you already have an existing OpenShift Data Foundation internal storage
                                  cluster.










    Last Updated: 2026-10-22



Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift
Data Foundation storage clusters

Instructions on how to deploy an external OpenShift Data Foundation storage cluster when you
already have an existing OpenShift Data Foundation internal storage cluster.

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

    Read this document for instructions about how to deploy an external OpenShift Data Foundation
    storage cluster when you already have an existing OpenShift Data Foundation internal storage
    cluster.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. OVERVIEW OF MULTIPLE STORAGE CLUSTER DEPLOYMENTS 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. PREPARING TO DEPLOY MULTIPLE OPENSHIFT DATA FOUNDATION STORAGE CLUSTERS
        6

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. DEPLOYING OPENSHIFT DATA FOUNDATION INTERNAL STORAGE CLUSTER 7

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. DEPLOYING OPENSHIFT DATA FOUNDATION EXTERNAL STORAGE CLUSTER 8

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. VERIFYING EXTERNAL OPENSHIFT DATA FOUNDATION STORAGE CLUSTER DEPLOYMENT
        13
  5.1. VERIFYING THE STATE OF THE PODS 13
  5.2. VERIFYING THAT THE OPENSHIFT DATA FOUNDATION CLUSTER IS HEALTHY 14
  5.3. VERIFYING THAT THE MULTICLOUD OBJECT GATEWAY IS HEALTHY 14
  5.4. VERIFYING THAT THE STORAGE CLASSES ARE CREATED AND LISTED 15
  5.5. VERIFYING THAT CEPH CLUSTER IS CONNECTED 15
  5.6. VERIFYING THAT STORAGE CLUSTER IS READY 15
  5.7. VERIFYING THE STATE OF THE PODS 16
  5.8. VERIFYING THAT THE OPENSHIFT DATA FOUNDATION CLUSTER IS HEALTHY 17
  5.9. VERIFYING THAT THE MULTICLOUD OBJECT GATEWAY IS HEALTHY 17
  5.10. VERIFYING THAT THE STORAGE CLASSES ARE CREATED AND LISTED 18
  5.11. VERIFYING THAT CEPH CLUSTER IS CONNECTED 19
  5.12. VERIFYING THAT STORAGE CLUSTER IS READY 19

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 6. MIGRATING APPLICATION WORKLOADS 20










    1

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters










2

    PREFACE

    PREFACE

This document describes how to deploy and manage multiple OpenShift Data Foundation storage
clusters.










3

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters

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

     CHAPTER 1. OVERVIEW OF MULTIPLE STORAGE CLUSTER DEPLOYMENTS

 CHAPTER 1. OVERVIEW OF MULTIPLE STORAGE CLUSTER
                   DEPLOYMENTS

Red Hat OpenShift Data Foundation provides the ability to deploy two storage clusters, where one is
internal mode and the other is in external mode. This can be achieved only with the first cluster installed
as internal in openshift-storage namespace while the second cluster is installed as the external in
openshift-storage-extended namespace. Clusters installed conversely is not currently supported.

Supported platforms

 Bare metal

 VMware VSphere

 OpenStack

 OpenShift Virtualization

 IBM Cloud

 IBM Power

 IBM Z










 5

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters

   CHAPTER 2. PREPARING TO DEPLOY MULTIPLE OPENSHIFT
        DATA FOUNDATION STORAGE CLUSTERS

  Before you begin the deployment of OpenShift Data Foundation using dynamic, local, or external
  storage, ensure that your resource requirements are met. See the Resource requirements section in the
  Planning guide.

        For the list of packages to be mirrored for installing OpenShift Data Foundation in
        disconnected environments, see the Disconnected environment section in the Planning guide.

  Things you should remember before installing multiple OpenShift Data Foundation storage clusters:

        openshift-storage and openshift-storage-extended are the exclusively supported
        namespaces.

        Internal storage cluster is restricted to the OpenShift Data Foundation operator namespace.

        External storage cluster is permissible in both operator and non-operator namespaces.

        Multiple storage clusters are not supported in the same namespace. Hence, the external storage
        system will not be visible under the OpenShift Data Foundation operator page as the operator
        is under openshift-storage namespace and the external storage system is not.

        Customers running external storage clusters in the operator namespace cannot utilize multiple
        storage clusters.

        Multicloud Object Gateway is supported solely within the operator namespace. It is ignored in
        other namespaces.

        RADOS Gateway (RGW) can be in either the operator namespace, a non-operator namespace,
        or both

        Network File System (NFS) is enabled as long as it is enabled for at least one of the clusters.

        Topology is enabled as long as it is enabled for at least one of the clusters.

        Topology domain labels are set as long as the internal cluster is present.

        The Topology view of the cluster is only supported for OpenShift Data Foundation internal
        mode deployments.

        Different multus settings are not supported for multiple storage clusters.










    6

CHAPTER 3. DEPLOYING OPENSHIFT DATA FOUNDATION INTERNAL STORAGE CLUSTER

CHAPTER 3. DEPLOYING OPENSHIFT DATA FOUNDATION
           INTERNAL STORAGE CLUSTER

To deploy and verify OpenShift Data Foundation storage cluster in the internal mode, refer to your
infrastructure specific deployment guides .










7

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters

                               CHAPTER 4. DEPLOYING OPENSHIFT DATA FOUNDATION
                                          EXTERNAL STORAGE CLUSTER

Deploy a second OpenShift Data Foundation storage cluster in external mode to add additional storage
or expand your current internal storage cluster.

Prerequisites

         An OpenShift Data Foundation cluster deployed in internal mode.

         Ensure that both the OpenShift Container Platform and OpenShift Data Foundation are
         upgraded to version 4.22.

Procedure

    1.   In the OpenShift Web Console, click Storage → External Systems.

    2.   Click Connect external systems.

    3.   Select Red Hat/IBM Ceph Cluster and click Next.

    4.   In the Connect Red Hat Ceph storage page, provide the necessary information:

         a.  Click on the Download Script link to download the python script for extracting Ceph cluster
             details.

         b.  For extracting the Red Hat Ceph Storage (RHCS) cluster details, run the downloaded
             python script on a Red Hat Ceph Storage node with the admin key.

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

             ii.   To retrieve the external cluster details from the RHCS cluster, run the following
                   command:

                     # python3 ceph-external-cluster-details-exporter.py \
                     --rbd-data-pool-name <rbd block pool name> [optional arguments]

                   For example:





8

CHAPTER 4. DEPLOYING OPENSHIFT DATA FOUNDATION EXTERNAL STORAGE CLUSTER


# python3 ceph-external-cluster-details-exporter.py --rbd-data-pool-name ceph-rbd --
monitoring-endpoint xxx.xxx.xxx.xxx --monitoring-endpoint-port xxxx --rgw-endpoint
xxx.xxx.xxx.xxx:xxxx --run-as-user client.ocs

In this example,


rbd-data-   A mandatory parameter that is used for providing block storage in OpenShift
pool-       Data Foundation.
name

rgw-        (Optional) This parameter is required only if the object storage is to be
endpoint    provisioned through Ceph Rados Gateway for OpenShift Data Foundation.
            Provide the endpoint in the following format: <ip_address>:<port>

             NOTE

             A fully-qualified domain name (FQDN) is also supported in
             the format <FQDN>:<PORT>.


monitorin   (Optional) This parameter accepts comma-separated list of IP addresses of
g-          active and standby mgrs reachable from the OpenShift Container Platform
endpoint    cluster. If not provided, the value is automatically populated.

monitorin   (Optional) It is the port associated with the ceph-mgr Prometheus exporter
g-          specified by --monitoring-endpoint. If not provided, the value is
endpoint-   automatically populated.
port

run-as-     (Optional) This parameter is used for providing name for the Ceph user which
user        is created by the script. If this parameter is not specified, a default user name
            client.healthchecker is created. The permissions for the new user is set as:

             caps: [mgr] allow command config

             caps: [mon] allow r, allow command quorum_status, allow command
             version

             caps: [osd] allow rwx pool=RGW_POOL_PREFIX.rgw.meta, allow r
             pool=.rgw.root, allow rw pool=RGW_POOL_PREFIX.rgw.control,
             allow rx pool=RGW_POOL_PREFIX.rgw.log, allow x
             pool=RGW_POOL_PREFIX.rgw.buckets.index



Additional flags:


rgw-pool-   (Optional) The prefix of the RGW pools. If not specified, the default prefix is
prefix      default.

rgw-tls-    (Optional) The file path of the RADOS Gateway endpoint TLS certificate.
cert-path




             9

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters


rgw-skip-     (Optional) This parameter ignores the TLS certification validation when a
tls           self-signed certificate is provided (NOT RECOMMENDED).

ceph-conf     (Optional) The name of the Ceph configuration file.

cluster-      (Optional) The Ceph cluster name.
name

output        (Optional) The file where the output is required to be stored.

cephfs-       (Optional) The name of the CephFS metadata pool.
metadata-
pool-
name

cephfs-       (Optional) The name of the CephFS data pool.
data-
pool-
name

cephfs-       (Optional) The name of the CephFS filesystem.
filesystem
-name

rbd-          (Optional) The name of the erasure coded RBD metadata pool.
metadata-
ec-pool-
name

dry-run       (Optional) This parameter helps to print the executed commands without
              running them.

restricted-   (Optional) This parameter restricts cephCSIKeyrings auth permissions to
auth-         specific pools and clusters. Mandatory flags that need to be set with this are
permission    rbd-data-pool-name and cluster-name. You can also pass thecephfs-
              filesystem-name flag if there is CephFS user restriction so that permission
              is restricted to a particular CephFS filesystem.

                           NOTE

                           This parameter must be applied only for the new
                           deployments. To restrict csi-users per pool and per cluster,
                           you need to create new csi-users and new secrets for those
                           csi-users.



Example with restricted auth permission:

  # python3 /etc/ceph/create-external-cluster-resources.py --cephfs-filesystem-name
  myfs --rbd-data-pool-name replicapool --cluster-name rookStorage --restricted-auth-
  permission true

10

            CHAPTER 4. DEPLOYING OPENSHIFT DATA FOUNDATION EXTERNAL STORAGE CLUSTER

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


              11

    Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters

             Additional advanced options may be available in the UI. Select only those that are required
             for the deployment policy and environment.

         v.  Enable automatic backup
             Allows automatic backup for the Multicloud Object Gateway metadata database.

              A.  Select the backup frequency: Daily, Weekly, or Monthly.

              B.  Specify the number of backups to retain.

                              NOTE

                              This option is disabled if the Use external PostgreSQL option is
                              selected.

     5.  Review the selections and click Connect.

Verification

         Navigate to Storage → Data Foundation → Storage Systemstab and verify that you can view
         all storage clusters.

         Verify that all components for the external OpenShift Data Foundation are successfully
         installed. See Verifying external OpenShift Data Foundation storage deployment for
         instructions.










    12

 CHAPTER 5. VERIFYING EXTERNAL OPENSHIFT DATA FOUNDATION STORAGE CLUSTER DEPLOYMENT

                             CHAPTER 5. VERIFYING EXTERNAL OPENSHIFT DATA
                                FOUNDATION STORAGE CLUSTER DEPLOYMENT

Verify the external OpenShift Data Foundation cluster deployment is healthy in a multiple-cluster
configuration.

5.1. VERIFYING THE STATE OF THE PODS

 1.   Click Workloads → Podsfrom the left pane of the OpenShift Web Console.

 2.   Select openshift-storage from the Project drop-down list.

              NOTE

              If the Show default projects option is disabled, use the toggle button to list all
              the default projects.

 3.   Verify that the following pods are in running state:
      Pods corresponding to OpenShift Data Foundation components

 Component        Corresponding pods

 OpenShift Data Foundation   ocs-operator-* (1 pod on any worker node)
 Operator
                             ocs-metrics-exporter-* (1 pod on any worker node)

                                    NOTE

                                    This pod must be present in openshift-storage-
                                    extended namespace as well such that there is 1 pod
                                    in each openshift-storage and openshift-
                                    storage extended namespace.

                             odf-operator-controller-manager-* (1 pod on any worker node)

                             odf-console-* (1 pod on any worker node)

                             csi-addons-controller-manager-* (1 pod on any worker node)


Rook-ceph Operator    rook-ceph-operator-* (1 pod on any worker node)










                                    13

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters


Component    Corresponding pods

CSI          cephfs
                openshift-storage.cephfs.csi.ceph.com-ctrlplugin-* (2
                pods distributed across storage nodes)

                openshift-storage.cephfs.csi.ceph.com-nodeplugin-* (1
                pod on each storage node)

             nfs
                openshift-storage.nfs.csi.ceph.com-ctrlplugin-* (2 pods
                distributed across storage nodes)

                openshift-storage.nfs.csi.ceph.com-nodeplugin-* (1 pod
                on each storage node)

             rbd
                openshift-storage.rbd.csi.ceph.com-ctrlplugin-* (2 pods
                distributed across storage nodes)

                openshift-storage.rbd.csi.ceph.com-nodeplugin-* (1 pod
                on each storage node)



5.2. VERIFYING THAT THE OPENSHIFT DATA FOUNDATION CLUSTER
IS HEALTHY

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and verify that you can now see
     two storage system links in the pop up that appears.

3.   Click each of the storage system links and verify the following:

      a.  In the Status card of the Block and File tab, verify that the Storage Cluster has a green tick.

      b.  In the Details card, verify that the cluster information is displayed.

5.3. VERIFYING THAT THE MULTICLOUD OBJECT GATEWAY IS
HEALTHY

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and then click the storage
     system link from the pop up that appears.

      a.  In the Status card of the Object tab, verify that both Object Service and Data Resiliency
          have a green tick.

      b.  In the Details card, verify that the Multicloud Object Gateway (MCG) information is
          displayed.


          NOTE
14

  CHAPTER 5. VERIFYING EXTERNAL OPENSHIFT DATA FOUNDATION STORAGE CLUSTER DEPLOYMENT

           NOTE

           The RADOS Object Gateway is only listed in case RADOS Object Gateway endpoint
           details are included while deploying OpenShift Data Foundation in external mode.

 5.4. VERIFYING THAT THE STORAGE CLASSES ARE CREATED AND
 LISTED

    1.   Click Storage → Storage Classesfrom the left pane of the OpenShift Web Console.

    2.   Verify that the following storage classes are created with the OpenShift Data Foundation
         cluster creation:

          ocs-external-storagecluster-ceph-rbd

          ocs-external-storagecluster-ceph-rgw

          ocs-external-storagecluster-cephfs

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

 5.5. VERIFYING THAT CEPH CLUSTER IS CONNECTED

 Run the following command to verify if the OpenShift Data Foundation cluster is connected to the
 external Red Hat Ceph Storage cluster.

  $ oc get cephcluster -n openshift-storage-extended
  NAME                   DATADIRHOSTPATH     MONCOUNT     AGE   PHASE MESSAGE
  HEALTH   EXTERNAL      FSID
  ocs-external-storagecluster-cephcluster        51m     Connected Cluster connected
successfully HEALTH_OK true 8f01d842-d4b2-11ee-b43c-0050568fb522

 5.6. VERIFYING THAT STORAGE CLUSTER IS READY

 Run the following command to verify if the storage cluster is ready and the External option is set to
 true.




                     15

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters

 $ oc get storagecluster -n openshift-storage-extended
 NAME                     AGE  PHASE EXTERNAL       CREATED AT          VERSION
 ocs-external-storagecluster   51m Ready   true     2024-02-28T10:05:54Z 4.22.0

5.7. VERIFYING THE STATE OF THE PODS

Verify that all expected OpenShift Data Foundation pods are running in the openshift-storage
namespace after deploying multiple storage clusters.

Procedure

 1.   Click Workloads → Podsfrom the left pane of the OpenShift Web Console.

 2.   Select openshift-storage from the Project drop-down list.

                      NOTE

                      If the Show default projects option is disabled, use the toggle button to list all
                      the default projects.

 3.   Verify that the following pods are in running state:
      Pods corresponding to OpenShift Data Foundation components:

         Component                 Corresponding pods

         OpenShift Data              ocs-operator-* (1 pod on any worker node)
         Foundation Operator
                                     ocs-metrics-exporter-* (1 pod on any worker node)

                                                    NOTE

                                                    This pod must be present in openshift-storage-
                                                    extended namespace as well such that there is 1
                                                    pod in each openshift-storage and openshift-
                                                    storage extended namespace.

                                     odf-operator-controller-manager-* (1 pod on any worker
                                     node)

                                     odf-console-* (1 pod on any worker node)

                                     csi-addons-controller-manager-* (1 pod on any worker node)


      Rook-ceph Operator             rook-ceph-operator-* (1 pod on any worker node)








16

CHAPTER 5. VERIFYING EXTERNAL OPENSHIFT DATA FOUNDATION STORAGE CLUSTER DEPLOYMENT


Component    Corresponding pods

CSI          cephfs
                openshift-storage.cephfs.csi.ceph.com-ctrlplugin-* (2
                pods distributed across storage nodes)

                openshift-storage.cephfs.csi.ceph.com-nodeplugin-*
                (1 pod on each storage node)

             nfs
                openshift-storage.nfs.csi.ceph.com-ctrlplugin-* (2
                pods distributed across storage nodes)

                openshift-storage.nfs.csi.ceph.com-nodeplugin-* (1
                pod on each storage node)

             rbd
                openshift-storage.rbd.csi.ceph.com-ctrlplugin-* (2
                pods distributed across storage nodes)

                openshift-storage.rbd.csi.ceph.com-nodeplugin-* (1
                pod on each storage node)




5.8. VERIFYING THAT THE OPENSHIFT DATA FOUNDATION CLUSTER
IS HEALTHY

Verify the health status of all storage systems after deploying multiple OpenShift Data Foundation
storage clusters in the OpenShift web console.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation.

2.   In the Status card of the Overview tab, click Storage System and verify that you can now see
     two storage system links in the pop up that appears.

3.   Click each of the storage system links and verify the following:

      a.  In the Status card of the Block and File tab, verify that the Storage Cluster has a green tick.

      b.  In the Details card, verify that the cluster information is displayed.
          For more information on the health of the OpenShift Data Foundation cluster using the
          Block and File dashboard, see the Monitoring OpenShift Data Foundation documentation.

Additional resources

     Monitoring OpenShift Data Foundation

5.9. VERIFYING THAT THE MULTICLOUD OBJECT GATEWAY IS
HEALTHY

                    17

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters

Verify that the Multicloud Object Gateway is healthy and both Object Service and Data Resiliency are
reporting green status in a multiple storage cluster deployment.

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

5.10. VERIFYING THAT THE STORAGE CLASSES ARE CREATED AND
LISTED

Verify that the expected OpenShift Data Foundation storage classes are created and listed after
deploying multiple storage clusters.

Procedure

1.   Click Storage → Storage Classesfrom the left pane of the OpenShift Web Console.

2.   Verify that the following storage classes are created with the OpenShift Data Foundation
     cluster creation:

          ocs-external-storagecluster-ceph-rbd

          ocs-external-storagecluster-ceph-rgw

          ocs-external-storagecluster-cephfs









                      NOTE
18

CHAPTER 5. VERIFYING EXTERNAL OPENSHIFT DATA FOUNDATION STORAGE CLUSTER DEPLOYMENT

                    NOTE

                             If an MDS is not deployed in the external cluster, ocs-external-
                             storagecluster-cephfs storage class will not be created.

                             If RGW is not deployed in the external cluster, the ocs-external-
                             storagecluster-ceph-rgw storage class will not be created.

                    For more information regarding MDS and RGW, see the Red Hat Ceph
                    Storage documentation.

Additional resources

     Red Hat Ceph Storage documentation

5.11. VERIFYING THAT CEPH CLUSTER IS CONNECTED

Verify that the OpenShift Data Foundation cluster is connected to the external Red Hat Ceph Storage
cluster in a multiple storage cluster deployment.

Procedure

1.   Run the following command to verify if the OpenShift Data Foundation cluster is connected to
     the external Red Hat Ceph Storage cluster:

         $ oc get cephcluster -n openshift-storage-extended
         NAME                     DATADIRHOSTPATH       MONCOUNT AGE PHASE
         MESSAGE             HEALTH                 EXTERNAL  FSID
         ocs-external-storagecluster-cephcluster        51m     Connected              Cluster
    connected successfully HEALTH_OK true 8f01d842-d4b2-11ee-b43c-0050568fb522

5.12. VERIFYING THAT STORAGE CLUSTER IS READY

Verify that the external OpenShift Data Foundation storage cluster is ready and has the External flag
set to true after deployment.

Procedure

1.   Run the following command to verify if the storage cluster is ready and the External option is
     set to true:

         $ oc get storagecluster -n openshift-storage-extended
         NAME                AGE  PHASE     EXTERNAL     CREATED AT                    VERSION
    ocs-external-storagecluster 51m Ready true 2024-02-28T10:05:54Z 4.22.0










                                                    19

Red Hat OpenShift Data Foundation 4.22 Deploying multiple OpenShift Data Foundation storage clusters

        CHAPTER 6. MIGRATING APPLICATION WORKLOADS

  You can migrate application workloads from the internal mode storage classes to the external mode
  storage classes using Migration Toolkit for Containers using the same cluster as source and target.










    20
