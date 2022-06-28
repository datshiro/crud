# Sample CRUD Application

This is a sample go application serving API Server for Create/Update/Delete/Read resource (E.g: User)

## Requirements

- Docker
- Go v1.12+

## How to run

- Init Docker, run:

  ```
  make docker/up
  ```

- Migrate DB, run:

  ```
  make db/up
  ```

- Run app:
  ```
  make app/run
  ```

