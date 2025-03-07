---
title: Deployments in ArangoGraph
menuTitle: Deployments
weight: 20
description: >-
   How to create and manage deployments in ArangoGraph
archetype: chapter
---
Below projects in the ArangoGraph deployment hierarchy are deployments. A deployment
contains an ArangoDB, configured as you choose. You can have any number of
deployments under one project.

**Organizations → Projects → <u>Deployments</u>**

Each deployment can also be backed up manually or automatically by schedules
you can define.

In essence, you can create whatever structure fits you for a given organization,
its projects and deployments.

![ArangoGraph Deployments](../../../images/arangograph-deployments-page.png)

## How to create a new deployment

1. If you do not have a project yet,
   [create a project](../projects.md#how-to-create-a-new-project) first.
2. In the main navigation, in the __Projects__ section, click the project for
   which you want to create a new deployment.
3. In the __Deployments__ page, you will see an empty list or a list with
   your project's existing deployments.
4. Click the __New deployment__ button.
5. Set up your deployment. The configuration options are described below.

![ArangoGraph New Deployment](../../../images/arangograph-new-deployment1.png)

{{< info >}}
Deployments contain exactly **one policy**. Within that policy, you can define
role bindings to regulate access control on a deployment level.
{{< /info >}}

### In the __General__ section

- Enter the name and optionally a short description for the deployment.

### In the __Location__ section

1. Select the __Provider__ and __Region__ of the provider.
   {{< warning >}}
   Once a deployment has been created, it is not possible to change the
   provider and region anymore.
   {{< /warning >}}
2. Select the __DB Version__.
   **Note**: If you don't know which DB version to select, leave the version
   selected by default.
3. In the __CA Certificate__ field
    - The default certificate created for your project will automatically be selected.
    - If you have no default certificate, or want to use a new certificate
      create a new certificate by typing the desired name for it and hitting
      enter or clicking the name when done.
    - Or, if you already have multiple certificates, select the desired one.
4. _Optional but strongly recommended:_ In the __IP allowlist__ field, select the
   desired one in case you want to limit access to your deployment to certain
   IP ranges. To create a allowlist, navigate to your project and select the
   __IP allowlists__ tab.

{{< security >}}
For any kind of production deployment we strongly advise to use an IP allowlist.
{{< /security >}}

### In the __Configuration__ section

Choose between a **OneShard**, **Sharded** or **Single Server** deployment.

- OneShard deployments are suitable when your data set fits in a single node.
  They are ideal for graph use cases.

- Sharded deployments are suitable when your data set is larger than a single
  node. The data will be sharded across multiple nodes.

- Single Server deployments are suitable when you want to try out ArangoDB without
  the need for high availability or scalability. The deployment will contain a
  single server only. Your data will not be replicated and your deployment can
  be restarted at any time.

{{< info >}}
Before you begin configuring your deployment, you first need to select the
provider and region in the Location section.
{{< /info >}}

#### OneShard

- Select the memory size of your node. The disk size is automatically set for you
  based on the selected memory size.

![ArangoGraph Deployment OneShard](../../../images/arangograph-new-deployment-oneshard.png)

#### Sharded

- In addition to the memory size as in the OneShard configuration, select
  the number of nodes for your deployment. The more nodes you have, the higher
  the replication factor can be. Same as in OneShard, the disk size is automatically
  set for you.

![ArangoGraph Deployment Sharded](../../../images/arangograph-new-deployment-sharded.png)

#### Single Server

- Like with OneShard and Sharded deployments, you can choose the memory size.
  Note, however, that the size you choose is for the entire deployment.
  For OneShard and Sharded deployments the chosen sizes are per node.

![ArangoGraph Deployment Single Server](../../../images/arangograph-new-deployment-singleserver.png)

### In the __Summary__ section

1. Review the configuration, and if you're ok with the setup press the
  __Create__ button.
2. You will be taken to the deployment overview page.
   **Note:** Your deployment is at that point being bootstrapped, this process
   will take a few minutes. Once it is ready, you will receive a confirmation email.

## How to access your deployment

1. In the main navigation, click the __Dashboard__ icon and then click __Projects__.
2. In the __Projects__ page, click the project for
   which you created a deployment earlier.
3. Alternatively, you can access your deployment by clicking __Deployments__ in the
   dashboard navigation. This page shows all deployments from all projects.
4. For each deployment in your project, you see the status. While your new
   deployment is being set up, it will display the __bootstrapping__ status.
5. Press the __View__ button to show the deployment page.
6. When a deployment displays a status of __OK__, you can access it.
7. Click the __Open web console__ button or on the endpoint URL property to open
   the dashboard of your new ArangoDB deployment.

At this point your ArangoDB deployment is available for you to use — **Have fun!**

If you have disabled the [auto-login option](#auto-login-to-database-ui) to the
database web interface, you need to follow the additional steps outlined below
to access your deployment:

1. Click the copy icon next to the root password. This will copy the deployment
   root password to your clipboard. You can also click the view icon to unmask
   the root password to see it.

   {{< security >}}
   Do not use the root username/password for everyday operations. It is recommended
   to use them only to create other user accounts with appropriate permissions.
   {{< /security >}}

2. Click the __Open web console__ button or on the endpoint URL property to open
   the dashboard of your new ArangoDB deployment.
3. In the __username__ field type `root`, and in the __password__ field paste the
   password that you copied earlier.
4. Press the __Login__ button.
5. Press the __Select DB: \_system__ button.

{{< info >}}
Each deployment is accessible on two ports:

- Port `8529` is the standard port recommended for use by web-browsers.
- Port `18529` is the alternate port that is recommended for use by automated services.

The difference between these ports is the certificate used. If you enable
__Use well-known certificate__, the certificates used on port `8529` is well-known
and automatically accepted by most web browsers. The certificate used on port
`18529` is a self-signed certificate. For securing automated services, the use of
a self-signed certificate is recommended. Read more on the
[Certificates](../security-and-access-control/x-509-certificates.md) page.
{{< /info >}}

## Password settings

### How to enable the automatic root user password rotation

Password rotation refers to changing passwords regularly - a security best
practice to reduce the vulnerability to password-based attacks and exploits
by limiting for how long passwords are valid. The ArangoGraph Insights Platform
can automatically change the `root` user password of an ArangoDB deployment
periodically to improve security.

1. Navigate to the __Deployment__ for which you want to enable an automatic
   password rotation for the root user.
2. In the __Quick start__ section, click the button with the __gear__ icon next to the
   __ROOT PASSWORD__.
3. In the __Password Settings__ dialog, turn the automatic password rotation on
   and click the __Confirm__ button.

   ![ArangoGraph Deployment Password Rotation](../../../images/arangograph-deployment-password-rotation.png)
4. You can expand the __Root password__ panel to see when the password was
   rotated last. The rotation takes place every three months.

### Auto login to database UI

ArangoGraph provides the ability to automatically login to your database using
your existing ArangoGraph credentials. This not only provides a seamless
experience, preventing you from having to manage multiple sets of credentials
but also improves the overall security of your database. As your credentials
are shared between ArangoGraph and your database, you can benefit from
end-to-end audit traceability for a given user, as well as integration with
ArangoGraph SSO.

You can enable this feature in the **Password Settings** dialog. Please note
that it may take a few minutes to get activated.
Once enabled, you no longer have to fill in the `root` user and password of
your ArangoDB deployment.

This feature can be disabled at any time. You may wish to consider explicitly
disabling this feature in the following situations:
- Your workflow requires you to access the database UI using different accounts
  with differing permission sets, as you cannot switch database users when
  automatic login is enabled.
- You need to give individuals access to a database's UI without giving them
  any access to ArangoGraph. Note, however, that it's possible to only give an
  ArangoGraph user database UI access, without other ArangoGraph permissions.

Before getting started, make sure you are signed into ArangoGraph as a user
with one of the following permissions in your project:
- `data.deployment.full-access`
- `data.deployment.read-only-access`

Organization owners have these permissions enabled by default.
The `deployment-full-access-user` and `deployment-read-only-user` roles which
contain these permissions can also be granted to other members of the
organization. See how to create a
[role binding](../security-and-access-control/_index.md#how-to-view-edit-or-remove-role-bindings-of-a-policy).

{{< warning >}}
This feature is available on `443` port only.
{{< /warning >}}

## How to edit a deployment

You can modify a deployment’s configuration, including the ArangoDB version
that is being used, change the memory size, or even switch from
a OneShard deployment to a Sharded one if your data set no longer fits in a
single node. 

{{< tip >}}
To edit an existing deployment, you must have the necessary set of permissions
attached to your role. Read more about [roles and permissions](../security-and-access-control/_index.md#roles).
{{< /tip >}}

1. Go to the **Projects** section and select an existing deployment from the list. 
2. Open the deployment you want to change. 
3. In the **Quick start** section, click the **Edit** button. 
4. In the **Version and Security** section, you can do the following:
   - Upgrade the ArangoDB version that is currently being used. See also
     [Upgrades and Versioning](upgrades-and-versioning.md)
   - Select a different CA certificate.
   - Add or remove an IP allowlist.
5. In the **Configuration** section, you can do the following:
   - Upgrade the memory size per node. The disk size is automatically set for you.
   - Change **OneShard** deployments into **Sharded** deployments. To do so,
     click **Sharded**. In addition to the other configuration options, you can
     select the number of nodes for your deployment. This can also be modified later on.

   {{< warning >}}
   Notice that you cannot switch from **Sharded** back to **OneShard**.
   {{< /warning >}}

   {{< warning >}}
   When upgrading the memory size in AWS deployments, the value gets locked and
   cannot be changed until the cloud provider rate limit is reset. 
   {{< /warning >}}

6. All changes are reflected in the **Summary** section. Review the new
   configuration and click **Save**. 

## How to connect a driver to your deployment

[ArangoDB drivers](../../develop/drivers/_index.md) allow you to use your ArangoGraph
deployment as a database system for your applications. Drivers act as interfaces
between different programming languages and ArangoDB, which enable you to
connect to and manipulate ArangoDB deployments from within compiled programs
or using scripting languages.

To get started, open a deployment.
In the **Quick start** section, click on the **Connecting drivers** button and
select your programming language. The code snippets provide examples on how to
connect to your instance.

{{< tip >}}
Note that ArangoGraph Insights Platform runs deployments in a cluster
configuration. To achieve the best possible availability, your client
application has to handle connection failures by retrying operations if needed.
{{< /tip >}}

![ArangoGraph Connecting Drivers Example](../../../images/arangograph-connecting-drivers-example.png)

## How to delete a deployment

{{< danger >}}
Deleting a deployment will delete all its data and backups.
This operation is **irreversible**. Please proceed with caution.
{{< /danger >}}

1. In the main navigation, in the __Projects__ section, click the project that
   holds the deployment you wish to delete.
2. In the __Deployments__ page, click the deployment you wish to delete.
3. Click the __Delete/Lock__ entry in the navigation.
4. Click the __Delete deployment__ button.
5. In the modal dialog, confirm the deletion by entering `Delete!` into the
   designated text field.
6. Confirm the deletion by pressing the __Yes__ button.
7. You will be taken back to the deployments page of the project.
   The deployment being deleted will display the __Deleting__ status until it has
   been successfully removed.
