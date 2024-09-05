import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import { provideCreateRoot } from "@gooddata/sdk-ui-ext";

// Include GoodData styles, needed for correct visualizations rendering
// You may exclude some file if you're not planning to use all visualizations
// For example, you can exclude sdk-ui-geo styles if you're not planning to use PushPin GeoChart
import "@gooddata/sdk-ui-kit/styles/css/main.css";
import "@gooddata/sdk-ui-filters/styles/css/main.css";
import "@gooddata/sdk-ui-charts/styles/css/main.css";
import "@gooddata/sdk-ui-pivot/styles/css/main.css";
import "@gooddata/sdk-ui-geo/styles/css/main.css";
import "@gooddata/sdk-ui-ext/styles/css/main.css";
import "@gooddata/sdk-ui-dashboard/styles/css/main.css";

import "./index.css";

// add provideCreateRoot to use React 18
provideCreateRoot(createRoot);

const container = document.getElementById('root')
const root = createRoot(container!)
root.render(<App />)
