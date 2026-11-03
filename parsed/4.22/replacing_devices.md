---
title: "replacing devices"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/replacing_devices"
source_file: "replacing_devices.pdf"
version: "4.22"
chapter: "replacing_devices"
---

               Red Hat OpenShift Data Foundation
                              4.22

    Replacing devices

Instructions for safely replacing operational or failed devices










    Last Updated: 2026-11-02



Red Hat OpenShift Data Foundation 4.22 Replacing devices

Instructions for safely replacing operational or failed devices

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

    This document explains how to safely replace storage devices for Red Hat OpenShift Data
    Foundation.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. REPLACING STORAGE DEVICES OVERVIEW 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON AWS 6
  2.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON AWS USER-PROVISIONED
  INFRASTRUCTURE 6
  2.2. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON AWS INSTALLER-PROVISIONED
  INFRASTRUCTURE 6

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON VMWARE 7
  3.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON VMWARE INFRASTRUCTURE 7

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON MICROSOFT AZURE 14
  4.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON AZURE INSTALLER-PROVISIONED
  INFRASTRUCTURE 14

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON GOOGLE CLOUD 15
  5.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON GOOGLE CLOUD INSTALLER-
  PROVISIONED INFRASTRUCTURE 15

CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE
. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
DEVICES 16
  6.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON CLUSTERS BACKED BY LOCAL
  STORAGE DEVICES 16
  6.2. REPLACING STORAGE DEVICES DURING A HARDWARE REFRESH 24
  6.3. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON IBM POWER 25

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 7. REMOVING OSDS USING THE OPENSHIFT DATA FOUNDATION CLI TOOL 37
  7.1. REMOVING OBJECT STORAGE DEVICES USING THE OPENSHIFT DATA FOUNDATION CLI TOOL 37










    1

Red Hat OpenShift Data Foundation 4.22 Replacing devices










2

    PREFACE

    PREFACE

This document describes how to replace operational and failed storage devices in your OpenShift Data
Foundation cluster.










3

Red Hat OpenShift Data Foundation 4.22 Replacing devices

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

        CHAPTER 1. REPLACING STORAGE DEVICES OVERVIEW

     CHAPTER 1. REPLACING STORAGE DEVICES OVERVIEW

Depending on the type of your deployment, you can choose one of the following procedures to replace
a storage device:

        NOTE

        OpenShift Data Foundation does not support heterogeneous OSD sizes.

        Replacing operational or failed storage devices on AWS user-provisioned infrastructure

        Replacing operational or failed storage devices on AWS installer-provisioned infrastructure

        Replacing operational or failed storage devices on VMware infrastructure

        Replacing operational or failed storage devices on Azure installer-provisioned infrastructure

        Replacing operational or failed storage devices on clusters backed by local storage devices

        Replacing operational or failed storage devices on IBM Power

Additional resources

        Section 2.1, “Replacing operational or failed storage devices on AWS user-provisioned
        infrastructure”

        Section 2.2, “Replacing operational or failed storage devices on AWS installer-provisioned
        infrastructure”

        Section 3.1, “Replacing operational or failed storage devices on VMware infrastructure”

        Section 4.1, “Replacing operational or failed storage devices on Azure installer-provisioned
        infrastructure”

        Section 6.1, “Replacing operational or failed storage devices on clusters backed by local storage
        devices”

        Section 6.3, “Replacing operational or failed storage devices on IBM Power”










    5

Red Hat OpenShift Data Foundation 4.22 Replacing devices

       CHAPTER 2. REPLACING DYNAMICALLY PROVISIONED
        STORAGE DEVICES ON AWS

  2.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON
  AWS USER-PROVISIONED INFRASTRUCTURE

  When you need to replace a device in a dynamically created storage cluster on an AWS user-provisioned
  infrastructure, you must replace the storage node.

  Additional resources

        Replacing an operational AWS node on user-provisioned infrastructure

        Replacing a failed AWS node on user-provisioned infrastructure

  2.2. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON
  AWS INSTALLER-PROVISIONED INFRASTRUCTURE

  When you need to replace a device in a dynamically created storage cluster on an AWS installer-
  provisioned infrastructure, you must replace the storage node.

  Additional resources

        Replacing an operational AWS node on installer-provisioned infrastructure

        Replacing a failed AWS node on installer-provisioned infrastructure










    6

             CHAPTER 3. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON VMWARE

 CHAPTER 3. REPLACING DYNAMICALLY PROVISIONED
                      STORAGE DEVICES ON VMWARE

3.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON
VMWARE INFRASTRUCTURE

Create a new Persistent Volume Claim (PVC) on a new volume, and remove the old object storage
device (OSD) when one or more virtual machine disks (VMDK) needs to be replaced in OpenShift Data
Foundation which is deployed dynamically on VMware infrastructure.

Prerequisites

      Ensure that the data is resilient.

          In the OpenShift Web Console, click Storage → Data Foundation.

          Click the Storage Systems tab, and then click ocs-storagecluster.

          In the Status card of Block and File dashboard, under the Overview tab, verify that Data
          Resiliency has a green tick mark.

Procedure

 1.   Identify the OSD that needs to be replaced and the OpenShift Container Platform node that
      has the OSD scheduled on it.

         $ oc get -n openshift-storage pods -l app=rook-ceph-osd -o wide

      Example output:

    rook-ceph-osd-0-6d77d6c7c6-m8xj6 0/1 CrashLoopBackOff 0 24h 10.129.0.16
         compute-2    <none>     <none>
         rook-ceph-osd-1-85d99fb95f-2svc7  1/1 Running     0         24h     10.128.2.24  compute-
         0   <none>    <none>
         rook-ceph-osd-2-6c66cdb977-jp542  1/1 Running     0     24h     10.130.0.18     compute-
         1   <none>    <none>

      In this example, rook-ceph-osd-0-6d77d6c7c6-m8xj6 needs to be replaced and      compute-2 is
      the OpenShift Container platform node on which the OSD is scheduled.

                     NOTE

                     The status of the pod is Running, if the OSD you want to replace is healthy.

 2.   Scale down the rook-ceph-operator and ocs-operator deployments.

         $ oc -n openshift-storage scale deployment rook-ceph-operator --replicas=0

         $ oc -n openshift-storage scale deployment ocs-operator --replicas=0

 3.   Scale down the OSD deployment for the OSD to be replaced.

      Each time you want to replace the OSD, update the osd_id_to_remove parameter with the
                          7

  Red Hat OpenShift Data Foundation 4.22 Replacing devices

    Each time you want to replace the OSD, update the osd_id_to_remove parameter with the
    OSD ID, and repeat this step.

       $ osd_id_to_remove=0

       $ oc scale -n openshift-storage deployment rook-ceph-osd-${osd_id_to_remove} --replicas=0

    where, osd_id_to_remove is the integer in the pod name immediately after the rook-ceph-osd
    prefix. In this example, the deployment name is rook-ceph-osd-0.

    Example output:

       deployment.extensions/rook-ceph-osd-0 scaled

