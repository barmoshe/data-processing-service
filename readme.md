# Data Processing Service

A practical demonstration of building cross-language microservices with [Temporal](https://temporal.io/). This project orchestrates activities written in **Go**, **Python**, and **TypeScript**—showcasing how Temporal can seamlessly coordinate different languages to build a robust, modular data processing workflow.

> **Related Article:**  
> [Building a Cross-Language Data Processing Service with Temporal: A Practical Guide](https://medium.com/@barmoshe/building-a-cross-language-data-processing-service-with-temporal-a-practical-guide-bf0fb1155d46)

---

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [Usage](#usage)
  - [Starting the Temporal Server](#starting-the-temporal-server)
  - [Running Workers](#running-workers)
  - [Running the Client](#running-the-client)
- [How It Works](#how-it-works)
- [Contributing](#contributing)

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
├── go.mod                         # Go module file.
├── go.sum                         # Go module checksums.
├── readme.md                      # This file.
├── workflow
│   └── workflow.go                # Go workflow definition invoking activities.
├── workers
│   ├── main.go                    # Go worker that executes the workflow and Go activities.
│   ├── python
│   │   └── python_worker.py       # Python worker for prefixing.
│   └── ts
│       ├── package-lock.json
│       ├── package.json
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

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/data-processing-service.git
   cd data-processing-service
   ```

2. **Install Node Dependencies:**

   Navigate to the TypeScript worker directory and install dependencies:

   ```bash
   cd workers/ts
   npm install
   cd ../..
   ```

---

## Usage

### Starting the Temporal Server

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

*(A screenshot of the Temporal UI in your browser is recommended.)*

---

### Running Workers

Each worker processes specific activities and listens on its designated task queue.

#### Go Worker

- **Location:** `workers/main.go`
- **Command:**

  ```bash
  go run workers/main.go
  ```

- **Notes:**  
  This worker registers the workflow (from the `workflow` folder) and the Go activity (`AddSuffixActivity` from the `activities` folder). It listens on the task queue `data-processing-task-queue`.

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
   The workflow (defined in `workflow/workflow.go`) sequentially calls activities across different languages:
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

## Contributing

Contributions are welcome! If you have suggestions or improvements, please open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push your branch.
4. Open a pull request explaining your changes.

---

Happy coding, and enjoy orchestrating your cross-language workflows with Temporal!

