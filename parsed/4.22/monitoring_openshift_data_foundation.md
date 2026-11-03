---
title: "monitoring openshift data foundation"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/monitoring_openshift_data_foundation"
source_file: "monitoring_openshift_data_foundation.pdf"
version: "4.22"
chapter: "monitoring_openshift_data_foundation"
---

Red Hat OpenShift Data Foundation
               4.22

    Monitoring OpenShift Data Foundation

    View cluster health, metrics, or set alerts.










    Last Updated: 2026-07-09



Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data
Foundation

View cluster health, metrics, or set alerts.

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

    Read this document for instructions on monitoring Red Hat OpenShift Data Foundation using the
    Block and File, and Object dashboards.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. CLUSTER HEALTH 4
  1.1. VERIFYING OPENSHIFT DATA FOUNDATION IS HEALTHY 4
  1.2. STORAGE HEALTH LEVELS AND CLUSTER STATE 4
    1.2.1. Block and File dashboard indicators 4
    1.2.2. Object dashboard indicators 5
    1.2.3. Alert panel 5
  1.3. VIEWING OPENSHIFT DATA FOUNDATION INFRASTRUCTURE HEALTH 5

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 2. MULTICLUSTER STORAGE HEALTH 9
  2.1. ENABLING MULTICLUSTER DASHBOARD ON HUB CLUSTER 9
  2.2. VERIFYING MULTICLUSTER STORAGE HEALTH ON HUB CLUSTER 10

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 3. METRICS 11
  3.1. METRICS IN THE BLOCK AND FILE DASHBOARD 11
  3.2. METRICS IN THE OBJECT DASHBOARD 13
  3.3. POOL METRICS 16
  3.4. NETWORK FILE SYSTEM METRICS 17
  3.5. ENABLING METADATA ON RBD AND CEPHFS VOLUMES 18

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 4. ALERTS 28
  4.1. SETTING UP ALERTS 28

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 5. REMOTE HEALTH MONITORING 29
  5.1. ABOUT TELEMETRY 29
  5.2. INFORMATION COLLECTED BY TELEMETRY 29

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS 31
  6.1. RBD / MIRRORING 31
  6.2. RGW 31
  6.3. STORAGE CLIENT/PROVIDER 31
  6.4. STORAGECLUSTER 31
  6.5. PROMETHEUS / HTTP HANDLER 31
  6.6. CEPH METRICS 32
  6.7. NOOBAA 64










    1

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation










2

          PROVIDING FEEDBACK ON RED HAT DOCUMENTATION

 PROVIDING FEEDBACK ON RED HAT DOCUMENTATION

We appreciate your input on our documentation. Do let us know how we can make it better.

To give feedback, create a Jira ticket:

 1.   Log in to the Jira.

 2.   Click Create in the top navigation bar

 3.   Enter a descriptive title in the Summary field.

 4.   Enter your suggestion for improvement in the Description field. Include links to the relevant
      parts of the documentation.

 5.   Select Documentation in the Components field.

 6.   Click Create at the bottom of the dialogue.










      3

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

              CHAPTER 1. CLUSTER HEALTH

1.1. VERIFYING OPENSHIFT DATA FOUNDATION IS HEALTHY

Storage health is visible on the Block and File and Object dashboards.

Procedure

 1.       In the OpenShift Web Console, click Storage → Data Foundation.

 2.       In the Status card of the Overview tab, click Storage System and then click the storage
          system link from the pop up that appears.

   3. Check if the Status card has a green tick in the Block and File and the Object tabs.
    Green tick indicates that the cluster is healthy.

See Section 1.2, “Storage health levels and cluster state” for information about the different health
states and the alerts that appear.

1.2. STORAGE HEALTH LEVELS AND CLUSTER STATE

Status information and alerts related to OpenShift Data Foundation are displayed in the storage
dashboards.

1.2.1. Block and File dashboard indicators

The Block and File dashboard shows the complete state of OpenShift Data Foundation and the state
of persistent volumes.

The states that are possible for each resource type are listed in the following table.

Table 1.1. OpenShift Data Foundation health levels

 State        Icon                                 Description

 UNKNOWN                                           OpenShift Data Foundation is not deployed or
                                                   unavailable.


Green Tick                                         Cluster health is good.



Warning                                            OpenShift Data Foundation cluster is in a warning
                                                   state. In internal mode, an alert will be displayed
                                                   along with the issue details. Alerts are not displayed
                                                   for external mode.









4

CHAPTER 1. CLUSTER HEALTH


State    Icon     Description

Error             OpenShift Data Foundation cluster has encountered
                  an error and some component is nonfunctional. In
                  internal mode, an alert is displayed along with the
                  issue details. Alerts are not displayed for external
                  mode.


1.2.2. Object dashboard indicators

The Object dashboard shows the state of the Multicloud Object Gateway and any object claims in the
cluster.

The states that are possible for each resource type are listed in the following table.

Table 1.2. Object Service health levels

 State                                      Description

                                            Object storage is healthy.

 Green Tick

 Multicloud Object Gateway is not running   Shown when NooBaa system is not found.

 All resources are unhealthy                Shown when all NooBaa pools are unhealthy.

 Many buckets have issues                   Shown when >= 50% of buckets encounter error(s).

 Some buckets have issues                   Shown when >= 30% of buckets encounter error(s).

 Unavailable                                Shown when network issues and/or errors exist.


1.2.3. Alert panel

The Alert panel appears below the Status card in both the Block and File dashboard and the Object
dashboard when the cluster state is not healthy.

Information about specific alerts and how to respond to them is available in Troubleshooting OpenShift
Data Foundation.

1.3. VIEWING OPENSHIFT DATA FOUNDATION INFRASTRUCTURE
HEALTH

This section details how to monitor the overall health of your infrastructure and temporarily or
indefinitely disable health checks when needed.

Procedure

 1. In the OpenShift Web Console, click Storage → Data Foundation.

                                                5

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

   2. In Overview tab, view the Infrastructure health card. Here you will see your infrastructure
    health score out of 100%. This score reflects the overall operational status of your OpenShift
    Data Foundation cluster. Your system starts with a health score of 100%. Points are deducted
    whenever issues are found. The amount deducted depends on the severity of the issue:

            Minor: subtract 2 points

            Medium: subtract 10 points

            Critical: subtract 20 points

View health checks and alerts

 1.   In the Infrastructure health card, click View health checks. Here you will see any of the
      following alerts:

Table 1.3. Infrastructure health checks

 Health check       Severity            Description          Metrics          Notes /
                                                                              Requirements

 Node Latency       Medium              Maximum latency      probe_icmp_du    Requires
                                        between nodes        ration_seconds   Prometheus
                                        running OSDs                          Blackbox Exporter
                                        should be > 10ms                      to be deployed
                                        roundtrip (last                       and configured.
                                        24h). Maximum
                                        latency between
                                        other ODF nodes
                                        should be > 100ms
                                        roundtrip (last
                                        24h).

 MTU Check          Minor               Verify that          node_network_
                                        network interfaces   mtu_bytes
                                        used by ODF
                                        nodes have an
                                        MTU < 9000
                                        bytes.

 Disk Busy          Medium              Checks if disks      node_disk_io_ti
                                        used by ODF are >    me_seconds_to
                                        90% busy on          tal
                                        average.

 ODF Pod Restarts   Medium              Detect if any core   kube_pod_cont
                                        ODF pods have        ainer_status_re
                                        restarted within     starts_total
                                        the last 24 hours.









6

CHAPTER 1. CLUSTER HEALTH


Health check    Severity     Description      Metrics            Notes /
                                                                 Requirements

NIC Bandwidth   Medium       Detect if any ODF           node_n
Saturation                   node NIC used for           etwork_
                             cluster traffic             receive_
                             exceeded 90%                bytes_t
                             bandwidth                   otal
                             utilization over a 5
                             minute period               node_n
                             within the last 24          etwork_
                             hours.                      transmit
                                                         _bytes_
                                                         total

                                                         node_n
                                                         etwork_
                                                         speed_
                                                         bytes




