# blobman

## Project structure

```rust
blobman
│
├── config.yaml
├── Dockerfile
├── docs
│   ├── LICENSE
│   ├── package.json
│   ├── package-lock.json
│   ├── spec
│   │   ├── components
│   │   │   ├── parameters
│   │   │   │   ├── pageLimitParam.yaml
│   │   │   │   ├── pageNumberParam.yaml
│   │   │   │   └── sortingParam.yaml
│   │   │   ├── README.md
│   │   │   └── schemas
│   │   │       ├── BlobKey.yaml
│   │   │       ├── BlobRequest.yaml
│   │   │       ├── Blob.yaml
│   │   │       ├── Errors.yaml
│   │   │       └── OwnerKey.yaml
│   │   ├── openapi.yaml
│   │   ├── paths
│   │   │   ├── blobs@{id}.yaml
│   │   │   └── blobs.yaml
│   │   └── README.md
│   ├── web
│   │   ├── favicon.png
│   │   ├── index.hbs
│   │   ├── index.html
│   │   └── redoc-config.yaml
│   ├── web_deploy
│   │   ├── favicon.png
│   │   ├── index.hbs
│   │   ├── index.html
│   │   ├── openapi.json
│   │   └── openapi.yaml
│   └── yarn.lock
├── generate.sh
├── go.mod
├── go.sum
├── internal
│   ├── assets
│   │   ├── main.go
│   │   └── migrations
│   │       └── 001_initial.sql
│   ├── cli
│   │   ├── main.go
│   │   └── migrate.go
│   ├── config
│   │   └── main.go
│   ├── data
│   │   ├── blobs.go
│   │   ├── owners.go
│   │   └── postgres
│   │       ├── blobs.go
│   │       └── owners.go
│   └── service
│       ├── handlers
│       │   ├── create_blob.go
│       │   ├── ctx.go
│       │   ├── delete_blob.go
│       │   ├── get_blob.go
│       │   └── get_blobs.go
│       ├── main.go
│       ├── requests
│       │   ├── create_blob.go
│       │   ├── delete_blob.go
│       │   ├── get_blob.go
│       │   └── get_blobs.go
│       └── router.go
├── main.go
├── README.md
└── resources
    ├── db.go
    ├── flag.go
    ├── included.go
    ├── model_blob_attributes.go
    ├── model_blob.go
    ├── model_blob_relationships.go
    ├── model_blob_request_attributes.go
    ├── model_blob_request.go
    ├── model_blob_request_relationships.go
    ├── model_details.go
    ├── model_key.go
    ├── model_links.go
    ├── model_relation_collection.go
    ├── model_relation.go
    └── model_resource_type.go
```

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
export KV_VIPER_FILE="./config.yaml"
```

```sh
go build -o blobman .
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

- migrate down

```sh
./blobman migrate down
```

- run blobman service

```sh
./blobman run service
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
