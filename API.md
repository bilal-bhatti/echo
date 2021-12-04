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
- body: `services.ContactRequest`
	- input: `string`

`responses`
- code: `200`, type: `services.ContactResponse`
	- input: `object`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/contacts/{id}: get</summary>


```

```

`path parameters`
- id: `integer`


`responses`
- code: `200`, type: `services.ContactResponse`
	- input: `object`
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
- body: `services.ThingRequest`
	- input: `string`

`responses`
- code: `200`, type: `services.ThingResponse`
	- input: `object`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/things/{id}: get</summary>


```

```

`path parameters`
- id: `integer`


`responses`
- code: `200`, type: `services.ThingResponse`
	- input: `object`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

