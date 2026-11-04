---
title: "deploying openshift data foundation using ibm cloud"
url: "https://docs.redhat.com/en/documentation/red_hat_openshift_data_foundation/4.22/html-single/deploying_openshift_data_foundation_using_ibm_cloud"
source_file: "deploying_openshift_data_foundation_using_ibm_cloud.pdf"
version: "4.22"
chapter: "deploying_openshift_data_foundation_using_ibm_cloud"
---

    Red Hat OpenShift Data Foundation
                   4.22

Deploying OpenShift Data Foundation using
                IBM Cloud

    Instructions on deploying Red Hat OpenShift Data Foundation using IBM Cloud










    Last Updated: 2026-10-22



Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data
Foundation using IBM Cloud

Instructions on deploying Red Hat OpenShift Data Foundation using IBM Cloud

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
Red Hat OpenShift Container Platform on IBM cloud clusters.

        Table of Contents

        Table of Contents

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PREFACE 3

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
PROVIDING FEEDBACK ON RED HAT DOCUMENTATION 4

. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
CHAPTER 1. DEPLOYING OPENSHIFT DATA FOUNDATION USING IBM CLOUD 5










    1

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Cloud










2

    PREFACE

    PREFACE

This document describes how to deploy Red Hat OpenShift Data Foundation using IBM Cloud in
connected environments.










3

Red Hat OpenShift Data Foundation 4.22 Deploying OpenShift Data Foundation using IBM Cloud

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

Red Hat OpenShift Data Foundation 4.12 supports deployment of Red Hat OpenShift on IBM Cloud
clusters in connected environments.










4

        CHAPTER 1. DEPLOYING OPENSHIFT DATA FOUNDATION USING IBM CLOUD

   CHAPTER 1. DEPLOYING OPENSHIFT DATA FOUNDATION
        USING IBM CLOUD

You can use Red Hat OpenShift Data Foundation for your workloads that run in IBM Cloud. These
workloads might run in Red Hat OpenShift on IBM Cloud clusters that are in the public cloud or in your
own IBM Cloud Satellite location.










    5
