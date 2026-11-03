---
title: "replacing nodes"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/replacing_nodes"
source_file: "replacing_nodes.pdf"
version: "4.22"
chapter: "replacing_nodes"
---

                      Red Hat OpenShift Data Foundation
                                     4.22

    Replacing nodes

Instructions for how to safely replace a node in an OpenShift Data Foundation
                                   cluster.










    Last Updated: 2026-11-02



Red Hat OpenShift Data Foundation 4.22 Replacing nodes

Instructions for how to safely replace a node in an OpenShift Data Foundation cluster.

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

    This document explains how to safely replace a node in a Red Hat OpenShift Data Foundation
    cluster.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. REPLACING STORAGE NODES OVERVIEW 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES 6
  2.1. REPLACING STORAGE NODES ON AWS 6
   2.1.1. Replacing an operational AWS node on user-provisioned infrastructure 6
   2.1.2. Replacing an operational AWS node on installer-provisioned infrastructure 8
   2.1.3. Replacing a failed AWS node on user-provisioned infrastructure 10
   2.1.4. Replacing a failed AWS node on installer-provisioned infrastructure 11
  2.2. REPLACING STORAGE NODES ON VMWARE 13
   2.2.1. Replacing an operational VMware node on user-provisioned infrastructure 13
   2.2.2. Replacing an operational VMware node on installer-provisioned infrastructure 15
   2.2.3. Replacing a failed VMware node on user-provisioned infrastructure 17
   2.2.4. Replacing a failed VMware node on installer-provisioned infrastructure 19
  2.3. REPLACING STORAGE NODES ON MICROSOFT AZURE 20
   2.3.1. Replacing operational nodes on Azure installer-provisioned infrastructure 20
   2.3.2. Replacing failed nodes on Azure installer-provisioned infrastructure 22
  2.4. REPLACING STORAGE NODES ON GOOGLE CLOUD 24
   2.4.1. Replacing operational nodes on Google Cloud installer-provisioned infrastructure 24
   2.4.2. Replacing failed nodes on Google Cloud installer-provisioned infrastructure 25

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES 28
  3.1. REPLACING STORAGE NODES ON BARE METAL INFRASTRUCTURE 28
   3.1.1. Replacing an operational node using local storage devices 28
   3.1.2. Replacing an operational node on bare metal user-provisioned infrastructure 28
   3.1.3. Replacing an operational node using local storage devices 34
   3.1.4. Replacing an operational node on bare metal user-provisioned infrastructure 34
   3.1.5. Replacing a failed node on bare metal user-provisioned infrastructure 34
  3.2. REPLACING STORAGE NODES ON IBM Z OR IBM® LINUXONE INFRASTRUCTURE 41
   3.2.1. Replacing operational nodes on IBM Z or IBM® LinuxONE infrastructure 41
   3.2.2. Replacing failed nodes on IBM Z or IBM® LinuxONE infrastructure 47
  3.3. REPLACING STORAGE NODES ON IBM POWER INFRASTRUCTURE 48
   3.3.1. Replacing an operational or failed storage node on IBM Power 48
  3.4. REPLACING STORAGE NODES ON VMWARE INFRASTRUCTURE 54
   3.4.1. Replacing an operational node using local storage devices 55
   3.4.2. Replacing an operational node on VMware user-provisioned infrastructure 55
   3.4.3. Replacing an operational node on bare metal user-provisioned infrastructure 55
   3.4.4. Replacing a failed node on bare metal user-provisioned infrastructure 55
   3.4.5. Replacing an operational node on VMware installer-provisioned infrastructure 63
   3.4.6. Replacing an operational node using local storage devices 69
   3.4.7. Replacing an operational node on VMware user-provisioned infrastructure 69
   3.4.8. Replacing an operational node on VMware installer-provisioned infrastructure 69
   3.4.9. Replacing a failed node on VMware user-provisioned infrastructure 69
   3.4.10. Replacing an operational node on bare metal user-provisioned infrastructure 69
   3.4.11. Replacing a failed node on bare metal user-provisioned infrastructure 69
   3.4.12. Replacing a failed node on VMware installer-provisioned infrastructure 80






    1

Red Hat OpenShift Data Foundation 4.22 Replacing nodes










2

    PREFACE

    PREFACE

This document describes how to replace operational and failed nodes in your OpenShift Data
Foundation cluster.










3

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

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

      CHAPTER 1. REPLACING STORAGE NODES OVERVIEW

CHAPTER 1. REPLACING STORAGE NODES OVERVIEW

For OpenShift Data Foundation, node replacement can be performed proactively for an operational
node and reactively for a failed node for supported deployments.

 For Amazon Web Services (AWS)

  User-provisioned infrastructure

  Installer-provisioned infrastructure

 For VMware

  User-provisioned infrastructure

  Installer-provisioned infrastructure

 For Microsoft Azure

  Installer-provisioned infrastructure

 For local storage devices

  Bare metal

  VMware

  IBM Power

Additional resources

 See the Red Hat Ceph Storage documentation for replacing your storage nodes in external
 mode.










  5

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

                                                    CHAPTER 2. REPLACING STORAGE NODES DEPLOYED
                    USING DYNAMIC DEVICES

2.1. REPLACING STORAGE NODES ON AWS

     To replace an operational node, see:

         Section 2.1.1, “Replacing an operational AWS node on user-provisioned infrastructure” .

         Section 2.1.2, “Replacing an operational AWS node on installer-provisioned infrastructure” .

     To replace a failed node, see:

         Section 2.1.3, “Replacing a failed AWS node on user-provisioned infrastructure” .

         Section 2.1.4, “Replacing a failed AWS node on installer-provisioned infrastructure” .

2.1.1. Replacing an operational AWS node on user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure and resources to
     the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

             NOTE

             When replacing an AWS node on user-provisioned infrastructure, the new node needs to
             be created in the same AWS zone as the original node.

Procedure

1.   Identify the node that you need to replace.

2.   Mark the node as unschedulable:

         $ oc adm cordon <node_name>

     <node_name>
         Specify the name of node that you need to replace.

3.   Drain the node:

             $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

                 IMPORTANT

                 This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                 during this period are temporary and are automatically resolved when you label
                 the new node, and it is functional.



6

              CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

 4.  Delete the node:

     $ oc delete nodes <node_name>

 5.  Create a new Amazon Web Service (AWS) machine instance with the required infrastructure.
     See Platform requirements.

 6.  Create a new OpenShift Container Platform node using the new AWS machine instance.

 7.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

     $ oc get csr

 8.  Approve all the required OpenShift Container Platform CSRs for the new node:

     $ oc adm certificate approve <certificate_name>

     <certificate_name>
     Specify the name of the CSR.

9. Click Compute → Nodes. Confirm that the new node is in Ready state.

10.  Apply the OpenShift Data Foundation label to the new node using one of the following:

     From the user interface

      a.  For the new node, click Action Menu (⋮) → Edit Labels.

      b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

     From the command-line interface

          Apply the OpenShift Data Foundation label to the new node:

         $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

     <new_node_name>
         Specify the name of the new node.


   Verification

 1.   Verify that the new node is present in the output:

      $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
      Running state:

       openshift-storage.cephfs.csi.ceph.com-*

       openshift-storage.rbd.csi.ceph.com-*


          7

    Red Hat OpenShift Data Foundation 4.22 Replacing nodes

    3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

    4.   Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

             $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

    5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
         are encrypted.
         For each of the new nodes identified in the previous step, do the following:

          a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

               $ oc debug node/<node_name>

               $ chroot /host

          b.  Display the list of available block devices:

               $ lsblk

              Check for the crypt keyword beside the one or more  ocs-deviceset names.

    6.   If the verification steps fail, contact Red Hat Support .

     2.1.2. Replacing an operational AWS node on installer-provisioned infrastructure

Procedure

    1.   Log in to the OpenShift Web Console, and click Compute → Nodes.

    2.   Identify the node that you need to replace. Take a note of its Machine Name.

    3.   Mark the node as unschedulable:

             $ oc adm cordon <node_name>

         <node_name>
             Specify the name of node that you need to replace.

    4.   Drain the node:

             $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

                       IMPORTANT

                       This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                       during this period are temporary and are automatically resolved when you label
                       the new node, and it is functional.

    5.   Click Compute → Machines. Search for the required machine.

    6.   Besides the required machine, click Action menu (⋮) → Delete Machine.

    8

           CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

 7.  Click Delete to confirm that the machine is deleted. A new machine is automatically created.

 8.  Wait for the new machine to start and transition into Running state.

           IMPORTANT

           This activity might take at least 5 - 10 minutes or more.

9. Click Compute → Nodes. Confirm that the new node is in Ready state.

10.  Apply the OpenShift Data Foundation label to the new node:

     From the user interface

      a.  For the new node, click Action Menu (⋮) → Edit Labels.

      b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

     From the command-line interface

          Apply the OpenShift Data Foundation label to the new node:

         $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

     <new_node_name>
         Specify the name of the new node.


   Verification

 1.   Verify that the new node is present in the output:

      $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
      Running state:

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

 3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

 4.   Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

      $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

 5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
      are encrypted.
      For each of the new nodes identified in the previous step, do the following:

      a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

              $ oc debug node/<node_name>


           9

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

             $ chroot /host

     b.  Display the list of available block devices:

             $ lsblk

         Check for the crypt keyword beside the one or more   ocs-deviceset names.

6.   If the verification steps fail, contact Red Hat Support .

2.1.3. Replacing a failed AWS node on user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure and resources to
     the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Identify the Amazon Web Service (AWS) machine instance of the node that you need to replace.

2.   Log in to AWS, and terminate the AWS machine instance that you identified.

3.   Create a new AWS machine instance with the required infrastructure. See Platform
     requirements.

4.   Create a new OpenShift Container Platform node using the new AWS machine instance.

5.   Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

         $ oc get csr

6.   Approve all the required OpenShift Container Platform CSRs for the new node:

         $ oc adm certificate approve <certificate_name>

     <certificate_name>
         Specify the name of the CSR.

   7. Click Compute → Nodes. Confirm that the new node is in Ready state.

8.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

     From the user interface

         a.  For the new node, click Action Menu (⋮) → Edit Labels.

         b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

     From the command-line interface

             Execute the following command to apply the OpenShift Data Foundation label to the
10

    CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

    Execute the following command to apply the OpenShift Data Foundation label to the
    new node:

    $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

    <new_node_name>
    Specify the name of the new node.


Verification

  1.    Verify that the new node is present in the output:

            $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

  2.    Click Workloads → Pods. Confirm that at least the following pods on the new node are in
        Running state:

             openshift-storage.cephfs.csi.ceph.com-*

             openshift-storage.rbd.csi.ceph.com-*

  3.    Verify that all the other required OpenShift Data Foundation pods are in Running state.

  4.    Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

            $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

  5.    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
        are encrypted.
        For each of the new nodes identified in previous step, do the following:

         a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

              $ oc debug node/<node_name>

              $ chroot /host

         b.  Display the list of available block devices:

              $ lsblk

             Check for the crypt keyword beside the one or more  ocs-deviceset names.

    6.  If the verification steps fail, contact Red Hat Support .

    2.1.4. Replacing a failed AWS node on installer-provisioned infrastructure

 Procedure

  1.    Log in to the OpenShift Web Console, and click Compute → Nodes.

  2.    Identify the faulty node, and click on its Machine Name.

                  11

    Red Hat OpenShift Data Foundation 4.22 Replacing nodes

     3.   Click Actions → Edit Annotations, and click Add More.

     4.   Add machine.openshift.io/exclude-node-draining, and click Save.

     5.   Click Actions → Delete Machine, and click Delete.

     6.   A new machine is automatically created, wait for new machine to start.

                        IMPORTANT

                        This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                        during this period are temporary and are automatically resolved when you label
                        the new node, and it is functional.

       7. Click Compute → Nodes. Confirm that the new node is in Ready state.

     8.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

          From the user interface

           a.  For the new node, click Action Menu (⋮) → Edit Labels.

           b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

          From the command-line interface

               Apply the OpenShift Data Foundation label to the new node:

             $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

          <new_node_name>
              Specify the name of the new node.

    9.    Optional: If the failed Amazon Web Service (AWS) instance is not removed automatically,
          terminate the instance from the AWS console.

Verification

     1.   Verify that the new node is present in the output:

          $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

    2.    Click Workloads → Pods. Confirm that at least the following pods on the new node are in
          Running state:

           openshift-storage.cephfs.csi.ceph.com-*

           openshift-storage.rbd.csi.ceph.com-*

     3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

     4.   Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

          $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

    12

CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES


