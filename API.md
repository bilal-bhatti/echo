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
<summary>/api/contacts: post</summary>


```

```

`body parameter`
- body: `ContactRequest`
	- input: `string`

`responses`
- code: `200`, type: `ContactResponse`
	- input: `object`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/api/contacts/{id}: get</summary>


```

```

`path parameters`
- id: `integer`


`responses`
- code: `200`, type: `ContactResponse`
	- input: `object`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/api/things: post</summary>


```

```

`body parameter`
- body: `ThingRequest`
	- input: `string`

`responses`
- code: `200`, type: `ThingResponse`
	- input: `object`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

<details>
<summary>/api/things/{id}: get</summary>


```

```

`path parameters`
- id: `integer`


`responses`
- code: `200`, type: `ThingResponse`
	- input: `object`
	- output: `string`
- `default`, type: `Error`
	- code: `integer`
	- status: `string`
</details>

