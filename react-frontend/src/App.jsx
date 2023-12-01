import React, {useEffect, useState, useReducer} from "react";
import "./App.css";
import MovieById from "./components/MovieById";
import AddMovie from "./components/AddMovie";
import MovieTable from "./components/MovieTable";

function App() {
  const [movies, setMovies] = useState([]);
  const [resData, setResData] = useState();

  const [reducerVal, forceUpdate] = useReducer(x => x + 1);
  const [redVal_2, forceUpdate_2] = useReducer(x => x + 1);
  const getMovies = async () => {
    const res = await fetch("http://localhost:8080/");
    const data = await res?.json();
    // console.log(data);
    setMovies(data);
  };

  useEffect(() => {
    getMovies();
  }, [reducerVal, redVal_2]);

  return (
    <div className='App'>
      <h2>CRUD API GO</h2>
      <div className='flex flex-col justify-center gap-y-10'>
        <div className='flex justify-between p-4 mx-3 mt-5 border border-black gap-x-3'>
          <div className='bg-purple-100 w-1/2'>
            {" "}
            <MovieById setData={setResData} />
            <AddMovie updateFunction={forceUpdate_2} />
          </div>

          <MovieTable
            movies={movies && movies}
            updateFunction={forceUpdate}
          />
        </div>
        <div className='mx-auto'>{resData && JSON.stringify(resData)}</div>
      </div>
    </div>
  );
}

export default App;
