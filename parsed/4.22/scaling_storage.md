---
title: "scaling storage"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/scaling_storage"
source_file: "scaling_storage.pdf"
version: "4.22"
chapter: "scaling_storage"
---

               Red Hat OpenShift Data Foundation
                              4.22

    Scaling storage

Instructions for scaling operations in OpenShift Data Foundation










    Last Updated: 2026-10-27



Red Hat OpenShift Data Foundation 4.22 Scaling storage

Instructions for scaling operations in OpenShift Data Foundation

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

    This document explains scaling options for Red Hat OpenShift Data Foundation.

                                                                         Table of Contents

       Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE        4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION                                              5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. INTRODUCTION TO SCALING STORAGE                                               6
 1.1. SUPPORTED DEPLOYMENTS FOR RED HAT OPENSHIFT DATA FOUNDATION                        6

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. REQUIREMENTS FOR SCALING STORAGE                                              8

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. SCALING STORAGE CAPACITY OF AWS OPENSHIFT DATA FOUNDATION CLUSTER             9
 3.1. SCALING UP STORAGE CAPACITY ON A CLUSTER                                           9
  3.1.1. Scaling up storage capacity on an AWS cluster by resizing existing OSDs        11
 3.2. SCALING OUT STORAGE CAPACITY ON A AWS CLUSTER                                     12
  3.2.1. Adding a node                                                                  13
   3.2.1.1. Adding a node to an installer-provisioned infrastructure                    13
   3.2.1.2. Adding a node to an user-provisioned infrastructure                         14
  3.2.2. Scaling up storage capacity                                                    15
 3.3. ENABLING AUTOMATIC CAPACITY SCALING                                               15
 3.4. SCALING NON-RESILIENT POOLS                                                       16

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. SCALING STORAGE OF BARE METAL OPENSHIFT DATA FOUNDATION CLUSTER              18
 4.1. SCALING UP A CLUSTER CREATED USING LOCAL STORAGE DEVICES                          18
  4.1.1. Scaling up storage capacity on a cluster by resizing existing OSDs             20
 4.2. SCALING OUT STORAGE CAPACITY ON A BARE METAL CLUSTER                              21
  4.2.1. Adding a node                                                                  22
   4.2.1.1. Adding a node to an installer-provisioned infrastructure                    22
   4.2.1.2. Adding a node using a local storage device                                  23
  4.2.2. Scaling up storage capacity                                                    25

CHAPTER 5. SCALING STORAGE USING MULTIPLE DEVICE CLASS IN THE SAME CLUSTER FOR LOCAL
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
STORAGE DEPLOYMENTS                                                                     26
 5.1. CREATING A NEW LOCALVOLUMESET                                                     26
 5.2. ATTACHING STORAGE FOR A NEW DEVICE SET                                            28
 5.3. ATTACHING STORAGE FOR A NEW DEVICE SET BY USING DISKS OF SAME TYPE ON THE SAME NODES
                                                                                        31

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER                  35
 6.1. SCALING UP STORAGE ON A DYNAMICALLY PROVISIONED VMWARE CLUSTER                    35
 6.2. SCALING UP A CLUSTER CREATED USING LOCAL STORAGE DEVICES                          37
  6.2.1. Scaling up storage capacity on a cluster by resizing existing OSDs             39
 6.3. ENABLING AUTOMATIC CAPACITY SCALING                                               40
 6.4. SCALING OUT STORAGE CAPACITY ON A VMWARE CLUSTER                                  41
  6.4.1. Adding a node to an installer-provisioned infrastructure                       41
  6.4.2. Adding a node to an user-provisioned infrastructure                            42
  6.4.3. Adding a node using a local storage device                                     43
  6.4.4. Scaling up storage capacity                                                    45
 6.5. SCALING NON-RESILIENT POOLS                                                       46

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 7. SCALING STORAGE OF MICROSOFT AZURE OPENSHIFT DATA FOUNDATION CLUSTER         48
 7.1. SCALING UP STORAGE CAPACITY ON A CLUSTER                                          48
  7.1.1. Scaling up storage capacity on a cluster by resizing existing OSDs             50
 7.2. ENABLING AUTOMATIC CAPACITY SCALING                                               51
 7.3. SCALING OUT STORAGE CAPACITY ON A MICROSOFT AZURE CLUSTER                         52
  7.3.1. Adding a node to an installer-provisioned infrastructure                       52

                                                                                         1

Red Hat OpenShift Data Foundation 4.22 Scaling storage

  7.3.2. Scaling up storage capacity                                            53
 7.4. SCALING NON-RESILIENT POOLS                                               53

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 8. SCALING STORAGE CAPACITY OF GCP OPENSHIFT DATA FOUNDATION CLUSTER    55
 8.1. SCALING UP STORAGE CAPACITY ON A CLUSTER                                  55
  8.1.1. Scaling up storage capacity on a cluster by resizing existing OSDs     57
 8.2. ENABLING AUTOMATIC CAPACITY SCALING                                       58
 8.3. SCALING OUT STORAGE CAPACITY ON A GCP CLUSTER                             59
  8.3.1. Adding a node                                                          59
   8.3.1.1. Adding a node to an installer-provisioned infrastructure            59
  8.3.2. Scaling up storage capacity                                            60
 8.4. SCALING NON-RESILIENT POOLS                                               60

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 9. SCALING STORAGE OF IBM Z OR IBM LINUXONE OPENSHIFT DATA FOUNDATION CLUSTER
                                                                                63
 9.1. SCALING UP STORAGE BY ADDING CAPACITY TO YOUR OPENSHIFT DATA FOUNDATION NODES ON
 IBM Z OR IBM LINUXONE INFRASTRUCTURE                                           63
 9.2. SCALING OUT STORAGE CAPACITY ON A IBM Z OR IBM LINUXONE CLUSTER           66
  9.2.1. Adding a node using a local storage device                             66
  9.2.2. Scaling up storage capacity                                            68

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 10. SCALING STORAGE OF IBM POWER OPENSHIFT DATA FOUNDATION CLUSTER      69
 10.1. SCALING UP STORAGE BY ADDING CAPACITY TO YOUR OPENSHIFT DATA FOUNDATION NODES ON
 IBM POWER INFRASTRUCTURE USING LOCAL STORAGE DEVICES                           69
 10.2. SCALING OUT STORAGE CAPACITY ON A IBM POWER CLUSTER                      72
  10.2.1. Adding a node using a local storage device on IBM Power               73
  10.2.2. Scaling up storage capacity                                           74

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 11. SCALING STORAGE CAPACITY OF IBM FLASHSYSTEM CLUSTER                 75
 11.1. SCALING UP STORAGE CAPACITY ON A CLUSTER                                 75
 11.2. SCALING OUT STORAGE CAPACITY ON AN IBM FLASHSYSTEM CLUSTER               77
  11.2.1. Adding a node                                                         77
   11.2.1.1. Adding a node to an installer-provisioned infrastructure           77
   11.2.1.2. Adding a node to an user-provisioned infrastructure                78
  11.2.2. Scaling up storage capacity                                           79










2

Table of Contents










3

Red Hat OpenShift Data Foundation 4.22 Scaling storage

    PREFACE

This document describes how to scale the storage capacity of your OpenShift Data Foundation cluster
on various platforms.










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

Red Hat OpenShift Data Foundation 4.22 Scaling storage

 CHAPTER 1. INTRODUCTION TO SCALING STORAGE

Red Hat OpenShift Data Foundation is a highly scalable storage system. OpenShift Data Foundation
allows you to scale by adding the disks in the multiple of three, or three or any number depending upon
the deployment type.

 For internal (dynamic provisioning) deployment mode, you can increase the capacity by adding
 3 disks at a time.

 For internal-attached (Local Storage Operator based) mode, you can deploy with less than 3
 failure domains.

With flexible scale deployment enabled, you can scale up by adding any number of disks. For deployment
with 3 failure domains, you will be able to scale up by adding disks in the multiple of 3.

For scaling your storage in external mode, see the Red Hat Ceph Storage documentation.

            NOTE

            You can use a maximum of twelve storage devices per node. The high number of storage
            devices will lead to a higher recovery time during the loss of a node. This
            recommendation ensures that nodes stay below the cloud provider dynamic storage
            device attachment limits, and limits the recovery time after node failure with local storage
            devices.

While scaling, you must ensure that there are enough CPU and Memory resources as per scaling
requirement.

Supported storage classes by default

 gp2-csi on AWS

 thin on VMware

 managed-csi on Microsoft Azure, including enabled performance plus

1.1. SUPPORTED DEPLOYMENTS FOR RED HAT OPENSHIFT DATA
FOUNDATION

 User-provisioned infrastructure:

  Amazon Web Services (AWS)

  VMware

  Bare metal

  IBM Power

  IBM Z or IBM® LinuxONE

 Installer-provisioned infrastructure:

  Amazon Web Services (AWS)



6

    CHAPTER 1. INTRODUCTION TO SCALING STORAGE

Microsoft Azure

VMware

Bare metal

Additional resources

Red Hat Ceph Storage documentation










7

Red Hat OpenShift Data Foundation 4.22 Scaling storage

CHAPTER 2. REQUIREMENTS FOR SCALING STORAGE

Before you proceed to scale the storage nodes, refer to the following sections to understand the node
requirements for your specific Red Hat OpenShift Data Foundation instance:

 Platform requirements

 Resource requirements

 Storage device requirements

  Dynamic storage devices

  Local storage devices

  Capacity planning

   IMPORTANT

   Always ensure that you have plenty of storage capacity.

   If storage ever fills completely, it is not possible to add capacity or delete or migrate
   content away from the storage to free up space completely. Full storage is very difficult
   to recover.

   Capacity alerts are issued when cluster storage capacity reaches 75% (near-full) and 85%
   (full) of total capacity. Always address capacity warnings promptly, and review your
   storage regularly to ensure that you do not run out of storage space.

   If storage capacity reaches 85% full state, Ceph may report HEALTH_ERR and prevent
   IO operations. In this case, you can increase the full ratio temporarily so that cluster
   rebalance can take place. For steps to increase the full ratio, see Setting Ceph OSD full
   thresholds using the ODF CLI tool.

   If you do run out of storage space completely, contact Red Hat Customer Support .










8

          CHAPTER 3. SCALING STORAGE CAPACITY OF AWS OPENSHIFT DATA FOUNDATION CLUSTER

             CHAPTER 3. SCALING STORAGE CAPACITY OF AWS
                 OPENSHIFT DATA FOUNDATION CLUSTER

To scale the storage capacity of your configured Red Hat OpenShift Data Foundation worker nodes on
AWS cluster, you can increase the capacity by adding three disks at a time. Three disks are needed since
OpenShift Data Foundation uses a replica count of 3 to maintain the high availability. So the amount of
storage consumed is three times the usable space.

             NOTE

             Usable space might vary when encryption is enabled or replica 2 pools are being used.

3.1. SCALING UP STORAGE CAPACITY ON A CLUSTER

To increase the storage capacity in a dynamically created storage cluster on an user-provisioned
infrastructure, you can add storage capacity and performance to your configured Red Hat OpenShift
Data Foundation worker nodes.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

     The disk should be of the same size and type as used during initial deployment.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Storage → Storage cluster.

      a.  Click Actions at the top right of the page to extend the options menu.

      b.  Select Add Capacity from the options menu.

      c.  Select the Storage Class. Choose the storage class which you wish to use to provision new
          storage devices.

      d.  Click Add.

3.   To check the status, navigate to Storage → Storage cluster and verify that the Storage
     Cluster in the Status card has a green tick.

Verification

     Verify the Raw Capacity card.

      a.  In the OpenShift Web Console, click Storage → Storage cluster.

      b.  In the Block and File tab, check the Raw Capacity card.
          Note that the capacity increases based on your selections.



             NOTE
                 9

    Red Hat OpenShift Data Foundation 4.22 Scaling storage

                  NOTE

                  The raw capacity does not take replication into account and shows the full
                  capacity.

    Verify that the new object storage devices (OSDs) and their corresponding new Persistent
    Volume Claims (PVCs) are created.

        To view the state of the newly created OSDs:

         a.  Click Workloads → Pods from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

        To view the state of the PVCs:

         a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.

    a.  Identify the nodes where the new OSD pods are running.

            $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
            pod-name>

        <OSD-pod-name>
            Is the name of the OSD pod.
            For example:

              $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
              ceph-osd-0-544db49d7f-qrgqm

            Example output:

              NODE
              compute-1


        b. For each of the nodes identified in the previous step, do the following:

        i. Create a debug pod and open a chroot environment for the selected hosts.

10

CHAPTER 3. SCALING STORAGE CAPACITY OF AWS OPENSHIFT DATA FOUNDATION CLUSTER

$ oc debug node/<node-name>

<node-name>
Is the name of the node.

$ chroot /host


          ii. Check for the crypt keyword beside the ocs-deviceset names.

                $ lsblk

             IMPORTANT

             Cluster reduction is supported only with the Red Hat Support Team’s assistance.

3.1.1. Scaling up storage capacity on an AWS cluster by resizing existing OSDs

To increase the storage capacity on a cluster, you can add storage capacity by resizing existing OSDs.

             IMPORTANT

             Resizing existing OSDs can only be done once every 6 hours due to AWS limitations. If
             you exceed this amount within the 6 hour time frame, you will receive a warning that
             includes the note You’ve reached the maximum modification rate per volume limit.
             Wait at least 6 hours between modifications per EBS volume.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Update the dataPVCTemplate size for the storageDeviceSets in the StorageCluster CR with
     the new desired size. For example:

         apiVersion: ocs.openshift.io/v1
         kind: StorageCluster
         metadata:
          namespace: openshift-storage
          name: example-storagecluster
         spec:
          manageNodes: false
          monPVCTemplate:
          spec:
              storageClassName:
              accessModes:
              - ReadWriteOnce
              resources:
              requests:
              storage: 10Gi

                  11