5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
     are encrypted.
     For each of the new nodes identified in the previous step, do the following:

      a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

      b.  Display the list of available block devices:

             $ lsblk

          Check for the crypt keyword beside the one or more  ocs-deviceset names.

6.   If the verification steps fail, contact Red Hat Support .

2.2. REPLACING STORAGE NODES ON VMWARE

     To replace an operational node, see:

          Section 2.2.1, “Replacing an operational VMware node on user-provisioned infrastructure” .

          Section 2.2.2, “Replacing an operational VMware node on installer-provisioned
          infrastructure”.

     To replace a failed node, see:

          Section 2.2.3, “Replacing a failed VMware node on user-provisioned infrastructure” .

          Section 2.2.4, “Replacing a failed VMware node on installer-provisioned infrastructure” .

2.2.1. Replacing an operational VMware node on user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure and resources to
     the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Identify the node and its Virtual Machine (VM) that you need replace.

2.   Mark the node as unschedulable:

         $ oc adm cordon <node_name>

     <node_name>
         Specify the name of node that you need to replace.


             13

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

3.    Drain the node:

      $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

                   IMPORTANT

                   This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                   during this period are temporary and are automatically resolved when you label
                   the new node, and it is functional.

4.    Delete the node:

      $ oc delete nodes <node_name>

5.    Log in to VMware vSphere, and terminate the VM that you identified:

                   IMPORTANT

                   Delete the VM only from the inventory and not from the disk.

6.    Create a new VM on VMware vSphere with the required infrastructure. See Platform
      requirements.

7.    Create a new OpenShift Container Platform worker node using the new VM.

8.    Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
      that are in Pending state:

      $ oc get csr

9.    Approve all the required OpenShift Container Platform CSRs for the new node:

      $ oc adm certificate approve <certificate_name>

      <certificate_name>
      Specify the name of the CSR.

  10. Click Compute → Nodes. Confirm that the new node is in Ready state.

11.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

      From the user interface

       a.  For the new node, click Action Menu (⋮) → Edit Labels.

       b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

      From the command-line interface

           Apply the OpenShift Data Foundation label to the new node:

          $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""


14

    CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

    <new_node_name>
    Specify the name of the new node.


Verification

  1.    Verify that the new node is present in the output:

            $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

  2.    Click Workloads → Pods. Confirm that at least the following pods on the new node are in
        Running state:

             openshift-storage.cephfs.csi.ceph.com-*

             openshift-storage.rbd.csi.ceph.com-*

  3.    Verify that all the other required OpenShift Data Foundation pods are in Running state.

  4.    Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

            $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

  5.    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
        are encrypted.
        For each of the new nodes identified in the previous step, do the following:

         a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

              $ oc debug node/<node_name>

              $ chroot /host

         b.  Display the list of available block devices:

              $ lsblk

             Check for the crypt keyword beside the one or more  ocs-deviceset names.

    6.  If the verification steps fail, contact Red Hat Support .

    2.2.2. Replacing an operational VMware node on installer-provisioned infrastructure

 Procedure

  1.    Log in to the OpenShift Web Console, and click Compute → Nodes.

  2.    Identify the node that you need to replace. Take a note of its Machine Name.

  3.    Mark the node as unschedulable:

            $ oc adm cordon <node_name>



              15

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

     <node_name>
     Specify the name of node that you need to replace.

4.   Drain the node:

     $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

           IMPORTANT

           This activity might take at least 5 - 10 minutes or more. Ceph errors generated
           during this period are temporary and are automatically resolved when you label
           the new node, and it is functional.

 5.  Click Compute → Machines. Search for the required machine.

 6.  Besides the required machine, click Action menu (⋮) → Delete Machine.

 7.  Click Delete to confirm the machine is deleted. A new machine is automatically created.

 8.  Wait for the new machine to start and transition into Running state.

           IMPORTANT

           This activity might take at least 5 - 10 minutes or more.

   9. Click Compute → Nodes. Confirm that the new node is in Ready state.

10.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

     From the user interface

      a.  For the new node, click Action Menu (⋮) → Edit Labels.

      b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

     From the command-line interface

          Apply the OpenShift Data Foundation label to the new node:

         $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

     <new_node_name>
         Specify the name of the new node.


Verification

1.   Verify that the new node is present in the output:

     $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
     Running state:

16

                                   CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

4.   Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

         $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
     are encrypted.
     For each of the new nodes identified in the previous step, do the following:

      a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

      b.  Display the list of available block devices:

             $ lsblk

          Check for the crypt keyword beside the one or more  ocs-deviceset names.

6.   If the verification steps fail, contact Red Hat Support .

2.2.3. Replacing a failed VMware node on user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure and resources to
     the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Identify the node and its Virtual Machine (VM) that you need to replace.

2.   Delete the node:

         $ oc delete nodes <node_name>

     <node_name>
         Specify the name of node that you need to replace.

3.   Log in to VMware vSphere and terminate the VM that you identified.

                    IMPORTANT

                    Delete the VM only from the inventory and not from the disk.

                                                                                                  17

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

4.  Create a new VM on VMware vSphere with the required infrastructure. See Platform
    requirements.

5.  Create a new OpenShift Container Platform worker node using the new VM.

6.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
    that are in Pending state:

    $ oc get csr

7.  Approve all the required OpenShift Container Platform CSRs for the new node:

    $ oc adm certificate approve <certificate_name>

    <certificate_name>
    Specify the name of the CSR.

   8. Click Compute → Nodes. Confirm that the new node is in Ready state.

9.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

    From the user interface

     a.  For the new node, click Action Menu (⋮) → Edit Labels.

     b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

    From the command-line interface

         Apply the OpenShift Data Foundation label to the new node:

        $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

    <new_node_name>
        Specify the name of the new node.


Verification

1.   Verify that the new node is present in the output:

     $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
     Running state:

      openshift-storage.cephfs.csi.ceph.com-*

      openshift-storage.rbd.csi.ceph.com-*

3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

4.   Verify that the new Object Storage Device (OSD) pods are running on the replacement node:



18

                         CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

             $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

    5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
         are encrypted.
         For each of the new nodes identified in the previous step, do the following:

          a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

                  $ oc debug node/<node_name>

                  $ chroot /host

          b.  Display the list of available block devices:

                  $ lsblk

              Check for the crypt keyword beside the one or more  ocs-deviceset names.

    6.   If the verification steps fail, contact Red Hat Support .

    2.2.4. Replacing a failed VMware node on installer-provisioned infrastructure

Procedure

    1.   Log in to the OpenShift Web Console, and click Compute → Nodes.

    2.   Identify the faulty node, and click on its Machine Name.

    3.   Click Actions → Edit Annotations, and click Add More.

    4.   Add machine.openshift.io/exclude-node-draining, and click Save.

    5.   Click Actions → Delete Machine, and click Delete.

    6.   A new machine is automatically created. Wait for te new machine to start.

                         IMPORTANT

                         This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                         during this period are temporary and are automatically resolved when you label
                         the new node, and it is functional.

     7. Click Compute → Nodes. Confirm that the new node is in Ready state.

    8.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

         From the user interface

              a.  For the new node, click Action Menu (⋮) → Edit Labels.

              b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

         From the command-line interface


                         19

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

               Apply the OpenShift Data Foundation label to the new node:

              $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

             <new_node_name>
              Specify the name of the new node.

9.       Optional: If the failed Virtual Machine (VM) is not removed automatically, terminate the VM
         from VMware vSphere.

Verification

1.       Verify that the new node is present in the output:

             $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

2.       Click Workloads → Pods. Confirm that at least the following pods on the new node are in
         Running state:

              openshift-storage.cephfs.csi.ceph.com-*

              openshift-storage.rbd.csi.ceph.com-*

3.       Verify that all the other required OpenShift Data Foundation pods are in Running state.

     4.  Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

           $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

5.       Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
         are encrypted.
         For each of the new nodes identified in the previous step, do the following:

          a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

               $ oc debug node/<node_name>

               $ chroot /host

          b.  Display the list of available block devices:

               $ lsblk

              Check for the crypt keyword beside the one or more ocs-deviceset names.

     6. If the verification steps fail, contact Red Hat Support .

2.3. REPLACING STORAGE NODES ON MICROSOFT AZURE

2.3.1. Replacing operational nodes on Azure installer-provisioned infrastructure


20

                                CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

   Procedure

 1.   Log in to the OpenShift Web Console, and click Compute → Nodes.

 2.   Identify the node that you need to replace. Take a note of its Machine Name.

 3.   Mark the node as unschedulable:

          $ oc adm cordon <node_name>

      <node_name>
          Specify the name of node that you need to replace.

 4.   Drain the node:

          $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

                 IMPORTANT

                 This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                 during this period are temporary and are automatically resolved when you label
                 the new node, and it is functional.

 5.   Click Compute → Machines. Search for the required machine.

 6.   Besides the required machine, click the Action menu (⋮) → Delete Machine.

 7.   Click Delete to confirm the machine is deleted. A new machine is automatically created.

 8.   Wait for the new machine to start and transition into Running state.

                 IMPORTANT

                 This activity might take at least 5 - 10 minutes or more.

    9. Click Compute → Nodes. Confirm that the new node is in Ready state.

10.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

      From the user interface

            a.  For the new node, click Action Menu (⋮) → Edit Labels.

            b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

      From the command-line interface

                Execute the following command to apply the OpenShift Data Foundation label to the
                new node:

               $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

          <new_node_name>
               Specify the name of the new node.

                                                                                               21

    Red Hat OpenShift Data Foundation 4.22 Replacing nodes

Verification

  1.    Verify that the new node is present in the output:

            $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

  2.    Click Workloads→ Pods. Confirm that at least the following pods on the new node are in
        Running state:

             openshift-storage.cephfs.csi.ceph.com-*

             openshift-storage.rbd.csi.ceph.com-*

  3.    Verify that all the other required OpenShift Data Foundation pods are in Running state.

  4.    Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

            $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

  5.    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
        are encrypted.
        For each of the new nodes identified in the previous step, do the following:

         a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

              $ oc debug node/<node_name>

              $ chroot /host

         b.  Display the list of available block devices:

              $ lsblk

             Check for the crypt keyword beside the one or more  ocs-deviceset names.

    6.  If the verification steps fail, contact Red Hat Support .

    2.3.2. Replacing failed nodes on Azure installer-provisioned infrastructure

 Procedure

  1.    Log in to the OpenShift Web Console, and click Compute → Nodes.

  2.    Identify the faulty node, and click on its Machine Name.

  3.    Click Actions → Edit Annotations, and click Add More.

  4.    Add machine.openshift.io/exclude-node-draining, and click Save.

  5.    Click Actions → Delete Machine, and click Delete.

  6.    A new machine is automatically created. Wait for the new machine to start.


              IMPORTANT
    22

                                         CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

                        IMPORTANT

                        This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                        during this period are temporary and are automatically resolved when you label
                        the new node, and it is functional.

     7. Click Compute → Nodes. Confirm that the new node is in Ready state.

     8.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

          From the user interface

              a.  For the new node, click Action Menu (⋮) → Edit Labels.

              b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

          From the command-line interface

                  Apply the OpenShift Data Foundation label to the new node:

              $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

          <new_node_name>
                 Specify the name of the new node.

    9.    Optional: If the failed Azure instance is not removed automatically, terminate the instance from
          the Azure console.

Verification

     1.   Verify that the new node is present in the output:

          $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

    2.    Click Workloads → Pods. Confirm that at least the following pods on the new node are in
          Running state:

              openshift-storage.cephfs.csi.ceph.com-*

              openshift-storage.rbd.csi.ceph.com-*

     3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

     4.   Verify that new the Object Storage Device (OSD) pods are running on the replacement node:

           $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

     5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
          are encrypted.
          For each of the new nodes identified in the previous step, do the following:

          a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

                  $ oc debug node/<node_name>

                                                                                                        23

Red Hat OpenShift Data Foundation 4.22 Replacing nodes



             $ chroot /host

         b.  Display the list of available block devices:

             $ lsblk

             Check for the crypt keyword beside the one or more   ocs-deviceset names.

6.       If the verification steps fail, contact Red Hat Support .

2.4. REPLACING STORAGE NODES ON GOOGLE CLOUD

2.4.1. Replacing operational nodes on Google Cloud installer-provisioned
infrastructure

Procedure

    1.   Log in to OpenShift Web Console and click Compute → Nodes.

    2.   Identify the node that needs to be replaced. Take a note of its Machine Name.

    3.   Mark the node as unschedulable using the following command:

            $ oc adm cordon <node_name>