4.  Verify that the rook-ceph-osd pod is terminated.

       $ oc get -n openshift-storage pods -l ceph-osd-id=${osd_id_to_remove}

    Example output:

       No resources found.

         IMPORTANT

         If the rook-ceph-osd pod is in terminating state, use the force option to delete
         the pod.

                   $ oc delete pod rook-ceph-osd-0-6d77d6c7c6-m8xj6 --force --grace-period=0

         Example output:

                   warning: Immediate deletion does not wait for confirmation that the running
                   resource has been terminated. The resource may continue to run on the
                   cluster indefinitely.
                    pod "rook-ceph-osd-0-6d77d6c7c6-m8xj6" force deleted

5. Remove the old OSD from the cluster so that you can add a new OSD.

    a.  Delete any old ocs-osd-removal jobs.

        $ oc delete -n openshift-storage job ocs-osd-removal-job

        Example output:

        job.batch "ocs-osd-removal-job" deleted

                    NOTE

                    If the above job does not reach Completed state after 10 minutes, then the
                    job must be deleted and rerun with FORCE_OSD_REMOVAL=true.


  8

     CHAPTER 3. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON VMWARE

b.  Navigate to the openshift-storage project.

     $ oc project openshift-storage

c.  Remove the old OSD from the cluster.

     $ oc process -n openshift-storage ocs-osd-removal -p
     FAILED_OSD_IDS=${osd_id_to_remove} FORCE_OSD_REMOVAL=false |oc create -n
     openshift-storage -f -

    Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
    insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
    Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.


     WARNING
Proceeding will permanently remove the OSD from the cluster and
destroy its data. Ensure osd_id_to_remove is correct. Do not process
multiple IDs from the same failure domain at once. Contact Red Hat
Technical Support if you need assistance.




6.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
    job pod.
    A status of Completed confirms that the OSD removal job succeeded.

     $ oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

7.  Ensure that the OSD removal is completed.

     $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
     removal'

    Example output:

     2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

             IMPORTANT

             If the ocs-osd-removal-job pod fails and the pod is not in the expected
             Completed state, check the pod logs for further debugging.

             For example:

                 # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1




8. If encryption was enabled at the time of install, remove dm-crypt managed device-mapper
                 9

Red Hat OpenShift Data Foundation 4.22 Replacing devices

8.  If encryption was enabled at the time of install, remove dm-crypt managed device-mapper
    mapping from the OSD devices that are removed from the respective OpenShift Data
    Foundation nodes.

     a.  Get the PVC name(s) of the replaced OSD(s) from the logs of ocs-osd-removal-job pod.

             $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 |egrep -i
             ‘pvc|deviceset’

         Example output:

             2021-05-12 14:31:34.666000 I | cephosd: removing the OSD PVC "ocs-deviceset-xxxx-
             xxx-xxx-xxx"

     b.  For each of the previously identified nodes, do the following:

         i.   Create a debug pod and chroot to the host on the storage node.

              $ oc debug node/<node name>

              <node name>
              Is the name of the node.

              $ chroot /host


         ii.   Find a relevant device name based on the PVC names identified in the previous step.

                $ dmsetup ls| grep <pvc name>

               <pvc name>
                Is the name of the PVC.
                Example output:

                ocs-deviceset-xxx-xxx-xxx-xxx-block-dmcrypt (253:0)


         iii. Remove the mapped device.

                $ cryptsetup luksClose --debug --verbose ocs-deviceset-xxx-xxx-xxx-xxx-block-
                dmcrypt










                IMPORTANT
10

           CHAPTER 3. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON VMWARE

                         IMPORTANT

                         If the above command gets stuck due to insufficient privileges, run the
                         following commands:

                          Press CTRL+Z to exit the above command.

                          Find the PID of the process which was stuck.

                          $ ps -ef | grep crypt

                          Terminate the process using the kill command.

                          $ kill -9 <PID>

                          <PID>
                          Is the process ID.

                          Verify that the device name is removed.

                          $ dmsetup ls

   9.     Delete the ocs-osd-removal job.

           $ oc delete -n openshift-storage job ocs-osd-removal-job

          Example output:

           job.batch "ocs-osd-removal-job" deleted

            NOTE

            When using an external key management system (KMS) with data encryption, the
            old OSD encryption key can be removed from the Vault server as it is now an
            orphan key.

  10.     Scale up the rook-ceph-operator and ocs-operator deployments.

           $ oc -n openshift-storage scale deployment rook-ceph-operator --replicas=1

           $ oc -n openshift-storage scale deployment ocs-operator --replicas=1

Verification

     1.   Verify that there is a new OSD running.

           $ oc get -n openshift-storage pods -l app=rook-ceph-osd

          Example output:




                          11

Red Hat OpenShift Data Foundation 4.22 Replacing devices

       rook-ceph-osd-0-5f7f4747d4-snshw                         1/1  Running   0  4m47s
       rook-ceph-osd-1-85d99fb95f-2svc7        1/1     Running               0    1d20h
       rook-ceph-osd-2-6c66cdb977-jp542                         1/1  Running   0  1d20h

