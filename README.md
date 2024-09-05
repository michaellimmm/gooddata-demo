# GoodData Demo

## Reference

- [Frontend integration series react](https://medium.com/gooddata-developers/frontend-integration-series-react-89b943086326)
- [Build a tool for data analysis data visualizations](https://medium.com/gooddata-developers/build-a-tool-for-data-analysis-data-visualizations-2d0e5c08d308)

## Notes

- remove React.StrictMode in `main.tsx`

```tsx
ReactDOM.createRoot(document.getElementById("root")!).render(
  // <React.StrictMode>
  <App />
  // </React.StrictMode>,
);
```

- add resolve `@gooddata` on vite.config.ts

```js
{
  // ...
  resolve: {
    alias: {
        "~@gooddata": "/node_modules/@gooddata",
    },
  }
  // ...
}

```

- add `provideCreateRoot(createRoot);` on main.tsx to use react 18 otherwise it will use react 17 as default

- add `@gooddata/sdk-ui-~ css` on main.tsx, to make ui more beautiful

```js
import "@gooddata/sdk-ui-kit/styles/css/main.css";
import "@gooddata/sdk-ui-filters/styles/css/main.css";
import "@gooddata/sdk-ui-charts/styles/css/main.css";
import "@gooddata/sdk-ui-pivot/styles/css/main.css";
import "@gooddata/sdk-ui-geo/styles/css/main.css";
import "@gooddata/sdk-ui-ext/styles/css/main.css";
import "@gooddata/sdk-ui-dashboard/styles/css/main.css";
```
