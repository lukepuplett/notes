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
- If you're just assigning to a variable, you can skip declaring it `let x;` and Svelte will do it.
  
More info on JS label syntax: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/label

#### Prefix stores with $ to access their values

https://svelte.dev/docs#component-format-script-4-prefix-stores-with-$-to-access-their-values
  
- A store is an object that allows reactive access to a value via a simple _store contract_.
- The `svelte/store` module contains minimal store implementations which fulfil this contract.
- When you have a reference to a store, access its value inside a component by prefixing with `$`.
- Assignments to `$`-prefixed variables require it is `writable` and calls the store's `.set(...)` method.
- Stores must be declared at the top level.
  
```JavaScript
<script>
	import { writable } from 'svelte/store';

	const count = writable(0);
	console.log($count); // logs 0

	count.set(1);
	console.log($count); // logs 1

	$count = 2;
	console.log($count); // logs 2
</script>
```
  
#### Store contract

https://svelte.dev/docs#component-format-script-4-prefix-stores-with-$-to-access-their-values-store-contract
  
```JavaScript
  store = { subscribe: (subscription: (value: any) => void) => (() => void), set?: (value: any) => void }
```

- Create your own without `svelte/store` by implementing the contract.
- Must contain `.subscribe(subscriptionFunction)` method taking a function which must be immediately called with the store's current value.
- "All" of the store's subscription functions must be called whenever the value changes.
- The `.subscribe(...)` method must return an unsubscribe function, which must stop its subscription and never call its sub function again.
- A store can have a `.set(...)` method taking a new value and calls all the subscription functions synchronously.
- For RxJS compat, the `.subscribe(...)` method is also allowed to return an object with an `.unsubscribe(...)` method.
  
### <script context="module">
	
https://svelte.dev/docs#component-format-script-context-module
	
- A `<script>` with `context="module"` runs once when the module first evaluates, rather than for each component instance.
- Values declared in this block are accessible from a regular `<script>` (and the component markup) but not vice versa.
- Any `exports` in this block will become exports of the compiled module.
- No `export default`.
	
### <style>
	
https://svelte.dev/docs#component-format-style
	
- CSS is scoped to that component (Svelte prepends a class).
- Wrap like this `:global(body)` to apply globally.
- Note the `keyframes` syntax, too.
	
```CSS
<style>
:global(body) {
	/* this will apply to <body> */
	margin: 0;
}

div :global(strong) {
	/* this will apply to all <strong> elements, in any
		 component, that are inside <div> elements belonging
		 to this component */
	color: goldenrod;
}

p:global(.red) {
	/* this will apply to all <p> elements belonging to this 
		 component with a class of red, even if class="red" does
		 not initially appear in the markup, and is instead 
		 added at runtime. This is useful when the class 
		 of the element is dynamically applied, for instance 
		 when updating the element's classList property directly. */
}
	
@keyframes -global-my-animation-name {...}
</style>
```

- Only one `<style>` top level per component.
- Can nest like so:
	
```html
<div>
<style>
	/* this style tag will be inserted as-is */
	div {
		/* this will apply to all `<div>` elements in the DOM */
		color: red;
	}
</style>
</div>
```

**Note** - nested styles are inserted verbatim.

### Template Syntax

