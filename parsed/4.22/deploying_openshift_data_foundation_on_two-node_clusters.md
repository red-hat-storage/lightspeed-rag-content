---
title: "deploying openshift data foundation on two-node clusters"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/deploying_openshift_data_foundation_on_two-node_clusters"
source_file: "deploying_openshift_data_foundation_on_two-node_clusters.pdf"
version: "4.22"
chapter: "deploying_openshift_data_foundation_on_two-node_clusters"
---

                      Red Hat OpenShift Data Foundation
                                     4.22

                 Deploying OpenShift Data Foundation on two-
                                node clusters

Instructions on deploying OpenShift Data Foundation on two-node clusters with
                                     DRBD










    Last Updated: 2026-10-22



Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data
Foundation on two-node clusters

Instructions on deploying OpenShift Data Foundation on two-node clusters with DRBD

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

Read this document for instructions about how to install Red Hat OpenShift Data Foundation on
two-node OpenShift clusters using Distributed Replicated Block Device (DRBD) for high availability.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. PREPARING TO DEPLOY TWO-NODE CLUSTER 6

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION ON TWO-NODE CLUSTER 8
  2.1. INSTALLING RED HAT OPENSHIFT DATA FOUNDATION OPERATOR 8
  2.2. CONFIGURING DRBD FOR TWO-NODE CLUSTER 9
  2.3. CREATING OPENSHIFT DATA FOUNDATION CLUSTER ON TWO-NODE CLUSTER 12
  2.4. TROUBLESHOOTING TWO-NODE DEPLOYMENTS 13
  2.5. FEATURES NOT SUPPORTED IN TWO-NODE DEPLOYMENTS 13










    1

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on two-node clusters










2

PREFACE

PREFACE

This document describes how to deploy Red Hat OpenShift Data Foundation on two-node clusters.










3

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on two-node clusters

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

Red Hat OpenShift Data Foundation supports deployment on two-node Red Hat OpenShift Container
Platform (RHOCP) bare metal clusters with fencing and Distributed Replicated Block Device (DRBD)
enabled. This architecture is suitable for distributed and edge deployments where the goal is to
minimize hardware footprint while maintaining High Availability (HA).

  NOTE

  Two-node cluster deployments with DRBD are supported only on bare metal
  infrastructure. See the Planning your deployment guide and the Preparing to deploy two-
  node cluster chapter for more information about deployment requirements.

To deploy OpenShift Data Foundation on a two-node cluster, follow the deployment process in the
Deploy using local storage devices with DRBD section.

Additional resources

 Planning your deployment

 Preparing to deploy two-node cluster










  5

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on two-node clusters

 CHAPTER 1. PREPARING TO DEPLOY TWO-NODE CLUSTER

 Before installing OpenShift Data Foundation on a two-node cluster, you must create persistent volumes
 manually for the OSDs on both nodes.

 Prerequisites

       You have a two-node OpenShift cluster deployed on bare metal infrastructure.
       Verify the cluster topology:

          oc get infrastructure cluster -o jsonpath='{.status.controlPlaneTopology}'

       The output should be DualReplica.

       You have identified the storage devices on both nodes that will be used for OSDs.

       Storage devices must be SSD type with minimum size of 500 GB.

       Storage device sizes must match across both nodes.

       For the list of packages to be mirrored for installing OpenShift Data Foundation in
       disconnected environments, see the Disconnected environment section in the Planning your
       deployment guide.

 Procedure

  1.   Create the storage class:

          oc create -f lso-storageclass.yml

       Example lso-storageclass.yml

  apiVersion: storage.k8s.io/v1
  kind: StorageClass
  metadata:
  name: localblock
  provisioner: kubernetes.io/no-provisioner
  volumeBindingMode: WaitForFirstConsumer

  1.   Create the persistent volumes:
       Edit the pv.yaml file to match your environment:

          Device path by disk-by-id (for example, /dev/disk/by-id/wwn-0x58ce38ee2311e971)

          Storage size (for example, storage: 100Gi)

          Node affinity (for example, values: - master-1)

              oc create -f pv.yml

          Example pv.yml

  apiVersion: v1
  kind: PersistentVolume

6

                CHAPTER 1. PREPARING TO DEPLOY TWO-NODE CLUSTER

 metadata:
  name: local-pv-1
 spec:
  capacity:
  storage: 500Gi
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
  storageClassName: localblock
  local:
  path: /dev/disk/by-id/wwn-0x58ce38ee2311e971
  nodeAffinity:
  required:
   nodeSelectorTerms:
   - matchExpressions:
    - key: kubernetes.io/hostname
      operator: In
      values:
      - master-0

+ Repeat this step to create PVs for all OSD disks on both nodes.