Red Hat OpenShift Data Foundation 4.22 Scaling storage

       storageDeviceSets:
       - name: example-deviceset
        count: 3
        resources: {}
        placement: {}
        dataPVCTemplate:
         spec:
          storageClassName:
          accessModes:
          - ReadWriteOnce
          volumeMode: Block
          resources:
          requests:
           storage: 1Ti
        portable: true

     In this example YAML, the storage parameter under storageDeviceSets reflects the size
     change of 1Ti.

 2.  Wait for the OSDs to restart.

 3.  Confirm that the resize took effect:

      $ oc get pvc -l ceph.rook.io/DeviceSet

     Example output:

      % oc get pvc -l ceph.rook.io/DeviceSet
      NAME     STATUS     VOLUME        CAPACITY     ACCESS MODES
      STORAGECLASS        AGE
    rook-ceph-osd-a4681 Bound pvc-c3868d38-6bd9-4c8d-91e1-a4681e4cb8c1 1Ti
      RWO     gp2-csi     176m

     Verify the resize is complete in the CAPACITY column.

4.   If the resize did not take effect, restart the OSD pods again. It may take multiple restarts for the
     resize to complete.

3.2. SCALING OUT STORAGE CAPACITY ON A AWS CLUSTER

OpenShift Data Foundation is highly scalable. It can be scaled out by adding new nodes with required
storage and enough hardware resources in terms of CPU and RAM. Practically there is no limit on the
number of nodes which can be added but from the support perspective 2000 nodes is the limit for
OpenShift Data Foundation.

Scaling out storage capacity can be broken down into two steps

     Adding new node

     Scaling up the storage capacity

         NOTE

         OpenShift Data Foundation does not support heterogeneous OSD/Disk sizes.


12

               CHAPTER 3. SCALING STORAGE CAPACITY OF AWS OPENSHIFT DATA FOUNDATION CLUSTER

3.2.1. Adding a node

You can add nodes to increase the storage capacity when existing worker nodes are already running at
their maximum supported OSDs or there are not enough resources to add new OSDs on the existing
nodes. It is always recommended to add nodes in the multiple of three, each of them in different failure
domains.

While it is recommended to add nodes in the multiple of three, you still have the flexibility to add one
node at a time in the flexible scaling deployment. Refer to the Knowledgebase article Verify if flexible
scaling is enabled.

                NOTE

                OpenShift Data Foundation does not support heterogeneous disk size and types. The
                new nodes to be added should have the disk of the same type and size which was used
                during OpenShift Data Foundation deployment.

3.2.1.1. Adding a node to an installer-provisioned infrastructure

Use this procedure to add a new node to an installer-provisioned infrastructure cluster.

Prerequisites

          You have administrative privilege to the OpenShift Container Platform Console.

          You have a running OpenShift Data Foundation Storage Cluster.

Procedure

     1.   Navigate to Compute → Machine Sets.

2.        On the machine set where you want to add nodes, select Edit Machine Count.

           a.  Add the amount of nodes, and click Save.

    b. Click Compute → Nodes and confirm if the new node is in Ready state.

     3.   Apply the OpenShift Data Foundation label to the new node.

           a.  For the new node, click Action menu (⋮) → Edit Labels.

           b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

                    NOTE

                    It is recommended to add 3 nodes, one each in different zones. You must add
                    3 nodes and perform this procedure for all of them. In case of bare metal
                    installer-provisioned infrastructure deployment, you must expand the cluster
                    first. For instructions, see Expanding the cluster.

Verification

     1.   Execute the following command in the terminal and verify that the new node is present in the
          output:



                    13

Red Hat OpenShift Data Foundation 4.22 Scaling storage

          $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

  2.  On the OpenShift web console, click Workloads → Pods, confirm that at least the following
      pods on the new node are in Running state:

           openshift-storage.cephfs.csi.ceph.com-*

           openshift-storage.rbd.csi.ceph.com-*

3.2.1.2. Adding a node to an user-provisioned infrastructure

Prerequisites

      You have administrative privilege to the OpenShift Container Platform Console.

      You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.    Depending on the type of infrastructure, perform the following steps:

       a.  Get a new machine with the required infrastructure. See Platform requirements.

       b.  Create a new OpenShift Container Platform worker node using the new machine.

2.    Check for certificate signing requests (CSRs) that are in Pending state.

          $ oc get csr

3.    Approve all the required CSRs for the new node.

          $ oc adm certificate approve <Certificate_Name>

      <Certificate_Name>
          Is the name of the CSR.

   4. Click Compute → Nodes, confirm if the new node is in Ready state.

5.    Apply the OpenShift Data Foundation label to the new node using any one of the following:

      From User interface

           a.  For the new node, click Action Menu (⋮) → Edit Labels.

           b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

      From Command line interface

               Apply the OpenShift Data Foundation label to the new node.

               $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

               <new_node_name>
               Is the name of the new node.

14

         CHAPTER 3. SCALING STORAGE CAPACITY OF AWS OPENSHIFT DATA FOUNDATION CLUSTER

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

         openshift-storage.cephfs.csi.ceph.com-*

         openshift-storage.rbd.csi.ceph.com-*

3.2.2. Scaling up storage capacity

To scale up storage capacity, see Scaling up storage capacity on a cluster .

3.3. ENABLING AUTOMATIC CAPACITY SCALING

You can enable automatic capacity scaling on clusters deployed using dynamic storage devices. When
automatic capacity scaling is enabled, additional raw capacity equivalent to the configured deployment
size is automatically added to the cluster when used capacity reaches 70%. This ensures your
deployment scales seamlessly to meet demand.

This option is disabled in lean profile mode, LSO deployment, and external mode deployment.

             IMPORTANT

             This may incur additional costs for the underlying storage.

Prerequisites

     Administrative privilege to the OpenShift Container Platform Console.

     A running OpenShift Data Foundation Storage Cluster.

             IMPORTANT

             IBM cloud 10iops-tier workloads are not supported.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Storage → Storage cluster.

3.   Click Actions at the top right of the page to extend the options menu.

4.   Select Automatic capacity scaling from the options menu.

5.   In the Automatic capacity scaling page, select the Enable automatic capacity scaling for
     your cluster checkbox.

6.   Set the cluster expansion limit from the dropdown. This is the maximum the cluster can expand
                 15

Red Hat OpenShift Data Foundation 4.22 Scaling storage

       6. Set the cluster expansion limit from the dropdown. This is the maximum the cluster can expand
        in the cloud. Automatic scaling is suspended if this limit is exceeded.

       7. Click Save changes.

  3.4. SCALING NON-RESILIENT POOLS

  After enabling the single replica feature, you can increase the number of non-resilient pools by updating
  the pool count in the StorageCluster custom resource. Each additional pool is created in a separate
  failure domain and requires one new storage device per failure domain.


     WARNING
    Non-resilient pools have no data replication. Increasing the pool count does not
    provide data redundancy. Ensure that your applications manage their own
    replication before scaling.




    Prerequisites

         The single replica feature is already enabled. For more information, see Storage class with single
         replica.

         For LSO-based deployments, ensure that the required number of unused local storage devices
         are correctly labeled on the nodes before scaling.

    Procedure

    1.   Run the following command to increase the non-resilient pool count:

             $ oc patch storagecluster ocs-storagecluster \
             -n openshift-storage \
             --type=json \
             --patch='[
               {
                 "op": "replace",
                 "path": "/spec/managedResources/cephNonResilientPools/count",
                 "value": 2
               }
             ]'

         Replace 2 with the desired number of non-resilient pools. The value must not exceed the
         number of available failure domains.

                       NOTE

                       For LSO-based deployments, the operator creates one OSD per local device per
                       failure domain. Ensure that the unused local devices are available on the nodes
                       before running this command.



    16

         CHAPTER 3. SCALING STORAGE CAPACITY OF AWS OPENSHIFT DATA FOUNDATION CLUSTER

    2.  Verify that the StorageCluster is in Ready state:

         $ oc get storagecluster -n openshift-storage

        Example output:

         NAME     AGE     PHASE EXTERNAL                 CREATED AT VERSION
         ocs-storagecluster 10m Ready     2024-02-05T13:56:15Z     4.22.0

3. Verify that the new CephBlockPools are created for each failure domain and are in Ready
   state:

         $ oc get cephblockpools -n openshift-storage

        Example output:

         NAME        PHASE
         ocs-storagecluster-cephblockpool  Ready
         ocs-storagecluster-cephblockpool-us-east-1a     Ready
         ocs-storagecluster-cephblockpool-us-east-1b     Ready
         ocs-storagecluster-cephblockpool-us-east-1c     Ready

    4.  Verify that the new OSD pods are in Running state:

         $ oc get pods -n openshift-storage | grep osd










                                                         17

Red Hat OpenShift Data Foundation 4.22 Scaling storage

        CHAPTER 4. SCALING STORAGE OF BARE METAL
        OPENSHIFT DATA FOUNDATION CLUSTER

  To scale the storage capacity of your configured Red Hat OpenShift Data Foundation worker nodes on
  your bare metal cluster, you can increase the capacity by adding three disks at a time. Three disks are
  needed since OpenShift Data Foundation uses a replica count of 3 to maintain the high availability. So
  the amount of storage consumed is three times the usable space.

        NOTE

        Usable space might vary when encryption is enabled or replica 2 pools are being used.

  4.1. SCALING UP A CLUSTER CREATED USING LOCAL STORAGE
  DEVICES

  To scale up an OpenShift Data Foundation cluster which was created using local storage devices, you
  need to add a new disk to the storage node. The new disks size must be of the same size as the disks
  used during the deployment because OpenShift Data Foundation does not support heterogeneous
  disks/OSDs.

  For deployments having three failure domains, you can scale up the storage by adding disks in the
  multiples of three, with the same number of disks coming from nodes in each of the failure domains. For
  example, if we scale by adding six disks, two disks are taken from nodes in each of the three failure
  domains. If the number of disks is not in multiples of three, it will only consume the disk to the maximum
  in the multiple of three while the remaining disks remain unused.

  For deployments having less than three failure domains, there is a flexibility to add any number of disks.
  Make sure to verify that flexible scaling is enabled. For information, refer to the Knowledgebase article
  Verify if flexible scaling is enabled .

        NOTE

        Flexible scaling features get enabled at the time of deployment and cannot be enabled
        or disabled later on.

  Prerequisites

        Administrative privilege to the OpenShift Container Platform Console.

        A running OpenShift Data Foundation Storage Cluster.

        Make sure that the disks to be used for scaling are attached to the storage node

        Make sure that LocalVolumeDiscovery and LocalVolumeSet objects are created.

  To add capacity, you can either use a storage class that you provisioned during the deployment or any
  other storage class that matches the filter.

  Procedure

        1. In the OpenShift Web Console, click Storage → Storage cluster.

        a. Click Actions at the top right of the page to extend the options menu.


    18

            CHAPTER 4. SCALING STORAGE OF BARE METAL OPENSHIFT DATA FOUNDATION CLUSTER

     b.  Select Add Capacity from the options menu.

     c.  Select the Storage Class for which you added disks or the new storage class depending on
         your requirement. Available Capacity displayed is based on the local disks available in
         storage class.

     d.  Click Add.

2.  To check the status, navigate to Storage → Storage cluster and verify that the Storage
    Cluster in the Status card has a green tick.

Verification

    Verify the Raw Capacity card.

     a.  In the OpenShift Web Console, click Storage → Storage cluster.

     b.  In the Block and File tab, check the Raw Capacity card.
         Note that the capacity increases based on your selections.

                       NOTE

                       The raw capacity does not take replication into account and shows the full
                       capacity.

    Verify that the new OSDs and their corresponding new Persistent Volume Claims (PVCs) are
    created.

         To view the state of the newly created OSDs:

            a.  Click Workloads → Pods from the OpenShift Web Console.

            b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

         To view the state of the PVCs:

            a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

            b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.

     a. Identify the nodes where the new OSD pods are running.


                        19

Red Hat OpenShift Data Foundation 4.22 Scaling storage

$ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
pod-name>

<OSD-pod-name>
Is the name of the OSD pod.
For example:

 $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
 ceph-osd-0-544db49d7f-qrgqm

Example output:

 NODE
 compute-1


b. For each of the nodes identified in the previous step, do the following:

   i.   Create a debug pod and open a chroot environment for the selected host(s).

        $ oc debug node/<node-name>

        <node-name>
        Is the name of the node.

        $ chroot /host


         ii. Check for the crypt keyword beside the ocs-deviceset names.

             $ lsblk

             IMPORTANT

             Cluster reduction is supported only with the Red Hat Support Team’s assistance.

4.1.1. Scaling up storage capacity on a cluster by resizing existing OSDs

To increase the storage capacity on a cluster, you can add storage capacity by resizing existing OSDs.

Prerequisites

       You have administrative privilege to the OpenShift Container Platform Console.

       You have a running OpenShift Data Foundation Storage Cluster.

Procedure

  1.   Update the dataPVCTemplate size for the storageDeviceSets in the StorageCluster CR with
       the new desired size. For example:

         apiVersion: ocs.openshift.io/v1

20

        CHAPTER 4. SCALING STORAGE OF BARE METAL OPENSHIFT DATA FOUNDATION CLUSTER

      kind: StorageCluster
      metadata:
       namespace: openshift-storage
       name: example-storagecluster
      spec:
       manageNodes: false
       monPVCTemplate:
        spec:
         storageClassName:
         accessModes:
         - ReadWriteOnce
         resources:
          requests:
           storage: 10Gi
       storageDeviceSets:
       - name: example-deviceset
        count: 3
        resources: {}
        placement: {}
        dataPVCTemplate:
         spec:
          storageClassName:
          accessModes:
          - ReadWriteOnce
          volumeMode: Block
          resources:
           requests:
            storage: 1Ti
        portable: true

     In this example YAML, the storage parameter under storageDeviceSets reflects the size
     change of 1Ti.

 2.  Wait for the OSDs to restart.

 3.  Confirm that the resize took effect:

      $ oc get pvc -l ceph.rook.io/DeviceSet

     Example output:

      % oc get pvc -l ceph.rook.io/DeviceSet
      NAME     STATUS     VOLUME        CAPACITY     ACCESS MODES
      STORAGECLASS        AGE
    rook-ceph-osd-a4681 Bound pvc-c3868d38-6bd9-4c8d-91e1-a4681e4cb8c1 1Ti
      RWO     gp2-csi     176m

     Verify the resize is complete in the CAPACITY column.