4.       Drain the node using the following command:

            $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

                    IMPORTANT

                    This activity may take at least 5-10 minutes or more. Ceph errors generated
                    during this period are temporary and are automatically resolved when the new
                    node is labeled and functional.

    5.   Click Compute → Machines. Search for the required machine.

    6.   Besides the required machine, click the Action menu (⋮) → Delete Machine.

    7.   Click Delete to confirm the machine deletion. A new machine is automatically created.

    8.   Wait for new machine to start and transition into Running state.

                    IMPORTANT

                    This activity may take at least 5-10 minutes or more.

 9. Click Compute → Nodes, confirm if the new node is in Ready state.

   10.   Apply the OpenShift Data Foundation label to the new node using any one of the following:


24

         CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

    From User interface

     a.  For the new node, click Action Menu (⋮) → Edit Labels

     b.  Add cluster.ocs.openshift.io/openshift-storage and click Save.

    From Command line interface

         Execute the following command to apply the OpenShift Data Foundation label to the
         new node:

         $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""



Verification

  1.   Verify that the new node is present in the output:

           $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

  2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
       Running state:

            openshift-storage.cephfs.csi.ceph.com-*

            openshift-storage.rbd.csi.ceph.com-*

  3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

  4.   Verify that the new Object Storage Device (OSD) pods are running on the replacement node:

           $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

  5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
       are encrypted.
       For each of the new nodes identified in the previous step, do the following:

        a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

        b.  Display the list of available block devices:

             $ lsblk

            Check for the crypt keyword beside the one or more  ocs-deviceset names.

    6. If the verification steps fail, contact Red Hat Support .

    2.4.2. Replacing failed nodes on Google Cloud installer-provisioned infrastructure

 Procedure
                 25

    Red Hat OpenShift Data Foundation 4.22 Replacing nodes

 Procedure

     1.   Log in to OpenShift Web Console and click Compute → Nodes.

     2.   Identify the faulty node and click on its Machine Name.

     3.   Click Actions → Edit Annotations, and click Add More.

     4.   Add machine.openshift.io/exclude-node-draining and click Save.

     5.   Click Actions → Delete Machine, and click Delete.

     6.   A new machine is automatically created, wait for new machine to start.

                        IMPORTANT

                        This activity may take at least 5-10 minutes or more. Ceph errors generated
                        during this period are temporary and are automatically resolved when the new
                        node is labeled and functional.

       7. Click Compute → Nodes, confirm if the new node is in Ready state.

     8.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

          From the web user interface

           a.  For the new node, click Action Menu (⋮) → Edit Labels

           b.  Add cluster.ocs.openshift.io/openshift-storage and click Save.

          From the command line interface

               Apply the OpenShift Data Foundation label to the new node:

              $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

          <new_node_name>
              Specify the name of the new node.

    9.    Optional: If the failed Google Cloud instance is not removed automatically, terminate the
          instance from Google Cloud console.

Verification

     1.   Verify that the new node is present in the output:

          $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

     2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
          Running state:

           openshift-storage.cephfs.csi.ceph.com-*

           openshift-storage.rbd.csi.ceph.com-*


    26

              CHAPTER 2. REPLACING STORAGE NODES DEPLOYED USING DYNAMIC DEVICES

3.  Verify that all the other required OpenShift Data Foundation pods are in Running state.

4.  Verify that new the Object Storage Device (OSD) pods are running on the replacement node:

        $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

5.  Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.
    For each of the new nodes identified in the previous step, do the following:

     a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

          $ oc debug node/<node_name>

          $ chroot /host

     b.  Display the list of available block devices:

          $ lsblk

         Check for the crypt keyword beside the one or more ocs-deviceset names.

6. If the verification steps fail, contact Red Hat Support .










          27

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

  CHAPTER 3. REPLACING STORAGE NODES DEPLOYED
             USING LOCAL STORAGE DEVICES

The procedure for bare‑metal deployments in this chapter also applies to agnostic deployments.

3.1. REPLACING STORAGE NODES ON BARE METAL
INFRASTRUCTURE

     To replace an operational node, see Section 3.1.2, “Replacing an operational node on bare metal
     user-provisioned infrastructure”.

     To replace a failed node, see Section 3.1.5, “Replacing a failed node on bare metal user-
     provisioned infrastructure”.

3.1.1. Replacing an operational node using local storage devices

Replace an operational node on VMware user-provisioned infrastructure to perform planned
maintenance or upgrades without service interruption.

3.1.2. Replacing an operational node on bare metal user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure, resources, and
     disks to the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Identify the node, and get the labels on the node that you need to replace:

         $ oc get nodes --show-labels | grep <node_name>

     <node_name>
         Specify the name of node that you need to replace.

2.   Identify the monitor pod (if any), and OSDs that are running in the node that you need to
     replace:

         $ oc get pods -n openshift-storage -o wide | grep -i <node_name>

3.   Scale down the deployments of the pods identified in the previous step:
     For example:

         $ oc scale deployment rook-ceph-mon-c --replicas=0 -n openshift-storage

         $ oc scale deployment rook-ceph-osd-0 --replicas=0 -n openshift-storage

         $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
         --replicas=0 -n openshift-storage

28

   CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES



 4.  Mark the node as unschedulable:

         $ oc adm cordon <node_name>

 5.  Drain the node:

     $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

 6.  Delete the node:

        $ oc delete node <node_name>

 7.  Get a new bare-metal machine with the required infrastructure. See Installing on bare metal .

           IMPORTANT

           For information about how to replace a master node when you have installed
           OpenShift Data Foundation on a three-node OpenShift compact bare-metal
           cluster, see the Backup and Restore guide in the OpenShift Container Platform
           documentation.

 8.  Create a new OpenShift Container Platform node using the new bare-metal machine.

 9.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

     $ oc get csr

10.  Approve all the required OpenShift Container Platform CSRs for the new node:

     $ oc adm certificate approve <certificate_name>

     <certificate_name>
     Specify the name of the CSR.

11. Click Compute → Nodes in the OpenShift Web Console. Confirm that the new node is in Ready
   state.

12.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

     From the user interface

      a.  For the new node, click Action Menu (⋮) → Edit Labels.

      b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

     From the command-line interface

          Apply the OpenShift Data Foundation label to the new node:

         $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""



           29

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

        <new_node_name>
         Specify the name of the new node.

13.  Identify the namespace where OpenShift local storage operator is installed, and assign it to the
     local_storage_project variable:

        $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

     For example:

        $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

        echo $local_storage_project

     Example output:

        openshift-local-storage

14.  Add a new worker node to the localVolumeDiscovery and localVolumeSet.

     a.  Update the localVolumeDiscovery definition to include the new node, and remove the
         failed node:

          # oc edit -n $local_storage_project localvolumediscovery auto-discover-devices

         Example output:

          [...]
          nodeSelector:
              nodeSelectorTerms:
              - matchExpressions:
              - key: kubernetes.io/hostname
                 operator: In
                 values:
                 - server1.example.com
                 - server2.example.com
                 #- server3.example.com
                 - newnode.example.com
          [...]

         Remember to save before exiting the editor.

         In this example, server3.example.com is removed, and newnode.example.com is the new
         node.

     b.  Determine the localVolumeSet to edit:

          # oc get -n $local_storage_project localvolumeset

         Example output:




30

          CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

          NAME       AGE
          localblock 25h

     c.  Update the localVolumeSet definition to include the new node, and remove the failed
         node:

          # oc edit -n $local_storage_project localvolumeset localblock

         Example output:

          [...]
          nodeSelector:
              nodeSelectorTerms:
              - matchExpressions:
              - key: kubernetes.io/hostname
                   operator: In
                   values:
                   - server1.example.com
                   - server2.example.com
                   #- server3.example.com
                   - newnode.example.com
          [...]

         Remember to save before exiting the editor.

         In the this example, server3.example.com is removed and newnode.example.com is the
         new node.

15.  Verify that the new localblock Persistent Volume (PV) is available:

        $oc get pv | grep localblock | grep Available

     Example output:

        local-pv-551d950  512Gi  RWO     Delete Available
        localblock 26s

16.  Navigate to the openshift-storage project:

        $ oc project openshift-storage

17.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required:

        $ oc process -n openshift-storage ocs-osd-removal \
        -p FAILED_OSD_IDS=<failed_osd_id> | oc create -f -

     <failed_osd_id>
        Is the integer in the pod name immediately after the rook-ceph-osd prefix.
        You can add comma separated OSD IDs in the command to remove more than one OSD, for
        example, FAILED_OSD_IDS=0,1,2.



        Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
                       31

 Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
      insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
      Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.

18.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
     job pod.
     A status of Completed confirms that the OSD removal job succeeded.

      # oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

19.  Ensure that the OSD removal is completed.

      $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
      removal'

     Example output:

      2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

                 IMPORTANT

                 If the ocs-osd-removal-job fails, and the pod is not in the expected Completed
                 state, check the pod logs for further debugging:

                 For example:

                    # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

20.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

     If there is a PV in Released state, delete it:

      # oc delete pv <persistent_volume>

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:

      persistentvolume "local-pv-d9c5cbd6" deleted

21.  Identify the crashcollector pod deployment:

 32

                               CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

        $ oc get deployment --selector=app=rook-ceph-
        crashcollector,node_name=<failed_node_name> -n openshift-storage

       If there is an existing crashcollector pod deployment, delete it:

        $ oc delete deployment --selector=app=rook-ceph-
        crashcollector,node_name=<failed_node_name> -n openshift-storage

22.    Delete the ocs-osd-removal-job:

        # oc delete -n openshift-storage job ocs-osd-removal-job

       Example output:

        job.batch "ocs-osd-removal-job" deleted

   Verification

 1.    Verify that the new node is present in the output:

        $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.    Click Workloads → Pods. Confirm that at least the following pods on the new node are in
       Running state:

         openshift-storage.cephfs.csi.ceph.com-*

         openshift-storage.rbd.csi.ceph.com-*

 3.    Verify that all other required OpenShift Data Foundation pods are in Running state.
       Ensure that the new incremental mon is created, and is in the Running state:

        $ oc get pod -n openshift-storage | grep mon

       Example output:

        rook-ceph-mon-a-cd575c89b-b6k66        2/2       Running
        0         38m
        rook-ceph-mon-b-6776bc469b-tzzt8       2/2       Running
        0         38m
        rook-ceph-mon-d-5ff5d488b5-7v8xh       2/2       Running
        0     4m8s

       OSD and monitor pod might take several minutes to get to the Running state.

   4.  Verify that new OSD pods are running on the replacement node:

        $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

 5.    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
       are encrypted.
       For each of the new nodes identified in the previous step, do the following:

                                                                                                    33

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

      b.  Display the list of available block devices:

             $ lsblk

          Check for the crypt keyword beside the one or more  ocs-deviceset names.

6.   If the verification steps fail, contact Red Hat Support .

3.1.3. Replacing an operational node using local storage devices

Replace an operational node on VMware user-provisioned infrastructure to perform planned
maintenance or upgrades without service interruption.

3.1.4. Replacing an operational node on bare metal user-provisioned infrastructure

3.1.5. Replacing a failed node on bare metal user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure, resources, and
     disks to the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Identify the node, and get the labels on the node that you need to replace:

               $ oc get nodes --show-labels | grep <node_name>

     <node_name>
         Specify the name of node that you need to replace.

2.   Identify the node, and get the labels on the node that you need to replace:

               $ oc get nodes --show-labels | grep <node_name>

     <node_name>
         Specify the name of node that you need to replace.

3.   Identify the monitor pod (if any), and OSDs that are running in the node that you need to
     replace:

         $ oc get pods -n openshift-storage -o wide | grep -i <node_name>



34

       CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

4.    Scale down the deployments of the pods identified in the previous step:
      For example:

       $ oc scale deployment rook-ceph-mon-c --replicas=0 -n openshift-storage

       $ oc scale deployment rook-ceph-osd-0 --replicas=0 -n openshift-storage

       $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
       --replicas=0 -n openshift-storage

 5.   Mark the node as unschedulable:

       $ oc adm cordon <node_name>

 6.   Remove the pods which are in Terminating state:

       $ oc get pods -A -o wide | grep -i <node_name> | awk '{if ($4 == "Terminating") system ("oc -
       n " $1 " delete pods " $2 " --grace-period=0 " " --force ")}'

 7.   Drain the node:

       $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

 8.   Delete the node:

       $ oc delete node <node_name>

 9.   Delete the node:

       $ oc delete node <node_name>

10.   Get a new bare-metal machine with the required infrastructure. See Installing on bare metal .

                  IMPORTANT

                  For information about how to replace a master node when you have installed
                  OpenShift Data Foundation on a three-node OpenShift compact bare-metal
                  cluster, see the Backup and Restore guide in the OpenShift Container Platform
                  documentation.

11.   Create a new OpenShift Container Platform node using the new bare-metal machine.

