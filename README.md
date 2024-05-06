<div align="center">
	<img src="https://github.com/dimaguy/cloudflare-email/assets/7207103/509b1b10-372a-4e54-9aff-75e05cbc4661" alt="Cloudflare Worker Email Server"/>
	<br>
  <h1>Cloudflare Worker Email Server</h1>
	<p>Send free transactional emails from your Cloudflare Workers using MailChannels.</p>
</div>


## Getting Started with local wrangler!

1. Clone this repository
2. Install the dependencies with `npm install`
3. Use the command `npx wrangler secret put --env production TOKEN` to deploy a securely stored token to Cloudflare. With this command, you will be prompted to enter a random secret value, which will be used to authenticate your requests with the HTTP `Authorization` header as described below. You can also set this encrypted value directly in your Cloudflare dashboard.
4. Repeat step 3 with the DKIM Secrets if need be
5. Deploy the worker with `npm run deploy`

Or deploy directly to Cloudflare and enable Github Actions

[![Deploy to Cloudflare Workers](https://deploy.workers.cloudflare.com/button)](https://deploy.workers.cloudflare.com/?url=https://github.com/robinbraemer/cloudflare-email)

## Setup SPF

SPF is a DNS record that helps prevent email spoofing. You will need to add an SPF record to your domain to allow MailChannels to send emails on your behalf.

1. Add a `TXT` record to your domain with the following values:

		- Name: `@`
		- Value: `v=spf1 a mx include:relay.mailchannels.net ~all`

Note: If you're facing [Domain Lockdown error](https://support.mailchannels.com/hc/en-us/articles/16918954360845-Secure-your-domain-name-against-spoofing-with-Domain-Lockdown), follow the below steps:

1. Create a `TXT` record with the following values:

		- Name: `_mailchannels`
		- Value: `v=mc1 cfid=yourdomain.workers.dev` (the value of `cfid` will also be present in the error response if you do it wrong)

## Setup DKIM

This step is optional, but highly recommended. DKIM is a DNS record that helps prevent email spoofing. You may follow the steps listed in the [MailChannels documentation](https://support.mailchannels.com/hc/en-us/articles/7122849237389-Adding-a-DKIM-Signature) to set up DKIM for your domain.

## Used Github/Cloudflare Secrets

- TOKEN - Client authentication
- DKIM_DOMAIN - DKIM Domain
- DKIM_PRIVATE_KEY - DKIM Private Key (Bare base64 encoded key without BEGIN and END statements and no newlines)
- DKIM_SELECTOR - DKIM Selector
- CF_ACCOUNT_ID - Cloudflare Account ID
- CF_API_TOKEN - Cloudflare API Token with enough capability to deploy workers

## Usage

Once you have deployed this worker function to Cloudflare Workers, you can send emails by making a `POST` request to the worker on the `/api/email` endpoint with the following parameters:

- Note you need to pass an `Authorization` header with the secure token you deployed. Like the following: `Authorization: TOKEN`

### Basic Email

The Most basic request would look like this:

```json
{
	"to": "john@example.com",
	"from": "me@example.com",
	"subject": "Hello World",
	"text": "Hello World"
}
```

### HTML Emails

You can also send HTML emails by adding an `html` parameter to the request. This can be used in conjunction with the `text` parameter to send a multi-part email.

```json
{
	"to": "john@example.com",
	"from": "me@example.com",
	"subject": "Hello World",
	"html": "<h1>Hello World</h1>"
}
```

### Sender and Recipient Name

You can also specify a sender and recipient name by adding a `name` parameter to the request. This can be used for both the `to` and `from` parameters.

```json
{
	"to": { "address": "john@example.com",  "name": "John Doe" },
	"from": { "address": "me@example.com", "name": "Jane Doe" },
	"subject": "Hello World",
	"text": "Hello World"
}
```

### Sending to Multiple Recipients

You may also send to multiple recipients by passing an array of emails, or an array of objects with `address` and `name` properties.

```json
{
	"to": [
		"john@example.com",
		"rose@example.com"
 	],
	"from": "me@example.com",
	"subject": "Hello World",
	"text": "Hello World"
}
```

or

```json
{
	"to": [
		{ "address": "john@example.com", "name": "John Doe" },
		{ "address": "rose@example.com", "name": "Rose Doe" }
 	],
	"from": "me@example.com",
	"subject": "Hello World",
	"text": "Hello World"
}
```

### Sending BCC and CC

You can also send BCC and CC emails by passing an array of emails, an object with `address` and `name` properties, or an array of either, similar to the `to` parameter.

```json
{
	"to": "john@example.com",
	"from": "me@example.com",
	"subject": "Hello World",
	"text": "Hello World",
	"cc": [
		"jim@example.com",
		"rose@example.com"
	],
	"bcc": [
		"gil@example.com"
	]
}
```

### Reply To

You can also specify a reply to email address by adding a `replyTo` parameter to the request. Again, you can use an email string, an object with `address` and `name` properties, or an array of either.

```json
{
	"to": "john@example.com",
	"from": "me@example.com",
	"replyTo": "support@example.com",
	"subject": "Hello World",
	"text": "Hello World"
}
```
