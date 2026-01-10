import './App.css';
import './style/animations.css';
import { useState } from 'react';
import Header from './_components/Header.jsx';
import MessageList from './_components/MessageList.jsx';
import MessageForm from './_components/MessageForm.jsx';

function App() {
  const [refreshTrigger, setRefreshTrigger] = useState(0);

  const handleMessagePosted = () => {
    setRefreshTrigger((prev) => prev + 1);
  };

  return (
    <div className="app">
      <Header />
      <main className="main-content layout">
        <section className="layout-left">
          <MessageForm onMessagePosted={handleMessagePosted} />
        </section>
        <section className="layout-right">
          <MessageList key={refreshTrigger} />
        </section>
      </main>
    </div>
  );
}

export default App;