12.   Get a new bare-metal machine with the required infrastructure. See Installing on bare metal .

                  IMPORTANT

                  For information about how to replace a master node when you have installed
                  OpenShift Data Foundation on a three-node OpenShift compact bare-metal
                  cluster, see the Backup and Restore guide in the OpenShift Container Platform
                  documentation.



                  35

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

13.  Create a new OpenShift Container Platform node using the new bare-metal machine.

14.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

     $ oc get csr

15.  Approve all the required OpenShift Container Platform CSRs for the new node:

     $ oc adm certificate approve <certificate_name>

     <certificate_name>
     Specify the name of the CSR.

16.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

     $ oc get csr

17.  Approve all the required OpenShift Container Platform CSRs for the new node:

     $ oc adm certificate approve <certificate_name>

     <certificate_name>
     Specify the name of the CSR.

  18. Click Compute → Nodes in the OpenShift Web Console. Confirm that the new node is in Ready
    state.

19.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

     From the user interface

      a.  For the new node, click Action Menu (⋮) → Edit Labels.

      b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

     From the command-line interface

          Apply the OpenShift Data Foundation label to the new node:

              $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

     <new_node_name>
         Specify the name of the new node.

20.  Identify the namespace where OpenShift local storage operator is installed, and assign it to the
     local_storage_project variable:

     $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

     For example:


36

                        CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

         $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

         echo $local_storage_project

     Example output:

         openshift-local-storage

21.  Add a new worker node to the localVolumeDiscovery and localVolumeSet.

      a.  Update the localVolumeDiscovery definition to include the new node, and remove the
          failed node:

           # oc edit -n $local_storage_project localvolumediscovery auto-discover-devices

          Example output:

           [...]
           nodeSelector:
               nodeSelectorTerms:
               - matchExpressions:
               - key: kubernetes.io/hostname
                operator: In
                values:
                - server1.example.com
                - server2.example.com
                #- server3.example.com
                - newnode.example.com
           [...]

          Remember to save before exiting the editor.

          In this example, server3.example.com is removed, and newnode.example.com is the new
          node.

      b.  Determine the localVolumeSet to edit:

           # oc get -n $local_storage_project localvolumeset

          Example output:

           NAME     AGE
           localblock 25h

      c.  Update the localVolumeSet definition to include the new node, and remove the failed
          node:

           # oc edit -n $local_storage_project localvolumeset localblock

          Example output:

           [...]
            nodeSelector:

                                                                                             37

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

           nodeSelectorTerms:
           - matchExpressions:
           - key: kubernetes.io/hostname
                 operator: In
                 values:
                 - server1.example.com
                 - server2.example.com
                 #- server3.example.com
                 - newnode.example.com
       [...]

       Remember to save before exiting the editor.

       In the this example, server3.example.com is removed and newnode.example.com is the
       new node.

22.  Verify that the new localblock Persistent Volume (PV) is available:

      $oc get pv | grep localblock | grep Available

     Example output:

      local-pv-551d950  512Gi RWO     Delete Available
      localblock 26s

23.  Navigate to the openshift-storage project:

      $ oc project openshift-storage

24.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required:

      $ oc process -n openshift-storage ocs-osd-removal \
      -p FAILED_OSD_IDS=<failed_osd_id> | oc create -f -

     <failed_osd_id>
      Is the integer in the pod name immediately after the rook-ceph-osd prefix.
      You can add comma separated OSD IDs in the command to remove more than one OSD, for
      example, FAILED_OSD_IDS=0,1,2.

      Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
      insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
      Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.

25.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
     job pod.
     A status of Completed confirms that the OSD removal job succeeded.

      # oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

26.  Ensure that the OSD removal is completed.




38

                 CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
      removal'

    Example output:

      2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

                 IMPORTANT

                 If the ocs-osd-removal-job fails, and the pod is not in the expected Completed
                 state, check the pod logs for further debugging:

                 For example:

                 # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

27.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

     If there is a PV in Released state, delete it:

      # oc delete pv <persistent_volume>

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:

      persistentvolume "local-pv-d9c5cbd6" deleted

28.  Identify the crashcollector pod deployment:

      $ oc get deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

     If there is an existing crashcollector pod deployment, delete it:

      $ oc delete deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

29.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released


                 39

   Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      Example output:

       local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
       0-6c5pw localblock 2d22h compute-1

      If there is a PV in Released state, delete it:

       # oc delete pv <persistent_volume>

      For example:

       # oc delete pv local-pv-d6bf175b

      Example output:

       persistentvolume "local-pv-d9c5cbd6" deleted

30.   Identify the crashcollector pod deployment:

       $ oc get deployment --selector=app=rook-ceph-
       crashcollector,node_name=<failed_node_name> -n openshift-storage

      If there is an existing crashcollector pod deployment, delete it:

       $ oc delete deployment --selector=app=rook-ceph-
       crashcollector,node_name=<failed_node_name> -n openshift-storage

31.   Delete the ocs-osd-removal-job:

       # oc delete -n openshift-storage job ocs-osd-removal-job

      Example output:

       job.batch "ocs-osd-removal-job" deleted

   Verification

 1.   Verify that the new node is present in the output:

       $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
      Running state:

        openshift-storage.cephfs.csi.ceph.com-*

        openshift-storage.rbd.csi.ceph.com-*

 3.   Verify that all other required OpenShift Data Foundation pods are in Running state.
      Ensure that the new incremental mon is created, and is in the Running state:



   40

            CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

          $ oc get pod -n openshift-storage | grep mon

      Example output:

          rook-ceph-mon-a-cd575c89b-b6k66    2/2       Running
          0 38m
          rook-ceph-mon-b-6776bc469b-tzzt8   2/2       Running
          0 38m
          rook-ceph-mon-d-5ff5d488b5-7v8xh   2/2       Running
          0 4m8s

      OSD and monitor pod might take several minutes to get to the Running state.

 4.   Verify that new OSD pods are running on the replacement node:

          $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

 5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
      are encrypted.
      For each of the new nodes identified in the previous step, do the following:

       a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

            $ oc debug node/<node_name>

            $ chroot /host

       b.  Display the list of available block devices:

            $ lsblk

           Check for the crypt keyword beside the one or more  ocs-deviceset names.

 6.   If the verification steps fail, contact Red Hat Support .

3.2. REPLACING STORAGE NODES ON IBM Z OR IBM® LINUXONE
INFRASTRUCTURE

You can choose one of the following procedures to replace storage nodes:

      Section 3.2.1, “Replacing operational nodes on IBM Z or IBM® LinuxONE infrastructure” .

      Section 3.2.2, “Replacing failed nodes on IBM Z or IBM® LinuxONE infrastructure” .

3.2.1. Replacing operational nodes on IBM Z or IBM® LinuxONE infrastructure

Use this procedure to replace an operational node on IBM Z or IBM® LinuxONE infrastructure.

Procedure

 1.   Identify the node and get labels on the node to be replaced. Make a note of the rack label.



                                                       41

   Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      $ oc get nodes --show-labels | grep <node_name>

 2.  Identify the mon (if any) and object storage device (OSD) pods that are running in the node to
     be replaced.

      $ oc get pods -n openshift-storage -o wide | grep -i <node_name>

 3.  Scale down the deployments of the pods identified in the previous step.
     For example:

      $ oc scale deployment rook-ceph-mon-c --replicas=0 -n openshift-storage
      $ oc scale deployment rook-ceph-osd-0 --replicas=0 -n openshift-storage
      $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
      --replicas=0 -n openshift-storage

 4.  Mark the nodes as unschedulable.

      $ oc adm cordon <node_name>

 5.  Remove the pods which are in the Terminating state.

      $ oc get pods -A -o wide | grep -i <node_name> | awk '{if ($4 == "Terminating") system ("oc -
      n " $1 " delete pods " $2 " --grace-period=0 " " --force ")}'

 6.  Drain the node.

      $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

 7.  Delete the node.

      $ oc delete node <node_name>

 8.  Get a new IBM Z storage node as a replacement.

 9.  Check for certificate signing requests (CSRs) related to OpenShift Data Foundation that are in
     Pending state:

      $ oc get csr

10.  Approve all required OpenShift Data Foundation CSRs for the new node:

      $ oc adm certificate approve <Certificate_Name>

      11. Click Compute → Nodes in OpenShift Web Console, confirm if the new node is in Ready state.

12.  Apply the openshift-storage label to the new node using any one of the following:

     From User interface

       a.  For the new node, click Action Menu (⋮) → Edit Labels

       b.  Add cluster.ocs.openshift.io/openshift-storage and click Save.


   42

CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

From Command line interface

Execute the following command to apply the OpenShift Data Foundation label to the
new node:

$ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""


13. Add a new worker node to localVolumeDiscovery and localVolumeSet.

    a.  Update the localVolumeDiscovery definition to include the new node and remove the
        failed node.

         # oc edit -n local-storage-project localvolumediscovery auto-discover-devices
         [...]
         nodeSelector:
             nodeSelectorTerms:
             - matchExpressions:
             - key: kubernetes.io/hostname
              operator: In
              values:
              - server1.example.com
              - server2.example.com
              #- server3.example.com
              - newnode.example.com
         [...]

        Remember to save before exiting the editor.

        In the above example, server3.example.com was removed and newnode.example.com is
        the new node.

    b.  Determine which localVolumeSet to edit.
        Replace local-storage-project in the following commands with the name of your local
        storage project. The default project name is openshift-local-storage in OpenShift Data
        Foundation 4.6 and later. Previous versions use local-storage by default.

         # oc get -n local-storage-project localvolumeset
         NAME     AGE
         localblock 25h

    c.  Update the localVolumeSet definition to include the new node and remove the failed
        node.

         # oc edit -n local-storage-project localvolumeset localblock
         [...]
          nodeSelector:
             nodeSelectorTerms:
             - matchExpressions:
             - key: kubernetes.io/hostname
              operator: In
              values:
              - server1.example.com
              - server2.example.com


              43

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

                    #- server3.example.com
                    - newnode.example.com
          [...]

         Remember to save before exiting the editor.

         In the above example, server3.example.com was removed and newnode.example.com is
         the new node.

14. Verify that the new localblock PV is available.

        $ oc get pv | grep localblock
             CAPA- ACCESS RECLAIM    STORAGE
        NAME     CITY MODES POLICY STATUS     CLAIM     CLASS     AGE
    local-pv- 931Gi RWO Delete Bound openshift-storage/ localblock 25h
        3e8964d3    ocs-deviceset-2-0
    -79j94
    local-pv- 931Gi RWO Delete Bound openshift-storage/ localblock 25h
        414755e0    ocs-deviceset-1-0
    -959rp
    local-pv- 931Gi RWO Delete Available localblock 3m24s b481410
    local-pv- 931Gi RWO Delete Bound openshift-storage/ localblock 25h
        d9c5cbd6            ocs-deviceset-0-0
                            -nvs68

15.  Change to the openshift-storage project.

        $ oc project openshift-storage

16.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required.

     a.  Identify the PVC as afterwards we need to delete PV associated with that specific PVC.

          $ osd_id_to_remove=1
          $ oc get -n openshift-storage -o yaml deployment rook-ceph-osd-${osd_id_to_remove} |
          grep ceph.rook.io/pvc

         where, osd_id_to_remove is the integer in the pod name immediately after the rook-ceph-
         osd prefix. In this example, the deployment name is rook-ceph-osd-1.

         Example output:

          ceph.rook.io/pvc: ocs-deviceset-localblock-0-data-0-g2mmc
             ceph.rook.io/pvc: ocs-deviceset-localblock-0-data-0-g2mmc

         In this example, the PVC name is ocs-deviceset-localblock-0-data-0-g2mmc.

     b.  Remove the failed OSD from the cluster.

          $ oc process -n openshift-storage ocs-osd-removal -p
          FAILED_OSD_IDS=${osd_id_to_remove} |oc create -f -

         You can remove more than one OSD by adding comma separated OSD IDs in the command.
         (For example: FAILED_OSD_IDS=0,1,2)