4.   If the resize did not take effect, restart the OSD pods again. It may take multiple restarts for the
     resize to complete.

4.2. SCALING OUT STORAGE CAPACITY ON A BARE METAL CLUSTER

OpenShift Data Foundation is highly scalable. It can be scaled out by adding new nodes with required

              21

Red Hat OpenShift Data Foundation 4.22 Scaling storage

storage and enough hardware resources in terms of CPU and RAM. There is no limit on the number of
nodes which can be added. However, from the technical support perspective, 2000 nodes is the limit for
OpenShift Data Foundation.

Scaling out storage capacity can be broken down into two steps

     Adding new node

     Scaling up the storage capacity

             NOTE

             OpenShift Data Foundation does not support heterogeneous OSD/Disk sizes.

4.2.1. Adding a node

You can add nodes to increase the storage capacity when existing worker nodes are already running at
their maximum supported OSDs or there are not enough resources to add new OSDs on the existing
nodes. It is always recommended to add nodes in the multiple of three, each of them in different failure
domains.

While it is recommended to add nodes in the multiple of three, you still have the flexibility to add one
node at a time in the flexible scaling deployment. Refer to the Knowledgebase article Verify if flexible
scaling is enabled.

             NOTE

             OpenShift Data Foundation does not support heterogeneous disk size and types. The
             new nodes to be added should have the disk of the same type and size which was used
             during OpenShift Data Foundation deployment.

4.2.1.1. Adding a node to an installer-provisioned infrastructure

Use this procedure to add a new node to an installer-provisioned infrastructure cluster.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Navigate to Compute → Machine Sets.

2.   On the machine set where you want to add nodes, select Edit Machine Count.

        a.  Add the amount of nodes, and click Save.

    b. Click Compute → Nodes and confirm if the new node is in Ready state.

3.   Apply the OpenShift Data Foundation label to the new node.

        a.  For the new node, click Action menu (⋮) → Edit Labels.



22

             CHAPTER 4. SCALING STORAGE OF BARE METAL OPENSHIFT DATA FOUNDATION CLUSTER

      b. Add cluster.ocs.openshift.io/openshift-storage, and click Save.

                 NOTE

                 It is recommended to add 3 nodes, one each in different zones. You must add
                 3 nodes and perform this procedure for all of them. In case of bare metal
                 installer-provisioned infrastructure deployment, you must expand the cluster
                 first. For instructions, see Expanding the cluster.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

4.2.1.2. Adding a node using a local storage device

You can add nodes to increase the storage capacity when existing worker nodes are already running at
their maximum supported OSDs or when there are not enough resources to add new OSDs on the
existing nodes.

Add nodes in the multiple of 3, each of them in different failure domains. Though it is recommended to
add nodes in multiples of 3 nodes, you have the flexibility to add one node at a time in flexible scaling
deployment. See Knowledgebase article Verify if flexible scaling is enabled

             NOTE

             OpenShift Data Foundation does not support heterogeneous disk size and types. The
             new nodes to be added should have the disk of the same type and size which was used
             during initial OpenShift Data Foundation deployment.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Depending on the type of infrastructure, perform the following steps:

      a.  Get a new machine with the required infrastructure. See Platform requirements.

      b.  Create a new OpenShift Container Platform worker node using the new machine.

2.   Check for certificate signing requests (CSRs) that are in Pending state.


                 23

Red Hat OpenShift Data Foundation 4.22 Scaling storage

        $ oc get csr

3.  Approve all the required CSRs for the new node.

        $ oc adm certificate approve <Certificate_Name>

    <Certificate_Name>
        Is the name of the CSR.

   4. Click Compute → Nodes, confirm if the new node is in Ready state.

5.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

    From User interface

         a.  For the new node, click Action Menu (⋮) → Edit Labels.

         b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

    From Command line interface

             Apply the OpenShift Data Foundation label to the new node.

             $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

             <new_node_name>
             Is the name of the new node.

6.  Click Ecosystem → Installed Operators from the OpenShift Web Console.
    From the Project drop-down list, make sure to select the project where the Local Storage
    Operator is installed.

7.  Click Local Storage.

8.  Click the Local Volume Discovery tab.

     a.  Beside the LocalVolumeDiscovery, click Action menu (⋮) → Edit Local Volume
         Discovery.

     b.  In the YAML, add the hostname of the new node in the values field under the node
         selector.

     c.  Click Save.

9.  Click the Local Volume Sets tab.

     a.  Beside the LocalVolumeSet, click Action menu (⋮) → Edit Local Volume Set.

     b.  In the YAML, add the hostname of the new node in the values field under the node
         selector.







24

CHAPTER 4. SCALING STORAGE OF BARE METAL OPENSHIFT DATA FOUNDATION CLUSTER










         YAML showing the addition of new hostnames

     c.  Click Save.

                    NOTE

                    It is recommended to add 3 nodes, one each in different zones. You must add
                    3 nodes and perform this procedure for all of them.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

        $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

         openshift-storage.cephfs.csi.ceph.com-*

         openshift-storage.rbd.csi.ceph.com-*

4.2.2. Scaling up storage capacity

To scale up storage capacity, see Scaling up storage by adding capacity .










                    25

Red Hat OpenShift Data Foundation 4.22 Scaling storage

                                                       CHAPTER 5. SCALING STORAGE USING MULTIPLE DEVICE
    CLASS IN THE SAME CLUSTER FOR LOCAL STORAGE
                   DEPLOYMENTS

OpenShift Data Foundation supports creating multiple device classes for OSDs within the same cluster.
Defining additional device classes provides flexibility in how storage devices are organized and allows
you to:

       Use different types of disks on the same nodes

       Use disks of the same type on the same nodes

       Use different-sized disks of the same type on the same or different nodes

       Isolate disks of the same type across different sets of nodes

       Combine different resource types, such as local disks and SAN-based LUNs

Overview of the required tasks

To scale storage by using multiple device classes, complete the following tasks:

1.     Prepare the disks.
       Attach new disks on the same or additional nodes and ensure they can be uniquely identified by
       a LocalVolumeSet.

               NOTE

                             Update the maxSize or disksFilter parameter of the existing LocalVolumeSet
               (such as localblock) to prevent it from claiming the new PVs.

2.     Create a new LocalVolumeSet.
       Define a LocalVolumeSet resource that represents the new device class.

3.     Attach new storage.
       Attach the storage created by the new LocalVolumeSet so that it can be consumed by the
       cluster.

5.1. CREATING A NEW LOCALVOLUMESET

Use this procedure when you want to use devices of the same type but with different sizes.

Prerequisites

       Update the maxSize parameter of the existing LocalVolumeSet (for example,         localblock) to
       ensure that it does not claim the newly created persistent volumes (PVs).

          $ oc -n openshift-local-storage patch localvolumesets.local.storage.openshift.io localblock \
        -n openshift-local-storage \
        -p '{"spec": {"deviceInclusionSpec": {"maxSize": "120Gi"}}}' \
        --type merge


In this example, the existing LocalVolumeSet created during deployment might not have a
26

       In this example, the existing LocalVolumeSet created during deployment might not have a
       maxSize value. Setting the limit to 120Gi ensures that new disks with a higher size (for example,
       130Gi) are claimed only by the new LocalVolumeSet and do not overlap with the existing one.

       When creating the new LocalVolumeSet, plan a unique filter to identify the intended disks. You
       can differentiate disks by node, disk size, or disk type.

       Add the new disks. For example, add three new SSD or NVMe disks sized 130Gi.

Procedure

  1.   In the OpenShift Web Console, click Ecosystem → Installed Operators.

     2. From the Project drop‑down list, select the project where the Local Storage Operator is
        installed.

  3.   Click Local Storage.

  4.   Click the Local Volume Sets tab.

  5.   On the Local Volume Sets page, click Create Local Volume Set.

  6.   Enter a name for the LocalVolumeSet and the associated StorageClass.

                    NOTE

                    The storage class name defaults to the LocalVolumeSet name, but it can be
                    modified.

  7.   Under Filter Disks By, select one of the following options:

           Disks on all nodes: Uses all matching disks across all nodes.

           Disks on selected nodes: Uses matching disks only on the nodes you choose.

  8.   From Disk Type, select SSD/NVMe from the list.

  9.   Expand the Advanced section and configure the following:

       a.  Volume Mode: Ensure that Block is selected.

       b.  Device Type: Select one or more device types.

       c.  Disk Size: Set the minimum and maximum disk size to include.

       d.  Maximum Disks Limit: Specify the maximum number of PVs that can be created per node.

                     NOTE

                     If this field is empty, PVs are created for all matching disks.

 10.   Click Create.

 11.   Wait for the newly created PVs associated with the new LocalVolumeSet to become available.

    Verification

                         27

Red Hat OpenShift Data Foundation 4.22 Scaling storage

 Verification

          Verify that the local volume set is created:

           $ oc get localvolumeset -n openshift-local-storage
           NAME         AGE
           localblock   16h
           localvolume2  43m

          Verify the local storage class

           oc get sc
           NAME                 PROVISIONER                                        RECLAIMPOLICY
     VOLUMEBINDINGMODE ALLOWVOLUMEEXPANSION AGE
           localblock           kubernetes.io/no-provisioner               Delete
           WaitForFirstConsumer  false      15h
           localvolume2          kubernetes.io/no-provisioner              Delete
           WaitForFirstConsumer  false      27m
           ocs-storagecluster-ceph-rbd                openshift-storage.rbd.csi.ceph.com  Delete  Immediate
           true          15h
           [...]

          Verify the PV by waiting for it to be available and it must be using the new storage class
          localvolume2:
          For example:

           $ oc get pv | grep localvolume2
           local-pv-14c0b1d          130Gi     RWO     Delete                  Available
           localvolume2          <unset>                 8m55s
           local-pv-41d0d077                     130Gi               RWO   Delete   Available
           localvolume2          <unset>                 7m24s
           local-pv-6c57a345                     130Gi               RWO   Delete   Available
           localvolume2          <unset>        5m4s

 5.2. ATTACHING STORAGE FOR A NEW DEVICE SET

 Prerequisites

          Make sure to label the required nodes in the new Local Volume Set with
          cluster.ocs.openshift.io/openshift-storage=''.

           oc label node <node-name > cluster.ocs.openshift.io/openshift-storage='' --overwrite

  Procedure

 1.       In the OpenShift Web Console, navigate to Storage → Storage cluster.

    2. Click Actions at the top right of the page and select Attach Storage.

      3.  From LSO StorageClass, select the newly created local storage class.

      4.  To enable encryption for the device set, select Enable encryption on device set.


           NOTE
 28

                   NOTE

                   Enabling or disabling OSD encryption using this option overrides the cluster‑wide
                   encryption setting for these new OSD storage devices.

 5.   Enter a name for the Device class.
      This is a label applied to OSDs to distinguish this set of devices from others (for example, by
      size, type, or node group). This option is available only when creating new devices of the same
      size on the same set of storage devices.

 6.   Select the Volume type.

 7.   Enter a name for the Pool.

 8.   Select the Data protection policy.
      It is recommended to choose 3‑way Replication.

 9.   Select the Data compression option to enable compression within replicas for improved
      storage efficiency.

10.   Select the Reclaim Policy.

11.   Select the Volume Binding Mode.

12.   Enter a name for the New User Storage Class.

13.   To generate an encryption key for each persistent volume created through this storage class,
      select Enable Encryption on StorageClass.

14.   Click Attach.

   Verification

      Verify that the PV and PVCs are in Bound state.
      For example:

       $ oc get pv | grep localvolume2
       local-pv-14c0b1d        130Gi              RWO          Delete      Bound    openshift-
       storage/localvolume2-0-data-0kp29f              localvolume2         <unset>
       31m
       local-pv-41d0d077                130Gi        RWO   Delete           Bound   openshift-
       storage/localvolume2-0-data-2vwk54                localvolume2       <unset>
       30m
       local-pv-6c57a345                130Gi        RWO   Delete           Bound   openshift-
       storage/localvolume2-0-data-1255ts              localvolume2         <unset>
       28m

       $ oc get pvc | grep localvolume2
       localvolume2-0-data-0kp29f             Bound    local-pv-14c0b1d             130Gi
       RWO         localvolume2               <unset>     19m
       localvolume2-0-data-1255ts             Bound    local-pv-6c57a345              130Gi
       RWO         localvolume2               <unset>     19m
       localvolume2-0-data-2vwk54             Bound    local-pv-41d0d077            130Gi
       RWO         localvolume2               <unset>     19m



                                                                                    29

