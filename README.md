# cnord-test

## run locally

1. Create PostgreSQL DB's "cnord" and "cnord_test"
2. Update POSTGRES_URL env variable at Makefile
3. Run DB migration 
   ```
   make migrationup
   ```
4. Generate .proto files with 
   ```
   make protos
   ```
6. Run service with
   ```
   make run
   ```

## testing
  Test API with 
  ```
  make testapi
  ```

## TODO
    - Swagger docs