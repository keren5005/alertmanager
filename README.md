# PromLink
**Team members:** Sharona Seleri, Yehontan Ailon, Keren Cohen, Amit Sindani

## Intro

PromLink is a comprehensive alert management solution designed to enhance incident response and monitoring capabilities for modern IT environments. By integrating Prometheus, AlertManager, and various communication applications, PromLink ensures real-time alerting, efficient incident management, and seamless communication within teams.

The project aims to streamline the monitoring process by providing a robust setup for generating, managing, and responding to alerts. Prometheus serves as the monitoring system that collects and stores metrics, while AlertManager processes these alerts to handle deduplication, grouping, and routing to the appropriate receiver integrations. Various communication applications, such as email, PagerDuty, and OpsGenie, are used to deliver alerts and notifications, ensuring that the relevant teams are promptly informed and can take immediate action.

PromLink also includes a detailed setup and configuration guide to help users deploy the system quickly and effectively. The optional integration with Prometheus allows users to leverage advanced monitoring features, further enhancing the capabilities of the alert management system.

## The Project Document
[The Project Document](https://docs.google.com/document/d/1Z8aP73-qfQTJvO1sbzkbGy7vpMwKQOuJ/edit)

## Installation Manuals
- AlertManager
- Prometheus Server Installation on Windows inviorment <br />
  [Prometheus Server Installation and Building in WSL](https://docs.google.com/document/d/1b3tM0dMIS8sU3uE8-a8AZPi6YwI0U_NX/edit)

## Alertmanager

Alertmanager processes alerts from client applications like the Prometheus server. It handles alert deduplication, grouping, and routing to appropriate receiver integrations, such as email, PagerDuty, or OpsGenie. It also supports alert silencing and inhibition. Alertmanager can be installed using precompiled binaries, Docker images, or by compiling the source code manually. Configuration is done via a YAML file, allowing for detailed alert routing, grouping, and inhibition rules.

For high availability, Alertmanager instances can be clustered, ensuring redundancy and reliability. The `amtool` CLI facilitates interaction with the Alertmanager API, allowing for alert management, silencing, and configuration testing. The Alertmanager API, generated via OpenAPI and Go Swagger, provides comprehensive endpoints for alert and configuration management.

## Architecture Diagram

![Alertmanager Architecture](https://raw.githubusercontent.com/keren5005/alertmanager/b0f54f0bc76114b03d81627c714a7d12e4138795/doc/arch.svg)

## Contributed Alertmanagers

### Alertmanager for Metarmost

<div style="background-color:#f0f0f0; padding:10px; border-radius:5px;">
<h4>Introduction</h4>
<p>Here are the Alertmanager setup for Metarmost. Below are the steps to clone the repository, build the project, and run the Alertmanager, as well as instructions for sending alerts using Postman.</p>

<h4>Setup Instructions</h4>
<ol>
<li><strong>Clone the Repository:</strong> First, you need to clone the repository from GitHub. Use the following command:</li>
<pre><code>git clone https://github.com/keren5005/alertmanager.git</code></pre>

<li><strong>Set Environment Variable:</strong> Set the <code>GO15VENDOREXPERIMENT</code> environment variable to 1. This is necessary for building the project.</li>
<pre><code>set GO15VENDOREXPERIMENT=1</code></pre>

<li><strong>Navigate to Alertmanager Directory:</strong> Change your directory to the <code>alertmanager</code> directory.</li>
<pre><code>cd alertmanager</code></pre>

<li><strong>Get Dependencies:</strong> Fetch the required dependencies for the Alertmanager.</li>
<pre><code>go get github.com/prometheus/alertmanager/cmd/...</code></pre>

<li><strong>Build the Alertmanager:</strong> Navigate to the <code>cmd/alertmanager</code> directory and build the project.</li>
<pre><code>cd cmd/alertmanager
go build</code></pre>

<li><strong>Copy Configuration File:</strong> Copy the sample configuration file to use it as your configuration.</li>
<pre><code>cd ../..
copy cmd/alertmanager/alertmanager.yml.sample cmd/alertmanager/alertmanager.yml</code></pre>

<li><strong>Run Alertmanager:</strong> Run the Alertmanager executable.</li>
<pre><code>cmd/alertmanager/alertmanager.exe</code></pre>

<li><strong>Access Alertmanager Web Interface:</strong> Open your web browser and go to: <a href="http://localhost:9093">http://localhost:9093</a></li>
</ol>

<h4>Development Instructions</h4>
<ol>
<li><strong>Open a New Terminal and Navigate to Alertmanager Directory:</strong> Open a new terminal window and navigate to the <code>alertmanager/cmd/alertmanager</code> directory.</li>
<pre><code>cd alertmanager/cmd/alertmanager</code></pre>

<li><strong>Build the Project:</strong> Build the Alertmanager project.</li>
<pre><code>go build</code></pre>
</ol>

<h4>Sending Alerts with Postman</h4>
<ol>
<li><strong>POST Request to Trigger an Alert:</strong> Send a POST request to <a href="[http://localhost:9093/api/v2/alerts](http://localhost:9093/api/v2/alerts)">http://localhost:9093/api/v2/alerts</a> with the following JSON body:</li>
<pre><code>[
{
    "status": "firing",
	"labels": {
		"alertname": "$name4",
		"service": "my-service4",
		"severity": "warning4",
		"instance": "$name.example.net4"
	},
	"annotations": {
		"summary": "High latency is high!"
	}
}
]
</code></pre>

<li><strong>Closing the Alert:</strong> Run the following command to stop the Alertmanager:</li>
<pre><code>cmd/alertmanager/alertmanager.exe</code></pre>
</ol>
</div>