Red Hat OpenShift Data Foundation 4.22 Scaling storage

  Verify that the new OSDs are created successfully and all the OSDs are running.
  For example:

   $ oc -n openshift-storage get pods -l app=rook-ceph-osd
   NAME                      READY      STATUS           RESTARTS AGE
   rook-ceph-osd-0-7899b89478-bmg2l      2/2          Running 2   16h
   rook-ceph-osd-1-6d7df8dbfc-z29jx  2/2           Running    2   16h
   rook-ceph-osd-2-66b9dc8cd7-vd77d      2/2          Running   2 16h
   rook-ceph-osd-3-79b59f44d7-cb9m7      2/2          Running 0   15m
   rook-ceph-osd-4-766bcd8646-z4zvz      2/2          Running 0   15m
   rook-ceph-osd-5-845554d78-qqcm4       2/2          Running 0   15m

  Verify that the storagecluster is in Ready state and the Ceph cluster health is      OK.
  For example:

   $ oc -n openshift-storage get storagecluster
   NAME              AGE     PHASE     EXTERNAL             CREATED AT                 VERSION
   ocs-storagecluster   2d2h    Ready                            2025-02-11T12:03:37Z  4.22.0

   $ oc -n openshift-storage get cephcluster
   NAME                      DATADIRHOSTPATH               MONCOUNT             AGE    PHASE MESSAGE
   HEALTH       EXTERNAL     FSID
   ocs-storagecluster-cephcluster   /var/lib/rook        3               2d2h Ready Cluster created
   successfully  HEALTH_OK                               8ec4e898-5386-42c7-b01b-169fe8f08ba4

  In the Ceph cluster, check the Ceph OSD tree to see if the new device classes are spread
  correctly and OSDs are up.
  For example:

   $ oc rsh -n openshift-storage $(oc get pods -o wide -n openshift-storage|grep tool|awk
   '{print$1}') ceph osd tree
   ID CLASS      WEIGHT      TYPE NAME                   STATUS REWEIGHT PRI-AFF
   -1          0.67406 root default
   -7          0.22469    host compute-0
    3 localvolume2 0.12700       osd.3                up           1.00000 1.00000
    0        ssd 0.09769        osd.0          up     1.00000 1.00000
   -3          0.22469    host compute-1
    4 localvolume2 0.12700       osd.4                up           1.00000 1.00000
    2        ssd 0.09769        osd.2          up     1.00000 1.00000
   -5          0.22469    host compute-2
    5 localvolume2 0.12700       osd.5                up           1.00000 1.00000
    1        ssd 0.09769        osd.1          up     1.00000 1.00000

  Verify that the user storageclass is created.
  For example:

   $ oc get sc
   NAME                      PROVISIONER                               RECLAIMPOLICY
    VOLUMEBINDINGMODE ALLOWVOLUMEEXPANSION AGE
   localblock      kubernetes.io/no-provisioner              Delete
   WaitForFirstConsumer      false             16h
   localvolume2              kubernetes.io/no-provisioner        Delete
   WaitForFirstConsumer      false             63m
   ocs-storagecluster-ceph-rbd     openshift-storage.rbd.csi.ceph.com         Delete       Immediate

30

  true          16h
  ocs-storagecluster-ceph-rgw                        openshift-storage.ceph.rook.io/bucket  Delete     Immediate
  false         16h
  ocs-storagecluster-cephfs       openshift-storage.cephfs.csi.ceph.com Delete                         Immediate
  true          16h
  openshift-storage.noobaa.io     openshift-storage.noobaa.io/obc      Delete                     Immediate
  false         16h
  ssd2                 openshift-storage.rbd.csi.ceph.com   Delete
  WaitForFirstConsumer            true           16m
  thin-csi             csi.vsphere.vmware.com                 Delete                        WaitForFirstConsumer
  true          17h
  thin-csi-odf                    csi.vsphere.vmware.com    Delete
  WaitForFirstConsumer            true           16h

5.3. ATTACHING STORAGE FOR A NEW DEVICE SET BY USING DISKS
OF SAME TYPE ON THE SAME NODES

Use this section when creating new devices of the same size on the same set of storage devices where
you need to define a custom device class.

Pools write data only to OSDs that match the selected device class. By default, all OSDs use the ssd
device class, so all pools can use all OSDs. If a custom device class is set, only new pools that explicitly
use that device class will write to those OSDs. Existing pools will continue using only the OSDs that
match their original device class.

Prerequisites

 Ensure that the worker nodes hosting the OSDs have a sufficient number of available disks.

 Make sure to label the required nodes in a new Local Volume Set with
 cluster.ocs.openshift.io/openshift-storage=''.

  $ oc label node <node-name > cluster.ocs.openshift.io/openshift-storage='' --overwrite

 Ensure that an adequate number of Persistent Volumes (PVs) are available to match the
 number of disks that can be used for OSD provisioning.
 To verify the available local block PVs, run the following command:

  $ oc get pv | grep localblock

 Example output:

  $ oc get pv | grep localblock
  local-pv-1217fe9b                          100Gi   RWO      Delete                         Available
  localblock                      <unset>            4m22s
  local-pv-18699823                          100Gi     RWO      Delete                        Bound    openshift-
  storage/ocs-deviceset-localblock-0-data-122gpk          localblock                          <unset>
  116m
  local-pv-7357d060                          100Gi     RWO      Delete                        Available
  localblock                      <unset>                             9m9s
  local-pv-c31f65ba                          100Gi   RWO      Delete                         Available
  localblock                      <unset>                             2m2s
  local-pv-d666f718                          100Gi   RWO      Delete                         Bound     openshift-
  storage/ocs-deviceset-localblock-0-data-0cgff6       localblock                             <unset>

                                                                                                           31

    Red Hat OpenShift Data Foundation 4.22 Scaling storage

         151m
         local-pv-fa54f87b        100Gi     RWO     Delete     Bound     openshift-
         storage/ocs-deviceset-localblock-0-data-2d277g  localblock      <unset>
         151m

        In this example, there are three available disks, which correspond to three PVs in the Available
        state.

Procedure

  1.    In the OpenShift Web Console, navigate to Storage → Storage cluster.

       2. Click Actions at the top right of the page and select Attach Storage.

  3.    From LSO StorageClass, select the newly created local storage class.

    4.  To enable encryption for the device set, select Enable encryption on device set.

                     NOTE

                     Enabling or disabling OSD encryption using this option overrides the cluster‑wide
                     encryption setting for these new OSD storage devices.

  5.    Enter a name for the Device class.
        This is a label applied to OSDs to distinguish this set of devices from others (for example, by
        size, type, or node group). This option is available only when creating new devices of the same
        size on the same set of storage devices.

  6.    Select the Volume type.

  7.    Enter a name for the Pool.

  8.    Select the Data protection policy.
        It is recommended to choose 3‑way Replication.

  9.    Select the Data compression option to enable compression within replicas for improved
        storage efficiency.

 10.    Select the Reclaim Policy.

 11.    Select the Volume Binding Mode.

 12.    Enter a name for the New User Storage Class.

 13.    To generate an encryption key for each persistent volume created through this storage class,
        select Enable encryption on StorageClass.

 14.    Click Attach.

    Verification

        Verify that the PV and PVCs are in Bound state.
        For example:

         $ oc get pv | grep localblock-1


    32

 local-pv-1217fe9b                    100Gi     RWO            Delete          Bound   openshift-
 storage/localblock-1-0-data-28ptwg              localblock                    <unset>
 84m
 local-pv-7357d060                      100Gi     RWO       Delete             Bound    openshift-
 storage/localblock-1-0-data-15td2h              localblock                    <unset>
 89m
 local-pv-c31f65ba                    100Gi     RWO            Delete          Bound   openshift-
 storage/localblock-1-0-data-08bjsj             localblock                    <unset>
 82m

 $ oc get pvc | grep localblock-1
 localblock-1-0-data-08bjsj      Bound          local-pv-c31f65ba                   100Gi     RWO
 localblock      <unset>                   79m
 localblock-1-0-data-15td2h             Bound       local-pv-7357d060                   100Gi
 RWO         localblock                <unset>           79m
 localblock-1-0-data-28ptwg             Bound       local-pv-1217fe9b                  100Gi
 RWO         localblock                <unset>           79m

Verify that the new OSDs are created successfully and all the OSDs are running.
For example:

 $ oc get pods -n openshift-storage -l app=rook-ceph-osd
 NAME                        READY      STATUS    RESTARTS   AGE
 rook-ceph-osd-0-85f6969cff-l67mz      2/2     Running  2    173m
 rook-ceph-osd-1-69c8f7c797-bvvpg         2/2   Running  2   172m
 rook-ceph-osd-2-5cbc9cc8cb-rjmbv       2/2     Running  2   172m
 rook-ceph-osd-3-76c84945b4-4mdv8      2/2      Running         0    84m
 rook-ceph-osd-4-78c6b8bf87-gkh52        2/2    Running     0   84m
 rook-ceph-osd-5-7876b5fb6f-m5j9q       2/2     Running  0   84m

Verify that the storagecluster is in Ready state and the Ceph cluster health is OK.
For example:

 $ oc -n openshift-storage get storagecluster
 NAME         AGE           PHASE     EXTERNAL   CREATED AT                    VERSION
 ocs-storagecluster   4h34m  Ready              2026-03-04T12:17:21Z          4.22.0

 $ oc -n openshift-storage get cephcluster
 NAME                       DATADIRHOSTPATH              MONCOUNT        AGE    PHASE   MESSAGE
 HEALTH           EXTERNAL   FSID
 ocs-storagecluster-cephcluster  /var/lib/rook     3      4h34m          Ready  Cluster created
 successfully HEALTH_OK                0daaeaae-d8cd-45e6-9df8-f23fe3a70263

In the Ceph cluster, check the Ceph OSD tree to see if the new device classes are spread
correctly and OSDs are up.
For example:

 $ oc rsh -n openshift-storage $(oc get pods -o wide -n openshift-storage|grep tool|awk
 '{print$1}') ceph osd tree
 ID CLASS     WEIGHT                TYPE NAME   STATUS REWEIGHT PRI-AFF
 -1          0.58612 root default
 -5          0.19537     host compute-0
    4 localblock-1 0.09769      osd.4      up            1.00000 1.00000


                                                                                        33

Red Hat OpenShift Data Foundation 4.22 Scaling storage

  1     ssd 0.09769       osd.1              up     1.00000 1.00000
 -7      0.19537        host compute-1
  5 localblock-1 0.09769     osd.5                  up 1.00000 1.00000
  2     ssd 0.09769       osd.2              up     1.00000 1.00000
 -3      0.19537        host compute-2
  3 localblock-1 0.09769     osd.3                  up 1.00000 1.00000
  0     ssd 0.09769       osd.0              up     1.00000 1.00000

Verify that the user storageclass is created.
For example:

 $ oc get sc
 NAME                  PROVISIONER                               RECLAIMPOLICY
    VOLUMEBINDINGMODE ALLOWVOLUMEEXPANSION AGE
 localblock           kubernetes.io/no-provisioner              Delete
 WaitForFirstConsumer    false               5h30m
 localblock-1-sc         openshift-storage.rbd.csi.ceph.com     Delete
 WaitForFirstConsumer    true                176m
 ocs-storagecluster-ceph-rbd        openshift-storage.rbd.csi.ceph.com  Delete  Immediate
 true             4h25m
 ocs-storagecluster-ceph-rgw    openshift-storage.ceph.rook.io/bucket   Delete  Immediate
 false            4h31m
 ocs-storagecluster-cephfs     openshift-storage.cephfs.csi.ceph.com    Delete  Immediate
 true             4h25m
 openshift-storage.noobaa.io    openshift-storage.noobaa.io/obc     Delete     Immediate
 false            4h18m
 thin-csi                       csi.vsphere.vmware.com  Delete          WaitForFirstConsumer
 true             6h8m
 thin-csi-odf         csi.vsphere.vmware.com      Delete
 WaitForFirstConsumer    true                4h31m










34

             CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER

CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT
              DATA FOUNDATION CLUSTER

6.1. SCALING UP STORAGE ON A DYNAMICALLY PROVISIONED
VMWARE CLUSTER

To increase the storage capacity in a dynamically created storage cluster on a VMware user-provisioned
infrastructure, you can add storage capacity and performance to your configured Red Hat OpenShift
Data Foundation worker nodes.

Prerequisites

     Administrative privilege to the OpenShift Container Platform Console.

     A running OpenShift Data Foundation Storage Cluster.

     Make sure that the disk is of the same size and type as the disk used during initial deployment.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Storage → Storage cluster.

      a.  Click Actions at the top right of the page to extend the options menu.

      b.  Select Add Capacity from the options menu.

      c.  Select the Storage Class. Choose the storage class which you wish to use to provision new
          storage devices.

      d.  Click Add.

3.   To check the status, navigate to Storage → Storage cluster and verify that Storage Cluster in
     the Status card has a green tick.

Verification

     Verify the Raw Capacity card.

      a.  In the OpenShift Web Console, click Storage → Storage cluster.

      b.  In the Block and File tab, check the Raw Capacity card.
          Note that the capacity increases based on your selections.

                    NOTE

                    The raw capacity does not take replication into account and shows the full
                    capacity.

     Verify that the new OSDs and their corresponding new Persistent Volume Claims (PVCs) are
     created.

          To view the state of the newly created OSDs:

                        35

Red Hat OpenShift Data Foundation 4.22 Scaling storage

     a.  Click Workloads → Pods from the OpenShift Web Console.

     b.  Select openshift-storage from the Project drop-down list.

                    NOTE

                    If the Show default projects option is disabled, use the toggle button to
                    list all the default projects.

To view the state of the PVCs:

     a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

     b.  Select openshift-storage from the Project drop-down list.

                    NOTE

                    If the Show default projects option is disabled, use the toggle button to
                    list all the default projects.

Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
are encrypted.

a.  Identify the nodes where the new OSD pods are running.

        $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
        pod-name>

    <OSD-pod-name>
        Is the name of the OSD pod.
        For example:

          $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
          ceph-osd-0-544db49d7f-qrgqm

        Example output:

          NODE
          compute-1


b. For each of the nodes identified in the previous step, do the following:

     i.   Create a debug pod and open a chroot environment for the selected hosts.

          $ oc debug node/<node-name>

          <node-name>
          Is the name of the node.

          $ chroot /host




36

               CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER

           ii. Check for the crypt keyword beside the ocs-deviceset names.

               $ lsblk

             IMPORTANT

             Cluster reduction is supported only with the Red Hat Support Team’s assistance.

6.2. SCALING UP A CLUSTER CREATED USING LOCAL STORAGE
DEVICES

To scale up an OpenShift Data Foundation cluster which was created using local storage devices, you
need to add a new disk to the storage node. The new disks size must be of the same size as the disks
used during the deployment because OpenShift Data Foundation does not support heterogeneous
disks/OSDs.

For deployments having three failure domains, you can scale up the storage by adding disks in the
multiples of three, with the same number of disks coming from nodes in each of the failure domains. For
example, if we scale by adding six disks, two disks are taken from nodes in each of the three failure
domains. If the number of disks is not in multiples of three, it will only consume the disk to the maximum
in the multiple of three while the remaining disks remain unused.

For deployments having less than three failure domains, there is a flexibility to add any number of disks.
Make sure to verify that flexible scaling is enabled. For information, refer to the Knowledgebase article
Verify if flexible scaling is enabled .

             NOTE

             Flexible scaling features get enabled at the time of deployment and cannot be enabled
             or disabled later on.

