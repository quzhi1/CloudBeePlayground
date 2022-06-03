# CloudBeePlayground
My personal playground for CloudBee

## Prerequisite
1. Put the environment key into `api_key` file.
2. Run:
```bash
pyenv virtualenv 3.8.5 cloud-bee-playground --force
pyenv activate cloud-bee-playground
pip install -r requirements.txt
```

## Run directly with CloudBee Go SDK
```bash
go run direct_use/main.go
```

## Run directly with CloudBee Python SDK
```bash
python direct_use/feature_flag.py
```

## Clean up
```bash
pyenv shell system
```