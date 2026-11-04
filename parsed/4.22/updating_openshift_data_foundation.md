---
title: "updating openshift data foundation"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/updating_openshift_data_foundation"
source_file: "updating_openshift_data_foundation.pdf"
version: "4.22"
chapter: "updating_openshift_data_foundation"
---

Red Hat OpenShift Data Foundation
        4.22

  Updating OpenShift Data Foundation

 Instructions for cluster and storage administrators regarding upgrading










    Last Updated: 2026-10-28



Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data
Foundation

Instructions for cluster and storage administrators regarding upgrading

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

    This document explains how to update previous versions of Red Hat OpenShift Data Foundation.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. OVERVIEW OF THE OPENSHIFT DATA FOUNDATION UPDATE PROCESS 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. OPENSHIFT DATA FOUNDATION UPGRADE CHANNELS AND RELEASES 8

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. UPDATING RED HAT OPENSHIFT DATA FOUNDATION 4.21 TO 4.22 9

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. UPDATING RED HAT OPENSHIFT DATA FOUNDATION 4.22.X TO 4.22.Y 13

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. CHANGING THE UPDATE APPROVAL STRATEGY 15

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 6. UPDATING THE OPENSHIFT DATA FOUNDATION EXTERNAL SECRET 16










    1

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation










2

    PREFACE

    PREFACE

This document describes how to update Red Hat OpenShift Data Foundation and its components to a
new version.










3

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

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

                                    CHAPTER 1. OVERVIEW OF THE OPENSHIFT DATA FOUNDATION UPDATE PROCESS

      CHAPTER 1. OVERVIEW OF THE OPENSHIFT DATA
              FOUNDATION UPDATE PROCESS

This chapter helps you to upgrade between the minor releases and z-streams for all Red Hat OpenShift
Data Foundation deployments (Internal, Internal-Attached and External). The upgrade process remains
the same for all deployments.

You can upgrade OpenShift Data Foundation and its components, either between minor releases like
4.21 and 4.22, or between z-stream updates like 4.22.0 and 4.22.1 by enabling automatic updates (if not
done so during operator installation) or performing manual updates. When a new z-stream release
becomes available, the upgrade process triggers automatically if the update strategy was set to
Automatic. It is recommended to take snapshots of all the volumes in use before running any
maintenance or upgrade procedures.

Extended Update Support (EUS)

EUS to EUS upgrade in OpenShift Data Foundation is sequential and it is aligned with OpenShift
upgrade. For more information, see Performing a Control Plane Only update and Control Plane Only
updates for layered products and Operators installed through Operator Lifecycle Manager in the
OpenShift Container Platform documentation.

For EUS upgrade of OpenShift Container Platform and OpenShift Data Foundation, make sure that
OpenShift Data Foundation is upgraded along with OpenShift Container Platform and compatibility
between OpenShift Data Foundation and OpenShift Container Platform is always maintained.

Example workflow of EUS upgrade:

1.   Pause the worker machine pools.

2.   Update OpenShift <4.y> → OpenShift <4.y+1>.

3.   Update OpenShift Data Foundation <4.y> → OpenShift Data Foundation <4.y+1>.

4.   Update OpenShift <4.y+1> → OpenShift <4.y+2>.

5.   Update to OpenShift Data Foundation <4.y+2>.

6.   Unpause the worker machine pools.

      NOTE

      You can update to ODF <4.y+2> either before or after worker machine pools are
      unpaused.

      IMPORTANT

      When you update OpenShift Data Foundation in external mode, make sure that the Red
      Had Ceph Storage and OpenShift Data Foundation versions are compatible. For more
      information about supported Red Had Ceph Storage version in external mode, refer to
      the Red Hat OpenShift Data Foundation Supportability and Interoperability Checker .
      Provide the required OpenShift Data Foundation version in the checker to see the
      supported Red Had Ceph version corresponding to the version in use.

You also need to upgrade the different parts of Red Hat OpenShift Data Foundation in the following
order for both internal and external mode deployments:

                                                                                                      5

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

1.   Update OpenShift Container Platform according to the Updating clusters documentation for
     OpenShift Container Platform.

