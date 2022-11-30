# Eta

## Intro

```
import * as eta from "https://deno.land/x/eta@v1.6.0/mod.ts"
```

 - Replace with pinned version.
 - Move to deps.
 - Remember XSS: https://eta.js.org/docs/syntax/auto-escaping

```
Eta.render('The answer to everything is <%= it.answer %>', { answer: 42 })
```

## Partials

### Named partials

 - Defined as template functions.

```
Eta.templates.define("mypartial", Eta.compile("PARTIAL SPEAKING"))
Eta.render('This is a partial: <%~ include("mypartial") %>', { name: "Person" })
```

### File partials

 - Looks in config.views for the templates.
 - Actually more nuanced, see: https://eta.js.org/docs/learn/file-handling

```
Eta.configure({
  views: path.join(__dirname, views)
})

let template = "<%~ includeFile('./footer.eta', data) %>"

Eta.render(template, data)
```

## Layouts

 - Different type of "part".
 - Call this from inside your template:

```
<% layout(filepath, options) %>
```

 - Your current template will become the `it.body` for the layout.
 - Options is optional and overrides the inherited template body (model?) stored in `it.body`.
 - More advanced explainer I don't really understand: https://eta.js.org/docs/learn/layouts

### Examples

```
// layout.eta
//
<!DOCTYPE html>
<html lang="en">
  <head>
    <title><%= it.title %></title>
  </head>
  <body>
    <%~ it.body %>
    <footer>Copyright SomeCo, Inc.</footer>
  </body>
</html>
```

Which is called from the index template, presumably making the body in the above be the div and p below.

```
// index.eta
//
<% layout('./layout') %>

<div id="content">
  <p><%= it.message %></p>
</div>
```

#### Conditional layouts

```
<% if (user.type === 'admin') { %>
  <% layout('./admin') %>
<% } else { %>
  <% layout('./user') %>
<% } %>

This is the template content
```

Can also be laid out as JS within a single set of tags as:

```
<% if (user.type === "admin") {
  layout("./admin")
} else {
  layout("./user")
} %>

This is the template content
```

## Configuring Eta

 - Merges current configuration with that sent in:

```
Eta.configure({
  cache: true // Make Eta cache templates
})
```
 
 - Current configuration is in `config` variable:

```
Eta.config.tags // ["<%", "%>"]
```

### Big list of configuration options

 - Here's the TS interface for the config (taken from source code):

```
interface EtaConfig {
  /** Whether or not to automatically XML-escape interpolations. Default true */
  autoEscape: boolean
  /** Configure automatic whitespace trimming. Default `[false, 'nl']` */
  autoTrim: trimConfig | [trimConfig, trimConfig]
  /** Compile to async function */
  async: boolean
  /** Whether or not to cache templates if `name` or `filename` is passed */
  cache: boolean
  /** XML-escaping function */
  e: (str: string) => string
  /** Parsing options */
  parse: {
    /** Which prefix to use for evaluation. Default `""` */
    exec: string
    /** Which prefix to use for interpolation. Default `"="` */
    interpolate: string
    /** Which prefix to use for raw interpolation. Default `"~"` */
    raw: string
  }
  /** Array of plugins */
  plugins: Array<{
    processFnString?: Function
    processAST?: Function
    processTemplate?: Function
  }>
  /** Remove all safe-to-remove whitespace */
  rmWhitespace: boolean
  /** Delimiters: by default `['<%', '%>']` */
  tags: [string, string]
  /** Holds template cache */
  templates: Cacher<TemplateFunction>
  /** Name of the data object. Default `it` */
  varName: string
  /** Absolute path to template file */
  filename?: string
  /** Holds cache of resolved filepaths. Set to `false` to disable */
  filepathCache?: Record<string, string> | false
  /** A filter function applied to every interpolation or raw interpolation */
  filter?: Function
  /** Function to include templates by name */
  include?: Function
  /** Function to include templates by filepath */
  includeFile?: Function
  /** Name of template */
  name?: string
  /** Where should absolute paths begin? Default '/' */
  root?: string
  /** Make data available on the global object instead of varName */
  useWith?: boolean
  /** Whether or not to cache templates if `name` or `filename` is passed: duplicate of `cache` */
  "view cache"?: boolean
  /** Directory or directories that contain templates */
  views?: string | Array<string>
  /** The config object can also have other properties, potentially added by plugins */
  [index: string]: any
}
```

### Default configuration

```
var config: EtaConfig = {
  async: false,
  autoEscape: true,
  autoTrim: [false, "nl"],
  cache: false,
  e: XMLEscape, // function defined elsewhere
  include: includeHelper, // function defined elsewhere
  includeFile: includeFileHelper, // function defined elsewhere
  parse: {
    exec: "",
    interpolate: "=",
    raw: "~"
  },
  plugins: [],
  rmWhitespace: false,
  tags: ["<%", "%>"],
  templates: templates,
  useWith: false,
  varName: "it"
}
```

## Plugins

 - You can create plugins: https://eta.js.org/docs/learn/plugins

## How does Eta resolve template files?

What happens when you call `renderFile(path, ...)` or `<%~ includeFile(path, ...)%>`?

 - If it's absolute, first looks in `config.views` and if that's a path to a directory, looks in there, but if it's an array of directory paths then it looks in each. If not found, looks in `config.root` (by default `/`, the file system base).
 - If it's relative, then if `includeFile()` was called from another template file, tries to resolve the new template based on that template's path. If not found, falls back using `config.views`.

 - See: https://eta.js.org/docs/learn/file-handling

## Async Support

- **Must** configure Eta to run in async mode.
- You'll need `await include(...)` else you'll die lonely.

```
function asyncFunc() {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve("HI FROM ASYNC")
    }, 20)
  })
}

let result = await Eta.render(
  "<%= it.name %>: <%= await it.asyncFunc() %>",
  { name: "Ada Lovelace", asyncFunc: asyncFunc },
  { async: true }
)
// 'Ada Lovelace: HI FROM ASYNC'
```

### Special functions

- Also `renderAsync` and `renderFileAsync` which are asynchronous versions of originals.

```
function asyncFunc() {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve("HI FROM ASYNC")
    }, 20)
  })
}

let result = await Eta.renderAsync(
  "<%= it.name %>: <%= await it.asyncFunc() %>",
  { name: "Ada Lovelace", asyncFunc }
)
```

## Examples

```
<ul>
<% it.users.forEach(function(user){ %>
  <li><%= user %></li>
<% }) %>
</ul>
```

```
<%= await getSomeValue() %>
    
<% /* Eta supports multiline comments,
which is really useful */ %>
<%= "<%" %>
```

```
<% /* Embedded JS templates mean that you can
write any valid JS expression inside interpolate tags: */ %>
<%= 2 + 4 %>
```

```
<%~ include("mypartial") %>
<%~ includeFile('./navbar', { pages: [
  'home',
  'about',
  'users'
] }) %>
```

```
<% layout("layouts/basic") %>
<p>This will be rendered into a layout</p>
```
