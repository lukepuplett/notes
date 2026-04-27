# React Notes - Traversy Media

## Core Concepts

A component is a function that returns JSX, takes in props, and encapsulates state internally.

- Use `useState` to define state for a component
- Use Context API, Redux, or Zustand for global state management

## Class vs. Functional Components

- Classes used to have lifecycle methods
- Hooks now allow you to hook into lifecycle functionality in a functional way
- Common hooks: `useState` and `useEffect`
- Other hooks exist but some are being retired

## Build & Development

React is typically built and hosted with **Vite**.

- Vite uses ES build
- Vite provides hot module replacement (HMR)

## Component Wrappers

- You can use wrapper components at the root level
- `React.StrictMode` is a common root wrapper component

## Component Declaration Syntax

Components are declared as functions. Multiple valid syntaxes:

```javascript
function MyComponent() { ... }
export const MyComponent = () => { ... }  // arrow function
export default function MyComponent() { ... }
```

## Styling

- CSS is imported at the top of the file and used like normal CSS imports
- Can use inline styles or import external CSS
- Inline styles use JavaScript objects (not CSS syntax)
- Object keys are camelCase versions of CSS property names (not snake-case)
  - Example: `backgroundColor` instead of `background-color`
- Define style objects as variables outside JSX: `<div style={myStyles}>`

## JSX Rules

### Single Element Rule
- Can only return a single element from a component
- Use a Fragment (`<>` and `</>`) to wrap multiple elements without adding an extra DOM node

### JSX Expressions
- Put expressions inside curly braces within JSX
- Use `array.map()` to render loops of HTML dynamically
  - Example: `{array.map((name) => <div>{name}</div>)}`
- Each HTML element in a list needs a `key` attribute or React will complain
  - Key is typically the array index: `<li key={index}>{item}</li>`

### Conditional Rendering
- Cannot use if statements inside JSX expressions (they must be one-liners)
- Use ternaries: `condition ? <div>True</div> : <div>False</div>`
- Use logical operators for simple conditionals: `condition && <div>Render this</div>`

## Component Architecture & Workflow

1. Design the full page in HTML first
2. Break it up into smaller, reusable components
3. Create a `components` folder to organize component files
4. Copy/paste HTML chunks into individual component files
5. Fix up imports as needed
6. Reference components from the parent/larger JSX file

This approach keeps non-dynamic parts simple and component-based.

**Note:** You can likely use a language model to break down a cohesive page into components quite easily.

## Asset Imports

You can import assets (like PNG images) using JavaScript import statements:

```javascript
import logo from './assets/logo.png'
```

## Props (Properties)

Props are component properties used to pass data into components.

```javascript
const MyComponent = (obj) => <h1>{obj.title}</h1>
```

### Destructuring Props

You can destructure props on the way in to make code cleaner:

```javascript
const MyComponent = ({ title, id }) => <h1>{title}</h1>
```

### Default Props

You can set default values during destructuring:

```javascript
const MyComponent = ({ title = "Default Title", id }) => ...
```

### Wrapper Components (Children)

Create wrapper components using a special `children` variable:

```javascript
const MyCard = ({ children }) => (
  <div className="card">
    {children}
  </div>
)
```

Use it like a standard HTML element:

```javascript
<MyCard>
  <p>Content goes here</p>
</MyCard>
```

### Props with Wrapper Components

You can pass additional props alongside children:

```javascript
const MyCard = ({ children, title, color }) => ...
```

Usage: `<MyCard title="My Title" color="blue">Content</MyCard>`

## State Management

### Types of State

- **Component State** - state managed within a single component
- **App State (Global State)** - state shared across multiple components

### Component State Use Cases

Toggle UI is a canonical example of good component state usage. A toggle component is the typical use case for component state.

### Using useState Hook

Import `useState` from React:

```javascript
import { useState } from 'react'
```

Declare state with destructuring:

```javascript
const [showThing, setShowThing] = useState(false)
```

- `showThing` - the current state value
- `setShowThing` - the function to update the state
- `false` - the initial/default value of the state

### Using useState with Events

Attach state updates to events like `onClick` on buttons or elements:

```javascript
onClick={() => setShowThing(!showThing)}
```

**Lambda Functions Required:** Must use a lambda (arrow) function to wrap the state update. Cannot call the function directly, or it will execute immediately on render.

- ❌ Wrong: `onClick={setShowThing()}`
- ✓ Correct: `onClick={() => setShowThing(!showThing)}`

