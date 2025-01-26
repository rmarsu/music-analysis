import React, { useState } from 'react';
import './App.css';

function App() {
  const [file, setFile] = useState(null);
  const [recommendations, setRecommendations] = useState([]);

  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData();
    formData.append('song', file);

    const fetchWithTimeout = (url, options = {}, timeout = 5000) => {
      return Promise.race([
        fetch(url, options),
        new Promise((_, reject) =>
          setTimeout(() => reject(new Error('Timeout')), timeout)
        )
      ]);
    };

    fetchWithTimeout('http://localhost:4001/api/get-recommend')
      .then(response => {
        const data = response.json();
        console.log(data);
        setRecommendations(data.recommendations);
      })
      .catch(error => {
        console.error('Fetch error:', error);
      });

  };

  return (
    <div className="app">
      <h1>🎀</h1>
      <h1>загрузите песню :3</h1>
      <form onSubmit={handleSubmit}>
        <input type="file" accept="audio/*" onChange={handleFileChange} />
        <button type="submit">получить рекомендации</button>
      </form>

      <h2>рекомендации:</h2>
      {recommendations.map((song) => (
        <div className="song-card" key={song.id}>
          <img src={song.cover} alt={`${song.title} обложка`} className="cover" />
          <div className="song-info">
            <h3>{song.title}</h3>
            <p>{song.description} <br /> {song.duration}</p>
          </div>
        </div>
      ))}
    </div>
  );

}

export default App;