2.  Verify that there is a new PVC created which is in Bound state.

       $ oc get -n openshift-storage pvc

    Example output:

       NAME        STATUS     VOLUME                                  CAPACITY    ACCESS
    MODES STORAGECLASS AGE
    ocs-deviceset-0-0-2s6w4 Bound pvc-7c9bcaf7-de68-40e1-95f9-0b0d7c0ae2fc 512Gi
       RWO        thin   5m
    ocs-deviceset-1-0-q8fwh Bound pvc-9e7e00cb-6b33-402e-9dc5-b8df4fd9010f 512Gi
       RWO        thin   1d20h
    ocs-deviceset-2-0-9v8lq Bound pvc-38cdfcee-ea7e-42a5-a6e1-aaa6d4924291 512Gi
       RWO        thin   1d20h

3.  Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
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

               $ oc debug node/<node name>

             <node name>
               Is the name of the node.

               $ chroot /host


        ii. Check for the crypt keyword beside the ocs-deviceset name(s).

12

        CHAPTER 3. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON VMWARE

        $ lsblk

4. Log in to OpenShift Web Console and view the storage dashboard.










    13

Red Hat OpenShift Data Foundation 4.22 Replacing devices

       CHAPTER 4. REPLACING DYNAMICALLY PROVISIONED
        STORAGE DEVICES ON MICROSOFT AZURE

  4.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON
  AZURE INSTALLER-PROVISIONED INFRASTRUCTURE

  When you need to replace a device in a dynamically created storage cluster on an Azure installer-
  provisioned infrastructure, you must replace the storage node.

  Additional resources

        Replacing operational nodes on Azure installer-provisioned infrastructure

        Replacing failed nodes on Azure installer-provisioned infrastructures










    14

        CHAPTER 5. REPLACING DYNAMICALLY PROVISIONED STORAGE DEVICES ON GOOGLE CLOUD

    CHAPTER 5. REPLACING DYNAMICALLY PROVISIONED
        STORAGE DEVICES ON GOOGLE CLOUD

5.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON
GOOGLE CLOUD INSTALLER-PROVISIONED INFRASTRUCTURE

When you need to replace a device in a dynamically created storage cluster on an Google Cloud
installer-provisioned infrastructure, you must replace the storage node.

Additional resources

      Replacing operational nodes on Google Cloud installer-provisioned infrastructure

      Replacing failed nodes on Google Cloud installer-provisioned infrastructures










    15

Red Hat OpenShift Data Foundation 4.22 Replacing devices

                            CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS
                                  DEPLOYED USING LOCAL STORAGE DEVICES

6.1. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON
CLUSTERS BACKED BY LOCAL STORAGE DEVICES

You can replace an object storage device (OSD) in OpenShift Data Foundation deployed using local
storage devices on the following infrastructures:

     Bare metal / Agnostic deployment

     VMware

     IBM Z and IBM® LinuxONE

             NOTE

             There might be a need to replace one or more underlying storage devices.

Prerequisites

     Red Hat recommends that replacement devices are configured with similar infrastructure and
     resources to the device being replaced.

     Ensure that the data is resilient.

          In the OpenShift Web Console, click Storage → Data Foundation.

          Click the Storage Systems tab, and then click ocs-storagecluster.

          In the Status card of Block and File dashboard, under the Overview tab, verify that Data
          Resiliency has a green tick mark.

Procedure

1.   Remove the underlying storage device from relevant worker node.

2.   Verify that relevant OSD Pod has moved to CrashLoopBackOff state.
     Identify the OSD that needs to be replaced and the OpenShift Container Platform node that
     has the OSD scheduled on it.

         $ oc get -n openshift-storage pods -l app=rook-ceph-osd -o wide

     Example output:

    rook-ceph-osd-0-6d77d6c7c6-m8xj6 0/1 CrashLoopBackOff 0 24h 10.129.0.16
         compute-2   <none>     <none>
         rook-ceph-osd-1-85d99fb95f-2svc7     1/1 Running 0          24h   10.128.2.24 compute-
         0   <none>   <none>
         rook-ceph-osd-2-6c66cdb977-jp542         1/1   Running 0     24h  10.130.0.18 compute-
         1   <none>   <none>


In this example, rook-ceph-osd-0-6d77d6c7c6-m8xj6 needs to be replaced and compute-2 is
16

CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

    In this example, rook-ceph-osd-0-6d77d6c7c6-m8xj6 needs to be replaced and   compute-2 is
    the OpenShift Container platform node on which the OSD is scheduled.

3.  Scale down the rook-ceph-operator and ocs-operator deployments.

       $ oc -n openshift-storage scale deployment rook-ceph-operator --replicas=0

       $ oc -n openshift-storage scale deployment ocs-operator --replicas=0

4.  Scale down the OSD deployment for the OSD to be replaced.

       $ osd_id_to_remove=0

       $ oc scale -n openshift-storage deployment rook-ceph-osd-${osd_id_to_remove} --replicas=0

    where, osd_id_to_remove is the integer in the pod name immediately after the rook-ceph-osd
    prefix. In this example, the deployment name is rook-ceph-osd-0.

    Example output:

       deployment.extensions/rook-ceph-osd-0 scaled

5.  Verify that the rook-ceph-osd pod is terminated.

       $ oc get -n openshift-storage pods -l ceph-osd-id=${osd_id_to_remove}

    Example output:

       No resources found in openshift-storage namespace.

        IMPORTANT

        If the rook-ceph-osd pod is in terminating state for more than a few minutes,
        use the force option to delete the pod.

                   $ oc delete -n openshift-storage pod rook-ceph-osd-0-6d77d6c7c6-m8xj6 --
                   grace-period=0 --force

        Example output:

                   warning: Immediate deletion does not wait for confirmation that the running
                   resource has been terminated. The resource may continue to run on the
                   cluster indefinitely.
                   pod "rook-ceph-osd-0-6d77d6c7c6-m8xj6" force deleted

6. Remove the old OSD from the cluster so that you can add a new OSD.

    a. Delete any old ocs-osd-removal jobs.

       $ oc delete -n openshift-storage job ocs-osd-removal-job


                   17

