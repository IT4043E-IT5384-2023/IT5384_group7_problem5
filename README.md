# Project Setup and Execution

## Step 1: Start Docker Compose

Navigate to the root directory of the project and run the following command in your terminal:

```bash
docker-compose up
```

## Step 2: Run API Service
Open a new terminal window, change the directory to the "pipeline" folder, and execute the following commands:
```bash
cd pipeline
pip install -r requirements.txt
python3 api.py
```

## Step 3: Run Kafka Service
Open another terminal window, change the directory to the "main" folder, and run the following command:
```bash
cd main
pip install -r requirements.txt
python3 kafka.py
```

## Step 4: Send Request using Postman
Open Postman and create a new request.

#### POST new transactions

```http
  POST http://localhost:3692/
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/` | `JSON`   | |

body=[{
    "id": "transaction_0x69769875b550f447c1008e4443960022dd53e69709a47eba5b0817fcd6506b98",
    "from_address": "0x67614Aaf3Fb2eAE60E009894d5d268a9bB70064b",
    "to_address": "0x632eC42e53736A382cA118F3f4Ae16525D0777AA",
    "value": 1683647532000000000,
    "block_timestamp": 1453409414
}]
