import { ColorModeScript, extendTheme, ThemeConfig } from '@chakra-ui/react';
import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import reportWebVitals from './reportWebVitals';

const config: ThemeConfig = {
  useSystemColorMode: false,
  initialColorMode: 'dark',
};

const theme = extendTheme({
  sizes: {
    screenH: '100vh',
    screenW: '100vw',
  },
  config,
});

ReactDOM.render(
  <React.StrictMode>
    <ColorModeScript initialColorMode={theme.config.initialColorMode} />
    <App />
  </React.StrictMode>,
  document.getElementById('root'),
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals(console.log);