Prerequisites

     Administrative privilege to the OpenShift Container Platform Console.

     A running OpenShift Data Foundation Storage Cluster.

     Make sure that the disks to be used for scaling are attached to the storage node

     Make sure that LocalVolumeDiscovery and LocalVolumeSet objects are created.

To add capacity, you can either use a storage class that you provisioned during the deployment or any
other storage class that matches the filter.

Procedure

1.   In the OpenShift Web Console, click Storage → Storage cluster.

      a.  Click Actions at the top right of the page to extend the options menu.

      b.  Select Add Capacity from the options menu.




      c. Select the Storage Class for which you added disks or the new storage class depending on
                 37

Red Hat OpenShift Data Foundation 4.22 Scaling storage

     c.  Select the Storage Class for which you added disks or the new storage class depending on
         your requirement. Available Capacity displayed is based on the local disks available in
         storage class.

     d.  Click Add.

2.  To check the status, navigate to Storage → Storage cluster and verify that the Storage
    Cluster in the Status card has a green tick.

Verification

    Verify the Raw Capacity card.

     a.  In the OpenShift Web Console, click Storage → Storage cluster.

     b.  In the Block and File tab, check the Raw Capacity card.
         Note that the capacity increases based on your selections.

                       NOTE

                       The raw capacity does not take replication into account and shows the full
                       capacity.

    Verify that the new OSDs and their corresponding new Persistent Volume Claims (PVCs) are
    created.

         To view the state of the newly created OSDs:

            a.  Click Workloads → Pods from the OpenShift Web Console.

            b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

         To view the state of the PVCs:

            a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

            b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.

     a. Identify the nodes where the new OSD pods are running.

               $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
               pod-name>

38

CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER


<OSD-pod-name>
Is the name of the OSD pod.
For example:

 $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
 ceph-osd-0-544db49d7f-qrgqm

Example output:

 NODE
 compute-1


b. For each of the nodes identified in the previous step, do the following:

   i.   Create a debug pod and open a chroot environment for the selected host(s).

        $ oc debug node/<node-name>

        <node-name>
        Is the name of the node.

        $ chroot /host


         ii. Check for the crypt keyword beside the ocs-deviceset names.

             $ lsblk

             IMPORTANT

             Cluster reduction is supported only with the Red Hat Support Team’s assistance.

6.2.1. Scaling up storage capacity on a cluster by resizing existing OSDs

To increase the storage capacity on a cluster, you can add storage capacity by resizing existing OSDs.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Update the dataPVCTemplate size for the storageDeviceSets in the StorageCluster CR with
     the new desired size. For example:

         apiVersion: ocs.openshift.io/v1
         kind: StorageCluster
         metadata:
         namespace: openshift-storage

                 39

Red Hat OpenShift Data Foundation 4.22 Scaling storage

        name: example-storagecluster
       spec:
        manageNodes: false
        monPVCTemplate:
         spec:
          storageClassName:
          accessModes:
          - ReadWriteOnce
          resources:
           requests:
            storage: 10Gi
        storageDeviceSets:
        - name: example-deviceset
         count: 3
         resources: {}
         placement: {}
         dataPVCTemplate:
          spec:
           storageClassName:
           accessModes:
           - ReadWriteOnce
           volumeMode: Block
           resources:
            requests:
             storage: 1Ti
         portable: true

      In this example YAML, the storage parameter under storageDeviceSets reflects the size
      change of 1Ti.

  2.  Wait for the OSDs to restart.

  3.  Confirm that the resize took effect:

       $ oc get pvc -l ceph.rook.io/DeviceSet

      Example output:

       % oc get pvc -l ceph.rook.io/DeviceSet
       NAME     STATUS     VOLUME        CAPACITY     ACCESS MODES
       STORAGECLASS        AGE
    rook-ceph-osd-a4681 Bound pvc-c3868d38-6bd9-4c8d-91e1-a4681e4cb8c1 1Ti
       RWO     gp2-csi    176m

      Verify the resize is complete in the CAPACITY column.

  4.  If the resize did not take effect, restart the OSD pods again. It may take multiple restarts for the
      resize to complete.

6.3. ENABLING AUTOMATIC CAPACITY SCALING

You can enable automatic capacity scaling on clusters deployed using dynamic storage devices. When
automatic capacity scaling is enabled, additional raw capacity equivalent to the configured deployment
size is automatically added to the cluster when used capacity reaches 70%. This ensures your
deployment scales seamlessly to meet demand.

40

             CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER

This option is disabled in lean profile mode, LSO deployment, and external mode deployment.

             IMPORTANT

             This may incur additional costs for the underlying storage.

Prerequisites

         Administrative privilege to the OpenShift Container Platform Console.

         A running OpenShift Data Foundation Storage Cluster.

             IMPORTANT

             IBM cloud 10iops-tier workloads are not supported.

Procedure

    1.   Log in to the OpenShift Web Console.

    2.   Click Storage → Storage cluster.

    3.   Click Actions at the top right of the page to extend the options menu.

    4.   Select Automatic capacity scaling from the options menu.

    5.   In the Automatic capacity scaling page, select the Enable automatic capacity scaling for
         your cluster checkbox.

    6.   Set the cluster expansion limit from the dropdown. This is the maximum the cluster can expand
         in the cloud. Automatic scaling is suspended if this limit is exceeded.

    7.   Click Save changes.

6.4. SCALING OUT STORAGE CAPACITY ON A VMWARE CLUSTER

6.4.1. Adding a node to an installer-provisioned infrastructure

Use this procedure to add a new node to an installer-provisioned infrastructure cluster.

Prerequisites

         You have administrative privilege to the OpenShift Container Platform Console.

         You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.       Navigate to Compute → Machine Sets.

    2.   On the machine set where you want to add nodes, select Edit Machine Count.

         a. Add the amount of nodes, and click Save.



             41

Red Hat OpenShift Data Foundation 4.22 Scaling storage

    b. Click Compute → Nodes and confirm if the new node is in Ready state.

3.  Apply the OpenShift Data Foundation label to the new node.

       a.  For the new node, click Action menu (⋮) → Edit Labels.

       b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

                     NOTE

                     It is recommended to add 3 nodes, one each in different zones. You must add
                     3 nodes and perform this procedure for all of them. In case of bare metal
                     installer-provisioned infrastructure deployment, you must expand the cluster
                     first. For instructions, see Expanding the cluster.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

           openshift-storage.cephfs.csi.ceph.com-*

           openshift-storage.rbd.csi.ceph.com-*

6.4.2. Adding a node to an user-provisioned infrastructure

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Depending on the type of infrastructure, perform the following steps:

       a.  Get a new machine with the required infrastructure. See Platform requirements.

       b.  Create a new OpenShift Container Platform worker node using the new machine.

2.   Check for certificate signing requests (CSRs) that are in Pending state.

         $ oc get csr

3.   Approve all the required CSRs for the new node.

         $ oc adm certificate approve <Certificate_Name>

     <Certificate_Name>


42

              CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER

        Is the name of the CSR.

4. Click Compute → Nodes, confirm if the new node is in Ready state.

    5.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

        From User interface

         a.  For the new node, click Action Menu (⋮) → Edit Labels.

         b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

        From Command line interface

             Apply the OpenShift Data Foundation label to the new node.

              $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

             <new_node_name>
              Is the name of the new node.


    Verification

    1.   Execute the following command in the terminal and verify that the new node is present in the
         output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

    2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
         pods on the new node are in Running state:

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

    6.4.3. Adding a node using a local storage device

    You can add nodes to increase the storage capacity when existing worker nodes are already running at
    their maximum supported OSDs or when there are not enough resources to add new OSDs on the
    existing nodes.

    Add nodes in the multiple of 3, each of them in different failure domains. Though it is recommended to
    add nodes in multiples of 3 nodes, you have the flexibility to add one node at a time in flexible scaling
    deployment. See Knowledgebase article Verify if flexible scaling is enabled

                 NOTE

                 OpenShift Data Foundation does not support heterogeneous disk size and types. The
                 new nodes to be added should have the disk of the same type and size which was used
                 during initial OpenShift Data Foundation deployment.

    Prerequisites

        You have administrative privilege to the OpenShift Container Platform Console.

                     43

    Red Hat OpenShift Data Foundation 4.22 Scaling storage

       You have a running OpenShift Data Foundation Storage Cluster.

Procedure

  1.   Depending on the type of infrastructure, perform the following steps:

       a.  Get a new machine with the required infrastructure. See Platform requirements.

       b.  Create a new OpenShift Container Platform worker node using the new machine.

  2.   Check for certificate signing requests (CSRs) that are in Pending state.

          $ oc get csr

  3.   Approve all the required CSRs for the new node.

          $ oc adm certificate approve <Certificate_Name>

       <Certificate_Name>
          Is the name of the CSR.

       4. Click Compute → Nodes, confirm if the new node is in Ready state.

  5.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

       From User interface

           a.  For the new node, click Action Menu (⋮) → Edit Labels.

           b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

       From Command line interface

               Apply the OpenShift Data Foundation label to the new node.

               $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

               <new_node_name>
               Is the name of the new node.

  6.   Click Ecosystem → Installed Operators from the OpenShift Web Console.
       From the Project drop-down list, make sure to select the project where the Local Storage
       Operator is installed.

  7.   Click Local Storage.

  8.   Click the Local Volume Discovery tab.

       a.  Beside the LocalVolumeDiscovery, click Action menu (⋮) → Edit Local Volume
           Discovery.

       b.  In the YAML, add the hostname of the new node in the values field under the node
           selector.



    44

       CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER

   c.  Click Save.

9. Click the Local Volume Sets tab.

   a.  Beside the LocalVolumeSet, click Action menu (⋮) → Edit Local Volume Set.

   b.  In the YAML, add the hostname of the new node in the values field under the node
       selector.










         YAML showing the addition of new hostnames

     c.  Click Save.

                    NOTE

                    It is recommended to add 3 nodes, one each in different zones. You must add
                    3 nodes and perform this procedure for all of them.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

        $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

         openshift-storage.cephfs.csi.ceph.com-*

         openshift-storage.rbd.csi.ceph.com-*

6.4.4. Scaling up storage capacity

To scale up storage capacity:


                    45

Red Hat OpenShift Data Foundation 4.22 Scaling storage

        For dynamic storage devices, see Scaling up storage capacity on a cluster .

        For local storage devices, see Scaling up a cluster created using local storage devices

  6.5. SCALING NON-RESILIENT POOLS

  After enabling the single replica feature, you can increase the number of non-resilient pools by updating
  the pool count in the StorageCluster custom resource. Each additional pool is created in a separate
  failure domain and requires one new storage device per failure domain.


     WARNING
    Non-resilient pools have no data replication. Increasing the pool count does not
    provide data redundancy. Ensure that your applications manage their own
    replication before scaling.




    Prerequisites

         The single replica feature is already enabled. For more information, see Storage class with single
         replica.

         For LSO-based deployments, ensure that the required number of unused local storage devices
         are correctly labeled on the nodes before scaling.

    Procedure

    1.   Run the following command to increase the non-resilient pool count:

             $ oc patch storagecluster ocs-storagecluster \
             -n openshift-storage \
             --type=json \
             --patch='[
               {
                 "op": "replace",
                 "path": "/spec/managedResources/cephNonResilientPools/count",
                 "value": 2
               }
             ]'

         Replace 2 with the desired number of non-resilient pools. The value must not exceed the
         number of available failure domains.

                       NOTE

                       For LSO-based deployments, the operator creates one OSD per local device per
                       failure domain. Ensure that the unused local devices are available on the nodes
                       before running this command.

    2.   Verify that the StorageCluster is in Ready state:


    46

             CHAPTER 6. SCALING STORAGE OF VMWARE OPENSHIFT DATA FOUNDATION CLUSTER

         $ oc get storagecluster -n openshift-storage

        Example output:

         NAME     AGE     PHASE EXTERNAL              CREATED AT VERSION
         ocs-storagecluster 10m Ready     2024-02-05T13:56:15Z   4.22.0

3. Verify that the new CephBlockPools are created for each failure domain and are in Ready
   state:

         $ oc get cephblockpools -n openshift-storage

        Example output:

         NAME        PHASE
         ocs-storagecluster-cephblockpool  Ready
         ocs-storagecluster-cephblockpool-us-east-1a  Ready
         ocs-storagecluster-cephblockpool-us-east-1b  Ready
         ocs-storagecluster-cephblockpool-us-east-1c  Ready

    4.  Verify that the new OSD pods are in Running state:

         $ oc get pods -n openshift-storage | grep osd










                                                      47

Red Hat OpenShift Data Foundation 4.22 Scaling storage

             CHAPTER 7. SCALING STORAGE OF MICROSOFT AZURE
                   OPENSHIFT DATA FOUNDATION CLUSTER

To scale the storage capacity of your configured Red Hat OpenShift Data Foundation worker nodes on
Microsoft Azure cluster, you can increase the capacity by adding three disks at a time. Three disks are
needed since OpenShift Data Foundation uses a replica count of 3 to maintain the high availability. So
the amount of storage consumed is three times the usable space.

             NOTE

             Usable space might vary when encryption is enabled or replica 2 pools are being used.

7.1. SCALING UP STORAGE CAPACITY ON A CLUSTER

To increase the storage capacity in a dynamically created storage cluster on an user-provisioned
infrastructure, you can add storage capacity and performance to your configured Red Hat OpenShift
Data Foundation worker nodes.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

     The disk should be of the same size and type as used during initial deployment.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Storage → Storage cluster.

      a.  Click Actions at the top right of the page to extend the options menu.

      b.  Select Add Capacity from the options menu.

      c.  Select the Storage Class. Choose the storage class which you wish to use to provision new
          storage devices.

      d.  Click Add.

3.   To check the status, navigate to Storage → Storage cluster and verify that the Storage
     Cluster in the Status card has a green tick.