### Updating State Based on Previous Value

You can pass a function to the setter that receives the current value:

```javascript
onClick={() => setShowThing(current => !current)}
```

- `current` represents the current/previous state value
- `!current` negates it (flips the boolean)

## React Router

### Imports

Import the following from `react-router-dom`:

- `createBrowserRouter`
- `createRoutesFromElements`
- `Route`
- `RouterProvider`

### Defining Routes

Create a router constant in your `App.jsx` file:

```javascript
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/about" element={<h1>Hello</h1>} />
  )
)
```

`Route` component takes two main attributes:

- `path` - the URL path (e.g., `/about`)
- `element` - the JSX to render when that path is visited
  - Can be pure JSX (e.g., `<h1>Hello</h1>`)
  - Can be a page component (e.g., `<AboutPage />`)

### Setting Up the App Component

```javascript
const App = () => {
  return <RouterProvider router={router} />
}
```

`RouterProvider` wraps your entire app and takes a `router` attribute set to the router you created.

### Nested Routes (Layouts)

`Route` components can have children to create layout structures:

```javascript
<Route path="/" element={<MainLayout />}>
  <Route index element={<HomePage />} />
</Route>
```

**Parent Route:**
- Has a `path` attribute (e.g., `/`)
- Has an `element` attribute containing the layout component (e.g., `<MainLayout />`)
- This layout component renders for all nested routes

**Child Routes:**
- Nested inside the parent `<Route>` tags
- Use the `index` attribute to define the default/home route
- `index` means this route renders when the path matches the parent path exactly
- Has an `element` attribute with the page component to render

**How It Works:**
When a user navigates to `/`, the `MainLayout` component renders. The `HomePage` component is rendered inside the `MainLayout` (via the `children` prop or `<Outlet />`). Child routes inherit the parent's layout.

### MainLayout Component

Import the `Outlet` component from `react-router-dom`:

```javascript
import { Outlet } from 'react-router-dom'
```

The `Outlet` component is a placeholder where child route content renders:

```javascript
const MainLayout = () => {
  return (
    <div>
      <header>
        <nav>Navigation here</nav>
      </header>
      <main>
        <Outlet />
      </main>
      <footer>
        <p>Footer here</p>
      </footer>
    </div>
  )
}
```

The `<Outlet />` is replaced by whatever child route's element is active (e.g., `<HomePage />`). This allows the layout structure (header, nav, footer) to persist while the main content changes based on the route.

### Link vs NavLink

Both are imported from `react-router-dom`.

**Link** - basic navigation component, renders as an `<a>` tag:

```javascript
<Link to="/about">About</Link>
```

**NavLink** - enhanced version of Link with active state styling. Automatically detects if the current route matches the link's path. Better for navigation menus where you want to highlight the active page.

#### NavLink with Dynamic className

`NavLink` can take a `className` that accepts a function. The function receives an object with `isActive` property:

```javascript
<NavLink 
  to="/about" 
  className={({ isActive }) => isActive ? 'nav-active' : ''}
>
  About
</NavLink>
```

Alternative approach (define separately for reusability):

```javascript
const navClassName = ({ isActive }) => isActive ? 'nav-active' : ''

<NavLink to="/about" className={navClassName}>
  About
</NavLink>
```

- `isActive` is `true` when the current route matches the link's `to` path
- `isActive` is `false` otherwise

## useEffect Hook

### Overview

Allows you to run side effects in functional components. "Side effects" refers to lifecycle events (component mount, update, cleanup).

### Syntax

```javascript
useEffect(() => {
  // work to do here
}, [dependencies])
```

### Dependencies Array

The second argument is a dependencies array in square brackets:

- Effect runs when any dependency in the array changes
- Empty array `[]` means the effect runs only once (on component mount)
- **Must always include a dependencies array**, or the effect will run on every render (causes problems)

### Async Functions in useEffect

Cannot pass an async function directly to `useEffect`. Must define an async function inside the effect and then call it:

```javascript
useEffect(() => {
  const fetchJobs = async () => {
    const response = await fetch('https://api.example.com/jobs')
    const data = await response.json()
    setJobs(data)
  }
  
  fetchJobs() // Call the async function
}, [])
```

**Why This Pattern:**
- `useEffect` itself cannot be async
- Define an async function inside the effect, then invoke it immediately
- This is the only way to use async/await within `useEffect`
