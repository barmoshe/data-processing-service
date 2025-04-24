# Data Processing Service
[![Temporal Code Exchange Featured](https://img.shields.io/badge/Temporal-Code_Exchange_Featured-blue?style=flat-square&logo=temporal&labelColor=141414&color=444CE7)](https://temporal.io/code-exchange/cross-language-data-processing-service-with-temporal)

A practical demonstration of building cross-language microservices with [Temporal](https://temporal.io/). This project orchestrates activities written in **Go**, **Python**, and **TypeScript**—showcasing how Temporal can seamlessly coordinate different languages to build a robust, modular data processing workflow.

> **Related Article:**  
> [Building a Cross-Language Data Processing Service with Temporal: A Practical Guide](https://medium.com/@barmoshe/building-a-cross-language-data-processing-service-with-temporal-a-practical-guide-bf0fb1155d46)

---

## Table of Contents

- [Data Processing Service](#data-processing-service)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Project Structure](#project-structure)
  - [Prerequisites](#prerequisites)
  - [Installation \& Setup](#installation--setup)
  - [Usage](#usage)
    - [Starting the Temporal Server](#starting-the-temporal-server)
    - [Running Workers](#running-workers)
      - [Go Worker](#go-worker)
      - [Python Worker](#python-worker)
      - [TypeScript Worker](#typescript-worker)
    - [Running the Client](#running-the-client)
  - [How It Works](#how-it-works)
  - [Summary](#summary)
  - [License](#license)

---

## Overview

This repository is based on the article [Building a Cross-Language Data Processing Service with Temporal: A Practical Guide](https://medium.com/@barmoshe/building-a-cross-language-data-processing-service-with-temporal-a-practical-guide-bf0fb1155d46). It demonstrates:

- **Temporal Server Setup:** Running a local Temporal server with a built-in UI and a local database file.
- **Multi-Language Workflow:** A workflow that invokes activities across Go, Python, and TypeScript.
- **Task Queue Coordination:** How each worker listens on a dedicated task queue.
- **Real vs. Demo Activities:** Starting with demo activities and progressively replacing them with real implementations.

---

## Project Structure

```
.
├── activities
│   ├── activities.go              # Go activity: AddSuffixActivity.
│   ├── activities_python.py       # Python activity: PythonAddRandomPrefixActivity.
│   └── activities_ts.ts           # TypeScript activity: toUpperCaseActivity.
├── client
│   └── main.go                    # Go client to submit workflows.
├── readme.md                      # This file.
├── workflows
│   └── workflow.go                # Go workflow definition invoking activities.
├── workers
│   ├── go
│   │   └── main.go                # Go worker that executes the workflow and Go activities.
│   ├── python
│   │   └── python_worker.py       # Python worker for prefixing.
│   └── ts
│       └── ts_worker.ts           # TypeScript worker for uppercasing.
└── your_temporal.db               # Local database file created by the Temporal server.
```

---

## Prerequisites

- **Temporal Server:** Follow [Temporal's Quick Start Guide](https://docs.temporal.io/docs/quick-start/) for installation.
- **Go:** Version 1.16 or higher.
- **Python:** Version 3.7 or above.
- **Node.js & npm:** Required for running the TypeScript worker.
- **Additional Tools:** `tsx` (or `ts-node`) for executing TypeScript files.

---

## Installation & Setup

1. **Install Temporal CLI:**

   - For macOS, run:
     ```bash
     brew install temporal
     ```
   - Or follow the [Temporal CLI installation guide](https://temporal.io/setup/install-temporal-cli) for other platforms.

2. **Clone the Repository & Navigate to the Project Folder:**

   ```bash
   git clone https://github.com/your-username/data-processing-service.git
   cd data-processing-service
   ```

3. **Initialize the Go Module & Install Dependencies:**

   ```bash
   go mod init data-processing-service
   go get go.temporal.io/sdk
   go mod tidy
   ```

4. **Install Node Dependencies for the TypeScript Worker:**

   ```bash
   cd workers/ts
   npm install
   cd ../..
   ```

5. **Start the Temporal Server:**

   Run the following command to start a local Temporal server. This command creates a local database file (`your_temporal.db`) and opens the Temporal UI on port `8080`:
   
   ```bash
   temporal server start-dev --db-filename your_temporal.db --ui-port 8080
   ```
   
   You should see output similar to:
   
   ```plaintext
   CLI 1.2.0 (Server 1.26.2, UI 2.34.0)
   Server:  localhost:7233
   UI:      http://localhost:8080
   Metrics: http://localhost:62564/metrics
   ```

6. **(Optional) Open the Project in VS Code:**

   ```bash
   code .
   ```

---

## Usage

### Starting the Temporal Server

Ensure the Temporal server is running (see the Installation & Setup section). The server uses the local database file (`your_temporal.db`) and provides a UI at [http://localhost:8080](http://localhost:8080).

---

### Running Workers

Each worker processes specific activities and listens on its designated task queue.

#### Go Worker

- **Location:** `workers/go/main.go`
- **Command:**

  ```bash
  go run workers/go/main.go
  ```

- **Notes:**  
  This worker registers the workflow (from the `workflows` folder) and the Go activity (`AddSuffixActivity` from the `activities` folder). It listens on the task queue `data-processing-task-queue`.

#### Python Worker

- **Location:** `workers/python/python_worker.py`
- **Command:**

  ```bash
  python3 -m workers.python.python_worker
  ```

- **Notes:**  
  This worker handles the Python activity (`PythonAddRandomPrefixActivity` from the `activities` folder) and polls the `python-task-queue`.

#### TypeScript Worker

- **Location:** `workers/ts/ts_worker.ts`
- **Command:**

  ```bash
  npx tsx workers/ts/ts_worker.ts
  ```

- **Notes:**  
  This worker processes the TypeScript activity (`TypeScriptToUppercaseActivity` from the `activities` folder) and polls the `typescript-task-queue`.

---

### Running the Client

The client submits a workflow to the Temporal server. To run the client:

```bash
cd client
go run main.go "sample-data"
```

Upon execution, the client will:
- Generate a unique workflow ID.
- Submit the workflow to the Temporal server.
- Display the processed result (for example, `LAMBDA-SAMPLE-DATA-SIX`).

Check the Temporal UI ([http://localhost:8080](http://localhost:8080)) to see the workflow's progress and details.

---

## How It Works

1. **Workflow Orchestration:**  
   The workflow (defined in `workflows/workflow.go`) sequentially calls activities across different languages:
   - **Python Activity:** Adds a random prefix to the input data (implemented in `activities/activities_python.py`).
   - **Go Activity:** Appends a suffix to the data (implemented in `activities/activities.go`).
   - **TypeScript Activity:** Converts the modified data to uppercase (implemented in `activities/activities_ts.ts`).
   
2. **Task Queue Management:**  
   Each activity is executed by a worker that listens on a specific task queue:
   - **`python-task-queue`** for the Python activity.
   - **`data-processing-task-queue`** for the Go activity and workflow.
   - **`typescript-task-queue`** for the TypeScript activity.
   
3. **Temporal's Role:**  
   Temporal ensures seamless communication between these activities, handling retries and state management, so you can focus on your business logic.

---

## Summary

This project demonstrates how to build a cross-language data processing service using Temporal.  
By splitting the implementation into language-specific workers and leveraging Temporal's workflow orchestration, you can create modular, scalable microservices that interact seamlessly—no matter the language.

---
## License
This project is licensed under the [MIT License](LICENSE).