2.   Update Red Hat OpenShift Data Foundation.

      a.  To prepare a disconnected environment for updates, see Using Operator Lifecycle
          Manager in disconnected environments to be able to update OpenShift Data Foundation as
          well as Local Storage Operator when in use. For the list of packages to be mirrored for
          upgrading OpenShift Data Foundation in disconnected environments, see the
          Disconnected environment section in the Planning your deployment guide.

      b.  For updating between minor releases, see Updating Red Hat OpenShift Data Foundation
          4.20 to 4.22.

      c.  For updating between z-stream releases, see Updating Red Hat OpenShift Data
          Foundation 4.22.x to 4.22.y.

      d.  For updating external mode deployments, you must also perform the steps from the
          Updating the Red Hat OpenShift Data Foundation external secret section.

      e.  If you use local storage, then update the Local Storage operator. See Checking for Local
          Storage Operator deployments if you are unsure.

           IMPORTANT

           If you have an existing setup of OpenShift Data Foundation 4.21 with disaster recovery
           (DR) enabled, ensure to update all your clusters in the environment at the same time and
           avoid updating a single cluster. This is to avoid any potential issues and maintain best
           compatibility. It is also important to maintain consistency across all OpenShift Data
           Foundation DR instances.

Update considerations

Review the following important considerations before you begin.

     The Red Hat OpenShift Container Platform version is the same as Red Hat OpenShift Data
     Foundation. See the Interoperability Matrix for more information about supported combinations
     of OpenShift Container Platform and Red Hat OpenShift Data Foundation.

     To know whether your cluster was deployed in internal or external mode, see How to determine
     if ODF cluster has storage in internal or external mode.

     The Local Storage Operator is fully supported only when the Local Storage Operator version
     matches the Red Hat OpenShift Container Platform version.

     In OpenShift Data Foundation clusters with disaster recovery (DR) enabled, during upgrade to
     version 4.22, previous bluestore-rdr OSDs are migrated to bluestore OSDs. bluestore backed
     OSDs now provide the same improved performance of bluestore-rdr based OSDs, which is
     important when the cluster is required to be used for Regional Disaster Recovery.
     During upgrade you can view the status of the OSD migration. In the OpenShift Web Console,
     navigate to Storage → Data Foundation → Storage System. In the Activity card of the Block
     and File tab you can view ongoing activities. Migrating cluster OSDsshows the status of the
     migration from bluestore-rdr to bluestore.




          IMPORTANT
6

     CHAPTER 1. OVERVIEW OF THE OPENSHIFT DATA FOUNDATION UPDATE PROCESS

 IMPORTANT

 To avoid data loss, it is recommended to take a backup of NooBaa DB PVC regularly. If
 NooBaa DB fails and cannot be recovered, then you can revert to the latest backed-up
 version. For instructions on backing up NooBaa DB, follow the steps in Perform a One-
 Time Backup of the Database for the Multicloud Object Gateway.

Additional resources

Performing a Control Plane Only update

Control Plane Only updates for layered products and Operators installed through Operator
Lifecycle Manager

Red Hat OpenShift Data Foundation Supportability and Interoperability Checker

Updating clusters

Using Operator Lifecycle Manager in disconnected environments

Disconnected environment

Updating Red Hat OpenShift Data Foundation 4.20 to 4.22

Updating Red Hat OpenShift Data Foundation 4.22.x to 4.22.y

Updating the Red Hat OpenShift Data Foundation external secret

Checking for Local Storage Operator deployments

Interoperability Matrix

How to determine if ODF cluster has storage in internal or external mode

