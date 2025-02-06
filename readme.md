# Data Processing Service

This Temporal project is organized as follows:

## Project Structure

- **client/**  
  Contains the main Go client that submits workflows.

- **processing/**  
  - `workflow.go` – Defines the main Temporal workflow.  
  - `activities.go` – Go activity implementations.  
  - `activities_python.py` – Python activity implementation.  
  - `activities_ts.ts` – TypeScript activity implementation.

- **workers/**  
  - **Go Worker:** `main.go` for Go activities.  
  - **python/** – Contains `python_worker.py` for Python activities.  
  - **ts/** – Contains `ts_worker.ts` for TypeScript activities.

- **your_temporal.db**  
  Local database (if applicable).

```

.
├── client
│   └── main.go
├── go.mod
├── go.sum
├── processing
│   ├── __pycache__
│   │   └── activities_python.cpython-313.pyc
│   ├── activities.go
│   ├── activities_python.py
│   ├── activities_ts.ts
│   └── workflow.go
├── workers
│   ├── main.go
│   ├── python
│   │   ├── __pycache__
│   │   │   └── python_worker.cpython-313.pyc
│   │   └── python_worker.py
│   └── ts
│       ├── package-lock.json
│       ├── package.json
│       └── ts_worker.ts
└── your_temporal.db

8 directories, 15 files
```
<br>  
## Order of Execution

1. **Temporal Server**  
   Ensure your Temporal server is running.

2. **Workers**  
   - Start the Go worker.  
   - Start the Python worker.  
   - Start the TypeScript worker.

3. **Client**  
   Run the client to submit workflows.