Red Hat OpenShift Data Foundation 4.22 Replacing devices

    Example output:

     job.batch "ocs-osd-removal-job" deleted

      NOTE

      The above command must reach Completed state before moving to the
      next steps. This can take more than ten minutes.

b.  Navigate to the openshift-storage project.

     $ oc project openshift-storage

c.  Remove the old OSD from the cluster.

     $ oc process -n openshift-storage ocs-osd-removal -p
     FAILED_OSD_IDS=${osd_id_to_remove} FORCE_OSD_REMOVAL=false |oc create -n
     openshift-storage -f -

    Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
    insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
    Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.


     WARNING
Proceeding will permanently remove the OSD from the cluster and
destroy its data. Ensure osd_id_to_remove is correct. Do not process
multiple IDs from the same failure domain at once. Contact Red Hat
Technical Support if you need assistance.




7.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
    job pod.
    A status of Completed confirms that the OSD removal job succeeded.

     $ oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

8.  Ensure that the OSD removal is completed.

     $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
     removal'

    Example output:

     2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0




     IMPORTANT
18

CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

               IMPORTANT

               If the ocs-osd-removal-job fails and the pod is not in the expected Completed
               state, check the pod logs for further debugging.

               For example:

               # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

9.  If encryption was enabled at the time of install, remove dm-crypt managed device-mapper
    mapping from the OSDs that are removed from the respective OpenShift Data Foundation
    nodes.

     a.  Get the Persistent Volume Claim (PVC) names of the replaced OSDs from the logs of ocs-
         osd-removal-job pod.

             $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 |egrep -i
             'pvc|deviceset'

         Example output:

             2021-05-12 14:31:34.666000 I | cephosd: removing the OSD PVC "ocs-deviceset-xxxx-
             xxx-xxx-xxx"

     b.  For each of the previously identified nodes, do the following:

         i.   Create a debug pod and chroot to the host on the storage node.

              $ oc debug node/<node name>

              <node name>
              Is the name of the node.

               $ chroot /host


         ii.   Find the relevant device name based on the PVC names identified in the previous step.

                $ dmsetup ls| grep <pvc name>

               <pvc name>
                Is the name of the PVC.
                Example output:

                ocs-deviceset-xxx-xxx-xxx-xxx-block-dmcrypt (253:0)


         iii.   Remove the mapped device.

                $ cryptsetup luksClose --debug --verbose <ocs-deviceset-name>

                <ocs-deviceset-name>

                Is the name of the relevant device based on the PVC names identified in the
                    19

Red Hat OpenShift Data Foundation 4.22 Replacing devices

    Is the name of the relevant device based on the PVC names identified in the
    previous step.

                  IMPORTANT

                  If the above command gets stuck due to insufficient privileges, run
                  the following commands:

                   Press CTRL+Z to exit the above command.

                   Find the PID of the process which was stuck.

                   $ ps -ef | grep crypt

                   Terminate the process using the kill command.

                   $ kill -9 <PID>

                   <PID>
                   Is the process ID.

                   Verify that the device name is removed.

                                            $ dmsetup ls



 10.   Find the persistent volume (PV) that need to be deleted.

        $ oc get pv -L kubernetes.io/hostname | grep <storageclass-name> | grep Released

       Example output:

        local-pv-d6bf175b 1490Gi     RWO     Delete     Released     openshift-
        storage/ocs-deviceset-0-data-0-6c5pw   localblock     2d22h compute-1

 11.   Delete the PV.

        $ oc delete pv <pv_name>

 12.   Physically add a new device to the node.

 13.   Track the provisioning of PVs for the devices that match the deviceInclusionSpec. It can take a
       few minutes to provision the PVs.

        $ oc -n openshift-local-storage describe localvolumeset <lvs-name>

       Example output:

        [...]
        Status:
        Conditions:
        Last Transition Time:   2020-11-17T05:03:32Z


    20

                                CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

          Message:                DiskMaker: Available, LocalProvisioner: Available
          Status:               True
          Type:                 DaemonSetsAvailable
          Last Transition Time:     2020-11-17T05:03:34Z
          Message:                Operator reconciled successfully.
          Status:               True
          Type:                 Available
         Observed Generation:    1
         Total Provisioned Device Count: 4
        Events:
        Type  Reason     Age    From       Message
        ----  ------  ----  ----          -------
    Normal Discovered 2m30s (x4 localvolumeset- node.example.com -
    NewDevice over 2m30s) symlink-controller found possible
    matching disk,
    waiting 1m to claim

    Normal FoundMatch 89s (x4 localvolumeset- node.example.com -
    ingDisk over 89s) symlink-controller symlinking matching
    disk

     Once the PV is provisioned, a new OSD pod is automatically created for the PV.

14.  Delete the ocs-osd-removal job(s).

        $ oc delete -n openshift-storage job ocs-osd-removal-job

     Example output:

        job.batch "ocs-osd-removal-job" deleted

                     NOTE

                     When using an external key management system (KMS) with data encryption, the
                     old OSD encryption key can be removed from the Vault server as it is now an
                     orphan key.

15.  Scale up the rook-ceph-operator and ocs-operator deployments.

        $ oc -n openshift-storage scale deployment rook-ceph-operator --replicas=1

        $ oc -n openshift-storage scale deployment ocs-operator --replicas=1