44

    CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES


 WARNING
        Proceeding will permanently remove the OSD from the cluster and
        destroy its data. Ensure osd_id_to_remove is correct. Do not process
        multiple IDs from the same failure domain at once. Contact Red Hat
        Technical Support if you need assistance. Also, make sure to check the
        state of the placement groups (PGs) and if those PGs are active and
        clean.




    17.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal
         pod.
         A status of Completed confirms that the OSD removal job succeeded.

             # oc get pod -l job-name=ocs-osd-removal-osd_id_to_remove -n openshift-storage

                NOTE

                If ocs-osd-removal fails and the pod is not in the expected Completed state,
                check the pod logs for further debugging. For example:

                          # oc logs -l job-name=ocs-osd-removal-osd_id_to_remove -n openshift-
                          storage --tail=-1

         It may be necessary to manually cleanup the removed OSD as follows:

             ceph osd crush remove osd.osd_id_to_remove
             ceph osd rm osd_id_to_remove
             ceph auth del osd.osd_id_to_remove
             ceph osd crush rm osd_id_to_remove

    18.  Delete the PV associated with the failed node.

          a.  Identify the PV associated with the PVC.
              The PVC name must be identical to the name that is obtained while removing the failed
              OSD from the cluster.

               # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released
               local-pv-5c9b8982 500Gi RWO Delete Released openshift-storage/ocs-deviceset-
               localblock-0-data-0-g2mmc localblock 24h worker-0

          b.  If there is a PV in Released state, delete it.

               # oc delete pv <persistent-volume>

              For example:




                          45

   Red Hat OpenShift Data Foundation 4.22 Replacing nodes

               # oc delete pv local-pv-5c9b8982
               persistentvolume "local-pv-5c9b8982" deleted

19.   Identify the crashcollector pod deployment.

          $ oc get deployment --selector=app=rook-ceph-crashcollector,node_name=
          <failed_node_name> -n openshift-storage

      If there is an existing crashcollector pod deployment, delete it.

          $ oc delete deployment --selector=app=rook-ceph-crashcollector,node_name=
          <failed_node_name> -n openshift-storage

20.   Delete the ocs-osd-removal job.

          # oc delete job ocs-osd-removal-${osd_id_to_remove}

      Example output:

          job.batch "ocs-osd-removal-0" deleted

   Verification

 1.   Verify that the new node is present in the output:

          $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
      Running state:

           openshift-storage.cephfs.csi.ceph.com-*

           openshift-storage.rbd.csi.ceph.com-*

 3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

 4.   Verify that new Object Storage Device (OSD) pods are running on the replacement node:

          $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

 5.   Optional: If data encryption is enabled on the cluster, verify that the new OSD devices are
      encrypted.
      For each of the new nodes identified in the previous step, do the following:

       a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

               $ oc debug node/<node_name>

               $ chroot /host

       b.  Display the list of available block devices:


   46

              CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

              $ lsblk

          Check for the crypt keyword beside the one or more     ocs-deviceset names.

    6.   If the verification steps fail, contact Red Hat Support .

    3.2.2. Replacing failed nodes on IBM Z or IBM® LinuxONE infrastructure

Procedure

    1.   Log in to the OpenShift Web Console, and click Compute → Nodes.

    2.   Identify the faulty node, and click on its Machine Name.

    3.   Click Actions → Edit Annotations, and click Add More.

    4.   Add machine.openshift.io/exclude-node-draining, and click Save.

    5.   Click Actions → Delete Machine, and click Delete.

    6.   A new machine is automatically created. Wait for new machine to start.

                     IMPORTANT

                     This activity might take at least 5 - 10 minutes or more. Ceph errors generated
                     during this period are temporary and are automatically resolved when you label
                     the new node, and it is functional.

     7. Click Compute → Nodes. Confirm that the new node is in Ready state.

    8.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

         From the user interface

          a.  For the new node, click Action Menu (⋮) → Edit Labels.

          b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

         From the command-line interface

              Apply the OpenShift Data Foundation label to the new node:

             $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

         <new_node_name>
             Specify the name of the new node.


Verification

     1. Verify that the new node is present in the output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= | cut -d' ' -f1


                     47

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
     Running state:

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

3.   Verify that all the other required OpenShift Data Foundation pods are in Running state.

4.   Verify that new Object Storage Device (OSD) pods are running on the replacement node:

         $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

5.   Optional: If data encryption is enabled on the cluster, verify that the new OSD devices are
     encrypted.
     For each of the new nodes identified in the previous step, do the following:

      a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

      b.  Display the list of available block devices:

             $ lsblk

          Check for the crypt keyword beside the one or more  ocs-deviceset names.

6.   If the verification steps fail, contact Red Hat Support .

3.3. REPLACING STORAGE NODES ON IBM POWER INFRASTRUCTURE

Replace operational or failed storage nodes on IBM Power infrastructure to maintain OpenShift Data
Foundation cluster health and data availability.

For OpenShift Data Foundation, you can perform node replacement proactively for an operational
node, and reactively for a failed node, for the deployments related to IBM Power.

3.3.1. Replacing an operational or failed storage node on IBM Power

Prerequisites

     Ensure that the replacement nodes are configured with the similar infrastructure and resources
     to the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Identify the node, and get the labels on the node that you need to replace:

               $ oc get nodes --show-labels | grep <node_name>


48

                             CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      <node_name>
       Specify the name of node that you need to replace.

 2.   Identify the mon (if any), and Object Storage Device (OSD) pods that are running in the node
      that you need to replace:

       $ oc get pods -n openshift-storage -o wide | grep -i <node_name>

 3.   Scale down the deployments of the pods identified in the previous step:
      For example:

       $ oc scale deployment rook-ceph-mon-a --replicas=0 -n openshift-storage

       $ oc scale deployment rook-ceph-osd-1 --replicas=0 -n openshift-storage

       $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
       --replicas=0 -n openshift-storage

 4.   Mark the node as unschedulable:

       $ oc adm cordon <node_name>

 5.   Remove the pods which are in Terminating state:

       $ oc get pods -A -o wide | grep -i <node_name> | awk '{if ($4 == "Terminating") system ("oc -
       n " $1 " delete pods " $2 " --grace-period=0 " " --force ")}'

 6.   Drain the node:

       $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

 7.   Delete the node:

       $ oc delete node <node_name>

 8.   Get a new IBM Power machine with the required infrastructure. See Installing a cluster on IBM
      Power.

 9.   Create a new OpenShift Container Platform node using the new IBM Power machine.

10.   Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
      that are in Pending state:

       $ oc get csr

11.   Approve all the required OpenShift Container Platform CSRs for the new node:

       $ oc adm certificate approve <certificate_name>

      <certificate_name>
       Specify the name of the CSR.

                                                                                                  49

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

   12. Click Compute → Nodes in the OpenShift Web Console. Confirm that the new node is in Ready
    state.

13.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

     From the user interface

          a.  For the new node, click Action Menu (⋮) → Edit Labels.

          b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

     From the command-line interface

          a.  Apply the OpenShift Data Foundation label to the new node:

             $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=''

         <new_node_name>
             Specify the name of the new node.

14.  Identify the namespace where OpenShift local storage operator is installed, and assign it to the
     local_storage_project variable:

         $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

     For example:

         $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

         echo $local_storage_project

     Example output:

         openshift-local-storage

15.  Add a newly added worker node to the localVolume.

      a.  Determine the localVolume you need to edit:

              # oc get -n $local_storage_project localvolume

          Example output:

              NAME     AGE
              localblock 25h

      b.  Update the localVolume definition to include the new node, and remove the failed node:

              # oc edit -n $local_storage_project localvolume localblock

          Example output:


50

          CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

          [...]
             nodeSelector:
                nodeSelectorTerms:
              - matchExpressions:
              - key: kubernetes.io/hostname
                   operator: In
                   values:
                   #- worker-0
                   - worker-1
                   - worker-2
                   - worker-3
          [...]

         Remember to save before exiting the editor.

         In the this example, worker-0 is removed and worker-3 is the new node.

16.  Verify that the new localblock Persistent Volume (PV) is available:

        $ oc get pv | grep localblock

     Example output:

        NAME        CAPACITY         ACCESSMODES RECLAIMPOLICY STATUS     CLAIM
        STORAGECLASS                  AGE
        local-pv-3e8964d3    500Gi     RWO     Delete  Bound     ocs-deviceset-localblock-2-
        data-0-mdbg9 localblock      25h
        local-pv-414755e0    500Gi     RWO     Delete  Bound     ocs-deviceset-localblock-1-
        data-0-4cslf localblock  25h
        local-pv-b481410   500Gi      RWO   Delete   Available
        localblock 3m24s
        local-pv-5c9b8982    500Gi     RWO     Delete  Bound     ocs-deviceset-localblock-0-
        data-0-g2mmc localblock       25h

17.  Navigate to the openshift-storage project:

        $ oc project openshift-storage

18.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required.

     a.  Identify the Persistent Volume Claim (PVC):

          $ osd_id_to_remove=1

          $ oc get -n openshift-storage -o yaml deployment rook-ceph-osd-${<osd_id_to_remove>}
          | grep ceph.rook.io/pvc

         where, <osd_id_to_remove> is the integer in the pod name immediately after the rook-
         ceph-osd prefix.

         In this example, the deployment name is rook-ceph-osd-1.

         Example output:



                                                                 51

 Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      ceph.rook.io/pvc: ocs-deviceset-localblock-0-data-0-g2mmc
      ceph.rook.io/pvc: ocs-deviceset-localblock-0-data-0-g2mmc

 b.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required:

      $ oc process -n openshift-storage ocs-osd-removal \
      -p FAILED_OSD_IDS=<failed_osd_id> | oc create -f -

     <failed_osd_id>
      Is the integer in the pod name immediately after the rook-ceph-osd prefix. You can add
      comma separated OSD IDs in the command to remove more than one OSD, for example,
      FAILED_OSD_IDS=0,1,2.
      Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
      insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
      Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.


       WARNING
 Proceeding will permanently remove the OSD from the cluster and
 destroy its data. Ensure osd_id_to_remove is correct. Do not
 process multiple IDs from the same failure domain at once. Contact
 Red Hat Technical Support if you need assistance.





19.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
     job pod.
     A status of Completed confirms that the OSD removal job has succeeded.

      # oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

20.  Ensure that the OSD removal is completed.

      $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
      removal'

     Example output:

      2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

              IMPORTANT

              If the ocs-osd-removal-job fails, and the pod is not in the expected Completed
              state, check the pod logs for further debugging.

              For example:

                  # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

 52

                                 CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

21.   Delete the PV associated with the failed node.

       a.  Identify the PV associated with the PVC:

            # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

           Example output:

            local-pv-5c9b8982 500Gi RWO Delete Released openshift-storage/ocs-deviceset-
            localblock-0-data-0-g2mmc localblock 24h worker-0

           The PVC name must be identical to the name that is obtained while removing the failed
           OSD from the cluster.

       b.  If there is a PV in Released state, delete it:

            # oc delete pv <persistent_volume>

           For example:

            # oc delete pv local-pv-5c9b8982

           Example output:

            persistentvolume "local-pv-5c9b8982" deleted

22.    Identify the crashcollector pod deployment:

          $ oc get deployment --selector=app=rook-ceph-
          crashcollector,node_name=<failed_node_name> -n openshift-storage

       If there is an existing crashcollector pod deployment, delete it:

          $ oc delete deployment --selector=app=rook-ceph-
          crashcollector,node_name=<failed_node_name> -n openshift-storage

23.    Delete the ocs-osd-removal-job:

          # oc delete -n openshift-storage job ocs-osd-removal-job

       Example output:

          job.batch "ocs-osd-removal-job" deleted

Verification

  1.   Verify that the new node is present in the output:

          $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

  2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
       Running state:

                                                                                                      53

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

         openshift-storage.cephfs.csi.ceph.com-*

         openshift-storage.rbd.csi.ceph.com-*

3.  Verify that all other required OpenShift Data Foundation pods are in Running state.
    Ensure that the new incremental mon is created and is in the  Running state:

        $ oc get pod -n openshift-storage | grep mon

    Example output:

        rook-ceph-mon-b-74f6dc9dd6-4llzq        1/1                   Running   0       6h14m
        rook-ceph-mon-c-74948755c-h7wtx                           1/1 Running    0  4h24m
        rook-ceph-mon-d-598f69869b-4bv49                          1/1 Running     0  162m

    The OSD and monitor pod might take several minutes to get to the Running state.

4.  Verify that the new OSD pods are running on the replacement node:

        $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

5.  Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
    are encrypted.
    For each of the new nodes identified in the previous step, do the following:

     a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

          $ oc debug node/<node_name>

          $ chroot /host

     b.  Display the list of available block devices:

          $ lsblk

         Check for the crypt keyword beside the one or more   ocs-deviceset names.

6.  If the verification steps fail, contact Red Hat Support .

                                        3.4. REPLACING STORAGE NODES ON VMWARE INFRASTRUCTURE

    To replace an operational node, see:

         Section 3.4.2, “Replacing an operational node on VMware user-provisioned infrastructure” .

         Section 3.4.5, “Replacing an operational node on VMware installer-provisioned
         infrastructure”.

    To replace a failed node,see:

         Section 3.4.9, “Replacing a failed node on VMware user-provisioned infrastructure” .

         Section 3.4.12, “Replacing a failed node on VMware installer-provisioned infrastructure” .


