# unero-frontend

This is the frontend of a person budgeting web application.

Roadmap:

- [ ] Login/Logout/Authentication
- [ ] Able to import .csv files
- [ ] Create budgets and track expenses
- [ ] Report on historical trends

## What I've learned

### Flex

I've almost always used grid layouts when I can, and have mostly avoided flexbox. I've learned that flexbox is a lot
more powerful than I thought, and I'm excited to use it more in the future.

Some notes on flex:

- `justify-items` sets the default horizontal alignment of flex items within a flex container
- `align-items` sets the default vertical alignment of flex items within a flex container.
- `justify-content` controls the alignment and spacing of flex items along the main axis of the flex container
- `align-content` controls the alignment and spacing of flex items along the cross axis of the flex container when there
  is extra space.

### CSS Variables

I have always used preprocessor variables, but I've learned that CSS variables are quite capable and in some cases more
flexible than preprocessor variables

Example:

 ```css
/* setting */
:root {
    --primary-color: #007bff;
}

/* using */
button {
    background-color: var(--primary-color);
    color: white;
}
```

My OS doesn't play well with the browser and CSS feature `prefer-color-scheme`, so I found that setting document
variables and conditionally setting the CSS variables based on the document variables was a good solution.

```js
document.documentElement.setAttribute("theme", "dark");
```

```css
:root[theme="dark"] {
    --color-background: var(--vt-c-black);
    --color-background-soft: var(--vt-c-black-soft);
    --color-background-mute: var(--vt-c-black-mute);
}
```