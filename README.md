# Dink Docs

Dink (**D**on't Th**ink**) Docs helps you keep your Markdown docs in sync with your code, 
other docs, or basically any folder, file or line inside a file on your computer.

# Use Case

## tldr
Docs and code get out of sync all the time. Dink detects changes in code and prompts you 
to update docs that depend on it.


## Example

Let's say you have two services: a REST API that lives in the `api` directory, and a React 
app that lives in the `app` directory.

```
.
├── README.md
├── api
│   ├── package.json
│   └── server.ts
└── app
    ├── index.jsx
    └── package.json

2 directories, 5 files
```

You add a `README.md` file to describe how to set-up a local dev environment, add 
TypeScript support, etc.

```Markdown
# Toast Delivery App

## Setting up a local development environment

### Server
Change into the `api` directory and start the API server:

`yarn install && yarn dev`


### Frontend
In a separate terminal, change into the `app` directory and start the React app:

`yarn install && yarn dev`
```

Two months later, you decide to rename the `app` directory to `dashboard` _cuz that's kool now_. 
You find and replace all the usages of `app` with `dashboard` in your code. That's great but 
you forgot about the docs, which are now silently out of date.

Something like this happens surprisingly often. Dink helps solve this by letting you establish 
dependencies between docs and code with a one line Markdown comment.


# How it works
Dink lets you sync the contents of a Markdown file using a tag that ties it to the 
contents of another folder, file or line inside a file. It does this by defining a _Dink 
Block_ around the Markdown code you want to keep in sync. 

```markdown
[//]: # (!#dink:main.go)
### Frontend
In a separate terminal, change into the `dashboard` directory and start the React app:

`yarn install && yarn dev`
```

If you now make any changes to the `main.go` file, Dink will prompt about the potentially 
out of date block of Markdown above.