16.  [Optional] Verify the Ceph cluster health and check if there are any crash reports associated
     with the previous failed OSD which has now been replaced.

     a.  List all crash reports:

             $ oc exec -it $(oc get pod -n openshift-storage -l app=rook-ceph-operator -o name) -n
             openshift-storage -- ceph crash ls -c /var/lib/rook/openshift-storage/openshift-
             storage.config

         Sample output:


                                       21

    Red Hat OpenShift Data Foundation 4.22 Replacing devices

                ID                          ENTITY NEW
                2025-09-25T07:12:45Z_1234567890        osd.5  *
                2025-09-25T08:20:11Z_abcdef1234        mon.a  *

           b.  Filter crashes for the failed OSD from the previous step. For example, if the failed OSD is
               osd.5, run:

                $ oc exec -it $(oc get pod -n openshift-storage -l app=rook-ceph-operator -o name) -n
                openshift-storage -- ceph crash ls -c /var/lib/rook/openshift-storage/openshift-
                storage.config | grep osd.5

               Note the specif crash report from the output.

           c.  Archive the specific crash. Use the crash ID from the previous step:

                $ oc exec -it $(oc get pod -n openshift-storage -l app=rook-ceph-operator -o name) -n
                openshift-storage -- ceph crash archive <crash-id> -c /var/lib/rook/openshift-
                storage/openshift-storage.config

               Replace crash-id with the crash report.

           d.  Verify the crash has been archived:

                $ oc exec -it $(oc get pod -n openshift-storage -l app=rook-ceph-operator -o name) -n
                openshift-storage -- ceph crash ls -c /var/lib/rook/openshift-storage/openshift-
                storage.config

Verification

     1.   Verify that there is a new OSD running.

              $ oc get -n openshift-storage pods -l app=rook-ceph-osd

          Example output:

              rook-ceph-osd-0-5f7f4747d4-snshw       1/1     Running   0  4m47s
              rook-ceph-osd-1-85d99fb95f-2svc7     1/1      Running   0   1d20h
              rook-ceph-osd-2-6c66cdb977-jp542       1/1     Running   0  1d20h

                      IMPORTANT

                      If the new OSD does not show as Running after a few minutes, restart the  rook-
                      ceph-operator pod to force a reconciliation.

                          $ oc delete pod -n openshift-storage -l app=rook-ceph-operator

                      Example output:

                          pod "rook-ceph-operator-6f74fb5bff-2d982" deleted

    2. Verify that a new PVC is created.


    22

    CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

           $ oc get -n openshift-storage pvc | grep <lvs-name>

        Example output:

           ocs-deviceset-0-0-c2mqb  Bound     local-pv-b481410    1490Gi   RWO   localblock
           5m
           ocs-deviceset-1-0-959rp  Bound    local-pv-414755e0    1490Gi   RWO   localblock
           1d20h
           ocs-deviceset-2-0-79j94  Bound    local-pv-3e8964d3    1490Gi   RWO   localblock
           1d20h

    3.  Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSDs are
        encrypted.

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

                  $ oc debug node/<node name>

                 <node name>
                  Is the name of the node.

                  $ chroot /host


        ii. Check for the crypt keyword beside the ocs-deviceset name(s).

        $ lsblk

4. Log in to OpenShift Web Console and check the OSD status on the storage dashboard.




            NOTE
                                                                                 23

Red Hat OpenShift Data Foundation 4.22 Replacing devices

             NOTE

             A full data recovery may take longer depending on the volume of data being recovered.

6.2. REPLACING STORAGE DEVICES DURING A HARDWARE REFRESH

When replacing a batch of smaller disks with larger disks as part of a hardware refresh, use the following
procedure. This approach differs from single-device replacement because it manages rebalancing
across multiple disks simultaneously, which reduces total migration time and cluster disruption.

Prerequisites

      Download the OpenShift Data Foundation command-line interface (CLI) tool. You can find a
      compatible version and download the CLI tool from the customer portal.

      Ensure the new replacement disks are available and accessible on the nodes.

Procedure

 1.   Pause rebalancing (optional, recommended for batch replacements).
      Use the OpenShift Data Foundation CLI to set nobackfill to prevent the cluster from
      rebalancing after each individual disk is added:

                 $ odf ceph osd set nobackfill

 2.   Add the new, larger capacity disks.
      Use the OpenShift Web Console to add the replacement disks:

       a.  Navigate to Storage → Data Foundation → Storage Systems.

       b.  Click the kebab menu (⋮) next to your storage system and select Add Capacity.

       c.  Select the appropriate Local Storage storage class and the number of new replacement
           disks.
           This creates the necessary persistent volumes (PVs) and persistent volume claims (PVCs)
           for the new OSDs automatically.

 3.   Remove the old, smaller disks.
      Use the odf purge-osd command to safely decommission each old OSD. This command checks
      whether it is safe to remove the OSD before draining and purging it:

                    $ odf purge-osd <OSD_ID>

      Replace <OSD_ID> with the ID of the OSD to remove. The OSD ID is the integer in the pod
      name immediately after the rook-ceph-osd prefix. For example, in the pod name rook-ceph-
      osd-0-6d77d6c7c6-m8xj6, the OSD ID is 0. Repeat this step for each old OSD being replaced.

                 NOTE

                 If removal of the OSD affects placement group (PG) status, you are prompted to
                 enter yes-force-destroy-osd to confirm.

 4.   Unset nobackfill and monitor recovery.


24

CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

        After all new OSDs are active and all old OSDs are purged, re-enable rebalancing:

        $ odf ceph osd unset nobackfill

        Monitor the data resiliency status from the Data Resiliency card on the OpenShift Data
        Foundation dashboard in the OpenShift Web Console. The card displays a progress indicator for
        the ongoing rebalancing operation.

    5.  (Optional) Adjust the recovery profile if migration is too slow.
        If the rebalancing rate is insufficient, use the OpenShift Data Foundation CLI to increase the
        recovery speed:

        $ odf set recovery-profile high_recovery_ops

                 NOTE

                 Setting a high recovery profile increases recovery throughput but may affect
                 cluster performance for running workloads. Reset the profile to the default when
                 recovery is complete.

6.3. REPLACING OPERATIONAL OR FAILED STORAGE DEVICES ON
IBM POWER

You can replace an object storage device (OSD) in OpenShift Data Foundation deployed using local
storage devices on IBM Power.

             NOTE

             There might be a need to replace one or more underlying storage devices.

Prerequisites

        Red Hat recommends that replacement devices are configured with similar infrastructure and
        resources to the device being replaced.

        Ensure that the data is resilient.

         In the OpenShift Web Console, click Storage → Data Foundation.

         Click the Storage Systems tab, and then click ocs-storagecluster.

         In the Status card of Block and File dashboard, under the Overview tab, verify that Data
         Resiliency has a green tick mark.

Procedure

