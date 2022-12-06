# Svelte

## Components

- Three sections, script, style and markup, all optional.

### <script>
  
- JavaScript that runs when component instance is created.
- Top-level variables declared and imports are accessible from component markup.
  
#### export creates a component prop
  
https://svelte.dev/docs#component-format-script-2-assignments-are-reactive
  
- The `export` keyword makes a variable a **prop** (property) which is accessible to component _consumers_.
- In Svelte _development mode_, a warning shows if no intial value is set and the consumer doesn't set one either.
- An exported `const`, `class` or `function` is readonly from outside the component.
- Functions are valid as prop values.
- Readonly props can be accessed as properties on the element, tied to the component using `bind:this` syntax.
- You can use reserved words as prop names, like `class`.

More info on that bind thing: https://svelte.dev/docs#template-syntax-component-directives-bind-this
  
#### Assignments are 'reactive'

https://svelte.dev/docs#component-format-script-2-assignments-are-reactive
  
- When you assign to a locally declare variable triggers a re-render.
- With arrays, using `.push()` and `.splice()` won't trigger, an assignment is needed.
- Svelte's `<script>` blocks only run on instantiation but function bodies obviously run whenever called.
  
More info on that array thing: https://svelte.dev/tutorial/updating-arrays-and-objects
  
#### $: marked a statement as reactive

https://svelte.dev/docs#component-format-script-3-$-marks-a-statement-as-reactive
  
- Any top-level statement can be made reactive with `$:` (JS label syntax) prefix.
- Reactive statements run after other script code and before the markup is rendered, whenever the values they depend on change.
- `$: document.title = title;` this will update whenever `title` prop changes.
- Only values which _directly_ appear in the `$:` statement will become reactive.
- Reactive blocks are ordered via static analysis at compile time: it's dumb and doesn't look inside/through functions.
- Order does matter and manual ordering is needed if functions obscure dependencies.
  
More info on JS label syntax: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/label
