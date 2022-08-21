# Dink Docs

Dink (**D**on't T**hink**) Docs helps you keep your Markdown docs in sync with your code, 
other docs, or basically any folder, file or line inside a file on your computer.

# Use Case
Let's say you have two services in your shiny new startup. A REST API that lives in the 
`api` directory and React app that lives in the app directory.

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

You add a `README.md` file to describe in painful detail how to set-up a local dev 
environment, add TypeScript support etc.

```Markdown
# Toast Delivery App

## Setting up a local development environment

### Server
Change into the api directory and start the API server:

`yarn install && yarn dev`


### Frontend
In a separate terminal, change into the `app` directory and start the React app:

`yarn install && yarn dev`
```

Two months later, you decide to rename the `app` directory to `dashboard` because that's 
what everyone else is doing. You find and replace all the usages of `app` with `dashboard`
in your code which is great but now, suddenly your docs are out of sync.

This is a contrived example but something like this happens surprisingly often. Dink helps
solve this by letting you link which docs link what code.


# How it works
Dink lets you sync the contents of a Markdown file using a tag that ties it to the 
contents of another folder, file or line inside a file. It does this by defining a _Dink 
Block_ around the Markdown code you want to keep in sync. 

```markdown
[//]: # (!#dink:main.go)
### Frontend
In a separate terminal, change into the `dashboard` directory and start the React app:

`yarn install && yarn dev`
[//]: # (!#donk)
```


