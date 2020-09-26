# Hasura Firebase Auth Webhook

![Go](https://github.com/sdil/hasura-firebase-auth-webhook/workflows/Go/badge.svg) ![Docker](https://github.com/sdil/hasura-firebase-auth-webhook/workflows/Docker/badge.svg?branch=master)

This project is highly inspired by [Hasura NodeJS auth webhook boilerplate](https://github.com/hasura/graphql-engine/blob/master/community/boilerplates/auth-webhooks/nodejs-firebase/firebase/firebaseHandler.js).

This webhook will verify the token.

## How Hasura authentication works

Generally, Hasura has 3 modes to authenticate users:

- Webhook*
- JWT (JSON Web Token)
- Unauthenticated / Public Access

You can see more how it works [here](https://hasura.io/docs/1.0/graphql/core/auth/authentication/index.html).

This Firebase Auth is intended to work with Webhook

![Webhook](https://hasura.io/docs/1.0/_images/auth-webhook-overview1.png)

## How Hasura Firebase Auth webhook helps

## How to use this

```shell
$ docker run -it ghcr.io/sdil/hasura-firebase-auth
```
