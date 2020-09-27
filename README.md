# Hasura Firebase Auth Webhook

![Go](https://github.com/sdil/hasura-firebase-auth-webhook/workflows/Go/badge.svg) ![Docker](https://github.com/sdil/hasura-firebase-auth-webhook/workflows/Docker/badge.svg?branch=master)

This project is highly inspired by [Hasura NodeJS auth webhook boilerplate](https://github.com/hasura/graphql-engine/blob/master/community/boilerplates/auth-webhooks/nodejs-firebase/firebase/firebaseHandler.js).

This webhook will verify the token.

## Table of Contents

- [How Hasura authentication works](#How-Hasura-authentication-works)
- [How Hasura Firebase Auth webhook helps](#How-Hasura-Firebase-Auth-webhook-helps)
- [How to use this](#How-to-use-this)
    - [Docker-compose](#Docker-compose-environment)
    - [Host Installation](#Host-Installation)

## How Hasura authentication works

Generally, Hasura has 3 modes to authenticate users:

- Webhook*
- JWT (JSON Web Token)
- Unauthenticated / Public Access

You can see more how it works [here](https://hasura.io/docs/1.0/graphql/core/auth/authentication/index.html).

This Firebase Auth is intended to work with Webhook

![Webhook](https://hasura.io/docs/1.0/_images/auth-webhook-overview1.png)

## How Hasura Firebase Auth webhook helps

If the user is successfully authenticated, the webhook will return the following result to the Hasura server:
```
$ curl -i localhost:8081 -H "Authorization: Bearer <Firebase token>"

HTTP 200 OK
Content-Type: application/json
Cache-Control: 300
{
    "X-Hasura-User-Id": <Firebase UID>,
    "X-Hasura-Role":    "user",
}
```

## Installation

### Docker-compose environment

1. Launch the webhook container:

    ```shell
    $ curl -o <docker compose manifest> -o docker-compose.yaml
    $ docker-compose up -d
    ```

2. Test the webhook server

### Host Installation

1. Download and launch the webhook server

    ```shell
    $ export GOOGLE_APPLICATION_CREDENTIALS="<path to service-account.json file>"
    $ curl -o <file to download> | sh
    ```

2. Start the Hasura GraphQL engine and point the authentication webhook to the hasura-firebase-auth webhook server

    ```shell
    graphql-engine --database-url <DB URL> serve --admin-secret <ADMIN_SECRET_KEY> --auth-hook localhost:8081
    ```

3. Test the GraphQL operation

## Roadmap

- Write a comprehensive unit tests
- Perform a benchmark & performance test
- Handle exceptions
- Add Kubernetes manifest example