54

         CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

3.4.1. Replacing an operational node using local storage devices

Replace an operational node on VMware user-provisioned infrastructure to perform planned
maintenance or upgrades without service interruption.

3.4.2. Replacing an operational node on VMware user-provisioned infrastructure

3.4.3. Replacing an operational node on bare metal user-provisioned infrastructure

3.4.4. Replacing a failed node on bare metal user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure, resources, and
     disks to the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Identify the node, and get the labels on the node that you need to replace:

         $ oc get nodes --show-labels | grep <node_name>

     <node_name>
         Specify the name of node that you need to replace.

2.   Identify the node, and get the labels on the node that you need to replace:

         $ oc get nodes --show-labels | grep <node_name>

     <node_name>
         Specify the name of node that you need to replace.

3.   Identify the node, and get the labels on the node that you need to replace:

         $ oc get nodes --show-labels | grep <node_name>

     <node_name>
         Specify the name of node that you need to replace.

4.   Identify the monitor pod (if any), and OSDs that are running in the node that you need to
     replace:

         $ oc get pods -n openshift-storage -o wide | grep -i <node_name>

5.   Scale down the deployments of the pods identified in the previous step:
     For example:

         $ oc scale deployment rook-ceph-mon-c --replicas=0 -n openshift-storage



     55

   Red Hat OpenShift Data Foundation 4.22 Replacing nodes

       $ oc scale deployment rook-ceph-osd-0 --replicas=0 -n openshift-storage

       $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
       --replicas=0 -n openshift-storage

 6.   Mark the node as unschedulable:

       $ oc adm cordon <node_name>

 7.   Remove the pods which are in Terminating state:

       $ oc get pods -A -o wide | grep -i <node_name> | awk '{if ($4 == "Terminating") system ("oc -
       n " $1 " delete pods " $2 " --grace-period=0 " " --force ")}'

 8.   Drain the node:

       $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

 9.   Delete the node:

       $ oc delete node <node_name>

10.   Delete the node:

       $ oc delete node <node_name>

11.   Delete the node:

       $ oc delete node <node_name>

12.   Log in to VMware vSphere and terminate the Virtual Machine (VM) that you have identified.

13.   Create a new VM on VMware vSphere with the required infrastructure. See Infrastructure
      requirements.

14.   Create a new OpenShift Container Platform worker node using the new VM.

15.   Get a new bare-metal machine with the required infrastructure. See Installing on bare metal .

                   IMPORTANT

                   For information about how to replace a master node when you have installed
                   OpenShift Data Foundation on a three-node OpenShift compact bare-metal
                   cluster, see the Backup and Restore guide in the OpenShift Container Platform
                   documentation.

16.   Create a new OpenShift Container Platform node using the new bare-metal machine.

17.   Get a new bare-metal machine with the required infrastructure. See Installing on bare metal .



                   IMPORTANT
   56

      CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

         IMPORTANT

         For information about how to replace a master node when you have installed
         OpenShift Data Foundation on a three-node OpenShift compact bare-metal
         cluster, see the Backup and Restore guide in the OpenShift Container Platform
         documentation.

 18.  Create a new OpenShift Container Platform node using the new bare-metal machine.

 19.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
      that are in Pending state:

      $ oc get csr

 20.  Approve all the required OpenShift Container Platform CSRs for the new node:

      $ oc adm certificate approve <certificate_name>

      <certificate_name>
      Specify the name of the CSR.

 21.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
      that are in Pending state:

      $ oc get csr

 22.  Approve all the required OpenShift Container Platform CSRs for the new node:

      $ oc adm certificate approve <certificate_name>

      <certificate_name>
      Specify the name of the CSR.

 23.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
      that are in Pending state:

      $ oc get csr

 24.  Approve all the required OpenShift Container Platform CSRs for the new node:

      $ oc adm certificate approve <certificate_name>

      <certificate_name>
      Specify the name of the CSR.

25. Click Compute → Nodes in the OpenShift Web Console. Confirm that the new node is in Ready
    state.

 26.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

      From the user interface

      a. For the new node, click Action Menu (⋮) → Edit Labels.

             57

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

           b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

       From the command-line interface

               Apply the OpenShift Data Foundation label to the new node:

              $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

          <new_node_name>
              Specify the name of the new node.

  27.  Identify the namespace where OpenShift local storage operator is installed, and assign it to the
       local_storage_project variable:

          $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

       For example:

          $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

          echo $local_storage_project

       Example output:

          openshift-local-storage

  28.  Add a new worker node to the localVolumeDiscovery and localVolumeSet.

       a.  Update the localVolumeDiscovery definition to include the new node, and remove the
           failed node:

               # oc edit -n $local_storage_project localvolumediscovery auto-discover-devices

           Example output:

               [...]
               nodeSelector:
                nodeSelectorTerms:
                   - matchExpressions:
                   - key: kubernetes.io/hostname
                    operator: In
                    values:
                    - server1.example.com
                    - server2.example.com
                    #- server3.example.com
                    - newnode.example.com
               [...]

           Remember to save before exiting the editor.

           In this example, server3.example.com is removed, and newnode.example.com is the new
           node.

58

                       CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      b.  Determine the localVolumeSet to edit:

           # oc get -n $local_storage_project localvolumeset

          Example output:

           NAME       AGE
           localblock 25h

      c.  Update the localVolumeSet definition to include the new node, and remove the failed
          node:

           # oc edit -n $local_storage_project localvolumeset localblock

          Example output:

           [...]
           nodeSelector:
               nodeSelectorTerms:
               - matchExpressions:
               - key: kubernetes.io/hostname
                    operator: In
                    values:
                    - server1.example.com
                    - server2.example.com
                    #- server3.example.com
                    - newnode.example.com
           [...]

          Remember to save before exiting the editor.

          In the this example, server3.example.com is removed and newnode.example.com is the
          new node.

29.  Verify that the new localblock Persistent Volume (PV) is available:

         $oc get pv | grep localblock | grep Available

     Example output:

         local-pv-551d950  512Gi  RWO     Delete Available
         localblock 26s

30.  Navigate to the openshift-storage project:

         $ oc project openshift-storage

31.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required:

         $ oc process -n openshift-storage ocs-osd-removal \
         -p FAILED_OSD_IDS=<failed_osd_id> | oc create -f -

     <failed_osd_id>

                                                                                            59

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      Is the integer in the pod name immediately after the rook-ceph-osd prefix.
      You can add comma separated OSD IDs in the command to remove more than one OSD, for
      example, FAILED_OSD_IDS=0,1,2.

      Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
      insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
      Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.

32.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
     job pod.
     A status of Completed confirms that the OSD removal job succeeded.

      # oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

33.  Ensure that the OSD removal is completed.

      $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
      removal'

     Example output:

      2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

                 IMPORTANT

                 If the ocs-osd-removal-job fails, and the pod is not in the expected Completed
                 state, check the pod logs for further debugging:

                 For example:

                    # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

34.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

     If there is a PV in Released state, delete it:

      # oc delete pv <persistent_volume>

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:


60

                    CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      persistentvolume "local-pv-d9c5cbd6" deleted

35.  Identify the crashcollector pod deployment:

      $ oc get deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

     If there is an existing crashcollector pod deployment, delete it:

      $ oc delete deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

36.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

     If there is a PV in Released state, delete it:

      # oc delete pv <persistent_volume>

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:

      persistentvolume "local-pv-d9c5cbd6" deleted

37.  Identify the crashcollector pod deployment:

      $ oc get deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

     If there is an existing crashcollector pod deployment, delete it:

      $ oc delete deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

38.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

                                                                                         61

 Red Hat OpenShift Data Foundation 4.22 Replacing nodes


      If there is a PV in Released state, delete it:

       # oc delete pv <persistent_volume>

      For example:

       # oc delete pv local-pv-d6bf175b

      Example output:

       persistentvolume "local-pv-d9c5cbd6" deleted

39.   Identify the crashcollector pod deployment:

       $ oc get deployment --selector=app=rook-ceph-
       crashcollector,node_name=<failed_node_name> -n openshift-storage

      If there is an existing crashcollector pod deployment, delete it:

       $ oc delete deployment --selector=app=rook-ceph-
       crashcollector,node_name=<failed_node_name> -n openshift-storage

 40.  Delete the ocs-osd-removal-job:

       # oc delete -n openshift-storage job ocs-osd-removal-job

      Example output:

       job.batch "ocs-osd-removal-job" deleted

 Verification

 1.   Verify that the new node is present in the output:

       $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
      Running state:

        openshift-storage.cephfs.csi.ceph.com-*

        openshift-storage.rbd.csi.ceph.com-*

 3.   Verify that all other required OpenShift Data Foundation pods are in Running state.
      Ensure that the new incremental mon is created, and is in the Running state:

       $ oc get pod -n openshift-storage | grep mon

      Example output:

       rook-ceph-mon-a-cd575c89b-b6k66     2/2     Running

 62

              CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

         0    38m
         rook-ceph-mon-b-6776bc469b-tzzt8   2/2       Running
         0    38m
         rook-ceph-mon-d-5ff5d488b5-7v8xh   2/2       Running
         0    4m8s

     OSD and monitor pod might take several minutes to get to the Running state.

4.   Verify that new OSD pods are running on the replacement node:

         $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
     are encrypted.
     For each of the new nodes identified in the previous step, do the following:

      a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

      b.  Display the list of available block devices:

             $ lsblk

          Check for the crypt keyword beside the one or more  ocs-deviceset names.

6.   If the verification steps fail, contact Red Hat Support .

3.4.5. Replacing an operational node on VMware installer-provisioned
infrastructure

Replace an operational node on VMware installer-provisioned infrastructure to perform planned
maintenance or upgrades without service interruption.

Prerequisites

     Ensure that the replacement nodes are configured with the similar infrastructure, resources,
     and disks to the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Log in to the OpenShift Web Console, and click Compute → Nodes.

2.   Identify the node that you need to replace. Take a note of its Machine Name.

3.   Get labels on the node:

         $ oc get nodes --show-labels | grep <node_name>



                                                      63

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      <node_name>
       Specify the name of node that you need to replace.

4.    Identify the mon (if any), and Object Storage Devices (OSDs) that are running in the node:

       $ oc get pods -n openshift-storage -o wide | grep -i <node_name>

 5.   Scale down the deployments of the pods that you identified in the previous step:
      For example:

       $ oc scale deployment rook-ceph-mon-c --replicas=0 -n openshift-storage

       $ oc scale deployment rook-ceph-osd-0 --replicas=0 -n openshift-storage

       $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
       --replicas=0 -n openshift-storage

 6.   Mark the node as unschedulable:

       $ oc adm cordon <node_name>

7.    Drain the node:

       $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

 8.   Click Compute → Machines. Search for the required machine.

 9.   Besides the required machine, click Action menu (⋮) → Delete Machine.

10.   Click Delete to confirm the machine deletion. A new machine is automatically created.

11.   Wait for the new machine to start and transition into Running state.

                  IMPORTANT

                  This activity might take at least 5 - 10 minutes or more.

   12. Click Compute → Nodes in the OpenShift Web Console. Confirm that the new node is in Ready
    state.

13.   Physically add a new device to the node.

14.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

      From the user interface

        a.  For the new node, click Action Menu (⋮) → Edit Labels.

        b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

      From the command-line interface

            Apply the OpenShift Data Foundation label to the new node:

64

                              CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

         $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

        <new_node_name>
         Specify the name of the new node.

15.  Identify the namespace where the OpenShift local storage operator is installed, and assign it to
     the local_storage_project variable:

        $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

     For example:

        $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

        echo $local_storage_project

     Example output:

        openshift-local-storage

16.  Add a new worker node to the localVolumeDiscovery and localVolumeSet.

     a.  Update the localVolumeDiscovery definition to include the new node and remove the
         failed node.

          # oc edit -n $local_storage_project localvolumediscovery auto-discover-devices

         Example output:

          [...]
          nodeSelector:
              nodeSelectorTerms:
              - matchExpressions:
              - key: kubernetes.io/hostname
                 operator: In
                 values:
                 - server1.example.com
                 - server2.example.com
                 #- server3.example.com
                 - newnode.example.com
          [...]

         Remember to save before exiting the editor.

         In this example, server3.example.com is removed, and newnode.example.com is the new
         node.

     b.  Determine the localVolumeSet you need to edit:

          # oc get -n $local_storage_project localvolumeset

         Example output:

                                                                                                   65

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

          NAME       AGE
          localblock 25h

     c.  Update the localVolumeSet definition to include the new node and remove the failed
         node:

          # oc edit -n $local_storage_project localvolumeset localblock

         Example output:

          [...]
          nodeSelector:
              nodeSelectorTerms:
              - matchExpressions:
              - key: kubernetes.io/hostname
                   operator: In
                   values:
                   - server1.example.com
                   - server2.example.com
                   #- server3.example.com
                   - newnode.example.com
          [...]

         Remember to save before exiting the editor.

         In this example, server3.example.com is removed, and newnode.example.com is the new
         node.