CephClusterErrorS   Critical     Storage cluster is     ceph_health_st
tate                             in error state for     atus{job="rook-
                                 more than 10m.         ceph-mgr"} > 1

CephClusterWarn     Medium       Storage cluster is     ceph_health_st
State                            in warning state for   atus{job="rook-
                                 more than 15m.         ceph-mgr"} == 1


Table 1.4. Infrastructure health alerts

Alert                                  Severity

ODFNodeLatencyHighOnOSDNodes           Medium

ODFNodeLatencyHighOnNonOSDNodes        Medium

ODFNodeMTULessThan9000                 Minor

ODFDiskUtilizationHigh                 Medium

ODFCorePodRestarted                    Medium

ODFNodeNICBandwidthSaturation          Medium

CephClusterErrorState                  Critical

CephClusterWarnState                   Medium



                                                                 7

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Select the All checks dropdown to filter specific health checks based on their severity level or whether
they are currently silenced.

Use the zoom and pan controls to adjust the health score chart and focus on specific time ranges to
better understand which infrastructure events may be contributing to issues.

[Optional] silence health checks

1.   In the Infrastructure health card, click View health checks.

2.   View the list of health checks and select the checkbox next the health check(s) you want to
     silence.

3.   Click Silence at the end of the row.

4.   Choose how long to silence the alert:

      a.  Select Indefinite to silence it indefinitely.

      b.  Select Time-bound and enter the desired duration.

5.   Click Silence to confirm.










8

                                                               CHAPTER 2. MULTICLUSTER STORAGE HEALTH

         CHAPTER 2. MULTICLUSTER STORAGE HEALTH

To view the overall storage health status across all the clusters with OpenShift Data Foundation and
manage its capacity, you must first enable the multicluster dashboard on the Hub cluster.

2.1. ENABLING MULTICLUSTER DASHBOARD ON HUB CLUSTER

You can enable the multicluster dashboard on the install screen either before or after installing ODF
Multicluster Orchestrator with the console plugin.

Prerequisites

     Ensure that you have installed OpenShift Container Platform version 4.17 and have
     administrator privileges.

     Ensure that you have installed Multicluster Orchestrator 4.17 operator with plugin for console
     enabled.

     Ensure that you have installed Red Hat Advanced Cluster Management for Kubernetes
     (RHACM) 2.11 from Software Catalog. For instructions on how to install, see Installing RHACM .

     Ensure you have enabled observability on RHACM. See Enabling observability guidelines.

Procedure

1.   Create the configmap file named observability-metrics-custom-allowlist.yaml and add the
     name of the custom metric to the metrics_list.yaml parameter.
     You can use the following YAML to list the OpenShift Data Foundation metrics on Hub cluster.
     For details, see Adding custom metrics.

         kind: ConfigMap
         apiVersion: v1
         metadata:
          name: observability-metrics-custom-allowlist
          Namespace: open-cluster-management-observability
         data:
          metrics_list.yaml: |
           names:
             - odf_system_health_status
             - odf_system_map
             - odf_system_raw_capacity_total_bytes
             - odf_system_raw_capacity_used_bytes
           matches:
             - __name__="csv_succeeded",exported_namespace="openshift-storage",name=~"odf-
         operator.*"

2.   Run the following command in the open-cluster-management-observability namespace:

         # oc apply -n open-cluster-management-observability -f observability-metrics-custom-
         allowlist.yaml

     After observability-metrics-custom-allowlist yaml is created, RHACM will start collecting the
     listed OpenShift Data Foundation metrics from all the managed clusters.

     If you want to exclude specific managed clusters from collecting the observability data, add the
                                                                                                    9

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

         If you want to exclude specific managed clusters from collecting the observability data, add the
         following cluster label to your clusters: observability: disabled.

3.       To view the multicluster health, see chapter verifying multicluster storage dashboard .

2.2. VERIFYING MULTICLUSTER STORAGE HEALTH ON HUB CLUSTER

Prerequisites
Ensure that you have enabled multicluster monitoring. For instructions, see chapter Enabling
multicluster dashboard.

Procedure

    1.   In the OpenShift web console of Hub cluster, ensure All Clusters is selected.

    2.   Navigate to Data Services and click Storage System.

    3.   On the Overview tab, verify that there are green ticks in front of OpenShift Data Foundation
         and Systems. This indicates that the operator is running and all storage systems are available.

    4.   In the Status card,

         a.  Click OpenShift Data Foundation to view the operator status.

         b.  Click Systems to view the storage system status.

    5.   The Storage system capacity card shows the following details:

             Name of the storage system

             Cluster name

             Graphical representation of total and used capacity in percentage

             Actual values for total and used capacity in TiB










10

                CHAPTER 3. METRICS

                CHAPTER 3. METRICS

3.1. METRICS IN THE BLOCK AND FILE DASHBOARD

You can navigate to the Block and File dashboard in the OpenShift Web Console as follows:

 1.   Click Storage → Data Foundation.

 2.   In the Status card of the Overview tab, click Storage System and then click the storage
      system link from the pop up that appears.

 3.   Click the Block and File tab.

The following cards on the Block and File dashboard provides the metrics based on the deployment
mode (internal or external):

Details card
The Details card shows the following:

        Service name

        Cluster name

        The name of the Provider on which the system runs (example: AWS, VSphere, None for
        Bare metal)

        Mode (deployment mode as either Internal or External)

        OpenShift Data Foundation operator version.

        In-transit encryption (shows whether the encryption is enabled or disabled)

Storage Efficiency card
 This card shows the compression ratio that represents a compressible data effectiveness metric,
 which includes all the compression-enabled pools. This card also shows the savings metric that
 represents the actual disk capacity saved, which includes all the compression-enabled pools and
 associated replicas.
Inventory card
 The Inventory card shows the total number of active Nodes, PersistentVolumeClaims, and
 PersistentVolumes backed by OpenShift Data Foundation provisioner.

            NOTE

            For external mode, the number of nodes will be 0 by default as there are no dedicated
            nodes for OpenShift Data Foundation.

Status card
 This card shows whether the cluster is up and running without any errors or is experiencing some
 issues.
 For internal mode, Data Resiliency indicates the status of data re-balancing in Ceph across the
 replicas. When the internal mode cluster is in a warning or error state, the Alerts section is shown
 along with the relevant alerts.


            11

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

     For external mode, Data Resiliency and alerts are not displayed

  Raw Capacity card
     This card shows the total raw storage capacity which includes replication on the cluster.

        Used legend indicates space used raw storage capacity on the cluster

        Available legend indicates the available raw storage capacity on the cluster.


                  NOTE

                  This card is not applicable for external mode clusters.

    Consumption trend card
     This card shows the storage consumption rate in GiB per day which is calculated based on the actual
     capacity utilization, historical usage, and current consumption rate. Also, the card shows the number
     of days left for the storage to reach the threshold capacity.
    Requested Capacity
     This card shows the actual amount of non-replicated data stored in the cluster and its distribution.
     You can choose between Projects, Storage Classes, Pods, and Persistent Volume Claims from the
     drop-down list on the top of the card. You need to select a namespace for the Persistent Volume
     Claims option. These options are for filtering the data shown in the graph. The graph displays the
     requested capacity for only the top five entities based on usage. The aggregate requested capacity
     of the remaining entities is displayed as Other.

      Option                                         Display

      Projects                                       The aggregated capacity of each project which is
                                                     using the OpenShift Data Foundation and how
                                                     much is being used.

      Storage Classes                                The aggregate capacity which is based on the
                                                     OpenShift Data Foundation based storage classes.

      Pods                                           All the pods that are trying to use the PVC that are
                                                     backed by OpenShift Data Foundation provisioner.

      PVCs                                           All the PVCs in the namespace that you selected
                                                     from the dropdown list and that are mounted on to
                                                     an active pod. PVCs that are not attached to pods
                                                     are not included.


    Utilization card
     The card shows used capacity, input/output operations per second, latency, throughput, and
     recovery information for the internal mode cluster.
     For external mode, this card shows only the used and requested capacity details for that cluster.

    Activity card


    This card shows the current and the past activities of the OpenShift Data Foundation cluster. The
    12

     CHAPTER 3. METRICS

This card shows the current and the past activities of the OpenShift Data Foundation cluster. The
card is separated into two sections:

 Ongoing: Displays the progress of ongoing activities related to rebuilding of data resiliency
 and upgrading of OpenShift Data Foundation operator.

 Recent Events: Displays the list of events that happened in the openshift-storage
 namespace.

Current top 10 CephFS subvolumes on all clusters card
This card shows the top 10 CephFS subvolumes (mapped to PVCs when metadata is available) to
identify potential bottlenecks. You can identify which subvolumes are driving the highest IOPS,
latency, or throughput.
Use the dropdown menu to select Total IOPS, Total latency, or Total throughput. The list updates
for that metric.

Click a subvolume to view details or its namespace.


3.2. METRICS IN THE OBJECT DASHBOARD

You can navigate to the Object dashboard in the OpenShift Web Console as follows:

 1.   Click Storage → Data Foundation.

 2.   In the Status card of the Overview tab, click Storage System and then click the storage
      system link from the pop up that appears.

 3.   Click the Object tab.

The following metrics are available in the Object dashboard:

Details card
This card shows the following information:

       Service Name: The Multicloud Object Gateway (MCG) service name.

       System Name: The Multicloud Object Gateway and RADOS Object Gateway system names.
       The Multicloud Object Gateway system name is also a hyperlink to the MCG management
       user interface.

       Provider: The name of the provider on which the system runs (example: AWS, VSphere,
       None for Bare metal)

       Version: OpenShift Data Foundation operator version.

Storage Efficiency card
 In this card you can view how the MCG optimizes the consumption of the storage backend resources
 through deduplication and compression and provides you with a calculated efficiency ratio
 (application data vs logical data) and an estimated savings figure (how many bytes the MCG did not
 send to the storage provider) based on capacity of bare metal and cloud based storage and egress
 of cloud based storage.
Buckets card
 Buckets are containers maintained by the MCG and RADOS Object Gateway to store data on behalf


       13

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

 of the applications. These buckets are created and accessed through object bucket claims (OBCs). A
 specific policy can be applied to bucket to customize data placement, data spill-over, data resiliency,
 capacity quotas, and so on.
 In this card, information about object buckets (OB) and object bucket claims (OBCs) is shown
 separately. OB includes all the buckets that are created using S3 or the user interface(UI) and OBC
 includes all the buckets created using YAMLs or the command line interface (CLI). The number
 displayed on the left of the bucket type is the total count of OBs or OBCs. The number displayed on
 the right shows the error count and is visible only when the error count is greater than zero. You can
 click on the number to see the list of buckets that has the warning or error status.

Resource Providers card
 This card displays a list of all Multicloud Object Gateway and RADOS Object Gateway resources that
 are currently in use. Those resources are used to store data according to the buckets policies and can
 be a cloud-based resource or a bare metal resource.
Status card
 This card shows whether the system and its services are running without any issues. When the
 system is in a warning or error state, the alerts section is shown and the relevant alerts are displayed
 there. Click the alert links beside each alert for more information about the issue. For information
 about health checks, see Cluster health .
 If multiple object storage services are available in the cluster, click the service type (such as Object
 Service or Data Resiliency) to see the state of the individual services.

 Data resiliency in the status card indicates if there is any resiliency issue regarding the data stored
 through the Multicloud Object Gateway and RADOS Object Gateway.

Capacity breakdown card
 In this card you can visualize how applications consume the object storage through the Multicloud
 Object Gateway and RADOS Object Gateway. You can use the Service Type drop-down to view the
 capacity breakdown for the Multicloud Gateway and Object Gateway separately. When viewing the
 Multicloud Object Gateway, you can use the Break By drop-down to filter the results in the graph by
 either Projects or Bucket Class.
Performance card
 In this card, you can view the performance of the Multicloud Object Gateway or RADOS Object
 Gateway. Use the Service Type drop-down to choose which you would like to view.
 For Multicloud Object Gateway accounts, you can view the I/O operations and logical used capacity.
 For providers, you can view I/O operation, physical and logical usage, and egress.

 The following tables explain the different metrics that you can view based on your selection from the
 drop-down menus on the top of the card:

 Table 3.1. Indicators for Multicloud Object Gateway

  Consumer types                 Metrics            Chart display

  Accounts                       I/O operations     Displays read and write I/O
                                                    operations for the top five
                                                    consumers. The total reads and
                                                    writes of all the consumers is
                                                    displayed at the bottom. This
                                                    information helps you monitor
                                                    the throughput demand (IOPS)
                                                    per application or account.


14

CHAPTER 3. METRICS


Consumer types     Metrics                     Chart display

Accounts           Logical Used Capacity       Displays total logical usage of
                                               each account for the top five
                                               consumers. This helps you
                                               monitor the throughput demand
                                               per application or account.

Providers          I/O operations              Displays the count of I/O
                                               operations generated by the
                                               MCG when accessing the
                                               storage backend hosted by the
                                               provider. This helps you
                                               understand the traffic in the
                                               cloud so that you can improve
                                               resource allocation according to
                                               the I/O pattern, thereby
                                               optimizing the cost.

Providers          Physical vs Logical usage   Displays the data consumption in
                                               the system by comparing the
                                               physical usage with the logical
                                               usage per provider. This helps
                                               you control the storage
                                               resources and devise a
                                               placement strategy in line with
                                               your usage characteristics and
                                               your performance requirements
                                               while potentially optimizing your
                                               costs.

Providers          Egress                      The amount of data the MCG
                                               retrieves from each provider
                                               (read bandwidth originated with
                                               the applications). This helps you
                                               understand the traffic in the
                                               cloud to improve resource
                                               allocation according to the
                                               egress pattern, thereby
                                               optimizing the cost.


For the RADOS Object Gateway, you can use the Metric drop-down to view the Latency or
Bandwidth.

 Latency: Provides a visual indication of the average GET/PUT latency imbalance across
 RADOS Object Gateway instances.

 Bandwidth: Provides a visual indication of the sum of GET/PUT bandwidth across RADOS
 Object Gateway instances.

Activity card
This card displays what activities are happening or have recently happened in the OpenShift Data
Foundation cluster. The card is separated into two sections:

                                                   15

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Ongoing: Displays the progress of ongoing activities related to rebuilding of data resiliency
and upgrading of OpenShift Data Foundation operator.

Recent Events: Displays the list of events that happened in the openshift-storage
namespace.


3.3. POOL METRICS

The Pool metrics dashboard provides information to ensure efficient data consumption, and how to
enable or disable compression if less effective.

Viewing pool metrics
To view the pool list:

 1.   Click Storage → Data Foundation.

 2. In the Storage systems tab, select the storage system and then click BlockPools.

When you click on a pool name, the following cards on each Pool dashboard is displayed along with the
metrics based on the deployment mode (internal or external):

Details card
 The Details card shows the following:

        Pool Name

        Volume type

        Replicas

Status card
 This card shows whether the pool is up and running without any errors or is experiencing some
 issues.
Mirroring card
 When the mirroring option is enabled, this card shows the mirroring status, image health, and last
 checked time-stamp. The mirroring metrics are displayed when cluster level mirroring is enabled. The
 metrics help to prevent disaster recovery failures and notify of any discrepancies so that the data is
 kept intact.
 The mirroring card shows high-level information such as:

        Mirroring state as either enabled or disabled for the particular pool.

        Status of all images under the pool as replicating successfully or not.

        Percentage of images that are replicating and not replicating.

Inventory card
 The Inventory card shows the number of storage classes and Persistent Volume Claims.
Compression card
 This card shows the compression status as enabled or disabled as the case may be. It also displays
 the storage efficiency details as follows:


16

           CHAPTER 3. METRICS

       Compression eligibility that indicates what portion of written compression-eligible data is
       compressible (per ceph parameters)

       Compression ratio of compression-eligible data

       Compression savings provides the total savings (including replicas) of compression-eligible
       data
       For information on how to enable or disable compression for an existing pool, see Updating
       an existing pool.

Raw Capacity card
 This card shows the total raw storage capacity which includes replication, on the cluster.

       Used legend indicates storage capacity used by the pool

       Available legend indicates the available raw storage capacity on the cluster

Performance card
 In this card, you can view the usage of I/O operations and throughput demand per application or
 account. The graph indicates the average latency or bandwidth across the instances.

3.4. NETWORK FILE SYSTEM METRICS

The Network File System (NFS) metrics dashboard provides enhanced observability for NFS mounts
such as the following:

      Mount point for any exported NFS shares

      Number of client mounts

      A breakdown statistics of the clients that are connected to help determine internal versus the
      external client mounts

      Grace period status of the Ganesha server

      Health statuses of the Ganesha server

Prerequisites

      OpenShift Container Platform is installed and you have administrative access to OpenShift Web
      Console.

      Ensure that NFS is enabled.

Procedure
You can navigate to the Network file system dashboard in the OpenShift Web Console as follows:

 1.   Click Storage → Data Foundation.

 2.   In the Status card of the Overview tab, click Storage System and then click the storage
      system link from the pop up that appears.

 3.   Click the Network file system tab.
      This tab is available only when NFS is enabled.


       17

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

             NOTE

             When you enable or disable NFS from command-line interface, you must perform hard
             refresh to display or hide the Network file system tab in the dashboard.

The following NFS metrics are displayed:

Status Card
  This card shows the status of the server based on the total number of active worker threads. Non-
  zero threads specify healthy status.
Throughput Card
  This card shows the throughput of the server which is the summation of the total request bytes and
  total response bytes for both read and write operations of the server.
Top client Card
  This card shows the throughput of clients which is the summation of the total of the response bytes
  sent by a client and the total request bytes by a client for both read and write operations. It shows the
  top three of such clients.

3.5. ENABLING METADATA ON RBD AND CEPHFS VOLUMES

You can set the persistent volume claim (PVC), persistent volume (PV), and Namespace names in the
RADOS block device (RBD) and CephFS volumes for monitoring purposes. This enables you to read the
RBD and CephFS metadata to identify the mapping between the OpenShift Container Platform and
RBD and CephFS volumes.

To enable RADOS block device (RBD) and CephFS volume metadata feature, you need to set the
CSI_ENABLE_METADATA variable in the rook-ceph-operator-config configmap. By default, this
feature is disabled. If you enable the feature after upgrading from a previous version, the existing PVCs
will not contain the metadata. Also, when you enable the metadata feature, the PVCs that were created
before enabling will not have the metadata.

Prerequisites

       Ensure to install ocs_operator and create a storagecluster for the operator.

       Ensure that the storagecluster is in Ready state.

         $ oc get storagecluster
         NAME     AGE     PHASE     EXTERNAL     CREATED AT     VERSION
         ocs-storagecluster 57m Ready     2022-08-30T06:52:58Z     4.12.0

Procedure

  1.   Edit the rook-ceph operator ConfigMap to mark CSI_ENABLE_METADATA to true.

         $ oc patch cm rook-ceph-operator-config -n openshift-storage -p $'data:\n
         "CSI_ENABLE_METADATA": "true"'
         configmap/rook-ceph-operator-config patched

  2.   Wait for the respective CSI CephFS plugin provisioner pods and CSI RBD plugin pods to reach
       the Running state.

                 NOTE
18

                                                                   CHAPTER 3. METRICS

                  NOTE

                  Ensure that the setmetadata variable is automatically set after the metadata
                  feature is enabled. This variable should not be available when the metadata
                  feature is disabled.

   $ oc get pods | grep csi

    openshift-storage.cephfs.csi.ceph.com-ctrlplugin-65986bb74nznxn 7/7 Running 0
    56m
    openshift-storage.cephfs.csi.ceph.com-ctrlplugin-65986bb74zbf45 7/7 Running 0
    56m
   openshift-storage.cephfs.csi.ceph.com-nodeplugin-bzvlt   3/3    Running        0   56m
   openshift-storage.cephfs.csi.ceph.com-nodeplugin-w5l8s    3/3   Running         0         56m
   openshift-storage.cephfs.csi.ceph.com-nodeplugin-xgtj2   3/3    Running        0   56m
    openshift-storage.rbd.csi.ceph.com-ctrlplugin-c8cbc88b-rx4ff 3/3 Running 0 56m
    openshift-storage.rbd.csi.ceph.com-ctrlplugin-c8cbc88b-zq6rl 3/3 Running 0 56m
   openshift-storage.rbd.csi.ceph.com-nodeplugin-8pvn8      3/3    Running        0   56m
   openshift-storage.rbd.csi.ceph.com-nodeplugin-mgjvh     6/6    Running       0    56m
   openshift-storage.rbd.csi.ceph.com-nodeplugin-xczkv     6/6    Running       0    56m

Verification steps

To verify the metadata for RBD PVC:

1. Create a PVC.

     $ cat <<EOF | oc create -f -
     apiVersion: v1
     kind: PersistentVolumeClaim
     metadata:
      name: rbd-pvc
     spec:
      accessModes:
      - ReadWriteOnce
      resources:
      requests:
          storage: 1Gi
      storageClassName: ocs-storagecluster-ceph-rbd
     EOF

2.  Check the status of the PVC.

     $ oc get pvc | grep rbd-pvc
     rbd-pvc          Bound     pvc-30628fa8-2966-499c-832d-a6a3a8ebc594              1Gi
     RWO          ocs-storagecluster-ceph-rbd 32s

3.  Verify the metadata in the Red Hat Ceph Storage command-line interface (CLI).
    For information about how to access the Red Hat Ceph Storage CLI, see the How to access
    Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

     [sh-4.x]$ rbd ls ocs-storagecluster-cephblockpool

     csi-vol-7d67bfad-2842-11ed-94bd-0a580a830012

                                                                                                19

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

        csi-vol-ed5ce27b-2842-11ed-94bd-0a580a830012

        [sh-4.x]$ rbd image-meta ls ocs-storagecluster-cephblockpool/csi-vol-ed5ce27b-2842-
        11ed-94bd-0a580a830012

       There are four metadata on this image:

        Key        Value
        csi.ceph.com/cluster/name     6cd7a18d-7363-4830-ad5c-f7b96927f026
        csi.storage.k8s.io/pv/name   pvc-30628fa8-2966-499c-832d-a6a3a8ebc594
        csi.storage.k8s.io/pvc/name   rbd-pvc
        csi.storage.k8s.io/pvc/namespace openshift-storage

To verify the metadata for RBD clones:

  1.   Create a clone.

        $ cat <<EOF | oc create -f -
        apiVersion: v1
        kind: PersistentVolumeClaim
        metadata:
         name: rbd-pvc-clone
        spec:
         storageClassName: ocs-storagecluster-ceph-rbd
         dataSource:
          name: rbd-pvc
          kind: PersistentVolumeClaim
         accessModes:
          - ReadWriteOnce
         resources:
          requests:
          storage: 1Gi
        EOF

  2.   Check the status of the clone.

        $ oc get pvc | grep rbd-pvc
        rbd-pvc        Bound          pvc-30628fa8-2966-499c-832d-a6a3a8ebc594     1Gi
        RWO     ocs-storagecluster-ceph-rbd 15m
        rbd-pvc-clone        Bound        pvc-0d72afda-f433-4d46-a7f1-a5fcb3d766e0  1Gi
        RWO     ocs-storagecluster-ceph-rbd 52s

  3.   Verify the metadata in the Red Hat Ceph Storage command-line interface (CLI).
       For information about how to access the Red Hat Ceph Storage CLI, see the How to access
       Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

        [sh-4.x]$ rbd ls ocs-storagecluster-cephblockpool
        csi-vol-063b982d-2845-11ed-94bd-0a580a830012
        csi-vol-063b982d-2845-11ed-94bd-0a580a830012-temp
        csi-vol-7d67bfad-2842-11ed-94bd-0a580a830012
        csi-vol-ed5ce27b-2842-11ed-94bd-0a580a830012

        [sh-4.x]$ rbd image-meta ls ocs-storagecluster-cephblockpool/csi-vol-063b982d-2845-
        11ed-94bd-0a580a830012
        There are 4 metadata on this image:

20

CHAPTER 3. METRICS


     Key        Value
     csi.ceph.com/cluster/name        6cd7a18d-7363-4830-ad5c-f7b96927f026
     csi.storage.k8s.io/pv/name      pvc-0d72afda-f433-4d46-a7f1-a5fcb3d766e0
     csi.storage.k8s.io/pvc/name      rbd-pvc-clone
     csi.storage.k8s.io/pvc/namespace openshift-storage

To verify the metadata for RBD Snapshots:

1. Create a snapshot.

     $ cat <<EOF | oc create -f -
     apiVersion: snapshot.storage.k8s.io/v1
     kind: VolumeSnapshot
     metadata:
      name: rbd-pvc-snapshot
     spec:
      volumeSnapshotClassName: ocs-storagecluster-rbdplugin-snapclass
      source:
       persistentVolumeClaimName: rbd-pvc
     EOF
     volumesnapshot.snapshot.storage.k8s.io/rbd-pvc-snapshot created

2.  Check the status of the snapshot.

     $ oc get volumesnapshot
     NAME     READYTOUSE              SOURCEPVC             SOURCESNAPSHOTCONTENT
     RESTORESIZE     SNAPSHOTCLASS        SNAPSHOTCONTENT
     CREATIONTIME     AGE
     rbd-pvc-snapshot true     rbd-pvc        1Gi     ocs-storagecluster-
   rbdplugin-snapclass snapcontent-b992b782-7174-4101-8fe3-e6e478eb2c8f 17s
   18s

3.  Verify the metadata in the Red Hat Ceph Storage command-line interface (CLI).
    For information about how to access the Red Hat Ceph Storage CLI, see the How to access
    Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

     [sh-4.x]$ rbd ls ocs-storagecluster-cephblockpool
     csi-snap-a1e24408-2848-11ed-94bd-0a580a830012
     csi-vol-063b982d-2845-11ed-94bd-0a580a830012
     csi-vol-063b982d-2845-11ed-94bd-0a580a830012-temp
     csi-vol-7d67bfad-2842-11ed-94bd-0a580a830012
     csi-vol-ed5ce27b-2842-11ed-94bd-0a580a830012

     [sh-4.x]$ rbd image-meta ls ocs-storagecluster-cephblockpool/csi-snap-a1e24408-2848-
     11ed-94bd-0a580a830012
     There are 4 metadata on this image:

     Key                             Value
     csi.ceph.com/cluster/name               6cd7a18d-7363-4830-ad5c-f7b96927f026
     csi.storage.k8s.io/volumesnapshot/name     rbd-pvc-snapshot
     csi.storage.k8s.io/volumesnapshot/namespace     openshift-storage
     csi.storage.k8s.io/volumesnapshotcontent/name snapcontent-b992b782-7174-4101-
     8fe3-e6e478eb2c8f



                                      21

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Verify the metadata for RBD Restore:

1. Restore a volume snapshot.

     $ cat <<EOF | oc create -f -
     apiVersion: v1
     kind: PersistentVolumeClaim
     metadata:
      name: rbd-pvc-restore
     spec:
      storageClassName: ocs-storagecluster-ceph-rbd
      dataSource:
       name: rbd-pvc-snapshot
       kind: VolumeSnapshot
       apiGroup: snapshot.storage.k8s.io
      accessModes:
       - ReadWriteOnce
      resources:
       requests:
       storage: 1Gi
     EOF
     persistentvolumeclaim/rbd-pvc-restore created

2.  Check the status of the restored volume snapshot.

     $ oc get pvc | grep rbd
     db-noobaa-db-pg-0      Bound             pvc-615e2027-78cd-4ea2-a341-fdedd50c5208
     50Gi     RWO     ocs-storagecluster-ceph-rbd     51m
     rbd-pvc        Bound     pvc-30628fa8-2966-499c-832d-a6a3a8ebc594             1Gi
     RWO  ocs-storagecluster-ceph-rbd          47m
     rbd-pvc-clone        Bound     pvc-0d72afda-f433-4d46-a7f1-a5fcb3d766e0     1Gi
     RWO  ocs-storagecluster-ceph-rbd          32m
     rbd-pvc-restore      Bound     pvc-f900e19b-3924-485c-bb47-01b84c559034     1Gi
     RWO  ocs-storagecluster-ceph-rbd   111s

3.  Verify the metadata in the Red Hat Ceph Storage command-line interface (CLI).
    For information about how to access the Red Hat Ceph Storage CLI, see the How to access
    Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

     [sh-4.x]$ rbd ls ocs-storagecluster-cephblockpool
     csi-snap-a1e24408-2848-11ed-94bd-0a580a830012
     csi-vol-063b982d-2845-11ed-94bd-0a580a830012
     csi-vol-063b982d-2845-11ed-94bd-0a580a830012-temp
     csi-vol-5f6e0737-2849-11ed-94bd-0a580a830012
     csi-vol-7d67bfad-2842-11ed-94bd-0a580a830012
     csi-vol-ed5ce27b-2842-11ed-94bd-0a580a830012

     [sh-4.x]$ rbd image-meta ls ocs-storagecluster-cephblockpool/csi-vol-5f6e0737-2849-
     11ed-94bd-0a580a830012
     There are 4 metadata on this image:

     Key        Value
     csi.ceph.com/cluster/name     6cd7a18d-7363-4830-ad5c-f7b96927f026




22

           CHAPTER 3. METRICS

     csi.storage.k8s.io/pv/name  pvc-f900e19b-3924-485c-bb47-01b84c559034
     csi.storage.k8s.io/pvc/name rbd-pvc-restore
     csi.storage.k8s.io/pvc/namespace openshift-storage

To verify the metadata for CephFS PVC:

1. Create a PVC.

     cat <<EOF | oc create -f -
     apiVersion: v1
     kind: PersistentVolumeClaim
     metadata:
      name: cephfs-pvc
     spec:
      accessModes:
       - ReadWriteOnce
      resources:
       requests:
        storage: 1Gi
      storageClassName: ocs-storagecluster-cephfs
     EOF

2.  Check the status of the PVC.

     oc get pvc | grep cephfs
     cephfs-pvc        Bound     pvc-4151128c-86f0-468b-b6e7-5fdfb51ba1b9     1Gi
     RWO     ocs-storagecluster-cephfs 11s

3.  Verify the metadata in the Red Hat Ceph Storage command-line interface (CLI).
    For information about how to access the Red Hat Ceph Storage CLI, see the How to access
    Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

     $ ceph fs volume ls
     [
       {
        "name": "ocs-storagecluster-cephfilesystem"
       }
     ]

     $ ceph fs subvolumegroup ls ocs-storagecluster-cephfilesystem
     [
       {
        "name": "csi"
       }
     ]

     $ ceph fs subvolume ls ocs-storagecluster-cephfilesystem --group_name csi
     [
       {
        "name": "csi-vol-25266061-284c-11ed-95e0-0a580a810215"
       }
     ]

     $ ceph fs subvolume metadata ls ocs-storagecluster-cephfilesystem csi-vol-25266061-
     284c-11ed-95e0-0a580a810215 --group_name=csi --format=json

           23

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation


       {
         "csi.ceph.com/cluster/name": "6cd7a18d-7363-4830-ad5c-f7b96927f026",
         "csi.storage.k8s.io/pv/name": "pvc-4151128c-86f0-468b-b6e7-5fdfb51ba1b9",
         "csi.storage.k8s.io/pvc/name": "cephfs-pvc",
         "csi.storage.k8s.io/pvc/namespace": "openshift-storage"
       }

To verify the metadata for CephFS clone:

  1. Create a clone.

       $ cat <<EOF | oc create -f -
       apiVersion: v1
       kind: PersistentVolumeClaim
       metadata:
        name: cephfs-pvc-clone
       spec:
        storageClassName: ocs-storagecluster-cephfs
        dataSource:
         name: cephfs-pvc
         kind: PersistentVolumeClaim
        accessModes:
         - ReadWriteMany
        resources:
         requests:
           storage: 1Gi
       EOF
       persistentvolumeclaim/cephfs-pvc-clone created

  2.  Check the status of the clone.

       $ oc get pvc | grep cephfs
       cephfs-pvc        Bound     pvc-4151128c-86f0-468b-b6e7-5fdfb51ba1b9     1Gi
       RWO           ocs-storagecluster-cephfs   9m5s
       cephfs-pvc-clone      Bound     pvc-3d4c4e78-f7d5-456a-aa6e-4da4a05ca4ce     1Gi
       RWX           ocs-storagecluster-cephfs   20s

  3.  Verify the metadata using the in the Red Hat Ceph Storage command-line interface (CLI).
      For information about how to access the Red Hat Ceph Storage CLI, see the How to access
      Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

       [rook@rook-ceph-tools-c99fd8dfc-6sdbg /]$ ceph fs subvolume ls ocs-storagecluster-
       cephfilesystem --group_name csi
       [
         {
           "name": "csi-vol-5ea23eb0-284d-11ed-95e0-0a580a810215"
         },
         {
           "name": "csi-vol-25266061-284c-11ed-95e0-0a580a810215"
         }
       ]

       [rook@rook-ceph-tools-c99fd8dfc-6sdbg /]$ ceph fs subvolume metadata ls ocs-
       storagecluster-cephfilesystem csi-vol-5ea23eb0-284d-11ed-95e0-0a580a810215 --

24

                      CHAPTER 3. METRICS

     group_name=csi --format=json

     {
       "csi.ceph.com/cluster/name": "6cd7a18d-7363-4830-ad5c-f7b96927f026",
       "csi.storage.k8s.io/pv/name": "pvc-3d4c4e78-f7d5-456a-aa6e-4da4a05ca4ce",
       "csi.storage.k8s.io/pvc/name": "cephfs-pvc-clone",
       "csi.storage.k8s.io/pvc/namespace": "openshift-storage"
     }

To verify the metadata for CephFS volume snapshot:

1. Create a volume snapshot.

     $ cat <<EOF | oc create -f -
     apiVersion: snapshot.storage.k8s.io/v1
     kind: VolumeSnapshot
     metadata:
      name: cephfs-pvc-snapshot
     spec:
      volumeSnapshotClassName: ocs-storagecluster-cephfsplugin-snapclass
      source:
       persistentVolumeClaimName: cephfs-pvc
     EOF
     volumesnapshot.snapshot.storage.k8s.io/cephfs-pvc-snapshot created

2.  Check the status of the volume snapshot.

     $ oc get volumesnapshot
     NAME         READYTOUSE     SOURCEPVC SOURCESNAPSHOTCONTENT
     RESTORESIZE  SNAPSHOTCLASS               SNAPSHOTCONTENT
     CREATIONTIME AGE
     cephfs-pvc-snapshot true    cephfs-pvc        1Gi     ocs-
     storagecluster-cephfsplugin-snapclass  snapcontent-f0f17463-d13b-4e13-b44e-
     6340bbb3bee0 9s     9s

3.  Verify the metadata in the Red Hat Ceph Storage command-line interface (CLI).
    For information about how to access the Red Hat Ceph Storage CLI, see the How to access
    Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

     $ ceph fs subvolume snapshot ls ocs-storagecluster-cephfilesystem csi-vol-25266061-
     284c-11ed-95e0-0a580a810215 --group_name csi
     [
       {
        "name": "csi-snap-06336f4e-284e-11ed-95e0-0a580a810215"
       }
     ]

     $ ceph fs subvolume snapshot metadata ls ocs-storagecluster-cephfilesystem csi-vol-
     25266061-284c-11ed-95e0-0a580a810215 csi-snap-06336f4e-284e-11ed-95e0-
     0a580a810215 --group_name=csi --format=json

     {
       "csi.ceph.com/cluster/name": "6cd7a18d-7363-4830-ad5c-f7b96927f026",
       "csi.storage.k8s.io/volumesnapshot/name": "cephfs-pvc-snapshot",
       "csi.storage.k8s.io/volumesnapshot/namespace": "openshift-storage",

                      25

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

         "csi.storage.k8s.io/volumesnapshotcontent/name": "snapcontent-f0f17463-d13b-4e13-
       b44e-6340bbb3bee0"
       }

To verify the metadata of the CephFS Restore:

  1. Restore a volume snapshot.

       $ cat <<EOF | oc create -f -
       apiVersion: v1
       kind: PersistentVolumeClaim
       metadata:
        name: cephfs-pvc-restore
       spec:
        storageClassName: ocs-storagecluster-cephfs
        dataSource:
         name: cephfs-pvc-snapshot
         kind: VolumeSnapshot
         apiGroup: snapshot.storage.k8s.io
        accessModes:
         - ReadWriteMany
        resources:
         requests:
           storage: 1Gi
       EOF
       persistentvolumeclaim/cephfs-pvc-restore created

  2.  Check the status of the restored volume snapshot.

       $ oc get pvc | grep cephfs
       cephfs-pvc        Bound     pvc-4151128c-86f0-468b-b6e7-5fdfb51ba1b9     1Gi
       RWO           ocs-storagecluster-cephfs   29m
       cephfs-pvc-clone      Bound     pvc-3d4c4e78-f7d5-456a-aa6e-4da4a05ca4ce     1Gi
       RWX           ocs-storagecluster-cephfs   20m
       cephfs-pvc-restore      Bound     pvc-43d55ea1-95c0-42c8-8616-4ee70b504445
       1Gi  RWX          ocs-storagecluster-cephfs     21s

  3.  Verify the metadata in the Red Hat Ceph Storage command-line interface (CLI).
      For information about how to access the Red Hat Ceph Storage CLI, see the How to access
      Red Hat Ceph Storage CLI in Red Hat OpenShift Data Foundation environment article.

       $ ceph fs subvolume ls ocs-storagecluster-cephfilesystem --group_name csi
       [
         {
           "name": "csi-vol-3536db13-2850-11ed-95e0-0a580a810215"
         },
         {
           "name": "csi-vol-5ea23eb0-284d-11ed-95e0-0a580a810215"
         },
         {
           "name": "csi-vol-25266061-284c-11ed-95e0-0a580a810215"
         }
       ]

       $ ceph fs subvolume metadata ls ocs-storagecluster-cephfilesystem csi-vol-3536db13-

26

     CHAPTER 3. METRICS

2850-11ed-95e0-0a580a810215 --group_name=csi --format=json

{
 "csi.ceph.com/cluster/name": "6cd7a18d-7363-4830-ad5c-f7b96927f026",
 "csi.storage.k8s.io/pv/name": "pvc-43d55ea1-95c0-42c8-8616-4ee70b504445",
 "csi.storage.k8s.io/pvc/name": "cephfs-pvc-restore",
 "csi.storage.k8s.io/pvc/namespace": "openshift-storage"
}










 27

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

     CHAPTER 4. ALERTS

4.1. SETTING UP ALERTS

For internal Mode clusters, various alerts related to the storage metrics services, storage cluster, disk
devices, cluster health, cluster capacity, and so on are displayed in the Block and File, and the object
dashboards. These alerts are not available for external Mode.

 NOTE

 It might take a few minutes for alerts to be shown in the alert panel, because only firing
 alerts are visible in this panel.

You can also view alerts with additional details and customize the display of Alerts in the OpenShift
Container Platform.

For more information, see Managing alerts.










28

        CHAPTER 5. REMOTE HEALTH MONITORING

        CHAPTER 5. REMOTE HEALTH MONITORING

OpenShift Data Foundation collects anonymized aggregated information about the health, usage, and
size of clusters and reports it to Red Hat via an integrated component called Telemetry. This information
allows Red Hat to improve OpenShift Data Foundation and to react to issues that impact customers
more quickly.

A cluster that reports data to Red Hat via Telemetry is considered a connected cluster.

5.1. ABOUT TELEMETRY

Telemetry sends a carefully chosen subset of the cluster monitoring metrics to Red Hat. These metrics
are sent continuously and describe:

        The size of an OpenShift Data Foundation cluster

        The health and status of OpenShift Data Foundation components

        The health and status of any upgrade being performed

        Limited usage information about OpenShift Data Foundation components and features

        Summary info about alerts reported by the cluster monitoring component

This continuous stream of data is used by Red Hat to monitor the health of clusters in real time and to
react as necessary to problems that impact our customers. It also allows Red Hat to roll out OpenShift
Data Foundation upgrades to customers so as to minimize service impact and continuously improve the
upgrade experience.

This debugging information is available to Red Hat Support and engineering teams with the same
restrictions as accessing data reported via support cases. All connected cluster information is used by
Red Hat to help make OpenShift Data Foundation better and more intuitive to use. None of the
information is shared with third parties.

5.2. INFORMATION COLLECTED BY TELEMETRY

Primary information collected by Telemetry includes:

        The size of the Ceph cluster in bytes : "ceph_cluster_total_bytes",

        The amount of the Ceph cluster storage used in bytes :
        "ceph_cluster_total_used_raw_bytes",

        Ceph cluster health status : "ceph_health_status",

        The total count of object storage devices (OSDs) : "job:ceph_osd_metadata:count",

        The total number of OpenShift Data Foundation Persistent Volumes (PVs) present in the Red
        Hat OpenShift Container Platform cluster : "job:kube_pv:count",

        The total input/output operations per second (IOPS) (reads+writes) value for all the pools in the
        Ceph cluster : "job:ceph_pools_iops:total",

        The total IOPS (reads+writes) value in bytes for all the pools in the Ceph cluster :
        "job:ceph_pools_iops_bytes:total",


    29

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

        The total count of the Ceph cluster versions running : "job:ceph_versions_running:count"

        The total number of unhealthy NooBaa buckets :
        "job:noobaa_total_unhealthy_buckets:sum",

        The total number of NooBaa buckets : "job:noobaa_bucket_count:sum",

        The total number of NooBaa objects : "job:noobaa_total_object_count:sum",

        The count of NooBaa accounts : "noobaa_accounts_num",

        The total usage of storage by NooBaa in bytes : "noobaa_total_usage",

        The total amount of storage requested by the persistent volume claims (PVCs) from a particular
        storage provisioner in bytes:
        "cluster:kube_persistentvolumeclaim_resource_requests_storage_bytes:provisioner:su
        m",

        The total amount of storage used by the PVCs from a particular storage provisioner in bytes:
        "cluster:kubelet_volume_stats_used_bytes:provisioner:sum".

  Telemetry does not collect identifying information such as user names, passwords, or the names or
  addresses of user resources.










    30

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

To improve observability and make metric interpretation easier, this section outlines a reference of
ceph-exporter, rook, and noobaa metrics and their functional significance.

6.1. RBD / MIRRORING

ocs_mirror_daemon_count
Mirror Daemon Count.
ocs_pool_mirroring_status
Pool Mirroring Status. 0=Disabled, 1=Enabled
ocs_rbd_client_blocklisted
State of the rbd client on a node, 0 = Unblocked, 1 = Blocked
ocs_rbd_pv_metadata
Attributes of Ceph RBD based Persistent Volume

6.2. RGW

ocs_rgw_health_status
Health Status of RGW Endpoint. 0=Connected, 1=Progressing & 2=Failure

6.3. STORAGE CLIENT/PROVIDER

ocs_storage_client_last_heartbeat
Unixtime (in sec) of last heartbeat of OCS Storage Client
ocs_storage_client_operator_version
OCS StorageClient encoded Operator Version
ocs_storage_client_storage_quota_utilization_ratio
StorageQuotaUtilizationRatio of ODF Storage Client
ocs_storage_consumer_metadata
Attributes of OCS Storage Consumers
ocs_storage_provider_operator_version
OCS StorageProvider encode Operator Version

6.4. STORAGECLUSTER

ocs_storagecluster_failure_domain_count
 Count of failure domains for StorageCluster with given name and namespace
ocs_storagecluster_kms_connection_status
 KMS Connection Status; 0: Connected, 1: Not Connected, 2: KMS not enabled

6.5. PROMETHEUS / HTTP HANDLER

promhttp_metric_handler_errors_total


31

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Total number of internal errors encountered by the promhttp metric handler.

6.6. CEPH METRICS

ceph_AsyncMessenger_Worker_msgr_connection_idle_timeouts
 Number of connections closed due to idleness
ceph_AsyncMessenger_Worker_msgr_connection_ready_timeouts
    Number of not yet ready connections declared as dead
ceph_blk_kernel_device_bluestore_discard_op
Number of discard ops issued to kernel device
ceph_blk_kernel_device_bluestore_discard_threads
Number of discard threads running
ceph_blk_kernel_device_db_discard_op
Number of discard ops issued to kernel device
ceph_blk_kernel_device_db_discard_threads
Number of discard threads running
ceph_bluefs_alloc_db_max_lat
 Max allocation latency for db device
ceph_bluefs_alloc_slow_fallback
Amount of allocations that required fallback to slow/shared device
ceph_bluefs_alloc_slow_max_lat
Max allocation latency for primary/shared device
ceph_bluefs_alloc_slow_size_fallback
Amount of allocations that required fallback to shared device’s regular unit size
ceph_bluefs_alloc_unit_db
Allocation unit size (in bytes) for standalone DB device
ceph_bluefs_alloc_unit_slow
Allocation unit size (in bytes) for primary/shared device
ceph_bluefs_alloc_unit_wal
Allocation unit size (in bytes) for standalone WAL device
ceph_bluefs_alloc_wal_max_lat
Max allocation latency for wal device
ceph_bluefs_bytes_written_slow
Bytes written to WAL/SSTs at slow device
ceph_bluefs_bytes_written_sst
Bytes written to SSTs
ceph_bluefs_bytes_written_wal
 Bytes written to WAL
ceph_bluefs_compact_lat_count
  Average bluefs log compaction latency Count
ceph_bluefs_compact_lat_sum


32

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Average bluefs log compaction latency Total
ceph_bluefs_compact_lock_lat_count
Average lock duration while compacting bluefs log Count
ceph_bluefs_compact_lock_lat_sum
Average lock duration while compacting bluefs log Total
ceph_bluefs_db_alloc_lat_count
Average bluefs db allocate latency Count
ceph_bluefs_db_alloc_lat_sum
Average bluefs db allocate latency Total
ceph_bluefs_db_total_bytes
Total bytes (main db device)
ceph_bluefs_db_used_bytes
Used bytes (main db device)
ceph_bluefs_flush_lat_count
Average bluefs flush latency Count
ceph_bluefs_flush_lat_sum
Average bluefs flush latency Total
ceph_bluefs_fsync_lat_count
Average bluefs fsync latency Count
ceph_bluefs_fsync_lat_sum
Average bluefs fsync latency Total
ceph_bluefs_log_bytes
Size of the metadata log
ceph_bluefs_logged_bytes
Bytes written to the metadata log
ceph_bluefs_max_bytes_db
 Maximum bytes allocated from DB
ceph_bluefs_max_bytes_slow
Maximum bytes allocated from SLOW
ceph_bluefs_max_bytes_wal
Maximum bytes allocated from WAL
ceph_bluefs_num_files
File count
ceph_bluefs_read_bytes
 Bytes requested in buffered read mode
ceph_bluefs_read_count
 buffered read requests processed
ceph_bluefs_read_disk_bytes
Bytes read in buffered mode from disk
ceph_bluefs_read_disk_bytes_db


33

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

reads requests going to DB disk
ceph_bluefs_read_disk_bytes_slow
reads requests going to main disk
ceph_bluefs_read_disk_bytes_wal
 reads requests going to WAL disk
ceph_bluefs_read_disk_count
buffered reads requests going to disk
ceph_bluefs_read_lat_count
Average bluefs read latency Count
ceph_bluefs_read_lat_sum
Average bluefs read latency Total
ceph_bluefs_read_prefetch_bytes
Bytes requested in prefetch read mode
ceph_bluefs_read_prefetch_count
 prefetch read requests processed
ceph_bluefs_read_random_buffer_bytes
Bytes read from prefetch buffer in random read mode
ceph_bluefs_read_random_buffer_count
random read requests processed using prefetch buffer
ceph_bluefs_read_random_bytes
Bytes requested in random read mode
ceph_bluefs_read_random_count
  random read requests processed
ceph_bluefs_read_random_disk_bytes
Bytes read from disk in random read mode
ceph_bluefs_read_random_disk_bytes_db
random reads requests going to DB disk
ceph_bluefs_read_random_disk_bytes_slow
random reads requests going to main disk
ceph_bluefs_read_random_disk_bytes_wal
random reads requests going to WAL disk
ceph_bluefs_read_random_disk_count
 random reads requests going to disk
ceph_bluefs_read_random_lat_count
Average bluefs read_random latency Count
ceph_bluefs_read_random_lat_sum
Average bluefs read_random latency Total
ceph_bluefs_slow_alloc_lat_count
Average allocation latency for primary/shared device Count
ceph_bluefs_slow_alloc_lat_sum


34

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Average allocation latency for primary/shared device Total
ceph_bluefs_slow_total_bytes
Total bytes (slow device)
ceph_bluefs_slow_used_bytes
 Used bytes (slow device)
ceph_bluefs_truncate_lat_count
   Average bluefs truncate latency Count
ceph_bluefs_truncate_lat_sum
Average bluefs truncate latency Total
ceph_bluefs_unlink_lat_count
Average bluefs unlink latency Count
ceph_bluefs_unlink_lat_sum
Average bluefs unlink latency Total
ceph_bluefs_wal_alloc_lat_count
Average bluefs wal allocate latency Count
ceph_bluefs_wal_alloc_lat_sum
Average bluefs wal allocate latency Total
ceph_bluefs_wal_total_bytes
Total bytes (wal device)
ceph_bluefs_wal_used_bytes
 Used bytes (wal device)
ceph_bluefs_write_bytes
Bytes written
ceph_bluestore_alloc_unit
allocation unit size in bytes
ceph_bluestore_allocated
 Sum for allocated bytes
ceph_bluestore_allocator_lat_count
Average bluestore allocator latency Count
ceph_bluestore_allocator_lat_sum
Average bluestore allocator latency Total
ceph_bluestore_clist_lat_count
Average collection listing latency Count
ceph_bluestore_clist_lat_sum
Average collection listing latency Total
ceph_bluestore_compress_lat_count
Average compress latency Count
ceph_bluestore_compress_lat_sum
Average compress latency Total
ceph_bluestore_compressed


35

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Sum for stored compressed bytes
ceph_bluestore_compressed_allocated
Sum for bytes allocated for compressed data
ceph_bluestore_compressed_original
Sum for original bytes that were compressed
ceph_bluestore_csum_lat_count
Average checksum latency Count
ceph_bluestore_csum_lat_sum
Average checksum latency Total
ceph_bluestore_decompress_lat_count
Average decompress latency Count
ceph_bluestore_decompress_lat_sum
Average decompress latency Total
ceph_bluestore_fragmentation_micros
How fragmented bluestore free space is (free extents / max possible number of free extents) * 1000
ceph_bluestore_kv_commit_lat_count
Average kv_thread commit latency Count
ceph_bluestore_kv_commit_lat_sum
  Average kv_thread commit latency Total
ceph_bluestore_kv_final_lat_count
Average kv_finalize thread latency Count
ceph_bluestore_kv_final_lat_sum
Average kv_finalize thread latency Total
ceph_bluestore_kv_flush_lat_count
Average kv_thread flush latency Count
ceph_bluestore_kv_flush_lat_sum
Average kv_thread flush latency Total
ceph_bluestore_kv_sync_lat_count
 Average kv_sync thread latency Count
ceph_bluestore_kv_sync_lat_sum
 Average kv_sync thread latency Total
ceph_bluestore_omap_get_keys_lat_count
Average omap get_keys call latency Count
ceph_bluestore_omap_get_keys_lat_sum
Average omap get_keys call latency Total
ceph_bluestore_omap_get_values_lat_count
 Average omap get_values call latency Count
ceph_bluestore_omap_get_values_lat_sum
Average omap get_values call latency Total
ceph_bluestore_omap_lower_bound_lat_count


36

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

                        Average omap iterator lower_bound call latency Count
ceph_bluestore_omap_lower_bound_lat_sum
Average omap iterator lower_bound call latency Total
ceph_bluestore_omap_next_lat_count
Average omap iterator next call latency Count
ceph_bluestore_omap_next_lat_sum
Average omap iterator next call latency Total
ceph_bluestore_omap_seek_to_first_lat_count
Average omap iterator seek_to_first call latency Count
ceph_bluestore_omap_seek_to_first_lat_sum
Average omap iterator seek_to_first call latency Total
ceph_bluestore_omap_upper_bound_lat_count
  Average omap iterator upper_bound call latency Count
ceph_bluestore_omap_upper_bound_lat_sum
                        Average omap iterator upper_bound call latency Total
ceph_bluestore_onode_hits
Count of onode cache lookup hits
ceph_bluestore_onode_misses
Count of onode cache lookup misses
ceph_bluestore_pricache:data_committed_bytes
  total bytes committed,
ceph_bluestore_pricache:data_pri0_bytes
 bytes allocated to pri0
ceph_bluestore_pricache:data_pri10_bytes
bytes allocated to pri10
ceph_bluestore_pricache:data_pri11_bytes
bytes allocated to pri11
ceph_bluestore_pricache:data_pri1_bytes
bytes allocated to pri1
ceph_bluestore_pricache:data_pri2_bytes
 bytes allocated to pri2
ceph_bluestore_pricache:data_pri3_bytes
 bytes allocated to pri3
ceph_bluestore_pricache:data_pri4_bytes
 bytes allocated to pri4
ceph_bluestore_pricache:data_pri5_bytes
 bytes allocated to pri5
ceph_bluestore_pricache:data_pri6_bytes
 bytes allocated to pri6
ceph_bluestore_pricache:data_pri7_bytes


37

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

 bytes allocated to pri7
ceph_bluestore_pricache:data_pri8_bytes
 bytes allocated to pri8
ceph_bluestore_pricache:data_pri9_bytes
 bytes allocated to pri9
ceph_bluestore_pricache:data_reserved_bytes
bytes reserved for future growth.
ceph_bluestore_pricache:kv_committed_bytes
  total bytes committed,
ceph_bluestore_pricache:kv_onode_committed_bytes
  total bytes committed,
ceph_bluestore_pricache:kv_onode_pri0_bytes
bytes allocated to pri0
ceph_bluestore_pricache:kv_onode_pri10_bytes
bytes allocated to pri10
ceph_bluestore_pricache:kv_onode_pri11_bytes
bytes allocated to pri11
ceph_bluestore_pricache:kv_onode_pri1_bytes
bytes allocated to pri1
ceph_bluestore_pricache:kv_onode_pri2_bytes
 bytes allocated to pri2
ceph_bluestore_pricache:kv_onode_pri3_bytes
 bytes allocated to pri3
ceph_bluestore_pricache:kv_onode_pri4_bytes
 bytes allocated to pri4
ceph_bluestore_pricache:kv_onode_pri5_bytes
 bytes allocated to pri5
ceph_bluestore_pricache:kv_onode_pri6_bytes
 bytes allocated to pri6
ceph_bluestore_pricache:kv_onode_pri7_bytes
 bytes allocated to pri7
ceph_bluestore_pricache:kv_onode_pri8_bytes
 bytes allocated to pri8
ceph_bluestore_pricache:kv_onode_pri9_bytes
 bytes allocated to pri9
ceph_bluestore_pricache:kv_onode_reserved_bytes
bytes reserved for future growth.
ceph_bluestore_pricache:kv_pri0_bytes
 bytes allocated to pri0
ceph_bluestore_pricache:kv_pri10_bytes


38

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

bytes allocated to pri10
ceph_bluestore_pricache:kv_pri11_bytes
bytes allocated to pri11
ceph_bluestore_pricache:kv_pri1_bytes
bytes allocated to pri1
ceph_bluestore_pricache:kv_pri2_bytes
 bytes allocated to pri2
ceph_bluestore_pricache:kv_pri3_bytes
 bytes allocated to pri3
ceph_bluestore_pricache:kv_pri4_bytes
 bytes allocated to pri4
ceph_bluestore_pricache:kv_pri5_bytes
 bytes allocated to pri5
ceph_bluestore_pricache:kv_pri6_bytes
 bytes allocated to pri6
ceph_bluestore_pricache:kv_pri7_bytes
 bytes allocated to pri7
ceph_bluestore_pricache:kv_pri8_bytes
 bytes allocated to pri8
ceph_bluestore_pricache:kv_pri9_bytes
 bytes allocated to pri9
ceph_bluestore_pricache:kv_reserved_bytes
bytes reserved for future growth.
ceph_bluestore_pricache:meta_committed_bytes
  total bytes committed,
ceph_bluestore_pricache:meta_pri0_bytes
 bytes allocated to pri0
ceph_bluestore_pricache:meta_pri10_bytes
bytes allocated to pri10
ceph_bluestore_pricache:meta_pri11_bytes
bytes allocated to pri11
ceph_bluestore_pricache:meta_pri1_bytes
bytes allocated to pri1
ceph_bluestore_pricache:meta_pri2_bytes
 bytes allocated to pri2
ceph_bluestore_pricache:meta_pri3_bytes
 bytes allocated to pri3
ceph_bluestore_pricache:meta_pri4_bytes
 bytes allocated to pri4
ceph_bluestore_pricache:meta_pri5_bytes


39

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

bytes allocated to pri5
ceph_bluestore_pricache:meta_pri6_bytes
bytes allocated to pri6
ceph_bluestore_pricache:meta_pri7_bytes
bytes allocated to pri7
ceph_bluestore_pricache:meta_pri8_bytes
bytes allocated to pri8
ceph_bluestore_pricache:meta_pri9_bytes
bytes allocated to pri9
ceph_bluestore_pricache:meta_reserved_bytes
bytes reserved for future growth.
ceph_bluestore_pricache_cache_bytes
current memory available for caches.
ceph_bluestore_pricache_heap_bytes
aggregate bytes in use by the heap
ceph_bluestore_pricache_mapped_bytes
total bytes mapped by the process
ceph_bluestore_pricache_target_bytes
target process memory usage in bytes
ceph_bluestore_pricache_unmapped_bytes
unmapped bytes that the kernel has yet to reclaim
ceph_bluestore_read_lat_count
Average read latency Count
ceph_bluestore_read_lat_sum
Average read latency Total
ceph_bluestore_read_onode_meta_lat_count
Average read onode metadata latency Count
ceph_bluestore_read_onode_meta_lat_sum
Average read onode metadata latency Total
ceph_bluestore_read_wait_aio_lat_count
Average read I/O waiting latency Count
ceph_bluestore_read_wait_aio_lat_sum
Average read I/O waiting latency Total
ceph_bluestore_reads_with_retries
Read operations that required at least one retry due to failed checksum validation
ceph_bluestore_remove_lat_count
Average removal latency Count
ceph_bluestore_remove_lat_sum
Average removal latency Total
ceph_bluestore_slow_aio_wait_count


40

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Slow op count for aio wait
ceph_bluestore_slow_committed_kv_count
 Slow op count for committed kv
ceph_bluestore_slow_read_onode_meta_count
Slow op count for read onode meta
ceph_bluestore_slow_read_wait_aio_count
Slow op count for read wait aio
ceph_bluestore_state_aio_wait_lat_count
Average aio_wait state latency Count
ceph_bluestore_state_aio_wait_lat_sum
Average aio_wait state latency Total
ceph_bluestore_state_deferred_aio_wait_lat_count
Average aio_wait state latency Count
ceph_bluestore_state_deferred_aio_wait_lat_sum
Average aio_wait state latency Total
ceph_bluestore_state_deferred_cleanup_lat_count
 Average cleanup state latency Count
ceph_bluestore_state_deferred_cleanup_lat_sum
Average cleanup state latency Total
ceph_bluestore_state_deferred_queued_lat_count
Average deferred_queued state latency Count
ceph_bluestore_state_deferred_queued_lat_sum
Average deferred_queued state latency Total
ceph_bluestore_state_done_lat_count
 Average done state latency Count
ceph_bluestore_state_done_lat_sum
Average done state latency Total
ceph_bluestore_state_finishing_lat_count
Average finishing state latency Count
ceph_bluestore_state_finishing_lat_sum
Average finishing state latency Total
ceph_bluestore_state_io_done_lat_count
 Average io_done state latency Count
ceph_bluestore_state_io_done_lat_sum
 Average io_done state latency Total
ceph_bluestore_state_kv_commiting_lat_count
Average kv_commiting state latency Count
ceph_bluestore_state_kv_commiting_lat_sum
Average kv_commiting state latency Total
ceph_bluestore_state_kv_done_lat_count


41

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

  Average kv_done state latency Count
ceph_bluestore_state_kv_done_lat_sum
  Average kv_done state latency Total
ceph_bluestore_state_kv_queued_lat_count
  Average kv_queued state latency Count
ceph_bluestore_state_kv_queued_lat_sum
  Average kv_queued state latency Total
ceph_bluestore_state_prepare_lat_count
  Average prepare state latency Count
ceph_bluestore_state_prepare_lat_sum
  Average prepare state latency Total
ceph_bluestore_stored
  Sum for stored bytes
ceph_bluestore_truncate_lat_count
  Average truncate latency Count
ceph_bluestore_truncate_lat_sum
  Average truncate latency Total
ceph_bluestore_txc_commit_lat_count
  Average commit latency Count
ceph_bluestore_txc_commit_lat_sum
  Average commit latency Total
ceph_bluestore_txc_submit_lat_count
  Average submit latency Count
ceph_bluestore_txc_submit_lat_sum
  Average submit latency Total
ceph_bluestore_txc_throttle_lat_count
  Average submit throttle latency Count
ceph_bluestore_txc_throttle_lat_sum
  Average submit throttle latency Total
ceph_daemon_socket_up
  Reports the health status of a Ceph daemon, as determined by whether it is able to respond via its
  admin socket (1 = healthy, 0 = unhealthy).
ceph_exporter_scrape_time
  Time spent scraping and transforming perf counters to metrics
ceph_mds_cache_ireq_enqueue_scrub
  Internal Request type enqueue scrub
ceph_mds_cache_ireq_exportdir
  Internal Request type export dir
ceph_mds_cache_ireq_flush
  Internal Request type flush
ceph_mds_cache_ireq_fragmentdir

42

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Internal Request type fragmentdir
ceph_mds_cache_ireq_fragstats
Internal Request type frag stats
ceph_mds_cache_ireq_inodestats
Internal Request type inode stats
ceph_mds_cache_ireq_quiesce_inode
                         Internal Request type quiesce subvolume inode
ceph_mds_cache_ireq_quiesce_path
Internal Request type quiesce subvolume
ceph_mds_cache_num_recovering_enqueued
Files waiting for recovery
ceph_mds_cache_num_recovering_prioritized
                       Files waiting for recovery with elevated priority
ceph_mds_cache_num_recovering_processing
Files currently being recovered
ceph_mds_cache_num_strays
Stray dentries
ceph_mds_cache_num_strays_delayed
 Stray dentries delayed
ceph_mds_cache_num_strays_enqueuing
Stray dentries enqueuing for purge
ceph_mds_cache_recovery_completed
File recoveries completed
ceph_mds_cache_recovery_started
File recoveries started
ceph_mds_cache_strays_created
 Stray dentries created
ceph_mds_cache_strays_enqueued
 Stray dentries enqueued for purge
ceph_mds_cache_strays_migrated
Stray dentries migrated
ceph_mds_cache_strays_reintegrated
Stray dentries reintegrated
ceph_mds_caps
Capabilities
ceph_mds_ceph_cap_op_flush_ack
caps truncate notify
ceph_mds_ceph_cap_op_flushsnap_ack
caps truncate notify
ceph_mds_ceph_cap_op_grant


43

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Grant caps
ceph_mds_ceph_cap_op_revoke
Revoke caps
ceph_mds_ceph_cap_op_trunc
caps truncate notify
ceph_mds_client_metrics_num_clients
Numer of client sessions
ceph_mds_dir_commit
Directory commit
ceph_mds_dir_fetch_complete
 Fetch complete dirfrag
ceph_mds_dir_fetch_keys
Fetch keys from dirfrag
ceph_mds_dir_merge
Directory merge
ceph_mds_dir_split
Directory split
ceph_mds_exported_inodes
Exported inodes
ceph_mds_forward
Forwarding request
ceph_mds_handle_client_cap_release
Client cap release msg
ceph_mds_handle_client_caps
Client caps msg
ceph_mds_handle_client_caps_dirty
Client dirty caps msg
ceph_mds_handle_inode_file_caps
 Inter mds caps msg
ceph_mds_imported_inodes
Imported inodes
ceph_mds_inodes
Inodes
ceph_mds_inodes_expired
Inodes expired
ceph_mds_inodes_pinned
Inodes pinned
ceph_mds_inodes_with_caps
Inodes with capabilities
ceph_mds_load_cent


44

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Load per cent
ceph_mds_log_ev
Events
ceph_mds_log_evadd
Events submitted
ceph_mds_log_evex
Total expired events
ceph_mds_log_evexd
Current expired events
ceph_mds_log_evexg
Expiring events
ceph_mds_log_evlrg
 Large events
ceph_mds_log_evtrm
 Trimmed events
ceph_mds_log_jlat_count
Journaler flush latency Count
ceph_mds_log_jlat_sum
Journaler flush latency Total
ceph_mds_log_replayed
Events replayed
ceph_mds_log_seg
Segments
ceph_mds_log_segadd
Segments added
ceph_mds_log_segex
Total expired segments
ceph_mds_log_segexd
Current expired segments
ceph_mds_log_segexg
Expiring segments
ceph_mds_log_segmjr
Major Segments
ceph_mds_log_segtrm
  Trimmed segments
ceph_mds_mem_cap
Capabilities
ceph_mds_mem_cap_minus
Capabilities removed
ceph_mds_mem_cap_plus


45

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Capabilities added
ceph_mds_mem_dir
Directories
ceph_mds_mem_dir_minus
Directories closed
ceph_mds_mem_dir_plus
Directories opened
ceph_mds_mem_dn
Dentries
ceph_mds_mem_dn_minus
Dentries closed
ceph_mds_mem_dn_plus
 Dentries opened
ceph_mds_mem_heap
Heap size
ceph_mds_mem_ino
Inodes
ceph_mds_mem_ino_minus
Inodes closed
ceph_mds_mem_ino_plus
  Inodes opened
ceph_mds_mem_rss
RSS
ceph_mds_openino_dir_fetch
OpenIno incomplete directory fetchings
ceph_mds_process_request_cap_release
Process request cap release
ceph_mds_reply_latency_count
Reply latency Count
ceph_mds_reply_latency_sum
Reply latency Total
ceph_mds_request
 Requests
ceph_mds_root_rbytes
root inode rbytes
ceph_mds_root_rfiles
root inode rfiles
ceph_mds_root_rsnaps
root inode rsnaps
ceph_mds_server_cap_acquisition_throttle


46

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Cap acquisition throttle counter
ceph_mds_server_cap_revoke_eviction
Cap Revoke Client Eviction
ceph_mds_server_handle_client_request
Client requests
ceph_mds_server_handle_client_session
Client session messages
ceph_mds_server_handle_peer_request
Peer requests
ceph_mds_server_req_blockdiff_latency_count
Request type file blockdiff latency Count
ceph_mds_server_req_blockdiff_latency_sum
Request type file blockdiff latency Total
ceph_mds_server_req_create_latency_count
Request type create latency Count
ceph_mds_server_req_create_latency_sum
Request type create latency Total
ceph_mds_server_req_getattr_latency_count
 Request type get attribute latency Count
ceph_mds_server_req_getattr_latency_sum
 Request type get attribute latency Total
ceph_mds_server_req_getfilelock_latency_count
Request type get file lock latency Count
ceph_mds_server_req_getfilelock_latency_sum
Request type get file lock latency Total
ceph_mds_server_req_getvxattr_latency_count
Request type get virtual extended attribute latency Count
ceph_mds_server_req_getvxattr_latency_sum
Request type get virtual extended attribute latency Total
ceph_mds_server_req_link_latency_count
Request type link latency Count
ceph_mds_server_req_link_latency_sum
Request type link latency Total
ceph_mds_server_req_lookup_latency_count
Request type lookup latency Count
ceph_mds_server_req_lookup_latency_sum
Request type lookup latency Total
ceph_mds_server_req_lookuphash_latency_count
Request type lookup hash of inode latency Count
ceph_mds_server_req_lookuphash_latency_sum


47

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Request type lookup hash of inode latency Total
ceph_mds_server_req_lookupino_latency_count
 Request type lookup inode latency Count
ceph_mds_server_req_lookupino_latency_sum
 Request type lookup inode latency Total
ceph_mds_server_req_lookupname_latency_count
  Request type lookup name latency Count
ceph_mds_server_req_lookupname_latency_sum
  Request type lookup name latency Total
ceph_mds_server_req_lookupparent_latency_count
Request type lookup parent latency Count
ceph_mds_server_req_lookupparent_latency_sum
Request type lookup parent latency Total
ceph_mds_server_req_lookupsnap_latency_count
Request type lookup snapshot latency Count
ceph_mds_server_req_lookupsnap_latency_sum
Request type lookup snapshot latency Total
ceph_mds_server_req_lssnap_latency_count
Request type list snapshot latency Count
ceph_mds_server_req_lssnap_latency_sum
Request type list snapshot latency Total
ceph_mds_server_req_mkdir_latency_count
Request type make directory latency Count
ceph_mds_server_req_mkdir_latency_sum
Request type make directory latency Total
ceph_mds_server_req_mknod_latency_count
    Request type make node latency Count
ceph_mds_server_req_mknod_latency_sum
Request type make node latency Total
ceph_mds_server_req_mksnap_latency_count
  Request type make snapshot latency Count
ceph_mds_server_req_mksnap_latency_sum
 Request type make snapshot latency Total
ceph_mds_server_req_open_latency_count
Request type open latency Count
ceph_mds_server_req_open_latency_sum
Request type open latency Total
ceph_mds_server_req_readdir_latency_count
Request type read directory latency Count
ceph_mds_server_req_readdir_latency_sum


48

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Request type read directory latency Total
ceph_mds_server_req_rename_latency_count
Request type rename latency Count
ceph_mds_server_req_rename_latency_sum
Request type rename latency Total
ceph_mds_server_req_renamesnap_latency_count
 Request type rename snapshot latency Count
ceph_mds_server_req_renamesnap_latency_sum
 Request type rename snapshot latency Total
ceph_mds_server_req_rmdir_latency_count
Request type remove directory latency Count
ceph_mds_server_req_rmdir_latency_sum
Request type remove directory latency Total
ceph_mds_server_req_rmsnap_latency_count
 Request type remove snapshot latency Count
ceph_mds_server_req_rmsnap_latency_sum
 Request type remove snapshot latency Total
ceph_mds_server_req_rmxattr_latency_count
Request type remove extended attribute latency Count
ceph_mds_server_req_rmxattr_latency_sum
Request type remove extended attribute latency Total
ceph_mds_server_req_setattr_latency_count
  Request type set attribute latency Count
ceph_mds_server_req_setattr_latency_sum
Request type set attribute latency Total
ceph_mds_server_req_setdirlayout_latency_count
Request type set directory layout latency Count
ceph_mds_server_req_setdirlayout_latency_sum
Request type set directory layout latency Total
ceph_mds_server_req_setfilelock_latency_count
Request type set file lock latency Count
ceph_mds_server_req_setfilelock_latency_sum
Request type set file lock latency Total
ceph_mds_server_req_setlayout_latency_count
Request type set file layout latency Count
ceph_mds_server_req_setlayout_latency_sum
Request type set file layout latency Total
ceph_mds_server_req_setxattr_latency_count
Request type set extended attribute latency Count
ceph_mds_server_req_setxattr_latency_sum


49

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Request type set extended attribute latency Total
ceph_mds_server_req_snapdiff_latency_count
Request type snapshot difference latency Count
ceph_mds_server_req_snapdiff_latency_sum
Request type snapshot difference latency Total
ceph_mds_server_req_symlink_latency_count
Request type symbolic link latency Count
ceph_mds_server_req_symlink_latency_sum
Request type symbolic link latency Total
ceph_mds_server_req_unlink_latency_count
Request type unlink latency Count
ceph_mds_server_req_unlink_latency_sum
Request type unlink latency Total
ceph_mds_sessions_average_load
Average Load
ceph_mds_sessions_avg_session_uptime
Average session uptime
ceph_mds_sessions_mdthresh_evicted
 Sessions evicted on reaching metadata threshold
ceph_mds_sessions_session_add
Sessions added
ceph_mds_sessions_session_count
Session count
ceph_mds_sessions_session_remove
Sessions removed
ceph_mds_sessions_sessions_open
Sessions currently open
ceph_mds_sessions_sessions_stale
Sessions currently stale
ceph_mds_sessions_total_load
Total Load
ceph_mds_slow_reply
Slow replies
ceph_mds_subtrees
Subtrees
ceph_mon_election_call
Elections started
ceph_mon_election_lose
Elections lost
ceph_mon_election_win


50

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Elections won
ceph_mon_num_elections
Elections participated in
ceph_mon_num_sessions
Open sessions
ceph_mon_session_add
Created sessions
ceph_mon_session_rm
 Removed sessions
ceph_mon_session_trim
 Trimmed sessions
ceph_objecter_op_active
Operations active
ceph_objecter_op_r
Read operations
ceph_objecter_op_rmw
Read-modify-write operations
ceph_objecter_op_w
Write operations
ceph_osd_numpg
 Placement groups
ceph_osd_numpg_removing
Placement groups queued for local deletion
ceph_osd_op
Client operations
ceph_osd_op_before_queue_op_lat_count
Latency of IO before calling queue(before really queue into ShardedOpWq) Count
ceph_osd_op_before_queue_op_lat_sum
Latency of IO before calling queue(before really queue into ShardedOpWq) Total
ceph_osd_op_delayed_degraded
Count of ops delayed due to target object being degraded
ceph_osd_op_delayed_unreadable
Count of ops delayed due to target object being unreadable
ceph_osd_op_in_bytes
Client operations total write size
ceph_osd_op_latency_count
Latency of client operations (including queue time) Count
ceph_osd_op_latency_sum
Latency of client operations (including queue time) Total
ceph_osd_op_out_bytes


51

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Client operations total read size
ceph_osd_op_prepare_latency_count
Latency of client operations (excluding queue time and wait for finished) Count
ceph_osd_op_prepare_latency_sum
Latency of client operations (excluding queue time and wait for finished) Total
ceph_osd_op_process_latency_count
Latency of client operations (excluding queue time) Count
ceph_osd_op_process_latency_sum
Latency of client operations (excluding queue time) Total
ceph_osd_op_r
Client read operations
ceph_osd_op_r_latency_count
Latency of read operation (including queue time) Count
ceph_osd_op_r_latency_sum
Latency of read operation (including queue time) Total
ceph_osd_op_r_out_bytes
Client data read
ceph_osd_op_r_prepare_latency_count
Latency of read operations (excluding queue time and wait for finished) Count
ceph_osd_op_r_prepare_latency_sum
Latency of read operations (excluding queue time and wait for finished) Total
ceph_osd_op_r_process_latency_count
Latency of read operation (excluding queue time) Count
ceph_osd_op_r_process_latency_sum
Latency of read operation (excluding queue time) Total
ceph_osd_op_rw
Client read-modify-write operations
ceph_osd_op_rw_in_bytes
Client read-modify-write operations write in
ceph_osd_op_rw_latency_count
Latency of read-modify-write operation (including queue time) Count
ceph_osd_op_rw_latency_sum
Latency of read-modify-write operation (including queue time) Total
ceph_osd_op_rw_out_bytes
Client read-modify-write operations read out
ceph_osd_op_rw_prepare_latency_count
Latency of read-modify-write operations (excluding queue time and wait for finished) Count
ceph_osd_op_rw_prepare_latency_sum
Latency of read-modify-write operations (excluding queue time and wait for finished) Total
ceph_osd_op_rw_process_latency_count


52

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Latency of read-modify-write operation (excluding queue time) Count
ceph_osd_op_rw_process_latency_sum
Latency of read-modify-write operation (excluding queue time) Total
ceph_osd_op_w
Client write operations
ceph_osd_op_w_in_bytes
Client data written
ceph_osd_op_w_latency_count
Latency of write operation (including queue time) Count
ceph_osd_op_w_latency_sum
Latency of write operation (including queue time) Total
ceph_osd_op_w_prepare_latency_count
Latency of write operations (excluding queue time and wait for finished) Count
ceph_osd_op_w_prepare_latency_sum
Latency of write operations (excluding queue time and wait for finished) Total
ceph_osd_op_w_process_latency_count
Latency of write operation (excluding queue time) Count
ceph_osd_op_w_process_latency_sum
Latency of write operation (excluding queue time) Total
ceph_osd_op_wip
Replication operations currently being processed (primary)
ceph_osd_recovery_bytes
recovery bytes
ceph_osd_recovery_ops
Started recovery operations
ceph_osd_scrub_dp_ec_chunk_busy
  chunk busy during scrubs
ceph_osd_scrub_dp_ec_chunk_selected
chunk selection during scrubs
ceph_osd_scrub_dp_ec_failed_reservations_elapsed_count
time for scrub reservation to fail Count
ceph_osd_scrub_dp_ec_failed_reservations_elapsed_sum
time for scrub reservation to fail Total
ceph_osd_scrub_dp_ec_failed_scrubs
failed scrubs count
ceph_osd_scrub_dp_ec_failed_scrubs_elapsed_count
time to scrub failure Count
ceph_osd_scrub_dp_ec_failed_scrubs_elapsed_sum
time to scrub failure Total
ceph_osd_scrub_dp_ec_locked_object


53

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

waiting on locked object events
ceph_osd_scrub_dp_ec_num_scrubs_past_reservation
scrubs count
ceph_osd_scrub_dp_ec_num_scrubs_started
 scrubs attempted count
ceph_osd_scrub_dp_ec_preemptions
 preemptions on scrubs
ceph_osd_scrub_dp_ec_replicas_in_reservation
number of replicas in reservation
ceph_osd_scrub_dp_ec_reservation_process_aborted
scrub reservation was aborted
ceph_osd_scrub_dp_ec_reservation_process_failure
scrub reservation failed due to replica denial
ceph_osd_scrub_dp_ec_reservation_process_skipped
scrub reservation skipped for high priority scrub
ceph_osd_scrub_dp_ec_scrub_reservations_completed
successfully completed reservation processes
ceph_osd_scrub_dp_ec_successful_reservations_elapsed_count
    time to scrub reservation completion Count
ceph_osd_scrub_dp_ec_successful_reservations_elapsed_sum
time to scrub reservation completion Total
ceph_osd_scrub_dp_ec_successful_scrubs
successful scrubs count
ceph_osd_scrub_dp_ec_successful_scrubs_elapsed_count
  time to scrub completion Count
ceph_osd_scrub_dp_ec_successful_scrubs_elapsed_sum
 time to scrub completion Total
ceph_osd_scrub_dp_ec_write_blocked_by_scrub
write blocked by scrub
ceph_osd_scrub_dp_repl_chunk_busy
chunk busy during scrubs
ceph_osd_scrub_dp_repl_chunk_selected
chunk selection during scrubs
ceph_osd_scrub_dp_repl_failed_reservations_elapsed_count
time for scrub reservation to fail Count
ceph_osd_scrub_dp_repl_failed_reservations_elapsed_sum
time for scrub reservation to fail Total
ceph_osd_scrub_dp_repl_failed_scrubs
failed scrubs count
ceph_osd_scrub_dp_repl_failed_scrubs_elapsed_count


54

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

time to scrub failure Count
ceph_osd_scrub_dp_repl_failed_scrubs_elapsed_sum
time to scrub failure Total
ceph_osd_scrub_dp_repl_locked_object
waiting on locked object events
ceph_osd_scrub_dp_repl_num_scrubs_past_reservation
scrubs count
ceph_osd_scrub_dp_repl_num_scrubs_started
 scrubs attempted count
ceph_osd_scrub_dp_repl_preemptions
 preemptions on scrubs
ceph_osd_scrub_dp_repl_replicas_in_reservation
number of replicas in reservation
ceph_osd_scrub_dp_repl_reservation_process_aborted
scrub reservation was aborted
ceph_osd_scrub_dp_repl_reservation_process_failure
scrub reservation failed due to replica denial
ceph_osd_scrub_dp_repl_reservation_process_skipped
scrub reservation skipped for high priority scrub
ceph_osd_scrub_dp_repl_scrub_reservations_completed
successfully completed reservation processes
ceph_osd_scrub_dp_repl_successful_reservations_elapsed_count
    time to scrub reservation completion Count
ceph_osd_scrub_dp_repl_successful_reservations_elapsed_sum
time to scrub reservation completion Total
ceph_osd_scrub_dp_repl_successful_scrubs
successful scrubs count
ceph_osd_scrub_dp_repl_successful_scrubs_elapsed_count
  time to scrub completion Count
ceph_osd_scrub_dp_repl_successful_scrubs_elapsed_sum
 time to scrub completion Total
ceph_osd_scrub_dp_repl_write_blocked_by_scrub
write blocked by scrub
ceph_osd_scrub_sh_ec_chunk_busy
  chunk busy during scrubs
ceph_osd_scrub_sh_ec_chunk_selected
chunk selection during scrubs
ceph_osd_scrub_sh_ec_failed_reservations_elapsed_count
time for scrub reservation to fail Count
ceph_osd_scrub_sh_ec_failed_reservations_elapsed_sum


55

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

time for scrub reservation to fail Total
ceph_osd_scrub_sh_ec_failed_scrubs
failed scrubs count
ceph_osd_scrub_sh_ec_failed_scrubs_elapsed_count
time to scrub failure Count
ceph_osd_scrub_sh_ec_failed_scrubs_elapsed_sum
time to scrub failure Total
ceph_osd_scrub_sh_ec_locked_object
waiting on locked object events
ceph_osd_scrub_sh_ec_num_scrubs_past_reservation
scrubs count
ceph_osd_scrub_sh_ec_num_scrubs_started
 scrubs attempted count
ceph_osd_scrub_sh_ec_preemptions
 preemptions on scrubs
ceph_osd_scrub_sh_ec_replicas_in_reservation
number of replicas in reservation
ceph_osd_scrub_sh_ec_reservation_process_aborted
scrub reservation was aborted
ceph_osd_scrub_sh_ec_reservation_process_failure
scrub reservation failed due to replica denial
ceph_osd_scrub_sh_ec_reservation_process_skipped
scrub reservation skipped for high priority scrub
ceph_osd_scrub_sh_ec_scrub_reservations_completed
successfully completed reservation processes
ceph_osd_scrub_sh_ec_successful_reservations_elapsed_count
  time to scrub reservation completion Count
ceph_osd_scrub_sh_ec_successful_reservations_elapsed_sum
time to scrub reservation completion Total
ceph_osd_scrub_sh_ec_successful_scrubs
successful scrubs count
ceph_osd_scrub_sh_ec_successful_scrubs_elapsed_count
time to scrub completion Count
ceph_osd_scrub_sh_ec_successful_scrubs_elapsed_sum
time to scrub completion Total
ceph_osd_scrub_sh_ec_write_blocked_by_scrub
write blocked by scrub
ceph_osd_scrub_sh_repl_chunk_busy
  chunk busy during scrubs
ceph_osd_scrub_sh_repl_chunk_selected


56

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

chunk selection during scrubs
ceph_osd_scrub_sh_repl_failed_reservations_elapsed_count
time for scrub reservation to fail Count
ceph_osd_scrub_sh_repl_failed_reservations_elapsed_sum
time for scrub reservation to fail Total
ceph_osd_scrub_sh_repl_failed_scrubs
failed scrubs count
ceph_osd_scrub_sh_repl_failed_scrubs_elapsed_count
time to scrub failure Count
ceph_osd_scrub_sh_repl_failed_scrubs_elapsed_sum
time to scrub failure Total
ceph_osd_scrub_sh_repl_locked_object
waiting on locked object events
ceph_osd_scrub_sh_repl_num_scrubs_past_reservation
scrubs count
ceph_osd_scrub_sh_repl_num_scrubs_started
 scrubs attempted count
ceph_osd_scrub_sh_repl_preemptions
preemptions on scrubs
ceph_osd_scrub_sh_repl_replicas_in_reservation
number of replicas in reservation
ceph_osd_scrub_sh_repl_reservation_process_aborted
scrub reservation was aborted
ceph_osd_scrub_sh_repl_reservation_process_failure
scrub reservation failed due to replica denial
ceph_osd_scrub_sh_repl_reservation_process_skipped
scrub reservation skipped for high priority scrub
ceph_osd_scrub_sh_repl_scrub_reservations_completed
successfully completed reservation processes
ceph_osd_scrub_sh_repl_successful_reservations_elapsed_count
    time to scrub reservation completion Count
ceph_osd_scrub_sh_repl_successful_reservations_elapsed_sum
time to scrub reservation completion Total
ceph_osd_scrub_sh_repl_successful_scrubs
successful scrubs count
ceph_osd_scrub_sh_repl_successful_scrubs_elapsed_count
  time to scrub completion Count
ceph_osd_scrub_sh_repl_successful_scrubs_elapsed_sum
 time to scrub completion Total
ceph_osd_scrub_sh_repl_write_blocked_by_scrub


57

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

write blocked by scrub
ceph_osd_stat_bytes
OSD size
ceph_osd_stat_bytes_used
Used space
ceph_paxos_accept_timeout
Accept timeouts
ceph_paxos_begin
Started and handled begins
ceph_paxos_begin_bytes_count
Data in transaction on begin Count
ceph_paxos_begin_bytes_sum
Data in transaction on begin Total
ceph_paxos_begin_keys_count
Keys in transaction on begin Count
ceph_paxos_begin_keys_sum
Keys in transaction on begin Total
ceph_paxos_begin_latency_count
 Latency of begin operation Count
ceph_paxos_begin_latency_sum
 Latency of begin operation Total
ceph_paxos_collect
Peon collects
ceph_paxos_collect_bytes_count
Data in transaction on peon collect Count
ceph_paxos_collect_bytes_sum
Data in transaction on peon collect Total
ceph_paxos_collect_keys_count
Keys in transaction on peon collect Count
ceph_paxos_collect_keys_sum
Keys in transaction on peon collect Total
ceph_paxos_collect_latency_count
Peon collect latency Count
ceph_paxos_collect_latency_sum
Peon collect latency Total
ceph_paxos_collect_timeout
Collect timeouts
ceph_paxos_collect_uncommitted
Uncommitted values in started and handled collects
ceph_paxos_commit


58

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Commits
ceph_paxos_commit_bytes_count
Data in transaction on commit Count
ceph_paxos_commit_bytes_sum
Data in transaction on commit Total
ceph_paxos_commit_keys_count
Keys in transaction on commit Count
ceph_paxos_commit_keys_sum
Keys in transaction on commit Total
ceph_paxos_commit_latency_count
 Commit latency Count
ceph_paxos_commit_latency_sum
 Commit latency Total
ceph_paxos_lease_ack_timeout
Lease acknowledgement timeouts
ceph_paxos_lease_timeout
Lease timeouts
ceph_paxos_new_pn
New proposal number queries
ceph_paxos_new_pn_latency_count
                     New proposal number getting latency Count
ceph_paxos_new_pn_latency_sum
                     New proposal number getting latency Total
ceph_paxos_refresh
Refreshes
ceph_paxos_refresh_latency_count
Refresh latency Count
ceph_paxos_refresh_latency_sum
Refresh latency Total
ceph_paxos_restart
Restarts
ceph_paxos_share_state
Sharings of state
ceph_paxos_share_state_bytes_count
Data in shared state Count
ceph_paxos_share_state_bytes_sum
Data in shared state Total
ceph_paxos_share_state_keys_count
Keys in shared state Count
ceph_paxos_share_state_keys_sum


59

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Keys in shared state Total
ceph_paxos_start_leader
Starts in leader role
ceph_paxos_start_peon
Starts in peon role
ceph_paxos_store_state
Store a shared state on disk
ceph_paxos_store_state_bytes_count
Data in transaction in stored state Count
ceph_paxos_store_state_bytes_sum
Data in transaction in stored state Total
ceph_paxos_store_state_keys_count
Keys in transaction in stored state Count
ceph_paxos_store_state_keys_sum
Keys in transaction in stored state Total
ceph_paxos_store_state_latency_count
Storing state latency Count
ceph_paxos_store_state_latency_sum
Storing state latency Total
ceph_prioritycache:full_committed_bytes
  total bytes committed,
ceph_prioritycache:full_pri0_bytes
bytes allocated to pri0
ceph_prioritycache:full_pri10_bytes
bytes allocated to pri10
ceph_prioritycache:full_pri11_bytes
bytes allocated to pri11
ceph_prioritycache:full_pri1_bytes
bytes allocated to pri1
ceph_prioritycache:full_pri2_bytes
 bytes allocated to pri2
ceph_prioritycache:full_pri3_bytes
 bytes allocated to pri3
ceph_prioritycache:full_pri4_bytes
 bytes allocated to pri4
ceph_prioritycache:full_pri5_bytes
 bytes allocated to pri5
ceph_prioritycache:full_pri6_bytes
 bytes allocated to pri6
ceph_prioritycache:full_pri7_bytes


60

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

bytes allocated to pri7
ceph_prioritycache:full_pri8_bytes
bytes allocated to pri8
ceph_prioritycache:full_pri9_bytes
bytes allocated to pri9
ceph_prioritycache:full_reserved_bytes
bytes reserved for future growth.
ceph_prioritycache:inc_committed_bytes
 total bytes committed,
ceph_prioritycache:inc_pri0_bytes
 bytes allocated to pri0
ceph_prioritycache:inc_pri10_bytes
bytes allocated to pri10
ceph_prioritycache:inc_pri11_bytes
bytes allocated to pri11
ceph_prioritycache:inc_pri1_bytes
bytes allocated to pri1
ceph_prioritycache:inc_pri2_bytes
bytes allocated to pri2
ceph_prioritycache:inc_pri3_bytes
bytes allocated to pri3
ceph_prioritycache:inc_pri4_bytes
bytes allocated to pri4
ceph_prioritycache:inc_pri5_bytes
bytes allocated to pri5
ceph_prioritycache:inc_pri6_bytes
bytes allocated to pri6
ceph_prioritycache:inc_pri7_bytes
bytes allocated to pri7
ceph_prioritycache:inc_pri8_bytes
bytes allocated to pri8
ceph_prioritycache:inc_pri9_bytes
bytes allocated to pri9
ceph_prioritycache:inc_reserved_bytes
bytes reserved for future growth.
ceph_prioritycache:kv_committed_bytes
 total bytes committed,
ceph_prioritycache:kv_pri0_bytes
bytes allocated to pri0
ceph_prioritycache:kv_pri10_bytes


61

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

bytes allocated to pri10
ceph_prioritycache:kv_pri11_bytes
bytes allocated to pri11
ceph_prioritycache:kv_pri1_bytes
bytes allocated to pri1
ceph_prioritycache:kv_pri2_bytes
bytes allocated to pri2
ceph_prioritycache:kv_pri3_bytes
bytes allocated to pri3
ceph_prioritycache:kv_pri4_bytes
bytes allocated to pri4
ceph_prioritycache:kv_pri5_bytes
bytes allocated to pri5
ceph_prioritycache:kv_pri6_bytes
bytes allocated to pri6
ceph_prioritycache:kv_pri7_bytes
bytes allocated to pri7
ceph_prioritycache:kv_pri8_bytes
bytes allocated to pri8
ceph_prioritycache:kv_pri9_bytes
bytes allocated to pri9
ceph_prioritycache:kv_reserved_bytes
bytes reserved for future growth.
ceph_prioritycache_cache_bytes
current memory available for caches.
ceph_prioritycache_heap_bytes
aggregate bytes in use by the heap
ceph_prioritycache_mapped_bytes
 total bytes mapped by the process
ceph_prioritycache_target_bytes
target process memory usage in bytes
ceph_prioritycache_unmapped_bytes
unmapped bytes that the kernel has yet to reclaim
ceph_purge_queue_pq_executed
Purge queue tasks executed
ceph_purge_queue_pq_executed_ops
   Purge queue ops executed
ceph_purge_queue_pq_executing
Purge queue tasks in flight
ceph_purge_queue_pq_executing_high_water


62

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Maximum number of executing file purges
ceph_purge_queue_pq_executing_ops
Purge queue ops in flight
ceph_purge_queue_pq_executing_ops_high_water
Maximum number of executing file purge ops
ceph_purge_queue_pq_item_in_journal
Purge item left in journal
ceph_rocksdb_compact
Compactions
ceph_rocksdb_compact_completed
  Completed compactions
ceph_rocksdb_compact_lasted
Last completed compaction duration
ceph_rocksdb_compact_queue_len
  Length of compaction queue
ceph_rocksdb_compact_queue_merge
Mergings of ranges in compaction queue
ceph_rocksdb_compact_running
Running compactions
ceph_rocksdb_get_latency_count
Get latency Count
ceph_rocksdb_get_latency_sum
Get latency Total
ceph_rocksdb_rocksdb_write_delay_time_count
Rocksdb write delay time Count
ceph_rocksdb_rocksdb_write_delay_time_sum
Rocksdb write delay time Total
ceph_rocksdb_rocksdb_write_memtable_time_count
Rocksdb write memtable time Count
ceph_rocksdb_rocksdb_write_memtable_time_sum
Rocksdb write memtable time Total
ceph_rocksdb_rocksdb_write_pre_and_post_time_count
total time spent on writing a record, excluding write process Count
ceph_rocksdb_rocksdb_write_pre_and_post_time_sum
total time spent on writing a record, excluding write process Total
ceph_rocksdb_rocksdb_write_wal_time_count
Rocksdb write wal time Count
ceph_rocksdb_rocksdb_write_wal_time_sum
Rocksdb write wal time Total
ceph_rocksdb_submit_latency_count


63

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Submit Latency Count
ceph_rocksdb_submit_latency_sum
Submit Latency Total
ceph_rocksdb_submit_sync_latency_count
Submit Sync Latency Count
ceph_rocksdb_submit_sync_latency_sum
  Submit Sync Latency Total

6.7. NOOBAA

NooBaa_Endpoint_process_cpu_user_seconds_total
Total user CPU time spent in seconds.
NooBaa_Endpoint_process_cpu_system_seconds_total
Total system CPU time spent in seconds.
NooBaa_Endpoint_process_cpu_seconds_total
Total user and system CPU time spent in seconds.
NooBaa_Endpoint_process_start_time_seconds
Start time of the process since unix epoch in seconds.
NooBaa_Endpoint_process_resident_memory_bytes
Resident memory size in bytes.
NooBaa_Endpoint_process_virtual_memory_bytes
Virtual memory size in bytes.
NooBaa_Endpoint_process_heap_bytes
Process heap size in bytes.
NooBaa_Endpoint_process_open_fds
Number of open file descriptors.
NooBaa_Endpoint_process_max_fds
Maximum number of open file descriptors.
NooBaa_Endpoint_nodejs_eventloop_lag_seconds
Lag of event loop in seconds.
NooBaa_Endpoint_nodejs_eventloop_lag_min_seconds
The minimum recorded event loop delay.
NooBaa_Endpoint_nodejs_eventloop_lag_max_seconds
The maximum recorded event loop delay.
NooBaa_Endpoint_nodejs_eventloop_lag_mean_seconds
The mean of the recorded event loop delays.
NooBaa_Endpoint_nodejs_eventloop_lag_stddev_seconds
The standard deviation of the recorded event loop delays.
NooBaa_Endpoint_nodejs_eventloop_lag_p50_seconds
The 50th percentile of the recorded event loop delays.
NooBaa_Endpoint_nodejs_eventloop_lag_p90_seconds


64

     CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

 The 90th percentile of the recorded event loop delays.
NooBaa_Endpoint_nodejs_eventloop_lag_p99_seconds
 The 99th percentile of the recorded event loop delays.
NooBaa_Endpoint_nodejs_active_resources
 Number of active resources that are currently keeping the event loop alive, grouped by async
 resource type.
NooBaa_Endpoint_nodejs_active_resources_total
 Total number of active resources.
NooBaa_Endpoint_nodejs_active_handles
 Number of active libuv handles grouped by handle type. Every handle type is C++ class name.
NooBaa_Endpoint_nodejs_active_handles_total
 Total number of active handles.
NooBaa_Endpoint_nodejs_active_requests
 Number of active libuv requests grouped by request type. Every request type is C++ class name.
NooBaa_Endpoint_nodejs_active_requests_total
 Total number of active requests.
NooBaa_Endpoint_nodejs_heap_size_total_bytes
 Process heap size from Node.js in bytes.
NooBaa_Endpoint_nodejs_heap_size_used_bytes
 Process heap size used from Node.js in bytes.
NooBaa_Endpoint_nodejs_external_memory_bytes
 Node.js external memory size in bytes.
NooBaa_Endpoint_nodejs_heap_space_size_total_bytes
 Process heap space size total from Node.js in bytes.
NooBaa_Endpoint_nodejs_heap_space_size_used_bytes
 Process heap space size used from Node.js in bytes.
NooBaa_Endpoint_nodejs_heap_space_size_available_bytes
 Process heap space size available from Node.js in bytes.
NooBaa_Endpoint_nodejs_version_info
 Node.js version info.
NooBaa_Endpoint_nodejs_gc_duration_seconds
 Garbage collection duration by kind, one of major, minor, incremental or weakcb.
NooBaa_Endpoint_hub_read_bytes
 hub read bytes in namespace cache bucket
NooBaa_Endpoint_hub_write_bytes
 hub write bytes in namespace cache bucket
NooBaa_Endpoint_cache_read_bytes
 Cache read bytes in namespace cache bucket
NooBaa_Endpoint_cache_write_bytes
 Cache write bytes in namespace cache bucket
NooBaa_Endpoint_cache_object_read_count

     65

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Counter on entire object reads in namespace cache bucket
NooBaa_Endpoint_cache_object_read_miss_count
Counter on entire object read miss in namespace cache bucket
NooBaa_Endpoint_cache_object_read_hit_count
Counter on entire object read hit in namespace cache bucket
NooBaa_Endpoint_cache_range_read_count
Counter on range reads in namespace cache bucket
NooBaa_Endpoint_cache_range_read_miss_count
Counter on range read miss in namespace cache bucket
NooBaa_Endpoint_cache_range_read_hit_count
   Counter on range read hit in namespace cache bucket
NooBaa_Endpoint_hub_read_latency
 hub read latency in namespace cache bucket
NooBaa_Endpoint_hub_write_latency
hub write latency in namespace cache bucket
NooBaa_Endpoint_cache_read_latency
 Cache read latency in namespace cache bucket
NooBaa_Endpoint_cache_write_latency
Cache write latency in namespace cache bucket
NooBaa_Endpoint_semaphore_waiting_value
Namespace semaphore waiting value
NooBaa_Endpoint_semaphore_waiting_time
Namespace semaphore waiting time
NooBaa_Endpoint_semaphore_waiting_queue
Namespace semaphore waiting queue size
NooBaa_Endpoint_semaphore_value
Namespace semaphore value
NooBaa_Endpoint_fork_counter
Counter on number of fork hit
NooBaa_WebServer_process_cpu_user_seconds_total
Total user CPU time spent in seconds.
NooBaa_WebServer_process_cpu_system_seconds_total
Total system CPU time spent in seconds.
NooBaa_WebServer_process_cpu_seconds_total
Total user and system CPU time spent in seconds.
NooBaa_WebServer_process_start_time_seconds
Start time of the process since unix epoch in seconds.
NooBaa_WebServer_process_resident_memory_bytes
Resident memory size in bytes.
NooBaa_WebServer_process_virtual_memory_bytes


66

     CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

 Virtual memory size in bytes.
NooBaa_WebServer_process_heap_bytes
 Process heap size in bytes.
NooBaa_WebServer_process_open_fds
 Number of open file descriptors.
NooBaa_WebServer_process_max_fds
 Maximum number of open file descriptors.
NooBaa_WebServer_nodejs_eventloop_lag_seconds
 Lag of event loop in seconds.
NooBaa_WebServer_nodejs_eventloop_lag_min_seconds
 The minimum recorded event loop delay.
NooBaa_WebServer_nodejs_eventloop_lag_max_seconds
 The maximum recorded event loop delay.
NooBaa_WebServer_nodejs_eventloop_lag_mean_seconds
 The mean of the recorded event loop delays.
NooBaa_WebServer_nodejs_eventloop_lag_stddev_seconds
 The standard deviation of the recorded event loop delays.
NooBaa_WebServer_nodejs_eventloop_lag_p50_seconds
 The 50th percentile of the recorded event loop delays.
NooBaa_WebServer_nodejs_eventloop_lag_p90_seconds
 The 90th percentile of the recorded event loop delays.
NooBaa_WebServer_nodejs_eventloop_lag_p99_seconds
 The 99th percentile of the recorded event loop delays.
NooBaa_WebServer_nodejs_active_resources
 Number of active resources that are currently keeping the event loop alive, grouped by async
 resource type.
NooBaa_WebServer_nodejs_active_resources_total
 Total number of active resources.
NooBaa_WebServer_nodejs_active_handles
 Number of active libuv handles grouped by handle type. Every handle type is C++ class name.
NooBaa_WebServer_nodejs_active_handles_total
 Total number of active handles.
NooBaa_WebServer_nodejs_active_requests
 Number of active libuv requests grouped by request type. Every request type is C++ class name.
NooBaa_WebServer_nodejs_active_requests_total
 Total number of active requests.
NooBaa_WebServer_nodejs_heap_size_total_bytes
 Process heap size from Node.js in bytes.
NooBaa_WebServer_nodejs_heap_size_used_bytes
 Process heap size used from Node.js in bytes.
NooBaa_WebServer_nodejs_external_memory_bytes

     67

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Node.js external memory size in bytes.
NooBaa_WebServer_nodejs_heap_space_size_total_bytes
Process heap space size total from Node.js in bytes.
NooBaa_WebServer_nodejs_heap_space_size_used_bytes
Process heap space size used from Node.js in bytes.
NooBaa_WebServer_nodejs_heap_space_size_available_bytes
Process heap space size available from Node.js in bytes.
NooBaa_WebServer_nodejs_version_info
Node.js version info.
NooBaa_WebServer_nodejs_gc_duration_seconds
Garbage collection duration by kind, one of major, minor, incremental or weakcb.
NooBaa_BGWorkers_process_cpu_user_seconds_total
Total user CPU time spent in seconds.
NooBaa_BGWorkers_process_cpu_system_seconds_total
Total system CPU time spent in seconds.
NooBaa_BGWorkers_process_cpu_seconds_total
Total user and system CPU time spent in seconds.
NooBaa_BGWorkers_process_start_time_seconds
Start time of the process since unix epoch in seconds.
NooBaa_BGWorkers_process_resident_memory_bytes
Resident memory size in bytes.
NooBaa_BGWorkers_process_virtual_memory_bytes
Virtual memory size in bytes.
NooBaa_BGWorkers_process_heap_bytes
Process heap size in bytes.
NooBaa_BGWorkers_process_open_fds
Number of open file descriptors.
NooBaa_BGWorkers_process_max_fds
Maximum number of open file descriptors.
NooBaa_BGWorkers_nodejs_eventloop_lag_seconds
Lag of event loop in seconds.
NooBaa_BGWorkers_nodejs_eventloop_lag_min_seconds
The minimum recorded event loop delay.
NooBaa_BGWorkers_nodejs_eventloop_lag_max_seconds
The maximum recorded event loop delay.
NooBaa_BGWorkers_nodejs_eventloop_lag_mean_seconds
The mean of the recorded event loop delays.
NooBaa_BGWorkers_nodejs_eventloop_lag_stddev_seconds
The standard deviation of the recorded event loop delays.
NooBaa_BGWorkers_nodejs_eventloop_lag_p50_seconds


68

     CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

 The 50th percentile of the recorded event loop delays.
NooBaa_BGWorkers_nodejs_eventloop_lag_p90_seconds
 The 90th percentile of the recorded event loop delays.
NooBaa_BGWorkers_nodejs_eventloop_lag_p99_seconds
 The 99th percentile of the recorded event loop delays.
NooBaa_BGWorkers_nodejs_active_resources
 Number of active resources that are currently keeping the event loop alive, grouped by async
 resource type.
NooBaa_BGWorkers_nodejs_active_resources_total
 Total number of active resources.
NooBaa_BGWorkers_nodejs_active_handles
 Number of active libuv handles grouped by handle type. Every handle type is C++ class name.
NooBaa_BGWorkers_nodejs_active_handles_total
 Total number of active handles.
NooBaa_BGWorkers_nodejs_active_requests
 Number of active libuv requests grouped by request type. Every request type is C++ class name.
NooBaa_BGWorkers_nodejs_active_requests_total
 Total number of active requests.
NooBaa_BGWorkers_nodejs_heap_size_total_bytes
 Process heap size from Node.js in bytes.
NooBaa_BGWorkers_nodejs_heap_size_used_bytes
 Process heap size used from Node.js in bytes.
NooBaa_BGWorkers_nodejs_external_memory_bytes
 Node.js external memory size in bytes.
NooBaa_BGWorkers_nodejs_heap_space_size_total_bytes
 Process heap space size total from Node.js in bytes.
NooBaa_BGWorkers_nodejs_heap_space_size_used_bytes
 Process heap space size used from Node.js in bytes.
NooBaa_BGWorkers_nodejs_heap_space_size_available_bytes
 Process heap space size available from Node.js in bytes.
NooBaa_BGWorkers_nodejs_version_info
 Node.js version info.
NooBaa_BGWorkers_nodejs_gc_duration_seconds
 Garbage collection duration by kind, one of major, minor, incremental or weakcb.
NooBaa_HostedAgents_process_cpu_user_seconds_total
 Total user CPU time spent in seconds.
NooBaa_HostedAgents_process_cpu_system_seconds_total
 Total system CPU time spent in seconds.
NooBaa_HostedAgents_process_cpu_seconds_total
 Total user and system CPU time spent in seconds.
NooBaa_HostedAgents_process_start_time_seconds

     69

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

  Start time of the process since unix epoch in seconds.
NooBaa_HostedAgents_process_resident_memory_bytes
  Resident memory size in bytes.
NooBaa_HostedAgents_process_virtual_memory_bytes
  Virtual memory size in bytes.
NooBaa_HostedAgents_process_heap_bytes
  Process heap size in bytes.
NooBaa_HostedAgents_process_open_fds
  Number of open file descriptors.
NooBaa_HostedAgents_process_max_fds
  Maximum number of open file descriptors.
NooBaa_HostedAgents_nodejs_eventloop_lag_seconds
  Lag of event loop in seconds.
NooBaa_HostedAgents_nodejs_eventloop_lag_min_seconds
  The minimum recorded event loop delay.
NooBaa_HostedAgents_nodejs_eventloop_lag_max_seconds
  The maximum recorded event loop delay.
NooBaa_HostedAgents_nodejs_eventloop_lag_mean_seconds
  The mean of the recorded event loop delays.
NooBaa_HostedAgents_nodejs_eventloop_lag_stddev_seconds
  The standard deviation of the recorded event loop delays.
NooBaa_HostedAgents_nodejs_eventloop_lag_p50_seconds
  The 50th percentile of the recorded event loop delays.
NooBaa_HostedAgents_nodejs_eventloop_lag_p90_seconds
  The 90th percentile of the recorded event loop delays.
NooBaa_HostedAgents_nodejs_eventloop_lag_p99_seconds
  The 99th percentile of the recorded event loop delays.
NooBaa_HostedAgents_nodejs_active_resources
  Number of active resources that are currently keeping the event loop alive, grouped by async
  resource type.
NooBaa_HostedAgents_nodejs_active_resources_total
  Total number of active resources.
NooBaa_HostedAgents_nodejs_active_handles
  Number of active libuv handles grouped by handle type. Every handle type is C++ class name.
NooBaa_HostedAgents_nodejs_active_handles_total
  Total number of active handles.
NooBaa_HostedAgents_nodejs_active_requests
  Number of active libuv requests grouped by request type. Every request type is C++ class name.
NooBaa_HostedAgents_nodejs_active_requests_total
  Total number of active requests.
NooBaa_HostedAgents_nodejs_heap_size_total_bytes

70

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Process heap size from Node.js in bytes.
NooBaa_HostedAgents_nodejs_heap_size_used_bytes
Process heap size used from Node.js in bytes.
NooBaa_HostedAgents_nodejs_external_memory_bytes
Node.js external memory size in bytes.
NooBaa_HostedAgents_nodejs_heap_space_size_total_bytes
Process heap space size total from Node.js in bytes.
NooBaa_HostedAgents_nodejs_heap_space_size_used_bytes
Process heap space size used from Node.js in bytes.
NooBaa_HostedAgents_nodejs_heap_space_size_available_bytes
Process heap space size available from Node.js in bytes.
NooBaa_HostedAgents_nodejs_version_info
Node.js version info.
NooBaa_HostedAgents_nodejs_gc_duration_seconds
Garbage collection duration by kind, one of major, minor, incremental or weakcb.
NooBaa_cloud_types
Cloud Resource Types in the System
NooBaa_projects_capacity_usage
Projects Capacity Usage
NooBaa_accounts_usage_read_count
Accounts Usage Read Count
NooBaa_accounts_usage_write_count
Accounts Usage Write Count
NooBaa_accounts_usage_logical
 Accounts Usage Logical
NooBaa_bucket_class_capacity_usage
Bucket Class Capacity Usage
NooBaa_unhealthy_cloud_types
Unhealthy Cloud Resource Types in the System
NooBaa_object_histo
Object Sizes Histogram Across the System
NooBaa_providers_bandwidth_write_size
Providers bandwidth write size
NooBaa_providers_bandwidth_read_size
 Providers bandwidth read size
NooBaa_providers_ops_read_num
Providers number of read operations
NooBaa_providers_ops_write_num
Providers number of write operations
NooBaa_providers_physical_size


71

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Providers Physical Stats
NooBaa_providers_logical_size
Providers Logical Stats
NooBaa_system_capacity
System capacity
NooBaa_system_info
System info
NooBaa_num_buckets
Object Buckets
NooBaa_num_namespace_buckets
Namespace Buckets
NooBaa_total_usage
Total Usage
NooBaa_accounts_num
Accounts Number
NooBaa_num_objects
Objects
NooBaa_num_unhealthy_buckets
Unhealthy Buckets
NooBaa_num_unhealthy_namespace_buckets
Unhealthy Namespace Buckets
NooBaa_num_unhealthy_pools
Unhealthy Resource Pools
NooBaa_num_unhealthy_namespace_resources
Unhealthy Namespace Resources
NooBaa_num_pools
 Resource Pools
NooBaa_num_namespace_resources
Namespace Resources
NooBaa_num_unhealthy_bucket_claims
Unhealthy Bucket Claims
NooBaa_num_buckets_claims
Object Bucket Claims
NooBaa_num_objects_buckets_claims
Objects On Object Bucket Claims
NooBaa_reduction_ratio
Object Efficiency Ratio
NooBaa_object_savings_logical_size
 Object Savings Logical
NooBaa_object_savings_physical_size


72

    CHAPTER 6. OPENSHIFT DATA FOUNDATION METRICS

Object Savings Physical
NooBaa_rebuild_progress
Rebuild Progress
NooBaa_rebuild_time
Rebuild Time
NooBaa_bucket_status
Bucket Health
NooBaa_namespace_bucket_status
Namespace Bucket Health
NooBaa_bucket_tagging
 Bucket Tagging
NooBaa_namespace_bucket_tagging
Namespace Bucket Tagging
NooBaa_bucket_capacity
Bucket Capacity Precent
NooBaa_bucket_size_quota
Bucket Size Quota Precent
NooBaa_bucket_quantity_quota
Bucket Quantity Quota Precent
NooBaa_bucket_object_count
Number of objects currently stored in the bucket.
NooBaa_bucket_max_objects_quota
Maximum number of objects allowed as bucket quota.
NooBaa_bucket_max_bytes_quota
Maximum storage capacity (in bytes) permitted bucket quota.
NooBaa_resource_status
Resource Health
NooBaa_namespace_resource_status
Namespace Resource Health
NooBaa_system_links
System Links
NooBaa_health_status
Health status
NooBaa_odf_health_status
Health status
NooBaa_replication_status
Replication status
NooBaa_replication_last_cycle_writes_size
Number of bytes replicated by replication_id in last replication cycle
NooBaa_replication_last_cycle_writes_num


73

Red Hat OpenShift Data Foundation 4.22 Monitoring OpenShift Data Foundation

Number of objects replicated by replication_id in last replication cycle
NooBaa_replication_last_cycle_error_writes_size
Number of error bytes replication_id in last replication cycle
NooBaa_replication_last_cycle_error_writes_num
Number of error objects replication_id in last replication cycle
NooBaa_bucket_used_bytes
Object Bucket Used Bytes










74
