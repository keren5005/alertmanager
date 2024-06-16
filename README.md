# PromLink
**Team members:** Sharona Seleri, Yehontan Ailon, Keren Cohen, Amit Sindani

## Intro

PromLink is a comprehensive alert management solution designed to enhance incident response and monitoring capabilities for modern IT environments. By integrating Prometheus, AlertManager, and various communication applications, PromLink ensures real-time alerting, efficient incident management, and seamless communication within teams.

The project aims to streamline the monitoring process by providing a robust setup for generating, managing, and responding to alerts. Prometheus serves as the monitoring system that collects and stores metrics, while AlertManager processes these alerts to handle deduplication, grouping, and routing to the appropriate receiver integrations. Various communication applications, such as email, PagerDuty, and OpsGenie, are used to deliver alerts and notifications, ensuring that the relevant teams are promptly informed and can take immediate action.

PromLink also includes a detailed setup and configuration guide to help users deploy the system quickly and effectively. The optional integration with Prometheus allows users to leverage advanced monitoring features, further enhancing the capabilities of the alert management system.

## The Project Document
[The Project Document](https://docs.google.com/document/d/1Z8aP73-qfQTJvO1sbzkbGy7vpMwKQOuJ/edit)

## Installations
- AlertManager
- Optional: Prometheus <br />
  [Prometheus Server Installation and Building in WSL](https://docs.google.com/document/d/1b3tM0dMIS8sU3uE8-a8AZPi6YwI0U_NX/edit)

## Alertmanager

Alertmanager processes alerts from client applications like the Prometheus server. It handles alert deduplication, grouping, and routing to appropriate receiver integrations, such as email, PagerDuty, or OpsGenie. It also supports alert silencing and inhibition. Alertmanager can be installed using precompiled binaries, Docker images, or by compiling the source code manually. Configuration is done via a YAML file, allowing for detailed alert routing, grouping, and inhibition rules.

For high availability, Alertmanager instances can be clustered, ensuring redundancy and reliability. The `amtool` CLI facilitates interaction with the Alertmanager API, allowing for alert management, silencing, and configuration testing. The Alertmanager API, generated via OpenAPI and Go Swagger, provides comprehensive endpoints for alert and configuration management.

## Architecture Diagram

![Alertmanager Architecture](https://raw.githubusercontent.com/keren5005/alertmanager/b0f54f0bc76114b03d81627c714a7d12e4138795/doc/arch.svg)

## Alertmanager for Metarmost

### Introduction

This repository contains the Alertmanager setup for Metarmost. Below are the steps to clone the repository, build the project, and run the Alertmanager, as well as instructions for sending alerts using Postman.

### Setup Instructions

1. **Clone the Repository:**
   First, you need to clone the repository from GitHub. Use the following command:
   ```sh
   git clone https://github.com/keren5005/alertmanager.git
