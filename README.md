# blobman

## Installation

```sh
git clone https://github.com/kenjitheman/blobman
```

## Usage

### Get the black hole to make it work

```sh
cd docs
```

- yarn

```sh
yarn install
```

```sh
yarn run build
```

- npm

```sh
npm install
```

```sh
npm run build
```

Now you can see API docs on [localhost:8080](http://localhost:8080)

### Install golang dependencies

- **in project root**

```sh
go mod tidy
```

### CLI usage

#### Build it

```sh
go build -o blobman .
```

### Or

```sh
go run main.go
```

```rust
usage: ./blobman [<flags>] <command> [<args> ...]

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  run service
    run service

  migrate up
    migrate db up

  migrate down
    migrate db down
```

#### Examples

- migrate up

```sh
./blobman migrate up
```

- or

```sh
go run main.go migrate up
```

- migrate down

```sh
./blobman migrate down
```

- or

```sh
go run main.go migrate down
```

- run blobman service

```sh
./blobman run service
```

- or

```sh
go run main.go run service
```

Now you succesfully launched blobman

### Blobman API Usage

#### Get List of Blobs

**Request:**

```http
GET http://localhost:18080/integrations/blobman/blobs
```

##### Query Parameters:

    - pageNumber: Page number for pagination
    - pageLimit: Number of items per page
    - sorting: Sorting options

```json
{
  "filter": {
    "value": "arbitrary text"
  }
}
```

**Response:**

```json
[
  {
    "id": "12345678",
    "attributes": {
      "value": "random_value_1"
    },
    "relationships": {
      "owner": "ce7cfb01-6fdf-47ce-82a8-26122eb0ab01"
    }
  },
  {
    "id": "12345679",
    "attributes": {
      "value": "random_value_2"
    },
    "relationships": {
      "owner": "ce7cfb01-6fdf-47ce-82a8-26122eb0ab02"
    }
  },
  // ... other blobs
]
```

#### Get a Specific Blob

**Request:**

```http
GET http://localhost:18080/integrations/blobman/blobs/{id}
```

**Response:**

```json
{
  "id": "12345678",
  "attributes": {
    "value": "random_value_1"
  },
  "relationships": {
    "owner": "ce7cfb01-6fdf-47ce-82a8-26122eb0ab01"
  }
}
```

#### Create a Blob

**Request:**

```http
POST http://localhost:18080/integrations/blobman/blobs
```

**Request Body:**

```json
{
  "data": {
    "attributes": {
      "value": "new_blob_value"
    },
    "relationships": {
      "owner": "ce7cfb01-6fdf-47ce-82a8-26122eb0ab03"
    }
  }
}
```

**Response:**

```json
{
  "id": "newly_generated_blob_id",
  "attributes": {
    "value": "new_blob_value"
  },
  "relationships": {
    "owner": "ce7cfb01-6fdf-47ce-82a8-26122eb0ab03"
  }
}
```

#### Delete a Blob

**Request:**

```http
DELETE http://localhost:18080/integrations/blobman/blobs/{id}
```

**Response:**

```http
204 No Content
```

## License

- [MIT](https://choosealicense.com/licenses/mit/)