Verification

     Verify the Raw Capacity card.

      a.  In the OpenShift Web Console, click Storage → Storage cluster.

      b.  In the Block and File tab, check the Raw Capacity card.
          Note that the capacity increases based on your selections.



             NOTE
48

    CHAPTER 7. SCALING STORAGE OF MICROSOFT AZURE OPENSHIFT DATA FOUNDATION CLUSTER

                  NOTE

                  The raw capacity does not take replication into account and shows the full
                  capacity.

    Verify that the new object storage devices (OSDs) and their corresponding new Persistent
    Volume Claims (PVCs) are created.

        To view the state of the newly created OSDs:

         a.  Click Workloads → Pods from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

        To view the state of the PVCs:

         a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.

    a.  Identify the nodes where the new OSD pods are running.

            $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
            pod-name>

        <OSD-pod-name>
            Is the name of the OSD pod.
            For example:

              $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
              ceph-osd-0-544db49d7f-qrgqm

            Example output:

              NODE
              compute-1


b. For each of the nodes identified in the previous step, do the following:

     i. Create a debug pod and open a chroot environment for the selected hosts.

        49

Red Hat OpenShift Data Foundation 4.22 Scaling storage

$ oc debug node/<node-name>

<node-name>
Is the name of the node.

$ chroot /host


           ii. Check for the crypt keyword beside the ocs-deviceset names.

                 $ lsblk

             IMPORTANT

             Cluster reduction is supported only with the Red Hat Support Team’s assistance.

7.1.1. Scaling up storage capacity on a cluster by resizing existing OSDs

To increase the storage capacity on a cluster, you can add storage capacity by resizing existing OSDs.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Update the dataPVCTemplate size for the storageDeviceSets in the StorageCluster CR with
     the new desired size. For example:

         apiVersion: ocs.openshift.io/v1
         kind: StorageCluster
         metadata:
          namespace: openshift-storage
          name: example-storagecluster
         spec:
          manageNodes: false
          monPVCTemplate:
           spec:
               storageClassName:
               accessModes:
               - ReadWriteOnce
               resources:
               requests:
               storage: 10Gi
          storageDeviceSets:
          - name: example-deviceset
           count: 3
           resources: {}
           placement: {}
           dataPVCTemplate:
               spec:


50

     CHAPTER 7. SCALING STORAGE OF MICROSOFT AZURE OPENSHIFT DATA FOUNDATION CLUSTER

              storageClassName:
              accessModes:
              - ReadWriteOnce
              volumeMode: Block
              resources:
              requests:
              storage: 1Ti
         portable: true

     In this example YAML, the storage parameter under storageDeviceSets reflects the size
     change of 1Ti.

 2.  Wait for the OSDs to restart.

 3.  Confirm that the resize took effect:

      $ oc get pvc -l ceph.rook.io/DeviceSet

     Example output:

      % oc get pvc -l ceph.rook.io/DeviceSet
      NAME        STATUS     VOLUME        CAPACITY     ACCESS MODES
      STORAGECLASS        AGE
    rook-ceph-osd-a4681 Bound pvc-c3868d38-6bd9-4c8d-91e1-a4681e4cb8c1 1Ti
      RWO         gp2-csi 176m

     Verify the resize is complete in the CAPACITY column.

4.   If the resize did not take effect, restart the OSD pods again. It may take multiple restarts for the
     resize to complete.

7.2. ENABLING AUTOMATIC CAPACITY SCALING

You can enable automatic capacity scaling on clusters deployed using dynamic storage devices. When
automatic capacity scaling is enabled, additional raw capacity equivalent to the configured deployment
size is automatically added to the cluster when used capacity reaches 70%. This ensures your
deployment scales seamlessly to meet demand.

This option is disabled in lean profile mode, LSO deployment, and external mode deployment.

             IMPORTANT

             This may incur additional costs for the underlying storage.

Prerequisites

     Administrative privilege to the OpenShift Container Platform Console.

     A running OpenShift Data Foundation Storage Cluster.

             IMPORTANT

             IBM cloud 10iops-tier workloads are not supported.



              51

Red Hat OpenShift Data Foundation 4.22 Scaling storage

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Storage → Storage cluster.

3.   Click Actions at the top right of the page to extend the options menu.

4.   Select Automatic capacity scaling from the options menu.

5.   In the Automatic capacity scaling page, select the Enable automatic capacity scaling for
     your cluster checkbox.

6.   Set the cluster expansion limit from the dropdown. This is the maximum the cluster can expand
     in the cloud. Automatic scaling is suspended if this limit is exceeded.

7.   Click Save changes.

7.3. SCALING OUT STORAGE CAPACITY ON A MICROSOFT AZURE
CLUSTER

OpenShift Data Foundation is highly scalable. It can be scaled out by adding new nodes with required
storage and enough hardware resources in terms of CPU and RAM. Practically there is no limit on the
number of nodes which can be added but from the support perspective 2000 nodes is the limit for
OpenShift Data Foundation.

Scaling out storage capacity can be broken down into two steps

     Adding new node

     Scaling up the storage capacity

             NOTE

             OpenShift Data Foundation does not support heterogeneous OSD/Disk sizes.

7.3.1. Adding a node to an installer-provisioned infrastructure

Use this procedure to add a new node to an installer-provisioned infrastructure cluster.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Navigate to Compute → Machine Sets.

2.   On the machine set where you want to add nodes, select Edit Machine Count.

       a.  Add the amount of nodes, and click Save.

    b. Click Compute → Nodes and confirm if the new node is in Ready state.


52

          CHAPTER 7. SCALING STORAGE OF MICROSOFT AZURE OPENSHIFT DATA FOUNDATION CLUSTER

3. Apply the OpenShift Data Foundation label to the new node.

           a.  For the new node, click Action menu (⋮) → Edit Labels.

           b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

                 NOTE

                 It is recommended to add 3 nodes, one each in different zones. You must add
                 3 nodes and perform this procedure for all of them. In case of bare metal
                 installer-provisioned infrastructure deployment, you must expand the cluster
                 first. For instructions, see Expanding the cluster.

Verification

     1.   Execute the following command in the terminal and verify that the new node is present in the
          output:

              $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

     2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
          pods on the new node are in Running state:

               openshift-storage.cephfs.csi.ceph.com-*

               openshift-storage.rbd.csi.ceph.com-*

7.3.2. Scaling up storage capacity

To scale up storage capacity, see Scaling up storage capacity on a cluster .

7.4. SCALING NON-RESILIENT POOLS

After enabling the single replica feature, you can increase the number of non-resilient pools by updating
the pool count in the StorageCluster custom resource. Each additional pool is created in a separate
failure domain and requires one new storage device per failure domain.


 WARNING
Non-resilient pools have no data replication. Increasing the pool count does not
provide data redundancy. Ensure that your applications manage their own
replication before scaling.




Prerequisites

          The single replica feature is already enabled. For more information, see Storage class with single
          replica.

          For LSO-based deployments, ensure that the required number of unused local storage devices
          are correctly labeled on the nodes before scaling.

                     53

Red Hat OpenShift Data Foundation 4.22 Scaling storage

Procedure

    1.  Run the following command to increase the non-resilient pool count:

         $ oc patch storagecluster ocs-storagecluster \
          -n openshift-storage \
          --type=json \
          --patch='[
            {
             "op": "replace",
             "path": "/spec/managedResources/cephNonResilientPools/count",
             "value": 2
            }
          ]'

        Replace 2 with the desired number of non-resilient pools. The value must not exceed the
        number of available failure domains.

                    NOTE

                    For LSO-based deployments, the operator creates one OSD per local device per
                    failure domain. Ensure that the unused local devices are available on the nodes
                    before running this command.

2.      Verify that the StorageCluster is in Ready state:

         $ oc get storagecluster -n openshift-storage

        Example output:

         NAME          AGE     PHASE     EXTERNAL        CREATED AT     VERSION
         ocs-storagecluster  10m Ready                 2024-02-05T13:56:15Z 4.22.0

   3. Verify that the new CephBlockPools are created for each failure domain and are in Ready
    state:

         $ oc get cephblockpools -n openshift-storage

        Example output:

         NAME           PHASE
         ocs-storagecluster-cephblockpool      Ready
         ocs-storagecluster-cephblockpool-us-east-1a     Ready
         ocs-storagecluster-cephblockpool-us-east-1b     Ready
         ocs-storagecluster-cephblockpool-us-east-1c     Ready

    4.  Verify that the new OSD pods are in Running state:

         $ oc get pods -n openshift-storage | grep osd






54

         CHAPTER 8. SCALING STORAGE CAPACITY OF GCP OPENSHIFT DATA FOUNDATION CLUSTER

     CHAPTER 8. SCALING STORAGE CAPACITY OF GCP
             OPENSHIFT DATA FOUNDATION CLUSTER

To scale the storage capacity of your configured Red Hat OpenShift Data Foundation worker nodes on
GCP cluster, you can increase the capacity by adding three disks at a time. Three disks are needed since
OpenShift Data Foundation uses a replica count of 3 to maintain the high availability. So the amount of
storage consumed is three times the usable space.

             NOTE

             Usable space might vary when encryption is enabled or replica 2 pools are being used.

8.1. SCALING UP STORAGE CAPACITY ON A CLUSTER

To increase the storage capacity in a dynamically created storage cluster on an user-provisioned
infrastructure, you can add storage capacity and performance to your configured Red Hat OpenShift
Data Foundation worker nodes.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

     The disk should be of the same size and type as used during initial deployment.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Storage → Storage cluster.

      a.  Click Actions at the top right of the page to extend the options menu.

      b.  Select Add Capacity from the options menu.

      c.  Select the Storage Class. Choose the storage class which you wish to use to provision new
          storage devices.

      d.  Click Add.

3.   To check the status, navigate to Storage → Storage cluster and verify that the Storage
     Cluster in the Status card has a green tick.

Verification

     Verify the Raw Capacity card.

      a.  In the OpenShift Web Console, click Storage → Storage cluster.

      b.  In the Block and File tab, check the Raw Capacity card.
          Note that the capacity increases based on your selections.



             NOTE
                 55

    Red Hat OpenShift Data Foundation 4.22 Scaling storage

                  NOTE

                  The raw capacity does not take replication into account and shows the full
                  capacity.

    Verify that the new object storage devices (OSDs) and their corresponding new Persistent
    Volume Claims (PVCs) are created.

        To view the state of the newly created OSDs:

         a.  Click Workloads → Pods from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

        To view the state of the PVCs:

         a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.

    a.  Identify the nodes where the new OSD pods are running.

            $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
            pod-name>

        <OSD-pod-name>
            Is the name of the OSD pod.
            For example:

              $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
              ceph-osd-0-544db49d7f-qrgqm

            Example output:

              NODE
              compute-1


        b. For each of the nodes identified in the previous step, do the following:

        i. Create a debug pod and open a chroot environment for the selected hosts.

56

CHAPTER 8. SCALING STORAGE CAPACITY OF GCP OPENSHIFT DATA FOUNDATION CLUSTER

$ oc debug node/<node-name>

<node-name>
Is the name of the node.

$ chroot /host


           ii. Check for the crypt keyword beside the ocs-deviceset names.

                 $ lsblk

             IMPORTANT

             Cluster reduction is supported only with the Red Hat Support Team’s assistance.

8.1.1. Scaling up storage capacity on a cluster by resizing existing OSDs

To increase the storage capacity on a cluster, you can add storage capacity by resizing existing OSDs.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Update the dataPVCTemplate size for the storageDeviceSets in the StorageCluster CR with
     the new desired size. For example:

         apiVersion: ocs.openshift.io/v1
         kind: StorageCluster
         metadata:
          namespace: openshift-storage
          name: example-storagecluster
         spec:
          manageNodes: false
          monPVCTemplate:
           spec:
               storageClassName:
               accessModes:
               - ReadWriteOnce
               resources:
               requests:
               storage: 10Gi
          storageDeviceSets:
          - name: example-deviceset
           count: 3
           resources: {}
           placement: {}
           dataPVCTemplate:
               spec:


               57

Red Hat OpenShift Data Foundation 4.22 Scaling storage

              storageClassName:
              accessModes:
              - ReadWriteOnce
              volumeMode: Block
              resources:
              requests:
              storage: 1Ti
         portable: true

     In this example YAML, the storage parameter under storageDeviceSets reflects the size
     change of 1Ti.

 2.  Wait for the OSDs to restart.

 3.  Confirm that the resize took effect:

      $ oc get pvc -l ceph.rook.io/DeviceSet

     Example output:

      % oc get pvc -l ceph.rook.io/DeviceSet
      NAME        STATUS     VOLUME        CAPACITY     ACCESS MODES
      STORAGECLASS        AGE
    rook-ceph-osd-a4681 Bound pvc-c3868d38-6bd9-4c8d-91e1-a4681e4cb8c1 1Ti
      RWO         gp2-csi 176m

     Verify the resize is complete in the CAPACITY column.

4.   If the resize did not take effect, restart the OSD pods again. It may take multiple restarts for the
     resize to complete.

8.2. ENABLING AUTOMATIC CAPACITY SCALING

You can enable automatic capacity scaling on clusters deployed using dynamic storage devices. When
automatic capacity scaling is enabled, additional raw capacity equivalent to the configured deployment
size is automatically added to the cluster when used capacity reaches 70%. This ensures your
deployment scales seamlessly to meet demand.

This option is disabled in lean profile mode, LSO deployment, and external mode deployment.

             IMPORTANT

             This may incur additional costs for the underlying storage.

Prerequisites

     Administrative privilege to the OpenShift Container Platform Console.

     A running OpenShift Data Foundation Storage Cluster.

             IMPORTANT

             IBM cloud 10iops-tier workloads are not supported.



58

         CHAPTER 8. SCALING STORAGE CAPACITY OF GCP OPENSHIFT DATA FOUNDATION CLUSTER

Procedure

    1.   Log in to the OpenShift Web Console.

    2.   Click Storage → Storage cluster.

    3.   Click Actions at the top right of the page to extend the options menu.

    4.   Select Automatic capacity scaling from the options menu.

    5.   In the Automatic capacity scaling page, select the Enable automatic capacity scaling for
         your cluster checkbox.

    6.   Set the cluster expansion limit from the dropdown. This is the maximum the cluster can expand
         in the cloud. Automatic scaling is suspended if this limit is exceeded.

    7.   Click Save changes.

