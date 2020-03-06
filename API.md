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
- name: `body`, type: `models.ContactRequest`
	- name: `input`, type: `string`

`responses`
- code: `200`, type: `models.ContactResponse`
	- name: `output`, type: `string`
- `default`, type: `Error`
	- name: `code`, type: `integer`
	- name: `status`, type: `string`
</details>

<details>
<summary>/contacts/{id}: get</summary>


```

```

`query parameters`
- name: `id`, type: `integer`


`responses`
- code: `200`, type: `models.ContactResponse`
	- name: `output`, type: `string`
- `default`, type: `Error`
	- name: `code`, type: `integer`
	- name: `status`, type: `string`
</details>

<details>
<summary>/echo/{str}: get</summary>


```

```

`path parameters`
- name: `str`, type: `string`


`responses`
- code: `200`, type: `services.EchoResponse`
	- name: `output`, type: `string`
- `default`, type: `Error`
	- name: `code`, type: `integer`
	- name: `status`, type: `string`
</details>

<details>
<summary>/things: post</summary>


```

```

`body parameter`
- name: `body`, type: `models.ThingRequest`
	- name: `input`, type: `string`

`responses`
- code: `200`, type: `models.ThingResponse`
	- name: `output`, type: `string`
- `default`, type: `Error`
	- name: `code`, type: `integer`
	- name: `status`, type: `string`
</details>

<details>
<summary>/things/{id}: get</summary>


```

```

`query parameters`
- name: `id`, type: `integer`


`responses`
- code: `200`, type: `models.ThingResponse`
	- name: `output`, type: `string`
- `default`, type: `Error`
	- name: `code`, type: `integer`
	- name: `status`, type: `string`
</details>

