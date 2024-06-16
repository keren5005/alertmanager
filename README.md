<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>PromLink README</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 20px;
            background-color: #f4f4f4;
        }
        .container {
            width: 80%;
            margin: auto;
            overflow: hidden;
        }
        header {
            background: #333;
            color: #fff;
            padding-top: 30px;
            min-height: 70px;
            border-bottom: #0779e4 3px solid;
        }
        header a, header img {
            color: #fff;
            text-decoration: none;
            text-transform: uppercase;
            font-size: 16px;
        }
        header img {
            display: inline-block;
            vertical-align: middle;
        }
        header h1 {
            display: inline-block;
            vertical-align: middle;
            margin: 0 0 0 10px;
        }
        .badges img {
            margin-right: 5px;
        }
        .content {
            background: #fff;
            padding: 20px;
            margin-top: 20px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
        }
        .content h2 {
            color: #333;
        }
        .content p, .content li {
            color: #555;
        }
        .content pre {
            background: #eee;
            padding: 10px;
            border: 1px solid #ccc;
            overflow: auto;
        }
        .content code {
            color: #d14;
            background: #f9f2f4;
            padding: 2px 4px;
            border-radius: 4px;
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>PromLink</h1>
            <div class="badges">
                <img src="https://img.shields.io/badge/license-GPL--3.0-blue" alt="License">
                <img src="https://img.shields.io/badge/commits-36%2Fmonth-brightgreen" alt="Commits">
                <img src="https://img.shields.io/badge/issues-104%20open-yellow" alt="Issues">
                <img src="https://img.shields.io/badge/build-passing-brightgreen" alt="Build Status">
            </div>
        </header>

        <div class="content">
            <h2>Intro</h2>
            <p>PromLink is a comprehensive alert management solution designed to enhance incident response and monitoring capabilities for modern IT environments. By integrating Prometheus, AlertManager, and various communication applications, PromLink ensures real-time alerting, efficient incident management, and seamless communication within teams.</p>

            <p>The project aims to streamline the monitoring process by providing a robust setup for generating, managing, and responding to alerts. Prometheus serves as the monitoring system that collects and stores metrics, while AlertManager processes these alerts to handle deduplication, grouping, and routing to the appropriate receiver integrations. Various communication applications, such as email, PagerDuty, and OpsGenie, are used to deliver alerts and notifications, ensuring that the relevant teams are promptly informed and can take immediate action.</p>

            <p>PromLink also includes a detailed setup and configuration guide to help users deploy the system quickly and effectively. The optional integration with Prometheus allows users to leverage advanced monitoring features, further enhancing the capabilities of the alert management system.</p>

            <h2>The Project Document</h2>
            <p><a href="https://docs.google.com/document/d/1Z8aP73-qfQTJvO1sbzkbGy7vpMwKQOuJ/edit">The Project Document</a></p>

            <h2>Installations</h2>
            <ul>
                <li>AlertManager</li>
                <li>Optional: Prometheus <br>
                    <a href="https://docs.google.com/document/d/1b3tM0dMIS8sU3uE8-a8AZPi6YwI0U_NX/edit">Prometheus Server Installation and Building in WSL</a>
                </li>
            </ul>

            <h2>Alertmanager</h2>
            <p>Alertmanager processes alerts from client applications like the Prometheus server. It handles alert deduplication, grouping, and routing to appropriate receiver integrations, such as email, PagerDuty, or OpsGenie. It also supports alert silencing and inhibition. Alertmanager can be installed using precompiled binaries, Docker images, or by compiling the source code manually. Configuration is done via a YAML file, allowing for detailed alert routing, grouping, and inhibition rules.</p>

            <p>For high availability, Alertmanager instances can be clustered, ensuring redundancy and reliability. The <code>amtool</code> CLI facilitates interaction with the Alertmanager API, allowing for alert management, silencing, and configuration testing. The Alertmanager API, generated via OpenAPI and Go Swagger, provides comprehensive endpoints for alert and configuration management.</p>

            <h2>Architecture Diagram</h2>
            <p><img src="https://raw.githubusercontent.com/keren5005/alertmanager/b0f54f0bc76114b03d81627c714a7d12e4138795/doc/arch.svg" alt="Alertmanager Architecture"></p>

            <h2>Contributed Alertmanagers</h2>
            <h3>Alertmanager for Metarmost</h3>
            <h4>Introduction</h4>
            <p>This repository contains the Alertmanager setup for Metarmost. Below are the steps to clone the repository, build the project, and run the Alertmanager, as well as instructions for sending alerts using Postman.</p>

            <h4>Setup Instructions</h4>
            <ol>
                <li><strong>Clone the Repository:</strong>
                    <pre><code>git clone https://github.com/keren5005/alertmanager.git</code></pre>
                </li>
                <li><strong>Set Environment Variable:</strong>
                    <pre><code>set GO15VENDOREXPERIMENT=1</code></pre>
                </li>
                <li><strong>Navigate to Alertmanager Directory:</strong>
                    <pre><code>cd alertmanager</code></pre>
                </li>
                <li><strong>Get Dependencies:</strong>
                    <pre><code>go get github.com/prometheus/alertmanager/cmd/...</code></pre>
                </li>
                <li><strong>Build the Alertmanager:</strong>
                    <pre><code>cd cmd/alertmanager
go build</code></pre>
                </li>
                <li><strong>Copy Configuration File:</strong>
                    <pre><code>cd ../..
copy cmd/alertmanager/alertmanager.yml.sample cmd/alertmanager/alertmanager.yml</code></pre>
                </li>
                <li><strong>Run Alertmanager:</strong>
                    <pre><code>cmd/alertmanager/alertmanager.exe</code></pre>
                </li>
                <li><strong>Access Alertmanager Web Interface:</strong>
                    <p>Open your web browser and go to: <a href="http://localhost:9093">http://localhost:9093</a></p>
                </li>
            </ol>

            <h4>Development Instructions</h4>
            <ol>
                <li><strong>Open a New Terminal and Navigate to Alertmanager Directory:</strong>
                    <pre><code>cd alertmanager/cmd/alertmanager</code></pre>
                </li>
                <li><strong>Build the Project:</strong>
                    <pre><code>go build</code></pre>
                </li>
            </ol>

            <h4>Sending Alerts with Postman</h4>
            <ol>
                <li><strong>POST Request to Trigger an Alert:</strong>
                    <pre><code>[
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
                    </code></pre>
                </li>
                <li><strong>Closing the Alert:</strong>
                    <pre><code>cmd/alertmanager/alertmanager.exe</code></pre>
                </li>
            </ol>
        </div>
    </div>
</body>
</html>