17.  Verify that the new localblock Persistent Volume (PV) is available:

        $ oc get pv | grep localblock | grep Available

     Example output:

        local-pv-551d950  512Gi  RWO     Delete Available
        localblock 26s

18.  Navigate to the openshift-storage project:

        $ oc project openshift-storage

19.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required:

        $ oc process -n openshift-storage ocs-osd-removal \
        -p FAILED_OSD_IDS=<failed_osd_id> | oc create -f -

     <failed_osd_id>
        Is the integer in the pod name immediately after the rook-ceph-osd prefix.
        You can add comma separated OSD IDs in the command to remove more than one OSD, for
        example, FAILED_OSD_IDS=0,1,2.



        Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
66

                 CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
      insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
      Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.

20.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
     job pod.
     A status of Completed confirms that the OSD removal job succeeded.

      # oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

21.  Ensure that the OSD removal is completed.

      $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
      removal'

     Example output:

      2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

                 IMPORTANT

                 If the ocs-osd-removal-job fails and the pod is not in the expected Completed
                 state, check the pod logs for further debugging.

                 For example:

                 # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

22.  Identify the PV associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

     If there is a PV in Released state, delete it:

      # oc delete pv <persistent_volume>

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:

      persistentvolume "local-pv-d9c5cbd6" deleted

23.  Identify the crashcollector pod deployment:


                 67

 Red Hat OpenShift Data Foundation 4.22 Replacing nodes

       $ oc get deployment --selector=app=rook-ceph-
       crashcollector,node_name=<failed_node_name> -n openshift-storage

      If there is an existing crashcollector pod deployment, delete it:

       $ oc delete deployment --selector=app=rook-ceph-
       crashcollector,node_name=<failed_node_name> -n openshift-storage

24.   Delete the ocs-osd-removal-job:

       # oc delete -n openshift-storage job ocs-osd-removal-job

      Example output:

       job.batch "ocs-osd-removal-job" deleted

 Verification

 1.   Verify that the new node is present in the output:

       $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.   Click Workloads → Pods. Confirm that at least the following pods on the new node are in
      Running state:

        openshift-storage.cephfs.csi.ceph.com-*

        openshift-storage.rbd.csi.ceph.com-*

 3.   Verify that all other required OpenShift Data Foundation pods are in Running state.
      Ensure that the new incremental mon is created and is in the Running state.

       $ oc get pod -n openshift-storage | grep mon

      Example output:

       rook-ceph-mon-a-cd575c89b-b6k66        2/2       Running
       0         38m
       rook-ceph-mon-b-6776bc469b-tzzt8       2/2       Running
       0         38m
       rook-ceph-mon-d-5ff5d488b5-7v8xh       2/2       Running
       0     4m8s

      OSD and monitor pod might take several minutes to get to the Running state.

 4.   Verify that new OSD pods are running on the replacement node:

       $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

 5.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
      are encrypted.
      For each of the new nodes identified in the previous step, do the following:

 68

                          CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

      b.  Display the list of available block devices:

             $ lsblk

          Check for the crypt keyword beside the one or more  ocs-deviceset names.

6.   If the verification steps fail, contact Red Hat Support .

3.4.6. Replacing an operational node using local storage devices

Replace an operational node on VMware user-provisioned infrastructure to perform planned
maintenance or upgrades without service interruption.

3.4.7. Replacing an operational node on VMware user-provisioned infrastructure

3.4.8. Replacing an operational node on VMware installer-provisioned
infrastructure

3.4.9. Replacing a failed node on VMware user-provisioned infrastructure

3.4.10. Replacing an operational node on bare metal user-provisioned infrastructure

3.4.11. Replacing a failed node on bare metal user-provisioned infrastructure

Prerequisites

     Ensure that the replacement nodes are configured with similar infrastructure, resources, and
     disks to the node that you replace.

     You must be logged into the OpenShift Container Platform cluster.

Procedure

1.   Log in to the OpenShift Web Console, and click Compute → Nodes.

2.   Identify the node that you need to replace. Take a note of its Machine Name.

3.   Get the labels on the node:

               $ oc get nodes --show-labels | grep <node_name>

     <node_name>
         Specify the name of node that you need to replace.

4.   Identify the node, and get the labels on the node that you need to replace:

                                                                                               69

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

       $ oc get nodes --show-labels | grep <node_name>

      <node_name>
       Specify the name of node that you need to replace.

 5.   Identify the node, and get the labels on the node that you need to replace:

       $ oc get nodes --show-labels | grep <node_name>

      <node_name>
       Specify the name of node that you need to replace.

6.    Identify the node, and get the labels on the node that you need to replace:

       $ oc get nodes --show-labels | grep <node_name>

      <node_name>
       Specify the name of node that you need to replace.

7.    Identify the node, and get the labels on the node that you need to replace:

       $ oc get nodes --show-labels | grep <node_name>

      <node_name>
       Specify the name of node that you need to replace.

 8.   Identify the monitor pod (if any), and OSDs that are running in the node that you need to
      replace:

       $ oc get pods -n openshift-storage -o wide | grep -i <node_name>

 9.   Scale down the deployments of the pods identified in the previous step:
      For example:

       $ oc scale deployment rook-ceph-mon-c --replicas=0 -n openshift-storage

       $ oc scale deployment rook-ceph-osd-0 --replicas=0 -n openshift-storage

       $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
       --replicas=0 -n openshift-storage

10.   Mark the node as unschedulable:

       $ oc adm cordon <node_name>

11.   Remove the pods which are in Terminating state:

       $ oc get pods -A -o wide | grep -i <node_name> | awk '{if ($4 == "Terminating") system ("oc -
       n " $1 " delete pods " $2 " --grace-period=0 " " --force ")}'




70

                            CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

12.  Remove the pods which are in Terminating state:

      $ oc get pods -A -o wide | grep -i <node_name> | awk '{if ($4 == "Terminating") system ("oc -
      n " $1 " delete pods " $2 " --grace-period=0 " " --force ")}'

13.  Drain the node:

      $ oc adm drain <node_name> --force --delete-emptydir-data=true --ignore-daemonsets

14.  Click Compute → Machines. Search for the required machine.

15.  Besides the required machine, click Action menu (⋮) → Delete Machine.

16.  Click Delete to confirm that the machine is deleted. A new machine is automatically created.

17.  Wait for the new machine to start and transition into Running state.

                  IMPORTANT

                  This activity might take at least 5 - 10 minutes or more.

18.  Delete the node:

      $ oc delete node <node_name>

19.  Delete the node:

      $ oc delete node <node_name>

20.  Delete the node:

      $ oc delete node <node_name>

21.  Delete the node:

      $ oc delete node <node_name>

22.  Log in to VMware vSphere and terminate the Virtual Machine (VM) that you have identified.

23.  Create a new VM on VMware vSphere with the required infrastructure. See Infrastructure
     requirements.

24.  Create a new OpenShift Container Platform worker node using the new VM.

25.  Log in to VMware vSphere and terminate the Virtual Machine (VM) that you have identified.

26.  Create a new VM on VMware vSphere with the required infrastructure. See Infrastructure
     requirements.

27.  Create a new OpenShift Container Platform worker node using the new VM.

28.  Get a new bare-metal machine with the required infrastructure. See Installing on bare metal .

                  IMPORTANT
                                                                                                 71

   Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      IMPORTANT

      For information about how to replace a master node when you have installed
      OpenShift Data Foundation on a three-node OpenShift compact bare-metal
      cluster, see the Backup and Restore guide in the OpenShift Container Platform
      documentation.

29.  Create a new OpenShift Container Platform node using the new bare-metal machine.

30.  Get a new bare-metal machine with the required infrastructure. See Installing on bare metal .

      IMPORTANT

      For information about how to replace a master node when you have installed
      OpenShift Data Foundation on a three-node OpenShift compact bare-metal
      cluster, see the Backup and Restore guide in the OpenShift Container Platform
      documentation.

31.  Create a new OpenShift Container Platform node using the new bare-metal machine.

32.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

     $ oc get csr

33.  Approve all the required OpenShift Container Platform CSRs for the new node:

     $ oc adm certificate approve <certificate_name>

     <certificate_name>
     Specify the name of the CSR.

34.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

     $ oc get csr

35.  Approve all the required OpenShift Container Platform CSRs for the new node:

     $ oc adm certificate approve <certificate_name>

     <certificate_name>
     Specify the name of the CSR.

36.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
     that are in Pending state:

     $ oc get csr

37.  Approve all the required OpenShift Container Platform CSRs for the new node:

     $ oc adm certificate approve <certificate_name>


   72

            CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      <certificate_name>
       Specify the name of the CSR.

 38.  Check for the Certificate Signing Requests (CSRs) related to OpenShift Container Platform
      that are in Pending state:

       $ oc get csr

 39.  Approve all the required OpenShift Container Platform CSRs for the new node:

       $ oc adm certificate approve <certificate_name>

      <certificate_name>
       Specify the name of the CSR.

40. Click Compute → Nodes in the OpenShift Web Console. Confirm that the new node is in Ready
     state.

 41.  Physically add a new device to the node.

 42.  Apply the OpenShift Data Foundation label to the new node using any one of the following:

      From the user interface

        a.  For the new node, click Action Menu (⋮) → Edit Labels.

        b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

      From the command-line interface

            Apply the OpenShift Data Foundation label to the new node:

           $ oc label node <new_node_name> cluster.ocs.openshift.io/openshift-storage=""

       <new_node_name>
           Specify the name of the new node.

 43.  Identify the namespace where OpenShift local storage operator is installed, and assign it to the
      local_storage_project variable:

       $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

      For example:

       $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

       echo $local_storage_project

      Example output:

       openshift-local-storage


            73

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

44. Add a new worker node to the localVolumeDiscovery and localVolumeSet.

    a.  Update the localVolumeDiscovery definition to include the new node, and remove the
        failed node:

         # oc edit -n $local_storage_project localvolumediscovery auto-discover-devices

        Example output:

         [...]
         nodeSelector:
             nodeSelectorTerms:
             - matchExpressions:
             - key: kubernetes.io/hostname
              operator: In
              values:
              - server1.example.com
              - server2.example.com
              #- server3.example.com
              - newnode.example.com
         [...]

        Remember to save before exiting the editor.

        In this example, server3.example.com is removed, and newnode.example.com is the new
        node.

    b.  Determine the localVolumeSet to edit:

         # oc get -n $local_storage_project localvolumeset

        Example output:

         NAME     AGE
         localblock 25h

    c.  Update the localVolumeSet definition to include the new node, and remove the failed
        node:

         # oc edit -n $local_storage_project localvolumeset localblock

        Example output:

         [...]
          nodeSelector:
             nodeSelectorTerms:
             - matchExpressions:
             - key: kubernetes.io/hostname
              operator: In
              values:
              - server1.example.com
              - server2.example.com




74

           CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

                 #- server3.example.com
                 - newnode.example.com
       [...]

       Remember to save before exiting the editor.

       In the this example, server3.example.com is removed and newnode.example.com is the
       new node.

45.  Verify that the new localblock Persistent Volume (PV) is available:

      $oc get pv | grep localblock | grep Available

     Example output:

      local-pv-551d950 512Gi RWO     Delete Available
      localblock 26s

46.  Navigate to the openshift-storage project:

      $ oc project openshift-storage

47.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required:

      $ oc process -n openshift-storage ocs-osd-removal \
      -p FAILED_OSD_IDS=<failed_osd_id> | oc create -f -

     <failed_osd_id>
      Is the integer in the pod name immediately after the rook-ceph-osd prefix.
      You can add comma separated OSD IDs in the command to remove more than one OSD, for
      example, FAILED_OSD_IDS=0,1,2.

      Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
      insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
      Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.

48.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
     job pod.
     A status of Completed confirms that the OSD removal job succeeded.

      # oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

49.  Ensure that the OSD removal is completed.

      $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
      removal'

     Example output:

      2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0


       IMPORTANT
                 75

   Red Hat OpenShift Data Foundation 4.22 Replacing nodes

                    IMPORTANT

                    If the ocs-osd-removal-job fails, and the pod is not in the expected Completed
                    state, check the pod logs for further debugging:

                    For example:

                       # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