Perform a One-Time Backup of the Database for the Multicloud Object Gateway










 7

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

       CHAPTER 2. OPENSHIFT DATA FOUNDATION UPGRADE
        CHANNELS AND RELEASES

  In OpenShift Container Platform 4.1, Red Hat introduced the concept of channels for recommending
  the appropriate release versions for cluster upgrades. By controlling the pace of upgrades, these
  upgrade channels allow you to choose an upgrade strategy. As OpenShift Data Foundation gets
  deployed as an operator in OpenShift Container Platform, it follows the same strategy to control the
  pace of upgrades by shipping the fixes in multiple channels. Upgrade channels are tied to a minor
  version of OpenShift Data Foundation.

  For example, OpenShift Data Foundation 4.22 upgrade channels recommend upgrades within 4.22.
  Upgrades to future releases is not recommended. This strategy ensures that administrators can
  explicitly decide to upgrade to the next minor version of OpenShift Data Foundation.

  Upgrade channels control only release selection and do not impact the version of the cluster that you
  install; the odf-operator decides the version of OpenShift Data Foundation to be installed. By default, it
  always installs the latest OpenShift Data Foundation release maintaining the compatibility with
  OpenShift Container Platform. So, on OpenShift Container Platform 4.22, OpenShift Data Foundation
  4.22 will be the latest version which can be installed.

  OpenShift Data Foundation upgrades are tied to the OpenShift Container Platform upgrade to ensure
  that compatibility and interoperability are maintained with the OpenShift Container Platform. For
  OpenShift Data Foundation 4.22, OpenShift Container Platform 4.22 and 4.23 (when generally
  available) are supported. OpenShift Container Platform 4.22 is supported to maintain forward
  compatibility of OpenShift Data Foundation with OpenShift Container Platform. Keep the OpenShift
  Data Foundation version the same as OpenShift Container Platform in order to get the benefit of all the
  features and enhancements in that release.

        IMPORTANT

        Due to fundamental Kubernetes design, all OpenShift Container Platform updates
        between minor versions must be serialized. You must update from OpenShift Container
        Platform 4.20 to 4.21 and then to 4.22. You cannot update from OpenShift Container
        Platform 4.20 to 4.22 directly. For more information, see Preparing to perform an EUS-
        to-EUS update in the Updating clusters guide in the OpenShift Container Platform
        documentation.

  OpenShift Data Foundation 4.22 offers the following upgrade channel:

        stable-4.22

        stable-4.21

  stable-4.22 channel

  Once a new version is Generally Available, the stable channel corresponding to the minor version gets
  updated with the new image which can be used to upgrade. You can use the stable-4.22 channel to
  upgrade from OpenShift Data Foundation 4.21 and upgrades within 4.22.

  stable-4.21

  You can use the stable-4.21 channel to upgrade from OpenShift Data Foundation 4.20 and upgrades
  within 4.21.




    8

                CHAPTER 3. UPDATING RED HAT OPENSHIFT DATA FOUNDATION 4.21 TO 4.22

 CHAPTER 3. UPDATING RED HAT OPENSHIFT DATA
                  FOUNDATION 4.21 TO 4.22

This chapter helps you to upgrade between the minor releases for all Red Hat OpenShift Data
Foundation deployments (Internal, Internal-Attached and External). The upgrade process remains the
same for all deployments. The Only difference is what gets upgraded and what’s not.

 For Internal and Internal-attached deployments, upgrading OpenShift Data Foundation
 upgrades all OpenShift Data Foundation services including the backend Red Hat Ceph Storage
 (RHCS) cluster.

 For External mode deployments, upgrading OpenShift Data Foundation only upgrades the
 OpenShift Data Foundation service while the backend Ceph storage cluster remains untouched
 and needs to be upgraded separately.
 You must upgrade Red Hat Ceph Storage along with OpenShift Data Foundation to get new
 feature support, security fixes, and other bug fixes. As there is no dependency on RHCS
 upgrade, you can upgrade the OpenShift Data Foundation operator first followed by RHCS
 upgrade or vice-versa. For more information about RHCS releases, see the knowledgebase
 solution, Red Hat Ceph Storage releases and corresponding Ceph package versions .

             IMPORTANT

             Upgrading to 4.22 directly from any version older than 4.21 is not supported.

             When upgrading from OpenShift Data Foundation 4.21.z to 4.22, you must upgrade to
             OpenShift Data Foundation 4.22.2 or later versions. Upgrades to OpenShift Data
             Foundation 4.22.0 or 4.22.1 are not supported and will fail.

