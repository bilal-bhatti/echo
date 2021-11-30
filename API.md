# API Summary

```
Version:     1.0.0
Title:       OpenAPI Version 2 Specification
Description: OpenAPI Version 2 Specification
Host:        api.host.com
BasePath:    /api
Consumes:    [application/json]
Produces:    [application/json]
```

<details>
<summary>/contacts: post</summary>


```

```

`body parameter`
- body: `models.ContactRequest`
	- input: `string`

`responses`
- code: `200`, type: `models.ContactResponse`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/contacts/{id}: get</summary>


```

```

`query parameters`
- id: `integer`


`responses`
- code: `200`, type: `models.ContactResponse`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/echo/{str}: get</summary>


```

```

`path parameters`
- str: `string`


`responses`
- code: `200`, type: `services.EchoResponse`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/things: post</summary>


```

```

`body parameter`
- body: `models.ThingRequest`
	- input: `string`

`responses`
- code: `200`, type: `models.ThingResponse`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/things/{id}: get</summary>


```

```

`query parameters`
- id: `integer`


`responses`
- code: `200`, type: `models.ThingResponse`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