8.3. SCALING OUT STORAGE CAPACITY ON A GCP CLUSTER

OpenShift Data Foundation is highly scalable. It can be scaled out by adding new nodes with required
storage and enough hardware resources in terms of CPU and RAM. Practically there is no limit on the
number of nodes which can be added but from the support perspective 2000 nodes is the limit for
OpenShift Data Foundation.

Scaling out storage capacity can be broken down into two steps

         Adding new node

         Scaling up the storage capacity

          NOTE

          OpenShift Data Foundation does not support heterogeneous OSD/Disk sizes.

8.3.1. Adding a node

You can add nodes to increase the storage capacity when existing worker nodes are already running at
their maximum supported OSDs or there are not enough resources to add new OSDs on the existing
nodes. It is always recommended to add nodes in the multiple of three, each of them in different failure
domains.

While it is recommended to add nodes in the multiple of three, you still have the flexibility to add one
node at a time in the flexible scaling deployment. Refer to the Knowledgebase article Verify if flexible
scaling is enabled.

          NOTE

          OpenShift Data Foundation does not support heterogeneous disk size and types. The
          new nodes to be added should have the disk of the same type and size which was used
          during OpenShift Data Foundation deployment.

8.3.1.1. Adding a node to an installer-provisioned infrastructure

Use this procedure to add a new node to an installer-provisioned infrastructure cluster.


          59

Red Hat OpenShift Data Foundation 4.22 Scaling storage

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Navigate to Compute → Machine Sets.

2.   On the machine set where you want to add nodes, select Edit Machine Count.

      a.  Add the amount of nodes, and click Save.

    b. Click Compute → Nodes and confirm if the new node is in Ready state.

3.   Apply the OpenShift Data Foundation label to the new node.

      a.  For the new node, click Action menu (⋮) → Edit Labels.

      b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

             NOTE

             It is recommended to add 3 nodes, one each in different zones. You must add
             3 nodes and perform this procedure for all of them. In case of bare metal
             installer-provisioned infrastructure deployment, you must expand the cluster
             first. For instructions, see Expanding the cluster.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

8.3.2. Scaling up storage capacity

To scale up storage capacity, see Scaling up storage capacity on a cluster .

8.4. SCALING NON-RESILIENT POOLS

After enabling the single replica feature, you can increase the number of non-resilient pools by updating
the pool count in the StorageCluster custom resource. Each additional pool is created in a separate
failure domain and requires one new storage device per failure domain.





60

    CHAPTER 8. SCALING STORAGE CAPACITY OF GCP OPENSHIFT DATA FOUNDATION CLUSTER


     WARNING
    Non-resilient pools have no data replication. Increasing the pool count does not
    provide data redundancy. Ensure that your applications manage their own
    replication before scaling.




    Prerequisites

         The single replica feature is already enabled. For more information, see Storage class with single
         replica.

         For LSO-based deployments, ensure that the required number of unused local storage devices
         are correctly labeled on the nodes before scaling.

Procedure

    1.   Run the following command to increase the non-resilient pool count:

          $ oc patch storagecluster ocs-storagecluster \
           -n openshift-storage \
           --type=json \
           --patch='[
             {
                 "op": "replace",
                 "path": "/spec/managedResources/cephNonResilientPools/count",
                 "value": 2
             }
           ]'

         Replace 2 with the desired number of non-resilient pools. The value must not exceed the
         number of available failure domains.

                     NOTE

                     For LSO-based deployments, the operator creates one OSD per local device per
                     failure domain. Ensure that the unused local devices are available on the nodes
                     before running this command.

    2.   Verify that the StorageCluster is in Ready state:

          $ oc get storagecluster -n openshift-storage

         Example output:

          NAME           AGE     PHASE          EXTERNAL   CREATED AT     VERSION
          ocs-storagecluster     10m Ready     2024-02-05T13:56:15Z     4.22.0

     3. Verify that the new CephBlockPools are created for each failure domain and are in Ready
        state:


                     61

Red Hat OpenShift Data Foundation 4.22 Scaling storage

     $ oc get cephblockpools -n openshift-storage

    Example output:

     NAME        PHASE
     ocs-storagecluster-cephblockpool      Ready
     ocs-storagecluster-cephblockpool-us-east-1a      Ready
     ocs-storagecluster-cephblockpool-us-east-1b      Ready
     ocs-storagecluster-cephblockpool-us-east-1c      Ready

4.  Verify that the new OSD pods are in Running state:

     $ oc get pods -n openshift-storage | grep osd










62

CHAPTER 9. SCALING STORAGE OF IBM Z OR IBM LINUXONE OPENSHIFT DATA FOUNDATION CLUSTER

CHAPTER 9. SCALING STORAGE OF IBM Z OR IBM LINUXONE
              OPENSHIFT DATA FOUNDATION CLUSTER

9.1. SCALING UP STORAGE BY ADDING CAPACITY TO YOUR
OPENSHIFT DATA FOUNDATION NODES ON IBM Z OR IBM LINUXONE
INFRASTRUCTURE

You can add storage capacity and performance to your configured Red Hat OpenShift Data Foundation
worker nodes.

             NOTE

             Flexible scaling features get enabled at the time of deployment and can not be enabled
             or disabled later on.

Prerequisites

     A running OpenShift Data Foundation Platform.

     Administrative privileges on the OpenShift Web Console.

     To scale using a storage class other than the one provisioned during deployment, first define an
     additional storage class. See Creating storage classes and pools for details.

Procedure

1.   Add additional hardware resources with zFCP disks.

     a.  List all the disks.

              $ lszdev

         Example output:

              TYPE        ID        ON     PERS NAMES
              zfcp-host   0.0.8204        yes yes
    zfcp-lun 0.0.8204:0x102107630b1b5060:0x4001402900000000 yes no sda sg0
    zfcp-lun 0.0.8204:0x500407630c0b50a4:0x3002b03000000000 yes yes sdb sg1
              qeth     0.0.bdd0:0.0.bdd1:0.0.bdd2        yes no encbdd0
              generic-ccw 0.0.0009        yes no

         A SCSI disk is represented as a zfcp-lun with the structure <device-id>:<wwpn>:<lun-id>
         in the ID section. The first disk is used for the operating system. The device id for the new
         disk can be the same.

     b.  Append a new SCSI disk.

              $ chzdev -e 0.0.8204:0x400506630b1b50a4:0x3001301a00000000





                         NOTE
                                                                                  63

Red Hat OpenShift Data Foundation 4.22 Scaling storage

                       NOTE

                       The device ID for the new disk must be the same as the disk to be replaced.
                       The new disk is identified with its WWPN and LUN ID.

    c.  List all the FCP devices to verify the new disk is configured.

               $ lszdev zfcp-lun
               TYPE    ID        ON     PERS NAMES
    zfcp-lun 0.0.8204:0x102107630b1b5060:0x4001402900000000 yes no sda sg0
    zfcp-lun 0.0.8204:0x500507630b1b50a4:0x4001302a00000000 yes yes sdb sg1
    zfcp-lun 0.0.8204:0x400506630b1b50a4:0x3001301a00000000 yes yes sdc sg2

2.  Navigate to the OpenShift Web Console.

3.  Click Storage → Storage cluster.

a.  Click Actions at the top right of the page to extend the options menu.

b.  Select Add Capacity from the options menu.
    The Raw Capacity field shows the size set during storage class creation. The total amount of
    storage consumed is three times this amount, because OpenShift Data Foundation uses a
    replica count of 3.

c.  Click Add.

    1.   To check the status, navigate to Storage → Storage cluster and verify that Storage
         Cluster in the Status card has a green tick.

Verification

    Verify the Raw Capacity card.

    a.   In the OpenShift Web Console, click Storage → Storage cluster.

    b.   In the Block and File tab, check the Raw Capacity card.
         Note that the capacity increases based on your selections.

                       NOTE

                       The raw capacity does not take replication into account and shows the full
                       capacity.

    Verify that the new OSDs and their corresponding new Persistent Volume Claims (PVCs) are
    created.

         To view the state of the newly created OSDs:

            a.  Click Workloads → Pods from the OpenShift Web Console.

            b.  Select openshift-storage from the Project drop-down list.





                       NOTE
64

CHAPTER 9. SCALING STORAGE OF IBM Z OR IBM LINUXONE OPENSHIFT DATA FOUNDATION CLUSTER

                    NOTE

                    If the Show default projects option is disabled, use the toggle button to
                    list all the default projects.

To view the state of the PVCs:

     a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

     b.  Select openshift-storage from the Project drop-down list.

                    NOTE

                    If the Show default projects option is disabled, use the toggle button to
                    list all the default projects.

Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
are encrypted.

a.  Identify the nodes where the new OSD pods are running.

        $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
        pod-name>

    <OSD-pod-name>
        Is the name of the OSD pod.
        For example:

          $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
          ceph-osd-0-544db49d7f-qrgqm

        Example output:

          NODE
          compute-1


b. For each of the nodes identified in the previous step, do the following:

     i.   Create a debug pod and open a chroot environment for the selected host(s).

          $ oc debug node/<node-name>

          <node-name>
          Is the name of the node.

          $ chroot /host


     ii. Check for the crypt keyword beside the ocs-deviceset names.

          $ lsblk



                    65

Red Hat OpenShift Data Foundation 4.22 Scaling storage

             IMPORTANT

             Cluster reduction is supported only with the Red Hat Support Team’s assistance.

9.2. SCALING OUT STORAGE CAPACITY ON A IBM Z OR IBM
LINUXONE CLUSTER

9.2.1. Adding a node using a local storage device

You can add nodes to increase the storage capacity when existing worker nodes are already running at
their maximum supported OSDs or when there are not enough resources to add new OSDs on the
existing nodes.

Add nodes in the multiple of 3, each of them in different failure domains. Though it is recommended to
add nodes in multiples of 3 nodes, you have the flexibility to add one node at a time in flexible scaling
deployment. See Knowledgebase article Verify if flexible scaling is enabled

             NOTE

             OpenShift Data Foundation does not support heterogeneous disk size and types. The
             new nodes to be added should have the disk of the same type and size which was used
             during initial OpenShift Data Foundation deployment.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Depending on the type of infrastructure, perform the following steps:

      a.  Get a new machine with the required infrastructure. See Platform requirements.

      b.  Create a new OpenShift Container Platform worker node using the new machine.

2.   Check for certificate signing requests (CSRs) that are in Pending state.

         $ oc get csr

3.   Approve all the required CSRs for the new node.

         $ oc adm certificate approve <Certificate_Name>

     <Certificate_Name>
         Is the name of the CSR.

   4. Click Compute → Nodes, confirm if the new node is in Ready state.

5.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

     From User interface


66

CHAPTER 9. SCALING STORAGE OF IBM Z OR IBM LINUXONE OPENSHIFT DATA FOUNDATION CLUSTER

         a.  For the new node, click Action Menu (⋮) → Edit Labels.

         b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

    From Command line interface

             Apply the OpenShift Data Foundation label to the new node.

             $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

             <new_node_name>
             Is the name of the new node.

6.  Click Ecosystem → Installed Operators from the OpenShift Web Console.
    From the Project drop-down list, make sure to select the project where the Local Storage
    Operator is installed.

7.  Click Local Storage.

8.  Click the Local Volume Discovery tab.

     a.  Beside the LocalVolumeDiscovery, click Action menu (⋮) → Edit Local Volume
         Discovery.

     b.  In the YAML, add the hostname of the new node in the values field under the node
         selector.

     c.  Click Save.

9.  Click the Local Volume Sets tab.

     a.  Beside the LocalVolumeSet, click Action menu (⋮) → Edit Local Volume Set.

     b.  In the YAML, add the hostname of the new node in the values field under the node
         selector.










             67

Red Hat OpenShift Data Foundation 4.22 Scaling storage

         YAML showing the addition of new hostnames

     c.  Click Save.

                    NOTE

                    It is recommended to add 3 nodes, one each in different zones. You must add
                    3 nodes and perform this procedure for all of them.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

        $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

         openshift-storage.cephfs.csi.ceph.com-*

         openshift-storage.rbd.csi.ceph.com-*

9.2.2. Scaling up storage capacity

To scale up storage capacity, see Scaling up storage capacity on a cluster .










68

             CHAPTER 10. SCALING STORAGE OF IBM POWER OPENSHIFT DATA FOUNDATION CLUSTER

CHAPTER 10. SCALING STORAGE OF IBM POWER OPENSHIFT
                 DATA FOUNDATION CLUSTER

To scale the storage capacity of your configured Red Hat OpenShift Data Foundation worker nodes on
IBM Power cluster, you can increase the capacity by adding three disks at a time. Three disks are
needed since OpenShift Data Foundation uses a replica count of 3 to maintain the high availability. So
the amount of storage consumed is three times the usable space.

             NOTE

             Usable space might vary when encryption is enabled or replica 2 pools are being used.

10.1. SCALING UP STORAGE BY ADDING CAPACITY TO YOUR
OPENSHIFT DATA FOUNDATION NODES ON IBM POWER
INFRASTRUCTURE USING LOCAL STORAGE DEVICES

In order to scale up an OpenShift Data Foundation cluster which was created using local storage devices,
a new disk needs to be added to the storage node. It is recommended to have the new disks of the same
size as used earlier during the deployment as OpenShift Data Foundation does not support
heterogeneous disks/OSD’s.

You can add storage capacity (additional storage devices) to your configured local storage based
OpenShift Data Foundation worker nodes on IBM Power infrastructures.

             NOTE

             Flexible scaling features get enabled at the time of deployment and can not be enabled
             or disabled later on.

Prerequisites

     You must be logged into the OpenShift Container Platform cluster.

     You must have installed the local storage operator. Use the following procedure:

         Installing Local Storage Operator on IBM Power

     You must have three OpenShift Container Platform worker nodes with the same storage type
     and size attached to each node (for example, 0.5TB SSD) as the original OpenShift Data
     Foundation StorageCluster was created with.