Prerequisites

 Ensure that the OpenShift Container Platform cluster has been updated to the latest stable
 release of version 4.22.x, see Updating Clusters.

 Ensure that the OpenShift Data Foundation cluster is healthy and data is resilient.

  Navigate to Storage → Data Foundation → Storage Systemstab and then click on the
  storage system name.

  Check for the green tick on the status card of both Overview - Block and File and Object
  tabs. Green tick indicates that the storage cluster , object service and data resiliency are all
  healthy.

 Ensure that all OpenShift Data Foundation Pods, including the operator pods, are in     Running
 state in the openshift-storage namespace.
 To view the state of the pods, on the OpenShift Web Console, click Workloads → Pods. Select
 openshift-storage from the Project drop-down list.

              NOTE

              If the Show default projects option is disabled, use the toggle button to list all
              the default projects.

 Ensure that you have sufficient time to complete the OpenShift Data Foundation update
 process, as the update time varies depending on the number of OSDs that run in the cluster.


              9

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

      Optional: To reduce the upgrade time for large clusters, it is encouraged to increase the
      maxUnavailable for NodePlugin. Use the following command:

 VAL=$(oc get nodes -o name | awk 'END {printf "%.0f\n", NR * 0.2}') && oc patch driver openshift-
 storage.rbd.csi.ceph.com -n openshift-storage --type='merge' -p "{\"spec\":{\"nodePlugin\":
 {\"updateStrategy\":{\"rollingUpdate\":{\"maxUnavailable\":$VAL}}}}}" && oc patch driver openshift-
 storage.cephfs.csi.ceph.com -n openshift-storage --type='merge' -p "{\"spec\":{\"nodePlugin\":
 {\"updateStrategy\":{\"rollingUpdate\":{\"maxUnavailable\":$VAL}}}}}"

This affects CephFS and CephRBD daemonsets and allows the pods to restart simultaneously or be
unavailable and reduce the upgrade time. For an optimal value, you can set the parameter values to
20%. However, if the value is too high, disruption for new volumes might be observed during the
upgrade.

      Prerequisite relevant only for OpenShift Data Foundation deployments on AWS using AWS
      Security Token Service (STS)
      Add another entry in the trust policy for noobaa-core account as follows:

       1.   Log into AWS web console where the AWS role resides using
            http://console.aws.amazon.com/.

       2.   Enter the IAM management tool and click Roles.

       3.   Find the name of the role created for AWS STS to support Multicloud Object Gateway
            (MCG) authentication using the following command in OpenShift CLI:

            $ oc get deployment noobaa-operator -o yaml -n openshift-storage | grep ROLEARN -A1
            value: arn:aws:iam::123456789101:role/your-role-name-here

       4.   Search for the role name that you obtained from the previous step in the tool and click on
            the role name.

       5.   Under the role summary, click Trust relationships.

       6.   In the Trusted entities tab, click Edit trust policy on the right.

       7.   Under the “Action”: “sts:AssumeRoleWithWebIdentity” field, there are two fields to
            enable access for two NooBaa service accounts noobaa and noobaa-endpoint. Add
            another entry for the core pod’s new service account name,
            system:serviceaccount:openshift-storage:noobaa-core.

       8.   Click Update policy at the bottom right of the page.
            The update might take about 5 minutes to get in place.

Procedure

 1.   On the OpenShift Web Console, navigate to Ecosystem → Installed Operators.

 2.   Select openshift-storage project.

 3.   Click the OpenShift Data Foundation operator name.

 4.   Click the Subscription tab and click the link under Update Channel.

 5.   Select the stable-4.22 update channel and Save it.


10

                                    CHAPTER 3. UPDATING RED HAT OPENSHIFT DATA FOUNDATION 4.21 TO 4.22

6.  If the Upgrade status shows requires approval, click on requires approval.

     a.  On the Install Plan Details page, click Preview Install Plan.

     b.  Review the install plan and click Approve.
         Wait for the Status to change from Unknown to Created.

7.  Navigate to Ecosystem → Installed Operators.

8.  Select the openshift-storage project.
    Wait for the OpenShift Data Foundation Operator Status to change to Up to date.

9.  After the operator is successfully upgraded, the web console automatically reloads to apply the
    changes. During this process, a temporary error message might appear on the page and this is
    expected and disappears after the refresh completes.

                    NOTE

                    After upgrading, if your cluster has five or more nodes, racks, or rooms, and when
                    there are five or more number of failure domains present in the deployment, you
                    can configure Ceph monitor counts based on the number of racks or zones. An
                    alert is displayed in the notification panel or Alert Center of the OpenShift Web
                    Console to indicate the option to increase the number of Ceph monitor counts.
                    You can use the Configure option in the alert to configure the Ceph monitor
                    counts. For more information, see Resolving low Ceph monitor count alert in the
                    Troubleshooting guide.

Verification

    Check the Version after the OpenShift Data Foundation name and check the operator status.

         Navigate to Ecosystem → Installed Operators and select the openshift-storage project.

         When the upgrade completes, the version updates to a new version number for OpenShift
         Data Foundation and status changes to Succeeded with a green tick.

    Verify that the OpenShift Data Foundation cluster is healthy and data is resilient.

         Navigate to Storage → Data Foundation → Storage Systems tab and then click on the
         storage system name.

         Check for the green tick on the status card of Overview- Block and File and Object tabs.
         Green tick indicates that the storage cluster, object service and data resiliency is healthy.

    If verification steps fail, contact Red Hat Support .

            IMPORTANT

            After updating external mode deployments, you must also update the external secret. For
            instructions, see Updating the OpenShift Data Foundation external secret .

Additional resources

    Updating the OpenShift Data Foundation external secret

    Commonly required logs for troubleshooting

                                                                                                    11

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

Resolving low Ceph monitor count alert










12

                 CHAPTER 4. UPDATING RED HAT OPENSHIFT DATA FOUNDATION 4.22.X TO 4.22.Y

     CHAPTER 4. UPDATING RED HAT OPENSHIFT DATA
                 FOUNDATION 4.22.X TO 4.22.Y

Apply a OpenShift Data Foundation z-stream update within the 4.22 release to receive bug fixes and
security patches without changing the minor version, using either the automatic approval strategy or by
manually approving the generated InstallPlan.

     For Internal and Internal-attached deployments, upgrading OpenShift Data Foundation
     upgrades all OpenShift Data Foundation services including the backend Red Hat Ceph Storage
     (RHCS) cluster.

     For External mode deployments, upgrading OpenShift Data Foundation only upgrades the
     OpenShift Data Foundation service while the backend Ceph storage cluster remains untouched
     and needs to be upgraded separately.
     Hence, we recommend upgrading RHCS along with OpenShift Data Foundation in order to get
     new feature support, security fixes, and other bug fixes. Since we do not have a strong
     dependency on RHCS upgrade, you can upgrade the OpenShift Data Foundation operator first
     followed by RHCS upgrade or vice-versa. See solution to know more about RHCS releases.

When a new z-stream release becomes available, the upgrade process triggers automatically if the
update strategy was set to Automatic. If the update strategy is set to Manual then use the following
procedure.

Prerequisites

     Ensure that the OpenShift Container Platform cluster has been updated to the latest stable
     release of version 4.22.x, see Updating Clusters.

     Ensure that the OpenShift Data Foundation cluster is healthy and data is resilient.

          Navigate to Storage → Data Foundation → Storage Systemstab and then click on the
          storage system name.

          Check for the green tick on the status card of Overview - Block and File and Object tabs.
          Green tick indicates that the storage cluster, object service and data resiliency is healthy.

     Ensure that all OpenShift Data Foundation Pods, including the operator pods, are in    Running
     state in the openshift-storage namespace.
     To view the state of the pods, on the OpenShift Web Console, click Workloads → Pods. Select
     openshift-storage from the Project drop-down list.

             NOTE

             If the Show default projects option is disabled, use the toggle button to list all
             the default projects.

     Ensure that you have sufficient time to complete the OpenShift Data Foundation update
     process, as the update time varies depending on the number of OSDs that run in the cluster.

Procedure

1.   On the OpenShift Web Console, navigate to Ecosystem → Installed Operators.

2.   Select openshift-storage project.


             13

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

            NOTE

            If the Show default projects option is disabled, use the toggle button to list all
            the default projects.

3.  Click the OpenShift Data Foundation operator name.

4.  Click the Subscription tab.

5.  If the Upgrade Status shows require approval, click on requires approval link.

6.  On the InstallPlan Details page, click Preview Install Plan.

7.  Review the install plan and click Approve.

8.  Wait for the Status to change from Unknown to Created.

9.  After the operator is successfully installed, the web console automatically reloads to apply the
    changes. During this process, a temporary error message might appear on the page and this is
    expected and disappears after the refresh completes.

Verification

    Check the Version after the OpenShift Data Foundation name and check the operator status.

     Navigate to Ecosystem → Installed Operators and select the openshift-storage project.

     When the upgrade completes, the version updates to a new version number for OpenShift
     Data Foundation and status changes to Succeeded with a green tick.

    Verify that the OpenShift Data Foundation cluster is healthy and data is resilient.

     Navigate to Storage → Data Foundation → Storage Systemstab and then click on the
     storage system name.

     Check for the green tick on the status card of Overview - Block and File and Object tabs.
     Green tick indicates that the storage cluster, object service and data resiliency is healthy

    If verification steps fail, contact Red Hat Support .










14

                CHAPTER 5. CHANGING THE UPDATE APPROVAL STRATEGY

CHAPTER 5. CHANGING THE UPDATE APPROVAL STRATEGY

To ensure that the storage system gets updated automatically when a new update is available in the
same channel, we recommend keeping the update approval strategy to Automatic. Changing the
update approval strategy to Manual will need manual approval for each upgrade.

Procedure

1.   Navigate to Ecosystem → Installed Operators.

2.   Select openshift-storage from the Project drop-down list.

            NOTE

            If the Show default projects option is disabled, use the toggle button to list all
            the default projects.

3.   Click on OpenShift Data Foundation operator name

4.   Go to the Subscription tab.

5.   Click on the pencil icon for changing the Update approval.

6.   Select the update approval strategy and click Save.

Verification

     Verify that the Update approval shows the newly selected approval strategy after it.










            15

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

CHAPTER 6. UPDATING THE OPENSHIFT DATA FOUNDATION
                 EXTERNAL SECRET

After upgrading to the latest version of OpenShift Data Foundation, update the OpenShift Data
Foundation external secret.

             NOTE

             Updating the external secret is not required for batch updates. For example, when
             updating from OpenShift Data Foundation 4.22.x to 4.22.y.

Prerequisites

     Update the OpenShift Container Platform cluster to the latest stable release of 4.22.z, see
     Updating Clusters.

     Ensure that the OpenShift Data Foundation cluster is healthy and the data is resilient. Navigate
     to Storage → Data Foundation → Storage Systems tab and then click on the storage system
     name.

          On the Overview - Block and File tab, check the Status card and confirm that the Storage
          cluster has a green tick indicating it is healthy.

          Click the Object tab and confirm Object Service and Data resiliency has a green tick
          indicating it is healthy. The RADOS Object Gateway is only listed in case RADOS Object
          Gateway endpoint details are included while deploying OpenShift Data Foundation in
          external mode.

     Red Hat Ceph Storage must have a Ceph dashboard installed and configured.

Procedure

1.   Download the OpenShift Data Foundation version of the ceph-external-cluster-details-
     exporter.py python script:

         # oc get cm rook-ceph-external-cluster-script-config -n openshift-storage -o
         jsonpath='{.data.script}' | base64 --decode > ceph-external-cluster-details-exporter.py

2.   Update permission caps on the external Red Hat Ceph Storage cluster by running ceph-
     external-cluster-details-exporter.py on any client node in the external Red Hat Ceph Storage
     cluster. You may need to ask your Red Hat Ceph Storage administrator to do this.

         # python3 ceph-external-cluster-details-exporter.py --upgrade

     The updated permissions for the user are set as:

         client.csi-cephfs-node
         key: AQCYz0piYgu/IRAAipji4C8+Lfymu9vOrox3zQ==
         caps: [mds] allow rw
         caps: [mgr] allow rw
         caps: [mon] allow r, allow command 'osd blocklist'
         caps: [osd] allow rw tag cephfs =
         client.csi-cephfs-provisioner
         key: AQCYz0piDUMSIxAARuGUyhLXFO9u4zQeRG65pQ==

16

            CHAPTER 6. UPDATING THE OPENSHIFT DATA FOUNDATION EXTERNAL SECRET

       caps: [mgr] allow rw
       caps: [mon] allow r, allow command 'osd blocklist'
       caps: [osd] allow rw tag cephfs metadata=*
       client.csi-rbd-node
       key: AQCYz0pi88IKHhAAvzRN4fD90nkb082ldrTaHA==
       caps: [mon] profile rbd, allow command 'osd blocklist'
       caps: [osd] profile rbd
       client.csi-rbd-provisioner
       key: AQCYz0pi6W8IIBAAgRJfrAW7kZfucNdqJqS9dQ==
       caps: [mgr] allow rw
       caps: [mon] profile rbd, allow command 'osd blocklist'
       caps: [osd] profile rbd

3.  Run the previously downloaded python script using one of the following options based on the
    method you used during deployment, either a configuration file or command-line flags.

    a.  Configuration file
        Create a config.ini file that includes all of the parameters used during initial deployment.
        Run the following command to get the configmap output which contains those parameters:

         $ oc get configmap -namespace openshift-storage external-cluster-user-command --
         output jsonpath='{.data.args}'

        Add the parameters from the previous output to the config.ini file. You can add additional
        parameters to the config.ini file to those used during deployment. See b for descriptions of
        the parameters.

        Example config.ini file:

         [Configurations]
         format = bash
         cephfs-filesystem-name = <filesystem-name>
         rbd-data-pool-name = <pool_name>
         ...

        Run the python script:

         # python3 ceph-external-cluster-details-exporter.py --config-file <config-file>

        Replace <config-file> with the path to the config.ini file.

    b.  Command-line flags
        Run the previously downloaded python script and pass the parameters for your deployment.
        Make sure to use all the flags that you used in the original deployment including any optional
        argument that you have used. You can also add additional flags to those used during
        deployment. See b for descriptions of the parameters.

         # python3 ceph-external-cluster-details-exporter.py --rbd-data-pool-name _<rbd block
         pool name>_ --monitoring-endpoint _<ceph mgr prometheus exporter endpoint>_ --
         monitoring-endpoint-port _<ceph mgr prometheus exporter port>_ --rgw-endpoint _<rgw
         endpoint>_ --run-as-user _<ocs_client_name>_ [optional arguments]

        Mandatory and optional parameters used during upgrade



         17

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation


Parameter                     Description

rbd-data-pool-name            (Mandatory) Used for providing block storage in
                              OpenShift Data Foundation.

rgw-endpoint                  (Optional) Provide this parameter if object storage is
                              to be provisioned through Ceph RADOS Gateway for
                              OpenShift Data Foundation. Provide the endpoint in
                              the following format: <ip_address>:<port>.

monitoring-endpoint           (Optional) Accepts comma separated list of IP
                              addresses of active and standby mgrs reachable
                              from the OpenShift Container Platform cluster. If not
                              provided, the value is automatically populated.

monitoring-endpoint-port      (Optional) The port associated with the ceph-mgr
                              Prometheus exporter specified by --monitoring-
                              endpoint. If not provided, the value is automatically
                              populated.

run-as-user                   (Mandatory) The client name used during OpenShift
                              Data Foundation cluster deployment. Use the
                              default client name client.healthchecker if a
                              different client name was not set.

rgw-pool-prefix               (Optional) The prefix of the RGW pools. If not
                              specified, the default prefix is default.

rgw-tls-cert-path             (Optional) The file path of the RADOS Gateway
                              endpoint TLS certificate.

rgw-skip-tls                  (Optional) This parameter ignores the TLS
                              certification validation when a self-signed certificate
                              is provided (NOT RECOMMENDED).

ceph-conf                     (Optional) The name of the Ceph configuration file.

k8s-cluster-name              (Optional) Kubernetes cluster name.

cluster-name                  (Optional) The Ceph cluster name.

output                        (Optional) The file where the output is required to be
                              stored.

cephfs-metadata-pool-name     (Optional) The name of the CephFS metadata pool.

cephfs-data-pool-name         (Optional) The name of the CephFS data pool.

cephfs-filesystem-name        (Optional) The name of the CephFS filesystem.



18

    CHAPTER 6. UPDATING THE OPENSHIFT DATA FOUNDATION EXTERNAL SECRET


    rbd-metadata-ec-pool-name     (Optional) The name of the erasure coded RBD
                                  metadata pool.


    dry-run                       (Optional) This parameter helps to print the
                                  executed commands without running them.


    1.   Save the JSON output generated after running the script in the previous step.
         Example output:

              [{"name": "rook-ceph-mon-endpoints", "kind": "ConfigMap", "data": {"data":
              "xxx.xxx.xxx.xxx:xxxx", "maxMonId": "0", "mapping": "{}"}}, {"name": "rook-ceph-mon", "kind":
              "Secret", "data": {"admin-secret": "admin-secret", "fsid": "<fs-id>", "mon-secret": "mon-
              secret"}}, {"name": "rook-ceph-operator-creds", "kind": "Secret", "data": {"userID": "<user-id>",
              "userKey": "<user-key>"}}, {"name": "rook-csi-rbd-node", "kind": "Secret", "data": {"userID":
              "csi-rbd-node", "userKey": "<user-key>"}}, {"name": "ceph-rbd", "kind": "StorageClass", "data":
              {"pool": "<pool>"}}, {"name": "monitoring-endpoint", "kind": "CephCluster", "data":
              {"MonitoringEndpoint": "xxx.xxx.xxx.xxxx", "MonitoringPort": "xxxx"}}, {"name": "rook-ceph-
              dashboard-link", "kind": "Secret", "data": {"userID": "ceph-dashboard-link", "userKey": "<user-
              key>"}}, {"name": "rook-csi-rbd-provisioner", "kind": "Secret", "data": {"userID": "csi-rbd-
              provisioner", "userKey": "<user-key>"}}, {"name": "rook-csi-cephfs-provisioner", "kind":
              "Secret", "data": {"adminID": "csi-cephfs-provisioner", "adminKey": "<admin-key>"}}, {"name":
              "rook-csi-cephfs-node", "kind": "Secret", "data": {"adminID": "csi-cephfs-node", "adminKey": "
              <admin-key>"}}, {"name": "cephfs", "kind": "StorageClass", "data": {"fsName": "cephfs",
              "pool": "cephfs_data"}}, {"name": "ceph-rgw", "kind": "StorageClass", "data": {"endpoint":
              "xxx.xxx.xxx.xxxx", "poolPrefix": "default"}}, {"name": "rgw-admin-ops-user", "kind": "Secret",
              "data": {"accessKey": "<access-key>", "secretKey": "<secret-key>"}}]

     2. Upload the generated JSON file.

          a.   Log in to the OpenShift Web Console.

          b.   Click Workloads → Secrets.

          c.   Set project to openshift-storage.

          d.   Click rook-ceph-external-cluster-details.

          e.   Click Actions (⋮) → Edit Secret.

          f.   Click Browse and upload the JSON file.

          g.   Click Save.

Verification

         To verify that the OpenShift Data Foundation cluster is healthy and data is resilient, navigate to
         Storage → Data Foundation → Storage Systems tab and then click on the storage system
         name.

               On the Overview → Block and File tab, check the Details card to verify that the RHCS
               dashboard link is available and also check the Status card to confirm that the Storage
               Cluster has a green tick indicating it is healthy.

               Click the Object tab and confirm Object Service and Data resiliency has a green tick

                                      19

Red Hat OpenShift Data Foundation 4.22 Updating OpenShift Data Foundation

indicating it is healthy. The RADOS Object Gateway is only listed in case RADOS Object
Gateway endpoint details are included while deploying OpenShift Data Foundation in
external mode.

If verification steps fail, contact Red Hat Support .










20