Verification

    Verify the creation of PVs:

        oc get pv

            NOTE

            If you are running OpenShift Container Platform with RHEL KVM, you must assign a serial
            number to your VM disk. Otherwise, the VM disk cannot be identified after reboot. You
            can use the virsh edit <VM> command to add the     <serial>mydisk</serial> definition.










            7

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on two-node clusters

             CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION ON
                            TWO-NODE CLUSTER

You can deploy OpenShift Data Foundation on a two-node bare metal cluster where OpenShift
Container Platform is already installed. When you deploy OpenShift Data Foundation on a two-node
OpenShift Container Platform cluster using local storage devices and Distributed Replicated Block
Device (DRBD), you create internal cluster resources with high availability. This approach uses DRBD for
the floating monitor to ensure data consistency across both nodes.

Perform the following steps to deploy OpenShift Data Foundation:

1.   Install the Red Hat OpenShift Data Foundation Operator .

2.   Configure DRBD for two-node cluster.

3.   Create an OpenShift Data Foundation cluster on two-node cluster .

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


8

                   CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION ON TWO-NODE CLUSTER

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

2.2. CONFIGURING DRBD FOR TWO-NODE CLUSTER

Configure DRBD (Distributed Replicated Block Device) for the floating monitor after installing the
OpenShift Data Foundation operator.

Prerequisites

         You have installed the OpenShift Data Foundation operator.

         You have a dedicated disk or partition on both nodes for the floating monitor with minimum size
         of 10 GB (recommended less than 50 GB).

         The disk must be SSD type and the same size on both nodes.

         You have oc CLI access to the cluster.

Procedure

1.       Fetch the DRBD setup script from the configmap:



               9

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on two-node clusters

         oc get configmap rook-ceph-drbd-setup-script -n openshift-storage -o jsonpath='{.data.script}'
         | base64 -d > drbd-setup && chmod +x drbd-setup

                NOTE

                The script must be run from a shell that has oc access to the cluster.

 2.     View the help text to understand available options:

         ./drbd-setup -h

        Example output

Usage:
 ./drbd-setup -d <path>
 ./drbd-setup -d0 <path0> -d1 <path1>
 ./drbd-setup -l

Backing paths are raw block device paths (e.g. /dev/sdb). Use the PATH column from -l.
Disks must be SSD-class (ROTA 0) and same size on both nodes.

 -d PATH
       One path used on both nodes. Choose this when each machine has the replica disk at the
       same device name (both nodes use e.g. /dev/sdb for the DRBD lower layer).

Use -d0/-d1 when the two nodes use different paths; do not combine -d with -d0/-d1.

 -d0 PATH  Path on node 0 only (first node name after sorting all cluster nodes).
 -d1 PATH  Path on node 1 only (second node).

Discovery:
 -l     List block devices on each node (NAME, PATH, SIZE, ROTA, TYPE, FSTYPE).

DRBD options:
 --drbd-conf-path PATH Host path to drbd.conf (default /etc/drbd.conf)
 --drbd-dir-path PATH   Host dir for resource snippets (default /etc/drbd.d)
 --drbd-resource NAME   Logical DRBD resource name in config (default r0)
 --drbd-device PATH     DRBD upper device node path, same on both nodes (default /dev/drbd0)
 --drbd-port N  TCP port for DRBD replication (default 7794)

General:
 -h     Show this help and exit

Environment:
 Defaults are documented on each assignment near the top of this script.
 OUTPUT_CM_NS / OUTPUT_CM_NAME — namespace and name of the summary ConfigMap
(floating mon).

   1.   List available disks on both nodes:

         ./drbd-setup -l

   2.   Run the setup script specifying the disk path for the floating monitor:


10

                     CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION ON TWO-NODE CLUSTER

      ./drbd-setup -d /dev/sdb

               NOTE

                     If the script fails intermittently at any step, re-run the script.

                     Use -d when both nodes use the same device path.

                     Use -d0 and -d1 when nodes use different device paths.

Verification

     Verify the DRBD configuration ConfigMap:
     A ConfigMap named drbd-configure is created in the openshift-storage namespace:

      oc get cm drbd-configure -n openshift-storage -o yaml

     Example output

apiVersion: v1
data:
 BLOCK_DEVICE_PATH_NODE_0: /dev/sdb
 BLOCK_DEVICE_PATH_NODE_1: /dev/sdb
 DISK_BY_ID_NODE_0: /dev/disk/by-id/scsi-0QEMU_QEMU_HARDDISK_drive-scsi1-0-0-0
 DISK_BY_ID_NODE_1: /dev/disk/by-id/scsi-0QEMU_QEMU_HARDDISK_drive-scsi1-0-0-0
 DRBD_CONF_PATH: /etc/drbd.conf
 DRBD_DEVICE_NAME: /dev/drbd0
 DRBD_DIR_PATH: /etc/drbd.d
 DRBD_PORT: "7794"
 DRBD_RESOURCE_NAME: r0
 DRBD_UTILS_IMAGE: registry.redhat.io/odf4/odf-drbd-
