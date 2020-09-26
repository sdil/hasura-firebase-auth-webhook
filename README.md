# Hasura Firebase Auth Webhook

![Go](https://github.com/sdil/hasura-firebase-auth-webhook/workflows/Go/badge.svg) ![Docker](https://github.com/sdil/hasura-firebase-auth-webhook/workflows/Docker/badge.svg?branch=master)

This project is highly inspired by [Hasura NodeJS auth webhook boilerplate](https://github.com/hasura/graphql-engine/blob/master/community/boilerplates/auth-webhooks/nodejs-firebase/firebase/firebaseHandler.js).

This webhook will verify the token.

## Table of Contents

- [How Hasura authentication works](#How-Hasura-authentication-works)
- [How Hasura Firebase Auth webhook helps](#How-Hasura-Firebase-Auth-webhook-helps)
- [How to use this](#How-to-use-this)

## How Hasura authentication works

Generally, Hasura has 3 modes to authenticate users:

- Webhook*
- JWT (JSON Web Token)
- Unauthenticated / Public Access

You can see more how it works [here](https://hasura.io/docs/1.0/graphql/core/auth/authentication/index.html).

This Firebase Auth is intended to work with Webhook

![Webhook](https://hasura.io/docs/1.0/_images/auth-webhook-overview1.png)

## How Hasura Firebase Auth webhook helps

If the user is authenticated, the webhook will return the following result to the Hasura server:
```
HTTP 200 OK
Cache-Control: 300
{
	"X-Hasura-User-Id": <Firebase UID>,
	"X-Hasura-Role":    "user",
}
```

## How to use this

```shell
$ docker run -it ghcr.io/sdil/hasura-firebase-auth-webhook:latest
```