1.      Identify the OSD that needs to be replaced and the OpenShift Container Platform node that
        has the OSD scheduled on it.

        $ oc get -n openshift-storage pods -l app=rook-ceph-osd -o wide

        Example output:



                 25

Red Hat OpenShift Data Foundation 4.22 Replacing devices

    rook-ceph-osd-0-86bf8cdc8-4nb5t 0/1 crashLoopBackOff 0 24h 10.129.2.26
    worker-0 <none> <none>
     rook-ceph-osd-1-7c99657cfb-jdzvz 1/1    Running    0    24h        10.128.2.46 worker-1
     <none>  <none>
     rook-ceph-osd-2-5f9f6dfb5b-2mnw9 1/1     Running     0         24h  10.131.0.33 worker-2
     <none>  <none>

    In this example, rook-ceph-osd-0-86bf8cdc8-4nb5t needs to be replaced and  worker-0 is the
    RHOCP node on which the OSD is scheduled.

             NOTE

             The status of the pod is Running if the OSD you want to replace is healthy.

2.  Scale down the rook-ceph-operator and ocs-operator deployments.

     $ oc -n openshift-storage scale deployment rook-ceph-operator --replicas=0

     $ oc -n openshift-storage scale deployment ocs-operator --replicas=0

3.  Scale down the OSD deployment for the OSD to be replaced.

     $ osd_id_to_remove=0

     $ oc scale -n openshift-storage deployment rook-ceph-osd-${osd_id_to_remove} --replicas=0

    where, osd_id_to_remove is the integer in the pod name immediately after the rook-ceph-osd
    prefix. In this example, the deployment name is rook-ceph-osd-0.

    Example output:

     deployment.extensions/rook-ceph-osd-0 scaled

4.  Verify that the rook-ceph-osd pod is terminated.

     $ oc get -n openshift-storage pods -l ceph-osd-id=${osd_id_to_remove}

    Example output:

     No resources found in openshift-storage namespace.










             IMPORTANT
26

CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

                IMPORTANT

                If the rook-ceph-osd pod is in terminating state for more than a few minutes,
                use the force option to delete the pod.

                    $ oc delete -n openshift-storage pod rook-ceph-osd-0-86bf8cdc8-4nb5t --
                    grace-period=0 --force

                Example output:

                    warning: Immediate deletion does not wait for confirmation that the running
                    resource has been terminated. The resource may continue to run on the
                    cluster indefinitely.
                    pod "rook-ceph-osd-0-86bf8cdc8-4nb5t" force deleted

5. Remove the old OSD from the cluster so that you can add a new OSD.

   a.  Identify the DeviceSet associated with the OSD to be replaced.

        $ oc get -n openshift-storage -o yaml deployment rook-ceph-osd-${osd_id_to_remove} |
        grep ceph.rook.io/pvc

       Example output:

        ceph.rook.io/pvc: ocs-deviceset-localblock-0-data-0-64xjl
        ceph.rook.io/pvc: ocs-deviceset-localblock-0-data-0-64xjl

       In this example, the Persistent Volume Claim (PVC) name is ocs-deviceset-localblock-0-
       data-0-64xjl.

   b.  Identify the Persistent Volume (PV) associated with the PVC.

        $ oc get -n openshift-storage pvc ocs-deviceset-<x>-<y>-<pvc-suffix>

       where, x, y, and pvc-suffix are the values in the DeviceSet identified in an earlier step.

       Example output:

        NAME          STATUS     VOLUME     CAPACITY     ACCESS MODES
    STORAGECLASS AGE
    ocs-deviceset-localblock-0-data-0-64xjl Bound local-pv-8137c873 256Gi RWO
    localblock 24h

       In this example, the associated PV is local-pv-8137c873.

   c.  Identify the name of the device to be replaced.

        $ oc get pv local-pv-<pv-suffix> -o yaml | grep path

       where, pv-suffix is the value in the PV name identified in an earlier step.

       Example output:



                      27

Red Hat OpenShift Data Foundation 4.22 Replacing devices

      path: /mnt/local-storage/localblock/vdc

     In this example, the device name is vdc.

d.   Identify the prepare-pod associated with the OSD to be replaced.

      $ oc describe -n openshift-storage pvc ocs-deviceset-<x>-<y>-<pvc-suffix> | grep Used

     where, x, y, and pvc-suffix are the values in the  DeviceSet identified in an earlier step.

     Example output:

      Used By:   rook-ceph-osd-prepare-ocs-deviceset-localblock-0-data-0-64knzkc

     In this example, the prepare-pod name is rook-ceph-osd-prepare-ocs-deviceset-
     localblock-0-data-0-64knzkc.

e.   Delete any old ocs-osd-removal jobs.

      $ oc delete -n openshift-storage job ocs-osd-removal-job

     Example output:

      job.batch "ocs-osd-removal-job" deleted

                 NOTE

                 The above command must reach Completed state before moving to the
                 next steps. This can take more than ten minutes.

f.   Change to the openshift-storage project.

      $ oc project openshift-storage

g.   Remove the old OSD from the cluster.

      $ oc process -n openshift-storage ocs-osd-removal -p
      FAILED_OSD_IDS=${osd_id_to_remove} FORCE_OSD_REMOVAL=false |oc create -n
      openshift-storage -f -

     Set FORCE_OSD_REMOVAL to true when removing an OSD leaves the cluster with
     insufficient OSDs to satisfy its data protection requirements. Without this parameter, the
     Ceph safe-to-destroy check does not pass and the removal job hangs indefinitely.










28

CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES


 WARNING
Proceeding will permanently remove the OSD from the cluster and
destroy its data. Ensure osd_id_to_remove is correct. Do not process
multiple IDs from the same failure domain at once. Contact Red Hat
Technical Support if you need assistance.




6.  Verify that the OSD was removed successfully by checking the status of the ocs-osd-removal-
    job pod.
    A status of Completed confirms that the OSD removal job succeeded.

        $ oc get pod -l job-name=ocs-osd-removal-job -n openshift-storage