rhel9@sha256:af2fc67f7690944008fe56bea65adf2bfa18d7fc437e2b2c21474de3520e3e90
 DRBD_VERSION: 9.2.18
 NODE_0_IP: 192.168.111.20
 NODE_0_NAME: master-0
 NODE_1_IP: 192.168.111.21
 NODE_1_NAME: master-1
 SETUP_TIMESTAMP: "2026-05-18T13:01:44Z"
kind: ConfigMap
metadata:
 name: drbd-configure
 namespace: openshift-storage

     Verify the drbd-autostart DaemonSet is present in the openshift-kmm namespace:

      oc get daemonset drbd-autostart -n openshift-kmm

     Check DRBD status on both nodes:

      oc debug -q node/master-0 -- chroot /host podman run --rm --privileged \
            -v /dev:/dev -v /etc/drbd.conf:/etc/drbd.conf -v /etc/drbd.d:/etc/drbd.d \
            --hostname master-0 --net host registry.redhat.io/odf4/odf-drbd-
      rhel9@sha256:af2fc67f7690944008fe56bea65adf2bfa18d7fc437e2b2c21474de3520e3e90 \
            drbdadm -c /etc/drbd.conf status r0

                         11

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on two-node clusters



     Repeat the command for master-1.

2.3. CREATING OPENSHIFT DATA FOUNDATION CLUSTER ON TWO-
NODE CLUSTER

Create a OpenShift Data Foundation storage cluster on a two-node deployment after completing
DRBD configuration, using the OpenShift Web Console to provision the storage system within the
constraints of the two-node architecture.

Prerequisites

     You have completed the steps in Preparing to deploy two-node cluster .

     You have installed the OpenShift Data Foundation operator.

     You have completed the steps in Configuring DRBD for two-node cluster.

     The drbd-configure ConfigMap exists in the openshift-storage namespace.

Procedure

1.   In the OpenShift Web Console, click Storage → Data Foundation → Storage Systems →
     Create StorageSystem.

 2. In the Backing storage page, select the storage class created earlier and click Next.

3.   Follow the wizard to configure the storage cluster according to your requirements.

4.   Wait for the StorageCluster status to become Ready.

Verification

     Verify the StorageCluster status:

         oc get storagecluster -n openshift-storage

     The status should be Ready.

     Verify all pods are running:

         oc get pods -n openshift-storage

     Verify the floating monitor is deployed correctly:

         oc get pods -n openshift-storage | grep rook-ceph-mon

     Check Ceph cluster health:

         oc rsh -n openshift-storage $(oc get pods -n openshift-storage -l app=rook-ceph-tools -o
         name) ceph status



         NOTE
12

                                       CHAPTER 2. DEPLOY OPENSHIFT DATA FOUNDATION ON TWO-NODE CLUSTER

          NOTE

              After node recovery, if Ceph status shows OSD pod down, try to restart the OSD
              pods to get them back.

              Two-node clusters support applications using RBD (RADOS Block Device) and
              CephFS storage with both ReadWriteOnce (RWO) and ReadWriteMany (RWX)
              access modes.

2.4. TROUBLESHOOTING TWO-NODE DEPLOYMENTS

During node failover events, the floating monitor might enter an unexpected drbd-initfailover state if
the incorrect node is marked as DRBD secondary.

Workaround

You can manually set both nodes to DRBD Secondary to reset the DRBD role state:

 oc debug node/"$Node" -- chroot /host sudo podman run --rm --privileged \
  -v /dev:/dev -v /etc/drbd.conf:/etc/drbd.conf -v /etc/drbd.d:/etc/drbd.d \
  --hostname "$Node" registry.redhat.io/odf4/odf-drbd-
 rhel9@sha256:af2fc67f7690944008fe56bea65adf2bfa18d7fc437e2b2c21474de3520e3e90 \
  drbdadm secondary r0

Replace $Node with the node name (for example, master-0 or master-1). Run this command for both
nodes.

2.5. FEATURES NOT SUPPORTED IN TWO-NODE DEPLOYMENTS

Several OpenShift Data Foundation features are not supported in the two-node deployment.

      Multicloud Object Gateway (NooBaa)

      Network File System (NFS)

      RADOS Object Gateway (RGW)

      Disaster Recovery (Regional DR and Metro-DR)

      PodDisruptionBudgets (PDBs)

      Monitor failover

      Multus networking

      External PostgreSQL

      Node selection

      Performance/resource profiles

      Automatic capacity scaling

      Host networking

      Erasure Coded (EC) pools

                                                                                                    13

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation on two-node clusters

External clients (HCI/Provider mode)










14