50.     Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

         # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

        Example output:

         local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
         0-6c5pw localblock 2d22h compute-1

        If there is a PV in Released state, delete it:

         # oc delete pv <persistent_volume>

        For example:

         # oc delete pv local-pv-d6bf175b

        Example output:

         persistentvolume "local-pv-d9c5cbd6" deleted

   51.  Identify the crashcollector pod deployment:

         $ oc get deployment --selector=app=rook-ceph-
         crashcollector,node_name=<failed_node_name> -n openshift-storage

        If there is an existing crashcollector pod deployment, delete it:

         $ oc delete deployment --selector=app=rook-ceph-
         crashcollector,node_name=<failed_node_name> -n openshift-storage

52.     Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

         # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

        Example output:

         local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
         0-6c5pw localblock 2d22h compute-1

        If there is a PV in Released state, delete it:



   76

      CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

      # oc delete pv <persistent_volume>

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:

      persistentvolume "local-pv-d9c5cbd6" deleted

53.  Identify the crashcollector pod deployment:

      $ oc get deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

     If there is an existing crashcollector pod deployment, delete it:

      $ oc delete deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

54.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

     If there is a PV in Released state, delete it:

      # oc delete pv <persistent_volume>

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:

      persistentvolume "local-pv-d9c5cbd6" deleted

55.  Identify the crashcollector pod deployment:

      $ oc get deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage

     If there is an existing crashcollector pod deployment, delete it:

      $ oc delete deployment --selector=app=rook-ceph-
      crashcollector,node_name=<failed_node_name> -n openshift-storage


      77

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

56.    Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

        # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

       Example output:

        local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
        0-6c5pw localblock 2d22h compute-1

       If there is a PV in Released state, delete it:

        # oc delete pv <persistent_volume>

       For example:

        # oc delete pv local-pv-d6bf175b

       Example output:

        persistentvolume "local-pv-d9c5cbd6" deleted

  57.  Identify the crashcollector pod deployment:

        $ oc get deployment --selector=app=rook-ceph-
        crashcollector,node_name=<failed_node_name> -n openshift-storage

       If there is an existing crashcollector pod deployment, delete it:

        $ oc delete deployment --selector=app=rook-ceph-
        crashcollector,node_name=<failed_node_name> -n openshift-storage

  58.  Identify the Persistent Volume (PV) associated with the Persistent Volume Claim (PVC):

        # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

       Example output:

        local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
        0-6c5pw localblock 2d22h compute-1

       If there is a PV in Released state, delete it:

        # oc delete pv <persistent_volume>

       For example:

        # oc delete pv local-pv-d6bf175b

       Example output:

        persistentvolume "local-pv-d9c5cbd6" deleted

78

              CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

59.     Identify the crashcollector pod deployment:

         $ oc get deployment --selector=app=rook-ceph-
         crashcollector,node_name=<failed_node_name> -n openshift-storage

        If there is an existing crashcollector pod deployment, delete it:

         $ oc delete deployment --selector=app=rook-ceph-
         crashcollector,node_name=<failed_node_name> -n openshift-storage

   60.  Delete the ocs-osd-removal-job:

         # oc delete -n openshift-storage job ocs-osd-removal-job

        Example output:

         job.batch "ocs-osd-removal-job" deleted

   Verification

 1.     Verify that the new node is present in the output:

         $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.     Click Workloads → Pods. Confirm that at least the following pods on the new node are in
        Running state:

          openshift-storage.cephfs.csi.ceph.com-*

          openshift-storage.rbd.csi.ceph.com-*

 3.     Verify that all other required OpenShift Data Foundation pods are in Running state.
        Ensure that the new incremental mon is created, and is in the Running state:

         $ oc get pod -n openshift-storage | grep mon

        Example output:

         rook-ceph-mon-a-cd575c89b-b6k66        2/2       Running
         0  38m
         rook-ceph-mon-b-6776bc469b-tzzt8       2/2       Running
         0  38m
         rook-ceph-mon-d-5ff5d488b5-7v8xh       2/2       Running
         0     4m8s

        OSD and monitor pod might take several minutes to get to the Running state.

 4.     Verify that new OSD pods are running on the replacement node:

         $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd



   5. Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
                                                              79

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

  5.  Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
      are encrypted.
      For each of the new nodes identified in the previous step, do the following:

       a.   Create a debug pod and open a chroot environment for the one or more selected hosts:

             $ oc debug node/<node_name>

             $ chroot /host

       b.   Display the list of available block devices:

             $ lsblk

            Check for the crypt keyword beside the one or more ocs-deviceset names.

6.    If the verification steps fail, contact Red Hat Support .

3.4.12. Replacing a failed node on VMware installer-provisioned infrastructure

Replace a failed node on VMware installer-provisioned infrastructure when the node is no longer
operational.

Prerequisites

      Ensure that the replacement nodes are configured with the similar infrastructure, resources,
      and disks to the node that you replace.

      You must be logged into the OpenShift Container Platform cluster.

Procedure

1.    Log in to the OpenShift Web Console, and click Compute → Nodes.

2.    Identify the node that you need to replace. Take a note of its Machine Name.

3.    Get the labels on the node:

          $ oc get nodes --show-labels | grep _<node_name>_

      <node_name>
          Specify the name of node that you need to replace.

  4.  Identify the mon (if any) and Object Storage Devices (OSDs) that are running in the node:

          $ oc get pods -n openshift-storage -o wide | grep -i _<node_name>_

5.    Scale down the deployments of the pods identified in the previous step:
      For example:

          $ oc scale deployment rook-ceph-mon-c --replicas=0 -n openshift-storage

          $ oc scale deployment rook-ceph-osd-0 --replicas=0 -n openshift-storage

80

    CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES



        $ oc scale deployment --selector=app=rook-ceph-crashcollector,node_name=<node_name>
        --replicas=0 -n openshift-storage

  6.   Mark the node as unschedulable:

        $ oc adm cordon _<node_name>_

  7.   Remove the pods which are in Terminating state:

        $ oc get pods -A -o wide | grep -i _<node_name>_ | awk '{if ($4 == "Terminating") system
        ("oc -n " $1 " delete pods " $2 " --grace-period=0 " " --force ")}'

  8.   Drain the node:

        $ oc adm drain _<node_name>_ --force --delete-emptydir-data=true --ignore-daemonsets

  9.   Click Compute → Machines. Search for the required machine.

 10.   Besides the required machine, click Action menu (⋮) → Delete Machine.

 11.   Click Delete to confirm the machine is deleted. A new machine is automatically created.

 12.   Wait for the new machine to start and transition into Running state.

              IMPORTANT

              This activity might take at least 5 - 10 minutes or more.

13. Click Compute → Nodes in the OpenShift Web Console. Confirm that the new node is in Ready
    state.

 14.   Physically add a new device to the node.

 15.   Apply the OpenShift Data Foundation label to the new node using any one of the following:

       From the user interface

         a.  For the new node, click Action Menu (⋮) → Edit Labels.

         b.  Add cluster.ocs.openshift.io/openshift-storage, and click Save.

       From the command-line interface

             Apply the OpenShift Data Foundation label to the new node:

            $ oc label node _<new_node_name>_ cluster.ocs.openshift.io/openshift-storage=""

        <new_node_name>
            Specify the name of the new node.




    16. Identify the namespace where the OpenShift local storage operator is installed, and assign it to
                  81

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

16.  Identify the namespace where the OpenShift local storage operator is installed, and assign it to
     the local_storage_project variable:

         $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

     For example:

         $ local_storage_project=$(oc get csv --all-namespaces | awk '{print $1}' | grep local)

         echo $local_storage_project

     Example output:

         openshift-local-storage

17.  Add a new worker node to the localVolumeDiscovery and localVolumeSet.

      a.  Update the localVolumeDiscovery definition to include the new node and remove the
          failed node:

           # oc edit -n $local_storage_project localvolumediscovery auto-discover-devices

          Example output:

           [...]
            nodeSelector:
               nodeSelectorTerms:
               - matchExpressions:
               - key: kubernetes.io/hostname
                 operator: In
                 values:
                 - server1.example.com
                 - server2.example.com
                 #- server3.example.com
                 - **newnode.example.com**
           [...]

          Remember to save before exiting the editor.

          In this example, server3.example.com is removed and newnode.example.com is the new
          node.

      b.  Determine the localVolumeSet you need to edit.

           # oc get -n $local_storage_project localvolumeset

          Example output:

           NAME      AGE
           localblock 25h

      c.  Update the localVolumeSet definition to include the new node and remove the failed
          node:

82

                         CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

        # oc edit -n $local_storage_project localvolumeset localblock

      Example output:

        [...]
            nodeSelector:
            nodeSelectorTerms:
            - matchExpressions:
            - key: kubernetes.io/hostname
                 operator: In
                 values:
                 - server1.example.com
                 - server2.example.com
                 #- server3.example.com
                 - **newnode.example.com**
        [...]

       Remember to save before exiting the editor.

       In this example, server3.example.com is removed and newnode.example.com is the new
       node.

18.  Verify that the new localblock PV is available:

      $ oc get pv | grep localblock | grep Available

     Example output:

      local-pv-551d950   512Gi RWO     Delete Available
      localblock 26s

19.  Navigate to the openshift-storage project:

      $ oc project openshift-storage

20.  Remove the failed OSD from the cluster. You can specify multiple failed OSDs if required:

      $ oc process -n openshift-storage ocs-osd-removal \
      -p FAILED_OSD_IDS=<failed_osd_id> | oc create -f -

     <failed_osd_id>
      Is the integer in the pod name immediately after the rook-ceph-osd prefix.
      You can add comma separated OSD IDs in the command to remove more than one OSD, for
      example, FAILED_OSD_IDS=0,1,2.

      Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
      insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
      Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.

21.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
     job pod.
     A status of Completed confirms that the OSD removal job succeeded.

                                                                                              83

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

      # oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

22.  Ensure that the OSD removal is completed.

      $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
      removal'

     Example output:

      2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

                 IMPORTANT

                 If the ocs-osd-removal-job fails and the pod is not in the expected Completed
                 state, check the pod logs for further debugging:

                 For example:

                    # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

23.  Identify the PV associated with the Persistent Volume Claim (PVC):

      # oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

     Example output:

      local-pv-d6bf175b 1490Gi RWO Delete Released openshift-storage/ocs-deviceset-0-data-
      0-6c5pw localblock 2d22h compute-1

     If there is a PV in Released state, delete it:

      # oc delete pv _<persistent_volume>_

     For example:

      # oc delete pv local-pv-d6bf175b

     Example output:

      persistentvolume "local-pv-d9c5cbd6" deleted

24.  Identify the crashcollector pod deployment:

      $ oc get deployment --selector=app=rook-ceph-
      crashcollector,node_name=_<failed_node_name>_ -n openshift-storage

     If there is an existing crashcollector pod deployment, delete it:

      $ oc delete deployment --selector=app=rook-ceph-
      crashcollector,node_name=_<failed_node_name>_ -n openshift-storage


84

               CHAPTER 3. REPLACING STORAGE NODES DEPLOYED USING LOCAL STORAGE DEVICES

25.    Delete the ocs-osd-removal-job:

           # oc delete -n openshift-storage job ocs-osd-removal-job

       Example output:

           job.batch "ocs-osd-removal-job" deleted

   Verification

 1.    Verify that the new node is present in the output:

           $ oc get nodes --show-labels | grep cluster.ocs.openshift.io/openshift-storage= |cut -d' ' -f1

 2.    Click Workloads → Pods. Confirm that at least the following pods on the new node are in
       Running state:

            openshift-storage.cephfs.csi.ceph.com-*

            openshift-storage.rbd.csi.ceph.com-*

 3.    Verify that all other required OpenShift Data Foundation pods are in Running state.
       Ensure that the new incremental mon is created, and is in the Running state:

           $ oc get pod -n openshift-storage | grep mon

       Example output:

           rook-ceph-mon-a-cd575c89b-b6k66        2/2    Running
           0      38m
           rook-ceph-mon-b-6776bc469b-tzzt8       2/2    Running
           0      38m
           rook-ceph-mon-d-5ff5d488b5-7v8xh       2/2    Running
           0   4m8s

       OSD and monitor pod might take several minutes to get to the Running state.

   4.  Verify that new OSD pods are running on the replacement node:

           $ oc get pods -o wide -n openshift-storage| egrep -i <new_node_name> | egrep osd

 5.    Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
       are encrypted.
       For each of the new nodes identified in the previous step, do the following:

        a.  Create a debug pod and open a chroot environment for the one or more selected hosts:

               $ oc debug node/<node_name>

               $ chroot /host

        b.  Display the list of available block devices:


                                                         85

Red Hat OpenShift Data Foundation 4.22 Replacing nodes

$ lsblk

Check for the crypt keyword beside the one or more    ocs-deviceset names.

6. If the verification steps fail, contact Red Hat Support .










86
