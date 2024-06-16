PromLink
Team members: Sharona Seleri, Yehontan Ailon, Keren Cohen, Amit Sindani

Intro
PromLink is a comprehensive alert management solution designed to enhance incident response and monitoring capabilities for modern IT environments. By integrating Prometheus, AlertManager, and various communication applications, PromLink ensures real-time alerting, efficient incident management, and seamless communication within teams.

The project aims to streamline the monitoring process by providing a robust setup for generating, managing, and responding to alerts. Prometheus serves as the monitoring system that collects and stores metrics, while AlertManager processes these alerts to handle deduplication, grouping, and routing to the appropriate receiver integrations. Various communication applications, such as email, PagerDuty, and OpsGenie, are used to deliver alerts and notifications, ensuring that the relevant teams are promptly informed and can take immediate action.

PromLink also includes a detailed setup and configuration guide to help users deploy the system quickly and effectively. The optional integration with Prometheus allows users to leverage advanced monitoring features, further enhancing the capabilities of the alert management system.

The Project Document
The Project Document

Installations
AlertManager
Optional: Prometheus <br />
Prometheus Server Installation and Building in WSL
Alertmanager
Alertmanager processes alerts from client applications like the Prometheus server. It handles alert deduplication, grouping, and routing to appropriate receiver integrations, such as email, PagerDuty, or OpsGenie. It also supports alert silencing and inhibition. Alertmanager can be installed using precompiled binaries, Docker images, or by compiling the source code manually. Configuration is done via a YAML file, allowing for detailed alert routing, grouping, and inhibition rules.

For high availability, Alertmanager instances can be clustered, ensuring redundancy and reliability. The amtool CLI facilitates interaction with the Alertmanager API, allowing for alert management, silencing, and configuration testing. The Alertmanager API, generated via OpenAPI and Go Swagger, provides comprehensive endpoints for alert and configuration management.

Architecture Diagram
Alertmanager for Metarmost
Introduction
This repository contains the Alertmanager setup for Metarmost. Below are the steps to clone the repository, build the project, and run the Alertmanager, as well as instructions for sending alerts using Postman.

Setup Instructions
Clone the Repository

sh
Copy code
git clone https://github.com/keren5005/alertmanager.git
Set Environment Variable

sh
Copy code
set GO15VENDOREXPERIMENT=1
Navigate to Alertmanager Directory

sh
Copy code
cd alertmanager
Get Dependencies

sh
Copy code
go get github.com/prometheus/alertmanager/cmd/...
Build the Alertmanager

sh
Copy code
cd .\cmd\alertmanager
go build
Copy Configuration File

sh
Copy code
cd .... 
copy .\alertmanager.yml.sample .\alertmanager.yml
Run Alertmanager

sh
Copy code
.\cmd\alertmanager\alertmanager.exe
Access Alertmanager Web Interface
Open your web browser and go to: http://localhost:9093

Development Instructions
Open a New Terminal and Navigate to Alertmanager Directory

sh
Copy code
cd alertmanager\cmd\alertmanager
Build the Project

sh
Copy code
go build
Sending Alerts with Postman
POST Request to Trigger an Alert
Send a POST request to http://localhost:9093/api/v2/alerts with the following JSON body:

json
Copy code
[
  {
    "labels": {
      "alertname": "TestAlert",
      "severity": "critical"
    },
    "annotations": {
      "summary": "This is a test alert"
    },
    "startsAt": "2024-06-16T00:00:00Z",
    "endsAt": "2024-06-16T01:00:00Z",
    "generatorURL": "http://prometheus.io"
  }
]
Closing the Alert
Run the following command to stop the Alertmanager:

sh
Copy code
cmd\alertmanager\alertmanager.exe
