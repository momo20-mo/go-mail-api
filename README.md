# Go Mail API

[![Build Status][badge_build_status]][link_build_status]
[![Release Status][badge_release_status]][link_build_status]
[![Repo size][badge_repo_size]][link_repo]
[![Image size][badge_size_latest]][link_docker_hub]
[![Docker Pulls][badge_docker_pulls]][link_docker_hub]

## Overview
Go Mail API is a simple REST API for sending email notifications via SMTP.

## Features
- Send emails using SMTP.
- Configurable SMTP settings.
- Support for environment variables and command-line flags for configuration.

## Installation
Download the appropriate binary for your operating system from the [GitHub Releases page](https://github.com/braveokafor/go-mail-api/releases).

## Configuration

The Go Mail API can be configured using environment variables or command-line flags. The following options are available:

| Environment Variable | Command-Line Flag | Description                    | Default Value |
|----------------------|-------------------|--------------------------------|---------------|
| `MAIL_API_PORT`      | `--api-port`      | Port for the API server        | `8080`        |
| `MAIL_SMTP_PORT`     | `--smtp-port`     | Port for the SMTP server       | `25`          |
| `MAIL_SMTP_HOST`     | `--smtp-host`     | Host for the SMTP server       | (required)    |
| `MAIL_SMTP_USER`     | `--smtp-user`     | User for the SMTP server       | (required)    |
| `MAIL_SMTP_PASS`     | `--smtp-pass`     | Password for the SMTP server   | (required)    |
| `MAIL_SMTP_USE_TLS`  | `--smtp-use-tls`  | Use TLS for the SMTP server    | `false`       |

Make sure to set the required configuration values (`MAIL_SMTP_HOST`, `MAIL_SMTP_USER`, `MAIL_SMTP_PASS`) either through environment variables or command-line flags.

## API Documentation

### Send Email

Send an email using the Go Mail API.

- **URL:** `/send`
- **Method:** `POST`

#### Request Body

| Field         | Type                | Required | Description                                                    |
|---------------|---------------------|----------|----------------------------------------------------------------|
| `from`        | string              | Yes      | The email address of the sender.                               |
| `to`          | array of strings    | Yes      | The email addresses of the recipients.                         |
| `cc`          | array of strings    | No       | The email addresses of the CC recipients.                      |
| `bcc`         | array of strings    | No       | The email addresses of the BCC recipients.                     |
| `subject`     | string              | Yes      | The subject of the email.                                      |
| `priority`    | string              | No       | The priority of the email (e.g., "high", "normal", "low").     |
| `headers`     | object              | No       | Additional headers to include in the email.                    |
| `text`        | string              | No       | The plain text version of the email body.                      |
| `html`        | string              | No       | The HTML version of the email body.                            |
| `attachments` | array of attachment objects | No       | Attachments to include in the email.                     |

##### Attachment Object

| Field         | Type    | Required | Description                                                         |
|---------------|---------|----------|---------------------------------------------------------------------|
| `filename`    | string  | Yes      | The name of the attachment file.                                    |
| `content_type`| string  | Yes      | The content type of the attachment.                                 |
| `content`     | string  | Yes      | The content of the attachment, either base64-encoded or plain text. |
| `encoded`     | boolean | No       | Indicates whether the content is base64-encoded. Defaults to `false`.|

#### Response

| Status Code | Body                                      |
|-------------|-------------------------------------------|
| `200 OK`    | `{"message": "Email sent successfully"}` |

### Example

1. Set the required configuration values:
   ```sh
   export MAIL_SMTP_HOST="smtp.gmail.com"
   export MAIL_SMTP_USER="youremail@gmail.com"
   export MAIL_SMTP_PASS="your app password"
   ```

2. Start the Go Mail API server:
   ```sh
   ./go-mail-api --api-port=8080 --smtp-port=587 --smtp-use-tls=true
   ```

3. Send a POST request to the `/send` endpoint:
   ```sh
   curl -X POST -H "Content-Type: application/json" -d '{
     "from": "youremail@gmail.com",
     "to": ["recipient1@example.com", "recipient2@example.com"],
     "subject": "Test Email",
     "priority": "high",
     "html": "<p>This is a <strong>test</strong> email.</p>",
     "attachments": [
       {
         "filename": "example.txt",
         "content_type": "text/plain",
         "content": "VGhpcyBpcyBhbiBleGFtcGxlIGF0dGFjaG1lbnQu",
         "encoded": true
       }
     ]
   }' http://localhost:8080/send
   ```

## Contributing
Contributions and suggestions are welcome.  
Please open an issue for discussion or propose improvements directly through a pull request.

## Support & Issues

[![Issues][badge_issues]][link_issues]
[![Issues][badge_pulls]][link_pulls]

For support or reporting issues, please open an issue in the GitHub repository or reach out on [LinkedIn](https://www.linkedin.com/in/braveokafor/).


[link_issues]:https://github.com/braveokafor/go-mail-api/issues
[link_pulls]:https://github.com/braveokafor/go-mail-api/pulls
[link_build_status]:https://github.com/braveokafor/go-mail-api/actions/workflows/ci.yaml
[link_build_status]:https://github.com/braveokafor/go-mail-api/actions/workflows/release.yaml
[link_docker_hub]:https://hub.docker.com/r/braveokafor/go-mail-api
[link_repo]:https://github.com/braveokafor/go-mail-api

[badge_issues]:https://img.shields.io/github/issues-raw/braveokafor/go-mail-api?style=flat-square&logo=GitHub
[badge_pulls]:https://img.shields.io/github/issues-pr/braveokafor/go-mail-api?style=flat-square&logo=GitHub
[badge_build_status]:https://img.shields.io/github/actions/workflow/status/braveokafor/go-mail-api/ci.yaml?style=flat-square&logo=GitHub&label=build
[badge_release_status]:https://img.shields.io/github/actions/workflow/status/braveokafor/go-mail-api/release.yaml?style=flat-square&logo=GitHub&label=release
[badge_size_latest]:https://img.shields.io/docker/image-size/braveokafor/go-mail-api/latest?style=flat-square&logo=Docker
[badge_docker_pulls]:https://img.shields.io/docker/pulls/braveokafor/go-mail-api?style=flat-square&logo=Docker
[badge_repo_size]:https://img.shields.io/github/repo-size/braveokafor/go-mail-api?style=flat-square&logo=GitHub
