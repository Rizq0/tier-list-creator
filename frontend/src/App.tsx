import React from 'react'
import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';

import './App.css'
import { theme } from './theme/Theme';
import ErrorBoundary from './components/utils/ErrorBoundary';
import Router from './router/Router';

function App() {

  return (
    <MantineProvider theme={theme}>
      <ErrorBoundary>
        <React.Suspense fallback="Loading...">
          <Router />
        </React.Suspense>
      </ErrorBoundary>
    </MantineProvider>
  )
}

export default App