7.  Ensure that the OSD removal is completed.

        $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 | egrep -i 'completed
        removal'

    Example output:

        2022-05-10 06:50:04.501511 I | cephosd: completed removal of OSD 0

                IMPORTANT

                If the ocs-osd-removal-job fails and the pod is not in the expected Completed
                state, check the pod logs for further debugging.

                For example:

                   # oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1

8.  If encryption was enabled at the time of install, remove dm-crypt managed device-mapper
    mapping from the OSDs that are removed from the respective OpenShift Data Foundation
    nodes.

     a.  Get the PVC name(s) of the replaced OSD(s) from the logs of ocs-osd-removal-job pod.

          $ oc logs -l job-name=ocs-osd-removal-job -n openshift-storage --tail=-1 |egrep -i
          ‘pvc|deviceset’

         Example output:

          2021-05-12 14:31:34.666000 I | cephosd: removing the OSD PVC "ocs-deviceset-xxxx-
          xxx-xxx-xxx"

     b.  For each of the previously identified nodes, do the following:



                29

Red Hat OpenShift Data Foundation 4.22 Replacing devices

i.   Create a debug pod and chroot to the host on the storage node.

     $ oc debug node/<node name>

     <node name>
     Is the name of the node.

     $ chroot /host


ii.   Find the relevant device name based on the PVC names identified in the previous step.

       $ dmsetup ls| grep <pvc name>

      <pvc name>
       Is the name of the PVC.
       Example output:

       ocs-deviceset-xxx-xxx-xxx-xxx-block-dmcrypt (253:0)


iii.          Remove the mapped device.

       $ cryptsetup luksClose --debug --verbose ocs-deviceset-xxx-xxx-xxx-xxx-block-
       dmcrypt

              IMPORTANT

              If the above command gets stuck due to insufficient privileges, run the
              following commands:

               Press CTRL+Z to exit the above command.

               Find the PID of the process which was stuck.

               $ ps -ef | grep crypt

               Terminate the process using the kill command.

               $ kill -9 <PID>

               <PID>
               Is the process ID.

               Verify that the device name is removed.

               $ dmsetup ls

9. Find the PV that need to be deleted.

   $ oc get pv -L kubernetes.io/hostname | grep localblock | grep Released

30

            CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

      Example output:

         local-pv-d6bf175b     1490Gi     RWO     Delete     Released openshift-
         storage/ocs-deviceset-0-data-0-6c5pw     localblock 2d22h    compute-1

10.   Delete the PV.

         $ oc delete pv <pv-name>

      <pv-name>
         Is the name of the PV.

11.   Replace the old device and use the new device to create a new OpenShift Container Platform
      PV.

      a.  Log in to the OpenShift Container Platform node with the device to be replaced. In this
          example, the OpenShift Container Platform node is worker-0.

           $ oc debug node/worker-0

          Example output:

           Starting pod/worker-0-debug ...
           To use host binaries, run `chroot /host`
           Pod IP: 192.168.88.21
           If you don't see a command prompt, try pressing enter.
           # chroot /host

       b. Record the /dev/disk that is to be replaced using the device name, vdc, identified earlier.

           # ls -alh /mnt/local-storage/localblock

          Example output:

           total 0
           drwxr-xr-x. 2 root root 17 Nov 18 15:23 .
           drwxr-xr-x. 3 root root 24 Nov 18 15:23 ..
           lrwxrwxrwx. 1 root root 8 Nov 18 15:23 vdc -> /dev/vdc

       c. Find the name of the LocalVolume CR, and remove or comment out the device /dev/disk
       that is to be replaced.

           $ oc get -n openshift-local-storage localvolume

          Example output:

           NAME     AGE
           localblock 25h

           # oc edit -n openshift-local-storage localvolume localblock

          Example output:

                                                                                               31

Red Hat OpenShift Data Foundation 4.22 Replacing devices

             [...]
              storageClassDevices:
              - devicePaths:
             #    - /dev/vdc
                  storageClassName: localblock
                  volumeMode: Block
             [...]

           Make sure to save the changes after editing the CR.

  12.  Log in to the OpenShift Container Platform node with the device to be replaced and remove
       the old symlink.

           $ oc debug node/worker-0

       Example output:

           Starting pod/worker-0-debug ...
           To use host binaries, run `chroot /host`
           Pod IP: 192.168.88.21
           If you don't see a command prompt, try pressing enter.
           # chroot /host

        a.  Identify the old symlink for the device name to be replaced. In this example, the device
            name is vdc.

             # ls -alh /mnt/local-storage/localblock

            Example output:

             total 0
             drwxr-xr-x. 2 root root 17 Nov 18 15:23 .
             drwxr-xr-x. 3 root root 24 Nov 18 15:23 ..
             lrwxrwxrwx. 1 root root 8 Nov 18 15:23 vdc -> /dev/vdc

        b.  Remove the symlink.

             # rm /mnt/local-storage/localblock/vdc

        c.  Verify that the symlink is removed.

             # ls -alh /mnt/local-storage/localblock

            Example output:

             total 0
             drwxr-xr-x. 2 root root 6 Nov 18 17:11 .
             drwxr-xr-x. 3 root root 24 Nov 18 15:23 ..

  13.  Replace the old device with the new device.

  14.  Log back into the correct OpenShift Container Platform node and identify the device name for
       the new drive. The device name must change unless you are resetting the same device.

32

CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

        # lsblk

     Example output:

        NAME                           MAJ:MIN RM SIZE RO TYPE MOUNTPOINT
        vda                   252:0        0 40G 0 disk
        |-vda1                 252:1   0      4M 0 part
        |-vda2                 252:2   0 384M 0 part /boot
        `-vda4                 252:4   0 39.6G 0 part
         `-coreos-luks-root-nocrypt 253:0   0 39.6G 0 dm  /sysroot
        vdb                   252:16   0 512B 1 disk
        vdd                   252:32   0 256G 0 disk

     In this example, the new device name is vdd.

