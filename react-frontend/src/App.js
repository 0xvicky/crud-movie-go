import React, {useEffect} from "react";
import "./App.css";

function App() {
  const getMovies = async () => {
    const res = await fetch("http://192.168.0.115:8000/movies/2");
    const data = await res.json();
    console.log(data);
  };

  useEffect(() => {
    getMovies();
  }, []);

  return <div className='App'>CRUD API UI</div>;
}

export default App;
