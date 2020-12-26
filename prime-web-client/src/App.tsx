import React from 'react';
import logo from './logo.svg';
import './App.css';
import Prime from './components/prime/prime'

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Prime/>
      </header>
    </div>
  );
}

export default App;