15.  After the new /dev/disk is available, you can add a new disk entry to the LocalVolume CR.

     a.  Edit the LocalVolume CR and add the new /dev/disk.
         In this example, the new device is /dev/vdd.

               # oc edit -n openshift-local-storage localvolume localblock

         Example output:

               [...]
                  storageClassDevices:
                  - devicePaths:
                  # - /dev/vdc
                    - /dev/vdd
                    storageClassName: localblock
                    volumeMode: Block
               [...]

         Make sure to save the changes after editing the CR.

16.  Verify that there is a new PV in Available state and of the correct size.

        $ oc get pv | grep 256Gi

     Example output:

    local-pv-1e31f771 256Gi RWO Delete Bound openshift-storage/ocs-deviceset-
    localblock-2-data-0-6xhkf localblock 24h
    local-pv-ec7f2b80 256Gi RWO Delete Bound openshift-storage/ocs-deviceset-
    localblock-1-data-0-hr2fx localblock 24h
    local-pv-8137c873 256Gi RWO Delete Available
    localblock 32m

17.  Create a new OSD for the new device.
     Deploy the new OSD. You need to restart the rook-ceph-operator to force operator
     reconciliation.

     a.  Identify the name of the rook-ceph-operator.


                                                33

Red Hat OpenShift Data Foundation 4.22 Replacing devices

            $ oc get -n openshift-storage pod -l app=rook-ceph-operator

          Example output:

            NAME        READY     STATUS     RESTARTS     AGE
            rook-ceph-operator-85f6494db4-sg62v 1/1     Running  0     1d20h

      b.  Delete the rook-ceph-operator.

            $ oc delete -n openshift-storage pod rook-ceph-operator-85f6494db4-sg62v

          Example output:

            pod "rook-ceph-operator-85f6494db4-sg62v" deleted

          In this example, the rook-ceph-operator pod name is rook-ceph-operator-85f6494db4-
          sg62v.

      c.  Verify that the rook-ceph-operator pod is restarted.

            $ oc get -n openshift-storage pod -l app=rook-ceph-operator

          Example output:

            NAME        READY                  STATUS   RESTARTS     AGE
            rook-ceph-operator-85f6494db4-wx9xx     1/1 Running  0     50s

          Creation of the new OSD may take several minutes after the operator restarts.

18.  Delete the ocs-osd-removal job(s).

         $ oc delete -n openshift-storage job ocs-osd-removal-job

     Example output:

         job.batch "ocs-osd-removal-job" deleted

                    NOTE

                    When using an external key management system (KMS) with data encryption, the
                    old OSD encryption key can be removed from the Vault server as it is now an
                    orphan key.

19.  Scale up the rook-ceph-operator and ocs-operator deployments.

         $ oc -n openshift-storage scale deployment rook-ceph-operator --replicas=1

         $ oc -n openshift-storage scale deployment ocs-operator --replicas=1

Verification


34

 CHAPTER 6. REPLACING STORAGE DEVICES ON CLUSTERS DEPLOYED USING LOCAL STORAGE DEVICES

1.   Verify that there is a new OSD running.

        $ oc get -n openshift-storage pods -l app=rook-ceph-osd

     Example output:

        rook-ceph-osd-0-76d8fb97f9-mn8qz    1/1    Running  0   23m
        rook-ceph-osd-1-7c99657cfb-jdzvz    1/1   Running  1   25h
        rook-ceph-osd-2-5f9f6dfb5b-2mnw9    1/1    Running  0   25h

 2.  Verify that a new PVC is created.

        $ oc get -n openshift-storage pvc | grep localblock

     Example output:

     ocs-deviceset-localblock-0-data-0-q4q6b Bound local-pv-8137c873 256Gi RWO
        localblock  10m
     ocs-deviceset-localblock-1-data-0-hr2fx Bound local-pv-ec7f2b80 256Gi RWO
        localblock  1d20h
     ocs-deviceset-localblock-2-data-0-6xhkf Bound local-pv-1e31f771 256Gi RWO
        localblock  1d20h

3.   Optional: If cluster-wide encryption is enabled on the cluster, verify that the new OSD devices
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


     b. For each of the previously identified nodes, do the following:

         i.   Create a debug pod and open a chroot environment for the selected host(s).

              $ oc debug node/<node name>

              <node name>
              Is the name of the node.


                                                                35

Red Hat OpenShift Data Foundation 4.22 Replacing devices

$ chroot /host


    ii. Check for the crypt keyword beside the ocs-deviceset name(s).

        $ lsblk

4.  Log in to OpenShift Web Console and check the status card in the OpenShift Data Foundation
    dashboard under Storage section.

        NOTE

        A full data recovery may take longer depending on the volume of data being
        recovered.










36

             CHAPTER 7. REMOVING OSDS USING THE OPENSHIFT DATA FOUNDATION CLI TOOL

         CHAPTER 7. REMOVING OSDS USING THE OPENSHIFT DATA
                        FOUNDATION CLI TOOL

7.1. REMOVING OBJECT STORAGE DEVICES USING THE OPENSHIFT
DATA FOUNDATION CLI TOOL

You can use the OpenShift Data Foundation command line interface (CLI) tool to automate the
process of object storage device (OSD) removal. This helps to avoid the possible data loss while
removing OSDs.

Prerequisites

      Download the OpenShift Data Foundation command line interface (CLI) tool. With the Data
      Foundation CLI tool, you can effectively manage and troubleshoot your Data Foundation
      environment from a terminal. You can find a compatible version and download the CLI tool from
      the customer portal.

Procedure

 1.   Identify the OSD that needs to be removed. Use this tool only if the OSD is in
      CrashLoopBackOff or Error state.

         $ oc get -n openshift-storage pods -l app=rook-ceph-osd -o wide

      Example output:

    rook-ceph-osd-0-6d77d6c7c6-m8xj6 0/1 CrashLoopBackOff 0 24h 10.129.0.16
         compute-2 <none> <none>

 2.   Run the following command to remove OSD 0:

         $ odf purge-osd 0

 3.   [Optional] If removal of the OSD affects placement group (PG) status, enter yes-force-
      destroy-osd.

 4.   Verify that the last line of the command output contains cephosd: completed removal of OSD
      0.

 5.   Verify the corresponding deployment is removed:

         $ oc get deployment rook-ceph-osd-0










         37
