---
title: Getting Started
menuTitle: Overview
weight: 30
description: >-
  Learn about ArangoDB's core concepts such as the high-level data
  organization and its data model
archetype: default
---
{{< description >}}

This chapter introduces ArangoDB's core concepts and covers the following:

- The high-level data organization
- Its data model (or rather multiple data models)
- Important terminology used throughout the database system and in this
  documentation

You can also find examples on how to interact with the database system
using ArangoDB's interactive command-line interface called [arangosh](../components/tools/arangodb-shell/_index.md).

See [How to Interact With ArangoDB](how-to-interact-with-arangodb.md) to learn
more about:
- Ways you can use and communicate with ArangoDB servers
- How to set up and deploy an ArangoDB instance
- Get data in and out of ArangoDB
- Back up and restore data in ArangoDB

Learn more about ArangoDB's [data structure](../concepts/data-structure/_index.md) and
[data models](../concepts/data-models.md).

## Modeling Data for ArangoDB

Plan for your data needs and map your conceptual model to the right features,
making the most of ArangoDB.

All of the following topics fall under the broader term of data modeling:

- Analyzing your project goals and existing data
- Designing logical models for your data and mapping them to ArangoDB, often
  by striking a balance between natural data structures and great performance
- Preparing and loading data into ArangoDB
- Transforming data once it is in ArangoDB

## Coming from SQL

If you worked with a database management system (RDBMS) such as MySQL,
MariaDB or PostgreSQL, you should be familiar with the SQL query language.

ArangoDB's query language is called AQL. There are some similarities between both
languages despite the different data models of the database systems. The most
notable difference is probably the concept of loops in AQL, which makes it feel
more like a programming language. It suits the schema-less model more natural
and makes the query language very powerful while remaining easy to read and write.

To get started with AQL, sign up for [ArangoDB University](https://university.arangodb.com/)
and get access to interactive courses powered by ArangoGraph. 

## Managed Cloud Service or On-premises

When you are ready for your own ArangoDB server, you can sign up for ArangoDB's
cloud service called ArangoGraph, which takes care of the setup and maintenance, so
that you can focus on building amazing things on top of ArangoDB. See
[Use ArangoDB in the Cloud](set-up-a-cloud-instance.md) to get started.

You can also install ArangoDB locally or on your own server hardware.
See [Install ArangoDB on-premises](on-premises-installation.md) for more details.

If you want to migrate from an on-premises system to the fully managed ArangoGraph
cloud platform, check out the [Cloud Migration Tool](../arangograph/migrate-to-the-cloud.md).

## Interactive Tutorials

To get started with ArangoDB and try out some of its features, you can use the
interactive tutorials. They provide information on various topics and allow you
to set up and access a cloud instance of ArangoDB - it’s free, easy to use, and
no installation is required.

**AQL Tutorials**

- [Game of Thrones AQL Tutorial](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/ArangoDB_GOT_Tutorial.ipynb)
- [AQL CRUD Part 1](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/AqlCrudTutorial.ipynb)
- [AQL CRUD Part 2](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/AqlPart2Tutorial.ipynb)
- [AQL Joins](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/AqlJoinTutorial.ipynb)
- [AQL Graph Traversal](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/AqlTraversalTutorial.ipynb)
- [AQL Geospatial](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/AqlGeospatialTutorial.ipynb)

**Graph Tutorials**

- [Property Graph Queries](https://colab.research.google.com/github/joerg84/Graph_Powered_ML_Workshop/blob/master/Graphs_Queries.ipynb)

**ArangoSearch Tutorials**

- [ArangoSearch](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/ArangoSearch.ipynb)
- [Fuzzy Search](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/FuzzySearch.ipynb)

**GraphML and Analytics Tutorials**

- [ArangoDB and NetworkX](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/ArangoDB_NetworkX_Interface_Introduction.ipynb)
- [Graph Analytics: Collaborative Filtering](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/Collaborative_Filtering.ipynb)
- [Graph Analytics: Fraud Detection](https://colab.research.google.com/github/joerg84/Graph_Powered_ML_Workshop/blob/master/Fraud_Detection.ipynb)
- [Graph Neural Networks with PyTorch](https://colab.research.google.com/github/arangodb/interactive_tutorials/blob/master/notebooks/arangoflix/predict_Movie_Rating_GNN.ipynb)

You can find all interactive tutorials on GitHub:

<https://github.com/arangodb/interactive_tutorials/>