Procedure

1.   To add storage capacity to OpenShift Container Platform nodes with OpenShift Data
     Foundation installed, you need to

     a.  Find the available devices that you want to add, that is, a minimum of one device per worker
         node. You can follow the procedure for finding available storage devices in the respective
         deployment guide.





             NOTE
                 69

Red Hat OpenShift Data Foundation 4.22 Scaling storage

                     NOTE

                     Make sure you perform this process for all the existing nodes (minimum of 3)
                     for which you want to add storage.

      b.  Add the additional disks to the LocalVolume custom resource (CR).

             $ oc edit -n openshift-local-storage localvolume localblock

          Example output:

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
             - /dev/sdx     # newly added device
            storageClassName: localblock
            volumeMode: Block

          Make sure to save the changes after editing the CR.

          Example output:

          localvolume.local.storage.openshift.io/localblock edited

          You can see in this CR that new devices are added.

          sdx

  2.  Display the newly created Persistent Volumes (PVs) with the storageclass name used in the
      localVolume CR.

         $ oc get pv | grep localblock | grep Available

      Example output:

         local-pv-a04ffd8   500Gi     RWO           Delete    Available   localblock   24s
         local-pv-a0ca996b              500Gi   RWO   Delete   Available   localblock   23s
         local-pv-c171754a              500Gi   RWO   Delete   Available   localblock   23s

  3.  Navigate to the OpenShift Web Console.

  4.  Click Storage → Storage cluster.

70

            CHAPTER 10. SCALING STORAGE OF IBM POWER OPENSHIFT DATA FOUNDATION CLUSTER

a.  Click Actions at the top right of the page to extend the options menu.

b.  Select Add Capacity from the options menu.
    From this dialog box, set the Storage Class name to the name used in the localVolume CR.
    Available Capacity displayed is based on the local disks available in storage class.

c.  Click Add.

    1.   To check the status, navigate to Storage → Storage cluster and verify that the Storage
         Cluster in the Status card has a green tick.

Verification

    Verify the available Capacity.

         In the OpenShift Web Console, click Storage → Storage cluster.

         Click the Block and File tab, then check the Raw Capacity card.
         Note that the capacity increases based on your selections.

                  NOTE

                  The raw capacity does not take replication into account and shows the full
                  capacity.

    Verify that the new OSDs and their corresponding new Persistent Volume Claims (PVCs) are
    created.

         To view the state of the newly created OSDs:

            a.  Click Workloads → Pods from the OpenShift Web Console.

            b.  Select openshift-storage from the Project drop-down list.

                   NOTE

                   If the Show default projects option is disabled, use the toggle button to
                   list all the default projects.

         To view the state of the PVCs:

            a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

            b.  Select openshift-storage from the Project drop-down list.

                   NOTE

                   If the Show default projects option is disabled, use the toggle button to
                   list all the default projects.

    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.

    a. Identify the nodes where the new OSD pods are running.


                   71

Red Hat OpenShift Data Foundation 4.22 Scaling storage

$ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
pod-name>

<OSD-pod-name>
Is the name of the OSD pod.
For example:

 $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
 ceph-osd-0-544db49d7f-qrgqm

Example output:

 NODE
 compute-1


b. For each of the nodes identified in the previous step, do the following:

   i.   Create a debug pod and open a chroot environment for the selected host(s).

        $ oc debug node/<node-name>

        <node-name>
        Is the name of the node.

        $ chroot /host


  ii. Check for the crypt keyword beside the ocs-deviceset names.

        $ lsblk

   IMPORTANT

   Cluster reduction is supported only with the Red Hat Support Team’s assistance.

10.2. SCALING OUT STORAGE CAPACITY ON A IBM POWER CLUSTER

OpenShift Data Foundation is highly scalable. It can be scaled out by adding new nodes with required
storage and enough hardware resources in terms of CPU and RAM. Practically there is no limit on the
number of nodes which can be added but from the support perspective 2000 nodes is the limit for
OpenShift Data Foundation.

Scaling out storage capacity can be broken down into two steps:

  Adding new node

  Scaling up the storage capacity

   NOTE

   OpenShift Data Foundation does not support heterogeneous OSD/Disk sizes.

72

             CHAPTER 10. SCALING STORAGE OF IBM POWER OPENSHIFT DATA FOUNDATION CLUSTER

10.2.1. Adding a node using a local storage device on IBM Power

You can add nodes to increase the storage capacity when existing worker nodes are already running at
their maximum supported OSDs or when there are not enough resources to add new OSDs on the
existing nodes.

Add nodes in the multiple of 3, each of them in different failure domains. Though it is recommended to
add nodes in multiples of 3 nodes, you have the flexibility to add one node at a time in flexible scaling
deployment. See Knowledgebase article Verify if flexible scaling is enabled

             NOTE

             OpenShift Data Foundation does not support heterogeneous disk size and types. The
             new nodes to be added should have the disk of the same type and size which was used
             during initial OpenShift Data Foundation deployment.

Prerequisites

     You must be logged into the OpenShift Container Platform cluster.

     You must have three OpenShift Container Platform worker nodes with the same storage type
     and size attached to each node (for example, 2TB SSD drive) as the original OpenShift Data
     Foundation StorageCluster was created with.

Procedure

1.   Get a new IBM Power machine with the required infrastructure. See Platform requirements.

2.   Create a new OpenShift Container Platform node using the new IBM Power machine.

      a.  Check for certificate signing requests (CSRs) that are in Pending state.

               $ oc get csr

      b.  Approve all the required CSRs for the new node.

               $ oc adm certificate approve <Certificate_Name>

          <Certificate_Name>
               Is the name of the CSR.

    c. Click Compute → Nodes, confirm if the new node is in Ready state.

      d.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

          From User interface

              i.   For the new node, click Action Menu (⋮) → Edit Labels.

             ii.   Add cluster.ocs.openshift.io/openshift-storage and click Save.

          From Command line interface

                   Apply the OpenShift Data Foundation label to the new node.



                   73

Red Hat OpenShift Data Foundation 4.22 Scaling storage

                  $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

           <new_node_name>
               Is the name of the new node.

      e.   Click Ecosystem → Installed Operators from the OpenShift Web Console.
           From the Project drop-down list, make sure to select the project where the Local Storage
           Operator is installed.

      f.   Click Local Storage.

      g.   Click the Local Volume tab.

             i.   Beside the LocalVolume, click Action menu (⋮) → Edit Local Volume.

            ii.   In the YAML, add the hostname of the new node in the values field under the node
                  selector.
                  YAML showing the addition of new hostnamesimage::images/ocs_yaml-showing-
                  addition-new-hostnames-localvolume-power.png[Screenshot of YAML showing the
                  addition of new hostnames.]

           iii.   Click Save.

                             NOTE

                             It is recommended to add 3 nodes, one each in different zones. You must
                             add 3 nodes and perform this procedure for all of them.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

          $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

           openshift-storage.cephfs.csi.ceph.com-*

           openshift-storage.rbd.csi.ceph.com-*

10.2.2. Scaling up storage capacity

To scale up storage capacity, see Scaling up storage capacity on a cluster .










74

                 CHAPTER 11. SCALING STORAGE CAPACITY OF IBM FLASHSYSTEM CLUSTER

      CHAPTER 11. SCALING STORAGE CAPACITY OF IBM
                 FLASHSYSTEM CLUSTER

To scale the storage capacity of your configured Red Hat OpenShift Data Foundation worker nodes on
a cluster using external IBM FlashSystem, you can increase the capacity by adding three disks at a time.
Three disks are needed since OpenShift Data Foundation uses a replica count of 3 to maintain the high
availability. So the amount of storage consumed is three times the usable space.

             NOTE

             Usable space might vary when encryption is enabled or replica 2 pools are being used.

11.1. SCALING UP STORAGE CAPACITY ON A CLUSTER

To increase the storage capacity in a dynamically created storage cluster on an user-provisioned
infrastructure, you can add storage capacity and performance to your configured Red Hat OpenShift
Data Foundation worker nodes.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

     The disk should be of the same size and type as used during initial deployment.

Procedure

1.   Log in to the OpenShift Web Console.

2.   Click Storage → Storage cluster.

      a.  Click Actions at the top right of the page to extend the options menu.

      b.  Select Add Capacity from the options menu.

      c.  Select the Storage Class. Choose the storage class which you wish to use to provision new
          storage devices.

      d.  Click Add.

3.   To check the status, navigate to Storage → Storage cluster and verify that the Storage
     Cluster in the Status card has a green tick.

Verification

     Verify the Raw Capacity card.

      a.  In the OpenShift Web Console, click Storage → Storage cluster.

      b.  In the Block and File tab, check the Raw Capacity card.
          Note that the capacity increases based on your selections.



             NOTE
                 75

    Red Hat OpenShift Data Foundation 4.22 Scaling storage

                  NOTE

                  The raw capacity does not take replication into account and shows the full
                  capacity.

    Verify that the new object storage devices (OSDs) and their corresponding new Persistent
    Volume Claims (PVCs) are created.

        To view the state of the newly created OSDs:

         a.  Click Workloads → Pods from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

        To view the state of the PVCs:

         a.  Click Storage → Persistent Volume Claims from the OpenShift Web Console.

         b.  Select openshift-storage from the Project drop-down list.

                        NOTE

                        If the Show default projects option is disabled, use the toggle button to
                        list all the default projects.

    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.

    a.  Identify the nodes where the new OSD pods are running.

            $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/<OSD-
            pod-name>

        <OSD-pod-name>
            Is the name of the OSD pod.
            For example:

              $ oc get -n openshift-storage -o=custom-columns=NODE:.spec.nodeName pod/rook-
              ceph-osd-0-544db49d7f-qrgqm

            Example output:

              NODE
              compute-1


        b. For each of the nodes identified in the previous step, do the following:

        i. Create a debug pod and open a chroot environment for the selected hosts.

76

CHAPTER 11. SCALING STORAGE CAPACITY OF IBM FLASHSYSTEM CLUSTER

$ oc debug node/<node-name>

<node-name>
Is the name of the node.

$ chroot /host


     ii. Check for the crypt keyword beside the ocs-deviceset names.

            $ lsblk

        IMPORTANT

        Cluster reduction is supported only with the Red Hat Support Team’s assistance.

11.2. SCALING OUT STORAGE CAPACITY ON AN IBM FLASHSYSTEM
CLUSTER

OpenShift Data Foundation is highly scalable. It can be scaled out by adding new nodes with required
storage and enough hardware resources in terms of CPU and RAM. Practically there is no limit on the
number of nodes which can be added but from the support perspective 2000 nodes is the limit for
OpenShift Data Foundation.

Scaling out storage capacity can be broken down into two steps

 Adding new node

 Scaling up the storage capacity

        NOTE

        OpenShift Data Foundation does not support heterogeneous OSD/Disk sizes.

11.2.1. Adding a node

You can add nodes to increase the storage capacity when existing worker nodes are already running at
their maximum supported OSDs or there are not enough resources to add new OSDs on the existing
nodes. It is always recommended to add nodes in the multiple of three, each of them in different failure
domains.

While we recommend adding nodes in the multiple of three, you still get the flexibility of adding one
node at a time in the flexible scaling deployment. Refer to the Knowledgebase article Verify if flexible
scaling is enabled.

        NOTE

        OpenShift Data Foundation does not support heterogeneous disk size and types. The
        new nodes to be added should have the disk of the same type and size which was used
        during OpenShift Data Foundation deployment.

11.2.1.1. Adding a node to an installer-provisioned infrastructure

            77

Red Hat OpenShift Data Foundation 4.22 Scaling storage

Use this procedure to add a new node to an installer-provisioned infrastructure cluster.

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure

1.   Navigate to Compute → Machine Sets.

2.   On the machine set where you want to add nodes, select Edit Machine Count.

      a.  Add the amount of nodes, and click Save.

    b. Click Compute → Nodes and confirm if the new node is in Ready state.

3.   Apply the OpenShift Data Foundation label to the new node.

      a.  For the new node, click Action menu (⋮) → Edit Labels.

      b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

             NOTE

             It is recommended to add 3 nodes, one each in different zones. You must add
             3 nodes and perform this procedure for all of them. In case of bare metal
             installer-provisioned infrastructure deployment, you must expand the cluster
             first. For instructions, see Expanding the cluster.

Verification

1.   Execute the following command in the terminal and verify that the new node is present in the
     output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
     pods on the new node are in Running state:

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

11.2.1.2. Adding a node to an user-provisioned infrastructure

Prerequisites

     You have administrative privilege to the OpenShift Container Platform Console.

     You have a running OpenShift Data Foundation Storage Cluster.

Procedure


78

                    CHAPTER 11. SCALING STORAGE CAPACITY OF IBM FLASHSYSTEM CLUSTER

  1.   Depending on the type of infrastructure, perform the following steps:

        a.  Get a new machine with the required infrastructure. See Platform requirements.

        b.  Create a new OpenShift Container Platform worker node using the new machine.

  2.   Check for certificate signing requests (CSRs) that are in Pending state.

           $ oc get csr

  3.   Approve all the required CSRs for the new node.

           $ oc adm certificate approve <Certificate_Name>

       <Certificate_Name>
           Is the name of the CSR.

4. Click Compute → Nodes, confirm if the new node is in Ready state.

  5.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

       From User interface

            a.  For the new node, click Action Menu (⋮) → Edit Labels.

            b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

       From Command line interface

                Apply the OpenShift Data Foundation label to the new node.

                $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

                <new_node_name>
                Is the name of the new node.


    Verification

    1.   Execute the following command in the terminal and verify that the new node is present in the
         output:

           $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

    2.   On the OpenShift web console, click Workloads → Pods, confirm that at least the following
         pods on the new node are in Running state:

            openshift-storage.cephfs.csi.ceph.com-*

            openshift-storage.rbd.csi.ceph.com-*

    11.2.2. Scaling up storage capacity

    To scale up storage capacity, see Scaling up a cluster .


                79

Red Hat OpenShift Data Foundation 4.22 Scaling storage










80
